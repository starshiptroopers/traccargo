// Copyright 2021 The Starship Troopers Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Traccargo is the Golang library for fetching devices information and device position updates from Traccar opensource GPS tracking system
//
// It supports only several requests to tracckar api: devices and device positions
// Also, traccargo support traccar websocket feed for continuous live device positions updates

package traccargo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/starshiptroopers/traccargo/models"
	"golang.org/x/net/websocket"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Traccargo struct {
	url                    *url.URL
	authToken              string
	user                   *models.User
	sessionCookies         []*http.Cookie
	httpClient             *http.Client
	ws                     *websocket.Conn
	wsMutex                sync.Mutex
	wsSubscription         func(message *WsMessage)
	LogWriter              io.Writer //for debug purpose. All significant operation and error events is written to this io.Writer if defined
	LogCommunicationWriter io.Writer //for debug purpose. All JSON responses will be written to this io.Writer if defined
}

var timeout = time.Second
var retries = 3
var logCommunicationWriter io.Writer = nil
var logWriter io.Writer = nil

var (
	TRACCAR_ERROR_UNREACHABLE = errors.New("server is unreachable")
	TRACCAR_ERROR_AUTH        = errors.New("server auth is failed")
	TRACCAR_ERROR_FAILED      = errors.New("operation can't be performed")
	TRACCAR_ERROR_UNKNOWN     = errors.New("unknown error")
	TRACCAR_ERROR_NOTFOUND    = errors.New("object not found")
)

//creates a traccar instance
func NewTraccar(apiUrl string, authToken string) (tr *Traccargo, err error) {
	u, err := url.Parse(apiUrl)
	if err != nil {
		return
	}
	tr = &Traccargo{
		url:                    u,
		authToken:              authToken,
		LogCommunicationWriter: logCommunicationWriter,
		LogWriter:              logWriter,
	}

	return
}

func (t *Traccargo) Close() {
	t.wsClose()
}

//subscribes to live traccar updates
//WebSocket connection to the traccar api endpoint will be established and keeps alive
//the callback function calls on every update
func (t *Traccargo) SubscribeUpdates(handler func(m *WsMessage)) (err error) {

	if t.wsSubscription != nil {
		return errors.New("already subscribed, only one subscriber is allowed")
	}
	t.wsSubscription = handler

	err = t.wsConnect()
	if err != nil {
		return
	}

	return
}

//finish websocket connection and live updates subscription
func (t *Traccargo) UnsubscribeUpdates() {

	t.wsSubscription = nil
	t.wsClose()

	return
}

//returns all devices positions from traccar server
func (t *Traccargo) Positions() (positions []*models.Position, err error) {

	if !t.isAuthorized() {
		if _, err = t.Session(); err != nil {
			return
		}
	}

	_, _, err = t.request(apiDescriptor{
		method:          "GET",
		path:            "/api/positions",
		request:         nil,
		requestEncoding: ENC_URI,
		response:        &positions,
	})
	if err != nil {
		return
	}

	return
}

//returns the device position from traccar server
func (t *Traccargo) Position(deviceId int64) (position models.Position, err error) {

	positions, err := t.Positions()
	if err != nil {
		return
	}
	err = TRACCAR_ERROR_NOTFOUND
	for _, devicePosition := range positions {
		if devicePosition.DeviceID == deviceId {
			return *devicePosition, nil
		}
	}
	return
}

//authorizes to the traccar server and returns User object
func (t *Traccargo) Session() (user models.User, err error) {

	_, res, err := t.request(apiDescriptor{
		method: "GET",
		path:   "/api/session",
		request: &struct {
			Token string `url:"token"`
		}{Token: t.authToken},
		requestEncoding: ENC_URI,
		response:        &user,
	})

	if err == TRACCAR_ERROR_UNKNOWN && res != nil && res.StatusCode == http.StatusNotFound {
		err = TRACCAR_ERROR_AUTH
		return
	}

	if err != nil {
		return
	}

	t.user = &user
	t.sessionCookies = res.Cookies()

	return
}

//do the request to the api endpoint
//errors TRACCAR_ERROR_UNREACHABLE | TRACCAR_ERROR_AUTH | url.Enc errors, http.request errors
func (t *Traccargo) request(dscr apiDescriptor) (responsePayload interface{}, httpResponse *http.Response, err error) {

	endpointURL := t.endpointURL(dscr.path)

	if dscr.requestEncoding == ENC_URI {
		values, encErr := query.Values(dscr.request)
		if encErr != nil {
			err = encErr
			return
		}
		if endpointURL, encErr = endpointURL.Parse("?" + values.Encode()); encErr != nil {
			err = encErr
			return
		}
	} else if dscr.requestEncoding == ENC_JSON {
		err = errors.New("JSON encoding isn't supported now")
	}

	req, err := http.NewRequest(dscr.method, endpointURL.String(), nil)
	if err != nil {
		return
	}

	for _, cookie := range t.sessionCookies {
		req.AddCookie(cookie)
	}

	var requestTimeout = timeout
	doRequest := func() (err error) {
		if t.httpClient == nil {
			t.httpClient = &http.Client{Timeout: requestTimeout}
		}
		httpResponse, err = t.httpClient.Do(req)
		t.debugPrint("We've sent a request to: " + dscr.path)
		if err == nil {
			return
		}
		if err, ok := err.(*url.Error); ok {
			if err.Timeout() {
				return err
			}
		}
		return TRACCAR_ERROR_UNREACHABLE
	}

	if err = t.doWithRetries(doRequest, retries); err != nil {
		return
	}

	buff, _ := ioutil.ReadAll(httpResponse.Body)
	_ = httpResponse.Body.Close()

	if httpResponse.StatusCode == http.StatusBadRequest {
		err = TRACCAR_ERROR_FAILED
		return
	}
	if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 300 {
		err = TRACCAR_ERROR_UNKNOWN
		return
	}

	//todo try to create a copy of dscr.response
	//response = reflect.New(reflect.TypeOf(dscr.response))
	r := dscr.response
	responsePayload = r
	err = json.Unmarshal(buff, responsePayload)

	if err == nil {
		t.debugPrintJSON("we'v got response: ", responsePayload)
	}
	return
}

//return an existent or create a new websocket connection
//err = TRACCAR_ERROR_UNREACHABLE on some kind of network connection error (timeout, wrong http response code and etc)
//      or another error on a pre-connection stage (for example wrong uri and etc)
func (t *Traccargo) wsConnect() (err error) {

	if t.ws != nil {
		return nil
	}

	config, err := t.wsPrepare()
	if err != nil {
		return
	}

	//starting the connection
	var readFailed = make(chan bool)
	var connectionError = make(chan error)

	//reconnection cycle
	go func() {
		defer func() {
			close(connectionError)
			close(readFailed)
		}()
		var attempt = 0

		for {
			attempt++
			t.wsMutex.Lock()
			err = t.doWithRetries(
				func() (err error) {
					t.debugPrint("websocket connection attempt: " + strconv.Itoa(attempt))
					t.ws, err = websocket.DialConfig(config)
					return
				},
				retries,
			)
			t.wsMutex.Unlock()
			//notify about a connection result
			if attempt == 1 {
				connectionError <- err
			}

			if err == nil {
				t.debugPrint("websocket connection established")
			} else {
				t.debugPrintError(fmt.Errorf("can't connect to the websocket: %v", err))
			}

			//the reading cycle
			go t.wsListen(readFailed)

			//waiting for reading stream is die
			<-readFailed

			if !t.wsClose() {
				//finish the loop when ws is nil
				t.debugPrint("websocket connection finished")
				break
			}

			time.Sleep(time.Second)
		}
	}()

	//wait for websocket connection is established
	return <-connectionError
}

//close websocket connection
func (t *Traccargo) wsClose() bool {
	t.wsMutex.Lock()
	defer t.wsMutex.Unlock()
	if t.ws == nil {
		return false
	}
	err := t.ws.Close()
	if err != nil {
		t.debugPrintError(fmt.Errorf("hm, error on websocket close: %v", err))
	} else {
		t.debugPrint("websocket connection closed")
	}
	t.ws = nil
	return true
}

//prepare websocket config struct
func (t *Traccargo) wsPrepare() (config *websocket.Config, err error) {

	//doing the authorization
	if !t.isAuthorized() {
		if _, err = t.Session(); err != nil {
			return
		}
	}

	//preparing the uri
	websocketUrl := "ws" + strings.TrimPrefix(t.endpointURL("/api/socket").String(), "http")
	config, err = websocket.NewConfig(websocketUrl, t.url.String())
	if err != nil {
		return
	}

	//adding auth cookies
	for _, cookie := range t.sessionCookies {
		encodedCookie := cookie.String()
		if c := config.Header.Get("Cookie"); c != "" {
			config.Header.Set("Cookie", c+"; "+encodedCookie)
		} else {
			config.Header.Set("Cookie", encodedCookie)
		}
	}
	return
}

//listen websocket for new messages, process message and calls subscriber callback if defined
func (t *Traccargo) wsListen(close chan bool) {
	ws := t.ws
	for {
		var m WsMessage
		err := websocket.JSON.Receive(ws, &m)

		if err == nil {
			t.debugPrintJSON("We'v got websocket data: ", &m)
			if t.wsSubscription != nil {
				go t.wsSubscription(&m)
			}
			continue
		}

		if strings.Contains(err.Error(), "json") {
			t.debugPrintError(fmt.Errorf("wrong json received"))
			//continue on json error
			continue
		}

		if strings.Contains(err.Error(), "EOF") {
			t.debugPrint("websocket connection closed by other side")
		} else if !strings.Contains(err.Error(), "closed") {
			t.debugPrintError(fmt.Errorf("unexpected websocket error: %v", err))
		}

		close <- true
		return
	}
}

//do the operation with N retries on error
//break on retries == 0 or err == nil or err == TRACCAR_ERROR_UNREACHABLE
func (t *Traccargo) doWithRetries(operation func() (err error), retries int) (err error) {
	err = operation()
	if err == nil || err == TRACCAR_ERROR_UNREACHABLE {
		return
	}

	if retries > 0 {
		return t.doWithRetries(operation, retries-1)
	} else {
		return TRACCAR_ERROR_UNREACHABLE
	}
}

//return the full endpoint URL by it's relative path
func (t *Traccargo) endpointURL(path string) (endpointURL *url.URL) {
	endpointURL, err := url.Parse(path)
	if err != nil {
		panic(err)
	}
	return t.url.ResolveReference(endpointURL)
}

//print debug message to the log writer
func (t *Traccargo) debugPrint(str string) {
	if t.LogWriter == nil {
		return
	}

	_, _ = fmt.Fprintln(t.LogWriter, str)
}

//print JSON struct to the log writer (communication)
func (t *Traccargo) debugPrintJSON(prefix string, data interface{}) {
	if t.LogCommunicationWriter == nil {
		return
	}
	t.debugPrint(prefix)
	b, _ := json.MarshalIndent(data, "", "	")
	t.debugPrint(string(b))
}

//print error to the debug log writer
func (t *Traccargo) debugPrintError(err error) {
	if t.LogWriter == nil {
		return
	}
	t.debugPrint(string(err.Error()))
}

//returns true when traccar session is defined
func (t *Traccargo) isAuthorized() bool {
	return t.user != nil
}

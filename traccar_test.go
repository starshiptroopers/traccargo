package traccargo

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testApiURL = ""
var testApiToken = ""
var testDeviceID int64 = 0

func TestTraccar(t *testing.T) {
	trc, err := NewTraccar(testApiURL, testApiToken)
	if err != nil {
		panic(err)
	}
	trc.LogCommunicationWriter = os.Stdout
	trc.LogWriter = os.Stdout

	position, err := trc.Position(testDeviceID)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Device %d: time %v, Position [%.6f %.6f]\n", position.DeviceID, position.FixTime, position.Latitude, position.Longitude)

	err = trc.SubscribeUpdates(func(m *WsMessage) {
		fmt.Printf("new message received\n")
	})

	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 10)

	trc.UnsubscribeUpdates()

	time.Sleep(time.Second * 5)
}

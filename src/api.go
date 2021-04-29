// Copyright 2021 The Starship Troopers Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package traccar

import "traccar/src/traccar/models"

type encoding uint8

const (
	ENC_URI 	encoding 	= 1
	ENC_JSON 	encoding 	= 2
)

type apiDescriptor struct {
	method 		string
	path		string
	request		interface{}
	requestEncoding	encoding
	response 	interface{}
}

type WsMessage struct {
	Devices		*[]models.Device
	Positions   *[]models.Position
	Event		*[]models.Event
}
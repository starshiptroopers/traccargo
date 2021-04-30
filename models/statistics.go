// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Statistics statistics
//
// swagger:model Statistics
type Statistics struct {

	// active devices
	ActiveDevices int64 `json:"activeDevices,omitempty"`

	// active users
	ActiveUsers int64 `json:"activeUsers,omitempty"`

	// in IS0 8601 format. eg. `1963-11-22T18:30:00Z`
	// Format: date-time
	CaptureTime strfmt.DateTime `json:"captureTime,omitempty"`

	// messages received
	MessagesReceived int64 `json:"messagesReceived,omitempty"`

	// messages stored
	MessagesStored int64 `json:"messagesStored,omitempty"`

	// requests
	Requests int64 `json:"requests,omitempty"`
}

// Validate validates this statistics
func (m *Statistics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCaptureTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Statistics) validateCaptureTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CaptureTime) { // not required
		return nil
	}

	if err := validate.FormatOf("captureTime", "body", "date-time", m.CaptureTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this statistics based on context it is used
func (m *Statistics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Statistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Statistics) UnmarshalBinary(b []byte) error {
	var res Statistics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
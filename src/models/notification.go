// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Notification notification
//
// swagger:model Notification
type Notification struct {

	// always
	Always bool `json:"always,omitempty"`

	// attributes
	Attributes interface{} `json:"attributes,omitempty"`

	// calendar Id
	CalendarID int64 `json:"calendarId,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// mail
	Mail bool `json:"mail,omitempty"`

	// sms
	Sms bool `json:"sms,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// web
	Web bool `json:"web,omitempty"`
}

// Validate validates this notification
func (m *Notification) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this notification based on context it is used
func (m *Notification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Notification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Notification) UnmarshalBinary(b []byte) error {
	var res Notification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
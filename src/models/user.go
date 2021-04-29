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

// User user
//
// swagger:model User
type User struct {

	// administrator
	Administrator bool `json:"administrator,omitempty"`

	// attributes
	Attributes interface{} `json:"attributes,omitempty"`

	// coordinate format
	CoordinateFormat string `json:"coordinateFormat,omitempty"`

	// device limit
	DeviceLimit int64 `json:"deviceLimit,omitempty"`

	// device readonly
	DeviceReadonly bool `json:"deviceReadonly,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// in IS0 8601 format. eg. `1963-11-22T18:30:00Z`
	// Format: date-time
	ExpirationTime strfmt.DateTime `json:"expirationTime,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// latitude
	Latitude float64 `json:"latitude,omitempty"`

	// limit commands
	LimitCommands bool `json:"limitCommands,omitempty"`

	// longitude
	Longitude float64 `json:"longitude,omitempty"`

	// map
	Map string `json:"map,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// poi layer
	PoiLayer string `json:"poiLayer,omitempty"`

	// readonly
	Readonly bool `json:"readonly,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// twelve hour format
	TwelveHourFormat bool `json:"twelveHourFormat,omitempty"`

	// user limit
	UserLimit int64 `json:"userLimit,omitempty"`

	// zoom
	Zoom int64 `json:"zoom,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpirationTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateExpirationTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpirationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("expirationTime", "body", "date-time", m.ExpirationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user based on context it is used
func (m *User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
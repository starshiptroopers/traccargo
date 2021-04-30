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

// Event event
//
// swagger:model Event
type Event struct {

	// attributes
	Attributes interface{} `json:"attributes,omitempty"`

	// device Id
	DeviceID int64 `json:"deviceId,omitempty"`

	// geofence Id
	GeofenceID int64 `json:"geofenceId,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// maintenance Id
	MaintenanceID int64 `json:"maintenanceId,omitempty"`

	// position Id
	PositionID int64 `json:"positionId,omitempty"`

	// in IS0 8601 format. eg. `1963-11-22T18:30:00Z`
	// Format: date-time
	ServerTime strfmt.DateTime `json:"serverTime,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServerTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateServerTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerTime) { // not required
		return nil
	}

	if err := validate.FormatOf("serverTime", "body", "date-time", m.ServerTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this event based on context it is used
func (m *Event) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
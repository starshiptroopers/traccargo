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

// Device device
//
// swagger:model Device
type Device struct {

	// attributes
	Attributes interface{} `json:"attributes,omitempty"`

	// category
	Category string `json:"category,omitempty"`

	// contact
	Contact string `json:"contact,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// geofence ids
	GeofenceIds []int64 `json:"geofenceIds"`

	// group Id
	GroupID int64 `json:"groupId,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// in IS0 8601 format. eg. `1963-11-22T18:30:00Z`
	// Format: date-time
	LastUpdate strfmt.DateTime `json:"lastUpdate,omitempty"`

	// model
	Model string `json:"model,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// position Id
	PositionID int64 `json:"positionId,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// unique Id
	UniqueID string `json:"uniqueId,omitempty"`
}

// Validate validates this device
func (m *Device) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastUpdate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Device) validateLastUpdate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdate", "body", "date-time", m.LastUpdate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this device based on context it is used
func (m *Device) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Device) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Device) UnmarshalBinary(b []byte) error {
	var res Device
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

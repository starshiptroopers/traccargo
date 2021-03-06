// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceAccumulators device accumulators
//
// swagger:model DeviceAccumulators
type DeviceAccumulators struct {

	// device Id
	DeviceID int64 `json:"deviceId,omitempty"`

	// hours
	Hours float64 `json:"hours,omitempty"`

	// in meters
	TotalDistance float64 `json:"totalDistance,omitempty"`
}

// Validate validates this device accumulators
func (m *DeviceAccumulators) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this device accumulators based on context it is used
func (m *DeviceAccumulators) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceAccumulators) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceAccumulators) UnmarshalBinary(b []byte) error {
	var res DeviceAccumulators
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReportSummary report summary
//
// swagger:model ReportSummary
type ReportSummary struct {

	// in knots
	AverageSpeed float64 `json:"averageSpeed,omitempty"`

	// device Id
	DeviceID int64 `json:"deviceId,omitempty"`

	// device name
	DeviceName string `json:"deviceName,omitempty"`

	// in meters
	Distance float64 `json:"distance,omitempty"`

	// engine hours
	EngineHours int64 `json:"engineHours,omitempty"`

	// in knots
	MaxSpeed float64 `json:"maxSpeed,omitempty"`

	// in liters
	SpentFuel float64 `json:"spentFuel,omitempty"`
}

// Validate validates this report summary
func (m *ReportSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this report summary based on context it is used
func (m *ReportSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportSummary) UnmarshalBinary(b []byte) error {
	var res ReportSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

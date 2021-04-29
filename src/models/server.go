// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Server server
//
// swagger:model Server
type Server struct {

	// attributes
	Attributes interface{} `json:"attributes,omitempty"`

	// bing key
	BingKey string `json:"bingKey,omitempty"`

	// coordinate format
	CoordinateFormat string `json:"coordinateFormat,omitempty"`

	// device readonly
	DeviceReadonly bool `json:"deviceReadonly,omitempty"`

	// force settings
	ForceSettings bool `json:"forceSettings,omitempty"`

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

	// map Url
	MapURL string `json:"mapUrl,omitempty"`

	// poi layer
	PoiLayer string `json:"poiLayer,omitempty"`

	// readonly
	Readonly bool `json:"readonly,omitempty"`

	// registration
	Registration bool `json:"registration,omitempty"`

	// twelve hour format
	TwelveHourFormat bool `json:"twelveHourFormat,omitempty"`

	// version
	Version string `json:"version,omitempty"`

	// zoom
	Zoom int64 `json:"zoom,omitempty"`
}

// Validate validates this server
func (m *Server) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this server based on context it is used
func (m *Server) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Server) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Server) UnmarshalBinary(b []byte) error {
	var res Server
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

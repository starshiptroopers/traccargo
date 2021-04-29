// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NotificationType notification type
//
// swagger:model NotificationType
type NotificationType struct {

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this notification type
func (m *NotificationType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this notification type based on context it is used
func (m *NotificationType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotificationType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationType) UnmarshalBinary(b []byte) error {
	var res NotificationType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
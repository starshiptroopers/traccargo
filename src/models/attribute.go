// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Attribute attribute
//
// swagger:model Attribute
type Attribute struct {

	// attribute
	Attribute string `json:"attribute,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// expression
	Expression string `json:"expression,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// String|Number|Boolean
	Type string `json:"type,omitempty"`
}

// Validate validates this attribute
func (m *Attribute) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this attribute based on context it is used
func (m *Attribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Attribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Attribute) UnmarshalBinary(b []byte) error {
	var res Attribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

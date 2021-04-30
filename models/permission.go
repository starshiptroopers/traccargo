// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Permission This is a permission map that contain two object indexes. It is used to link/unlink objects. Order is important. Example: { deviceId:8, geofenceId: 16 }
//
// swagger:model Permission
type Permission struct {

	// Computed Attribute Id, can be second parameter only
	AttributeID int64 `json:"attributeId,omitempty"`

	// Calendar Id, can be second parameter only and only in combination with userId
	CalendarID int64 `json:"calendarId,omitempty"`

	// Device Id, can be first parameter or second only in combination with userId
	DeviceID int64 `json:"deviceId,omitempty"`

	// Driver Id, can be second parameter only
	DriverID int64 `json:"driverId,omitempty"`

	// Geofence Id, can be second parameter only
	GeofenceID int64 `json:"geofenceId,omitempty"`

	// Group Id, can be first parameter or second only in combination with userId
	GroupID int64 `json:"groupId,omitempty"`

	// User Id, can be second parameter only and only in combination with userId
	ManagedUserID int64 `json:"managedUserId,omitempty"`

	// User Id, can be only first parameter
	UserID int64 `json:"userId,omitempty"`
}

// Validate validates this permission
func (m *Permission) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this permission based on context it is used
func (m *Permission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Permission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Permission) UnmarshalBinary(b []byte) error {
	var res Permission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
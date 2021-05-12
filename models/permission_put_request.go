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

// PermissionPutRequest Specifies the permission level to assign.
//
// swagger:model permission_put_request
type PermissionPutRequest struct {

	// permission level
	// Required: true
	PermissionLevel *PermissionLevel `json:"permission_level"`
}

// Validate validates this permission put request
func (m *PermissionPutRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissionLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PermissionPutRequest) validatePermissionLevel(formats strfmt.Registry) error {

	if err := validate.Required("permission_level", "body", m.PermissionLevel); err != nil {
		return err
	}

	if err := validate.Required("permission_level", "body", m.PermissionLevel); err != nil {
		return err
	}

	if m.PermissionLevel != nil {
		if err := m.PermissionLevel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permission_level")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this permission put request based on the context it is used
func (m *PermissionPutRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePermissionLevel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PermissionPutRequest) contextValidatePermissionLevel(ctx context.Context, formats strfmt.Registry) error {

	if m.PermissionLevel != nil {
		if err := m.PermissionLevel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permission_level")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PermissionPutRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PermissionPutRequest) UnmarshalBinary(b []byte) error {
	var res PermissionPutRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

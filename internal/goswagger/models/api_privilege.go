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

// APIPrivilege Api privilege
//
// swagger:model ApiPrivilege
type APIPrivilege struct {

	// description
	Description string `json:"description,omitempty"`

	// The name of the privilege.  This value cannot be changed.
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name,omitempty"`

	// Indicates whether the privilege can be changed. External values supplied to this will be ignored by the system.
	ReadOnly bool `json:"readOnly,omitempty"`

	// The type of privilege, each type covers different portions of the system. External values supplied to this will be ignored by the system.
	Type string `json:"type,omitempty"`
}

// Validate validates this Api privilege
func (m *APIPrivilege) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIPrivilege) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Api privilege based on context it is used
func (m *APIPrivilege) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIPrivilege) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPrivilege) UnmarshalBinary(b []byte) error {
	var res APIPrivilege
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
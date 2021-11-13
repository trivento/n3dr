// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GenericBlobStoreAPIResponse generic blob store Api response
//
// swagger:model GenericBlobStoreApiResponse
type GenericBlobStoreAPIResponse struct {

	// available space in bytes
	AvailableSpaceInBytes int64 `json:"availableSpaceInBytes,omitempty"`

	// blob count
	BlobCount int64 `json:"blobCount,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Settings to control the soft quota
	SoftQuota *BlobStoreAPISoftQuota `json:"softQuota,omitempty"`

	// total size in bytes
	TotalSizeInBytes int64 `json:"totalSizeInBytes,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// unavailable
	Unavailable bool `json:"unavailable,omitempty"`
}

// Validate validates this generic blob store Api response
func (m *GenericBlobStoreAPIResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSoftQuota(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenericBlobStoreAPIResponse) validateSoftQuota(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftQuota) { // not required
		return nil
	}

	if m.SoftQuota != nil {
		if err := m.SoftQuota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softQuota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("softQuota")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this generic blob store Api response based on the context it is used
func (m *GenericBlobStoreAPIResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSoftQuota(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenericBlobStoreAPIResponse) contextValidateSoftQuota(ctx context.Context, formats strfmt.Registry) error {

	if m.SoftQuota != nil {
		if err := m.SoftQuota.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softQuota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("softQuota")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GenericBlobStoreAPIResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericBlobStoreAPIResponse) UnmarshalBinary(b []byte) error {
	var res GenericBlobStoreAPIResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

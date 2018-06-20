// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CertificateVariableUsageResourceTenantResource certificate variable usage resource tenant resource
// swagger:model CertificateVariableUsageResource[TenantResource]
type CertificateVariableUsageResourceTenantResource struct {

	// owner
	Owner *TenantResource `json:"Owner,omitempty"`

	// variables
	Variables []string `json:"Variables"`
}

// Validate validates this certificate variable usage resource tenant resource
func (m *CertificateVariableUsageResourceTenantResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateVariableUsageResourceTenantResource) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Owner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertificateVariableUsageResourceTenantResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateVariableUsageResourceTenantResource) UnmarshalBinary(b []byte) error {
	var res CertificateVariableUsageResourceTenantResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

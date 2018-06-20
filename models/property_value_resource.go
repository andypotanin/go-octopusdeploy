// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PropertyValueResource property value resource
// swagger:model PropertyValueResource
type PropertyValueResource struct {

	// is sensitive
	// Read Only: true
	IsSensitive *bool `json:"IsSensitive,omitempty"`

	// sensitive value
	// Read Only: true
	SensitiveValue *SensitiveValue `json:"SensitiveValue,omitempty"`

	// value
	// Read Only: true
	Value string `json:"Value,omitempty"`
}

// Validate validates this property value resource
func (m *PropertyValueResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSensitiveValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PropertyValueResource) validateSensitiveValue(formats strfmt.Registry) error {

	if swag.IsZero(m.SensitiveValue) { // not required
		return nil
	}

	if m.SensitiveValue != nil {
		if err := m.SensitiveValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SensitiveValue")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PropertyValueResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PropertyValueResource) UnmarshalBinary(b []byte) error {
	var res PropertyValueResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

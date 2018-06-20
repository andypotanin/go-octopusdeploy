// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// FormElement form element
// swagger:model FormElement
type FormElement struct {

	// control
	// Read Only: true
	Control Control `json:"Control,omitempty"`

	// is value required
	// Read Only: true
	IsValueRequired *bool `json:"IsValueRequired,omitempty"`

	// name
	// Read Only: true
	Name string `json:"Name,omitempty"`
}

// Validate validates this form element
func (m *FormElement) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FormElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FormElement) UnmarshalBinary(b []byte) error {
	var res FormElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TenantVariableResourceLibraryVariables tenant variable resource library variables
// swagger:model tenantVariableResourceLibraryVariables
type TenantVariableResourceLibraryVariables map[string]Library

// Validate validates this tenant variable resource library variables
func (m TenantVariableResourceLibraryVariables) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", TenantVariableResourceLibraryVariables(m)); err != nil {
		return err
	}

	for k := range m {

		if val, ok := m[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

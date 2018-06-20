// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServerConfigurationSettingsResource server configuration settings resource
// swagger:model ServerConfigurationSettingsResource
type ServerConfigurationSettingsResource struct {

	// configuration set
	ConfigurationSet string `json:"ConfigurationSet,omitempty"`

	// configuration values
	ConfigurationValues []*ServerConfigurationValueResource `json:"ConfigurationValues"`
}

// Validate validates this server configuration settings resource
func (m *ServerConfigurationSettingsResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigurationValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerConfigurationSettingsResource) validateConfigurationValues(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigurationValues) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfigurationValues); i++ {
		if swag.IsZero(m.ConfigurationValues[i]) { // not required
			continue
		}

		if m.ConfigurationValues[i] != nil {
			if err := m.ConfigurationValues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ConfigurationValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerConfigurationSettingsResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerConfigurationSettingsResource) UnmarshalBinary(b []byte) error {
	var res ServerConfigurationSettingsResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ReleaseProgressionResourceDeployments release progression resource deployments
// swagger:model releaseProgressionResourceDeployments
type ReleaseProgressionResourceDeployments map[string][]DashboardItemResource

// Validate validates this release progression resource deployments
func (m ReleaseProgressionResourceDeployments) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", ReleaseProgressionResourceDeployments(m)); err != nil {
		return err
	}

	for k := range m {

		if err := validate.Required(k, "body", m[k]); err != nil {
			return err
		}

		for i := 0; i < len(m[k]); i++ {

			if err := m[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

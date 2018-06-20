// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountUsageResource account usage resource
// swagger:model AccountUsageResource
type AccountUsageResource struct {

	// deployment processes
	DeploymentProcesses []*StepUsage `json:"DeploymentProcesses"`

	// Id
	ID string `json:"Id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// library variable sets
	LibraryVariableSets []*LibraryVariableSetUsageEntry `json:"LibraryVariableSets"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// project variable sets
	ProjectVariableSets []*ProjectVariableSetUsage `json:"ProjectVariableSets"`

	// releases
	Releases []*ReleaseUsage `json:"Releases"`

	// targets
	Targets []*TargetUsageEntry `json:"Targets"`
}

// Validate validates this account usage resource
func (m *AccountUsageResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeploymentProcesses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLibraryVariableSets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectVariableSets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountUsageResource) validateDeploymentProcesses(formats strfmt.Registry) error {

	if swag.IsZero(m.DeploymentProcesses) { // not required
		return nil
	}

	for i := 0; i < len(m.DeploymentProcesses); i++ {
		if swag.IsZero(m.DeploymentProcesses[i]) { // not required
			continue
		}

		if m.DeploymentProcesses[i] != nil {
			if err := m.DeploymentProcesses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DeploymentProcesses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountUsageResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountUsageResource) validateLibraryVariableSets(formats strfmt.Registry) error {

	if swag.IsZero(m.LibraryVariableSets) { // not required
		return nil
	}

	for i := 0; i < len(m.LibraryVariableSets); i++ {
		if swag.IsZero(m.LibraryVariableSets[i]) { // not required
			continue
		}

		if m.LibraryVariableSets[i] != nil {
			if err := m.LibraryVariableSets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("LibraryVariableSets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountUsageResource) validateProjectVariableSets(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectVariableSets) { // not required
		return nil
	}

	for i := 0; i < len(m.ProjectVariableSets); i++ {
		if swag.IsZero(m.ProjectVariableSets[i]) { // not required
			continue
		}

		if m.ProjectVariableSets[i] != nil {
			if err := m.ProjectVariableSets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ProjectVariableSets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountUsageResource) validateReleases(formats strfmt.Registry) error {

	if swag.IsZero(m.Releases) { // not required
		return nil
	}

	for i := 0; i < len(m.Releases); i++ {
		if swag.IsZero(m.Releases[i]) { // not required
			continue
		}

		if m.Releases[i] != nil {
			if err := m.Releases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Releases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountUsageResource) validateTargets(formats strfmt.Registry) error {

	if swag.IsZero(m.Targets) { // not required
		return nil
	}

	for i := 0; i < len(m.Targets); i++ {
		if swag.IsZero(m.Targets[i]) { // not required
			continue
		}

		if m.Targets[i] != nil {
			if err := m.Targets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountUsageResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountUsageResource) UnmarshalBinary(b []byte) error {
	var res AccountUsageResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

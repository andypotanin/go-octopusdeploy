// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TriggerActionResource trigger action resource
// swagger:model TriggerActionResource
type TriggerActionResource struct {

	// action type
	// Read Only: true
	// Enum: [AutoDeploy]
	ActionType int32 `json:"ActionType,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`
}

// Validate validates this trigger action resource
func (m *TriggerActionResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var triggerActionResourceTypeActionTypePropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`["AutoDeploy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		triggerActionResourceTypeActionTypePropEnum = append(triggerActionResourceTypeActionTypePropEnum, v)
	}
}

// prop value enum
func (m *TriggerActionResource) validateActionTypeEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, triggerActionResourceTypeActionTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TriggerActionResource) validateActionType(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionTypeEnum("ActionType", "body", m.ActionType); err != nil {
		return err
	}

	return nil
}

func (m *TriggerActionResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TriggerActionResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TriggerActionResource) UnmarshalBinary(b []byte) error {
	var res TriggerActionResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

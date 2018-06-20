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

// SubscriptionResource subscription resource
// swagger:model SubscriptionResource
type SubscriptionResource struct {

	// event notification subscription
	EventNotificationSubscription *EventNotificationSubscription `json:"EventNotificationSubscription,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// is disabled
	IsDisabled bool `json:"IsDisabled,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// type
	// Enum: [Event]
	Type int32 `json:"Type,omitempty"`
}

// Validate validates this subscription resource
func (m *SubscriptionResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventNotificationSubscription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionResource) validateEventNotificationSubscription(formats strfmt.Registry) error {

	if swag.IsZero(m.EventNotificationSubscription) { // not required
		return nil
	}

	if m.EventNotificationSubscription != nil {
		if err := m.EventNotificationSubscription.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EventNotificationSubscription")
			}
			return err
		}
	}

	return nil
}

func (m *SubscriptionResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

var subscriptionResourceTypeTypePropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`["Event"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionResourceTypeTypePropEnum = append(subscriptionResourceTypeTypePropEnum, v)
	}
}

// prop value enum
func (m *SubscriptionResource) validateTypeEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, subscriptionResourceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionResource) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionResource) UnmarshalBinary(b []byte) error {
	var res SubscriptionResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

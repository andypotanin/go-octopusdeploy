// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserRoleResource user role resource
// swagger:model UserRoleResource
type UserRoleResource struct {

	// can be deleted
	CanBeDeleted bool `json:"CanBeDeleted,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// granted permissions
	GrantedPermissions []int32 `json:"GrantedPermissions"`

	// Id
	ID string `json:"Id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// permission descriptions
	PermissionDescriptions []string `json:"PermissionDescriptions"`

	// supported restrictions
	SupportedRestrictions []string `json:"SupportedRestrictions"`
}

// Validate validates this user role resource
func (m *UserRoleResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGrantedPermissions(formats); err != nil {
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

var userRoleResourceGrantedPermissionsItemsEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`["None","AdministerSystem","ProjectEdit","ProjectView","ProjectCreate","ProjectDelete","ProcessView","ProcessEdit","VariableEdit","VariableEditUnscoped","VariableView","VariableViewUnscoped","ReleaseCreate","ReleaseView","ReleaseEdit","ReleaseDelete","DefectReport","DefectResolve","DeploymentCreate","DeploymentDelete","DeploymentView","EnvironmentView","EnvironmentCreate","EnvironmentEdit","EnvironmentDelete","MachineCreate","MachineEdit","MachineView","MachineDelete","ArtifactView","ArtifactCreate","ArtifactEdit","ArtifactDelete","FeedView","EventView","LibraryVariableSetView","LibraryVariableSetCreate","LibraryVariableSetEdit","LibraryVariableSetDelete","ProjectGroupView","ProjectGroupCreate","ProjectGroupEdit","ProjectGroupDelete","TeamCreate","TeamView","TeamEdit","TeamDelete","UserView","UserInvite","UserRoleView","UserRoleEdit","TaskView","TaskViewLog","TaskCreate","TaskCancel","TaskEdit","InterruptionView","InterruptionSubmit","InterruptionViewSubmitResponsible","BuiltInFeedPush","BuiltInFeedAdminister","BuiltInFeedDownload","ActionTemplateView","ActionTemplateCreate","ActionTemplateEdit","ActionTemplateDelete","LifecycleCreate","LifecycleView","LifecycleEdit","LifecycleDelete","AccountView","AccountEdit","AccountCreate","AccountDelete","AuditView","TenantCreate","TenantEdit","TenantView","TenantDelete","TagSetCreate","TagSetEdit","TagSetDelete","MachinePolicyCreate","MachinePolicyView","MachinePolicyEdit","MachinePolicyDelete","ProxyCreate","ProxyView","ProxyEdit","ProxyDelete","SubscriptionCreate","SubscriptionView","SubscriptionEdit","SubscriptionDelete","TriggerCreate","TriggerView","TriggerEdit","TriggerDelete","CertificateView","CertificateCreate","CertificateEdit","CertificateDelete","CertificateExportPrivateKey","UserEdit","ConfigureServer","FeedEdit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userRoleResourceGrantedPermissionsItemsEnum = append(userRoleResourceGrantedPermissionsItemsEnum, v)
	}
}

func (m *UserRoleResource) validateGrantedPermissionsItemsEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, userRoleResourceGrantedPermissionsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *UserRoleResource) validateGrantedPermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.GrantedPermissions) { // not required
		return nil
	}

	for i := 0; i < len(m.GrantedPermissions); i++ {

		// value enum
		if err := m.validateGrantedPermissionsItemsEnum("GrantedPermissions"+"."+strconv.Itoa(i), "body", m.GrantedPermissions[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *UserRoleResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserRoleResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserRoleResource) UnmarshalBinary(b []byte) error {
	var res UserRoleResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

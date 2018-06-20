// Code generated by go-swagger; DO NOT EDIT.

package maintenance_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions maintenance configuration update action operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions maintenance configuration update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions maintenance configuration update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions maintenance configuration update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions maintenance configuration update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions maintenance configuration update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions maintenance configuration update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsMaintenanceConfigurationUpdateActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

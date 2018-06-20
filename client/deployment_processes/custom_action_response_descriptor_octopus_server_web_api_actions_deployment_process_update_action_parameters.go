// Code generated by go-swagger; DO NOT EDIT.

package deployment_processes

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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions deployment process update action operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions deployment process update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions deployment process update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions deployment process update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions deployment process update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions deployment process update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions deployment process update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom action response descriptor octopus server web Api actions deployment process update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams) WithID(id string) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom action response descriptor octopus server web Api actions deployment process update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

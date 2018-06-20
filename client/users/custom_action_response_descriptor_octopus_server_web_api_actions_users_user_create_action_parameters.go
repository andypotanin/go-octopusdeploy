// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions users user create action operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions users user create action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions users user create action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions users user create action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions users user create action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions users user create action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions users user create action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserCreateActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

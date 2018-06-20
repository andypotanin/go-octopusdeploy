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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions users user logout action operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions users user logout action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions users user logout action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions users user logout action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions users user logout action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions users user logout action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions users user logout action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLogoutActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

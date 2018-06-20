// Code generated by go-swagger; DO NOT EDIT.

package lets_encrypt

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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions lets encrypt configuration update action operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions lets encrypt configuration update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions lets encrypt configuration update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions lets encrypt configuration update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions lets encrypt configuration update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions lets encrypt configuration update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions lets encrypt configuration update action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLetsEncryptConfigurationUpdateActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

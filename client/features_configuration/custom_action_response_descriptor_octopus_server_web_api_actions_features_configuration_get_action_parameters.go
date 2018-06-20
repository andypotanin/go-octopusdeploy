// Code generated by go-swagger; DO NOT EDIT.

package features_configuration

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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions features configuration get action operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions features configuration get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions features configuration get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions features configuration get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions features configuration get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions features configuration get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions features configuration get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsFeaturesConfigurationGetActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

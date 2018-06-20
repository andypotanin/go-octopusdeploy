// Code generated by go-swagger; DO NOT EDIT.

package environments

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

// NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams creates a new CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams object
// with the default values initialized.
func NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams() *CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams {

	return &CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParamsWithTimeout creates a new CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParamsWithTimeout(timeout time.Duration) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams {

	return &CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams{

		timeout: timeout,
	}
}

// NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParamsWithContext creates a new CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParamsWithContext(ctx context.Context) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams {

	return &CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams{

		Context: ctx,
	}
}

// NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParamsWithHTTPClient creates a new CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParamsWithHTTPClient(client *http.Client) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams {

	return &CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams{
		HTTPClient: client,
	}
}

/*CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams contains all the parameters to send to the API endpoint
for the custom query response descriptor octopus server web Api actions environments summary responder operation typically these are written to a http.Request
*/
type CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom query response descriptor octopus server web Api actions environments summary responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams) WithTimeout(timeout time.Duration) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom query response descriptor octopus server web Api actions environments summary responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom query response descriptor octopus server web Api actions environments summary responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams) WithContext(ctx context.Context) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom query response descriptor octopus server web Api actions environments summary responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom query response descriptor octopus server web Api actions environments summary responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams) WithHTTPClient(client *http.Client) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom query response descriptor octopus server web Api actions environments summary responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsEnvironmentsSummaryResponderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package interruptions

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

// NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams creates a new CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams object
// with the default values initialized.
func NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams() *CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams {

	return &CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParamsWithTimeout creates a new CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParamsWithTimeout(timeout time.Duration) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams {

	return &CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams{

		timeout: timeout,
	}
}

// NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParamsWithContext creates a new CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParamsWithContext(ctx context.Context) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams {

	return &CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams{

		Context: ctx,
	}
}

// NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParamsWithHTTPClient creates a new CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParamsWithHTTPClient(client *http.Client) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams {

	return &CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams{
		HTTPClient: client,
	}
}

/*CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams contains all the parameters to send to the API endpoint
for the custom query response descriptor octopus server web Api actions list interruptions responder operation typically these are written to a http.Request
*/
type CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom query response descriptor octopus server web Api actions list interruptions responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams) WithTimeout(timeout time.Duration) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom query response descriptor octopus server web Api actions list interruptions responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom query response descriptor octopus server web Api actions list interruptions responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams) WithContext(ctx context.Context) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom query response descriptor octopus server web Api actions list interruptions responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom query response descriptor octopus server web Api actions list interruptions responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams) WithHTTPClient(client *http.Client) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom query response descriptor octopus server web Api actions list interruptions responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsListInterruptionsResponderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

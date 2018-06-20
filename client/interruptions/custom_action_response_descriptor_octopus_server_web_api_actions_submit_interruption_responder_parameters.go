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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions submit interruption responder operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions submit interruption responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions submit interruption responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions submit interruption responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions submit interruption responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions submit interruption responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions submit interruption responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom action response descriptor octopus server web Api actions submit interruption responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams) WithID(id string) *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom action response descriptor octopus server web Api actions submit interruption responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSubmitInterruptionResponderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

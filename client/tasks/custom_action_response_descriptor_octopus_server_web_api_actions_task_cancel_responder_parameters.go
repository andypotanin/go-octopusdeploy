// Code generated by go-swagger; DO NOT EDIT.

package tasks

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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions task cancel responder operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions task cancel responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions task cancel responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions task cancel responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions task cancel responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions task cancel responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions task cancel responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom action response descriptor octopus server web Api actions task cancel responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams) WithID(id string) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom action response descriptor octopus server web Api actions task cancel responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskCancelResponderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

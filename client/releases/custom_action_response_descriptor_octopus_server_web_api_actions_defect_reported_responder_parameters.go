// Code generated by go-swagger; DO NOT EDIT.

package releases

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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions defect reported responder operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions defect reported responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions defect reported responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions defect reported responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions defect reported responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions defect reported responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions defect reported responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom action response descriptor octopus server web Api actions defect reported responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams) WithID(id string) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom action response descriptor octopus server web Api actions defect reported responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDefectReportedResponderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

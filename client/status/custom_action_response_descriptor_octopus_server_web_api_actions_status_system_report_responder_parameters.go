// Code generated by go-swagger; DO NOT EDIT.

package status

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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions status system report responder operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions status system report responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions status system report responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions status system report responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions status system report responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions status system report responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions status system report responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package certificates

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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions certificates certificate un archive action operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions certificates certificate un archive action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions certificates certificate un archive action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions certificates certificate un archive action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions certificates certificate un archive action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions certificates certificate un archive action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions certificates certificate un archive action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom action response descriptor octopus server web Api actions certificates certificate un archive action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams) WithID(id string) *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom action response descriptor octopus server web Api actions certificates certificate un archive action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificateUnArchiveActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package proxies

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

// NewListAllResponseDescriptorEnvironmentsProxyProxyResourceParams creates a new ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams object
// with the default values initialized.
func NewListAllResponseDescriptorEnvironmentsProxyProxyResourceParams() *ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams {

	return &ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllResponseDescriptorEnvironmentsProxyProxyResourceParamsWithTimeout creates a new ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllResponseDescriptorEnvironmentsProxyProxyResourceParamsWithTimeout(timeout time.Duration) *ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams {

	return &ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams{

		timeout: timeout,
	}
}

// NewListAllResponseDescriptorEnvironmentsProxyProxyResourceParamsWithContext creates a new ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllResponseDescriptorEnvironmentsProxyProxyResourceParamsWithContext(ctx context.Context) *ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams {

	return &ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams{

		Context: ctx,
	}
}

// NewListAllResponseDescriptorEnvironmentsProxyProxyResourceParamsWithHTTPClient creates a new ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllResponseDescriptorEnvironmentsProxyProxyResourceParamsWithHTTPClient(client *http.Client) *ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams {

	return &ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams{
		HTTPClient: client,
	}
}

/*ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams contains all the parameters to send to the API endpoint
for the list all response descriptor environments proxy proxy resource operation typically these are written to a http.Request
*/
type ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all response descriptor environments proxy proxy resource params
func (o *ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams) WithTimeout(timeout time.Duration) *ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all response descriptor environments proxy proxy resource params
func (o *ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all response descriptor environments proxy proxy resource params
func (o *ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams) WithContext(ctx context.Context) *ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all response descriptor environments proxy proxy resource params
func (o *ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all response descriptor environments proxy proxy resource params
func (o *ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams) WithHTTPClient(client *http.Client) *ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all response descriptor environments proxy proxy resource params
func (o *ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllResponseDescriptorEnvironmentsProxyProxyResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

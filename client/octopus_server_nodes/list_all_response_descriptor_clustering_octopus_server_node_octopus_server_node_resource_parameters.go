// Code generated by go-swagger; DO NOT EDIT.

package octopus_server_nodes

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

// NewListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams creates a new ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams object
// with the default values initialized.
func NewListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams() *ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams {

	return &ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParamsWithTimeout creates a new ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParamsWithTimeout(timeout time.Duration) *ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams {

	return &ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams{

		timeout: timeout,
	}
}

// NewListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParamsWithContext creates a new ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParamsWithContext(ctx context.Context) *ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams {

	return &ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams{

		Context: ctx,
	}
}

// NewListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParamsWithHTTPClient creates a new ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParamsWithHTTPClient(client *http.Client) *ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams {

	return &ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams{
		HTTPClient: client,
	}
}

/*ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams contains all the parameters to send to the API endpoint
for the list all response descriptor clustering octopus server node octopus server node resource operation typically these are written to a http.Request
*/
type ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all response descriptor clustering octopus server node octopus server node resource params
func (o *ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) WithTimeout(timeout time.Duration) *ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all response descriptor clustering octopus server node octopus server node resource params
func (o *ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all response descriptor clustering octopus server node octopus server node resource params
func (o *ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) WithContext(ctx context.Context) *ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all response descriptor clustering octopus server node octopus server node resource params
func (o *ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all response descriptor clustering octopus server node octopus server node resource params
func (o *ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) WithHTTPClient(client *http.Client) *ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all response descriptor clustering octopus server node octopus server node resource params
func (o *ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

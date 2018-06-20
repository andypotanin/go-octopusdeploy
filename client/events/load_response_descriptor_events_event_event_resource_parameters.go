// Code generated by go-swagger; DO NOT EDIT.

package events

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

// NewLoadResponseDescriptorEventsEventEventResourceParams creates a new LoadResponseDescriptorEventsEventEventResourceParams object
// with the default values initialized.
func NewLoadResponseDescriptorEventsEventEventResourceParams() *LoadResponseDescriptorEventsEventEventResourceParams {
	var ()
	return &LoadResponseDescriptorEventsEventEventResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoadResponseDescriptorEventsEventEventResourceParamsWithTimeout creates a new LoadResponseDescriptorEventsEventEventResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoadResponseDescriptorEventsEventEventResourceParamsWithTimeout(timeout time.Duration) *LoadResponseDescriptorEventsEventEventResourceParams {
	var ()
	return &LoadResponseDescriptorEventsEventEventResourceParams{

		timeout: timeout,
	}
}

// NewLoadResponseDescriptorEventsEventEventResourceParamsWithContext creates a new LoadResponseDescriptorEventsEventEventResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoadResponseDescriptorEventsEventEventResourceParamsWithContext(ctx context.Context) *LoadResponseDescriptorEventsEventEventResourceParams {
	var ()
	return &LoadResponseDescriptorEventsEventEventResourceParams{

		Context: ctx,
	}
}

// NewLoadResponseDescriptorEventsEventEventResourceParamsWithHTTPClient creates a new LoadResponseDescriptorEventsEventEventResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoadResponseDescriptorEventsEventEventResourceParamsWithHTTPClient(client *http.Client) *LoadResponseDescriptorEventsEventEventResourceParams {
	var ()
	return &LoadResponseDescriptorEventsEventEventResourceParams{
		HTTPClient: client,
	}
}

/*LoadResponseDescriptorEventsEventEventResourceParams contains all the parameters to send to the API endpoint
for the load response descriptor events event event resource operation typically these are written to a http.Request
*/
type LoadResponseDescriptorEventsEventEventResourceParams struct {

	/*ID
	  ID of the EventResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the load response descriptor events event event resource params
func (o *LoadResponseDescriptorEventsEventEventResourceParams) WithTimeout(timeout time.Duration) *LoadResponseDescriptorEventsEventEventResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the load response descriptor events event event resource params
func (o *LoadResponseDescriptorEventsEventEventResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the load response descriptor events event event resource params
func (o *LoadResponseDescriptorEventsEventEventResourceParams) WithContext(ctx context.Context) *LoadResponseDescriptorEventsEventEventResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the load response descriptor events event event resource params
func (o *LoadResponseDescriptorEventsEventEventResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the load response descriptor events event event resource params
func (o *LoadResponseDescriptorEventsEventEventResourceParams) WithHTTPClient(client *http.Client) *LoadResponseDescriptorEventsEventEventResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the load response descriptor events event event resource params
func (o *LoadResponseDescriptorEventsEventEventResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the load response descriptor events event event resource params
func (o *LoadResponseDescriptorEventsEventEventResourceParams) WithID(id string) *LoadResponseDescriptorEventsEventEventResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the load response descriptor events event event resource params
func (o *LoadResponseDescriptorEventsEventEventResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LoadResponseDescriptorEventsEventEventResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

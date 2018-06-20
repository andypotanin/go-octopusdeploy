// Code generated by go-swagger; DO NOT EDIT.

package api_keys

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

// NewDeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams creates a new DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams object
// with the default values initialized.
func NewDeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams() *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParamsWithTimeout creates a new DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParamsWithTimeout(timeout time.Duration) *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams{

		timeout: timeout,
	}
}

// NewDeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParamsWithContext creates a new DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParamsWithContext(ctx context.Context) *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams{

		Context: ctx,
	}
}

// NewDeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParamsWithHTTPClient creates a new DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParamsWithHTTPClient(client *http.Client) *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams{
		HTTPClient: client,
	}
}

/*DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams contains all the parameters to send to the API endpoint
for the delete on background response descriptor users Api key Api key resource operation typically these are written to a http.Request
*/
type DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams struct {

	/*ID
	  ID of the ApiKeyResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete on background response descriptor users Api key Api key resource params
func (o *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams) WithTimeout(timeout time.Duration) *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete on background response descriptor users Api key Api key resource params
func (o *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete on background response descriptor users Api key Api key resource params
func (o *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams) WithContext(ctx context.Context) *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete on background response descriptor users Api key Api key resource params
func (o *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete on background response descriptor users Api key Api key resource params
func (o *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams) WithHTTPClient(client *http.Client) *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete on background response descriptor users Api key Api key resource params
func (o *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete on background response descriptor users Api key Api key resource params
func (o *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams) WithID(id string) *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete on background response descriptor users Api key Api key resource params
func (o *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOnBackgroundResponseDescriptorUsersAPIKeyAPIKeyResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

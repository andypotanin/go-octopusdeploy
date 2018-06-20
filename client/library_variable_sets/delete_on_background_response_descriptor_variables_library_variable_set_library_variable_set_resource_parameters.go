// Code generated by go-swagger; DO NOT EDIT.

package library_variable_sets

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

// NewDeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams creates a new DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams object
// with the default values initialized.
func NewDeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams() *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParamsWithTimeout creates a new DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParamsWithTimeout(timeout time.Duration) *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams{

		timeout: timeout,
	}
}

// NewDeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParamsWithContext creates a new DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParamsWithContext(ctx context.Context) *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams{

		Context: ctx,
	}
}

// NewDeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParamsWithHTTPClient creates a new DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParamsWithHTTPClient(client *http.Client) *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams{
		HTTPClient: client,
	}
}

/*DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams contains all the parameters to send to the API endpoint
for the delete on background response descriptor variables library variable set library variable set resource operation typically these are written to a http.Request
*/
type DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams struct {

	/*ID
	  ID of the LibraryVariableSetResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete on background response descriptor variables library variable set library variable set resource params
func (o *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams) WithTimeout(timeout time.Duration) *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete on background response descriptor variables library variable set library variable set resource params
func (o *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete on background response descriptor variables library variable set library variable set resource params
func (o *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams) WithContext(ctx context.Context) *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete on background response descriptor variables library variable set library variable set resource params
func (o *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete on background response descriptor variables library variable set library variable set resource params
func (o *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams) WithHTTPClient(client *http.Client) *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete on background response descriptor variables library variable set library variable set resource params
func (o *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete on background response descriptor variables library variable set library variable set resource params
func (o *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams) WithID(id string) *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete on background response descriptor variables library variable set library variable set resource params
func (o *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOnBackgroundResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

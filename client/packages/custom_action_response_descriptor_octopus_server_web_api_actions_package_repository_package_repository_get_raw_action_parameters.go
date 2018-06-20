// Code generated by go-swagger; DO NOT EDIT.

package packages

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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions package repository package repository get raw action operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions package repository package repository get raw action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions package repository package repository get raw action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions package repository package repository get raw action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions package repository package repository get raw action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions package repository package repository get raw action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions package repository package repository get raw action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom action response descriptor octopus server web Api actions package repository package repository get raw action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams) WithID(id string) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom action response descriptor octopus server web Api actions package repository package repository get raw action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetRawActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

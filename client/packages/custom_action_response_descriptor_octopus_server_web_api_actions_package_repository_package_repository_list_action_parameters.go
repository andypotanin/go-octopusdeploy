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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions package repository package repository list action operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions package repository package repository list action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions package repository package repository list action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions package repository package repository list action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions package repository package repository list action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions package repository package repository list action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions package repository package repository list action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryListActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

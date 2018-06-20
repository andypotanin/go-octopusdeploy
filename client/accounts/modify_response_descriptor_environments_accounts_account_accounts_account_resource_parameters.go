// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// NewModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams creates a new ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams object
// with the default values initialized.
func NewModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams() *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams {
	var ()
	return &ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParamsWithTimeout creates a new ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParamsWithTimeout(timeout time.Duration) *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams {
	var ()
	return &ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams{

		timeout: timeout,
	}
}

// NewModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParamsWithContext creates a new ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParamsWithContext(ctx context.Context) *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams {
	var ()
	return &ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams{

		Context: ctx,
	}
}

// NewModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParamsWithHTTPClient creates a new ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParamsWithHTTPClient(client *http.Client) *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams {
	var ()
	return &ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams{
		HTTPClient: client,
	}
}

/*ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams contains all the parameters to send to the API endpoint
for the modify response descriptor environments accounts account accounts account resource operation typically these are written to a http.Request
*/
type ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams struct {

	/*AccountResource
	  The AccountResource resource to create

	*/
	AccountResource *models.AccountResource
	/*ID
	  ID of the AccountResource to modify

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify response descriptor environments accounts account accounts account resource params
func (o *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams) WithTimeout(timeout time.Duration) *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify response descriptor environments accounts account accounts account resource params
func (o *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify response descriptor environments accounts account accounts account resource params
func (o *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams) WithContext(ctx context.Context) *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify response descriptor environments accounts account accounts account resource params
func (o *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify response descriptor environments accounts account accounts account resource params
func (o *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams) WithHTTPClient(client *http.Client) *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify response descriptor environments accounts account accounts account resource params
func (o *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountResource adds the accountResource to the modify response descriptor environments accounts account accounts account resource params
func (o *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams) WithAccountResource(accountResource *models.AccountResource) *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams {
	o.SetAccountResource(accountResource)
	return o
}

// SetAccountResource adds the accountResource to the modify response descriptor environments accounts account accounts account resource params
func (o *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams) SetAccountResource(accountResource *models.AccountResource) {
	o.AccountResource = accountResource
}

// WithID adds the id to the modify response descriptor environments accounts account accounts account resource params
func (o *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams) WithID(id string) *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the modify response descriptor environments accounts account accounts account resource params
func (o *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountResource != nil {
		if err := r.SetBodyParam(o.AccountResource); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

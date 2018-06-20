// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams creates a new ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams object
// with the default values initialized.
func NewChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams() *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams {
	var ()
	return &ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParamsWithTimeout creates a new ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParamsWithTimeout(timeout time.Duration) *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams {
	var ()
	return &ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams{

		timeout: timeout,
	}
}

// NewChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParamsWithContext creates a new ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParamsWithContext(ctx context.Context) *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams {
	var ()
	return &ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams{

		Context: ctx,
	}
}

// NewChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParamsWithHTTPClient creates a new ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParamsWithHTTPClient(client *http.Client) *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams {
	var ()
	return &ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams{
		HTTPClient: client,
	}
}

/*ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams contains all the parameters to send to the API endpoint
for the child index response descriptor projects project projects channel project resource channel resource operation typically these are written to a http.Request
*/
type ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams struct {

	/*Skip
	  Number of items to skip

	*/
	Skip *int32
	/*Take
	  Number of items to take

	*/
	Take *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the child index response descriptor projects project projects channel project resource channel resource params
func (o *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams) WithTimeout(timeout time.Duration) *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the child index response descriptor projects project projects channel project resource channel resource params
func (o *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the child index response descriptor projects project projects channel project resource channel resource params
func (o *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams) WithContext(ctx context.Context) *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the child index response descriptor projects project projects channel project resource channel resource params
func (o *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the child index response descriptor projects project projects channel project resource channel resource params
func (o *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams) WithHTTPClient(client *http.Client) *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the child index response descriptor projects project projects channel project resource channel resource params
func (o *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSkip adds the skip to the child index response descriptor projects project projects channel project resource channel resource params
func (o *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams) WithSkip(skip *int32) *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the child index response descriptor projects project projects channel project resource channel resource params
func (o *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the child index response descriptor projects project projects channel project resource channel resource params
func (o *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams) WithTake(take *int32) *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the child index response descriptor projects project projects channel project resource channel resource params
func (o *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Skip != nil {

		// query param skip
		var qrSkip int32
		if o.Skip != nil {
			qrSkip = *o.Skip
		}
		qSkip := swag.FormatInt32(qrSkip)
		if qSkip != "" {
			if err := r.SetQueryParam("skip", qSkip); err != nil {
				return err
			}
		}

	}

	if o.Take != nil {

		// query param take
		var qrTake int32
		if o.Take != nil {
			qrTake = *o.Take
		}
		qTake := swag.FormatInt32(qrTake)
		if qTake != "" {
			if err := r.SetQueryParam("take", qTake); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

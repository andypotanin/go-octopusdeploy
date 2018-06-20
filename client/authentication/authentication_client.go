// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new authentication API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authentication API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CustomActionResponseDescriptorOctopusServerWebAPIActionsAuthenticationAuthenticationResponder Provides the details of the enabled authentication providers.
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomActionResponseDescriptorOctopusServerWebAPIActionsAuthenticationAuthenticationResponder(params *CustomActionResponseDescriptorOctopusServerWebAPIActionsAuthenticationAuthenticationResponderParams) (*CustomActionResponseDescriptorOctopusServerWebAPIActionsAuthenticationAuthenticationResponderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAuthenticationAuthenticationResponderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomActionResponseDescriptor[Octopus.Server.Web.Api.Actions.Authentication.AuthenticationResponder]",
		Method:             "GET",
		PathPattern:        "/api/authentication",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomActionResponseDescriptorOctopusServerWebAPIActionsAuthenticationAuthenticationResponderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomActionResponseDescriptorOctopusServerWebAPIActionsAuthenticationAuthenticationResponderOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new tenants API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tenants API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateResponseDescriptorTenantsTenantTenantResource creates a tenant resource

Creates a new tenant.
*/
func (a *Client) CreateResponseDescriptorTenantsTenantTenantResource(params *CreateResponseDescriptorTenantsTenantTenantResourceParams, authInfo runtime.ClientAuthInfoWriter) (*CreateResponseDescriptorTenantsTenantTenantResourceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateResponseDescriptorTenantsTenantTenantResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateResponseDescriptor[Tenants.Tenant,TenantResource]",
		Method:             "POST",
		PathPattern:        "/api/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateResponseDescriptorTenantsTenantTenantResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateResponseDescriptorTenantsTenantTenantResourceCreated), nil

}

/*
CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantLogoGetResponder Gets the logo associated with the tenant.
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantLogoGetResponder(params *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantLogoGetResponderParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantLogoGetResponderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantLogoGetResponderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomActionResponseDescriptor[Octopus.Server.Web.Api.Actions.TenantLogoGetResponder]",
		Method:             "GET",
		PathPattern:        "/api/tenants/{id}/logo",
		ProducesMediaTypes: []string{"image/png"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantLogoGetResponderReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantLogoGetResponderOK), nil

}

/*
CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantLogoPutResponder Updates the logo associated with the tenant.
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantLogoPutResponder(params *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantLogoPutResponderParams, authInfo runtime.ClientAuthInfoWriter) (*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantLogoPutResponderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantLogoPutResponderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomActionResponseDescriptor[Octopus.Server.Web.Api.Actions.TenantLogoPutResponder]",
		Method:             "PUT",
		PathPattern:        "/api/tenants/{id}/logo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantLogoPutResponderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantLogoPutResponderOK), nil

}

/*
CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantTagTestAction Checks tenants for matching tags rule
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantTagTestAction(params *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantTagTestActionParams, authInfo runtime.ClientAuthInfoWriter) (*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantTagTestActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantTagTestActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomActionResponseDescriptor[Octopus.Server.Web.Api.Actions.TenantTagTestAction]",
		Method:             "GET",
		PathPattern:        "/api/tenants/tag-test",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantTagTestActionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantTagTestActionOK), nil

}

/*
CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariableSetTestAction Checks tenants for matching variable set rule
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariableSetTestAction(params *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariableSetTestActionParams, authInfo runtime.ClientAuthInfoWriter) (*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariableSetTestActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariableSetTestActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomActionResponseDescriptor[Octopus.Server.Web.Api.Actions.TenantVariableSetTestAction]",
		Method:             "GET",
		PathPattern:        "/api/tenants/variableset-test",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariableSetTestActionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariableSetTestActionOK), nil

}

/*
CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponder Gets all the available variables associated with the tenant.
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponder(params *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams, authInfo runtime.ClientAuthInfoWriter) (*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomActionResponseDescriptor[Octopus.Server.Web.Api.Actions.TenantVariablesGetResponder]",
		Method:             "GET",
		PathPattern:        "/api/tenants/{id}/variables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderOK), nil

}

/*
CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesMissingAction Returns list of tenants who are missing required variables
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesMissingAction(params *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesMissingActionParams, authInfo runtime.ClientAuthInfoWriter) (*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesMissingActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesMissingActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomActionResponseDescriptor[Octopus.Server.Web.Api.Actions.TenantVariablesMissingAction]",
		Method:             "GET",
		PathPattern:        "/api/tenants/variables-missing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesMissingActionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesMissingActionOK), nil

}

/*
CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponder Updates the variables associated with the tenant.
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponder(params *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderParams, authInfo runtime.ClientAuthInfoWriter) (*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomActionResponseDescriptor[Octopus.Server.Web.Api.Actions.TenantVariablesPutResponder]",
		Method:             "PUT",
		PathPattern:        "/api/tenants/{id}/variables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderOK), nil

}

/*
DeleteOnBackgroundResponseDescriptorTenantsTenantTenantResource deletes a tenant resource by ID

Deletes an existing tenant.
*/
func (a *Client) DeleteOnBackgroundResponseDescriptorTenantsTenantTenantResource(params *DeleteOnBackgroundResponseDescriptorTenantsTenantTenantResourceParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOnBackgroundResponseDescriptorTenantsTenantTenantResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOnBackgroundResponseDescriptorTenantsTenantTenantResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteOnBackgroundResponseDescriptor[Tenants.Tenant,TenantResource]",
		Method:             "DELETE",
		PathPattern:        "/api/tenants/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteOnBackgroundResponseDescriptorTenantsTenantTenantResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOnBackgroundResponseDescriptorTenantsTenantTenantResourceOK), nil

}

/*
IndexResponseDescriptorTenantsTenantTenantResource gets a list of tenant resources

Lists all of the tenants in the current Octopus installation. The results will be sorted alphabetically by name, and returned 30 at a time.
*/
func (a *Client) IndexResponseDescriptorTenantsTenantTenantResource(params *IndexResponseDescriptorTenantsTenantTenantResourceParams, authInfo runtime.ClientAuthInfoWriter) (*IndexResponseDescriptorTenantsTenantTenantResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexResponseDescriptorTenantsTenantTenantResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "IndexResponseDescriptor[Tenants.Tenant,TenantResource]",
		Method:             "GET",
		PathPattern:        "/api/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexResponseDescriptorTenantsTenantTenantResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IndexResponseDescriptorTenantsTenantTenantResourceOK), nil

}

/*
ListAllResponseDescriptorTenantsTenantTenantResource gets a list of tenant resources

Lists all of the tenants in the current Octopus installation. The results will be sorted alphabetically by name.
*/
func (a *Client) ListAllResponseDescriptorTenantsTenantTenantResource(params *ListAllResponseDescriptorTenantsTenantTenantResourceParams, authInfo runtime.ClientAuthInfoWriter) (*ListAllResponseDescriptorTenantsTenantTenantResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllResponseDescriptorTenantsTenantTenantResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListAllResponseDescriptor[Tenants.Tenant,TenantResource]",
		Method:             "GET",
		PathPattern:        "/api/tenants/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAllResponseDescriptorTenantsTenantTenantResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListAllResponseDescriptorTenantsTenantTenantResourceOK), nil

}

/*
LoadResponseDescriptorTenantsTenantTenantResource gets a tenant resource by ID

Gets a single tenant by ID.
*/
func (a *Client) LoadResponseDescriptorTenantsTenantTenantResource(params *LoadResponseDescriptorTenantsTenantTenantResourceParams, authInfo runtime.ClientAuthInfoWriter) (*LoadResponseDescriptorTenantsTenantTenantResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoadResponseDescriptorTenantsTenantTenantResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LoadResponseDescriptor[Tenants.Tenant,TenantResource]",
		Method:             "GET",
		PathPattern:        "/api/tenants/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoadResponseDescriptorTenantsTenantTenantResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LoadResponseDescriptorTenantsTenantTenantResourceOK), nil

}

/*
ModifyResponseDescriptorTenantsTenantTenantResource modifies a tenant resource by ID

Modifies an existing tenant.
*/
func (a *Client) ModifyResponseDescriptorTenantsTenantTenantResource(params *ModifyResponseDescriptorTenantsTenantTenantResourceParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyResponseDescriptorTenantsTenantTenantResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyResponseDescriptorTenantsTenantTenantResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyResponseDescriptor[Tenants.Tenant,TenantResource]",
		Method:             "PUT",
		PathPattern:        "/api/tenants/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyResponseDescriptorTenantsTenantTenantResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyResponseDescriptorTenantsTenantTenantResourceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

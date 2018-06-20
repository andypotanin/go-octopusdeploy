// Code generated by go-swagger; DO NOT EDIT.

package deployment_processes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new deployment processes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deployment processes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateAction Modifies a deployment process. Only allowed for deployment processes owned by a project (cannot be used to change the deployment process owned by a release).
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateAction(params *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams, authInfo runtime.ClientAuthInfoWriter) (*CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomActionResponseDescriptor[Octopus.Server.Web.Api.Actions.DeploymentProcessUpdateAction]",
		Method:             "PUT",
		PathPattern:        "/api/deploymentprocesses/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionOK), nil

}

/*
CustomActionResponseDescriptorOctopusServerWebAPIActionsReleaseTemplateAction Gets all of the information necessary for creating or editing a release using this deployment process.
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomActionResponseDescriptorOctopusServerWebAPIActionsReleaseTemplateAction(params *CustomActionResponseDescriptorOctopusServerWebAPIActionsReleaseTemplateActionParams, authInfo runtime.ClientAuthInfoWriter) (*CustomActionResponseDescriptorOctopusServerWebAPIActionsReleaseTemplateActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomActionResponseDescriptorOctopusServerWebAPIActionsReleaseTemplateActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomActionResponseDescriptor[Octopus.Server.Web.Api.Actions.ReleaseTemplateAction]",
		Method:             "GET",
		PathPattern:        "/api/deploymentprocesses/{id}/template",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomActionResponseDescriptorOctopusServerWebAPIActionsReleaseTemplateActionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomActionResponseDescriptorOctopusServerWebAPIActionsReleaseTemplateActionOK), nil

}

/*
IndexResponseDescriptorProjectsDeploymentProcessDeploymentProcessResource gets a list of deployment process resources

Lists all the deployment processeses in the current Octopus installation, sorted by Id
*/
func (a *Client) IndexResponseDescriptorProjectsDeploymentProcessDeploymentProcessResource(params *IndexResponseDescriptorProjectsDeploymentProcessDeploymentProcessResourceParams, authInfo runtime.ClientAuthInfoWriter) (*IndexResponseDescriptorProjectsDeploymentProcessDeploymentProcessResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexResponseDescriptorProjectsDeploymentProcessDeploymentProcessResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "IndexResponseDescriptor[Projects.DeploymentProcess,DeploymentProcessResource]",
		Method:             "GET",
		PathPattern:        "/api/deploymentprocesses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexResponseDescriptorProjectsDeploymentProcessDeploymentProcessResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IndexResponseDescriptorProjectsDeploymentProcessDeploymentProcessResourceOK), nil

}

/*
LoadResponseDescriptorProjectsDeploymentProcessDeploymentProcessResource gets a deployment process resource by ID

Gets a deployment process by ID.
*/
func (a *Client) LoadResponseDescriptorProjectsDeploymentProcessDeploymentProcessResource(params *LoadResponseDescriptorProjectsDeploymentProcessDeploymentProcessResourceParams, authInfo runtime.ClientAuthInfoWriter) (*LoadResponseDescriptorProjectsDeploymentProcessDeploymentProcessResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoadResponseDescriptorProjectsDeploymentProcessDeploymentProcessResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LoadResponseDescriptor[Projects.DeploymentProcess,DeploymentProcessResource]",
		Method:             "GET",
		PathPattern:        "/api/deploymentprocesses/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoadResponseDescriptorProjectsDeploymentProcessDeploymentProcessResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LoadResponseDescriptorProjectsDeploymentProcessDeploymentProcessResourceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

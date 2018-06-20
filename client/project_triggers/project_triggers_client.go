// Code generated by go-swagger; DO NOT EDIT.

package project_triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new project triggers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for project triggers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ChildIndexResponseDescriptorProjectsProjectProjectsProjectTriggerProjectTriggerProjectResourceProjectTriggerResource gets a list of project trigger resources

Lists all the project triggers for the given project
*/
func (a *Client) ChildIndexResponseDescriptorProjectsProjectProjectsProjectTriggerProjectTriggerProjectResourceProjectTriggerResource(params *ChildIndexResponseDescriptorProjectsProjectProjectsProjectTriggerProjectTriggerProjectResourceProjectTriggerResourceParams, authInfo runtime.ClientAuthInfoWriter) (*ChildIndexResponseDescriptorProjectsProjectProjectsProjectTriggerProjectTriggerProjectResourceProjectTriggerResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChildIndexResponseDescriptorProjectsProjectProjectsProjectTriggerProjectTriggerProjectResourceProjectTriggerResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChildIndexResponseDescriptor[Projects.Project,Projects.ProjectTrigger.ProjectTrigger,ProjectResource,ProjectTriggerResource]",
		Method:             "GET",
		PathPattern:        "/api/projects/{id}/triggers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChildIndexResponseDescriptorProjectsProjectProjectsProjectTriggerProjectTriggerProjectResourceProjectTriggerResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChildIndexResponseDescriptorProjectsProjectProjectsProjectTriggerProjectTriggerProjectResourceProjectTriggerResourceOK), nil

}

/*
CreateResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResource creates a project trigger resource

Creates a new project trigger
*/
func (a *Client) CreateResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResource(params *CreateResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceParams, authInfo runtime.ClientAuthInfoWriter) (*CreateResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateResponseDescriptor[Projects.ProjectTrigger.ProjectTrigger,ProjectTriggerResource]",
		Method:             "POST",
		PathPattern:        "/api/projecttriggers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceCreated), nil

}

/*
DeleteOnBackgroundResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResource deletes a project trigger resource by ID

Deletes an existing project trigger.
*/
func (a *Client) DeleteOnBackgroundResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResource(params *DeleteOnBackgroundResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOnBackgroundResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOnBackgroundResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteOnBackgroundResponseDescriptor[Projects.ProjectTrigger.ProjectTrigger,ProjectTriggerResource]",
		Method:             "DELETE",
		PathPattern:        "/api/projecttriggers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteOnBackgroundResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOnBackgroundResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK), nil

}

/*
IndexResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResource gets a list of project trigger resources

Gets all the project triggers in the current Octopus installation, sorted by Id
*/
func (a *Client) IndexResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResource(params *IndexResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceParams, authInfo runtime.ClientAuthInfoWriter) (*IndexResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIndexResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "IndexResponseDescriptor[Projects.ProjectTrigger.ProjectTrigger,ProjectTriggerResource]",
		Method:             "GET",
		PathPattern:        "/api/projecttriggers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IndexResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IndexResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK), nil

}

/*
LoadResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResource gets a project trigger resource by ID

Get a project trigger
*/
func (a *Client) LoadResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResource(params *LoadResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceParams, authInfo runtime.ClientAuthInfoWriter) (*LoadResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoadResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LoadResponseDescriptor[Projects.ProjectTrigger.ProjectTrigger,ProjectTriggerResource]",
		Method:             "GET",
		PathPattern:        "/api/projecttriggers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoadResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LoadResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK), nil

}

/*
ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResource modifies a project trigger resource by ID

Updates an existing project trigger
*/
func (a *Client) ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResource(params *ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyResponseDescriptor[Projects.ProjectTrigger.ProjectTrigger,ProjectTriggerResource]",
		Method:             "PUT",
		PathPattern:        "/api/projecttriggers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

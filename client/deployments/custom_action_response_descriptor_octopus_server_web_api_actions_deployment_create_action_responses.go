// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateAction structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionCreated creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionCreated with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionCreated() *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionCreated {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionCreated{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionCreated handles this case with default header values.

DeploymentResource resource returned
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionCreated struct {
	Payload *models.DeploymentResource
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionCreated) Error() string {
	return fmt.Sprintf("[POST /api/deployments][%d] customActionResponseDescriptorOctopusServerWebApiActionsDeploymentCreateActionCreated  %+v", 201, o.Payload)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionBadRequest creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionBadRequest with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionBadRequest() *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionBadRequest {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionBadRequest{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionBadRequest handles this case with default header values.

Deployment Schedule is invalid.
License is not compliant.
Deployment creation failed.
No request body was supplied.
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionBadRequest struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/deployments][%d] customActionResponseDescriptorOctopusServerWebApiActionsDeploymentCreateActionBadRequest ", 400)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

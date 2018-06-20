// Code generated by go-swagger; DO NOT EDIT.

package deployment_processes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateAction structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionOK handles this case with default header values.

DeploymentProcessResource resource returned
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionOK struct {
	Payload *models.DeploymentProcessResource
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionOK) Error() string {
	return fmt.Sprintf("[PUT /api/deploymentprocesses/{id}][%d] customActionResponseDescriptorOctopusServerWebApiActionsDeploymentProcessUpdateActionOK  %+v", 200, o.Payload)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentProcessResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionBadRequest creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionBadRequest with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionBadRequest() *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionBadRequest {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionBadRequest{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionBadRequest handles this case with default header values.

Validation failed.
You cannot modify this deployment process because it is frozen. You'll need to create a new deployment process instead.
Changes to this deployment process could not be saved, because another user has made changes to the deployment process between when you started editing and when you saved your changes. Please reload or open a new tab to make your changes.
No request body was supplied.
No id parameter was provided.
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionBadRequest struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/deploymentprocesses/{id}][%d] customActionResponseDescriptorOctopusServerWebApiActionsDeploymentProcessUpdateActionBadRequest ", 400)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentProcessUpdateActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

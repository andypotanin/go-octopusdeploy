// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponder structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderOK handles this case with default header values.

TenantVariableResource resource returned
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderOK struct {
	Payload *models.TenantVariableResource
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderOK) Error() string {
	return fmt.Sprintf("[PUT /api/tenants/{id}/variables][%d] customActionResponseDescriptorOctopusServerWebApiActionsTenantVariablesPutResponderOK  %+v", 200, o.Payload)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantVariableResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderBadRequest creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderBadRequest with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderBadRequest() *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderBadRequest {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderBadRequest{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderBadRequest handles this case with default header values.

No request body was supplied.
No id parameter was provided.
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderBadRequest struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/tenants/{id}/variables][%d] customActionResponseDescriptorOctopusServerWebApiActionsTenantVariablesPutResponderBadRequest ", 400)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesPutResponderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

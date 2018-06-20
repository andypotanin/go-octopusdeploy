// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListAction structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionOK handles this case with default header values.

AzureResourceGroupResource[] resource returned
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionOK struct {
	Payload []*models.AzureResourceGroupResource
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionOK) Error() string {
	return fmt.Sprintf("[GET /api/accounts/{id}/resourceGroups][%d] customActionResponseDescriptorOctopusServerWebApiActionsAzureResourceGroupsListActionOK  %+v", 200, o.Payload)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionBadRequest creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionBadRequest with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionBadRequest() *CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionBadRequest {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionBadRequest{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionBadRequest handles this case with default header values.

No id parameter was provided.
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionBadRequest struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/accounts/{id}/resourceGroups][%d] customActionResponseDescriptorOctopusServerWebApiActionsAzureResourceGroupsListActionBadRequest ", 400)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionInternalServerError creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionInternalServerError with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionInternalServerError() *CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionInternalServerError {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionInternalServerError{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionInternalServerError handles this case with default header values.

Internal Exception
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionInternalServerError struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/accounts/{id}/resourceGroups][%d] customActionResponseDescriptorOctopusServerWebApiActionsAzureResourceGroupsListActionInternalServerError ", 500)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsAzureResourceGroupsListActionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

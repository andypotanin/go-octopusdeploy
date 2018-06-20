// Code generated by go-swagger; DO NOT EDIT.

package progression

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewAction structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionOK handles this case with default header values.

ProgressionResource resource returned
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionOK struct {
	Payload *models.ProgressionResource
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionOK) Error() string {
	return fmt.Sprintf("[GET /api/progression/{id}][%d] customActionResponseDescriptorOctopusServerWebApiActionsProgressionViewActionOK  %+v", 200, o.Payload)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProgressionResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionBadRequest creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionBadRequest with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionBadRequest() *CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionBadRequest {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionBadRequest{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionBadRequest handles this case with default header values.

No id parameter was provided.
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionBadRequest struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/progression/{id}][%d] customActionResponseDescriptorOctopusServerWebApiActionsProgressionViewActionBadRequest ", 400)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsProgressionViewActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

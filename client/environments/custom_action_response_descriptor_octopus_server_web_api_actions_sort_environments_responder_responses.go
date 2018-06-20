// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsSortEnvironmentsResponderReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsSortEnvironmentsResponder structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsSortEnvironmentsResponderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortEnvironmentsResponderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSortEnvironmentsResponderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSortEnvironmentsResponderOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsSortEnvironmentsResponderOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSortEnvironmentsResponderOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortEnvironmentsResponderOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsSortEnvironmentsResponderOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsSortEnvironmentsResponderOK handles this case with default header values.

OK - Default
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsSortEnvironmentsResponderOK struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortEnvironmentsResponderOK) Error() string {
	return fmt.Sprintf("[PUT /api/environments/sortorder][%d] customActionResponseDescriptorOctopusServerWebApiActionsSortEnvironmentsResponderOK ", 200)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortEnvironmentsResponderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

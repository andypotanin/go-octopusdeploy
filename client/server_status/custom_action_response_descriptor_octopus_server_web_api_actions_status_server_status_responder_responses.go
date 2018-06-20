// Code generated by go-swagger; DO NOT EDIT.

package server_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusServerStatusResponderReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusServerStatusResponder structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusServerStatusResponderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusServerStatusResponderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsStatusServerStatusResponderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsStatusServerStatusResponderOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusServerStatusResponderOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsStatusServerStatusResponderOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusServerStatusResponderOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusServerStatusResponderOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusServerStatusResponderOK handles this case with default header values.

OK - Default
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusServerStatusResponderOK struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusServerStatusResponderOK) Error() string {
	return fmt.Sprintf("[GET /api/serverstatus][%d] customActionResponseDescriptorOctopusServerWebApiActionsStatusServerStatusResponderOK ", 200)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusServerStatusResponderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

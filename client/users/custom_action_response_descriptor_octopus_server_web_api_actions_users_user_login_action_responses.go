// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLoginActionReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLoginAction structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLoginActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLoginActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLoginActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLoginActionOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLoginActionOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLoginActionOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLoginActionOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLoginActionOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLoginActionOK handles this case with default header values.

OK - Default
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLoginActionOK struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLoginActionOK) Error() string {
	return fmt.Sprintf("[POST /api/users/login][%d] customActionResponseDescriptorOctopusServerWebApiActionsUsersUserLoginActionOK ", 200)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserLoginActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

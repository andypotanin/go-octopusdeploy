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

// DeleteOnBackgroundResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceReader is a Reader for the DeleteOnBackgroundResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResource structure.
type DeleteOnBackgroundResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOnBackgroundResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteOnBackgroundResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteOnBackgroundResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK creates a DeleteOnBackgroundResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK with default headers values
func NewDeleteOnBackgroundResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK() *DeleteOnBackgroundResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK {
	return &DeleteOnBackgroundResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK{}
}

/*DeleteOnBackgroundResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK handles this case with default header values.

AccountResource Deleted
*/
type DeleteOnBackgroundResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK struct {
	Payload *models.TaskResource
}

func (o *DeleteOnBackgroundResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK) Error() string {
	return fmt.Sprintf("[DELETE /api/accounts/{id}][%d] deleteOnBackgroundResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK  %+v", 200, o.Payload)
}

func (o *DeleteOnBackgroundResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

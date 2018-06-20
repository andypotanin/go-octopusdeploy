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

// LoadResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceReader is a Reader for the LoadResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResource structure.
type LoadResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoadResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLoadResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLoadResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK creates a LoadResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK with default headers values
func NewLoadResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK() *LoadResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK {
	return &LoadResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK{}
}

/*LoadResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK handles this case with default header values.

Load AccountResource by ID
*/
type LoadResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK struct {
	Payload *models.AccountResource
}

func (o *LoadResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/accounts/{id}][%d] loadResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK  %+v", 200, o.Payload)
}

func (o *LoadResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

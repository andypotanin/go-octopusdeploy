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

// IndexResponseDescriptorTenantsTenantTenantResourceReader is a Reader for the IndexResponseDescriptorTenantsTenantTenantResource structure.
type IndexResponseDescriptorTenantsTenantTenantResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexResponseDescriptorTenantsTenantTenantResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIndexResponseDescriptorTenantsTenantTenantResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIndexResponseDescriptorTenantsTenantTenantResourceOK creates a IndexResponseDescriptorTenantsTenantTenantResourceOK with default headers values
func NewIndexResponseDescriptorTenantsTenantTenantResourceOK() *IndexResponseDescriptorTenantsTenantTenantResourceOK {
	return &IndexResponseDescriptorTenantsTenantTenantResourceOK{}
}

/*IndexResponseDescriptorTenantsTenantTenantResourceOK handles this case with default header values.

Load TenantResource collection
*/
type IndexResponseDescriptorTenantsTenantTenantResourceOK struct {
	Payload *models.ResourceCollectionTenantResource
}

func (o *IndexResponseDescriptorTenantsTenantTenantResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/tenants][%d] indexResponseDescriptorTenantsTenantTenantResourceOK  %+v", 200, o.Payload)
}

func (o *IndexResponseDescriptorTenantsTenantTenantResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionTenantResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package certificate_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// IndexResponseDescriptorConfigurationCertificateConfigurationCertificateConfigurationResourceReader is a Reader for the IndexResponseDescriptorConfigurationCertificateConfigurationCertificateConfigurationResource structure.
type IndexResponseDescriptorConfigurationCertificateConfigurationCertificateConfigurationResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexResponseDescriptorConfigurationCertificateConfigurationCertificateConfigurationResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIndexResponseDescriptorConfigurationCertificateConfigurationCertificateConfigurationResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIndexResponseDescriptorConfigurationCertificateConfigurationCertificateConfigurationResourceOK creates a IndexResponseDescriptorConfigurationCertificateConfigurationCertificateConfigurationResourceOK with default headers values
func NewIndexResponseDescriptorConfigurationCertificateConfigurationCertificateConfigurationResourceOK() *IndexResponseDescriptorConfigurationCertificateConfigurationCertificateConfigurationResourceOK {
	return &IndexResponseDescriptorConfigurationCertificateConfigurationCertificateConfigurationResourceOK{}
}

/*IndexResponseDescriptorConfigurationCertificateConfigurationCertificateConfigurationResourceOK handles this case with default header values.

Load CertificateConfigurationResource collection
*/
type IndexResponseDescriptorConfigurationCertificateConfigurationCertificateConfigurationResourceOK struct {
	Payload *models.ResourceCollectionCertificateConfigurationResource
}

func (o *IndexResponseDescriptorConfigurationCertificateConfigurationCertificateConfigurationResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/configuration/certificates][%d] indexResponseDescriptorConfigurationCertificateConfigurationCertificateConfigurationResourceOK  %+v", 200, o.Payload)
}

func (o *IndexResponseDescriptorConfigurationCertificateConfigurationCertificateConfigurationResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionCertificateConfigurationResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

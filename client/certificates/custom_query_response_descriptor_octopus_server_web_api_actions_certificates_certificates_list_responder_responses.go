// Code generated by go-swagger; DO NOT EDIT.

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// CustomQueryResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificatesListResponderReader is a Reader for the CustomQueryResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificatesListResponder structure.
type CustomQueryResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificatesListResponderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificatesListResponderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificatesListResponderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificatesListResponderOK creates a CustomQueryResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificatesListResponderOK with default headers values
func NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificatesListResponderOK() *CustomQueryResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificatesListResponderOK {
	return &CustomQueryResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificatesListResponderOK{}
}

/*CustomQueryResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificatesListResponderOK handles this case with default header values.

ResourceCollection[CertificateResource] resource returned
*/
type CustomQueryResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificatesListResponderOK struct {
	Payload *models.ResourceCollectionCertificateResource
}

func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificatesListResponderOK) Error() string {
	return fmt.Sprintf("[GET /api/certificates][%d] customQueryResponseDescriptorOctopusServerWebApiActionsCertificatesCertificatesListResponderOK  %+v", 200, o.Payload)
}

func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsCertificatesCertificatesListResponderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionCertificateResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// IndexResponseDescriptorProjectsReleaseReleaseResourceReader is a Reader for the IndexResponseDescriptorProjectsReleaseReleaseResource structure.
type IndexResponseDescriptorProjectsReleaseReleaseResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexResponseDescriptorProjectsReleaseReleaseResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIndexResponseDescriptorProjectsReleaseReleaseResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIndexResponseDescriptorProjectsReleaseReleaseResourceOK creates a IndexResponseDescriptorProjectsReleaseReleaseResourceOK with default headers values
func NewIndexResponseDescriptorProjectsReleaseReleaseResourceOK() *IndexResponseDescriptorProjectsReleaseReleaseResourceOK {
	return &IndexResponseDescriptorProjectsReleaseReleaseResourceOK{}
}

/*IndexResponseDescriptorProjectsReleaseReleaseResourceOK handles this case with default header values.

Load ReleaseResource collection
*/
type IndexResponseDescriptorProjectsReleaseReleaseResourceOK struct {
	Payload *models.ResourceCollectionReleaseResource
}

func (o *IndexResponseDescriptorProjectsReleaseReleaseResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/releases][%d] indexResponseDescriptorProjectsReleaseReleaseResourceOK  %+v", 200, o.Payload)
}

func (o *IndexResponseDescriptorProjectsReleaseReleaseResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionReleaseResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

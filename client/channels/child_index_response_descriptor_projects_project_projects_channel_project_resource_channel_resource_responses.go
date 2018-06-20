// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceReader is a Reader for the ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResource structure.
type ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceOK creates a ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceOK with default headers values
func NewChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceOK() *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceOK {
	return &ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceOK{}
}

/*ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceOK handles this case with default header values.

Load ProjectResource collection
*/
type ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceOK struct {
	Payload *models.ResourceCollectionProjectResource
}

func (o *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/projects/{id}/channels][%d] childIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceOK  %+v", 200, o.Payload)
}

func (o *ChildIndexResponseDescriptorProjectsProjectProjectsChannelProjectResourceChannelResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionProjectResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

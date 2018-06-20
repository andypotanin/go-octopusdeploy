// Code generated by go-swagger; DO NOT EDIT.

package artifacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// LoadResponseDescriptorServerTasksArtifactArtifactResourceReader is a Reader for the LoadResponseDescriptorServerTasksArtifactArtifactResource structure.
type LoadResponseDescriptorServerTasksArtifactArtifactResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoadResponseDescriptorServerTasksArtifactArtifactResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLoadResponseDescriptorServerTasksArtifactArtifactResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLoadResponseDescriptorServerTasksArtifactArtifactResourceOK creates a LoadResponseDescriptorServerTasksArtifactArtifactResourceOK with default headers values
func NewLoadResponseDescriptorServerTasksArtifactArtifactResourceOK() *LoadResponseDescriptorServerTasksArtifactArtifactResourceOK {
	return &LoadResponseDescriptorServerTasksArtifactArtifactResourceOK{}
}

/*LoadResponseDescriptorServerTasksArtifactArtifactResourceOK handles this case with default header values.

Load ArtifactResource by ID
*/
type LoadResponseDescriptorServerTasksArtifactArtifactResourceOK struct {
	Payload *models.ArtifactResource
}

func (o *LoadResponseDescriptorServerTasksArtifactArtifactResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/artifacts/{id}][%d] loadResponseDescriptorServerTasksArtifactArtifactResourceOK  %+v", 200, o.Payload)
}

func (o *LoadResponseDescriptorServerTasksArtifactArtifactResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ArtifactResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

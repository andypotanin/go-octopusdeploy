// Code generated by go-swagger; DO NOT EDIT.

package dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsDynamicDashboardViewActionReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsDynamicDashboardViewAction structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsDynamicDashboardViewActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDynamicDashboardViewActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDynamicDashboardViewActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDynamicDashboardViewActionOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsDynamicDashboardViewActionOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDynamicDashboardViewActionOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsDynamicDashboardViewActionOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDynamicDashboardViewActionOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsDynamicDashboardViewActionOK handles this case with default header values.

DashboardResource resource returned
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsDynamicDashboardViewActionOK struct {
	Payload *models.DashboardResource
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDynamicDashboardViewActionOK) Error() string {
	return fmt.Sprintf("[GET /api/dashboard/dynamic][%d] customActionResponseDescriptorOctopusServerWebApiActionsDynamicDashboardViewActionOK  %+v", 200, o.Payload)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDynamicDashboardViewActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

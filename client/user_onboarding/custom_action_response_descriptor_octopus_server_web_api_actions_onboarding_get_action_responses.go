// Code generated by go-swagger; DO NOT EDIT.

package user_onboarding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsOnboardingGetActionReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsOnboardingGetAction structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsOnboardingGetActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsOnboardingGetActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsOnboardingGetActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsOnboardingGetActionOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsOnboardingGetActionOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsOnboardingGetActionOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsOnboardingGetActionOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsOnboardingGetActionOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsOnboardingGetActionOK handles this case with default header values.

OnboardingResource resource returned
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsOnboardingGetActionOK struct {
	Payload *models.OnboardingResource
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsOnboardingGetActionOK) Error() string {
	return fmt.Sprintf("[GET /api/useronboarding][%d] customActionResponseDescriptorOctopusServerWebApiActionsOnboardingGetActionOK  %+v", 200, o.Payload)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsOnboardingGetActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OnboardingResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package octopus_server_nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponder structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 418:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderIMATeapot()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderOK handles this case with default header values.

OK
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderOK struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderOK) Error() string {
	return fmt.Sprintf("[GET /api/octopusservernodes/ping][%d] customActionResponseDescriptorOctopusServerWebApiActionsLoadBalancerPingResponderOK ", 200)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderIMATeapot creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderIMATeapot with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderIMATeapot() *CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderIMATeapot {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderIMATeapot{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderIMATeapot handles this case with default header values.

ImATeapot
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderIMATeapot struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderIMATeapot) Error() string {
	return fmt.Sprintf("[GET /api/octopusservernodes/ping][%d] customActionResponseDescriptorOctopusServerWebApiActionsLoadBalancerPingResponderIMATeapot ", 418)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLoadBalancerPingResponderIMATeapot) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

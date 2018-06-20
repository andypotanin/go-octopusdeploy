// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadAction structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionOK(writer io.Writer) *CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionOK{
		Payload: writer,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionOK handles this case with default header values.

OK
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionOK struct {
	Payload io.Writer
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionOK) Error() string {
	return fmt.Sprintf("[GET /api/accounts/{id}/pk][%d] customActionResponseDescriptorOctopusServerWebApiActionsAccountPublicKeyDownloadActionOK  %+v", 200, o.Payload)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionBadRequest creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionBadRequest with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionBadRequest() *CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionBadRequest {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionBadRequest{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionBadRequest handles this case with default header values.

No id parameter was provided.
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionBadRequest struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/accounts/{id}/pk][%d] customActionResponseDescriptorOctopusServerWebApiActionsAccountPublicKeyDownloadActionBadRequest ", 400)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsAccountPublicKeyDownloadActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

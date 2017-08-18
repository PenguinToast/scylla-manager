// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetExceptionsReader is a Reader for the GetExceptions structure.
type GetExceptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExceptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetExceptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetExceptionsOK creates a GetExceptionsOK with default headers values
func NewGetExceptionsOK() *GetExceptionsOK {
	return &GetExceptionsOK{}
}

/*GetExceptionsOK handles this case with default header values.

GetExceptionsOK get exceptions o k
*/
type GetExceptionsOK struct {
	Payload int32
}

func (o *GetExceptionsOK) Error() string {
	return fmt.Sprintf("[GET /storage_service/metrics/exceptions][%d] getExceptionsOK  %+v", 200, o.Payload)
}

func (o *GetExceptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

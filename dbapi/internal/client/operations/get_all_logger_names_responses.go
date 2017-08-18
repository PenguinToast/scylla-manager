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

// GetAllLoggerNamesReader is a Reader for the GetAllLoggerNames structure.
type GetAllLoggerNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllLoggerNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllLoggerNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllLoggerNamesOK creates a GetAllLoggerNamesOK with default headers values
func NewGetAllLoggerNamesOK() *GetAllLoggerNamesOK {
	return &GetAllLoggerNamesOK{}
}

/*GetAllLoggerNamesOK handles this case with default header values.

GetAllLoggerNamesOK get all logger names o k
*/
type GetAllLoggerNamesOK struct {
	Payload []string
}

func (o *GetAllLoggerNamesOK) Error() string {
	return fmt.Sprintf("[GET /system/logger][%d] getAllLoggerNamesOK  %+v", 200, o.Payload)
}

func (o *GetAllLoggerNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetAllMinRowSizeReader is a Reader for the GetAllMinRowSize structure.
type GetAllMinRowSizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllMinRowSizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllMinRowSizeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllMinRowSizeOK creates a GetAllMinRowSizeOK with default headers values
func NewGetAllMinRowSizeOK() *GetAllMinRowSizeOK {
	return &GetAllMinRowSizeOK{}
}

/*GetAllMinRowSizeOK handles this case with default header values.

GetAllMinRowSizeOK get all min row size o k
*/
type GetAllMinRowSizeOK struct {
	Payload int64
}

func (o *GetAllMinRowSizeOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/min_row_size][%d] getAllMinRowSizeOK  %+v", 200, o.Payload)
}

func (o *GetAllMinRowSizeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

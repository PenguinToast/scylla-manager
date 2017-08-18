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

// GetBytesCompactedReader is a Reader for the GetBytesCompacted structure.
type GetBytesCompactedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBytesCompactedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBytesCompactedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBytesCompactedOK creates a GetBytesCompactedOK with default headers values
func NewGetBytesCompactedOK() *GetBytesCompactedOK {
	return &GetBytesCompactedOK{}
}

/*GetBytesCompactedOK handles this case with default header values.

GetBytesCompactedOK get bytes compacted o k
*/
type GetBytesCompactedOK struct {
	Payload int32
}

func (o *GetBytesCompactedOK) Error() string {
	return fmt.Sprintf("[GET /compaction_manager/metrics/bytes_compacted][%d] getBytesCompactedOK  %+v", 200, o.Payload)
}

func (o *GetBytesCompactedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

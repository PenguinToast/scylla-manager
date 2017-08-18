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

// SetMinimumCompactionThresholdReader is a Reader for the SetMinimumCompactionThreshold structure.
type SetMinimumCompactionThresholdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetMinimumCompactionThresholdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetMinimumCompactionThresholdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetMinimumCompactionThresholdOK creates a SetMinimumCompactionThresholdOK with default headers values
func NewSetMinimumCompactionThresholdOK() *SetMinimumCompactionThresholdOK {
	return &SetMinimumCompactionThresholdOK{}
}

/*SetMinimumCompactionThresholdOK handles this case with default header values.

SetMinimumCompactionThresholdOK set minimum compaction threshold o k
*/
type SetMinimumCompactionThresholdOK struct {
	Payload string
}

func (o *SetMinimumCompactionThresholdOK) Error() string {
	return fmt.Sprintf("[POST /column_family/minimum_compaction/{name}][%d] setMinimumCompactionThresholdOK  %+v", 200, o.Payload)
}

func (o *SetMinimumCompactionThresholdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

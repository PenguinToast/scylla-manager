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

// GetSstableCountPerLevelReader is a Reader for the GetSstableCountPerLevel structure.
type GetSstableCountPerLevelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSstableCountPerLevelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSstableCountPerLevelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSstableCountPerLevelOK creates a GetSstableCountPerLevelOK with default headers values
func NewGetSstableCountPerLevelOK() *GetSstableCountPerLevelOK {
	return &GetSstableCountPerLevelOK{}
}

/*GetSstableCountPerLevelOK handles this case with default header values.

GetSstableCountPerLevelOK get sstable count per level o k
*/
type GetSstableCountPerLevelOK struct {
	Payload []int32
}

func (o *GetSstableCountPerLevelOK) Error() string {
	return fmt.Sprintf("[GET /column_family/sstables/per_level/{name}][%d] getSstableCountPerLevelOK  %+v", 200, o.Payload)
}

func (o *GetSstableCountPerLevelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

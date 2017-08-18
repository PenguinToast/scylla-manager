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

// GetTombstoneScannedHistogramReader is a Reader for the GetTombstoneScannedHistogram structure.
type GetTombstoneScannedHistogramReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTombstoneScannedHistogramReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTombstoneScannedHistogramOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTombstoneScannedHistogramOK creates a GetTombstoneScannedHistogramOK with default headers values
func NewGetTombstoneScannedHistogramOK() *GetTombstoneScannedHistogramOK {
	return &GetTombstoneScannedHistogramOK{}
}

/*GetTombstoneScannedHistogramOK handles this case with default header values.

GetTombstoneScannedHistogramOK get tombstone scanned histogram o k
*/
type GetTombstoneScannedHistogramOK struct {
	Payload []int32
}

func (o *GetTombstoneScannedHistogramOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/tombstone_scanned_histogram/{name}][%d] getTombstoneScannedHistogramOK  %+v", 200, o.Payload)
}

func (o *GetTombstoneScannedHistogramOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

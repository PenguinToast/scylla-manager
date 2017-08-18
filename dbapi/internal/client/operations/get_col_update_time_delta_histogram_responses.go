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

// GetColUpdateTimeDeltaHistogramReader is a Reader for the GetColUpdateTimeDeltaHistogram structure.
type GetColUpdateTimeDeltaHistogramReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetColUpdateTimeDeltaHistogramReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetColUpdateTimeDeltaHistogramOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetColUpdateTimeDeltaHistogramOK creates a GetColUpdateTimeDeltaHistogramOK with default headers values
func NewGetColUpdateTimeDeltaHistogramOK() *GetColUpdateTimeDeltaHistogramOK {
	return &GetColUpdateTimeDeltaHistogramOK{}
}

/*GetColUpdateTimeDeltaHistogramOK handles this case with default header values.

GetColUpdateTimeDeltaHistogramOK get col update time delta histogram o k
*/
type GetColUpdateTimeDeltaHistogramOK struct {
	Payload []int32
}

func (o *GetColUpdateTimeDeltaHistogramOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/col_update_time_delta_histogram/{name}][%d] getColUpdateTimeDeltaHistogramOK  %+v", 200, o.Payload)
}

func (o *GetColUpdateTimeDeltaHistogramOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

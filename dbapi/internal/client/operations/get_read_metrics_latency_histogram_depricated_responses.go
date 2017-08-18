// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/scylladb/mermaid/dbapi/internal/models"
)

// GetReadMetricsLatencyHistogramDepricatedReader is a Reader for the GetReadMetricsLatencyHistogramDepricated structure.
type GetReadMetricsLatencyHistogramDepricatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReadMetricsLatencyHistogramDepricatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReadMetricsLatencyHistogramDepricatedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReadMetricsLatencyHistogramDepricatedOK creates a GetReadMetricsLatencyHistogramDepricatedOK with default headers values
func NewGetReadMetricsLatencyHistogramDepricatedOK() *GetReadMetricsLatencyHistogramDepricatedOK {
	return &GetReadMetricsLatencyHistogramDepricatedOK{}
}

/*GetReadMetricsLatencyHistogramDepricatedOK handles this case with default header values.

GetReadMetricsLatencyHistogramDepricatedOK get read metrics latency histogram depricated o k
*/
type GetReadMetricsLatencyHistogramDepricatedOK struct {
	Payload *models.Histogram
}

func (o *GetReadMetricsLatencyHistogramDepricatedOK) Error() string {
	return fmt.Sprintf("[GET /storage_proxy/metrics/read/histogram][%d] getReadMetricsLatencyHistogramDepricatedOK  %+v", 200, o.Payload)
}

func (o *GetReadMetricsLatencyHistogramDepricatedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Histogram)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

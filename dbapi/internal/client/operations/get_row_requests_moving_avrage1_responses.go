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

// GetRowRequestsMovingAvrage1Reader is a Reader for the GetRowRequestsMovingAvrage1 structure.
type GetRowRequestsMovingAvrage1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRowRequestsMovingAvrage1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRowRequestsMovingAvrage1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRowRequestsMovingAvrage1OK creates a GetRowRequestsMovingAvrage1OK with default headers values
func NewGetRowRequestsMovingAvrage1OK() *GetRowRequestsMovingAvrage1OK {
	return &GetRowRequestsMovingAvrage1OK{}
}

/*GetRowRequestsMovingAvrage1OK handles this case with default header values.

GetRowRequestsMovingAvrage1OK get row requests moving avrage1 o k
*/
type GetRowRequestsMovingAvrage1OK struct {
	Payload *models.RateMovingAverage
}

func (o *GetRowRequestsMovingAvrage1OK) Error() string {
	return fmt.Sprintf("[GET /cache_service/metrics/row/requests_moving_avrage][%d] getRowRequestsMovingAvrage1OK  %+v", 200, o.Payload)
}

func (o *GetRowRequestsMovingAvrage1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RateMovingAverage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

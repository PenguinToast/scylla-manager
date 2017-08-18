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

// GetHintedHandoffEnabledByDcReader is a Reader for the GetHintedHandoffEnabledByDc structure.
type GetHintedHandoffEnabledByDcReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHintedHandoffEnabledByDcReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetHintedHandoffEnabledByDcOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetHintedHandoffEnabledByDcOK creates a GetHintedHandoffEnabledByDcOK with default headers values
func NewGetHintedHandoffEnabledByDcOK() *GetHintedHandoffEnabledByDcOK {
	return &GetHintedHandoffEnabledByDcOK{}
}

/*GetHintedHandoffEnabledByDcOK handles this case with default header values.

GetHintedHandoffEnabledByDcOK get hinted handoff enabled by dc o k
*/
type GetHintedHandoffEnabledByDcOK struct {
	Payload []*models.MapperList
}

func (o *GetHintedHandoffEnabledByDcOK) Error() string {
	return fmt.Sprintf("[GET /storage_proxy/hinted_handoff_enabled_by_dc][%d] getHintedHandoffEnabledByDcOK  %+v", 200, o.Payload)
}

func (o *GetHintedHandoffEnabledByDcOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

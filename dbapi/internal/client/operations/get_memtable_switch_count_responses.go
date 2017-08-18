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

// GetMemtableSwitchCountReader is a Reader for the GetMemtableSwitchCount structure.
type GetMemtableSwitchCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMemtableSwitchCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMemtableSwitchCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMemtableSwitchCountOK creates a GetMemtableSwitchCountOK with default headers values
func NewGetMemtableSwitchCountOK() *GetMemtableSwitchCountOK {
	return &GetMemtableSwitchCountOK{}
}

/*GetMemtableSwitchCountOK handles this case with default header values.

GetMemtableSwitchCountOK get memtable switch count o k
*/
type GetMemtableSwitchCountOK struct {
	Payload int32
}

func (o *GetMemtableSwitchCountOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/memtable_switch_count/{name}][%d] getMemtableSwitchCountOK  %+v", 200, o.Payload)
}

func (o *GetMemtableSwitchCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

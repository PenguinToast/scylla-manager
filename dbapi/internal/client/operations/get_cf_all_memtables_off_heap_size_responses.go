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

// GetCfAllMemtablesOffHeapSizeReader is a Reader for the GetCfAllMemtablesOffHeapSize structure.
type GetCfAllMemtablesOffHeapSizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCfAllMemtablesOffHeapSizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCfAllMemtablesOffHeapSizeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCfAllMemtablesOffHeapSizeOK creates a GetCfAllMemtablesOffHeapSizeOK with default headers values
func NewGetCfAllMemtablesOffHeapSizeOK() *GetCfAllMemtablesOffHeapSizeOK {
	return &GetCfAllMemtablesOffHeapSizeOK{}
}

/*GetCfAllMemtablesOffHeapSizeOK handles this case with default header values.

GetCfAllMemtablesOffHeapSizeOK get cf all memtables off heap size o k
*/
type GetCfAllMemtablesOffHeapSizeOK struct {
	Payload int64
}

func (o *GetCfAllMemtablesOffHeapSizeOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/all_memtables_off_heap_size/{name}][%d] getCfAllMemtablesOffHeapSizeOK  %+v", 200, o.Payload)
}

func (o *GetCfAllMemtablesOffHeapSizeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

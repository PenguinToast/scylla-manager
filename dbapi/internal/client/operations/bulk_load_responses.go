// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// BulkLoadReader is a Reader for the BulkLoad structure.
type BulkLoadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkLoadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBulkLoadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBulkLoadOK creates a BulkLoadOK with default headers values
func NewBulkLoadOK() *BulkLoadOK {
	return &BulkLoadOK{}
}

/*BulkLoadOK handles this case with default header values.

BulkLoadOK bulk load o k
*/
type BulkLoadOK struct {
}

func (o *BulkLoadOK) Error() string {
	return fmt.Sprintf("[POST /storage_service/bulk_load/{path}][%d] bulkLoadOK ", 200)
}

func (o *BulkLoadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

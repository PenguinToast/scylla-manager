// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DecommissionReader is a Reader for the Decommission structure.
type DecommissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DecommissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDecommissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDecommissionOK creates a DecommissionOK with default headers values
func NewDecommissionOK() *DecommissionOK {
	return &DecommissionOK{}
}

/*DecommissionOK handles this case with default header values.

DecommissionOK decommission o k
*/
type DecommissionOK struct {
}

func (o *DecommissionOK) Error() string {
	return fmt.Sprintf("[POST /storage_service/decommission][%d] decommissionOK ", 200)
}

func (o *DecommissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

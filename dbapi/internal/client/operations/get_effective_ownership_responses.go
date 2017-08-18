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

// GetEffectiveOwnershipReader is a Reader for the GetEffectiveOwnership structure.
type GetEffectiveOwnershipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEffectiveOwnershipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEffectiveOwnershipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEffectiveOwnershipOK creates a GetEffectiveOwnershipOK with default headers values
func NewGetEffectiveOwnershipOK() *GetEffectiveOwnershipOK {
	return &GetEffectiveOwnershipOK{}
}

/*GetEffectiveOwnershipOK handles this case with default header values.

GetEffectiveOwnershipOK get effective ownership o k
*/
type GetEffectiveOwnershipOK struct {
	Payload []*models.Mapper
}

func (o *GetEffectiveOwnershipOK) Error() string {
	return fmt.Sprintf("[GET /storage_service/ownership/{keyspace}][%d] getEffectiveOwnershipOK  %+v", 200, o.Payload)
}

func (o *GetEffectiveOwnershipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

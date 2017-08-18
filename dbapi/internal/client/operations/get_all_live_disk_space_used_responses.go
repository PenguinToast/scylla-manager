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

// GetAllLiveDiskSpaceUsedReader is a Reader for the GetAllLiveDiskSpaceUsed structure.
type GetAllLiveDiskSpaceUsedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllLiveDiskSpaceUsedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllLiveDiskSpaceUsedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllLiveDiskSpaceUsedOK creates a GetAllLiveDiskSpaceUsedOK with default headers values
func NewGetAllLiveDiskSpaceUsedOK() *GetAllLiveDiskSpaceUsedOK {
	return &GetAllLiveDiskSpaceUsedOK{}
}

/*GetAllLiveDiskSpaceUsedOK handles this case with default header values.

GetAllLiveDiskSpaceUsedOK get all live disk space used o k
*/
type GetAllLiveDiskSpaceUsedOK struct {
	Payload int32
}

func (o *GetAllLiveDiskSpaceUsedOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/live_disk_space_used][%d] getAllLiveDiskSpaceUsedOK  %+v", 200, o.Payload)
}

func (o *GetAllLiveDiskSpaceUsedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

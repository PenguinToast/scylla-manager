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

// GetAllTrueSnapshotsSizeReader is a Reader for the GetAllTrueSnapshotsSize structure.
type GetAllTrueSnapshotsSizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTrueSnapshotsSizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllTrueSnapshotsSizeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllTrueSnapshotsSizeOK creates a GetAllTrueSnapshotsSizeOK with default headers values
func NewGetAllTrueSnapshotsSizeOK() *GetAllTrueSnapshotsSizeOK {
	return &GetAllTrueSnapshotsSizeOK{}
}

/*GetAllTrueSnapshotsSizeOK handles this case with default header values.

GetAllTrueSnapshotsSizeOK get all true snapshots size o k
*/
type GetAllTrueSnapshotsSizeOK struct {
	Payload int64
}

func (o *GetAllTrueSnapshotsSizeOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/true_snapshots_size][%d] getAllTrueSnapshotsSizeOK  %+v", 200, o.Payload)
}

func (o *GetAllTrueSnapshotsSizeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

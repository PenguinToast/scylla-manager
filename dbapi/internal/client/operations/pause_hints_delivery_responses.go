// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PauseHintsDeliveryReader is a Reader for the PauseHintsDelivery structure.
type PauseHintsDeliveryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PauseHintsDeliveryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPauseHintsDeliveryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPauseHintsDeliveryOK creates a PauseHintsDeliveryOK with default headers values
func NewPauseHintsDeliveryOK() *PauseHintsDeliveryOK {
	return &PauseHintsDeliveryOK{}
}

/*PauseHintsDeliveryOK handles this case with default header values.

PauseHintsDeliveryOK pause hints delivery o k
*/
type PauseHintsDeliveryOK struct {
}

func (o *PauseHintsDeliveryOK) Error() string {
	return fmt.Sprintf("[POST /hinted_handoff/pause][%d] pauseHintsDeliveryOK ", 200)
}

func (o *PauseHintsDeliveryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPauseHintsDeliveryParams creates a new PauseHintsDeliveryParams object
// with the default values initialized.
func NewPauseHintsDeliveryParams() *PauseHintsDeliveryParams {
	var ()
	return &PauseHintsDeliveryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPauseHintsDeliveryParamsWithTimeout creates a new PauseHintsDeliveryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPauseHintsDeliveryParamsWithTimeout(timeout time.Duration) *PauseHintsDeliveryParams {
	var ()
	return &PauseHintsDeliveryParams{

		timeout: timeout,
	}
}

// NewPauseHintsDeliveryParamsWithContext creates a new PauseHintsDeliveryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPauseHintsDeliveryParamsWithContext(ctx context.Context) *PauseHintsDeliveryParams {
	var ()
	return &PauseHintsDeliveryParams{

		Context: ctx,
	}
}

// NewPauseHintsDeliveryParamsWithHTTPClient creates a new PauseHintsDeliveryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPauseHintsDeliveryParamsWithHTTPClient(client *http.Client) *PauseHintsDeliveryParams {
	var ()
	return &PauseHintsDeliveryParams{
		HTTPClient: client,
	}
}

/*PauseHintsDeliveryParams contains all the parameters to send to the API endpoint
for the pause hints delivery operation typically these are written to a http.Request
*/
type PauseHintsDeliveryParams struct {

	/*Pause
	  pause status

	*/
	Pause bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pause hints delivery params
func (o *PauseHintsDeliveryParams) WithTimeout(timeout time.Duration) *PauseHintsDeliveryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pause hints delivery params
func (o *PauseHintsDeliveryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pause hints delivery params
func (o *PauseHintsDeliveryParams) WithContext(ctx context.Context) *PauseHintsDeliveryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pause hints delivery params
func (o *PauseHintsDeliveryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pause hints delivery params
func (o *PauseHintsDeliveryParams) WithHTTPClient(client *http.Client) *PauseHintsDeliveryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pause hints delivery params
func (o *PauseHintsDeliveryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPause adds the pause to the pause hints delivery params
func (o *PauseHintsDeliveryParams) WithPause(pause bool) *PauseHintsDeliveryParams {
	o.SetPause(pause)
	return o
}

// SetPause adds the pause to the pause hints delivery params
func (o *PauseHintsDeliveryParams) SetPause(pause bool) {
	o.Pause = pause
}

// WriteToRequest writes these params to a swagger request
func (o *PauseHintsDeliveryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param pause
	qrPause := o.Pause
	qPause := swag.FormatBool(qrPause)
	if qPause != "" {
		if err := r.SetQueryParam("pause", qPause); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

	strfmt "github.com/go-openapi/strfmt"
)

// NewTrueSnapshotsSizeParams creates a new TrueSnapshotsSizeParams object
// with the default values initialized.
func NewTrueSnapshotsSizeParams() *TrueSnapshotsSizeParams {

	return &TrueSnapshotsSizeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTrueSnapshotsSizeParamsWithTimeout creates a new TrueSnapshotsSizeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTrueSnapshotsSizeParamsWithTimeout(timeout time.Duration) *TrueSnapshotsSizeParams {

	return &TrueSnapshotsSizeParams{

		timeout: timeout,
	}
}

// NewTrueSnapshotsSizeParamsWithContext creates a new TrueSnapshotsSizeParams object
// with the default values initialized, and the ability to set a context for a request
func NewTrueSnapshotsSizeParamsWithContext(ctx context.Context) *TrueSnapshotsSizeParams {

	return &TrueSnapshotsSizeParams{

		Context: ctx,
	}
}

// NewTrueSnapshotsSizeParamsWithHTTPClient creates a new TrueSnapshotsSizeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTrueSnapshotsSizeParamsWithHTTPClient(client *http.Client) *TrueSnapshotsSizeParams {

	return &TrueSnapshotsSizeParams{
		HTTPClient: client,
	}
}

/*TrueSnapshotsSizeParams contains all the parameters to send to the API endpoint
for the true snapshots size operation typically these are written to a http.Request
*/
type TrueSnapshotsSizeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the true snapshots size params
func (o *TrueSnapshotsSizeParams) WithTimeout(timeout time.Duration) *TrueSnapshotsSizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the true snapshots size params
func (o *TrueSnapshotsSizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the true snapshots size params
func (o *TrueSnapshotsSizeParams) WithContext(ctx context.Context) *TrueSnapshotsSizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the true snapshots size params
func (o *TrueSnapshotsSizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the true snapshots size params
func (o *TrueSnapshotsSizeParams) WithHTTPClient(client *http.Client) *TrueSnapshotsSizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the true snapshots size params
func (o *TrueSnapshotsSizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TrueSnapshotsSizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

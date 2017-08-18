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

// NewInvalidateCounterCache1Params creates a new InvalidateCounterCache1Params object
// with the default values initialized.
func NewInvalidateCounterCache1Params() *InvalidateCounterCache1Params {

	return &InvalidateCounterCache1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewInvalidateCounterCache1ParamsWithTimeout creates a new InvalidateCounterCache1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewInvalidateCounterCache1ParamsWithTimeout(timeout time.Duration) *InvalidateCounterCache1Params {

	return &InvalidateCounterCache1Params{

		timeout: timeout,
	}
}

// NewInvalidateCounterCache1ParamsWithContext creates a new InvalidateCounterCache1Params object
// with the default values initialized, and the ability to set a context for a request
func NewInvalidateCounterCache1ParamsWithContext(ctx context.Context) *InvalidateCounterCache1Params {

	return &InvalidateCounterCache1Params{

		Context: ctx,
	}
}

// NewInvalidateCounterCache1ParamsWithHTTPClient creates a new InvalidateCounterCache1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInvalidateCounterCache1ParamsWithHTTPClient(client *http.Client) *InvalidateCounterCache1Params {

	return &InvalidateCounterCache1Params{
		HTTPClient: client,
	}
}

/*InvalidateCounterCache1Params contains all the parameters to send to the API endpoint
for the invalidate counter cache1 operation typically these are written to a http.Request
*/
type InvalidateCounterCache1Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the invalidate counter cache1 params
func (o *InvalidateCounterCache1Params) WithTimeout(timeout time.Duration) *InvalidateCounterCache1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invalidate counter cache1 params
func (o *InvalidateCounterCache1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invalidate counter cache1 params
func (o *InvalidateCounterCache1Params) WithContext(ctx context.Context) *InvalidateCounterCache1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invalidate counter cache1 params
func (o *InvalidateCounterCache1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invalidate counter cache1 params
func (o *InvalidateCounterCache1Params) WithHTTPClient(client *http.Client) *InvalidateCounterCache1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invalidate counter cache1 params
func (o *InvalidateCounterCache1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *InvalidateCounterCache1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

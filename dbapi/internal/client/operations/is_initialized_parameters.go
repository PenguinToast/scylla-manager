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

// NewIsInitializedParams creates a new IsInitializedParams object
// with the default values initialized.
func NewIsInitializedParams() *IsInitializedParams {

	return &IsInitializedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIsInitializedParamsWithTimeout creates a new IsInitializedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIsInitializedParamsWithTimeout(timeout time.Duration) *IsInitializedParams {

	return &IsInitializedParams{

		timeout: timeout,
	}
}

// NewIsInitializedParamsWithContext creates a new IsInitializedParams object
// with the default values initialized, and the ability to set a context for a request
func NewIsInitializedParamsWithContext(ctx context.Context) *IsInitializedParams {

	return &IsInitializedParams{

		Context: ctx,
	}
}

// NewIsInitializedParamsWithHTTPClient creates a new IsInitializedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIsInitializedParamsWithHTTPClient(client *http.Client) *IsInitializedParams {

	return &IsInitializedParams{
		HTTPClient: client,
	}
}

/*IsInitializedParams contains all the parameters to send to the API endpoint
for the is initialized operation typically these are written to a http.Request
*/
type IsInitializedParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the is initialized params
func (o *IsInitializedParams) WithTimeout(timeout time.Duration) *IsInitializedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the is initialized params
func (o *IsInitializedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the is initialized params
func (o *IsInitializedParams) WithContext(ctx context.Context) *IsInitializedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the is initialized params
func (o *IsInitializedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the is initialized params
func (o *IsInitializedParams) WithHTTPClient(client *http.Client) *IsInitializedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the is initialized params
func (o *IsInitializedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IsInitializedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

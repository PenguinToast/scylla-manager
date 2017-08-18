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

// NewGetKeyCapacity1Params creates a new GetKeyCapacity1Params object
// with the default values initialized.
func NewGetKeyCapacity1Params() *GetKeyCapacity1Params {

	return &GetKeyCapacity1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKeyCapacity1ParamsWithTimeout creates a new GetKeyCapacity1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKeyCapacity1ParamsWithTimeout(timeout time.Duration) *GetKeyCapacity1Params {

	return &GetKeyCapacity1Params{

		timeout: timeout,
	}
}

// NewGetKeyCapacity1ParamsWithContext creates a new GetKeyCapacity1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetKeyCapacity1ParamsWithContext(ctx context.Context) *GetKeyCapacity1Params {

	return &GetKeyCapacity1Params{

		Context: ctx,
	}
}

// NewGetKeyCapacity1ParamsWithHTTPClient creates a new GetKeyCapacity1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKeyCapacity1ParamsWithHTTPClient(client *http.Client) *GetKeyCapacity1Params {

	return &GetKeyCapacity1Params{
		HTTPClient: client,
	}
}

/*GetKeyCapacity1Params contains all the parameters to send to the API endpoint
for the get key capacity1 operation typically these are written to a http.Request
*/
type GetKeyCapacity1Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get key capacity1 params
func (o *GetKeyCapacity1Params) WithTimeout(timeout time.Duration) *GetKeyCapacity1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get key capacity1 params
func (o *GetKeyCapacity1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get key capacity1 params
func (o *GetKeyCapacity1Params) WithContext(ctx context.Context) *GetKeyCapacity1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get key capacity1 params
func (o *GetKeyCapacity1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get key capacity1 params
func (o *GetKeyCapacity1Params) WithHTTPClient(client *http.Client) *GetKeyCapacity1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get key capacity1 params
func (o *GetKeyCapacity1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetKeyCapacity1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

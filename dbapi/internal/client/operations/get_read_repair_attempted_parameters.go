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

// NewGetReadRepairAttemptedParams creates a new GetReadRepairAttemptedParams object
// with the default values initialized.
func NewGetReadRepairAttemptedParams() *GetReadRepairAttemptedParams {

	return &GetReadRepairAttemptedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReadRepairAttemptedParamsWithTimeout creates a new GetReadRepairAttemptedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReadRepairAttemptedParamsWithTimeout(timeout time.Duration) *GetReadRepairAttemptedParams {

	return &GetReadRepairAttemptedParams{

		timeout: timeout,
	}
}

// NewGetReadRepairAttemptedParamsWithContext creates a new GetReadRepairAttemptedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReadRepairAttemptedParamsWithContext(ctx context.Context) *GetReadRepairAttemptedParams {

	return &GetReadRepairAttemptedParams{

		Context: ctx,
	}
}

// NewGetReadRepairAttemptedParamsWithHTTPClient creates a new GetReadRepairAttemptedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReadRepairAttemptedParamsWithHTTPClient(client *http.Client) *GetReadRepairAttemptedParams {

	return &GetReadRepairAttemptedParams{
		HTTPClient: client,
	}
}

/*GetReadRepairAttemptedParams contains all the parameters to send to the API endpoint
for the get read repair attempted operation typically these are written to a http.Request
*/
type GetReadRepairAttemptedParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get read repair attempted params
func (o *GetReadRepairAttemptedParams) WithTimeout(timeout time.Duration) *GetReadRepairAttemptedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get read repair attempted params
func (o *GetReadRepairAttemptedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get read repair attempted params
func (o *GetReadRepairAttemptedParams) WithContext(ctx context.Context) *GetReadRepairAttemptedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get read repair attempted params
func (o *GetReadRepairAttemptedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get read repair attempted params
func (o *GetReadRepairAttemptedParams) WithHTTPClient(client *http.Client) *GetReadRepairAttemptedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get read repair attempted params
func (o *GetReadRepairAttemptedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetReadRepairAttemptedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

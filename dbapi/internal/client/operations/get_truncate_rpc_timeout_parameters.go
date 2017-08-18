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

// NewGetTruncateRPCTimeoutParams creates a new GetTruncateRPCTimeoutParams object
// with the default values initialized.
func NewGetTruncateRPCTimeoutParams() *GetTruncateRPCTimeoutParams {

	return &GetTruncateRPCTimeoutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTruncateRPCTimeoutParamsWithTimeout creates a new GetTruncateRPCTimeoutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTruncateRPCTimeoutParamsWithTimeout(timeout time.Duration) *GetTruncateRPCTimeoutParams {

	return &GetTruncateRPCTimeoutParams{

		timeout: timeout,
	}
}

// NewGetTruncateRPCTimeoutParamsWithContext creates a new GetTruncateRPCTimeoutParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTruncateRPCTimeoutParamsWithContext(ctx context.Context) *GetTruncateRPCTimeoutParams {

	return &GetTruncateRPCTimeoutParams{

		Context: ctx,
	}
}

// NewGetTruncateRPCTimeoutParamsWithHTTPClient creates a new GetTruncateRPCTimeoutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTruncateRPCTimeoutParamsWithHTTPClient(client *http.Client) *GetTruncateRPCTimeoutParams {

	return &GetTruncateRPCTimeoutParams{
		HTTPClient: client,
	}
}

/*GetTruncateRPCTimeoutParams contains all the parameters to send to the API endpoint
for the get truncate rpc timeout operation typically these are written to a http.Request
*/
type GetTruncateRPCTimeoutParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get truncate rpc timeout params
func (o *GetTruncateRPCTimeoutParams) WithTimeout(timeout time.Duration) *GetTruncateRPCTimeoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get truncate rpc timeout params
func (o *GetTruncateRPCTimeoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get truncate rpc timeout params
func (o *GetTruncateRPCTimeoutParams) WithContext(ctx context.Context) *GetTruncateRPCTimeoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get truncate rpc timeout params
func (o *GetTruncateRPCTimeoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get truncate rpc timeout params
func (o *GetTruncateRPCTimeoutParams) WithHTTPClient(client *http.Client) *GetTruncateRPCTimeoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get truncate rpc timeout params
func (o *GetTruncateRPCTimeoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetTruncateRPCTimeoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

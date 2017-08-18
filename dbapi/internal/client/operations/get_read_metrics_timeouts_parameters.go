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

// NewGetReadMetricsTimeoutsParams creates a new GetReadMetricsTimeoutsParams object
// with the default values initialized.
func NewGetReadMetricsTimeoutsParams() *GetReadMetricsTimeoutsParams {

	return &GetReadMetricsTimeoutsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReadMetricsTimeoutsParamsWithTimeout creates a new GetReadMetricsTimeoutsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReadMetricsTimeoutsParamsWithTimeout(timeout time.Duration) *GetReadMetricsTimeoutsParams {

	return &GetReadMetricsTimeoutsParams{

		timeout: timeout,
	}
}

// NewGetReadMetricsTimeoutsParamsWithContext creates a new GetReadMetricsTimeoutsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReadMetricsTimeoutsParamsWithContext(ctx context.Context) *GetReadMetricsTimeoutsParams {

	return &GetReadMetricsTimeoutsParams{

		Context: ctx,
	}
}

// NewGetReadMetricsTimeoutsParamsWithHTTPClient creates a new GetReadMetricsTimeoutsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReadMetricsTimeoutsParamsWithHTTPClient(client *http.Client) *GetReadMetricsTimeoutsParams {

	return &GetReadMetricsTimeoutsParams{
		HTTPClient: client,
	}
}

/*GetReadMetricsTimeoutsParams contains all the parameters to send to the API endpoint
for the get read metrics timeouts operation typically these are written to a http.Request
*/
type GetReadMetricsTimeoutsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get read metrics timeouts params
func (o *GetReadMetricsTimeoutsParams) WithTimeout(timeout time.Duration) *GetReadMetricsTimeoutsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get read metrics timeouts params
func (o *GetReadMetricsTimeoutsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get read metrics timeouts params
func (o *GetReadMetricsTimeoutsParams) WithContext(ctx context.Context) *GetReadMetricsTimeoutsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get read metrics timeouts params
func (o *GetReadMetricsTimeoutsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get read metrics timeouts params
func (o *GetReadMetricsTimeoutsParams) WithHTTPClient(client *http.Client) *GetReadMetricsTimeoutsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get read metrics timeouts params
func (o *GetReadMetricsTimeoutsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetReadMetricsTimeoutsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

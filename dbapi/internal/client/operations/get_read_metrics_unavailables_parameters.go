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

// NewGetReadMetricsUnavailablesParams creates a new GetReadMetricsUnavailablesParams object
// with the default values initialized.
func NewGetReadMetricsUnavailablesParams() *GetReadMetricsUnavailablesParams {

	return &GetReadMetricsUnavailablesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReadMetricsUnavailablesParamsWithTimeout creates a new GetReadMetricsUnavailablesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReadMetricsUnavailablesParamsWithTimeout(timeout time.Duration) *GetReadMetricsUnavailablesParams {

	return &GetReadMetricsUnavailablesParams{

		timeout: timeout,
	}
}

// NewGetReadMetricsUnavailablesParamsWithContext creates a new GetReadMetricsUnavailablesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReadMetricsUnavailablesParamsWithContext(ctx context.Context) *GetReadMetricsUnavailablesParams {

	return &GetReadMetricsUnavailablesParams{

		Context: ctx,
	}
}

// NewGetReadMetricsUnavailablesParamsWithHTTPClient creates a new GetReadMetricsUnavailablesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReadMetricsUnavailablesParamsWithHTTPClient(client *http.Client) *GetReadMetricsUnavailablesParams {

	return &GetReadMetricsUnavailablesParams{
		HTTPClient: client,
	}
}

/*GetReadMetricsUnavailablesParams contains all the parameters to send to the API endpoint
for the get read metrics unavailables operation typically these are written to a http.Request
*/
type GetReadMetricsUnavailablesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get read metrics unavailables params
func (o *GetReadMetricsUnavailablesParams) WithTimeout(timeout time.Duration) *GetReadMetricsUnavailablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get read metrics unavailables params
func (o *GetReadMetricsUnavailablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get read metrics unavailables params
func (o *GetReadMetricsUnavailablesParams) WithContext(ctx context.Context) *GetReadMetricsUnavailablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get read metrics unavailables params
func (o *GetReadMetricsUnavailablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get read metrics unavailables params
func (o *GetReadMetricsUnavailablesParams) WithHTTPClient(client *http.Client) *GetReadMetricsUnavailablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get read metrics unavailables params
func (o *GetReadMetricsUnavailablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetReadMetricsUnavailablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

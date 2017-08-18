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

// NewGetRangeMetricsLatencyHistogramDepricatedParams creates a new GetRangeMetricsLatencyHistogramDepricatedParams object
// with the default values initialized.
func NewGetRangeMetricsLatencyHistogramDepricatedParams() *GetRangeMetricsLatencyHistogramDepricatedParams {

	return &GetRangeMetricsLatencyHistogramDepricatedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRangeMetricsLatencyHistogramDepricatedParamsWithTimeout creates a new GetRangeMetricsLatencyHistogramDepricatedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRangeMetricsLatencyHistogramDepricatedParamsWithTimeout(timeout time.Duration) *GetRangeMetricsLatencyHistogramDepricatedParams {

	return &GetRangeMetricsLatencyHistogramDepricatedParams{

		timeout: timeout,
	}
}

// NewGetRangeMetricsLatencyHistogramDepricatedParamsWithContext creates a new GetRangeMetricsLatencyHistogramDepricatedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRangeMetricsLatencyHistogramDepricatedParamsWithContext(ctx context.Context) *GetRangeMetricsLatencyHistogramDepricatedParams {

	return &GetRangeMetricsLatencyHistogramDepricatedParams{

		Context: ctx,
	}
}

// NewGetRangeMetricsLatencyHistogramDepricatedParamsWithHTTPClient creates a new GetRangeMetricsLatencyHistogramDepricatedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRangeMetricsLatencyHistogramDepricatedParamsWithHTTPClient(client *http.Client) *GetRangeMetricsLatencyHistogramDepricatedParams {

	return &GetRangeMetricsLatencyHistogramDepricatedParams{
		HTTPClient: client,
	}
}

/*GetRangeMetricsLatencyHistogramDepricatedParams contains all the parameters to send to the API endpoint
for the get range metrics latency histogram depricated operation typically these are written to a http.Request
*/
type GetRangeMetricsLatencyHistogramDepricatedParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get range metrics latency histogram depricated params
func (o *GetRangeMetricsLatencyHistogramDepricatedParams) WithTimeout(timeout time.Duration) *GetRangeMetricsLatencyHistogramDepricatedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get range metrics latency histogram depricated params
func (o *GetRangeMetricsLatencyHistogramDepricatedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get range metrics latency histogram depricated params
func (o *GetRangeMetricsLatencyHistogramDepricatedParams) WithContext(ctx context.Context) *GetRangeMetricsLatencyHistogramDepricatedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get range metrics latency histogram depricated params
func (o *GetRangeMetricsLatencyHistogramDepricatedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get range metrics latency histogram depricated params
func (o *GetRangeMetricsLatencyHistogramDepricatedParams) WithHTTPClient(client *http.Client) *GetRangeMetricsLatencyHistogramDepricatedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get range metrics latency histogram depricated params
func (o *GetRangeMetricsLatencyHistogramDepricatedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRangeMetricsLatencyHistogramDepricatedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

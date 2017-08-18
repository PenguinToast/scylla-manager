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

// NewGetRangeLatencyEstimatedHistogramParams creates a new GetRangeLatencyEstimatedHistogramParams object
// with the default values initialized.
func NewGetRangeLatencyEstimatedHistogramParams() *GetRangeLatencyEstimatedHistogramParams {
	var ()
	return &GetRangeLatencyEstimatedHistogramParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRangeLatencyEstimatedHistogramParamsWithTimeout creates a new GetRangeLatencyEstimatedHistogramParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRangeLatencyEstimatedHistogramParamsWithTimeout(timeout time.Duration) *GetRangeLatencyEstimatedHistogramParams {
	var ()
	return &GetRangeLatencyEstimatedHistogramParams{

		timeout: timeout,
	}
}

// NewGetRangeLatencyEstimatedHistogramParamsWithContext creates a new GetRangeLatencyEstimatedHistogramParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRangeLatencyEstimatedHistogramParamsWithContext(ctx context.Context) *GetRangeLatencyEstimatedHistogramParams {
	var ()
	return &GetRangeLatencyEstimatedHistogramParams{

		Context: ctx,
	}
}

// NewGetRangeLatencyEstimatedHistogramParamsWithHTTPClient creates a new GetRangeLatencyEstimatedHistogramParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRangeLatencyEstimatedHistogramParamsWithHTTPClient(client *http.Client) *GetRangeLatencyEstimatedHistogramParams {
	var ()
	return &GetRangeLatencyEstimatedHistogramParams{
		HTTPClient: client,
	}
}

/*GetRangeLatencyEstimatedHistogramParams contains all the parameters to send to the API endpoint
for the get range latency estimated histogram operation typically these are written to a http.Request
*/
type GetRangeLatencyEstimatedHistogramParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get range latency estimated histogram params
func (o *GetRangeLatencyEstimatedHistogramParams) WithTimeout(timeout time.Duration) *GetRangeLatencyEstimatedHistogramParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get range latency estimated histogram params
func (o *GetRangeLatencyEstimatedHistogramParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get range latency estimated histogram params
func (o *GetRangeLatencyEstimatedHistogramParams) WithContext(ctx context.Context) *GetRangeLatencyEstimatedHistogramParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get range latency estimated histogram params
func (o *GetRangeLatencyEstimatedHistogramParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get range latency estimated histogram params
func (o *GetRangeLatencyEstimatedHistogramParams) WithHTTPClient(client *http.Client) *GetRangeLatencyEstimatedHistogramParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get range latency estimated histogram params
func (o *GetRangeLatencyEstimatedHistogramParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get range latency estimated histogram params
func (o *GetRangeLatencyEstimatedHistogramParams) WithName(name string) *GetRangeLatencyEstimatedHistogramParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get range latency estimated histogram params
func (o *GetRangeLatencyEstimatedHistogramParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetRangeLatencyEstimatedHistogramParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

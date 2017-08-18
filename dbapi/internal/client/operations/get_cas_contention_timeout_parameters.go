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

// NewGetCasContentionTimeoutParams creates a new GetCasContentionTimeoutParams object
// with the default values initialized.
func NewGetCasContentionTimeoutParams() *GetCasContentionTimeoutParams {

	return &GetCasContentionTimeoutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCasContentionTimeoutParamsWithTimeout creates a new GetCasContentionTimeoutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCasContentionTimeoutParamsWithTimeout(timeout time.Duration) *GetCasContentionTimeoutParams {

	return &GetCasContentionTimeoutParams{

		timeout: timeout,
	}
}

// NewGetCasContentionTimeoutParamsWithContext creates a new GetCasContentionTimeoutParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCasContentionTimeoutParamsWithContext(ctx context.Context) *GetCasContentionTimeoutParams {

	return &GetCasContentionTimeoutParams{

		Context: ctx,
	}
}

// NewGetCasContentionTimeoutParamsWithHTTPClient creates a new GetCasContentionTimeoutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCasContentionTimeoutParamsWithHTTPClient(client *http.Client) *GetCasContentionTimeoutParams {

	return &GetCasContentionTimeoutParams{
		HTTPClient: client,
	}
}

/*GetCasContentionTimeoutParams contains all the parameters to send to the API endpoint
for the get cas contention timeout operation typically these are written to a http.Request
*/
type GetCasContentionTimeoutParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cas contention timeout params
func (o *GetCasContentionTimeoutParams) WithTimeout(timeout time.Duration) *GetCasContentionTimeoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cas contention timeout params
func (o *GetCasContentionTimeoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cas contention timeout params
func (o *GetCasContentionTimeoutParams) WithContext(ctx context.Context) *GetCasContentionTimeoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cas contention timeout params
func (o *GetCasContentionTimeoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cas contention timeout params
func (o *GetCasContentionTimeoutParams) WithHTTPClient(client *http.Client) *GetCasContentionTimeoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cas contention timeout params
func (o *GetCasContentionTimeoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCasContentionTimeoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

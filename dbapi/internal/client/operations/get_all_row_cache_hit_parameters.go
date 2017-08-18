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

// NewGetAllRowCacheHitParams creates a new GetAllRowCacheHitParams object
// with the default values initialized.
func NewGetAllRowCacheHitParams() *GetAllRowCacheHitParams {

	return &GetAllRowCacheHitParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllRowCacheHitParamsWithTimeout creates a new GetAllRowCacheHitParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllRowCacheHitParamsWithTimeout(timeout time.Duration) *GetAllRowCacheHitParams {

	return &GetAllRowCacheHitParams{

		timeout: timeout,
	}
}

// NewGetAllRowCacheHitParamsWithContext creates a new GetAllRowCacheHitParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllRowCacheHitParamsWithContext(ctx context.Context) *GetAllRowCacheHitParams {

	return &GetAllRowCacheHitParams{

		Context: ctx,
	}
}

// NewGetAllRowCacheHitParamsWithHTTPClient creates a new GetAllRowCacheHitParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllRowCacheHitParamsWithHTTPClient(client *http.Client) *GetAllRowCacheHitParams {

	return &GetAllRowCacheHitParams{
		HTTPClient: client,
	}
}

/*GetAllRowCacheHitParams contains all the parameters to send to the API endpoint
for the get all row cache hit operation typically these are written to a http.Request
*/
type GetAllRowCacheHitParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all row cache hit params
func (o *GetAllRowCacheHitParams) WithTimeout(timeout time.Duration) *GetAllRowCacheHitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all row cache hit params
func (o *GetAllRowCacheHitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all row cache hit params
func (o *GetAllRowCacheHitParams) WithContext(ctx context.Context) *GetAllRowCacheHitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all row cache hit params
func (o *GetAllRowCacheHitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all row cache hit params
func (o *GetAllRowCacheHitParams) WithHTTPClient(client *http.Client) *GetAllRowCacheHitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all row cache hit params
func (o *GetAllRowCacheHitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllRowCacheHitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

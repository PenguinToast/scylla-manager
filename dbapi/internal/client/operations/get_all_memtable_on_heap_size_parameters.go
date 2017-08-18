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

// NewGetAllMemtableOnHeapSizeParams creates a new GetAllMemtableOnHeapSizeParams object
// with the default values initialized.
func NewGetAllMemtableOnHeapSizeParams() *GetAllMemtableOnHeapSizeParams {

	return &GetAllMemtableOnHeapSizeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllMemtableOnHeapSizeParamsWithTimeout creates a new GetAllMemtableOnHeapSizeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllMemtableOnHeapSizeParamsWithTimeout(timeout time.Duration) *GetAllMemtableOnHeapSizeParams {

	return &GetAllMemtableOnHeapSizeParams{

		timeout: timeout,
	}
}

// NewGetAllMemtableOnHeapSizeParamsWithContext creates a new GetAllMemtableOnHeapSizeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllMemtableOnHeapSizeParamsWithContext(ctx context.Context) *GetAllMemtableOnHeapSizeParams {

	return &GetAllMemtableOnHeapSizeParams{

		Context: ctx,
	}
}

// NewGetAllMemtableOnHeapSizeParamsWithHTTPClient creates a new GetAllMemtableOnHeapSizeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllMemtableOnHeapSizeParamsWithHTTPClient(client *http.Client) *GetAllMemtableOnHeapSizeParams {

	return &GetAllMemtableOnHeapSizeParams{
		HTTPClient: client,
	}
}

/*GetAllMemtableOnHeapSizeParams contains all the parameters to send to the API endpoint
for the get all memtable on heap size operation typically these are written to a http.Request
*/
type GetAllMemtableOnHeapSizeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all memtable on heap size params
func (o *GetAllMemtableOnHeapSizeParams) WithTimeout(timeout time.Duration) *GetAllMemtableOnHeapSizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all memtable on heap size params
func (o *GetAllMemtableOnHeapSizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all memtable on heap size params
func (o *GetAllMemtableOnHeapSizeParams) WithContext(ctx context.Context) *GetAllMemtableOnHeapSizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all memtable on heap size params
func (o *GetAllMemtableOnHeapSizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all memtable on heap size params
func (o *GetAllMemtableOnHeapSizeParams) WithHTTPClient(client *http.Client) *GetAllMemtableOnHeapSizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all memtable on heap size params
func (o *GetAllMemtableOnHeapSizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllMemtableOnHeapSizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

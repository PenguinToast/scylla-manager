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

// NewGetAllMemtableOffHeapSizeParams creates a new GetAllMemtableOffHeapSizeParams object
// with the default values initialized.
func NewGetAllMemtableOffHeapSizeParams() *GetAllMemtableOffHeapSizeParams {

	return &GetAllMemtableOffHeapSizeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllMemtableOffHeapSizeParamsWithTimeout creates a new GetAllMemtableOffHeapSizeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllMemtableOffHeapSizeParamsWithTimeout(timeout time.Duration) *GetAllMemtableOffHeapSizeParams {

	return &GetAllMemtableOffHeapSizeParams{

		timeout: timeout,
	}
}

// NewGetAllMemtableOffHeapSizeParamsWithContext creates a new GetAllMemtableOffHeapSizeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllMemtableOffHeapSizeParamsWithContext(ctx context.Context) *GetAllMemtableOffHeapSizeParams {

	return &GetAllMemtableOffHeapSizeParams{

		Context: ctx,
	}
}

// NewGetAllMemtableOffHeapSizeParamsWithHTTPClient creates a new GetAllMemtableOffHeapSizeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllMemtableOffHeapSizeParamsWithHTTPClient(client *http.Client) *GetAllMemtableOffHeapSizeParams {

	return &GetAllMemtableOffHeapSizeParams{
		HTTPClient: client,
	}
}

/*GetAllMemtableOffHeapSizeParams contains all the parameters to send to the API endpoint
for the get all memtable off heap size operation typically these are written to a http.Request
*/
type GetAllMemtableOffHeapSizeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all memtable off heap size params
func (o *GetAllMemtableOffHeapSizeParams) WithTimeout(timeout time.Duration) *GetAllMemtableOffHeapSizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all memtable off heap size params
func (o *GetAllMemtableOffHeapSizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all memtable off heap size params
func (o *GetAllMemtableOffHeapSizeParams) WithContext(ctx context.Context) *GetAllMemtableOffHeapSizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all memtable off heap size params
func (o *GetAllMemtableOffHeapSizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all memtable off heap size params
func (o *GetAllMemtableOffHeapSizeParams) WithHTTPClient(client *http.Client) *GetAllMemtableOffHeapSizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all memtable off heap size params
func (o *GetAllMemtableOffHeapSizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllMemtableOffHeapSizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

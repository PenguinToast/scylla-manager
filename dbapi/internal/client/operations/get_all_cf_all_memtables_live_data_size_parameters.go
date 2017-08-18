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

// NewGetAllCfAllMemtablesLiveDataSizeParams creates a new GetAllCfAllMemtablesLiveDataSizeParams object
// with the default values initialized.
func NewGetAllCfAllMemtablesLiveDataSizeParams() *GetAllCfAllMemtablesLiveDataSizeParams {

	return &GetAllCfAllMemtablesLiveDataSizeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllCfAllMemtablesLiveDataSizeParamsWithTimeout creates a new GetAllCfAllMemtablesLiveDataSizeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllCfAllMemtablesLiveDataSizeParamsWithTimeout(timeout time.Duration) *GetAllCfAllMemtablesLiveDataSizeParams {

	return &GetAllCfAllMemtablesLiveDataSizeParams{

		timeout: timeout,
	}
}

// NewGetAllCfAllMemtablesLiveDataSizeParamsWithContext creates a new GetAllCfAllMemtablesLiveDataSizeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllCfAllMemtablesLiveDataSizeParamsWithContext(ctx context.Context) *GetAllCfAllMemtablesLiveDataSizeParams {

	return &GetAllCfAllMemtablesLiveDataSizeParams{

		Context: ctx,
	}
}

// NewGetAllCfAllMemtablesLiveDataSizeParamsWithHTTPClient creates a new GetAllCfAllMemtablesLiveDataSizeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllCfAllMemtablesLiveDataSizeParamsWithHTTPClient(client *http.Client) *GetAllCfAllMemtablesLiveDataSizeParams {

	return &GetAllCfAllMemtablesLiveDataSizeParams{
		HTTPClient: client,
	}
}

/*GetAllCfAllMemtablesLiveDataSizeParams contains all the parameters to send to the API endpoint
for the get all cf all memtables live data size operation typically these are written to a http.Request
*/
type GetAllCfAllMemtablesLiveDataSizeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all cf all memtables live data size params
func (o *GetAllCfAllMemtablesLiveDataSizeParams) WithTimeout(timeout time.Duration) *GetAllCfAllMemtablesLiveDataSizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all cf all memtables live data size params
func (o *GetAllCfAllMemtablesLiveDataSizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all cf all memtables live data size params
func (o *GetAllCfAllMemtablesLiveDataSizeParams) WithContext(ctx context.Context) *GetAllCfAllMemtablesLiveDataSizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all cf all memtables live data size params
func (o *GetAllCfAllMemtablesLiveDataSizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all cf all memtables live data size params
func (o *GetAllCfAllMemtablesLiveDataSizeParams) WithHTTPClient(client *http.Client) *GetAllCfAllMemtablesLiveDataSizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all cf all memtables live data size params
func (o *GetAllCfAllMemtablesLiveDataSizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllCfAllMemtablesLiveDataSizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

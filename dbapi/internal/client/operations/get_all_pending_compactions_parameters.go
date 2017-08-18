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

// NewGetAllPendingCompactionsParams creates a new GetAllPendingCompactionsParams object
// with the default values initialized.
func NewGetAllPendingCompactionsParams() *GetAllPendingCompactionsParams {

	return &GetAllPendingCompactionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllPendingCompactionsParamsWithTimeout creates a new GetAllPendingCompactionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllPendingCompactionsParamsWithTimeout(timeout time.Duration) *GetAllPendingCompactionsParams {

	return &GetAllPendingCompactionsParams{

		timeout: timeout,
	}
}

// NewGetAllPendingCompactionsParamsWithContext creates a new GetAllPendingCompactionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllPendingCompactionsParamsWithContext(ctx context.Context) *GetAllPendingCompactionsParams {

	return &GetAllPendingCompactionsParams{

		Context: ctx,
	}
}

// NewGetAllPendingCompactionsParamsWithHTTPClient creates a new GetAllPendingCompactionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllPendingCompactionsParamsWithHTTPClient(client *http.Client) *GetAllPendingCompactionsParams {

	return &GetAllPendingCompactionsParams{
		HTTPClient: client,
	}
}

/*GetAllPendingCompactionsParams contains all the parameters to send to the API endpoint
for the get all pending compactions operation typically these are written to a http.Request
*/
type GetAllPendingCompactionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all pending compactions params
func (o *GetAllPendingCompactionsParams) WithTimeout(timeout time.Duration) *GetAllPendingCompactionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all pending compactions params
func (o *GetAllPendingCompactionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all pending compactions params
func (o *GetAllPendingCompactionsParams) WithContext(ctx context.Context) *GetAllPendingCompactionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all pending compactions params
func (o *GetAllPendingCompactionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all pending compactions params
func (o *GetAllPendingCompactionsParams) WithHTTPClient(client *http.Client) *GetAllPendingCompactionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all pending compactions params
func (o *GetAllPendingCompactionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllPendingCompactionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

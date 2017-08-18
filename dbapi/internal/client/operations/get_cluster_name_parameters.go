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

// NewGetClusterNameParams creates a new GetClusterNameParams object
// with the default values initialized.
func NewGetClusterNameParams() *GetClusterNameParams {

	return &GetClusterNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterNameParamsWithTimeout creates a new GetClusterNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterNameParamsWithTimeout(timeout time.Duration) *GetClusterNameParams {

	return &GetClusterNameParams{

		timeout: timeout,
	}
}

// NewGetClusterNameParamsWithContext creates a new GetClusterNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterNameParamsWithContext(ctx context.Context) *GetClusterNameParams {

	return &GetClusterNameParams{

		Context: ctx,
	}
}

// NewGetClusterNameParamsWithHTTPClient creates a new GetClusterNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterNameParamsWithHTTPClient(client *http.Client) *GetClusterNameParams {

	return &GetClusterNameParams{
		HTTPClient: client,
	}
}

/*GetClusterNameParams contains all the parameters to send to the API endpoint
for the get cluster name operation typically these are written to a http.Request
*/
type GetClusterNameParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster name params
func (o *GetClusterNameParams) WithTimeout(timeout time.Duration) *GetClusterNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster name params
func (o *GetClusterNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster name params
func (o *GetClusterNameParams) WithContext(ctx context.Context) *GetClusterNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster name params
func (o *GetClusterNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster name params
func (o *GetClusterNameParams) WithHTTPClient(client *http.Client) *GetClusterNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster name params
func (o *GetClusterNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

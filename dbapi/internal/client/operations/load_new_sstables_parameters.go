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

// NewLoadNewSstablesParams creates a new LoadNewSstablesParams object
// with the default values initialized.
func NewLoadNewSstablesParams() *LoadNewSstablesParams {
	var ()
	return &LoadNewSstablesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoadNewSstablesParamsWithTimeout creates a new LoadNewSstablesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoadNewSstablesParamsWithTimeout(timeout time.Duration) *LoadNewSstablesParams {
	var ()
	return &LoadNewSstablesParams{

		timeout: timeout,
	}
}

// NewLoadNewSstablesParamsWithContext creates a new LoadNewSstablesParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoadNewSstablesParamsWithContext(ctx context.Context) *LoadNewSstablesParams {
	var ()
	return &LoadNewSstablesParams{

		Context: ctx,
	}
}

// NewLoadNewSstablesParamsWithHTTPClient creates a new LoadNewSstablesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoadNewSstablesParamsWithHTTPClient(client *http.Client) *LoadNewSstablesParams {
	var ()
	return &LoadNewSstablesParams{
		HTTPClient: client,
	}
}

/*LoadNewSstablesParams contains all the parameters to send to the API endpoint
for the load new sstables operation typically these are written to a http.Request
*/
type LoadNewSstablesParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the load new sstables params
func (o *LoadNewSstablesParams) WithTimeout(timeout time.Duration) *LoadNewSstablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the load new sstables params
func (o *LoadNewSstablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the load new sstables params
func (o *LoadNewSstablesParams) WithContext(ctx context.Context) *LoadNewSstablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the load new sstables params
func (o *LoadNewSstablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the load new sstables params
func (o *LoadNewSstablesParams) WithHTTPClient(client *http.Client) *LoadNewSstablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the load new sstables params
func (o *LoadNewSstablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the load new sstables params
func (o *LoadNewSstablesParams) WithName(name string) *LoadNewSstablesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the load new sstables params
func (o *LoadNewSstablesParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *LoadNewSstablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

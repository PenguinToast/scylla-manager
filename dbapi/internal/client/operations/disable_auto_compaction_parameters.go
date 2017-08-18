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

// NewDisableAutoCompactionParams creates a new DisableAutoCompactionParams object
// with the default values initialized.
func NewDisableAutoCompactionParams() *DisableAutoCompactionParams {
	var ()
	return &DisableAutoCompactionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDisableAutoCompactionParamsWithTimeout creates a new DisableAutoCompactionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDisableAutoCompactionParamsWithTimeout(timeout time.Duration) *DisableAutoCompactionParams {
	var ()
	return &DisableAutoCompactionParams{

		timeout: timeout,
	}
}

// NewDisableAutoCompactionParamsWithContext creates a new DisableAutoCompactionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDisableAutoCompactionParamsWithContext(ctx context.Context) *DisableAutoCompactionParams {
	var ()
	return &DisableAutoCompactionParams{

		Context: ctx,
	}
}

// NewDisableAutoCompactionParamsWithHTTPClient creates a new DisableAutoCompactionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDisableAutoCompactionParamsWithHTTPClient(client *http.Client) *DisableAutoCompactionParams {
	var ()
	return &DisableAutoCompactionParams{
		HTTPClient: client,
	}
}

/*DisableAutoCompactionParams contains all the parameters to send to the API endpoint
for the disable auto compaction operation typically these are written to a http.Request
*/
type DisableAutoCompactionParams struct {

	/*Cf
	  Comma seperated column family names

	*/
	Cf *string
	/*Keyspace
	  The keyspace

	*/
	Keyspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the disable auto compaction params
func (o *DisableAutoCompactionParams) WithTimeout(timeout time.Duration) *DisableAutoCompactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disable auto compaction params
func (o *DisableAutoCompactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disable auto compaction params
func (o *DisableAutoCompactionParams) WithContext(ctx context.Context) *DisableAutoCompactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disable auto compaction params
func (o *DisableAutoCompactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disable auto compaction params
func (o *DisableAutoCompactionParams) WithHTTPClient(client *http.Client) *DisableAutoCompactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disable auto compaction params
func (o *DisableAutoCompactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCf adds the cf to the disable auto compaction params
func (o *DisableAutoCompactionParams) WithCf(cf *string) *DisableAutoCompactionParams {
	o.SetCf(cf)
	return o
}

// SetCf adds the cf to the disable auto compaction params
func (o *DisableAutoCompactionParams) SetCf(cf *string) {
	o.Cf = cf
}

// WithKeyspace adds the keyspace to the disable auto compaction params
func (o *DisableAutoCompactionParams) WithKeyspace(keyspace string) *DisableAutoCompactionParams {
	o.SetKeyspace(keyspace)
	return o
}

// SetKeyspace adds the keyspace to the disable auto compaction params
func (o *DisableAutoCompactionParams) SetKeyspace(keyspace string) {
	o.Keyspace = keyspace
}

// WriteToRequest writes these params to a swagger request
func (o *DisableAutoCompactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cf != nil {

		// query param cf
		var qrCf string
		if o.Cf != nil {
			qrCf = *o.Cf
		}
		qCf := qrCf
		if qCf != "" {
			if err := r.SetQueryParam("cf", qCf); err != nil {
				return err
			}
		}

	}

	// path param keyspace
	if err := r.SetPathParam("keyspace", o.Keyspace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

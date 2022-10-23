// Code generated by go-swagger; DO NOT EDIT.

package script

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDelete1Params creates a new Delete1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDelete1Params() *Delete1Params {
	return &Delete1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDelete1ParamsWithTimeout creates a new Delete1Params object
// with the ability to set a timeout on a request.
func NewDelete1ParamsWithTimeout(timeout time.Duration) *Delete1Params {
	return &Delete1Params{
		timeout: timeout,
	}
}

// NewDelete1ParamsWithContext creates a new Delete1Params object
// with the ability to set a context for a request.
func NewDelete1ParamsWithContext(ctx context.Context) *Delete1Params {
	return &Delete1Params{
		Context: ctx,
	}
}

// NewDelete1ParamsWithHTTPClient creates a new Delete1Params object
// with the ability to set a custom HTTPClient for a request.
func NewDelete1ParamsWithHTTPClient(client *http.Client) *Delete1Params {
	return &Delete1Params{
		HTTPClient: client,
	}
}

/* Delete1Params contains all the parameters to send to the API endpoint
   for the delete 1 operation.

   Typically these are written to a http.Request.
*/
type Delete1Params struct {

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Delete1Params) WithDefaults() *Delete1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Delete1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete 1 params
func (o *Delete1Params) WithTimeout(timeout time.Duration) *Delete1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete 1 params
func (o *Delete1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete 1 params
func (o *Delete1Params) WithContext(ctx context.Context) *Delete1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete 1 params
func (o *Delete1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete 1 params
func (o *Delete1Params) WithHTTPClient(client *http.Client) *Delete1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete 1 params
func (o *Delete1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete 1 params
func (o *Delete1Params) WithName(name string) *Delete1Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete 1 params
func (o *Delete1Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *Delete1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
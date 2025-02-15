// Code generated by go-swagger; DO NOT EDIT.

package tasks

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

// NewRunParams creates a new RunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRunParams() *RunParams {
	return &RunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRunParamsWithTimeout creates a new RunParams object
// with the ability to set a timeout on a request.
func NewRunParamsWithTimeout(timeout time.Duration) *RunParams {
	return &RunParams{
		timeout: timeout,
	}
}

// NewRunParamsWithContext creates a new RunParams object
// with the ability to set a context for a request.
func NewRunParamsWithContext(ctx context.Context) *RunParams {
	return &RunParams{
		Context: ctx,
	}
}

// NewRunParamsWithHTTPClient creates a new RunParams object
// with the ability to set a custom HTTPClient for a request.
func NewRunParamsWithHTTPClient(client *http.Client) *RunParams {
	return &RunParams{
		HTTPClient: client,
	}
}

/* RunParams contains all the parameters to send to the API endpoint
   for the run operation.

   Typically these are written to a http.Request.
*/
type RunParams struct {

	/* ID.

	   Id of the task to run
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RunParams) WithDefaults() *RunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the run params
func (o *RunParams) WithTimeout(timeout time.Duration) *RunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the run params
func (o *RunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the run params
func (o *RunParams) WithContext(ctx context.Context) *RunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the run params
func (o *RunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the run params
func (o *RunParams) WithHTTPClient(client *http.Client) *RunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the run params
func (o *RunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the run params
func (o *RunParams) WithID(id string) *RunParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the run params
func (o *RunParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

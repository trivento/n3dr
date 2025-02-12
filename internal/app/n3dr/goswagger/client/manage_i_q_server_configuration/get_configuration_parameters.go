// Code generated by go-swagger; DO NOT EDIT.

package manage_i_q_server_configuration

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

// NewGetConfigurationParams creates a new GetConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConfigurationParams() *GetConfigurationParams {
	return &GetConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfigurationParamsWithTimeout creates a new GetConfigurationParams object
// with the ability to set a timeout on a request.
func NewGetConfigurationParamsWithTimeout(timeout time.Duration) *GetConfigurationParams {
	return &GetConfigurationParams{
		timeout: timeout,
	}
}

// NewGetConfigurationParamsWithContext creates a new GetConfigurationParams object
// with the ability to set a context for a request.
func NewGetConfigurationParamsWithContext(ctx context.Context) *GetConfigurationParams {
	return &GetConfigurationParams{
		Context: ctx,
	}
}

// NewGetConfigurationParamsWithHTTPClient creates a new GetConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConfigurationParamsWithHTTPClient(client *http.Client) *GetConfigurationParams {
	return &GetConfigurationParams{
		HTTPClient: client,
	}
}

/* GetConfigurationParams contains all the parameters to send to the API endpoint
   for the get configuration operation.

   Typically these are written to a http.Request.
*/
type GetConfigurationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConfigurationParams) WithDefaults() *GetConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get configuration params
func (o *GetConfigurationParams) WithTimeout(timeout time.Duration) *GetConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get configuration params
func (o *GetConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get configuration params
func (o *GetConfigurationParams) WithContext(ctx context.Context) *GetConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get configuration params
func (o *GetConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get configuration params
func (o *GetConfigurationParams) WithHTTPClient(client *http.Client) *GetConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get configuration params
func (o *GetConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

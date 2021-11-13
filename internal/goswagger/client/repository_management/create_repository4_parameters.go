// Code generated by go-swagger; DO NOT EDIT.

package repository_management

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

	"github.com/030/n3dr/internal/goswagger/models"
)

// NewCreateRepository4Params creates a new CreateRepository4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRepository4Params() *CreateRepository4Params {
	return &CreateRepository4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRepository4ParamsWithTimeout creates a new CreateRepository4Params object
// with the ability to set a timeout on a request.
func NewCreateRepository4ParamsWithTimeout(timeout time.Duration) *CreateRepository4Params {
	return &CreateRepository4Params{
		timeout: timeout,
	}
}

// NewCreateRepository4ParamsWithContext creates a new CreateRepository4Params object
// with the ability to set a context for a request.
func NewCreateRepository4ParamsWithContext(ctx context.Context) *CreateRepository4Params {
	return &CreateRepository4Params{
		Context: ctx,
	}
}

// NewCreateRepository4ParamsWithHTTPClient creates a new CreateRepository4Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRepository4ParamsWithHTTPClient(client *http.Client) *CreateRepository4Params {
	return &CreateRepository4Params{
		HTTPClient: client,
	}
}

/* CreateRepository4Params contains all the parameters to send to the API endpoint
   for the create repository 4 operation.

   Typically these are written to a http.Request.
*/
type CreateRepository4Params struct {

	// Body.
	Body *models.AptProxyRepositoryAPIRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create repository 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository4Params) WithDefaults() *CreateRepository4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create repository 4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRepository4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create repository 4 params
func (o *CreateRepository4Params) WithTimeout(timeout time.Duration) *CreateRepository4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create repository 4 params
func (o *CreateRepository4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create repository 4 params
func (o *CreateRepository4Params) WithContext(ctx context.Context) *CreateRepository4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create repository 4 params
func (o *CreateRepository4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create repository 4 params
func (o *CreateRepository4Params) WithHTTPClient(client *http.Client) *CreateRepository4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create repository 4 params
func (o *CreateRepository4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create repository 4 params
func (o *CreateRepository4Params) WithBody(body *models.AptProxyRepositoryAPIRequest) *CreateRepository4Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create repository 4 params
func (o *CreateRepository4Params) SetBody(body *models.AptProxyRepositoryAPIRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRepository4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

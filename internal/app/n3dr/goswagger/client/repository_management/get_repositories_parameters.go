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
)

// NewGetRepositoriesParams creates a new GetRepositoriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRepositoriesParams() *GetRepositoriesParams {
	return &GetRepositoriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesParamsWithTimeout creates a new GetRepositoriesParams object
// with the ability to set a timeout on a request.
func NewGetRepositoriesParamsWithTimeout(timeout time.Duration) *GetRepositoriesParams {
	return &GetRepositoriesParams{
		timeout: timeout,
	}
}

// NewGetRepositoriesParamsWithContext creates a new GetRepositoriesParams object
// with the ability to set a context for a request.
func NewGetRepositoriesParamsWithContext(ctx context.Context) *GetRepositoriesParams {
	return &GetRepositoriesParams{
		Context: ctx,
	}
}

// NewGetRepositoriesParamsWithHTTPClient creates a new GetRepositoriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRepositoriesParamsWithHTTPClient(client *http.Client) *GetRepositoriesParams {
	return &GetRepositoriesParams{
		HTTPClient: client,
	}
}

/* GetRepositoriesParams contains all the parameters to send to the API endpoint
   for the get repositories operation.

   Typically these are written to a http.Request.
*/
type GetRepositoriesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repositories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesParams) WithDefaults() *GetRepositoriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repositories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepositoriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repositories params
func (o *GetRepositoriesParams) WithTimeout(timeout time.Duration) *GetRepositoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories params
func (o *GetRepositoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories params
func (o *GetRepositoriesParams) WithContext(ctx context.Context) *GetRepositoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories params
func (o *GetRepositoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories params
func (o *GetRepositoriesParams) WithHTTPClient(client *http.Client) *GetRepositoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories params
func (o *GetRepositoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
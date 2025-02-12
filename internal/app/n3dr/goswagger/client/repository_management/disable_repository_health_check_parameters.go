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

// NewDisableRepositoryHealthCheckParams creates a new DisableRepositoryHealthCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDisableRepositoryHealthCheckParams() *DisableRepositoryHealthCheckParams {
	return &DisableRepositoryHealthCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDisableRepositoryHealthCheckParamsWithTimeout creates a new DisableRepositoryHealthCheckParams object
// with the ability to set a timeout on a request.
func NewDisableRepositoryHealthCheckParamsWithTimeout(timeout time.Duration) *DisableRepositoryHealthCheckParams {
	return &DisableRepositoryHealthCheckParams{
		timeout: timeout,
	}
}

// NewDisableRepositoryHealthCheckParamsWithContext creates a new DisableRepositoryHealthCheckParams object
// with the ability to set a context for a request.
func NewDisableRepositoryHealthCheckParamsWithContext(ctx context.Context) *DisableRepositoryHealthCheckParams {
	return &DisableRepositoryHealthCheckParams{
		Context: ctx,
	}
}

// NewDisableRepositoryHealthCheckParamsWithHTTPClient creates a new DisableRepositoryHealthCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewDisableRepositoryHealthCheckParamsWithHTTPClient(client *http.Client) *DisableRepositoryHealthCheckParams {
	return &DisableRepositoryHealthCheckParams{
		HTTPClient: client,
	}
}

/* DisableRepositoryHealthCheckParams contains all the parameters to send to the API endpoint
   for the disable repository health check operation.

   Typically these are written to a http.Request.
*/
type DisableRepositoryHealthCheckParams struct {

	/* RepositoryName.

	   Name of the repository to disable Repository Health Check for
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the disable repository health check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisableRepositoryHealthCheckParams) WithDefaults() *DisableRepositoryHealthCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the disable repository health check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisableRepositoryHealthCheckParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the disable repository health check params
func (o *DisableRepositoryHealthCheckParams) WithTimeout(timeout time.Duration) *DisableRepositoryHealthCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disable repository health check params
func (o *DisableRepositoryHealthCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disable repository health check params
func (o *DisableRepositoryHealthCheckParams) WithContext(ctx context.Context) *DisableRepositoryHealthCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disable repository health check params
func (o *DisableRepositoryHealthCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disable repository health check params
func (o *DisableRepositoryHealthCheckParams) WithHTTPClient(client *http.Client) *DisableRepositoryHealthCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disable repository health check params
func (o *DisableRepositoryHealthCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the disable repository health check params
func (o *DisableRepositoryHealthCheckParams) WithRepositoryName(repositoryName string) *DisableRepositoryHealthCheckParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the disable repository health check params
func (o *DisableRepositoryHealthCheckParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *DisableRepositoryHealthCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repositoryName
	if err := r.SetPathParam("repositoryName", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

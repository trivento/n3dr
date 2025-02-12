// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

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

// NewGetPrivilegeParams creates a new GetPrivilegeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPrivilegeParams() *GetPrivilegeParams {
	return &GetPrivilegeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivilegeParamsWithTimeout creates a new GetPrivilegeParams object
// with the ability to set a timeout on a request.
func NewGetPrivilegeParamsWithTimeout(timeout time.Duration) *GetPrivilegeParams {
	return &GetPrivilegeParams{
		timeout: timeout,
	}
}

// NewGetPrivilegeParamsWithContext creates a new GetPrivilegeParams object
// with the ability to set a context for a request.
func NewGetPrivilegeParamsWithContext(ctx context.Context) *GetPrivilegeParams {
	return &GetPrivilegeParams{
		Context: ctx,
	}
}

// NewGetPrivilegeParamsWithHTTPClient creates a new GetPrivilegeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPrivilegeParamsWithHTTPClient(client *http.Client) *GetPrivilegeParams {
	return &GetPrivilegeParams{
		HTTPClient: client,
	}
}

/* GetPrivilegeParams contains all the parameters to send to the API endpoint
   for the get privilege operation.

   Typically these are written to a http.Request.
*/
type GetPrivilegeParams struct {

	/* PrivilegeName.

	   The name of the privilege to retrieve.
	*/
	PrivilegeName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get privilege params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPrivilegeParams) WithDefaults() *GetPrivilegeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get privilege params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPrivilegeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get privilege params
func (o *GetPrivilegeParams) WithTimeout(timeout time.Duration) *GetPrivilegeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get privilege params
func (o *GetPrivilegeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get privilege params
func (o *GetPrivilegeParams) WithContext(ctx context.Context) *GetPrivilegeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get privilege params
func (o *GetPrivilegeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get privilege params
func (o *GetPrivilegeParams) WithHTTPClient(client *http.Client) *GetPrivilegeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get privilege params
func (o *GetPrivilegeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPrivilegeName adds the privilegeName to the get privilege params
func (o *GetPrivilegeParams) WithPrivilegeName(privilegeName string) *GetPrivilegeParams {
	o.SetPrivilegeName(privilegeName)
	return o
}

// SetPrivilegeName adds the privilegeName to the get privilege params
func (o *GetPrivilegeParams) SetPrivilegeName(privilegeName string) {
	o.PrivilegeName = privilegeName
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivilegeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param privilegeName
	if err := r.SetPathParam("privilegeName", o.PrivilegeName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

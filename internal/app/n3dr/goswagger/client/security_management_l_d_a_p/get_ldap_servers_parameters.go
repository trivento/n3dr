// Code generated by go-swagger; DO NOT EDIT.

package security_management_l_d_a_p

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

// NewGetLdapServersParams creates a new GetLdapServersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLdapServersParams() *GetLdapServersParams {
	return &GetLdapServersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLdapServersParamsWithTimeout creates a new GetLdapServersParams object
// with the ability to set a timeout on a request.
func NewGetLdapServersParamsWithTimeout(timeout time.Duration) *GetLdapServersParams {
	return &GetLdapServersParams{
		timeout: timeout,
	}
}

// NewGetLdapServersParamsWithContext creates a new GetLdapServersParams object
// with the ability to set a context for a request.
func NewGetLdapServersParamsWithContext(ctx context.Context) *GetLdapServersParams {
	return &GetLdapServersParams{
		Context: ctx,
	}
}

// NewGetLdapServersParamsWithHTTPClient creates a new GetLdapServersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLdapServersParamsWithHTTPClient(client *http.Client) *GetLdapServersParams {
	return &GetLdapServersParams{
		HTTPClient: client,
	}
}

/* GetLdapServersParams contains all the parameters to send to the API endpoint
   for the get ldap servers operation.

   Typically these are written to a http.Request.
*/
type GetLdapServersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get ldap servers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLdapServersParams) WithDefaults() *GetLdapServersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get ldap servers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLdapServersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get ldap servers params
func (o *GetLdapServersParams) WithTimeout(timeout time.Duration) *GetLdapServersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ldap servers params
func (o *GetLdapServersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ldap servers params
func (o *GetLdapServersParams) WithContext(ctx context.Context) *GetLdapServersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ldap servers params
func (o *GetLdapServersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ldap servers params
func (o *GetLdapServersParams) WithHTTPClient(client *http.Client) *GetLdapServersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ldap servers params
func (o *GetLdapServersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLdapServersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

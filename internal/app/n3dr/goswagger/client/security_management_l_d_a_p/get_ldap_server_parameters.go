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

// NewGetLdapServerParams creates a new GetLdapServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLdapServerParams() *GetLdapServerParams {
	return &GetLdapServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLdapServerParamsWithTimeout creates a new GetLdapServerParams object
// with the ability to set a timeout on a request.
func NewGetLdapServerParamsWithTimeout(timeout time.Duration) *GetLdapServerParams {
	return &GetLdapServerParams{
		timeout: timeout,
	}
}

// NewGetLdapServerParamsWithContext creates a new GetLdapServerParams object
// with the ability to set a context for a request.
func NewGetLdapServerParamsWithContext(ctx context.Context) *GetLdapServerParams {
	return &GetLdapServerParams{
		Context: ctx,
	}
}

// NewGetLdapServerParamsWithHTTPClient creates a new GetLdapServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLdapServerParamsWithHTTPClient(client *http.Client) *GetLdapServerParams {
	return &GetLdapServerParams{
		HTTPClient: client,
	}
}

/* GetLdapServerParams contains all the parameters to send to the API endpoint
   for the get ldap server operation.

   Typically these are written to a http.Request.
*/
type GetLdapServerParams struct {

	/* Name.

	   Name of the LDAP server to retrieve
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get ldap server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLdapServerParams) WithDefaults() *GetLdapServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get ldap server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLdapServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get ldap server params
func (o *GetLdapServerParams) WithTimeout(timeout time.Duration) *GetLdapServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ldap server params
func (o *GetLdapServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ldap server params
func (o *GetLdapServerParams) WithContext(ctx context.Context) *GetLdapServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ldap server params
func (o *GetLdapServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ldap server params
func (o *GetLdapServerParams) WithHTTPClient(client *http.Client) *GetLdapServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ldap server params
func (o *GetLdapServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get ldap server params
func (o *GetLdapServerParams) WithName(name string) *GetLdapServerParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get ldap server params
func (o *GetLdapServerParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetLdapServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

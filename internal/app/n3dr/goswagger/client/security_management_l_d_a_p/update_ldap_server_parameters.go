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

	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
)

// NewUpdateLdapServerParams creates a new UpdateLdapServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateLdapServerParams() *UpdateLdapServerParams {
	return &UpdateLdapServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLdapServerParamsWithTimeout creates a new UpdateLdapServerParams object
// with the ability to set a timeout on a request.
func NewUpdateLdapServerParamsWithTimeout(timeout time.Duration) *UpdateLdapServerParams {
	return &UpdateLdapServerParams{
		timeout: timeout,
	}
}

// NewUpdateLdapServerParamsWithContext creates a new UpdateLdapServerParams object
// with the ability to set a context for a request.
func NewUpdateLdapServerParamsWithContext(ctx context.Context) *UpdateLdapServerParams {
	return &UpdateLdapServerParams{
		Context: ctx,
	}
}

// NewUpdateLdapServerParamsWithHTTPClient creates a new UpdateLdapServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateLdapServerParamsWithHTTPClient(client *http.Client) *UpdateLdapServerParams {
	return &UpdateLdapServerParams{
		HTTPClient: client,
	}
}

/* UpdateLdapServerParams contains all the parameters to send to the API endpoint
   for the update ldap server operation.

   Typically these are written to a http.Request.
*/
type UpdateLdapServerParams struct {

	/* Body.

	   Updated values of LDAP server
	*/
	Body *models.UpdateLdapServerXo

	/* Name.

	   Name of the LDAP server to update
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update ldap server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLdapServerParams) WithDefaults() *UpdateLdapServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update ldap server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLdapServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update ldap server params
func (o *UpdateLdapServerParams) WithTimeout(timeout time.Duration) *UpdateLdapServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update ldap server params
func (o *UpdateLdapServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update ldap server params
func (o *UpdateLdapServerParams) WithContext(ctx context.Context) *UpdateLdapServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update ldap server params
func (o *UpdateLdapServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update ldap server params
func (o *UpdateLdapServerParams) WithHTTPClient(client *http.Client) *UpdateLdapServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update ldap server params
func (o *UpdateLdapServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update ldap server params
func (o *UpdateLdapServerParams) WithBody(body *models.UpdateLdapServerXo) *UpdateLdapServerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update ldap server params
func (o *UpdateLdapServerParams) SetBody(body *models.UpdateLdapServerXo) {
	o.Body = body
}

// WithName adds the name to the update ldap server params
func (o *UpdateLdapServerParams) WithName(name string) *UpdateLdapServerParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update ldap server params
func (o *UpdateLdapServerParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLdapServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
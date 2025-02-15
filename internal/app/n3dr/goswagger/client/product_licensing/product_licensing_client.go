// Code generated by go-swagger; DO NOT EDIT.

package product_licensing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new product licensing API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for product licensing API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetLicenseStatus(params *GetLicenseStatusParams, opts ...ClientOption) (*GetLicenseStatusOK, error)

	RemoveLicense(params *RemoveLicenseParams, opts ...ClientOption) error

	SetLicense(params *SetLicenseParams, opts ...ClientOption) (*SetLicenseOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetLicenseStatus gets the current license status
*/
func (a *Client) GetLicenseStatus(params *GetLicenseStatusParams, opts ...ClientOption) (*GetLicenseStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLicenseStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLicenseStatus",
		Method:             "GET",
		PathPattern:        "/v1/system/license",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLicenseStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLicenseStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLicenseStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveLicense uninstalls license if present
*/
func (a *Client) RemoveLicense(params *RemoveLicenseParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveLicenseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeLicense",
		Method:             "DELETE",
		PathPattern:        "/v1/system/license",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemoveLicenseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  SetLicense uploads a new license file

  Server must be restarted to take effect
*/
func (a *Client) SetLicense(params *SetLicenseParams, opts ...ClientOption) (*SetLicenseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetLicenseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setLicense",
		Method:             "POST",
		PathPattern:        "/v1/system/license",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetLicenseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetLicenseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setLicense: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

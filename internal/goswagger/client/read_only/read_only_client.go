// Code generated by go-swagger; DO NOT EDIT.

package read_only

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new read only API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for read only API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ForceRelease(params *ForceReleaseParams, opts ...ClientOption) (*ForceReleaseNoContent, error)

	Freeze(params *FreezeParams, opts ...ClientOption) (*FreezeNoContent, error)

	Get(params *GetParams, opts ...ClientOption) (*GetOK, error)

	Release(params *ReleaseParams, opts ...ClientOption) (*ReleaseNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ForceRelease forciblies release read only

  Forcibly release read-only status, including System initiated tasks. Warning: may result in data loss.
*/
func (a *Client) ForceRelease(params *ForceReleaseParams, opts ...ClientOption) (*ForceReleaseNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForceReleaseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "forceRelease",
		Method:             "POST",
		PathPattern:        "/v1/read-only/force-release",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ForceReleaseReader{formats: a.formats},
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
	success, ok := result.(*ForceReleaseNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for forceRelease: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Freeze enables read only
*/
func (a *Client) Freeze(params *FreezeParams, opts ...ClientOption) (*FreezeNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFreezeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "freeze",
		Method:             "POST",
		PathPattern:        "/v1/read-only/freeze",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FreezeReader{formats: a.formats},
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
	success, ok := result.(*FreezeNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for freeze: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Get gets read only state
*/
func (a *Client) Get(params *GetParams, opts ...ClientOption) (*GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get",
		Method:             "GET",
		PathPattern:        "/v1/read-only",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReader{formats: a.formats},
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
	success, ok := result.(*GetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Release releases read only

  Release administrator initiated read-only status. Will not release read-only caused by system tasks.
*/
func (a *Client) Release(params *ReleaseParams, opts ...ClientOption) (*ReleaseNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReleaseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "release",
		Method:             "POST",
		PathPattern:        "/v1/read-only/release",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReleaseReader{formats: a.formats},
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
	success, ok := result.(*ReleaseNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for release: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

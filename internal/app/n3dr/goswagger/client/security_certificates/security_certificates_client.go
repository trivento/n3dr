// Code generated by go-swagger; DO NOT EDIT.

package security_certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new security certificates API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for security certificates API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddCertificate(params *AddCertificateParams, opts ...ClientOption) (*AddCertificateCreated, error)

	GetTrustStoreCertificates(params *GetTrustStoreCertificatesParams, opts ...ClientOption) (*GetTrustStoreCertificatesOK, error)

	RemoveCertificate(params *RemoveCertificateParams, opts ...ClientOption) error

	RetrieveCertificate(params *RetrieveCertificateParams, opts ...ClientOption) (*RetrieveCertificateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddCertificate adds a certificate to the trust store
*/
func (a *Client) AddCertificate(params *AddCertificateParams, opts ...ClientOption) (*AddCertificateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddCertificateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addCertificate",
		Method:             "POST",
		PathPattern:        "/v1/security/ssl/truststore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddCertificateReader{formats: a.formats},
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
	success, ok := result.(*AddCertificateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addCertificate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTrustStoreCertificates retrieves a list of certificates added to the trust store
*/
func (a *Client) GetTrustStoreCertificates(params *GetTrustStoreCertificatesParams, opts ...ClientOption) (*GetTrustStoreCertificatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTrustStoreCertificatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTrustStoreCertificates",
		Method:             "GET",
		PathPattern:        "/v1/security/ssl/truststore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTrustStoreCertificatesReader{formats: a.formats},
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
	success, ok := result.(*GetTrustStoreCertificatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTrustStoreCertificates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveCertificate removes a certificate in the trust store
*/
func (a *Client) RemoveCertificate(params *RemoveCertificateParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveCertificateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeCertificate",
		Method:             "DELETE",
		PathPattern:        "/v1/security/ssl/truststore/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemoveCertificateReader{formats: a.formats},
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
  RetrieveCertificate helpers method to retrieve certificate details from a remote system
*/
func (a *Client) RetrieveCertificate(params *RetrieveCertificateParams, opts ...ClientOption) (*RetrieveCertificateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveCertificateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "retrieveCertificate",
		Method:             "GET",
		PathPattern:        "/v1/security/ssl",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RetrieveCertificateReader{formats: a.formats},
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
	success, ok := result.(*RetrieveCertificateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for retrieveCertificate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

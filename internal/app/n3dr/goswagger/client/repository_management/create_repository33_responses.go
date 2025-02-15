// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository33Reader is a Reader for the CreateRepository33 structure.
type CreateRepository33Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository33Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository33Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository33Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository33Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateRepository33MethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository33Created creates a CreateRepository33Created with default headers values
func NewCreateRepository33Created() *CreateRepository33Created {
	return &CreateRepository33Created{}
}

/* CreateRepository33Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository33Created struct {
}

func (o *CreateRepository33Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/r/proxy][%d] createRepository33Created ", 201)
}

func (o *CreateRepository33Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository33Unauthorized creates a CreateRepository33Unauthorized with default headers values
func NewCreateRepository33Unauthorized() *CreateRepository33Unauthorized {
	return &CreateRepository33Unauthorized{}
}

/* CreateRepository33Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository33Unauthorized struct {
}

func (o *CreateRepository33Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/r/proxy][%d] createRepository33Unauthorized ", 401)
}

func (o *CreateRepository33Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository33Forbidden creates a CreateRepository33Forbidden with default headers values
func NewCreateRepository33Forbidden() *CreateRepository33Forbidden {
	return &CreateRepository33Forbidden{}
}

/* CreateRepository33Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository33Forbidden struct {
}

func (o *CreateRepository33Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/r/proxy][%d] createRepository33Forbidden ", 403)
}

func (o *CreateRepository33Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository33MethodNotAllowed creates a CreateRepository33MethodNotAllowed with default headers values
func NewCreateRepository33MethodNotAllowed() *CreateRepository33MethodNotAllowed {
	return &CreateRepository33MethodNotAllowed{}
}

/* CreateRepository33MethodNotAllowed describes a response with status code 405, with default header values.

Feature is disabled in High Availability
*/
type CreateRepository33MethodNotAllowed struct {
}

func (o *CreateRepository33MethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/r/proxy][%d] createRepository33MethodNotAllowed ", 405)
}

func (o *CreateRepository33MethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

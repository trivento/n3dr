// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository7Reader is a Reader for the CreateRepository7 structure.
type CreateRepository7Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository7Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository7Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository7Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository7Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository7Created creates a CreateRepository7Created with default headers values
func NewCreateRepository7Created() *CreateRepository7Created {
	return &CreateRepository7Created{}
}

/* CreateRepository7Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository7Created struct {
}

func (o *CreateRepository7Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/proxy][%d] createRepository7Created ", 201)
}

func (o *CreateRepository7Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository7Unauthorized creates a CreateRepository7Unauthorized with default headers values
func NewCreateRepository7Unauthorized() *CreateRepository7Unauthorized {
	return &CreateRepository7Unauthorized{}
}

/* CreateRepository7Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository7Unauthorized struct {
}

func (o *CreateRepository7Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/proxy][%d] createRepository7Unauthorized ", 401)
}

func (o *CreateRepository7Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository7Forbidden creates a CreateRepository7Forbidden with default headers values
func NewCreateRepository7Forbidden() *CreateRepository7Forbidden {
	return &CreateRepository7Forbidden{}
}

/* CreateRepository7Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository7Forbidden struct {
}

func (o *CreateRepository7Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/raw/proxy][%d] createRepository7Forbidden ", 403)
}

func (o *CreateRepository7Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

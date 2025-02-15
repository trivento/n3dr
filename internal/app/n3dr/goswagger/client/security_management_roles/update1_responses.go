// Code generated by go-swagger; DO NOT EDIT.

package security_management_roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Update1Reader is a Reader for the Update1 structure.
type Update1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Update1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 403:
		result := NewUpdate1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdate1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdate1Forbidden creates a Update1Forbidden with default headers values
func NewUpdate1Forbidden() *Update1Forbidden {
	return &Update1Forbidden{}
}

/* Update1Forbidden describes a response with status code 403, with default header values.

Insufficient permissions to update role
*/
type Update1Forbidden struct {
}

func (o *Update1Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/security/roles/{id}][%d] update1Forbidden ", 403)
}

func (o *Update1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdate1NotFound creates a Update1NotFound with default headers values
func NewUpdate1NotFound() *Update1NotFound {
	return &Update1NotFound{}
}

/* Update1NotFound describes a response with status code 404, with default header values.

Role not found
*/
type Update1NotFound struct {
}

func (o *Update1NotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/security/roles/{id}][%d] update1NotFound ", 404)
}

func (o *Update1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdatePrivilege3Reader is a Reader for the UpdatePrivilege3 structure.
type UpdatePrivilege3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePrivilege3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewUpdatePrivilege3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdatePrivilege3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatePrivilege3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdatePrivilege3BadRequest creates a UpdatePrivilege3BadRequest with default headers values
func NewUpdatePrivilege3BadRequest() *UpdatePrivilege3BadRequest {
	return &UpdatePrivilege3BadRequest{}
}

/* UpdatePrivilege3BadRequest describes a response with status code 400, with default header values.

Privilege object not configured properly.
*/
type UpdatePrivilege3BadRequest struct {
}

func (o *UpdatePrivilege3BadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-admin/{privilegeName}][%d] updatePrivilege3BadRequest ", 400)
}

func (o *UpdatePrivilege3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePrivilege3Forbidden creates a UpdatePrivilege3Forbidden with default headers values
func NewUpdatePrivilege3Forbidden() *UpdatePrivilege3Forbidden {
	return &UpdatePrivilege3Forbidden{}
}

/* UpdatePrivilege3Forbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type UpdatePrivilege3Forbidden struct {
}

func (o *UpdatePrivilege3Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-admin/{privilegeName}][%d] updatePrivilege3Forbidden ", 403)
}

func (o *UpdatePrivilege3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePrivilege3NotFound creates a UpdatePrivilege3NotFound with default headers values
func NewUpdatePrivilege3NotFound() *UpdatePrivilege3NotFound {
	return &UpdatePrivilege3NotFound{}
}

/* UpdatePrivilege3NotFound describes a response with status code 404, with default header values.

Privilege not found in the system.
*/
type UpdatePrivilege3NotFound struct {
}

func (o *UpdatePrivilege3NotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/security/privileges/repository-admin/{privilegeName}][%d] updatePrivilege3NotFound ", 404)
}

func (o *UpdatePrivilege3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
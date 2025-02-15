// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreatePrivilege5Reader is a Reader for the CreatePrivilege5 structure.
type CreatePrivilege5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePrivilege5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewCreatePrivilege5BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreatePrivilege5Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePrivilege5BadRequest creates a CreatePrivilege5BadRequest with default headers values
func NewCreatePrivilege5BadRequest() *CreatePrivilege5BadRequest {
	return &CreatePrivilege5BadRequest{}
}

/* CreatePrivilege5BadRequest describes a response with status code 400, with default header values.

Privilege object not configured properly.
*/
type CreatePrivilege5BadRequest struct {
}

func (o *CreatePrivilege5BadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/script][%d] createPrivilege5BadRequest ", 400)
}

func (o *CreatePrivilege5BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePrivilege5Forbidden creates a CreatePrivilege5Forbidden with default headers values
func NewCreatePrivilege5Forbidden() *CreatePrivilege5Forbidden {
	return &CreatePrivilege5Forbidden{}
}

/* CreatePrivilege5Forbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type CreatePrivilege5Forbidden struct {
}

func (o *CreatePrivilege5Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/security/privileges/script][%d] createPrivilege5Forbidden ", 403)
}

func (o *CreatePrivilege5Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

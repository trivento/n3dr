// Code generated by go-swagger; DO NOT EDIT.

package security_management_privileges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/goswagger/models"
)

// GetPrivilegeReader is a Reader for the GetPrivilege structure.
type GetPrivilegeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivilegeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivilegeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetPrivilegeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPrivilegeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrivilegeOK creates a GetPrivilegeOK with default headers values
func NewGetPrivilegeOK() *GetPrivilegeOK {
	return &GetPrivilegeOK{}
}

/* GetPrivilegeOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPrivilegeOK struct {
	Payload *models.APIPrivilege
}

func (o *GetPrivilegeOK) Error() string {
	return fmt.Sprintf("[GET /v1/security/privileges/{privilegeId}][%d] getPrivilegeOK  %+v", 200, o.Payload)
}
func (o *GetPrivilegeOK) GetPayload() *models.APIPrivilege {
	return o.Payload
}

func (o *GetPrivilegeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPrivilege)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrivilegeForbidden creates a GetPrivilegeForbidden with default headers values
func NewGetPrivilegeForbidden() *GetPrivilegeForbidden {
	return &GetPrivilegeForbidden{}
}

/* GetPrivilegeForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type GetPrivilegeForbidden struct {
}

func (o *GetPrivilegeForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/security/privileges/{privilegeId}][%d] getPrivilegeForbidden ", 403)
}

func (o *GetPrivilegeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPrivilegeNotFound creates a GetPrivilegeNotFound with default headers values
func NewGetPrivilegeNotFound() *GetPrivilegeNotFound {
	return &GetPrivilegeNotFound{}
}

/* GetPrivilegeNotFound describes a response with status code 404, with default header values.

Privilege not found in the system.
*/
type GetPrivilegeNotFound struct {
}

func (o *GetPrivilegeNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/security/privileges/{privilegeId}][%d] getPrivilegeNotFound ", 404)
}

func (o *GetPrivilegeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

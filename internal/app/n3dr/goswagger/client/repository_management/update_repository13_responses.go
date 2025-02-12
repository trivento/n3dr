// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository13Reader is a Reader for the UpdateRepository13 structure.
type UpdateRepository13Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository13Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository13NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository13Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository13Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository13NoContent creates a UpdateRepository13NoContent with default headers values
func NewUpdateRepository13NoContent() *UpdateRepository13NoContent {
	return &UpdateRepository13NoContent{}
}

/* UpdateRepository13NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository13NoContent struct {
}

func (o *UpdateRepository13NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/proxy/{repositoryName}][%d] updateRepository13NoContent ", 204)
}

func (o *UpdateRepository13NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository13Unauthorized creates a UpdateRepository13Unauthorized with default headers values
func NewUpdateRepository13Unauthorized() *UpdateRepository13Unauthorized {
	return &UpdateRepository13Unauthorized{}
}

/* UpdateRepository13Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository13Unauthorized struct {
}

func (o *UpdateRepository13Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/proxy/{repositoryName}][%d] updateRepository13Unauthorized ", 401)
}

func (o *UpdateRepository13Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository13Forbidden creates a UpdateRepository13Forbidden with default headers values
func NewUpdateRepository13Forbidden() *UpdateRepository13Forbidden {
	return &UpdateRepository13Forbidden{}
}

/* UpdateRepository13Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository13Forbidden struct {
}

func (o *UpdateRepository13Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/nuget/proxy/{repositoryName}][%d] updateRepository13Forbidden ", 403)
}

func (o *UpdateRepository13Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

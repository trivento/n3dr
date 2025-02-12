// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository35Reader is a Reader for the UpdateRepository35 structure.
type UpdateRepository35Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository35Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository35NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository35Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository35Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRepository35NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository35NoContent creates a UpdateRepository35NoContent with default headers values
func NewUpdateRepository35NoContent() *UpdateRepository35NoContent {
	return &UpdateRepository35NoContent{}
}

/* UpdateRepository35NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository35NoContent struct {
}

func (o *UpdateRepository35NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/go/group/{repositoryName}][%d] updateRepository35NoContent ", 204)
}

func (o *UpdateRepository35NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository35Unauthorized creates a UpdateRepository35Unauthorized with default headers values
func NewUpdateRepository35Unauthorized() *UpdateRepository35Unauthorized {
	return &UpdateRepository35Unauthorized{}
}

/* UpdateRepository35Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository35Unauthorized struct {
}

func (o *UpdateRepository35Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/go/group/{repositoryName}][%d] updateRepository35Unauthorized ", 401)
}

func (o *UpdateRepository35Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository35Forbidden creates a UpdateRepository35Forbidden with default headers values
func NewUpdateRepository35Forbidden() *UpdateRepository35Forbidden {
	return &UpdateRepository35Forbidden{}
}

/* UpdateRepository35Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository35Forbidden struct {
}

func (o *UpdateRepository35Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/go/group/{repositoryName}][%d] updateRepository35Forbidden ", 403)
}

func (o *UpdateRepository35Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository35NotFound creates a UpdateRepository35NotFound with default headers values
func NewUpdateRepository35NotFound() *UpdateRepository35NotFound {
	return &UpdateRepository35NotFound{}
}

/* UpdateRepository35NotFound describes a response with status code 404, with default header values.

Repository not found
*/
type UpdateRepository35NotFound struct {
}

func (o *UpdateRepository35NotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/go/group/{repositoryName}][%d] updateRepository35NotFound ", 404)
}

func (o *UpdateRepository35NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

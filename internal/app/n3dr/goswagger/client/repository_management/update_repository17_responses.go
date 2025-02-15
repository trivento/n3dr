// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository17Reader is a Reader for the UpdateRepository17 structure.
type UpdateRepository17Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository17Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository17NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository17Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository17Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRepository17NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository17NoContent creates a UpdateRepository17NoContent with default headers values
func NewUpdateRepository17NoContent() *UpdateRepository17NoContent {
	return &UpdateRepository17NoContent{}
}

/* UpdateRepository17NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository17NoContent struct {
}

func (o *UpdateRepository17NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/group/{repositoryName}][%d] updateRepository17NoContent ", 204)
}

func (o *UpdateRepository17NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository17Unauthorized creates a UpdateRepository17Unauthorized with default headers values
func NewUpdateRepository17Unauthorized() *UpdateRepository17Unauthorized {
	return &UpdateRepository17Unauthorized{}
}

/* UpdateRepository17Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository17Unauthorized struct {
}

func (o *UpdateRepository17Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/group/{repositoryName}][%d] updateRepository17Unauthorized ", 401)
}

func (o *UpdateRepository17Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository17Forbidden creates a UpdateRepository17Forbidden with default headers values
func NewUpdateRepository17Forbidden() *UpdateRepository17Forbidden {
	return &UpdateRepository17Forbidden{}
}

/* UpdateRepository17Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository17Forbidden struct {
}

func (o *UpdateRepository17Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/group/{repositoryName}][%d] updateRepository17Forbidden ", 403)
}

func (o *UpdateRepository17Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository17NotFound creates a UpdateRepository17NotFound with default headers values
func NewUpdateRepository17NotFound() *UpdateRepository17NotFound {
	return &UpdateRepository17NotFound{}
}

/* UpdateRepository17NotFound describes a response with status code 404, with default header values.

Repository not found
*/
type UpdateRepository17NotFound struct {
}

func (o *UpdateRepository17NotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/docker/group/{repositoryName}][%d] updateRepository17NotFound ", 404)
}

func (o *UpdateRepository17NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/goswagger/models"
)

// GetTaskByIDReader is a Reader for the GetTaskByID structure.
type GetTaskByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTaskByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTaskByIDOK creates a GetTaskByIDOK with default headers values
func NewGetTaskByIDOK() *GetTaskByIDOK {
	return &GetTaskByIDOK{}
}

/* GetTaskByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTaskByIDOK struct {
	Payload *models.TaskXO
}

func (o *GetTaskByIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskByIdOK  %+v", 200, o.Payload)
}
func (o *GetTaskByIDOK) GetPayload() *models.TaskXO {
	return o.Payload
}

func (o *GetTaskByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskByIDNotFound creates a GetTaskByIDNotFound with default headers values
func NewGetTaskByIDNotFound() *GetTaskByIDNotFound {
	return &GetTaskByIDNotFound{}
}

/* GetTaskByIDNotFound describes a response with status code 404, with default header values.

Task not found
*/
type GetTaskByIDNotFound struct {
}

func (o *GetTaskByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/tasks/{id}][%d] getTaskByIdNotFound ", 404)
}

func (o *GetTaskByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
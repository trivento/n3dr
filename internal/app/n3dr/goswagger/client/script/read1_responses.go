// Code generated by go-swagger; DO NOT EDIT.

package script

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
)

// Read1Reader is a Reader for the Read1 structure.
type Read1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Read1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRead1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRead1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRead1OK creates a Read1OK with default headers values
func NewRead1OK() *Read1OK {
	return &Read1OK{}
}

/* Read1OK describes a response with status code 200, with default header values.

successful operation
*/
type Read1OK struct {
	Payload *models.ScriptXO
}

func (o *Read1OK) Error() string {
	return fmt.Sprintf("[GET /v1/script/{name}][%d] read1OK  %+v", 200, o.Payload)
}
func (o *Read1OK) GetPayload() *models.ScriptXO {
	return o.Payload
}

func (o *Read1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScriptXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRead1NotFound creates a Read1NotFound with default headers values
func NewRead1NotFound() *Read1NotFound {
	return &Read1NotFound{}
}

/* Read1NotFound describes a response with status code 404, with default header values.

No script with the specified name
*/
type Read1NotFound struct {
}

func (o *Read1NotFound) Error() string {
	return fmt.Sprintf("[GET /v1/script/{name}][%d] read1NotFound ", 404)
}

func (o *Read1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

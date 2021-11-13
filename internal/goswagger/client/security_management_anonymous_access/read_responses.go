// Code generated by go-swagger; DO NOT EDIT.

package security_management_anonymous_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/goswagger/models"
)

// ReadReader is a Reader for the Read structure.
type ReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewReadForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadOK creates a ReadOK with default headers values
func NewReadOK() *ReadOK {
	return &ReadOK{}
}

/* ReadOK describes a response with status code 200, with default header values.

successful operation
*/
type ReadOK struct {
	Payload *models.AnonymousAccessSettingsXO
}

func (o *ReadOK) Error() string {
	return fmt.Sprintf("[GET /v1/security/anonymous][%d] readOK  %+v", 200, o.Payload)
}
func (o *ReadOK) GetPayload() *models.AnonymousAccessSettingsXO {
	return o.Payload
}

func (o *ReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnonymousAccessSettingsXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadForbidden creates a ReadForbidden with default headers values
func NewReadForbidden() *ReadForbidden {
	return &ReadForbidden{}
}

/* ReadForbidden describes a response with status code 403, with default header values.

Insufficient permissions to update settings
*/
type ReadForbidden struct {
}

func (o *ReadForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/security/anonymous][%d] readForbidden ", 403)
}

func (o *ReadForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

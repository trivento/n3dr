// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IsWritableReader is a Reader for the IsWritable structure.
type IsWritableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsWritableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIsWritableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 503:
		result := NewIsWritableServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIsWritableOK creates a IsWritableOK with default headers values
func NewIsWritableOK() *IsWritableOK {
	return &IsWritableOK{}
}

/* IsWritableOK describes a response with status code 200, with default header values.

Available to service requests
*/
type IsWritableOK struct {
}

func (o *IsWritableOK) Error() string {
	return fmt.Sprintf("[GET /v1/status/writable][%d] isWritableOK ", 200)
}

func (o *IsWritableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsWritableServiceUnavailable creates a IsWritableServiceUnavailable with default headers values
func NewIsWritableServiceUnavailable() *IsWritableServiceUnavailable {
	return &IsWritableServiceUnavailable{}
}

/* IsWritableServiceUnavailable describes a response with status code 503, with default header values.

Unavailable to service requests
*/
type IsWritableServiceUnavailable struct {
}

func (o *IsWritableServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/status/writable][%d] isWritableServiceUnavailable ", 503)
}

func (o *IsWritableServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
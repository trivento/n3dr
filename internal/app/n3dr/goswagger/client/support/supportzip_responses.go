// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SupportzipReader is a Reader for the Supportzip structure.
type SupportzipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SupportzipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewSupportzipDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewSupportzipDefault creates a SupportzipDefault with default headers values
func NewSupportzipDefault(code int) *SupportzipDefault {
	return &SupportzipDefault{
		_statusCode: code,
	}
}

/* SupportzipDefault describes a response with status code -1, with default header values.

successful operation
*/
type SupportzipDefault struct {
	_statusCode int
}

// Code gets the status code for the supportzip default response
func (o *SupportzipDefault) Code() int {
	return o._statusCode
}

func (o *SupportzipDefault) Error() string {
	return fmt.Sprintf("[POST /v1/support/supportzip][%d] supportzip default ", o._statusCode)
}

func (o *SupportzipDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

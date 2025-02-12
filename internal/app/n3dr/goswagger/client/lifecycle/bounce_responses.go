// Code generated by go-swagger; DO NOT EDIT.

package lifecycle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// BounceReader is a Reader for the Bounce structure.
type BounceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BounceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewBounceDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewBounceDefault creates a BounceDefault with default headers values
func NewBounceDefault(code int) *BounceDefault {
	return &BounceDefault{
		_statusCode: code,
	}
}

/* BounceDefault describes a response with status code -1, with default header values.

successful operation
*/
type BounceDefault struct {
	_statusCode int
}

// Code gets the status code for the bounce default response
func (o *BounceDefault) Code() int {
	return o._statusCode
}

func (o *BounceDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/lifecycle/bounce][%d] bounce default ", o._statusCode)
}

func (o *BounceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package security_management_realms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SetActiveRealmsReader is a Reader for the SetActiveRealms structure.
type SetActiveRealmsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetActiveRealmsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewSetActiveRealmsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewSetActiveRealmsDefault creates a SetActiveRealmsDefault with default headers values
func NewSetActiveRealmsDefault(code int) *SetActiveRealmsDefault {
	return &SetActiveRealmsDefault{
		_statusCode: code,
	}
}

/* SetActiveRealmsDefault describes a response with status code -1, with default header values.

successful operation
*/
type SetActiveRealmsDefault struct {
	_statusCode int
}

// Code gets the status code for the set active realms default response
func (o *SetActiveRealmsDefault) Code() int {
	return o._statusCode
}

func (o *SetActiveRealmsDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/security/realms/active][%d] setActiveRealms default ", o._statusCode)
}

func (o *SetActiveRealmsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

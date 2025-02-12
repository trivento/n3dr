// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
)

// SearchReader is a Reader for the Search structure.
type SearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchOK creates a SearchOK with default headers values
func NewSearchOK() *SearchOK {
	return &SearchOK{}
}

/* SearchOK describes a response with status code 200, with default header values.

successful operation
*/
type SearchOK struct {
	Payload *models.PageComponentXO
}

func (o *SearchOK) Error() string {
	return fmt.Sprintf("[GET /v1/search][%d] searchOK  %+v", 200, o.Payload)
}
func (o *SearchOK) GetPayload() *models.PageComponentXO {
	return o.Payload
}

func (o *SearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageComponentXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

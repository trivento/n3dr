// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
)

// GetRepository27Reader is a Reader for the GetRepository27 structure.
type GetRepository27Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository27Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository27OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepository27OK creates a GetRepository27OK with default headers values
func NewGetRepository27OK() *GetRepository27OK {
	return &GetRepository27OK{}
}

/* GetRepository27OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository27OK struct {
	Payload *models.SimpleAPIGroupRepository
}

func (o *GetRepository27OK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/pypi/group/{repositoryName}][%d] getRepository27OK  %+v", 200, o.Payload)
}
func (o *GetRepository27OK) GetPayload() *models.SimpleAPIGroupRepository {
	return o.Payload
}

func (o *GetRepository27OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIGroupRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

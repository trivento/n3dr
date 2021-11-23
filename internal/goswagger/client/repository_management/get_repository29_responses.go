// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/goswagger/models"
)

// GetRepository29Reader is a Reader for the GetRepository29 structure.
type GetRepository29Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository29Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository29OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepository29OK creates a GetRepository29OK with default headers values
func NewGetRepository29OK() *GetRepository29OK {
	return &GetRepository29OK{}
}

/* GetRepository29OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository29OK struct {
	Payload *models.SimpleAPIProxyRepository
}

func (o *GetRepository29OK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/conda/proxy/{repositoryName}][%d] getRepository29OK  %+v", 200, o.Payload)
}
func (o *GetRepository29OK) GetPayload() *models.SimpleAPIProxyRepository {
	return o.Payload
}

func (o *GetRepository29OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIProxyRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
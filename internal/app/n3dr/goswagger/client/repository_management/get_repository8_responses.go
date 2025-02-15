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

// GetRepository8Reader is a Reader for the GetRepository8 structure.
type GetRepository8Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepository8Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepository8OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepository8OK creates a GetRepository8OK with default headers values
func NewGetRepository8OK() *GetRepository8OK {
	return &GetRepository8OK{}
}

/* GetRepository8OK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepository8OK struct {
	Payload *models.SimpleAPIProxyRepository
}

func (o *GetRepository8OK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/raw/proxy/{repositoryName}][%d] getRepository8OK  %+v", 200, o.Payload)
}
func (o *GetRepository8OK) GetPayload() *models.SimpleAPIProxyRepository {
	return o.Payload
}

func (o *GetRepository8OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIProxyRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

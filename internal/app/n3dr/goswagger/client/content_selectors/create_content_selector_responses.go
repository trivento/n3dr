// Code generated by go-swagger; DO NOT EDIT.

package content_selectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateContentSelectorReader is a Reader for the CreateContentSelector structure.
type CreateContentSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateContentSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCreateContentSelectorNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateContentSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateContentSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateContentSelectorNoContent creates a CreateContentSelectorNoContent with default headers values
func NewCreateContentSelectorNoContent() *CreateContentSelectorNoContent {
	return &CreateContentSelectorNoContent{}
}

/* CreateContentSelectorNoContent describes a response with status code 204, with default header values.

Content selector successfully created
*/
type CreateContentSelectorNoContent struct {
}

func (o *CreateContentSelectorNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/security/content-selectors][%d] createContentSelectorNoContent ", 204)
}

func (o *CreateContentSelectorNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateContentSelectorBadRequest creates a CreateContentSelectorBadRequest with default headers values
func NewCreateContentSelectorBadRequest() *CreateContentSelectorBadRequest {
	return &CreateContentSelectorBadRequest{}
}

/* CreateContentSelectorBadRequest describes a response with status code 400, with default header values.

Invalid request
*/
type CreateContentSelectorBadRequest struct {
}

func (o *CreateContentSelectorBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/security/content-selectors][%d] createContentSelectorBadRequest ", 400)
}

func (o *CreateContentSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateContentSelectorForbidden creates a CreateContentSelectorForbidden with default headers values
func NewCreateContentSelectorForbidden() *CreateContentSelectorForbidden {
	return &CreateContentSelectorForbidden{}
}

/* CreateContentSelectorForbidden describes a response with status code 403, with default header values.

Insufficient permissions to create content selectors
*/
type CreateContentSelectorForbidden struct {
}

func (o *CreateContentSelectorForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/security/content-selectors][%d] createContentSelectorForbidden ", 403)
}

func (o *CreateContentSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
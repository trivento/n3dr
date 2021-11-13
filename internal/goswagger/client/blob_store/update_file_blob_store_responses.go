// Code generated by go-swagger; DO NOT EDIT.

package blob_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateFileBlobStoreReader is a Reader for the UpdateFileBlobStore structure.
type UpdateFileBlobStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFileBlobStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateFileBlobStoreNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateFileBlobStoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateFileBlobStoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateFileBlobStoreNoContent creates a UpdateFileBlobStoreNoContent with default headers values
func NewUpdateFileBlobStoreNoContent() *UpdateFileBlobStoreNoContent {
	return &UpdateFileBlobStoreNoContent{}
}

/* UpdateFileBlobStoreNoContent describes a response with status code 204, with default header values.

Success
*/
type UpdateFileBlobStoreNoContent struct {
}

func (o *UpdateFileBlobStoreNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/blobstores/file/{name}][%d] updateFileBlobStoreNoContent ", 204)
}

func (o *UpdateFileBlobStoreNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateFileBlobStoreForbidden creates a UpdateFileBlobStoreForbidden with default headers values
func NewUpdateFileBlobStoreForbidden() *UpdateFileBlobStoreForbidden {
	return &UpdateFileBlobStoreForbidden{}
}

/* UpdateFileBlobStoreForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateFileBlobStoreForbidden struct {
}

func (o *UpdateFileBlobStoreForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/blobstores/file/{name}][%d] updateFileBlobStoreForbidden ", 403)
}

func (o *UpdateFileBlobStoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateFileBlobStoreNotFound creates a UpdateFileBlobStoreNotFound with default headers values
func NewUpdateFileBlobStoreNotFound() *UpdateFileBlobStoreNotFound {
	return &UpdateFileBlobStoreNotFound{}
}

/* UpdateFileBlobStoreNotFound describes a response with status code 404, with default header values.

Blob store not found
*/
type UpdateFileBlobStoreNotFound struct {
}

func (o *UpdateFileBlobStoreNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/blobstores/file/{name}][%d] updateFileBlobStoreNotFound ", 404)
}

func (o *UpdateFileBlobStoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

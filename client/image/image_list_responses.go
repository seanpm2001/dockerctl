// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-swagger/dockerctl/models"
)

// ImageListReader is a Reader for the ImageList structure.
type ImageListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewImageListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageListOK creates a ImageListOK with default headers values
func NewImageListOK() *ImageListOK {
	return &ImageListOK{}
}

/* ImageListOK describes a response with status code 200, with default header values.

Summary image data for the images matching the query
*/
type ImageListOK struct {
	Payload []*models.ImageSummary
}

func (o *ImageListOK) Error() string {
	return fmt.Sprintf("[GET /images/json][%d] imageListOK  %+v", 200, o.Payload)
}
func (o *ImageListOK) GetPayload() []*models.ImageSummary {
	return o.Payload
}

func (o *ImageListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageListInternalServerError creates a ImageListInternalServerError with default headers values
func NewImageListInternalServerError() *ImageListInternalServerError {
	return &ImageListInternalServerError{}
}

/* ImageListInternalServerError describes a response with status code 500, with default header values.

server error
*/
type ImageListInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ImageListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images/json][%d] imageListInternalServerError  %+v", 500, o.Payload)
}
func (o *ImageListInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ImageListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// ImageGetReader is a Reader for the ImageGet structure.
type ImageGetReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *ImageGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageGetOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewImageGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageGetOK creates a ImageGetOK with default headers values
func NewImageGetOK(writer io.Writer) *ImageGetOK {
	return &ImageGetOK{

		Payload: writer,
	}
}

/* ImageGetOK describes a response with status code 200, with default header values.

no error
*/
type ImageGetOK struct {
	Payload io.Writer
}

// IsSuccess returns true when this image get o k response has a 2xx status code
func (o *ImageGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image get o k response has a 3xx status code
func (o *ImageGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image get o k response has a 4xx status code
func (o *ImageGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image get o k response has a 5xx status code
func (o *ImageGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image get o k response a status code equal to that given
func (o *ImageGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImageGetOK) Error() string {
	return fmt.Sprintf("[GET /images/{name}/get][%d] imageGetOK  %+v", 200, o.Payload)
}

func (o *ImageGetOK) String() string {
	return fmt.Sprintf("[GET /images/{name}/get][%d] imageGetOK  %+v", 200, o.Payload)
}

func (o *ImageGetOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *ImageGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageGetInternalServerError creates a ImageGetInternalServerError with default headers values
func NewImageGetInternalServerError() *ImageGetInternalServerError {
	return &ImageGetInternalServerError{}
}

/* ImageGetInternalServerError describes a response with status code 500, with default header values.

server error
*/
type ImageGetInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this image get internal server error response has a 2xx status code
func (o *ImageGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image get internal server error response has a 3xx status code
func (o *ImageGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image get internal server error response has a 4xx status code
func (o *ImageGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image get internal server error response has a 5xx status code
func (o *ImageGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image get internal server error response a status code equal to that given
func (o *ImageGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImageGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images/{name}/get][%d] imageGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /images/{name}/get][%d] imageGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageGetInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ImageGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-swagger/dockerctl/models"
)

// ContainerAttachReader is a Reader for the ContainerAttach structure.
type ContainerAttachReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerAttachReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 101:
		result := NewContainerAttachSwitchingProtocols()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 200:
		result := NewContainerAttachOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewContainerAttachBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewContainerAttachNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerAttachInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerAttachSwitchingProtocols creates a ContainerAttachSwitchingProtocols with default headers values
func NewContainerAttachSwitchingProtocols() *ContainerAttachSwitchingProtocols {
	return &ContainerAttachSwitchingProtocols{}
}

/* ContainerAttachSwitchingProtocols describes a response with status code 101, with default header values.

no error, hints proxy about hijacking
*/
type ContainerAttachSwitchingProtocols struct {
}

// IsSuccess returns true when this container attach switching protocols response has a 2xx status code
func (o *ContainerAttachSwitchingProtocols) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container attach switching protocols response has a 3xx status code
func (o *ContainerAttachSwitchingProtocols) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container attach switching protocols response has a 4xx status code
func (o *ContainerAttachSwitchingProtocols) IsClientError() bool {
	return false
}

// IsServerError returns true when this container attach switching protocols response has a 5xx status code
func (o *ContainerAttachSwitchingProtocols) IsServerError() bool {
	return false
}

// IsCode returns true when this container attach switching protocols response a status code equal to that given
func (o *ContainerAttachSwitchingProtocols) IsCode(code int) bool {
	return code == 101
}

func (o *ContainerAttachSwitchingProtocols) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachSwitchingProtocols ", 101)
}

func (o *ContainerAttachSwitchingProtocols) String() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachSwitchingProtocols ", 101)
}

func (o *ContainerAttachSwitchingProtocols) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerAttachOK creates a ContainerAttachOK with default headers values
func NewContainerAttachOK() *ContainerAttachOK {
	return &ContainerAttachOK{}
}

/* ContainerAttachOK describes a response with status code 200, with default header values.

no error, no upgrade header found
*/
type ContainerAttachOK struct {
}

// IsSuccess returns true when this container attach o k response has a 2xx status code
func (o *ContainerAttachOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container attach o k response has a 3xx status code
func (o *ContainerAttachOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container attach o k response has a 4xx status code
func (o *ContainerAttachOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this container attach o k response has a 5xx status code
func (o *ContainerAttachOK) IsServerError() bool {
	return false
}

// IsCode returns true when this container attach o k response a status code equal to that given
func (o *ContainerAttachOK) IsCode(code int) bool {
	return code == 200
}

func (o *ContainerAttachOK) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachOK ", 200)
}

func (o *ContainerAttachOK) String() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachOK ", 200)
}

func (o *ContainerAttachOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerAttachBadRequest creates a ContainerAttachBadRequest with default headers values
func NewContainerAttachBadRequest() *ContainerAttachBadRequest {
	return &ContainerAttachBadRequest{}
}

/* ContainerAttachBadRequest describes a response with status code 400, with default header values.

bad parameter
*/
type ContainerAttachBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this container attach bad request response has a 2xx status code
func (o *ContainerAttachBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container attach bad request response has a 3xx status code
func (o *ContainerAttachBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container attach bad request response has a 4xx status code
func (o *ContainerAttachBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this container attach bad request response has a 5xx status code
func (o *ContainerAttachBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this container attach bad request response a status code equal to that given
func (o *ContainerAttachBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ContainerAttachBadRequest) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachBadRequest  %+v", 400, o.Payload)
}

func (o *ContainerAttachBadRequest) String() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachBadRequest  %+v", 400, o.Payload)
}

func (o *ContainerAttachBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerAttachBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerAttachNotFound creates a ContainerAttachNotFound with default headers values
func NewContainerAttachNotFound() *ContainerAttachNotFound {
	return &ContainerAttachNotFound{}
}

/* ContainerAttachNotFound describes a response with status code 404, with default header values.

no such container
*/
type ContainerAttachNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this container attach not found response has a 2xx status code
func (o *ContainerAttachNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container attach not found response has a 3xx status code
func (o *ContainerAttachNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container attach not found response has a 4xx status code
func (o *ContainerAttachNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container attach not found response has a 5xx status code
func (o *ContainerAttachNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container attach not found response a status code equal to that given
func (o *ContainerAttachNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ContainerAttachNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachNotFound  %+v", 404, o.Payload)
}

func (o *ContainerAttachNotFound) String() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachNotFound  %+v", 404, o.Payload)
}

func (o *ContainerAttachNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerAttachNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerAttachInternalServerError creates a ContainerAttachInternalServerError with default headers values
func NewContainerAttachInternalServerError() *ContainerAttachInternalServerError {
	return &ContainerAttachInternalServerError{}
}

/* ContainerAttachInternalServerError describes a response with status code 500, with default header values.

server error
*/
type ContainerAttachInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this container attach internal server error response has a 2xx status code
func (o *ContainerAttachInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container attach internal server error response has a 3xx status code
func (o *ContainerAttachInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container attach internal server error response has a 4xx status code
func (o *ContainerAttachInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container attach internal server error response has a 5xx status code
func (o *ContainerAttachInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container attach internal server error response a status code equal to that given
func (o *ContainerAttachInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerAttachInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerAttachInternalServerError) String() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerAttachInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerAttachInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

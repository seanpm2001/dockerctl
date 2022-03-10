// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-swagger/dockerctl/models"
)

// TaskLogsReader is a Reader for the TaskLogs structure.
type TaskLogsReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *TaskLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaskLogsOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewTaskLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTaskLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewTaskLogsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTaskLogsOK creates a TaskLogsOK with default headers values
func NewTaskLogsOK(writer io.Writer) *TaskLogsOK {
	return &TaskLogsOK{

		Payload: writer,
	}
}

/* TaskLogsOK describes a response with status code 200, with default header values.

logs returned as a stream in response body
*/
type TaskLogsOK struct {
	Payload io.Writer
}

// IsSuccess returns true when this task logs o k response has a 2xx status code
func (o *TaskLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this task logs o k response has a 3xx status code
func (o *TaskLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this task logs o k response has a 4xx status code
func (o *TaskLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this task logs o k response has a 5xx status code
func (o *TaskLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this task logs o k response a status code equal to that given
func (o *TaskLogsOK) IsCode(code int) bool {
	return code == 200
}

func (o *TaskLogsOK) Error() string {
	return fmt.Sprintf("[GET /tasks/{id}/logs][%d] taskLogsOK  %+v", 200, o.Payload)
}

func (o *TaskLogsOK) String() string {
	return fmt.Sprintf("[GET /tasks/{id}/logs][%d] taskLogsOK  %+v", 200, o.Payload)
}

func (o *TaskLogsOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *TaskLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaskLogsNotFound creates a TaskLogsNotFound with default headers values
func NewTaskLogsNotFound() *TaskLogsNotFound {
	return &TaskLogsNotFound{}
}

/* TaskLogsNotFound describes a response with status code 404, with default header values.

no such task
*/
type TaskLogsNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this task logs not found response has a 2xx status code
func (o *TaskLogsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this task logs not found response has a 3xx status code
func (o *TaskLogsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this task logs not found response has a 4xx status code
func (o *TaskLogsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this task logs not found response has a 5xx status code
func (o *TaskLogsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this task logs not found response a status code equal to that given
func (o *TaskLogsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *TaskLogsNotFound) Error() string {
	return fmt.Sprintf("[GET /tasks/{id}/logs][%d] taskLogsNotFound  %+v", 404, o.Payload)
}

func (o *TaskLogsNotFound) String() string {
	return fmt.Sprintf("[GET /tasks/{id}/logs][%d] taskLogsNotFound  %+v", 404, o.Payload)
}

func (o *TaskLogsNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TaskLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaskLogsInternalServerError creates a TaskLogsInternalServerError with default headers values
func NewTaskLogsInternalServerError() *TaskLogsInternalServerError {
	return &TaskLogsInternalServerError{}
}

/* TaskLogsInternalServerError describes a response with status code 500, with default header values.

server error
*/
type TaskLogsInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this task logs internal server error response has a 2xx status code
func (o *TaskLogsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this task logs internal server error response has a 3xx status code
func (o *TaskLogsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this task logs internal server error response has a 4xx status code
func (o *TaskLogsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this task logs internal server error response has a 5xx status code
func (o *TaskLogsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this task logs internal server error response a status code equal to that given
func (o *TaskLogsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TaskLogsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tasks/{id}/logs][%d] taskLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *TaskLogsInternalServerError) String() string {
	return fmt.Sprintf("[GET /tasks/{id}/logs][%d] taskLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *TaskLogsInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TaskLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaskLogsServiceUnavailable creates a TaskLogsServiceUnavailable with default headers values
func NewTaskLogsServiceUnavailable() *TaskLogsServiceUnavailable {
	return &TaskLogsServiceUnavailable{}
}

/* TaskLogsServiceUnavailable describes a response with status code 503, with default header values.

node is not part of a swarm
*/
type TaskLogsServiceUnavailable struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this task logs service unavailable response has a 2xx status code
func (o *TaskLogsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this task logs service unavailable response has a 3xx status code
func (o *TaskLogsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this task logs service unavailable response has a 4xx status code
func (o *TaskLogsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this task logs service unavailable response has a 5xx status code
func (o *TaskLogsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this task logs service unavailable response a status code equal to that given
func (o *TaskLogsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *TaskLogsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /tasks/{id}/logs][%d] taskLogsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *TaskLogsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /tasks/{id}/logs][%d] taskLogsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *TaskLogsServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TaskLogsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// TaskListReader is a Reader for the TaskList structure.
type TaskListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaskListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaskListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewTaskListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewTaskListServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTaskListOK creates a TaskListOK with default headers values
func NewTaskListOK() *TaskListOK {
	return &TaskListOK{}
}

/* TaskListOK describes a response with status code 200, with default header values.

no error
*/
type TaskListOK struct {
	Payload []*models.Task
}

// IsSuccess returns true when this task list o k response has a 2xx status code
func (o *TaskListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this task list o k response has a 3xx status code
func (o *TaskListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this task list o k response has a 4xx status code
func (o *TaskListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this task list o k response has a 5xx status code
func (o *TaskListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this task list o k response a status code equal to that given
func (o *TaskListOK) IsCode(code int) bool {
	return code == 200
}

func (o *TaskListOK) Error() string {
	return fmt.Sprintf("[GET /tasks][%d] taskListOK  %+v", 200, o.Payload)
}

func (o *TaskListOK) String() string {
	return fmt.Sprintf("[GET /tasks][%d] taskListOK  %+v", 200, o.Payload)
}

func (o *TaskListOK) GetPayload() []*models.Task {
	return o.Payload
}

func (o *TaskListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaskListInternalServerError creates a TaskListInternalServerError with default headers values
func NewTaskListInternalServerError() *TaskListInternalServerError {
	return &TaskListInternalServerError{}
}

/* TaskListInternalServerError describes a response with status code 500, with default header values.

server error
*/
type TaskListInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this task list internal server error response has a 2xx status code
func (o *TaskListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this task list internal server error response has a 3xx status code
func (o *TaskListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this task list internal server error response has a 4xx status code
func (o *TaskListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this task list internal server error response has a 5xx status code
func (o *TaskListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this task list internal server error response a status code equal to that given
func (o *TaskListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TaskListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tasks][%d] taskListInternalServerError  %+v", 500, o.Payload)
}

func (o *TaskListInternalServerError) String() string {
	return fmt.Sprintf("[GET /tasks][%d] taskListInternalServerError  %+v", 500, o.Payload)
}

func (o *TaskListInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TaskListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaskListServiceUnavailable creates a TaskListServiceUnavailable with default headers values
func NewTaskListServiceUnavailable() *TaskListServiceUnavailable {
	return &TaskListServiceUnavailable{}
}

/* TaskListServiceUnavailable describes a response with status code 503, with default header values.

node is not part of a swarm
*/
type TaskListServiceUnavailable struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this task list service unavailable response has a 2xx status code
func (o *TaskListServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this task list service unavailable response has a 3xx status code
func (o *TaskListServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this task list service unavailable response has a 4xx status code
func (o *TaskListServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this task list service unavailable response has a 5xx status code
func (o *TaskListServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this task list service unavailable response a status code equal to that given
func (o *TaskListServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *TaskListServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /tasks][%d] taskListServiceUnavailable  %+v", 503, o.Payload)
}

func (o *TaskListServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /tasks][%d] taskListServiceUnavailable  %+v", 503, o.Payload)
}

func (o *TaskListServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TaskListServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

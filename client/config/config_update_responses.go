// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-swagger/dockerctl/models"
)

// ConfigUpdateReader is a Reader for the ConfigUpdate structure.
type ConfigUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConfigUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewConfigUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewConfigUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewConfigUpdateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConfigUpdateOK creates a ConfigUpdateOK with default headers values
func NewConfigUpdateOK() *ConfigUpdateOK {
	return &ConfigUpdateOK{}
}

/* ConfigUpdateOK describes a response with status code 200, with default header values.

no error
*/
type ConfigUpdateOK struct {
}

// IsSuccess returns true when this config update o k response has a 2xx status code
func (o *ConfigUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this config update o k response has a 3xx status code
func (o *ConfigUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this config update o k response has a 4xx status code
func (o *ConfigUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this config update o k response has a 5xx status code
func (o *ConfigUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this config update o k response a status code equal to that given
func (o *ConfigUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ConfigUpdateOK) Error() string {
	return fmt.Sprintf("[POST /configs/{id}/update][%d] configUpdateOK ", 200)
}

func (o *ConfigUpdateOK) String() string {
	return fmt.Sprintf("[POST /configs/{id}/update][%d] configUpdateOK ", 200)
}

func (o *ConfigUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConfigUpdateBadRequest creates a ConfigUpdateBadRequest with default headers values
func NewConfigUpdateBadRequest() *ConfigUpdateBadRequest {
	return &ConfigUpdateBadRequest{}
}

/* ConfigUpdateBadRequest describes a response with status code 400, with default header values.

bad parameter
*/
type ConfigUpdateBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this config update bad request response has a 2xx status code
func (o *ConfigUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this config update bad request response has a 3xx status code
func (o *ConfigUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this config update bad request response has a 4xx status code
func (o *ConfigUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this config update bad request response has a 5xx status code
func (o *ConfigUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this config update bad request response a status code equal to that given
func (o *ConfigUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ConfigUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /configs/{id}/update][%d] configUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ConfigUpdateBadRequest) String() string {
	return fmt.Sprintf("[POST /configs/{id}/update][%d] configUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ConfigUpdateBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConfigUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigUpdateNotFound creates a ConfigUpdateNotFound with default headers values
func NewConfigUpdateNotFound() *ConfigUpdateNotFound {
	return &ConfigUpdateNotFound{}
}

/* ConfigUpdateNotFound describes a response with status code 404, with default header values.

no such config
*/
type ConfigUpdateNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this config update not found response has a 2xx status code
func (o *ConfigUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this config update not found response has a 3xx status code
func (o *ConfigUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this config update not found response has a 4xx status code
func (o *ConfigUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this config update not found response has a 5xx status code
func (o *ConfigUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this config update not found response a status code equal to that given
func (o *ConfigUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ConfigUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /configs/{id}/update][%d] configUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ConfigUpdateNotFound) String() string {
	return fmt.Sprintf("[POST /configs/{id}/update][%d] configUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ConfigUpdateNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConfigUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigUpdateInternalServerError creates a ConfigUpdateInternalServerError with default headers values
func NewConfigUpdateInternalServerError() *ConfigUpdateInternalServerError {
	return &ConfigUpdateInternalServerError{}
}

/* ConfigUpdateInternalServerError describes a response with status code 500, with default header values.

server error
*/
type ConfigUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this config update internal server error response has a 2xx status code
func (o *ConfigUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this config update internal server error response has a 3xx status code
func (o *ConfigUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this config update internal server error response has a 4xx status code
func (o *ConfigUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this config update internal server error response has a 5xx status code
func (o *ConfigUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this config update internal server error response a status code equal to that given
func (o *ConfigUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ConfigUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /configs/{id}/update][%d] configUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *ConfigUpdateInternalServerError) String() string {
	return fmt.Sprintf("[POST /configs/{id}/update][%d] configUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *ConfigUpdateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConfigUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigUpdateServiceUnavailable creates a ConfigUpdateServiceUnavailable with default headers values
func NewConfigUpdateServiceUnavailable() *ConfigUpdateServiceUnavailable {
	return &ConfigUpdateServiceUnavailable{}
}

/* ConfigUpdateServiceUnavailable describes a response with status code 503, with default header values.

node is not part of a swarm
*/
type ConfigUpdateServiceUnavailable struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this config update service unavailable response has a 2xx status code
func (o *ConfigUpdateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this config update service unavailable response has a 3xx status code
func (o *ConfigUpdateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this config update service unavailable response has a 4xx status code
func (o *ConfigUpdateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this config update service unavailable response has a 5xx status code
func (o *ConfigUpdateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this config update service unavailable response a status code equal to that given
func (o *ConfigUpdateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *ConfigUpdateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /configs/{id}/update][%d] configUpdateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ConfigUpdateServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /configs/{id}/update][%d] configUpdateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ConfigUpdateServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConfigUpdateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

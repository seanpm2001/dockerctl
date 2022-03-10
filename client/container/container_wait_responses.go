// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/go-swagger/dockerctl/models"
)

// ContainerWaitReader is a Reader for the ContainerWait structure.
type ContainerWaitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerWaitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerWaitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerWaitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerWaitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerWaitOK creates a ContainerWaitOK with default headers values
func NewContainerWaitOK() *ContainerWaitOK {
	return &ContainerWaitOK{}
}

/* ContainerWaitOK describes a response with status code 200, with default header values.

The container has exit.
*/
type ContainerWaitOK struct {
	Payload *ContainerWaitOKBody
}

// IsSuccess returns true when this container wait o k response has a 2xx status code
func (o *ContainerWaitOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container wait o k response has a 3xx status code
func (o *ContainerWaitOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container wait o k response has a 4xx status code
func (o *ContainerWaitOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this container wait o k response has a 5xx status code
func (o *ContainerWaitOK) IsServerError() bool {
	return false
}

// IsCode returns true when this container wait o k response a status code equal to that given
func (o *ContainerWaitOK) IsCode(code int) bool {
	return code == 200
}

func (o *ContainerWaitOK) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/wait][%d] containerWaitOK  %+v", 200, o.Payload)
}

func (o *ContainerWaitOK) String() string {
	return fmt.Sprintf("[POST /containers/{id}/wait][%d] containerWaitOK  %+v", 200, o.Payload)
}

func (o *ContainerWaitOK) GetPayload() *ContainerWaitOKBody {
	return o.Payload
}

func (o *ContainerWaitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerWaitOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerWaitNotFound creates a ContainerWaitNotFound with default headers values
func NewContainerWaitNotFound() *ContainerWaitNotFound {
	return &ContainerWaitNotFound{}
}

/* ContainerWaitNotFound describes a response with status code 404, with default header values.

no such container
*/
type ContainerWaitNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this container wait not found response has a 2xx status code
func (o *ContainerWaitNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container wait not found response has a 3xx status code
func (o *ContainerWaitNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container wait not found response has a 4xx status code
func (o *ContainerWaitNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container wait not found response has a 5xx status code
func (o *ContainerWaitNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container wait not found response a status code equal to that given
func (o *ContainerWaitNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ContainerWaitNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/wait][%d] containerWaitNotFound  %+v", 404, o.Payload)
}

func (o *ContainerWaitNotFound) String() string {
	return fmt.Sprintf("[POST /containers/{id}/wait][%d] containerWaitNotFound  %+v", 404, o.Payload)
}

func (o *ContainerWaitNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerWaitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerWaitInternalServerError creates a ContainerWaitInternalServerError with default headers values
func NewContainerWaitInternalServerError() *ContainerWaitInternalServerError {
	return &ContainerWaitInternalServerError{}
}

/* ContainerWaitInternalServerError describes a response with status code 500, with default header values.

server error
*/
type ContainerWaitInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this container wait internal server error response has a 2xx status code
func (o *ContainerWaitInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container wait internal server error response has a 3xx status code
func (o *ContainerWaitInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container wait internal server error response has a 4xx status code
func (o *ContainerWaitInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container wait internal server error response has a 5xx status code
func (o *ContainerWaitInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container wait internal server error response a status code equal to that given
func (o *ContainerWaitInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerWaitInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/wait][%d] containerWaitInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerWaitInternalServerError) String() string {
	return fmt.Sprintf("[POST /containers/{id}/wait][%d] containerWaitInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerWaitInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerWaitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerWaitOKBody ContainerWaitResponse
//
// OK response to ContainerWait operation
swagger:model ContainerWaitOKBody
*/
type ContainerWaitOKBody struct {

	// error
	Error *ContainerWaitOKBodyError `json:"Error,omitempty"`

	// Exit code of the container
	// Required: true
	StatusCode int64 `json:"StatusCode"`
}

// Validate validates this container wait o k body
func (o *ContainerWaitOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatusCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ContainerWaitOKBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerWaitOK" + "." + "Error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerWaitOK" + "." + "Error")
			}
			return err
		}
	}

	return nil
}

func (o *ContainerWaitOKBody) validateStatusCode(formats strfmt.Registry) error {

	if err := validate.Required("containerWaitOK"+"."+"StatusCode", "body", int64(o.StatusCode)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this container wait o k body based on the context it is used
func (o *ContainerWaitOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ContainerWaitOKBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerWaitOK" + "." + "Error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerWaitOK" + "." + "Error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ContainerWaitOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerWaitOKBody) UnmarshalBinary(b []byte) error {
	var res ContainerWaitOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerWaitOKBodyError container waiting error, if any
swagger:model ContainerWaitOKBodyError
*/
type ContainerWaitOKBodyError struct {

	// Details of an error
	Message string `json:"Message,omitempty"`
}

// Validate validates this container wait o k body error
func (o *ContainerWaitOKBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container wait o k body error based on context it is used
func (o *ContainerWaitOKBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerWaitOKBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerWaitOKBodyError) UnmarshalBinary(b []byte) error {
	var res ContainerWaitOKBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

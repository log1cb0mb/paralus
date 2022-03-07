// Code generated by go-swagger; DO NOT EDIT.

package bootstrap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/RafaySystems/rcloud-base/api/def/clients/sentry/models"
)

// BootstrapGetBootstrapAgentTemplateReader is a Reader for the BootstrapGetBootstrapAgentTemplate structure.
type BootstrapGetBootstrapAgentTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BootstrapGetBootstrapAgentTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBootstrapGetBootstrapAgentTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewBootstrapGetBootstrapAgentTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBootstrapGetBootstrapAgentTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBootstrapGetBootstrapAgentTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBootstrapGetBootstrapAgentTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBootstrapGetBootstrapAgentTemplateOK creates a BootstrapGetBootstrapAgentTemplateOK with default headers values
func NewBootstrapGetBootstrapAgentTemplateOK() *BootstrapGetBootstrapAgentTemplateOK {
	return &BootstrapGetBootstrapAgentTemplateOK{}
}

/* BootstrapGetBootstrapAgentTemplateOK describes a response with status code 200, with default header values.

A successful response.
*/
type BootstrapGetBootstrapAgentTemplateOK struct {
	Payload *models.SentryBootstrapAgentTemplate
}

func (o *BootstrapGetBootstrapAgentTemplateOK) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/bootstrap/template/{metadata.name}][%d] bootstrapGetBootstrapAgentTemplateOK  %+v", 200, o.Payload)
}
func (o *BootstrapGetBootstrapAgentTemplateOK) GetPayload() *models.SentryBootstrapAgentTemplate {
	return o.Payload
}

func (o *BootstrapGetBootstrapAgentTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SentryBootstrapAgentTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBootstrapGetBootstrapAgentTemplateForbidden creates a BootstrapGetBootstrapAgentTemplateForbidden with default headers values
func NewBootstrapGetBootstrapAgentTemplateForbidden() *BootstrapGetBootstrapAgentTemplateForbidden {
	return &BootstrapGetBootstrapAgentTemplateForbidden{}
}

/* BootstrapGetBootstrapAgentTemplateForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type BootstrapGetBootstrapAgentTemplateForbidden struct {
	Payload interface{}
}

func (o *BootstrapGetBootstrapAgentTemplateForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/bootstrap/template/{metadata.name}][%d] bootstrapGetBootstrapAgentTemplateForbidden  %+v", 403, o.Payload)
}
func (o *BootstrapGetBootstrapAgentTemplateForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *BootstrapGetBootstrapAgentTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBootstrapGetBootstrapAgentTemplateNotFound creates a BootstrapGetBootstrapAgentTemplateNotFound with default headers values
func NewBootstrapGetBootstrapAgentTemplateNotFound() *BootstrapGetBootstrapAgentTemplateNotFound {
	return &BootstrapGetBootstrapAgentTemplateNotFound{}
}

/* BootstrapGetBootstrapAgentTemplateNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type BootstrapGetBootstrapAgentTemplateNotFound struct {
	Payload interface{}
}

func (o *BootstrapGetBootstrapAgentTemplateNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/bootstrap/template/{metadata.name}][%d] bootstrapGetBootstrapAgentTemplateNotFound  %+v", 404, o.Payload)
}
func (o *BootstrapGetBootstrapAgentTemplateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *BootstrapGetBootstrapAgentTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBootstrapGetBootstrapAgentTemplateInternalServerError creates a BootstrapGetBootstrapAgentTemplateInternalServerError with default headers values
func NewBootstrapGetBootstrapAgentTemplateInternalServerError() *BootstrapGetBootstrapAgentTemplateInternalServerError {
	return &BootstrapGetBootstrapAgentTemplateInternalServerError{}
}

/* BootstrapGetBootstrapAgentTemplateInternalServerError describes a response with status code 500, with default header values.

Returned for internal server error
*/
type BootstrapGetBootstrapAgentTemplateInternalServerError struct {
	Payload interface{}
}

func (o *BootstrapGetBootstrapAgentTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/bootstrap/template/{metadata.name}][%d] bootstrapGetBootstrapAgentTemplateInternalServerError  %+v", 500, o.Payload)
}
func (o *BootstrapGetBootstrapAgentTemplateInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *BootstrapGetBootstrapAgentTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBootstrapGetBootstrapAgentTemplateDefault creates a BootstrapGetBootstrapAgentTemplateDefault with default headers values
func NewBootstrapGetBootstrapAgentTemplateDefault(code int) *BootstrapGetBootstrapAgentTemplateDefault {
	return &BootstrapGetBootstrapAgentTemplateDefault{
		_statusCode: code,
	}
}

/* BootstrapGetBootstrapAgentTemplateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BootstrapGetBootstrapAgentTemplateDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the bootstrap get bootstrap agent template default response
func (o *BootstrapGetBootstrapAgentTemplateDefault) Code() int {
	return o._statusCode
}

func (o *BootstrapGetBootstrapAgentTemplateDefault) Error() string {
	return fmt.Sprintf("[GET /v2/sentry/bootstrap/template/{metadata.name}][%d] Bootstrap_GetBootstrapAgentTemplate default  %+v", o._statusCode, o.Payload)
}
func (o *BootstrapGetBootstrapAgentTemplateDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *BootstrapGetBootstrapAgentTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
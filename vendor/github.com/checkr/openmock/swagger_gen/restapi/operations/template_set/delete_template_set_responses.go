// Code generated by go-swagger; DO NOT EDIT.

package template_set

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/checkr/openmock/swagger_gen/models"
)

// DeleteTemplateSetNoContentCode is the HTTP code returned for type DeleteTemplateSetNoContent
const DeleteTemplateSetNoContentCode int = 204

/*DeleteTemplateSetNoContent when successfully deleted

swagger:response deleteTemplateSetNoContent
*/
type DeleteTemplateSetNoContent struct {
}

// NewDeleteTemplateSetNoContent creates DeleteTemplateSetNoContent with default headers values
func NewDeleteTemplateSetNoContent() *DeleteTemplateSetNoContent {

	return &DeleteTemplateSetNoContent{}
}

// WriteResponse to the client
func (o *DeleteTemplateSetNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*DeleteTemplateSetDefault server error

swagger:response deleteTemplateSetDefault
*/
type DeleteTemplateSetDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTemplateSetDefault creates DeleteTemplateSetDefault with default headers values
func NewDeleteTemplateSetDefault(code int) *DeleteTemplateSetDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteTemplateSetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete template set default response
func (o *DeleteTemplateSetDefault) WithStatusCode(code int) *DeleteTemplateSetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete template set default response
func (o *DeleteTemplateSetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete template set default response
func (o *DeleteTemplateSetDefault) WithPayload(payload *models.Error) *DeleteTemplateSetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete template set default response
func (o *DeleteTemplateSetDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTemplateSetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

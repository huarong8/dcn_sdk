// Code generated by go-swagger; DO NOT EDIT.

package index_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ErrorPageUsingGETOKCode is the HTTP code returned for type ErrorPageUsingGETOK
const ErrorPageUsingGETOKCode int = 200

/*ErrorPageUsingGETOK OK

swagger:response errorPageUsingGETOK
*/
type ErrorPageUsingGETOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewErrorPageUsingGETOK creates ErrorPageUsingGETOK with default headers values
func NewErrorPageUsingGETOK() *ErrorPageUsingGETOK {

	return &ErrorPageUsingGETOK{}
}

// WithPayload adds the payload to the error page using g e t o k response
func (o *ErrorPageUsingGETOK) WithPayload(payload string) *ErrorPageUsingGETOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the error page using g e t o k response
func (o *ErrorPageUsingGETOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ErrorPageUsingGETOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ErrorPageUsingGETUnauthorizedCode is the HTTP code returned for type ErrorPageUsingGETUnauthorized
const ErrorPageUsingGETUnauthorizedCode int = 401

/*ErrorPageUsingGETUnauthorized Unauthorized

swagger:response errorPageUsingGETUnauthorized
*/
type ErrorPageUsingGETUnauthorized struct {
}

// NewErrorPageUsingGETUnauthorized creates ErrorPageUsingGETUnauthorized with default headers values
func NewErrorPageUsingGETUnauthorized() *ErrorPageUsingGETUnauthorized {

	return &ErrorPageUsingGETUnauthorized{}
}

// WriteResponse to the client
func (o *ErrorPageUsingGETUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ErrorPageUsingGETForbiddenCode is the HTTP code returned for type ErrorPageUsingGETForbidden
const ErrorPageUsingGETForbiddenCode int = 403

/*ErrorPageUsingGETForbidden Forbidden

swagger:response errorPageUsingGETForbidden
*/
type ErrorPageUsingGETForbidden struct {
}

// NewErrorPageUsingGETForbidden creates ErrorPageUsingGETForbidden with default headers values
func NewErrorPageUsingGETForbidden() *ErrorPageUsingGETForbidden {

	return &ErrorPageUsingGETForbidden{}
}

// WriteResponse to the client
func (o *ErrorPageUsingGETForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// ErrorPageUsingGETNotFoundCode is the HTTP code returned for type ErrorPageUsingGETNotFound
const ErrorPageUsingGETNotFoundCode int = 404

/*ErrorPageUsingGETNotFound Not Found

swagger:response errorPageUsingGETNotFound
*/
type ErrorPageUsingGETNotFound struct {
}

// NewErrorPageUsingGETNotFound creates ErrorPageUsingGETNotFound with default headers values
func NewErrorPageUsingGETNotFound() *ErrorPageUsingGETNotFound {

	return &ErrorPageUsingGETNotFound{}
}

// WriteResponse to the client
func (o *ErrorPageUsingGETNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
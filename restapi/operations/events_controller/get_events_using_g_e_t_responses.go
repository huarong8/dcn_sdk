// Code generated by go-swagger; DO NOT EDIT.

package events_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"dcn_sdk/models"
)

// GetEventsUsingGETOKCode is the HTTP code returned for type GetEventsUsingGETOK
const GetEventsUsingGETOKCode int = 200

/*GetEventsUsingGETOK OK

swagger:response getEventsUsingGETOK
*/
type GetEventsUsingGETOK struct {

	/*
	  In: Body
	*/
	Payload *models.EventsModel `json:"body,omitempty"`
}

// NewGetEventsUsingGETOK creates GetEventsUsingGETOK with default headers values
func NewGetEventsUsingGETOK() *GetEventsUsingGETOK {

	return &GetEventsUsingGETOK{}
}

// WithPayload adds the payload to the get events using g e t o k response
func (o *GetEventsUsingGETOK) WithPayload(payload *models.EventsModel) *GetEventsUsingGETOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get events using g e t o k response
func (o *GetEventsUsingGETOK) SetPayload(payload *models.EventsModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventsUsingGETOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventsUsingGETUnauthorizedCode is the HTTP code returned for type GetEventsUsingGETUnauthorized
const GetEventsUsingGETUnauthorizedCode int = 401

/*GetEventsUsingGETUnauthorized Unauthorized

swagger:response getEventsUsingGETUnauthorized
*/
type GetEventsUsingGETUnauthorized struct {
}

// NewGetEventsUsingGETUnauthorized creates GetEventsUsingGETUnauthorized with default headers values
func NewGetEventsUsingGETUnauthorized() *GetEventsUsingGETUnauthorized {

	return &GetEventsUsingGETUnauthorized{}
}

// WriteResponse to the client
func (o *GetEventsUsingGETUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetEventsUsingGETForbiddenCode is the HTTP code returned for type GetEventsUsingGETForbidden
const GetEventsUsingGETForbiddenCode int = 403

/*GetEventsUsingGETForbidden Forbidden

swagger:response getEventsUsingGETForbidden
*/
type GetEventsUsingGETForbidden struct {
}

// NewGetEventsUsingGETForbidden creates GetEventsUsingGETForbidden with default headers values
func NewGetEventsUsingGETForbidden() *GetEventsUsingGETForbidden {

	return &GetEventsUsingGETForbidden{}
}

// WriteResponse to the client
func (o *GetEventsUsingGETForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetEventsUsingGETNotFoundCode is the HTTP code returned for type GetEventsUsingGETNotFound
const GetEventsUsingGETNotFoundCode int = 404

/*GetEventsUsingGETNotFound Not Found

swagger:response getEventsUsingGETNotFound
*/
type GetEventsUsingGETNotFound struct {
}

// NewGetEventsUsingGETNotFound creates GetEventsUsingGETNotFound with default headers values
func NewGetEventsUsingGETNotFound() *GetEventsUsingGETNotFound {

	return &GetEventsUsingGETNotFound{}
}

// WriteResponse to the client
func (o *GetEventsUsingGETNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

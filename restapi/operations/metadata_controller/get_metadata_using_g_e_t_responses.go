// Code generated by go-swagger; DO NOT EDIT.

package metadata_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"dcn_sdk/models"
)

// GetMetadataUsingGETOKCode is the HTTP code returned for type GetMetadataUsingGETOK
const GetMetadataUsingGETOKCode int = 200

/*GetMetadataUsingGETOK OK

swagger:response getMetadataUsingGETOK
*/
type GetMetadataUsingGETOK struct {

	/*
	  In: Body
	*/
	Payload *models.MetadataModel `json:"body,omitempty"`
}

// NewGetMetadataUsingGETOK creates GetMetadataUsingGETOK with default headers values
func NewGetMetadataUsingGETOK() *GetMetadataUsingGETOK {

	return &GetMetadataUsingGETOK{}
}

// WithPayload adds the payload to the get metadata using g e t o k response
func (o *GetMetadataUsingGETOK) WithPayload(payload *models.MetadataModel) *GetMetadataUsingGETOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get metadata using g e t o k response
func (o *GetMetadataUsingGETOK) SetPayload(payload *models.MetadataModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMetadataUsingGETOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMetadataUsingGETUnauthorizedCode is the HTTP code returned for type GetMetadataUsingGETUnauthorized
const GetMetadataUsingGETUnauthorizedCode int = 401

/*GetMetadataUsingGETUnauthorized Unauthorized

swagger:response getMetadataUsingGETUnauthorized
*/
type GetMetadataUsingGETUnauthorized struct {
}

// NewGetMetadataUsingGETUnauthorized creates GetMetadataUsingGETUnauthorized with default headers values
func NewGetMetadataUsingGETUnauthorized() *GetMetadataUsingGETUnauthorized {

	return &GetMetadataUsingGETUnauthorized{}
}

// WriteResponse to the client
func (o *GetMetadataUsingGETUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetMetadataUsingGETForbiddenCode is the HTTP code returned for type GetMetadataUsingGETForbidden
const GetMetadataUsingGETForbiddenCode int = 403

/*GetMetadataUsingGETForbidden Forbidden

swagger:response getMetadataUsingGETForbidden
*/
type GetMetadataUsingGETForbidden struct {
}

// NewGetMetadataUsingGETForbidden creates GetMetadataUsingGETForbidden with default headers values
func NewGetMetadataUsingGETForbidden() *GetMetadataUsingGETForbidden {

	return &GetMetadataUsingGETForbidden{}
}

// WriteResponse to the client
func (o *GetMetadataUsingGETForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetMetadataUsingGETNotFoundCode is the HTTP code returned for type GetMetadataUsingGETNotFound
const GetMetadataUsingGETNotFoundCode int = 404

/*GetMetadataUsingGETNotFound Not Found

swagger:response getMetadataUsingGETNotFound
*/
type GetMetadataUsingGETNotFound struct {
}

// NewGetMetadataUsingGETNotFound creates GetMetadataUsingGETNotFound with default headers values
func NewGetMetadataUsingGETNotFound() *GetMetadataUsingGETNotFound {

	return &GetMetadataUsingGETNotFound{}
}

// WriteResponse to the client
func (o *GetMetadataUsingGETNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

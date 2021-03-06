// Code generated by go-swagger; DO NOT EDIT.

package block_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"dcn_sdk/models"
)

// GetBlockUsingGETOKCode is the HTTP code returned for type GetBlockUsingGETOK
const GetBlockUsingGETOKCode int = 200

/*GetBlockUsingGETOK OK

swagger:response getBlockUsingGETOK
*/
type GetBlockUsingGETOK struct {

	/*
	  In: Body
	*/
	Payload *models.BlockModel `json:"body,omitempty"`
}

// NewGetBlockUsingGETOK creates GetBlockUsingGETOK with default headers values
func NewGetBlockUsingGETOK() *GetBlockUsingGETOK {

	return &GetBlockUsingGETOK{}
}

// WithPayload adds the payload to the get block using g e t o k response
func (o *GetBlockUsingGETOK) WithPayload(payload *models.BlockModel) *GetBlockUsingGETOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get block using g e t o k response
func (o *GetBlockUsingGETOK) SetPayload(payload *models.BlockModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBlockUsingGETOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBlockUsingGETUnauthorizedCode is the HTTP code returned for type GetBlockUsingGETUnauthorized
const GetBlockUsingGETUnauthorizedCode int = 401

/*GetBlockUsingGETUnauthorized Unauthorized

swagger:response getBlockUsingGETUnauthorized
*/
type GetBlockUsingGETUnauthorized struct {
}

// NewGetBlockUsingGETUnauthorized creates GetBlockUsingGETUnauthorized with default headers values
func NewGetBlockUsingGETUnauthorized() *GetBlockUsingGETUnauthorized {

	return &GetBlockUsingGETUnauthorized{}
}

// WriteResponse to the client
func (o *GetBlockUsingGETUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetBlockUsingGETForbiddenCode is the HTTP code returned for type GetBlockUsingGETForbidden
const GetBlockUsingGETForbiddenCode int = 403

/*GetBlockUsingGETForbidden Forbidden

swagger:response getBlockUsingGETForbidden
*/
type GetBlockUsingGETForbidden struct {
}

// NewGetBlockUsingGETForbidden creates GetBlockUsingGETForbidden with default headers values
func NewGetBlockUsingGETForbidden() *GetBlockUsingGETForbidden {

	return &GetBlockUsingGETForbidden{}
}

// WriteResponse to the client
func (o *GetBlockUsingGETForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetBlockUsingGETNotFoundCode is the HTTP code returned for type GetBlockUsingGETNotFound
const GetBlockUsingGETNotFoundCode int = 404

/*GetBlockUsingGETNotFound Not Found

swagger:response getBlockUsingGETNotFound
*/
type GetBlockUsingGETNotFound struct {
}

// NewGetBlockUsingGETNotFound creates GetBlockUsingGETNotFound with default headers values
func NewGetBlockUsingGETNotFound() *GetBlockUsingGETNotFound {

	return &GetBlockUsingGETNotFound{}
}

// WriteResponse to the client
func (o *GetBlockUsingGETNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

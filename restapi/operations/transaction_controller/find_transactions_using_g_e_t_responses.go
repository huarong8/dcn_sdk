// Code generated by go-swagger; DO NOT EDIT.

package transaction_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"dcn_sdk/models"
)

// FindTransactionsUsingGETOKCode is the HTTP code returned for type FindTransactionsUsingGETOK
const FindTransactionsUsingGETOKCode int = 200

/*FindTransactionsUsingGETOK OK

swagger:response findTransactionsUsingGETOK
*/
type FindTransactionsUsingGETOK struct {

	/*
	  In: Body
	*/
	Payload *models.LatestTransactionModel `json:"body,omitempty"`
}

// NewFindTransactionsUsingGETOK creates FindTransactionsUsingGETOK with default headers values
func NewFindTransactionsUsingGETOK() *FindTransactionsUsingGETOK {

	return &FindTransactionsUsingGETOK{}
}

// WithPayload adds the payload to the find transactions using g e t o k response
func (o *FindTransactionsUsingGETOK) WithPayload(payload *models.LatestTransactionModel) *FindTransactionsUsingGETOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find transactions using g e t o k response
func (o *FindTransactionsUsingGETOK) SetPayload(payload *models.LatestTransactionModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindTransactionsUsingGETOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FindTransactionsUsingGETUnauthorizedCode is the HTTP code returned for type FindTransactionsUsingGETUnauthorized
const FindTransactionsUsingGETUnauthorizedCode int = 401

/*FindTransactionsUsingGETUnauthorized Unauthorized

swagger:response findTransactionsUsingGETUnauthorized
*/
type FindTransactionsUsingGETUnauthorized struct {
}

// NewFindTransactionsUsingGETUnauthorized creates FindTransactionsUsingGETUnauthorized with default headers values
func NewFindTransactionsUsingGETUnauthorized() *FindTransactionsUsingGETUnauthorized {

	return &FindTransactionsUsingGETUnauthorized{}
}

// WriteResponse to the client
func (o *FindTransactionsUsingGETUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// FindTransactionsUsingGETForbiddenCode is the HTTP code returned for type FindTransactionsUsingGETForbidden
const FindTransactionsUsingGETForbiddenCode int = 403

/*FindTransactionsUsingGETForbidden Forbidden

swagger:response findTransactionsUsingGETForbidden
*/
type FindTransactionsUsingGETForbidden struct {
}

// NewFindTransactionsUsingGETForbidden creates FindTransactionsUsingGETForbidden with default headers values
func NewFindTransactionsUsingGETForbidden() *FindTransactionsUsingGETForbidden {

	return &FindTransactionsUsingGETForbidden{}
}

// WriteResponse to the client
func (o *FindTransactionsUsingGETForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// FindTransactionsUsingGETNotFoundCode is the HTTP code returned for type FindTransactionsUsingGETNotFound
const FindTransactionsUsingGETNotFoundCode int = 404

/*FindTransactionsUsingGETNotFound Not Found

swagger:response findTransactionsUsingGETNotFound
*/
type FindTransactionsUsingGETNotFound struct {
}

// NewFindTransactionsUsingGETNotFound creates FindTransactionsUsingGETNotFound with default headers values
func NewFindTransactionsUsingGETNotFound() *FindTransactionsUsingGETNotFound {

	return &FindTransactionsUsingGETNotFound{}
}

// WriteResponse to the client
func (o *FindTransactionsUsingGETNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

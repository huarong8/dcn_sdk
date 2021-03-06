// Code generated by go-swagger; DO NOT EDIT.

package gas_oracle_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"dcn_sdk/models"
)

// GasPriceInfoUsingGETOKCode is the HTTP code returned for type GasPriceInfoUsingGETOK
const GasPriceInfoUsingGETOKCode int = 200

/*GasPriceInfoUsingGETOK OK

swagger:response gasPriceInfoUsingGETOK
*/
type GasPriceInfoUsingGETOK struct {

	/*
	  In: Body
	*/
	Payload *models.GasPriceModel `json:"body,omitempty"`
}

// NewGasPriceInfoUsingGETOK creates GasPriceInfoUsingGETOK with default headers values
func NewGasPriceInfoUsingGETOK() *GasPriceInfoUsingGETOK {

	return &GasPriceInfoUsingGETOK{}
}

// WithPayload adds the payload to the gas price info using g e t o k response
func (o *GasPriceInfoUsingGETOK) WithPayload(payload *models.GasPriceModel) *GasPriceInfoUsingGETOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the gas price info using g e t o k response
func (o *GasPriceInfoUsingGETOK) SetPayload(payload *models.GasPriceModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GasPriceInfoUsingGETOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GasPriceInfoUsingGETUnauthorizedCode is the HTTP code returned for type GasPriceInfoUsingGETUnauthorized
const GasPriceInfoUsingGETUnauthorizedCode int = 401

/*GasPriceInfoUsingGETUnauthorized Unauthorized

swagger:response gasPriceInfoUsingGETUnauthorized
*/
type GasPriceInfoUsingGETUnauthorized struct {
}

// NewGasPriceInfoUsingGETUnauthorized creates GasPriceInfoUsingGETUnauthorized with default headers values
func NewGasPriceInfoUsingGETUnauthorized() *GasPriceInfoUsingGETUnauthorized {

	return &GasPriceInfoUsingGETUnauthorized{}
}

// WriteResponse to the client
func (o *GasPriceInfoUsingGETUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GasPriceInfoUsingGETForbiddenCode is the HTTP code returned for type GasPriceInfoUsingGETForbidden
const GasPriceInfoUsingGETForbiddenCode int = 403

/*GasPriceInfoUsingGETForbidden Forbidden

swagger:response gasPriceInfoUsingGETForbidden
*/
type GasPriceInfoUsingGETForbidden struct {
}

// NewGasPriceInfoUsingGETForbidden creates GasPriceInfoUsingGETForbidden with default headers values
func NewGasPriceInfoUsingGETForbidden() *GasPriceInfoUsingGETForbidden {

	return &GasPriceInfoUsingGETForbidden{}
}

// WriteResponse to the client
func (o *GasPriceInfoUsingGETForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GasPriceInfoUsingGETNotFoundCode is the HTTP code returned for type GasPriceInfoUsingGETNotFound
const GasPriceInfoUsingGETNotFoundCode int = 404

/*GasPriceInfoUsingGETNotFound Not Found

swagger:response gasPriceInfoUsingGETNotFound
*/
type GasPriceInfoUsingGETNotFound struct {
}

// NewGasPriceInfoUsingGETNotFound creates GasPriceInfoUsingGETNotFound with default headers values
func NewGasPriceInfoUsingGETNotFound() *GasPriceInfoUsingGETNotFound {

	return &GasPriceInfoUsingGETNotFound{}
}

// WriteResponse to the client
func (o *GasPriceInfoUsingGETNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

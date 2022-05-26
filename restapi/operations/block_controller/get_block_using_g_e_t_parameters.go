// Code generated by go-swagger; DO NOT EDIT.

package block_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetBlockUsingGETParams creates a new GetBlockUsingGETParams object
//
// There are no default values defined in the spec.
func NewGetBlockUsingGETParams() GetBlockUsingGETParams {

	return GetBlockUsingGETParams{}
}

// GetBlockUsingGETParams contains all the bound params for the get block using g e t operation
// typically these are obtained from a http.Request
//
// swagger:parameters getBlockUsingGET
type GetBlockUsingGETParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The hash identifying the block.
	  In: path
	*/
	BlockHash *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetBlockUsingGETParams() beforehand.
func (o *GetBlockUsingGETParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rBlockHash, rhkBlockHash, _ := route.Params.GetOK("blockHash")
	if err := o.bindBlockHash(rBlockHash, rhkBlockHash, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBlockHash binds and validates parameter BlockHash from path.
func (o *GetBlockUsingGETParams) bindBlockHash(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// Parameter is provided by construction from the route
	o.BlockHash = &raw

	return nil
}
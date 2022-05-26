// Code generated by go-swagger; DO NOT EDIT.

package account_mapping_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewDeleteAccountUsingDELETEParams creates a new DeleteAccountUsingDELETEParams object
//
// There are no default values defined in the spec.
func NewDeleteAccountUsingDELETEParams() DeleteAccountUsingDELETEParams {

	return DeleteAccountUsingDELETEParams{}
}

// DeleteAccountUsingDELETEParams contains all the bound params for the delete account using d e l e t e operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteAccountUsingDELETE
type DeleteAccountUsingDELETEParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Account address to be mapped
	  In: path
	*/
	ID *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteAccountUsingDELETEParams() beforehand.
func (o *DeleteAccountUsingDELETEParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *DeleteAccountUsingDELETEParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// Parameter is provided by construction from the route
	o.ID = &raw

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package account_mapping_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"dcn_sdk/models"
)

// NewEditAccountUsingPUTParams creates a new EditAccountUsingPUTParams object
//
// There are no default values defined in the spec.
func NewEditAccountUsingPUTParams() EditAccountUsingPUTParams {

	return EditAccountUsingPUTParams{}
}

// EditAccountUsingPUTParams contains all the bound params for the edit account using p u t operation
// typically these are obtained from a http.Request
//
// swagger:parameters editAccountUsingPUT
type EditAccountUsingPUTParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*account
	  Required: true
	  In: body
	*/
	Account *models.EditAccountMapping
	/*Label to be mapped
	  In: path
	*/
	ID *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewEditAccountUsingPUTParams() beforehand.
func (o *EditAccountUsingPUTParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.EditAccountMapping
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("account", "body", ""))
			} else {
				res = append(res, errors.NewParseError("account", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Account = &body
			}
		}
	} else {
		res = append(res, errors.Required("account", "body", ""))
	}

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
func (o *EditAccountUsingPUTParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// Parameter is provided by construction from the route
	o.ID = &raw

	return nil
}
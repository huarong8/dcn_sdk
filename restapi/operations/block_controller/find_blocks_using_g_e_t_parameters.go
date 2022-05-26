// Code generated by go-swagger; DO NOT EDIT.

package block_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewFindBlocksUsingGETParams creates a new FindBlocksUsingGETParams object
//
// There are no default values defined in the spec.
func NewFindBlocksUsingGETParams() FindBlocksUsingGETParams {

	return FindBlocksUsingGETParams{}
}

// FindBlocksUsingGETParams contains all the bound params for the find blocks using g e t operation
// typically these are obtained from a http.Request
//
// swagger:parameters findBlocksUsingGET
type FindBlocksUsingGETParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	Direction *string
	/*
	  In: query
	*/
	Filter *string
	/*
	  In: query
	*/
	Page *int32
	/*
	  In: query
	*/
	Size *int32
	/*
	  In: query
	*/
	Sort *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewFindBlocksUsingGETParams() beforehand.
func (o *FindBlocksUsingGETParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDirection, qhkDirection, _ := qs.GetOK("direction")
	if err := o.bindDirection(qDirection, qhkDirection, route.Formats); err != nil {
		res = append(res, err)
	}

	qFilter, qhkFilter, _ := qs.GetOK("filter")
	if err := o.bindFilter(qFilter, qhkFilter, route.Formats); err != nil {
		res = append(res, err)
	}

	qPage, qhkPage, _ := qs.GetOK("page")
	if err := o.bindPage(qPage, qhkPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qSize, qhkSize, _ := qs.GetOK("size")
	if err := o.bindSize(qSize, qhkSize, route.Formats); err != nil {
		res = append(res, err)
	}

	qSort, qhkSort, _ := qs.GetOK("sort")
	if err := o.bindSort(qSort, qhkSort, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDirection binds and validates parameter Direction from query.
func (o *FindBlocksUsingGETParams) bindDirection(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Direction = &raw

	if err := o.validateDirection(formats); err != nil {
		return err
	}

	return nil
}

// validateDirection carries on validations for parameter Direction
func (o *FindBlocksUsingGETParams) validateDirection(formats strfmt.Registry) error {

	if err := validate.EnumCase("direction", "query", *o.Direction, []interface{}{"ASC", "DESC"}, true); err != nil {
		return err
	}

	return nil
}

// bindFilter binds and validates parameter Filter from query.
func (o *FindBlocksUsingGETParams) bindFilter(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Filter = &raw

	return nil
}

// bindPage binds and validates parameter Page from query.
func (o *FindBlocksUsingGETParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("page", "query", "int32", raw)
	}
	o.Page = &value

	return nil
}

// bindSize binds and validates parameter Size from query.
func (o *FindBlocksUsingGETParams) bindSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("size", "query", "int32", raw)
	}
	o.Size = &value

	return nil
}

// bindSort binds and validates parameter Sort from query.
func (o *FindBlocksUsingGETParams) bindSort(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Sort = &raw

	if err := o.validateSort(formats); err != nil {
		return err
	}

	return nil
}

// validateSort carries on validations for parameter Sort
func (o *FindBlocksUsingGETParams) validateSort(formats strfmt.Registry) error {

	if err := validate.EnumCase("sort", "query", *o.Sort, []interface{}{"number", "timestampISO", "transactionCount"}, true); err != nil {
		return err
	}

	return nil
}

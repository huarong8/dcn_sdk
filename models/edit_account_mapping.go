// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EditAccountMapping EditAccountMapping
//
// swagger:model EditAccountMapping
type EditAccountMapping struct {

	// label
	// Required: true
	Label *string `json:"label"`
}

// Validate validates this edit account mapping
func (m *EditAccountMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EditAccountMapping) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this edit account mapping based on context it is used
func (m *EditAccountMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EditAccountMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EditAccountMapping) UnmarshalBinary(b []byte) error {
	var res EditAccountMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

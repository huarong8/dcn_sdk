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

// SearchRecommendation SearchRecommendation
//
// swagger:model SearchRecommendation
type SearchRecommendation struct {

	// link
	// Required: true
	Link *Link `json:"link"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this search recommendation
func (m *SearchRecommendation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchRecommendation) validateLink(formats strfmt.Registry) error {

	if err := validate.Required("link", "body", m.Link); err != nil {
		return err
	}

	if m.Link != nil {
		if err := m.Link.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

func (m *SearchRecommendation) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this search recommendation based on the context it is used
func (m *SearchRecommendation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLink(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchRecommendation) contextValidateLink(ctx context.Context, formats strfmt.Registry) error {

	if m.Link != nil {
		if err := m.Link.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchRecommendation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchRecommendation) UnmarshalBinary(b []byte) error {
	var res SearchRecommendation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

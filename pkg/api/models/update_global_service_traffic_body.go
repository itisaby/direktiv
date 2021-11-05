// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateGlobalServiceTrafficBody UpdateGlobalServiceTrafficBody update global service traffic body
// Example: {"values":[{"percent":60,"revision":"global-fast-request-00002"},{"percent":40,"revision":"global-fast-request-00001"}]}
//
// swagger:model UpdateGlobalServiceTrafficBody
type UpdateGlobalServiceTrafficBody struct {

	// List of revision traffic targets
	// Required: true
	Values []*UpdateGlobalServiceTrafficParamsBodyValuesItems0 `json:"values"`
}

// Validate validates this update global service traffic body
func (m *UpdateGlobalServiceTrafficBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateGlobalServiceTrafficBody) validateValues(formats strfmt.Registry) error {

	if err := validate.Required("values", "body", m.Values); err != nil {
		return err
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update global service traffic body based on the context it is used
func (m *UpdateGlobalServiceTrafficBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateGlobalServiceTrafficBody) contextValidateValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Values); i++ {

		if m.Values[i] != nil {
			if err := m.Values[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateGlobalServiceTrafficBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateGlobalServiceTrafficBody) UnmarshalBinary(b []byte) error {
	var res UpdateGlobalServiceTrafficBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
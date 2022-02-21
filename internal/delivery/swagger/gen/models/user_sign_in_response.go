// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserSignInResponse user sign in response
//
// swagger:model user.signInResponse
type UserSignInResponse struct {

	// tokens
	Tokens *UserTokensResponse `json:"tokens,omitempty"`

	// user
	User *UserUser `json:"user,omitempty"`
}

// Validate validates this user sign in response
func (m *UserSignInResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTokens(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserSignInResponse) validateTokens(formats strfmt.Registry) error {
	if swag.IsZero(m.Tokens) { // not required
		return nil
	}

	if m.Tokens != nil {
		if err := m.Tokens.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tokens")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tokens")
			}
			return err
		}
	}

	return nil
}

func (m *UserSignInResponse) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this user sign in response based on the context it is used
func (m *UserSignInResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTokens(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserSignInResponse) contextValidateTokens(ctx context.Context, formats strfmt.Registry) error {

	if m.Tokens != nil {
		if err := m.Tokens.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tokens")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tokens")
			}
			return err
		}
	}

	return nil
}

func (m *UserSignInResponse) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserSignInResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserSignInResponse) UnmarshalBinary(b []byte) error {
	var res UserSignInResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

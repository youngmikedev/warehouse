// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserTokensResponse user tokens response
//
// swagger:model user.tokensResponse
type UserTokensResponse struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`

	// refresh token
	RefreshToken string `json:"refresh_token,omitempty"`
}

// Validate validates this user tokens response
func (m *UserTokensResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user tokens response based on context it is used
func (m *UserTokensResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserTokensResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserTokensResponse) UnmarshalBinary(b []byte) error {
	var res UserTokensResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

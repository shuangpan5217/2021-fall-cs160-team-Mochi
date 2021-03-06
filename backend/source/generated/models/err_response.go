// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ErrResponse err response
//
// swagger:model errResponse
type ErrResponse struct {

	// error message
	ErrMessage string `json:"errMessage,omitempty"`

	// http error code
	StatusCode int32 `json:"status_code,omitempty"`
}

// Validate validates this err response
func (m *ErrResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this err response based on context it is used
func (m *ErrResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrResponse) UnmarshalBinary(b []byte) error {
	var res ErrResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

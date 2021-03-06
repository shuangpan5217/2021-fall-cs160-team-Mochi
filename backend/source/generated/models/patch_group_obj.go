// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchGroupObj patch group obj
//
// swagger:model patchGroupObj
type PatchGroupObj struct {

	// group's desciption
	Description string `json:"description,omitempty"`

	// group name
	GroupName string `json:"group_name,omitempty"`
}

// Validate validates this patch group obj
func (m *PatchGroupObj) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch group obj based on context it is used
func (m *PatchGroupObj) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchGroupObj) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchGroupObj) UnmarshalBinary(b []byte) error {
	var res PatchGroupObj
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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
)

// NotesGetResponse array of note
//
// swagger:model notesGetResponse
type NotesGetResponse struct {

	// notes
	Notes []*NoteObjectResponse `json:"notes"`
}

// Validate validates this notes get response
func (m *NotesGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotesGetResponse) validateNotes(formats strfmt.Registry) error {
	if swag.IsZero(m.Notes) { // not required
		return nil
	}

	for i := 0; i < len(m.Notes); i++ {
		if swag.IsZero(m.Notes[i]) { // not required
			continue
		}

		if m.Notes[i] != nil {
			if err := m.Notes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this notes get response based on the context it is used
func (m *NotesGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNotes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotesGetResponse) contextValidateNotes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Notes); i++ {

		if m.Notes[i] != nil {
			if err := m.Notes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotesGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotesGetResponse) UnmarshalBinary(b []byte) error {
	var res NotesGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

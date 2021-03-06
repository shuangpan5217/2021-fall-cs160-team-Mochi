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

// GetAllGroupsObject get all groups object
//
// swagger:model getAllGroupsObject
type GetAllGroupsObject struct {

	// array of groups by username
	AllGroups []*GroupObj `json:"allGroups"`
}

// Validate validates this get all groups object
func (m *GetAllGroupsObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetAllGroupsObject) validateAllGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.AllGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.AllGroups); i++ {
		if swag.IsZero(m.AllGroups[i]) { // not required
			continue
		}

		if m.AllGroups[i] != nil {
			if err := m.AllGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get all groups object based on the context it is used
func (m *GetAllGroupsObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetAllGroupsObject) contextValidateAllGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AllGroups); i++ {

		if m.AllGroups[i] != nil {
			if err := m.AllGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetAllGroupsObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAllGroupsObject) UnmarshalBinary(b []byte) error {
	var res GetAllGroupsObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// Secret secret
//
// swagger:model Secret
type Secret struct {

	// created at
	// Example: 2017-07-20T13:55:28.678958722Z
	CreatedAt string `json:"CreatedAt,omitempty"`

	// ID
	// Example: blt1owaxmitz71s9v5zh81zun
	ID string `json:"ID,omitempty"`

	// spec
	Spec *SecretSpec `json:"Spec,omitempty"`

	// updated at
	// Example: 2017-07-20T13:55:28.678958722Z
	UpdatedAt string `json:"UpdatedAt,omitempty"`

	// version
	Version *ObjectVersion `json:"Version,omitempty"`
}

// Validate validates this secret
func (m *Secret) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secret) validateSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Spec")
			}
			return err
		}
	}

	return nil
}

func (m *Secret) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Version")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secret based on the context it is used
func (m *Secret) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secret) contextValidateSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.Spec != nil {
		if err := m.Spec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Spec")
			}
			return err
		}
	}

	return nil
}

func (m *Secret) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.Version != nil {
		if err := m.Version.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secret) UnmarshalBinary(b []byte) error {
	var res Secret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

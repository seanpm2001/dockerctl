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

// ContainerSummary container summary
//
// swagger:model ContainerSummary
type ContainerSummary []*ContainerSummaryItems0

// Validate validates this container summary
func (m ContainerSummary) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this container summary based on the context it is used
func (m ContainerSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if m[i] != nil {
			if err := m[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContainerSummaryItems0 container summary items0
//
// swagger:model ContainerSummaryItems0
type ContainerSummaryItems0 struct {

	// Command to run when starting the container
	Command string `json:"Command,omitempty"`

	// When the container was created
	Created int64 `json:"Created,omitempty"`

	// host config
	HostConfig *ContainerSummaryItems0HostConfig `json:"HostConfig,omitempty"`

	// The ID of this container
	ID string `json:"Id,omitempty"`

	// The name of the image used when creating this container
	Image string `json:"Image,omitempty"`

	// The ID of the image that this container was created from
	ImageID string `json:"ImageID,omitempty"`

	// User-defined key/value metadata.
	Labels map[string]string `json:"Labels,omitempty"`

	// mounts
	Mounts []*Mount `json:"Mounts"`

	// The names that this container has been given
	Names []string `json:"Names"`

	// network settings
	NetworkSettings *ContainerSummaryItems0NetworkSettings `json:"NetworkSettings,omitempty"`

	// The ports exposed by this container
	Ports []*Port `json:"Ports"`

	// The total size of all the files in this container
	SizeRootFs int64 `json:"SizeRootFs,omitempty"`

	// The size of files that have been created or changed by this container
	SizeRw int64 `json:"SizeRw,omitempty"`

	// The state of this container (e.g. `Exited`)
	State string `json:"State,omitempty"`

	// Additional human-readable status of this container (e.g. `Exit 0`)
	Status string `json:"Status,omitempty"`
}

// Validate validates this container summary items0
func (m *ContainerSummaryItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerSummaryItems0) validateHostConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.HostConfig) { // not required
		return nil
	}

	if m.HostConfig != nil {
		if err := m.HostConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HostConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("HostConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerSummaryItems0) validateMounts(formats strfmt.Registry) error {
	if swag.IsZero(m.Mounts) { // not required
		return nil
	}

	for i := 0; i < len(m.Mounts); i++ {
		if swag.IsZero(m.Mounts[i]) { // not required
			continue
		}

		if m.Mounts[i] != nil {
			if err := m.Mounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Mounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Mounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContainerSummaryItems0) validateNetworkSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkSettings) { // not required
		return nil
	}

	if m.NetworkSettings != nil {
		if err := m.NetworkSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NetworkSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NetworkSettings")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerSummaryItems0) validatePorts(formats strfmt.Registry) error {
	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for i := 0; i < len(m.Ports); i++ {
		if swag.IsZero(m.Ports[i]) { // not required
			continue
		}

		if m.Ports[i] != nil {
			if err := m.Ports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this container summary items0 based on the context it is used
func (m *ContainerSummaryItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHostConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerSummaryItems0) contextValidateHostConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.HostConfig != nil {
		if err := m.HostConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HostConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("HostConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerSummaryItems0) contextValidateMounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Mounts); i++ {

		if m.Mounts[i] != nil {
			if err := m.Mounts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Mounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Mounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContainerSummaryItems0) contextValidateNetworkSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkSettings != nil {
		if err := m.NetworkSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NetworkSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NetworkSettings")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerSummaryItems0) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ports); i++ {

		if m.Ports[i] != nil {
			if err := m.Ports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerSummaryItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerSummaryItems0) UnmarshalBinary(b []byte) error {
	var res ContainerSummaryItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ContainerSummaryItems0HostConfig container summary items0 host config
//
// swagger:model ContainerSummaryItems0HostConfig
type ContainerSummaryItems0HostConfig struct {

	// network mode
	NetworkMode string `json:"NetworkMode,omitempty"`
}

// Validate validates this container summary items0 host config
func (m *ContainerSummaryItems0HostConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container summary items0 host config based on context it is used
func (m *ContainerSummaryItems0HostConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContainerSummaryItems0HostConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerSummaryItems0HostConfig) UnmarshalBinary(b []byte) error {
	var res ContainerSummaryItems0HostConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ContainerSummaryItems0NetworkSettings A summary of the container's network settings
//
// swagger:model ContainerSummaryItems0NetworkSettings
type ContainerSummaryItems0NetworkSettings struct {

	// networks
	Networks map[string]EndpointSettings `json:"Networks,omitempty"`
}

// Validate validates this container summary items0 network settings
func (m *ContainerSummaryItems0NetworkSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerSummaryItems0NetworkSettings) validateNetworks(formats strfmt.Registry) error {
	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for k := range m.Networks {

		if err := validate.Required("NetworkSettings"+"."+"Networks"+"."+k, "body", m.Networks[k]); err != nil {
			return err
		}
		if val, ok := m.Networks[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NetworkSettings" + "." + "Networks" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NetworkSettings" + "." + "Networks" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this container summary items0 network settings based on the context it is used
func (m *ContainerSummaryItems0NetworkSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerSummaryItems0NetworkSettings) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Networks {

		if val, ok := m.Networks[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerSummaryItems0NetworkSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerSummaryItems0NetworkSettings) UnmarshalBinary(b []byte) error {
	var res ContainerSummaryItems0NetworkSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

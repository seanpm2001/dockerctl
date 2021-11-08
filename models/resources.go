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

// Resources A container's resources (cgroups config, ulimits, etc)
//
// swagger:model Resources
type Resources struct {

	// Limit read rate (bytes per second) from a device, in the form `[{"Path": "device_path", "Rate": rate}]`.
	//
	BlkioDeviceReadBps []*ThrottleDevice `json:"BlkioDeviceReadBps"`

	// Limit read rate (IO per second) from a device, in the form `[{"Path": "device_path", "Rate": rate}]`.
	//
	BlkioDeviceReadIOps []*ThrottleDevice `json:"BlkioDeviceReadIOps"`

	// Limit write rate (bytes per second) to a device, in the form `[{"Path": "device_path", "Rate": rate}]`.
	//
	BlkioDeviceWriteBps []*ThrottleDevice `json:"BlkioDeviceWriteBps"`

	// Limit write rate (IO per second) to a device, in the form `[{"Path": "device_path", "Rate": rate}]`.
	//
	BlkioDeviceWriteIOps []*ThrottleDevice `json:"BlkioDeviceWriteIOps"`

	// Block IO weight (relative weight).
	// Maximum: 1000
	// Minimum: 0
	BlkioWeight *int64 `json:"BlkioWeight,omitempty"`

	// Block IO weight (relative device weight) in the form `[{"Path": "device_path", "Weight": weight}]`.
	//
	BlkioWeightDevice []*ResourcesBlkioWeightDeviceItems0 `json:"BlkioWeightDevice"`

	// Path to `cgroups` under which the container's `cgroup` is created. If the path is not absolute, the path is considered to be relative to the `cgroups` path of the init process. Cgroups are created if they do not already exist.
	CgroupParent string `json:"CgroupParent,omitempty"`

	// The number of usable CPUs (Windows only).
	//
	// On Windows Server containers, the processor resource controls are mutually exclusive. The order of precedence is `CPUCount` first, then `CPUShares`, and `CPUPercent` last.
	//
	CPUCount int64 `json:"CpuCount,omitempty"`

	// The usable percentage of the available CPUs (Windows only).
	//
	// On Windows Server containers, the processor resource controls are mutually exclusive. The order of precedence is `CPUCount` first, then `CPUShares`, and `CPUPercent` last.
	//
	CPUPercent int64 `json:"CpuPercent,omitempty"`

	// The length of a CPU period in microseconds.
	CPUPeriod int64 `json:"CpuPeriod,omitempty"`

	// Microseconds of CPU time that the container can get in a CPU period.
	CPUQuota int64 `json:"CpuQuota,omitempty"`

	// The length of a CPU real-time period in microseconds. Set to 0 to allocate no time allocated to real-time tasks.
	CPURealtimePeriod int64 `json:"CpuRealtimePeriod,omitempty"`

	// The length of a CPU real-time runtime in microseconds. Set to 0 to allocate no time allocated to real-time tasks.
	CPURealtimeRuntime int64 `json:"CpuRealtimeRuntime,omitempty"`

	// An integer value representing this container's relative CPU weight versus other containers.
	CPUShares int64 `json:"CpuShares,omitempty"`

	// CPUs in which to allow execution (e.g., `0-3`, `0,1`)
	// Example: 0-3
	CpusetCpus string `json:"CpusetCpus,omitempty"`

	// Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.
	CpusetMems string `json:"CpusetMems,omitempty"`

	// a list of cgroup rules to apply to the container
	DeviceCgroupRules []string `json:"DeviceCgroupRules"`

	// a list of requests for devices to be sent to device drivers
	DeviceRequests []*DeviceRequest `json:"DeviceRequests"`

	// A list of devices to add to the container.
	Devices []*DeviceMapping `json:"Devices"`

	// Maximum IO in bytes per second for the container system drive (Windows only)
	IOMaximumBandwidth int64 `json:"IOMaximumBandwidth,omitempty"`

	// Maximum IOps for the container system drive (Windows only)
	IOMaximumIOps int64 `json:"IOMaximumIOps,omitempty"`

	// Run an init inside the container that forwards signals and reaps processes. This field is omitted if empty, and the default (as configured on the daemon) is used.
	Init *bool `json:"Init,omitempty"`

	// Kernel memory limit in bytes.
	// Example: 209715200
	KernelMemory int64 `json:"KernelMemory,omitempty"`

	// Hard limit for kernel TCP buffer memory (in bytes).
	KernelMemoryTCP int64 `json:"KernelMemoryTCP,omitempty"`

	// Memory limit in bytes.
	Memory int64 `json:"Memory,omitempty"`

	// Memory soft limit in bytes.
	MemoryReservation int64 `json:"MemoryReservation,omitempty"`

	// Total memory limit (memory + swap). Set as `-1` to enable unlimited swap.
	MemorySwap int64 `json:"MemorySwap,omitempty"`

	// Tune a container's memory swappiness behavior. Accepts an integer between 0 and 100.
	// Maximum: 100
	// Minimum: 0
	MemorySwappiness *int64 `json:"MemorySwappiness,omitempty"`

	// CPU quota in units of 10<sup>-9</sup> CPUs.
	NanoCPUs int64 `json:"NanoCPUs,omitempty"`

	// Disable OOM Killer for the container.
	OomKillDisable bool `json:"OomKillDisable,omitempty"`

	// Tune a container's PIDs limit. Set `0` or `-1` for unlimited, or `null` to not change.
	//
	PidsLimit *int64 `json:"PidsLimit,omitempty"`

	// A list of resource limits to set in the container. For example: `{"Name": "nofile", "Soft": 1024, "Hard": 2048}`"
	//
	Ulimits []*ResourcesUlimitsItems0 `json:"Ulimits"`
}

// Validate validates this resources
func (m *Resources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlkioDeviceReadBps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlkioDeviceReadIOps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlkioDeviceWriteBps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlkioDeviceWriteIOps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlkioWeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlkioWeightDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemorySwappiness(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUlimits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resources) validateBlkioDeviceReadBps(formats strfmt.Registry) error {
	if swag.IsZero(m.BlkioDeviceReadBps) { // not required
		return nil
	}

	for i := 0; i < len(m.BlkioDeviceReadBps); i++ {
		if swag.IsZero(m.BlkioDeviceReadBps[i]) { // not required
			continue
		}

		if m.BlkioDeviceReadBps[i] != nil {
			if err := m.BlkioDeviceReadBps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlkioDeviceReadBps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BlkioDeviceReadBps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) validateBlkioDeviceReadIOps(formats strfmt.Registry) error {
	if swag.IsZero(m.BlkioDeviceReadIOps) { // not required
		return nil
	}

	for i := 0; i < len(m.BlkioDeviceReadIOps); i++ {
		if swag.IsZero(m.BlkioDeviceReadIOps[i]) { // not required
			continue
		}

		if m.BlkioDeviceReadIOps[i] != nil {
			if err := m.BlkioDeviceReadIOps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlkioDeviceReadIOps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BlkioDeviceReadIOps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) validateBlkioDeviceWriteBps(formats strfmt.Registry) error {
	if swag.IsZero(m.BlkioDeviceWriteBps) { // not required
		return nil
	}

	for i := 0; i < len(m.BlkioDeviceWriteBps); i++ {
		if swag.IsZero(m.BlkioDeviceWriteBps[i]) { // not required
			continue
		}

		if m.BlkioDeviceWriteBps[i] != nil {
			if err := m.BlkioDeviceWriteBps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlkioDeviceWriteBps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BlkioDeviceWriteBps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) validateBlkioDeviceWriteIOps(formats strfmt.Registry) error {
	if swag.IsZero(m.BlkioDeviceWriteIOps) { // not required
		return nil
	}

	for i := 0; i < len(m.BlkioDeviceWriteIOps); i++ {
		if swag.IsZero(m.BlkioDeviceWriteIOps[i]) { // not required
			continue
		}

		if m.BlkioDeviceWriteIOps[i] != nil {
			if err := m.BlkioDeviceWriteIOps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlkioDeviceWriteIOps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BlkioDeviceWriteIOps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) validateBlkioWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.BlkioWeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("BlkioWeight", "body", *m.BlkioWeight, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("BlkioWeight", "body", *m.BlkioWeight, 1000, false); err != nil {
		return err
	}

	return nil
}

func (m *Resources) validateBlkioWeightDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.BlkioWeightDevice) { // not required
		return nil
	}

	for i := 0; i < len(m.BlkioWeightDevice); i++ {
		if swag.IsZero(m.BlkioWeightDevice[i]) { // not required
			continue
		}

		if m.BlkioWeightDevice[i] != nil {
			if err := m.BlkioWeightDevice[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlkioWeightDevice" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BlkioWeightDevice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) validateDeviceRequests(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceRequests) { // not required
		return nil
	}

	for i := 0; i < len(m.DeviceRequests); i++ {
		if swag.IsZero(m.DeviceRequests[i]) { // not required
			continue
		}

		if m.DeviceRequests[i] != nil {
			if err := m.DeviceRequests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DeviceRequests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DeviceRequests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) validateDevices(formats strfmt.Registry) error {
	if swag.IsZero(m.Devices) { // not required
		return nil
	}

	for i := 0; i < len(m.Devices); i++ {
		if swag.IsZero(m.Devices[i]) { // not required
			continue
		}

		if m.Devices[i] != nil {
			if err := m.Devices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) validateMemorySwappiness(formats strfmt.Registry) error {
	if swag.IsZero(m.MemorySwappiness) { // not required
		return nil
	}

	if err := validate.MinimumInt("MemorySwappiness", "body", *m.MemorySwappiness, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("MemorySwappiness", "body", *m.MemorySwappiness, 100, false); err != nil {
		return err
	}

	return nil
}

func (m *Resources) validateUlimits(formats strfmt.Registry) error {
	if swag.IsZero(m.Ulimits) { // not required
		return nil
	}

	for i := 0; i < len(m.Ulimits); i++ {
		if swag.IsZero(m.Ulimits[i]) { // not required
			continue
		}

		if m.Ulimits[i] != nil {
			if err := m.Ulimits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Ulimits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Ulimits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this resources based on the context it is used
func (m *Resources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBlkioDeviceReadBps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBlkioDeviceReadIOps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBlkioDeviceWriteBps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBlkioDeviceWriteIOps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBlkioWeightDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUlimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resources) contextValidateBlkioDeviceReadBps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BlkioDeviceReadBps); i++ {

		if m.BlkioDeviceReadBps[i] != nil {
			if err := m.BlkioDeviceReadBps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlkioDeviceReadBps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BlkioDeviceReadBps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) contextValidateBlkioDeviceReadIOps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BlkioDeviceReadIOps); i++ {

		if m.BlkioDeviceReadIOps[i] != nil {
			if err := m.BlkioDeviceReadIOps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlkioDeviceReadIOps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BlkioDeviceReadIOps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) contextValidateBlkioDeviceWriteBps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BlkioDeviceWriteBps); i++ {

		if m.BlkioDeviceWriteBps[i] != nil {
			if err := m.BlkioDeviceWriteBps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlkioDeviceWriteBps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BlkioDeviceWriteBps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) contextValidateBlkioDeviceWriteIOps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BlkioDeviceWriteIOps); i++ {

		if m.BlkioDeviceWriteIOps[i] != nil {
			if err := m.BlkioDeviceWriteIOps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlkioDeviceWriteIOps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BlkioDeviceWriteIOps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) contextValidateBlkioWeightDevice(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BlkioWeightDevice); i++ {

		if m.BlkioWeightDevice[i] != nil {
			if err := m.BlkioWeightDevice[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BlkioWeightDevice" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BlkioWeightDevice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) contextValidateDeviceRequests(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DeviceRequests); i++ {

		if m.DeviceRequests[i] != nil {
			if err := m.DeviceRequests[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DeviceRequests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DeviceRequests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) contextValidateDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Devices); i++ {

		if m.Devices[i] != nil {
			if err := m.Devices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Resources) contextValidateUlimits(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ulimits); i++ {

		if m.Ulimits[i] != nil {
			if err := m.Ulimits[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Ulimits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Ulimits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Resources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Resources) UnmarshalBinary(b []byte) error {
	var res Resources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResourcesBlkioWeightDeviceItems0 resources blkio weight device items0
//
// swagger:model ResourcesBlkioWeightDeviceItems0
type ResourcesBlkioWeightDeviceItems0 struct {

	// path
	Path string `json:"Path,omitempty"`

	// weight
	// Minimum: 0
	Weight *int64 `json:"Weight,omitempty"`
}

// Validate validates this resources blkio weight device items0
func (m *ResourcesBlkioWeightDeviceItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourcesBlkioWeightDeviceItems0) validateWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.Weight) { // not required
		return nil
	}

	if err := validate.MinimumInt("Weight", "body", *m.Weight, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this resources blkio weight device items0 based on context it is used
func (m *ResourcesBlkioWeightDeviceItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourcesBlkioWeightDeviceItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourcesBlkioWeightDeviceItems0) UnmarshalBinary(b []byte) error {
	var res ResourcesBlkioWeightDeviceItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResourcesUlimitsItems0 resources ulimits items0
//
// swagger:model ResourcesUlimitsItems0
type ResourcesUlimitsItems0 struct {

	// Hard limit
	Hard int64 `json:"Hard,omitempty"`

	// Name of ulimit
	Name string `json:"Name,omitempty"`

	// Soft limit
	Soft int64 `json:"Soft,omitempty"`
}

// Validate validates this resources ulimits items0
func (m *ResourcesUlimitsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this resources ulimits items0 based on context it is used
func (m *ResourcesUlimitsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourcesUlimitsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourcesUlimitsItems0) UnmarshalBinary(b []byte) error {
	var res ResourcesUlimitsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for MountPoint

// register flags to command
func registerModelMountPointFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerMountPointDestination(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountPointDriver(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountPointMode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountPointName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountPointPropagation(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountPointRW(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountPointSource(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountPointType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerMountPointDestination(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	DestinationDescription := ``

	var DestinationFlagName string
	if cmdPrefix == "" {
		DestinationFlagName = "Destination"
	} else {
		DestinationFlagName = fmt.Sprintf("%v.Destination", cmdPrefix)
	}

	var DestinationFlagDefault string

	_ = cmd.PersistentFlags().String(DestinationFlagName, DestinationFlagDefault, DestinationDescription)

	return nil
}

func registerMountPointDriver(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	DriverDescription := ``

	var DriverFlagName string
	if cmdPrefix == "" {
		DriverFlagName = "Driver"
	} else {
		DriverFlagName = fmt.Sprintf("%v.Driver", cmdPrefix)
	}

	var DriverFlagDefault string

	_ = cmd.PersistentFlags().String(DriverFlagName, DriverFlagDefault, DriverDescription)

	return nil
}

func registerMountPointMode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ModeDescription := ``

	var ModeFlagName string
	if cmdPrefix == "" {
		ModeFlagName = "Mode"
	} else {
		ModeFlagName = fmt.Sprintf("%v.Mode", cmdPrefix)
	}

	var ModeFlagDefault string

	_ = cmd.PersistentFlags().String(ModeFlagName, ModeFlagDefault, ModeDescription)

	return nil
}

func registerMountPointName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	NameDescription := ``

	var NameFlagName string
	if cmdPrefix == "" {
		NameFlagName = "Name"
	} else {
		NameFlagName = fmt.Sprintf("%v.Name", cmdPrefix)
	}

	var NameFlagDefault string

	_ = cmd.PersistentFlags().String(NameFlagName, NameFlagDefault, NameDescription)

	return nil
}

func registerMountPointPropagation(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	PropagationDescription := ``

	var PropagationFlagName string
	if cmdPrefix == "" {
		PropagationFlagName = "Propagation"
	} else {
		PropagationFlagName = fmt.Sprintf("%v.Propagation", cmdPrefix)
	}

	var PropagationFlagDefault string

	_ = cmd.PersistentFlags().String(PropagationFlagName, PropagationFlagDefault, PropagationDescription)

	return nil
}

func registerMountPointRW(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	RWDescription := ``

	var RWFlagName string
	if cmdPrefix == "" {
		RWFlagName = "RW"
	} else {
		RWFlagName = fmt.Sprintf("%v.RW", cmdPrefix)
	}

	var RWFlagDefault bool

	_ = cmd.PersistentFlags().Bool(RWFlagName, RWFlagDefault, RWDescription)

	return nil
}

func registerMountPointSource(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	SourceDescription := ``

	var SourceFlagName string
	if cmdPrefix == "" {
		SourceFlagName = "Source"
	} else {
		SourceFlagName = fmt.Sprintf("%v.Source", cmdPrefix)
	}

	var SourceFlagDefault string

	_ = cmd.PersistentFlags().String(SourceFlagName, SourceFlagDefault, SourceDescription)

	return nil
}

func registerMountPointType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	TypeDescription := ``

	var TypeFlagName string
	if cmdPrefix == "" {
		TypeFlagName = "Type"
	} else {
		TypeFlagName = fmt.Sprintf("%v.Type", cmdPrefix)
	}

	var TypeFlagDefault string

	_ = cmd.PersistentFlags().String(TypeFlagName, TypeFlagDefault, TypeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelMountPointFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, DestinationAdded := retrieveMountPointDestinationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DestinationAdded

	err, DriverAdded := retrieveMountPointDriverFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DriverAdded

	err, ModeAdded := retrieveMountPointModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ModeAdded

	err, NameAdded := retrieveMountPointNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || NameAdded

	err, PropagationAdded := retrieveMountPointPropagationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PropagationAdded

	err, RWAdded := retrieveMountPointRWFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || RWAdded

	err, SourceAdded := retrieveMountPointSourceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SourceAdded

	err, TypeAdded := retrieveMountPointTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || TypeAdded

	return nil, retAdded
}

func retrieveMountPointDestinationFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	DestinationFlagName := fmt.Sprintf("%v.Destination", cmdPrefix)
	if cmd.Flags().Changed(DestinationFlagName) {

		var DestinationFlagName string
		if cmdPrefix == "" {
			DestinationFlagName = "Destination"
		} else {
			DestinationFlagName = fmt.Sprintf("%v.Destination", cmdPrefix)
		}

		DestinationFlagValue, err := cmd.Flags().GetString(DestinationFlagName)
		if err != nil {
			return err, false
		}
		m.Destination = DestinationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountPointDriverFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	DriverFlagName := fmt.Sprintf("%v.Driver", cmdPrefix)
	if cmd.Flags().Changed(DriverFlagName) {

		var DriverFlagName string
		if cmdPrefix == "" {
			DriverFlagName = "Driver"
		} else {
			DriverFlagName = fmt.Sprintf("%v.Driver", cmdPrefix)
		}

		DriverFlagValue, err := cmd.Flags().GetString(DriverFlagName)
		if err != nil {
			return err, false
		}
		m.Driver = DriverFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountPointModeFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ModeFlagName := fmt.Sprintf("%v.Mode", cmdPrefix)
	if cmd.Flags().Changed(ModeFlagName) {

		var ModeFlagName string
		if cmdPrefix == "" {
			ModeFlagName = "Mode"
		} else {
			ModeFlagName = fmt.Sprintf("%v.Mode", cmdPrefix)
		}

		ModeFlagValue, err := cmd.Flags().GetString(ModeFlagName)
		if err != nil {
			return err, false
		}
		m.Mode = ModeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountPointNameFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	NameFlagName := fmt.Sprintf("%v.Name", cmdPrefix)
	if cmd.Flags().Changed(NameFlagName) {

		var NameFlagName string
		if cmdPrefix == "" {
			NameFlagName = "Name"
		} else {
			NameFlagName = fmt.Sprintf("%v.Name", cmdPrefix)
		}

		NameFlagValue, err := cmd.Flags().GetString(NameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = NameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountPointPropagationFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	PropagationFlagName := fmt.Sprintf("%v.Propagation", cmdPrefix)
	if cmd.Flags().Changed(PropagationFlagName) {

		var PropagationFlagName string
		if cmdPrefix == "" {
			PropagationFlagName = "Propagation"
		} else {
			PropagationFlagName = fmt.Sprintf("%v.Propagation", cmdPrefix)
		}

		PropagationFlagValue, err := cmd.Flags().GetString(PropagationFlagName)
		if err != nil {
			return err, false
		}
		m.Propagation = PropagationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountPointRWFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	RWFlagName := fmt.Sprintf("%v.RW", cmdPrefix)
	if cmd.Flags().Changed(RWFlagName) {

		var RWFlagName string
		if cmdPrefix == "" {
			RWFlagName = "RW"
		} else {
			RWFlagName = fmt.Sprintf("%v.RW", cmdPrefix)
		}

		RWFlagValue, err := cmd.Flags().GetBool(RWFlagName)
		if err != nil {
			return err, false
		}
		m.RW = RWFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountPointSourceFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	SourceFlagName := fmt.Sprintf("%v.Source", cmdPrefix)
	if cmd.Flags().Changed(SourceFlagName) {

		var SourceFlagName string
		if cmdPrefix == "" {
			SourceFlagName = "Source"
		} else {
			SourceFlagName = fmt.Sprintf("%v.Source", cmdPrefix)
		}

		SourceFlagValue, err := cmd.Flags().GetString(SourceFlagName)
		if err != nil {
			return err, false
		}
		m.Source = SourceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountPointTypeFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	TypeFlagName := fmt.Sprintf("%v.Type", cmdPrefix)
	if cmd.Flags().Changed(TypeFlagName) {

		var TypeFlagName string
		if cmdPrefix == "" {
			TypeFlagName = "Type"
		} else {
			TypeFlagName = fmt.Sprintf("%v.Type", cmdPrefix)
		}

		TypeFlagValue, err := cmd.Flags().GetString(TypeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = TypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

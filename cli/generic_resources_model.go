// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/go-swagger/dockerctl/models"

	"github.com/spf13/cobra"
)

// Schema cli for []*GenericResourcesItems0

// Name: [GenericResources], Type:[[]*GenericResourcesItems0], register and retrieve functions are not rendered by go-swagger cli

// Extra schema cli for GenericResourcesItems0

// register flags to command
func registerModelGenericResourcesItems0Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGenericResourcesItems0DiscreteResourceSpec(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGenericResourcesItems0NamedResourceSpec(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGenericResourcesItems0DiscreteResourceSpec(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var DiscreteResourceSpecFlagName string
	if cmdPrefix == "" {
		DiscreteResourceSpecFlagName = "DiscreteResourceSpec"
	} else {
		DiscreteResourceSpecFlagName = fmt.Sprintf("%v.DiscreteResourceSpec", cmdPrefix)
	}

	if err := registerModelGenericResourcesItems0DiscreteResourceSpecFlags(depth+1, DiscreteResourceSpecFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerGenericResourcesItems0NamedResourceSpec(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var NamedResourceSpecFlagName string
	if cmdPrefix == "" {
		NamedResourceSpecFlagName = "NamedResourceSpec"
	} else {
		NamedResourceSpecFlagName = fmt.Sprintf("%v.NamedResourceSpec", cmdPrefix)
	}

	if err := registerModelGenericResourcesItems0NamedResourceSpecFlags(depth+1, NamedResourceSpecFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGenericResourcesItems0Flags(depth int, m *models.GenericResourcesItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, DiscreteResourceSpecAdded := retrieveGenericResourcesItems0DiscreteResourceSpecFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DiscreteResourceSpecAdded

	err, NamedResourceSpecAdded := retrieveGenericResourcesItems0NamedResourceSpecFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || NamedResourceSpecAdded

	return nil, retAdded
}

func retrieveGenericResourcesItems0DiscreteResourceSpecFlags(depth int, m *models.GenericResourcesItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	DiscreteResourceSpecFlagName := fmt.Sprintf("%v.DiscreteResourceSpec", cmdPrefix)
	if cmd.Flags().Changed(DiscreteResourceSpecFlagName) {
		// info: complex object DiscreteResourceSpec GenericResourcesItems0DiscreteResourceSpec is retrieved outside this Changed() block
	}
	DiscreteResourceSpecFlagValue := m.DiscreteResourceSpec
	if swag.IsZero(DiscreteResourceSpecFlagValue) {
		DiscreteResourceSpecFlagValue = &models.GenericResourcesItems0DiscreteResourceSpec{}
	}

	err, DiscreteResourceSpecAdded := retrieveModelGenericResourcesItems0DiscreteResourceSpecFlags(depth+1, DiscreteResourceSpecFlagValue, DiscreteResourceSpecFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DiscreteResourceSpecAdded
	if DiscreteResourceSpecAdded {
		m.DiscreteResourceSpec = DiscreteResourceSpecFlagValue
	}

	return nil, retAdded
}

func retrieveGenericResourcesItems0NamedResourceSpecFlags(depth int, m *models.GenericResourcesItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	NamedResourceSpecFlagName := fmt.Sprintf("%v.NamedResourceSpec", cmdPrefix)
	if cmd.Flags().Changed(NamedResourceSpecFlagName) {
		// info: complex object NamedResourceSpec GenericResourcesItems0NamedResourceSpec is retrieved outside this Changed() block
	}
	NamedResourceSpecFlagValue := m.NamedResourceSpec
	if swag.IsZero(NamedResourceSpecFlagValue) {
		NamedResourceSpecFlagValue = &models.GenericResourcesItems0NamedResourceSpec{}
	}

	err, NamedResourceSpecAdded := retrieveModelGenericResourcesItems0NamedResourceSpecFlags(depth+1, NamedResourceSpecFlagValue, NamedResourceSpecFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || NamedResourceSpecAdded
	if NamedResourceSpecAdded {
		m.NamedResourceSpec = NamedResourceSpecFlagValue
	}

	return nil, retAdded
}

// Extra schema cli for GenericResourcesItems0DiscreteResourceSpec

// register flags to command
func registerModelGenericResourcesItems0DiscreteResourceSpecFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGenericResourcesItems0DiscreteResourceSpecKind(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGenericResourcesItems0DiscreteResourceSpecValue(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGenericResourcesItems0DiscreteResourceSpecKind(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	KindDescription := ``

	var KindFlagName string
	if cmdPrefix == "" {
		KindFlagName = "Kind"
	} else {
		KindFlagName = fmt.Sprintf("%v.Kind", cmdPrefix)
	}

	var KindFlagDefault string

	_ = cmd.PersistentFlags().String(KindFlagName, KindFlagDefault, KindDescription)

	return nil
}

func registerGenericResourcesItems0DiscreteResourceSpecValue(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ValueDescription := ``

	var ValueFlagName string
	if cmdPrefix == "" {
		ValueFlagName = "Value"
	} else {
		ValueFlagName = fmt.Sprintf("%v.Value", cmdPrefix)
	}

	var ValueFlagDefault int64

	_ = cmd.PersistentFlags().Int64(ValueFlagName, ValueFlagDefault, ValueDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGenericResourcesItems0DiscreteResourceSpecFlags(depth int, m *models.GenericResourcesItems0DiscreteResourceSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, KindAdded := retrieveGenericResourcesItems0DiscreteResourceSpecKindFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || KindAdded

	err, ValueAdded := retrieveGenericResourcesItems0DiscreteResourceSpecValueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ValueAdded

	return nil, retAdded
}

func retrieveGenericResourcesItems0DiscreteResourceSpecKindFlags(depth int, m *models.GenericResourcesItems0DiscreteResourceSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	KindFlagName := fmt.Sprintf("%v.Kind", cmdPrefix)
	if cmd.Flags().Changed(KindFlagName) {

		var KindFlagName string
		if cmdPrefix == "" {
			KindFlagName = "Kind"
		} else {
			KindFlagName = fmt.Sprintf("%v.Kind", cmdPrefix)
		}

		KindFlagValue, err := cmd.Flags().GetString(KindFlagName)
		if err != nil {
			return err, false
		}
		m.Kind = KindFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveGenericResourcesItems0DiscreteResourceSpecValueFlags(depth int, m *models.GenericResourcesItems0DiscreteResourceSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ValueFlagName := fmt.Sprintf("%v.Value", cmdPrefix)
	if cmd.Flags().Changed(ValueFlagName) {

		var ValueFlagName string
		if cmdPrefix == "" {
			ValueFlagName = "Value"
		} else {
			ValueFlagName = fmt.Sprintf("%v.Value", cmdPrefix)
		}

		ValueFlagValue, err := cmd.Flags().GetInt64(ValueFlagName)
		if err != nil {
			return err, false
		}
		m.Value = ValueFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for GenericResourcesItems0NamedResourceSpec

// register flags to command
func registerModelGenericResourcesItems0NamedResourceSpecFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGenericResourcesItems0NamedResourceSpecKind(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGenericResourcesItems0NamedResourceSpecValue(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGenericResourcesItems0NamedResourceSpecKind(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	KindDescription := ``

	var KindFlagName string
	if cmdPrefix == "" {
		KindFlagName = "Kind"
	} else {
		KindFlagName = fmt.Sprintf("%v.Kind", cmdPrefix)
	}

	var KindFlagDefault string

	_ = cmd.PersistentFlags().String(KindFlagName, KindFlagDefault, KindDescription)

	return nil
}

func registerGenericResourcesItems0NamedResourceSpecValue(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ValueDescription := ``

	var ValueFlagName string
	if cmdPrefix == "" {
		ValueFlagName = "Value"
	} else {
		ValueFlagName = fmt.Sprintf("%v.Value", cmdPrefix)
	}

	var ValueFlagDefault string

	_ = cmd.PersistentFlags().String(ValueFlagName, ValueFlagDefault, ValueDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGenericResourcesItems0NamedResourceSpecFlags(depth int, m *models.GenericResourcesItems0NamedResourceSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, KindAdded := retrieveGenericResourcesItems0NamedResourceSpecKindFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || KindAdded

	err, ValueAdded := retrieveGenericResourcesItems0NamedResourceSpecValueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ValueAdded

	return nil, retAdded
}

func retrieveGenericResourcesItems0NamedResourceSpecKindFlags(depth int, m *models.GenericResourcesItems0NamedResourceSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	KindFlagName := fmt.Sprintf("%v.Kind", cmdPrefix)
	if cmd.Flags().Changed(KindFlagName) {

		var KindFlagName string
		if cmdPrefix == "" {
			KindFlagName = "Kind"
		} else {
			KindFlagName = fmt.Sprintf("%v.Kind", cmdPrefix)
		}

		KindFlagValue, err := cmd.Flags().GetString(KindFlagName)
		if err != nil {
			return err, false
		}
		m.Kind = KindFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveGenericResourcesItems0NamedResourceSpecValueFlags(depth int, m *models.GenericResourcesItems0NamedResourceSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ValueFlagName := fmt.Sprintf("%v.Value", cmdPrefix)
	if cmd.Flags().Changed(ValueFlagName) {

		var ValueFlagName string
		if cmdPrefix == "" {
			ValueFlagName = "Value"
		} else {
			ValueFlagName = fmt.Sprintf("%v.Value", cmdPrefix)
		}

		ValueFlagValue, err := cmd.Flags().GetString(ValueFlagName)
		if err != nil {
			return err, false
		}
		m.Value = ValueFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/go-swagger/dockerctl/models"

	"github.com/spf13/cobra"
)

// Schema cli for Service

// register flags to command
func registerModelServiceFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServiceCreatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceEndpoint(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceServiceStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceSpec(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceUpdateStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceUpdatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceCreatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	CreatedAtDescription := ``

	var CreatedAtFlagName string
	if cmdPrefix == "" {
		CreatedAtFlagName = "CreatedAt"
	} else {
		CreatedAtFlagName = fmt.Sprintf("%v.CreatedAt", cmdPrefix)
	}

	var CreatedAtFlagDefault string

	_ = cmd.PersistentFlags().String(CreatedAtFlagName, CreatedAtFlagDefault, CreatedAtDescription)

	return nil
}

func registerServiceEndpoint(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var EndpointFlagName string
	if cmdPrefix == "" {
		EndpointFlagName = "Endpoint"
	} else {
		EndpointFlagName = fmt.Sprintf("%v.Endpoint", cmdPrefix)
	}

	if err := registerModelServiceEndpointFlags(depth+1, EndpointFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	IDDescription := ``

	var IDFlagName string
	if cmdPrefix == "" {
		IDFlagName = "ID"
	} else {
		IDFlagName = fmt.Sprintf("%v.ID", cmdPrefix)
	}

	var IDFlagDefault string

	_ = cmd.PersistentFlags().String(IDFlagName, IDFlagDefault, IDDescription)

	return nil
}

func registerServiceServiceStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var ServiceStatusFlagName string
	if cmdPrefix == "" {
		ServiceStatusFlagName = "ServiceStatus"
	} else {
		ServiceStatusFlagName = fmt.Sprintf("%v.ServiceStatus", cmdPrefix)
	}

	if err := registerModelServiceServiceStatusFlags(depth+1, ServiceStatusFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceSpec(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var SpecFlagName string
	if cmdPrefix == "" {
		SpecFlagName = "Spec"
	} else {
		SpecFlagName = fmt.Sprintf("%v.Spec", cmdPrefix)
	}

	if err := registerModelServiceSpecFlags(depth+1, SpecFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceUpdateStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var UpdateStatusFlagName string
	if cmdPrefix == "" {
		UpdateStatusFlagName = "UpdateStatus"
	} else {
		UpdateStatusFlagName = fmt.Sprintf("%v.UpdateStatus", cmdPrefix)
	}

	if err := registerModelServiceUpdateStatusFlags(depth+1, UpdateStatusFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceUpdatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	UpdatedAtDescription := ``

	var UpdatedAtFlagName string
	if cmdPrefix == "" {
		UpdatedAtFlagName = "UpdatedAt"
	} else {
		UpdatedAtFlagName = fmt.Sprintf("%v.UpdatedAt", cmdPrefix)
	}

	var UpdatedAtFlagDefault string

	_ = cmd.PersistentFlags().String(UpdatedAtFlagName, UpdatedAtFlagDefault, UpdatedAtDescription)

	return nil
}

func registerServiceVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var VersionFlagName string
	if cmdPrefix == "" {
		VersionFlagName = "Version"
	} else {
		VersionFlagName = fmt.Sprintf("%v.Version", cmdPrefix)
	}

	if err := registerModelObjectVersionFlags(depth+1, VersionFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServiceFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, CreatedAtAdded := retrieveServiceCreatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || CreatedAtAdded

	err, EndpointAdded := retrieveServiceEndpointFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || EndpointAdded

	err, IDAdded := retrieveServiceIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || IDAdded

	err, ServiceStatusAdded := retrieveServiceServiceStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ServiceStatusAdded

	err, SpecAdded := retrieveServiceSpecFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SpecAdded

	err, UpdateStatusAdded := retrieveServiceUpdateStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || UpdateStatusAdded

	err, UpdatedAtAdded := retrieveServiceUpdatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || UpdatedAtAdded

	err, VersionAdded := retrieveServiceVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || VersionAdded

	return nil, retAdded
}

func retrieveServiceCreatedAtFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	CreatedAtFlagName := fmt.Sprintf("%v.CreatedAt", cmdPrefix)
	if cmd.Flags().Changed(CreatedAtFlagName) {

		var CreatedAtFlagName string
		if cmdPrefix == "" {
			CreatedAtFlagName = "CreatedAt"
		} else {
			CreatedAtFlagName = fmt.Sprintf("%v.CreatedAt", cmdPrefix)
		}

		CreatedAtFlagValue, err := cmd.Flags().GetString(CreatedAtFlagName)
		if err != nil {
			return err, false
		}
		m.CreatedAt = CreatedAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceEndpointFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	EndpointFlagName := fmt.Sprintf("%v.Endpoint", cmdPrefix)
	if cmd.Flags().Changed(EndpointFlagName) {
		// info: complex object Endpoint ServiceEndpoint is retrieved outside this Changed() block
	}
	EndpointFlagValue := m.Endpoint
	if swag.IsZero(EndpointFlagValue) {
		EndpointFlagValue = &models.ServiceEndpoint{}
	}

	err, EndpointAdded := retrieveModelServiceEndpointFlags(depth+1, EndpointFlagValue, EndpointFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || EndpointAdded
	if EndpointAdded {
		m.Endpoint = EndpointFlagValue
	}

	return nil, retAdded
}

func retrieveServiceIDFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	IDFlagName := fmt.Sprintf("%v.ID", cmdPrefix)
	if cmd.Flags().Changed(IDFlagName) {

		var IDFlagName string
		if cmdPrefix == "" {
			IDFlagName = "ID"
		} else {
			IDFlagName = fmt.Sprintf("%v.ID", cmdPrefix)
		}

		IDFlagValue, err := cmd.Flags().GetString(IDFlagName)
		if err != nil {
			return err, false
		}
		m.ID = IDFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceServiceStatusFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ServiceStatusFlagName := fmt.Sprintf("%v.ServiceStatus", cmdPrefix)
	if cmd.Flags().Changed(ServiceStatusFlagName) {
		// info: complex object ServiceStatus ServiceServiceStatus is retrieved outside this Changed() block
	}
	ServiceStatusFlagValue := m.ServiceStatus
	if swag.IsZero(ServiceStatusFlagValue) {
		ServiceStatusFlagValue = &models.ServiceServiceStatus{}
	}

	err, ServiceStatusAdded := retrieveModelServiceServiceStatusFlags(depth+1, ServiceStatusFlagValue, ServiceStatusFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ServiceStatusAdded
	if ServiceStatusAdded {
		m.ServiceStatus = ServiceStatusFlagValue
	}

	return nil, retAdded
}

func retrieveServiceSpecFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	SpecFlagName := fmt.Sprintf("%v.Spec", cmdPrefix)
	if cmd.Flags().Changed(SpecFlagName) {
		// info: complex object Spec ServiceSpec is retrieved outside this Changed() block
	}
	SpecFlagValue := m.Spec
	if swag.IsZero(SpecFlagValue) {
		SpecFlagValue = &models.ServiceSpec{}
	}

	err, SpecAdded := retrieveModelServiceSpecFlags(depth+1, SpecFlagValue, SpecFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SpecAdded
	if SpecAdded {
		m.Spec = SpecFlagValue
	}

	return nil, retAdded
}

func retrieveServiceUpdateStatusFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	UpdateStatusFlagName := fmt.Sprintf("%v.UpdateStatus", cmdPrefix)
	if cmd.Flags().Changed(UpdateStatusFlagName) {
		// info: complex object UpdateStatus ServiceUpdateStatus is retrieved outside this Changed() block
	}
	UpdateStatusFlagValue := m.UpdateStatus
	if swag.IsZero(UpdateStatusFlagValue) {
		UpdateStatusFlagValue = &models.ServiceUpdateStatus{}
	}

	err, UpdateStatusAdded := retrieveModelServiceUpdateStatusFlags(depth+1, UpdateStatusFlagValue, UpdateStatusFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || UpdateStatusAdded
	if UpdateStatusAdded {
		m.UpdateStatus = UpdateStatusFlagValue
	}

	return nil, retAdded
}

func retrieveServiceUpdatedAtFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	UpdatedAtFlagName := fmt.Sprintf("%v.UpdatedAt", cmdPrefix)
	if cmd.Flags().Changed(UpdatedAtFlagName) {

		var UpdatedAtFlagName string
		if cmdPrefix == "" {
			UpdatedAtFlagName = "UpdatedAt"
		} else {
			UpdatedAtFlagName = fmt.Sprintf("%v.UpdatedAt", cmdPrefix)
		}

		UpdatedAtFlagValue, err := cmd.Flags().GetString(UpdatedAtFlagName)
		if err != nil {
			return err, false
		}
		m.UpdatedAt = UpdatedAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceVersionFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	VersionFlagName := fmt.Sprintf("%v.Version", cmdPrefix)
	if cmd.Flags().Changed(VersionFlagName) {
		// info: complex object Version ObjectVersion is retrieved outside this Changed() block
	}
	VersionFlagValue := m.Version
	if swag.IsZero(VersionFlagValue) {
		VersionFlagValue = &models.ObjectVersion{}
	}

	err, VersionAdded := retrieveModelObjectVersionFlags(depth+1, VersionFlagValue, VersionFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || VersionAdded
	if VersionAdded {
		m.Version = VersionFlagValue
	}

	return nil, retAdded
}

// Extra schema cli for ServiceEndpoint

// register flags to command
func registerModelServiceEndpointFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServiceEndpointPorts(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceEndpointSpec(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceEndpointVirtualIPs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceEndpointPorts(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Ports []*EndpointPortConfig array type is not supported by go-swagger cli yet

	return nil
}

func registerServiceEndpointSpec(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var SpecFlagName string
	if cmdPrefix == "" {
		SpecFlagName = "Spec"
	} else {
		SpecFlagName = fmt.Sprintf("%v.Spec", cmdPrefix)
	}

	if err := registerModelEndpointSpecFlags(depth+1, SpecFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceEndpointVirtualIPs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: VirtualIPs []*ServiceEndpointVirtualIPsItems0 array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServiceEndpointFlags(depth int, m *models.ServiceEndpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, PortsAdded := retrieveServiceEndpointPortsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PortsAdded

	err, SpecAdded := retrieveServiceEndpointSpecFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SpecAdded

	err, VirtualIPsAdded := retrieveServiceEndpointVirtualIPsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || VirtualIPsAdded

	return nil, retAdded
}

func retrieveServiceEndpointPortsFlags(depth int, m *models.ServiceEndpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	PortsFlagName := fmt.Sprintf("%v.Ports", cmdPrefix)
	if cmd.Flags().Changed(PortsFlagName) {
		// warning: Ports array type []*EndpointPortConfig is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveServiceEndpointSpecFlags(depth int, m *models.ServiceEndpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	SpecFlagName := fmt.Sprintf("%v.Spec", cmdPrefix)
	if cmd.Flags().Changed(SpecFlagName) {
		// info: complex object Spec EndpointSpec is retrieved outside this Changed() block
	}
	SpecFlagValue := m.Spec
	if swag.IsZero(SpecFlagValue) {
		SpecFlagValue = &models.EndpointSpec{}
	}

	err, SpecAdded := retrieveModelEndpointSpecFlags(depth+1, SpecFlagValue, SpecFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SpecAdded
	if SpecAdded {
		m.Spec = SpecFlagValue
	}

	return nil, retAdded
}

func retrieveServiceEndpointVirtualIPsFlags(depth int, m *models.ServiceEndpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	VirtualIPsFlagName := fmt.Sprintf("%v.VirtualIPs", cmdPrefix)
	if cmd.Flags().Changed(VirtualIPsFlagName) {
		// warning: VirtualIPs array type []*ServiceEndpointVirtualIPsItems0 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

// Extra schema cli for ServiceEndpointVirtualIPsItems0

// register flags to command
func registerModelServiceEndpointVirtualIPsItems0Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServiceEndpointVirtualIPsItems0Addr(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceEndpointVirtualIPsItems0NetworkID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceEndpointVirtualIPsItems0Addr(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	AddrDescription := ``

	var AddrFlagName string
	if cmdPrefix == "" {
		AddrFlagName = "Addr"
	} else {
		AddrFlagName = fmt.Sprintf("%v.Addr", cmdPrefix)
	}

	var AddrFlagDefault string

	_ = cmd.PersistentFlags().String(AddrFlagName, AddrFlagDefault, AddrDescription)

	return nil
}

func registerServiceEndpointVirtualIPsItems0NetworkID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	NetworkIDDescription := ``

	var NetworkIDFlagName string
	if cmdPrefix == "" {
		NetworkIDFlagName = "NetworkID"
	} else {
		NetworkIDFlagName = fmt.Sprintf("%v.NetworkID", cmdPrefix)
	}

	var NetworkIDFlagDefault string

	_ = cmd.PersistentFlags().String(NetworkIDFlagName, NetworkIDFlagDefault, NetworkIDDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServiceEndpointVirtualIPsItems0Flags(depth int, m *models.ServiceEndpointVirtualIPsItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, AddrAdded := retrieveServiceEndpointVirtualIPsItems0AddrFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || AddrAdded

	err, NetworkIDAdded := retrieveServiceEndpointVirtualIPsItems0NetworkIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || NetworkIDAdded

	return nil, retAdded
}

func retrieveServiceEndpointVirtualIPsItems0AddrFlags(depth int, m *models.ServiceEndpointVirtualIPsItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	AddrFlagName := fmt.Sprintf("%v.Addr", cmdPrefix)
	if cmd.Flags().Changed(AddrFlagName) {

		var AddrFlagName string
		if cmdPrefix == "" {
			AddrFlagName = "Addr"
		} else {
			AddrFlagName = fmt.Sprintf("%v.Addr", cmdPrefix)
		}

		AddrFlagValue, err := cmd.Flags().GetString(AddrFlagName)
		if err != nil {
			return err, false
		}
		m.Addr = AddrFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceEndpointVirtualIPsItems0NetworkIDFlags(depth int, m *models.ServiceEndpointVirtualIPsItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	NetworkIDFlagName := fmt.Sprintf("%v.NetworkID", cmdPrefix)
	if cmd.Flags().Changed(NetworkIDFlagName) {

		var NetworkIDFlagName string
		if cmdPrefix == "" {
			NetworkIDFlagName = "NetworkID"
		} else {
			NetworkIDFlagName = fmt.Sprintf("%v.NetworkID", cmdPrefix)
		}

		NetworkIDFlagValue, err := cmd.Flags().GetString(NetworkIDFlagName)
		if err != nil {
			return err, false
		}
		m.NetworkID = NetworkIDFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for ServiceServiceStatus

// register flags to command
func registerModelServiceServiceStatusFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServiceServiceStatusDesiredTasks(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceServiceStatusRunningTasks(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceServiceStatusDesiredTasks(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primitive DesiredTasks uint64 is not supported by go-swagger cli yet

	return nil
}

func registerServiceServiceStatusRunningTasks(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primitive RunningTasks uint64 is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServiceServiceStatusFlags(depth int, m *models.ServiceServiceStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, DesiredTasksAdded := retrieveServiceServiceStatusDesiredTasksFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DesiredTasksAdded

	err, RunningTasksAdded := retrieveServiceServiceStatusRunningTasksFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || RunningTasksAdded

	return nil, retAdded
}

func retrieveServiceServiceStatusDesiredTasksFlags(depth int, m *models.ServiceServiceStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	DesiredTasksFlagName := fmt.Sprintf("%v.DesiredTasks", cmdPrefix)
	if cmd.Flags().Changed(DesiredTasksFlagName) {

		// warning: primitive DesiredTasks uint64 is not supported by go-swagger cli yet

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceServiceStatusRunningTasksFlags(depth int, m *models.ServiceServiceStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	RunningTasksFlagName := fmt.Sprintf("%v.RunningTasks", cmdPrefix)
	if cmd.Flags().Changed(RunningTasksFlagName) {

		// warning: primitive RunningTasks uint64 is not supported by go-swagger cli yet

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for ServiceUpdateStatus

// register flags to command
func registerModelServiceUpdateStatusFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServiceUpdateStatusCompletedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceUpdateStatusMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceUpdateStatusStartedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceUpdateStatusState(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceUpdateStatusCompletedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	CompletedAtDescription := ``

	var CompletedAtFlagName string
	if cmdPrefix == "" {
		CompletedAtFlagName = "CompletedAt"
	} else {
		CompletedAtFlagName = fmt.Sprintf("%v.CompletedAt", cmdPrefix)
	}

	var CompletedAtFlagDefault string

	_ = cmd.PersistentFlags().String(CompletedAtFlagName, CompletedAtFlagDefault, CompletedAtDescription)

	return nil
}

func registerServiceUpdateStatusMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	MessageDescription := ``

	var MessageFlagName string
	if cmdPrefix == "" {
		MessageFlagName = "Message"
	} else {
		MessageFlagName = fmt.Sprintf("%v.Message", cmdPrefix)
	}

	var MessageFlagDefault string

	_ = cmd.PersistentFlags().String(MessageFlagName, MessageFlagDefault, MessageDescription)

	return nil
}

func registerServiceUpdateStatusStartedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	StartedAtDescription := ``

	var StartedAtFlagName string
	if cmdPrefix == "" {
		StartedAtFlagName = "StartedAt"
	} else {
		StartedAtFlagName = fmt.Sprintf("%v.StartedAt", cmdPrefix)
	}

	var StartedAtFlagDefault string

	_ = cmd.PersistentFlags().String(StartedAtFlagName, StartedAtFlagDefault, StartedAtDescription)

	return nil
}

func registerServiceUpdateStatusState(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	StateDescription := `Enum: ["updating","paused","completed"]. `

	var StateFlagName string
	if cmdPrefix == "" {
		StateFlagName = "State"
	} else {
		StateFlagName = fmt.Sprintf("%v.State", cmdPrefix)
	}

	var StateFlagDefault string

	_ = cmd.PersistentFlags().String(StateFlagName, StateFlagDefault, StateDescription)

	if err := cmd.RegisterFlagCompletionFunc(StateFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["updating","paused","completed"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServiceUpdateStatusFlags(depth int, m *models.ServiceUpdateStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, CompletedAtAdded := retrieveServiceUpdateStatusCompletedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || CompletedAtAdded

	err, MessageAdded := retrieveServiceUpdateStatusMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || MessageAdded

	err, StartedAtAdded := retrieveServiceUpdateStatusStartedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || StartedAtAdded

	err, StateAdded := retrieveServiceUpdateStatusStateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || StateAdded

	return nil, retAdded
}

func retrieveServiceUpdateStatusCompletedAtFlags(depth int, m *models.ServiceUpdateStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	CompletedAtFlagName := fmt.Sprintf("%v.CompletedAt", cmdPrefix)
	if cmd.Flags().Changed(CompletedAtFlagName) {

		var CompletedAtFlagName string
		if cmdPrefix == "" {
			CompletedAtFlagName = "CompletedAt"
		} else {
			CompletedAtFlagName = fmt.Sprintf("%v.CompletedAt", cmdPrefix)
		}

		CompletedAtFlagValue, err := cmd.Flags().GetString(CompletedAtFlagName)
		if err != nil {
			return err, false
		}
		m.CompletedAt = CompletedAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceUpdateStatusMessageFlags(depth int, m *models.ServiceUpdateStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	MessageFlagName := fmt.Sprintf("%v.Message", cmdPrefix)
	if cmd.Flags().Changed(MessageFlagName) {

		var MessageFlagName string
		if cmdPrefix == "" {
			MessageFlagName = "Message"
		} else {
			MessageFlagName = fmt.Sprintf("%v.Message", cmdPrefix)
		}

		MessageFlagValue, err := cmd.Flags().GetString(MessageFlagName)
		if err != nil {
			return err, false
		}
		m.Message = MessageFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceUpdateStatusStartedAtFlags(depth int, m *models.ServiceUpdateStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	StartedAtFlagName := fmt.Sprintf("%v.StartedAt", cmdPrefix)
	if cmd.Flags().Changed(StartedAtFlagName) {

		var StartedAtFlagName string
		if cmdPrefix == "" {
			StartedAtFlagName = "StartedAt"
		} else {
			StartedAtFlagName = fmt.Sprintf("%v.StartedAt", cmdPrefix)
		}

		StartedAtFlagValue, err := cmd.Flags().GetString(StartedAtFlagName)
		if err != nil {
			return err, false
		}
		m.StartedAt = StartedAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceUpdateStatusStateFlags(depth int, m *models.ServiceUpdateStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	StateFlagName := fmt.Sprintf("%v.State", cmdPrefix)
	if cmd.Flags().Changed(StateFlagName) {

		var StateFlagName string
		if cmdPrefix == "" {
			StateFlagName = "State"
		} else {
			StateFlagName = fmt.Sprintf("%v.State", cmdPrefix)
		}

		StateFlagValue, err := cmd.Flags().GetString(StateFlagName)
		if err != nil {
			return err, false
		}
		m.State = StateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

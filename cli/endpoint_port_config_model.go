// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for EndpointPortConfig

// register flags to command
func registerModelEndpointPortConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEndpointPortConfigName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointPortConfigProtocol(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointPortConfigPublishMode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointPortConfigPublishedPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointPortConfigTargetPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEndpointPortConfigName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerEndpointPortConfigProtocol(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ProtocolDescription := `Enum: ["tcp","udp","sctp"]. `

	var ProtocolFlagName string
	if cmdPrefix == "" {
		ProtocolFlagName = "Protocol"
	} else {
		ProtocolFlagName = fmt.Sprintf("%v.Protocol", cmdPrefix)
	}

	var ProtocolFlagDefault string

	_ = cmd.PersistentFlags().String(ProtocolFlagName, ProtocolFlagDefault, ProtocolDescription)

	if err := cmd.RegisterFlagCompletionFunc(ProtocolFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["tcp","udp","sctp"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerEndpointPortConfigPublishMode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	PublishModeDescription := `Enum: ["ingress","host"]. The mode in which port is published.

<p><br /></p>

- "ingress" makes the target port accessible on every node,
  regardless of whether there is a task for the service running on
  that node or not.
- "host" bypasses the routing mesh and publish the port directly on
  the swarm node where that service is running.
`

	var PublishModeFlagName string
	if cmdPrefix == "" {
		PublishModeFlagName = "PublishMode"
	} else {
		PublishModeFlagName = fmt.Sprintf("%v.PublishMode", cmdPrefix)
	}

	var PublishModeFlagDefault string = "ingress"

	_ = cmd.PersistentFlags().String(PublishModeFlagName, PublishModeFlagDefault, PublishModeDescription)

	if err := cmd.RegisterFlagCompletionFunc(PublishModeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["ingress","host"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerEndpointPortConfigPublishedPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	PublishedPortDescription := `The port on the swarm hosts.`

	var PublishedPortFlagName string
	if cmdPrefix == "" {
		PublishedPortFlagName = "PublishedPort"
	} else {
		PublishedPortFlagName = fmt.Sprintf("%v.PublishedPort", cmdPrefix)
	}

	var PublishedPortFlagDefault int64

	_ = cmd.PersistentFlags().Int64(PublishedPortFlagName, PublishedPortFlagDefault, PublishedPortDescription)

	return nil
}

func registerEndpointPortConfigTargetPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	TargetPortDescription := `The port inside the container.`

	var TargetPortFlagName string
	if cmdPrefix == "" {
		TargetPortFlagName = "TargetPort"
	} else {
		TargetPortFlagName = fmt.Sprintf("%v.TargetPort", cmdPrefix)
	}

	var TargetPortFlagDefault int64

	_ = cmd.PersistentFlags().Int64(TargetPortFlagName, TargetPortFlagDefault, TargetPortDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEndpointPortConfigFlags(depth int, m *models.EndpointPortConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, NameAdded := retrieveEndpointPortConfigNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || NameAdded

	err, ProtocolAdded := retrieveEndpointPortConfigProtocolFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ProtocolAdded

	err, PublishModeAdded := retrieveEndpointPortConfigPublishModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PublishModeAdded

	err, PublishedPortAdded := retrieveEndpointPortConfigPublishedPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PublishedPortAdded

	err, TargetPortAdded := retrieveEndpointPortConfigTargetPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || TargetPortAdded

	return nil, retAdded
}

func retrieveEndpointPortConfigNameFlags(depth int, m *models.EndpointPortConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveEndpointPortConfigProtocolFlags(depth int, m *models.EndpointPortConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ProtocolFlagName := fmt.Sprintf("%v.Protocol", cmdPrefix)
	if cmd.Flags().Changed(ProtocolFlagName) {

		var ProtocolFlagName string
		if cmdPrefix == "" {
			ProtocolFlagName = "Protocol"
		} else {
			ProtocolFlagName = fmt.Sprintf("%v.Protocol", cmdPrefix)
		}

		ProtocolFlagValue, err := cmd.Flags().GetString(ProtocolFlagName)
		if err != nil {
			return err, false
		}
		m.Protocol = ProtocolFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEndpointPortConfigPublishModeFlags(depth int, m *models.EndpointPortConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	PublishModeFlagName := fmt.Sprintf("%v.PublishMode", cmdPrefix)
	if cmd.Flags().Changed(PublishModeFlagName) {

		var PublishModeFlagName string
		if cmdPrefix == "" {
			PublishModeFlagName = "PublishMode"
		} else {
			PublishModeFlagName = fmt.Sprintf("%v.PublishMode", cmdPrefix)
		}

		PublishModeFlagValue, err := cmd.Flags().GetString(PublishModeFlagName)
		if err != nil {
			return err, false
		}
		m.PublishMode = &PublishModeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEndpointPortConfigPublishedPortFlags(depth int, m *models.EndpointPortConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	PublishedPortFlagName := fmt.Sprintf("%v.PublishedPort", cmdPrefix)
	if cmd.Flags().Changed(PublishedPortFlagName) {

		var PublishedPortFlagName string
		if cmdPrefix == "" {
			PublishedPortFlagName = "PublishedPort"
		} else {
			PublishedPortFlagName = fmt.Sprintf("%v.PublishedPort", cmdPrefix)
		}

		PublishedPortFlagValue, err := cmd.Flags().GetInt64(PublishedPortFlagName)
		if err != nil {
			return err, false
		}
		m.PublishedPort = PublishedPortFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEndpointPortConfigTargetPortFlags(depth int, m *models.EndpointPortConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	TargetPortFlagName := fmt.Sprintf("%v.TargetPort", cmdPrefix)
	if cmd.Flags().Changed(TargetPortFlagName) {

		var TargetPortFlagName string
		if cmdPrefix == "" {
			TargetPortFlagName = "TargetPort"
		} else {
			TargetPortFlagName = fmt.Sprintf("%v.TargetPort", cmdPrefix)
		}

		TargetPortFlagValue, err := cmd.Flags().GetInt64(TargetPortFlagName)
		if err != nil {
			return err, false
		}
		m.TargetPort = TargetPortFlagValue

		retAdded = true
	}

	return nil, retAdded
}

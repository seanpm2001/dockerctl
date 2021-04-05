// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// register flags to command
func registerModelEndpointIPAMConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEndpointIPAMConfigIPV4Address(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointIPAMConfigIPV6Address(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointIPAMConfigLinkLocalIPs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEndpointIPAMConfigIPV4Address(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipv4AddressDescription := ``

	var ipv4AddressFlagName string
	if cmdPrefix == "" {
		ipv4AddressFlagName = "IPv4Address"
	} else {
		ipv4AddressFlagName = fmt.Sprintf("%v.IPv4Address", cmdPrefix)
	}

	var ipv4AddressFlagDefault string

	_ = cmd.PersistentFlags().String(ipv4AddressFlagName, ipv4AddressFlagDefault, ipv4AddressDescription)

	return nil
}

func registerEndpointIPAMConfigIPV6Address(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipv6AddressDescription := ``

	var ipv6AddressFlagName string
	if cmdPrefix == "" {
		ipv6AddressFlagName = "IPv6Address"
	} else {
		ipv6AddressFlagName = fmt.Sprintf("%v.IPv6Address", cmdPrefix)
	}

	var ipv6AddressFlagDefault string

	_ = cmd.PersistentFlags().String(ipv6AddressFlagName, ipv6AddressFlagDefault, ipv6AddressDescription)

	return nil
}

func registerEndpointIPAMConfigLinkLocalIPs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: LinkLocalIPs []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEndpointIPAMConfigFlags(depth int, m *models.EndpointIPAMConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, ipv4AddressAdded := retrieveEndpointIPAMConfigIPV4AddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipv4AddressAdded

	err, ipv6AddressAdded := retrieveEndpointIPAMConfigIPV6AddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipv6AddressAdded

	err, linkLocalIPsAdded := retrieveEndpointIPAMConfigLinkLocalIPsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || linkLocalIPsAdded

	return nil, retAdded
}

func retrieveEndpointIPAMConfigIPV4AddressFlags(depth int, m *models.EndpointIPAMConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	ipv4AddressFlagName := fmt.Sprintf("%v.IPv4Address", cmdPrefix)
	if cmd.Flags().Changed(ipv4AddressFlagName) {

		var ipv4AddressFlagName string
		if cmdPrefix == "" {
			ipv4AddressFlagName = "IPv4Address"
		} else {
			ipv4AddressFlagName = fmt.Sprintf("%v.IPv4Address", cmdPrefix)
		}

		ipv4AddressFlagValue, err := cmd.Flags().GetString(ipv4AddressFlagName)
		if err != nil {
			return err, false
		}
		m.IPV4Address = ipv4AddressFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveEndpointIPAMConfigIPV6AddressFlags(depth int, m *models.EndpointIPAMConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	ipv6AddressFlagName := fmt.Sprintf("%v.IPv6Address", cmdPrefix)
	if cmd.Flags().Changed(ipv6AddressFlagName) {

		var ipv6AddressFlagName string
		if cmdPrefix == "" {
			ipv6AddressFlagName = "IPv6Address"
		} else {
			ipv6AddressFlagName = fmt.Sprintf("%v.IPv6Address", cmdPrefix)
		}

		ipv6AddressFlagValue, err := cmd.Flags().GetString(ipv6AddressFlagName)
		if err != nil {
			return err, false
		}
		m.IPV6Address = ipv6AddressFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveEndpointIPAMConfigLinkLocalIPsFlags(depth int, m *models.EndpointIPAMConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	linkLocalIPsFlagName := fmt.Sprintf("%v.LinkLocalIPs", cmdPrefix)
	if cmd.Flags().Changed(linkLocalIPsFlagName) {
		// warning: LinkLocalIPs array type []string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

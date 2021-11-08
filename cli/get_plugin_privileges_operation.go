// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/plugin"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationPluginGetPluginPrivilegesCmd returns a cmd to handle operation getPluginPrivileges
func makeOperationPluginGetPluginPrivilegesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "GetPluginPrivileges",
		Short: ``,
		RunE:  runOperationPluginGetPluginPrivileges,
	}

	if err := registerOperationPluginGetPluginPrivilegesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPluginGetPluginPrivileges uses cmd flags to call endpoint api
func runOperationPluginGetPluginPrivileges(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := plugin.NewGetPluginPrivilegesParams()
	if err, _ := retrieveOperationPluginGetPluginPrivilegesRemoteFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationPluginGetPluginPrivilegesResult(appCli.Plugin.GetPluginPrivileges(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationPluginGetPluginPrivilegesParamFlags registers all flags needed to fill params
func registerOperationPluginGetPluginPrivilegesParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPluginGetPluginPrivilegesRemoteParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPluginGetPluginPrivilegesRemoteParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	remoteDescription := `Required. The name of the plugin. The ` + "`" + `:latest` + "`" + ` tag is optional, and is the default if omitted.`

	var remoteFlagName string
	if cmdPrefix == "" {
		remoteFlagName = "remote"
	} else {
		remoteFlagName = fmt.Sprintf("%v.remote", cmdPrefix)
	}

	var remoteFlagDefault string

	_ = cmd.PersistentFlags().String(remoteFlagName, remoteFlagDefault, remoteDescription)

	return nil
}

func retrieveOperationPluginGetPluginPrivilegesRemoteFlag(m *plugin.GetPluginPrivilegesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("remote") {

		var remoteFlagName string
		if cmdPrefix == "" {
			remoteFlagName = "remote"
		} else {
			remoteFlagName = fmt.Sprintf("%v.remote", cmdPrefix)
		}

		remoteFlagValue, err := cmd.Flags().GetString(remoteFlagName)
		if err != nil {
			return err, false
		}
		m.Remote = remoteFlagValue

	}
	return nil, retAdded
}

// parseOperationPluginGetPluginPrivilegesResult parses request result and return the string content
func parseOperationPluginGetPluginPrivilegesResult(resp0 *plugin.GetPluginPrivilegesOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*plugin.GetPluginPrivilegesOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*plugin.GetPluginPrivilegesInternalServerError)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}

// register flags to command
func registerModelGetPluginPrivilegesOKBodyItems0Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGetPluginPrivilegesOKBodyItems0Description(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGetPluginPrivilegesOKBodyItems0Name(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGetPluginPrivilegesOKBodyItems0Value(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGetPluginPrivilegesOKBodyItems0Description(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := ``

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "Description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.Description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerGetPluginPrivilegesOKBodyItems0Name(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "Name"
	} else {
		nameFlagName = fmt.Sprintf("%v.Name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerGetPluginPrivilegesOKBodyItems0Value(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Value []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetPluginPrivilegesOKBodyItems0Flags(depth int, m *plugin.GetPluginPrivilegesOKBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrieveGetPluginPrivilegesOKBodyItems0DescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, nameAdded := retrieveGetPluginPrivilegesOKBodyItems0NameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, valueAdded := retrieveGetPluginPrivilegesOKBodyItems0ValueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || valueAdded

	return nil, retAdded
}

func retrieveGetPluginPrivilegesOKBodyItems0DescriptionFlags(depth int, m *plugin.GetPluginPrivilegesOKBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.Description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "Description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.Description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveGetPluginPrivilegesOKBodyItems0NameFlags(depth int, m *plugin.GetPluginPrivilegesOKBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.Name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "Name"
		} else {
			nameFlagName = fmt.Sprintf("%v.Name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveGetPluginPrivilegesOKBodyItems0ValueFlags(depth int, m *plugin.GetPluginPrivilegesOKBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	valueFlagName := fmt.Sprintf("%v.Value", cmdPrefix)
	if cmd.Flags().Changed(valueFlagName) {
		// warning: Value array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

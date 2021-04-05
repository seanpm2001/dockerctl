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

// makeOperationPluginPluginPullCmd returns a cmd to handle operation pluginPull
func makeOperationPluginPluginPullCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "PluginPull",
		Short: `Pulls and installs a plugin. After the plugin is installed, it can be enabled using the [` + "`" + `POST /plugins/{name}/enable` + "`" + ` endpoint](#operation/PostPluginsEnable).
`,
		RunE: runOperationPluginPluginPull,
	}

	if err := registerOperationPluginPluginPullParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPluginPluginPull uses cmd flags to call endpoint api
func runOperationPluginPluginPull(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := plugin.NewPluginPullParams()
	if err, _ := retrieveOperationPluginPluginPullXRegistryAuthFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPluginPluginPullBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPluginPluginPullNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPluginPluginPullRemoteFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationPluginPluginPullResult(appCli.Plugin.PluginPull(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationPluginPluginPullXRegistryAuthFlag(m *plugin.PluginPullParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("X-Registry-Auth") {

		var xRegistryAuthFlagName string
		if cmdPrefix == "" {
			xRegistryAuthFlagName = "X-Registry-Auth"
		} else {
			xRegistryAuthFlagName = fmt.Sprintf("%v.X-Registry-Auth", cmdPrefix)
		}

		xRegistryAuthFlagValue, err := cmd.Flags().GetString(xRegistryAuthFlagName)
		if err != nil {
			return err, false
		}
		m.XRegistryAuth = &xRegistryAuthFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPluginPluginPullBodyFlag(m *plugin.PluginPullParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// warning: body array type []*PluginPullParamsBodyItems0 is not supported by go-swagger cli yet
	}
	return nil, retAdded
}
func retrieveOperationPluginPluginPullNameFlag(m *plugin.PluginPullParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("name") {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = &nameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPluginPluginPullRemoteFlag(m *plugin.PluginPullParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationPluginPluginPullResult prints output to stdout
func printOperationPluginPluginPullResult(resp0 *plugin.PluginPullNoContent, respErr error) error {
	if respErr != nil {

		// Non schema case: warning pluginPullNoContent is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*plugin.PluginPullInternalServerError)
		if ok {
			if !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return err
				}
				fmt.Println(string(msgStr))
				return nil
			}
		}

		return respErr
	}

	// warning: non schema response pluginPullNoContent is not supported by go-swagger cli yet.

	return nil
}

// registerOperationPluginPluginPullParamFlags registers all flags needed to fill params
func registerOperationPluginPluginPullParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPluginPluginPullXRegistryAuthParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPluginPluginPullBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPluginPluginPullNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPluginPluginPullRemoteParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPluginPluginPullXRegistryAuthParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	xRegistryAuthDescription := `A base64-encoded auth configuration to use when pulling a plugin from a registry. [See the authentication section for details.](#section/Authentication)`

	var xRegistryAuthFlagName string
	if cmdPrefix == "" {
		xRegistryAuthFlagName = "X-Registry-Auth"
	} else {
		xRegistryAuthFlagName = fmt.Sprintf("%v.X-Registry-Auth", cmdPrefix)
	}

	var xRegistryAuthFlagDefault string

	_ = cmd.PersistentFlags().String(xRegistryAuthFlagName, xRegistryAuthFlagDefault, xRegistryAuthDescription)

	return nil
}
func registerOperationPluginPluginPullBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {
	// warning: go type []*PluginPullParamsBodyItems0 is not supported by go-swagger cli yet.
	return nil
}
func registerOperationPluginPluginPullNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Local name for the pulled plugin.

The ` + "`" + `:latest` + "`" + ` tag is optional, and is used as the default if omitted.
`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}
func registerOperationPluginPluginPullRemoteParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	remoteDescription := `Required. Remote reference for plugin to install.

The ` + "`" + `:latest` + "`" + ` tag is optional, and is used as the default if omitted.
`

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

// register flags to command
func registerModelPluginPullParamsBodyItems0Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPluginPullParamsBodyItems0Description(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPluginPullParamsBodyItems0Name(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPluginPullParamsBodyItems0Value(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPluginPullParamsBodyItems0Description(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerPluginPullParamsBodyItems0Name(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerPluginPullParamsBodyItems0Value(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Value []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPluginPullParamsBodyItems0Flags(depth int, m *plugin.PluginPullParamsBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrievePluginPullParamsBodyItems0DescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, nameAdded := retrievePluginPullParamsBodyItems0NameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, valueAdded := retrievePluginPullParamsBodyItems0ValueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || valueAdded

	return nil, retAdded
}

func retrievePluginPullParamsBodyItems0DescriptionFlags(depth int, m *plugin.PluginPullParamsBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePluginPullParamsBodyItems0NameFlags(depth int, m *plugin.PluginPullParamsBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePluginPullParamsBodyItems0ValueFlags(depth int, m *plugin.PluginPullParamsBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

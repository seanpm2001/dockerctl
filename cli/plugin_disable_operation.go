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

// makeOperationPluginPluginDisableCmd returns a cmd to handle operation pluginDisable
func makeOperationPluginPluginDisableCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "PluginDisable",
		Short: ``,
		RunE:  runOperationPluginPluginDisable,
	}

	if err := registerOperationPluginPluginDisableParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPluginPluginDisable uses cmd flags to call endpoint api
func runOperationPluginPluginDisable(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := plugin.NewPluginDisableParams()
	if err, _ := retrieveOperationPluginPluginDisableNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationPluginPluginDisableResult(appCli.Plugin.PluginDisable(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationPluginPluginDisableParamFlags registers all flags needed to fill params
func registerOperationPluginPluginDisableParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPluginPluginDisableNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPluginPluginDisableNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	NameDescription := `Required. The name of the plugin. The ` + "`" + `:latest` + "`" + ` tag is optional, and is the default if omitted.`

	var NameFlagName string
	if cmdPrefix == "" {
		NameFlagName = "name"
	} else {
		NameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var NameFlagDefault string

	_ = cmd.PersistentFlags().String(NameFlagName, NameFlagDefault, NameDescription)

	return nil
}

func retrieveOperationPluginPluginDisableNameFlag(m *plugin.PluginDisableParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("name") {

		var NameFlagName string
		if cmdPrefix == "" {
			NameFlagName = "name"
		} else {
			NameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		NameFlagValue, err := cmd.Flags().GetString(NameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = NameFlagValue

	}
	return nil, retAdded
}

// parseOperationPluginPluginDisableResult parses request result and return the string content
func parseOperationPluginPluginDisableResult(resp0 *plugin.PluginDisableOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning pluginDisableOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*plugin.PluginDisableNotFound)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*plugin.PluginDisableInternalServerError)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response pluginDisableOK is not supported by go-swagger cli yet.

	return "", nil
}

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

// makeOperationPluginPluginListCmd returns a cmd to handle operation pluginList
func makeOperationPluginPluginListCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "PluginList",
		Short: `Returns information about installed plugins.`,
		RunE:  runOperationPluginPluginList,
	}

	if err := registerOperationPluginPluginListParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPluginPluginList uses cmd flags to call endpoint api
func runOperationPluginPluginList(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := plugin.NewPluginListParams()
	if err, _ := retrieveOperationPluginPluginListFiltersFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationPluginPluginListResult(appCli.Plugin.PluginList(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationPluginPluginListParamFlags registers all flags needed to fill params
func registerOperationPluginPluginListParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPluginPluginListFiltersParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPluginPluginListFiltersParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	FiltersDescription := `A JSON encoded value of the filters (a ` + "`" + `map[string][]string` + "`" + `) to process on the plugin list. Available filters:

- ` + "`" + `capability=<capability name>` + "`" + `
- ` + "`" + `enable=<true>|<false>` + "`" + `
`

	var FiltersFlagName string
	if cmdPrefix == "" {
		FiltersFlagName = "filters"
	} else {
		FiltersFlagName = fmt.Sprintf("%v.filters", cmdPrefix)
	}

	var FiltersFlagDefault string

	_ = cmd.PersistentFlags().String(FiltersFlagName, FiltersFlagDefault, FiltersDescription)

	return nil
}

func retrieveOperationPluginPluginListFiltersFlag(m *plugin.PluginListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filters") {

		var FiltersFlagName string
		if cmdPrefix == "" {
			FiltersFlagName = "filters"
		} else {
			FiltersFlagName = fmt.Sprintf("%v.filters", cmdPrefix)
		}

		FiltersFlagValue, err := cmd.Flags().GetString(FiltersFlagName)
		if err != nil {
			return err, false
		}
		m.Filters = &FiltersFlagValue

	}
	return nil, retAdded
}

// parseOperationPluginPluginListResult parses request result and return the string content
func parseOperationPluginPluginListResult(resp0 *plugin.PluginListOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*plugin.PluginListOK)
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
		resp1, ok := iResp1.(*plugin.PluginListInternalServerError)
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

// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/volume"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationVolumeVolumeListCmd returns a cmd to handle operation volumeList
func makeOperationVolumeVolumeListCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "VolumeList",
		Short: ``,
		RunE:  runOperationVolumeVolumeList,
	}

	if err := registerOperationVolumeVolumeListParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVolumeVolumeList uses cmd flags to call endpoint api
func runOperationVolumeVolumeList(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := volume.NewVolumeListParams()
	if err, _ := retrieveOperationVolumeVolumeListFiltersFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationVolumeVolumeListResult(appCli.Volume.VolumeList(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationVolumeVolumeListParamFlags registers all flags needed to fill params
func registerOperationVolumeVolumeListParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVolumeVolumeListFiltersParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVolumeVolumeListFiltersParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filtersDescription := `JSON encoded value of the filters (a ` + "`" + `map[string][]string` + "`" + `) to
process on the volumes list. Available filters:

- ` + "`" + `dangling=<boolean>` + "`" + ` When set to ` + "`" + `true` + "`" + ` (or ` + "`" + `1` + "`" + `), returns all
   volumes that are not in use by a container. When set to ` + "`" + `false` + "`" + `
   (or ` + "`" + `0` + "`" + `), only volumes that are in use by one or more
   containers are returned.
- ` + "`" + `driver=<volume-driver-name>` + "`" + ` Matches volumes based on their driver.
- ` + "`" + `label=<key>` + "`" + ` or ` + "`" + `label=<key>:<value>` + "`" + ` Matches volumes based on
   the presence of a ` + "`" + `label` + "`" + ` alone or a ` + "`" + `label` + "`" + ` and a value.
- ` + "`" + `name=<volume-name>` + "`" + ` Matches all or part of a volume name.
`

	var filtersFlagName string
	if cmdPrefix == "" {
		filtersFlagName = "filters"
	} else {
		filtersFlagName = fmt.Sprintf("%v.filters", cmdPrefix)
	}

	var filtersFlagDefault string

	_ = cmd.PersistentFlags().String(filtersFlagName, filtersFlagDefault, filtersDescription)

	return nil
}

func retrieveOperationVolumeVolumeListFiltersFlag(m *volume.VolumeListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filters") {

		var filtersFlagName string
		if cmdPrefix == "" {
			filtersFlagName = "filters"
		} else {
			filtersFlagName = fmt.Sprintf("%v.filters", cmdPrefix)
		}

		filtersFlagValue, err := cmd.Flags().GetString(filtersFlagName)
		if err != nil {
			return err, false
		}
		m.Filters = &filtersFlagValue

	}
	return nil, retAdded
}

// parseOperationVolumeVolumeListResult parses request result and return the string content
func parseOperationVolumeVolumeListResult(resp0 *volume.VolumeListOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*volume.VolumeListOK)
		if ok {
			if !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*volume.VolumeListInternalServerError)
		if ok {
			if !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}

// register flags to command
func registerModelVolumeListOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVolumeListOKBodyVolumes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVolumeListOKBodyWarnings(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVolumeListOKBodyVolumes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Volumes []*models.Volume array type is not supported by go-swagger cli yet

	return nil
}

func registerVolumeListOKBodyWarnings(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Warnings []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVolumeListOKBodyFlags(depth int, m *volume.VolumeListOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, volumesAdded := retrieveVolumeListOKBodyVolumesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || volumesAdded

	err, warningsAdded := retrieveVolumeListOKBodyWarningsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || warningsAdded

	return nil, retAdded
}

func retrieveVolumeListOKBodyVolumesFlags(depth int, m *volume.VolumeListOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	volumesFlagName := fmt.Sprintf("%v.Volumes", cmdPrefix)
	if cmd.Flags().Changed(volumesFlagName) {
		// warning: Volumes array type []*models.Volume is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveVolumeListOKBodyWarningsFlags(depth int, m *volume.VolumeListOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	warningsFlagName := fmt.Sprintf("%v.Warnings", cmdPrefix)
	if cmd.Flags().Changed(warningsFlagName) {
		// warning: Warnings array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

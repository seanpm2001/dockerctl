// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/container"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationContainerContainerPruneCmd returns a cmd to handle operation containerPrune
func makeOperationContainerContainerPruneCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ContainerPrune",
		Short: ``,
		RunE:  runOperationContainerContainerPrune,
	}

	if err := registerOperationContainerContainerPruneParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerContainerPrune uses cmd flags to call endpoint api
func runOperationContainerContainerPrune(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewContainerPruneParams()
	if err, _ := retrieveOperationContainerContainerPruneFiltersFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationContainerContainerPruneResult(appCli.Container.ContainerPrune(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationContainerContainerPruneFiltersFlag(m *container.ContainerPruneParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationContainerContainerPruneResult prints output to stdout
func printOperationContainerContainerPruneResult(resp0 *container.ContainerPruneOK, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*container.ContainerPruneOK)
		if ok {
			if !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return err
				}
				fmt.Println(string(msgStr))
				return nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*container.ContainerPruneInternalServerError)
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

	if !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return err
		}
		fmt.Println(string(msgStr))
	}

	return nil
}

// registerOperationContainerContainerPruneParamFlags registers all flags needed to fill params
func registerOperationContainerContainerPruneParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerContainerPruneFiltersParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerContainerPruneFiltersParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filtersDescription := `Filters to process on the prune list, encoded as JSON (a ` + "`" + `map[string][]string` + "`" + `).

Available filters:
- ` + "`" + `until=<timestamp>` + "`" + ` Prune containers created before this timestamp. The ` + "`" + `<timestamp>` + "`" + ` can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. ` + "`" + `10m` + "`" + `, ` + "`" + `1h30m` + "`" + `) computed relative to the daemon machine’s time.
- ` + "`" + `label` + "`" + ` (` + "`" + `label=<key>` + "`" + `, ` + "`" + `label=<key>=<value>` + "`" + `, ` + "`" + `label!=<key>` + "`" + `, or ` + "`" + `label!=<key>=<value>` + "`" + `) Prune containers with (or without, in case ` + "`" + `label!=...` + "`" + ` is used) the specified labels.
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

// register flags to command
func registerModelContainerPruneOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerContainerPruneOKBodyContainersDeleted(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerPruneOKBodySpaceReclaimed(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainerPruneOKBodyContainersDeleted(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: ContainersDeleted []string array type is not supported by go-swagger cli yet

	return nil
}

func registerContainerPruneOKBodySpaceReclaimed(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	spaceReclaimedDescription := `Disk space reclaimed in bytes`

	var spaceReclaimedFlagName string
	if cmdPrefix == "" {
		spaceReclaimedFlagName = "SpaceReclaimed"
	} else {
		spaceReclaimedFlagName = fmt.Sprintf("%v.SpaceReclaimed", cmdPrefix)
	}

	var spaceReclaimedFlagDefault int64

	_ = cmd.PersistentFlags().Int64(spaceReclaimedFlagName, spaceReclaimedFlagDefault, spaceReclaimedDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainerPruneOKBodyFlags(depth int, m *container.ContainerPruneOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, containersDeletedAdded := retrieveContainerPruneOKBodyContainersDeletedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || containersDeletedAdded

	err, spaceReclaimedAdded := retrieveContainerPruneOKBodySpaceReclaimedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || spaceReclaimedAdded

	return nil, retAdded
}

func retrieveContainerPruneOKBodyContainersDeletedFlags(depth int, m *container.ContainerPruneOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	containersDeletedFlagName := fmt.Sprintf("%v.ContainersDeleted", cmdPrefix)
	if cmd.Flags().Changed(containersDeletedFlagName) {
		// warning: ContainersDeleted array type []string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

func retrieveContainerPruneOKBodySpaceReclaimedFlags(depth int, m *container.ContainerPruneOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	spaceReclaimedFlagName := fmt.Sprintf("%v.SpaceReclaimed", cmdPrefix)
	if cmd.Flags().Changed(spaceReclaimedFlagName) {

		var spaceReclaimedFlagName string
		if cmdPrefix == "" {
			spaceReclaimedFlagName = "SpaceReclaimed"
		} else {
			spaceReclaimedFlagName = fmt.Sprintf("%v.SpaceReclaimed", cmdPrefix)
		}

		spaceReclaimedFlagValue, err := cmd.Flags().GetInt64(spaceReclaimedFlagName)
		if err != nil {
			return err, false
		}
		m.SpaceReclaimed = spaceReclaimedFlagValue

		retAdded = true
	}
	return nil, retAdded
}

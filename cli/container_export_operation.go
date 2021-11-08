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

// makeOperationContainerContainerExportCmd returns a cmd to handle operation containerExport
func makeOperationContainerContainerExportCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ContainerExport",
		Short: `Export the contents of a container as a tarball.`,
		RunE:  runOperationContainerContainerExport,
	}

	if err := registerOperationContainerContainerExportParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerContainerExport uses cmd flags to call endpoint api
func runOperationContainerContainerExport(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewContainerExportParams()
	if err, _ := retrieveOperationContainerContainerExportIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationContainerContainerExportResult(appCli.Container.ContainerExport(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationContainerContainerExportParamFlags registers all flags needed to fill params
func registerOperationContainerContainerExportParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerContainerExportIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerContainerExportIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. ID or name of the container`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func retrieveOperationContainerContainerExportIDFlag(m *container.ContainerExportParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

	}
	return nil, retAdded
}

// parseOperationContainerContainerExportResult parses request result and return the string content
func parseOperationContainerContainerExportResult(resp0 *container.ContainerExportOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning containerExportOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*container.ContainerExportNotFound)
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
		resp2, ok := iResp2.(*container.ContainerExportInternalServerError)
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

	// warning: non schema response containerExportOK is not supported by go-swagger cli yet.

	return "", nil
}

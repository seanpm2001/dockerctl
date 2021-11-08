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

// makeOperationVolumeVolumeDeleteCmd returns a cmd to handle operation volumeDelete
func makeOperationVolumeVolumeDeleteCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "VolumeDelete",
		Short: `Instruct the driver to remove the volume.`,
		RunE:  runOperationVolumeVolumeDelete,
	}

	if err := registerOperationVolumeVolumeDeleteParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVolumeVolumeDelete uses cmd flags to call endpoint api
func runOperationVolumeVolumeDelete(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := volume.NewVolumeDeleteParams()
	if err, _ := retrieveOperationVolumeVolumeDeleteForceFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationVolumeVolumeDeleteNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationVolumeVolumeDeleteResult(appCli.Volume.VolumeDelete(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationVolumeVolumeDeleteParamFlags registers all flags needed to fill params
func registerOperationVolumeVolumeDeleteParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVolumeVolumeDeleteForceParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationVolumeVolumeDeleteNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVolumeVolumeDeleteForceParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	forceDescription := `Force the removal of the volume`

	var forceFlagName string
	if cmdPrefix == "" {
		forceFlagName = "force"
	} else {
		forceFlagName = fmt.Sprintf("%v.force", cmdPrefix)
	}

	var forceFlagDefault bool

	_ = cmd.PersistentFlags().Bool(forceFlagName, forceFlagDefault, forceDescription)

	return nil
}
func registerOperationVolumeVolumeDeleteNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Volume name or ID`

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

func retrieveOperationVolumeVolumeDeleteForceFlag(m *volume.VolumeDeleteParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("force") {

		var forceFlagName string
		if cmdPrefix == "" {
			forceFlagName = "force"
		} else {
			forceFlagName = fmt.Sprintf("%v.force", cmdPrefix)
		}

		forceFlagValue, err := cmd.Flags().GetBool(forceFlagName)
		if err != nil {
			return err, false
		}
		m.Force = &forceFlagValue

	}
	return nil, retAdded
}
func retrieveOperationVolumeVolumeDeleteNameFlag(m *volume.VolumeDeleteParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Name = nameFlagValue

	}
	return nil, retAdded
}

// parseOperationVolumeVolumeDeleteResult parses request result and return the string content
func parseOperationVolumeVolumeDeleteResult(resp0 *volume.VolumeDeleteNoContent, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning volumeDeleteNoContent is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*volume.VolumeDeleteNotFound)
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
		resp2, ok := iResp2.(*volume.VolumeDeleteConflict)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*volume.VolumeDeleteInternalServerError)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response volumeDeleteNoContent is not supported by go-swagger cli yet.

	return "", nil
}

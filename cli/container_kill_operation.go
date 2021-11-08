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

// makeOperationContainerContainerKillCmd returns a cmd to handle operation containerKill
func makeOperationContainerContainerKillCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ContainerKill",
		Short: `Send a POSIX signal to a container, defaulting to killing to the container.`,
		RunE:  runOperationContainerContainerKill,
	}

	if err := registerOperationContainerContainerKillParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerContainerKill uses cmd flags to call endpoint api
func runOperationContainerContainerKill(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewContainerKillParams()
	if err, _ := retrieveOperationContainerContainerKillIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerKillSignalFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationContainerContainerKillResult(appCli.Container.ContainerKill(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationContainerContainerKillParamFlags registers all flags needed to fill params
func registerOperationContainerContainerKillParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerContainerKillIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerKillSignalParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerContainerKillIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationContainerContainerKillSignalParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	signalDescription := `Signal to send to the container as an integer or string (e.g. ` + "`" + `SIGINT` + "`" + `)`

	var signalFlagName string
	if cmdPrefix == "" {
		signalFlagName = "signal"
	} else {
		signalFlagName = fmt.Sprintf("%v.signal", cmdPrefix)
	}

	var signalFlagDefault string = "SIGKILL"

	_ = cmd.PersistentFlags().String(signalFlagName, signalFlagDefault, signalDescription)

	return nil
}

func retrieveOperationContainerContainerKillIDFlag(m *container.ContainerKillParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerContainerKillSignalFlag(m *container.ContainerKillParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("signal") {

		var signalFlagName string
		if cmdPrefix == "" {
			signalFlagName = "signal"
		} else {
			signalFlagName = fmt.Sprintf("%v.signal", cmdPrefix)
		}

		signalFlagValue, err := cmd.Flags().GetString(signalFlagName)
		if err != nil {
			return err, false
		}
		m.Signal = &signalFlagValue

	}
	return nil, retAdded
}

// parseOperationContainerContainerKillResult parses request result and return the string content
func parseOperationContainerContainerKillResult(resp0 *container.ContainerKillNoContent, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning containerKillNoContent is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*container.ContainerKillNotFound)
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
		resp2, ok := iResp2.(*container.ContainerKillConflict)
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
		resp3, ok := iResp3.(*container.ContainerKillInternalServerError)
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

	// warning: non schema response containerKillNoContent is not supported by go-swagger cli yet.

	return "", nil
}

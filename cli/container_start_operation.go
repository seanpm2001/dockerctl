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

// makeOperationContainerContainerStartCmd returns a cmd to handle operation containerStart
func makeOperationContainerContainerStartCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ContainerStart",
		Short: ``,
		RunE:  runOperationContainerContainerStart,
	}

	if err := registerOperationContainerContainerStartParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerContainerStart uses cmd flags to call endpoint api
func runOperationContainerContainerStart(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewContainerStartParams()
	if err, _ := retrieveOperationContainerContainerStartDetachKeysFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerStartIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationContainerContainerStartResult(appCli.Container.ContainerStart(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationContainerContainerStartParamFlags registers all flags needed to fill params
func registerOperationContainerContainerStartParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerContainerStartDetachKeysParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerStartIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerContainerStartDetachKeysParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	detachKeysDescription := `Override the key sequence for detaching a container. Format is a single character ` + "`" + `[a-Z]` + "`" + ` or ` + "`" + `ctrl-<value>` + "`" + ` where ` + "`" + `<value>` + "`" + ` is one of: ` + "`" + `a-z` + "`" + `, ` + "`" + `@` + "`" + `, ` + "`" + `^` + "`" + `, ` + "`" + `[` + "`" + `, ` + "`" + `,` + "`" + ` or ` + "`" + `_` + "`" + `.`

	var detachKeysFlagName string
	if cmdPrefix == "" {
		detachKeysFlagName = "detachKeys"
	} else {
		detachKeysFlagName = fmt.Sprintf("%v.detachKeys", cmdPrefix)
	}

	var detachKeysFlagDefault string

	_ = cmd.PersistentFlags().String(detachKeysFlagName, detachKeysFlagDefault, detachKeysDescription)

	return nil
}
func registerOperationContainerContainerStartIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationContainerContainerStartDetachKeysFlag(m *container.ContainerStartParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("detachKeys") {

		var detachKeysFlagName string
		if cmdPrefix == "" {
			detachKeysFlagName = "detachKeys"
		} else {
			detachKeysFlagName = fmt.Sprintf("%v.detachKeys", cmdPrefix)
		}

		detachKeysFlagValue, err := cmd.Flags().GetString(detachKeysFlagName)
		if err != nil {
			return err, false
		}
		m.DetachKeys = &detachKeysFlagValue

	}
	return nil, retAdded
}
func retrieveOperationContainerContainerStartIDFlag(m *container.ContainerStartParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationContainerContainerStartResult parses request result and return the string content
func parseOperationContainerContainerStartResult(resp0 *container.ContainerStartNoContent, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning containerStartNoContent is not supported

		// Non schema case: warning containerStartNotModified is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*container.ContainerStartNotFound)
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
		resp3, ok := iResp3.(*container.ContainerStartInternalServerError)
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

	// warning: non schema response containerStartNoContent is not supported by go-swagger cli yet.

	return "", nil
}

// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/network"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationNetworkNetworkDeleteCmd returns a cmd to handle operation networkDelete
func makeOperationNetworkNetworkDeleteCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "NetworkDelete",
		Short: ``,
		RunE:  runOperationNetworkNetworkDelete,
	}

	if err := registerOperationNetworkNetworkDeleteParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationNetworkNetworkDelete uses cmd flags to call endpoint api
func runOperationNetworkNetworkDelete(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := network.NewNetworkDeleteParams()
	if err, _ := retrieveOperationNetworkNetworkDeleteIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationNetworkNetworkDeleteResult(appCli.Network.NetworkDelete(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationNetworkNetworkDeleteParamFlags registers all flags needed to fill params
func registerOperationNetworkNetworkDeleteParamFlags(cmd *cobra.Command) error {
	if err := registerOperationNetworkNetworkDeleteIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationNetworkNetworkDeleteIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. Network ID or name`

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

func retrieveOperationNetworkNetworkDeleteIDFlag(m *network.NetworkDeleteParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationNetworkNetworkDeleteResult parses request result and return the string content
func parseOperationNetworkNetworkDeleteResult(resp0 *network.NetworkDeleteNoContent, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning networkDeleteNoContent is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*network.NetworkDeleteForbidden)
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
		resp2, ok := iResp2.(*network.NetworkDeleteNotFound)
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
		resp3, ok := iResp3.(*network.NetworkDeleteInternalServerError)
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

	// warning: non schema response networkDeleteNoContent is not supported by go-swagger cli yet.

	return "", nil
}

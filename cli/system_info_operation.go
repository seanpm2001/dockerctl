// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/system"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSystemSystemInfoCmd returns a cmd to handle operation systemInfo
func makeOperationSystemSystemInfoCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "SystemInfo",
		Short: ``,
		RunE:  runOperationSystemSystemInfo,
	}

	if err := registerOperationSystemSystemInfoParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSystemSystemInfo uses cmd flags to call endpoint api
func runOperationSystemSystemInfo(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := system.NewSystemInfoParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSystemSystemInfoResult(appCli.System.SystemInfo(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSystemSystemInfoParamFlags registers all flags needed to fill params
func registerOperationSystemSystemInfoParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationSystemSystemInfoResult parses request result and return the string content
func parseOperationSystemSystemInfoResult(resp0 *system.SystemInfoOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*system.SystemInfoOK)
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
		resp1, ok := iResp1.(*system.SystemInfoInternalServerError)
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

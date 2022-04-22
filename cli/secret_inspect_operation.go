// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/secret"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSecretSecretInspectCmd returns a cmd to handle operation secretInspect
func makeOperationSecretSecretInspectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "SecretInspect",
		Short: ``,
		RunE:  runOperationSecretSecretInspect,
	}

	if err := registerOperationSecretSecretInspectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSecretSecretInspect uses cmd flags to call endpoint api
func runOperationSecretSecretInspect(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := secret.NewSecretInspectParams()
	if err, _ := retrieveOperationSecretSecretInspectIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSecretSecretInspectResult(appCli.Secret.SecretInspect(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSecretSecretInspectParamFlags registers all flags needed to fill params
func registerOperationSecretSecretInspectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSecretSecretInspectIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSecretSecretInspectIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	IDDescription := `Required. ID of the secret`

	var IDFlagName string
	if cmdPrefix == "" {
		IDFlagName = "id"
	} else {
		IDFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var IDFlagDefault string

	_ = cmd.PersistentFlags().String(IDFlagName, IDFlagDefault, IDDescription)

	return nil
}

func retrieveOperationSecretSecretInspectIDFlag(m *secret.SecretInspectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var IDFlagName string
		if cmdPrefix == "" {
			IDFlagName = "id"
		} else {
			IDFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		IDFlagValue, err := cmd.Flags().GetString(IDFlagName)
		if err != nil {
			return err, false
		}
		m.ID = IDFlagValue

	}
	return nil, retAdded
}

// parseOperationSecretSecretInspectResult parses request result and return the string content
func parseOperationSecretSecretInspectResult(resp0 *secret.SecretInspectOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*secret.SecretInspectOK)
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
		resp1, ok := iResp1.(*secret.SecretInspectNotFound)
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
		resp2, ok := iResp2.(*secret.SecretInspectInternalServerError)
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
		resp3, ok := iResp3.(*secret.SecretInspectServiceUnavailable)
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

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}

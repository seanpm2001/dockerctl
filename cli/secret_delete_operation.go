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

// makeOperationSecretSecretDeleteCmd returns a cmd to handle operation secretDelete
func makeOperationSecretSecretDeleteCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "SecretDelete",
		Short: ``,
		RunE:  runOperationSecretSecretDelete,
	}

	if err := registerOperationSecretSecretDeleteParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSecretSecretDelete uses cmd flags to call endpoint api
func runOperationSecretSecretDelete(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := secret.NewSecretDeleteParams()
	if err, _ := retrieveOperationSecretSecretDeleteIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSecretSecretDeleteResult(appCli.Secret.SecretDelete(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSecretSecretDeleteParamFlags registers all flags needed to fill params
func registerOperationSecretSecretDeleteParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSecretSecretDeleteIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSecretSecretDeleteIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. ID of the secret`

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

func retrieveOperationSecretSecretDeleteIDFlag(m *secret.SecretDeleteParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSecretSecretDeleteResult parses request result and return the string content
func parseOperationSecretSecretDeleteResult(resp0 *secret.SecretDeleteNoContent, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning secretDeleteNoContent is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*secret.SecretDeleteNotFound)
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
		resp2, ok := iResp2.(*secret.SecretDeleteInternalServerError)
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
		resp3, ok := iResp3.(*secret.SecretDeleteServiceUnavailable)
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

	// warning: non schema response secretDeleteNoContent is not supported by go-swagger cli yet.

	return "", nil
}

// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/swarm"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSwarmSwarmInspectCmd returns a cmd to handle operation swarmInspect
func makeOperationSwarmSwarmInspectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "SwarmInspect",
		Short: ``,
		RunE:  runOperationSwarmSwarmInspect,
	}

	if err := registerOperationSwarmSwarmInspectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSwarmSwarmInspect uses cmd flags to call endpoint api
func runOperationSwarmSwarmInspect(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := swarm.NewSwarmInspectParams()
	// make request and then print result
	if err := printOperationSwarmSwarmInspectResult(appCli.Swarm.SwarmInspect(params)); err != nil {
		return err
	}
	return nil
}

// printOperationSwarmSwarmInspectResult prints output to stdout
func printOperationSwarmSwarmInspectResult(resp0 *swarm.SwarmInspectOK, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*swarm.SwarmInspectOK)
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
		resp1, ok := iResp1.(*swarm.SwarmInspectNotFound)
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

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*swarm.SwarmInspectInternalServerError)
		if ok {
			if !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return err
				}
				fmt.Println(string(msgStr))
				return nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*swarm.SwarmInspectServiceUnavailable)
		if ok {
			if !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
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

// registerOperationSwarmSwarmInspectParamFlags registers all flags needed to fill params
func registerOperationSwarmSwarmInspectParamFlags(cmd *cobra.Command) error {
	return nil
}

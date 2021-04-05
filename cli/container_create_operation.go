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

// makeOperationContainerContainerCreateCmd returns a cmd to handle operation containerCreate
func makeOperationContainerContainerCreateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ContainerCreate",
		Short: ``,
		RunE:  runOperationContainerContainerCreate,
	}

	if err := registerOperationContainerContainerCreateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerContainerCreate uses cmd flags to call endpoint api
func runOperationContainerContainerCreate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewContainerCreateParams()
	if err, _ := retrieveOperationContainerContainerCreateBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerCreateNameFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationContainerContainerCreateResult(appCli.Container.ContainerCreate(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationContainerContainerCreateBodyFlag(m *container.ContainerCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := container.ContainerCreateBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in ContainerCreateBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = container.ContainerCreateBody{}
	}
	err, added := retrieveModelContainerCreateBodyFlags(0, &bodyValueModel, "containerCreateBody", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = bodyValueModel
	}
	bodyValueDebugBytes, err := json.Marshal(m.Body)
	if err != nil {
		return err, false
	}
	logDebugf("Body payload: %v", string(bodyValueDebugBytes))
	return nil, retAdded
}
func retrieveOperationContainerContainerCreateNameFlag(m *container.ContainerCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Name = &nameFlagValue

	}
	return nil, retAdded
}

// printOperationContainerContainerCreateResult prints output to stdout
func printOperationContainerContainerCreateResult(resp0 *container.ContainerCreateCreated, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*container.ContainerCreateCreated)
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
		resp1, ok := iResp1.(*container.ContainerCreateBadRequest)
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
		resp2, ok := iResp2.(*container.ContainerCreateNotFound)
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
		resp3, ok := iResp3.(*container.ContainerCreateConflict)
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

		var iResp4 interface{} = respErr
		resp4, ok := iResp4.(*container.ContainerCreateInternalServerError)
		if ok {
			if !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
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

// registerOperationContainerContainerCreateParamFlags registers all flags needed to fill params
func registerOperationContainerContainerCreateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerContainerCreateBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerCreateNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerContainerCreateBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	exampleBodyStr := "go-swagger TODO"
	_ = cmd.PersistentFlags().String(bodyFlagName, "", fmt.Sprintf("Optional json string for [body] of form %v.Container to create", string(exampleBodyStr)))

	// add flags for body
	if err := registerModelContainerCreateBodyFlags(0, "containerCreateBody", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationContainerContainerCreateNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Assign the specified name to the container. Must match ` + "`" + `/?[a-zA-Z0-9][a-zA-Z0-9_.-]+` + "`" + `.`

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

// register flags to command
func registerModelContainerCreateBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// allOf ContainerCreateParamsBodyAO0 is not supported by go-swwagger cli yet

	// allOf ContainerCreateParamsBodyAO1 is not supported by go-swwagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainerCreateBodyFlags(depth int, m *container.ContainerCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// allOf ContainerCreateParamsBodyAO0 is not supported by go-swwagger cli yet

	// allOf ContainerCreateParamsBodyAO1 is not supported by go-swwagger cli yet

	return nil, retAdded
}

// register flags to command
func registerModelContainerCreateCreatedBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerContainerCreateCreatedBodyID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerCreateCreatedBodyWarnings(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainerCreateCreatedBodyID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. The ID of the created container`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "Id"
	} else {
		idFlagName = fmt.Sprintf("%v.Id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerContainerCreateCreatedBodyWarnings(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Warnings []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainerCreateCreatedBodyFlags(depth int, m *container.ContainerCreateCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveContainerCreateCreatedBodyIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, warningsAdded := retrieveContainerCreateCreatedBodyWarningsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || warningsAdded

	return nil, retAdded
}

func retrieveContainerCreateCreatedBodyIDFlags(depth int, m *container.ContainerCreateCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	idFlagName := fmt.Sprintf("%v.Id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "Id"
		} else {
			idFlagName = fmt.Sprintf("%v.Id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveContainerCreateCreatedBodyWarningsFlags(depth int, m *container.ContainerCreateCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	warningsFlagName := fmt.Sprintf("%v.Warnings", cmdPrefix)
	if cmd.Flags().Changed(warningsFlagName) {
		// warning: Warnings array type []string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

// register flags to command
func registerModelContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfigEndpointsConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfigEndpointsConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: EndpointsConfig map[string]models.EndpointSettings map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfigFlags(depth int, m *container.ContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, endpointsConfigAdded := retrieveContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfigEndpointsConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endpointsConfigAdded

	return nil, retAdded
}

func retrieveContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfigEndpointsConfigFlags(depth int, m *container.ContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	endpointsConfigFlagName := fmt.Sprintf("%v.EndpointsConfig", cmdPrefix)
	if cmd.Flags().Changed(endpointsConfigFlagName) {
		// warning: EndpointsConfig map type map[string]models.EndpointSettings is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

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

// makeOperationContainerContainerArchiveCmd returns a cmd to handle operation containerArchive
func makeOperationContainerContainerArchiveCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ContainerArchive",
		Short: `Get a tar archive of a resource in the filesystem of container id.`,
		RunE:  runOperationContainerContainerArchive,
	}

	if err := registerOperationContainerContainerArchiveParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerContainerArchive uses cmd flags to call endpoint api
func runOperationContainerContainerArchive(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewContainerArchiveParams()
	if err, _ := retrieveOperationContainerContainerArchiveIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerArchivePathFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationContainerContainerArchiveResult(appCli.Container.ContainerArchive(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationContainerContainerArchiveParamFlags registers all flags needed to fill params
func registerOperationContainerContainerArchiveParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerContainerArchiveIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerArchivePathParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerContainerArchiveIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	IDDescription := `Required. ID or name of the container`

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
func registerOperationContainerContainerArchivePathParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	PathDescription := `Required. Resource in the container’s filesystem to archive.`

	var PathFlagName string
	if cmdPrefix == "" {
		PathFlagName = "path"
	} else {
		PathFlagName = fmt.Sprintf("%v.path", cmdPrefix)
	}

	var PathFlagDefault string

	_ = cmd.PersistentFlags().String(PathFlagName, PathFlagDefault, PathDescription)

	return nil
}

func retrieveOperationContainerContainerArchiveIDFlag(m *container.ContainerArchiveParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerContainerArchivePathFlag(m *container.ContainerArchiveParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("path") {

		var PathFlagName string
		if cmdPrefix == "" {
			PathFlagName = "path"
		} else {
			PathFlagName = fmt.Sprintf("%v.path", cmdPrefix)
		}

		PathFlagValue, err := cmd.Flags().GetString(PathFlagName)
		if err != nil {
			return err, false
		}
		m.Path = PathFlagValue

	}
	return nil, retAdded
}

// parseOperationContainerContainerArchiveResult parses request result and return the string content
func parseOperationContainerContainerArchiveResult(resp0 *container.ContainerArchiveOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning containerArchiveOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*container.ContainerArchiveBadRequest)
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
		resp2, ok := iResp2.(*container.ContainerArchiveNotFound)
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
		resp3, ok := iResp3.(*container.ContainerArchiveInternalServerError)
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

	// warning: non schema response containerArchiveOK is not supported by go-swagger cli yet.

	return "", nil
}

// register flags to command
func registerModelContainerArchiveBadRequestBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// register embedded models.ErrorResponse flags

	if err := registerModelErrorResponseFlags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	// register anonymous fields for containerArchiveBadRequestBodyAO1

	if err := registerContainerArchiveBadRequestBodyAnonContainerArchiveBadRequestBodyAO1Message(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

// inline definition name containerArchiveBadRequestBodyAO1, type

func registerContainerArchiveBadRequestBodyAnonContainerArchiveBadRequestBodyAO1Message(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	MessageDescription := `The error message. Either "must specify path parameter" (path cannot be empty) or "not a directory" (path was asserted to be a directory but exists as a file).`

	var MessageFlagName string
	if cmdPrefix == "" {
		MessageFlagName = "message"
	} else {
		MessageFlagName = fmt.Sprintf("%v.message", cmdPrefix)
	}

	var MessageFlagDefault string

	_ = cmd.PersistentFlags().String(MessageFlagName, MessageFlagDefault, MessageDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainerArchiveBadRequestBodyFlags(depth int, m *container.ContainerArchiveBadRequestBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// retrieve model models.ErrorResponse
	err, ContainerArchiveBadRequestBodyAO0Added := retrieveModelErrorResponseFlags(depth, &m.ErrorResponse, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ContainerArchiveBadRequestBodyAO0Added

	// retrieve allOf containerArchiveBadRequestBodyAO1 fields

	err, MessageAdded := retrieveContainerArchiveBadRequestBodyAnonContainerArchiveBadRequestBodyAO1MessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || MessageAdded

	return nil, retAdded
}

// define retrieve functions for fields for inline definition name containerArchiveBadRequestBodyAO1

func retrieveContainerArchiveBadRequestBodyAnonContainerArchiveBadRequestBodyAO1MessageFlags(depth int, m *container.ContainerArchiveBadRequestBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	MessageFlagName := fmt.Sprintf("%v.message", cmdPrefix)
	if cmd.Flags().Changed(MessageFlagName) {

		var MessageFlagName string
		if cmdPrefix == "" {
			MessageFlagName = "message"
		} else {
			MessageFlagName = fmt.Sprintf("%v.message", cmdPrefix)
		}

		MessageFlagValue, err := cmd.Flags().GetString(MessageFlagName)
		if err != nil {
			return err, false
		}
		m.Message = MessageFlagValue

		retAdded = true
	}

	return nil, retAdded
}

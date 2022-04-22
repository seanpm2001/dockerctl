// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for ServiceUpdateResponse

// register flags to command
func registerModelServiceUpdateResponseFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServiceUpdateResponseWarnings(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceUpdateResponseWarnings(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Warnings []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServiceUpdateResponseFlags(depth int, m *models.ServiceUpdateResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, WarningsAdded := retrieveServiceUpdateResponseWarningsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || WarningsAdded

	return nil, retAdded
}

func retrieveServiceUpdateResponseWarningsFlags(depth int, m *models.ServiceUpdateResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	WarningsFlagName := fmt.Sprintf("%v.Warnings", cmdPrefix)
	if cmd.Flags().Changed(WarningsFlagName) {
		// warning: Warnings array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

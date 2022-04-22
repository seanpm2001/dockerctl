// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for IDResponse

// register flags to command
func registerModelIDResponseFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIDResponseID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIDResponseID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	IDDescription := `Required. The id of the newly created object.`

	var IDFlagName string
	if cmdPrefix == "" {
		IDFlagName = "Id"
	} else {
		IDFlagName = fmt.Sprintf("%v.Id", cmdPrefix)
	}

	var IDFlagDefault string

	_ = cmd.PersistentFlags().String(IDFlagName, IDFlagDefault, IDDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIDResponseFlags(depth int, m *models.IDResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, IDAdded := retrieveIDResponseIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || IDAdded

	return nil, retAdded
}

func retrieveIDResponseIDFlags(depth int, m *models.IDResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	IDFlagName := fmt.Sprintf("%v.Id", cmdPrefix)
	if cmd.Flags().Changed(IDFlagName) {

		var IDFlagName string
		if cmdPrefix == "" {
			IDFlagName = "Id"
		} else {
			IDFlagName = fmt.Sprintf("%v.Id", cmdPrefix)
		}

		IDFlagValue, err := cmd.Flags().GetString(IDFlagName)
		if err != nil {
			return err, false
		}
		m.ID = IDFlagValue

		retAdded = true
	}

	return nil, retAdded
}

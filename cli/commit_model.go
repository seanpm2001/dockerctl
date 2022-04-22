// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for Commit

// register flags to command
func registerModelCommitFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCommitExpected(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCommitID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCommitExpected(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ExpectedDescription := `Commit ID of external tool expected by dockerd as set at build time.
`

	var ExpectedFlagName string
	if cmdPrefix == "" {
		ExpectedFlagName = "Expected"
	} else {
		ExpectedFlagName = fmt.Sprintf("%v.Expected", cmdPrefix)
	}

	var ExpectedFlagDefault string

	_ = cmd.PersistentFlags().String(ExpectedFlagName, ExpectedFlagDefault, ExpectedDescription)

	return nil
}

func registerCommitID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	IDDescription := `Actual commit ID of external tool.`

	var IDFlagName string
	if cmdPrefix == "" {
		IDFlagName = "ID"
	} else {
		IDFlagName = fmt.Sprintf("%v.ID", cmdPrefix)
	}

	var IDFlagDefault string

	_ = cmd.PersistentFlags().String(IDFlagName, IDFlagDefault, IDDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCommitFlags(depth int, m *models.Commit, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, ExpectedAdded := retrieveCommitExpectedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ExpectedAdded

	err, IDAdded := retrieveCommitIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || IDAdded

	return nil, retAdded
}

func retrieveCommitExpectedFlags(depth int, m *models.Commit, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ExpectedFlagName := fmt.Sprintf("%v.Expected", cmdPrefix)
	if cmd.Flags().Changed(ExpectedFlagName) {

		var ExpectedFlagName string
		if cmdPrefix == "" {
			ExpectedFlagName = "Expected"
		} else {
			ExpectedFlagName = fmt.Sprintf("%v.Expected", cmdPrefix)
		}

		ExpectedFlagValue, err := cmd.Flags().GetString(ExpectedFlagName)
		if err != nil {
			return err, false
		}
		m.Expected = ExpectedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCommitIDFlags(depth int, m *models.Commit, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	IDFlagName := fmt.Sprintf("%v.ID", cmdPrefix)
	if cmd.Flags().Changed(IDFlagName) {

		var IDFlagName string
		if cmdPrefix == "" {
			IDFlagName = "ID"
		} else {
			IDFlagName = fmt.Sprintf("%v.ID", cmdPrefix)
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

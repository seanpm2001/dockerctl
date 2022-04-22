// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/exec"
	"github.com/go-swagger/dockerctl/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationExecExecInspectCmd returns a cmd to handle operation execInspect
func makeOperationExecExecInspectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ExecInspect",
		Short: `Return low-level information about an exec instance.`,
		RunE:  runOperationExecExecInspect,
	}

	if err := registerOperationExecExecInspectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationExecExecInspect uses cmd flags to call endpoint api
func runOperationExecExecInspect(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := exec.NewExecInspectParams()
	if err, _ := retrieveOperationExecExecInspectIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationExecExecInspectResult(appCli.Exec.ExecInspect(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationExecExecInspectParamFlags registers all flags needed to fill params
func registerOperationExecExecInspectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationExecExecInspectIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationExecExecInspectIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	IDDescription := `Required. Exec instance ID`

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

func retrieveOperationExecExecInspectIDFlag(m *exec.ExecInspectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationExecExecInspectResult parses request result and return the string content
func parseOperationExecExecInspectResult(resp0 *exec.ExecInspectOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*exec.ExecInspectOK)
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
		resp1, ok := iResp1.(*exec.ExecInspectNotFound)
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
		resp2, ok := iResp2.(*exec.ExecInspectInternalServerError)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
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

// register flags to command
func registerModelExecInspectOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerExecInspectOKBodyCanRemove(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyContainerID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyDetachKeys(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyExitCode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyOpenStderr(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyOpenStdin(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyOpenStdout(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyPid(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyProcessConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyRunning(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerExecInspectOKBodyCanRemove(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	CanRemoveDescription := ``

	var CanRemoveFlagName string
	if cmdPrefix == "" {
		CanRemoveFlagName = "CanRemove"
	} else {
		CanRemoveFlagName = fmt.Sprintf("%v.CanRemove", cmdPrefix)
	}

	var CanRemoveFlagDefault bool

	_ = cmd.PersistentFlags().Bool(CanRemoveFlagName, CanRemoveFlagDefault, CanRemoveDescription)

	return nil
}

func registerExecInspectOKBodyContainerID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ContainerIDDescription := ``

	var ContainerIDFlagName string
	if cmdPrefix == "" {
		ContainerIDFlagName = "ContainerID"
	} else {
		ContainerIDFlagName = fmt.Sprintf("%v.ContainerID", cmdPrefix)
	}

	var ContainerIDFlagDefault string

	_ = cmd.PersistentFlags().String(ContainerIDFlagName, ContainerIDFlagDefault, ContainerIDDescription)

	return nil
}

func registerExecInspectOKBodyDetachKeys(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	DetachKeysDescription := ``

	var DetachKeysFlagName string
	if cmdPrefix == "" {
		DetachKeysFlagName = "DetachKeys"
	} else {
		DetachKeysFlagName = fmt.Sprintf("%v.DetachKeys", cmdPrefix)
	}

	var DetachKeysFlagDefault string

	_ = cmd.PersistentFlags().String(DetachKeysFlagName, DetachKeysFlagDefault, DetachKeysDescription)

	return nil
}

func registerExecInspectOKBodyExitCode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ExitCodeDescription := ``

	var ExitCodeFlagName string
	if cmdPrefix == "" {
		ExitCodeFlagName = "ExitCode"
	} else {
		ExitCodeFlagName = fmt.Sprintf("%v.ExitCode", cmdPrefix)
	}

	var ExitCodeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(ExitCodeFlagName, ExitCodeFlagDefault, ExitCodeDescription)

	return nil
}

func registerExecInspectOKBodyID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	IDDescription := ``

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

func registerExecInspectOKBodyOpenStderr(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	OpenStderrDescription := ``

	var OpenStderrFlagName string
	if cmdPrefix == "" {
		OpenStderrFlagName = "OpenStderr"
	} else {
		OpenStderrFlagName = fmt.Sprintf("%v.OpenStderr", cmdPrefix)
	}

	var OpenStderrFlagDefault bool

	_ = cmd.PersistentFlags().Bool(OpenStderrFlagName, OpenStderrFlagDefault, OpenStderrDescription)

	return nil
}

func registerExecInspectOKBodyOpenStdin(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	OpenStdinDescription := ``

	var OpenStdinFlagName string
	if cmdPrefix == "" {
		OpenStdinFlagName = "OpenStdin"
	} else {
		OpenStdinFlagName = fmt.Sprintf("%v.OpenStdin", cmdPrefix)
	}

	var OpenStdinFlagDefault bool

	_ = cmd.PersistentFlags().Bool(OpenStdinFlagName, OpenStdinFlagDefault, OpenStdinDescription)

	return nil
}

func registerExecInspectOKBodyOpenStdout(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	OpenStdoutDescription := ``

	var OpenStdoutFlagName string
	if cmdPrefix == "" {
		OpenStdoutFlagName = "OpenStdout"
	} else {
		OpenStdoutFlagName = fmt.Sprintf("%v.OpenStdout", cmdPrefix)
	}

	var OpenStdoutFlagDefault bool

	_ = cmd.PersistentFlags().Bool(OpenStdoutFlagName, OpenStdoutFlagDefault, OpenStdoutDescription)

	return nil
}

func registerExecInspectOKBodyPid(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	PidDescription := `The system process ID for the exec process.`

	var PidFlagName string
	if cmdPrefix == "" {
		PidFlagName = "Pid"
	} else {
		PidFlagName = fmt.Sprintf("%v.Pid", cmdPrefix)
	}

	var PidFlagDefault int64

	_ = cmd.PersistentFlags().Int64(PidFlagName, PidFlagDefault, PidDescription)

	return nil
}

func registerExecInspectOKBodyProcessConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var ProcessConfigFlagName string
	if cmdPrefix == "" {
		ProcessConfigFlagName = "ProcessConfig"
	} else {
		ProcessConfigFlagName = fmt.Sprintf("%v.ProcessConfig", cmdPrefix)
	}

	if err := registerModelProcessConfigFlags(depth+1, ProcessConfigFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerExecInspectOKBodyRunning(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	RunningDescription := ``

	var RunningFlagName string
	if cmdPrefix == "" {
		RunningFlagName = "Running"
	} else {
		RunningFlagName = fmt.Sprintf("%v.Running", cmdPrefix)
	}

	var RunningFlagDefault bool

	_ = cmd.PersistentFlags().Bool(RunningFlagName, RunningFlagDefault, RunningDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelExecInspectOKBodyFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, CanRemoveAdded := retrieveExecInspectOKBodyCanRemoveFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || CanRemoveAdded

	err, ContainerIDAdded := retrieveExecInspectOKBodyContainerIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ContainerIDAdded

	err, DetachKeysAdded := retrieveExecInspectOKBodyDetachKeysFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DetachKeysAdded

	err, ExitCodeAdded := retrieveExecInspectOKBodyExitCodeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ExitCodeAdded

	err, IDAdded := retrieveExecInspectOKBodyIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || IDAdded

	err, OpenStderrAdded := retrieveExecInspectOKBodyOpenStderrFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || OpenStderrAdded

	err, OpenStdinAdded := retrieveExecInspectOKBodyOpenStdinFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || OpenStdinAdded

	err, OpenStdoutAdded := retrieveExecInspectOKBodyOpenStdoutFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || OpenStdoutAdded

	err, PidAdded := retrieveExecInspectOKBodyPidFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PidAdded

	err, ProcessConfigAdded := retrieveExecInspectOKBodyProcessConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ProcessConfigAdded

	err, RunningAdded := retrieveExecInspectOKBodyRunningFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || RunningAdded

	return nil, retAdded
}

func retrieveExecInspectOKBodyCanRemoveFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	CanRemoveFlagName := fmt.Sprintf("%v.CanRemove", cmdPrefix)
	if cmd.Flags().Changed(CanRemoveFlagName) {

		var CanRemoveFlagName string
		if cmdPrefix == "" {
			CanRemoveFlagName = "CanRemove"
		} else {
			CanRemoveFlagName = fmt.Sprintf("%v.CanRemove", cmdPrefix)
		}

		CanRemoveFlagValue, err := cmd.Flags().GetBool(CanRemoveFlagName)
		if err != nil {
			return err, false
		}
		m.CanRemove = CanRemoveFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExecInspectOKBodyContainerIDFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ContainerIDFlagName := fmt.Sprintf("%v.ContainerID", cmdPrefix)
	if cmd.Flags().Changed(ContainerIDFlagName) {

		var ContainerIDFlagName string
		if cmdPrefix == "" {
			ContainerIDFlagName = "ContainerID"
		} else {
			ContainerIDFlagName = fmt.Sprintf("%v.ContainerID", cmdPrefix)
		}

		ContainerIDFlagValue, err := cmd.Flags().GetString(ContainerIDFlagName)
		if err != nil {
			return err, false
		}
		m.ContainerID = ContainerIDFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExecInspectOKBodyDetachKeysFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	DetachKeysFlagName := fmt.Sprintf("%v.DetachKeys", cmdPrefix)
	if cmd.Flags().Changed(DetachKeysFlagName) {

		var DetachKeysFlagName string
		if cmdPrefix == "" {
			DetachKeysFlagName = "DetachKeys"
		} else {
			DetachKeysFlagName = fmt.Sprintf("%v.DetachKeys", cmdPrefix)
		}

		DetachKeysFlagValue, err := cmd.Flags().GetString(DetachKeysFlagName)
		if err != nil {
			return err, false
		}
		m.DetachKeys = DetachKeysFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExecInspectOKBodyExitCodeFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ExitCodeFlagName := fmt.Sprintf("%v.ExitCode", cmdPrefix)
	if cmd.Flags().Changed(ExitCodeFlagName) {

		var ExitCodeFlagName string
		if cmdPrefix == "" {
			ExitCodeFlagName = "ExitCode"
		} else {
			ExitCodeFlagName = fmt.Sprintf("%v.ExitCode", cmdPrefix)
		}

		ExitCodeFlagValue, err := cmd.Flags().GetInt64(ExitCodeFlagName)
		if err != nil {
			return err, false
		}
		m.ExitCode = ExitCodeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExecInspectOKBodyIDFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveExecInspectOKBodyOpenStderrFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	OpenStderrFlagName := fmt.Sprintf("%v.OpenStderr", cmdPrefix)
	if cmd.Flags().Changed(OpenStderrFlagName) {

		var OpenStderrFlagName string
		if cmdPrefix == "" {
			OpenStderrFlagName = "OpenStderr"
		} else {
			OpenStderrFlagName = fmt.Sprintf("%v.OpenStderr", cmdPrefix)
		}

		OpenStderrFlagValue, err := cmd.Flags().GetBool(OpenStderrFlagName)
		if err != nil {
			return err, false
		}
		m.OpenStderr = OpenStderrFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExecInspectOKBodyOpenStdinFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	OpenStdinFlagName := fmt.Sprintf("%v.OpenStdin", cmdPrefix)
	if cmd.Flags().Changed(OpenStdinFlagName) {

		var OpenStdinFlagName string
		if cmdPrefix == "" {
			OpenStdinFlagName = "OpenStdin"
		} else {
			OpenStdinFlagName = fmt.Sprintf("%v.OpenStdin", cmdPrefix)
		}

		OpenStdinFlagValue, err := cmd.Flags().GetBool(OpenStdinFlagName)
		if err != nil {
			return err, false
		}
		m.OpenStdin = OpenStdinFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExecInspectOKBodyOpenStdoutFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	OpenStdoutFlagName := fmt.Sprintf("%v.OpenStdout", cmdPrefix)
	if cmd.Flags().Changed(OpenStdoutFlagName) {

		var OpenStdoutFlagName string
		if cmdPrefix == "" {
			OpenStdoutFlagName = "OpenStdout"
		} else {
			OpenStdoutFlagName = fmt.Sprintf("%v.OpenStdout", cmdPrefix)
		}

		OpenStdoutFlagValue, err := cmd.Flags().GetBool(OpenStdoutFlagName)
		if err != nil {
			return err, false
		}
		m.OpenStdout = OpenStdoutFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExecInspectOKBodyPidFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	PidFlagName := fmt.Sprintf("%v.Pid", cmdPrefix)
	if cmd.Flags().Changed(PidFlagName) {

		var PidFlagName string
		if cmdPrefix == "" {
			PidFlagName = "Pid"
		} else {
			PidFlagName = fmt.Sprintf("%v.Pid", cmdPrefix)
		}

		PidFlagValue, err := cmd.Flags().GetInt64(PidFlagName)
		if err != nil {
			return err, false
		}
		m.Pid = PidFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExecInspectOKBodyProcessConfigFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ProcessConfigFlagName := fmt.Sprintf("%v.ProcessConfig", cmdPrefix)
	if cmd.Flags().Changed(ProcessConfigFlagName) {
		// info: complex object ProcessConfig models.ProcessConfig is retrieved outside this Changed() block
	}
	ProcessConfigFlagValue := m.ProcessConfig
	if swag.IsZero(ProcessConfigFlagValue) {
		ProcessConfigFlagValue = &models.ProcessConfig{}
	}

	err, ProcessConfigAdded := retrieveModelProcessConfigFlags(depth+1, ProcessConfigFlagValue, ProcessConfigFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ProcessConfigAdded
	if ProcessConfigAdded {
		m.ProcessConfig = ProcessConfigFlagValue
	}

	return nil, retAdded
}

func retrieveExecInspectOKBodyRunningFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	RunningFlagName := fmt.Sprintf("%v.Running", cmdPrefix)
	if cmd.Flags().Changed(RunningFlagName) {

		var RunningFlagName string
		if cmdPrefix == "" {
			RunningFlagName = "Running"
		} else {
			RunningFlagName = fmt.Sprintf("%v.Running", cmdPrefix)
		}

		RunningFlagValue, err := cmd.Flags().GetBool(RunningFlagName)
		if err != nil {
			return err, false
		}
		m.Running = RunningFlagValue

		retAdded = true
	}

	return nil, retAdded
}

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

// makeOperationSystemSystemVersionCmd returns a cmd to handle operation systemVersion
func makeOperationSystemSystemVersionCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "SystemVersion",
		Short: `Returns the version of Docker that is running and various information about the system that Docker is running on.`,
		RunE:  runOperationSystemSystemVersion,
	}

	if err := registerOperationSystemSystemVersionParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSystemSystemVersion uses cmd flags to call endpoint api
func runOperationSystemSystemVersion(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := system.NewSystemVersionParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSystemSystemVersionResult(appCli.System.SystemVersion(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSystemSystemVersionParamFlags registers all flags needed to fill params
func registerOperationSystemSystemVersionParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationSystemSystemVersionResult parses request result and return the string content
func parseOperationSystemSystemVersionResult(resp0 *system.SystemVersionOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*system.SystemVersionOK)
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
		resp1, ok := iResp1.(*system.SystemVersionInternalServerError)
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

// register flags to command
func registerModelComponentVersionFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerComponentVersionDetails(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerComponentVersionName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerComponentVersionVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerComponentVersionDetails(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Details interface{} map type is not supported by go-swagger cli yet

	return nil
}

func registerComponentVersionName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. `

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "Name"
	} else {
		nameFlagName = fmt.Sprintf("%v.Name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerComponentVersionVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	versionDescription := `Required. `

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "Version"
	} else {
		versionFlagName = fmt.Sprintf("%v.Version", cmdPrefix)
	}

	var versionFlagDefault string

	_ = cmd.PersistentFlags().String(versionFlagName, versionFlagDefault, versionDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelComponentVersionFlags(depth int, m *system.ComponentVersion, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, detailsAdded := retrieveComponentVersionDetailsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || detailsAdded

	err, nameAdded := retrieveComponentVersionNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, versionAdded := retrieveComponentVersionVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	return nil, retAdded
}

func retrieveComponentVersionDetailsFlags(depth int, m *system.ComponentVersion, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	detailsFlagName := fmt.Sprintf("%v.Details", cmdPrefix)
	if cmd.Flags().Changed(detailsFlagName) {
		// warning: Details map type interface{} is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveComponentVersionNameFlags(depth int, m *system.ComponentVersion, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.Name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "Name"
		} else {
			nameFlagName = fmt.Sprintf("%v.Name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = &nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveComponentVersionVersionFlags(depth int, m *system.ComponentVersion, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	versionFlagName := fmt.Sprintf("%v.Version", cmdPrefix)
	if cmd.Flags().Changed(versionFlagName) {

		var versionFlagName string
		if cmdPrefix == "" {
			versionFlagName = "Version"
		} else {
			versionFlagName = fmt.Sprintf("%v.Version", cmdPrefix)
		}

		versionFlagValue, err := cmd.Flags().GetString(versionFlagName)
		if err != nil {
			return err, false
		}
		m.Version = versionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// register flags to command
func registerModelSystemVersionOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSystemVersionOKBodyAPIVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemVersionOKBodyArch(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemVersionOKBodyBuildTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemVersionOKBodyComponents(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemVersionOKBodyExperimental(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemVersionOKBodyGitCommit(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemVersionOKBodyGoVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemVersionOKBodyKernelVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemVersionOKBodyMinAPIVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemVersionOKBodyOs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemVersionOKBodyPlatform(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemVersionOKBodyVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSystemVersionOKBodyAPIVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	apiVersionDescription := ``

	var apiVersionFlagName string
	if cmdPrefix == "" {
		apiVersionFlagName = "ApiVersion"
	} else {
		apiVersionFlagName = fmt.Sprintf("%v.ApiVersion", cmdPrefix)
	}

	var apiVersionFlagDefault string

	_ = cmd.PersistentFlags().String(apiVersionFlagName, apiVersionFlagDefault, apiVersionDescription)

	return nil
}

func registerSystemVersionOKBodyArch(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	archDescription := ``

	var archFlagName string
	if cmdPrefix == "" {
		archFlagName = "Arch"
	} else {
		archFlagName = fmt.Sprintf("%v.Arch", cmdPrefix)
	}

	var archFlagDefault string

	_ = cmd.PersistentFlags().String(archFlagName, archFlagDefault, archDescription)

	return nil
}

func registerSystemVersionOKBodyBuildTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	buildTimeDescription := ``

	var buildTimeFlagName string
	if cmdPrefix == "" {
		buildTimeFlagName = "BuildTime"
	} else {
		buildTimeFlagName = fmt.Sprintf("%v.BuildTime", cmdPrefix)
	}

	var buildTimeFlagDefault string

	_ = cmd.PersistentFlags().String(buildTimeFlagName, buildTimeFlagDefault, buildTimeDescription)

	return nil
}

func registerSystemVersionOKBodyComponents(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Components []*ComponentVersion array type is not supported by go-swagger cli yet

	return nil
}

func registerSystemVersionOKBodyExperimental(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	experimentalDescription := ``

	var experimentalFlagName string
	if cmdPrefix == "" {
		experimentalFlagName = "Experimental"
	} else {
		experimentalFlagName = fmt.Sprintf("%v.Experimental", cmdPrefix)
	}

	var experimentalFlagDefault bool

	_ = cmd.PersistentFlags().Bool(experimentalFlagName, experimentalFlagDefault, experimentalDescription)

	return nil
}

func registerSystemVersionOKBodyGitCommit(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	gitCommitDescription := ``

	var gitCommitFlagName string
	if cmdPrefix == "" {
		gitCommitFlagName = "GitCommit"
	} else {
		gitCommitFlagName = fmt.Sprintf("%v.GitCommit", cmdPrefix)
	}

	var gitCommitFlagDefault string

	_ = cmd.PersistentFlags().String(gitCommitFlagName, gitCommitFlagDefault, gitCommitDescription)

	return nil
}

func registerSystemVersionOKBodyGoVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	goVersionDescription := ``

	var goVersionFlagName string
	if cmdPrefix == "" {
		goVersionFlagName = "GoVersion"
	} else {
		goVersionFlagName = fmt.Sprintf("%v.GoVersion", cmdPrefix)
	}

	var goVersionFlagDefault string

	_ = cmd.PersistentFlags().String(goVersionFlagName, goVersionFlagDefault, goVersionDescription)

	return nil
}

func registerSystemVersionOKBodyKernelVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	kernelVersionDescription := ``

	var kernelVersionFlagName string
	if cmdPrefix == "" {
		kernelVersionFlagName = "KernelVersion"
	} else {
		kernelVersionFlagName = fmt.Sprintf("%v.KernelVersion", cmdPrefix)
	}

	var kernelVersionFlagDefault string

	_ = cmd.PersistentFlags().String(kernelVersionFlagName, kernelVersionFlagDefault, kernelVersionDescription)

	return nil
}

func registerSystemVersionOKBodyMinAPIVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	minApiVersionDescription := ``

	var minApiVersionFlagName string
	if cmdPrefix == "" {
		minApiVersionFlagName = "MinAPIVersion"
	} else {
		minApiVersionFlagName = fmt.Sprintf("%v.MinAPIVersion", cmdPrefix)
	}

	var minApiVersionFlagDefault string

	_ = cmd.PersistentFlags().String(minApiVersionFlagName, minApiVersionFlagDefault, minApiVersionDescription)

	return nil
}

func registerSystemVersionOKBodyOs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	osDescription := ``

	var osFlagName string
	if cmdPrefix == "" {
		osFlagName = "Os"
	} else {
		osFlagName = fmt.Sprintf("%v.Os", cmdPrefix)
	}

	var osFlagDefault string

	_ = cmd.PersistentFlags().String(osFlagName, osFlagDefault, osDescription)

	return nil
}

func registerSystemVersionOKBodyPlatform(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var platformFlagName string
	if cmdPrefix == "" {
		platformFlagName = "Platform"
	} else {
		platformFlagName = fmt.Sprintf("%v.Platform", cmdPrefix)
	}

	if err := registerModelSystemVersionOKBodyPlatformFlags(depth+1, platformFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSystemVersionOKBodyVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	versionDescription := ``

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "Version"
	} else {
		versionFlagName = fmt.Sprintf("%v.Version", cmdPrefix)
	}

	var versionFlagDefault string

	_ = cmd.PersistentFlags().String(versionFlagName, versionFlagDefault, versionDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSystemVersionOKBodyFlags(depth int, m *system.SystemVersionOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, apiVersionAdded := retrieveSystemVersionOKBodyAPIVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || apiVersionAdded

	err, archAdded := retrieveSystemVersionOKBodyArchFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || archAdded

	err, buildTimeAdded := retrieveSystemVersionOKBodyBuildTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || buildTimeAdded

	err, componentsAdded := retrieveSystemVersionOKBodyComponentsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || componentsAdded

	err, experimentalAdded := retrieveSystemVersionOKBodyExperimentalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || experimentalAdded

	err, gitCommitAdded := retrieveSystemVersionOKBodyGitCommitFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || gitCommitAdded

	err, goVersionAdded := retrieveSystemVersionOKBodyGoVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || goVersionAdded

	err, kernelVersionAdded := retrieveSystemVersionOKBodyKernelVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || kernelVersionAdded

	err, minApiVersionAdded := retrieveSystemVersionOKBodyMinAPIVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || minApiVersionAdded

	err, osAdded := retrieveSystemVersionOKBodyOsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || osAdded

	err, platformAdded := retrieveSystemVersionOKBodyPlatformFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || platformAdded

	err, versionAdded := retrieveSystemVersionOKBodyVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	return nil, retAdded
}

func retrieveSystemVersionOKBodyAPIVersionFlags(depth int, m *system.SystemVersionOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	apiVersionFlagName := fmt.Sprintf("%v.ApiVersion", cmdPrefix)
	if cmd.Flags().Changed(apiVersionFlagName) {

		var apiVersionFlagName string
		if cmdPrefix == "" {
			apiVersionFlagName = "ApiVersion"
		} else {
			apiVersionFlagName = fmt.Sprintf("%v.ApiVersion", cmdPrefix)
		}

		apiVersionFlagValue, err := cmd.Flags().GetString(apiVersionFlagName)
		if err != nil {
			return err, false
		}
		m.APIVersion = apiVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSystemVersionOKBodyArchFlags(depth int, m *system.SystemVersionOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	archFlagName := fmt.Sprintf("%v.Arch", cmdPrefix)
	if cmd.Flags().Changed(archFlagName) {

		var archFlagName string
		if cmdPrefix == "" {
			archFlagName = "Arch"
		} else {
			archFlagName = fmt.Sprintf("%v.Arch", cmdPrefix)
		}

		archFlagValue, err := cmd.Flags().GetString(archFlagName)
		if err != nil {
			return err, false
		}
		m.Arch = archFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSystemVersionOKBodyBuildTimeFlags(depth int, m *system.SystemVersionOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	buildTimeFlagName := fmt.Sprintf("%v.BuildTime", cmdPrefix)
	if cmd.Flags().Changed(buildTimeFlagName) {

		var buildTimeFlagName string
		if cmdPrefix == "" {
			buildTimeFlagName = "BuildTime"
		} else {
			buildTimeFlagName = fmt.Sprintf("%v.BuildTime", cmdPrefix)
		}

		buildTimeFlagValue, err := cmd.Flags().GetString(buildTimeFlagName)
		if err != nil {
			return err, false
		}
		m.BuildTime = buildTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSystemVersionOKBodyComponentsFlags(depth int, m *system.SystemVersionOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	componentsFlagName := fmt.Sprintf("%v.Components", cmdPrefix)
	if cmd.Flags().Changed(componentsFlagName) {
		// warning: Components array type []*ComponentVersion is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveSystemVersionOKBodyExperimentalFlags(depth int, m *system.SystemVersionOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	experimentalFlagName := fmt.Sprintf("%v.Experimental", cmdPrefix)
	if cmd.Flags().Changed(experimentalFlagName) {

		var experimentalFlagName string
		if cmdPrefix == "" {
			experimentalFlagName = "Experimental"
		} else {
			experimentalFlagName = fmt.Sprintf("%v.Experimental", cmdPrefix)
		}

		experimentalFlagValue, err := cmd.Flags().GetBool(experimentalFlagName)
		if err != nil {
			return err, false
		}
		m.Experimental = experimentalFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSystemVersionOKBodyGitCommitFlags(depth int, m *system.SystemVersionOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	gitCommitFlagName := fmt.Sprintf("%v.GitCommit", cmdPrefix)
	if cmd.Flags().Changed(gitCommitFlagName) {

		var gitCommitFlagName string
		if cmdPrefix == "" {
			gitCommitFlagName = "GitCommit"
		} else {
			gitCommitFlagName = fmt.Sprintf("%v.GitCommit", cmdPrefix)
		}

		gitCommitFlagValue, err := cmd.Flags().GetString(gitCommitFlagName)
		if err != nil {
			return err, false
		}
		m.GitCommit = gitCommitFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSystemVersionOKBodyGoVersionFlags(depth int, m *system.SystemVersionOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	goVersionFlagName := fmt.Sprintf("%v.GoVersion", cmdPrefix)
	if cmd.Flags().Changed(goVersionFlagName) {

		var goVersionFlagName string
		if cmdPrefix == "" {
			goVersionFlagName = "GoVersion"
		} else {
			goVersionFlagName = fmt.Sprintf("%v.GoVersion", cmdPrefix)
		}

		goVersionFlagValue, err := cmd.Flags().GetString(goVersionFlagName)
		if err != nil {
			return err, false
		}
		m.GoVersion = goVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSystemVersionOKBodyKernelVersionFlags(depth int, m *system.SystemVersionOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	kernelVersionFlagName := fmt.Sprintf("%v.KernelVersion", cmdPrefix)
	if cmd.Flags().Changed(kernelVersionFlagName) {

		var kernelVersionFlagName string
		if cmdPrefix == "" {
			kernelVersionFlagName = "KernelVersion"
		} else {
			kernelVersionFlagName = fmt.Sprintf("%v.KernelVersion", cmdPrefix)
		}

		kernelVersionFlagValue, err := cmd.Flags().GetString(kernelVersionFlagName)
		if err != nil {
			return err, false
		}
		m.KernelVersion = kernelVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSystemVersionOKBodyMinAPIVersionFlags(depth int, m *system.SystemVersionOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	minApiVersionFlagName := fmt.Sprintf("%v.MinAPIVersion", cmdPrefix)
	if cmd.Flags().Changed(minApiVersionFlagName) {

		var minApiVersionFlagName string
		if cmdPrefix == "" {
			minApiVersionFlagName = "MinAPIVersion"
		} else {
			minApiVersionFlagName = fmt.Sprintf("%v.MinAPIVersion", cmdPrefix)
		}

		minApiVersionFlagValue, err := cmd.Flags().GetString(minApiVersionFlagName)
		if err != nil {
			return err, false
		}
		m.MinAPIVersion = minApiVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSystemVersionOKBodyOsFlags(depth int, m *system.SystemVersionOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	osFlagName := fmt.Sprintf("%v.Os", cmdPrefix)
	if cmd.Flags().Changed(osFlagName) {

		var osFlagName string
		if cmdPrefix == "" {
			osFlagName = "Os"
		} else {
			osFlagName = fmt.Sprintf("%v.Os", cmdPrefix)
		}

		osFlagValue, err := cmd.Flags().GetString(osFlagName)
		if err != nil {
			return err, false
		}
		m.Os = osFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSystemVersionOKBodyPlatformFlags(depth int, m *system.SystemVersionOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	platformFlagName := fmt.Sprintf("%v.Platform", cmdPrefix)
	if cmd.Flags().Changed(platformFlagName) {
		// info: complex object Platform SystemVersionOKBodyPlatform is retrieved outside this Changed() block
	}
	platformFlagValue := m.Platform
	if swag.IsZero(platformFlagValue) {
		platformFlagValue = &system.SystemVersionOKBodyPlatform{}
	}

	err, platformAdded := retrieveModelSystemVersionOKBodyPlatformFlags(depth+1, platformFlagValue, platformFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || platformAdded
	if platformAdded {
		m.Platform = platformFlagValue
	}

	return nil, retAdded
}

func retrieveSystemVersionOKBodyVersionFlags(depth int, m *system.SystemVersionOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	versionFlagName := fmt.Sprintf("%v.Version", cmdPrefix)
	if cmd.Flags().Changed(versionFlagName) {

		var versionFlagName string
		if cmdPrefix == "" {
			versionFlagName = "Version"
		} else {
			versionFlagName = fmt.Sprintf("%v.Version", cmdPrefix)
		}

		versionFlagValue, err := cmd.Flags().GetString(versionFlagName)
		if err != nil {
			return err, false
		}
		m.Version = versionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// register flags to command
func registerModelSystemVersionOKBodyPlatformFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSystemVersionOKBodyPlatformName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSystemVersionOKBodyPlatformName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. `

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "Name"
	} else {
		nameFlagName = fmt.Sprintf("%v.Name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSystemVersionOKBodyPlatformFlags(depth int, m *system.SystemVersionOKBodyPlatform, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, nameAdded := retrieveSystemVersionOKBodyPlatformNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	return nil, retAdded
}

func retrieveSystemVersionOKBodyPlatformNameFlags(depth int, m *system.SystemVersionOKBodyPlatform, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.Name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "Name"
		} else {
			nameFlagName = fmt.Sprintf("%v.Name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = &nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

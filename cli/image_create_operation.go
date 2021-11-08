// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/image"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationImageImageCreateCmd returns a cmd to handle operation imageCreate
func makeOperationImageImageCreateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ImageCreate",
		Short: `Create an image by either pulling it from a registry or importing it.`,
		RunE:  runOperationImageImageCreate,
	}

	if err := registerOperationImageImageCreateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationImageImageCreate uses cmd flags to call endpoint api
func runOperationImageImageCreate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := image.NewImageCreateParams()
	if err, _ := retrieveOperationImageImageCreateXRegistryAuthFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageCreateFromImageFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageCreateFromSrcFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageCreateInputImageFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageCreatePlatformFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageCreateRepoFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageCreateTagFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationImageImageCreateResult(appCli.Image.ImageCreate(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationImageImageCreateParamFlags registers all flags needed to fill params
func registerOperationImageImageCreateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationImageImageCreateXRegistryAuthParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageCreateFromImageParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageCreateFromSrcParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageCreateInputImageParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageCreatePlatformParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageCreateRepoParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageCreateTagParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationImageImageCreateXRegistryAuthParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	xRegistryAuthDescription := `A base64-encoded auth configuration. [See the authentication section for details.](#section/Authentication)`

	var xRegistryAuthFlagName string
	if cmdPrefix == "" {
		xRegistryAuthFlagName = "X-Registry-Auth"
	} else {
		xRegistryAuthFlagName = fmt.Sprintf("%v.X-Registry-Auth", cmdPrefix)
	}

	var xRegistryAuthFlagDefault string

	_ = cmd.PersistentFlags().String(xRegistryAuthFlagName, xRegistryAuthFlagDefault, xRegistryAuthDescription)

	return nil
}
func registerOperationImageImageCreateFromImageParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	fromImageDescription := `Name of the image to pull. The name may include a tag or digest. This parameter may only be used when pulling an image. The pull is cancelled if the HTTP connection is closed.`

	var fromImageFlagName string
	if cmdPrefix == "" {
		fromImageFlagName = "fromImage"
	} else {
		fromImageFlagName = fmt.Sprintf("%v.fromImage", cmdPrefix)
	}

	var fromImageFlagDefault string

	_ = cmd.PersistentFlags().String(fromImageFlagName, fromImageFlagDefault, fromImageDescription)

	return nil
}
func registerOperationImageImageCreateFromSrcParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	fromSrcDescription := `Source to import. The value may be a URL from which the image can be retrieved or ` + "`" + `-` + "`" + ` to read the image from the request body. This parameter may only be used when importing an image.`

	var fromSrcFlagName string
	if cmdPrefix == "" {
		fromSrcFlagName = "fromSrc"
	} else {
		fromSrcFlagName = fmt.Sprintf("%v.fromSrc", cmdPrefix)
	}

	var fromSrcFlagDefault string

	_ = cmd.PersistentFlags().String(fromSrcFlagName, fromSrcFlagDefault, fromSrcDescription)

	return nil
}
func registerOperationImageImageCreateInputImageParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	inputImageDescription := `Image content if the value ` + "`" + `-` + "`" + ` has been specified in fromSrc query parameter`

	var inputImageFlagName string
	if cmdPrefix == "" {
		inputImageFlagName = "inputImage"
	} else {
		inputImageFlagName = fmt.Sprintf("%v.inputImage", cmdPrefix)
	}

	var inputImageFlagDefault string

	_ = cmd.PersistentFlags().String(inputImageFlagName, inputImageFlagDefault, inputImageDescription)

	return nil
}
func registerOperationImageImageCreatePlatformParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	platformDescription := `Platform in the format os[/arch[/variant]]`

	var platformFlagName string
	if cmdPrefix == "" {
		platformFlagName = "platform"
	} else {
		platformFlagName = fmt.Sprintf("%v.platform", cmdPrefix)
	}

	var platformFlagDefault string

	_ = cmd.PersistentFlags().String(platformFlagName, platformFlagDefault, platformDescription)

	return nil
}
func registerOperationImageImageCreateRepoParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	repoDescription := `Repository name given to an image when it is imported. The repo may include a tag. This parameter may only be used when importing an image.`

	var repoFlagName string
	if cmdPrefix == "" {
		repoFlagName = "repo"
	} else {
		repoFlagName = fmt.Sprintf("%v.repo", cmdPrefix)
	}

	var repoFlagDefault string

	_ = cmd.PersistentFlags().String(repoFlagName, repoFlagDefault, repoDescription)

	return nil
}
func registerOperationImageImageCreateTagParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tagDescription := `Tag or digest. If empty when pulling an image, this causes all tags for the given image to be pulled.`

	var tagFlagName string
	if cmdPrefix == "" {
		tagFlagName = "tag"
	} else {
		tagFlagName = fmt.Sprintf("%v.tag", cmdPrefix)
	}

	var tagFlagDefault string

	_ = cmd.PersistentFlags().String(tagFlagName, tagFlagDefault, tagDescription)

	return nil
}

func retrieveOperationImageImageCreateXRegistryAuthFlag(m *image.ImageCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("X-Registry-Auth") {

		var xRegistryAuthFlagName string
		if cmdPrefix == "" {
			xRegistryAuthFlagName = "X-Registry-Auth"
		} else {
			xRegistryAuthFlagName = fmt.Sprintf("%v.X-Registry-Auth", cmdPrefix)
		}

		xRegistryAuthFlagValue, err := cmd.Flags().GetString(xRegistryAuthFlagName)
		if err != nil {
			return err, false
		}
		m.XRegistryAuth = &xRegistryAuthFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageCreateFromImageFlag(m *image.ImageCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("fromImage") {

		var fromImageFlagName string
		if cmdPrefix == "" {
			fromImageFlagName = "fromImage"
		} else {
			fromImageFlagName = fmt.Sprintf("%v.fromImage", cmdPrefix)
		}

		fromImageFlagValue, err := cmd.Flags().GetString(fromImageFlagName)
		if err != nil {
			return err, false
		}
		m.FromImage = &fromImageFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageCreateFromSrcFlag(m *image.ImageCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("fromSrc") {

		var fromSrcFlagName string
		if cmdPrefix == "" {
			fromSrcFlagName = "fromSrc"
		} else {
			fromSrcFlagName = fmt.Sprintf("%v.fromSrc", cmdPrefix)
		}

		fromSrcFlagValue, err := cmd.Flags().GetString(fromSrcFlagName)
		if err != nil {
			return err, false
		}
		m.FromSrc = &fromSrcFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageCreateInputImageFlag(m *image.ImageCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("inputImage") {

		var inputImageFlagName string
		if cmdPrefix == "" {
			inputImageFlagName = "inputImage"
		} else {
			inputImageFlagName = fmt.Sprintf("%v.inputImage", cmdPrefix)
		}

		inputImageFlagValue, err := cmd.Flags().GetString(inputImageFlagName)
		if err != nil {
			return err, false
		}
		m.InputImage = inputImageFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageCreatePlatformFlag(m *image.ImageCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("platform") {

		var platformFlagName string
		if cmdPrefix == "" {
			platformFlagName = "platform"
		} else {
			platformFlagName = fmt.Sprintf("%v.platform", cmdPrefix)
		}

		platformFlagValue, err := cmd.Flags().GetString(platformFlagName)
		if err != nil {
			return err, false
		}
		m.Platform = &platformFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageCreateRepoFlag(m *image.ImageCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("repo") {

		var repoFlagName string
		if cmdPrefix == "" {
			repoFlagName = "repo"
		} else {
			repoFlagName = fmt.Sprintf("%v.repo", cmdPrefix)
		}

		repoFlagValue, err := cmd.Flags().GetString(repoFlagName)
		if err != nil {
			return err, false
		}
		m.Repo = &repoFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageCreateTagFlag(m *image.ImageCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("tag") {

		var tagFlagName string
		if cmdPrefix == "" {
			tagFlagName = "tag"
		} else {
			tagFlagName = fmt.Sprintf("%v.tag", cmdPrefix)
		}

		tagFlagValue, err := cmd.Flags().GetString(tagFlagName)
		if err != nil {
			return err, false
		}
		m.Tag = &tagFlagValue

	}
	return nil, retAdded
}

// parseOperationImageImageCreateResult parses request result and return the string content
func parseOperationImageImageCreateResult(resp0 *image.ImageCreateOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning imageCreateOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*image.ImageCreateNotFound)
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
		resp2, ok := iResp2.(*image.ImageCreateInternalServerError)
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

	// warning: non schema response imageCreateOK is not supported by go-swagger cli yet.

	return "", nil
}

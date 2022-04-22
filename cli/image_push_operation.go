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

// makeOperationImageImagePushCmd returns a cmd to handle operation imagePush
func makeOperationImageImagePushCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "ImagePush",
		Short: `Push an image to a registry.

If you wish to push an image on to a private registry, that image must already have a tag which references the registry. For example, ` + "`" + `registry.example.com/myimage:latest` + "`" + `.

The push is cancelled if the HTTP connection is closed.
`,
		RunE: runOperationImageImagePush,
	}

	if err := registerOperationImageImagePushParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationImageImagePush uses cmd flags to call endpoint api
func runOperationImageImagePush(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := image.NewImagePushParams()
	if err, _ := retrieveOperationImageImagePushXRegistryAuthFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImagePushNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImagePushTagFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationImageImagePushResult(appCli.Image.ImagePush(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationImageImagePushParamFlags registers all flags needed to fill params
func registerOperationImageImagePushParamFlags(cmd *cobra.Command) error {
	if err := registerOperationImageImagePushXRegistryAuthParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImagePushNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImagePushTagParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationImageImagePushXRegistryAuthParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	XRegistryAuthDescription := `Required. A base64-encoded auth configuration. [See the authentication section for details.](#section/Authentication)`

	var XRegistryAuthFlagName string
	if cmdPrefix == "" {
		XRegistryAuthFlagName = "X-Registry-Auth"
	} else {
		XRegistryAuthFlagName = fmt.Sprintf("%v.X-Registry-Auth", cmdPrefix)
	}

	var XRegistryAuthFlagDefault string

	_ = cmd.PersistentFlags().String(XRegistryAuthFlagName, XRegistryAuthFlagDefault, XRegistryAuthDescription)

	return nil
}
func registerOperationImageImagePushNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	NameDescription := `Required. Image name or ID.`

	var NameFlagName string
	if cmdPrefix == "" {
		NameFlagName = "name"
	} else {
		NameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var NameFlagDefault string

	_ = cmd.PersistentFlags().String(NameFlagName, NameFlagDefault, NameDescription)

	return nil
}
func registerOperationImageImagePushTagParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	TagDescription := `The tag to associate with the image on the registry.`

	var TagFlagName string
	if cmdPrefix == "" {
		TagFlagName = "tag"
	} else {
		TagFlagName = fmt.Sprintf("%v.tag", cmdPrefix)
	}

	var TagFlagDefault string

	_ = cmd.PersistentFlags().String(TagFlagName, TagFlagDefault, TagDescription)

	return nil
}

func retrieveOperationImageImagePushXRegistryAuthFlag(m *image.ImagePushParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("X-Registry-Auth") {

		var XRegistryAuthFlagName string
		if cmdPrefix == "" {
			XRegistryAuthFlagName = "X-Registry-Auth"
		} else {
			XRegistryAuthFlagName = fmt.Sprintf("%v.X-Registry-Auth", cmdPrefix)
		}

		XRegistryAuthFlagValue, err := cmd.Flags().GetString(XRegistryAuthFlagName)
		if err != nil {
			return err, false
		}
		m.XRegistryAuth = XRegistryAuthFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImagePushNameFlag(m *image.ImagePushParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("name") {

		var NameFlagName string
		if cmdPrefix == "" {
			NameFlagName = "name"
		} else {
			NameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		NameFlagValue, err := cmd.Flags().GetString(NameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = NameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImagePushTagFlag(m *image.ImagePushParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("tag") {

		var TagFlagName string
		if cmdPrefix == "" {
			TagFlagName = "tag"
		} else {
			TagFlagName = fmt.Sprintf("%v.tag", cmdPrefix)
		}

		TagFlagValue, err := cmd.Flags().GetString(TagFlagName)
		if err != nil {
			return err, false
		}
		m.Tag = &TagFlagValue

	}
	return nil, retAdded
}

// parseOperationImageImagePushResult parses request result and return the string content
func parseOperationImageImagePushResult(resp0 *image.ImagePushOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning imagePushOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*image.ImagePushNotFound)
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
		resp2, ok := iResp2.(*image.ImagePushInternalServerError)
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

	// warning: non schema response imagePushOK is not supported by go-swagger cli yet.

	return "", nil
}

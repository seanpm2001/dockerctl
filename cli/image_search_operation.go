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

// makeOperationImageImageSearchCmd returns a cmd to handle operation imageSearch
func makeOperationImageImageSearchCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ImageSearch",
		Short: `Search for an image on Docker Hub.`,
		RunE:  runOperationImageImageSearch,
	}

	if err := registerOperationImageImageSearchParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationImageImageSearch uses cmd flags to call endpoint api
func runOperationImageImageSearch(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := image.NewImageSearchParams()
	if err, _ := retrieveOperationImageImageSearchFiltersFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageSearchLimitFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageSearchTermFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationImageImageSearchResult(appCli.Image.ImageSearch(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationImageImageSearchParamFlags registers all flags needed to fill params
func registerOperationImageImageSearchParamFlags(cmd *cobra.Command) error {
	if err := registerOperationImageImageSearchFiltersParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageSearchLimitParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageSearchTermParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationImageImageSearchFiltersParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filtersDescription := `A JSON encoded value of the filters (a ` + "`" + `map[string][]string` + "`" + `) to process on the images list. Available filters:

- ` + "`" + `is-automated=(true|false)` + "`" + `
- ` + "`" + `is-official=(true|false)` + "`" + `
- ` + "`" + `stars=<number>` + "`" + ` Matches images that has at least 'number' stars.
`

	var filtersFlagName string
	if cmdPrefix == "" {
		filtersFlagName = "filters"
	} else {
		filtersFlagName = fmt.Sprintf("%v.filters", cmdPrefix)
	}

	var filtersFlagDefault string

	_ = cmd.PersistentFlags().String(filtersFlagName, filtersFlagDefault, filtersDescription)

	return nil
}
func registerOperationImageImageSearchLimitParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	limitDescription := `Maximum number of results to return`

	var limitFlagName string
	if cmdPrefix == "" {
		limitFlagName = "limit"
	} else {
		limitFlagName = fmt.Sprintf("%v.limit", cmdPrefix)
	}

	var limitFlagDefault int64

	_ = cmd.PersistentFlags().Int64(limitFlagName, limitFlagDefault, limitDescription)

	return nil
}
func registerOperationImageImageSearchTermParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	termDescription := `Required. Term to search`

	var termFlagName string
	if cmdPrefix == "" {
		termFlagName = "term"
	} else {
		termFlagName = fmt.Sprintf("%v.term", cmdPrefix)
	}

	var termFlagDefault string

	_ = cmd.PersistentFlags().String(termFlagName, termFlagDefault, termDescription)

	return nil
}

func retrieveOperationImageImageSearchFiltersFlag(m *image.ImageSearchParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filters") {

		var filtersFlagName string
		if cmdPrefix == "" {
			filtersFlagName = "filters"
		} else {
			filtersFlagName = fmt.Sprintf("%v.filters", cmdPrefix)
		}

		filtersFlagValue, err := cmd.Flags().GetString(filtersFlagName)
		if err != nil {
			return err, false
		}
		m.Filters = &filtersFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageSearchLimitFlag(m *image.ImageSearchParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("limit") {

		var limitFlagName string
		if cmdPrefix == "" {
			limitFlagName = "limit"
		} else {
			limitFlagName = fmt.Sprintf("%v.limit", cmdPrefix)
		}

		limitFlagValue, err := cmd.Flags().GetInt64(limitFlagName)
		if err != nil {
			return err, false
		}
		m.Limit = &limitFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageSearchTermFlag(m *image.ImageSearchParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("term") {

		var termFlagName string
		if cmdPrefix == "" {
			termFlagName = "term"
		} else {
			termFlagName = fmt.Sprintf("%v.term", cmdPrefix)
		}

		termFlagValue, err := cmd.Flags().GetString(termFlagName)
		if err != nil {
			return err, false
		}
		m.Term = termFlagValue

	}
	return nil, retAdded
}

// parseOperationImageImageSearchResult parses request result and return the string content
func parseOperationImageImageSearchResult(resp0 *image.ImageSearchOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*image.ImageSearchOK)
		if ok {
			if !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*image.ImageSearchInternalServerError)
		if ok {
			if !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}

// register flags to command
func registerModelImageSearchOKBodyItems0Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerImageSearchOKBodyItems0Description(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageSearchOKBodyItems0IsAutomated(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageSearchOKBodyItems0IsOfficial(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageSearchOKBodyItems0Name(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageSearchOKBodyItems0StarCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageSearchOKBodyItems0Description(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := ``

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerImageSearchOKBodyItems0IsAutomated(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isAutomatedDescription := ``

	var isAutomatedFlagName string
	if cmdPrefix == "" {
		isAutomatedFlagName = "is_automated"
	} else {
		isAutomatedFlagName = fmt.Sprintf("%v.is_automated", cmdPrefix)
	}

	var isAutomatedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isAutomatedFlagName, isAutomatedFlagDefault, isAutomatedDescription)

	return nil
}

func registerImageSearchOKBodyItems0IsOfficial(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isOfficialDescription := ``

	var isOfficialFlagName string
	if cmdPrefix == "" {
		isOfficialFlagName = "is_official"
	} else {
		isOfficialFlagName = fmt.Sprintf("%v.is_official", cmdPrefix)
	}

	var isOfficialFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isOfficialFlagName, isOfficialFlagDefault, isOfficialDescription)

	return nil
}

func registerImageSearchOKBodyItems0Name(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

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

func registerImageSearchOKBodyItems0StarCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	starCountDescription := ``

	var starCountFlagName string
	if cmdPrefix == "" {
		starCountFlagName = "star_count"
	} else {
		starCountFlagName = fmt.Sprintf("%v.star_count", cmdPrefix)
	}

	var starCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(starCountFlagName, starCountFlagDefault, starCountDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelImageSearchOKBodyItems0Flags(depth int, m *image.ImageSearchOKBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrieveImageSearchOKBodyItems0DescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, isAutomatedAdded := retrieveImageSearchOKBodyItems0IsAutomatedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isAutomatedAdded

	err, isOfficialAdded := retrieveImageSearchOKBodyItems0IsOfficialFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isOfficialAdded

	err, nameAdded := retrieveImageSearchOKBodyItems0NameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, starCountAdded := retrieveImageSearchOKBodyItems0StarCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || starCountAdded

	return nil, retAdded
}

func retrieveImageSearchOKBodyItems0DescriptionFlags(depth int, m *image.ImageSearchOKBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageSearchOKBodyItems0IsAutomatedFlags(depth int, m *image.ImageSearchOKBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isAutomatedFlagName := fmt.Sprintf("%v.is_automated", cmdPrefix)
	if cmd.Flags().Changed(isAutomatedFlagName) {

		var isAutomatedFlagName string
		if cmdPrefix == "" {
			isAutomatedFlagName = "is_automated"
		} else {
			isAutomatedFlagName = fmt.Sprintf("%v.is_automated", cmdPrefix)
		}

		isAutomatedFlagValue, err := cmd.Flags().GetBool(isAutomatedFlagName)
		if err != nil {
			return err, false
		}
		m.IsAutomated = isAutomatedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageSearchOKBodyItems0IsOfficialFlags(depth int, m *image.ImageSearchOKBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isOfficialFlagName := fmt.Sprintf("%v.is_official", cmdPrefix)
	if cmd.Flags().Changed(isOfficialFlagName) {

		var isOfficialFlagName string
		if cmdPrefix == "" {
			isOfficialFlagName = "is_official"
		} else {
			isOfficialFlagName = fmt.Sprintf("%v.is_official", cmdPrefix)
		}

		isOfficialFlagValue, err := cmd.Flags().GetBool(isOfficialFlagName)
		if err != nil {
			return err, false
		}
		m.IsOfficial = isOfficialFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageSearchOKBodyItems0NameFlags(depth int, m *image.ImageSearchOKBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

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
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageSearchOKBodyItems0StarCountFlags(depth int, m *image.ImageSearchOKBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	starCountFlagName := fmt.Sprintf("%v.star_count", cmdPrefix)
	if cmd.Flags().Changed(starCountFlagName) {

		var starCountFlagName string
		if cmdPrefix == "" {
			starCountFlagName = "star_count"
		} else {
			starCountFlagName = fmt.Sprintf("%v.star_count", cmdPrefix)
		}

		starCountFlagValue, err := cmd.Flags().GetInt64(starCountFlagName)
		if err != nil {
			return err, false
		}
		m.StarCount = starCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/go-swagger/dockerctl/models"

	"github.com/spf13/cobra"
)

// Schema cli for Image

// register flags to command
func registerModelImageFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerImageArchitecture(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageAuthor(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageComment(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageContainer(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageContainerConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageCreated(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageDockerVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageGraphDriver(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageMetadata(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageOs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageOsVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageParent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageRepoDigests(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageRepoTags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageRootFS(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageVirtualSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageArchitecture(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ArchitectureDescription := `Required. `

	var ArchitectureFlagName string
	if cmdPrefix == "" {
		ArchitectureFlagName = "Architecture"
	} else {
		ArchitectureFlagName = fmt.Sprintf("%v.Architecture", cmdPrefix)
	}

	var ArchitectureFlagDefault string

	_ = cmd.PersistentFlags().String(ArchitectureFlagName, ArchitectureFlagDefault, ArchitectureDescription)

	return nil
}

func registerImageAuthor(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	AuthorDescription := `Required. `

	var AuthorFlagName string
	if cmdPrefix == "" {
		AuthorFlagName = "Author"
	} else {
		AuthorFlagName = fmt.Sprintf("%v.Author", cmdPrefix)
	}

	var AuthorFlagDefault string

	_ = cmd.PersistentFlags().String(AuthorFlagName, AuthorFlagDefault, AuthorDescription)

	return nil
}

func registerImageComment(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	CommentDescription := `Required. `

	var CommentFlagName string
	if cmdPrefix == "" {
		CommentFlagName = "Comment"
	} else {
		CommentFlagName = fmt.Sprintf("%v.Comment", cmdPrefix)
	}

	var CommentFlagDefault string

	_ = cmd.PersistentFlags().String(CommentFlagName, CommentFlagDefault, CommentDescription)

	return nil
}

func registerImageConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var ConfigFlagName string
	if cmdPrefix == "" {
		ConfigFlagName = "Config"
	} else {
		ConfigFlagName = fmt.Sprintf("%v.Config", cmdPrefix)
	}

	if err := registerModelContainerConfigFlags(depth+1, ConfigFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageContainer(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ContainerDescription := `Required. `

	var ContainerFlagName string
	if cmdPrefix == "" {
		ContainerFlagName = "Container"
	} else {
		ContainerFlagName = fmt.Sprintf("%v.Container", cmdPrefix)
	}

	var ContainerFlagDefault string

	_ = cmd.PersistentFlags().String(ContainerFlagName, ContainerFlagDefault, ContainerDescription)

	return nil
}

func registerImageContainerConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var ContainerConfigFlagName string
	if cmdPrefix == "" {
		ContainerConfigFlagName = "ContainerConfig"
	} else {
		ContainerConfigFlagName = fmt.Sprintf("%v.ContainerConfig", cmdPrefix)
	}

	if err := registerModelContainerConfigFlags(depth+1, ContainerConfigFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageCreated(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	CreatedDescription := `Required. `

	var CreatedFlagName string
	if cmdPrefix == "" {
		CreatedFlagName = "Created"
	} else {
		CreatedFlagName = fmt.Sprintf("%v.Created", cmdPrefix)
	}

	var CreatedFlagDefault string

	_ = cmd.PersistentFlags().String(CreatedFlagName, CreatedFlagDefault, CreatedDescription)

	return nil
}

func registerImageDockerVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	DockerVersionDescription := `Required. `

	var DockerVersionFlagName string
	if cmdPrefix == "" {
		DockerVersionFlagName = "DockerVersion"
	} else {
		DockerVersionFlagName = fmt.Sprintf("%v.DockerVersion", cmdPrefix)
	}

	var DockerVersionFlagDefault string

	_ = cmd.PersistentFlags().String(DockerVersionFlagName, DockerVersionFlagDefault, DockerVersionDescription)

	return nil
}

func registerImageGraphDriver(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var GraphDriverFlagName string
	if cmdPrefix == "" {
		GraphDriverFlagName = "GraphDriver"
	} else {
		GraphDriverFlagName = fmt.Sprintf("%v.GraphDriver", cmdPrefix)
	}

	if err := registerModelGraphDriverDataFlags(depth+1, GraphDriverFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	IDDescription := `Required. `

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

func registerImageMetadata(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var MetadataFlagName string
	if cmdPrefix == "" {
		MetadataFlagName = "Metadata"
	} else {
		MetadataFlagName = fmt.Sprintf("%v.Metadata", cmdPrefix)
	}

	if err := registerModelImageMetadataFlags(depth+1, MetadataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageOs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	OsDescription := `Required. `

	var OsFlagName string
	if cmdPrefix == "" {
		OsFlagName = "Os"
	} else {
		OsFlagName = fmt.Sprintf("%v.Os", cmdPrefix)
	}

	var OsFlagDefault string

	_ = cmd.PersistentFlags().String(OsFlagName, OsFlagDefault, OsDescription)

	return nil
}

func registerImageOsVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	OsVersionDescription := ``

	var OsVersionFlagName string
	if cmdPrefix == "" {
		OsVersionFlagName = "OsVersion"
	} else {
		OsVersionFlagName = fmt.Sprintf("%v.OsVersion", cmdPrefix)
	}

	var OsVersionFlagDefault string

	_ = cmd.PersistentFlags().String(OsVersionFlagName, OsVersionFlagDefault, OsVersionDescription)

	return nil
}

func registerImageParent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ParentDescription := `Required. `

	var ParentFlagName string
	if cmdPrefix == "" {
		ParentFlagName = "Parent"
	} else {
		ParentFlagName = fmt.Sprintf("%v.Parent", cmdPrefix)
	}

	var ParentFlagDefault string

	_ = cmd.PersistentFlags().String(ParentFlagName, ParentFlagDefault, ParentDescription)

	return nil
}

func registerImageRepoDigests(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: RepoDigests []string array type is not supported by go-swagger cli yet

	return nil
}

func registerImageRepoTags(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: RepoTags []string array type is not supported by go-swagger cli yet

	return nil
}

func registerImageRootFS(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var RootFSFlagName string
	if cmdPrefix == "" {
		RootFSFlagName = "RootFS"
	} else {
		RootFSFlagName = fmt.Sprintf("%v.RootFS", cmdPrefix)
	}

	if err := registerModelImageRootFSFlags(depth+1, RootFSFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	SizeDescription := `Required. `

	var SizeFlagName string
	if cmdPrefix == "" {
		SizeFlagName = "Size"
	} else {
		SizeFlagName = fmt.Sprintf("%v.Size", cmdPrefix)
	}

	var SizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(SizeFlagName, SizeFlagDefault, SizeDescription)

	return nil
}

func registerImageVirtualSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	VirtualSizeDescription := `Required. `

	var VirtualSizeFlagName string
	if cmdPrefix == "" {
		VirtualSizeFlagName = "VirtualSize"
	} else {
		VirtualSizeFlagName = fmt.Sprintf("%v.VirtualSize", cmdPrefix)
	}

	var VirtualSizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(VirtualSizeFlagName, VirtualSizeFlagDefault, VirtualSizeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelImageFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, ArchitectureAdded := retrieveImageArchitectureFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ArchitectureAdded

	err, AuthorAdded := retrieveImageAuthorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || AuthorAdded

	err, CommentAdded := retrieveImageCommentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || CommentAdded

	err, ConfigAdded := retrieveImageConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ConfigAdded

	err, ContainerAdded := retrieveImageContainerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ContainerAdded

	err, ContainerConfigAdded := retrieveImageContainerConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ContainerConfigAdded

	err, CreatedAdded := retrieveImageCreatedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || CreatedAdded

	err, DockerVersionAdded := retrieveImageDockerVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DockerVersionAdded

	err, GraphDriverAdded := retrieveImageGraphDriverFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || GraphDriverAdded

	err, IDAdded := retrieveImageIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || IDAdded

	err, MetadataAdded := retrieveImageMetadataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || MetadataAdded

	err, OsAdded := retrieveImageOsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || OsAdded

	err, OsVersionAdded := retrieveImageOsVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || OsVersionAdded

	err, ParentAdded := retrieveImageParentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ParentAdded

	err, RepoDigestsAdded := retrieveImageRepoDigestsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || RepoDigestsAdded

	err, RepoTagsAdded := retrieveImageRepoTagsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || RepoTagsAdded

	err, RootFSAdded := retrieveImageRootFSFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || RootFSAdded

	err, SizeAdded := retrieveImageSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SizeAdded

	err, VirtualSizeAdded := retrieveImageVirtualSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || VirtualSizeAdded

	return nil, retAdded
}

func retrieveImageArchitectureFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ArchitectureFlagName := fmt.Sprintf("%v.Architecture", cmdPrefix)
	if cmd.Flags().Changed(ArchitectureFlagName) {

		var ArchitectureFlagName string
		if cmdPrefix == "" {
			ArchitectureFlagName = "Architecture"
		} else {
			ArchitectureFlagName = fmt.Sprintf("%v.Architecture", cmdPrefix)
		}

		ArchitectureFlagValue, err := cmd.Flags().GetString(ArchitectureFlagName)
		if err != nil {
			return err, false
		}
		m.Architecture = ArchitectureFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageAuthorFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	AuthorFlagName := fmt.Sprintf("%v.Author", cmdPrefix)
	if cmd.Flags().Changed(AuthorFlagName) {

		var AuthorFlagName string
		if cmdPrefix == "" {
			AuthorFlagName = "Author"
		} else {
			AuthorFlagName = fmt.Sprintf("%v.Author", cmdPrefix)
		}

		AuthorFlagValue, err := cmd.Flags().GetString(AuthorFlagName)
		if err != nil {
			return err, false
		}
		m.Author = AuthorFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageCommentFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	CommentFlagName := fmt.Sprintf("%v.Comment", cmdPrefix)
	if cmd.Flags().Changed(CommentFlagName) {

		var CommentFlagName string
		if cmdPrefix == "" {
			CommentFlagName = "Comment"
		} else {
			CommentFlagName = fmt.Sprintf("%v.Comment", cmdPrefix)
		}

		CommentFlagValue, err := cmd.Flags().GetString(CommentFlagName)
		if err != nil {
			return err, false
		}
		m.Comment = CommentFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageConfigFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ConfigFlagName := fmt.Sprintf("%v.Config", cmdPrefix)
	if cmd.Flags().Changed(ConfigFlagName) {
		// info: complex object Config ContainerConfig is retrieved outside this Changed() block
	}
	ConfigFlagValue := m.Config
	if swag.IsZero(ConfigFlagValue) {
		ConfigFlagValue = &models.ContainerConfig{}
	}

	err, ConfigAdded := retrieveModelContainerConfigFlags(depth+1, ConfigFlagValue, ConfigFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ConfigAdded
	if ConfigAdded {
		m.Config = ConfigFlagValue
	}

	return nil, retAdded
}

func retrieveImageContainerFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ContainerFlagName := fmt.Sprintf("%v.Container", cmdPrefix)
	if cmd.Flags().Changed(ContainerFlagName) {

		var ContainerFlagName string
		if cmdPrefix == "" {
			ContainerFlagName = "Container"
		} else {
			ContainerFlagName = fmt.Sprintf("%v.Container", cmdPrefix)
		}

		ContainerFlagValue, err := cmd.Flags().GetString(ContainerFlagName)
		if err != nil {
			return err, false
		}
		m.Container = ContainerFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageContainerConfigFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ContainerConfigFlagName := fmt.Sprintf("%v.ContainerConfig", cmdPrefix)
	if cmd.Flags().Changed(ContainerConfigFlagName) {
		// info: complex object ContainerConfig ContainerConfig is retrieved outside this Changed() block
	}
	ContainerConfigFlagValue := m.ContainerConfig
	if swag.IsZero(ContainerConfigFlagValue) {
		ContainerConfigFlagValue = &models.ContainerConfig{}
	}

	err, ContainerConfigAdded := retrieveModelContainerConfigFlags(depth+1, ContainerConfigFlagValue, ContainerConfigFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ContainerConfigAdded
	if ContainerConfigAdded {
		m.ContainerConfig = ContainerConfigFlagValue
	}

	return nil, retAdded
}

func retrieveImageCreatedFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	CreatedFlagName := fmt.Sprintf("%v.Created", cmdPrefix)
	if cmd.Flags().Changed(CreatedFlagName) {

		var CreatedFlagName string
		if cmdPrefix == "" {
			CreatedFlagName = "Created"
		} else {
			CreatedFlagName = fmt.Sprintf("%v.Created", cmdPrefix)
		}

		CreatedFlagValue, err := cmd.Flags().GetString(CreatedFlagName)
		if err != nil {
			return err, false
		}
		m.Created = CreatedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageDockerVersionFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	DockerVersionFlagName := fmt.Sprintf("%v.DockerVersion", cmdPrefix)
	if cmd.Flags().Changed(DockerVersionFlagName) {

		var DockerVersionFlagName string
		if cmdPrefix == "" {
			DockerVersionFlagName = "DockerVersion"
		} else {
			DockerVersionFlagName = fmt.Sprintf("%v.DockerVersion", cmdPrefix)
		}

		DockerVersionFlagValue, err := cmd.Flags().GetString(DockerVersionFlagName)
		if err != nil {
			return err, false
		}
		m.DockerVersion = DockerVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageGraphDriverFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	GraphDriverFlagName := fmt.Sprintf("%v.GraphDriver", cmdPrefix)
	if cmd.Flags().Changed(GraphDriverFlagName) {
		// info: complex object GraphDriver GraphDriverData is retrieved outside this Changed() block
	}
	GraphDriverFlagValue := m.GraphDriver
	if swag.IsZero(GraphDriverFlagValue) {
		GraphDriverFlagValue = &models.GraphDriverData{}
	}

	err, GraphDriverAdded := retrieveModelGraphDriverDataFlags(depth+1, GraphDriverFlagValue, GraphDriverFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || GraphDriverAdded
	if GraphDriverAdded {
		m.GraphDriver = GraphDriverFlagValue
	}

	return nil, retAdded
}

func retrieveImageIDFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveImageMetadataFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	MetadataFlagName := fmt.Sprintf("%v.Metadata", cmdPrefix)
	if cmd.Flags().Changed(MetadataFlagName) {
		// info: complex object Metadata ImageMetadata is retrieved outside this Changed() block
	}
	MetadataFlagValue := m.Metadata
	if swag.IsZero(MetadataFlagValue) {
		MetadataFlagValue = &models.ImageMetadata{}
	}

	err, MetadataAdded := retrieveModelImageMetadataFlags(depth+1, MetadataFlagValue, MetadataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || MetadataAdded
	if MetadataAdded {
		m.Metadata = MetadataFlagValue
	}

	return nil, retAdded
}

func retrieveImageOsFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	OsFlagName := fmt.Sprintf("%v.Os", cmdPrefix)
	if cmd.Flags().Changed(OsFlagName) {

		var OsFlagName string
		if cmdPrefix == "" {
			OsFlagName = "Os"
		} else {
			OsFlagName = fmt.Sprintf("%v.Os", cmdPrefix)
		}

		OsFlagValue, err := cmd.Flags().GetString(OsFlagName)
		if err != nil {
			return err, false
		}
		m.Os = OsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageOsVersionFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	OsVersionFlagName := fmt.Sprintf("%v.OsVersion", cmdPrefix)
	if cmd.Flags().Changed(OsVersionFlagName) {

		var OsVersionFlagName string
		if cmdPrefix == "" {
			OsVersionFlagName = "OsVersion"
		} else {
			OsVersionFlagName = fmt.Sprintf("%v.OsVersion", cmdPrefix)
		}

		OsVersionFlagValue, err := cmd.Flags().GetString(OsVersionFlagName)
		if err != nil {
			return err, false
		}
		m.OsVersion = OsVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageParentFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ParentFlagName := fmt.Sprintf("%v.Parent", cmdPrefix)
	if cmd.Flags().Changed(ParentFlagName) {

		var ParentFlagName string
		if cmdPrefix == "" {
			ParentFlagName = "Parent"
		} else {
			ParentFlagName = fmt.Sprintf("%v.Parent", cmdPrefix)
		}

		ParentFlagValue, err := cmd.Flags().GetString(ParentFlagName)
		if err != nil {
			return err, false
		}
		m.Parent = ParentFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageRepoDigestsFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	RepoDigestsFlagName := fmt.Sprintf("%v.RepoDigests", cmdPrefix)
	if cmd.Flags().Changed(RepoDigestsFlagName) {
		// warning: RepoDigests array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveImageRepoTagsFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	RepoTagsFlagName := fmt.Sprintf("%v.RepoTags", cmdPrefix)
	if cmd.Flags().Changed(RepoTagsFlagName) {
		// warning: RepoTags array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveImageRootFSFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	RootFSFlagName := fmt.Sprintf("%v.RootFS", cmdPrefix)
	if cmd.Flags().Changed(RootFSFlagName) {
		// info: complex object RootFS ImageRootFS is retrieved outside this Changed() block
	}
	RootFSFlagValue := m.RootFS
	if swag.IsZero(RootFSFlagValue) {
		RootFSFlagValue = &models.ImageRootFS{}
	}

	err, RootFSAdded := retrieveModelImageRootFSFlags(depth+1, RootFSFlagValue, RootFSFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || RootFSAdded
	if RootFSAdded {
		m.RootFS = RootFSFlagValue
	}

	return nil, retAdded
}

func retrieveImageSizeFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	SizeFlagName := fmt.Sprintf("%v.Size", cmdPrefix)
	if cmd.Flags().Changed(SizeFlagName) {

		var SizeFlagName string
		if cmdPrefix == "" {
			SizeFlagName = "Size"
		} else {
			SizeFlagName = fmt.Sprintf("%v.Size", cmdPrefix)
		}

		SizeFlagValue, err := cmd.Flags().GetInt64(SizeFlagName)
		if err != nil {
			return err, false
		}
		m.Size = SizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageVirtualSizeFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	VirtualSizeFlagName := fmt.Sprintf("%v.VirtualSize", cmdPrefix)
	if cmd.Flags().Changed(VirtualSizeFlagName) {

		var VirtualSizeFlagName string
		if cmdPrefix == "" {
			VirtualSizeFlagName = "VirtualSize"
		} else {
			VirtualSizeFlagName = fmt.Sprintf("%v.VirtualSize", cmdPrefix)
		}

		VirtualSizeFlagValue, err := cmd.Flags().GetInt64(VirtualSizeFlagName)
		if err != nil {
			return err, false
		}
		m.VirtualSize = VirtualSizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for ImageMetadata

// register flags to command
func registerModelImageMetadataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerImageMetadataLastTagTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageMetadataLastTagTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	LastTagTimeDescription := ``

	var LastTagTimeFlagName string
	if cmdPrefix == "" {
		LastTagTimeFlagName = "LastTagTime"
	} else {
		LastTagTimeFlagName = fmt.Sprintf("%v.LastTagTime", cmdPrefix)
	}

	var LastTagTimeFlagDefault string

	_ = cmd.PersistentFlags().String(LastTagTimeFlagName, LastTagTimeFlagDefault, LastTagTimeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelImageMetadataFlags(depth int, m *models.ImageMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, LastTagTimeAdded := retrieveImageMetadataLastTagTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || LastTagTimeAdded

	return nil, retAdded
}

func retrieveImageMetadataLastTagTimeFlags(depth int, m *models.ImageMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	LastTagTimeFlagName := fmt.Sprintf("%v.LastTagTime", cmdPrefix)
	if cmd.Flags().Changed(LastTagTimeFlagName) {

		var LastTagTimeFlagName string
		if cmdPrefix == "" {
			LastTagTimeFlagName = "LastTagTime"
		} else {
			LastTagTimeFlagName = fmt.Sprintf("%v.LastTagTime", cmdPrefix)
		}

		LastTagTimeFlagValue, err := cmd.Flags().GetString(LastTagTimeFlagName)
		if err != nil {
			return err, false
		}
		m.LastTagTime = LastTagTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for ImageRootFS

// register flags to command
func registerModelImageRootFSFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerImageRootFSBaseLayer(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageRootFSLayers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageRootFSType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageRootFSBaseLayer(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	BaseLayerDescription := ``

	var BaseLayerFlagName string
	if cmdPrefix == "" {
		BaseLayerFlagName = "BaseLayer"
	} else {
		BaseLayerFlagName = fmt.Sprintf("%v.BaseLayer", cmdPrefix)
	}

	var BaseLayerFlagDefault string

	_ = cmd.PersistentFlags().String(BaseLayerFlagName, BaseLayerFlagDefault, BaseLayerDescription)

	return nil
}

func registerImageRootFSLayers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Layers []string array type is not supported by go-swagger cli yet

	return nil
}

func registerImageRootFSType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	TypeDescription := `Required. `

	var TypeFlagName string
	if cmdPrefix == "" {
		TypeFlagName = "Type"
	} else {
		TypeFlagName = fmt.Sprintf("%v.Type", cmdPrefix)
	}

	var TypeFlagDefault string

	_ = cmd.PersistentFlags().String(TypeFlagName, TypeFlagDefault, TypeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelImageRootFSFlags(depth int, m *models.ImageRootFS, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, BaseLayerAdded := retrieveImageRootFSBaseLayerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || BaseLayerAdded

	err, LayersAdded := retrieveImageRootFSLayersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || LayersAdded

	err, TypeAdded := retrieveImageRootFSTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || TypeAdded

	return nil, retAdded
}

func retrieveImageRootFSBaseLayerFlags(depth int, m *models.ImageRootFS, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	BaseLayerFlagName := fmt.Sprintf("%v.BaseLayer", cmdPrefix)
	if cmd.Flags().Changed(BaseLayerFlagName) {

		var BaseLayerFlagName string
		if cmdPrefix == "" {
			BaseLayerFlagName = "BaseLayer"
		} else {
			BaseLayerFlagName = fmt.Sprintf("%v.BaseLayer", cmdPrefix)
		}

		BaseLayerFlagValue, err := cmd.Flags().GetString(BaseLayerFlagName)
		if err != nil {
			return err, false
		}
		m.BaseLayer = BaseLayerFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageRootFSLayersFlags(depth int, m *models.ImageRootFS, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	LayersFlagName := fmt.Sprintf("%v.Layers", cmdPrefix)
	if cmd.Flags().Changed(LayersFlagName) {
		// warning: Layers array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveImageRootFSTypeFlags(depth int, m *models.ImageRootFS, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	TypeFlagName := fmt.Sprintf("%v.Type", cmdPrefix)
	if cmd.Flags().Changed(TypeFlagName) {

		var TypeFlagName string
		if cmdPrefix == "" {
			TypeFlagName = "Type"
		} else {
			TypeFlagName = fmt.Sprintf("%v.Type", cmdPrefix)
		}

		TypeFlagValue, err := cmd.Flags().GetString(TypeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = TypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

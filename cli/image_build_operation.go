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

// makeOperationImageImageBuildCmd returns a cmd to handle operation imageBuild
func makeOperationImageImageBuildCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "ImageBuild",
		Short: `Build an image from a tar archive with a ` + "`" + `Dockerfile` + "`" + ` in it.

The ` + "`" + `Dockerfile` + "`" + ` specifies how the image is built from the tar archive. It is typically in the archive's root, but can be at a different path or have a different name by specifying the ` + "`" + `dockerfile` + "`" + ` parameter. [See the ` + "`" + `Dockerfile` + "`" + ` reference for more information](https://docs.docker.com/engine/reference/builder/).

The Docker daemon performs a preliminary validation of the ` + "`" + `Dockerfile` + "`" + ` before starting the build, and returns an error if the syntax is incorrect. After that, each instruction is run one-by-one until the ID of the new image is output.

The build is canceled if the client drops the connection by quitting or being killed.
`,
		RunE: runOperationImageImageBuild,
	}

	if err := registerOperationImageImageBuildParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationImageImageBuild uses cmd flags to call endpoint api
func runOperationImageImageBuild(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := image.NewImageBuildParams()
	if err, _ := retrieveOperationImageImageBuildContentTypeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildXRegistryConfigFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildBuildargsFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildCachefromFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildCpuperiodFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildCpuquotaFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildCpusetcpusFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildCpusharesFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildDockerfileFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildExtrahostsFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildForcermFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildInputStreamFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildLabelsFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildMemoryFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildMemswapFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildNetworkmodeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildNocacheFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildOutputsFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildPlatformFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildPullFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildQFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildRemoteFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildRmFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildShmsizeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildSquashFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildTFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageBuildTargetFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationImageImageBuildResult(appCli.Image.ImageBuild(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationImageImageBuildParamFlags registers all flags needed to fill params
func registerOperationImageImageBuildParamFlags(cmd *cobra.Command) error {
	if err := registerOperationImageImageBuildContentTypeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildXRegistryConfigParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildBuildargsParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildCachefromParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildCpuperiodParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildCpuquotaParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildCpusetcpusParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildCpusharesParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildDockerfileParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildExtrahostsParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildForcermParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildInputStreamParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildLabelsParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildMemoryParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildMemswapParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildNetworkmodeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildNocacheParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildOutputsParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildPlatformParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildPullParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildQParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildRemoteParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildRmParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildShmsizeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildSquashParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildTParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageBuildTargetParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationImageImageBuildContentTypeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	ContentTypeDescription := `Enum: ["application/x-tar"]. `

	var ContentTypeFlagName string
	if cmdPrefix == "" {
		ContentTypeFlagName = "Content-type"
	} else {
		ContentTypeFlagName = fmt.Sprintf("%v.Content-type", cmdPrefix)
	}

	var ContentTypeFlagDefault string = "application/x-tar"

	_ = cmd.PersistentFlags().String(ContentTypeFlagName, ContentTypeFlagDefault, ContentTypeDescription)

	if err := cmd.RegisterFlagCompletionFunc(ContentTypeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["application/x-tar"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}
func registerOperationImageImageBuildXRegistryConfigParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	XRegistryConfigDescription := `This is a base64-encoded JSON object with auth configurations for multiple registries that a build may refer to.

The key is a registry URL, and the value is an auth configuration object, [as described in the authentication section](#section/Authentication). For example:

` + "`" + `` + "`" + `` + "`" + `
{
  "docker.example.com": {
    "username": "janedoe",
    "password": "hunter2"
  },
  "https://index.docker.io/v1/": {
    "username": "mobydock",
    "password": "conta1n3rize14"
  }
}
` + "`" + `` + "`" + `` + "`" + `

Only the registry domain name (and port if not the default 443) are required. However, for legacy reasons, the Docker Hub registry must be specified with both a ` + "`" + `https://` + "`" + ` prefix and a ` + "`" + `/v1/` + "`" + ` suffix even though Docker will prefer to use the v2 registry API.
`

	var XRegistryConfigFlagName string
	if cmdPrefix == "" {
		XRegistryConfigFlagName = "X-Registry-Config"
	} else {
		XRegistryConfigFlagName = fmt.Sprintf("%v.X-Registry-Config", cmdPrefix)
	}

	var XRegistryConfigFlagDefault string

	_ = cmd.PersistentFlags().String(XRegistryConfigFlagName, XRegistryConfigFlagDefault, XRegistryConfigDescription)

	return nil
}
func registerOperationImageImageBuildBuildargsParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	BuildargsDescription := `JSON map of string pairs for build-time variables. Users pass these values at build-time. Docker uses the buildargs as the environment context for commands run via the ` + "`" + `Dockerfile` + "`" + ` RUN instruction, or for variable expansion in other ` + "`" + `Dockerfile` + "`" + ` instructions. This is not meant for passing secret values.

For example, the build arg ` + "`" + `FOO=bar` + "`" + ` would become ` + "`" + `{"FOO":"bar"}` + "`" + ` in JSON. This would result in the the query parameter ` + "`" + `buildargs={"FOO":"bar"}` + "`" + `. Note that ` + "`" + `{"FOO":"bar"}` + "`" + ` should be URI component encoded.

[Read more about the buildargs instruction.](https://docs.docker.com/engine/reference/builder/#arg)
`

	var BuildargsFlagName string
	if cmdPrefix == "" {
		BuildargsFlagName = "buildargs"
	} else {
		BuildargsFlagName = fmt.Sprintf("%v.buildargs", cmdPrefix)
	}

	var BuildargsFlagDefault string

	_ = cmd.PersistentFlags().String(BuildargsFlagName, BuildargsFlagDefault, BuildargsDescription)

	return nil
}
func registerOperationImageImageBuildCachefromParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	CachefromDescription := `JSON array of images used for build cache resolution.`

	var CachefromFlagName string
	if cmdPrefix == "" {
		CachefromFlagName = "cachefrom"
	} else {
		CachefromFlagName = fmt.Sprintf("%v.cachefrom", cmdPrefix)
	}

	var CachefromFlagDefault string

	_ = cmd.PersistentFlags().String(CachefromFlagName, CachefromFlagDefault, CachefromDescription)

	return nil
}
func registerOperationImageImageBuildCpuperiodParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	CpuperiodDescription := `The length of a CPU period in microseconds.`

	var CpuperiodFlagName string
	if cmdPrefix == "" {
		CpuperiodFlagName = "cpuperiod"
	} else {
		CpuperiodFlagName = fmt.Sprintf("%v.cpuperiod", cmdPrefix)
	}

	var CpuperiodFlagDefault int64

	_ = cmd.PersistentFlags().Int64(CpuperiodFlagName, CpuperiodFlagDefault, CpuperiodDescription)

	return nil
}
func registerOperationImageImageBuildCpuquotaParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	CpuquotaDescription := `Microseconds of CPU time that the container can get in a CPU period.`

	var CpuquotaFlagName string
	if cmdPrefix == "" {
		CpuquotaFlagName = "cpuquota"
	} else {
		CpuquotaFlagName = fmt.Sprintf("%v.cpuquota", cmdPrefix)
	}

	var CpuquotaFlagDefault int64

	_ = cmd.PersistentFlags().Int64(CpuquotaFlagName, CpuquotaFlagDefault, CpuquotaDescription)

	return nil
}
func registerOperationImageImageBuildCpusetcpusParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	CpusetcpusDescription := `CPUs in which to allow execution (e.g., ` + "`" + `0-3` + "`" + `, ` + "`" + `0,1` + "`" + `).`

	var CpusetcpusFlagName string
	if cmdPrefix == "" {
		CpusetcpusFlagName = "cpusetcpus"
	} else {
		CpusetcpusFlagName = fmt.Sprintf("%v.cpusetcpus", cmdPrefix)
	}

	var CpusetcpusFlagDefault string

	_ = cmd.PersistentFlags().String(CpusetcpusFlagName, CpusetcpusFlagDefault, CpusetcpusDescription)

	return nil
}
func registerOperationImageImageBuildCpusharesParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	CpusharesDescription := `CPU shares (relative weight).`

	var CpusharesFlagName string
	if cmdPrefix == "" {
		CpusharesFlagName = "cpushares"
	} else {
		CpusharesFlagName = fmt.Sprintf("%v.cpushares", cmdPrefix)
	}

	var CpusharesFlagDefault int64

	_ = cmd.PersistentFlags().Int64(CpusharesFlagName, CpusharesFlagDefault, CpusharesDescription)

	return nil
}
func registerOperationImageImageBuildDockerfileParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	DockerfileDescription := `Path within the build context to the ` + "`" + `Dockerfile` + "`" + `. This is ignored if ` + "`" + `remote` + "`" + ` is specified and points to an external ` + "`" + `Dockerfile` + "`" + `.`

	var DockerfileFlagName string
	if cmdPrefix == "" {
		DockerfileFlagName = "dockerfile"
	} else {
		DockerfileFlagName = fmt.Sprintf("%v.dockerfile", cmdPrefix)
	}

	var DockerfileFlagDefault string = "Dockerfile"

	_ = cmd.PersistentFlags().String(DockerfileFlagName, DockerfileFlagDefault, DockerfileDescription)

	return nil
}
func registerOperationImageImageBuildExtrahostsParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	ExtrahostsDescription := `Extra hosts to add to /etc/hosts`

	var ExtrahostsFlagName string
	if cmdPrefix == "" {
		ExtrahostsFlagName = "extrahosts"
	} else {
		ExtrahostsFlagName = fmt.Sprintf("%v.extrahosts", cmdPrefix)
	}

	var ExtrahostsFlagDefault string

	_ = cmd.PersistentFlags().String(ExtrahostsFlagName, ExtrahostsFlagDefault, ExtrahostsDescription)

	return nil
}
func registerOperationImageImageBuildForcermParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	ForcermDescription := `Always remove intermediate containers, even upon failure.`

	var ForcermFlagName string
	if cmdPrefix == "" {
		ForcermFlagName = "forcerm"
	} else {
		ForcermFlagName = fmt.Sprintf("%v.forcerm", cmdPrefix)
	}

	var ForcermFlagDefault bool

	_ = cmd.PersistentFlags().Bool(ForcermFlagName, ForcermFlagDefault, ForcermDescription)

	return nil
}
func registerOperationImageImageBuildInputStreamParamFlags(cmdPrefix string, cmd *cobra.Command) error {
	// warning: go type io.ReadCloser is not supported by go-swagger cli yet.
	return nil
}
func registerOperationImageImageBuildLabelsParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	LabelsDescription := `Arbitrary key/value labels to set on the image, as a JSON map of string pairs.`

	var LabelsFlagName string
	if cmdPrefix == "" {
		LabelsFlagName = "labels"
	} else {
		LabelsFlagName = fmt.Sprintf("%v.labels", cmdPrefix)
	}

	var LabelsFlagDefault string

	_ = cmd.PersistentFlags().String(LabelsFlagName, LabelsFlagDefault, LabelsDescription)

	return nil
}
func registerOperationImageImageBuildMemoryParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	MemoryDescription := `Set memory limit for build.`

	var MemoryFlagName string
	if cmdPrefix == "" {
		MemoryFlagName = "memory"
	} else {
		MemoryFlagName = fmt.Sprintf("%v.memory", cmdPrefix)
	}

	var MemoryFlagDefault int64

	_ = cmd.PersistentFlags().Int64(MemoryFlagName, MemoryFlagDefault, MemoryDescription)

	return nil
}
func registerOperationImageImageBuildMemswapParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	MemswapDescription := `Total memory (memory + swap). Set as ` + "`" + `-1` + "`" + ` to disable swap.`

	var MemswapFlagName string
	if cmdPrefix == "" {
		MemswapFlagName = "memswap"
	} else {
		MemswapFlagName = fmt.Sprintf("%v.memswap", cmdPrefix)
	}

	var MemswapFlagDefault int64

	_ = cmd.PersistentFlags().Int64(MemswapFlagName, MemswapFlagDefault, MemswapDescription)

	return nil
}
func registerOperationImageImageBuildNetworkmodeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	NetworkmodeDescription := `Sets the networking mode for the run commands during build. Supported
standard values are: ` + "`" + `bridge` + "`" + `, ` + "`" + `host` + "`" + `, ` + "`" + `none` + "`" + `, and ` + "`" + `container:<name|id>` + "`" + `.
Any other value is taken as a custom network's name or ID to which this
container should connect to.
`

	var NetworkmodeFlagName string
	if cmdPrefix == "" {
		NetworkmodeFlagName = "networkmode"
	} else {
		NetworkmodeFlagName = fmt.Sprintf("%v.networkmode", cmdPrefix)
	}

	var NetworkmodeFlagDefault string

	_ = cmd.PersistentFlags().String(NetworkmodeFlagName, NetworkmodeFlagDefault, NetworkmodeDescription)

	return nil
}
func registerOperationImageImageBuildNocacheParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	NocacheDescription := `Do not use the cache when building the image.`

	var NocacheFlagName string
	if cmdPrefix == "" {
		NocacheFlagName = "nocache"
	} else {
		NocacheFlagName = fmt.Sprintf("%v.nocache", cmdPrefix)
	}

	var NocacheFlagDefault bool

	_ = cmd.PersistentFlags().Bool(NocacheFlagName, NocacheFlagDefault, NocacheDescription)

	return nil
}
func registerOperationImageImageBuildOutputsParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	OutputsDescription := `BuildKit output configuration`

	var OutputsFlagName string
	if cmdPrefix == "" {
		OutputsFlagName = "outputs"
	} else {
		OutputsFlagName = fmt.Sprintf("%v.outputs", cmdPrefix)
	}

	var OutputsFlagDefault string

	_ = cmd.PersistentFlags().String(OutputsFlagName, OutputsFlagDefault, OutputsDescription)

	return nil
}
func registerOperationImageImageBuildPlatformParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	PlatformDescription := `Platform in the format os[/arch[/variant]]`

	var PlatformFlagName string
	if cmdPrefix == "" {
		PlatformFlagName = "platform"
	} else {
		PlatformFlagName = fmt.Sprintf("%v.platform", cmdPrefix)
	}

	var PlatformFlagDefault string

	_ = cmd.PersistentFlags().String(PlatformFlagName, PlatformFlagDefault, PlatformDescription)

	return nil
}
func registerOperationImageImageBuildPullParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	PullDescription := `Attempt to pull the image even if an older image exists locally.`

	var PullFlagName string
	if cmdPrefix == "" {
		PullFlagName = "pull"
	} else {
		PullFlagName = fmt.Sprintf("%v.pull", cmdPrefix)
	}

	var PullFlagDefault string

	_ = cmd.PersistentFlags().String(PullFlagName, PullFlagDefault, PullDescription)

	return nil
}
func registerOperationImageImageBuildQParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	QDescription := `Suppress verbose build output.`

	var QFlagName string
	if cmdPrefix == "" {
		QFlagName = "q"
	} else {
		QFlagName = fmt.Sprintf("%v.q", cmdPrefix)
	}

	var QFlagDefault bool

	_ = cmd.PersistentFlags().Bool(QFlagName, QFlagDefault, QDescription)

	return nil
}
func registerOperationImageImageBuildRemoteParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	RemoteDescription := `A Git repository URI or HTTP/HTTPS context URI. If the URI points to a single text file, the file’s contents are placed into a file called ` + "`" + `Dockerfile` + "`" + ` and the image is built from that file. If the URI points to a tarball, the file is downloaded by the daemon and the contents therein used as the context for the build. If the URI points to a tarball and the ` + "`" + `dockerfile` + "`" + ` parameter is also specified, there must be a file with the corresponding path inside the tarball.`

	var RemoteFlagName string
	if cmdPrefix == "" {
		RemoteFlagName = "remote"
	} else {
		RemoteFlagName = fmt.Sprintf("%v.remote", cmdPrefix)
	}

	var RemoteFlagDefault string

	_ = cmd.PersistentFlags().String(RemoteFlagName, RemoteFlagDefault, RemoteDescription)

	return nil
}
func registerOperationImageImageBuildRmParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	RmDescription := `Remove intermediate containers after a successful build.`

	var RmFlagName string
	if cmdPrefix == "" {
		RmFlagName = "rm"
	} else {
		RmFlagName = fmt.Sprintf("%v.rm", cmdPrefix)
	}

	var RmFlagDefault bool = true

	_ = cmd.PersistentFlags().Bool(RmFlagName, RmFlagDefault, RmDescription)

	return nil
}
func registerOperationImageImageBuildShmsizeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	ShmsizeDescription := `Size of ` + "`" + `/dev/shm` + "`" + ` in bytes. The size must be greater than 0. If omitted the system uses 64MB.`

	var ShmsizeFlagName string
	if cmdPrefix == "" {
		ShmsizeFlagName = "shmsize"
	} else {
		ShmsizeFlagName = fmt.Sprintf("%v.shmsize", cmdPrefix)
	}

	var ShmsizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(ShmsizeFlagName, ShmsizeFlagDefault, ShmsizeDescription)

	return nil
}
func registerOperationImageImageBuildSquashParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	SquashDescription := `Squash the resulting images layers into a single layer. *(Experimental release only.)*`

	var SquashFlagName string
	if cmdPrefix == "" {
		SquashFlagName = "squash"
	} else {
		SquashFlagName = fmt.Sprintf("%v.squash", cmdPrefix)
	}

	var SquashFlagDefault bool

	_ = cmd.PersistentFlags().Bool(SquashFlagName, SquashFlagDefault, SquashDescription)

	return nil
}
func registerOperationImageImageBuildTParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	TDescription := `A name and optional tag to apply to the image in the ` + "`" + `name:tag` + "`" + ` format. If you omit the tag the default ` + "`" + `latest` + "`" + ` value is assumed. You can provide several ` + "`" + `t` + "`" + ` parameters.`

	var TFlagName string
	if cmdPrefix == "" {
		TFlagName = "t"
	} else {
		TFlagName = fmt.Sprintf("%v.t", cmdPrefix)
	}

	var TFlagDefault string

	_ = cmd.PersistentFlags().String(TFlagName, TFlagDefault, TDescription)

	return nil
}
func registerOperationImageImageBuildTargetParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	TargetDescription := `Target build stage`

	var TargetFlagName string
	if cmdPrefix == "" {
		TargetFlagName = "target"
	} else {
		TargetFlagName = fmt.Sprintf("%v.target", cmdPrefix)
	}

	var TargetFlagDefault string

	_ = cmd.PersistentFlags().String(TargetFlagName, TargetFlagDefault, TargetDescription)

	return nil
}

func retrieveOperationImageImageBuildContentTypeFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("Content-type") {

		var ContentTypeFlagName string
		if cmdPrefix == "" {
			ContentTypeFlagName = "Content-type"
		} else {
			ContentTypeFlagName = fmt.Sprintf("%v.Content-type", cmdPrefix)
		}

		ContentTypeFlagValue, err := cmd.Flags().GetString(ContentTypeFlagName)
		if err != nil {
			return err, false
		}
		m.ContentType = &ContentTypeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildXRegistryConfigFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("X-Registry-Config") {

		var XRegistryConfigFlagName string
		if cmdPrefix == "" {
			XRegistryConfigFlagName = "X-Registry-Config"
		} else {
			XRegistryConfigFlagName = fmt.Sprintf("%v.X-Registry-Config", cmdPrefix)
		}

		XRegistryConfigFlagValue, err := cmd.Flags().GetString(XRegistryConfigFlagName)
		if err != nil {
			return err, false
		}
		m.XRegistryConfig = &XRegistryConfigFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildBuildargsFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("buildargs") {

		var BuildargsFlagName string
		if cmdPrefix == "" {
			BuildargsFlagName = "buildargs"
		} else {
			BuildargsFlagName = fmt.Sprintf("%v.buildargs", cmdPrefix)
		}

		BuildargsFlagValue, err := cmd.Flags().GetString(BuildargsFlagName)
		if err != nil {
			return err, false
		}
		m.Buildargs = &BuildargsFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildCachefromFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("cachefrom") {

		var CachefromFlagName string
		if cmdPrefix == "" {
			CachefromFlagName = "cachefrom"
		} else {
			CachefromFlagName = fmt.Sprintf("%v.cachefrom", cmdPrefix)
		}

		CachefromFlagValue, err := cmd.Flags().GetString(CachefromFlagName)
		if err != nil {
			return err, false
		}
		m.Cachefrom = &CachefromFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildCpuperiodFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("cpuperiod") {

		var CpuperiodFlagName string
		if cmdPrefix == "" {
			CpuperiodFlagName = "cpuperiod"
		} else {
			CpuperiodFlagName = fmt.Sprintf("%v.cpuperiod", cmdPrefix)
		}

		CpuperiodFlagValue, err := cmd.Flags().GetInt64(CpuperiodFlagName)
		if err != nil {
			return err, false
		}
		m.Cpuperiod = &CpuperiodFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildCpuquotaFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("cpuquota") {

		var CpuquotaFlagName string
		if cmdPrefix == "" {
			CpuquotaFlagName = "cpuquota"
		} else {
			CpuquotaFlagName = fmt.Sprintf("%v.cpuquota", cmdPrefix)
		}

		CpuquotaFlagValue, err := cmd.Flags().GetInt64(CpuquotaFlagName)
		if err != nil {
			return err, false
		}
		m.Cpuquota = &CpuquotaFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildCpusetcpusFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("cpusetcpus") {

		var CpusetcpusFlagName string
		if cmdPrefix == "" {
			CpusetcpusFlagName = "cpusetcpus"
		} else {
			CpusetcpusFlagName = fmt.Sprintf("%v.cpusetcpus", cmdPrefix)
		}

		CpusetcpusFlagValue, err := cmd.Flags().GetString(CpusetcpusFlagName)
		if err != nil {
			return err, false
		}
		m.Cpusetcpus = &CpusetcpusFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildCpusharesFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("cpushares") {

		var CpusharesFlagName string
		if cmdPrefix == "" {
			CpusharesFlagName = "cpushares"
		} else {
			CpusharesFlagName = fmt.Sprintf("%v.cpushares", cmdPrefix)
		}

		CpusharesFlagValue, err := cmd.Flags().GetInt64(CpusharesFlagName)
		if err != nil {
			return err, false
		}
		m.Cpushares = &CpusharesFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildDockerfileFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("dockerfile") {

		var DockerfileFlagName string
		if cmdPrefix == "" {
			DockerfileFlagName = "dockerfile"
		} else {
			DockerfileFlagName = fmt.Sprintf("%v.dockerfile", cmdPrefix)
		}

		DockerfileFlagValue, err := cmd.Flags().GetString(DockerfileFlagName)
		if err != nil {
			return err, false
		}
		m.Dockerfile = &DockerfileFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildExtrahostsFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("extrahosts") {

		var ExtrahostsFlagName string
		if cmdPrefix == "" {
			ExtrahostsFlagName = "extrahosts"
		} else {
			ExtrahostsFlagName = fmt.Sprintf("%v.extrahosts", cmdPrefix)
		}

		ExtrahostsFlagValue, err := cmd.Flags().GetString(ExtrahostsFlagName)
		if err != nil {
			return err, false
		}
		m.Extrahosts = &ExtrahostsFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildForcermFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("forcerm") {

		var ForcermFlagName string
		if cmdPrefix == "" {
			ForcermFlagName = "forcerm"
		} else {
			ForcermFlagName = fmt.Sprintf("%v.forcerm", cmdPrefix)
		}

		ForcermFlagValue, err := cmd.Flags().GetBool(ForcermFlagName)
		if err != nil {
			return err, false
		}
		m.Forcerm = &ForcermFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildInputStreamFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("inputStream") {
		// warning: io.ReadCloser is not supported by go-swagger cli yet
	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildLabelsFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("labels") {

		var LabelsFlagName string
		if cmdPrefix == "" {
			LabelsFlagName = "labels"
		} else {
			LabelsFlagName = fmt.Sprintf("%v.labels", cmdPrefix)
		}

		LabelsFlagValue, err := cmd.Flags().GetString(LabelsFlagName)
		if err != nil {
			return err, false
		}
		m.Labels = &LabelsFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildMemoryFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("memory") {

		var MemoryFlagName string
		if cmdPrefix == "" {
			MemoryFlagName = "memory"
		} else {
			MemoryFlagName = fmt.Sprintf("%v.memory", cmdPrefix)
		}

		MemoryFlagValue, err := cmd.Flags().GetInt64(MemoryFlagName)
		if err != nil {
			return err, false
		}
		m.Memory = &MemoryFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildMemswapFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("memswap") {

		var MemswapFlagName string
		if cmdPrefix == "" {
			MemswapFlagName = "memswap"
		} else {
			MemswapFlagName = fmt.Sprintf("%v.memswap", cmdPrefix)
		}

		MemswapFlagValue, err := cmd.Flags().GetInt64(MemswapFlagName)
		if err != nil {
			return err, false
		}
		m.Memswap = &MemswapFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildNetworkmodeFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("networkmode") {

		var NetworkmodeFlagName string
		if cmdPrefix == "" {
			NetworkmodeFlagName = "networkmode"
		} else {
			NetworkmodeFlagName = fmt.Sprintf("%v.networkmode", cmdPrefix)
		}

		NetworkmodeFlagValue, err := cmd.Flags().GetString(NetworkmodeFlagName)
		if err != nil {
			return err, false
		}
		m.Networkmode = &NetworkmodeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildNocacheFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("nocache") {

		var NocacheFlagName string
		if cmdPrefix == "" {
			NocacheFlagName = "nocache"
		} else {
			NocacheFlagName = fmt.Sprintf("%v.nocache", cmdPrefix)
		}

		NocacheFlagValue, err := cmd.Flags().GetBool(NocacheFlagName)
		if err != nil {
			return err, false
		}
		m.Nocache = &NocacheFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildOutputsFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("outputs") {

		var OutputsFlagName string
		if cmdPrefix == "" {
			OutputsFlagName = "outputs"
		} else {
			OutputsFlagName = fmt.Sprintf("%v.outputs", cmdPrefix)
		}

		OutputsFlagValue, err := cmd.Flags().GetString(OutputsFlagName)
		if err != nil {
			return err, false
		}
		m.Outputs = &OutputsFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildPlatformFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("platform") {

		var PlatformFlagName string
		if cmdPrefix == "" {
			PlatformFlagName = "platform"
		} else {
			PlatformFlagName = fmt.Sprintf("%v.platform", cmdPrefix)
		}

		PlatformFlagValue, err := cmd.Flags().GetString(PlatformFlagName)
		if err != nil {
			return err, false
		}
		m.Platform = &PlatformFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildPullFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("pull") {

		var PullFlagName string
		if cmdPrefix == "" {
			PullFlagName = "pull"
		} else {
			PullFlagName = fmt.Sprintf("%v.pull", cmdPrefix)
		}

		PullFlagValue, err := cmd.Flags().GetString(PullFlagName)
		if err != nil {
			return err, false
		}
		m.Pull = &PullFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildQFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("q") {

		var QFlagName string
		if cmdPrefix == "" {
			QFlagName = "q"
		} else {
			QFlagName = fmt.Sprintf("%v.q", cmdPrefix)
		}

		QFlagValue, err := cmd.Flags().GetBool(QFlagName)
		if err != nil {
			return err, false
		}
		m.Q = &QFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildRemoteFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("remote") {

		var RemoteFlagName string
		if cmdPrefix == "" {
			RemoteFlagName = "remote"
		} else {
			RemoteFlagName = fmt.Sprintf("%v.remote", cmdPrefix)
		}

		RemoteFlagValue, err := cmd.Flags().GetString(RemoteFlagName)
		if err != nil {
			return err, false
		}
		m.Remote = &RemoteFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildRmFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("rm") {

		var RmFlagName string
		if cmdPrefix == "" {
			RmFlagName = "rm"
		} else {
			RmFlagName = fmt.Sprintf("%v.rm", cmdPrefix)
		}

		RmFlagValue, err := cmd.Flags().GetBool(RmFlagName)
		if err != nil {
			return err, false
		}
		m.Rm = &RmFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildShmsizeFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("shmsize") {

		var ShmsizeFlagName string
		if cmdPrefix == "" {
			ShmsizeFlagName = "shmsize"
		} else {
			ShmsizeFlagName = fmt.Sprintf("%v.shmsize", cmdPrefix)
		}

		ShmsizeFlagValue, err := cmd.Flags().GetInt64(ShmsizeFlagName)
		if err != nil {
			return err, false
		}
		m.Shmsize = &ShmsizeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildSquashFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("squash") {

		var SquashFlagName string
		if cmdPrefix == "" {
			SquashFlagName = "squash"
		} else {
			SquashFlagName = fmt.Sprintf("%v.squash", cmdPrefix)
		}

		SquashFlagValue, err := cmd.Flags().GetBool(SquashFlagName)
		if err != nil {
			return err, false
		}
		m.Squash = &SquashFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildTFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("t") {

		var TFlagName string
		if cmdPrefix == "" {
			TFlagName = "t"
		} else {
			TFlagName = fmt.Sprintf("%v.t", cmdPrefix)
		}

		TFlagValue, err := cmd.Flags().GetString(TFlagName)
		if err != nil {
			return err, false
		}
		m.T = &TFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageBuildTargetFlag(m *image.ImageBuildParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("target") {

		var TargetFlagName string
		if cmdPrefix == "" {
			TargetFlagName = "target"
		} else {
			TargetFlagName = fmt.Sprintf("%v.target", cmdPrefix)
		}

		TargetFlagValue, err := cmd.Flags().GetString(TargetFlagName)
		if err != nil {
			return err, false
		}
		m.Target = &TargetFlagValue

	}
	return nil, retAdded
}

// parseOperationImageImageBuildResult parses request result and return the string content
func parseOperationImageImageBuildResult(resp0 *image.ImageBuildOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning imageBuildOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*image.ImageBuildBadRequest)
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
		resp2, ok := iResp2.(*image.ImageBuildInternalServerError)
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

	// warning: non schema response imageBuildOK is not supported by go-swagger cli yet.

	return "", nil
}

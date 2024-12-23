package vsphere

import (
	"bytes"
	"fmt"
	"path"
	"regexp"

	_ "embed" // Allow embedding state

	"github.com/hashicorp/terraform-provider-vsphere/vsphere"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/info"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	tks "github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumi/pulumi-vsphere/provider/v4/pkg/version"
)

const (
	vspherePkg = "vsphere"
	vsphereMod = "index"
)

//go:embed cmd/pulumi-resource-vsphere/bridge-metadata.json
var metadata []byte

func vsphereDataSource(mod string, res string) tks.ModuleMember {
	return tfbridge.MakeDataSource(vspherePkg, mod, res)
}

func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(vsphere.Provider())
	prov := tfbridge.ProviderInfo{
		P:                p,
		Name:             "vsphere",
		Description:      "A Pulumi package for creating vsphere resources",
		Keywords:         []string{"pulumi", "vsphere"},
		License:          "Apache-2.0",
		Homepage:         "https://pulumi.io",
		Repository:       "https://github.com/pulumi/pulumi-vsphere",
		GitHubOrg:        "hashicorp",
		UpstreamRepoPath: "./upstream",
		Version:          version.Version,
		MetadataInfo:     tfbridge.NewProviderMetadata(metadata),
		DocRules:         &info.DocRule{EditRules: editRules},
		Config: map[string]*tfbridge.SchemaInfo{
			"allow_unverified_ssl": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_ALLOW_UNVERIFIED_SSL"},
				},
			},
			"client_debug": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_CLIENT_DEBUG"},
				},
			},
			"client_debug_path_run": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_CLIENT_DEBUG_PATH_RUN"},
				},
			},
			"client_debug_path": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_CLIENT_DEBUG_PATH"},
				},
			},
			"persist_session": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_PERSIST_SESSION"},
				},
			},
			"vim_session_path": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_VIM_SESSION_PATH"},
				},
			},
			"rest_session_path": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_REST_SESSION_PATH"},
				},
			},
			"vim_keep_alive": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VSPHERE_VIM_KEEP_ALIVE"},
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"vsphere_storage_policy": {Tok: vsphereDataSource(vsphereMod, "getPolicy")},
			"vsphere_datastore_stats": {Docs: &tfbridge.DocInfo{
				Source: "datastore_stats.html .markdown", // Note extra space.
			}},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			DevDependencies: map[string]string{
				"@types/node": "^10.9.2",
			},
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			RespectSchemaVersion: true,

			PyProject: struct{ Enabled bool }{true},
		},

		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", vspherePkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				vspherePkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				"vsphere": "VSphere",
			},
		},
	}

	prov.MustComputeTokens(tokens.SingleModule("vsphere_", vsphereMod,
		tokens.MakeStandard(vspherePkg)))
	prov.SetAutonaming(255, "-")
	prov.MustApplyAutoAliases()

	return prov
}

func editRules(defaults []info.DocsEdit) []info.DocsEdit {
	textReplace := func(old, new string) info.DocsEdit {
		o, n := []byte(old), []byte(new)
		return info.DocsEdit{
			Path: "*",
			Edit: func(_ string, content []byte) ([]byte, error) {
				return bytes.ReplaceAll(content, o, n), nil
			},
		}
	}

	docImportLink := regexp.MustCompile(`\[([^\]]+)\]\[docs-import\]`)
	stripLink := info.DocsEdit{
		Path: "*",
		Edit: func(_ string, content []byte) ([]byte, error) {
			return docImportLink.ReplaceAll(content, []byte("$1")), nil
		},
	}
	return append(defaults,
		textReplace("## Importing", "## Import"),
		textReplace(`"terraform-`, `"pulumi-`),
		stripLink, // Strip labeled markdown linkes of the form [<something>][docs-import]
		textReplace("\n[docs-import]: https://www.terraform.io/docs/import/index.html\n", "\n"),
		textReplace("\nAll information will be added to the Terraform state after import.\n", "\n"),
	)
}

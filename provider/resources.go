// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"path"

	_ "embed" // Embed bridge metadata

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"

	"github.com/pulumi/pulumi-local/provider/pkg/version"
)

//go:embed cmd/pulumi-resource-local/bridge-metadata.json
var metadata []byte

const providerName = "local"

var pkgVersion = version.Version

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:                ShimmedProvider(),
		Name:             providerName,
		MetadataInfo:     tfbridge.NewProviderMetadata(metadata),
		Version:          pkgVersion,
		UpstreamRepoPath: "./upstream",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Local",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher:   "Pulumi",
		Description: "A Pulumi package for creating and managing Local cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com/",
		Repository: "https://github.com/pulumi/pulumi-local",
		// The GitHub Org hosting the upstream provider - defaults to `terraform-providers`. Note that
		// this should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "terraform-providers",
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^16.0.0", // so we can access strongly typed node definitions.
			},
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			RespectSchemaVersion: true,
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
			PyProject: struct{ Enabled bool }{true},
		},

		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				"github.com/pulumi/pulumi-local/sdk",
				tfbridge.GetModuleMajorVersion(pkgVersion),
				"go",
				providerName,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.MustComputeTokens(tokens.SingleModule(providerName+"_", "index",
		tokens.MakeStandard(providerName)))

	prov.SetAutonaming(255, "-")

	prov.MustApplyAutoAliases()

	return prov
}

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

package dbtcloud

import (
	_ "embed" // to embed bridge metadata
	"fmt"
	"path/filepath"

	"github.com/a-schot/pulumi-dbtcloud/provider/pkg/version"
	dbtcloud "github.com/dbt-labs/terraform-provider-dbtcloud/pkg/provider"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tfbridgetokens "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
)

//go:embed cmd/pulumi-resource-dbtcloud/bridge-metadata.json
var bridgeMetadata []byte

// all of the token components used below.
const (
	mainPkg = "dbtcloud"
	mainMod = "index"
)

// Provider returns additional overlaid schema and metadata associated with the provider
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(dbtcloud.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "dbtcloud",
		DisplayName:       "dbt Cloud",
		Publisher:         "a-schot",
		LogoURL:           "https://docs.getdbt.com/img/dbt-logo-light.svg",
		PluginDownloadURL: "https://github.com/a-schot/pulumi-dbtcloud/releases/download/v${VERSION}",
		Description:       "A Pulumi package for creating and managing dbt Cloud resources.",
		Keywords:          []string{"pulumi", "dbtcloud", "dbt", "cloud", "category/cloud", "category/database"},
		License:           "Apache-2.0",
		Homepage:          "https://www.pulumi.com",
		Repository:        "https://github.com/a-schot/pulumi-dbtcloud",
		GitHubOrg:         "dbt-labs",
		TFProviderLicense: tfbridge.SetProviderLicense(tfbridge.MITLicenseType),
		MetadataInfo:      tfbridge.NewProviderMetadata(bridgeMetadata),
		DocRules:          &tfbridge.DocRuleInfo{EditRules: editRules},

		// Ignore legacy tokens (prefix: "dbt_cloud_" instead of "dbtcloud_")
		// These will be removed from upstream at v0.3
		IgnoreMappings: []string{
			"dbt_cloud_bigquery_connection",
			"dbt_cloud_bigquery_credential",
			"dbt_cloud_connection",
			"dbt_cloud_databricks_credential",
			"dbt_cloud_environment",
			"dbt_cloud_environment_variable",
			"dbt_cloud_group",
			"dbt_cloud_job",
			"dbt_cloud_postgres_credential",
			"dbt_cloud_project",
			"dbt_cloud_project_artefacts",
			"dbt_cloud_project_connection",
			"dbt_cloud_project_repository",
			"dbt_cloud_repository",
			"dbt_cloud_service_token",
			"dbt_cloud_snowflake_credential",
			"dbt_cloud_webhook",
			"dbt_cloud_group",
			"dbt_cloud_user",
			"dbt_cloud_privatelink_endpoint",
		},
		Config: map[string]*tfbridge.SchemaInfo{
			"account_id": {
				Default: &tfbridge.DefaultInfo{EnvVars: []string{"DBT_CLOUD_ACCOUNT_ID"}},
			},
			"token": {
				Default: &tfbridge.DefaultInfo{EnvVars: []string{"DBT_CLOUD_TOKEN"}},
			},
			"host_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"DBT_CLOUD_HOST_URL"},
					Value:   "https://cloud.getdbt.com/api",
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"dbtcloud_bigquery_credential": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "BigQueryCredential"),
			},
			"dbtcloud_bigquery_connection": {
				Tok:  tfbridge.MakeResource(mainPkg, mainMod, "BigQueryConnection"),
				Docs: &tfbridge.DocInfo{Markdown: []byte(``)},
			},
			"dbtcloud_extended_attributes": {
				Tok:        tfbridge.MakeResource(mainPkg, mainMod, "ExtendedAttributes"),
				CSharpName: "NameExtendedAttributes",
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"dbtcloud_bigquery_credential": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBigQueryCredential"),
			},
			"dbtcloud_bigquery_connection": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBigQueryConnection"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@aschot/pulumi-dbtcloud",

			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "aschot_pulumi_dbtcloud",

			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/a-schot/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "ASchot.Pulumi",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.MustComputeTokens(
		tfbridgetokens.SingleModule(
			"dbtcloud_",
			mainMod,
			tfbridgetokens.MakeStandard(mainPkg),
		).Ignore("dbt_cloud"), // to ignore legacy resources; this is doing nothing for now
	)

	prov.MustApplyAutoAliases()

	prov.SetAutonaming(255, "-")

	return prov
}

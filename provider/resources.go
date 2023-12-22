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
	_ "embed"
	"fmt"
	"path/filepath"

	"github.com/a-schot/pulumi-dbtcloud/provider/pkg/version"
	dbtcloud "github.com/dbt-labs/terraform-provider-dbtcloud/pkg/provider"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tfbridgetokens "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

//go:embed cmd/pulumi-resource-dbtcloud/bridge-metadata.json
var bridgeMetadata []byte

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "dbtcloud"
	// modules:
	mainMod = "index" // the dbtcloud module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(dbtcloud.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                    p,
		Name:                 "dbtcloud",
		DisplayName:          "dbt Cloud",
		Publisher:            "a-schot",
		LogoURL:              "",
		PluginDownloadURL:    "https://github.com/a-schot/pulumi-dbtcloud/releases/download/v${VERSION}",
		Description:          "A Pulumi package for creating and managing dbt Cloud resources.",
		Keywords:             []string{"pulumi", "dbtcloud", "category/cloud", "dbt", "cloud"},
		License:              "Apache-2.0",
		Homepage:             "https://www.pulumi.com",
		Repository:           "https://github.com/a-schot/pulumi-dbtcloud",
		GitHubOrg:            "dbt-labs",
		TFProviderLicense:    tfbridge.SetProviderLicense(tfbridge.MITLicenseType),
		MetadataInfo:         tfbridge.NewProviderMetadata(bridgeMetadata),
		PreConfigureCallback: preConfigureCallback,

		IgnoreMappings: []string{
			// Ignore legacy tokens (prefix: dbt_cloud_ instead of dbtcloud_)
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
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			// "dbtcloud_job":                               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Job")},
			// "dbtcloud_project":                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Project")},
			// "dbtcloud_project_connection":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ProjectConnection")},
			// "dbtcloud_project_repository":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ProjectRepository")},
			// "dbtcloud_project_artefacts":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ProjectArtefacts")},
			// "dbtcloud_environment":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Environment")},
			// "dbtcloud_environment_variable":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "EnvironmentVariable")},
			// "dbtcloud_databricks_credential":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DatabricksCredential")},
			// "dbtcloud_snowflake_credential":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SnowflakeCredential")},
			"dbtcloud_bigquery_credential": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "BigQueryCredential")},
			// "dbtcloud_postgres_credential":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "PostgresCredential")},
			// "dbtcloud_connection":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Connection")},
			"dbtcloud_bigquery_connection": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "BigQueryConnection")},
			// "dbtcloud_repository":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Repository")},
			// "dbtcloud_group":                             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Group")},
			// "dbtcloud_service_token":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ServiceToken")},
			// "dbtcloud_webhook":                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Webhook")},
			// "dbtcloud_notification":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Notification")},
			// "dbtcloud_user_groups":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "UserGroups")},
			// "dbtcloud_license_map":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LicenseMap")},
			// "dbtcloud_environment_variable_job_override": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "EnvironmentVariableJobOverride")},
			// "dbtcloud_fabric_connection":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FabricConnection")},
			// "dbtcloud_fabric_credential":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FabricCredential")},
			"dbtcloud_extended_attributes": {
				Tok:        tfbridge.MakeResource(mainPkg, mainMod, "ExtendedAttributes"),
				CSharpName: "NameExtendedAttributes",
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// "dbtcloud_group":                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getGroup")},
			// "dbtcloud_job":                      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getJob")},
			// "dbtcloud_project":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getProject")},
			// "dbtcloud_environment":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getEnvironment")},
			// "dbtcloud_environment_variable":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getEnvironmentVariable")},
			// "dbtcloud_snowflake_credential":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSnowflakeCredential")},
			"dbtcloud_bigquery_credential": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBigQueryCredential")},
			// "dbtcloud_postgres_credential":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getPostgresCredential")},
			// "dbtcloud_databricks_credential":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDatabricksCredential")},
			// "dbtcloud_connection":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getConnection")},
			"dbtcloud_bigquery_connection": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBigQueryConnection")},
			// "dbtcloud_repository":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRepository")},
			// "dbtcloud_user":                     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUser")},
			// "dbtcloud_service_token":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getServiceToken")},
			// "dbtcloud_webhook":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getWebhook")},
			// "dbtcloud_privatelink_endpoint":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getPrivatelinkEndpoint")},
			// "dbtcloud_notification":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNotification")},
			// "dbtcloud_user_groups":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUserGroups")},
			// "dbtcloud_extended_attributes":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getExtendedAttributes")},
			// "dbtcloud_group_users":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getGroupUsers")},
			// "dbtcloud_azure_dev_ops_project":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAzureDevOpsProject")},
			// "dbtcloud_azure_dev_ops_repository": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAzureDevOpsRepository")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@aschot/pulumi-dbtcloud",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "aschot_pulumi_dbtcloud",
			// List any Python dependencies and their version ranges
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

	// These are new API's that you may opt to use to automatically compute resource tokens,
	// and apply auto aliasing for full backwards compatibility.
	// For more information, please reference: https://pkg.go.dev/github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge#ProviderInfo.ComputeTokens

	var strat tfbridge.Strategy = tfbridgetokens.SingleModule(
		"dbtcloud_",
		mainMod,
		tfbridgetokens.MakeStandard(mainPkg),
	).Ignore("dbt_cloud")

	prov.MustComputeTokens(strat)

	prov.MustApplyAutoAliases()

	prov.SetAutonaming(255, "-")

	return prov
}

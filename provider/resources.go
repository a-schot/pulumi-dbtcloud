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
		P:    p,
		Name: "dbtcloud",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "dbt Cloud",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "a-schot",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "https://github.com/a-schot/pulumi-dbtcloud/releases/download/v${VERSION}",
		Description:       "A Pulumi package for creating and managing dbt Cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "dbtcloud", "category/cloud", "dbt", "cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/a-schot/pulumi-dbtcloud",
		GitHubOrg:  "dbt-labs",
		Config:     map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },
			"dbtcloud_job":                               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Job")},
			"dbtcloud_project":                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Project")},
			"dbtcloud_project_connection":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ProjectConnection")},
			"dbtcloud_project_repository":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ProjectRepository")},
			"dbtcloud_project_artefacts":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ProjectArtefacts")},
			"dbtcloud_environment":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Environment")},
			"dbtcloud_environment_variable":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "EnvironmentVariable")},
			"dbtcloud_databricks_credential":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DatabricksCredential")},
			"dbtcloud_snowflake_credential":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SnowflakeCredential")},
			"dbtcloud_bigquery_credential":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "BigqueryCredential")},
			"dbtcloud_postgres_credential":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "PostgresCredential")},
			"dbtcloud_connection":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Connection")},
			"dbtcloud_bigquery_connection":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "BigqueryConnection")},
			"dbtcloud_repository":                        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Repository")},
			"dbtcloud_group":                             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Group")},
			"dbtcloud_service_token":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ServiceToken")},
			"dbtcloud_webhook":                           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Webhook")},
			"dbtcloud_notification":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Notification")},
			"dbtcloud_user_groups":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "UserGroups")},
			"dbtcloud_license_map":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LicenseMap")},
			"dbtcloud_extended_attributes":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ExtendedAttributes")},
			"dbtcloud_environment_variable_job_override": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "EnvironmentVariableJobOverride")},
			"dbtcloud_fabric_connection":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FabricConnection")},
			"dbtcloud_fabric_credential":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FabricCredential")},
			// legacy tokens will be removed in 0.3 for the dbt Cloud TF provider
			"dbt_cloud_job":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacyJob")},
			"dbt_cloud_project":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacyProject")},
			"dbt_cloud_project_connection":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacyProjectConnection")},
			"dbt_cloud_project_repository":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacyProjectRepository")},
			"dbt_cloud_project_artefacts":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacyProjectArtefacts")},
			"dbt_cloud_environment":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacyEnvironment")},
			"dbt_cloud_environment_variable":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacyEnvironmentVariable")},
			"dbt_cloud_databricks_credential": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacyDatabricksCredential")},
			"dbt_cloud_snowflake_credential":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacySnowflakeCredential")},
			"dbt_cloud_bigquery_credential":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacyBigqueryCredential")},
			"dbt_cloud_postgres_credential":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacyPostgresCredential")},
			"dbt_cloud_connection":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacyConnection")},
			"dbt_cloud_bigquery_connection":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacyBigqueryConnection")},
			"dbt_cloud_repository":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacyRepository")},
			"dbt_cloud_group":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacyGroup")},
			"dbt_cloud_service_token":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacyServiceToken")},
			"dbt_cloud_webhook":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LegacyWebhook")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
			"dbtcloud_group":                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getGroup")},
			"dbtcloud_job":                      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getJob")},
			"dbtcloud_project":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getProject")},
			"dbtcloud_environment":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getEnvironment")},
			"dbtcloud_environment_variable":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getEnvironmentVariable")},
			"dbtcloud_snowflake_credential":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSnowflakeCredential")},
			"dbtcloud_bigquery_credential":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBigQueryCredential")},
			"dbtcloud_postgres_credential":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getPostgresCredential")},
			"dbtcloud_databricks_credential":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDatabricksCredential")},
			"dbtcloud_connection":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getConnection")},
			"dbtcloud_bigquery_connection":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBigqueryConnection")},
			"dbtcloud_repository":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRepository")},
			"dbtcloud_user":                     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUser")},
			"dbtcloud_service_token":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getServiceToken")},
			"dbtcloud_webhook":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getWebhook")},
			"dbtcloud_privatelink_endpoint":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getPrivatelinkEndpoint")},
			"dbtcloud_notification":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNotification")},
			"dbtcloud_user_groups":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUserGroups")},
			"dbtcloud_extended_attributes":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getExtendedAttributes")},
			"dbtcloud_group_users":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getGroupUsers")},
			"dbtcloud_azure_dev_ops_project":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAzureDevOpsProject")},
			"dbtcloud_azure_dev_ops_repository": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAzureDevOpsRepository")},
			// legacy tokens will be removed in 0.3 for the dbt Cloud TF provider
			"dbt_cloud_group":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "LegacyGetGroup")},
			"dbt_cloud_job":                   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "LegacyGetJob")},
			"dbt_cloud_project":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "LegacyGetProject")},
			"dbt_cloud_environment":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "LegacyGetEnvironment")},
			"dbt_cloud_environment_variable":  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "LegacyGetEnvironmentVariable")},
			"dbt_cloud_snowflake_credential":  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "LegacyGetSnowflakeCredential")},
			"dbt_cloud_bigquery_credential":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "LegacyGetBigqueryCredential")},
			"dbt_cloud_postgres_credential":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "LegacyGetPostgresCredential")},
			"dbt_cloud_databricks_credential": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "LegacyGetDatabricksCredential")},
			"dbt_cloud_connection":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "LegacyGetConnection")},
			"dbt_cloud_bigquery_connection":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "LegacyGetBigqueryConnection")},
			"dbt_cloud_repository":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "LegacyGetRepository")},
			"dbt_cloud_user":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "LegacyGetUser")},
			"dbt_cloud_service_token":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "LegacyGetServiceToken")},
			"dbt_cloud_webhook":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "LegacyGetWebhook")},
			"dbt_cloud_privatelink_endpoint":  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "LegacyGetPrivatelinkEndpoint")},
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
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
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

	prov.MustComputeTokens(tfbridgetokens.SingleModule(
		"dbtcloud_",
		mainMod,
		tfbridgetokens.MakeStandard(mainPkg),
	))
	// prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}

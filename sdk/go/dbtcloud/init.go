// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbtcloud

import (
	"fmt"

	"github.com/a-schot/pulumi-dbtcloud/sdk/go/dbtcloud/internal"
	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "dbtcloud:index/bigQueryConnection:BigQueryConnection":
		r = &BigQueryConnection{}
	case "dbtcloud:index/bigQueryCredential:BigQueryCredential":
		r = &BigQueryCredential{}
	case "dbtcloud:index/connection:Connection":
		r = &Connection{}
	case "dbtcloud:index/databricksCredential:DatabricksCredential":
		r = &DatabricksCredential{}
	case "dbtcloud:index/environment:Environment":
		r = &Environment{}
	case "dbtcloud:index/environmentVariable:EnvironmentVariable":
		r = &EnvironmentVariable{}
	case "dbtcloud:index/environmentVariableJobOverride:EnvironmentVariableJobOverride":
		r = &EnvironmentVariableJobOverride{}
	case "dbtcloud:index/extendedAttributes:ExtendedAttributes":
		r = &ExtendedAttributes{}
	case "dbtcloud:index/fabricConnection:FabricConnection":
		r = &FabricConnection{}
	case "dbtcloud:index/fabricCredential:FabricCredential":
		r = &FabricCredential{}
	case "dbtcloud:index/group:Group":
		r = &Group{}
	case "dbtcloud:index/job:Job":
		r = &Job{}
	case "dbtcloud:index/legacyBigqueryConnection:LegacyBigqueryConnection":
		r = &LegacyBigqueryConnection{}
	case "dbtcloud:index/legacyBigqueryCredential:LegacyBigqueryCredential":
		r = &LegacyBigqueryCredential{}
	case "dbtcloud:index/legacyConnection:LegacyConnection":
		r = &LegacyConnection{}
	case "dbtcloud:index/legacyDatabricksCredential:LegacyDatabricksCredential":
		r = &LegacyDatabricksCredential{}
	case "dbtcloud:index/legacyEnvironment:LegacyEnvironment":
		r = &LegacyEnvironment{}
	case "dbtcloud:index/legacyEnvironmentVariable:LegacyEnvironmentVariable":
		r = &LegacyEnvironmentVariable{}
	case "dbtcloud:index/legacyGroup:LegacyGroup":
		r = &LegacyGroup{}
	case "dbtcloud:index/legacyJob:LegacyJob":
		r = &LegacyJob{}
	case "dbtcloud:index/legacyPostgresCredential:LegacyPostgresCredential":
		r = &LegacyPostgresCredential{}
	case "dbtcloud:index/legacyProject:LegacyProject":
		r = &LegacyProject{}
	case "dbtcloud:index/legacyProjectArtefacts:LegacyProjectArtefacts":
		r = &LegacyProjectArtefacts{}
	case "dbtcloud:index/legacyProjectConnection:LegacyProjectConnection":
		r = &LegacyProjectConnection{}
	case "dbtcloud:index/legacyProjectRepository:LegacyProjectRepository":
		r = &LegacyProjectRepository{}
	case "dbtcloud:index/legacyRepository:LegacyRepository":
		r = &LegacyRepository{}
	case "dbtcloud:index/legacyServiceToken:LegacyServiceToken":
		r = &LegacyServiceToken{}
	case "dbtcloud:index/legacySnowflakeCredential:LegacySnowflakeCredential":
		r = &LegacySnowflakeCredential{}
	case "dbtcloud:index/legacyWebhook:LegacyWebhook":
		r = &LegacyWebhook{}
	case "dbtcloud:index/licenseMap:LicenseMap":
		r = &LicenseMap{}
	case "dbtcloud:index/notification:Notification":
		r = &Notification{}
	case "dbtcloud:index/postgresCredential:PostgresCredential":
		r = &PostgresCredential{}
	case "dbtcloud:index/project:Project":
		r = &Project{}
	case "dbtcloud:index/projectArtefacts:ProjectArtefacts":
		r = &ProjectArtefacts{}
	case "dbtcloud:index/projectConnection:ProjectConnection":
		r = &ProjectConnection{}
	case "dbtcloud:index/projectRepository:ProjectRepository":
		r = &ProjectRepository{}
	case "dbtcloud:index/repository:Repository":
		r = &Repository{}
	case "dbtcloud:index/serviceToken:ServiceToken":
		r = &ServiceToken{}
	case "dbtcloud:index/snowflakeCredential:SnowflakeCredential":
		r = &SnowflakeCredential{}
	case "dbtcloud:index/userGroups:UserGroups":
		r = &UserGroups{}
	case "dbtcloud:index/webhook:Webhook":
		r = &Webhook{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:dbtcloud" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/bigQueryConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/bigQueryCredential",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/connection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/databricksCredential",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/environment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/environmentVariable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/environmentVariableJobOverride",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/extendedAttributes",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/fabricConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/fabricCredential",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/job",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacyBigqueryConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacyBigqueryCredential",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacyConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacyDatabricksCredential",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacyEnvironment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacyEnvironmentVariable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacyGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacyJob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacyPostgresCredential",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacyProject",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacyProjectArtefacts",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacyProjectConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacyProjectRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacyRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacyServiceToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacySnowflakeCredential",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/legacyWebhook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/licenseMap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/notification",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/postgresCredential",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/project",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/projectArtefacts",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/projectConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/projectRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/repository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/serviceToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/snowflakeCredential",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/userGroups",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"dbtcloud",
		"index/webhook",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"dbtcloud",
		&pkg{version},
	)
}

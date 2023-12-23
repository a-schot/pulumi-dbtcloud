// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/a-schot/pulumi-dbtcloud/sdk/go/dbtcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// Account identifier for your dbt Cloud implementation. Instead of setting the parameter, you can set the environment
// variable `DBT_CLOUD_ACCOUNT_ID`
func GetAccountId(ctx *pulumi.Context) int {
	v, err := config.TryInt(ctx, "dbtcloud:accountId")
	if err == nil {
		return v
	}
	var value int
	if d := internal.GetEnvOrDefault(nil, internal.ParseEnvInt, "DBT_CLOUD_ACCOUNT_ID"); d != nil {
		value = d.(int)
	}
	return value
}

// URL for your dbt Cloud deployment. Instead of setting the parameter, you can set the environment variable
// `DBT_CLOUD_HOST_URL` - Defaults to https://cloud.getdbt.com/api
func GetHostUrl(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "dbtcloud:hostUrl")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault("https://cloud.getdbt.com/api", nil, "DBT_CLOUD_HOST_URL"); d != nil {
		value = d.(string)
	}
	return value
}

// API token for your dbt Cloud. Instead of setting the parameter, you can set the environment variable `DBT_CLOUD_TOKEN`
func GetToken(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "dbtcloud:token")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault(nil, nil, "DBT_CLOUD_TOKEN"); d != nil {
		value = d.(string)
	}
	return value
}

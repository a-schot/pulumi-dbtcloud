// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbtcloud

import (
	"context"
	"reflect"

	"github.com/a-schot/pulumi-dbtcloud/sdk/go/dbtcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentResult
	err := ctx.Invoke("dbtcloud:index/getEnvironment:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEnvironment.
type LookupEnvironmentArgs struct {
	// ID of the environment
	EnvironmentId int `pulumi:"environmentId"`
	// Project ID to create the environment in
	ProjectId int `pulumi:"projectId"`
}

// A collection of values returned by getEnvironment.
type LookupEnvironmentResult struct {
	// Credential ID to create the environment with
	CredentialId int `pulumi:"credentialId"`
	// Which custom branch to use in this environment
	CustomBranch string `pulumi:"customBranch"`
	// Version number of dbt to use in this environment, usually in the format 1.2.0-latest rather than core versions
	DbtVersion string `pulumi:"dbtVersion"`
	// The type of deployment environment (currently 'production' or empty)
	DeploymentType string `pulumi:"deploymentType"`
	// ID of the environment
	EnvironmentId int `pulumi:"environmentId"`
	// The ID of the extended attributes applied
	ExtendedAttributesId int `pulumi:"extendedAttributesId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Whether the environment is active
	IsActive bool `pulumi:"isActive"`
	// Environment name
	Name string `pulumi:"name"`
	// Project ID to create the environment in
	ProjectId int `pulumi:"projectId"`
	// The type of environment (must be either development or deployment)
	Type string `pulumi:"type"`
	// Whether to use a custom git branch in this environment
	UseCustomBranch bool `pulumi:"useCustomBranch"`
}

func LookupEnvironmentOutput(ctx *pulumi.Context, args LookupEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentResult, error) {
			args := v.(LookupEnvironmentArgs)
			r, err := LookupEnvironment(ctx, &args, opts...)
			var s LookupEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentResultOutput)
}

// A collection of arguments for invoking getEnvironment.
type LookupEnvironmentOutputArgs struct {
	// ID of the environment
	EnvironmentId pulumi.IntInput `pulumi:"environmentId"`
	// Project ID to create the environment in
	ProjectId pulumi.IntInput `pulumi:"projectId"`
}

func (LookupEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentArgs)(nil)).Elem()
}

// A collection of values returned by getEnvironment.
type LookupEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentResult)(nil)).Elem()
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutput() LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutputWithContext(ctx context.Context) LookupEnvironmentResultOutput {
	return o
}

// Credential ID to create the environment with
func (o LookupEnvironmentResultOutput) CredentialId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) int { return v.CredentialId }).(pulumi.IntOutput)
}

// Which custom branch to use in this environment
func (o LookupEnvironmentResultOutput) CustomBranch() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.CustomBranch }).(pulumi.StringOutput)
}

// Version number of dbt to use in this environment, usually in the format 1.2.0-latest rather than core versions
func (o LookupEnvironmentResultOutput) DbtVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.DbtVersion }).(pulumi.StringOutput)
}

// The type of deployment environment (currently 'production' or empty)
func (o LookupEnvironmentResultOutput) DeploymentType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.DeploymentType }).(pulumi.StringOutput)
}

// ID of the environment
func (o LookupEnvironmentResultOutput) EnvironmentId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) int { return v.EnvironmentId }).(pulumi.IntOutput)
}

// The ID of the extended attributes applied
func (o LookupEnvironmentResultOutput) ExtendedAttributesId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) int { return v.ExtendedAttributesId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether the environment is active
func (o LookupEnvironmentResultOutput) IsActive() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) bool { return v.IsActive }).(pulumi.BoolOutput)
}

// Environment name
func (o LookupEnvironmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Project ID to create the environment in
func (o LookupEnvironmentResultOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) int { return v.ProjectId }).(pulumi.IntOutput)
}

// The type of environment (must be either development or deployment)
func (o LookupEnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

// Whether to use a custom git branch in this environment
func (o LookupEnvironmentResultOutput) UseCustomBranch() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) bool { return v.UseCustomBranch }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentResultOutput{})
}
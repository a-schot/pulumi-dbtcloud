// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbtcloud

import (
	"context"
	"reflect"

	"github.com/a-schot/pulumi-dbtcloud/sdk/go/dbtcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deprecated: Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource
func LegacyGetEnvironmentVariable(ctx *pulumi.Context, args *LegacyGetEnvironmentVariableArgs, opts ...pulumi.InvokeOption) (*LegacyGetEnvironmentVariableResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LegacyGetEnvironmentVariableResult
	err := ctx.Invoke("dbtcloud:index/legacyGetEnvironmentVariable:LegacyGetEnvironmentVariable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking LegacyGetEnvironmentVariable.
type LegacyGetEnvironmentVariableArgs struct {
	Name      string `pulumi:"name"`
	ProjectId int    `pulumi:"projectId"`
}

// A collection of values returned by LegacyGetEnvironmentVariable.
type LegacyGetEnvironmentVariableResult struct {
	EnvironmentValues map[string]interface{} `pulumi:"environmentValues"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	Name      string `pulumi:"name"`
	ProjectId int    `pulumi:"projectId"`
}

func LegacyGetEnvironmentVariableOutput(ctx *pulumi.Context, args LegacyGetEnvironmentVariableOutputArgs, opts ...pulumi.InvokeOption) LegacyGetEnvironmentVariableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LegacyGetEnvironmentVariableResult, error) {
			args := v.(LegacyGetEnvironmentVariableArgs)
			r, err := LegacyGetEnvironmentVariable(ctx, &args, opts...)
			var s LegacyGetEnvironmentVariableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LegacyGetEnvironmentVariableResultOutput)
}

// A collection of arguments for invoking LegacyGetEnvironmentVariable.
type LegacyGetEnvironmentVariableOutputArgs struct {
	Name      pulumi.StringInput `pulumi:"name"`
	ProjectId pulumi.IntInput    `pulumi:"projectId"`
}

func (LegacyGetEnvironmentVariableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LegacyGetEnvironmentVariableArgs)(nil)).Elem()
}

// A collection of values returned by LegacyGetEnvironmentVariable.
type LegacyGetEnvironmentVariableResultOutput struct{ *pulumi.OutputState }

func (LegacyGetEnvironmentVariableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LegacyGetEnvironmentVariableResult)(nil)).Elem()
}

func (o LegacyGetEnvironmentVariableResultOutput) ToLegacyGetEnvironmentVariableResultOutput() LegacyGetEnvironmentVariableResultOutput {
	return o
}

func (o LegacyGetEnvironmentVariableResultOutput) ToLegacyGetEnvironmentVariableResultOutputWithContext(ctx context.Context) LegacyGetEnvironmentVariableResultOutput {
	return o
}

func (o LegacyGetEnvironmentVariableResultOutput) EnvironmentValues() pulumi.MapOutput {
	return o.ApplyT(func(v LegacyGetEnvironmentVariableResult) map[string]interface{} { return v.EnvironmentValues }).(pulumi.MapOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LegacyGetEnvironmentVariableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LegacyGetEnvironmentVariableResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LegacyGetEnvironmentVariableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LegacyGetEnvironmentVariableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LegacyGetEnvironmentVariableResultOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v LegacyGetEnvironmentVariableResult) int { return v.ProjectId }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LegacyGetEnvironmentVariableResultOutput{})
}

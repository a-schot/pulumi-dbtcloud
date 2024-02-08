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
func LegacyGetJob(ctx *pulumi.Context, args *LegacyGetJobArgs, opts ...pulumi.InvokeOption) (*LegacyGetJobResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LegacyGetJobResult
	err := ctx.Invoke("dbtcloud:index/legacyGetJob:LegacyGetJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking LegacyGetJob.
type LegacyGetJobArgs struct {
	JobId     int `pulumi:"jobId"`
	ProjectId int `pulumi:"projectId"`
}

// A collection of values returned by LegacyGetJob.
type LegacyGetJobResult struct {
	DeferringEnvironmentId int    `pulumi:"deferringEnvironmentId"`
	DeferringJobId         int    `pulumi:"deferringJobId"`
	Description            string `pulumi:"description"`
	EnvironmentId          int    `pulumi:"environmentId"`
	// The provider-assigned unique ID for this managed resource.
	Id                             string                                      `pulumi:"id"`
	JobCompletionTriggerConditions []LegacyGetJobJobCompletionTriggerCondition `pulumi:"jobCompletionTriggerConditions"`
	JobId                          int                                         `pulumi:"jobId"`
	Name                           string                                      `pulumi:"name"`
	ProjectId                      int                                         `pulumi:"projectId"`
	SelfDeferring                  bool                                        `pulumi:"selfDeferring"`
	TimeoutSeconds                 int                                         `pulumi:"timeoutSeconds"`
	Triggers                       map[string]bool                             `pulumi:"triggers"`
	TriggersOnDraftPr              bool                                        `pulumi:"triggersOnDraftPr"`
}

func LegacyGetJobOutput(ctx *pulumi.Context, args LegacyGetJobOutputArgs, opts ...pulumi.InvokeOption) LegacyGetJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LegacyGetJobResult, error) {
			args := v.(LegacyGetJobArgs)
			r, err := LegacyGetJob(ctx, &args, opts...)
			var s LegacyGetJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LegacyGetJobResultOutput)
}

// A collection of arguments for invoking LegacyGetJob.
type LegacyGetJobOutputArgs struct {
	JobId     pulumi.IntInput `pulumi:"jobId"`
	ProjectId pulumi.IntInput `pulumi:"projectId"`
}

func (LegacyGetJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LegacyGetJobArgs)(nil)).Elem()
}

// A collection of values returned by LegacyGetJob.
type LegacyGetJobResultOutput struct{ *pulumi.OutputState }

func (LegacyGetJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LegacyGetJobResult)(nil)).Elem()
}

func (o LegacyGetJobResultOutput) ToLegacyGetJobResultOutput() LegacyGetJobResultOutput {
	return o
}

func (o LegacyGetJobResultOutput) ToLegacyGetJobResultOutputWithContext(ctx context.Context) LegacyGetJobResultOutput {
	return o
}

func (o LegacyGetJobResultOutput) DeferringEnvironmentId() pulumi.IntOutput {
	return o.ApplyT(func(v LegacyGetJobResult) int { return v.DeferringEnvironmentId }).(pulumi.IntOutput)
}

func (o LegacyGetJobResultOutput) DeferringJobId() pulumi.IntOutput {
	return o.ApplyT(func(v LegacyGetJobResult) int { return v.DeferringJobId }).(pulumi.IntOutput)
}

func (o LegacyGetJobResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LegacyGetJobResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LegacyGetJobResultOutput) EnvironmentId() pulumi.IntOutput {
	return o.ApplyT(func(v LegacyGetJobResult) int { return v.EnvironmentId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LegacyGetJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LegacyGetJobResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LegacyGetJobResultOutput) JobCompletionTriggerConditions() LegacyGetJobJobCompletionTriggerConditionArrayOutput {
	return o.ApplyT(func(v LegacyGetJobResult) []LegacyGetJobJobCompletionTriggerCondition {
		return v.JobCompletionTriggerConditions
	}).(LegacyGetJobJobCompletionTriggerConditionArrayOutput)
}

func (o LegacyGetJobResultOutput) JobId() pulumi.IntOutput {
	return o.ApplyT(func(v LegacyGetJobResult) int { return v.JobId }).(pulumi.IntOutput)
}

func (o LegacyGetJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LegacyGetJobResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LegacyGetJobResultOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v LegacyGetJobResult) int { return v.ProjectId }).(pulumi.IntOutput)
}

func (o LegacyGetJobResultOutput) SelfDeferring() pulumi.BoolOutput {
	return o.ApplyT(func(v LegacyGetJobResult) bool { return v.SelfDeferring }).(pulumi.BoolOutput)
}

func (o LegacyGetJobResultOutput) TimeoutSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v LegacyGetJobResult) int { return v.TimeoutSeconds }).(pulumi.IntOutput)
}

func (o LegacyGetJobResultOutput) Triggers() pulumi.BoolMapOutput {
	return o.ApplyT(func(v LegacyGetJobResult) map[string]bool { return v.Triggers }).(pulumi.BoolMapOutput)
}

func (o LegacyGetJobResultOutput) TriggersOnDraftPr() pulumi.BoolOutput {
	return o.ApplyT(func(v LegacyGetJobResult) bool { return v.TriggersOnDraftPr }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LegacyGetJobResultOutput{})
}

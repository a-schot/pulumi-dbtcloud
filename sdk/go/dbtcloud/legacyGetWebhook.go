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
func LegacyGetWebhook(ctx *pulumi.Context, args *LegacyGetWebhookArgs, opts ...pulumi.InvokeOption) (*LegacyGetWebhookResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LegacyGetWebhookResult
	err := ctx.Invoke("dbtcloud:index/legacyGetWebhook:LegacyGetWebhook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking LegacyGetWebhook.
type LegacyGetWebhookArgs struct {
	WebhookId string `pulumi:"webhookId"`
}

// A collection of values returned by LegacyGetWebhook.
type LegacyGetWebhookResult struct {
	AccountIdentifier string   `pulumi:"accountIdentifier"`
	Active            bool     `pulumi:"active"`
	ClientUrl         string   `pulumi:"clientUrl"`
	Description       string   `pulumi:"description"`
	EventTypes        []string `pulumi:"eventTypes"`
	HttpStatusCode    string   `pulumi:"httpStatusCode"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	JobIds    []int  `pulumi:"jobIds"`
	Name      string `pulumi:"name"`
	WebhookId string `pulumi:"webhookId"`
}

func LegacyGetWebhookOutput(ctx *pulumi.Context, args LegacyGetWebhookOutputArgs, opts ...pulumi.InvokeOption) LegacyGetWebhookResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LegacyGetWebhookResult, error) {
			args := v.(LegacyGetWebhookArgs)
			r, err := LegacyGetWebhook(ctx, &args, opts...)
			var s LegacyGetWebhookResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LegacyGetWebhookResultOutput)
}

// A collection of arguments for invoking LegacyGetWebhook.
type LegacyGetWebhookOutputArgs struct {
	WebhookId pulumi.StringInput `pulumi:"webhookId"`
}

func (LegacyGetWebhookOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LegacyGetWebhookArgs)(nil)).Elem()
}

// A collection of values returned by LegacyGetWebhook.
type LegacyGetWebhookResultOutput struct{ *pulumi.OutputState }

func (LegacyGetWebhookResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LegacyGetWebhookResult)(nil)).Elem()
}

func (o LegacyGetWebhookResultOutput) ToLegacyGetWebhookResultOutput() LegacyGetWebhookResultOutput {
	return o
}

func (o LegacyGetWebhookResultOutput) ToLegacyGetWebhookResultOutputWithContext(ctx context.Context) LegacyGetWebhookResultOutput {
	return o
}

func (o LegacyGetWebhookResultOutput) AccountIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LegacyGetWebhookResult) string { return v.AccountIdentifier }).(pulumi.StringOutput)
}

func (o LegacyGetWebhookResultOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v LegacyGetWebhookResult) bool { return v.Active }).(pulumi.BoolOutput)
}

func (o LegacyGetWebhookResultOutput) ClientUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LegacyGetWebhookResult) string { return v.ClientUrl }).(pulumi.StringOutput)
}

func (o LegacyGetWebhookResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LegacyGetWebhookResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LegacyGetWebhookResultOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LegacyGetWebhookResult) []string { return v.EventTypes }).(pulumi.StringArrayOutput)
}

func (o LegacyGetWebhookResultOutput) HttpStatusCode() pulumi.StringOutput {
	return o.ApplyT(func(v LegacyGetWebhookResult) string { return v.HttpStatusCode }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LegacyGetWebhookResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LegacyGetWebhookResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LegacyGetWebhookResultOutput) JobIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LegacyGetWebhookResult) []int { return v.JobIds }).(pulumi.IntArrayOutput)
}

func (o LegacyGetWebhookResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LegacyGetWebhookResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LegacyGetWebhookResultOutput) WebhookId() pulumi.StringOutput {
	return o.ApplyT(func(v LegacyGetWebhookResult) string { return v.WebhookId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LegacyGetWebhookResultOutput{})
}

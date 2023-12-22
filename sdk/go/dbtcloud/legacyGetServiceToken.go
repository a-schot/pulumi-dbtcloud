// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbtcloud

import (
	"context"
	"reflect"

	"github.com/a-schot/pulumi-dbtcloud/sdk/go/dbtcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LegacyGetServiceToken(ctx *pulumi.Context, args *LegacyGetServiceTokenArgs, opts ...pulumi.InvokeOption) (*LegacyGetServiceTokenResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LegacyGetServiceTokenResult
	err := ctx.Invoke("dbtcloud:index/legacyGetServiceToken:LegacyGetServiceToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking LegacyGetServiceToken.
type LegacyGetServiceTokenArgs struct {
	ServiceTokenId int `pulumi:"serviceTokenId"`
}

// A collection of values returned by LegacyGetServiceToken.
type LegacyGetServiceTokenResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                      string                                        `pulumi:"id"`
	Name                    string                                        `pulumi:"name"`
	ServiceTokenId          int                                           `pulumi:"serviceTokenId"`
	ServiceTokenPermissions []LegacyGetServiceTokenServiceTokenPermission `pulumi:"serviceTokenPermissions"`
	Uid                     string                                        `pulumi:"uid"`
}

func LegacyGetServiceTokenOutput(ctx *pulumi.Context, args LegacyGetServiceTokenOutputArgs, opts ...pulumi.InvokeOption) LegacyGetServiceTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LegacyGetServiceTokenResult, error) {
			args := v.(LegacyGetServiceTokenArgs)
			r, err := LegacyGetServiceToken(ctx, &args, opts...)
			var s LegacyGetServiceTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LegacyGetServiceTokenResultOutput)
}

// A collection of arguments for invoking LegacyGetServiceToken.
type LegacyGetServiceTokenOutputArgs struct {
	ServiceTokenId pulumi.IntInput `pulumi:"serviceTokenId"`
}

func (LegacyGetServiceTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LegacyGetServiceTokenArgs)(nil)).Elem()
}

// A collection of values returned by LegacyGetServiceToken.
type LegacyGetServiceTokenResultOutput struct{ *pulumi.OutputState }

func (LegacyGetServiceTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LegacyGetServiceTokenResult)(nil)).Elem()
}

func (o LegacyGetServiceTokenResultOutput) ToLegacyGetServiceTokenResultOutput() LegacyGetServiceTokenResultOutput {
	return o
}

func (o LegacyGetServiceTokenResultOutput) ToLegacyGetServiceTokenResultOutputWithContext(ctx context.Context) LegacyGetServiceTokenResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LegacyGetServiceTokenResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LegacyGetServiceTokenResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LegacyGetServiceTokenResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LegacyGetServiceTokenResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LegacyGetServiceTokenResultOutput) ServiceTokenId() pulumi.IntOutput {
	return o.ApplyT(func(v LegacyGetServiceTokenResult) int { return v.ServiceTokenId }).(pulumi.IntOutput)
}

func (o LegacyGetServiceTokenResultOutput) ServiceTokenPermissions() LegacyGetServiceTokenServiceTokenPermissionArrayOutput {
	return o.ApplyT(func(v LegacyGetServiceTokenResult) []LegacyGetServiceTokenServiceTokenPermission {
		return v.ServiceTokenPermissions
	}).(LegacyGetServiceTokenServiceTokenPermissionArrayOutput)
}

func (o LegacyGetServiceTokenResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LegacyGetServiceTokenResult) string { return v.Uid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LegacyGetServiceTokenResultOutput{})
}
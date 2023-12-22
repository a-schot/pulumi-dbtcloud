// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbtcloud

import (
	"context"
	"reflect"

	"github.com/a-schot/pulumi-dbtcloud/sdk/go/dbtcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/a-schot/pulumi-dbtcloud/sdk/go/dbtcloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dbtcloud.LookupUserGroups(ctx, &dbtcloud.LookupUserGroupsArgs{
//				UserId: 12345,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupUserGroups(ctx *pulumi.Context, args *LookupUserGroupsArgs, opts ...pulumi.InvokeOption) (*LookupUserGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserGroupsResult
	err := ctx.Invoke("dbtcloud:index/getUserGroups:getUserGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUserGroups.
type LookupUserGroupsArgs struct {
	// ID of the user
	UserId int `pulumi:"userId"`
}

// A collection of values returned by getUserGroups.
type LookupUserGroupsResult struct {
	// IDs of the groups assigned to the user
	GroupIds []int `pulumi:"groupIds"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the user
	UserId int `pulumi:"userId"`
}

func LookupUserGroupsOutput(ctx *pulumi.Context, args LookupUserGroupsOutputArgs, opts ...pulumi.InvokeOption) LookupUserGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserGroupsResult, error) {
			args := v.(LookupUserGroupsArgs)
			r, err := LookupUserGroups(ctx, &args, opts...)
			var s LookupUserGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserGroupsResultOutput)
}

// A collection of arguments for invoking getUserGroups.
type LookupUserGroupsOutputArgs struct {
	// ID of the user
	UserId pulumi.IntInput `pulumi:"userId"`
}

func (LookupUserGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getUserGroups.
type LookupUserGroupsResultOutput struct{ *pulumi.OutputState }

func (LookupUserGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserGroupsResult)(nil)).Elem()
}

func (o LookupUserGroupsResultOutput) ToLookupUserGroupsResultOutput() LookupUserGroupsResultOutput {
	return o
}

func (o LookupUserGroupsResultOutput) ToLookupUserGroupsResultOutputWithContext(ctx context.Context) LookupUserGroupsResultOutput {
	return o
}

// IDs of the groups assigned to the user
func (o LookupUserGroupsResultOutput) GroupIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupUserGroupsResult) []int { return v.GroupIds }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupUserGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the user
func (o LookupUserGroupsResultOutput) UserId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupUserGroupsResult) int { return v.UserId }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserGroupsResultOutput{})
}

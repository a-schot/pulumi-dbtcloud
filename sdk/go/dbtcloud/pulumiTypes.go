// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbtcloud

import (
	"context"
	"reflect"

	"github.com/a-schot/pulumi-dbtcloud/sdk/go/dbtcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type GroupGroupPermission struct {
	// Whether or not to apply this permission to all projects for this group
	AllProjects bool `pulumi:"allProjects"`
	// Set of permissions to apply
	PermissionSet string `pulumi:"permissionSet"`
	// Project ID to apply this permission to for this group
	ProjectId *int `pulumi:"projectId"`
}

// GroupGroupPermissionInput is an input type that accepts GroupGroupPermissionArgs and GroupGroupPermissionOutput values.
// You can construct a concrete instance of `GroupGroupPermissionInput` via:
//
//	GroupGroupPermissionArgs{...}
type GroupGroupPermissionInput interface {
	pulumi.Input

	ToGroupGroupPermissionOutput() GroupGroupPermissionOutput
	ToGroupGroupPermissionOutputWithContext(context.Context) GroupGroupPermissionOutput
}

type GroupGroupPermissionArgs struct {
	// Whether or not to apply this permission to all projects for this group
	AllProjects pulumi.BoolInput `pulumi:"allProjects"`
	// Set of permissions to apply
	PermissionSet pulumi.StringInput `pulumi:"permissionSet"`
	// Project ID to apply this permission to for this group
	ProjectId pulumi.IntPtrInput `pulumi:"projectId"`
}

func (GroupGroupPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupGroupPermission)(nil)).Elem()
}

func (i GroupGroupPermissionArgs) ToGroupGroupPermissionOutput() GroupGroupPermissionOutput {
	return i.ToGroupGroupPermissionOutputWithContext(context.Background())
}

func (i GroupGroupPermissionArgs) ToGroupGroupPermissionOutputWithContext(ctx context.Context) GroupGroupPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupGroupPermissionOutput)
}

// GroupGroupPermissionArrayInput is an input type that accepts GroupGroupPermissionArray and GroupGroupPermissionArrayOutput values.
// You can construct a concrete instance of `GroupGroupPermissionArrayInput` via:
//
//	GroupGroupPermissionArray{ GroupGroupPermissionArgs{...} }
type GroupGroupPermissionArrayInput interface {
	pulumi.Input

	ToGroupGroupPermissionArrayOutput() GroupGroupPermissionArrayOutput
	ToGroupGroupPermissionArrayOutputWithContext(context.Context) GroupGroupPermissionArrayOutput
}

type GroupGroupPermissionArray []GroupGroupPermissionInput

func (GroupGroupPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupGroupPermission)(nil)).Elem()
}

func (i GroupGroupPermissionArray) ToGroupGroupPermissionArrayOutput() GroupGroupPermissionArrayOutput {
	return i.ToGroupGroupPermissionArrayOutputWithContext(context.Background())
}

func (i GroupGroupPermissionArray) ToGroupGroupPermissionArrayOutputWithContext(ctx context.Context) GroupGroupPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupGroupPermissionArrayOutput)
}

type GroupGroupPermissionOutput struct{ *pulumi.OutputState }

func (GroupGroupPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupGroupPermission)(nil)).Elem()
}

func (o GroupGroupPermissionOutput) ToGroupGroupPermissionOutput() GroupGroupPermissionOutput {
	return o
}

func (o GroupGroupPermissionOutput) ToGroupGroupPermissionOutputWithContext(ctx context.Context) GroupGroupPermissionOutput {
	return o
}

// Whether or not to apply this permission to all projects for this group
func (o GroupGroupPermissionOutput) AllProjects() pulumi.BoolOutput {
	return o.ApplyT(func(v GroupGroupPermission) bool { return v.AllProjects }).(pulumi.BoolOutput)
}

// Set of permissions to apply
func (o GroupGroupPermissionOutput) PermissionSet() pulumi.StringOutput {
	return o.ApplyT(func(v GroupGroupPermission) string { return v.PermissionSet }).(pulumi.StringOutput)
}

// Project ID to apply this permission to for this group
func (o GroupGroupPermissionOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GroupGroupPermission) *int { return v.ProjectId }).(pulumi.IntPtrOutput)
}

type GroupGroupPermissionArrayOutput struct{ *pulumi.OutputState }

func (GroupGroupPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupGroupPermission)(nil)).Elem()
}

func (o GroupGroupPermissionArrayOutput) ToGroupGroupPermissionArrayOutput() GroupGroupPermissionArrayOutput {
	return o
}

func (o GroupGroupPermissionArrayOutput) ToGroupGroupPermissionArrayOutputWithContext(ctx context.Context) GroupGroupPermissionArrayOutput {
	return o
}

func (o GroupGroupPermissionArrayOutput) Index(i pulumi.IntInput) GroupGroupPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupGroupPermission {
		return vs[0].([]GroupGroupPermission)[vs[1].(int)]
	}).(GroupGroupPermissionOutput)
}

type ServiceTokenServiceTokenPermission struct {
	// Whether or not to apply this permission to all projects for this service token
	AllProjects bool `pulumi:"allProjects"`
	// Set of permissions to apply
	PermissionSet string `pulumi:"permissionSet"`
	// Project ID to apply this permission to for this service token
	ProjectId *int `pulumi:"projectId"`
}

// ServiceTokenServiceTokenPermissionInput is an input type that accepts ServiceTokenServiceTokenPermissionArgs and ServiceTokenServiceTokenPermissionOutput values.
// You can construct a concrete instance of `ServiceTokenServiceTokenPermissionInput` via:
//
//	ServiceTokenServiceTokenPermissionArgs{...}
type ServiceTokenServiceTokenPermissionInput interface {
	pulumi.Input

	ToServiceTokenServiceTokenPermissionOutput() ServiceTokenServiceTokenPermissionOutput
	ToServiceTokenServiceTokenPermissionOutputWithContext(context.Context) ServiceTokenServiceTokenPermissionOutput
}

type ServiceTokenServiceTokenPermissionArgs struct {
	// Whether or not to apply this permission to all projects for this service token
	AllProjects pulumi.BoolInput `pulumi:"allProjects"`
	// Set of permissions to apply
	PermissionSet pulumi.StringInput `pulumi:"permissionSet"`
	// Project ID to apply this permission to for this service token
	ProjectId pulumi.IntPtrInput `pulumi:"projectId"`
}

func (ServiceTokenServiceTokenPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTokenServiceTokenPermission)(nil)).Elem()
}

func (i ServiceTokenServiceTokenPermissionArgs) ToServiceTokenServiceTokenPermissionOutput() ServiceTokenServiceTokenPermissionOutput {
	return i.ToServiceTokenServiceTokenPermissionOutputWithContext(context.Background())
}

func (i ServiceTokenServiceTokenPermissionArgs) ToServiceTokenServiceTokenPermissionOutputWithContext(ctx context.Context) ServiceTokenServiceTokenPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTokenServiceTokenPermissionOutput)
}

// ServiceTokenServiceTokenPermissionArrayInput is an input type that accepts ServiceTokenServiceTokenPermissionArray and ServiceTokenServiceTokenPermissionArrayOutput values.
// You can construct a concrete instance of `ServiceTokenServiceTokenPermissionArrayInput` via:
//
//	ServiceTokenServiceTokenPermissionArray{ ServiceTokenServiceTokenPermissionArgs{...} }
type ServiceTokenServiceTokenPermissionArrayInput interface {
	pulumi.Input

	ToServiceTokenServiceTokenPermissionArrayOutput() ServiceTokenServiceTokenPermissionArrayOutput
	ToServiceTokenServiceTokenPermissionArrayOutputWithContext(context.Context) ServiceTokenServiceTokenPermissionArrayOutput
}

type ServiceTokenServiceTokenPermissionArray []ServiceTokenServiceTokenPermissionInput

func (ServiceTokenServiceTokenPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceTokenServiceTokenPermission)(nil)).Elem()
}

func (i ServiceTokenServiceTokenPermissionArray) ToServiceTokenServiceTokenPermissionArrayOutput() ServiceTokenServiceTokenPermissionArrayOutput {
	return i.ToServiceTokenServiceTokenPermissionArrayOutputWithContext(context.Background())
}

func (i ServiceTokenServiceTokenPermissionArray) ToServiceTokenServiceTokenPermissionArrayOutputWithContext(ctx context.Context) ServiceTokenServiceTokenPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTokenServiceTokenPermissionArrayOutput)
}

type ServiceTokenServiceTokenPermissionOutput struct{ *pulumi.OutputState }

func (ServiceTokenServiceTokenPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTokenServiceTokenPermission)(nil)).Elem()
}

func (o ServiceTokenServiceTokenPermissionOutput) ToServiceTokenServiceTokenPermissionOutput() ServiceTokenServiceTokenPermissionOutput {
	return o
}

func (o ServiceTokenServiceTokenPermissionOutput) ToServiceTokenServiceTokenPermissionOutputWithContext(ctx context.Context) ServiceTokenServiceTokenPermissionOutput {
	return o
}

// Whether or not to apply this permission to all projects for this service token
func (o ServiceTokenServiceTokenPermissionOutput) AllProjects() pulumi.BoolOutput {
	return o.ApplyT(func(v ServiceTokenServiceTokenPermission) bool { return v.AllProjects }).(pulumi.BoolOutput)
}

// Set of permissions to apply
func (o ServiceTokenServiceTokenPermissionOutput) PermissionSet() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceTokenServiceTokenPermission) string { return v.PermissionSet }).(pulumi.StringOutput)
}

// Project ID to apply this permission to for this service token
func (o ServiceTokenServiceTokenPermissionOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceTokenServiceTokenPermission) *int { return v.ProjectId }).(pulumi.IntPtrOutput)
}

type ServiceTokenServiceTokenPermissionArrayOutput struct{ *pulumi.OutputState }

func (ServiceTokenServiceTokenPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceTokenServiceTokenPermission)(nil)).Elem()
}

func (o ServiceTokenServiceTokenPermissionArrayOutput) ToServiceTokenServiceTokenPermissionArrayOutput() ServiceTokenServiceTokenPermissionArrayOutput {
	return o
}

func (o ServiceTokenServiceTokenPermissionArrayOutput) ToServiceTokenServiceTokenPermissionArrayOutputWithContext(ctx context.Context) ServiceTokenServiceTokenPermissionArrayOutput {
	return o
}

func (o ServiceTokenServiceTokenPermissionArrayOutput) Index(i pulumi.IntInput) ServiceTokenServiceTokenPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceTokenServiceTokenPermission {
		return vs[0].([]ServiceTokenServiceTokenPermission)[vs[1].(int)]
	}).(ServiceTokenServiceTokenPermissionOutput)
}

type GetGroupUsersUser struct {
	Email string `pulumi:"email"`
	Id    int    `pulumi:"id"`
}

// GetGroupUsersUserInput is an input type that accepts GetGroupUsersUserArgs and GetGroupUsersUserOutput values.
// You can construct a concrete instance of `GetGroupUsersUserInput` via:
//
//	GetGroupUsersUserArgs{...}
type GetGroupUsersUserInput interface {
	pulumi.Input

	ToGetGroupUsersUserOutput() GetGroupUsersUserOutput
	ToGetGroupUsersUserOutputWithContext(context.Context) GetGroupUsersUserOutput
}

type GetGroupUsersUserArgs struct {
	Email pulumi.StringInput `pulumi:"email"`
	Id    pulumi.IntInput    `pulumi:"id"`
}

func (GetGroupUsersUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupUsersUser)(nil)).Elem()
}

func (i GetGroupUsersUserArgs) ToGetGroupUsersUserOutput() GetGroupUsersUserOutput {
	return i.ToGetGroupUsersUserOutputWithContext(context.Background())
}

func (i GetGroupUsersUserArgs) ToGetGroupUsersUserOutputWithContext(ctx context.Context) GetGroupUsersUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGroupUsersUserOutput)
}

// GetGroupUsersUserArrayInput is an input type that accepts GetGroupUsersUserArray and GetGroupUsersUserArrayOutput values.
// You can construct a concrete instance of `GetGroupUsersUserArrayInput` via:
//
//	GetGroupUsersUserArray{ GetGroupUsersUserArgs{...} }
type GetGroupUsersUserArrayInput interface {
	pulumi.Input

	ToGetGroupUsersUserArrayOutput() GetGroupUsersUserArrayOutput
	ToGetGroupUsersUserArrayOutputWithContext(context.Context) GetGroupUsersUserArrayOutput
}

type GetGroupUsersUserArray []GetGroupUsersUserInput

func (GetGroupUsersUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGroupUsersUser)(nil)).Elem()
}

func (i GetGroupUsersUserArray) ToGetGroupUsersUserArrayOutput() GetGroupUsersUserArrayOutput {
	return i.ToGetGroupUsersUserArrayOutputWithContext(context.Background())
}

func (i GetGroupUsersUserArray) ToGetGroupUsersUserArrayOutputWithContext(ctx context.Context) GetGroupUsersUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGroupUsersUserArrayOutput)
}

type GetGroupUsersUserOutput struct{ *pulumi.OutputState }

func (GetGroupUsersUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupUsersUser)(nil)).Elem()
}

func (o GetGroupUsersUserOutput) ToGetGroupUsersUserOutput() GetGroupUsersUserOutput {
	return o
}

func (o GetGroupUsersUserOutput) ToGetGroupUsersUserOutputWithContext(ctx context.Context) GetGroupUsersUserOutput {
	return o
}

func (o GetGroupUsersUserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupUsersUser) string { return v.Email }).(pulumi.StringOutput)
}

func (o GetGroupUsersUserOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetGroupUsersUser) int { return v.Id }).(pulumi.IntOutput)
}

type GetGroupUsersUserArrayOutput struct{ *pulumi.OutputState }

func (GetGroupUsersUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGroupUsersUser)(nil)).Elem()
}

func (o GetGroupUsersUserArrayOutput) ToGetGroupUsersUserArrayOutput() GetGroupUsersUserArrayOutput {
	return o
}

func (o GetGroupUsersUserArrayOutput) ToGetGroupUsersUserArrayOutputWithContext(ctx context.Context) GetGroupUsersUserArrayOutput {
	return o
}

func (o GetGroupUsersUserArrayOutput) Index(i pulumi.IntInput) GetGroupUsersUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetGroupUsersUser {
		return vs[0].([]GetGroupUsersUser)[vs[1].(int)]
	}).(GetGroupUsersUserOutput)
}

type GetServiceTokenServiceTokenPermission struct {
	AllProjects   bool   `pulumi:"allProjects"`
	PermissionSet string `pulumi:"permissionSet"`
	ProjectId     int    `pulumi:"projectId"`
}

// GetServiceTokenServiceTokenPermissionInput is an input type that accepts GetServiceTokenServiceTokenPermissionArgs and GetServiceTokenServiceTokenPermissionOutput values.
// You can construct a concrete instance of `GetServiceTokenServiceTokenPermissionInput` via:
//
//	GetServiceTokenServiceTokenPermissionArgs{...}
type GetServiceTokenServiceTokenPermissionInput interface {
	pulumi.Input

	ToGetServiceTokenServiceTokenPermissionOutput() GetServiceTokenServiceTokenPermissionOutput
	ToGetServiceTokenServiceTokenPermissionOutputWithContext(context.Context) GetServiceTokenServiceTokenPermissionOutput
}

type GetServiceTokenServiceTokenPermissionArgs struct {
	AllProjects   pulumi.BoolInput   `pulumi:"allProjects"`
	PermissionSet pulumi.StringInput `pulumi:"permissionSet"`
	ProjectId     pulumi.IntInput    `pulumi:"projectId"`
}

func (GetServiceTokenServiceTokenPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceTokenServiceTokenPermission)(nil)).Elem()
}

func (i GetServiceTokenServiceTokenPermissionArgs) ToGetServiceTokenServiceTokenPermissionOutput() GetServiceTokenServiceTokenPermissionOutput {
	return i.ToGetServiceTokenServiceTokenPermissionOutputWithContext(context.Background())
}

func (i GetServiceTokenServiceTokenPermissionArgs) ToGetServiceTokenServiceTokenPermissionOutputWithContext(ctx context.Context) GetServiceTokenServiceTokenPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetServiceTokenServiceTokenPermissionOutput)
}

// GetServiceTokenServiceTokenPermissionArrayInput is an input type that accepts GetServiceTokenServiceTokenPermissionArray and GetServiceTokenServiceTokenPermissionArrayOutput values.
// You can construct a concrete instance of `GetServiceTokenServiceTokenPermissionArrayInput` via:
//
//	GetServiceTokenServiceTokenPermissionArray{ GetServiceTokenServiceTokenPermissionArgs{...} }
type GetServiceTokenServiceTokenPermissionArrayInput interface {
	pulumi.Input

	ToGetServiceTokenServiceTokenPermissionArrayOutput() GetServiceTokenServiceTokenPermissionArrayOutput
	ToGetServiceTokenServiceTokenPermissionArrayOutputWithContext(context.Context) GetServiceTokenServiceTokenPermissionArrayOutput
}

type GetServiceTokenServiceTokenPermissionArray []GetServiceTokenServiceTokenPermissionInput

func (GetServiceTokenServiceTokenPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetServiceTokenServiceTokenPermission)(nil)).Elem()
}

func (i GetServiceTokenServiceTokenPermissionArray) ToGetServiceTokenServiceTokenPermissionArrayOutput() GetServiceTokenServiceTokenPermissionArrayOutput {
	return i.ToGetServiceTokenServiceTokenPermissionArrayOutputWithContext(context.Background())
}

func (i GetServiceTokenServiceTokenPermissionArray) ToGetServiceTokenServiceTokenPermissionArrayOutputWithContext(ctx context.Context) GetServiceTokenServiceTokenPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetServiceTokenServiceTokenPermissionArrayOutput)
}

type GetServiceTokenServiceTokenPermissionOutput struct{ *pulumi.OutputState }

func (GetServiceTokenServiceTokenPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceTokenServiceTokenPermission)(nil)).Elem()
}

func (o GetServiceTokenServiceTokenPermissionOutput) ToGetServiceTokenServiceTokenPermissionOutput() GetServiceTokenServiceTokenPermissionOutput {
	return o
}

func (o GetServiceTokenServiceTokenPermissionOutput) ToGetServiceTokenServiceTokenPermissionOutputWithContext(ctx context.Context) GetServiceTokenServiceTokenPermissionOutput {
	return o
}

func (o GetServiceTokenServiceTokenPermissionOutput) AllProjects() pulumi.BoolOutput {
	return o.ApplyT(func(v GetServiceTokenServiceTokenPermission) bool { return v.AllProjects }).(pulumi.BoolOutput)
}

func (o GetServiceTokenServiceTokenPermissionOutput) PermissionSet() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceTokenServiceTokenPermission) string { return v.PermissionSet }).(pulumi.StringOutput)
}

func (o GetServiceTokenServiceTokenPermissionOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v GetServiceTokenServiceTokenPermission) int { return v.ProjectId }).(pulumi.IntOutput)
}

type GetServiceTokenServiceTokenPermissionArrayOutput struct{ *pulumi.OutputState }

func (GetServiceTokenServiceTokenPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetServiceTokenServiceTokenPermission)(nil)).Elem()
}

func (o GetServiceTokenServiceTokenPermissionArrayOutput) ToGetServiceTokenServiceTokenPermissionArrayOutput() GetServiceTokenServiceTokenPermissionArrayOutput {
	return o
}

func (o GetServiceTokenServiceTokenPermissionArrayOutput) ToGetServiceTokenServiceTokenPermissionArrayOutputWithContext(ctx context.Context) GetServiceTokenServiceTokenPermissionArrayOutput {
	return o
}

func (o GetServiceTokenServiceTokenPermissionArrayOutput) Index(i pulumi.IntInput) GetServiceTokenServiceTokenPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetServiceTokenServiceTokenPermission {
		return vs[0].([]GetServiceTokenServiceTokenPermission)[vs[1].(int)]
	}).(GetServiceTokenServiceTokenPermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupGroupPermissionInput)(nil)).Elem(), GroupGroupPermissionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupGroupPermissionArrayInput)(nil)).Elem(), GroupGroupPermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTokenServiceTokenPermissionInput)(nil)).Elem(), ServiceTokenServiceTokenPermissionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTokenServiceTokenPermissionArrayInput)(nil)).Elem(), ServiceTokenServiceTokenPermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetGroupUsersUserInput)(nil)).Elem(), GetGroupUsersUserArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetGroupUsersUserArrayInput)(nil)).Elem(), GetGroupUsersUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetServiceTokenServiceTokenPermissionInput)(nil)).Elem(), GetServiceTokenServiceTokenPermissionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetServiceTokenServiceTokenPermissionArrayInput)(nil)).Elem(), GetServiceTokenServiceTokenPermissionArray{})
	pulumi.RegisterOutputType(GroupGroupPermissionOutput{})
	pulumi.RegisterOutputType(GroupGroupPermissionArrayOutput{})
	pulumi.RegisterOutputType(ServiceTokenServiceTokenPermissionOutput{})
	pulumi.RegisterOutputType(ServiceTokenServiceTokenPermissionArrayOutput{})
	pulumi.RegisterOutputType(GetGroupUsersUserOutput{})
	pulumi.RegisterOutputType(GetGroupUsersUserArrayOutput{})
	pulumi.RegisterOutputType(GetServiceTokenServiceTokenPermissionOutput{})
	pulumi.RegisterOutputType(GetServiceTokenServiceTokenPermissionArrayOutput{})
}

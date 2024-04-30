// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbtcloud

import (
	"context"
	"reflect"

	"github.com/a-schot/pulumi-dbtcloud/sdk/go/dbtcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// *Note*: Groups currently do not support updates, as per both the API and the UI.
//
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
//			// NOTE for customers using the LEGACY dbt_cloud provider:
//			_, err := dbtcloud.NewGroup(ctx, "tfGroup1", &dbtcloud.GroupArgs{
//				GroupPermissions: dbtcloud.GroupGroupPermissionArray{
//					&dbtcloud.GroupGroupPermissionArgs{
//						PermissionSet: pulumi.String("member"),
//						AllProjects:   pulumi.Bool(true),
//					},
//					&dbtcloud.GroupGroupPermissionArgs{
//						PermissionSet: pulumi.String("developer"),
//						AllProjects:   pulumi.Bool(false),
//						ProjectId:     pulumi.Any(dbtcloud_project.Dbt_project.Id),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Import using a group ID found in the URL or via the API.
//
// ```sh
// $ pulumi import dbtcloud:index/group:Group test_group "group_id"
// ```
//
// ```sh
// $ pulumi import dbtcloud:index/group:Group test_group 12345
// ```
type Group struct {
	pulumi.CustomResourceState

	// Whether or not to assign this group to users by default
	AssignByDefault  pulumi.BoolPtrOutput            `pulumi:"assignByDefault"`
	GroupPermissions GroupGroupPermissionArrayOutput `pulumi:"groupPermissions"`
	// Whether the group is active
	IsActive pulumi.BoolPtrOutput `pulumi:"isActive"`
	// Group name
	Name pulumi.StringOutput `pulumi:"name"`
	// SSO mapping group names for this group
	SsoMappingGroups pulumi.StringArrayOutput `pulumi:"ssoMappingGroups"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		args = &GroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Group
	err := ctx.RegisterResource("dbtcloud:index/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("dbtcloud:index/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// Whether or not to assign this group to users by default
	AssignByDefault  *bool                  `pulumi:"assignByDefault"`
	GroupPermissions []GroupGroupPermission `pulumi:"groupPermissions"`
	// Whether the group is active
	IsActive *bool `pulumi:"isActive"`
	// Group name
	Name *string `pulumi:"name"`
	// SSO mapping group names for this group
	SsoMappingGroups []string `pulumi:"ssoMappingGroups"`
}

type GroupState struct {
	// Whether or not to assign this group to users by default
	AssignByDefault  pulumi.BoolPtrInput
	GroupPermissions GroupGroupPermissionArrayInput
	// Whether the group is active
	IsActive pulumi.BoolPtrInput
	// Group name
	Name pulumi.StringPtrInput
	// SSO mapping group names for this group
	SsoMappingGroups pulumi.StringArrayInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// Whether or not to assign this group to users by default
	AssignByDefault  *bool                  `pulumi:"assignByDefault"`
	GroupPermissions []GroupGroupPermission `pulumi:"groupPermissions"`
	// Whether the group is active
	IsActive *bool `pulumi:"isActive"`
	// Group name
	Name *string `pulumi:"name"`
	// SSO mapping group names for this group
	SsoMappingGroups []string `pulumi:"ssoMappingGroups"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// Whether or not to assign this group to users by default
	AssignByDefault  pulumi.BoolPtrInput
	GroupPermissions GroupGroupPermissionArrayInput
	// Whether the group is active
	IsActive pulumi.BoolPtrInput
	// Group name
	Name pulumi.StringPtrInput
	// SSO mapping group names for this group
	SsoMappingGroups pulumi.StringArrayInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//	GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//	GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

// Whether or not to assign this group to users by default
func (o GroupOutput) AssignByDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolPtrOutput { return v.AssignByDefault }).(pulumi.BoolPtrOutput)
}

func (o GroupOutput) GroupPermissions() GroupGroupPermissionArrayOutput {
	return o.ApplyT(func(v *Group) GroupGroupPermissionArrayOutput { return v.GroupPermissions }).(GroupGroupPermissionArrayOutput)
}

// Whether the group is active
func (o GroupOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.BoolPtrOutput { return v.IsActive }).(pulumi.BoolPtrOutput)
}

// Group name
func (o GroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// SSO mapping group names for this group
func (o GroupOutput) SsoMappingGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Group) pulumi.StringArrayOutput { return v.SsoMappingGroups }).(pulumi.StringArrayOutput)
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Group {
		return vs[0].([]*Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Group {
		return vs[0].(map[string]*Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupArrayInput)(nil)).Elem(), GroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMapInput)(nil)).Elem(), GroupMap{})
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}

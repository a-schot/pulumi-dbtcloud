// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbtcloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/a-schot/pulumi-dbtcloud/sdk/go/dbtcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Assigns a set of dbt Cloud groups to a given User ID.
//
// > If additional groups were assigned manually in dbt Cloud, they will be removed. The full list of groups need to be provided as config.
//
// > This resource does not currently support deletion (e.g. a deleted resource will stay as-is in dbt Cloud).
// This is intentional in order to prevent accidental deletion of all users groups assigned to a user.
// If you would like a different behavior, please open an issue on GitHub. To remove all groups for a user, set "groupIds" to the empty set "[]".
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
//			_, err := dbtcloud.NewUserGroups(ctx, "myUserGroups", &dbtcloud.UserGroupsArgs{
//				UserId: pulumi.Any(dbtcloud_user.My_user.Id),
//				GroupIds: pulumi.IntArray{
//					pulumi.Int(1234),
//					dbtcloud_group.My_group.Id,
//					local.My_group_id,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dbtcloud.NewUserGroups(ctx, "myOtherUserGroups", &dbtcloud.UserGroupsArgs{
//				UserId:   pulumi.Int(123456),
//				GroupIds: pulumi.IntArray{},
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
// Import using the User ID The User ID can be retrieved from the dbt Cloud UI or with the data source dbtcloud_user
//
// ```sh
//
//	$ pulumi import dbtcloud:index/userGroups:UserGroups my_user_groups "user_id"
//
// ```
//
// ```sh
//
//	$ pulumi import dbtcloud:index/userGroups:UserGroups my_user_groups 123456
//
// ```
type UserGroups struct {
	pulumi.CustomResourceState

	// IDs of the groups to assign to the user. If additional groups were assigned manually in dbt Cloud, they will be removed.
	GroupIds pulumi.IntArrayOutput `pulumi:"groupIds"`
	// The internal ID of a dbt Cloud user
	UserId pulumi.IntOutput `pulumi:"userId"`
}

// NewUserGroups registers a new resource with the given unique name, arguments, and options.
func NewUserGroups(ctx *pulumi.Context,
	name string, args *UserGroupsArgs, opts ...pulumi.ResourceOption) (*UserGroups, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupIds == nil {
		return nil, errors.New("invalid value for required argument 'GroupIds'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserGroups
	err := ctx.RegisterResource("dbtcloud:index/userGroups:UserGroups", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserGroups gets an existing UserGroups resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserGroups(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserGroupsState, opts ...pulumi.ResourceOption) (*UserGroups, error) {
	var resource UserGroups
	err := ctx.ReadResource("dbtcloud:index/userGroups:UserGroups", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserGroups resources.
type userGroupsState struct {
	// IDs of the groups to assign to the user. If additional groups were assigned manually in dbt Cloud, they will be removed.
	GroupIds []int `pulumi:"groupIds"`
	// The internal ID of a dbt Cloud user
	UserId *int `pulumi:"userId"`
}

type UserGroupsState struct {
	// IDs of the groups to assign to the user. If additional groups were assigned manually in dbt Cloud, they will be removed.
	GroupIds pulumi.IntArrayInput
	// The internal ID of a dbt Cloud user
	UserId pulumi.IntPtrInput
}

func (UserGroupsState) ElementType() reflect.Type {
	return reflect.TypeOf((*userGroupsState)(nil)).Elem()
}

type userGroupsArgs struct {
	// IDs of the groups to assign to the user. If additional groups were assigned manually in dbt Cloud, they will be removed.
	GroupIds []int `pulumi:"groupIds"`
	// The internal ID of a dbt Cloud user
	UserId int `pulumi:"userId"`
}

// The set of arguments for constructing a UserGroups resource.
type UserGroupsArgs struct {
	// IDs of the groups to assign to the user. If additional groups were assigned manually in dbt Cloud, they will be removed.
	GroupIds pulumi.IntArrayInput
	// The internal ID of a dbt Cloud user
	UserId pulumi.IntInput
}

func (UserGroupsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userGroupsArgs)(nil)).Elem()
}

type UserGroupsInput interface {
	pulumi.Input

	ToUserGroupsOutput() UserGroupsOutput
	ToUserGroupsOutputWithContext(ctx context.Context) UserGroupsOutput
}

func (*UserGroups) ElementType() reflect.Type {
	return reflect.TypeOf((**UserGroups)(nil)).Elem()
}

func (i *UserGroups) ToUserGroupsOutput() UserGroupsOutput {
	return i.ToUserGroupsOutputWithContext(context.Background())
}

func (i *UserGroups) ToUserGroupsOutputWithContext(ctx context.Context) UserGroupsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGroupsOutput)
}

// UserGroupsArrayInput is an input type that accepts UserGroupsArray and UserGroupsArrayOutput values.
// You can construct a concrete instance of `UserGroupsArrayInput` via:
//
//	UserGroupsArray{ UserGroupsArgs{...} }
type UserGroupsArrayInput interface {
	pulumi.Input

	ToUserGroupsArrayOutput() UserGroupsArrayOutput
	ToUserGroupsArrayOutputWithContext(context.Context) UserGroupsArrayOutput
}

type UserGroupsArray []UserGroupsInput

func (UserGroupsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserGroups)(nil)).Elem()
}

func (i UserGroupsArray) ToUserGroupsArrayOutput() UserGroupsArrayOutput {
	return i.ToUserGroupsArrayOutputWithContext(context.Background())
}

func (i UserGroupsArray) ToUserGroupsArrayOutputWithContext(ctx context.Context) UserGroupsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGroupsArrayOutput)
}

// UserGroupsMapInput is an input type that accepts UserGroupsMap and UserGroupsMapOutput values.
// You can construct a concrete instance of `UserGroupsMapInput` via:
//
//	UserGroupsMap{ "key": UserGroupsArgs{...} }
type UserGroupsMapInput interface {
	pulumi.Input

	ToUserGroupsMapOutput() UserGroupsMapOutput
	ToUserGroupsMapOutputWithContext(context.Context) UserGroupsMapOutput
}

type UserGroupsMap map[string]UserGroupsInput

func (UserGroupsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserGroups)(nil)).Elem()
}

func (i UserGroupsMap) ToUserGroupsMapOutput() UserGroupsMapOutput {
	return i.ToUserGroupsMapOutputWithContext(context.Background())
}

func (i UserGroupsMap) ToUserGroupsMapOutputWithContext(ctx context.Context) UserGroupsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGroupsMapOutput)
}

type UserGroupsOutput struct{ *pulumi.OutputState }

func (UserGroupsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserGroups)(nil)).Elem()
}

func (o UserGroupsOutput) ToUserGroupsOutput() UserGroupsOutput {
	return o
}

func (o UserGroupsOutput) ToUserGroupsOutputWithContext(ctx context.Context) UserGroupsOutput {
	return o
}

// IDs of the groups to assign to the user. If additional groups were assigned manually in dbt Cloud, they will be removed.
func (o UserGroupsOutput) GroupIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *UserGroups) pulumi.IntArrayOutput { return v.GroupIds }).(pulumi.IntArrayOutput)
}

// The internal ID of a dbt Cloud user
func (o UserGroupsOutput) UserId() pulumi.IntOutput {
	return o.ApplyT(func(v *UserGroups) pulumi.IntOutput { return v.UserId }).(pulumi.IntOutput)
}

type UserGroupsArrayOutput struct{ *pulumi.OutputState }

func (UserGroupsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserGroups)(nil)).Elem()
}

func (o UserGroupsArrayOutput) ToUserGroupsArrayOutput() UserGroupsArrayOutput {
	return o
}

func (o UserGroupsArrayOutput) ToUserGroupsArrayOutputWithContext(ctx context.Context) UserGroupsArrayOutput {
	return o
}

func (o UserGroupsArrayOutput) Index(i pulumi.IntInput) UserGroupsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserGroups {
		return vs[0].([]*UserGroups)[vs[1].(int)]
	}).(UserGroupsOutput)
}

type UserGroupsMapOutput struct{ *pulumi.OutputState }

func (UserGroupsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserGroups)(nil)).Elem()
}

func (o UserGroupsMapOutput) ToUserGroupsMapOutput() UserGroupsMapOutput {
	return o
}

func (o UserGroupsMapOutput) ToUserGroupsMapOutputWithContext(ctx context.Context) UserGroupsMapOutput {
	return o
}

func (o UserGroupsMapOutput) MapIndex(k pulumi.StringInput) UserGroupsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserGroups {
		return vs[0].(map[string]*UserGroups)[vs[1].(string)]
	}).(UserGroupsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserGroupsInput)(nil)).Elem(), &UserGroups{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserGroupsArrayInput)(nil)).Elem(), UserGroupsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserGroupsMapInput)(nil)).Elem(), UserGroupsMap{})
	pulumi.RegisterOutputType(UserGroupsOutput{})
	pulumi.RegisterOutputType(UserGroupsArrayOutput{})
	pulumi.RegisterOutputType(UserGroupsMapOutput{})
}

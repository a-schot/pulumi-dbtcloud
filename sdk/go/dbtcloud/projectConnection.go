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
//			_, err := dbtcloud.NewProjectConnection(ctx, "dbtProjectConnection", &dbtcloud.ProjectConnectionArgs{
//				ProjectId:    pulumi.Any(dbtcloud_project.Dbt_project.Id),
//				ConnectionId: pulumi.Any(dbtcloud_connection.Dbt_connection.Connection_id),
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
// Import using a project ID and Connection ID found in the URL or via the API.
//
// ```sh
//
//	$ pulumi import dbtcloud:index/projectConnection:ProjectConnection my_project "project_id:connection_id"
//
// ```
//
// ```sh
//
//	$ pulumi import dbtcloud:index/projectConnection:ProjectConnection my_project 12345:5678
//
// ```
type ProjectConnection struct {
	pulumi.CustomResourceState

	// Connection ID
	ConnectionId pulumi.IntOutput `pulumi:"connectionId"`
	// Project ID
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
}

// NewProjectConnection registers a new resource with the given unique name, arguments, and options.
func NewProjectConnection(ctx *pulumi.Context,
	name string, args *ProjectConnectionArgs, opts ...pulumi.ResourceOption) (*ProjectConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionId'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectConnection
	err := ctx.RegisterResource("dbtcloud:index/projectConnection:ProjectConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectConnection gets an existing ProjectConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectConnectionState, opts ...pulumi.ResourceOption) (*ProjectConnection, error) {
	var resource ProjectConnection
	err := ctx.ReadResource("dbtcloud:index/projectConnection:ProjectConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectConnection resources.
type projectConnectionState struct {
	// Connection ID
	ConnectionId *int `pulumi:"connectionId"`
	// Project ID
	ProjectId *int `pulumi:"projectId"`
}

type ProjectConnectionState struct {
	// Connection ID
	ConnectionId pulumi.IntPtrInput
	// Project ID
	ProjectId pulumi.IntPtrInput
}

func (ProjectConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectConnectionState)(nil)).Elem()
}

type projectConnectionArgs struct {
	// Connection ID
	ConnectionId int `pulumi:"connectionId"`
	// Project ID
	ProjectId int `pulumi:"projectId"`
}

// The set of arguments for constructing a ProjectConnection resource.
type ProjectConnectionArgs struct {
	// Connection ID
	ConnectionId pulumi.IntInput
	// Project ID
	ProjectId pulumi.IntInput
}

func (ProjectConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectConnectionArgs)(nil)).Elem()
}

type ProjectConnectionInput interface {
	pulumi.Input

	ToProjectConnectionOutput() ProjectConnectionOutput
	ToProjectConnectionOutputWithContext(ctx context.Context) ProjectConnectionOutput
}

func (*ProjectConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectConnection)(nil)).Elem()
}

func (i *ProjectConnection) ToProjectConnectionOutput() ProjectConnectionOutput {
	return i.ToProjectConnectionOutputWithContext(context.Background())
}

func (i *ProjectConnection) ToProjectConnectionOutputWithContext(ctx context.Context) ProjectConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectConnectionOutput)
}

// ProjectConnectionArrayInput is an input type that accepts ProjectConnectionArray and ProjectConnectionArrayOutput values.
// You can construct a concrete instance of `ProjectConnectionArrayInput` via:
//
//	ProjectConnectionArray{ ProjectConnectionArgs{...} }
type ProjectConnectionArrayInput interface {
	pulumi.Input

	ToProjectConnectionArrayOutput() ProjectConnectionArrayOutput
	ToProjectConnectionArrayOutputWithContext(context.Context) ProjectConnectionArrayOutput
}

type ProjectConnectionArray []ProjectConnectionInput

func (ProjectConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectConnection)(nil)).Elem()
}

func (i ProjectConnectionArray) ToProjectConnectionArrayOutput() ProjectConnectionArrayOutput {
	return i.ToProjectConnectionArrayOutputWithContext(context.Background())
}

func (i ProjectConnectionArray) ToProjectConnectionArrayOutputWithContext(ctx context.Context) ProjectConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectConnectionArrayOutput)
}

// ProjectConnectionMapInput is an input type that accepts ProjectConnectionMap and ProjectConnectionMapOutput values.
// You can construct a concrete instance of `ProjectConnectionMapInput` via:
//
//	ProjectConnectionMap{ "key": ProjectConnectionArgs{...} }
type ProjectConnectionMapInput interface {
	pulumi.Input

	ToProjectConnectionMapOutput() ProjectConnectionMapOutput
	ToProjectConnectionMapOutputWithContext(context.Context) ProjectConnectionMapOutput
}

type ProjectConnectionMap map[string]ProjectConnectionInput

func (ProjectConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectConnection)(nil)).Elem()
}

func (i ProjectConnectionMap) ToProjectConnectionMapOutput() ProjectConnectionMapOutput {
	return i.ToProjectConnectionMapOutputWithContext(context.Background())
}

func (i ProjectConnectionMap) ToProjectConnectionMapOutputWithContext(ctx context.Context) ProjectConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectConnectionMapOutput)
}

type ProjectConnectionOutput struct{ *pulumi.OutputState }

func (ProjectConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectConnection)(nil)).Elem()
}

func (o ProjectConnectionOutput) ToProjectConnectionOutput() ProjectConnectionOutput {
	return o
}

func (o ProjectConnectionOutput) ToProjectConnectionOutputWithContext(ctx context.Context) ProjectConnectionOutput {
	return o
}

// Connection ID
func (o ProjectConnectionOutput) ConnectionId() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectConnection) pulumi.IntOutput { return v.ConnectionId }).(pulumi.IntOutput)
}

// Project ID
func (o ProjectConnectionOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectConnection) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

type ProjectConnectionArrayOutput struct{ *pulumi.OutputState }

func (ProjectConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectConnection)(nil)).Elem()
}

func (o ProjectConnectionArrayOutput) ToProjectConnectionArrayOutput() ProjectConnectionArrayOutput {
	return o
}

func (o ProjectConnectionArrayOutput) ToProjectConnectionArrayOutputWithContext(ctx context.Context) ProjectConnectionArrayOutput {
	return o
}

func (o ProjectConnectionArrayOutput) Index(i pulumi.IntInput) ProjectConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectConnection {
		return vs[0].([]*ProjectConnection)[vs[1].(int)]
	}).(ProjectConnectionOutput)
}

type ProjectConnectionMapOutput struct{ *pulumi.OutputState }

func (ProjectConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectConnection)(nil)).Elem()
}

func (o ProjectConnectionMapOutput) ToProjectConnectionMapOutput() ProjectConnectionMapOutput {
	return o
}

func (o ProjectConnectionMapOutput) ToProjectConnectionMapOutputWithContext(ctx context.Context) ProjectConnectionMapOutput {
	return o
}

func (o ProjectConnectionMapOutput) MapIndex(k pulumi.StringInput) ProjectConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectConnection {
		return vs[0].(map[string]*ProjectConnection)[vs[1].(string)]
	}).(ProjectConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectConnectionInput)(nil)).Elem(), &ProjectConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectConnectionArrayInput)(nil)).Elem(), ProjectConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectConnectionMapInput)(nil)).Elem(), ProjectConnectionMap{})
	pulumi.RegisterOutputType(ProjectConnectionOutput{})
	pulumi.RegisterOutputType(ProjectConnectionArrayOutput{})
	pulumi.RegisterOutputType(ProjectConnectionMapOutput{})
}

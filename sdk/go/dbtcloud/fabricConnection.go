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

// Resource to create Microsoft Fabric connections in dbt Cloud
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
//			_, err := dbtcloud.NewFabricConnection(ctx, "myFabricConnection", &dbtcloud.FabricConnectionArgs{
//				ProjectId:    pulumi.Any(dbtcloud_project.Dbt_project.Id),
//				Server:       pulumi.String("my-server"),
//				Database:     pulumi.String("my-database"),
//				Port:         pulumi.Int(1234),
//				LoginTimeout: pulumi.Int(30),
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
// ```sh
//
//	$ pulumi import dbtcloud:index/fabricConnection:FabricConnection my_connection "project_id:connection_id"
//
// ```
//
// ```sh
//
//	$ pulumi import dbtcloud:index/fabricConnection:FabricConnection my_connection 12345:6789
//
// ```
type FabricConnection struct {
	pulumi.CustomResourceState

	// Adapter id created for the Fabric connection
	AdapterId pulumi.IntOutput `pulumi:"adapterId"`
	// Connection Identifier
	ConnectionId pulumi.IntOutput `pulumi:"connectionId"`
	// The database to connect to for this connection.
	Database pulumi.StringOutput `pulumi:"database"`
	// The number of seconds used to establish a connection before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
	LoginTimeout pulumi.IntPtrOutput `pulumi:"loginTimeout"`
	// Connection name
	Name pulumi.StringOutput `pulumi:"name"`
	// The port to connect to for this connection.
	Port pulumi.IntOutput `pulumi:"port"`
	// Project ID to create the connection in
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// The number of seconds used to wait for a query before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
	QueryTimeout pulumi.IntPtrOutput `pulumi:"queryTimeout"`
	// The number of automatic times to retry a query before failing. Defaults to 1. Queries with syntax errors will not be retried. This setting can be used to overcome intermittent network issues.
	Retries pulumi.IntPtrOutput `pulumi:"retries"`
	// The server hostname.
	Server pulumi.StringOutput `pulumi:"server"`
}

// NewFabricConnection registers a new resource with the given unique name, arguments, and options.
func NewFabricConnection(ctx *pulumi.Context,
	name string, args *FabricConnectionArgs, opts ...pulumi.ResourceOption) (*FabricConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Database == nil {
		return nil, errors.New("invalid value for required argument 'Database'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Server == nil {
		return nil, errors.New("invalid value for required argument 'Server'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FabricConnection
	err := ctx.RegisterResource("dbtcloud:index/fabricConnection:FabricConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFabricConnection gets an existing FabricConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFabricConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FabricConnectionState, opts ...pulumi.ResourceOption) (*FabricConnection, error) {
	var resource FabricConnection
	err := ctx.ReadResource("dbtcloud:index/fabricConnection:FabricConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FabricConnection resources.
type fabricConnectionState struct {
	// Adapter id created for the Fabric connection
	AdapterId *int `pulumi:"adapterId"`
	// Connection Identifier
	ConnectionId *int `pulumi:"connectionId"`
	// The database to connect to for this connection.
	Database *string `pulumi:"database"`
	// The number of seconds used to establish a connection before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
	LoginTimeout *int `pulumi:"loginTimeout"`
	// Connection name
	Name *string `pulumi:"name"`
	// The port to connect to for this connection.
	Port *int `pulumi:"port"`
	// Project ID to create the connection in
	ProjectId *int `pulumi:"projectId"`
	// The number of seconds used to wait for a query before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
	QueryTimeout *int `pulumi:"queryTimeout"`
	// The number of automatic times to retry a query before failing. Defaults to 1. Queries with syntax errors will not be retried. This setting can be used to overcome intermittent network issues.
	Retries *int `pulumi:"retries"`
	// The server hostname.
	Server *string `pulumi:"server"`
}

type FabricConnectionState struct {
	// Adapter id created for the Fabric connection
	AdapterId pulumi.IntPtrInput
	// Connection Identifier
	ConnectionId pulumi.IntPtrInput
	// The database to connect to for this connection.
	Database pulumi.StringPtrInput
	// The number of seconds used to establish a connection before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
	LoginTimeout pulumi.IntPtrInput
	// Connection name
	Name pulumi.StringPtrInput
	// The port to connect to for this connection.
	Port pulumi.IntPtrInput
	// Project ID to create the connection in
	ProjectId pulumi.IntPtrInput
	// The number of seconds used to wait for a query before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
	QueryTimeout pulumi.IntPtrInput
	// The number of automatic times to retry a query before failing. Defaults to 1. Queries with syntax errors will not be retried. This setting can be used to overcome intermittent network issues.
	Retries pulumi.IntPtrInput
	// The server hostname.
	Server pulumi.StringPtrInput
}

func (FabricConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*fabricConnectionState)(nil)).Elem()
}

type fabricConnectionArgs struct {
	// The database to connect to for this connection.
	Database string `pulumi:"database"`
	// The number of seconds used to establish a connection before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
	LoginTimeout *int `pulumi:"loginTimeout"`
	// Connection name
	Name *string `pulumi:"name"`
	// The port to connect to for this connection.
	Port int `pulumi:"port"`
	// Project ID to create the connection in
	ProjectId int `pulumi:"projectId"`
	// The number of seconds used to wait for a query before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
	QueryTimeout *int `pulumi:"queryTimeout"`
	// The number of automatic times to retry a query before failing. Defaults to 1. Queries with syntax errors will not be retried. This setting can be used to overcome intermittent network issues.
	Retries *int `pulumi:"retries"`
	// The server hostname.
	Server string `pulumi:"server"`
}

// The set of arguments for constructing a FabricConnection resource.
type FabricConnectionArgs struct {
	// The database to connect to for this connection.
	Database pulumi.StringInput
	// The number of seconds used to establish a connection before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
	LoginTimeout pulumi.IntPtrInput
	// Connection name
	Name pulumi.StringPtrInput
	// The port to connect to for this connection.
	Port pulumi.IntInput
	// Project ID to create the connection in
	ProjectId pulumi.IntInput
	// The number of seconds used to wait for a query before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
	QueryTimeout pulumi.IntPtrInput
	// The number of automatic times to retry a query before failing. Defaults to 1. Queries with syntax errors will not be retried. This setting can be used to overcome intermittent network issues.
	Retries pulumi.IntPtrInput
	// The server hostname.
	Server pulumi.StringInput
}

func (FabricConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fabricConnectionArgs)(nil)).Elem()
}

type FabricConnectionInput interface {
	pulumi.Input

	ToFabricConnectionOutput() FabricConnectionOutput
	ToFabricConnectionOutputWithContext(ctx context.Context) FabricConnectionOutput
}

func (*FabricConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**FabricConnection)(nil)).Elem()
}

func (i *FabricConnection) ToFabricConnectionOutput() FabricConnectionOutput {
	return i.ToFabricConnectionOutputWithContext(context.Background())
}

func (i *FabricConnection) ToFabricConnectionOutputWithContext(ctx context.Context) FabricConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FabricConnectionOutput)
}

// FabricConnectionArrayInput is an input type that accepts FabricConnectionArray and FabricConnectionArrayOutput values.
// You can construct a concrete instance of `FabricConnectionArrayInput` via:
//
//	FabricConnectionArray{ FabricConnectionArgs{...} }
type FabricConnectionArrayInput interface {
	pulumi.Input

	ToFabricConnectionArrayOutput() FabricConnectionArrayOutput
	ToFabricConnectionArrayOutputWithContext(context.Context) FabricConnectionArrayOutput
}

type FabricConnectionArray []FabricConnectionInput

func (FabricConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FabricConnection)(nil)).Elem()
}

func (i FabricConnectionArray) ToFabricConnectionArrayOutput() FabricConnectionArrayOutput {
	return i.ToFabricConnectionArrayOutputWithContext(context.Background())
}

func (i FabricConnectionArray) ToFabricConnectionArrayOutputWithContext(ctx context.Context) FabricConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FabricConnectionArrayOutput)
}

// FabricConnectionMapInput is an input type that accepts FabricConnectionMap and FabricConnectionMapOutput values.
// You can construct a concrete instance of `FabricConnectionMapInput` via:
//
//	FabricConnectionMap{ "key": FabricConnectionArgs{...} }
type FabricConnectionMapInput interface {
	pulumi.Input

	ToFabricConnectionMapOutput() FabricConnectionMapOutput
	ToFabricConnectionMapOutputWithContext(context.Context) FabricConnectionMapOutput
}

type FabricConnectionMap map[string]FabricConnectionInput

func (FabricConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FabricConnection)(nil)).Elem()
}

func (i FabricConnectionMap) ToFabricConnectionMapOutput() FabricConnectionMapOutput {
	return i.ToFabricConnectionMapOutputWithContext(context.Background())
}

func (i FabricConnectionMap) ToFabricConnectionMapOutputWithContext(ctx context.Context) FabricConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FabricConnectionMapOutput)
}

type FabricConnectionOutput struct{ *pulumi.OutputState }

func (FabricConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FabricConnection)(nil)).Elem()
}

func (o FabricConnectionOutput) ToFabricConnectionOutput() FabricConnectionOutput {
	return o
}

func (o FabricConnectionOutput) ToFabricConnectionOutputWithContext(ctx context.Context) FabricConnectionOutput {
	return o
}

// Adapter id created for the Fabric connection
func (o FabricConnectionOutput) AdapterId() pulumi.IntOutput {
	return o.ApplyT(func(v *FabricConnection) pulumi.IntOutput { return v.AdapterId }).(pulumi.IntOutput)
}

// Connection Identifier
func (o FabricConnectionOutput) ConnectionId() pulumi.IntOutput {
	return o.ApplyT(func(v *FabricConnection) pulumi.IntOutput { return v.ConnectionId }).(pulumi.IntOutput)
}

// The database to connect to for this connection.
func (o FabricConnectionOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricConnection) pulumi.StringOutput { return v.Database }).(pulumi.StringOutput)
}

// The number of seconds used to establish a connection before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
func (o FabricConnectionOutput) LoginTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FabricConnection) pulumi.IntPtrOutput { return v.LoginTimeout }).(pulumi.IntPtrOutput)
}

// Connection name
func (o FabricConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The port to connect to for this connection.
func (o FabricConnectionOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *FabricConnection) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Project ID to create the connection in
func (o FabricConnectionOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *FabricConnection) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

// The number of seconds used to wait for a query before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
func (o FabricConnectionOutput) QueryTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FabricConnection) pulumi.IntPtrOutput { return v.QueryTimeout }).(pulumi.IntPtrOutput)
}

// The number of automatic times to retry a query before failing. Defaults to 1. Queries with syntax errors will not be retried. This setting can be used to overcome intermittent network issues.
func (o FabricConnectionOutput) Retries() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FabricConnection) pulumi.IntPtrOutput { return v.Retries }).(pulumi.IntPtrOutput)
}

// The server hostname.
func (o FabricConnectionOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *FabricConnection) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

type FabricConnectionArrayOutput struct{ *pulumi.OutputState }

func (FabricConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FabricConnection)(nil)).Elem()
}

func (o FabricConnectionArrayOutput) ToFabricConnectionArrayOutput() FabricConnectionArrayOutput {
	return o
}

func (o FabricConnectionArrayOutput) ToFabricConnectionArrayOutputWithContext(ctx context.Context) FabricConnectionArrayOutput {
	return o
}

func (o FabricConnectionArrayOutput) Index(i pulumi.IntInput) FabricConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FabricConnection {
		return vs[0].([]*FabricConnection)[vs[1].(int)]
	}).(FabricConnectionOutput)
}

type FabricConnectionMapOutput struct{ *pulumi.OutputState }

func (FabricConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FabricConnection)(nil)).Elem()
}

func (o FabricConnectionMapOutput) ToFabricConnectionMapOutput() FabricConnectionMapOutput {
	return o
}

func (o FabricConnectionMapOutput) ToFabricConnectionMapOutputWithContext(ctx context.Context) FabricConnectionMapOutput {
	return o
}

func (o FabricConnectionMapOutput) MapIndex(k pulumi.StringInput) FabricConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FabricConnection {
		return vs[0].(map[string]*FabricConnection)[vs[1].(string)]
	}).(FabricConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FabricConnectionInput)(nil)).Elem(), &FabricConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*FabricConnectionArrayInput)(nil)).Elem(), FabricConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FabricConnectionMapInput)(nil)).Elem(), FabricConnectionMap{})
	pulumi.RegisterOutputType(FabricConnectionOutput{})
	pulumi.RegisterOutputType(FabricConnectionArrayOutput{})
	pulumi.RegisterOutputType(FabricConnectionMapOutput{})
}

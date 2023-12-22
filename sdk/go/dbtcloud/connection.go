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

// Create a Data Warehouse connection for your project in dbt Cloud. The connection will need to be linked to the dbt Cloud project via a `ProjectConnection` resource.
//
// This resource can be used for Databricks, Postgres, Redshift, Snowflake and AlloyDB connections.
// For BigQuery, due to the list of fields being very different, you can use the `BigqueryConnection` resource.
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
//			_, err := dbtcloud.NewConnection(ctx, "databricks", &dbtcloud.ConnectionArgs{
//				ProjectId: pulumi.Any(dbtcloud_project.Dbt_project.Id),
//				Type:      pulumi.String("adapter"),
//				Database:  pulumi.String(""),
//				HostName:  pulumi.String("my-databricks-host.cloud.databricks.com"),
//				HttpPath:  pulumi.String("/my/path"),
//				Catalog:   pulumi.String("moo"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dbtcloud.NewConnection(ctx, "redshift", &dbtcloud.ConnectionArgs{
//				ProjectId: pulumi.Any(dbtcloud_project.Dbt_project.Id),
//				Type:      pulumi.String("redshift"),
//				Database:  pulumi.String("my-database"),
//				Port:      pulumi.Int(5439),
//				HostName:  pulumi.String("my-redshift-hostname"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dbtcloud.NewConnection(ctx, "snowflake", &dbtcloud.ConnectionArgs{
//				ProjectId: pulumi.Any(dbtcloud_project.Dbt_project.Id),
//				Type:      pulumi.String("snowflake"),
//				Account:   pulumi.String("my-snowflake-account"),
//				Database:  pulumi.String("MY_DATABASE"),
//				Role:      pulumi.String("MY_ROLE"),
//				Warehouse: pulumi.String("MY_WAREHOUSE"),
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
// Import using a project ID and connection ID found in the URL or via the API.
//
// ```sh
//
//	$ pulumi import dbtcloud:index/connection:Connection test_connection "project_id:connection_id"
//
// ```
//
// ```sh
//
//	$ pulumi import dbtcloud:index/connection:Connection test_connection 12345:6789
//
// ```
type Connection struct {
	pulumi.CustomResourceState

	// Account name for the connection
	Account pulumi.StringPtrOutput `pulumi:"account"`
	// Adapter id created for the Databricks connection
	AdapterId pulumi.IntOutput `pulumi:"adapterId"`
	// Whether or not the connection should allow client session keep alive
	AllowKeepAlive pulumi.BoolPtrOutput `pulumi:"allowKeepAlive"`
	// Whether or not the connection should allow SSO
	AllowSso pulumi.BoolPtrOutput `pulumi:"allowSso"`
	// Catalog name if Unity Catalog is enabled in your Databricks workspace
	Catalog pulumi.StringPtrOutput `pulumi:"catalog"`
	// Connection Identifier
	ConnectionId pulumi.IntOutput `pulumi:"connectionId"`
	// Database name for the connection
	Database pulumi.StringOutput `pulumi:"database"`
	// Host name for the connection, including Databricks cluster
	HostName pulumi.StringPtrOutput `pulumi:"hostName"`
	// The HTTP path of the Databricks cluster or SQL warehouse
	HttpPath pulumi.StringPtrOutput `pulumi:"httpPath"`
	// Whether the connection is active
	IsActive pulumi.BoolPtrOutput `pulumi:"isActive"`
	// Connection name
	Name pulumi.StringOutput `pulumi:"name"`
	// OAuth client identifier
	OauthClientId pulumi.StringPtrOutput `pulumi:"oauthClientId"`
	// OAuth client secret
	OauthClientSecret pulumi.StringPtrOutput `pulumi:"oauthClientSecret"`
	// Port number to connect via
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The ID of the PrivateLink connection. This ID can be found using the `privatelinkEndpoint` data source
	PrivateLinkEndpointId pulumi.StringPtrOutput `pulumi:"privateLinkEndpointId"`
	// Project ID to create the connection in
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// Role name for the connection
	Role pulumi.StringPtrOutput `pulumi:"role"`
	// Whether or not tunneling should be enabled on your database connection
	TunnelEnabled pulumi.BoolPtrOutput `pulumi:"tunnelEnabled"`
	// The type of connection
	Type pulumi.StringOutput `pulumi:"type"`
	// Warehouse name for the connection
	Warehouse pulumi.StringPtrOutput `pulumi:"warehouse"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Database == nil {
		return nil, errors.New("invalid value for required argument 'Database'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Connection
	err := ctx.RegisterResource("dbtcloud:index/connection:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("dbtcloud:index/connection:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
	// Account name for the connection
	Account *string `pulumi:"account"`
	// Adapter id created for the Databricks connection
	AdapterId *int `pulumi:"adapterId"`
	// Whether or not the connection should allow client session keep alive
	AllowKeepAlive *bool `pulumi:"allowKeepAlive"`
	// Whether or not the connection should allow SSO
	AllowSso *bool `pulumi:"allowSso"`
	// Catalog name if Unity Catalog is enabled in your Databricks workspace
	Catalog *string `pulumi:"catalog"`
	// Connection Identifier
	ConnectionId *int `pulumi:"connectionId"`
	// Database name for the connection
	Database *string `pulumi:"database"`
	// Host name for the connection, including Databricks cluster
	HostName *string `pulumi:"hostName"`
	// The HTTP path of the Databricks cluster or SQL warehouse
	HttpPath *string `pulumi:"httpPath"`
	// Whether the connection is active
	IsActive *bool `pulumi:"isActive"`
	// Connection name
	Name *string `pulumi:"name"`
	// OAuth client identifier
	OauthClientId *string `pulumi:"oauthClientId"`
	// OAuth client secret
	OauthClientSecret *string `pulumi:"oauthClientSecret"`
	// Port number to connect via
	Port *int `pulumi:"port"`
	// The ID of the PrivateLink connection. This ID can be found using the `privatelinkEndpoint` data source
	PrivateLinkEndpointId *string `pulumi:"privateLinkEndpointId"`
	// Project ID to create the connection in
	ProjectId *int `pulumi:"projectId"`
	// Role name for the connection
	Role *string `pulumi:"role"`
	// Whether or not tunneling should be enabled on your database connection
	TunnelEnabled *bool `pulumi:"tunnelEnabled"`
	// The type of connection
	Type *string `pulumi:"type"`
	// Warehouse name for the connection
	Warehouse *string `pulumi:"warehouse"`
}

type ConnectionState struct {
	// Account name for the connection
	Account pulumi.StringPtrInput
	// Adapter id created for the Databricks connection
	AdapterId pulumi.IntPtrInput
	// Whether or not the connection should allow client session keep alive
	AllowKeepAlive pulumi.BoolPtrInput
	// Whether or not the connection should allow SSO
	AllowSso pulumi.BoolPtrInput
	// Catalog name if Unity Catalog is enabled in your Databricks workspace
	Catalog pulumi.StringPtrInput
	// Connection Identifier
	ConnectionId pulumi.IntPtrInput
	// Database name for the connection
	Database pulumi.StringPtrInput
	// Host name for the connection, including Databricks cluster
	HostName pulumi.StringPtrInput
	// The HTTP path of the Databricks cluster or SQL warehouse
	HttpPath pulumi.StringPtrInput
	// Whether the connection is active
	IsActive pulumi.BoolPtrInput
	// Connection name
	Name pulumi.StringPtrInput
	// OAuth client identifier
	OauthClientId pulumi.StringPtrInput
	// OAuth client secret
	OauthClientSecret pulumi.StringPtrInput
	// Port number to connect via
	Port pulumi.IntPtrInput
	// The ID of the PrivateLink connection. This ID can be found using the `privatelinkEndpoint` data source
	PrivateLinkEndpointId pulumi.StringPtrInput
	// Project ID to create the connection in
	ProjectId pulumi.IntPtrInput
	// Role name for the connection
	Role pulumi.StringPtrInput
	// Whether or not tunneling should be enabled on your database connection
	TunnelEnabled pulumi.BoolPtrInput
	// The type of connection
	Type pulumi.StringPtrInput
	// Warehouse name for the connection
	Warehouse pulumi.StringPtrInput
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// Account name for the connection
	Account *string `pulumi:"account"`
	// Whether or not the connection should allow client session keep alive
	AllowKeepAlive *bool `pulumi:"allowKeepAlive"`
	// Whether or not the connection should allow SSO
	AllowSso *bool `pulumi:"allowSso"`
	// Catalog name if Unity Catalog is enabled in your Databricks workspace
	Catalog *string `pulumi:"catalog"`
	// Database name for the connection
	Database string `pulumi:"database"`
	// Host name for the connection, including Databricks cluster
	HostName *string `pulumi:"hostName"`
	// The HTTP path of the Databricks cluster or SQL warehouse
	HttpPath *string `pulumi:"httpPath"`
	// Whether the connection is active
	IsActive *bool `pulumi:"isActive"`
	// Connection name
	Name *string `pulumi:"name"`
	// OAuth client identifier
	OauthClientId *string `pulumi:"oauthClientId"`
	// OAuth client secret
	OauthClientSecret *string `pulumi:"oauthClientSecret"`
	// Port number to connect via
	Port *int `pulumi:"port"`
	// The ID of the PrivateLink connection. This ID can be found using the `privatelinkEndpoint` data source
	PrivateLinkEndpointId *string `pulumi:"privateLinkEndpointId"`
	// Project ID to create the connection in
	ProjectId int `pulumi:"projectId"`
	// Role name for the connection
	Role *string `pulumi:"role"`
	// Whether or not tunneling should be enabled on your database connection
	TunnelEnabled *bool `pulumi:"tunnelEnabled"`
	// The type of connection
	Type string `pulumi:"type"`
	// Warehouse name for the connection
	Warehouse *string `pulumi:"warehouse"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// Account name for the connection
	Account pulumi.StringPtrInput
	// Whether or not the connection should allow client session keep alive
	AllowKeepAlive pulumi.BoolPtrInput
	// Whether or not the connection should allow SSO
	AllowSso pulumi.BoolPtrInput
	// Catalog name if Unity Catalog is enabled in your Databricks workspace
	Catalog pulumi.StringPtrInput
	// Database name for the connection
	Database pulumi.StringInput
	// Host name for the connection, including Databricks cluster
	HostName pulumi.StringPtrInput
	// The HTTP path of the Databricks cluster or SQL warehouse
	HttpPath pulumi.StringPtrInput
	// Whether the connection is active
	IsActive pulumi.BoolPtrInput
	// Connection name
	Name pulumi.StringPtrInput
	// OAuth client identifier
	OauthClientId pulumi.StringPtrInput
	// OAuth client secret
	OauthClientSecret pulumi.StringPtrInput
	// Port number to connect via
	Port pulumi.IntPtrInput
	// The ID of the PrivateLink connection. This ID can be found using the `privatelinkEndpoint` data source
	PrivateLinkEndpointId pulumi.StringPtrInput
	// Project ID to create the connection in
	ProjectId pulumi.IntInput
	// Role name for the connection
	Role pulumi.StringPtrInput
	// Whether or not tunneling should be enabled on your database connection
	TunnelEnabled pulumi.BoolPtrInput
	// The type of connection
	Type pulumi.StringInput
	// Warehouse name for the connection
	Warehouse pulumi.StringPtrInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

// ConnectionArrayInput is an input type that accepts ConnectionArray and ConnectionArrayOutput values.
// You can construct a concrete instance of `ConnectionArrayInput` via:
//
//	ConnectionArray{ ConnectionArgs{...} }
type ConnectionArrayInput interface {
	pulumi.Input

	ToConnectionArrayOutput() ConnectionArrayOutput
	ToConnectionArrayOutputWithContext(context.Context) ConnectionArrayOutput
}

type ConnectionArray []ConnectionInput

func (ConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connection)(nil)).Elem()
}

func (i ConnectionArray) ToConnectionArrayOutput() ConnectionArrayOutput {
	return i.ToConnectionArrayOutputWithContext(context.Background())
}

func (i ConnectionArray) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionArrayOutput)
}

// ConnectionMapInput is an input type that accepts ConnectionMap and ConnectionMapOutput values.
// You can construct a concrete instance of `ConnectionMapInput` via:
//
//	ConnectionMap{ "key": ConnectionArgs{...} }
type ConnectionMapInput interface {
	pulumi.Input

	ToConnectionMapOutput() ConnectionMapOutput
	ToConnectionMapOutputWithContext(context.Context) ConnectionMapOutput
}

type ConnectionMap map[string]ConnectionInput

func (ConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connection)(nil)).Elem()
}

func (i ConnectionMap) ToConnectionMapOutput() ConnectionMapOutput {
	return i.ToConnectionMapOutputWithContext(context.Background())
}

func (i ConnectionMap) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMapOutput)
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

// Account name for the connection
func (o ConnectionOutput) Account() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.Account }).(pulumi.StringPtrOutput)
}

// Adapter id created for the Databricks connection
func (o ConnectionOutput) AdapterId() pulumi.IntOutput {
	return o.ApplyT(func(v *Connection) pulumi.IntOutput { return v.AdapterId }).(pulumi.IntOutput)
}

// Whether or not the connection should allow client session keep alive
func (o ConnectionOutput) AllowKeepAlive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.BoolPtrOutput { return v.AllowKeepAlive }).(pulumi.BoolPtrOutput)
}

// Whether or not the connection should allow SSO
func (o ConnectionOutput) AllowSso() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.BoolPtrOutput { return v.AllowSso }).(pulumi.BoolPtrOutput)
}

// Catalog name if Unity Catalog is enabled in your Databricks workspace
func (o ConnectionOutput) Catalog() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.Catalog }).(pulumi.StringPtrOutput)
}

// Connection Identifier
func (o ConnectionOutput) ConnectionId() pulumi.IntOutput {
	return o.ApplyT(func(v *Connection) pulumi.IntOutput { return v.ConnectionId }).(pulumi.IntOutput)
}

// Database name for the connection
func (o ConnectionOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Database }).(pulumi.StringOutput)
}

// Host name for the connection, including Databricks cluster
func (o ConnectionOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.HostName }).(pulumi.StringPtrOutput)
}

// The HTTP path of the Databricks cluster or SQL warehouse
func (o ConnectionOutput) HttpPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.HttpPath }).(pulumi.StringPtrOutput)
}

// Whether the connection is active
func (o ConnectionOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.BoolPtrOutput { return v.IsActive }).(pulumi.BoolPtrOutput)
}

// Connection name
func (o ConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// OAuth client identifier
func (o ConnectionOutput) OauthClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.OauthClientId }).(pulumi.StringPtrOutput)
}

// OAuth client secret
func (o ConnectionOutput) OauthClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.OauthClientSecret }).(pulumi.StringPtrOutput)
}

// Port number to connect via
func (o ConnectionOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// The ID of the PrivateLink connection. This ID can be found using the `privatelinkEndpoint` data source
func (o ConnectionOutput) PrivateLinkEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.PrivateLinkEndpointId }).(pulumi.StringPtrOutput)
}

// Project ID to create the connection in
func (o ConnectionOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *Connection) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

// Role name for the connection
func (o ConnectionOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.Role }).(pulumi.StringPtrOutput)
}

// Whether or not tunneling should be enabled on your database connection
func (o ConnectionOutput) TunnelEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.BoolPtrOutput { return v.TunnelEnabled }).(pulumi.BoolPtrOutput)
}

// The type of connection
func (o ConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Warehouse name for the connection
func (o ConnectionOutput) Warehouse() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.Warehouse }).(pulumi.StringPtrOutput)
}

type ConnectionArrayOutput struct{ *pulumi.OutputState }

func (ConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connection)(nil)).Elem()
}

func (o ConnectionArrayOutput) ToConnectionArrayOutput() ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) Index(i pulumi.IntInput) ConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Connection {
		return vs[0].([]*Connection)[vs[1].(int)]
	}).(ConnectionOutput)
}

type ConnectionMapOutput struct{ *pulumi.OutputState }

func (ConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connection)(nil)).Elem()
}

func (o ConnectionMapOutput) ToConnectionMapOutput() ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) MapIndex(k pulumi.StringInput) ConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Connection {
		return vs[0].(map[string]*Connection)[vs[1].(string)]
	}).(ConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), &Connection{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionArrayInput)(nil)).Elem(), ConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionMapInput)(nil)).Elem(), ConnectionMap{})
	pulumi.RegisterOutputType(ConnectionOutput{})
	pulumi.RegisterOutputType(ConnectionArrayOutput{})
	pulumi.RegisterOutputType(ConnectionMapOutput{})
}
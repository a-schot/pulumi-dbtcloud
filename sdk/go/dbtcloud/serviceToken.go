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
//			// NOTE for customers using the LEGACY dbt_cloud provider:
//			_, err := dbtcloud.NewServiceToken(ctx, "testServiceToken", &dbtcloud.ServiceTokenArgs{
//				ServiceTokenPermissions: dbtcloud.ServiceTokenServiceTokenPermissionArray{
//					&dbtcloud.ServiceTokenServiceTokenPermissionArgs{
//						PermissionSet: pulumi.String("git_admin"),
//						AllProjects:   pulumi.Bool(true),
//					},
//					&dbtcloud.ServiceTokenServiceTokenPermissionArgs{
//						PermissionSet: pulumi.String("job_admin"),
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
// $ pulumi import dbtcloud:index/serviceToken:ServiceToken test_service_token "service_token_id"
// ```
//
// ```sh
// $ pulumi import dbtcloud:index/serviceToken:ServiceToken test_service_token 12345
// ```
type ServiceToken struct {
	pulumi.CustomResourceState

	// Service token name
	Name pulumi.StringOutput `pulumi:"name"`
	// Permissions set for the service token
	ServiceTokenPermissions ServiceTokenServiceTokenPermissionArrayOutput `pulumi:"serviceTokenPermissions"`
	// Service token state (1 is active, 2 is inactive)
	State pulumi.IntPtrOutput `pulumi:"state"`
	// Service token secret value (only accessible on creation))
	TokenString pulumi.StringOutput `pulumi:"tokenString"`
	// Service token UID (part of the token)
	Uid pulumi.StringOutput `pulumi:"uid"`
}

// NewServiceToken registers a new resource with the given unique name, arguments, and options.
func NewServiceToken(ctx *pulumi.Context,
	name string, args *ServiceTokenArgs, opts ...pulumi.ResourceOption) (*ServiceToken, error) {
	if args == nil {
		args = &ServiceTokenArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tokenString",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceToken
	err := ctx.RegisterResource("dbtcloud:index/serviceToken:ServiceToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceToken gets an existing ServiceToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceTokenState, opts ...pulumi.ResourceOption) (*ServiceToken, error) {
	var resource ServiceToken
	err := ctx.ReadResource("dbtcloud:index/serviceToken:ServiceToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceToken resources.
type serviceTokenState struct {
	// Service token name
	Name *string `pulumi:"name"`
	// Permissions set for the service token
	ServiceTokenPermissions []ServiceTokenServiceTokenPermission `pulumi:"serviceTokenPermissions"`
	// Service token state (1 is active, 2 is inactive)
	State *int `pulumi:"state"`
	// Service token secret value (only accessible on creation))
	TokenString *string `pulumi:"tokenString"`
	// Service token UID (part of the token)
	Uid *string `pulumi:"uid"`
}

type ServiceTokenState struct {
	// Service token name
	Name pulumi.StringPtrInput
	// Permissions set for the service token
	ServiceTokenPermissions ServiceTokenServiceTokenPermissionArrayInput
	// Service token state (1 is active, 2 is inactive)
	State pulumi.IntPtrInput
	// Service token secret value (only accessible on creation))
	TokenString pulumi.StringPtrInput
	// Service token UID (part of the token)
	Uid pulumi.StringPtrInput
}

func (ServiceTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTokenState)(nil)).Elem()
}

type serviceTokenArgs struct {
	// Service token name
	Name *string `pulumi:"name"`
	// Permissions set for the service token
	ServiceTokenPermissions []ServiceTokenServiceTokenPermission `pulumi:"serviceTokenPermissions"`
	// Service token state (1 is active, 2 is inactive)
	State *int `pulumi:"state"`
}

// The set of arguments for constructing a ServiceToken resource.
type ServiceTokenArgs struct {
	// Service token name
	Name pulumi.StringPtrInput
	// Permissions set for the service token
	ServiceTokenPermissions ServiceTokenServiceTokenPermissionArrayInput
	// Service token state (1 is active, 2 is inactive)
	State pulumi.IntPtrInput
}

func (ServiceTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTokenArgs)(nil)).Elem()
}

type ServiceTokenInput interface {
	pulumi.Input

	ToServiceTokenOutput() ServiceTokenOutput
	ToServiceTokenOutputWithContext(ctx context.Context) ServiceTokenOutput
}

func (*ServiceToken) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceToken)(nil)).Elem()
}

func (i *ServiceToken) ToServiceTokenOutput() ServiceTokenOutput {
	return i.ToServiceTokenOutputWithContext(context.Background())
}

func (i *ServiceToken) ToServiceTokenOutputWithContext(ctx context.Context) ServiceTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTokenOutput)
}

// ServiceTokenArrayInput is an input type that accepts ServiceTokenArray and ServiceTokenArrayOutput values.
// You can construct a concrete instance of `ServiceTokenArrayInput` via:
//
//	ServiceTokenArray{ ServiceTokenArgs{...} }
type ServiceTokenArrayInput interface {
	pulumi.Input

	ToServiceTokenArrayOutput() ServiceTokenArrayOutput
	ToServiceTokenArrayOutputWithContext(context.Context) ServiceTokenArrayOutput
}

type ServiceTokenArray []ServiceTokenInput

func (ServiceTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceToken)(nil)).Elem()
}

func (i ServiceTokenArray) ToServiceTokenArrayOutput() ServiceTokenArrayOutput {
	return i.ToServiceTokenArrayOutputWithContext(context.Background())
}

func (i ServiceTokenArray) ToServiceTokenArrayOutputWithContext(ctx context.Context) ServiceTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTokenArrayOutput)
}

// ServiceTokenMapInput is an input type that accepts ServiceTokenMap and ServiceTokenMapOutput values.
// You can construct a concrete instance of `ServiceTokenMapInput` via:
//
//	ServiceTokenMap{ "key": ServiceTokenArgs{...} }
type ServiceTokenMapInput interface {
	pulumi.Input

	ToServiceTokenMapOutput() ServiceTokenMapOutput
	ToServiceTokenMapOutputWithContext(context.Context) ServiceTokenMapOutput
}

type ServiceTokenMap map[string]ServiceTokenInput

func (ServiceTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceToken)(nil)).Elem()
}

func (i ServiceTokenMap) ToServiceTokenMapOutput() ServiceTokenMapOutput {
	return i.ToServiceTokenMapOutputWithContext(context.Background())
}

func (i ServiceTokenMap) ToServiceTokenMapOutputWithContext(ctx context.Context) ServiceTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTokenMapOutput)
}

type ServiceTokenOutput struct{ *pulumi.OutputState }

func (ServiceTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceToken)(nil)).Elem()
}

func (o ServiceTokenOutput) ToServiceTokenOutput() ServiceTokenOutput {
	return o
}

func (o ServiceTokenOutput) ToServiceTokenOutputWithContext(ctx context.Context) ServiceTokenOutput {
	return o
}

// Service token name
func (o ServiceTokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceToken) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Permissions set for the service token
func (o ServiceTokenOutput) ServiceTokenPermissions() ServiceTokenServiceTokenPermissionArrayOutput {
	return o.ApplyT(func(v *ServiceToken) ServiceTokenServiceTokenPermissionArrayOutput { return v.ServiceTokenPermissions }).(ServiceTokenServiceTokenPermissionArrayOutput)
}

// Service token state (1 is active, 2 is inactive)
func (o ServiceTokenOutput) State() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceToken) pulumi.IntPtrOutput { return v.State }).(pulumi.IntPtrOutput)
}

// Service token secret value (only accessible on creation))
func (o ServiceTokenOutput) TokenString() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceToken) pulumi.StringOutput { return v.TokenString }).(pulumi.StringOutput)
}

// Service token UID (part of the token)
func (o ServiceTokenOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceToken) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

type ServiceTokenArrayOutput struct{ *pulumi.OutputState }

func (ServiceTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceToken)(nil)).Elem()
}

func (o ServiceTokenArrayOutput) ToServiceTokenArrayOutput() ServiceTokenArrayOutput {
	return o
}

func (o ServiceTokenArrayOutput) ToServiceTokenArrayOutputWithContext(ctx context.Context) ServiceTokenArrayOutput {
	return o
}

func (o ServiceTokenArrayOutput) Index(i pulumi.IntInput) ServiceTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceToken {
		return vs[0].([]*ServiceToken)[vs[1].(int)]
	}).(ServiceTokenOutput)
}

type ServiceTokenMapOutput struct{ *pulumi.OutputState }

func (ServiceTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceToken)(nil)).Elem()
}

func (o ServiceTokenMapOutput) ToServiceTokenMapOutput() ServiceTokenMapOutput {
	return o
}

func (o ServiceTokenMapOutput) ToServiceTokenMapOutputWithContext(ctx context.Context) ServiceTokenMapOutput {
	return o
}

func (o ServiceTokenMapOutput) MapIndex(k pulumi.StringInput) ServiceTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceToken {
		return vs[0].(map[string]*ServiceToken)[vs[1].(string)]
	}).(ServiceTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTokenInput)(nil)).Elem(), &ServiceToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTokenArrayInput)(nil)).Elem(), ServiceTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTokenMapInput)(nil)).Elem(), ServiceTokenMap{})
	pulumi.RegisterOutputType(ServiceTokenOutput{})
	pulumi.RegisterOutputType(ServiceTokenArrayOutput{})
	pulumi.RegisterOutputType(ServiceTokenMapOutput{})
}

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

type LegacyPostgresCredential struct {
	pulumi.CustomResourceState

	// The system Postgres/Redshift/AlloyDB credential ID
	CredentialId pulumi.IntOutput `pulumi:"credentialId"`
	// Default schema name
	DefaultSchema pulumi.StringOutput `pulumi:"defaultSchema"`
	// Whether the Postgres/Redshift/AlloyDB credential is active
	IsActive pulumi.BoolPtrOutput `pulumi:"isActive"`
	// Number of threads to use
	NumThreads pulumi.IntPtrOutput `pulumi:"numThreads"`
	// Password for Postgres/Redshift/AlloyDB
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Project ID to create the Postgres/Redshift/AlloyDB credential in
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// Default schema name
	TargetName pulumi.StringPtrOutput `pulumi:"targetName"`
	// Type of connection. One of (postgres/redshift). Use postgres for alloydb connections
	Type pulumi.StringOutput `pulumi:"type"`
	// Username for Postgres/Redshift/AlloyDB
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewLegacyPostgresCredential registers a new resource with the given unique name, arguments, and options.
func NewLegacyPostgresCredential(ctx *pulumi.Context,
	name string, args *LegacyPostgresCredentialArgs, opts ...pulumi.ResourceOption) (*LegacyPostgresCredential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultSchema == nil {
		return nil, errors.New("invalid value for required argument 'DefaultSchema'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LegacyPostgresCredential
	err := ctx.RegisterResource("dbtcloud:index/legacyPostgresCredential:LegacyPostgresCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLegacyPostgresCredential gets an existing LegacyPostgresCredential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLegacyPostgresCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LegacyPostgresCredentialState, opts ...pulumi.ResourceOption) (*LegacyPostgresCredential, error) {
	var resource LegacyPostgresCredential
	err := ctx.ReadResource("dbtcloud:index/legacyPostgresCredential:LegacyPostgresCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LegacyPostgresCredential resources.
type legacyPostgresCredentialState struct {
	// The system Postgres/Redshift/AlloyDB credential ID
	CredentialId *int `pulumi:"credentialId"`
	// Default schema name
	DefaultSchema *string `pulumi:"defaultSchema"`
	// Whether the Postgres/Redshift/AlloyDB credential is active
	IsActive *bool `pulumi:"isActive"`
	// Number of threads to use
	NumThreads *int `pulumi:"numThreads"`
	// Password for Postgres/Redshift/AlloyDB
	Password *string `pulumi:"password"`
	// Project ID to create the Postgres/Redshift/AlloyDB credential in
	ProjectId *int `pulumi:"projectId"`
	// Default schema name
	TargetName *string `pulumi:"targetName"`
	// Type of connection. One of (postgres/redshift). Use postgres for alloydb connections
	Type *string `pulumi:"type"`
	// Username for Postgres/Redshift/AlloyDB
	Username *string `pulumi:"username"`
}

type LegacyPostgresCredentialState struct {
	// The system Postgres/Redshift/AlloyDB credential ID
	CredentialId pulumi.IntPtrInput
	// Default schema name
	DefaultSchema pulumi.StringPtrInput
	// Whether the Postgres/Redshift/AlloyDB credential is active
	IsActive pulumi.BoolPtrInput
	// Number of threads to use
	NumThreads pulumi.IntPtrInput
	// Password for Postgres/Redshift/AlloyDB
	Password pulumi.StringPtrInput
	// Project ID to create the Postgres/Redshift/AlloyDB credential in
	ProjectId pulumi.IntPtrInput
	// Default schema name
	TargetName pulumi.StringPtrInput
	// Type of connection. One of (postgres/redshift). Use postgres for alloydb connections
	Type pulumi.StringPtrInput
	// Username for Postgres/Redshift/AlloyDB
	Username pulumi.StringPtrInput
}

func (LegacyPostgresCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*legacyPostgresCredentialState)(nil)).Elem()
}

type legacyPostgresCredentialArgs struct {
	// Default schema name
	DefaultSchema string `pulumi:"defaultSchema"`
	// Whether the Postgres/Redshift/AlloyDB credential is active
	IsActive *bool `pulumi:"isActive"`
	// Number of threads to use
	NumThreads *int `pulumi:"numThreads"`
	// Password for Postgres/Redshift/AlloyDB
	Password *string `pulumi:"password"`
	// Project ID to create the Postgres/Redshift/AlloyDB credential in
	ProjectId int `pulumi:"projectId"`
	// Default schema name
	TargetName *string `pulumi:"targetName"`
	// Type of connection. One of (postgres/redshift). Use postgres for alloydb connections
	Type string `pulumi:"type"`
	// Username for Postgres/Redshift/AlloyDB
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a LegacyPostgresCredential resource.
type LegacyPostgresCredentialArgs struct {
	// Default schema name
	DefaultSchema pulumi.StringInput
	// Whether the Postgres/Redshift/AlloyDB credential is active
	IsActive pulumi.BoolPtrInput
	// Number of threads to use
	NumThreads pulumi.IntPtrInput
	// Password for Postgres/Redshift/AlloyDB
	Password pulumi.StringPtrInput
	// Project ID to create the Postgres/Redshift/AlloyDB credential in
	ProjectId pulumi.IntInput
	// Default schema name
	TargetName pulumi.StringPtrInput
	// Type of connection. One of (postgres/redshift). Use postgres for alloydb connections
	Type pulumi.StringInput
	// Username for Postgres/Redshift/AlloyDB
	Username pulumi.StringInput
}

func (LegacyPostgresCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*legacyPostgresCredentialArgs)(nil)).Elem()
}

type LegacyPostgresCredentialInput interface {
	pulumi.Input

	ToLegacyPostgresCredentialOutput() LegacyPostgresCredentialOutput
	ToLegacyPostgresCredentialOutputWithContext(ctx context.Context) LegacyPostgresCredentialOutput
}

func (*LegacyPostgresCredential) ElementType() reflect.Type {
	return reflect.TypeOf((**LegacyPostgresCredential)(nil)).Elem()
}

func (i *LegacyPostgresCredential) ToLegacyPostgresCredentialOutput() LegacyPostgresCredentialOutput {
	return i.ToLegacyPostgresCredentialOutputWithContext(context.Background())
}

func (i *LegacyPostgresCredential) ToLegacyPostgresCredentialOutputWithContext(ctx context.Context) LegacyPostgresCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegacyPostgresCredentialOutput)
}

// LegacyPostgresCredentialArrayInput is an input type that accepts LegacyPostgresCredentialArray and LegacyPostgresCredentialArrayOutput values.
// You can construct a concrete instance of `LegacyPostgresCredentialArrayInput` via:
//
//	LegacyPostgresCredentialArray{ LegacyPostgresCredentialArgs{...} }
type LegacyPostgresCredentialArrayInput interface {
	pulumi.Input

	ToLegacyPostgresCredentialArrayOutput() LegacyPostgresCredentialArrayOutput
	ToLegacyPostgresCredentialArrayOutputWithContext(context.Context) LegacyPostgresCredentialArrayOutput
}

type LegacyPostgresCredentialArray []LegacyPostgresCredentialInput

func (LegacyPostgresCredentialArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LegacyPostgresCredential)(nil)).Elem()
}

func (i LegacyPostgresCredentialArray) ToLegacyPostgresCredentialArrayOutput() LegacyPostgresCredentialArrayOutput {
	return i.ToLegacyPostgresCredentialArrayOutputWithContext(context.Background())
}

func (i LegacyPostgresCredentialArray) ToLegacyPostgresCredentialArrayOutputWithContext(ctx context.Context) LegacyPostgresCredentialArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegacyPostgresCredentialArrayOutput)
}

// LegacyPostgresCredentialMapInput is an input type that accepts LegacyPostgresCredentialMap and LegacyPostgresCredentialMapOutput values.
// You can construct a concrete instance of `LegacyPostgresCredentialMapInput` via:
//
//	LegacyPostgresCredentialMap{ "key": LegacyPostgresCredentialArgs{...} }
type LegacyPostgresCredentialMapInput interface {
	pulumi.Input

	ToLegacyPostgresCredentialMapOutput() LegacyPostgresCredentialMapOutput
	ToLegacyPostgresCredentialMapOutputWithContext(context.Context) LegacyPostgresCredentialMapOutput
}

type LegacyPostgresCredentialMap map[string]LegacyPostgresCredentialInput

func (LegacyPostgresCredentialMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LegacyPostgresCredential)(nil)).Elem()
}

func (i LegacyPostgresCredentialMap) ToLegacyPostgresCredentialMapOutput() LegacyPostgresCredentialMapOutput {
	return i.ToLegacyPostgresCredentialMapOutputWithContext(context.Background())
}

func (i LegacyPostgresCredentialMap) ToLegacyPostgresCredentialMapOutputWithContext(ctx context.Context) LegacyPostgresCredentialMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegacyPostgresCredentialMapOutput)
}

type LegacyPostgresCredentialOutput struct{ *pulumi.OutputState }

func (LegacyPostgresCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LegacyPostgresCredential)(nil)).Elem()
}

func (o LegacyPostgresCredentialOutput) ToLegacyPostgresCredentialOutput() LegacyPostgresCredentialOutput {
	return o
}

func (o LegacyPostgresCredentialOutput) ToLegacyPostgresCredentialOutputWithContext(ctx context.Context) LegacyPostgresCredentialOutput {
	return o
}

// The system Postgres/Redshift/AlloyDB credential ID
func (o LegacyPostgresCredentialOutput) CredentialId() pulumi.IntOutput {
	return o.ApplyT(func(v *LegacyPostgresCredential) pulumi.IntOutput { return v.CredentialId }).(pulumi.IntOutput)
}

// Default schema name
func (o LegacyPostgresCredentialOutput) DefaultSchema() pulumi.StringOutput {
	return o.ApplyT(func(v *LegacyPostgresCredential) pulumi.StringOutput { return v.DefaultSchema }).(pulumi.StringOutput)
}

// Whether the Postgres/Redshift/AlloyDB credential is active
func (o LegacyPostgresCredentialOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LegacyPostgresCredential) pulumi.BoolPtrOutput { return v.IsActive }).(pulumi.BoolPtrOutput)
}

// Number of threads to use
func (o LegacyPostgresCredentialOutput) NumThreads() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LegacyPostgresCredential) pulumi.IntPtrOutput { return v.NumThreads }).(pulumi.IntPtrOutput)
}

// Password for Postgres/Redshift/AlloyDB
func (o LegacyPostgresCredentialOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LegacyPostgresCredential) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Project ID to create the Postgres/Redshift/AlloyDB credential in
func (o LegacyPostgresCredentialOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *LegacyPostgresCredential) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

// Default schema name
func (o LegacyPostgresCredentialOutput) TargetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LegacyPostgresCredential) pulumi.StringPtrOutput { return v.TargetName }).(pulumi.StringPtrOutput)
}

// Type of connection. One of (postgres/redshift). Use postgres for alloydb connections
func (o LegacyPostgresCredentialOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LegacyPostgresCredential) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Username for Postgres/Redshift/AlloyDB
func (o LegacyPostgresCredentialOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *LegacyPostgresCredential) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type LegacyPostgresCredentialArrayOutput struct{ *pulumi.OutputState }

func (LegacyPostgresCredentialArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LegacyPostgresCredential)(nil)).Elem()
}

func (o LegacyPostgresCredentialArrayOutput) ToLegacyPostgresCredentialArrayOutput() LegacyPostgresCredentialArrayOutput {
	return o
}

func (o LegacyPostgresCredentialArrayOutput) ToLegacyPostgresCredentialArrayOutputWithContext(ctx context.Context) LegacyPostgresCredentialArrayOutput {
	return o
}

func (o LegacyPostgresCredentialArrayOutput) Index(i pulumi.IntInput) LegacyPostgresCredentialOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LegacyPostgresCredential {
		return vs[0].([]*LegacyPostgresCredential)[vs[1].(int)]
	}).(LegacyPostgresCredentialOutput)
}

type LegacyPostgresCredentialMapOutput struct{ *pulumi.OutputState }

func (LegacyPostgresCredentialMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LegacyPostgresCredential)(nil)).Elem()
}

func (o LegacyPostgresCredentialMapOutput) ToLegacyPostgresCredentialMapOutput() LegacyPostgresCredentialMapOutput {
	return o
}

func (o LegacyPostgresCredentialMapOutput) ToLegacyPostgresCredentialMapOutputWithContext(ctx context.Context) LegacyPostgresCredentialMapOutput {
	return o
}

func (o LegacyPostgresCredentialMapOutput) MapIndex(k pulumi.StringInput) LegacyPostgresCredentialOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LegacyPostgresCredential {
		return vs[0].(map[string]*LegacyPostgresCredential)[vs[1].(string)]
	}).(LegacyPostgresCredentialOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LegacyPostgresCredentialInput)(nil)).Elem(), &LegacyPostgresCredential{})
	pulumi.RegisterInputType(reflect.TypeOf((*LegacyPostgresCredentialArrayInput)(nil)).Elem(), LegacyPostgresCredentialArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LegacyPostgresCredentialMapInput)(nil)).Elem(), LegacyPostgresCredentialMap{})
	pulumi.RegisterOutputType(LegacyPostgresCredentialOutput{})
	pulumi.RegisterOutputType(LegacyPostgresCredentialArrayOutput{})
	pulumi.RegisterOutputType(LegacyPostgresCredentialMapOutput{})
}
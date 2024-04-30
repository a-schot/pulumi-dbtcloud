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
//			// NOTE for customers using the LEGACY dbt_cloud provider:
//			// when using the Databricks adapter
//			_, err := dbtcloud.NewDatabricksCredential(ctx, "myDatabricksCred", &dbtcloud.DatabricksCredentialArgs{
//				ProjectId:   pulumi.Any(dbtcloud_project.Dbt_project.Id),
//				AdapterId:   pulumi.Any(dbtcloud_connection.My_databricks_connection.Adapter_id),
//				TargetName:  pulumi.String("prod"),
//				Token:       pulumi.String("abcdefgh"),
//				Schema:      pulumi.String("my_schema"),
//				AdapterType: pulumi.String("databricks"),
//			})
//			if err != nil {
//				return err
//			}
//			// when using the Spark adapter
//			_, err = dbtcloud.NewDatabricksCredential(ctx, "mySparkCred", &dbtcloud.DatabricksCredentialArgs{
//				ProjectId:   pulumi.Any(dbtcloud_project.Dbt_project.Id),
//				AdapterId:   pulumi.Any(dbtcloud_connection.My_databricks_connection.Adapter_id),
//				TargetName:  pulumi.String("prod"),
//				Token:       pulumi.String("abcdefgh"),
//				Schema:      pulumi.String("my_schema"),
//				AdapterType: pulumi.String("spark"),
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
// Import using a project ID and credential ID found in the URL or via the API.
//
// ```sh
// $ pulumi import dbtcloud:index/databricksCredential:DatabricksCredential my_databricks_credential "project_id:credential_id"
// ```
//
// ```sh
// $ pulumi import dbtcloud:index/databricksCredential:DatabricksCredential my_databricks_credential 12345:6789
// ```
type DatabricksCredential struct {
	pulumi.CustomResourceState

	// Databricks adapter ID for the credential
	AdapterId pulumi.IntOutput `pulumi:"adapterId"`
	// The type of the adapter (databricks or spark)
	AdapterType pulumi.StringOutput `pulumi:"adapterType"`
	// The catalog where to create models (only for the databricks adapter)
	Catalog pulumi.StringPtrOutput `pulumi:"catalog"`
	// The system Databricks credential ID
	CredentialId pulumi.IntOutput `pulumi:"credentialId"`
	// Project ID to create the Databricks credential in
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// The schema where to create models
	Schema pulumi.StringOutput `pulumi:"schema"`
	// Target name
	TargetName pulumi.StringPtrOutput `pulumi:"targetName"`
	// Token for Databricks user
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewDatabricksCredential registers a new resource with the given unique name, arguments, and options.
func NewDatabricksCredential(ctx *pulumi.Context,
	name string, args *DatabricksCredentialArgs, opts ...pulumi.ResourceOption) (*DatabricksCredential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdapterId == nil {
		return nil, errors.New("invalid value for required argument 'AdapterId'")
	}
	if args.AdapterType == nil {
		return nil, errors.New("invalid value for required argument 'AdapterType'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Schema == nil {
		return nil, errors.New("invalid value for required argument 'Schema'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatabricksCredential
	err := ctx.RegisterResource("dbtcloud:index/databricksCredential:DatabricksCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabricksCredential gets an existing DatabricksCredential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabricksCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabricksCredentialState, opts ...pulumi.ResourceOption) (*DatabricksCredential, error) {
	var resource DatabricksCredential
	err := ctx.ReadResource("dbtcloud:index/databricksCredential:DatabricksCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabricksCredential resources.
type databricksCredentialState struct {
	// Databricks adapter ID for the credential
	AdapterId *int `pulumi:"adapterId"`
	// The type of the adapter (databricks or spark)
	AdapterType *string `pulumi:"adapterType"`
	// The catalog where to create models (only for the databricks adapter)
	Catalog *string `pulumi:"catalog"`
	// The system Databricks credential ID
	CredentialId *int `pulumi:"credentialId"`
	// Project ID to create the Databricks credential in
	ProjectId *int `pulumi:"projectId"`
	// The schema where to create models
	Schema *string `pulumi:"schema"`
	// Target name
	TargetName *string `pulumi:"targetName"`
	// Token for Databricks user
	Token *string `pulumi:"token"`
}

type DatabricksCredentialState struct {
	// Databricks adapter ID for the credential
	AdapterId pulumi.IntPtrInput
	// The type of the adapter (databricks or spark)
	AdapterType pulumi.StringPtrInput
	// The catalog where to create models (only for the databricks adapter)
	Catalog pulumi.StringPtrInput
	// The system Databricks credential ID
	CredentialId pulumi.IntPtrInput
	// Project ID to create the Databricks credential in
	ProjectId pulumi.IntPtrInput
	// The schema where to create models
	Schema pulumi.StringPtrInput
	// Target name
	TargetName pulumi.StringPtrInput
	// Token for Databricks user
	Token pulumi.StringPtrInput
}

func (DatabricksCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*databricksCredentialState)(nil)).Elem()
}

type databricksCredentialArgs struct {
	// Databricks adapter ID for the credential
	AdapterId int `pulumi:"adapterId"`
	// The type of the adapter (databricks or spark)
	AdapterType string `pulumi:"adapterType"`
	// The catalog where to create models (only for the databricks adapter)
	Catalog *string `pulumi:"catalog"`
	// Project ID to create the Databricks credential in
	ProjectId int `pulumi:"projectId"`
	// The schema where to create models
	Schema string `pulumi:"schema"`
	// Target name
	TargetName *string `pulumi:"targetName"`
	// Token for Databricks user
	Token string `pulumi:"token"`
}

// The set of arguments for constructing a DatabricksCredential resource.
type DatabricksCredentialArgs struct {
	// Databricks adapter ID for the credential
	AdapterId pulumi.IntInput
	// The type of the adapter (databricks or spark)
	AdapterType pulumi.StringInput
	// The catalog where to create models (only for the databricks adapter)
	Catalog pulumi.StringPtrInput
	// Project ID to create the Databricks credential in
	ProjectId pulumi.IntInput
	// The schema where to create models
	Schema pulumi.StringInput
	// Target name
	TargetName pulumi.StringPtrInput
	// Token for Databricks user
	Token pulumi.StringInput
}

func (DatabricksCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databricksCredentialArgs)(nil)).Elem()
}

type DatabricksCredentialInput interface {
	pulumi.Input

	ToDatabricksCredentialOutput() DatabricksCredentialOutput
	ToDatabricksCredentialOutputWithContext(ctx context.Context) DatabricksCredentialOutput
}

func (*DatabricksCredential) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabricksCredential)(nil)).Elem()
}

func (i *DatabricksCredential) ToDatabricksCredentialOutput() DatabricksCredentialOutput {
	return i.ToDatabricksCredentialOutputWithContext(context.Background())
}

func (i *DatabricksCredential) ToDatabricksCredentialOutputWithContext(ctx context.Context) DatabricksCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksCredentialOutput)
}

// DatabricksCredentialArrayInput is an input type that accepts DatabricksCredentialArray and DatabricksCredentialArrayOutput values.
// You can construct a concrete instance of `DatabricksCredentialArrayInput` via:
//
//	DatabricksCredentialArray{ DatabricksCredentialArgs{...} }
type DatabricksCredentialArrayInput interface {
	pulumi.Input

	ToDatabricksCredentialArrayOutput() DatabricksCredentialArrayOutput
	ToDatabricksCredentialArrayOutputWithContext(context.Context) DatabricksCredentialArrayOutput
}

type DatabricksCredentialArray []DatabricksCredentialInput

func (DatabricksCredentialArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabricksCredential)(nil)).Elem()
}

func (i DatabricksCredentialArray) ToDatabricksCredentialArrayOutput() DatabricksCredentialArrayOutput {
	return i.ToDatabricksCredentialArrayOutputWithContext(context.Background())
}

func (i DatabricksCredentialArray) ToDatabricksCredentialArrayOutputWithContext(ctx context.Context) DatabricksCredentialArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksCredentialArrayOutput)
}

// DatabricksCredentialMapInput is an input type that accepts DatabricksCredentialMap and DatabricksCredentialMapOutput values.
// You can construct a concrete instance of `DatabricksCredentialMapInput` via:
//
//	DatabricksCredentialMap{ "key": DatabricksCredentialArgs{...} }
type DatabricksCredentialMapInput interface {
	pulumi.Input

	ToDatabricksCredentialMapOutput() DatabricksCredentialMapOutput
	ToDatabricksCredentialMapOutputWithContext(context.Context) DatabricksCredentialMapOutput
}

type DatabricksCredentialMap map[string]DatabricksCredentialInput

func (DatabricksCredentialMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabricksCredential)(nil)).Elem()
}

func (i DatabricksCredentialMap) ToDatabricksCredentialMapOutput() DatabricksCredentialMapOutput {
	return i.ToDatabricksCredentialMapOutputWithContext(context.Background())
}

func (i DatabricksCredentialMap) ToDatabricksCredentialMapOutputWithContext(ctx context.Context) DatabricksCredentialMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabricksCredentialMapOutput)
}

type DatabricksCredentialOutput struct{ *pulumi.OutputState }

func (DatabricksCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabricksCredential)(nil)).Elem()
}

func (o DatabricksCredentialOutput) ToDatabricksCredentialOutput() DatabricksCredentialOutput {
	return o
}

func (o DatabricksCredentialOutput) ToDatabricksCredentialOutputWithContext(ctx context.Context) DatabricksCredentialOutput {
	return o
}

// Databricks adapter ID for the credential
func (o DatabricksCredentialOutput) AdapterId() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabricksCredential) pulumi.IntOutput { return v.AdapterId }).(pulumi.IntOutput)
}

// The type of the adapter (databricks or spark)
func (o DatabricksCredentialOutput) AdapterType() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabricksCredential) pulumi.StringOutput { return v.AdapterType }).(pulumi.StringOutput)
}

// The catalog where to create models (only for the databricks adapter)
func (o DatabricksCredentialOutput) Catalog() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabricksCredential) pulumi.StringPtrOutput { return v.Catalog }).(pulumi.StringPtrOutput)
}

// The system Databricks credential ID
func (o DatabricksCredentialOutput) CredentialId() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabricksCredential) pulumi.IntOutput { return v.CredentialId }).(pulumi.IntOutput)
}

// Project ID to create the Databricks credential in
func (o DatabricksCredentialOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabricksCredential) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

// The schema where to create models
func (o DatabricksCredentialOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabricksCredential) pulumi.StringOutput { return v.Schema }).(pulumi.StringOutput)
}

// Target name
func (o DatabricksCredentialOutput) TargetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabricksCredential) pulumi.StringPtrOutput { return v.TargetName }).(pulumi.StringPtrOutput)
}

// Token for Databricks user
func (o DatabricksCredentialOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabricksCredential) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

type DatabricksCredentialArrayOutput struct{ *pulumi.OutputState }

func (DatabricksCredentialArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabricksCredential)(nil)).Elem()
}

func (o DatabricksCredentialArrayOutput) ToDatabricksCredentialArrayOutput() DatabricksCredentialArrayOutput {
	return o
}

func (o DatabricksCredentialArrayOutput) ToDatabricksCredentialArrayOutputWithContext(ctx context.Context) DatabricksCredentialArrayOutput {
	return o
}

func (o DatabricksCredentialArrayOutput) Index(i pulumi.IntInput) DatabricksCredentialOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabricksCredential {
		return vs[0].([]*DatabricksCredential)[vs[1].(int)]
	}).(DatabricksCredentialOutput)
}

type DatabricksCredentialMapOutput struct{ *pulumi.OutputState }

func (DatabricksCredentialMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabricksCredential)(nil)).Elem()
}

func (o DatabricksCredentialMapOutput) ToDatabricksCredentialMapOutput() DatabricksCredentialMapOutput {
	return o
}

func (o DatabricksCredentialMapOutput) ToDatabricksCredentialMapOutputWithContext(ctx context.Context) DatabricksCredentialMapOutput {
	return o
}

func (o DatabricksCredentialMapOutput) MapIndex(k pulumi.StringInput) DatabricksCredentialOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabricksCredential {
		return vs[0].(map[string]*DatabricksCredential)[vs[1].(string)]
	}).(DatabricksCredentialOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabricksCredentialInput)(nil)).Elem(), &DatabricksCredential{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabricksCredentialArrayInput)(nil)).Elem(), DatabricksCredentialArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabricksCredentialMapInput)(nil)).Elem(), DatabricksCredentialMap{})
	pulumi.RegisterOutputType(DatabricksCredentialOutput{})
	pulumi.RegisterOutputType(DatabricksCredentialArrayOutput{})
	pulumi.RegisterOutputType(DatabricksCredentialMapOutput{})
}

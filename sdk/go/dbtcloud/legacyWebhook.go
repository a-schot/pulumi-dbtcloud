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

// Deprecated: Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource
type LegacyWebhook struct {
	pulumi.CustomResourceState

	// Webhooks Account Identifier
	AccountIdentifier pulumi.StringOutput `pulumi:"accountIdentifier"`
	// Webhooks active flag
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// Webhooks Client URL
	ClientUrl pulumi.StringOutput `pulumi:"clientUrl"`
	// Webhooks Description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Webhooks Event Types
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// Secret key for the webhook. Can be used to validate the authenticity of the webhook.
	HmacSecret pulumi.StringOutput `pulumi:"hmacSecret"`
	// Latest HTTP status of the webhook
	HttpStatusCode pulumi.StringOutput `pulumi:"httpStatusCode"`
	// List of job IDs to trigger the webhook, An empty list will trigger on all jobs
	JobIds pulumi.IntArrayOutput `pulumi:"jobIds"`
	// Webhooks Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Webhooks ID
	WebhookId pulumi.StringOutput `pulumi:"webhookId"`
}

// NewLegacyWebhook registers a new resource with the given unique name, arguments, and options.
func NewLegacyWebhook(ctx *pulumi.Context,
	name string, args *LegacyWebhookArgs, opts ...pulumi.ResourceOption) (*LegacyWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientUrl == nil {
		return nil, errors.New("invalid value for required argument 'ClientUrl'")
	}
	if args.EventTypes == nil {
		return nil, errors.New("invalid value for required argument 'EventTypes'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"hmacSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LegacyWebhook
	err := ctx.RegisterResource("dbtcloud:index/legacyWebhook:LegacyWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLegacyWebhook gets an existing LegacyWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLegacyWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LegacyWebhookState, opts ...pulumi.ResourceOption) (*LegacyWebhook, error) {
	var resource LegacyWebhook
	err := ctx.ReadResource("dbtcloud:index/legacyWebhook:LegacyWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LegacyWebhook resources.
type legacyWebhookState struct {
	// Webhooks Account Identifier
	AccountIdentifier *string `pulumi:"accountIdentifier"`
	// Webhooks active flag
	Active *bool `pulumi:"active"`
	// Webhooks Client URL
	ClientUrl *string `pulumi:"clientUrl"`
	// Webhooks Description
	Description *string `pulumi:"description"`
	// Webhooks Event Types
	EventTypes []string `pulumi:"eventTypes"`
	// Secret key for the webhook. Can be used to validate the authenticity of the webhook.
	HmacSecret *string `pulumi:"hmacSecret"`
	// Latest HTTP status of the webhook
	HttpStatusCode *string `pulumi:"httpStatusCode"`
	// List of job IDs to trigger the webhook, An empty list will trigger on all jobs
	JobIds []int `pulumi:"jobIds"`
	// Webhooks Name
	Name *string `pulumi:"name"`
	// Webhooks ID
	WebhookId *string `pulumi:"webhookId"`
}

type LegacyWebhookState struct {
	// Webhooks Account Identifier
	AccountIdentifier pulumi.StringPtrInput
	// Webhooks active flag
	Active pulumi.BoolPtrInput
	// Webhooks Client URL
	ClientUrl pulumi.StringPtrInput
	// Webhooks Description
	Description pulumi.StringPtrInput
	// Webhooks Event Types
	EventTypes pulumi.StringArrayInput
	// Secret key for the webhook. Can be used to validate the authenticity of the webhook.
	HmacSecret pulumi.StringPtrInput
	// Latest HTTP status of the webhook
	HttpStatusCode pulumi.StringPtrInput
	// List of job IDs to trigger the webhook, An empty list will trigger on all jobs
	JobIds pulumi.IntArrayInput
	// Webhooks Name
	Name pulumi.StringPtrInput
	// Webhooks ID
	WebhookId pulumi.StringPtrInput
}

func (LegacyWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*legacyWebhookState)(nil)).Elem()
}

type legacyWebhookArgs struct {
	// Webhooks active flag
	Active *bool `pulumi:"active"`
	// Webhooks Client URL
	ClientUrl string `pulumi:"clientUrl"`
	// Webhooks Description
	Description *string `pulumi:"description"`
	// Webhooks Event Types
	EventTypes []string `pulumi:"eventTypes"`
	// List of job IDs to trigger the webhook, An empty list will trigger on all jobs
	JobIds []int `pulumi:"jobIds"`
	// Webhooks Name
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a LegacyWebhook resource.
type LegacyWebhookArgs struct {
	// Webhooks active flag
	Active pulumi.BoolPtrInput
	// Webhooks Client URL
	ClientUrl pulumi.StringInput
	// Webhooks Description
	Description pulumi.StringPtrInput
	// Webhooks Event Types
	EventTypes pulumi.StringArrayInput
	// List of job IDs to trigger the webhook, An empty list will trigger on all jobs
	JobIds pulumi.IntArrayInput
	// Webhooks Name
	Name pulumi.StringPtrInput
}

func (LegacyWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*legacyWebhookArgs)(nil)).Elem()
}

type LegacyWebhookInput interface {
	pulumi.Input

	ToLegacyWebhookOutput() LegacyWebhookOutput
	ToLegacyWebhookOutputWithContext(ctx context.Context) LegacyWebhookOutput
}

func (*LegacyWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**LegacyWebhook)(nil)).Elem()
}

func (i *LegacyWebhook) ToLegacyWebhookOutput() LegacyWebhookOutput {
	return i.ToLegacyWebhookOutputWithContext(context.Background())
}

func (i *LegacyWebhook) ToLegacyWebhookOutputWithContext(ctx context.Context) LegacyWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegacyWebhookOutput)
}

// LegacyWebhookArrayInput is an input type that accepts LegacyWebhookArray and LegacyWebhookArrayOutput values.
// You can construct a concrete instance of `LegacyWebhookArrayInput` via:
//
//	LegacyWebhookArray{ LegacyWebhookArgs{...} }
type LegacyWebhookArrayInput interface {
	pulumi.Input

	ToLegacyWebhookArrayOutput() LegacyWebhookArrayOutput
	ToLegacyWebhookArrayOutputWithContext(context.Context) LegacyWebhookArrayOutput
}

type LegacyWebhookArray []LegacyWebhookInput

func (LegacyWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LegacyWebhook)(nil)).Elem()
}

func (i LegacyWebhookArray) ToLegacyWebhookArrayOutput() LegacyWebhookArrayOutput {
	return i.ToLegacyWebhookArrayOutputWithContext(context.Background())
}

func (i LegacyWebhookArray) ToLegacyWebhookArrayOutputWithContext(ctx context.Context) LegacyWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegacyWebhookArrayOutput)
}

// LegacyWebhookMapInput is an input type that accepts LegacyWebhookMap and LegacyWebhookMapOutput values.
// You can construct a concrete instance of `LegacyWebhookMapInput` via:
//
//	LegacyWebhookMap{ "key": LegacyWebhookArgs{...} }
type LegacyWebhookMapInput interface {
	pulumi.Input

	ToLegacyWebhookMapOutput() LegacyWebhookMapOutput
	ToLegacyWebhookMapOutputWithContext(context.Context) LegacyWebhookMapOutput
}

type LegacyWebhookMap map[string]LegacyWebhookInput

func (LegacyWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LegacyWebhook)(nil)).Elem()
}

func (i LegacyWebhookMap) ToLegacyWebhookMapOutput() LegacyWebhookMapOutput {
	return i.ToLegacyWebhookMapOutputWithContext(context.Background())
}

func (i LegacyWebhookMap) ToLegacyWebhookMapOutputWithContext(ctx context.Context) LegacyWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegacyWebhookMapOutput)
}

type LegacyWebhookOutput struct{ *pulumi.OutputState }

func (LegacyWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LegacyWebhook)(nil)).Elem()
}

func (o LegacyWebhookOutput) ToLegacyWebhookOutput() LegacyWebhookOutput {
	return o
}

func (o LegacyWebhookOutput) ToLegacyWebhookOutputWithContext(ctx context.Context) LegacyWebhookOutput {
	return o
}

// Webhooks Account Identifier
func (o LegacyWebhookOutput) AccountIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *LegacyWebhook) pulumi.StringOutput { return v.AccountIdentifier }).(pulumi.StringOutput)
}

// Webhooks active flag
func (o LegacyWebhookOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LegacyWebhook) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

// Webhooks Client URL
func (o LegacyWebhookOutput) ClientUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *LegacyWebhook) pulumi.StringOutput { return v.ClientUrl }).(pulumi.StringOutput)
}

// Webhooks Description
func (o LegacyWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LegacyWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Webhooks Event Types
func (o LegacyWebhookOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LegacyWebhook) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// Secret key for the webhook. Can be used to validate the authenticity of the webhook.
func (o LegacyWebhookOutput) HmacSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *LegacyWebhook) pulumi.StringOutput { return v.HmacSecret }).(pulumi.StringOutput)
}

// Latest HTTP status of the webhook
func (o LegacyWebhookOutput) HttpStatusCode() pulumi.StringOutput {
	return o.ApplyT(func(v *LegacyWebhook) pulumi.StringOutput { return v.HttpStatusCode }).(pulumi.StringOutput)
}

// List of job IDs to trigger the webhook, An empty list will trigger on all jobs
func (o LegacyWebhookOutput) JobIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *LegacyWebhook) pulumi.IntArrayOutput { return v.JobIds }).(pulumi.IntArrayOutput)
}

// Webhooks Name
func (o LegacyWebhookOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LegacyWebhook) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Webhooks ID
func (o LegacyWebhookOutput) WebhookId() pulumi.StringOutput {
	return o.ApplyT(func(v *LegacyWebhook) pulumi.StringOutput { return v.WebhookId }).(pulumi.StringOutput)
}

type LegacyWebhookArrayOutput struct{ *pulumi.OutputState }

func (LegacyWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LegacyWebhook)(nil)).Elem()
}

func (o LegacyWebhookArrayOutput) ToLegacyWebhookArrayOutput() LegacyWebhookArrayOutput {
	return o
}

func (o LegacyWebhookArrayOutput) ToLegacyWebhookArrayOutputWithContext(ctx context.Context) LegacyWebhookArrayOutput {
	return o
}

func (o LegacyWebhookArrayOutput) Index(i pulumi.IntInput) LegacyWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LegacyWebhook {
		return vs[0].([]*LegacyWebhook)[vs[1].(int)]
	}).(LegacyWebhookOutput)
}

type LegacyWebhookMapOutput struct{ *pulumi.OutputState }

func (LegacyWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LegacyWebhook)(nil)).Elem()
}

func (o LegacyWebhookMapOutput) ToLegacyWebhookMapOutput() LegacyWebhookMapOutput {
	return o
}

func (o LegacyWebhookMapOutput) ToLegacyWebhookMapOutputWithContext(ctx context.Context) LegacyWebhookMapOutput {
	return o
}

func (o LegacyWebhookMapOutput) MapIndex(k pulumi.StringInput) LegacyWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LegacyWebhook {
		return vs[0].(map[string]*LegacyWebhook)[vs[1].(string)]
	}).(LegacyWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LegacyWebhookInput)(nil)).Elem(), &LegacyWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*LegacyWebhookArrayInput)(nil)).Elem(), LegacyWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LegacyWebhookMapInput)(nil)).Elem(), LegacyWebhookMap{})
	pulumi.RegisterOutputType(LegacyWebhookOutput{})
	pulumi.RegisterOutputType(LegacyWebhookArrayOutput{})
	pulumi.RegisterOutputType(LegacyWebhookMapOutput{})
}

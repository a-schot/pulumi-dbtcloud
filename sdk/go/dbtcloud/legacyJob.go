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
type LegacyJob struct {
	pulumi.CustomResourceState

	// Version number of dbt to use in this job, usually in the format 1.2.0-latest rather than core versions
	DbtVersion pulumi.StringPtrOutput `pulumi:"dbtVersion"`
	// Environment identifier that this job defers to (new deferring approach)
	DeferringEnvironmentId pulumi.IntPtrOutput `pulumi:"deferringEnvironmentId"`
	// Job identifier that this job defers to (legacy deferring approach)
	DeferringJobId pulumi.IntPtrOutput `pulumi:"deferringJobId"`
	// Description for the job
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Environment ID to create the job in
	EnvironmentId pulumi.IntOutput `pulumi:"environmentId"`
	// List of commands to execute for the job
	ExecuteSteps pulumi.StringArrayOutput `pulumi:"executeSteps"`
	// Flag for whether the job should generate documentation
	GenerateDocs pulumi.BoolPtrOutput `pulumi:"generateDocs"`
	// Flag for whether the job is marked active or deleted. To create/keep a job in a 'deactivated' state, check the
	// `triggers` config.
	IsActive pulumi.BoolPtrOutput `pulumi:"isActive"`
	// Which other job should trigger this job when it finishes, and on which conditions (sometimes referred as 'job
	// chaining').
	JobCompletionTriggerCondition LegacyJobJobCompletionTriggerConditionPtrOutput `pulumi:"jobCompletionTriggerCondition"`
	// Job name
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of threads to use in the job
	NumThreads pulumi.IntPtrOutput `pulumi:"numThreads"`
	// Project ID to create the job in
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// Flag for whether the job should add a `dbt source freshness` step to the job. The difference between manually adding a
	// step with `dbt source freshness` in the job steps or using this flag is that with this flag, a failed freshness will
	// still allow the following steps to run.
	RunGenerateSources pulumi.BoolPtrOutput `pulumi:"runGenerateSources"`
	// Custom cron expression for schedule
	ScheduleCron pulumi.StringPtrOutput `pulumi:"scheduleCron"`
	// List of days of week as numbers (0 = Sunday, 7 = Saturday) to execute the job at if running on a schedule
	ScheduleDays pulumi.IntArrayOutput `pulumi:"scheduleDays"`
	// List of hours to execute the job at if running on a schedule
	ScheduleHours pulumi.IntArrayOutput `pulumi:"scheduleHours"`
	// Number of hours between job executions if running on a schedule
	ScheduleInterval pulumi.IntPtrOutput `pulumi:"scheduleInterval"`
	// Type of schedule to use, one of every_day/ days_of_week/ custom_cron
	ScheduleType pulumi.StringPtrOutput `pulumi:"scheduleType"`
	// Whether this job defers on a previous run of itself
	SelfDeferring pulumi.BoolPtrOutput `pulumi:"selfDeferring"`
	// Target name for the dbt profile
	TargetName pulumi.StringPtrOutput `pulumi:"targetName"`
	// Number of seconds to allow the job to run before timing out
	TimeoutSeconds pulumi.IntPtrOutput `pulumi:"timeoutSeconds"`
	// Flags for which types of triggers to use, the values are `github_webhook`, `git_provider_webhook`, `schedule` and
	// `custom_branch_only`. <br>`custom_branch_only` is only relevant for CI jobs triggered automatically on PR creation to
	// only trigger a job on a PR to the custom branch of the environment. To create a job in a 'deactivated' state, set all to
	// `false`.
	Triggers pulumi.BoolMapOutput `pulumi:"triggers"`
	// Whether the CI job should be automatically triggered on draft PRs
	TriggersOnDraftPr pulumi.BoolPtrOutput `pulumi:"triggersOnDraftPr"`
}

// NewLegacyJob registers a new resource with the given unique name, arguments, and options.
func NewLegacyJob(ctx *pulumi.Context,
	name string, args *LegacyJobArgs, opts ...pulumi.ResourceOption) (*LegacyJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.ExecuteSteps == nil {
		return nil, errors.New("invalid value for required argument 'ExecuteSteps'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Triggers == nil {
		return nil, errors.New("invalid value for required argument 'Triggers'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LegacyJob
	err := ctx.RegisterResource("dbtcloud:index/legacyJob:LegacyJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLegacyJob gets an existing LegacyJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLegacyJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LegacyJobState, opts ...pulumi.ResourceOption) (*LegacyJob, error) {
	var resource LegacyJob
	err := ctx.ReadResource("dbtcloud:index/legacyJob:LegacyJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LegacyJob resources.
type legacyJobState struct {
	// Version number of dbt to use in this job, usually in the format 1.2.0-latest rather than core versions
	DbtVersion *string `pulumi:"dbtVersion"`
	// Environment identifier that this job defers to (new deferring approach)
	DeferringEnvironmentId *int `pulumi:"deferringEnvironmentId"`
	// Job identifier that this job defers to (legacy deferring approach)
	DeferringJobId *int `pulumi:"deferringJobId"`
	// Description for the job
	Description *string `pulumi:"description"`
	// Environment ID to create the job in
	EnvironmentId *int `pulumi:"environmentId"`
	// List of commands to execute for the job
	ExecuteSteps []string `pulumi:"executeSteps"`
	// Flag for whether the job should generate documentation
	GenerateDocs *bool `pulumi:"generateDocs"`
	// Flag for whether the job is marked active or deleted. To create/keep a job in a 'deactivated' state, check the
	// `triggers` config.
	IsActive *bool `pulumi:"isActive"`
	// Which other job should trigger this job when it finishes, and on which conditions (sometimes referred as 'job
	// chaining').
	JobCompletionTriggerCondition *LegacyJobJobCompletionTriggerCondition `pulumi:"jobCompletionTriggerCondition"`
	// Job name
	Name *string `pulumi:"name"`
	// Number of threads to use in the job
	NumThreads *int `pulumi:"numThreads"`
	// Project ID to create the job in
	ProjectId *int `pulumi:"projectId"`
	// Flag for whether the job should add a `dbt source freshness` step to the job. The difference between manually adding a
	// step with `dbt source freshness` in the job steps or using this flag is that with this flag, a failed freshness will
	// still allow the following steps to run.
	RunGenerateSources *bool `pulumi:"runGenerateSources"`
	// Custom cron expression for schedule
	ScheduleCron *string `pulumi:"scheduleCron"`
	// List of days of week as numbers (0 = Sunday, 7 = Saturday) to execute the job at if running on a schedule
	ScheduleDays []int `pulumi:"scheduleDays"`
	// List of hours to execute the job at if running on a schedule
	ScheduleHours []int `pulumi:"scheduleHours"`
	// Number of hours between job executions if running on a schedule
	ScheduleInterval *int `pulumi:"scheduleInterval"`
	// Type of schedule to use, one of every_day/ days_of_week/ custom_cron
	ScheduleType *string `pulumi:"scheduleType"`
	// Whether this job defers on a previous run of itself
	SelfDeferring *bool `pulumi:"selfDeferring"`
	// Target name for the dbt profile
	TargetName *string `pulumi:"targetName"`
	// Number of seconds to allow the job to run before timing out
	TimeoutSeconds *int `pulumi:"timeoutSeconds"`
	// Flags for which types of triggers to use, the values are `github_webhook`, `git_provider_webhook`, `schedule` and
	// `custom_branch_only`. <br>`custom_branch_only` is only relevant for CI jobs triggered automatically on PR creation to
	// only trigger a job on a PR to the custom branch of the environment. To create a job in a 'deactivated' state, set all to
	// `false`.
	Triggers map[string]bool `pulumi:"triggers"`
	// Whether the CI job should be automatically triggered on draft PRs
	TriggersOnDraftPr *bool `pulumi:"triggersOnDraftPr"`
}

type LegacyJobState struct {
	// Version number of dbt to use in this job, usually in the format 1.2.0-latest rather than core versions
	DbtVersion pulumi.StringPtrInput
	// Environment identifier that this job defers to (new deferring approach)
	DeferringEnvironmentId pulumi.IntPtrInput
	// Job identifier that this job defers to (legacy deferring approach)
	DeferringJobId pulumi.IntPtrInput
	// Description for the job
	Description pulumi.StringPtrInput
	// Environment ID to create the job in
	EnvironmentId pulumi.IntPtrInput
	// List of commands to execute for the job
	ExecuteSteps pulumi.StringArrayInput
	// Flag for whether the job should generate documentation
	GenerateDocs pulumi.BoolPtrInput
	// Flag for whether the job is marked active or deleted. To create/keep a job in a 'deactivated' state, check the
	// `triggers` config.
	IsActive pulumi.BoolPtrInput
	// Which other job should trigger this job when it finishes, and on which conditions (sometimes referred as 'job
	// chaining').
	JobCompletionTriggerCondition LegacyJobJobCompletionTriggerConditionPtrInput
	// Job name
	Name pulumi.StringPtrInput
	// Number of threads to use in the job
	NumThreads pulumi.IntPtrInput
	// Project ID to create the job in
	ProjectId pulumi.IntPtrInput
	// Flag for whether the job should add a `dbt source freshness` step to the job. The difference between manually adding a
	// step with `dbt source freshness` in the job steps or using this flag is that with this flag, a failed freshness will
	// still allow the following steps to run.
	RunGenerateSources pulumi.BoolPtrInput
	// Custom cron expression for schedule
	ScheduleCron pulumi.StringPtrInput
	// List of days of week as numbers (0 = Sunday, 7 = Saturday) to execute the job at if running on a schedule
	ScheduleDays pulumi.IntArrayInput
	// List of hours to execute the job at if running on a schedule
	ScheduleHours pulumi.IntArrayInput
	// Number of hours between job executions if running on a schedule
	ScheduleInterval pulumi.IntPtrInput
	// Type of schedule to use, one of every_day/ days_of_week/ custom_cron
	ScheduleType pulumi.StringPtrInput
	// Whether this job defers on a previous run of itself
	SelfDeferring pulumi.BoolPtrInput
	// Target name for the dbt profile
	TargetName pulumi.StringPtrInput
	// Number of seconds to allow the job to run before timing out
	TimeoutSeconds pulumi.IntPtrInput
	// Flags for which types of triggers to use, the values are `github_webhook`, `git_provider_webhook`, `schedule` and
	// `custom_branch_only`. <br>`custom_branch_only` is only relevant for CI jobs triggered automatically on PR creation to
	// only trigger a job on a PR to the custom branch of the environment. To create a job in a 'deactivated' state, set all to
	// `false`.
	Triggers pulumi.BoolMapInput
	// Whether the CI job should be automatically triggered on draft PRs
	TriggersOnDraftPr pulumi.BoolPtrInput
}

func (LegacyJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*legacyJobState)(nil)).Elem()
}

type legacyJobArgs struct {
	// Version number of dbt to use in this job, usually in the format 1.2.0-latest rather than core versions
	DbtVersion *string `pulumi:"dbtVersion"`
	// Environment identifier that this job defers to (new deferring approach)
	DeferringEnvironmentId *int `pulumi:"deferringEnvironmentId"`
	// Job identifier that this job defers to (legacy deferring approach)
	DeferringJobId *int `pulumi:"deferringJobId"`
	// Description for the job
	Description *string `pulumi:"description"`
	// Environment ID to create the job in
	EnvironmentId int `pulumi:"environmentId"`
	// List of commands to execute for the job
	ExecuteSteps []string `pulumi:"executeSteps"`
	// Flag for whether the job should generate documentation
	GenerateDocs *bool `pulumi:"generateDocs"`
	// Flag for whether the job is marked active or deleted. To create/keep a job in a 'deactivated' state, check the
	// `triggers` config.
	IsActive *bool `pulumi:"isActive"`
	// Which other job should trigger this job when it finishes, and on which conditions (sometimes referred as 'job
	// chaining').
	JobCompletionTriggerCondition *LegacyJobJobCompletionTriggerCondition `pulumi:"jobCompletionTriggerCondition"`
	// Job name
	Name *string `pulumi:"name"`
	// Number of threads to use in the job
	NumThreads *int `pulumi:"numThreads"`
	// Project ID to create the job in
	ProjectId int `pulumi:"projectId"`
	// Flag for whether the job should add a `dbt source freshness` step to the job. The difference between manually adding a
	// step with `dbt source freshness` in the job steps or using this flag is that with this flag, a failed freshness will
	// still allow the following steps to run.
	RunGenerateSources *bool `pulumi:"runGenerateSources"`
	// Custom cron expression for schedule
	ScheduleCron *string `pulumi:"scheduleCron"`
	// List of days of week as numbers (0 = Sunday, 7 = Saturday) to execute the job at if running on a schedule
	ScheduleDays []int `pulumi:"scheduleDays"`
	// List of hours to execute the job at if running on a schedule
	ScheduleHours []int `pulumi:"scheduleHours"`
	// Number of hours between job executions if running on a schedule
	ScheduleInterval *int `pulumi:"scheduleInterval"`
	// Type of schedule to use, one of every_day/ days_of_week/ custom_cron
	ScheduleType *string `pulumi:"scheduleType"`
	// Whether this job defers on a previous run of itself
	SelfDeferring *bool `pulumi:"selfDeferring"`
	// Target name for the dbt profile
	TargetName *string `pulumi:"targetName"`
	// Number of seconds to allow the job to run before timing out
	TimeoutSeconds *int `pulumi:"timeoutSeconds"`
	// Flags for which types of triggers to use, the values are `github_webhook`, `git_provider_webhook`, `schedule` and
	// `custom_branch_only`. <br>`custom_branch_only` is only relevant for CI jobs triggered automatically on PR creation to
	// only trigger a job on a PR to the custom branch of the environment. To create a job in a 'deactivated' state, set all to
	// `false`.
	Triggers map[string]bool `pulumi:"triggers"`
	// Whether the CI job should be automatically triggered on draft PRs
	TriggersOnDraftPr *bool `pulumi:"triggersOnDraftPr"`
}

// The set of arguments for constructing a LegacyJob resource.
type LegacyJobArgs struct {
	// Version number of dbt to use in this job, usually in the format 1.2.0-latest rather than core versions
	DbtVersion pulumi.StringPtrInput
	// Environment identifier that this job defers to (new deferring approach)
	DeferringEnvironmentId pulumi.IntPtrInput
	// Job identifier that this job defers to (legacy deferring approach)
	DeferringJobId pulumi.IntPtrInput
	// Description for the job
	Description pulumi.StringPtrInput
	// Environment ID to create the job in
	EnvironmentId pulumi.IntInput
	// List of commands to execute for the job
	ExecuteSteps pulumi.StringArrayInput
	// Flag for whether the job should generate documentation
	GenerateDocs pulumi.BoolPtrInput
	// Flag for whether the job is marked active or deleted. To create/keep a job in a 'deactivated' state, check the
	// `triggers` config.
	IsActive pulumi.BoolPtrInput
	// Which other job should trigger this job when it finishes, and on which conditions (sometimes referred as 'job
	// chaining').
	JobCompletionTriggerCondition LegacyJobJobCompletionTriggerConditionPtrInput
	// Job name
	Name pulumi.StringPtrInput
	// Number of threads to use in the job
	NumThreads pulumi.IntPtrInput
	// Project ID to create the job in
	ProjectId pulumi.IntInput
	// Flag for whether the job should add a `dbt source freshness` step to the job. The difference between manually adding a
	// step with `dbt source freshness` in the job steps or using this flag is that with this flag, a failed freshness will
	// still allow the following steps to run.
	RunGenerateSources pulumi.BoolPtrInput
	// Custom cron expression for schedule
	ScheduleCron pulumi.StringPtrInput
	// List of days of week as numbers (0 = Sunday, 7 = Saturday) to execute the job at if running on a schedule
	ScheduleDays pulumi.IntArrayInput
	// List of hours to execute the job at if running on a schedule
	ScheduleHours pulumi.IntArrayInput
	// Number of hours between job executions if running on a schedule
	ScheduleInterval pulumi.IntPtrInput
	// Type of schedule to use, one of every_day/ days_of_week/ custom_cron
	ScheduleType pulumi.StringPtrInput
	// Whether this job defers on a previous run of itself
	SelfDeferring pulumi.BoolPtrInput
	// Target name for the dbt profile
	TargetName pulumi.StringPtrInput
	// Number of seconds to allow the job to run before timing out
	TimeoutSeconds pulumi.IntPtrInput
	// Flags for which types of triggers to use, the values are `github_webhook`, `git_provider_webhook`, `schedule` and
	// `custom_branch_only`. <br>`custom_branch_only` is only relevant for CI jobs triggered automatically on PR creation to
	// only trigger a job on a PR to the custom branch of the environment. To create a job in a 'deactivated' state, set all to
	// `false`.
	Triggers pulumi.BoolMapInput
	// Whether the CI job should be automatically triggered on draft PRs
	TriggersOnDraftPr pulumi.BoolPtrInput
}

func (LegacyJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*legacyJobArgs)(nil)).Elem()
}

type LegacyJobInput interface {
	pulumi.Input

	ToLegacyJobOutput() LegacyJobOutput
	ToLegacyJobOutputWithContext(ctx context.Context) LegacyJobOutput
}

func (*LegacyJob) ElementType() reflect.Type {
	return reflect.TypeOf((**LegacyJob)(nil)).Elem()
}

func (i *LegacyJob) ToLegacyJobOutput() LegacyJobOutput {
	return i.ToLegacyJobOutputWithContext(context.Background())
}

func (i *LegacyJob) ToLegacyJobOutputWithContext(ctx context.Context) LegacyJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegacyJobOutput)
}

// LegacyJobArrayInput is an input type that accepts LegacyJobArray and LegacyJobArrayOutput values.
// You can construct a concrete instance of `LegacyJobArrayInput` via:
//
//	LegacyJobArray{ LegacyJobArgs{...} }
type LegacyJobArrayInput interface {
	pulumi.Input

	ToLegacyJobArrayOutput() LegacyJobArrayOutput
	ToLegacyJobArrayOutputWithContext(context.Context) LegacyJobArrayOutput
}

type LegacyJobArray []LegacyJobInput

func (LegacyJobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LegacyJob)(nil)).Elem()
}

func (i LegacyJobArray) ToLegacyJobArrayOutput() LegacyJobArrayOutput {
	return i.ToLegacyJobArrayOutputWithContext(context.Background())
}

func (i LegacyJobArray) ToLegacyJobArrayOutputWithContext(ctx context.Context) LegacyJobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegacyJobArrayOutput)
}

// LegacyJobMapInput is an input type that accepts LegacyJobMap and LegacyJobMapOutput values.
// You can construct a concrete instance of `LegacyJobMapInput` via:
//
//	LegacyJobMap{ "key": LegacyJobArgs{...} }
type LegacyJobMapInput interface {
	pulumi.Input

	ToLegacyJobMapOutput() LegacyJobMapOutput
	ToLegacyJobMapOutputWithContext(context.Context) LegacyJobMapOutput
}

type LegacyJobMap map[string]LegacyJobInput

func (LegacyJobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LegacyJob)(nil)).Elem()
}

func (i LegacyJobMap) ToLegacyJobMapOutput() LegacyJobMapOutput {
	return i.ToLegacyJobMapOutputWithContext(context.Background())
}

func (i LegacyJobMap) ToLegacyJobMapOutputWithContext(ctx context.Context) LegacyJobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegacyJobMapOutput)
}

type LegacyJobOutput struct{ *pulumi.OutputState }

func (LegacyJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LegacyJob)(nil)).Elem()
}

func (o LegacyJobOutput) ToLegacyJobOutput() LegacyJobOutput {
	return o
}

func (o LegacyJobOutput) ToLegacyJobOutputWithContext(ctx context.Context) LegacyJobOutput {
	return o
}

// Version number of dbt to use in this job, usually in the format 1.2.0-latest rather than core versions
func (o LegacyJobOutput) DbtVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.StringPtrOutput { return v.DbtVersion }).(pulumi.StringPtrOutput)
}

// Environment identifier that this job defers to (new deferring approach)
func (o LegacyJobOutput) DeferringEnvironmentId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.IntPtrOutput { return v.DeferringEnvironmentId }).(pulumi.IntPtrOutput)
}

// Job identifier that this job defers to (legacy deferring approach)
func (o LegacyJobOutput) DeferringJobId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.IntPtrOutput { return v.DeferringJobId }).(pulumi.IntPtrOutput)
}

// Description for the job
func (o LegacyJobOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Environment ID to create the job in
func (o LegacyJobOutput) EnvironmentId() pulumi.IntOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.IntOutput { return v.EnvironmentId }).(pulumi.IntOutput)
}

// List of commands to execute for the job
func (o LegacyJobOutput) ExecuteSteps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.StringArrayOutput { return v.ExecuteSteps }).(pulumi.StringArrayOutput)
}

// Flag for whether the job should generate documentation
func (o LegacyJobOutput) GenerateDocs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.BoolPtrOutput { return v.GenerateDocs }).(pulumi.BoolPtrOutput)
}

// Flag for whether the job is marked active or deleted. To create/keep a job in a 'deactivated' state, check the
// `triggers` config.
func (o LegacyJobOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.BoolPtrOutput { return v.IsActive }).(pulumi.BoolPtrOutput)
}

// Which other job should trigger this job when it finishes, and on which conditions (sometimes referred as 'job
// chaining').
func (o LegacyJobOutput) JobCompletionTriggerCondition() LegacyJobJobCompletionTriggerConditionPtrOutput {
	return o.ApplyT(func(v *LegacyJob) LegacyJobJobCompletionTriggerConditionPtrOutput {
		return v.JobCompletionTriggerCondition
	}).(LegacyJobJobCompletionTriggerConditionPtrOutput)
}

// Job name
func (o LegacyJobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of threads to use in the job
func (o LegacyJobOutput) NumThreads() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.IntPtrOutput { return v.NumThreads }).(pulumi.IntPtrOutput)
}

// Project ID to create the job in
func (o LegacyJobOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

// Flag for whether the job should add a `dbt source freshness` step to the job. The difference between manually adding a
// step with `dbt source freshness` in the job steps or using this flag is that with this flag, a failed freshness will
// still allow the following steps to run.
func (o LegacyJobOutput) RunGenerateSources() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.BoolPtrOutput { return v.RunGenerateSources }).(pulumi.BoolPtrOutput)
}

// Custom cron expression for schedule
func (o LegacyJobOutput) ScheduleCron() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.StringPtrOutput { return v.ScheduleCron }).(pulumi.StringPtrOutput)
}

// List of days of week as numbers (0 = Sunday, 7 = Saturday) to execute the job at if running on a schedule
func (o LegacyJobOutput) ScheduleDays() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.IntArrayOutput { return v.ScheduleDays }).(pulumi.IntArrayOutput)
}

// List of hours to execute the job at if running on a schedule
func (o LegacyJobOutput) ScheduleHours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.IntArrayOutput { return v.ScheduleHours }).(pulumi.IntArrayOutput)
}

// Number of hours between job executions if running on a schedule
func (o LegacyJobOutput) ScheduleInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.IntPtrOutput { return v.ScheduleInterval }).(pulumi.IntPtrOutput)
}

// Type of schedule to use, one of every_day/ days_of_week/ custom_cron
func (o LegacyJobOutput) ScheduleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.StringPtrOutput { return v.ScheduleType }).(pulumi.StringPtrOutput)
}

// Whether this job defers on a previous run of itself
func (o LegacyJobOutput) SelfDeferring() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.BoolPtrOutput { return v.SelfDeferring }).(pulumi.BoolPtrOutput)
}

// Target name for the dbt profile
func (o LegacyJobOutput) TargetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.StringPtrOutput { return v.TargetName }).(pulumi.StringPtrOutput)
}

// Number of seconds to allow the job to run before timing out
func (o LegacyJobOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.IntPtrOutput { return v.TimeoutSeconds }).(pulumi.IntPtrOutput)
}

// Flags for which types of triggers to use, the values are `github_webhook`, `git_provider_webhook`, `schedule` and
// `custom_branch_only`. <br>`custom_branch_only` is only relevant for CI jobs triggered automatically on PR creation to
// only trigger a job on a PR to the custom branch of the environment. To create a job in a 'deactivated' state, set all to
// `false`.
func (o LegacyJobOutput) Triggers() pulumi.BoolMapOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.BoolMapOutput { return v.Triggers }).(pulumi.BoolMapOutput)
}

// Whether the CI job should be automatically triggered on draft PRs
func (o LegacyJobOutput) TriggersOnDraftPr() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LegacyJob) pulumi.BoolPtrOutput { return v.TriggersOnDraftPr }).(pulumi.BoolPtrOutput)
}

type LegacyJobArrayOutput struct{ *pulumi.OutputState }

func (LegacyJobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LegacyJob)(nil)).Elem()
}

func (o LegacyJobArrayOutput) ToLegacyJobArrayOutput() LegacyJobArrayOutput {
	return o
}

func (o LegacyJobArrayOutput) ToLegacyJobArrayOutputWithContext(ctx context.Context) LegacyJobArrayOutput {
	return o
}

func (o LegacyJobArrayOutput) Index(i pulumi.IntInput) LegacyJobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LegacyJob {
		return vs[0].([]*LegacyJob)[vs[1].(int)]
	}).(LegacyJobOutput)
}

type LegacyJobMapOutput struct{ *pulumi.OutputState }

func (LegacyJobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LegacyJob)(nil)).Elem()
}

func (o LegacyJobMapOutput) ToLegacyJobMapOutput() LegacyJobMapOutput {
	return o
}

func (o LegacyJobMapOutput) ToLegacyJobMapOutputWithContext(ctx context.Context) LegacyJobMapOutput {
	return o
}

func (o LegacyJobMapOutput) MapIndex(k pulumi.StringInput) LegacyJobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LegacyJob {
		return vs[0].(map[string]*LegacyJob)[vs[1].(string)]
	}).(LegacyJobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LegacyJobInput)(nil)).Elem(), &LegacyJob{})
	pulumi.RegisterInputType(reflect.TypeOf((*LegacyJobArrayInput)(nil)).Elem(), LegacyJobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LegacyJobMapInput)(nil)).Elem(), LegacyJobMap{})
	pulumi.RegisterOutputType(LegacyJobOutput{})
	pulumi.RegisterOutputType(LegacyJobArrayOutput{})
	pulumi.RegisterOutputType(LegacyJobMapOutput{})
}

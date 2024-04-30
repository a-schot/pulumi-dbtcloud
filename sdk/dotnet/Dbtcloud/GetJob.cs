// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ASchot.Pulumi.Dbtcloud
{
    public static class GetJob
    {
        public static Task<GetJobResult> InvokeAsync(GetJobArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetJobResult>("dbtcloud:index/getJob:getJob", args ?? new GetJobArgs(), options.WithDefaults());

        public static Output<GetJobResult> Invoke(GetJobInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetJobResult>("dbtcloud:index/getJob:getJob", args ?? new GetJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetJobArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the job
        /// </summary>
        [Input("jobId", required: true)]
        public int JobId { get; set; }

        /// <summary>
        /// ID of the project the job is in
        /// </summary>
        [Input("projectId", required: true)]
        public int ProjectId { get; set; }

        public GetJobArgs()
        {
        }
        public static new GetJobArgs Empty => new GetJobArgs();
    }

    public sealed class GetJobInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the job
        /// </summary>
        [Input("jobId", required: true)]
        public Input<int> JobId { get; set; } = null!;

        /// <summary>
        /// ID of the project the job is in
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        public GetJobInvokeArgs()
        {
        }
        public static new GetJobInvokeArgs Empty => new GetJobInvokeArgs();
    }


    [OutputType]
    public sealed class GetJobResult
    {
        /// <summary>
        /// ID of the environment this job defers to
        /// </summary>
        public readonly int DeferringEnvironmentId;
        /// <summary>
        /// ID of the job this job defers to
        /// </summary>
        public readonly int DeferringJobId;
        /// <summary>
        /// Long description for the job
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// ID of the environment the job is in
        /// </summary>
        public readonly int EnvironmentId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Which other job should trigger this job when it finishes, and on which conditions.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetJobJobCompletionTriggerConditionResult> CompletionTriggerCondition;
        /// <summary>
        /// ID of the job
        /// </summary>
        public readonly int JobId;
        /// <summary>
        /// Given name for the job
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID of the project the job is in
        /// </summary>
        public readonly int ProjectId;
        /// <summary>
        /// Whether this job defers on a previous run of itself (overrides value in deferring*job*id)
        /// </summary>
        public readonly bool SelfDeferring;
        /// <summary>
        /// Number of seconds before the job times out
        /// </summary>
        public readonly int TimeoutSeconds;
        /// <summary>
        /// Flags for which types of triggers to use, keys of github*webhook, git*provider*webhook, schedule, custom*branch_only
        /// </summary>
        public readonly ImmutableDictionary<string, bool> Triggers;
        /// <summary>
        /// Whether the CI job should be automatically triggered on draft PRs
        /// </summary>
        public readonly bool TriggersOnDraftPr;

        [OutputConstructor]
        private GetJobResult(
            int deferringEnvironmentId,

            int deferringJobId,

            string description,

            int environmentId,

            string id,

            ImmutableArray<Outputs.GetJobJobCompletionTriggerConditionResult> jobCompletionTriggerConditions,

            int jobId,

            string name,

            int projectId,

            bool selfDeferring,

            int timeoutSeconds,

            ImmutableDictionary<string, bool> triggers,

            bool triggersOnDraftPr)
        {
            DeferringEnvironmentId = deferringEnvironmentId;
            DeferringJobId = deferringJobId;
            Description = description;
            EnvironmentId = environmentId;
            Id = id;
            CompletionTriggerCondition = jobCompletionTriggerConditions;
            JobId = jobId;
            Name = name;
            ProjectId = projectId;
            SelfDeferring = selfDeferring;
            TimeoutSeconds = timeoutSeconds;
            Triggers = triggers;
            TriggersOnDraftPr = triggersOnDraftPr;
        }
    }
}

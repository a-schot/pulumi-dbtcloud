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
    [Obsolete(@"Do not use! This resource is mapped from the legacy Terraform `dbt_cloud_`-prefixed resource/datasource")]
    public static class LegacyGetJob
    {
        public static Task<LegacyGetJobResult> InvokeAsync(LegacyGetJobArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<LegacyGetJobResult>("dbtcloud:index/legacyGetJob:LegacyGetJob", args ?? new LegacyGetJobArgs(), options.WithDefaults());

        public static Output<LegacyGetJobResult> Invoke(LegacyGetJobInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<LegacyGetJobResult>("dbtcloud:index/legacyGetJob:LegacyGetJob", args ?? new LegacyGetJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class LegacyGetJobArgs : global::Pulumi.InvokeArgs
    {
        [Input("jobId", required: true)]
        public int JobId { get; set; }

        [Input("projectId", required: true)]
        public int ProjectId { get; set; }

        public LegacyGetJobArgs()
        {
        }
        public static new LegacyGetJobArgs Empty => new LegacyGetJobArgs();
    }

    public sealed class LegacyGetJobInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("jobId", required: true)]
        public Input<int> JobId { get; set; } = null!;

        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        public LegacyGetJobInvokeArgs()
        {
        }
        public static new LegacyGetJobInvokeArgs Empty => new LegacyGetJobInvokeArgs();
    }


    [OutputType]
    public sealed class LegacyGetJobResult
    {
        public readonly int DeferringEnvironmentId;
        public readonly int DeferringJobId;
        public readonly string Description;
        public readonly int EnvironmentId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.LegacyGetJobJobCompletionTriggerConditionResult> CompletionTriggerCondition;
        public readonly int JobId;
        public readonly string Name;
        public readonly int ProjectId;
        public readonly bool SelfDeferring;
        public readonly int TimeoutSeconds;
        public readonly ImmutableDictionary<string, bool> Triggers;
        public readonly bool TriggersOnDraftPr;

        [OutputConstructor]
        private LegacyGetJobResult(
            int deferringEnvironmentId,

            int deferringJobId,

            string description,

            int environmentId,

            string id,

            ImmutableArray<Outputs.LegacyGetJobJobCompletionTriggerConditionResult> jobCompletionTriggerConditions,

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

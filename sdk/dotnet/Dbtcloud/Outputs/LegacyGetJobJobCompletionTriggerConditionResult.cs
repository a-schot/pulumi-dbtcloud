// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ASchot.Pulumi.Dbtcloud.Outputs
{

    [OutputType]
    public sealed class LegacyGetJobJobCompletionTriggerConditionResult
    {
        /// <summary>
        /// The ID of the job that would trigger this job after completion.
        /// </summary>
        public readonly int JobId;
        /// <summary>
        /// The ID of the project where the trigger job is running in.
        /// </summary>
        public readonly int ProjectId;
        /// <summary>
        /// List of statuses to trigger the job on.
        /// </summary>
        public readonly ImmutableArray<string> Statuses;

        [OutputConstructor]
        private LegacyGetJobJobCompletionTriggerConditionResult(
            int jobId,

            int projectId,

            ImmutableArray<string> statuses)
        {
            JobId = jobId;
            ProjectId = projectId;
            Statuses = statuses;
        }
    }
}

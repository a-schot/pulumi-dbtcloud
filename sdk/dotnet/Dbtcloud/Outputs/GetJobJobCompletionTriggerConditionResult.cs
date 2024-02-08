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
    public sealed class GetJobJobCompletionTriggerConditionResult
    {
        public readonly int JobId;
        public readonly int ProjectId;
        public readonly ImmutableArray<string> Statuses;

        [OutputConstructor]
        private GetJobJobCompletionTriggerConditionResult(
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

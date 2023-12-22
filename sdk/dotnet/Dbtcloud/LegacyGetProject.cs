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
    public static class LegacyGetProject
    {
        public static Task<LegacyGetProjectResult> InvokeAsync(LegacyGetProjectArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<LegacyGetProjectResult>("dbtcloud:index/legacyGetProject:LegacyGetProject", args ?? new LegacyGetProjectArgs(), options.WithDefaults());

        public static Output<LegacyGetProjectResult> Invoke(LegacyGetProjectInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<LegacyGetProjectResult>("dbtcloud:index/legacyGetProject:LegacyGetProject", args ?? new LegacyGetProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class LegacyGetProjectArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public string? Name { get; set; }

        [Input("projectId")]
        public int? ProjectId { get; set; }

        public LegacyGetProjectArgs()
        {
        }
        public static new LegacyGetProjectArgs Empty => new LegacyGetProjectArgs();
    }

    public sealed class LegacyGetProjectInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        public LegacyGetProjectInvokeArgs()
        {
        }
        public static new LegacyGetProjectInvokeArgs Empty => new LegacyGetProjectInvokeArgs();
    }


    [OutputType]
    public sealed class LegacyGetProjectResult
    {
        public readonly int ConnectionId;
        public readonly int DocsJobId;
        public readonly int FreshnessJobId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly int? ProjectId;
        public readonly int RepositoryId;
        public readonly int State;

        [OutputConstructor]
        private LegacyGetProjectResult(
            int connectionId,

            int docsJobId,

            int freshnessJobId,

            string id,

            string name,

            int? projectId,

            int repositoryId,

            int state)
        {
            ConnectionId = connectionId;
            DocsJobId = docsJobId;
            FreshnessJobId = freshnessJobId;
            Id = id;
            Name = name;
            ProjectId = projectId;
            RepositoryId = repositoryId;
            State = state;
        }
    }
}
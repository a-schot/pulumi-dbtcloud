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
    public static class LegacyGetConnection
    {
        public static Task<LegacyGetConnectionResult> InvokeAsync(LegacyGetConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<LegacyGetConnectionResult>("dbtcloud:index/legacyGetConnection:LegacyGetConnection", args ?? new LegacyGetConnectionArgs(), options.WithDefaults());

        public static Output<LegacyGetConnectionResult> Invoke(LegacyGetConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<LegacyGetConnectionResult>("dbtcloud:index/legacyGetConnection:LegacyGetConnection", args ?? new LegacyGetConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class LegacyGetConnectionArgs : global::Pulumi.InvokeArgs
    {
        [Input("connectionId", required: true)]
        public int ConnectionId { get; set; }

        [Input("projectId", required: true)]
        public int ProjectId { get; set; }

        public LegacyGetConnectionArgs()
        {
        }
        public static new LegacyGetConnectionArgs Empty => new LegacyGetConnectionArgs();
    }

    public sealed class LegacyGetConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("connectionId", required: true)]
        public Input<int> ConnectionId { get; set; } = null!;

        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        public LegacyGetConnectionInvokeArgs()
        {
        }
        public static new LegacyGetConnectionInvokeArgs Empty => new LegacyGetConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class LegacyGetConnectionResult
    {
        public readonly string Account;
        public readonly bool AllowKeepAlive;
        public readonly bool AllowSso;
        public readonly int ConnectionId;
        public readonly string Database;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsActive;
        public readonly string Name;
        public readonly string PrivateLinkEndpointId;
        public readonly int ProjectId;
        public readonly string Role;
        public readonly string Type;
        public readonly string Warehouse;

        [OutputConstructor]
        private LegacyGetConnectionResult(
            string account,

            bool allowKeepAlive,

            bool allowSso,

            int connectionId,

            string database,

            string id,

            bool isActive,

            string name,

            string privateLinkEndpointId,

            int projectId,

            string role,

            string type,

            string warehouse)
        {
            Account = account;
            AllowKeepAlive = allowKeepAlive;
            AllowSso = allowSso;
            ConnectionId = connectionId;
            Database = database;
            Id = id;
            IsActive = isActive;
            Name = name;
            PrivateLinkEndpointId = privateLinkEndpointId;
            ProjectId = projectId;
            Role = role;
            Type = type;
            Warehouse = warehouse;
        }
    }
}
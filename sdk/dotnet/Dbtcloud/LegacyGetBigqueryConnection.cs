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
    public static class LegacyGetBigqueryConnection
    {
        public static Task<LegacyGetBigqueryConnectionResult> InvokeAsync(LegacyGetBigqueryConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<LegacyGetBigqueryConnectionResult>("dbtcloud:index/legacyGetBigqueryConnection:LegacyGetBigqueryConnection", args ?? new LegacyGetBigqueryConnectionArgs(), options.WithDefaults());

        public static Output<LegacyGetBigqueryConnectionResult> Invoke(LegacyGetBigqueryConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<LegacyGetBigqueryConnectionResult>("dbtcloud:index/legacyGetBigqueryConnection:LegacyGetBigqueryConnection", args ?? new LegacyGetBigqueryConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class LegacyGetBigqueryConnectionArgs : global::Pulumi.InvokeArgs
    {
        [Input("connectionId", required: true)]
        public int ConnectionId { get; set; }

        [Input("projectId", required: true)]
        public int ProjectId { get; set; }

        public LegacyGetBigqueryConnectionArgs()
        {
        }
        public static new LegacyGetBigqueryConnectionArgs Empty => new LegacyGetBigqueryConnectionArgs();
    }

    public sealed class LegacyGetBigqueryConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("connectionId", required: true)]
        public Input<int> ConnectionId { get; set; } = null!;

        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        public LegacyGetBigqueryConnectionInvokeArgs()
        {
        }
        public static new LegacyGetBigqueryConnectionInvokeArgs Empty => new LegacyGetBigqueryConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class LegacyGetBigqueryConnectionResult
    {
        public readonly string AuthProviderX509CertUrl;
        public readonly string AuthUri;
        public readonly string ClientEmail;
        public readonly string ClientId;
        public readonly string ClientX509CertUrl;
        public readonly int ConnectionId;
        public readonly string DataprocClusterName;
        public readonly string DataprocRegion;
        public readonly string ExecutionProject;
        public readonly string GcpProjectId;
        public readonly string GcsBucket;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsActive;
        public readonly bool IsConfiguredForOauth;
        public readonly string Location;
        public readonly int MaximumBytesBilled;
        public readonly string Name;
        public readonly string Priority;
        public readonly string PrivateKey;
        public readonly string PrivateKeyId;
        public readonly int ProjectId;
        public readonly int Retries;
        public readonly int TimeoutSeconds;
        public readonly string TokenUri;
        public readonly string Type;

        [OutputConstructor]
        private LegacyGetBigqueryConnectionResult(
            string authProviderX509CertUrl,

            string authUri,

            string clientEmail,

            string clientId,

            string clientX509CertUrl,

            int connectionId,

            string dataprocClusterName,

            string dataprocRegion,

            string executionProject,

            string gcpProjectId,

            string gcsBucket,

            string id,

            bool isActive,

            bool isConfiguredForOauth,

            string location,

            int maximumBytesBilled,

            string name,

            string priority,

            string privateKey,

            string privateKeyId,

            int projectId,

            int retries,

            int timeoutSeconds,

            string tokenUri,

            string type)
        {
            AuthProviderX509CertUrl = authProviderX509CertUrl;
            AuthUri = authUri;
            ClientEmail = clientEmail;
            ClientId = clientId;
            ClientX509CertUrl = clientX509CertUrl;
            ConnectionId = connectionId;
            DataprocClusterName = dataprocClusterName;
            DataprocRegion = dataprocRegion;
            ExecutionProject = executionProject;
            GcpProjectId = gcpProjectId;
            GcsBucket = gcsBucket;
            Id = id;
            IsActive = isActive;
            IsConfiguredForOauth = isConfiguredForOauth;
            Location = location;
            MaximumBytesBilled = maximumBytesBilled;
            Name = name;
            Priority = priority;
            PrivateKey = privateKey;
            PrivateKeyId = privateKeyId;
            ProjectId = projectId;
            Retries = retries;
            TimeoutSeconds = timeoutSeconds;
            TokenUri = tokenUri;
            Type = type;
        }
    }
}
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
    public static class LegacyGetSnowflakeCredential
    {
        public static Task<LegacyGetSnowflakeCredentialResult> InvokeAsync(LegacyGetSnowflakeCredentialArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<LegacyGetSnowflakeCredentialResult>("dbtcloud:index/legacyGetSnowflakeCredential:LegacyGetSnowflakeCredential", args ?? new LegacyGetSnowflakeCredentialArgs(), options.WithDefaults());

        public static Output<LegacyGetSnowflakeCredentialResult> Invoke(LegacyGetSnowflakeCredentialInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<LegacyGetSnowflakeCredentialResult>("dbtcloud:index/legacyGetSnowflakeCredential:LegacyGetSnowflakeCredential", args ?? new LegacyGetSnowflakeCredentialInvokeArgs(), options.WithDefaults());
    }


    public sealed class LegacyGetSnowflakeCredentialArgs : global::Pulumi.InvokeArgs
    {
        [Input("credentialId", required: true)]
        public int CredentialId { get; set; }

        [Input("projectId", required: true)]
        public int ProjectId { get; set; }

        public LegacyGetSnowflakeCredentialArgs()
        {
        }
        public static new LegacyGetSnowflakeCredentialArgs Empty => new LegacyGetSnowflakeCredentialArgs();
    }

    public sealed class LegacyGetSnowflakeCredentialInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("credentialId", required: true)]
        public Input<int> CredentialId { get; set; } = null!;

        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        public LegacyGetSnowflakeCredentialInvokeArgs()
        {
        }
        public static new LegacyGetSnowflakeCredentialInvokeArgs Empty => new LegacyGetSnowflakeCredentialInvokeArgs();
    }


    [OutputType]
    public sealed class LegacyGetSnowflakeCredentialResult
    {
        public readonly string AuthType;
        public readonly int CredentialId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsActive;
        public readonly int NumThreads;
        public readonly int ProjectId;
        public readonly string Schema;
        public readonly string User;

        [OutputConstructor]
        private LegacyGetSnowflakeCredentialResult(
            string authType,

            int credentialId,

            string id,

            bool isActive,

            int numThreads,

            int projectId,

            string schema,

            string user)
        {
            AuthType = authType;
            CredentialId = credentialId;
            Id = id;
            IsActive = isActive;
            NumThreads = numThreads;
            ProjectId = projectId;
            Schema = schema;
            User = user;
        }
    }
}

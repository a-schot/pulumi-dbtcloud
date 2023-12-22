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
    public static class GetRepository
    {
        public static Task<GetRepositoryResult> InvokeAsync(GetRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryResult>("dbtcloud:index/getRepository:getRepository", args ?? new GetRepositoryArgs(), options.WithDefaults());

        public static Output<GetRepositoryResult> Invoke(GetRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryResult>("dbtcloud:index/getRepository:getRepository", args ?? new GetRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether we should return the public deploy key
        /// </summary>
        [Input("fetchDeployKey")]
        public bool? FetchDeployKey { get; set; }

        /// <summary>
        /// Project ID to create the repository in
        /// </summary>
        [Input("projectId", required: true)]
        public int ProjectId { get; set; }

        /// <summary>
        /// ID for the repository
        /// </summary>
        [Input("repositoryId", required: true)]
        public int RepositoryId { get; set; }

        public GetRepositoryArgs()
        {
        }
        public static new GetRepositoryArgs Empty => new GetRepositoryArgs();
    }

    public sealed class GetRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether we should return the public deploy key
        /// </summary>
        [Input("fetchDeployKey")]
        public Input<bool>? FetchDeployKey { get; set; }

        /// <summary>
        /// Project ID to create the repository in
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        /// <summary>
        /// ID for the repository
        /// </summary>
        [Input("repositoryId", required: true)]
        public Input<int> RepositoryId { get; set; } = null!;

        public GetRepositoryInvokeArgs()
        {
        }
        public static new GetRepositoryInvokeArgs Empty => new GetRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoryResult
    {
        /// <summary>
        /// Public key generated by dbt when using `deploy_key` clone strategy
        /// </summary>
        public readonly string DeployKey;
        /// <summary>
        /// Whether we should return the public deploy key
        /// </summary>
        public readonly bool? FetchDeployKey;
        /// <summary>
        /// Git clone strategy for the repository
        /// </summary>
        public readonly string GitCloneStrategy;
        /// <summary>
        /// Identifier for the GitHub installation
        /// </summary>
        public readonly int GithubInstallationId;
        /// <summary>
        /// Identifier for the Gitlab project
        /// </summary>
        public readonly int GitlabProjectId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether the repository is active
        /// </summary>
        public readonly bool IsActive;
        /// <summary>
        /// Project ID to create the repository in
        /// </summary>
        public readonly int ProjectId;
        /// <summary>
        /// Connection name
        /// </summary>
        public readonly string RemoteUrl;
        /// <summary>
        /// Credentials ID for the repository (From the repository side not the dbt Cloud ID)
        /// </summary>
        public readonly int RepositoryCredentialsId;
        /// <summary>
        /// ID for the repository
        /// </summary>
        public readonly int RepositoryId;

        [OutputConstructor]
        private GetRepositoryResult(
            string deployKey,

            bool? fetchDeployKey,

            string gitCloneStrategy,

            int githubInstallationId,

            int gitlabProjectId,

            string id,

            bool isActive,

            int projectId,

            string remoteUrl,

            int repositoryCredentialsId,

            int repositoryId)
        {
            DeployKey = deployKey;
            FetchDeployKey = fetchDeployKey;
            GitCloneStrategy = gitCloneStrategy;
            GithubInstallationId = githubInstallationId;
            GitlabProjectId = gitlabProjectId;
            Id = id;
            IsActive = isActive;
            ProjectId = projectId;
            RemoteUrl = remoteUrl;
            RepositoryCredentialsId = repositoryCredentialsId;
            RepositoryId = repositoryId;
        }
    }
}

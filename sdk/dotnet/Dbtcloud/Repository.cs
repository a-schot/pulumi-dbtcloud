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
    /// <summary>
    /// *Note*: Some upstream resources can be slow to create, so if creating a project at
    /// the same time as the repository, it's recommended to use the `depends_on` meta argument.
    /// 
    /// In order to find the `github_installation_id`, you can log in to dbt Cloud, replace `&lt;dbt_cloud_url&gt;` by your dbt Cloud
    /// URL and run the following commands in the Google Chrome console:
    /// 
    /// Alternatively, you can go to the page `https://&lt;dbt_cloud_url&gt;/api/v2/integrations/github/installations/` and read the
    /// value of `id`  or use the `http` provider to retrieve it automatically like in the example below.
    /// 
    /// ## Example Usage
    /// ### repo cloned via the GitHub integration, manually entering the `github_installation_id`
    /// resource "dbtcloud_repository" "github_repo" {
    ///   project_id             = dbtcloud_project.dbt_project.id
    ///   remote_url             = "git@github.com:&lt;github_org&gt;/&lt;github_repo&gt;.git"
    ///   github_installation_id = 9876
    ///   git_clone_strategy     = "github_app"
    /// }
    /// 
    /// ### repo cloned via the GitHub integration, with auto-retrieval of the `github_installation_id`
    /// # here, we assume that `token` and `host_url` are respectively accessible via `var.dbt_token` and `var.dbt_host_url`
    /// # NOTE: the following requires connecting via a user token and can't be retrieved with a service token
    /// data "http" "github_installations_response" {
    ///   url = format("%s/v2/integrations/github/installations/", var.dbt_host_url)
    ///   request_headers = {
    ///     Authorization = format("Bearer %s", var.dbt_token)
    ///   }
    /// }
    /// 
    /// locals {
    ///   github_installation_id = jsondecode(data.http.github_installations_response.response_body)[0].id
    /// }
    /// 
    /// resource "dbtcloud_repository" "github_repo_other" {
    ///   project_id             = dbtcloud_project.dbt_project.id
    ///   remote_url             = "git@github.com:&lt;github_org&gt;/&lt;github_repo&gt;.git"
    ///   github_installation_id = local.github_installation_id
    ///   git_clone_strategy     = "github_app"
    /// }
    /// 
    /// ### repo cloned via the GitLab integration
    /// # as of 15 Sept 2023 this resource requires using a user token and can't be set with a service token - CC-791
    /// resource "dbtcloud_repository" "gitlab_repo" {
    ///   project_id         = dbtcloud_project.dbt_project.id
    ///   remote_url         = "&lt;gitlab-group&gt;/&lt;gitlab-project&gt;"
    ///   gitlab_project_id  = 8765
    ///   git_clone_strategy = "deploy_token"
    /// }
    /// 
    /// ### repo cloned via the deploy token strategy
    /// resource "dbtcloud_repository" "deploy_repo" {
    ///   project_id         = dbtcloud_project.dbt_project.id
    ///   remote_url         = "git://github.com/&lt;github_org&gt;/&lt;github_repo&gt;.git"
    ///   git_clone_strategy = "deploy_key"
    /// }
    /// 
    /// ## Import
    /// 
    /// Import using a project ID and repository ID found in the URL or via the API.
    /// 
    /// ```sh
    ///  $ pulumi import dbtcloud:index/repository:Repository test_repository "project_id:repository_id"
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import dbtcloud:index/repository:Repository test_repository 12345:6789
    /// ```
    /// </summary>
    [DbtcloudResourceType("dbtcloud:index/repository:Repository")]
    public partial class Repository : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure Dev Ops project ID. It can be retrieved via the Azure API or using the data source `dbtcloud.getAzureDevOpsProject` and the project name - (for ADO native integration only)
        /// </summary>
        [Output("azureActiveDirectoryProjectId")]
        public Output<string?> AzureActiveDirectoryProjectId { get; private set; } = null!;

        /// <summary>
        /// The Azure Dev Ops repository ID. It can be retrieved via the Azure API or using the data source `dbtcloud.getAzureDevOpsRepository` along with the ADO Project ID and the repository name - (for ADO native integration only)
        /// </summary>
        [Output("azureActiveDirectoryRepositoryId")]
        public Output<string?> AzureActiveDirectoryRepositoryId { get; private set; } = null!;

        /// <summary>
        /// If set to False (the default), the connection will fail if the service user doesn't have access to set webhooks (required for auto-triggering CI jobs). If set to True, the connection will be successful but no automated CI job will be triggered - (for ADO native integration only)
        /// </summary>
        [Output("azureBypassWebhookRegistrationFailure")]
        public Output<bool?> AzureBypassWebhookRegistrationFailure { get; private set; } = null!;

        /// <summary>
        /// Public key generated by dbt when using `deploy_key` clone strategy
        /// </summary>
        [Output("deployKey")]
        public Output<string> DeployKey { get; private set; } = null!;

        /// <summary>
        /// Whether we should return the public deploy key - (for the `deploy_key` strategy)
        /// </summary>
        [Output("fetchDeployKey")]
        public Output<bool?> FetchDeployKey { get; private set; } = null!;

        /// <summary>
        /// Git clone strategy for the repository. Can be `deploy_key` (default) for cloning via SSH Deploy Key, `github_app` for GitHub native integration, `deploy_token` for the GitLab native integration and `azure_active_directory_app` for ADO native integration
        /// </summary>
        [Output("gitCloneStrategy")]
        public Output<string?> GitCloneStrategy { get; private set; } = null!;

        /// <summary>
        /// Identifier for the GitHub App - (for GitHub native integration only)
        /// </summary>
        [Output("githubInstallationId")]
        public Output<int?> GithubInstallationId { get; private set; } = null!;

        /// <summary>
        /// Identifier for the Gitlab project -  (for GitLab native integration only)
        /// </summary>
        [Output("gitlabProjectId")]
        public Output<int?> GitlabProjectId { get; private set; } = null!;

        /// <summary>
        /// Whether the repository is active
        /// </summary>
        [Output("isActive")]
        public Output<bool?> IsActive { get; private set; } = null!;

        /// <summary>
        /// Project ID to create the repository in
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Git URL for the repository or &lt;Group&gt;/&lt;Project&gt; for Gitlab
        /// </summary>
        [Output("remoteUrl")]
        public Output<string> RemoteUrl { get; private set; } = null!;

        /// <summary>
        /// Credentials ID for the repository (From the repository side not the dbt Cloud ID)
        /// </summary>
        [Output("repositoryCredentialsId")]
        public Output<int> RepositoryCredentialsId { get; private set; } = null!;

        /// <summary>
        /// Repository Identifier
        /// </summary>
        [Output("repositoryId")]
        public Output<int> RepositoryId { get; private set; } = null!;


        /// <summary>
        /// Create a Repository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Repository(string name, RepositoryArgs args, CustomResourceOptions? options = null)
            : base("dbtcloud:index/repository:Repository", name, args ?? new RepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Repository(string name, Input<string> id, RepositoryState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/repository:Repository", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/a-schot/pulumi-dbtcloud/releases/download/v${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Repository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Repository Get(string name, Input<string> id, RepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new Repository(name, id, state, options);
        }
    }

    public sealed class RepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure Dev Ops project ID. It can be retrieved via the Azure API or using the data source `dbtcloud.getAzureDevOpsProject` and the project name - (for ADO native integration only)
        /// </summary>
        [Input("azureActiveDirectoryProjectId")]
        public Input<string>? AzureActiveDirectoryProjectId { get; set; }

        /// <summary>
        /// The Azure Dev Ops repository ID. It can be retrieved via the Azure API or using the data source `dbtcloud.getAzureDevOpsRepository` along with the ADO Project ID and the repository name - (for ADO native integration only)
        /// </summary>
        [Input("azureActiveDirectoryRepositoryId")]
        public Input<string>? AzureActiveDirectoryRepositoryId { get; set; }

        /// <summary>
        /// If set to False (the default), the connection will fail if the service user doesn't have access to set webhooks (required for auto-triggering CI jobs). If set to True, the connection will be successful but no automated CI job will be triggered - (for ADO native integration only)
        /// </summary>
        [Input("azureBypassWebhookRegistrationFailure")]
        public Input<bool>? AzureBypassWebhookRegistrationFailure { get; set; }

        /// <summary>
        /// Whether we should return the public deploy key - (for the `deploy_key` strategy)
        /// </summary>
        [Input("fetchDeployKey")]
        public Input<bool>? FetchDeployKey { get; set; }

        /// <summary>
        /// Git clone strategy for the repository. Can be `deploy_key` (default) for cloning via SSH Deploy Key, `github_app` for GitHub native integration, `deploy_token` for the GitLab native integration and `azure_active_directory_app` for ADO native integration
        /// </summary>
        [Input("gitCloneStrategy")]
        public Input<string>? GitCloneStrategy { get; set; }

        /// <summary>
        /// Identifier for the GitHub App - (for GitHub native integration only)
        /// </summary>
        [Input("githubInstallationId")]
        public Input<int>? GithubInstallationId { get; set; }

        /// <summary>
        /// Identifier for the Gitlab project -  (for GitLab native integration only)
        /// </summary>
        [Input("gitlabProjectId")]
        public Input<int>? GitlabProjectId { get; set; }

        /// <summary>
        /// Whether the repository is active
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Project ID to create the repository in
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        /// <summary>
        /// Git URL for the repository or &lt;Group&gt;/&lt;Project&gt; for Gitlab
        /// </summary>
        [Input("remoteUrl", required: true)]
        public Input<string> RemoteUrl { get; set; } = null!;

        public RepositoryArgs()
        {
        }
        public static new RepositoryArgs Empty => new RepositoryArgs();
    }

    public sealed class RepositoryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure Dev Ops project ID. It can be retrieved via the Azure API or using the data source `dbtcloud.getAzureDevOpsProject` and the project name - (for ADO native integration only)
        /// </summary>
        [Input("azureActiveDirectoryProjectId")]
        public Input<string>? AzureActiveDirectoryProjectId { get; set; }

        /// <summary>
        /// The Azure Dev Ops repository ID. It can be retrieved via the Azure API or using the data source `dbtcloud.getAzureDevOpsRepository` along with the ADO Project ID and the repository name - (for ADO native integration only)
        /// </summary>
        [Input("azureActiveDirectoryRepositoryId")]
        public Input<string>? AzureActiveDirectoryRepositoryId { get; set; }

        /// <summary>
        /// If set to False (the default), the connection will fail if the service user doesn't have access to set webhooks (required for auto-triggering CI jobs). If set to True, the connection will be successful but no automated CI job will be triggered - (for ADO native integration only)
        /// </summary>
        [Input("azureBypassWebhookRegistrationFailure")]
        public Input<bool>? AzureBypassWebhookRegistrationFailure { get; set; }

        /// <summary>
        /// Public key generated by dbt when using `deploy_key` clone strategy
        /// </summary>
        [Input("deployKey")]
        public Input<string>? DeployKey { get; set; }

        /// <summary>
        /// Whether we should return the public deploy key - (for the `deploy_key` strategy)
        /// </summary>
        [Input("fetchDeployKey")]
        public Input<bool>? FetchDeployKey { get; set; }

        /// <summary>
        /// Git clone strategy for the repository. Can be `deploy_key` (default) for cloning via SSH Deploy Key, `github_app` for GitHub native integration, `deploy_token` for the GitLab native integration and `azure_active_directory_app` for ADO native integration
        /// </summary>
        [Input("gitCloneStrategy")]
        public Input<string>? GitCloneStrategy { get; set; }

        /// <summary>
        /// Identifier for the GitHub App - (for GitHub native integration only)
        /// </summary>
        [Input("githubInstallationId")]
        public Input<int>? GithubInstallationId { get; set; }

        /// <summary>
        /// Identifier for the Gitlab project -  (for GitLab native integration only)
        /// </summary>
        [Input("gitlabProjectId")]
        public Input<int>? GitlabProjectId { get; set; }

        /// <summary>
        /// Whether the repository is active
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Project ID to create the repository in
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Git URL for the repository or &lt;Group&gt;/&lt;Project&gt; for Gitlab
        /// </summary>
        [Input("remoteUrl")]
        public Input<string>? RemoteUrl { get; set; }

        /// <summary>
        /// Credentials ID for the repository (From the repository side not the dbt Cloud ID)
        /// </summary>
        [Input("repositoryCredentialsId")]
        public Input<int>? RepositoryCredentialsId { get; set; }

        /// <summary>
        /// Repository Identifier
        /// </summary>
        [Input("repositoryId")]
        public Input<int>? RepositoryId { get; set; }

        public RepositoryState()
        {
        }
        public static new RepositoryState Empty => new RepositoryState();
    }
}
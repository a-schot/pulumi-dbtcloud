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
    [DbtcloudResourceType("dbtcloud:index/legacyEnvironment:LegacyEnvironment")]
    public partial class LegacyEnvironment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Credential ID to create the environment with. A credential is not required for development environments but is required
        /// for deployment environments
        /// </summary>
        [Output("credentialId")]
        public Output<int?> CredentialId { get; private set; } = null!;

        /// <summary>
        /// Which custom branch to use in this environment
        /// </summary>
        [Output("customBranch")]
        public Output<string?> CustomBranch { get; private set; } = null!;

        /// <summary>
        /// Version number of dbt to use in this environment. It needs to be in the format `major.minor.0-latest` or
        /// `major.minor.0-pre`, e.g. `1.5.0-latest`
        /// </summary>
        [Output("dbtVersion")]
        public Output<string> DbtVersion { get; private set; } = null!;

        /// <summary>
        /// The type of environment. Only valid for environments of type 'deployment' and for now can only be empty or set to
        /// 'production'
        /// </summary>
        [Output("deploymentType")]
        public Output<string?> DeploymentType { get; private set; } = null!;

        /// <summary>
        /// Environment ID within the project
        /// </summary>
        [Output("environmentId")]
        public Output<int> EnvironmentId { get; private set; } = null!;

        /// <summary>
        /// ID of the extended attributes for the environment
        /// </summary>
        [Output("extendedAttributesId")]
        public Output<int?> ExtendedAttributesId { get; private set; } = null!;

        /// <summary>
        /// Whether the environment is active
        /// </summary>
        [Output("isActive")]
        public Output<bool?> IsActive { get; private set; } = null!;

        /// <summary>
        /// Environment name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Project ID to create the environment in
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The type of environment (must be either development or deployment)
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Whether to use a custom git branch in this environment
        /// </summary>
        [Output("useCustomBranch")]
        public Output<bool?> UseCustomBranch { get; private set; } = null!;


        /// <summary>
        /// Create a LegacyEnvironment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LegacyEnvironment(string name, LegacyEnvironmentArgs args, CustomResourceOptions? options = null)
            : base("dbtcloud:index/legacyEnvironment:LegacyEnvironment", name, args ?? new LegacyEnvironmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LegacyEnvironment(string name, Input<string> id, LegacyEnvironmentState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/legacyEnvironment:LegacyEnvironment", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/a-schot/pulumi-dbtcloud",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LegacyEnvironment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LegacyEnvironment Get(string name, Input<string> id, LegacyEnvironmentState? state = null, CustomResourceOptions? options = null)
        {
            return new LegacyEnvironment(name, id, state, options);
        }
    }

    public sealed class LegacyEnvironmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Credential ID to create the environment with. A credential is not required for development environments but is required
        /// for deployment environments
        /// </summary>
        [Input("credentialId")]
        public Input<int>? CredentialId { get; set; }

        /// <summary>
        /// Which custom branch to use in this environment
        /// </summary>
        [Input("customBranch")]
        public Input<string>? CustomBranch { get; set; }

        /// <summary>
        /// Version number of dbt to use in this environment. It needs to be in the format `major.minor.0-latest` or
        /// `major.minor.0-pre`, e.g. `1.5.0-latest`
        /// </summary>
        [Input("dbtVersion", required: true)]
        public Input<string> DbtVersion { get; set; } = null!;

        /// <summary>
        /// The type of environment. Only valid for environments of type 'deployment' and for now can only be empty or set to
        /// 'production'
        /// </summary>
        [Input("deploymentType")]
        public Input<string>? DeploymentType { get; set; }

        /// <summary>
        /// ID of the extended attributes for the environment
        /// </summary>
        [Input("extendedAttributesId")]
        public Input<int>? ExtendedAttributesId { get; set; }

        /// <summary>
        /// Whether the environment is active
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Environment name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Project ID to create the environment in
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        /// <summary>
        /// The type of environment (must be either development or deployment)
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Whether to use a custom git branch in this environment
        /// </summary>
        [Input("useCustomBranch")]
        public Input<bool>? UseCustomBranch { get; set; }

        public LegacyEnvironmentArgs()
        {
        }
        public static new LegacyEnvironmentArgs Empty => new LegacyEnvironmentArgs();
    }

    public sealed class LegacyEnvironmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Credential ID to create the environment with. A credential is not required for development environments but is required
        /// for deployment environments
        /// </summary>
        [Input("credentialId")]
        public Input<int>? CredentialId { get; set; }

        /// <summary>
        /// Which custom branch to use in this environment
        /// </summary>
        [Input("customBranch")]
        public Input<string>? CustomBranch { get; set; }

        /// <summary>
        /// Version number of dbt to use in this environment. It needs to be in the format `major.minor.0-latest` or
        /// `major.minor.0-pre`, e.g. `1.5.0-latest`
        /// </summary>
        [Input("dbtVersion")]
        public Input<string>? DbtVersion { get; set; }

        /// <summary>
        /// The type of environment. Only valid for environments of type 'deployment' and for now can only be empty or set to
        /// 'production'
        /// </summary>
        [Input("deploymentType")]
        public Input<string>? DeploymentType { get; set; }

        /// <summary>
        /// Environment ID within the project
        /// </summary>
        [Input("environmentId")]
        public Input<int>? EnvironmentId { get; set; }

        /// <summary>
        /// ID of the extended attributes for the environment
        /// </summary>
        [Input("extendedAttributesId")]
        public Input<int>? ExtendedAttributesId { get; set; }

        /// <summary>
        /// Whether the environment is active
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Environment name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Project ID to create the environment in
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// The type of environment (must be either development or deployment)
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Whether to use a custom git branch in this environment
        /// </summary>
        [Input("useCustomBranch")]
        public Input<bool>? UseCustomBranch { get; set; }

        public LegacyEnvironmentState()
        {
        }
        public static new LegacyEnvironmentState Empty => new LegacyEnvironmentState();
    }
}

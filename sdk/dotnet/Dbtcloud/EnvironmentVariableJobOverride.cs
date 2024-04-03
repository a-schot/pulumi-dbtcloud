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
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Dbtcloud = ASchot.Pulumi.Dbtcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myEnvVarJobOverride = new Dbtcloud.EnvironmentVariableJobOverride("myEnvVarJobOverride", new()
    ///     {
    ///         ProjectId = dbtcloud_project.Dbt_project.Id,
    ///         JobDefinitionId = dbtcloud_job.Daily_job.Id,
    ///         RawValue = "my_override_value",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Import using a project ID, a job ID and the environment variable override ID
    /// 
    /// ```sh
    ///  $ pulumi import dbtcloud:index/environmentVariableJobOverride:EnvironmentVariableJobOverride test_environment_variable_job_override "project_id:job_id:environment_variable_override_id"
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import dbtcloud:index/environmentVariableJobOverride:EnvironmentVariableJobOverride test_environment_variable_job_override 12345:678:123456
    /// ```
    /// </summary>
    [DbtcloudResourceType("dbtcloud:index/environmentVariableJobOverride:EnvironmentVariableJobOverride")]
    public partial class EnvironmentVariableJobOverride : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the environment variable job override
        /// </summary>
        [Output("environmentVariableJobOverrideId")]
        public Output<int> EnvironmentVariableJobOverrideId { get; private set; } = null!;

        /// <summary>
        /// The job ID for which the environment variable is being overridden
        /// </summary>
        [Output("jobDefinitionId")]
        public Output<int> JobDefinitionId { get; private set; } = null!;

        /// <summary>
        /// The environment variable name to override
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project ID for which the environment variable is being overridden
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The value for the override of the environment variable
        /// </summary>
        [Output("rawValue")]
        public Output<string> RawValue { get; private set; } = null!;


        /// <summary>
        /// Create a EnvironmentVariableJobOverride resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EnvironmentVariableJobOverride(string name, EnvironmentVariableJobOverrideArgs args, CustomResourceOptions? options = null)
            : base("dbtcloud:index/environmentVariableJobOverride:EnvironmentVariableJobOverride", name, args ?? new EnvironmentVariableJobOverrideArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EnvironmentVariableJobOverride(string name, Input<string> id, EnvironmentVariableJobOverrideState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/environmentVariableJobOverride:EnvironmentVariableJobOverride", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EnvironmentVariableJobOverride resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EnvironmentVariableJobOverride Get(string name, Input<string> id, EnvironmentVariableJobOverrideState? state = null, CustomResourceOptions? options = null)
        {
            return new EnvironmentVariableJobOverride(name, id, state, options);
        }
    }

    public sealed class EnvironmentVariableJobOverrideArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The job ID for which the environment variable is being overridden
        /// </summary>
        [Input("jobDefinitionId", required: true)]
        public Input<int> JobDefinitionId { get; set; } = null!;

        /// <summary>
        /// The environment variable name to override
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project ID for which the environment variable is being overridden
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        /// <summary>
        /// The value for the override of the environment variable
        /// </summary>
        [Input("rawValue", required: true)]
        public Input<string> RawValue { get; set; } = null!;

        public EnvironmentVariableJobOverrideArgs()
        {
        }
        public static new EnvironmentVariableJobOverrideArgs Empty => new EnvironmentVariableJobOverrideArgs();
    }

    public sealed class EnvironmentVariableJobOverrideState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the environment variable job override
        /// </summary>
        [Input("environmentVariableJobOverrideId")]
        public Input<int>? EnvironmentVariableJobOverrideId { get; set; }

        /// <summary>
        /// The job ID for which the environment variable is being overridden
        /// </summary>
        [Input("jobDefinitionId")]
        public Input<int>? JobDefinitionId { get; set; }

        /// <summary>
        /// The environment variable name to override
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project ID for which the environment variable is being overridden
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// The value for the override of the environment variable
        /// </summary>
        [Input("rawValue")]
        public Input<string>? RawValue { get; set; }

        public EnvironmentVariableJobOverrideState()
        {
        }
        public static new EnvironmentVariableJobOverrideState Empty => new EnvironmentVariableJobOverrideState();
    }
}

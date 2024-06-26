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
    [DbtcloudResourceType("dbtcloud:index/legacyProjectRepository:LegacyProjectRepository")]
    public partial class LegacyProjectRepository : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Project ID
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Repository ID
        /// </summary>
        [Output("repositoryId")]
        public Output<int> RepositoryId { get; private set; } = null!;


        /// <summary>
        /// Create a LegacyProjectRepository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LegacyProjectRepository(string name, LegacyProjectRepositoryArgs args, CustomResourceOptions? options = null)
            : base("dbtcloud:index/legacyProjectRepository:LegacyProjectRepository", name, args ?? new LegacyProjectRepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LegacyProjectRepository(string name, Input<string> id, LegacyProjectRepositoryState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/legacyProjectRepository:LegacyProjectRepository", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LegacyProjectRepository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LegacyProjectRepository Get(string name, Input<string> id, LegacyProjectRepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new LegacyProjectRepository(name, id, state, options);
        }
    }

    public sealed class LegacyProjectRepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Project ID
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        /// <summary>
        /// Repository ID
        /// </summary>
        [Input("repositoryId", required: true)]
        public Input<int> RepositoryId { get; set; } = null!;

        public LegacyProjectRepositoryArgs()
        {
        }
        public static new LegacyProjectRepositoryArgs Empty => new LegacyProjectRepositoryArgs();
    }

    public sealed class LegacyProjectRepositoryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Project ID
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Repository ID
        /// </summary>
        [Input("repositoryId")]
        public Input<int>? RepositoryId { get; set; }

        public LegacyProjectRepositoryState()
        {
        }
        public static new LegacyProjectRepositoryState Empty => new LegacyProjectRepositoryState();
    }
}

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
    [DbtcloudResourceType("dbtcloud:index/legacyProjectArtefacts:LegacyProjectArtefacts")]
    public partial class LegacyProjectArtefacts : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Docs Job ID
        /// </summary>
        [Output("docsJobId")]
        public Output<int?> DocsJobId { get; private set; } = null!;

        /// <summary>
        /// Freshness Job ID
        /// </summary>
        [Output("freshnessJobId")]
        public Output<int?> FreshnessJobId { get; private set; } = null!;

        /// <summary>
        /// Project ID
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;


        /// <summary>
        /// Create a LegacyProjectArtefacts resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LegacyProjectArtefacts(string name, LegacyProjectArtefactsArgs args, CustomResourceOptions? options = null)
            : base("dbtcloud:index/legacyProjectArtefacts:LegacyProjectArtefacts", name, args ?? new LegacyProjectArtefactsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LegacyProjectArtefacts(string name, Input<string> id, LegacyProjectArtefactsState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/legacyProjectArtefacts:LegacyProjectArtefacts", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LegacyProjectArtefacts resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LegacyProjectArtefacts Get(string name, Input<string> id, LegacyProjectArtefactsState? state = null, CustomResourceOptions? options = null)
        {
            return new LegacyProjectArtefacts(name, id, state, options);
        }
    }

    public sealed class LegacyProjectArtefactsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Docs Job ID
        /// </summary>
        [Input("docsJobId")]
        public Input<int>? DocsJobId { get; set; }

        /// <summary>
        /// Freshness Job ID
        /// </summary>
        [Input("freshnessJobId")]
        public Input<int>? FreshnessJobId { get; set; }

        /// <summary>
        /// Project ID
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        public LegacyProjectArtefactsArgs()
        {
        }
        public static new LegacyProjectArtefactsArgs Empty => new LegacyProjectArtefactsArgs();
    }

    public sealed class LegacyProjectArtefactsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Docs Job ID
        /// </summary>
        [Input("docsJobId")]
        public Input<int>? DocsJobId { get; set; }

        /// <summary>
        /// Freshness Job ID
        /// </summary>
        [Input("freshnessJobId")]
        public Input<int>? FreshnessJobId { get; set; }

        /// <summary>
        /// Project ID
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        public LegacyProjectArtefactsState()
        {
        }
        public static new LegacyProjectArtefactsState Empty => new LegacyProjectArtefactsState();
    }
}

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
    ///     // NOTE for customers using the LEGACY dbt_cloud provider:
    ///     var myCredential = new Dbtcloud.BigQueryCredential("myCredential", new()
    ///     {
    ///         ProjectId = dbtcloud_project.Dbt_project.Id,
    ///         Dataset = "my_bq_dataset",
    ///         NumThreads = 16,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import dbtcloud:index/bigQueryCredential:BigQueryCredential my_credential "project_id:credential_id"
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import dbtcloud:index/bigQueryCredential:BigQueryCredential my_credential 12345:5678
    /// ```
    /// </summary>
    [DbtcloudResourceType("dbtcloud:index/bigQueryCredential:BigQueryCredential")]
    public partial class BigQueryCredential : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The system BigQuery credential ID
        /// </summary>
        [Output("credentialId")]
        public Output<int> CredentialId { get; private set; } = null!;

        /// <summary>
        /// Default dataset name
        /// </summary>
        [Output("dataset")]
        public Output<string> Dataset { get; private set; } = null!;

        /// <summary>
        /// Whether the BigQuery credential is active
        /// </summary>
        [Output("isActive")]
        public Output<bool?> IsActive { get; private set; } = null!;

        /// <summary>
        /// Number of threads to use
        /// </summary>
        [Output("numThreads")]
        public Output<int> NumThreads { get; private set; } = null!;

        /// <summary>
        /// Project ID to create the BigQuery credential in
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;


        /// <summary>
        /// Create a BigQueryCredential resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BigQueryCredential(string name, BigQueryCredentialArgs args, CustomResourceOptions? options = null)
            : base("dbtcloud:index/bigQueryCredential:BigQueryCredential", name, args ?? new BigQueryCredentialArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BigQueryCredential(string name, Input<string> id, BigQueryCredentialState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/bigQueryCredential:BigQueryCredential", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BigQueryCredential resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BigQueryCredential Get(string name, Input<string> id, BigQueryCredentialState? state = null, CustomResourceOptions? options = null)
        {
            return new BigQueryCredential(name, id, state, options);
        }
    }

    public sealed class BigQueryCredentialArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default dataset name
        /// </summary>
        [Input("dataset", required: true)]
        public Input<string> Dataset { get; set; } = null!;

        /// <summary>
        /// Whether the BigQuery credential is active
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Number of threads to use
        /// </summary>
        [Input("numThreads", required: true)]
        public Input<int> NumThreads { get; set; } = null!;

        /// <summary>
        /// Project ID to create the BigQuery credential in
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        public BigQueryCredentialArgs()
        {
        }
        public static new BigQueryCredentialArgs Empty => new BigQueryCredentialArgs();
    }

    public sealed class BigQueryCredentialState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The system BigQuery credential ID
        /// </summary>
        [Input("credentialId")]
        public Input<int>? CredentialId { get; set; }

        /// <summary>
        /// Default dataset name
        /// </summary>
        [Input("dataset")]
        public Input<string>? Dataset { get; set; }

        /// <summary>
        /// Whether the BigQuery credential is active
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Number of threads to use
        /// </summary>
        [Input("numThreads")]
        public Input<int>? NumThreads { get; set; }

        /// <summary>
        /// Project ID to create the BigQuery credential in
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        public BigQueryCredentialState()
        {
        }
        public static new BigQueryCredentialState Empty => new BigQueryCredentialState();
    }
}

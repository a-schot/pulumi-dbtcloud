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
    ///     // use dbt_cloud_databricks_credential instead of dbtcloud_databricks_credential for the legacy resource names
    ///     // legacy names will be removed from 0.3 onwards
    ///     // when using the Databricks adapter
    ///     var myDatabricksCred = new Dbtcloud.DatabricksCredential("myDatabricksCred", new()
    ///     {
    ///         ProjectId = dbtcloud_project.Dbt_project.Id,
    ///         AdapterId = 123,
    ///         TargetName = "prod",
    ///         Token = "abcdefgh",
    ///         Schema = "my_schema",
    ///         AdapterType = "databricks",
    ///     });
    /// 
    ///     // when using the Spark adapter
    ///     var mySparkCred = new Dbtcloud.DatabricksCredential("mySparkCred", new()
    ///     {
    ///         ProjectId = dbtcloud_project.Dbt_project.Id,
    ///         AdapterId = 456,
    ///         TargetName = "prod",
    ///         Token = "abcdefgh",
    ///         Schema = "my_schema",
    ///         AdapterType = "spark",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Import using a project ID and credential ID found in the URL or via the API.
    /// 
    /// ```sh
    ///  $ pulumi import dbtcloud:index/databricksCredential:DatabricksCredential my_databricks_credential "project_id:credential_id"
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import dbtcloud:index/databricksCredential:DatabricksCredential my_databricks_credential 12345:6789
    /// ```
    /// </summary>
    [DbtcloudResourceType("dbtcloud:index/databricksCredential:DatabricksCredential")]
    public partial class DatabricksCredential : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Databricks adapter ID for the credential
        /// </summary>
        [Output("adapterId")]
        public Output<int> AdapterId { get; private set; } = null!;

        /// <summary>
        /// The type of the adapter (databricks or spark)
        /// </summary>
        [Output("adapterType")]
        public Output<string> AdapterType { get; private set; } = null!;

        /// <summary>
        /// The catalog where to create models (only for the databricks adapter)
        /// </summary>
        [Output("catalog")]
        public Output<string?> Catalog { get; private set; } = null!;

        /// <summary>
        /// The system Databricks credential ID
        /// </summary>
        [Output("credentialId")]
        public Output<int> CredentialId { get; private set; } = null!;

        /// <summary>
        /// Project ID to create the Databricks credential in
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The schema where to create models
        /// </summary>
        [Output("schema")]
        public Output<string> Schema { get; private set; } = null!;

        /// <summary>
        /// Target name
        /// </summary>
        [Output("targetName")]
        public Output<string?> TargetName { get; private set; } = null!;

        /// <summary>
        /// Token for Databricks user
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;


        /// <summary>
        /// Create a DatabricksCredential resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabricksCredential(string name, DatabricksCredentialArgs args, CustomResourceOptions? options = null)
            : base("dbtcloud:index/databricksCredential:DatabricksCredential", name, args ?? new DatabricksCredentialArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabricksCredential(string name, Input<string> id, DatabricksCredentialState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/databricksCredential:DatabricksCredential", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/a-schot/pulumi-dbtcloud/releases/download/v${VERSION}",
                AdditionalSecretOutputs =
                {
                    "token",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DatabricksCredential resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabricksCredential Get(string name, Input<string> id, DatabricksCredentialState? state = null, CustomResourceOptions? options = null)
        {
            return new DatabricksCredential(name, id, state, options);
        }
    }

    public sealed class DatabricksCredentialArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Databricks adapter ID for the credential
        /// </summary>
        [Input("adapterId", required: true)]
        public Input<int> AdapterId { get; set; } = null!;

        /// <summary>
        /// The type of the adapter (databricks or spark)
        /// </summary>
        [Input("adapterType", required: true)]
        public Input<string> AdapterType { get; set; } = null!;

        /// <summary>
        /// The catalog where to create models (only for the databricks adapter)
        /// </summary>
        [Input("catalog")]
        public Input<string>? Catalog { get; set; }

        /// <summary>
        /// Project ID to create the Databricks credential in
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        /// <summary>
        /// The schema where to create models
        /// </summary>
        [Input("schema", required: true)]
        public Input<string> Schema { get; set; } = null!;

        /// <summary>
        /// Target name
        /// </summary>
        [Input("targetName")]
        public Input<string>? TargetName { get; set; }

        [Input("token", required: true)]
        private Input<string>? _token;

        /// <summary>
        /// Token for Databricks user
        /// </summary>
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public DatabricksCredentialArgs()
        {
        }
        public static new DatabricksCredentialArgs Empty => new DatabricksCredentialArgs();
    }

    public sealed class DatabricksCredentialState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Databricks adapter ID for the credential
        /// </summary>
        [Input("adapterId")]
        public Input<int>? AdapterId { get; set; }

        /// <summary>
        /// The type of the adapter (databricks or spark)
        /// </summary>
        [Input("adapterType")]
        public Input<string>? AdapterType { get; set; }

        /// <summary>
        /// The catalog where to create models (only for the databricks adapter)
        /// </summary>
        [Input("catalog")]
        public Input<string>? Catalog { get; set; }

        /// <summary>
        /// The system Databricks credential ID
        /// </summary>
        [Input("credentialId")]
        public Input<int>? CredentialId { get; set; }

        /// <summary>
        /// Project ID to create the Databricks credential in
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// The schema where to create models
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// Target name
        /// </summary>
        [Input("targetName")]
        public Input<string>? TargetName { get; set; }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// Token for Databricks user
        /// </summary>
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public DatabricksCredentialState()
        {
        }
        public static new DatabricksCredentialState Empty => new DatabricksCredentialState();
    }
}

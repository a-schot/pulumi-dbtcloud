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
    ///     var postgresProdCredential = new Dbtcloud.PostgresCredential("postgresProdCredential", new()
    ///     {
    ///         IsActive = true,
    ///         ProjectId = dbtcloud_project.Dbt_project.Id,
    ///         Type = "postgres",
    ///         DefaultSchema = "my_schema",
    ///         Username = "my_username",
    ///         Password = "my_password",
    ///         NumThreads = 16,
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
    ///  $ pulumi import dbtcloud:index/postgresCredential:PostgresCredential my_credential "project_id:credential_id"
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import dbtcloud:index/postgresCredential:PostgresCredential my_credential 12345:6789
    /// ```
    /// </summary>
    [DbtcloudResourceType("dbtcloud:index/postgresCredential:PostgresCredential")]
    public partial class PostgresCredential : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The system Postgres/Redshift/AlloyDB credential ID
        /// </summary>
        [Output("credentialId")]
        public Output<int> CredentialId { get; private set; } = null!;

        /// <summary>
        /// Default schema name
        /// </summary>
        [Output("defaultSchema")]
        public Output<string> DefaultSchema { get; private set; } = null!;

        /// <summary>
        /// Whether the Postgres/Redshift/AlloyDB credential is active
        /// </summary>
        [Output("isActive")]
        public Output<bool?> IsActive { get; private set; } = null!;

        /// <summary>
        /// Number of threads to use
        /// </summary>
        [Output("numThreads")]
        public Output<int?> NumThreads { get; private set; } = null!;

        /// <summary>
        /// Password for Postgres/Redshift/AlloyDB
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Project ID to create the Postgres/Redshift/AlloyDB credential in
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Default schema name
        /// </summary>
        [Output("targetName")]
        public Output<string?> TargetName { get; private set; } = null!;

        /// <summary>
        /// Type of connection. One of (postgres/redshift). Use postgres for alloydb connections
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Username for Postgres/Redshift/AlloyDB
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a PostgresCredential resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PostgresCredential(string name, PostgresCredentialArgs args, CustomResourceOptions? options = null)
            : base("dbtcloud:index/postgresCredential:PostgresCredential", name, args ?? new PostgresCredentialArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PostgresCredential(string name, Input<string> id, PostgresCredentialState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/postgresCredential:PostgresCredential", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github::api.github.com/a-schot/pulumi-dbtcloud",
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PostgresCredential resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PostgresCredential Get(string name, Input<string> id, PostgresCredentialState? state = null, CustomResourceOptions? options = null)
        {
            return new PostgresCredential(name, id, state, options);
        }
    }

    public sealed class PostgresCredentialArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default schema name
        /// </summary>
        [Input("defaultSchema", required: true)]
        public Input<string> DefaultSchema { get; set; } = null!;

        /// <summary>
        /// Whether the Postgres/Redshift/AlloyDB credential is active
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Number of threads to use
        /// </summary>
        [Input("numThreads")]
        public Input<int>? NumThreads { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for Postgres/Redshift/AlloyDB
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Project ID to create the Postgres/Redshift/AlloyDB credential in
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        /// <summary>
        /// Default schema name
        /// </summary>
        [Input("targetName")]
        public Input<string>? TargetName { get; set; }

        /// <summary>
        /// Type of connection. One of (postgres/redshift). Use postgres for alloydb connections
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Username for Postgres/Redshift/AlloyDB
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public PostgresCredentialArgs()
        {
        }
        public static new PostgresCredentialArgs Empty => new PostgresCredentialArgs();
    }

    public sealed class PostgresCredentialState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The system Postgres/Redshift/AlloyDB credential ID
        /// </summary>
        [Input("credentialId")]
        public Input<int>? CredentialId { get; set; }

        /// <summary>
        /// Default schema name
        /// </summary>
        [Input("defaultSchema")]
        public Input<string>? DefaultSchema { get; set; }

        /// <summary>
        /// Whether the Postgres/Redshift/AlloyDB credential is active
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Number of threads to use
        /// </summary>
        [Input("numThreads")]
        public Input<int>? NumThreads { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for Postgres/Redshift/AlloyDB
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Project ID to create the Postgres/Redshift/AlloyDB credential in
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Default schema name
        /// </summary>
        [Input("targetName")]
        public Input<string>? TargetName { get; set; }

        /// <summary>
        /// Type of connection. One of (postgres/redshift). Use postgres for alloydb connections
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Username for Postgres/Redshift/AlloyDB
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public PostgresCredentialState()
        {
        }
        public static new PostgresCredentialState Empty => new PostgresCredentialState();
    }
}

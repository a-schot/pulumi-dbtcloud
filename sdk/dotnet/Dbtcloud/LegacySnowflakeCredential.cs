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
    [DbtcloudResourceType("dbtcloud:index/legacySnowflakeCredential:LegacySnowflakeCredential")]
    public partial class LegacySnowflakeCredential : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The type of Snowflake credential ('password' or 'keypair')
        /// </summary>
        [Output("authType")]
        public Output<string> AuthType { get; private set; } = null!;

        /// <summary>
        /// The system Snowflake credential ID
        /// </summary>
        [Output("credentialId")]
        public Output<int> CredentialId { get; private set; } = null!;

        /// <summary>
        /// Database to connect to
        /// </summary>
        [Output("database")]
        public Output<string?> Database { get; private set; } = null!;

        /// <summary>
        /// Whether the Snowflake credential is active
        /// </summary>
        [Output("isActive")]
        public Output<bool?> IsActive { get; private set; } = null!;

        /// <summary>
        /// Number of threads to use
        /// </summary>
        [Output("numThreads")]
        public Output<int> NumThreads { get; private set; } = null!;

        /// <summary>
        /// Password for Snowflake
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Private key for Snowflake
        /// </summary>
        [Output("privateKey")]
        public Output<string?> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// Private key passphrase for Snowflake
        /// </summary>
        [Output("privateKeyPassphrase")]
        public Output<string?> PrivateKeyPassphrase { get; private set; } = null!;

        /// <summary>
        /// Project ID to create the Snowflake credential in
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Role to assume
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;

        /// <summary>
        /// Default schema name
        /// </summary>
        [Output("schema")]
        public Output<string> Schema { get; private set; } = null!;

        /// <summary>
        /// Username for Snowflake
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;

        /// <summary>
        /// Warehouse to use
        /// </summary>
        [Output("warehouse")]
        public Output<string?> Warehouse { get; private set; } = null!;


        /// <summary>
        /// Create a LegacySnowflakeCredential resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LegacySnowflakeCredential(string name, LegacySnowflakeCredentialArgs args, CustomResourceOptions? options = null)
            : base("dbtcloud:index/legacySnowflakeCredential:LegacySnowflakeCredential", name, args ?? new LegacySnowflakeCredentialArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LegacySnowflakeCredential(string name, Input<string> id, LegacySnowflakeCredentialState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/legacySnowflakeCredential:LegacySnowflakeCredential", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/a-schot/pulumi-dbtcloud",
                AdditionalSecretOutputs =
                {
                    "password",
                    "privateKey",
                    "privateKeyPassphrase",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LegacySnowflakeCredential resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LegacySnowflakeCredential Get(string name, Input<string> id, LegacySnowflakeCredentialState? state = null, CustomResourceOptions? options = null)
        {
            return new LegacySnowflakeCredential(name, id, state, options);
        }
    }

    public sealed class LegacySnowflakeCredentialArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of Snowflake credential ('password' or 'keypair')
        /// </summary>
        [Input("authType", required: true)]
        public Input<string> AuthType { get; set; } = null!;

        /// <summary>
        /// Database to connect to
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Whether the Snowflake credential is active
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Number of threads to use
        /// </summary>
        [Input("numThreads", required: true)]
        public Input<int> NumThreads { get; set; } = null!;

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for Snowflake
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

        [Input("privateKey")]
        private Input<string>? _privateKey;

        /// <summary>
        /// Private key for Snowflake
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKeyPassphrase")]
        private Input<string>? _privateKeyPassphrase;

        /// <summary>
        /// Private key passphrase for Snowflake
        /// </summary>
        public Input<string>? PrivateKeyPassphrase
        {
            get => _privateKeyPassphrase;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKeyPassphrase = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Project ID to create the Snowflake credential in
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        /// <summary>
        /// Role to assume
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Default schema name
        /// </summary>
        [Input("schema", required: true)]
        public Input<string> Schema { get; set; } = null!;

        /// <summary>
        /// Username for Snowflake
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        /// <summary>
        /// Warehouse to use
        /// </summary>
        [Input("warehouse")]
        public Input<string>? Warehouse { get; set; }

        public LegacySnowflakeCredentialArgs()
        {
        }
        public static new LegacySnowflakeCredentialArgs Empty => new LegacySnowflakeCredentialArgs();
    }

    public sealed class LegacySnowflakeCredentialState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of Snowflake credential ('password' or 'keypair')
        /// </summary>
        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        /// <summary>
        /// The system Snowflake credential ID
        /// </summary>
        [Input("credentialId")]
        public Input<int>? CredentialId { get; set; }

        /// <summary>
        /// Database to connect to
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Whether the Snowflake credential is active
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
        /// Password for Snowflake
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

        [Input("privateKey")]
        private Input<string>? _privateKey;

        /// <summary>
        /// Private key for Snowflake
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKeyPassphrase")]
        private Input<string>? _privateKeyPassphrase;

        /// <summary>
        /// Private key passphrase for Snowflake
        /// </summary>
        public Input<string>? PrivateKeyPassphrase
        {
            get => _privateKeyPassphrase;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKeyPassphrase = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Project ID to create the Snowflake credential in
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Role to assume
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Default schema name
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// Username for Snowflake
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        /// <summary>
        /// Warehouse to use
        /// </summary>
        [Input("warehouse")]
        public Input<string>? Warehouse { get; set; }

        public LegacySnowflakeCredentialState()
        {
        }
        public static new LegacySnowflakeCredentialState Empty => new LegacySnowflakeCredentialState();
    }
}

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
    [DbtcloudResourceType("dbtcloud:index/legacyBigqueryConnection:LegacyBigqueryConnection")]
    public partial class LegacyBigqueryConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Application ID for BQ OAuth
        /// </summary>
        [Output("applicationId")]
        public Output<string?> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// The Application Secret for BQ OAuth
        /// </summary>
        [Output("applicationSecret")]
        public Output<string?> ApplicationSecret { get; private set; } = null!;

        /// <summary>
        /// Auth Provider X509 Cert URL for the Service Account
        /// </summary>
        [Output("authProviderX509CertUrl")]
        public Output<string> AuthProviderX509CertUrl { get; private set; } = null!;

        /// <summary>
        /// Auth URI for the Service Account
        /// </summary>
        [Output("authUri")]
        public Output<string> AuthUri { get; private set; } = null!;

        /// <summary>
        /// Service Account email
        /// </summary>
        [Output("clientEmail")]
        public Output<string> ClientEmail { get; private set; } = null!;

        /// <summary>
        /// Client ID of the Service Account
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// Client X509 Cert URL for the Service Account
        /// </summary>
        [Output("clientX509CertUrl")]
        public Output<string> ClientX509CertUrl { get; private set; } = null!;

        /// <summary>
        /// Connection Identifier
        /// </summary>
        [Output("connectionId")]
        public Output<int> ConnectionId { get; private set; } = null!;

        /// <summary>
        /// Dataproc cluster name for PySpark workloads
        /// </summary>
        [Output("dataprocClusterName")]
        public Output<string?> DataprocClusterName { get; private set; } = null!;

        /// <summary>
        /// Google Cloud region for PySpark workloads on Dataproc
        /// </summary>
        [Output("dataprocRegion")]
        public Output<string?> DataprocRegion { get; private set; } = null!;

        /// <summary>
        /// Project to bill for query execution
        /// </summary>
        [Output("executionProject")]
        public Output<string?> ExecutionProject { get; private set; } = null!;

        /// <summary>
        /// GCP project ID
        /// </summary>
        [Output("gcpProjectId")]
        public Output<string> GcpProjectId { get; private set; } = null!;

        /// <summary>
        /// URI for a Google Cloud Storage bucket to host Python code executed via Datapro
        /// </summary>
        [Output("gcsBucket")]
        public Output<string?> GcsBucket { get; private set; } = null!;

        /// <summary>
        /// Whether the connection is active
        /// </summary>
        [Output("isActive")]
        public Output<bool?> IsActive { get; private set; } = null!;

        /// <summary>
        /// Whether the connection is configured for OAuth or not
        /// </summary>
        [Output("isConfiguredForOauth")]
        public Output<bool> IsConfiguredForOauth { get; private set; } = null!;

        /// <summary>
        /// Location to create new Datasets in
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Max number of bytes that can be billed for a given BigQuery query
        /// </summary>
        [Output("maximumBytesBilled")]
        public Output<int?> MaximumBytesBilled { get; private set; } = null!;

        /// <summary>
        /// Connection name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The priority with which to execute BigQuery queries (batch or interactive)
        /// </summary>
        [Output("priority")]
        public Output<string?> Priority { get; private set; } = null!;

        /// <summary>
        /// Private key of the Service Account
        /// </summary>
        [Output("privateKey")]
        public Output<string> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// Private key ID of the Service Account
        /// </summary>
        [Output("privateKeyId")]
        public Output<string> PrivateKeyId { get; private set; } = null!;

        /// <summary>
        /// Project ID to create the connection in
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Number of retries for queries
        /// </summary>
        [Output("retries")]
        public Output<int?> Retries { get; private set; } = null!;

        /// <summary>
        /// Timeout in seconds for queries
        /// </summary>
        [Output("timeoutSeconds")]
        public Output<int> TimeoutSeconds { get; private set; } = null!;

        /// <summary>
        /// Token URI for the Service Account
        /// </summary>
        [Output("tokenUri")]
        public Output<string> TokenUri { get; private set; } = null!;

        /// <summary>
        /// The type of connection
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a LegacyBigqueryConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LegacyBigqueryConnection(string name, LegacyBigqueryConnectionArgs args, CustomResourceOptions? options = null)
            : base("dbtcloud:index/legacyBigqueryConnection:LegacyBigqueryConnection", name, args ?? new LegacyBigqueryConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LegacyBigqueryConnection(string name, Input<string> id, LegacyBigqueryConnectionState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/legacyBigqueryConnection:LegacyBigqueryConnection", name, state, MakeResourceOptions(options, id))
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
                    "applicationId",
                    "applicationSecret",
                    "privateKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LegacyBigqueryConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LegacyBigqueryConnection Get(string name, Input<string> id, LegacyBigqueryConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new LegacyBigqueryConnection(name, id, state, options);
        }
    }

    public sealed class LegacyBigqueryConnectionArgs : global::Pulumi.ResourceArgs
    {
        [Input("applicationId")]
        private Input<string>? _applicationId;

        /// <summary>
        /// The Application ID for BQ OAuth
        /// </summary>
        public Input<string>? ApplicationId
        {
            get => _applicationId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _applicationId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("applicationSecret")]
        private Input<string>? _applicationSecret;

        /// <summary>
        /// The Application Secret for BQ OAuth
        /// </summary>
        public Input<string>? ApplicationSecret
        {
            get => _applicationSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _applicationSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Auth Provider X509 Cert URL for the Service Account
        /// </summary>
        [Input("authProviderX509CertUrl", required: true)]
        public Input<string> AuthProviderX509CertUrl { get; set; } = null!;

        /// <summary>
        /// Auth URI for the Service Account
        /// </summary>
        [Input("authUri", required: true)]
        public Input<string> AuthUri { get; set; } = null!;

        /// <summary>
        /// Service Account email
        /// </summary>
        [Input("clientEmail", required: true)]
        public Input<string> ClientEmail { get; set; } = null!;

        /// <summary>
        /// Client ID of the Service Account
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// Client X509 Cert URL for the Service Account
        /// </summary>
        [Input("clientX509CertUrl", required: true)]
        public Input<string> ClientX509CertUrl { get; set; } = null!;

        /// <summary>
        /// Dataproc cluster name for PySpark workloads
        /// </summary>
        [Input("dataprocClusterName")]
        public Input<string>? DataprocClusterName { get; set; }

        /// <summary>
        /// Google Cloud region for PySpark workloads on Dataproc
        /// </summary>
        [Input("dataprocRegion")]
        public Input<string>? DataprocRegion { get; set; }

        /// <summary>
        /// Project to bill for query execution
        /// </summary>
        [Input("executionProject")]
        public Input<string>? ExecutionProject { get; set; }

        /// <summary>
        /// GCP project ID
        /// </summary>
        [Input("gcpProjectId", required: true)]
        public Input<string> GcpProjectId { get; set; } = null!;

        /// <summary>
        /// URI for a Google Cloud Storage bucket to host Python code executed via Datapro
        /// </summary>
        [Input("gcsBucket")]
        public Input<string>? GcsBucket { get; set; }

        /// <summary>
        /// Whether the connection is active
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Location to create new Datasets in
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Max number of bytes that can be billed for a given BigQuery query
        /// </summary>
        [Input("maximumBytesBilled")]
        public Input<int>? MaximumBytesBilled { get; set; }

        /// <summary>
        /// Connection name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The priority with which to execute BigQuery queries (batch or interactive)
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        [Input("privateKey", required: true)]
        private Input<string>? _privateKey;

        /// <summary>
        /// Private key of the Service Account
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

        /// <summary>
        /// Private key ID of the Service Account
        /// </summary>
        [Input("privateKeyId", required: true)]
        public Input<string> PrivateKeyId { get; set; } = null!;

        /// <summary>
        /// Project ID to create the connection in
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        /// <summary>
        /// Number of retries for queries
        /// </summary>
        [Input("retries")]
        public Input<int>? Retries { get; set; }

        /// <summary>
        /// Timeout in seconds for queries
        /// </summary>
        [Input("timeoutSeconds", required: true)]
        public Input<int> TimeoutSeconds { get; set; } = null!;

        /// <summary>
        /// Token URI for the Service Account
        /// </summary>
        [Input("tokenUri", required: true)]
        public Input<string> TokenUri { get; set; } = null!;

        /// <summary>
        /// The type of connection
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public LegacyBigqueryConnectionArgs()
        {
        }
        public static new LegacyBigqueryConnectionArgs Empty => new LegacyBigqueryConnectionArgs();
    }

    public sealed class LegacyBigqueryConnectionState : global::Pulumi.ResourceArgs
    {
        [Input("applicationId")]
        private Input<string>? _applicationId;

        /// <summary>
        /// The Application ID for BQ OAuth
        /// </summary>
        public Input<string>? ApplicationId
        {
            get => _applicationId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _applicationId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("applicationSecret")]
        private Input<string>? _applicationSecret;

        /// <summary>
        /// The Application Secret for BQ OAuth
        /// </summary>
        public Input<string>? ApplicationSecret
        {
            get => _applicationSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _applicationSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Auth Provider X509 Cert URL for the Service Account
        /// </summary>
        [Input("authProviderX509CertUrl")]
        public Input<string>? AuthProviderX509CertUrl { get; set; }

        /// <summary>
        /// Auth URI for the Service Account
        /// </summary>
        [Input("authUri")]
        public Input<string>? AuthUri { get; set; }

        /// <summary>
        /// Service Account email
        /// </summary>
        [Input("clientEmail")]
        public Input<string>? ClientEmail { get; set; }

        /// <summary>
        /// Client ID of the Service Account
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// Client X509 Cert URL for the Service Account
        /// </summary>
        [Input("clientX509CertUrl")]
        public Input<string>? ClientX509CertUrl { get; set; }

        /// <summary>
        /// Connection Identifier
        /// </summary>
        [Input("connectionId")]
        public Input<int>? ConnectionId { get; set; }

        /// <summary>
        /// Dataproc cluster name for PySpark workloads
        /// </summary>
        [Input("dataprocClusterName")]
        public Input<string>? DataprocClusterName { get; set; }

        /// <summary>
        /// Google Cloud region for PySpark workloads on Dataproc
        /// </summary>
        [Input("dataprocRegion")]
        public Input<string>? DataprocRegion { get; set; }

        /// <summary>
        /// Project to bill for query execution
        /// </summary>
        [Input("executionProject")]
        public Input<string>? ExecutionProject { get; set; }

        /// <summary>
        /// GCP project ID
        /// </summary>
        [Input("gcpProjectId")]
        public Input<string>? GcpProjectId { get; set; }

        /// <summary>
        /// URI for a Google Cloud Storage bucket to host Python code executed via Datapro
        /// </summary>
        [Input("gcsBucket")]
        public Input<string>? GcsBucket { get; set; }

        /// <summary>
        /// Whether the connection is active
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Whether the connection is configured for OAuth or not
        /// </summary>
        [Input("isConfiguredForOauth")]
        public Input<bool>? IsConfiguredForOauth { get; set; }

        /// <summary>
        /// Location to create new Datasets in
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Max number of bytes that can be billed for a given BigQuery query
        /// </summary>
        [Input("maximumBytesBilled")]
        public Input<int>? MaximumBytesBilled { get; set; }

        /// <summary>
        /// Connection name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The priority with which to execute BigQuery queries (batch or interactive)
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        [Input("privateKey")]
        private Input<string>? _privateKey;

        /// <summary>
        /// Private key of the Service Account
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

        /// <summary>
        /// Private key ID of the Service Account
        /// </summary>
        [Input("privateKeyId")]
        public Input<string>? PrivateKeyId { get; set; }

        /// <summary>
        /// Project ID to create the connection in
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Number of retries for queries
        /// </summary>
        [Input("retries")]
        public Input<int>? Retries { get; set; }

        /// <summary>
        /// Timeout in seconds for queries
        /// </summary>
        [Input("timeoutSeconds")]
        public Input<int>? TimeoutSeconds { get; set; }

        /// <summary>
        /// Token URI for the Service Account
        /// </summary>
        [Input("tokenUri")]
        public Input<string>? TokenUri { get; set; }

        /// <summary>
        /// The type of connection
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public LegacyBigqueryConnectionState()
        {
        }
        public static new LegacyBigqueryConnectionState Empty => new LegacyBigqueryConnectionState();
    }
}
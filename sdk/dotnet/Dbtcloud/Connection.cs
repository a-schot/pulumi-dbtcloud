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
    /// Create a Data Warehouse connection for your project in dbt Cloud. The connection will need to be linked to the dbt Cloud project via a `dbtcloud.ProjectConnection` resource.
    /// 
    /// This resource can be used for Databricks, Postgres, Redshift, Snowflake and AlloyDB connections.
    /// For BigQuery, due to the list of fields being very different, you can use the `dbtcloud.BigqueryConnection` resource.
    /// 
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
    ///     // use dbt_cloud_connection instead of dbtcloud_connection for the legacy resource names
    ///     // legacy names will be removed from 0.3 onwards
    ///     var databricks = new Dbtcloud.Connection("databricks", new()
    ///     {
    ///         ProjectId = dbtcloud_project.Dbt_project.Id,
    ///         Type = "adapter",
    ///         Database = "",
    ///         HostName = "my-databricks-host.cloud.databricks.com",
    ///         HttpPath = "/my/path",
    ///         Catalog = "moo",
    ///     });
    /// 
    ///     var redshift = new Dbtcloud.Connection("redshift", new()
    ///     {
    ///         ProjectId = dbtcloud_project.Dbt_project.Id,
    ///         Type = "redshift",
    ///         Database = "my-database",
    ///         Port = 5439,
    ///         HostName = "my-redshift-hostname",
    ///     });
    /// 
    ///     var snowflake = new Dbtcloud.Connection("snowflake", new()
    ///     {
    ///         ProjectId = dbtcloud_project.Dbt_project.Id,
    ///         Type = "snowflake",
    ///         Account = "my-snowflake-account",
    ///         Database = "MY_DATABASE",
    ///         Role = "MY_ROLE",
    ///         Warehouse = "MY_WAREHOUSE",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Import using a project ID and connection ID found in the URL or via the API.
    /// 
    /// ```sh
    ///  $ pulumi import dbtcloud:index/connection:Connection test_connection "project_id:connection_id"
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import dbtcloud:index/connection:Connection test_connection 12345:6789
    /// ```
    /// </summary>
    [DbtcloudResourceType("dbtcloud:index/connection:Connection")]
    public partial class Connection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Account name for the connection
        /// </summary>
        [Output("account")]
        public Output<string?> Account { get; private set; } = null!;

        /// <summary>
        /// Adapter id created for the Databricks connection
        /// </summary>
        [Output("adapterId")]
        public Output<int> AdapterId { get; private set; } = null!;

        /// <summary>
        /// Whether or not the connection should allow client session keep alive
        /// </summary>
        [Output("allowKeepAlive")]
        public Output<bool?> AllowKeepAlive { get; private set; } = null!;

        /// <summary>
        /// Whether or not the connection should allow SSO
        /// </summary>
        [Output("allowSso")]
        public Output<bool?> AllowSso { get; private set; } = null!;

        /// <summary>
        /// Catalog name if Unity Catalog is enabled in your Databricks workspace
        /// </summary>
        [Output("catalog")]
        public Output<string?> Catalog { get; private set; } = null!;

        /// <summary>
        /// Connection Identifier
        /// </summary>
        [Output("connectionId")]
        public Output<int> ConnectionId { get; private set; } = null!;

        /// <summary>
        /// Database name for the connection
        /// </summary>
        [Output("database")]
        public Output<string> Database { get; private set; } = null!;

        /// <summary>
        /// Host name for the connection, including Databricks cluster
        /// </summary>
        [Output("hostName")]
        public Output<string?> HostName { get; private set; } = null!;

        /// <summary>
        /// The HTTP path of the Databricks cluster or SQL warehouse
        /// </summary>
        [Output("httpPath")]
        public Output<string?> HttpPath { get; private set; } = null!;

        /// <summary>
        /// Whether the connection is active
        /// </summary>
        [Output("isActive")]
        public Output<bool?> IsActive { get; private set; } = null!;

        /// <summary>
        /// Connection name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// OAuth client identifier
        /// </summary>
        [Output("oauthClientId")]
        public Output<string?> OauthClientId { get; private set; } = null!;

        /// <summary>
        /// OAuth client secret
        /// </summary>
        [Output("oauthClientSecret")]
        public Output<string?> OauthClientSecret { get; private set; } = null!;

        /// <summary>
        /// Port number to connect via
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// The ID of the PrivateLink connection. This ID can be found using the `privatelink_endpoint` data source
        /// </summary>
        [Output("privateLinkEndpointId")]
        public Output<string?> PrivateLinkEndpointId { get; private set; } = null!;

        /// <summary>
        /// Project ID to create the connection in
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Role name for the connection
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;

        /// <summary>
        /// Whether or not tunneling should be enabled on your database connection
        /// </summary>
        [Output("tunnelEnabled")]
        public Output<bool?> TunnelEnabled { get; private set; } = null!;

        /// <summary>
        /// The type of connection
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Warehouse name for the connection
        /// </summary>
        [Output("warehouse")]
        public Output<string?> Warehouse { get; private set; } = null!;


        /// <summary>
        /// Create a Connection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Connection(string name, ConnectionArgs args, CustomResourceOptions? options = null)
            : base("dbtcloud:index/connection:Connection", name, args ?? new ConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Connection(string name, Input<string> id, ConnectionState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/connection:Connection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Connection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Connection Get(string name, Input<string> id, ConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new Connection(name, id, state, options);
        }
    }

    public sealed class ConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Account name for the connection
        /// </summary>
        [Input("account")]
        public Input<string>? Account { get; set; }

        /// <summary>
        /// Whether or not the connection should allow client session keep alive
        /// </summary>
        [Input("allowKeepAlive")]
        public Input<bool>? AllowKeepAlive { get; set; }

        /// <summary>
        /// Whether or not the connection should allow SSO
        /// </summary>
        [Input("allowSso")]
        public Input<bool>? AllowSso { get; set; }

        /// <summary>
        /// Catalog name if Unity Catalog is enabled in your Databricks workspace
        /// </summary>
        [Input("catalog")]
        public Input<string>? Catalog { get; set; }

        /// <summary>
        /// Database name for the connection
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// Host name for the connection, including Databricks cluster
        /// </summary>
        [Input("hostName")]
        public Input<string>? HostName { get; set; }

        /// <summary>
        /// The HTTP path of the Databricks cluster or SQL warehouse
        /// </summary>
        [Input("httpPath")]
        public Input<string>? HttpPath { get; set; }

        /// <summary>
        /// Whether the connection is active
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Connection name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// OAuth client identifier
        /// </summary>
        [Input("oauthClientId")]
        public Input<string>? OauthClientId { get; set; }

        /// <summary>
        /// OAuth client secret
        /// </summary>
        [Input("oauthClientSecret")]
        public Input<string>? OauthClientSecret { get; set; }

        /// <summary>
        /// Port number to connect via
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The ID of the PrivateLink connection. This ID can be found using the `privatelink_endpoint` data source
        /// </summary>
        [Input("privateLinkEndpointId")]
        public Input<string>? PrivateLinkEndpointId { get; set; }

        /// <summary>
        /// Project ID to create the connection in
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        /// <summary>
        /// Role name for the connection
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Whether or not tunneling should be enabled on your database connection
        /// </summary>
        [Input("tunnelEnabled")]
        public Input<bool>? TunnelEnabled { get; set; }

        /// <summary>
        /// The type of connection
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Warehouse name for the connection
        /// </summary>
        [Input("warehouse")]
        public Input<string>? Warehouse { get; set; }

        public ConnectionArgs()
        {
        }
        public static new ConnectionArgs Empty => new ConnectionArgs();
    }

    public sealed class ConnectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Account name for the connection
        /// </summary>
        [Input("account")]
        public Input<string>? Account { get; set; }

        /// <summary>
        /// Adapter id created for the Databricks connection
        /// </summary>
        [Input("adapterId")]
        public Input<int>? AdapterId { get; set; }

        /// <summary>
        /// Whether or not the connection should allow client session keep alive
        /// </summary>
        [Input("allowKeepAlive")]
        public Input<bool>? AllowKeepAlive { get; set; }

        /// <summary>
        /// Whether or not the connection should allow SSO
        /// </summary>
        [Input("allowSso")]
        public Input<bool>? AllowSso { get; set; }

        /// <summary>
        /// Catalog name if Unity Catalog is enabled in your Databricks workspace
        /// </summary>
        [Input("catalog")]
        public Input<string>? Catalog { get; set; }

        /// <summary>
        /// Connection Identifier
        /// </summary>
        [Input("connectionId")]
        public Input<int>? ConnectionId { get; set; }

        /// <summary>
        /// Database name for the connection
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Host name for the connection, including Databricks cluster
        /// </summary>
        [Input("hostName")]
        public Input<string>? HostName { get; set; }

        /// <summary>
        /// The HTTP path of the Databricks cluster or SQL warehouse
        /// </summary>
        [Input("httpPath")]
        public Input<string>? HttpPath { get; set; }

        /// <summary>
        /// Whether the connection is active
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Connection name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// OAuth client identifier
        /// </summary>
        [Input("oauthClientId")]
        public Input<string>? OauthClientId { get; set; }

        /// <summary>
        /// OAuth client secret
        /// </summary>
        [Input("oauthClientSecret")]
        public Input<string>? OauthClientSecret { get; set; }

        /// <summary>
        /// Port number to connect via
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The ID of the PrivateLink connection. This ID can be found using the `privatelink_endpoint` data source
        /// </summary>
        [Input("privateLinkEndpointId")]
        public Input<string>? PrivateLinkEndpointId { get; set; }

        /// <summary>
        /// Project ID to create the connection in
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Role name for the connection
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Whether or not tunneling should be enabled on your database connection
        /// </summary>
        [Input("tunnelEnabled")]
        public Input<bool>? TunnelEnabled { get; set; }

        /// <summary>
        /// The type of connection
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Warehouse name for the connection
        /// </summary>
        [Input("warehouse")]
        public Input<string>? Warehouse { get; set; }

        public ConnectionState()
        {
        }
        public static new ConnectionState Empty => new ConnectionState();
    }
}

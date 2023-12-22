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
    /// The provider type for the dbtcloud package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [DbtcloudResourceType("pulumi:providers:dbtcloud")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// URL for your dbt Cloud deployment. Instead of setting the parameter, you can set the environment variable
        /// `DBT_CLOUD_HOST_URL` - Defaults to https://cloud.getdbt.com/api
        /// </summary>
        [Output("hostUrl")]
        public Output<string?> HostUrl { get; private set; } = null!;

        /// <summary>
        /// API token for your dbt Cloud. Instead of setting the parameter, you can set the environment variable `DBT_CLOUD_TOKEN`
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs args, CustomResourceOptions? options = null)
            : base("dbtcloud", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
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
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Account identifier for your dbt Cloud implementation. Instead of setting the parameter, you can set the environment
        /// variable `DBT_CLOUD_ACCOUNT_ID`
        /// </summary>
        [Input("accountId", required: true, json: true)]
        public Input<int> AccountId { get; set; } = null!;

        /// <summary>
        /// URL for your dbt Cloud deployment. Instead of setting the parameter, you can set the environment variable
        /// `DBT_CLOUD_HOST_URL` - Defaults to https://cloud.getdbt.com/api
        /// </summary>
        [Input("hostUrl")]
        public Input<string>? HostUrl { get; set; }

        /// <summary>
        /// API token for your dbt Cloud. Instead of setting the parameter, you can set the environment variable `DBT_CLOUD_TOKEN`
        /// </summary>
        [Input("token", required: true)]
        public Input<string> Token { get; set; } = null!;

        public ProviderArgs()
        {
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
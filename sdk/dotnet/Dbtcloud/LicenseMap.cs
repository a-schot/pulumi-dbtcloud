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
    ///     // Developer license group mapping
    ///     var devLicenseMap = new Dbtcloud.LicenseMap("devLicenseMap", new()
    ///     {
    ///         LicenseType = "developer",
    ///         SsoLicenseMappingGroups = new[]
    ///         {
    ///             "DEV-SSO-GROUP",
    ///         },
    ///     });
    /// 
    ///     // Read-only license mapping
    ///     var readOnlyLicenseMap = new Dbtcloud.LicenseMap("readOnlyLicenseMap", new()
    ///     {
    ///         LicenseType = "read_only",
    ///         SsoLicenseMappingGroups = new[]
    ///         {
    ///             "READ-ONLY-SSO-GROUP",
    ///         },
    ///     });
    /// 
    ///     // IT license mapping
    ///     var itLicenseMap = new Dbtcloud.LicenseMap("itLicenseMap", new()
    ///     {
    ///         LicenseType = "it",
    ///         SsoLicenseMappingGroups = new[]
    ///         {
    ///             "IT-SSO-GROUP",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Import using a license map ID found in the URL or via the API.
    /// 
    /// ```sh
    ///  $ pulumi import dbtcloud:index/licenseMap:LicenseMap test_license_map "license_map_id"
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import dbtcloud:index/licenseMap:LicenseMap test_license_map 12345
    /// ```
    /// </summary>
    [DbtcloudResourceType("dbtcloud:index/licenseMap:LicenseMap")]
    public partial class LicenseMap : global::Pulumi.CustomResource
    {
        /// <summary>
        /// License type
        /// </summary>
        [Output("licenseType")]
        public Output<string> LicenseType { get; private set; } = null!;

        /// <summary>
        /// SSO license mapping group names for this group
        /// </summary>
        [Output("ssoLicenseMappingGroups")]
        public Output<ImmutableArray<string>> SsoLicenseMappingGroups { get; private set; } = null!;


        /// <summary>
        /// Create a LicenseMap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LicenseMap(string name, LicenseMapArgs args, CustomResourceOptions? options = null)
            : base("dbtcloud:index/licenseMap:LicenseMap", name, args ?? new LicenseMapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LicenseMap(string name, Input<string> id, LicenseMapState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/licenseMap:LicenseMap", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LicenseMap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LicenseMap Get(string name, Input<string> id, LicenseMapState? state = null, CustomResourceOptions? options = null)
        {
            return new LicenseMap(name, id, state, options);
        }
    }

    public sealed class LicenseMapArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// License type
        /// </summary>
        [Input("licenseType", required: true)]
        public Input<string> LicenseType { get; set; } = null!;

        [Input("ssoLicenseMappingGroups")]
        private InputList<string>? _ssoLicenseMappingGroups;

        /// <summary>
        /// SSO license mapping group names for this group
        /// </summary>
        public InputList<string> SsoLicenseMappingGroups
        {
            get => _ssoLicenseMappingGroups ?? (_ssoLicenseMappingGroups = new InputList<string>());
            set => _ssoLicenseMappingGroups = value;
        }

        public LicenseMapArgs()
        {
        }
        public static new LicenseMapArgs Empty => new LicenseMapArgs();
    }

    public sealed class LicenseMapState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// License type
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        [Input("ssoLicenseMappingGroups")]
        private InputList<string>? _ssoLicenseMappingGroups;

        /// <summary>
        /// SSO license mapping group names for this group
        /// </summary>
        public InputList<string> SsoLicenseMappingGroups
        {
            get => _ssoLicenseMappingGroups ?? (_ssoLicenseMappingGroups = new InputList<string>());
            set => _ssoLicenseMappingGroups = value;
        }

        public LicenseMapState()
        {
        }
        public static new LicenseMapState Empty => new LicenseMapState();
    }
}

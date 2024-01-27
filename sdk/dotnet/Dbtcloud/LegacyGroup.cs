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
    [DbtcloudResourceType("dbtcloud:index/legacyGroup:LegacyGroup")]
    public partial class LegacyGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether or not to assign this group to users by default
        /// </summary>
        [Output("assignByDefault")]
        public Output<bool?> AssignByDefault { get; private set; } = null!;

        [Output("groupPermissions")]
        public Output<ImmutableArray<Outputs.LegacyGroupGroupPermission>> GroupPermissions { get; private set; } = null!;

        /// <summary>
        /// Whether the group is active
        /// </summary>
        [Output("isActive")]
        public Output<bool?> IsActive { get; private set; } = null!;

        /// <summary>
        /// Group name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// SSO mapping group names for this group
        /// </summary>
        [Output("ssoMappingGroups")]
        public Output<ImmutableArray<string>> SsoMappingGroups { get; private set; } = null!;


        /// <summary>
        /// Create a LegacyGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LegacyGroup(string name, LegacyGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/legacyGroup:LegacyGroup", name, args ?? new LegacyGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LegacyGroup(string name, Input<string> id, LegacyGroupState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/legacyGroup:LegacyGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LegacyGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LegacyGroup Get(string name, Input<string> id, LegacyGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new LegacyGroup(name, id, state, options);
        }
    }

    public sealed class LegacyGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not to assign this group to users by default
        /// </summary>
        [Input("assignByDefault")]
        public Input<bool>? AssignByDefault { get; set; }

        [Input("groupPermissions")]
        private InputList<Inputs.LegacyGroupGroupPermissionArgs>? _groupPermissions;
        public InputList<Inputs.LegacyGroupGroupPermissionArgs> GroupPermissions
        {
            get => _groupPermissions ?? (_groupPermissions = new InputList<Inputs.LegacyGroupGroupPermissionArgs>());
            set => _groupPermissions = value;
        }

        /// <summary>
        /// Whether the group is active
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Group name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("ssoMappingGroups")]
        private InputList<string>? _ssoMappingGroups;

        /// <summary>
        /// SSO mapping group names for this group
        /// </summary>
        public InputList<string> SsoMappingGroups
        {
            get => _ssoMappingGroups ?? (_ssoMappingGroups = new InputList<string>());
            set => _ssoMappingGroups = value;
        }

        public LegacyGroupArgs()
        {
        }
        public static new LegacyGroupArgs Empty => new LegacyGroupArgs();
    }

    public sealed class LegacyGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not to assign this group to users by default
        /// </summary>
        [Input("assignByDefault")]
        public Input<bool>? AssignByDefault { get; set; }

        [Input("groupPermissions")]
        private InputList<Inputs.LegacyGroupGroupPermissionGetArgs>? _groupPermissions;
        public InputList<Inputs.LegacyGroupGroupPermissionGetArgs> GroupPermissions
        {
            get => _groupPermissions ?? (_groupPermissions = new InputList<Inputs.LegacyGroupGroupPermissionGetArgs>());
            set => _groupPermissions = value;
        }

        /// <summary>
        /// Whether the group is active
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// Group name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("ssoMappingGroups")]
        private InputList<string>? _ssoMappingGroups;

        /// <summary>
        /// SSO mapping group names for this group
        /// </summary>
        public InputList<string> SsoMappingGroups
        {
            get => _ssoMappingGroups ?? (_ssoMappingGroups = new InputList<string>());
            set => _ssoMappingGroups = value;
        }

        public LegacyGroupState()
        {
        }
        public static new LegacyGroupState Empty => new LegacyGroupState();
    }
}

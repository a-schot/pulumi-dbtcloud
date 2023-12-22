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
    /// *Note*: Groups currently do not support updates, as per both the API and the UI.
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
    ///     // use dbt_cloud_group instead of dbtcloud_group for the legacy resource names
    ///     // legacy names will be removed from 0.3 onwards
    ///     var tfGroup1 = new Dbtcloud.Group("tfGroup1", new()
    ///     {
    ///         GroupPermissions = new[]
    ///         {
    ///             new Dbtcloud.Inputs.GroupGroupPermissionArgs
    ///             {
    ///                 PermissionSet = "member",
    ///                 AllProjects = true,
    ///             },
    ///             new Dbtcloud.Inputs.GroupGroupPermissionArgs
    ///             {
    ///                 PermissionSet = "developer",
    ///                 AllProjects = false,
    ///                 ProjectId = dbtcloud_project.Dbt_project.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Import using a group ID found in the URL or via the API.
    /// 
    /// ```sh
    ///  $ pulumi import dbtcloud:index/group:Group test_group "group_id"
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import dbtcloud:index/group:Group test_group 12345
    /// ```
    /// </summary>
    [DbtcloudResourceType("dbtcloud:index/group:Group")]
    public partial class Group : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether or not to assign this group to users by default
        /// </summary>
        [Output("assignByDefault")]
        public Output<bool?> AssignByDefault { get; private set; } = null!;

        [Output("groupPermissions")]
        public Output<ImmutableArray<Outputs.GroupGroupPermission>> GroupPermissions { get; private set; } = null!;

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
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs? args = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/group:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/group:Group", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not to assign this group to users by default
        /// </summary>
        [Input("assignByDefault")]
        public Input<bool>? AssignByDefault { get; set; }

        [Input("groupPermissions")]
        private InputList<Inputs.GroupGroupPermissionArgs>? _groupPermissions;
        public InputList<Inputs.GroupGroupPermissionArgs> GroupPermissions
        {
            get => _groupPermissions ?? (_groupPermissions = new InputList<Inputs.GroupGroupPermissionArgs>());
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

        public GroupArgs()
        {
        }
        public static new GroupArgs Empty => new GroupArgs();
    }

    public sealed class GroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not to assign this group to users by default
        /// </summary>
        [Input("assignByDefault")]
        public Input<bool>? AssignByDefault { get; set; }

        [Input("groupPermissions")]
        private InputList<Inputs.GroupGroupPermissionGetArgs>? _groupPermissions;
        public InputList<Inputs.GroupGroupPermissionGetArgs> GroupPermissions
        {
            get => _groupPermissions ?? (_groupPermissions = new InputList<Inputs.GroupGroupPermissionGetArgs>());
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

        public GroupState()
        {
        }
        public static new GroupState Empty => new GroupState();
    }
}
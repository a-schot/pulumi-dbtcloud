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
    /// Assigns a set of dbt Cloud groups to a given User ID.
    /// 
    /// &gt; If additional groups were assigned manually in dbt Cloud, they will be removed. The full list of groups need to be provided as config.
    /// 
    /// &gt; This resource does not currently support deletion (e.g. a deleted resource will stay as-is in dbt Cloud).
    /// This is intentional in order to prevent accidental deletion of all users groups assigned to a user.
    /// If you would like a different behavior, please open an issue on GitHub. To remove all groups for a user, set "group_ids" to the empty set "[]".
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
    ///     // we can assign groups to users
    ///     var myUserGroups = new Dbtcloud.UserGroups("myUserGroups", new()
    ///     {
    ///         UserId = data.Dbtcloud_user.My_user.Id,
    ///         GroupIds = new[]
    ///         {
    ///             1234,
    ///             dbtcloud_group.My_group.Id,
    ///             local.My_group_id,
    ///         },
    ///     });
    /// 
    ///     // as Delete is not handled currently, by design, removing all groups from a user can be done with
    ///     var myOtherUserGroups = new Dbtcloud.UserGroups("myOtherUserGroups", new()
    ///     {
    ///         UserId = 123456,
    ///         GroupIds = new[] {},
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Import using the User ID
    /// 
    /// The User ID can be retrieved from the dbt Cloud UI or with the data source dbtcloud_user
    /// 
    /// ```sh
    /// $ pulumi import dbtcloud:index/userGroups:UserGroups my_user_groups "user_id"
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import dbtcloud:index/userGroups:UserGroups my_user_groups 123456
    /// ```
    /// </summary>
    [DbtcloudResourceType("dbtcloud:index/userGroups:UserGroups")]
    public partial class UserGroups : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IDs of the groups to assign to the user. If additional groups were assigned manually in dbt Cloud, they will be removed.
        /// </summary>
        [Output("groupIds")]
        public Output<ImmutableArray<int>> GroupIds { get; private set; } = null!;

        /// <summary>
        /// The internal ID of a dbt Cloud user
        /// </summary>
        [Output("userId")]
        public Output<int> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a UserGroups resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserGroups(string name, UserGroupsArgs args, CustomResourceOptions? options = null)
            : base("dbtcloud:index/userGroups:UserGroups", name, args ?? new UserGroupsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserGroups(string name, Input<string> id, UserGroupsState? state = null, CustomResourceOptions? options = null)
            : base("dbtcloud:index/userGroups:UserGroups", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserGroups resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserGroups Get(string name, Input<string> id, UserGroupsState? state = null, CustomResourceOptions? options = null)
        {
            return new UserGroups(name, id, state, options);
        }
    }

    public sealed class UserGroupsArgs : global::Pulumi.ResourceArgs
    {
        [Input("groupIds", required: true)]
        private InputList<int>? _groupIds;

        /// <summary>
        /// IDs of the groups to assign to the user. If additional groups were assigned manually in dbt Cloud, they will be removed.
        /// </summary>
        public InputList<int> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<int>());
            set => _groupIds = value;
        }

        /// <summary>
        /// The internal ID of a dbt Cloud user
        /// </summary>
        [Input("userId", required: true)]
        public Input<int> UserId { get; set; } = null!;

        public UserGroupsArgs()
        {
        }
        public static new UserGroupsArgs Empty => new UserGroupsArgs();
    }

    public sealed class UserGroupsState : global::Pulumi.ResourceArgs
    {
        [Input("groupIds")]
        private InputList<int>? _groupIds;

        /// <summary>
        /// IDs of the groups to assign to the user. If additional groups were assigned manually in dbt Cloud, they will be removed.
        /// </summary>
        public InputList<int> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<int>());
            set => _groupIds = value;
        }

        /// <summary>
        /// The internal ID of a dbt Cloud user
        /// </summary>
        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public UserGroupsState()
        {
        }
        public static new UserGroupsState Empty => new UserGroupsState();
    }
}

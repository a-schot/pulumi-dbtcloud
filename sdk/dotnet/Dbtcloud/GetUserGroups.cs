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
    public static class GetUserGroups
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Dbtcloud = Pulumi.Dbtcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myUserGroups = Dbtcloud.GetUserGroups.Invoke(new()
        ///     {
        ///         UserId = 12345,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUserGroupsResult> InvokeAsync(GetUserGroupsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserGroupsResult>("dbtcloud:index/getUserGroups:getUserGroups", args ?? new GetUserGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Dbtcloud = Pulumi.Dbtcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myUserGroups = Dbtcloud.GetUserGroups.Invoke(new()
        ///     {
        ///         UserId = 12345,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetUserGroupsResult> Invoke(GetUserGroupsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserGroupsResult>("dbtcloud:index/getUserGroups:getUserGroups", args ?? new GetUserGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserGroupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the user
        /// </summary>
        [Input("userId", required: true)]
        public int UserId { get; set; }

        public GetUserGroupsArgs()
        {
        }
        public static new GetUserGroupsArgs Empty => new GetUserGroupsArgs();
    }

    public sealed class GetUserGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the user
        /// </summary>
        [Input("userId", required: true)]
        public Input<int> UserId { get; set; } = null!;

        public GetUserGroupsInvokeArgs()
        {
        }
        public static new GetUserGroupsInvokeArgs Empty => new GetUserGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserGroupsResult
    {
        /// <summary>
        /// IDs of the groups assigned to the user
        /// </summary>
        public readonly ImmutableArray<int> GroupIds;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// ID of the user
        /// </summary>
        public readonly int UserId;

        [OutputConstructor]
        private GetUserGroupsResult(
            ImmutableArray<int> groupIds,

            string id,

            int userId)
        {
            GroupIds = groupIds;
            Id = id;
            UserId = userId;
        }
    }
}

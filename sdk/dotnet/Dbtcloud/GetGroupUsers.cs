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
    public static class GetGroupUsers
    {
        /// <summary>
        /// Returns a list of users assigned to a specific dbt Cloud group
        /// 
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
        ///     var myGroupUsers = Dbtcloud.GetGroupUsers.Invoke(new()
        ///     {
        ///         GroupId = 1234,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGroupUsersResult> InvokeAsync(GetGroupUsersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupUsersResult>("dbtcloud:index/getGroupUsers:getGroupUsers", args ?? new GetGroupUsersArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a list of users assigned to a specific dbt Cloud group
        /// 
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
        ///     var myGroupUsers = Dbtcloud.GetGroupUsers.Invoke(new()
        ///     {
        ///         GroupId = 1234,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetGroupUsersResult> Invoke(GetGroupUsersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupUsersResult>("dbtcloud:index/getGroupUsers:getGroupUsers", args ?? new GetGroupUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupUsersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the group
        /// </summary>
        [Input("groupId", required: true)]
        public int GroupId { get; set; }

        public GetGroupUsersArgs()
        {
        }
        public static new GetGroupUsersArgs Empty => new GetGroupUsersArgs();
    }

    public sealed class GetGroupUsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the group
        /// </summary>
        [Input("groupId", required: true)]
        public Input<int> GroupId { get; set; } = null!;

        public GetGroupUsersInvokeArgs()
        {
        }
        public static new GetGroupUsersInvokeArgs Empty => new GetGroupUsersInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupUsersResult
    {
        /// <summary>
        /// ID of the group
        /// </summary>
        public readonly int GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of users (map of ID and email) in the group
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupUsersUserResult> Users;

        [OutputConstructor]
        private GetGroupUsersResult(
            int groupId,

            string id,

            ImmutableArray<Outputs.GetGroupUsersUserResult> users)
        {
            GroupId = groupId;
            Id = id;
            Users = users;
        }
    }
}

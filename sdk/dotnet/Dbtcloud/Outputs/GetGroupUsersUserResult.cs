// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ASchot.Pulumi.Dbtcloud.Outputs
{

    [OutputType]
    public sealed class GetGroupUsersUserResult
    {
        public readonly string Email;
        public readonly int Id;

        [OutputConstructor]
        private GetGroupUsersUserResult(
            string email,

            int id)
        {
            Email = email;
            Id = id;
        }
    }
}
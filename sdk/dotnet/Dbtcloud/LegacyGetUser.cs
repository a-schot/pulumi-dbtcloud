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
    public static class LegacyGetUser
    {
        public static Task<LegacyGetUserResult> InvokeAsync(LegacyGetUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<LegacyGetUserResult>("dbtcloud:index/legacyGetUser:LegacyGetUser", args ?? new LegacyGetUserArgs(), options.WithDefaults());

        public static Output<LegacyGetUserResult> Invoke(LegacyGetUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<LegacyGetUserResult>("dbtcloud:index/legacyGetUser:LegacyGetUser", args ?? new LegacyGetUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class LegacyGetUserArgs : global::Pulumi.InvokeArgs
    {
        [Input("email", required: true)]
        public string Email { get; set; } = null!;

        public LegacyGetUserArgs()
        {
        }
        public static new LegacyGetUserArgs Empty => new LegacyGetUserArgs();
    }

    public sealed class LegacyGetUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        public LegacyGetUserInvokeArgs()
        {
        }
        public static new LegacyGetUserInvokeArgs Empty => new LegacyGetUserInvokeArgs();
    }


    [OutputType]
    public sealed class LegacyGetUserResult
    {
        public readonly string Email;
        public readonly int Id;

        [OutputConstructor]
        private LegacyGetUserResult(
            string email,

            int id)
        {
            Email = email;
            Id = id;
        }
    }
}
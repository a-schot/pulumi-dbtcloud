// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ASchot.Pulumi.Dbtcloud.Inputs
{

    public sealed class LegacyServiceTokenServiceTokenPermissionArgs : global::Pulumi.ResourceArgs
    {
        [Input("allProjects", required: true)]
        public Input<bool> AllProjects { get; set; } = null!;

        [Input("permissionSet", required: true)]
        public Input<string> PermissionSet { get; set; } = null!;

        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        public LegacyServiceTokenServiceTokenPermissionArgs()
        {
        }
        public static new LegacyServiceTokenServiceTokenPermissionArgs Empty => new LegacyServiceTokenServiceTokenPermissionArgs();
    }
}
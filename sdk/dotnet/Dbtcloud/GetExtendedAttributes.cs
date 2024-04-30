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
    public static class GetExtendedAttributes
    {
        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Dbtcloud = Pulumi.Dbtcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myExtendedAttributes = Dbtcloud.GetExtendedAttributes.Invoke(new()
        ///     {
        ///         ExtendedAttributesId = 12345,
        ///         ProjectId = 6789,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetExtendedAttributesResult> InvokeAsync(GetExtendedAttributesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetExtendedAttributesResult>("dbtcloud:index/getExtendedAttributes:getExtendedAttributes", args ?? new GetExtendedAttributesArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Dbtcloud = Pulumi.Dbtcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myExtendedAttributes = Dbtcloud.GetExtendedAttributes.Invoke(new()
        ///     {
        ///         ExtendedAttributesId = 12345,
        ///         ProjectId = 6789,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetExtendedAttributesResult> Invoke(GetExtendedAttributesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetExtendedAttributesResult>("dbtcloud:index/getExtendedAttributes:getExtendedAttributes", args ?? new GetExtendedAttributesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetExtendedAttributesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the extended attributes
        /// </summary>
        [Input("extendedAttributesId", required: true)]
        public int ExtendedAttributesId { get; set; }

        /// <summary>
        /// Project ID the extended attributes refers to
        /// </summary>
        [Input("projectId", required: true)]
        public int ProjectId { get; set; }

        public GetExtendedAttributesArgs()
        {
        }
        public static new GetExtendedAttributesArgs Empty => new GetExtendedAttributesArgs();
    }

    public sealed class GetExtendedAttributesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the extended attributes
        /// </summary>
        [Input("extendedAttributesId", required: true)]
        public Input<int> ExtendedAttributesId { get; set; } = null!;

        /// <summary>
        /// Project ID the extended attributes refers to
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        public GetExtendedAttributesInvokeArgs()
        {
        }
        public static new GetExtendedAttributesInvokeArgs Empty => new GetExtendedAttributesInvokeArgs();
    }


    [OutputType]
    public sealed class GetExtendedAttributesResult
    {
        /// <summary>
        /// A JSON string listing the extended attributes mapping
        /// </summary>
        public readonly string ExtendedAttributes;
        /// <summary>
        /// ID of the extended attributes
        /// </summary>
        public readonly int ExtendedAttributesId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Project ID the extended attributes refers to
        /// </summary>
        public readonly int ProjectId;
        /// <summary>
        /// The state of the extended attributes (1 = active, 2 = inactive)
        /// </summary>
        public readonly int State;

        [OutputConstructor]
        private GetExtendedAttributesResult(
            string extendedAttributes,

            int extendedAttributesId,

            string id,

            int projectId,

            int state)
        {
            ExtendedAttributes = extendedAttributes;
            ExtendedAttributesId = extendedAttributesId;
            Id = id;
            ProjectId = projectId;
            State = state;
        }
    }
}

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
    public static class GetAzureDevOpsProject
    {
        /// <summary>
        /// Use this data source to retrieve the ID of an Azure Dev Ops project 
        /// based on its name.
        /// 		
        /// This data source requires connecting with a user token and doesn't work with a service token.
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
        ///     var myAdoProject = Dbtcloud.GetAzureDevOpsProject.Invoke(new()
        ///     {
        ///         Name = "my-project-name",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAzureDevOpsProjectResult> InvokeAsync(GetAzureDevOpsProjectArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAzureDevOpsProjectResult>("dbtcloud:index/getAzureDevOpsProject:getAzureDevOpsProject", args ?? new GetAzureDevOpsProjectArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve the ID of an Azure Dev Ops project 
        /// based on its name.
        /// 		
        /// This data source requires connecting with a user token and doesn't work with a service token.
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
        ///     var myAdoProject = Dbtcloud.GetAzureDevOpsProject.Invoke(new()
        ///     {
        ///         Name = "my-project-name",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAzureDevOpsProjectResult> Invoke(GetAzureDevOpsProjectInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAzureDevOpsProjectResult>("dbtcloud:index/getAzureDevOpsProject:getAzureDevOpsProject", args ?? new GetAzureDevOpsProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAzureDevOpsProjectArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the ADO project
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetAzureDevOpsProjectArgs()
        {
        }
        public static new GetAzureDevOpsProjectArgs Empty => new GetAzureDevOpsProjectArgs();
    }

    public sealed class GetAzureDevOpsProjectInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the ADO project
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetAzureDevOpsProjectInvokeArgs()
        {
        }
        public static new GetAzureDevOpsProjectInvokeArgs Empty => new GetAzureDevOpsProjectInvokeArgs();
    }


    [OutputType]
    public sealed class GetAzureDevOpsProjectResult
    {
        /// <summary>
        /// The internal Azure Dev Ops ID of the ADO Project
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the ADO project
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The URL of the ADO project
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetAzureDevOpsProjectResult(
            string id,

            string name,

            string url)
        {
            Id = id;
            Name = name;
            Url = url;
        }
    }
}

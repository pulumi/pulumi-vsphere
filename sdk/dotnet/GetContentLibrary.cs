// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetContentLibrary
    {
        /// <summary>
        /// The `vsphere.ContentLibrary` data source can be used to discover the ID of a Content Library.
        /// 
        /// &gt; **NOTE:** This resource requires vCenter and is not available on direct ESXi
        /// connections.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using VSphere = Pulumi.VSphere;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var library = Output.Create(VSphere.GetContentLibrary.InvokeAsync(new VSphere.GetContentLibraryArgs
        ///         {
        ///             Name = "Content Library Test",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetContentLibraryResult> InvokeAsync(GetContentLibraryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetContentLibraryResult>("vsphere:index/getContentLibrary:getContentLibrary", args ?? new GetContentLibraryArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.ContentLibrary` data source can be used to discover the ID of a Content Library.
        /// 
        /// &gt; **NOTE:** This resource requires vCenter and is not available on direct ESXi
        /// connections.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using VSphere = Pulumi.VSphere;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var library = Output.Create(VSphere.GetContentLibrary.InvokeAsync(new VSphere.GetContentLibraryArgs
        ///         {
        ///             Name = "Content Library Test",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetContentLibraryResult> Invoke(GetContentLibraryInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetContentLibraryResult>("vsphere:index/getContentLibrary:getContentLibrary", args ?? new GetContentLibraryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContentLibraryArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Content Library.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetContentLibraryArgs()
        {
        }
    }

    public sealed class GetContentLibraryInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Content Library.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetContentLibraryInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetContentLibraryResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetContentLibraryResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}

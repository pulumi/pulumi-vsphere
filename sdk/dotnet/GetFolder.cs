// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetFolder
    {
        /// <summary>
        /// The `vsphere.Folder` data source can be used to get the general attributes of a
        /// vSphere inventory folder. Paths are absolute and include must include the
        /// datacenter.  
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
        ///         var folder = Output.Create(VSphere.GetFolder.InvokeAsync(new VSphere.GetFolderArgs
        ///         {
        ///             Path = "/dc-01/datastore-01/folder-01",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFolderResult> InvokeAsync(GetFolderArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFolderResult>("vsphere:index/getFolder:getFolder", args ?? new GetFolderArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.Folder` data source can be used to get the general attributes of a
        /// vSphere inventory folder. Paths are absolute and include must include the
        /// datacenter.  
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
        ///         var folder = Output.Create(VSphere.GetFolder.InvokeAsync(new VSphere.GetFolderArgs
        ///         {
        ///             Path = "/dc-01/datastore-01/folder-01",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFolderResult> Invoke(GetFolderInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFolderResult>("vsphere:index/getFolder:getFolder", args ?? new GetFolderInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFolderArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The absolute path of the folder. For example, given a
        /// default datacenter of `default-dc`, a folder of type `vm`, and a folder name
        /// of `test-folder`, the resulting path would be
        /// `/default-dc/vm/test-folder`. The valid folder types to be used in
        /// the path are: `vm`, `host`, `datacenter`, `datastore`, or `network`.
        /// </summary>
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        public GetFolderArgs()
        {
        }
    }

    public sealed class GetFolderInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The absolute path of the folder. For example, given a
        /// default datacenter of `default-dc`, a folder of type `vm`, and a folder name
        /// of `test-folder`, the resulting path would be
        /// `/default-dc/vm/test-folder`. The valid folder types to be used in
        /// the path are: `vm`, `host`, `datacenter`, `datastore`, or `network`.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public GetFolderInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFolderResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Path;

        [OutputConstructor]
        private GetFolderResult(
            string id,

            string path)
        {
            Id = id;
            Path = path;
        }
    }
}

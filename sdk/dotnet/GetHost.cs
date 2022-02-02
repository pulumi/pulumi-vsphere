// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetHost
    {
        /// <summary>
        /// The `vsphere.Host` data source can be used to discover the ID of a vSphere
        /// host. This can then be used with resources or data sources that require a host
        /// managed object reference ID.
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
        ///         var datacenter = Output.Create(VSphere.GetDatacenter.InvokeAsync(new VSphere.GetDatacenterArgs
        ///         {
        ///             Name = "dc1",
        ///         }));
        ///         var host = datacenter.Apply(datacenter =&gt; Output.Create(VSphere.GetHost.InvokeAsync(new VSphere.GetHostArgs
        ///         {
        ///             DatacenterId = datacenter.Id,
        ///             Name = "esxi1",
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetHostResult> InvokeAsync(GetHostArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHostResult>("vsphere:index/getHost:getHost", args ?? new GetHostArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.Host` data source can be used to discover the ID of a vSphere
        /// host. This can then be used with resources or data sources that require a host
        /// managed object reference ID.
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
        ///         var datacenter = Output.Create(VSphere.GetDatacenter.InvokeAsync(new VSphere.GetDatacenterArgs
        ///         {
        ///             Name = "dc1",
        ///         }));
        ///         var host = datacenter.Apply(datacenter =&gt; Output.Create(VSphere.GetHost.InvokeAsync(new VSphere.GetHostArgs
        ///         {
        ///             DatacenterId = datacenter.Id,
        ///             Name = "esxi1",
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetHostResult> Invoke(GetHostInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetHostResult>("vsphere:index/getHost:getHost", args ?? new GetHostInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHostArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The managed object reference
        /// ID of a datacenter.
        /// </summary>
        [Input("datacenterId", required: true)]
        public string DatacenterId { get; set; } = null!;

        /// <summary>
        /// The name of the host. This can be a name or path. Can be
        /// omitted if there is only one host in your inventory.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetHostArgs()
        {
        }
    }

    public sealed class GetHostInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The managed object reference
        /// ID of a datacenter.
        /// </summary>
        [Input("datacenterId", required: true)]
        public Input<string> DatacenterId { get; set; } = null!;

        /// <summary>
        /// The name of the host. This can be a name or path. Can be
        /// omitted if there is only one host in your inventory.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetHostInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetHostResult
    {
        public readonly string DatacenterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        /// <summary>
        /// The managed object ID of the host's
        /// root resource pool.
        /// </summary>
        public readonly string ResourcePoolId;

        [OutputConstructor]
        private GetHostResult(
            string datacenterId,

            string id,

            string? name,

            string resourcePoolId)
        {
            DatacenterId = datacenterId;
            Id = id;
            Name = name;
            ResourcePoolId = resourcePoolId;
        }
    }
}

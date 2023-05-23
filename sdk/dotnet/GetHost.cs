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
        /// The `vsphere.Host` data source can be used to discover the ID of an ESXi host.
        /// This can then be used with resources or data sources that require an ESX
        /// host's managed object reference ID.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using VSphere = Pulumi.VSphere;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var datacenter = VSphere.GetDatacenter.Invoke(new()
        ///     {
        ///         Name = "dc-01",
        ///     });
        /// 
        ///     var host = VSphere.GetHost.Invoke(new()
        ///     {
        ///         Name = "esxi-01.example.com",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetHostResult> InvokeAsync(GetHostArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHostResult>("vsphere:index/getHost:getHost", args ?? new GetHostArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.Host` data source can be used to discover the ID of an ESXi host.
        /// This can then be used with resources or data sources that require an ESX
        /// host's managed object reference ID.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using VSphere = Pulumi.VSphere;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var datacenter = VSphere.GetDatacenter.Invoke(new()
        ///     {
        ///         Name = "dc-01",
        ///     });
        /// 
        ///     var host = VSphere.GetHost.Invoke(new()
        ///     {
        ///         Name = "esxi-01.example.com",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetHostResult> Invoke(GetHostInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHostResult>("vsphere:index/getHost:getHost", args ?? new GetHostInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHostArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The managed object reference ID
        /// of a vSphere datacenter object.
        /// </summary>
        [Input("datacenterId", required: true)]
        public string DatacenterId { get; set; } = null!;

        /// <summary>
        /// The name of the ESXI host. This can be a name or path.
        /// Can be omitted if there is only one host in your inventory.
        /// 
        /// &gt; **NOTE:** When used against an ESXi host directly, this data source _always_
        /// returns the ESXi host's object ID, regardless of what is entered into `name`.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetHostArgs()
        {
        }
        public static new GetHostArgs Empty => new GetHostArgs();
    }

    public sealed class GetHostInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The managed object reference ID
        /// of a vSphere datacenter object.
        /// </summary>
        [Input("datacenterId", required: true)]
        public Input<string> DatacenterId { get; set; } = null!;

        /// <summary>
        /// The name of the ESXI host. This can be a name or path.
        /// Can be omitted if there is only one host in your inventory.
        /// 
        /// &gt; **NOTE:** When used against an ESXi host directly, this data source _always_
        /// returns the ESXi host's object ID, regardless of what is entered into `name`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetHostInvokeArgs()
        {
        }
        public static new GetHostInvokeArgs Empty => new GetHostInvokeArgs();
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
        /// The managed object ID of the ESXi
        /// host's root resource pool.
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

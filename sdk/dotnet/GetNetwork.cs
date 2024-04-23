// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetNetwork
    {
        /// <summary>
        /// The `vsphere.getNetwork` data source can be used to discover the ID of a network
        /// in vSphere. This can be any network that can be used as the backing for a
        /// network interface for `vsphere.VirtualMachine` or any other vSphere resource
        /// that requires a network. This includes standard (host-based) port groups,
        /// distributed port groups, or opaque networks such as those managed by NSX.
        /// 
        /// ## Example Usage
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
        ///     var network = VSphere.GetNetwork.Invoke(new()
        ///     {
        ///         Name = "VM Network",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetNetworkResult> InvokeAsync(GetNetworkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkResult>("vsphere:index/getNetwork:getNetwork", args ?? new GetNetworkArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.getNetwork` data source can be used to discover the ID of a network
        /// in vSphere. This can be any network that can be used as the backing for a
        /// network interface for `vsphere.VirtualMachine` or any other vSphere resource
        /// that requires a network. This includes standard (host-based) port groups,
        /// distributed port groups, or opaque networks such as those managed by NSX.
        /// 
        /// ## Example Usage
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
        ///     var network = VSphere.GetNetwork.Invoke(new()
        ///     {
        ///         Name = "VM Network",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetNetworkResult> Invoke(GetNetworkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkResult>("vsphere:index/getNetwork:getNetwork", args ?? new GetNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The managed object reference ID
        /// of the datacenter the network is located in. This can be omitted if the
        /// search path used in `name` is an absolute path. For default datacenters,
        /// use the `id` attribute from an empty `vsphere.Datacenter` data source.
        /// </summary>
        [Input("datacenterId")]
        public string? DatacenterId { get; set; }

        /// <summary>
        /// For distributed port group type
        /// network objects, the ID of the distributed virtual switch for which the port
        /// group belongs. It is useful to differentiate port groups with same name
        /// using the distributed virtual switch ID.
        /// </summary>
        [Input("distributedVirtualSwitchUuid")]
        public string? DistributedVirtualSwitchUuid { get; set; }

        /// <summary>
        /// The name of the network. This can be a name or path.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetNetworkArgs()
        {
        }
        public static new GetNetworkArgs Empty => new GetNetworkArgs();
    }

    public sealed class GetNetworkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The managed object reference ID
        /// of the datacenter the network is located in. This can be omitted if the
        /// search path used in `name` is an absolute path. For default datacenters,
        /// use the `id` attribute from an empty `vsphere.Datacenter` data source.
        /// </summary>
        [Input("datacenterId")]
        public Input<string>? DatacenterId { get; set; }

        /// <summary>
        /// For distributed port group type
        /// network objects, the ID of the distributed virtual switch for which the port
        /// group belongs. It is useful to differentiate port groups with same name
        /// using the distributed virtual switch ID.
        /// </summary>
        [Input("distributedVirtualSwitchUuid")]
        public Input<string>? DistributedVirtualSwitchUuid { get; set; }

        /// <summary>
        /// The name of the network. This can be a name or path.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetNetworkInvokeArgs()
        {
        }
        public static new GetNetworkInvokeArgs Empty => new GetNetworkInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkResult
    {
        public readonly string? DatacenterId;
        public readonly string? DistributedVirtualSwitchUuid;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The managed object type for the discovered network. This will be one
        /// of `DistributedVirtualPortgroup` for distributed port groups, `Network` for
        /// standard (host-based) port groups, or `OpaqueNetwork` for networks managed
        /// externally, such as those managed by NSX.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetNetworkResult(
            string? datacenterId,

            string? distributedVirtualSwitchUuid,

            string id,

            string name,

            string type)
        {
            DatacenterId = datacenterId;
            DistributedVirtualSwitchUuid = distributedVirtualSwitchUuid;
            Id = id;
            Name = name;
            Type = type;
        }
    }
}

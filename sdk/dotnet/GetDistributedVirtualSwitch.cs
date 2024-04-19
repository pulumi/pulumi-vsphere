// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetDistributedVirtualSwitch
    {
        /// <summary>
        /// The `vsphere.DistributedVirtualSwitch` data source can be used to discover
        /// the ID and uplink data of a of a vSphere distributed switch (VDS). This
        /// can then be used with resources or data sources that require a VDS, such as the
        /// `vsphere.DistributedPortGroup` resource, for which
        /// an example is shown below.
        /// 
        /// &gt; **NOTE:** This data source requires vCenter Server and is not available on
        /// direct ESXi host connections.
        /// 
        /// ## Example Usage
        /// 
        /// The following example locates a distributed switch named `vds-01`, in the
        /// datacenter `dc-01`. It then uses this distributed switch to set up a
        /// `vsphere.DistributedPortGroup` resource that uses the first uplink as a
        /// primary uplink and the second uplink as a secondary.
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
        ///     var vds = VSphere.GetDistributedVirtualSwitch.Invoke(new()
        ///     {
        ///         Name = "vds-01",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        ///     var dvpg = new VSphere.DistributedPortGroup("dvpg", new()
        ///     {
        ///         DistributedVirtualSwitchUuid = vds.Apply(getDistributedVirtualSwitchResult =&gt; getDistributedVirtualSwitchResult.Id),
        ///         ActiveUplinks = new[]
        ///         {
        ///             vds.Apply(getDistributedVirtualSwitchResult =&gt; getDistributedVirtualSwitchResult.Uplinks[0]),
        ///         },
        ///         StandbyUplinks = new[]
        ///         {
        ///             vds.Apply(getDistributedVirtualSwitchResult =&gt; getDistributedVirtualSwitchResult.Uplinks[1]),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDistributedVirtualSwitchResult> InvokeAsync(GetDistributedVirtualSwitchArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDistributedVirtualSwitchResult>("vsphere:index/getDistributedVirtualSwitch:getDistributedVirtualSwitch", args ?? new GetDistributedVirtualSwitchArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.DistributedVirtualSwitch` data source can be used to discover
        /// the ID and uplink data of a of a vSphere distributed switch (VDS). This
        /// can then be used with resources or data sources that require a VDS, such as the
        /// `vsphere.DistributedPortGroup` resource, for which
        /// an example is shown below.
        /// 
        /// &gt; **NOTE:** This data source requires vCenter Server and is not available on
        /// direct ESXi host connections.
        /// 
        /// ## Example Usage
        /// 
        /// The following example locates a distributed switch named `vds-01`, in the
        /// datacenter `dc-01`. It then uses this distributed switch to set up a
        /// `vsphere.DistributedPortGroup` resource that uses the first uplink as a
        /// primary uplink and the second uplink as a secondary.
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
        ///     var vds = VSphere.GetDistributedVirtualSwitch.Invoke(new()
        ///     {
        ///         Name = "vds-01",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        ///     var dvpg = new VSphere.DistributedPortGroup("dvpg", new()
        ///     {
        ///         DistributedVirtualSwitchUuid = vds.Apply(getDistributedVirtualSwitchResult =&gt; getDistributedVirtualSwitchResult.Id),
        ///         ActiveUplinks = new[]
        ///         {
        ///             vds.Apply(getDistributedVirtualSwitchResult =&gt; getDistributedVirtualSwitchResult.Uplinks[0]),
        ///         },
        ///         StandbyUplinks = new[]
        ///         {
        ///             vds.Apply(getDistributedVirtualSwitchResult =&gt; getDistributedVirtualSwitchResult.Uplinks[1]),
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDistributedVirtualSwitchResult> Invoke(GetDistributedVirtualSwitchInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDistributedVirtualSwitchResult>("vsphere:index/getDistributedVirtualSwitch:getDistributedVirtualSwitch", args ?? new GetDistributedVirtualSwitchInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDistributedVirtualSwitchArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The managed object reference ID
        /// of the datacenter the VDS is located in. This can be omitted if the search
        /// path used in `name` is an absolute path. For default datacenters, use the `id`
        /// attribute from an empty `vsphere.Datacenter` data source.
        /// </summary>
        [Input("datacenterId")]
        public string? DatacenterId { get; set; }

        /// <summary>
        /// The name of the VDS. This can be a name or path.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetDistributedVirtualSwitchArgs()
        {
        }
        public static new GetDistributedVirtualSwitchArgs Empty => new GetDistributedVirtualSwitchArgs();
    }

    public sealed class GetDistributedVirtualSwitchInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The managed object reference ID
        /// of the datacenter the VDS is located in. This can be omitted if the search
        /// path used in `name` is an absolute path. For default datacenters, use the `id`
        /// attribute from an empty `vsphere.Datacenter` data source.
        /// </summary>
        [Input("datacenterId")]
        public Input<string>? DatacenterId { get; set; }

        /// <summary>
        /// The name of the VDS. This can be a name or path.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetDistributedVirtualSwitchInvokeArgs()
        {
        }
        public static new GetDistributedVirtualSwitchInvokeArgs Empty => new GetDistributedVirtualSwitchInvokeArgs();
    }


    [OutputType]
    public sealed class GetDistributedVirtualSwitchResult
    {
        public readonly string? DatacenterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The list of the uplinks on this vSphere distributed switch, as per the
        /// `uplinks` argument to the
        /// `vsphere.DistributedVirtualSwitch`
        /// resource.
        /// </summary>
        public readonly ImmutableArray<string> Uplinks;

        [OutputConstructor]
        private GetDistributedVirtualSwitchResult(
            string? datacenterId,

            string id,

            string name,

            ImmutableArray<string> uplinks)
        {
            DatacenterId = datacenterId;
            Id = id;
            Name = name;
            Uplinks = uplinks;
        }
    }
}

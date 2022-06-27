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
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// The following example locates a distributed switch named `vds-01`, in the
        /// datacenter `dc-01`. It then uses this distributed switch to set up a
        /// `vsphere.DistributedPortGroup` resource that uses the first uplink as a
        /// primary uplink and the second uplink as a secondary.
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
        ///             Name = "dc-01",
        ///         }));
        ///         var vds = datacenter.Apply(datacenter =&gt; Output.Create(VSphere.GetDistributedVirtualSwitch.InvokeAsync(new VSphere.GetDistributedVirtualSwitchArgs
        ///         {
        ///             Name = "vds-01",
        ///             DatacenterId = datacenter.Id,
        ///         })));
        ///         var dvpg = new VSphere.DistributedPortGroup("dvpg", new VSphere.DistributedPortGroupArgs
        ///         {
        ///             DistributedVirtualSwitchUuid = vds.Apply(vds =&gt; vds.Id),
        ///             ActiveUplinks = 
        ///             {
        ///                 vds.Apply(vds =&gt; vds.Uplinks?[0]),
        ///             },
        ///             StandbyUplinks = 
        ///             {
        ///                 vds.Apply(vds =&gt; vds.Uplinks?[1]),
        ///             },
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDistributedVirtualSwitchResult> InvokeAsync(GetDistributedVirtualSwitchArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDistributedVirtualSwitchResult>("vsphere:index/getDistributedVirtualSwitch:getDistributedVirtualSwitch", args ?? new GetDistributedVirtualSwitchArgs(), options.WithDefaults());

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
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// The following example locates a distributed switch named `vds-01`, in the
        /// datacenter `dc-01`. It then uses this distributed switch to set up a
        /// `vsphere.DistributedPortGroup` resource that uses the first uplink as a
        /// primary uplink and the second uplink as a secondary.
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
        ///             Name = "dc-01",
        ///         }));
        ///         var vds = datacenter.Apply(datacenter =&gt; Output.Create(VSphere.GetDistributedVirtualSwitch.InvokeAsync(new VSphere.GetDistributedVirtualSwitchArgs
        ///         {
        ///             Name = "vds-01",
        ///             DatacenterId = datacenter.Id,
        ///         })));
        ///         var dvpg = new VSphere.DistributedPortGroup("dvpg", new VSphere.DistributedPortGroupArgs
        ///         {
        ///             DistributedVirtualSwitchUuid = vds.Apply(vds =&gt; vds.Id),
        ///             ActiveUplinks = 
        ///             {
        ///                 vds.Apply(vds =&gt; vds.Uplinks?[0]),
        ///             },
        ///             StandbyUplinks = 
        ///             {
        ///                 vds.Apply(vds =&gt; vds.Uplinks?[1]),
        ///             },
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDistributedVirtualSwitchResult> Invoke(GetDistributedVirtualSwitchInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDistributedVirtualSwitchResult>("vsphere:index/getDistributedVirtualSwitch:getDistributedVirtualSwitch", args ?? new GetDistributedVirtualSwitchInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDistributedVirtualSwitchArgs : Pulumi.InvokeArgs
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
    }

    public sealed class GetDistributedVirtualSwitchInvokeArgs : Pulumi.InvokeArgs
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

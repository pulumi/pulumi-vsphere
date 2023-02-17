// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetComputeClusterHostGroup
    {
        /// <summary>
        /// The `vsphere.ComputeClusterHostGroup` data source can be used to discover
        /// the IDs ESXi hosts in a host group and return host group attributes to other
        /// resources.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using VSphere = Pulumi.VSphere;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var datacenter = VSphere.GetDatacenter.Invoke(new()
        ///     {
        ///         Name = @var.Vsphere_datacenter,
        ///     });
        /// 
        ///     var cluster = VSphere.GetComputeCluster.Invoke(new()
        ///     {
        ///         Name = @var.Vsphere_cluster,
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        ///     var hostGroup1 = VSphere.GetComputeClusterHostGroup.Invoke(new()
        ///     {
        ///         Name = "host_group1",
        ///         ComputeClusterId = cluster.Apply(getComputeClusterResult =&gt; getComputeClusterResult.Id),
        ///     });
        /// 
        ///     var hostRule1 = new VSphere.ComputeClusterVmHostRule("hostRule1", new()
        ///     {
        ///         ComputeClusterId = cluster.Apply(getComputeClusterResult =&gt; getComputeClusterResult.Id),
        ///         VmGroupName = "vm_group1",
        ///         AffinityHostGroupName = hostGroup1.Apply(getComputeClusterHostGroupResult =&gt; getComputeClusterHostGroupResult.Name),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetComputeClusterHostGroupResult> InvokeAsync(GetComputeClusterHostGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetComputeClusterHostGroupResult>("vsphere:index/getComputeClusterHostGroup:getComputeClusterHostGroup", args ?? new GetComputeClusterHostGroupArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.ComputeClusterHostGroup` data source can be used to discover
        /// the IDs ESXi hosts in a host group and return host group attributes to other
        /// resources.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using VSphere = Pulumi.VSphere;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var datacenter = VSphere.GetDatacenter.Invoke(new()
        ///     {
        ///         Name = @var.Vsphere_datacenter,
        ///     });
        /// 
        ///     var cluster = VSphere.GetComputeCluster.Invoke(new()
        ///     {
        ///         Name = @var.Vsphere_cluster,
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        ///     var hostGroup1 = VSphere.GetComputeClusterHostGroup.Invoke(new()
        ///     {
        ///         Name = "host_group1",
        ///         ComputeClusterId = cluster.Apply(getComputeClusterResult =&gt; getComputeClusterResult.Id),
        ///     });
        /// 
        ///     var hostRule1 = new VSphere.ComputeClusterVmHostRule("hostRule1", new()
        ///     {
        ///         ComputeClusterId = cluster.Apply(getComputeClusterResult =&gt; getComputeClusterResult.Id),
        ///         VmGroupName = "vm_group1",
        ///         AffinityHostGroupName = hostGroup1.Apply(getComputeClusterHostGroupResult =&gt; getComputeClusterHostGroupResult.Name),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetComputeClusterHostGroupResult> Invoke(GetComputeClusterHostGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetComputeClusterHostGroupResult>("vsphere:index/getComputeClusterHostGroup:getComputeClusterHostGroup", args ?? new GetComputeClusterHostGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetComputeClusterHostGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The [managed object reference ID][docs-about-morefs]
        /// of the compute cluster for the host group.
        /// </summary>
        [Input("computeClusterId", required: true)]
        public string ComputeClusterId { get; set; } = null!;

        /// <summary>
        /// The name of the host group.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetComputeClusterHostGroupArgs()
        {
        }
        public static new GetComputeClusterHostGroupArgs Empty => new GetComputeClusterHostGroupArgs();
    }

    public sealed class GetComputeClusterHostGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The [managed object reference ID][docs-about-morefs]
        /// of the compute cluster for the host group.
        /// </summary>
        [Input("computeClusterId", required: true)]
        public Input<string> ComputeClusterId { get; set; } = null!;

        /// <summary>
        /// The name of the host group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetComputeClusterHostGroupInvokeArgs()
        {
        }
        public static new GetComputeClusterHostGroupInvokeArgs Empty => new GetComputeClusterHostGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetComputeClusterHostGroupResult
    {
        public readonly string ComputeClusterId;
        /// <summary>
        /// The [managed object reference ID][docs-about-morefs] of
        /// the ESXi hosts in the host group.
        /// </summary>
        public readonly ImmutableArray<string> HostSystemIds;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetComputeClusterHostGroupResult(
            string computeClusterId,

            ImmutableArray<string> hostSystemIds,

            string id,

            string name)
        {
            ComputeClusterId = computeClusterId;
            HostSystemIds = hostSystemIds;
            Id = id;
            Name = name;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetComputeCluster
    {
        /// <summary>
        /// The `vsphere.ComputeCluster` data source can be used to discover the ID of a
        /// cluster in vSphere. This is useful to fetch the ID of a cluster that you want
        /// to use for virtual machine placement via the
        /// `vsphere.VirtualMachine` resource, allowing
        /// you to specify the cluster's root resource pool directly versus using the alias
        /// available through the `vsphere.ResourcePool`
        /// data source.
        /// 
        /// &gt; You may also wish to see the
        /// `vsphere.ComputeCluster` resource for further
        /// details about clusters or how to work with them.
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
        ///         var computeCluster = Output.Create(VSphere.GetComputeCluster.InvokeAsync(new VSphere.GetComputeClusterArgs
        ///         {
        ///             DatacenterId = data.Vsphere_datacenter.Dc.Id,
        ///             Name = "compute-cluster1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetComputeClusterResult> InvokeAsync(GetComputeClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetComputeClusterResult>("vsphere:index/getComputeCluster:getComputeCluster", args ?? new GetComputeClusterArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.ComputeCluster` data source can be used to discover the ID of a
        /// cluster in vSphere. This is useful to fetch the ID of a cluster that you want
        /// to use for virtual machine placement via the
        /// `vsphere.VirtualMachine` resource, allowing
        /// you to specify the cluster's root resource pool directly versus using the alias
        /// available through the `vsphere.ResourcePool`
        /// data source.
        /// 
        /// &gt; You may also wish to see the
        /// `vsphere.ComputeCluster` resource for further
        /// details about clusters or how to work with them.
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
        ///         var computeCluster = Output.Create(VSphere.GetComputeCluster.InvokeAsync(new VSphere.GetComputeClusterArgs
        ///         {
        ///             DatacenterId = data.Vsphere_datacenter.Dc.Id,
        ///             Name = "compute-cluster1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetComputeClusterResult> Invoke(GetComputeClusterInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetComputeClusterResult>("vsphere:index/getComputeCluster:getComputeCluster", args ?? new GetComputeClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetComputeClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The managed object reference
        /// ID of the datacenter the cluster is located in.  This can
        /// be omitted if the search path used in `name` is an absolute path.  For
        /// default datacenters, use the id attribute from an empty `vsphere.Datacenter`
        /// data source.
        /// </summary>
        [Input("datacenterId")]
        public string? DatacenterId { get; set; }

        /// <summary>
        /// The name or absolute path to the cluster.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetComputeClusterArgs()
        {
        }
    }

    public sealed class GetComputeClusterInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The managed object reference
        /// ID of the datacenter the cluster is located in.  This can
        /// be omitted if the search path used in `name` is an absolute path.  For
        /// default datacenters, use the id attribute from an empty `vsphere.Datacenter`
        /// data source.
        /// </summary>
        [Input("datacenterId")]
        public Input<string>? DatacenterId { get; set; }

        /// <summary>
        /// The name or absolute path to the cluster.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetComputeClusterInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetComputeClusterResult
    {
        public readonly string? DatacenterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ResourcePoolId;

        [OutputConstructor]
        private GetComputeClusterResult(
            string? datacenterId,

            string id,

            string name,

            string resourcePoolId)
        {
            DatacenterId = datacenterId;
            Id = id;
            Name = name;
            ResourcePoolId = resourcePoolId;
        }
    }
}

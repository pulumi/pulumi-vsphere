// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.VSphere
{
    public static class GetResourcePool
    {
        /// <summary>
        /// The `vsphere.ResourcePool` data source can be used to discover the ID of a
        /// resource pool in vSphere. This is useful to fetch the ID of a resource pool
        /// that you want to use to create virtual machines in using the
        /// `vsphere.VirtualMachine` resource. 
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
        ///         var pool = datacenter.Apply(datacenter =&gt; Output.Create(VSphere.GetResourcePool.InvokeAsync(new VSphere.GetResourcePoolArgs
        ///         {
        ///             DatacenterId = datacenter.Id,
        ///             Name = "resource-pool-1",
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% example %}}
        /// ### Specifying the root resource pool for a standalone host
        /// 
        /// &gt; **NOTE:** Fetching the root resource pool for a cluster can now be done
        /// directly via the `vsphere.ComputeCluster`
        /// data source.
        /// 
        /// All compute resources in vSphere (clusters, standalone hosts, and standalone
        /// ESXi) have a resource pool, even if one has not been explicitly created. This
        /// resource pool is referred to as the _root resource pool_ and can be looked up
        /// by specifying the path as per the example below:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using VSphere = Pulumi.VSphere;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var pool = Output.Create(VSphere.GetResourcePool.InvokeAsync(new VSphere.GetResourcePoolArgs
        ///         {
        ///             DatacenterId = data.Vsphere_datacenter.Dc.Id,
        ///             Name = "esxi1/Resources",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// For more information on the root resource pool, see [Managing Resource
        /// Pools][vmware-docs-resource-pools] in the vSphere documentation.
        /// 
        /// [vmware-docs-resource-pools]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-60077B40-66FF-4625-934A-641703ED7601.html
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetResourcePoolResult> InvokeAsync(GetResourcePoolArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourcePoolResult>("vsphere:index/getResourcePool:getResourcePool", args ?? new GetResourcePoolArgs(), options.WithVersion());

        /// <summary>
        /// The `vsphere.ResourcePool` data source can be used to discover the ID of a
        /// resource pool in vSphere. This is useful to fetch the ID of a resource pool
        /// that you want to use to create virtual machines in using the
        /// `vsphere.VirtualMachine` resource. 
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
        ///         var pool = datacenter.Apply(datacenter =&gt; Output.Create(VSphere.GetResourcePool.InvokeAsync(new VSphere.GetResourcePoolArgs
        ///         {
        ///             DatacenterId = datacenter.Id,
        ///             Name = "resource-pool-1",
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% example %}}
        /// ### Specifying the root resource pool for a standalone host
        /// 
        /// &gt; **NOTE:** Fetching the root resource pool for a cluster can now be done
        /// directly via the `vsphere.ComputeCluster`
        /// data source.
        /// 
        /// All compute resources in vSphere (clusters, standalone hosts, and standalone
        /// ESXi) have a resource pool, even if one has not been explicitly created. This
        /// resource pool is referred to as the _root resource pool_ and can be looked up
        /// by specifying the path as per the example below:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using VSphere = Pulumi.VSphere;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var pool = Output.Create(VSphere.GetResourcePool.InvokeAsync(new VSphere.GetResourcePoolArgs
        ///         {
        ///             DatacenterId = data.Vsphere_datacenter.Dc.Id,
        ///             Name = "esxi1/Resources",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// For more information on the root resource pool, see [Managing Resource
        /// Pools][vmware-docs-resource-pools] in the vSphere documentation.
        /// 
        /// [vmware-docs-resource-pools]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-60077B40-66FF-4625-934A-641703ED7601.html
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetResourcePoolResult> Invoke(GetResourcePoolInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetResourcePoolResult>("vsphere:index/getResourcePool:getResourcePool", args ?? new GetResourcePoolInvokeArgs(), options.WithVersion());
    }


    public sealed class GetResourcePoolArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The managed object reference
        /// ID of the datacenter the resource pool is located in.
        /// This can be omitted if the search path used in `name` is an absolute path.
        /// For default datacenters, use the id attribute from an empty
        /// `vsphere.Datacenter` data source.
        /// </summary>
        [Input("datacenterId")]
        public string? DatacenterId { get; set; }

        /// <summary>
        /// The name of the resource pool. This can be a name or
        /// path. This is required when using vCenter.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetResourcePoolArgs()
        {
        }
    }

    public sealed class GetResourcePoolInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The managed object reference
        /// ID of the datacenter the resource pool is located in.
        /// This can be omitted if the search path used in `name` is an absolute path.
        /// For default datacenters, use the id attribute from an empty
        /// `vsphere.Datacenter` data source.
        /// </summary>
        [Input("datacenterId")]
        public Input<string>? DatacenterId { get; set; }

        /// <summary>
        /// The name of the resource pool. This can be a name or
        /// path. This is required when using vCenter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetResourcePoolInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResourcePoolResult
    {
        public readonly string? DatacenterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;

        [OutputConstructor]
        private GetResourcePoolResult(
            string? datacenterId,

            string id,

            string? name)
        {
            DatacenterId = datacenterId;
            Id = id;
            Name = name;
        }
    }
}

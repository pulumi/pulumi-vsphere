// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetDatastoreCluster
    {
        /// <summary>
        /// The `vsphere.DatastoreCluster` data source can be used to discover the ID of a
        /// vSphere datastore cluster object. This can then be used with resources or data sources
        /// that require a datastore. For example, to assign datastores using the
        /// `vsphere.NasDatastore` or `vsphere.VmfsDatastore` resources, or to create virtual machines in using the `vsphere.VirtualMachine` resource.
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
        ///     var datastoreCluster = VSphere.GetDatastoreCluster.Invoke(new()
        ///     {
        ///         Name = "datastore-cluster-01",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatastoreClusterResult> InvokeAsync(GetDatastoreClusterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatastoreClusterResult>("vsphere:index/getDatastoreCluster:getDatastoreCluster", args ?? new GetDatastoreClusterArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.DatastoreCluster` data source can be used to discover the ID of a
        /// vSphere datastore cluster object. This can then be used with resources or data sources
        /// that require a datastore. For example, to assign datastores using the
        /// `vsphere.NasDatastore` or `vsphere.VmfsDatastore` resources, or to create virtual machines in using the `vsphere.VirtualMachine` resource.
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
        ///     var datastoreCluster = VSphere.GetDatastoreCluster.Invoke(new()
        ///     {
        ///         Name = "datastore-cluster-01",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDatastoreClusterResult> Invoke(GetDatastoreClusterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatastoreClusterResult>("vsphere:index/getDatastoreCluster:getDatastoreCluster", args ?? new GetDatastoreClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatastoreClusterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The managed object reference
        /// ID of the datacenter the datastore cluster is located in.
        /// This can be omitted if the search path used in `name` is an absolute path.
        /// For default datacenters, use the id attribute from an empty
        /// `vsphere.Datacenter` data source.
        /// </summary>
        [Input("datacenterId")]
        public string? DatacenterId { get; set; }

        /// <summary>
        /// The name or absolute path to the datastore cluster.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetDatastoreClusterArgs()
        {
        }
        public static new GetDatastoreClusterArgs Empty => new GetDatastoreClusterArgs();
    }

    public sealed class GetDatastoreClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The managed object reference
        /// ID of the datacenter the datastore cluster is located in.
        /// This can be omitted if the search path used in `name` is an absolute path.
        /// For default datacenters, use the id attribute from an empty
        /// `vsphere.Datacenter` data source.
        /// </summary>
        [Input("datacenterId")]
        public Input<string>? DatacenterId { get; set; }

        /// <summary>
        /// The name or absolute path to the datastore cluster.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetDatastoreClusterInvokeArgs()
        {
        }
        public static new GetDatastoreClusterInvokeArgs Empty => new GetDatastoreClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatastoreClusterResult
    {
        public readonly string? DatacenterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetDatastoreClusterResult(
            string? datacenterId,

            string id,

            string name)
        {
            DatacenterId = datacenterId;
            Id = id;
            Name = name;
        }
    }
}

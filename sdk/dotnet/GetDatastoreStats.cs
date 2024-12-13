// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetDatastoreStats
    {
        /// <summary>
        /// The `vsphere.getDatastoreStats` data source can be used to retrieve the usage
        /// stats of all vSphere datastore objects in a datacenter. This can then be used as
        /// a standalone data source to get information required as input to other data
        /// sources.
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
        ///     var datastoreStats = VSphere.GetDatastoreStats.Invoke(new()
        ///     {
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// A useful example of this data source would be to determine the datastore with
        /// the most free space. For example, in addition to the above:
        /// 
        /// Create an `outputs.tf` like that:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["maxFreeSpaceName"] = theirMaxFreeSpaceName,
        ///         ["maxFreeSpace"] = theirMaxFreeSpace,
        ///     };
        /// });
        /// ```
        /// 
        /// and a `locals.tf` like that:
        /// </summary>
        public static Task<GetDatastoreStatsResult> InvokeAsync(GetDatastoreStatsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatastoreStatsResult>("vsphere:index/getDatastoreStats:getDatastoreStats", args ?? new GetDatastoreStatsArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.getDatastoreStats` data source can be used to retrieve the usage
        /// stats of all vSphere datastore objects in a datacenter. This can then be used as
        /// a standalone data source to get information required as input to other data
        /// sources.
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
        ///     var datastoreStats = VSphere.GetDatastoreStats.Invoke(new()
        ///     {
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// A useful example of this data source would be to determine the datastore with
        /// the most free space. For example, in addition to the above:
        /// 
        /// Create an `outputs.tf` like that:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["maxFreeSpaceName"] = theirMaxFreeSpaceName,
        ///         ["maxFreeSpace"] = theirMaxFreeSpace,
        ///     };
        /// });
        /// ```
        /// 
        /// and a `locals.tf` like that:
        /// </summary>
        public static Output<GetDatastoreStatsResult> Invoke(GetDatastoreStatsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatastoreStatsResult>("vsphere:index/getDatastoreStats:getDatastoreStats", args ?? new GetDatastoreStatsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.getDatastoreStats` data source can be used to retrieve the usage
        /// stats of all vSphere datastore objects in a datacenter. This can then be used as
        /// a standalone data source to get information required as input to other data
        /// sources.
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
        ///     var datastoreStats = VSphere.GetDatastoreStats.Invoke(new()
        ///     {
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// A useful example of this data source would be to determine the datastore with
        /// the most free space. For example, in addition to the above:
        /// 
        /// Create an `outputs.tf` like that:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["maxFreeSpaceName"] = theirMaxFreeSpaceName,
        ///         ["maxFreeSpace"] = theirMaxFreeSpace,
        ///     };
        /// });
        /// ```
        /// 
        /// and a `locals.tf` like that:
        /// </summary>
        public static Output<GetDatastoreStatsResult> Invoke(GetDatastoreStatsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatastoreStatsResult>("vsphere:index/getDatastoreStats:getDatastoreStats", args ?? new GetDatastoreStatsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatastoreStatsArgs : global::Pulumi.InvokeArgs
    {
        [Input("capacity")]
        private Dictionary<string, string>? _capacity;

        /// <summary>
        /// A mapping of the capacity for all datastore in the datacenter,
        /// where the name of the datastore is used as key and the capacity as value.
        /// </summary>
        public Dictionary<string, string> Capacity
        {
            get => _capacity ?? (_capacity = new Dictionary<string, string>());
            set => _capacity = value;
        }

        /// <summary>
        /// The
        /// [managed object reference ID][docs-about-morefs] of the datacenter the
        /// datastores are located in. For default datacenters, use the `id` attribute
        /// from an empty `vsphere.Datacenter` data source.
        /// 
        /// [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
        /// </summary>
        [Input("datacenterId", required: true)]
        public string DatacenterId { get; set; } = null!;

        [Input("freeSpace")]
        private Dictionary<string, string>? _freeSpace;

        /// <summary>
        /// A mapping of the free space for each datastore in the
        /// datacenter, where the name of the datastore is used as key and the free space
        /// as value.
        /// </summary>
        public Dictionary<string, string> FreeSpace
        {
            get => _freeSpace ?? (_freeSpace = new Dictionary<string, string>());
            set => _freeSpace = value;
        }

        public GetDatastoreStatsArgs()
        {
        }
        public static new GetDatastoreStatsArgs Empty => new GetDatastoreStatsArgs();
    }

    public sealed class GetDatastoreStatsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("capacity")]
        private InputMap<string>? _capacity;

        /// <summary>
        /// A mapping of the capacity for all datastore in the datacenter,
        /// where the name of the datastore is used as key and the capacity as value.
        /// </summary>
        public InputMap<string> Capacity
        {
            get => _capacity ?? (_capacity = new InputMap<string>());
            set => _capacity = value;
        }

        /// <summary>
        /// The
        /// [managed object reference ID][docs-about-morefs] of the datacenter the
        /// datastores are located in. For default datacenters, use the `id` attribute
        /// from an empty `vsphere.Datacenter` data source.
        /// 
        /// [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
        /// </summary>
        [Input("datacenterId", required: true)]
        public Input<string> DatacenterId { get; set; } = null!;

        [Input("freeSpace")]
        private InputMap<string>? _freeSpace;

        /// <summary>
        /// A mapping of the free space for each datastore in the
        /// datacenter, where the name of the datastore is used as key and the free space
        /// as value.
        /// </summary>
        public InputMap<string> FreeSpace
        {
            get => _freeSpace ?? (_freeSpace = new InputMap<string>());
            set => _freeSpace = value;
        }

        public GetDatastoreStatsInvokeArgs()
        {
        }
        public static new GetDatastoreStatsInvokeArgs Empty => new GetDatastoreStatsInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatastoreStatsResult
    {
        /// <summary>
        /// A mapping of the capacity for all datastore in the datacenter,
        /// where the name of the datastore is used as key and the capacity as value.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Capacity;
        /// <summary>
        /// The [managed object reference ID][docs-about-morefs] of the
        /// datacenter the datastores are located in.
        /// </summary>
        public readonly string DatacenterId;
        /// <summary>
        /// A mapping of the free space for each datastore in the
        /// datacenter, where the name of the datastore is used as key and the free space
        /// as value.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? FreeSpace;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDatastoreStatsResult(
            ImmutableDictionary<string, string>? capacity,

            string datacenterId,

            ImmutableDictionary<string, string>? freeSpace,

            string id)
        {
            Capacity = capacity;
            DatacenterId = datacenterId;
            FreeSpace = freeSpace;
            Id = id;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetDatastore
    {
        /// <summary>
        /// The `vsphere.getDatastore` data source can be used to discover the ID of a
        /// vSphere datastore object. This can then be used with resources or data sources
        /// that require a datastore. For example, to create virtual machines in using the
        /// `vsphere.VirtualMachine` resource.
        /// 
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
        ///     var datastore = VSphere.GetDatastore.Invoke(new()
        ///     {
        ///         Name = "datastore-01",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDatastoreResult> InvokeAsync(GetDatastoreArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatastoreResult>("vsphere:index/getDatastore:getDatastore", args ?? new GetDatastoreArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.getDatastore` data source can be used to discover the ID of a
        /// vSphere datastore object. This can then be used with resources or data sources
        /// that require a datastore. For example, to create virtual machines in using the
        /// `vsphere.VirtualMachine` resource.
        /// 
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
        ///     var datastore = VSphere.GetDatastore.Invoke(new()
        ///     {
        ///         Name = "datastore-01",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatastoreResult> Invoke(GetDatastoreInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatastoreResult>("vsphere:index/getDatastore:getDatastore", args ?? new GetDatastoreInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.getDatastore` data source can be used to discover the ID of a
        /// vSphere datastore object. This can then be used with resources or data sources
        /// that require a datastore. For example, to create virtual machines in using the
        /// `vsphere.VirtualMachine` resource.
        /// 
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
        ///     var datastore = VSphere.GetDatastore.Invoke(new()
        ///     {
        ///         Name = "datastore-01",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatastoreResult> Invoke(GetDatastoreInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatastoreResult>("vsphere:index/getDatastore:getDatastore", args ?? new GetDatastoreInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatastoreArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The managed object reference ID
        /// of the datacenter the datastore is located in. This can be omitted if the
        /// search path used in `name` is an absolute path. For default datacenters, use
        /// the `id` attribute from an empty `vsphere.Datacenter` data source.
        /// </summary>
        [Input("datacenterId")]
        public string? DatacenterId { get; set; }

        /// <summary>
        /// The name of the datastore. This can be a name or path.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("stats")]
        private Dictionary<string, string>? _stats;

        /// <summary>
        /// The disk space usage statistics for the specific datastore. The
        /// total datastore capacity is represented as `capacity` and the free remaining
        /// disk is represented as `free`.
        /// </summary>
        public Dictionary<string, string> Stats
        {
            get => _stats ?? (_stats = new Dictionary<string, string>());
            set => _stats = value;
        }

        public GetDatastoreArgs()
        {
        }
        public static new GetDatastoreArgs Empty => new GetDatastoreArgs();
    }

    public sealed class GetDatastoreInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The managed object reference ID
        /// of the datacenter the datastore is located in. This can be omitted if the
        /// search path used in `name` is an absolute path. For default datacenters, use
        /// the `id` attribute from an empty `vsphere.Datacenter` data source.
        /// </summary>
        [Input("datacenterId")]
        public Input<string>? DatacenterId { get; set; }

        /// <summary>
        /// The name of the datastore. This can be a name or path.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("stats")]
        private InputMap<string>? _stats;

        /// <summary>
        /// The disk space usage statistics for the specific datastore. The
        /// total datastore capacity is represented as `capacity` and the free remaining
        /// disk is represented as `free`.
        /// </summary>
        public InputMap<string> Stats
        {
            get => _stats ?? (_stats = new InputMap<string>());
            set => _stats = value;
        }

        public GetDatastoreInvokeArgs()
        {
        }
        public static new GetDatastoreInvokeArgs Empty => new GetDatastoreInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatastoreResult
    {
        public readonly string? DatacenterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The disk space usage statistics for the specific datastore. The
        /// total datastore capacity is represented as `capacity` and the free remaining
        /// disk is represented as `free`.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Stats;

        [OutputConstructor]
        private GetDatastoreResult(
            string? datacenterId,

            string id,

            string name,

            ImmutableDictionary<string, string>? stats)
        {
            DatacenterId = datacenterId;
            Id = id;
            Name = name;
            Stats = stats;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    /// <summary>
    /// The `vsphere..DatastoreClusterVmAntiAffinityRule` resource can be used to
    /// manage VM anti-affinity rules in a datastore cluster, either created by the
    /// `vsphere..DatastoreCluster` resource or looked up
    /// by the `vsphere..DatastoreCluster` data source.
    /// 
    /// This rule can be used to tell a set to virtual machines to run on different
    /// datastores within a cluster, useful for preventing single points of failure in
    /// application cluster scenarios. When configured, Storage DRS will make a best effort to
    /// ensure that the virtual machines run on different datastores, or prevent any
    /// operation that would keep that from happening, depending on the value of the
    /// `mandatory` flag.
    /// 
    /// &gt; **NOTE:** This resource requires vCenter and is not available on direct ESXi
    /// connections.
    /// 
    /// &gt; **NOTE:** Storage DRS requires a vSphere Enterprise Plus license.
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using VSphere = Pulumi.VSphere;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var dc = Output.Create(VSphere.GetDatacenter.InvokeAsync(new VSphere.GetDatacenterArgs
    ///         {
    ///             Name = "dc1",
    ///         }));
    ///         var datastoreCluster = dc.Apply(dc =&gt; Output.Create(VSphere.GetDatastoreCluster.InvokeAsync(new VSphere.GetDatastoreClusterArgs
    ///         {
    ///             DatacenterId = dc.Id,
    ///             Name = "datastore-cluster1",
    ///         })));
    ///         var cluster = dc.Apply(dc =&gt; Output.Create(VSphere.GetComputeCluster.InvokeAsync(new VSphere.GetComputeClusterArgs
    ///         {
    ///             DatacenterId = dc.Id,
    ///             Name = "cluster1",
    ///         })));
    ///         var network = dc.Apply(dc =&gt; Output.Create(VSphere.GetNetwork.InvokeAsync(new VSphere.GetNetworkArgs
    ///         {
    ///             DatacenterId = dc.Id,
    ///             Name = "network1",
    ///         })));
    ///         var vm = new List&lt;VSphere.VirtualMachine&gt;();
    ///         for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
    ///         {
    ///             var range = new { Value = rangeIndex };
    ///             vm.Add(new VSphere.VirtualMachine($"vm-{range.Value}", new VSphere.VirtualMachineArgs
    ///             {
    ///                 DatastoreClusterId = datastoreCluster.Apply(datastoreCluster =&gt; datastoreCluster.Id),
    ///                 Disks = 
    ///                 {
    ///                     new VSphere.Inputs.VirtualMachineDiskArgs
    ///                     {
    ///                         Label = "disk0",
    ///                         Size = 20,
    ///                     },
    ///                 },
    ///                 GuestId = "other3xLinux64Guest",
    ///                 Memory = 2048,
    ///                 NetworkInterfaces = 
    ///                 {
    ///                     new VSphere.Inputs.VirtualMachineNetworkInterfaceArgs
    ///                     {
    ///                         NetworkId = network.Apply(network =&gt; network.Id),
    ///                     },
    ///                 },
    ///                 NumCpus = 2,
    ///                 ResourcePoolId = cluster.Apply(cluster =&gt; cluster.ResourcePoolId),
    ///             }));
    ///         }
    ///         var clusterVmAntiAffinityRule = new VSphere.DatastoreClusterVmAntiAffinityRule("clusterVmAntiAffinityRule", new VSphere.DatastoreClusterVmAntiAffinityRuleArgs
    ///         {
    ///             DatastoreClusterId = datastoreCluster.Apply(datastoreCluster =&gt; datastoreCluster.Id),
    ///             VirtualMachineIds = vm.Select(__item =&gt; __item.Id).ToList(),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class DatastoreClusterVmAntiAffinityRule : Pulumi.CustomResource
    {
        /// <summary>
        /// The managed object reference
        /// ID of the datastore cluster to put the group in.  Forces
        /// a new resource if changed.
        /// </summary>
        [Output("datastoreClusterId")]
        public Output<string> DatastoreClusterId { get; private set; } = null!;

        /// <summary>
        /// Enable this rule in the cluster. Default: `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// When this value is `true`, prevents any virtual
        /// machine operations that may violate this rule. Default: `false`.
        /// </summary>
        [Output("mandatory")]
        public Output<bool?> Mandatory { get; private set; } = null!;

        /// <summary>
        /// The name of the rule. This must be unique in the cluster.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The UUIDs of the virtual machines to run
        /// on different datastores from each other.
        /// </summary>
        [Output("virtualMachineIds")]
        public Output<ImmutableArray<string>> VirtualMachineIds { get; private set; } = null!;


        /// <summary>
        /// Create a DatastoreClusterVmAntiAffinityRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatastoreClusterVmAntiAffinityRule(string name, DatastoreClusterVmAntiAffinityRuleArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/datastoreClusterVmAntiAffinityRule:DatastoreClusterVmAntiAffinityRule", name, args ?? new DatastoreClusterVmAntiAffinityRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatastoreClusterVmAntiAffinityRule(string name, Input<string> id, DatastoreClusterVmAntiAffinityRuleState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/datastoreClusterVmAntiAffinityRule:DatastoreClusterVmAntiAffinityRule", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DatastoreClusterVmAntiAffinityRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatastoreClusterVmAntiAffinityRule Get(string name, Input<string> id, DatastoreClusterVmAntiAffinityRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new DatastoreClusterVmAntiAffinityRule(name, id, state, options);
        }
    }

    public sealed class DatastoreClusterVmAntiAffinityRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The managed object reference
        /// ID of the datastore cluster to put the group in.  Forces
        /// a new resource if changed.
        /// </summary>
        [Input("datastoreClusterId", required: true)]
        public Input<string> DatastoreClusterId { get; set; } = null!;

        /// <summary>
        /// Enable this rule in the cluster. Default: `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// When this value is `true`, prevents any virtual
        /// machine operations that may violate this rule. Default: `false`.
        /// </summary>
        [Input("mandatory")]
        public Input<bool>? Mandatory { get; set; }

        /// <summary>
        /// The name of the rule. This must be unique in the cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("virtualMachineIds", required: true)]
        private InputList<string>? _virtualMachineIds;

        /// <summary>
        /// The UUIDs of the virtual machines to run
        /// on different datastores from each other.
        /// </summary>
        public InputList<string> VirtualMachineIds
        {
            get => _virtualMachineIds ?? (_virtualMachineIds = new InputList<string>());
            set => _virtualMachineIds = value;
        }

        public DatastoreClusterVmAntiAffinityRuleArgs()
        {
        }
    }

    public sealed class DatastoreClusterVmAntiAffinityRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The managed object reference
        /// ID of the datastore cluster to put the group in.  Forces
        /// a new resource if changed.
        /// </summary>
        [Input("datastoreClusterId")]
        public Input<string>? DatastoreClusterId { get; set; }

        /// <summary>
        /// Enable this rule in the cluster. Default: `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// When this value is `true`, prevents any virtual
        /// machine operations that may violate this rule. Default: `false`.
        /// </summary>
        [Input("mandatory")]
        public Input<bool>? Mandatory { get; set; }

        /// <summary>
        /// The name of the rule. This must be unique in the cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("virtualMachineIds")]
        private InputList<string>? _virtualMachineIds;

        /// <summary>
        /// The UUIDs of the virtual machines to run
        /// on different datastores from each other.
        /// </summary>
        public InputList<string> VirtualMachineIds
        {
            get => _virtualMachineIds ?? (_virtualMachineIds = new InputList<string>());
            set => _virtualMachineIds = value;
        }

        public DatastoreClusterVmAntiAffinityRuleState()
        {
        }
    }
}

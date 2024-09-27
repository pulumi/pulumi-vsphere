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
    /// The `vsphere.ComputeClusterVmAffinityRule` resource can be used to
    /// manage virtual machine affinity rules in a cluster, either created by the
    /// `vsphere.ComputeCluster` resource or looked up
    /// by the `vsphere.ComputeCluster` data source.
    /// 
    /// This rule can be used to tell a set of virtual machines to run together on the
    /// same host within a cluster. When configured, DRS will make a best effort to
    /// ensure that the virtual machines run on the same host, or prevent any operation
    /// that would keep that from happening, depending on the value of the
    /// `mandatory` flag.
    /// 
    /// &gt; An affinity rule can only be used to place virtual machines on the same
    /// _non-specific_ hosts. It cannot be used to pin virtual machines to a host.
    /// To enable this capability, use the
    /// `vsphere.ComputeClusterVmHostRule`
    /// resource.
    /// 
    /// &gt; **NOTE:** This resource requires vCenter Server and is not available on
    /// direct ESXi host connections.
    /// 
    /// ## Example Usage
    /// 
    /// The following example creates two virtual machines in a cluster using the
    /// `vsphere.VirtualMachine` resource, creating the
    /// virtual machines in the cluster looked up by the
    /// `vsphere.ComputeCluster` data source. It
    /// then creates an affinity rule for these two virtual machines, ensuring they
    /// will run on the same host whenever possible.
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
    ///     var cluster = VSphere.GetComputeCluster.Invoke(new()
    ///     {
    ///         Name = "cluster-01",
    ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
    ///     });
    /// 
    ///     var network = VSphere.GetNetwork.Invoke(new()
    ///     {
    ///         Name = "VM Network",
    ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
    ///     });
    /// 
    ///     var vm = new List&lt;VSphere.VirtualMachine&gt;();
    ///     for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
    ///     {
    ///         var range = new { Value = rangeIndex };
    ///         vm.Add(new VSphere.VirtualMachine($"vm-{range.Value}", new()
    ///         {
    ///             Name = $"foo-{range.Value}",
    ///             ResourcePoolId = cluster.Apply(getComputeClusterResult =&gt; getComputeClusterResult.ResourcePoolId),
    ///             DatastoreId = datastore.Apply(getDatastoreResult =&gt; getDatastoreResult.Id),
    ///             NumCpus = 1,
    ///             Memory = 1024,
    ///             GuestId = "otherLinux64Guest",
    ///             NetworkInterfaces = new[]
    ///             {
    ///                 new VSphere.Inputs.VirtualMachineNetworkInterfaceArgs
    ///                 {
    ///                     NetworkId = network.Apply(getNetworkResult =&gt; getNetworkResult.Id),
    ///                 },
    ///             },
    ///             Disks = new[]
    ///             {
    ///                 new VSphere.Inputs.VirtualMachineDiskArgs
    ///                 {
    ///                     Label = "disk0",
    ///                     Size = 20,
    ///                 },
    ///             },
    ///         }));
    ///     }
    ///     var vmAffinityRule = new VSphere.ComputeClusterVmAffinityRule("vm_affinity_rule", new()
    ///     {
    ///         Name = "vm-affinity-rule",
    ///         ComputeClusterId = cluster.Apply(getComputeClusterResult =&gt; getComputeClusterResult.Id),
    ///         VirtualMachineIds = vm.Select((value, i) =&gt; new { Key = i.ToString(), Value = pair.Value }).Select(v =&gt; 
    ///         {
    ///             return v.Id;
    ///         }).ToList(),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// The following example creates an affinity rule for a set of virtual machines
    /// in the cluster by looking up the virtual machine UUIDs from the
    /// `vsphere.VirtualMachine` data source.
    /// 
    /// ## Import
    /// 
    /// An existing rule can be imported into this resource by supplying
    /// 
    /// both the path to the cluster, and the name the rule. If the name or cluster is
    /// 
    /// not found, or if the rule is of a different type, an error will be returned. An
    /// 
    /// example is below:
    /// 
    /// ```sh
    /// $ pulumi import vsphere:index/computeClusterVmAffinityRule:ComputeClusterVmAffinityRule vm_affinity_rule \
    /// ```
    /// 
    ///   '{"compute_cluster_path": "/dc-01/host/cluster-01", \
    /// 
    ///   "name": "vm-affinity-rule"}'
    /// </summary>
    [VSphereResourceType("vsphere:index/computeClusterVmAffinityRule:ComputeClusterVmAffinityRule")]
    public partial class ComputeClusterVmAffinityRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The managed object reference
        /// ID of the cluster to put the group in.  Forces a new
        /// resource if changed.
        /// </summary>
        [Output("computeClusterId")]
        public Output<string> ComputeClusterId { get; private set; } = null!;

        /// <summary>
        /// Enable this rule in the cluster. Default: `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// When this value is `true`, prevents any virtual
        /// machine operations that may violate this rule. Default: `false`.
        /// 
        /// &gt; **NOTE:** The namespace for rule names on this resource (defined by the
        /// `name` argument) is shared with all rules in the cluster - consider
        /// this when naming your rules.
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
        /// on the same host together.
        /// </summary>
        [Output("virtualMachineIds")]
        public Output<ImmutableArray<string>> VirtualMachineIds { get; private set; } = null!;


        /// <summary>
        /// Create a ComputeClusterVmAffinityRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ComputeClusterVmAffinityRule(string name, ComputeClusterVmAffinityRuleArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/computeClusterVmAffinityRule:ComputeClusterVmAffinityRule", name, args ?? new ComputeClusterVmAffinityRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ComputeClusterVmAffinityRule(string name, Input<string> id, ComputeClusterVmAffinityRuleState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/computeClusterVmAffinityRule:ComputeClusterVmAffinityRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ComputeClusterVmAffinityRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ComputeClusterVmAffinityRule Get(string name, Input<string> id, ComputeClusterVmAffinityRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new ComputeClusterVmAffinityRule(name, id, state, options);
        }
    }

    public sealed class ComputeClusterVmAffinityRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The managed object reference
        /// ID of the cluster to put the group in.  Forces a new
        /// resource if changed.
        /// </summary>
        [Input("computeClusterId", required: true)]
        public Input<string> ComputeClusterId { get; set; } = null!;

        /// <summary>
        /// Enable this rule in the cluster. Default: `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// When this value is `true`, prevents any virtual
        /// machine operations that may violate this rule. Default: `false`.
        /// 
        /// &gt; **NOTE:** The namespace for rule names on this resource (defined by the
        /// `name` argument) is shared with all rules in the cluster - consider
        /// this when naming your rules.
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
        /// on the same host together.
        /// </summary>
        public InputList<string> VirtualMachineIds
        {
            get => _virtualMachineIds ?? (_virtualMachineIds = new InputList<string>());
            set => _virtualMachineIds = value;
        }

        public ComputeClusterVmAffinityRuleArgs()
        {
        }
        public static new ComputeClusterVmAffinityRuleArgs Empty => new ComputeClusterVmAffinityRuleArgs();
    }

    public sealed class ComputeClusterVmAffinityRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The managed object reference
        /// ID of the cluster to put the group in.  Forces a new
        /// resource if changed.
        /// </summary>
        [Input("computeClusterId")]
        public Input<string>? ComputeClusterId { get; set; }

        /// <summary>
        /// Enable this rule in the cluster. Default: `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// When this value is `true`, prevents any virtual
        /// machine operations that may violate this rule. Default: `false`.
        /// 
        /// &gt; **NOTE:** The namespace for rule names on this resource (defined by the
        /// `name` argument) is shared with all rules in the cluster - consider
        /// this when naming your rules.
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
        /// on the same host together.
        /// </summary>
        public InputList<string> VirtualMachineIds
        {
            get => _virtualMachineIds ?? (_virtualMachineIds = new InputList<string>());
            set => _virtualMachineIds = value;
        }

        public ComputeClusterVmAffinityRuleState()
        {
        }
        public static new ComputeClusterVmAffinityRuleState Empty => new ComputeClusterVmAffinityRuleState();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    /// <summary>
    /// The `vsphere..ComputeClusterVmGroup` resource can be used to manage groups of
    /// virtual machines in a cluster, either created by the
    /// [`vsphere..ComputeCluster`][tf-vsphere-cluster-resource] resource or looked up
    /// by the [`vsphere..ComputeCluster`][tf-vsphere-cluster-data-source] data source.
    /// 
    /// [tf-vsphere-cluster-resource]: /docs/providers/vsphere/r/compute_cluster.html
    /// [tf-vsphere-cluster-data-source]: /docs/providers/vsphere/d/compute_cluster.html
    /// 
    /// This resource mainly serves as an input to the
    /// [`vsphere..ComputeClusterVmDependencyRule`][tf-vsphere-cluster-vm-dependency-rule-resource]
    /// and
    /// [`vsphere..ComputeClusterVmHostRule`][tf-vsphere-cluster-vm-host-rule-resource]
    /// resources. See the individual resource documentation pages for more information.
    /// 
    /// [tf-vsphere-cluster-vm-dependency-rule-resource]: /docs/providers/vsphere/r/compute_cluster_vm_dependency_rule.html
    /// [tf-vsphere-cluster-vm-host-rule-resource]: /docs/providers/vsphere/r/compute_cluster_vm_host_rule.html
    /// 
    /// &gt; **NOTE:** This resource requires vCenter and is not available on direct ESXi
    /// connections.
    /// 
    /// &gt; **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/compute_cluster_vm_group.html.markdown.
    /// </summary>
    public partial class ComputeClusterVmGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The [managed object reference
        /// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
        /// resource if changed.
        /// </summary>
        [Output("computeClusterId")]
        public Output<string> ComputeClusterId { get; private set; } = null!;

        /// <summary>
        /// The name of the VM group. This must be unique in the
        /// cluster. Forces a new resource if changed.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The UUIDs of the virtual machines in this
        /// group.
        /// </summary>
        [Output("virtualMachineIds")]
        public Output<ImmutableArray<string>> VirtualMachineIds { get; private set; } = null!;


        /// <summary>
        /// Create a ComputeClusterVmGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ComputeClusterVmGroup(string name, ComputeClusterVmGroupArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/computeClusterVmGroup:ComputeClusterVmGroup", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ComputeClusterVmGroup(string name, Input<string> id, ComputeClusterVmGroupState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/computeClusterVmGroup:ComputeClusterVmGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ComputeClusterVmGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ComputeClusterVmGroup Get(string name, Input<string> id, ComputeClusterVmGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ComputeClusterVmGroup(name, id, state, options);
        }
    }

    public sealed class ComputeClusterVmGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [managed object reference
        /// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
        /// resource if changed.
        /// </summary>
        [Input("computeClusterId", required: true)]
        public Input<string> ComputeClusterId { get; set; } = null!;

        /// <summary>
        /// The name of the VM group. This must be unique in the
        /// cluster. Forces a new resource if changed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("virtualMachineIds")]
        private InputList<string>? _virtualMachineIds;

        /// <summary>
        /// The UUIDs of the virtual machines in this
        /// group.
        /// </summary>
        public InputList<string> VirtualMachineIds
        {
            get => _virtualMachineIds ?? (_virtualMachineIds = new InputList<string>());
            set => _virtualMachineIds = value;
        }

        public ComputeClusterVmGroupArgs()
        {
        }
    }

    public sealed class ComputeClusterVmGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [managed object reference
        /// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
        /// resource if changed.
        /// </summary>
        [Input("computeClusterId")]
        public Input<string>? ComputeClusterId { get; set; }

        /// <summary>
        /// The name of the VM group. This must be unique in the
        /// cluster. Forces a new resource if changed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("virtualMachineIds")]
        private InputList<string>? _virtualMachineIds;

        /// <summary>
        /// The UUIDs of the virtual machines in this
        /// group.
        /// </summary>
        public InputList<string> VirtualMachineIds
        {
            get => _virtualMachineIds ?? (_virtualMachineIds = new InputList<string>());
            set => _virtualMachineIds = value;
        }

        public ComputeClusterVmGroupState()
        {
        }
    }
}

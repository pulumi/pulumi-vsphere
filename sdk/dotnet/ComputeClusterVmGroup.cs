// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    [VSphereResourceType("vsphere:index/computeClusterVmGroup:ComputeClusterVmGroup")]
    public partial class ComputeClusterVmGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The managed object reference
        /// ID of the cluster to put the group in.  Forces a new
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
        /// 
        /// &gt; **NOTE:** The namespace for cluster names on this resource (defined by the
        /// `name` argument) is shared with the
        /// `vsphere.ComputeClusterHostGroup`
        /// resource. Make sure your names are unique across both resources.
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
            : base("vsphere:index/computeClusterVmGroup:ComputeClusterVmGroup", name, args ?? new ComputeClusterVmGroupArgs(), MakeResourceOptions(options, ""))
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

    public sealed class ComputeClusterVmGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The managed object reference
        /// ID of the cluster to put the group in.  Forces a new
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
        /// 
        /// &gt; **NOTE:** The namespace for cluster names on this resource (defined by the
        /// `name` argument) is shared with the
        /// `vsphere.ComputeClusterHostGroup`
        /// resource. Make sure your names are unique across both resources.
        /// </summary>
        public InputList<string> VirtualMachineIds
        {
            get => _virtualMachineIds ?? (_virtualMachineIds = new InputList<string>());
            set => _virtualMachineIds = value;
        }

        public ComputeClusterVmGroupArgs()
        {
        }
        public static new ComputeClusterVmGroupArgs Empty => new ComputeClusterVmGroupArgs();
    }

    public sealed class ComputeClusterVmGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The managed object reference
        /// ID of the cluster to put the group in.  Forces a new
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
        /// 
        /// &gt; **NOTE:** The namespace for cluster names on this resource (defined by the
        /// `name` argument) is shared with the
        /// `vsphere.ComputeClusterHostGroup`
        /// resource. Make sure your names are unique across both resources.
        /// </summary>
        public InputList<string> VirtualMachineIds
        {
            get => _virtualMachineIds ?? (_virtualMachineIds = new InputList<string>());
            set => _virtualMachineIds = value;
        }

        public ComputeClusterVmGroupState()
        {
        }
        public static new ComputeClusterVmGroupState Empty => new ComputeClusterVmGroupState();
    }
}

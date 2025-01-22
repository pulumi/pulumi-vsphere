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
    /// The `vsphere.VirtualMachineSnapshot` resource can be used to manage snapshots
    /// for a virtual machine.
    /// 
    /// For more information on managing snapshots and how they work in VMware, see
    /// [here][ext-vm-snapshot-management].
    /// 
    /// [ext-vm-snapshot-management]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-single-host-management-vmware-host-client-8-0/virtual-machine-management-with-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/managing-virtual-machines-in-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/using-snapshots-to-manage-virtual-machines-vSphereSingleHostManagementVMwareHostClient.html
    /// 
    /// &gt; **NOTE:** A snapshot in VMware differs from traditional disk snapshots, and
    /// can contain the actual running state of the virtual machine, data for all disks
    /// that have not been set to be independent from the snapshot (including ones that
    /// have been attached via the `attach`
    /// parameter to the `vsphere.VirtualMachine` `disk` block), and even the
    /// configuration of the virtual machine at the time of the snapshot. Virtual
    /// machine, disk activity, and configuration changes post-snapshot are not
    /// included in the original state. Use this resource with care! Neither VMware nor
    /// HashiCorp recommends retaining snapshots for a extended period of time and does
    /// NOT recommend using them as as backup feature. For more information on the
    /// limitation of virtual machine snapshots, see [here][ext-vm-snap-limitations].
    /// 
    /// [ext-vm-snap-limitations]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-virtual-machine-administration-guide-8-0/managing-virtual-machinesvsphere-vm-admin/using-snapshots-to-manage-virtual-machinesvsphere-vm-admin/snapshot-file-names-and-descriptionvsphere-vm-admin.html
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
    ///     var demo1 = new VSphere.VirtualMachineSnapshot("demo1", new()
    ///     {
    ///         VirtualMachineUuid = "9aac5551-a351-4158-8c5c-15a71e8ec5c9",
    ///         SnapshotName = "Snapshot Name",
    ///         Description = "This is Demo Snapshot",
    ///         Memory = true,
    ///         Quiesce = true,
    ///         RemoveChildren = false,
    ///         Consolidate = true,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VSphereResourceType("vsphere:index/virtualMachineSnapshot:VirtualMachineSnapshot")]
    public partial class VirtualMachineSnapshot : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If set to `true`, the delta disks involved in this
        /// snapshot will be consolidated into the parent when this resource is
        /// destroyed.
        /// </summary>
        [Output("consolidate")]
        public Output<bool?> Consolidate { get; private set; } = null!;

        /// <summary>
        /// A description for the snapshot.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, a dump of the internal state of the
        /// virtual machine is included in the snapshot.
        /// </summary>
        [Output("memory")]
        public Output<bool> Memory { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, and the virtual machine is powered
        /// on when the snapshot is taken, VMware Tools is used to quiesce the file
        /// system in the virtual machine.
        /// </summary>
        [Output("quiesce")]
        public Output<bool> Quiesce { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, the entire snapshot subtree
        /// is removed when this resource is destroyed.
        /// </summary>
        [Output("removeChildren")]
        public Output<bool?> RemoveChildren { get; private set; } = null!;

        /// <summary>
        /// The name of the snapshot.
        /// </summary>
        [Output("snapshotName")]
        public Output<string> SnapshotName { get; private set; } = null!;

        /// <summary>
        /// The virtual machine UUID.
        /// </summary>
        [Output("virtualMachineUuid")]
        public Output<string> VirtualMachineUuid { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualMachineSnapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualMachineSnapshot(string name, VirtualMachineSnapshotArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/virtualMachineSnapshot:VirtualMachineSnapshot", name, args ?? new VirtualMachineSnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualMachineSnapshot(string name, Input<string> id, VirtualMachineSnapshotState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/virtualMachineSnapshot:VirtualMachineSnapshot", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualMachineSnapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualMachineSnapshot Get(string name, Input<string> id, VirtualMachineSnapshotState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualMachineSnapshot(name, id, state, options);
        }
    }

    public sealed class VirtualMachineSnapshotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to `true`, the delta disks involved in this
        /// snapshot will be consolidated into the parent when this resource is
        /// destroyed.
        /// </summary>
        [Input("consolidate")]
        public Input<bool>? Consolidate { get; set; }

        /// <summary>
        /// A description for the snapshot.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// If set to `true`, a dump of the internal state of the
        /// virtual machine is included in the snapshot.
        /// </summary>
        [Input("memory", required: true)]
        public Input<bool> Memory { get; set; } = null!;

        /// <summary>
        /// If set to `true`, and the virtual machine is powered
        /// on when the snapshot is taken, VMware Tools is used to quiesce the file
        /// system in the virtual machine.
        /// </summary>
        [Input("quiesce", required: true)]
        public Input<bool> Quiesce { get; set; } = null!;

        /// <summary>
        /// If set to `true`, the entire snapshot subtree
        /// is removed when this resource is destroyed.
        /// </summary>
        [Input("removeChildren")]
        public Input<bool>? RemoveChildren { get; set; }

        /// <summary>
        /// The name of the snapshot.
        /// </summary>
        [Input("snapshotName", required: true)]
        public Input<string> SnapshotName { get; set; } = null!;

        /// <summary>
        /// The virtual machine UUID.
        /// </summary>
        [Input("virtualMachineUuid", required: true)]
        public Input<string> VirtualMachineUuid { get; set; } = null!;

        public VirtualMachineSnapshotArgs()
        {
        }
        public static new VirtualMachineSnapshotArgs Empty => new VirtualMachineSnapshotArgs();
    }

    public sealed class VirtualMachineSnapshotState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to `true`, the delta disks involved in this
        /// snapshot will be consolidated into the parent when this resource is
        /// destroyed.
        /// </summary>
        [Input("consolidate")]
        public Input<bool>? Consolidate { get; set; }

        /// <summary>
        /// A description for the snapshot.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If set to `true`, a dump of the internal state of the
        /// virtual machine is included in the snapshot.
        /// </summary>
        [Input("memory")]
        public Input<bool>? Memory { get; set; }

        /// <summary>
        /// If set to `true`, and the virtual machine is powered
        /// on when the snapshot is taken, VMware Tools is used to quiesce the file
        /// system in the virtual machine.
        /// </summary>
        [Input("quiesce")]
        public Input<bool>? Quiesce { get; set; }

        /// <summary>
        /// If set to `true`, the entire snapshot subtree
        /// is removed when this resource is destroyed.
        /// </summary>
        [Input("removeChildren")]
        public Input<bool>? RemoveChildren { get; set; }

        /// <summary>
        /// The name of the snapshot.
        /// </summary>
        [Input("snapshotName")]
        public Input<string>? SnapshotName { get; set; }

        /// <summary>
        /// The virtual machine UUID.
        /// </summary>
        [Input("virtualMachineUuid")]
        public Input<string>? VirtualMachineUuid { get; set; }

        public VirtualMachineSnapshotState()
        {
        }
        public static new VirtualMachineSnapshotState Empty => new VirtualMachineSnapshotState();
    }
}

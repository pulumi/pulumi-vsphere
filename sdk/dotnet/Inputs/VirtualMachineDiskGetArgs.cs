// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class VirtualMachineDiskGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If this is true, the disk is attached instead of created. Implies keep_on_remove.
        /// </summary>
        [Input("attach")]
        public Input<bool>? Attach { get; set; }

        /// <summary>
        /// The type of controller the disk should be connected to. Must be 'scsi', 'sata', 'nvme', or 'ide'.
        /// </summary>
        [Input("controllerType")]
        public Input<string>? ControllerType { get; set; }

        /// <summary>
        /// The datastore ID for this virtual disk, if different than the virtual machine.
        /// </summary>
        [Input("datastoreId")]
        public Input<string>? DatastoreId { get; set; }

        /// <summary>
        /// The internally-computed address of this device, such as scsi:0:1, denoting scsi bus #0 and device unit 1.
        /// </summary>
        [Input("deviceAddress")]
        public Input<string>? DeviceAddress { get; set; }

        /// <summary>
        /// The mode of this this virtual disk for purposes of writes and snapshotting. Can be one of append, independent_nonpersistent, independent_persistent, nonpersistent, persistent, or undoable.
        /// </summary>
        [Input("diskMode")]
        public Input<string>? DiskMode { get; set; }

        /// <summary>
        /// The sharing mode of this virtual disk. Can be one of sharingMultiWriter or sharingNone.
        /// </summary>
        [Input("diskSharing")]
        public Input<string>? DiskSharing { get; set; }

        /// <summary>
        /// The virtual disk file zeroing policy when thin_provision is not true. The default is false, which lazily-zeros the disk, speeding up thick-provisioned disk creation time.
        /// </summary>
        [Input("eagerlyScrub")]
        public Input<bool>? EagerlyScrub { get; set; }

        /// <summary>
        /// The upper limit of IOPS that this disk can use.
        /// </summary>
        [Input("ioLimit")]
        public Input<int>? IoLimit { get; set; }

        /// <summary>
        /// The I/O guarantee that this disk has, in IOPS.
        /// </summary>
        [Input("ioReservation")]
        public Input<int>? IoReservation { get; set; }

        /// <summary>
        /// The share count for this disk when the share level is custom.
        /// </summary>
        [Input("ioShareCount")]
        public Input<int>? IoShareCount { get; set; }

        /// <summary>
        /// The share allocation level for this disk. Can be one of low, normal, high, or custom.
        /// </summary>
        [Input("ioShareLevel")]
        public Input<string>? IoShareLevel { get; set; }

        /// <summary>
        /// Set to true to keep the underlying VMDK file when removing this virtual disk from configuration.
        /// </summary>
        [Input("keepOnRemove")]
        public Input<bool>? KeepOnRemove { get; set; }

        /// <summary>
        /// The ID of the device within the virtual machine.
        /// </summary>
        [Input("key")]
        public Input<int>? Key { get; set; }

        /// <summary>
        /// A unique label for this disk.
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// The full path of the virtual disk. This can only be provided if attach is set to true, otherwise it is a read-only value.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The size of the disk, in GB.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The ID of the storage policy to assign to the virtual disk in VM.
        /// </summary>
        [Input("storagePolicyId")]
        public Input<string>? StoragePolicyId { get; set; }

        /// <summary>
        /// If true, this disk is thin provisioned, with space for the file being allocated on an as-needed basis.
        /// </summary>
        [Input("thinProvisioned")]
        public Input<bool>? ThinProvisioned { get; set; }

        /// <summary>
        /// The unique device number for this disk. This number determines where on the SCSI bus this device will be attached.
        /// </summary>
        [Input("unitNumber")]
        public Input<int>? UnitNumber { get; set; }

        /// <summary>
        /// The UUID of the virtual machine. Also exposed as the `id` of the resource.
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// If true, writes for this disk are sent directly to the filesystem immediately instead of being buffered.
        /// </summary>
        [Input("writeThrough")]
        public Input<bool>? WriteThrough { get; set; }

        public VirtualMachineDiskGetArgs()
        {
        }
        public static new VirtualMachineDiskGetArgs Empty => new VirtualMachineDiskGetArgs();
    }
}

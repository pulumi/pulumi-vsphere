// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class VirtualMachineDiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Attach an external disk instead of creating a new one. Implies and conflicts with `keep_on_remove`. If set, you cannot set `size`, `eagerly_scrub`, or `thin_provisioned`. Must set `path` if used.
        /// 
        /// &gt; **NOTE:** External disks cannot be attached when `datastore_cluster_id` is used.
        /// </summary>
        [Input("attach")]
        public Input<bool>? Attach { get; set; }

        /// <summary>
        /// The type of storage controller to attach the  disk to. Can be `scsi`, `sata`, or `ide`. You must have the appropriate number of controllers enabled for the selected type. Default `scsi`.
        /// </summary>
        [Input("controllerType")]
        public Input<string>? ControllerType { get; set; }

        /// <summary>
        /// The [managed object reference ID][docs-about-morefs] for the datastore on which the virtual disk is placed. The default is to use the datastore of the virtual machine. See the section on virtual machine migration for information on modifying this value.
        /// 
        /// &gt; **NOTE:** Datastores cannot be assigned to individual disks when `datastore_cluster_id` is used.
        /// </summary>
        [Input("datastoreId")]
        public Input<string>? DatastoreId { get; set; }

        [Input("deviceAddress")]
        public Input<string>? DeviceAddress { get; set; }

        /// <summary>
        /// The mode of this this virtual disk for purposes of writes and snapshots. One of `append`, `independent_nonpersistent`, `independent_persistent`, `nonpersistent`, `persistent`, or `undoable`. Default: `persistent`. For more information on these option, please refer to the [product documentation][vmware-docs-disk-mode].
        /// </summary>
        [Input("diskMode")]
        public Input<string>? DiskMode { get; set; }

        /// <summary>
        /// The sharing mode of this virtual disk. One of `sharingMultiWriter` or `sharingNone`. Default: `sharingNone`.
        /// 
        /// &gt; **NOTE:** Disk sharing is only available on vSphere 6.0 and later.
        /// </summary>
        [Input("diskSharing")]
        public Input<string>? DiskSharing { get; set; }

        /// <summary>
        /// If set to `true`, the disk space is zeroed out when the virtual machine is created. This will delay the creation of the virtual disk. Cannot be set to `true` when `thin_provisioned` is `true`.  See the section on picking a disk type for more information.  Default: `false`.
        /// </summary>
        [Input("eagerlyScrub")]
        public Input<bool>? EagerlyScrub { get; set; }

        /// <summary>
        /// The upper limit of IOPS that this disk can use. The default is no limit.
        /// </summary>
        [Input("ioLimit")]
        public Input<int>? IoLimit { get; set; }

        /// <summary>
        /// The I/O reservation (guarantee) for the virtual disk has, in IOPS.  The default is no reservation.
        /// </summary>
        [Input("ioReservation")]
        public Input<int>? IoReservation { get; set; }

        /// <summary>
        /// The share count for the virtual disk when the share level is `custom`.
        /// </summary>
        [Input("ioShareCount")]
        public Input<int>? IoShareCount { get; set; }

        /// <summary>
        /// The share allocation level for the virtual disk. One of `low`, `normal`, `high`, or `custom`. Default: `normal`.
        /// </summary>
        [Input("ioShareLevel")]
        public Input<string>? IoShareLevel { get; set; }

        /// <summary>
        /// Keep this disk when removing the device or destroying the virtual machine. Default: `false`.
        /// </summary>
        [Input("keepOnRemove")]
        public Input<bool>? KeepOnRemove { get; set; }

        /// <summary>
        /// The ID of the device within the virtual machine.
        /// </summary>
        [Input("key")]
        public Input<int>? Key { get; set; }

        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// When using `attach`, this parameter controls the path of a virtual disk to attach externally. Otherwise, it is a computed attribute that contains the virtual disk filename.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The size of the disk, in GB. Must be a whole number.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The UUID of the storage policy to assign to the virtual disk.
        /// </summary>
        [Input("storagePolicyId")]
        public Input<string>? StoragePolicyId { get; set; }

        /// <summary>
        /// If `true`, the disk is thin provisioned, with space for the file being allocated on an as-needed basis. Cannot be set to `true` when `eagerly_scrub` is `true`. See the section on selecting a disk type for more information. Default: `true`.
        /// </summary>
        [Input("thinProvisioned")]
        public Input<bool>? ThinProvisioned { get; set; }

        /// <summary>
        /// The disk number on the storage bus. The maximum value for this setting is the value of the controller count times the controller capacity (15 for SCSI, 30 for SATA, and 2 for IDE). Duplicate unit numbers are not allowed. Default `0`, for which one disk must be set to.
        /// </summary>
        [Input("unitNumber")]
        public Input<int>? UnitNumber { get; set; }

        /// <summary>
        /// The UUID of the virtual disk VMDK file. This is used to track the virtual disk on the virtual machine.
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// If `true`, writes for this disk are sent directly to the filesystem immediately instead of being buffered. Default: `false`.
        /// </summary>
        [Input("writeThrough")]
        public Input<bool>? WriteThrough { get; set; }

        public VirtualMachineDiskArgs()
        {
        }
        public static new VirtualMachineDiskArgs Empty => new VirtualMachineDiskArgs();
    }
}

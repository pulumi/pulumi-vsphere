// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Outputs
{

    [OutputType]
    public sealed class VirtualMachineDisk
    {
        /// <summary>
        /// If this is true, the disk is attached instead of created. Implies keep_on_remove.
        /// </summary>
        public readonly bool? Attach;
        /// <summary>
        /// The type of controller the disk should be connected to. Must be 'scsi', 'sata', 'nvme', or 'ide'.
        /// </summary>
        public readonly string? ControllerType;
        /// <summary>
        /// The datastore ID for this virtual disk, if different than the virtual machine.
        /// </summary>
        public readonly string? DatastoreId;
        /// <summary>
        /// The internally-computed address of this device, such as scsi:0:1, denoting scsi bus #0 and device unit 1.
        /// </summary>
        public readonly string? DeviceAddress;
        /// <summary>
        /// The mode of this this virtual disk for purposes of writes and snapshotting. Can be one of append, independent_nonpersistent, independent_persistent, nonpersistent, persistent, or undoable.
        /// </summary>
        public readonly string? DiskMode;
        /// <summary>
        /// The sharing mode of this virtual disk. Can be one of sharingMultiWriter or sharingNone.
        /// </summary>
        public readonly string? DiskSharing;
        /// <summary>
        /// The virtual disk file zeroing policy when thin_provision is not true. The default is false, which lazily-zeros the disk, speeding up thick-provisioned disk creation time.
        /// </summary>
        public readonly bool? EagerlyScrub;
        /// <summary>
        /// The upper limit of IOPS that this disk can use.
        /// </summary>
        public readonly int? IoLimit;
        /// <summary>
        /// The I/O guarantee that this disk has, in IOPS.
        /// </summary>
        public readonly int? IoReservation;
        /// <summary>
        /// The share count for this disk when the share level is custom.
        /// </summary>
        public readonly int? IoShareCount;
        /// <summary>
        /// The share allocation level for this disk. Can be one of low, normal, high, or custom.
        /// </summary>
        public readonly string? IoShareLevel;
        /// <summary>
        /// Set to true to keep the underlying VMDK file when removing this virtual disk from configuration.
        /// </summary>
        public readonly bool? KeepOnRemove;
        /// <summary>
        /// The ID of the device within the virtual machine.
        /// </summary>
        public readonly int? Key;
        /// <summary>
        /// A unique label for this disk.
        /// </summary>
        public readonly string Label;
        /// <summary>
        /// The full path of the virtual disk. This can only be provided if attach is set to true, otherwise it is a read-only value.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// The size of the disk, in GB.
        /// </summary>
        public readonly int? Size;
        /// <summary>
        /// The ID of the storage policy to assign to the virtual disk in VM.
        /// </summary>
        public readonly string? StoragePolicyId;
        /// <summary>
        /// If true, this disk is thin provisioned, with space for the file being allocated on an as-needed basis.
        /// </summary>
        public readonly bool? ThinProvisioned;
        /// <summary>
        /// The unique device number for this disk. This number determines where on the SCSI bus this device will be attached.
        /// </summary>
        public readonly int? UnitNumber;
        /// <summary>
        /// The UUID of the virtual machine. Also exposed as the `id` of the resource.
        /// </summary>
        public readonly string? Uuid;
        /// <summary>
        /// If true, writes for this disk are sent directly to the filesystem immediately instead of being buffered.
        /// </summary>
        public readonly bool? WriteThrough;

        [OutputConstructor]
        private VirtualMachineDisk(
            bool? attach,

            string? controllerType,

            string? datastoreId,

            string? deviceAddress,

            string? diskMode,

            string? diskSharing,

            bool? eagerlyScrub,

            int? ioLimit,

            int? ioReservation,

            int? ioShareCount,

            string? ioShareLevel,

            bool? keepOnRemove,

            int? key,

            string label,

            string? path,

            int? size,

            string? storagePolicyId,

            bool? thinProvisioned,

            int? unitNumber,

            string? uuid,

            bool? writeThrough)
        {
            Attach = attach;
            ControllerType = controllerType;
            DatastoreId = datastoreId;
            DeviceAddress = deviceAddress;
            DiskMode = diskMode;
            DiskSharing = diskSharing;
            EagerlyScrub = eagerlyScrub;
            IoLimit = ioLimit;
            IoReservation = ioReservation;
            IoShareCount = ioShareCount;
            IoShareLevel = ioShareLevel;
            KeepOnRemove = keepOnRemove;
            Key = key;
            Label = label;
            Path = path;
            Size = size;
            StoragePolicyId = storagePolicyId;
            ThinProvisioned = thinProvisioned;
            UnitNumber = unitNumber;
            Uuid = uuid;
            WriteThrough = writeThrough;
        }
    }
}

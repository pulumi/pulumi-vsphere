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
    public sealed class VirtualMachineCdrom
    {
        /// <summary>
        /// Indicates whether the device should be backed by
        /// remote client device. Conflicts with `datastore_id` and `path`.
        /// </summary>
        public readonly bool? ClientDevice;
        /// <summary>
        /// The datastore ID that the ISO is located in.
        /// Requried for using a datastore ISO. Conflicts with `client_device`.
        /// </summary>
        public readonly string? DatastoreId;
        /// <summary>
        /// An address internal to this provider that helps locate the
        /// device when `key` is unavailable. This follows a convention of
        /// `CONTROLLER_TYPE:BUS_NUMBER:UNIT_NUMBER`. Example: `scsi:0:1` means device
        /// unit 1 on SCSI bus 0.
        /// </summary>
        public readonly string? DeviceAddress;
        /// <summary>
        /// The ID of the device within the virtual machine.
        /// </summary>
        public readonly int? Key;
        /// <summary>
        /// The path to the ISO file. Required for using a datastore
        /// ISO. Conflicts with `client_device`.
        /// </summary>
        public readonly string? Path;

        [OutputConstructor]
        private VirtualMachineCdrom(
            bool? clientDevice,

            string? datastoreId,

            string? deviceAddress,

            int? key,

            string? path)
        {
            ClientDevice = clientDevice;
            DatastoreId = datastoreId;
            DeviceAddress = deviceAddress;
            Key = key;
            Path = path;
        }
    }
}
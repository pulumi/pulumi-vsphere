// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Outputs
{

    [OutputType]
    public sealed class VirtualMachineNetworkInterface
    {
        /// <summary>
        /// The controller type. Can be one of e1000, e1000e, sriov, vmxnet3, or vrdma.
        /// </summary>
        public readonly string? AdapterType;
        /// <summary>
        /// The upper bandwidth limit of this network interface, in Mbits/sec.
        /// </summary>
        public readonly int? BandwidthLimit;
        /// <summary>
        /// The bandwidth reservation of this network interface, in Mbits/sec.
        /// </summary>
        public readonly int? BandwidthReservation;
        /// <summary>
        /// The share count for this network interface when the share level is custom.
        /// </summary>
        public readonly int? BandwidthShareCount;
        /// <summary>
        /// The bandwidth share allocation level for this interface. Can be one of low, normal, high, or custom.
        /// </summary>
        public readonly string? BandwidthShareLevel;
        /// <summary>
        /// The internally-computed address of this device, such as scsi:0:1, denoting scsi bus #0 and device unit 1.
        /// </summary>
        public readonly string? DeviceAddress;
        /// <summary>
        /// The ID of the device within the virtual machine.
        /// </summary>
        public readonly int? Key;
        /// <summary>
        /// The MAC address of this network interface. Can only be manually set if use_static_mac is true.
        /// </summary>
        public readonly string? MacAddress;
        /// <summary>
        /// The ID of the network to connect this network interface to.
        /// </summary>
        public readonly string NetworkId;
        /// <summary>
        /// Mapping of network interface to OVF network.
        /// </summary>
        public readonly string? OvfMapping;
        /// <summary>
        /// The ID of the Physical SR-IOV NIC to attach to, e.g. '0000:d8:00.0'
        /// </summary>
        public readonly string? PhysicalFunction;
        /// <summary>
        /// If true, the mac_address field is treated as a static MAC address and set accordingly.
        /// </summary>
        public readonly bool? UseStaticMac;

        [OutputConstructor]
        private VirtualMachineNetworkInterface(
            string? adapterType,

            int? bandwidthLimit,

            int? bandwidthReservation,

            int? bandwidthShareCount,

            string? bandwidthShareLevel,

            string? deviceAddress,

            int? key,

            string? macAddress,

            string networkId,

            string? ovfMapping,

            string? physicalFunction,

            bool? useStaticMac)
        {
            AdapterType = adapterType;
            BandwidthLimit = bandwidthLimit;
            BandwidthReservation = bandwidthReservation;
            BandwidthShareCount = bandwidthShareCount;
            BandwidthShareLevel = bandwidthShareLevel;
            DeviceAddress = deviceAddress;
            Key = key;
            MacAddress = macAddress;
            NetworkId = networkId;
            OvfMapping = ovfMapping;
            PhysicalFunction = physicalFunction;
            UseStaticMac = useStaticMac;
        }
    }
}

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
    public sealed class VirtualMachineNetworkInterface
    {
        /// <summary>
        /// The network interface type. One of `e1000`, `e1000e`, `sriov`, or `vmxnet3`. Default: `vmxnet3`.
        /// </summary>
        public readonly string? AdapterType;
        /// <summary>
        /// The upper bandwidth limit of the network interface, in Mbits/sec. The default is no limit. Ignored if `adapter_type` is set to `sriov`.
        /// </summary>
        public readonly int? BandwidthLimit;
        /// <summary>
        /// The bandwidth reservation of the network interface, in Mbits/sec. The default is no reservation.
        /// </summary>
        public readonly int? BandwidthReservation;
        /// <summary>
        /// The share count for the network interface when the share level is `custom`. Ignored if `adapter_type` is set to `sriov`.
        /// </summary>
        public readonly int? BandwidthShareCount;
        /// <summary>
        /// The bandwidth share allocation level for the network interface. One of `low`, `normal`, `high`, or `custom`. Default: `normal`. Ignored if `adapter_type` is set to `sriov`.
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
        /// The MAC address of the network interface. Can only be manually set if `use_static_mac` is `true`. Otherwise, the value is computed and presents the assigned MAC address for the interface.
        /// </summary>
        public readonly string? MacAddress;
        /// <summary>
        /// The [managed object reference ID][docs-about-morefs] of the network on which to connect the virtual machine network interface.
        /// </summary>
        public readonly string NetworkId;
        /// <summary>
        /// Specifies which NIC in an OVF/OVA the `network_interface` should be associated. Only applies at creation when deploying from an OVF/OVA.
        /// </summary>
        public readonly string? OvfMapping;
        /// <summary>
        /// The ID of the Physical SR-IOV NIC to attach to, e.g. '0000:d8:00.0'
        /// </summary>
        public readonly string? PhysicalFunction;
        /// <summary>
        /// If true, the `mac_address` field is treated as a static MAC address and set accordingly. Setting this to `true` requires `mac_address` to be set. Default: `false`.
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

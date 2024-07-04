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
    public sealed class GetVirtualMachineNetworkInterfaceResult
    {
        /// <summary>
        /// The network interface types for each network interface found
        /// on the virtual machine, in device bus order. Will be one of `e1000`,
        /// `e1000e`, `vmxnet3vrdma`, or `vmxnet3`.
        /// </summary>
        public readonly string AdapterType;
        /// <summary>
        /// The upper bandwidth limit of this network interface,
        /// in Mbits/sec.
        /// </summary>
        public readonly int? BandwidthLimit;
        /// <summary>
        /// The bandwidth reservation of this network
        /// interface, in Mbits/sec.
        /// </summary>
        public readonly int? BandwidthReservation;
        /// <summary>
        /// The share count for this network interface when the
        /// share level is custom.
        /// </summary>
        public readonly int BandwidthShareCount;
        /// <summary>
        /// The bandwidth share allocation level for this
        /// interface. Can be one of `low`, `normal`, `high`, or `custom`.
        /// </summary>
        public readonly string? BandwidthShareLevel;
        /// <summary>
        /// The MAC address of this network interface.
        /// </summary>
        public readonly string MacAddress;
        /// <summary>
        /// The managed object reference ID of the network this interface
        /// is connected to.
        /// </summary>
        public readonly string NetworkId;
        /// <summary>
        /// The ID of the Physical SR-IOV NIC to attach to, e.g. '0000:d8:00.0'
        /// </summary>
        public readonly string PhysicalFunction;

        [OutputConstructor]
        private GetVirtualMachineNetworkInterfaceResult(
            string adapterType,

            int? bandwidthLimit,

            int? bandwidthReservation,

            int bandwidthShareCount,

            string? bandwidthShareLevel,

            string macAddress,

            string networkId,

            string physicalFunction)
        {
            AdapterType = adapterType;
            BandwidthLimit = bandwidthLimit;
            BandwidthReservation = bandwidthReservation;
            BandwidthShareCount = bandwidthShareCount;
            BandwidthShareLevel = bandwidthShareLevel;
            MacAddress = macAddress;
            NetworkId = networkId;
            PhysicalFunction = physicalFunction;
        }
    }
}

// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class VirtualMachineNetworkInterfaceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The controller type. Can be one of e1000, e1000e, sriov, vmxnet3, or vrdma.
        /// </summary>
        [Input("adapterType")]
        public Input<string>? AdapterType { get; set; }

        /// <summary>
        /// The upper bandwidth limit of this network interface, in Mbits/sec.
        /// </summary>
        [Input("bandwidthLimit")]
        public Input<int>? BandwidthLimit { get; set; }

        /// <summary>
        /// The bandwidth reservation of this network interface, in Mbits/sec.
        /// </summary>
        [Input("bandwidthReservation")]
        public Input<int>? BandwidthReservation { get; set; }

        /// <summary>
        /// The share count for this network interface when the share level is custom.
        /// </summary>
        [Input("bandwidthShareCount")]
        public Input<int>? BandwidthShareCount { get; set; }

        /// <summary>
        /// The bandwidth share allocation level for this interface. Can be one of low, normal, high, or custom.
        /// </summary>
        [Input("bandwidthShareLevel")]
        public Input<string>? BandwidthShareLevel { get; set; }

        /// <summary>
        /// The internally-computed address of this device, such as scsi:0:1, denoting scsi bus #0 and device unit 1.
        /// </summary>
        [Input("deviceAddress")]
        public Input<string>? DeviceAddress { get; set; }

        /// <summary>
        /// The ID of the device within the virtual machine.
        /// </summary>
        [Input("key")]
        public Input<int>? Key { get; set; }

        /// <summary>
        /// The MAC address of this network interface. Can only be manually set if use_static_mac is true.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// The ID of the network to connect this network interface to.
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        /// <summary>
        /// Mapping of network interface to OVF network.
        /// </summary>
        [Input("ovfMapping")]
        public Input<string>? OvfMapping { get; set; }

        /// <summary>
        /// The ID of the Physical SR-IOV NIC to attach to, e.g. '0000:d8:00.0'
        /// </summary>
        [Input("physicalFunction")]
        public Input<string>? PhysicalFunction { get; set; }

        /// <summary>
        /// If true, the mac_address field is treated as a static MAC address and set accordingly.
        /// </summary>
        [Input("useStaticMac")]
        public Input<bool>? UseStaticMac { get; set; }

        public VirtualMachineNetworkInterfaceGetArgs()
        {
        }
        public static new VirtualMachineNetworkInterfaceGetArgs Empty => new VirtualMachineNetworkInterfaceGetArgs();
    }
}

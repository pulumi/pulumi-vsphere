// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
        /// The network interface type. One of `e1000`, `e1000e`, `sriov`, or `vmxnet3`. Default: `vmxnet3`.
        /// </summary>
        [Input("adapterType")]
        public Input<string>? AdapterType { get; set; }

        /// <summary>
        /// The upper bandwidth limit of the network interface, in Mbits/sec. The default is no limit. Ignored if `adapter_type` is set to `sriov`.
        /// </summary>
        [Input("bandwidthLimit")]
        public Input<int>? BandwidthLimit { get; set; }

        /// <summary>
        /// The bandwidth reservation of the network interface, in Mbits/sec. The default is no reservation.
        /// </summary>
        [Input("bandwidthReservation")]
        public Input<int>? BandwidthReservation { get; set; }

        /// <summary>
        /// The share count for the network interface when the share level is `custom`. Ignored if `adapter_type` is set to `sriov`.
        /// </summary>
        [Input("bandwidthShareCount")]
        public Input<int>? BandwidthShareCount { get; set; }

        /// <summary>
        /// The bandwidth share allocation level for the network interface. One of `low`, `normal`, `high`, or `custom`. Default: `normal`. Ignored if `adapter_type` is set to `sriov`.
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
        /// The MAC address of the network interface. Can only be manually set if `use_static_mac` is `true`. Otherwise, the value is computed and presents the assigned MAC address for the interface.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// The [managed object reference ID][docs-about-morefs] of the network on which to connect the virtual machine network interface.
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        /// <summary>
        /// Specifies which NIC in an OVF/OVA the `network_interface` should be associated. Only applies at creation when deploying from an OVF/OVA.
        /// </summary>
        [Input("ovfMapping")]
        public Input<string>? OvfMapping { get; set; }

        /// <summary>
        /// The ID of the Physical SR-IOV NIC to attach to, e.g. '0000:d8:00.0'
        /// </summary>
        [Input("physicalFunction")]
        public Input<string>? PhysicalFunction { get; set; }

        /// <summary>
        /// If true, the `mac_address` field is treated as a static MAC address and set accordingly. Setting this to `true` requires `mac_address` to be set. Default: `false`.
        /// </summary>
        [Input("useStaticMac")]
        public Input<bool>? UseStaticMac { get; set; }

        public VirtualMachineNetworkInterfaceGetArgs()
        {
        }
        public static new VirtualMachineNetworkInterfaceGetArgs Empty => new VirtualMachineNetworkInterfaceGetArgs();
    }
}

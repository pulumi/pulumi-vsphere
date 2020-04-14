// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class VirtualMachineCloneCustomizeArgs : Pulumi.ResourceArgs
    {
        [Input("dnsServerLists")]
        private InputList<string>? _dnsServerLists;
        public InputList<string> DnsServerLists
        {
            get => _dnsServerLists ?? (_dnsServerLists = new InputList<string>());
            set => _dnsServerLists = value;
        }

        [Input("dnsSuffixLists")]
        private InputList<string>? _dnsSuffixLists;
        public InputList<string> DnsSuffixLists
        {
            get => _dnsSuffixLists ?? (_dnsSuffixLists = new InputList<string>());
            set => _dnsSuffixLists = value;
        }

        [Input("ipv4Gateway")]
        public Input<string>? Ipv4Gateway { get; set; }

        [Input("ipv6Gateway")]
        public Input<string>? Ipv6Gateway { get; set; }

        [Input("linuxOptions")]
        public Input<Inputs.VirtualMachineCloneCustomizeLinuxOptionsArgs>? LinuxOptions { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.VirtualMachineCloneCustomizeNetworkInterfaceArgs>? _networkInterfaces;

        /// <summary>
        /// A specification for a virtual NIC on this
        /// virtual machine. See network interface options
        /// below.
        /// </summary>
        public InputList<Inputs.VirtualMachineCloneCustomizeNetworkInterfaceArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.VirtualMachineCloneCustomizeNetworkInterfaceArgs>());
            set => _networkInterfaces = value;
        }

        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        [Input("windowsOptions")]
        public Input<Inputs.VirtualMachineCloneCustomizeWindowsOptionsArgs>? WindowsOptions { get; set; }

        [Input("windowsSysprepText")]
        public Input<string>? WindowsSysprepText { get; set; }

        public VirtualMachineCloneCustomizeArgs()
        {
        }
    }
}

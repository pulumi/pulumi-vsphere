// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class VnicIpv6Args : Pulumi.ResourceArgs
    {
        [Input("addresses")]
        private InputList<string>? _addresses;

        /// <summary>
        /// List of IPv6 addresses
        /// </summary>
        public InputList<string> Addresses
        {
            get => _addresses ?? (_addresses = new InputList<string>());
            set => _addresses = value;
        }

        /// <summary>
        /// Use IPv6 Autoconfiguration (RFC2462).
        /// </summary>
        [Input("autoconfig")]
        public Input<bool>? Autoconfig { get; set; }

        /// <summary>
        /// Use DHCP to configure the interface's IPv4 stack.
        /// </summary>
        [Input("dhcp")]
        public Input<bool>? Dhcp { get; set; }

        /// <summary>
        /// IP address of the default gateway, if DHCP or autoconfig is not set.
        /// </summary>
        [Input("gw")]
        public Input<string>? Gw { get; set; }

        public VnicIpv6Args()
        {
        }
    }
}
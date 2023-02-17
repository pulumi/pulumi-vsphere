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
    public sealed class VnicIpv4
    {
        /// <summary>
        /// Use DHCP to configure the interface's IPv4 stack.
        /// </summary>
        public readonly bool? Dhcp;
        /// <summary>
        /// IP address of the default gateway, if DHCP is not set.
        /// </summary>
        public readonly string? Gw;
        /// <summary>
        /// Address of the interface, if DHCP is not set.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// Netmask of the interface, if DHCP is not set.
        /// </summary>
        public readonly string? Netmask;

        [OutputConstructor]
        private VnicIpv4(
            bool? dhcp,

            string? gw,

            string? ip,

            string? netmask)
        {
            Dhcp = dhcp;
            Gw = gw;
            Ip = ip;
            Netmask = netmask;
        }
    }
}

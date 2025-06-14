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
    public sealed class SupervisorManagementNetwork
    {
        /// <summary>
        /// Number of addresses to allocate. Starts from 'starting_address'
        /// </summary>
        public readonly int AddressCount;
        /// <summary>
        /// Gateway IP address.
        /// </summary>
        public readonly string Gateway;
        /// <summary>
        /// ID of the network. (e.g. a distributed port group).
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Starting address of the management network range.
        /// </summary>
        public readonly string StartingAddress;
        /// <summary>
        /// Subnet mask.
        /// </summary>
        public readonly string SubnetMask;

        [OutputConstructor]
        private SupervisorManagementNetwork(
            int addressCount,

            string gateway,

            string network,

            string startingAddress,

            string subnetMask)
        {
            AddressCount = addressCount;
            Gateway = gateway;
            Network = network;
            StartingAddress = startingAddress;
            SubnetMask = subnetMask;
        }
    }
}

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
    public sealed class DistributedVirtualSwitchHost
    {
        /// <summary>
        /// The list of NIC devices to map to uplinks on the VDS,
        /// added in order they are specified.
        /// </summary>
        public readonly ImmutableArray<string> Devices;
        /// <summary>
        /// The host system ID of the host to add to the
        /// VDS.
        /// </summary>
        public readonly string HostSystemId;

        [OutputConstructor]
        private DistributedVirtualSwitchHost(
            ImmutableArray<string> devices,

            string hostSystemId)
        {
            Devices = devices;
            HostSystemId = hostSystemId;
        }
    }
}

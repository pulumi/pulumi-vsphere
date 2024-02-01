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
    public sealed class DistributedVirtualSwitchVlanRange
    {
        /// <summary>
        /// The minimum VLAN to use in the range.
        /// </summary>
        public readonly int MaxVlan;
        /// <summary>
        /// The minimum VLAN to use in the range.
        /// </summary>
        public readonly int MinVlan;

        [OutputConstructor]
        private DistributedVirtualSwitchVlanRange(
            int maxVlan,

            int minVlan)
        {
            MaxVlan = maxVlan;
            MinVlan = minVlan;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class DistributedVirtualSwitchPvlanMappingGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The primary VLAN ID. The VLAN IDs of 0 and
        /// 4095 are reserved and cannot be used in this property.
        /// </summary>
        [Input("primaryVlanId", required: true)]
        public Input<int> PrimaryVlanId { get; set; } = null!;

        /// <summary>
        /// The private VLAN type. Valid values are
        /// promiscuous, community and isolated.
        /// </summary>
        [Input("pvlanType", required: true)]
        public Input<string> PvlanType { get; set; } = null!;

        /// <summary>
        /// The secondary VLAN ID. The VLAN IDs of 0
        /// and 4095 are reserved and cannot be used in this property.
        /// </summary>
        [Input("secondaryVlanId", required: true)]
        public Input<int> SecondaryVlanId { get; set; } = null!;

        public DistributedVirtualSwitchPvlanMappingGetArgs()
        {
        }
        public static new DistributedVirtualSwitchPvlanMappingGetArgs Empty => new DistributedVirtualSwitchPvlanMappingGetArgs();
    }
}

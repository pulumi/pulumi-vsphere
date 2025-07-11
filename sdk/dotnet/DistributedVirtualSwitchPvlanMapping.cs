// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    [VSphereResourceType("vsphere:index/distributedVirtualSwitchPvlanMapping:DistributedVirtualSwitchPvlanMapping")]
    public partial class DistributedVirtualSwitchPvlanMapping : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the distributed virtual switch to attach this mapping to.
        /// </summary>
        [Output("distributedVirtualSwitchId")]
        public Output<string> DistributedVirtualSwitchId { get; private set; } = null!;

        /// <summary>
        /// The primary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        /// </summary>
        [Output("primaryVlanId")]
        public Output<int> PrimaryVlanId { get; private set; } = null!;

        /// <summary>
        /// The private VLAN type. Valid values are promiscuous, community and isolated.
        /// </summary>
        [Output("pvlanType")]
        public Output<string> PvlanType { get; private set; } = null!;

        /// <summary>
        /// The secondary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        /// </summary>
        [Output("secondaryVlanId")]
        public Output<int> SecondaryVlanId { get; private set; } = null!;


        /// <summary>
        /// Create a DistributedVirtualSwitchPvlanMapping resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DistributedVirtualSwitchPvlanMapping(string name, DistributedVirtualSwitchPvlanMappingArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/distributedVirtualSwitchPvlanMapping:DistributedVirtualSwitchPvlanMapping", name, args ?? new DistributedVirtualSwitchPvlanMappingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DistributedVirtualSwitchPvlanMapping(string name, Input<string> id, DistributedVirtualSwitchPvlanMappingState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/distributedVirtualSwitchPvlanMapping:DistributedVirtualSwitchPvlanMapping", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DistributedVirtualSwitchPvlanMapping resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DistributedVirtualSwitchPvlanMapping Get(string name, Input<string> id, DistributedVirtualSwitchPvlanMappingState? state = null, CustomResourceOptions? options = null)
        {
            return new DistributedVirtualSwitchPvlanMapping(name, id, state, options);
        }
    }

    public sealed class DistributedVirtualSwitchPvlanMappingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the distributed virtual switch to attach this mapping to.
        /// </summary>
        [Input("distributedVirtualSwitchId", required: true)]
        public Input<string> DistributedVirtualSwitchId { get; set; } = null!;

        /// <summary>
        /// The primary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        /// </summary>
        [Input("primaryVlanId", required: true)]
        public Input<int> PrimaryVlanId { get; set; } = null!;

        /// <summary>
        /// The private VLAN type. Valid values are promiscuous, community and isolated.
        /// </summary>
        [Input("pvlanType", required: true)]
        public Input<string> PvlanType { get; set; } = null!;

        /// <summary>
        /// The secondary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        /// </summary>
        [Input("secondaryVlanId", required: true)]
        public Input<int> SecondaryVlanId { get; set; } = null!;

        public DistributedVirtualSwitchPvlanMappingArgs()
        {
        }
        public static new DistributedVirtualSwitchPvlanMappingArgs Empty => new DistributedVirtualSwitchPvlanMappingArgs();
    }

    public sealed class DistributedVirtualSwitchPvlanMappingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the distributed virtual switch to attach this mapping to.
        /// </summary>
        [Input("distributedVirtualSwitchId")]
        public Input<string>? DistributedVirtualSwitchId { get; set; }

        /// <summary>
        /// The primary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        /// </summary>
        [Input("primaryVlanId")]
        public Input<int>? PrimaryVlanId { get; set; }

        /// <summary>
        /// The private VLAN type. Valid values are promiscuous, community and isolated.
        /// </summary>
        [Input("pvlanType")]
        public Input<string>? PvlanType { get; set; }

        /// <summary>
        /// The secondary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        /// </summary>
        [Input("secondaryVlanId")]
        public Input<int>? SecondaryVlanId { get; set; }

        public DistributedVirtualSwitchPvlanMappingState()
        {
        }
        public static new DistributedVirtualSwitchPvlanMappingState Empty => new DistributedVirtualSwitchPvlanMappingState();
    }
}

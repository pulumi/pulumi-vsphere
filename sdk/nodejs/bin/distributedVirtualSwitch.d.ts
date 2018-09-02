import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_distributed_virtual_switch` resource can be used to manage VMware
 * Distributed Virtual Switches.
 *
 * An essential component of a distributed, scalable VMware datacenter, the
 * vSphere Distributed Virtual Switch (DVS) provides centralized management and
 * monitoring of the networking configuration of all the hosts that are associated
 * with the switch. In addition to adding port groups (see the
 * [`vsphere_distributed_port_group`][distributed-port-group] resource) that can
 * be used as networks for virtual machines, a DVS can be configured to perform
 * advanced high availability, traffic shaping, network monitoring, and more.
 *
 * For an overview on vSphere networking concepts, see [this
 * page][ref-vsphere-net-concepts]. For more information on vSphere DVS, see [this
 * page][ref-vsphere-dvs].
 *
 * [distributed-port-group]: /docs/providers/vsphere/r/distributed_port_group.html
 * [ref-vsphere-net-concepts]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.networking.doc/GUID-2B11DBB8-CB3C-4AFF-8885-EFEA0FC562F4.html
 * [ref-vsphere-dvs]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.networking.doc/GUID-375B45C7-684C-4C51-BA3C-70E48DFABF04.html
 *
 * ~> **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 */
export declare class DistributedVirtualSwitch extends pulumi.CustomResource {
    /**
     * Get an existing DistributedVirtualSwitch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DistributedVirtualSwitchState): DistributedVirtualSwitch;
    /**
     * A list of active uplinks to be used in load
     * balancing. These uplinks need to match the definitions in the
     * `uplinks` DVS argument. See
     * here for more details.
     */
    readonly activeUplinks: pulumi.Output<string[]>;
    /**
     * Controls whether or not a virtual
     * network adapter is allowed to send network traffic with a different MAC
     * address than that of its own.
     */
    readonly allowForgedTransmits: pulumi.Output<boolean>;
    /**
     * Controls whether or not the Media Access
     * Control (MAC) address can be changed.
     */
    readonly allowMacChanges: pulumi.Output<boolean>;
    /**
     * Enable promiscuous mode on the network. This
     * flag indicates whether or not all traffic is seen on a given port.
     */
    readonly allowPromiscuous: pulumi.Output<boolean>;
    /**
     * Shuts down all ports in the port groups that
     * this policy applies to, effectively blocking all network access to connected
     * virtual devices.
     */
    readonly blockAllPorts: pulumi.Output<boolean>;
    /**
     * Enables beacon probing as an additional measure
     * to detect NIC failure.
     */
    readonly checkBeacon: pulumi.Output<boolean>;
    /**
     * The version string of the configuration that this spec is trying to change.
     */
    readonly configVersion: pulumi.Output<string>;
    /**
     * The detailed contact information for the person
     * who is responsible for the DVS.
     */
    readonly contactDetail: pulumi.Output<string | undefined>;
    /**
     * The name of the person who is responsible for the
     * DVS.
     */
    readonly contactName: pulumi.Output<string | undefined>;
    /**
     * Map of custom attribute ids to attribute
     * value strings to set for virtual switch. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     */
    readonly customAttributes: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * The ID of the datacenter where the distributed
     * virtual switch will be created. Forces a new resource if changed.
     */
    readonly datacenterId: pulumi.Output<string>;
    /**
     * A detailed description for the DVS.
     */
    readonly description: pulumi.Output<string | undefined>;
    /**
     * Allow VMDirectPath Gen2 for the ports
     * for which this policy applies to.
     */
    readonly directpathGen2Allowed: pulumi.Output<boolean>;
    /**
     * The average bandwidth in bits
     * per second if egress traffic shaping is enabled on the port.
     */
    readonly egressShapingAverageBandwidth: pulumi.Output<number>;
    /**
     * The maximum burst size allowed in
     * bytes if egress traffic shaping is enabled on the port.
     */
    readonly egressShapingBurstSize: pulumi.Output<number>;
    /**
     * `true` if the traffic shaper is enabled
     * on the port for egress traffic.
     */
    readonly egressShapingEnabled: pulumi.Output<boolean>;
    /**
     * The peak bandwidth during bursts
     * in bits per second if egress traffic shaping is enabled on the port.
     */
    readonly egressShapingPeakBandwidth: pulumi.Output<number>;
    /**
     * If `true`, the teaming policy will re-activate failed
     * uplinks higher in precedence when they come back up.
     */
    readonly failback: pulumi.Output<boolean>;
    /**
     * The maximum allowed usage for the faultTolerance traffic class, in Mbits/sec.
     */
    readonly faulttoleranceMaximumMbit: pulumi.Output<number>;
    /**
     * The amount of guaranteed bandwidth for the faultTolerance traffic class, in Mbits/sec.
     */
    readonly faulttoleranceReservationMbit: pulumi.Output<number>;
    /**
     * The amount of shares to allocate to the faultTolerance traffic class for a custom share level.
     */
    readonly faulttoleranceShareCount: pulumi.Output<number>;
    /**
     * The allocation level for the faultTolerance traffic class. Can be one of high, low, normal, or custom.
     */
    readonly faulttoleranceShareLevel: pulumi.Output<string>;
    /**
     * The folder to create the DVS in. Forces a new resource
     * if changed.
     */
    readonly folder: pulumi.Output<string | undefined>;
    /**
     * The maximum allowed usage for the hbr traffic class, in Mbits/sec.
     */
    readonly hbrMaximumMbit: pulumi.Output<number>;
    /**
     * The amount of guaranteed bandwidth for the hbr traffic class, in Mbits/sec.
     */
    readonly hbrReservationMbit: pulumi.Output<number>;
    /**
     * The amount of shares to allocate to the hbr traffic class for a custom share level.
     */
    readonly hbrShareCount: pulumi.Output<number>;
    /**
     * The allocation level for the hbr traffic class. Can be one of high, low, normal, or custom.
     */
    readonly hbrShareLevel: pulumi.Output<string>;
    /**
     * Use the `host` block to declare a host specification. The
     * options are:
     */
    readonly hosts: pulumi.Output<{
        devices: string[];
        hostSystemId: string;
    }[] | undefined>;
    /**
     * The average bandwidth in
     * bits per second if ingress traffic shaping is enabled on the port.
     */
    readonly ingressShapingAverageBandwidth: pulumi.Output<number>;
    /**
     * The maximum burst size allowed in
     * bytes if ingress traffic shaping is enabled on the port.
     */
    readonly ingressShapingBurstSize: pulumi.Output<number>;
    /**
     * `true` if the traffic shaper is
     * enabled on the port for ingress traffic.
     */
    readonly ingressShapingEnabled: pulumi.Output<boolean>;
    /**
     * The peak bandwidth during
     * bursts in bits per second if ingress traffic shaping is enabled on the port.
     */
    readonly ingressShapingPeakBandwidth: pulumi.Output<number>;
    /**
     * An IPv4 address to identify the switch. This is
     * mostly useful when used with the Netflow arguments found
     * below.
     */
    readonly ipv4Address: pulumi.Output<string | undefined>;
    /**
     * The maximum allowed usage for the iSCSI traffic class, in Mbits/sec.
     */
    readonly iscsiMaximumMbit: pulumi.Output<number>;
    /**
     * The amount of guaranteed bandwidth for the iSCSI traffic class, in Mbits/sec.
     */
    readonly iscsiReservationMbit: pulumi.Output<number>;
    /**
     * The amount of shares to allocate to the iSCSI traffic class for a custom share level.
     */
    readonly iscsiShareCount: pulumi.Output<number>;
    /**
     * The allocation level for the iSCSI traffic class. Can be one of high, low, normal, or custom.
     */
    readonly iscsiShareLevel: pulumi.Output<string>;
    /**
     * The Link Aggregation Control Protocol group
     * version to use with the switch. Possible values are `singleLag` and
     * `multipleLag`.
     */
    readonly lacpApiVersion: pulumi.Output<string>;
    /**
     * Enables LACP for the ports that this policy
     * applies to.
     */
    readonly lacpEnabled: pulumi.Output<boolean>;
    /**
     * The LACP mode. Can be one of `active` or `passive`.
     */
    readonly lacpMode: pulumi.Output<string>;
    /**
     * Whether to `advertise` or `listen`
     * for link discovery traffic.
     */
    readonly linkDiscoveryOperation: pulumi.Output<string | undefined>;
    /**
     * The discovery protocol type. Valid
     * types are `cdp` and `lldp`.
     */
    readonly linkDiscoveryProtocol: pulumi.Output<string | undefined>;
    /**
     * The maximum allowed usage for the management traffic class, in Mbits/sec.
     */
    readonly managementMaximumMbit: pulumi.Output<number>;
    /**
     * The amount of guaranteed bandwidth for the management traffic class, in Mbits/sec.
     */
    readonly managementReservationMbit: pulumi.Output<number>;
    /**
     * The amount of shares to allocate to the management traffic class for a custom share level.
     */
    readonly managementShareCount: pulumi.Output<number>;
    /**
     * The allocation level for the management traffic class. Can be one of high, low, normal, or custom.
     */
    readonly managementShareLevel: pulumi.Output<string>;
    /**
     * The maximum transmission unit (MTU) for the virtual
     * switch.
     */
    readonly maxMtu: pulumi.Output<number>;
    /**
     * The multicast filtering mode to use
     * with the switch. Can be one of `legacyFiltering` or `snooping`.
     */
    readonly multicastFilteringMode: pulumi.Output<string>;
    /**
     * The name of the distributed virtual switch.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The number of seconds after which
     * active flows are forced to be exported to the collector. Allowed range is
     * `60` to `3600`. Default: `60`.
     */
    readonly netflowActiveFlowTimeout: pulumi.Output<number | undefined>;
    /**
     * IP address for the Netflow
     * collector, using IPv4 or IPv6. IPv6 is supported in vSphere Distributed
     * Switch Version 6.0 or later. Must be set before Netflow can be enabled.
     */
    readonly netflowCollectorIpAddress: pulumi.Output<string | undefined>;
    /**
     * Port for the Netflow collector. This
     * must be set before Netflow can be enabled.
     */
    readonly netflowCollectorPort: pulumi.Output<number | undefined>;
    /**
     * Enables Netflow on all ports that this policy
     * applies to.
     */
    readonly netflowEnabled: pulumi.Output<boolean>;
    /**
     * The number of seconds after which
     * idle flows are forced to be exported to the collector. Allowed range is `10`
     * to `600`. Default: `15`.
     */
    readonly netflowIdleFlowTimeout: pulumi.Output<number | undefined>;
    /**
     * Whether to limit analysis to
     * traffic that has both source and destination served by the same host.
     * Default: `false`.
     */
    readonly netflowInternalFlowsOnly: pulumi.Output<boolean | undefined>;
    /**
     * The observation domain ID for
     * the Netflow collector.
     */
    readonly netflowObservationDomainId: pulumi.Output<number | undefined>;
    /**
     * The ratio of total number of packets to
     * the number of packets analyzed. The default is `0`, which indicates that the
     * switch should analyze all packets. The maximum value is `1000`, which
     * indicates an analysis rate of 0.001%.
     */
    readonly netflowSamplingRate: pulumi.Output<number | undefined>;
    /**
     * Set to `true` to enable
     * network I/O control. Default: `false`.
     */
    readonly networkResourceControlEnabled: pulumi.Output<boolean | undefined>;
    /**
     * The version of network I/O
     * control to use. Can be one of `version2` or `version3`. Default: `version2`.
     */
    readonly networkResourceControlVersion: pulumi.Output<string>;
    /**
     * The maximum allowed usage for the nfs traffic class, in Mbits/sec.
     */
    readonly nfsMaximumMbit: pulumi.Output<number>;
    /**
     * The amount of guaranteed bandwidth for the nfs traffic class, in Mbits/sec.
     */
    readonly nfsReservationMbit: pulumi.Output<number>;
    /**
     * The amount of shares to allocate to the nfs traffic class for a custom share level.
     */
    readonly nfsShareCount: pulumi.Output<number>;
    /**
     * The allocation level for the nfs traffic class. Can be one of high, low, normal, or custom.
     */
    readonly nfsShareLevel: pulumi.Output<string>;
    /**
     * If `true`, the teaming policy will notify the
     * broadcast network of an uplink failover, triggering cache updates.
     */
    readonly notifySwitches: pulumi.Output<boolean>;
    /**
     * Used to define a secondary VLAN
     * ID when using private VLANs.
     */
    readonly portPrivateSecondaryVlanId: pulumi.Output<number>;
    /**
     * A list of standby uplinks to be used in
     * failover. These uplinks need to match the definitions in the
     * `uplinks` DVS argument. See
     * here for more details.
     */
    readonly standbyUplinks: pulumi.Output<string[]>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags: pulumi.Output<string[] | undefined>;
    /**
     * The uplink teaming policy. Can be one of
     * `loadbalance_ip`, `loadbalance_srcmac`, `loadbalance_srcid`, or
     * `failover_explicit`.
     */
    readonly teamingPolicy: pulumi.Output<string>;
    /**
     * Forward all traffic transmitted by ports for which
     * this policy applies to its DVS uplinks.
     */
    readonly txUplink: pulumi.Output<boolean>;
    /**
     * A list of strings that uniquely identifies the names
     * of the uplinks on the DVS across hosts. The number of items in this list
     * controls the number of uplinks that exist on the DVS, in addition to the
     * names.  See here for an example on how to
     * use this option.
     */
    readonly uplinks: pulumi.Output<string[]>;
    /**
     * The maximum allowed usage for the vdp traffic class, in Mbits/sec.
     */
    readonly vdpMaximumMbit: pulumi.Output<number>;
    /**
     * The amount of guaranteed bandwidth for the vdp traffic class, in Mbits/sec.
     */
    readonly vdpReservationMbit: pulumi.Output<number>;
    /**
     * The amount of shares to allocate to the vdp traffic class for a custom share level.
     */
    readonly vdpShareCount: pulumi.Output<number>;
    /**
     * The allocation level for the vdp traffic class. Can be one of high, low, normal, or custom.
     */
    readonly vdpShareLevel: pulumi.Output<string>;
    /**
     * - The version of the DVS to create. The default is to
     * create the DVS at the latest version supported by the version of vSphere
     * being used. A DVS can be upgraded to another version, but cannot be
     * downgraded.
     */
    readonly version: pulumi.Output<string>;
    /**
     * The maximum allowed usage for the virtualMachine traffic class, in Mbits/sec.
     */
    readonly virtualmachineMaximumMbit: pulumi.Output<number>;
    /**
     * The amount of guaranteed bandwidth for the virtualMachine traffic class, in Mbits/sec.
     */
    readonly virtualmachineReservationMbit: pulumi.Output<number>;
    /**
     * The amount of shares to allocate to the virtualMachine traffic class for a custom share level.
     */
    readonly virtualmachineShareCount: pulumi.Output<number>;
    /**
     * The allocation level for the virtualMachine traffic class. Can be one of high, low, normal, or custom.
     */
    readonly virtualmachineShareLevel: pulumi.Output<string>;
    /**
     * The VLAN ID for single VLAN mode. 0 denotes no VLAN.
     */
    readonly vlanId: pulumi.Output<number>;
    /**
     * Used to denote VLAN trunking. Use the `min_vlan`
     * and `max_vlan` sub-arguments to define the tagged VLAN range. Multiple
     * `vlan_range` definitions are allowed, but they must not overlap. Example
     * below:
     */
    readonly vlanRanges: pulumi.Output<{
        maxVlan: number;
        minVlan: number;
    }[]>;
    /**
     * The maximum allowed usage for the vmotion traffic class, in Mbits/sec.
     */
    readonly vmotionMaximumMbit: pulumi.Output<number>;
    /**
     * The amount of guaranteed bandwidth for the vmotion traffic class, in Mbits/sec.
     */
    readonly vmotionReservationMbit: pulumi.Output<number>;
    /**
     * The amount of shares to allocate to the vmotion traffic class for a custom share level.
     */
    readonly vmotionShareCount: pulumi.Output<number>;
    /**
     * The allocation level for the vmotion traffic class. Can be one of high, low, normal, or custom.
     */
    readonly vmotionShareLevel: pulumi.Output<string>;
    /**
     * The maximum allowed usage for the vsan traffic class, in Mbits/sec.
     */
    readonly vsanMaximumMbit: pulumi.Output<number>;
    /**
     * The amount of guaranteed bandwidth for the vsan traffic class, in Mbits/sec.
     */
    readonly vsanReservationMbit: pulumi.Output<number>;
    /**
     * The amount of shares to allocate to the vsan traffic class for a custom share level.
     */
    readonly vsanShareCount: pulumi.Output<number>;
    /**
     * The allocation level for the vsan traffic class. Can be one of high, low, normal, or custom.
     */
    readonly vsanShareLevel: pulumi.Output<string>;
    /**
     * Create a DistributedVirtualSwitch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DistributedVirtualSwitchArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering DistributedVirtualSwitch resources.
 */
export interface DistributedVirtualSwitchState {
    /**
     * A list of active uplinks to be used in load
     * balancing. These uplinks need to match the definitions in the
     * `uplinks` DVS argument. See
     * here for more details.
     */
    readonly activeUplinks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Controls whether or not a virtual
     * network adapter is allowed to send network traffic with a different MAC
     * address than that of its own.
     */
    readonly allowForgedTransmits?: pulumi.Input<boolean>;
    /**
     * Controls whether or not the Media Access
     * Control (MAC) address can be changed.
     */
    readonly allowMacChanges?: pulumi.Input<boolean>;
    /**
     * Enable promiscuous mode on the network. This
     * flag indicates whether or not all traffic is seen on a given port.
     */
    readonly allowPromiscuous?: pulumi.Input<boolean>;
    /**
     * Shuts down all ports in the port groups that
     * this policy applies to, effectively blocking all network access to connected
     * virtual devices.
     */
    readonly blockAllPorts?: pulumi.Input<boolean>;
    /**
     * Enables beacon probing as an additional measure
     * to detect NIC failure.
     */
    readonly checkBeacon?: pulumi.Input<boolean>;
    /**
     * The version string of the configuration that this spec is trying to change.
     */
    readonly configVersion?: pulumi.Input<string>;
    /**
     * The detailed contact information for the person
     * who is responsible for the DVS.
     */
    readonly contactDetail?: pulumi.Input<string>;
    /**
     * The name of the person who is responsible for the
     * DVS.
     */
    readonly contactName?: pulumi.Input<string>;
    /**
     * Map of custom attribute ids to attribute
     * value strings to set for virtual switch. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     */
    readonly customAttributes?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The ID of the datacenter where the distributed
     * virtual switch will be created. Forces a new resource if changed.
     */
    readonly datacenterId?: pulumi.Input<string>;
    /**
     * A detailed description for the DVS.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Allow VMDirectPath Gen2 for the ports
     * for which this policy applies to.
     */
    readonly directpathGen2Allowed?: pulumi.Input<boolean>;
    /**
     * The average bandwidth in bits
     * per second if egress traffic shaping is enabled on the port.
     */
    readonly egressShapingAverageBandwidth?: pulumi.Input<number>;
    /**
     * The maximum burst size allowed in
     * bytes if egress traffic shaping is enabled on the port.
     */
    readonly egressShapingBurstSize?: pulumi.Input<number>;
    /**
     * `true` if the traffic shaper is enabled
     * on the port for egress traffic.
     */
    readonly egressShapingEnabled?: pulumi.Input<boolean>;
    /**
     * The peak bandwidth during bursts
     * in bits per second if egress traffic shaping is enabled on the port.
     */
    readonly egressShapingPeakBandwidth?: pulumi.Input<number>;
    /**
     * If `true`, the teaming policy will re-activate failed
     * uplinks higher in precedence when they come back up.
     */
    readonly failback?: pulumi.Input<boolean>;
    /**
     * The maximum allowed usage for the faultTolerance traffic class, in Mbits/sec.
     */
    readonly faulttoleranceMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the faultTolerance traffic class, in Mbits/sec.
     */
    readonly faulttoleranceReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the faultTolerance traffic class for a custom share level.
     */
    readonly faulttoleranceShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the faultTolerance traffic class. Can be one of high, low, normal, or custom.
     */
    readonly faulttoleranceShareLevel?: pulumi.Input<string>;
    /**
     * The folder to create the DVS in. Forces a new resource
     * if changed.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * The maximum allowed usage for the hbr traffic class, in Mbits/sec.
     */
    readonly hbrMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the hbr traffic class, in Mbits/sec.
     */
    readonly hbrReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the hbr traffic class for a custom share level.
     */
    readonly hbrShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the hbr traffic class. Can be one of high, low, normal, or custom.
     */
    readonly hbrShareLevel?: pulumi.Input<string>;
    /**
     * Use the `host` block to declare a host specification. The
     * options are:
     */
    readonly hosts?: pulumi.Input<pulumi.Input<{
        devices: pulumi.Input<pulumi.Input<string>[]>;
        hostSystemId: pulumi.Input<string>;
    }>[]>;
    /**
     * The average bandwidth in
     * bits per second if ingress traffic shaping is enabled on the port.
     */
    readonly ingressShapingAverageBandwidth?: pulumi.Input<number>;
    /**
     * The maximum burst size allowed in
     * bytes if ingress traffic shaping is enabled on the port.
     */
    readonly ingressShapingBurstSize?: pulumi.Input<number>;
    /**
     * `true` if the traffic shaper is
     * enabled on the port for ingress traffic.
     */
    readonly ingressShapingEnabled?: pulumi.Input<boolean>;
    /**
     * The peak bandwidth during
     * bursts in bits per second if ingress traffic shaping is enabled on the port.
     */
    readonly ingressShapingPeakBandwidth?: pulumi.Input<number>;
    /**
     * An IPv4 address to identify the switch. This is
     * mostly useful when used with the Netflow arguments found
     * below.
     */
    readonly ipv4Address?: pulumi.Input<string>;
    /**
     * The maximum allowed usage for the iSCSI traffic class, in Mbits/sec.
     */
    readonly iscsiMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the iSCSI traffic class, in Mbits/sec.
     */
    readonly iscsiReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the iSCSI traffic class for a custom share level.
     */
    readonly iscsiShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the iSCSI traffic class. Can be one of high, low, normal, or custom.
     */
    readonly iscsiShareLevel?: pulumi.Input<string>;
    /**
     * The Link Aggregation Control Protocol group
     * version to use with the switch. Possible values are `singleLag` and
     * `multipleLag`.
     */
    readonly lacpApiVersion?: pulumi.Input<string>;
    /**
     * Enables LACP for the ports that this policy
     * applies to.
     */
    readonly lacpEnabled?: pulumi.Input<boolean>;
    /**
     * The LACP mode. Can be one of `active` or `passive`.
     */
    readonly lacpMode?: pulumi.Input<string>;
    /**
     * Whether to `advertise` or `listen`
     * for link discovery traffic.
     */
    readonly linkDiscoveryOperation?: pulumi.Input<string>;
    /**
     * The discovery protocol type. Valid
     * types are `cdp` and `lldp`.
     */
    readonly linkDiscoveryProtocol?: pulumi.Input<string>;
    /**
     * The maximum allowed usage for the management traffic class, in Mbits/sec.
     */
    readonly managementMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the management traffic class, in Mbits/sec.
     */
    readonly managementReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the management traffic class for a custom share level.
     */
    readonly managementShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the management traffic class. Can be one of high, low, normal, or custom.
     */
    readonly managementShareLevel?: pulumi.Input<string>;
    /**
     * The maximum transmission unit (MTU) for the virtual
     * switch.
     */
    readonly maxMtu?: pulumi.Input<number>;
    /**
     * The multicast filtering mode to use
     * with the switch. Can be one of `legacyFiltering` or `snooping`.
     */
    readonly multicastFilteringMode?: pulumi.Input<string>;
    /**
     * The name of the distributed virtual switch.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The number of seconds after which
     * active flows are forced to be exported to the collector. Allowed range is
     * `60` to `3600`. Default: `60`.
     */
    readonly netflowActiveFlowTimeout?: pulumi.Input<number>;
    /**
     * IP address for the Netflow
     * collector, using IPv4 or IPv6. IPv6 is supported in vSphere Distributed
     * Switch Version 6.0 or later. Must be set before Netflow can be enabled.
     */
    readonly netflowCollectorIpAddress?: pulumi.Input<string>;
    /**
     * Port for the Netflow collector. This
     * must be set before Netflow can be enabled.
     */
    readonly netflowCollectorPort?: pulumi.Input<number>;
    /**
     * Enables Netflow on all ports that this policy
     * applies to.
     */
    readonly netflowEnabled?: pulumi.Input<boolean>;
    /**
     * The number of seconds after which
     * idle flows are forced to be exported to the collector. Allowed range is `10`
     * to `600`. Default: `15`.
     */
    readonly netflowIdleFlowTimeout?: pulumi.Input<number>;
    /**
     * Whether to limit analysis to
     * traffic that has both source and destination served by the same host.
     * Default: `false`.
     */
    readonly netflowInternalFlowsOnly?: pulumi.Input<boolean>;
    /**
     * The observation domain ID for
     * the Netflow collector.
     */
    readonly netflowObservationDomainId?: pulumi.Input<number>;
    /**
     * The ratio of total number of packets to
     * the number of packets analyzed. The default is `0`, which indicates that the
     * switch should analyze all packets. The maximum value is `1000`, which
     * indicates an analysis rate of 0.001%.
     */
    readonly netflowSamplingRate?: pulumi.Input<number>;
    /**
     * Set to `true` to enable
     * network I/O control. Default: `false`.
     */
    readonly networkResourceControlEnabled?: pulumi.Input<boolean>;
    /**
     * The version of network I/O
     * control to use. Can be one of `version2` or `version3`. Default: `version2`.
     */
    readonly networkResourceControlVersion?: pulumi.Input<string>;
    /**
     * The maximum allowed usage for the nfs traffic class, in Mbits/sec.
     */
    readonly nfsMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the nfs traffic class, in Mbits/sec.
     */
    readonly nfsReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the nfs traffic class for a custom share level.
     */
    readonly nfsShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the nfs traffic class. Can be one of high, low, normal, or custom.
     */
    readonly nfsShareLevel?: pulumi.Input<string>;
    /**
     * If `true`, the teaming policy will notify the
     * broadcast network of an uplink failover, triggering cache updates.
     */
    readonly notifySwitches?: pulumi.Input<boolean>;
    /**
     * Used to define a secondary VLAN
     * ID when using private VLANs.
     */
    readonly portPrivateSecondaryVlanId?: pulumi.Input<number>;
    /**
     * A list of standby uplinks to be used in
     * failover. These uplinks need to match the definitions in the
     * `uplinks` DVS argument. See
     * here for more details.
     */
    readonly standbyUplinks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The uplink teaming policy. Can be one of
     * `loadbalance_ip`, `loadbalance_srcmac`, `loadbalance_srcid`, or
     * `failover_explicit`.
     */
    readonly teamingPolicy?: pulumi.Input<string>;
    /**
     * Forward all traffic transmitted by ports for which
     * this policy applies to its DVS uplinks.
     */
    readonly txUplink?: pulumi.Input<boolean>;
    /**
     * A list of strings that uniquely identifies the names
     * of the uplinks on the DVS across hosts. The number of items in this list
     * controls the number of uplinks that exist on the DVS, in addition to the
     * names.  See here for an example on how to
     * use this option.
     */
    readonly uplinks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum allowed usage for the vdp traffic class, in Mbits/sec.
     */
    readonly vdpMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the vdp traffic class, in Mbits/sec.
     */
    readonly vdpReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the vdp traffic class for a custom share level.
     */
    readonly vdpShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the vdp traffic class. Can be one of high, low, normal, or custom.
     */
    readonly vdpShareLevel?: pulumi.Input<string>;
    /**
     * - The version of the DVS to create. The default is to
     * create the DVS at the latest version supported by the version of vSphere
     * being used. A DVS can be upgraded to another version, but cannot be
     * downgraded.
     */
    readonly version?: pulumi.Input<string>;
    /**
     * The maximum allowed usage for the virtualMachine traffic class, in Mbits/sec.
     */
    readonly virtualmachineMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the virtualMachine traffic class, in Mbits/sec.
     */
    readonly virtualmachineReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the virtualMachine traffic class for a custom share level.
     */
    readonly virtualmachineShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the virtualMachine traffic class. Can be one of high, low, normal, or custom.
     */
    readonly virtualmachineShareLevel?: pulumi.Input<string>;
    /**
     * The VLAN ID for single VLAN mode. 0 denotes no VLAN.
     */
    readonly vlanId?: pulumi.Input<number>;
    /**
     * Used to denote VLAN trunking. Use the `min_vlan`
     * and `max_vlan` sub-arguments to define the tagged VLAN range. Multiple
     * `vlan_range` definitions are allowed, but they must not overlap. Example
     * below:
     */
    readonly vlanRanges?: pulumi.Input<pulumi.Input<{
        maxVlan: pulumi.Input<number>;
        minVlan: pulumi.Input<number>;
    }>[]>;
    /**
     * The maximum allowed usage for the vmotion traffic class, in Mbits/sec.
     */
    readonly vmotionMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the vmotion traffic class, in Mbits/sec.
     */
    readonly vmotionReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the vmotion traffic class for a custom share level.
     */
    readonly vmotionShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the vmotion traffic class. Can be one of high, low, normal, or custom.
     */
    readonly vmotionShareLevel?: pulumi.Input<string>;
    /**
     * The maximum allowed usage for the vsan traffic class, in Mbits/sec.
     */
    readonly vsanMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the vsan traffic class, in Mbits/sec.
     */
    readonly vsanReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the vsan traffic class for a custom share level.
     */
    readonly vsanShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the vsan traffic class. Can be one of high, low, normal, or custom.
     */
    readonly vsanShareLevel?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a DistributedVirtualSwitch resource.
 */
export interface DistributedVirtualSwitchArgs {
    /**
     * A list of active uplinks to be used in load
     * balancing. These uplinks need to match the definitions in the
     * `uplinks` DVS argument. See
     * here for more details.
     */
    readonly activeUplinks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Controls whether or not a virtual
     * network adapter is allowed to send network traffic with a different MAC
     * address than that of its own.
     */
    readonly allowForgedTransmits?: pulumi.Input<boolean>;
    /**
     * Controls whether or not the Media Access
     * Control (MAC) address can be changed.
     */
    readonly allowMacChanges?: pulumi.Input<boolean>;
    /**
     * Enable promiscuous mode on the network. This
     * flag indicates whether or not all traffic is seen on a given port.
     */
    readonly allowPromiscuous?: pulumi.Input<boolean>;
    /**
     * Shuts down all ports in the port groups that
     * this policy applies to, effectively blocking all network access to connected
     * virtual devices.
     */
    readonly blockAllPorts?: pulumi.Input<boolean>;
    /**
     * Enables beacon probing as an additional measure
     * to detect NIC failure.
     */
    readonly checkBeacon?: pulumi.Input<boolean>;
    /**
     * The detailed contact information for the person
     * who is responsible for the DVS.
     */
    readonly contactDetail?: pulumi.Input<string>;
    /**
     * The name of the person who is responsible for the
     * DVS.
     */
    readonly contactName?: pulumi.Input<string>;
    /**
     * Map of custom attribute ids to attribute
     * value strings to set for virtual switch. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     */
    readonly customAttributes?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The ID of the datacenter where the distributed
     * virtual switch will be created. Forces a new resource if changed.
     */
    readonly datacenterId: pulumi.Input<string>;
    /**
     * A detailed description for the DVS.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Allow VMDirectPath Gen2 for the ports
     * for which this policy applies to.
     */
    readonly directpathGen2Allowed?: pulumi.Input<boolean>;
    /**
     * The average bandwidth in bits
     * per second if egress traffic shaping is enabled on the port.
     */
    readonly egressShapingAverageBandwidth?: pulumi.Input<number>;
    /**
     * The maximum burst size allowed in
     * bytes if egress traffic shaping is enabled on the port.
     */
    readonly egressShapingBurstSize?: pulumi.Input<number>;
    /**
     * `true` if the traffic shaper is enabled
     * on the port for egress traffic.
     */
    readonly egressShapingEnabled?: pulumi.Input<boolean>;
    /**
     * The peak bandwidth during bursts
     * in bits per second if egress traffic shaping is enabled on the port.
     */
    readonly egressShapingPeakBandwidth?: pulumi.Input<number>;
    /**
     * If `true`, the teaming policy will re-activate failed
     * uplinks higher in precedence when they come back up.
     */
    readonly failback?: pulumi.Input<boolean>;
    /**
     * The maximum allowed usage for the faultTolerance traffic class, in Mbits/sec.
     */
    readonly faulttoleranceMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the faultTolerance traffic class, in Mbits/sec.
     */
    readonly faulttoleranceReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the faultTolerance traffic class for a custom share level.
     */
    readonly faulttoleranceShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the faultTolerance traffic class. Can be one of high, low, normal, or custom.
     */
    readonly faulttoleranceShareLevel?: pulumi.Input<string>;
    /**
     * The folder to create the DVS in. Forces a new resource
     * if changed.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * The maximum allowed usage for the hbr traffic class, in Mbits/sec.
     */
    readonly hbrMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the hbr traffic class, in Mbits/sec.
     */
    readonly hbrReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the hbr traffic class for a custom share level.
     */
    readonly hbrShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the hbr traffic class. Can be one of high, low, normal, or custom.
     */
    readonly hbrShareLevel?: pulumi.Input<string>;
    /**
     * Use the `host` block to declare a host specification. The
     * options are:
     */
    readonly hosts?: pulumi.Input<pulumi.Input<{
        devices: pulumi.Input<pulumi.Input<string>[]>;
        hostSystemId: pulumi.Input<string>;
    }>[]>;
    /**
     * The average bandwidth in
     * bits per second if ingress traffic shaping is enabled on the port.
     */
    readonly ingressShapingAverageBandwidth?: pulumi.Input<number>;
    /**
     * The maximum burst size allowed in
     * bytes if ingress traffic shaping is enabled on the port.
     */
    readonly ingressShapingBurstSize?: pulumi.Input<number>;
    /**
     * `true` if the traffic shaper is
     * enabled on the port for ingress traffic.
     */
    readonly ingressShapingEnabled?: pulumi.Input<boolean>;
    /**
     * The peak bandwidth during
     * bursts in bits per second if ingress traffic shaping is enabled on the port.
     */
    readonly ingressShapingPeakBandwidth?: pulumi.Input<number>;
    /**
     * An IPv4 address to identify the switch. This is
     * mostly useful when used with the Netflow arguments found
     * below.
     */
    readonly ipv4Address?: pulumi.Input<string>;
    /**
     * The maximum allowed usage for the iSCSI traffic class, in Mbits/sec.
     */
    readonly iscsiMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the iSCSI traffic class, in Mbits/sec.
     */
    readonly iscsiReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the iSCSI traffic class for a custom share level.
     */
    readonly iscsiShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the iSCSI traffic class. Can be one of high, low, normal, or custom.
     */
    readonly iscsiShareLevel?: pulumi.Input<string>;
    /**
     * The Link Aggregation Control Protocol group
     * version to use with the switch. Possible values are `singleLag` and
     * `multipleLag`.
     */
    readonly lacpApiVersion?: pulumi.Input<string>;
    /**
     * Enables LACP for the ports that this policy
     * applies to.
     */
    readonly lacpEnabled?: pulumi.Input<boolean>;
    /**
     * The LACP mode. Can be one of `active` or `passive`.
     */
    readonly lacpMode?: pulumi.Input<string>;
    /**
     * Whether to `advertise` or `listen`
     * for link discovery traffic.
     */
    readonly linkDiscoveryOperation?: pulumi.Input<string>;
    /**
     * The discovery protocol type. Valid
     * types are `cdp` and `lldp`.
     */
    readonly linkDiscoveryProtocol?: pulumi.Input<string>;
    /**
     * The maximum allowed usage for the management traffic class, in Mbits/sec.
     */
    readonly managementMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the management traffic class, in Mbits/sec.
     */
    readonly managementReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the management traffic class for a custom share level.
     */
    readonly managementShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the management traffic class. Can be one of high, low, normal, or custom.
     */
    readonly managementShareLevel?: pulumi.Input<string>;
    /**
     * The maximum transmission unit (MTU) for the virtual
     * switch.
     */
    readonly maxMtu?: pulumi.Input<number>;
    /**
     * The multicast filtering mode to use
     * with the switch. Can be one of `legacyFiltering` or `snooping`.
     */
    readonly multicastFilteringMode?: pulumi.Input<string>;
    /**
     * The name of the distributed virtual switch.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The number of seconds after which
     * active flows are forced to be exported to the collector. Allowed range is
     * `60` to `3600`. Default: `60`.
     */
    readonly netflowActiveFlowTimeout?: pulumi.Input<number>;
    /**
     * IP address for the Netflow
     * collector, using IPv4 or IPv6. IPv6 is supported in vSphere Distributed
     * Switch Version 6.0 or later. Must be set before Netflow can be enabled.
     */
    readonly netflowCollectorIpAddress?: pulumi.Input<string>;
    /**
     * Port for the Netflow collector. This
     * must be set before Netflow can be enabled.
     */
    readonly netflowCollectorPort?: pulumi.Input<number>;
    /**
     * Enables Netflow on all ports that this policy
     * applies to.
     */
    readonly netflowEnabled?: pulumi.Input<boolean>;
    /**
     * The number of seconds after which
     * idle flows are forced to be exported to the collector. Allowed range is `10`
     * to `600`. Default: `15`.
     */
    readonly netflowIdleFlowTimeout?: pulumi.Input<number>;
    /**
     * Whether to limit analysis to
     * traffic that has both source and destination served by the same host.
     * Default: `false`.
     */
    readonly netflowInternalFlowsOnly?: pulumi.Input<boolean>;
    /**
     * The observation domain ID for
     * the Netflow collector.
     */
    readonly netflowObservationDomainId?: pulumi.Input<number>;
    /**
     * The ratio of total number of packets to
     * the number of packets analyzed. The default is `0`, which indicates that the
     * switch should analyze all packets. The maximum value is `1000`, which
     * indicates an analysis rate of 0.001%.
     */
    readonly netflowSamplingRate?: pulumi.Input<number>;
    /**
     * Set to `true` to enable
     * network I/O control. Default: `false`.
     */
    readonly networkResourceControlEnabled?: pulumi.Input<boolean>;
    /**
     * The version of network I/O
     * control to use. Can be one of `version2` or `version3`. Default: `version2`.
     */
    readonly networkResourceControlVersion?: pulumi.Input<string>;
    /**
     * The maximum allowed usage for the nfs traffic class, in Mbits/sec.
     */
    readonly nfsMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the nfs traffic class, in Mbits/sec.
     */
    readonly nfsReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the nfs traffic class for a custom share level.
     */
    readonly nfsShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the nfs traffic class. Can be one of high, low, normal, or custom.
     */
    readonly nfsShareLevel?: pulumi.Input<string>;
    /**
     * If `true`, the teaming policy will notify the
     * broadcast network of an uplink failover, triggering cache updates.
     */
    readonly notifySwitches?: pulumi.Input<boolean>;
    /**
     * Used to define a secondary VLAN
     * ID when using private VLANs.
     */
    readonly portPrivateSecondaryVlanId?: pulumi.Input<number>;
    /**
     * A list of standby uplinks to be used in
     * failover. These uplinks need to match the definitions in the
     * `uplinks` DVS argument. See
     * here for more details.
     */
    readonly standbyUplinks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The uplink teaming policy. Can be one of
     * `loadbalance_ip`, `loadbalance_srcmac`, `loadbalance_srcid`, or
     * `failover_explicit`.
     */
    readonly teamingPolicy?: pulumi.Input<string>;
    /**
     * Forward all traffic transmitted by ports for which
     * this policy applies to its DVS uplinks.
     */
    readonly txUplink?: pulumi.Input<boolean>;
    /**
     * A list of strings that uniquely identifies the names
     * of the uplinks on the DVS across hosts. The number of items in this list
     * controls the number of uplinks that exist on the DVS, in addition to the
     * names.  See here for an example on how to
     * use this option.
     */
    readonly uplinks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum allowed usage for the vdp traffic class, in Mbits/sec.
     */
    readonly vdpMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the vdp traffic class, in Mbits/sec.
     */
    readonly vdpReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the vdp traffic class for a custom share level.
     */
    readonly vdpShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the vdp traffic class. Can be one of high, low, normal, or custom.
     */
    readonly vdpShareLevel?: pulumi.Input<string>;
    /**
     * - The version of the DVS to create. The default is to
     * create the DVS at the latest version supported by the version of vSphere
     * being used. A DVS can be upgraded to another version, but cannot be
     * downgraded.
     */
    readonly version?: pulumi.Input<string>;
    /**
     * The maximum allowed usage for the virtualMachine traffic class, in Mbits/sec.
     */
    readonly virtualmachineMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the virtualMachine traffic class, in Mbits/sec.
     */
    readonly virtualmachineReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the virtualMachine traffic class for a custom share level.
     */
    readonly virtualmachineShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the virtualMachine traffic class. Can be one of high, low, normal, or custom.
     */
    readonly virtualmachineShareLevel?: pulumi.Input<string>;
    /**
     * The VLAN ID for single VLAN mode. 0 denotes no VLAN.
     */
    readonly vlanId?: pulumi.Input<number>;
    /**
     * Used to denote VLAN trunking. Use the `min_vlan`
     * and `max_vlan` sub-arguments to define the tagged VLAN range. Multiple
     * `vlan_range` definitions are allowed, but they must not overlap. Example
     * below:
     */
    readonly vlanRanges?: pulumi.Input<pulumi.Input<{
        maxVlan: pulumi.Input<number>;
        minVlan: pulumi.Input<number>;
    }>[]>;
    /**
     * The maximum allowed usage for the vmotion traffic class, in Mbits/sec.
     */
    readonly vmotionMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the vmotion traffic class, in Mbits/sec.
     */
    readonly vmotionReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the vmotion traffic class for a custom share level.
     */
    readonly vmotionShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the vmotion traffic class. Can be one of high, low, normal, or custom.
     */
    readonly vmotionShareLevel?: pulumi.Input<string>;
    /**
     * The maximum allowed usage for the vsan traffic class, in Mbits/sec.
     */
    readonly vsanMaximumMbit?: pulumi.Input<number>;
    /**
     * The amount of guaranteed bandwidth for the vsan traffic class, in Mbits/sec.
     */
    readonly vsanReservationMbit?: pulumi.Input<number>;
    /**
     * The amount of shares to allocate to the vsan traffic class for a custom share level.
     */
    readonly vsanShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for the vsan traffic class. Can be one of high, low, normal, or custom.
     */
    readonly vsanShareLevel?: pulumi.Input<string>;
}

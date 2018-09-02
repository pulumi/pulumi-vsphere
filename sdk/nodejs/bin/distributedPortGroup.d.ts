import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_distributed_port_group` resource can be used to manage vSphere
 * distributed virtual port groups. These port groups are connected to distributed
 * virtual switches, which can be managed by the
 * [`vsphere_distributed_virtual_switch`][distributed-virtual-switch] resource.
 *
 * Distributed port groups can be used as networks for virtual machines, allowing
 * VMs to use the networking supplied by a distributed virtual switch (DVS), with
 * a set of policies that apply to that individual newtork, if desired.
 *
 * For an overview on vSphere networking concepts, see [this
 * page][ref-vsphere-net-concepts]. For more information on vSphere DVS
 * portgroups, see [this page][ref-vsphere-dvportgroup].
 *
 * [distributed-virtual-switch]: /docs/providers/vsphere/r/distributed_virtual_switch.html
 * [ref-vsphere-net-concepts]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.networking.doc/GUID-2B11DBB8-CB3C-4AFF-8885-EFEA0FC562F4.html
 * [ref-vsphere-dvportgroup]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.networking.doc/GUID-69933F6E-2442-46CF-AA17-1196CB9A0A09.html
 *
 * ~> **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 */
export declare class DistributedPortGroup extends pulumi.CustomResource {
    /**
     * Get an existing DistributedPortGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DistributedPortGroupState): DistributedPortGroup;
    /**
     * List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
     */
    readonly activeUplinks: pulumi.Output<string[]>;
    /**
     * Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address
     * than that of its own.
     */
    readonly allowForgedTransmits: pulumi.Output<boolean>;
    /**
     * Controls whether or not the Media Access Control (MAC) address can be changed.
     */
    readonly allowMacChanges: pulumi.Output<boolean>;
    /**
     * Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
     */
    readonly allowPromiscuous: pulumi.Output<boolean>;
    /**
     * Allows the port group to create additional ports
     * past the limit specified in `number_of_ports` if necessary. Default: `true`.
     */
    readonly autoExpand: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether to block all ports by default.
     */
    readonly blockAllPorts: pulumi.Output<boolean>;
    /**
     * Allow the [port shutdown
     * policy][port-shutdown-policy] to be overridden on an individual port.
     */
    readonly blockOverrideAllowed: pulumi.Output<boolean | undefined>;
    /**
     * Enable beacon probing on the ports this policy applies to.
     */
    readonly checkBeacon: pulumi.Output<boolean>;
    /**
     * Version string of the configuration that this spec is trying to change.
     */
    readonly configVersion: pulumi.Output<string>;
    /**
     * Map of custom attribute ids to attribute
     * value string to set for port group. See [here][docs-setting-custom-attributes]
     * for a reference on how to set values for custom attributes.
     */
    readonly customAttributes: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * An optional description for the port group.
     */
    readonly description: pulumi.Output<string | undefined>;
    /**
     * Allow VMDirectPath Gen2 on the ports this policy applies to.
     */
    readonly directpathGen2Allowed: pulumi.Output<boolean>;
    /**
     * The ID of the DVS to add the
     * port group to. Forces a new resource if changed.
     */
    readonly distributedVirtualSwitchUuid: pulumi.Output<string>;
    /**
     * The average egress bandwidth in bits per second if egress shaping is enabled on the port.
     */
    readonly egressShapingAverageBandwidth: pulumi.Output<number>;
    /**
     * The maximum egress burst size allowed in bytes if egress shaping is enabled on the port.
     */
    readonly egressShapingBurstSize: pulumi.Output<number>;
    /**
     * True if the traffic shaper is enabled for egress traffic on the port.
     */
    readonly egressShapingEnabled: pulumi.Output<boolean>;
    /**
     * The peak egress bandwidth during bursts in bits per second if egress traffic shaping is enabled on the port.
     */
    readonly egressShapingPeakBandwidth: pulumi.Output<number>;
    /**
     * If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
     */
    readonly failback: pulumi.Output<boolean>;
    /**
     * The average ingress bandwidth in bits per second if ingress shaping is enabled on the port.
     */
    readonly ingressShapingAverageBandwidth: pulumi.Output<number>;
    /**
     * The maximum ingress burst size allowed in bytes if ingress shaping is enabled on the port.
     */
    readonly ingressShapingBurstSize: pulumi.Output<number>;
    /**
     * True if the traffic shaper is enabled for ingress traffic on the port.
     */
    readonly ingressShapingEnabled: pulumi.Output<boolean>;
    /**
     * The peak ingress bandwidth during bursts in bits per second if ingress traffic shaping is enabled on the port.
     */
    readonly ingressShapingPeakBandwidth: pulumi.Output<number>;
    /**
     * The generated UUID of the portgroup.
     */
    readonly key: pulumi.Output<string>;
    /**
     * Whether or not to enable LACP on all uplink ports.
     */
    readonly lacpEnabled: pulumi.Output<boolean>;
    /**
     * The uplink LACP mode to use. Can be one of active or passive.
     */
    readonly lacpMode: pulumi.Output<string>;
    /**
     * Allow a port in this port group to be
     * moved to another port group while it is connected.
     */
    readonly livePortMovingAllowed: pulumi.Output<boolean | undefined>;
    /**
     * The name of the port group.
     */
    readonly name: pulumi.Output<string>;
    /**
     * Indicates whether to enable netflow on all ports.
     */
    readonly netflowEnabled: pulumi.Output<boolean>;
    /**
     * Allow the [Netflow
     * policy][netflow-policy] on this port group to be overridden on an individual
     * port.
     */
    readonly netflowOverrideAllowed: pulumi.Output<boolean | undefined>;
    /**
     * The key of a network resource pool
     * to associate with this port group. The default is `-1`, which implies no
     * association.
     */
    readonly networkResourcePoolKey: pulumi.Output<string | undefined>;
    /**
     * Allow the network
     * resource pool set on this port group to be overridden on an individual port.
     */
    readonly networkResourcePoolOverrideAllowed: pulumi.Output<boolean | undefined>;
    /**
     * If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
     */
    readonly notifySwitches: pulumi.Output<boolean>;
    /**
     * The number of ports available on this port
     * group. Cannot be decreased below the amount of used ports on the port group.
     */
    readonly numberOfPorts: pulumi.Output<number>;
    /**
     * Reset a port's settings to the
     * settings defined on this port group policy when the port disconnects.
     */
    readonly portConfigResetAtDisconnect: pulumi.Output<boolean | undefined>;
    /**
     * An optional formatting policy for naming of
     * the ports in this port group. See the `portNameFormat` attribute listed
     * [here][ext-vsphere-portname-format] for details on the format syntax.
     */
    readonly portNameFormat: pulumi.Output<string | undefined>;
    /**
     * The secondary VLAN ID for this port.
     */
    readonly portPrivateSecondaryVlanId: pulumi.Output<number>;
    /**
     * Allow the [security policy
     * settings][sec-policy-settings] defined in this port group policy to be
     * overridden on an individual port.
     */
    readonly securityPolicyOverrideAllowed: pulumi.Output<boolean | undefined>;
    /**
     * Allow the [traffic shaping
     * options][traffic-shaping-settings] on this port group policy to be overridden
     * on an individual port.
     */
    readonly shapingOverrideAllowed: pulumi.Output<boolean | undefined>;
    /**
     * List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
     */
    readonly standbyUplinks: pulumi.Output<string[]>;
    /**
     * A list of tag IDs to apply to this object.
     */
    readonly tags: pulumi.Output<string[] | undefined>;
    /**
     * The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid,
     * failover_explicit, or loadbalance_loadbased.
     */
    readonly teamingPolicy: pulumi.Output<string>;
    /**
     * Allow any traffic filters on
     * this port group to be overridden on an individual port.
     */
    readonly trafficFilterOverrideAllowed: pulumi.Output<boolean | undefined>;
    /**
     * If true, a copy of packets sent to the switch will always be forwarded to an uplink in addition to the regular
     * packet forwarded done by the switch.
     */
    readonly txUplink: pulumi.Output<boolean>;
    /**
     * The port group type. Can be one of `earlyBinding` (static
     * binding) or `ephemeral`. Default: `earlyBinding`.
     */
    readonly type: pulumi.Output<string | undefined>;
    /**
     * Allow the [uplink teaming
     * options][uplink-teaming-settings] on this port group to be overridden on an
     * individual port.
     */
    readonly uplinkTeamingOverrideAllowed: pulumi.Output<boolean | undefined>;
    /**
     * The VLAN ID for single VLAN mode. 0 denotes no VLAN.
     */
    readonly vlanId: pulumi.Output<number>;
    /**
     * Allow the [VLAN settings][vlan-settings]
     * on this port group to be overridden on an individual port.
     */
    readonly vlanOverrideAllowed: pulumi.Output<boolean | undefined>;
    /**
     * The VLAN ID for single VLAN mode. 0 denotes no VLAN.
     */
    readonly vlanRanges: pulumi.Output<{
        maxVlan: number;
        minVlan: number;
    }[]>;
    /**
     * Create a DistributedPortGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DistributedPortGroupArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering DistributedPortGroup resources.
 */
export interface DistributedPortGroupState {
    /**
     * List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
     */
    readonly activeUplinks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address
     * than that of its own.
     */
    readonly allowForgedTransmits?: pulumi.Input<boolean>;
    /**
     * Controls whether or not the Media Access Control (MAC) address can be changed.
     */
    readonly allowMacChanges?: pulumi.Input<boolean>;
    /**
     * Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
     */
    readonly allowPromiscuous?: pulumi.Input<boolean>;
    /**
     * Allows the port group to create additional ports
     * past the limit specified in `number_of_ports` if necessary. Default: `true`.
     */
    readonly autoExpand?: pulumi.Input<boolean>;
    /**
     * Indicates whether to block all ports by default.
     */
    readonly blockAllPorts?: pulumi.Input<boolean>;
    /**
     * Allow the [port shutdown
     * policy][port-shutdown-policy] to be overridden on an individual port.
     */
    readonly blockOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * Enable beacon probing on the ports this policy applies to.
     */
    readonly checkBeacon?: pulumi.Input<boolean>;
    /**
     * Version string of the configuration that this spec is trying to change.
     */
    readonly configVersion?: pulumi.Input<string>;
    /**
     * Map of custom attribute ids to attribute
     * value string to set for port group. See [here][docs-setting-custom-attributes]
     * for a reference on how to set values for custom attributes.
     */
    readonly customAttributes?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * An optional description for the port group.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Allow VMDirectPath Gen2 on the ports this policy applies to.
     */
    readonly directpathGen2Allowed?: pulumi.Input<boolean>;
    /**
     * The ID of the DVS to add the
     * port group to. Forces a new resource if changed.
     */
    readonly distributedVirtualSwitchUuid?: pulumi.Input<string>;
    /**
     * The average egress bandwidth in bits per second if egress shaping is enabled on the port.
     */
    readonly egressShapingAverageBandwidth?: pulumi.Input<number>;
    /**
     * The maximum egress burst size allowed in bytes if egress shaping is enabled on the port.
     */
    readonly egressShapingBurstSize?: pulumi.Input<number>;
    /**
     * True if the traffic shaper is enabled for egress traffic on the port.
     */
    readonly egressShapingEnabled?: pulumi.Input<boolean>;
    /**
     * The peak egress bandwidth during bursts in bits per second if egress traffic shaping is enabled on the port.
     */
    readonly egressShapingPeakBandwidth?: pulumi.Input<number>;
    /**
     * If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
     */
    readonly failback?: pulumi.Input<boolean>;
    /**
     * The average ingress bandwidth in bits per second if ingress shaping is enabled on the port.
     */
    readonly ingressShapingAverageBandwidth?: pulumi.Input<number>;
    /**
     * The maximum ingress burst size allowed in bytes if ingress shaping is enabled on the port.
     */
    readonly ingressShapingBurstSize?: pulumi.Input<number>;
    /**
     * True if the traffic shaper is enabled for ingress traffic on the port.
     */
    readonly ingressShapingEnabled?: pulumi.Input<boolean>;
    /**
     * The peak ingress bandwidth during bursts in bits per second if ingress traffic shaping is enabled on the port.
     */
    readonly ingressShapingPeakBandwidth?: pulumi.Input<number>;
    /**
     * The generated UUID of the portgroup.
     */
    readonly key?: pulumi.Input<string>;
    /**
     * Whether or not to enable LACP on all uplink ports.
     */
    readonly lacpEnabled?: pulumi.Input<boolean>;
    /**
     * The uplink LACP mode to use. Can be one of active or passive.
     */
    readonly lacpMode?: pulumi.Input<string>;
    /**
     * Allow a port in this port group to be
     * moved to another port group while it is connected.
     */
    readonly livePortMovingAllowed?: pulumi.Input<boolean>;
    /**
     * The name of the port group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Indicates whether to enable netflow on all ports.
     */
    readonly netflowEnabled?: pulumi.Input<boolean>;
    /**
     * Allow the [Netflow
     * policy][netflow-policy] on this port group to be overridden on an individual
     * port.
     */
    readonly netflowOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * The key of a network resource pool
     * to associate with this port group. The default is `-1`, which implies no
     * association.
     */
    readonly networkResourcePoolKey?: pulumi.Input<string>;
    /**
     * Allow the network
     * resource pool set on this port group to be overridden on an individual port.
     */
    readonly networkResourcePoolOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
     */
    readonly notifySwitches?: pulumi.Input<boolean>;
    /**
     * The number of ports available on this port
     * group. Cannot be decreased below the amount of used ports on the port group.
     */
    readonly numberOfPorts?: pulumi.Input<number>;
    /**
     * Reset a port's settings to the
     * settings defined on this port group policy when the port disconnects.
     */
    readonly portConfigResetAtDisconnect?: pulumi.Input<boolean>;
    /**
     * An optional formatting policy for naming of
     * the ports in this port group. See the `portNameFormat` attribute listed
     * [here][ext-vsphere-portname-format] for details on the format syntax.
     */
    readonly portNameFormat?: pulumi.Input<string>;
    /**
     * The secondary VLAN ID for this port.
     */
    readonly portPrivateSecondaryVlanId?: pulumi.Input<number>;
    /**
     * Allow the [security policy
     * settings][sec-policy-settings] defined in this port group policy to be
     * overridden on an individual port.
     */
    readonly securityPolicyOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * Allow the [traffic shaping
     * options][traffic-shaping-settings] on this port group policy to be overridden
     * on an individual port.
     */
    readonly shapingOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
     */
    readonly standbyUplinks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of tag IDs to apply to this object.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid,
     * failover_explicit, or loadbalance_loadbased.
     */
    readonly teamingPolicy?: pulumi.Input<string>;
    /**
     * Allow any traffic filters on
     * this port group to be overridden on an individual port.
     */
    readonly trafficFilterOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * If true, a copy of packets sent to the switch will always be forwarded to an uplink in addition to the regular
     * packet forwarded done by the switch.
     */
    readonly txUplink?: pulumi.Input<boolean>;
    /**
     * The port group type. Can be one of `earlyBinding` (static
     * binding) or `ephemeral`. Default: `earlyBinding`.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Allow the [uplink teaming
     * options][uplink-teaming-settings] on this port group to be overridden on an
     * individual port.
     */
    readonly uplinkTeamingOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * The VLAN ID for single VLAN mode. 0 denotes no VLAN.
     */
    readonly vlanId?: pulumi.Input<number>;
    /**
     * Allow the [VLAN settings][vlan-settings]
     * on this port group to be overridden on an individual port.
     */
    readonly vlanOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * The VLAN ID for single VLAN mode. 0 denotes no VLAN.
     */
    readonly vlanRanges?: pulumi.Input<pulumi.Input<{
        maxVlan: pulumi.Input<number>;
        minVlan: pulumi.Input<number>;
    }>[]>;
}
/**
 * The set of arguments for constructing a DistributedPortGroup resource.
 */
export interface DistributedPortGroupArgs {
    /**
     * List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
     */
    readonly activeUplinks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address
     * than that of its own.
     */
    readonly allowForgedTransmits?: pulumi.Input<boolean>;
    /**
     * Controls whether or not the Media Access Control (MAC) address can be changed.
     */
    readonly allowMacChanges?: pulumi.Input<boolean>;
    /**
     * Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
     */
    readonly allowPromiscuous?: pulumi.Input<boolean>;
    /**
     * Allows the port group to create additional ports
     * past the limit specified in `number_of_ports` if necessary. Default: `true`.
     */
    readonly autoExpand?: pulumi.Input<boolean>;
    /**
     * Indicates whether to block all ports by default.
     */
    readonly blockAllPorts?: pulumi.Input<boolean>;
    /**
     * Allow the [port shutdown
     * policy][port-shutdown-policy] to be overridden on an individual port.
     */
    readonly blockOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * Enable beacon probing on the ports this policy applies to.
     */
    readonly checkBeacon?: pulumi.Input<boolean>;
    /**
     * Map of custom attribute ids to attribute
     * value string to set for port group. See [here][docs-setting-custom-attributes]
     * for a reference on how to set values for custom attributes.
     */
    readonly customAttributes?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * An optional description for the port group.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Allow VMDirectPath Gen2 on the ports this policy applies to.
     */
    readonly directpathGen2Allowed?: pulumi.Input<boolean>;
    /**
     * The ID of the DVS to add the
     * port group to. Forces a new resource if changed.
     */
    readonly distributedVirtualSwitchUuid: pulumi.Input<string>;
    /**
     * The average egress bandwidth in bits per second if egress shaping is enabled on the port.
     */
    readonly egressShapingAverageBandwidth?: pulumi.Input<number>;
    /**
     * The maximum egress burst size allowed in bytes if egress shaping is enabled on the port.
     */
    readonly egressShapingBurstSize?: pulumi.Input<number>;
    /**
     * True if the traffic shaper is enabled for egress traffic on the port.
     */
    readonly egressShapingEnabled?: pulumi.Input<boolean>;
    /**
     * The peak egress bandwidth during bursts in bits per second if egress traffic shaping is enabled on the port.
     */
    readonly egressShapingPeakBandwidth?: pulumi.Input<number>;
    /**
     * If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
     */
    readonly failback?: pulumi.Input<boolean>;
    /**
     * The average ingress bandwidth in bits per second if ingress shaping is enabled on the port.
     */
    readonly ingressShapingAverageBandwidth?: pulumi.Input<number>;
    /**
     * The maximum ingress burst size allowed in bytes if ingress shaping is enabled on the port.
     */
    readonly ingressShapingBurstSize?: pulumi.Input<number>;
    /**
     * True if the traffic shaper is enabled for ingress traffic on the port.
     */
    readonly ingressShapingEnabled?: pulumi.Input<boolean>;
    /**
     * The peak ingress bandwidth during bursts in bits per second if ingress traffic shaping is enabled on the port.
     */
    readonly ingressShapingPeakBandwidth?: pulumi.Input<number>;
    /**
     * Whether or not to enable LACP on all uplink ports.
     */
    readonly lacpEnabled?: pulumi.Input<boolean>;
    /**
     * The uplink LACP mode to use. Can be one of active or passive.
     */
    readonly lacpMode?: pulumi.Input<string>;
    /**
     * Allow a port in this port group to be
     * moved to another port group while it is connected.
     */
    readonly livePortMovingAllowed?: pulumi.Input<boolean>;
    /**
     * The name of the port group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Indicates whether to enable netflow on all ports.
     */
    readonly netflowEnabled?: pulumi.Input<boolean>;
    /**
     * Allow the [Netflow
     * policy][netflow-policy] on this port group to be overridden on an individual
     * port.
     */
    readonly netflowOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * The key of a network resource pool
     * to associate with this port group. The default is `-1`, which implies no
     * association.
     */
    readonly networkResourcePoolKey?: pulumi.Input<string>;
    /**
     * Allow the network
     * resource pool set on this port group to be overridden on an individual port.
     */
    readonly networkResourcePoolOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
     */
    readonly notifySwitches?: pulumi.Input<boolean>;
    /**
     * The number of ports available on this port
     * group. Cannot be decreased below the amount of used ports on the port group.
     */
    readonly numberOfPorts?: pulumi.Input<number>;
    /**
     * Reset a port's settings to the
     * settings defined on this port group policy when the port disconnects.
     */
    readonly portConfigResetAtDisconnect?: pulumi.Input<boolean>;
    /**
     * An optional formatting policy for naming of
     * the ports in this port group. See the `portNameFormat` attribute listed
     * [here][ext-vsphere-portname-format] for details on the format syntax.
     */
    readonly portNameFormat?: pulumi.Input<string>;
    /**
     * The secondary VLAN ID for this port.
     */
    readonly portPrivateSecondaryVlanId?: pulumi.Input<number>;
    /**
     * Allow the [security policy
     * settings][sec-policy-settings] defined in this port group policy to be
     * overridden on an individual port.
     */
    readonly securityPolicyOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * Allow the [traffic shaping
     * options][traffic-shaping-settings] on this port group policy to be overridden
     * on an individual port.
     */
    readonly shapingOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
     */
    readonly standbyUplinks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of tag IDs to apply to this object.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid,
     * failover_explicit, or loadbalance_loadbased.
     */
    readonly teamingPolicy?: pulumi.Input<string>;
    /**
     * Allow any traffic filters on
     * this port group to be overridden on an individual port.
     */
    readonly trafficFilterOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * If true, a copy of packets sent to the switch will always be forwarded to an uplink in addition to the regular
     * packet forwarded done by the switch.
     */
    readonly txUplink?: pulumi.Input<boolean>;
    /**
     * The port group type. Can be one of `earlyBinding` (static
     * binding) or `ephemeral`. Default: `earlyBinding`.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Allow the [uplink teaming
     * options][uplink-teaming-settings] on this port group to be overridden on an
     * individual port.
     */
    readonly uplinkTeamingOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * The VLAN ID for single VLAN mode. 0 denotes no VLAN.
     */
    readonly vlanId?: pulumi.Input<number>;
    /**
     * Allow the [VLAN settings][vlan-settings]
     * on this port group to be overridden on an individual port.
     */
    readonly vlanOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * The VLAN ID for single VLAN mode. 0 denotes no VLAN.
     */
    readonly vlanRanges?: pulumi.Input<pulumi.Input<{
        maxVlan: pulumi.Input<number>;
        minVlan: pulumi.Input<number>;
    }>[]>;
}

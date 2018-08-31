import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_host_virtual_switch` resource can be used to manage vSphere
 * standard switches on an ESXi host. These switches can be used as a backing for
 * standard port groups, which can be managed by the
 * [`vsphere_host_port_group`][host-port-group] resource.
 *
 * For an overview on vSphere networking concepts, see [this
 * page][ref-vsphere-net-concepts].
 *
 * [host-port-group]: /docs/providers/vsphere/r/host_port_group.html
 * [ref-vsphere-net-concepts]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.networking.doc/GUID-2B11DBB8-CB3C-4AFF-8885-EFEA0FC562F4.html
 */
export declare class HostVirtualSwitch extends pulumi.CustomResource {
    /**
     * Get an existing HostVirtualSwitch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostVirtualSwitchState): HostVirtualSwitch;
    /**
     * The list of active network adapters used for load
     * balancing.
     */
    readonly activeNics: pulumi.Output<string[]>;
    /**
     * Controls whether or not the virtual
     * network adapter is allowed to send network traffic with a different MAC
     * address than that of its own. Default: `true`.
     */
    readonly allowForgedTransmits: pulumi.Output<boolean | undefined>;
    /**
     * Controls whether or not the Media Access
     * Control (MAC) address can be changed. Default: `true`.
     */
    readonly allowMacChanges: pulumi.Output<boolean | undefined>;
    /**
     * Enable promiscuous mode on the network. This
     * flag indicates whether or not all traffic is seen on a given port. Default:
     * `false`.
     */
    readonly allowPromiscuous: pulumi.Output<boolean | undefined>;
    /**
     * The interval, in seconds, that a NIC beacon
     * packet is sent out. This can be used with `check_beacon` to
     * offer link failure capability beyond link status only. Default: `1`.
     */
    readonly beaconInterval: pulumi.Output<number | undefined>;
    /**
     * Enable beacon probing - this requires that the
     * `beacon_interval` option has been set in the bridge
     * options. If this is set to `false`, only link status is used to check for
     * failed NICs.  Default: `false`.
     */
    readonly checkBeacon: pulumi.Output<boolean | undefined>;
    /**
     * If set to `true`, the teaming policy will re-activate
     * failed interfaces higher in precedence when they come back up.  Default:
     * `true`.
     */
    readonly failback: pulumi.Output<boolean | undefined>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the host to set the virtual switch up on. Forces a new resource if changed.
     */
    readonly hostSystemId: pulumi.Output<string>;
    /**
     * Whether to `advertise` or `listen`
     * for link discovery traffic. Default: `listen`.
     */
    readonly linkDiscoveryOperation: pulumi.Output<string | undefined>;
    /**
     * The discovery protocol type.  Valid
     * types are `cpd` and `lldp`. Default: `cdp`.
     */
    readonly linkDiscoveryProtocol: pulumi.Output<string | undefined>;
    /**
     * The maximum transmission unit (MTU) for the virtual
     * switch. Default: `1500`.
     */
    readonly mtu: pulumi.Output<number | undefined>;
    /**
     * The name of the virtual switch. Forces a new resource if
     * changed.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The network interfaces to bind to the bridge.
     */
    readonly networkAdapters: pulumi.Output<string[]>;
    /**
     * If set to `true`, the teaming policy will
     * notify the broadcast network of a NIC failover, triggering cache updates.
     * Default: `true`.
     */
    readonly notifySwitches: pulumi.Output<boolean | undefined>;
    /**
     * The number of ports to create with this
     * virtual switch. Default: `128`.
     */
    readonly numberOfPorts: pulumi.Output<number | undefined>;
    /**
     * The average bandwidth in bits per
     * second if traffic shaping is enabled. Default: `0`
     */
    readonly shapingAverageBandwidth: pulumi.Output<number | undefined>;
    /**
     * The maximum burst size allowed in bytes if
     * shaping is enabled. Default: `0`
     */
    readonly shapingBurstSize: pulumi.Output<number | undefined>;
    /**
     * Set to `true` to enable the traffic shaper for
     * ports managed by this virtual switch. Default: `false`.
     */
    readonly shapingEnabled: pulumi.Output<boolean | undefined>;
    /**
     * The peak bandwidth during bursts in
     * bits per second if traffic shaping is enabled. Default: `0`
     */
    readonly shapingPeakBandwidth: pulumi.Output<number | undefined>;
    /**
     * The list of standby network adapters used for
     * failover.
     */
    readonly standbyNics: pulumi.Output<string[]>;
    /**
     * The network adapter teaming policy. Can be one
     * of `loadbalance_ip`, `loadbalance_srcmac`, `loadbalance_srcid`, or
     * `failover_explicit`. Default: `loadbalance_srcid`.
     */
    readonly teamingPolicy: pulumi.Output<string | undefined>;
    /**
     * Create a HostVirtualSwitch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostVirtualSwitchArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering HostVirtualSwitch resources.
 */
export interface HostVirtualSwitchState {
    /**
     * The list of active network adapters used for load
     * balancing.
     */
    readonly activeNics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Controls whether or not the virtual
     * network adapter is allowed to send network traffic with a different MAC
     * address than that of its own. Default: `true`.
     */
    readonly allowForgedTransmits?: pulumi.Input<boolean>;
    /**
     * Controls whether or not the Media Access
     * Control (MAC) address can be changed. Default: `true`.
     */
    readonly allowMacChanges?: pulumi.Input<boolean>;
    /**
     * Enable promiscuous mode on the network. This
     * flag indicates whether or not all traffic is seen on a given port. Default:
     * `false`.
     */
    readonly allowPromiscuous?: pulumi.Input<boolean>;
    /**
     * The interval, in seconds, that a NIC beacon
     * packet is sent out. This can be used with `check_beacon` to
     * offer link failure capability beyond link status only. Default: `1`.
     */
    readonly beaconInterval?: pulumi.Input<number>;
    /**
     * Enable beacon probing - this requires that the
     * `beacon_interval` option has been set in the bridge
     * options. If this is set to `false`, only link status is used to check for
     * failed NICs.  Default: `false`.
     */
    readonly checkBeacon?: pulumi.Input<boolean>;
    /**
     * If set to `true`, the teaming policy will re-activate
     * failed interfaces higher in precedence when they come back up.  Default:
     * `true`.
     */
    readonly failback?: pulumi.Input<boolean>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the host to set the virtual switch up on. Forces a new resource if changed.
     */
    readonly hostSystemId?: pulumi.Input<string>;
    /**
     * Whether to `advertise` or `listen`
     * for link discovery traffic. Default: `listen`.
     */
    readonly linkDiscoveryOperation?: pulumi.Input<string>;
    /**
     * The discovery protocol type.  Valid
     * types are `cpd` and `lldp`. Default: `cdp`.
     */
    readonly linkDiscoveryProtocol?: pulumi.Input<string>;
    /**
     * The maximum transmission unit (MTU) for the virtual
     * switch. Default: `1500`.
     */
    readonly mtu?: pulumi.Input<number>;
    /**
     * The name of the virtual switch. Forces a new resource if
     * changed.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The network interfaces to bind to the bridge.
     */
    readonly networkAdapters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set to `true`, the teaming policy will
     * notify the broadcast network of a NIC failover, triggering cache updates.
     * Default: `true`.
     */
    readonly notifySwitches?: pulumi.Input<boolean>;
    /**
     * The number of ports to create with this
     * virtual switch. Default: `128`.
     */
    readonly numberOfPorts?: pulumi.Input<number>;
    /**
     * The average bandwidth in bits per
     * second if traffic shaping is enabled. Default: `0`
     */
    readonly shapingAverageBandwidth?: pulumi.Input<number>;
    /**
     * The maximum burst size allowed in bytes if
     * shaping is enabled. Default: `0`
     */
    readonly shapingBurstSize?: pulumi.Input<number>;
    /**
     * Set to `true` to enable the traffic shaper for
     * ports managed by this virtual switch. Default: `false`.
     */
    readonly shapingEnabled?: pulumi.Input<boolean>;
    /**
     * The peak bandwidth during bursts in
     * bits per second if traffic shaping is enabled. Default: `0`
     */
    readonly shapingPeakBandwidth?: pulumi.Input<number>;
    /**
     * The list of standby network adapters used for
     * failover.
     */
    readonly standbyNics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The network adapter teaming policy. Can be one
     * of `loadbalance_ip`, `loadbalance_srcmac`, `loadbalance_srcid`, or
     * `failover_explicit`. Default: `loadbalance_srcid`.
     */
    readonly teamingPolicy?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a HostVirtualSwitch resource.
 */
export interface HostVirtualSwitchArgs {
    /**
     * The list of active network adapters used for load
     * balancing.
     */
    readonly activeNics: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Controls whether or not the virtual
     * network adapter is allowed to send network traffic with a different MAC
     * address than that of its own. Default: `true`.
     */
    readonly allowForgedTransmits?: pulumi.Input<boolean>;
    /**
     * Controls whether or not the Media Access
     * Control (MAC) address can be changed. Default: `true`.
     */
    readonly allowMacChanges?: pulumi.Input<boolean>;
    /**
     * Enable promiscuous mode on the network. This
     * flag indicates whether or not all traffic is seen on a given port. Default:
     * `false`.
     */
    readonly allowPromiscuous?: pulumi.Input<boolean>;
    /**
     * The interval, in seconds, that a NIC beacon
     * packet is sent out. This can be used with `check_beacon` to
     * offer link failure capability beyond link status only. Default: `1`.
     */
    readonly beaconInterval?: pulumi.Input<number>;
    /**
     * Enable beacon probing - this requires that the
     * `beacon_interval` option has been set in the bridge
     * options. If this is set to `false`, only link status is used to check for
     * failed NICs.  Default: `false`.
     */
    readonly checkBeacon?: pulumi.Input<boolean>;
    /**
     * If set to `true`, the teaming policy will re-activate
     * failed interfaces higher in precedence when they come back up.  Default:
     * `true`.
     */
    readonly failback?: pulumi.Input<boolean>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the host to set the virtual switch up on. Forces a new resource if changed.
     */
    readonly hostSystemId: pulumi.Input<string>;
    /**
     * Whether to `advertise` or `listen`
     * for link discovery traffic. Default: `listen`.
     */
    readonly linkDiscoveryOperation?: pulumi.Input<string>;
    /**
     * The discovery protocol type.  Valid
     * types are `cpd` and `lldp`. Default: `cdp`.
     */
    readonly linkDiscoveryProtocol?: pulumi.Input<string>;
    /**
     * The maximum transmission unit (MTU) for the virtual
     * switch. Default: `1500`.
     */
    readonly mtu?: pulumi.Input<number>;
    /**
     * The name of the virtual switch. Forces a new resource if
     * changed.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The network interfaces to bind to the bridge.
     */
    readonly networkAdapters: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set to `true`, the teaming policy will
     * notify the broadcast network of a NIC failover, triggering cache updates.
     * Default: `true`.
     */
    readonly notifySwitches?: pulumi.Input<boolean>;
    /**
     * The number of ports to create with this
     * virtual switch. Default: `128`.
     */
    readonly numberOfPorts?: pulumi.Input<number>;
    /**
     * The average bandwidth in bits per
     * second if traffic shaping is enabled. Default: `0`
     */
    readonly shapingAverageBandwidth?: pulumi.Input<number>;
    /**
     * The maximum burst size allowed in bytes if
     * shaping is enabled. Default: `0`
     */
    readonly shapingBurstSize?: pulumi.Input<number>;
    /**
     * Set to `true` to enable the traffic shaper for
     * ports managed by this virtual switch. Default: `false`.
     */
    readonly shapingEnabled?: pulumi.Input<boolean>;
    /**
     * The peak bandwidth during bursts in
     * bits per second if traffic shaping is enabled. Default: `0`
     */
    readonly shapingPeakBandwidth?: pulumi.Input<number>;
    /**
     * The list of standby network adapters used for
     * failover.
     */
    readonly standbyNics: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The network adapter teaming policy. Can be one
     * of `loadbalance_ip`, `loadbalance_srcmac`, `loadbalance_srcid`, or
     * `failover_explicit`. Default: `loadbalance_srcid`.
     */
    readonly teamingPolicy?: pulumi.Input<string>;
}

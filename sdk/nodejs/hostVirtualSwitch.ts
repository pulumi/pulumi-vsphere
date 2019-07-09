// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class HostVirtualSwitch extends pulumi.CustomResource {
    /**
     * Get an existing HostVirtualSwitch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostVirtualSwitchState, opts?: pulumi.CustomResourceOptions): HostVirtualSwitch {
        return new HostVirtualSwitch(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/hostVirtualSwitch:HostVirtualSwitch';

    /**
     * Returns true if the given object is an instance of HostVirtualSwitch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HostVirtualSwitch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HostVirtualSwitch.__pulumiType;
    }

    /**
     * The list of active network adapters used for load
     * balancing.
     */
    public readonly activeNics!: pulumi.Output<string[]>;
    /**
     * Controls whether or not the virtual
     * network adapter is allowed to send network traffic with a different MAC
     * address than that of its own. Default: `true`.
     */
    public readonly allowForgedTransmits!: pulumi.Output<boolean | undefined>;
    /**
     * Controls whether or not the Media Access
     * Control (MAC) address can be changed. Default: `true`.
     */
    public readonly allowMacChanges!: pulumi.Output<boolean | undefined>;
    /**
     * Enable promiscuous mode on the network. This
     * flag indicates whether or not all traffic is seen on a given port. Default:
     * `false`.
     */
    public readonly allowPromiscuous!: pulumi.Output<boolean | undefined>;
    /**
     * The interval, in seconds, that a NIC beacon
     * packet is sent out. This can be used with `check_beacon` to
     * offer link failure capability beyond link status only. Default: `1`.
     */
    public readonly beaconInterval!: pulumi.Output<number | undefined>;
    /**
     * Enable beacon probing - this requires that the
     * `beacon_interval` option has been set in the bridge
     * options. If this is set to `false`, only link status is used to check for
     * failed NICs.  Default: `false`.
     */
    public readonly checkBeacon!: pulumi.Output<boolean | undefined>;
    /**
     * If set to `true`, the teaming policy will re-activate
     * failed interfaces higher in precedence when they come back up.  Default:
     * `true`.
     */
    public readonly failback!: pulumi.Output<boolean | undefined>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the host to set the virtual switch up on. Forces a new resource if changed.
     */
    public readonly hostSystemId!: pulumi.Output<string>;
    /**
     * Whether to `advertise` or `listen`
     * for link discovery traffic. Default: `listen`.
     */
    public readonly linkDiscoveryOperation!: pulumi.Output<string | undefined>;
    /**
     * The discovery protocol type.  Valid
     * types are `cpd` and `lldp`. Default: `cdp`.
     */
    public readonly linkDiscoveryProtocol!: pulumi.Output<string | undefined>;
    /**
     * The maximum transmission unit (MTU) for the virtual
     * switch. Default: `1500`.
     */
    public readonly mtu!: pulumi.Output<number | undefined>;
    /**
     * The name of the virtual switch. Forces a new resource if
     * changed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The network interfaces to bind to the bridge.
     */
    public readonly networkAdapters!: pulumi.Output<string[]>;
    /**
     * If set to `true`, the teaming policy will
     * notify the broadcast network of a NIC failover, triggering cache updates.
     * Default: `true`.
     */
    public readonly notifySwitches!: pulumi.Output<boolean | undefined>;
    /**
     * The number of ports to create with this
     * virtual switch. Default: `128`.
     */
    public readonly numberOfPorts!: pulumi.Output<number | undefined>;
    /**
     * The average bandwidth in bits per
     * second if traffic shaping is enabled. Default: `0`
     */
    public readonly shapingAverageBandwidth!: pulumi.Output<number | undefined>;
    /**
     * The maximum burst size allowed in bytes if
     * shaping is enabled. Default: `0`
     */
    public readonly shapingBurstSize!: pulumi.Output<number | undefined>;
    /**
     * Set to `true` to enable the traffic shaper for
     * ports managed by this virtual switch. Default: `false`.
     */
    public readonly shapingEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The peak bandwidth during bursts in
     * bits per second if traffic shaping is enabled. Default: `0`
     */
    public readonly shapingPeakBandwidth!: pulumi.Output<number | undefined>;
    /**
     * The list of standby network adapters used for
     * failover.
     */
    public readonly standbyNics!: pulumi.Output<string[]>;
    /**
     * The network adapter teaming policy. Can be one
     * of `loadbalance_ip`, `loadbalance_srcmac`, `loadbalance_srcid`, or
     * `failover_explicit`. Default: `loadbalance_srcid`.
     */
    public readonly teamingPolicy!: pulumi.Output<string | undefined>;

    /**
     * Create a HostVirtualSwitch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostVirtualSwitchArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostVirtualSwitchArgs | HostVirtualSwitchState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as HostVirtualSwitchState | undefined;
            inputs["activeNics"] = state ? state.activeNics : undefined;
            inputs["allowForgedTransmits"] = state ? state.allowForgedTransmits : undefined;
            inputs["allowMacChanges"] = state ? state.allowMacChanges : undefined;
            inputs["allowPromiscuous"] = state ? state.allowPromiscuous : undefined;
            inputs["beaconInterval"] = state ? state.beaconInterval : undefined;
            inputs["checkBeacon"] = state ? state.checkBeacon : undefined;
            inputs["failback"] = state ? state.failback : undefined;
            inputs["hostSystemId"] = state ? state.hostSystemId : undefined;
            inputs["linkDiscoveryOperation"] = state ? state.linkDiscoveryOperation : undefined;
            inputs["linkDiscoveryProtocol"] = state ? state.linkDiscoveryProtocol : undefined;
            inputs["mtu"] = state ? state.mtu : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkAdapters"] = state ? state.networkAdapters : undefined;
            inputs["notifySwitches"] = state ? state.notifySwitches : undefined;
            inputs["numberOfPorts"] = state ? state.numberOfPorts : undefined;
            inputs["shapingAverageBandwidth"] = state ? state.shapingAverageBandwidth : undefined;
            inputs["shapingBurstSize"] = state ? state.shapingBurstSize : undefined;
            inputs["shapingEnabled"] = state ? state.shapingEnabled : undefined;
            inputs["shapingPeakBandwidth"] = state ? state.shapingPeakBandwidth : undefined;
            inputs["standbyNics"] = state ? state.standbyNics : undefined;
            inputs["teamingPolicy"] = state ? state.teamingPolicy : undefined;
        } else {
            const args = argsOrState as HostVirtualSwitchArgs | undefined;
            if (!args || args.activeNics === undefined) {
                throw new Error("Missing required property 'activeNics'");
            }
            if (!args || args.hostSystemId === undefined) {
                throw new Error("Missing required property 'hostSystemId'");
            }
            if (!args || args.networkAdapters === undefined) {
                throw new Error("Missing required property 'networkAdapters'");
            }
            if (!args || args.standbyNics === undefined) {
                throw new Error("Missing required property 'standbyNics'");
            }
            inputs["activeNics"] = args ? args.activeNics : undefined;
            inputs["allowForgedTransmits"] = args ? args.allowForgedTransmits : undefined;
            inputs["allowMacChanges"] = args ? args.allowMacChanges : undefined;
            inputs["allowPromiscuous"] = args ? args.allowPromiscuous : undefined;
            inputs["beaconInterval"] = args ? args.beaconInterval : undefined;
            inputs["checkBeacon"] = args ? args.checkBeacon : undefined;
            inputs["failback"] = args ? args.failback : undefined;
            inputs["hostSystemId"] = args ? args.hostSystemId : undefined;
            inputs["linkDiscoveryOperation"] = args ? args.linkDiscoveryOperation : undefined;
            inputs["linkDiscoveryProtocol"] = args ? args.linkDiscoveryProtocol : undefined;
            inputs["mtu"] = args ? args.mtu : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkAdapters"] = args ? args.networkAdapters : undefined;
            inputs["notifySwitches"] = args ? args.notifySwitches : undefined;
            inputs["numberOfPorts"] = args ? args.numberOfPorts : undefined;
            inputs["shapingAverageBandwidth"] = args ? args.shapingAverageBandwidth : undefined;
            inputs["shapingBurstSize"] = args ? args.shapingBurstSize : undefined;
            inputs["shapingEnabled"] = args ? args.shapingEnabled : undefined;
            inputs["shapingPeakBandwidth"] = args ? args.shapingPeakBandwidth : undefined;
            inputs["standbyNics"] = args ? args.standbyNics : undefined;
            inputs["teamingPolicy"] = args ? args.teamingPolicy : undefined;
        }
        super(HostVirtualSwitch.__pulumiType, name, inputs, opts);
    }
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

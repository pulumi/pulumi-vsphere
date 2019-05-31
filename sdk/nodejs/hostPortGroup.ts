// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere_host_port_group` resource can be used to manage vSphere standard
 * port groups on an ESXi host. These port groups are connected to standard
 * virtual switches, which can be managed by the
 * [`vsphere_host_virtual_switch`][host-virtual-switch] resource.
 * 
 * For an overview on vSphere networking concepts, see [this page][ref-vsphere-net-concepts].
 * 
 * [host-virtual-switch]: /docs/providers/vsphere/r/host_virtual_switch.html
 * [ref-vsphere-net-concepts]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.networking.doc/GUID-2B11DBB8-CB3C-4AFF-8885-EFEA0FC562F4.html
 * 
 * ## Example Usages
 * 
 * **Create a virtual switch and bind a port group to it:**
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 * 
 * const datacenter = pulumi.output(vsphere.getDatacenter({
 *     name: "dc1",
 * }));
 * const esxiHost = datacenter.apply(datacenter => vsphere.getHost({
 *     datacenterId: datacenter.id,
 *     name: "esxi1",
 * }));
 * const switchHostVirtualSwitch = new vsphere.HostVirtualSwitch("switch", {
 *     activeNics: ["vmnic0"],
 *     hostSystemId: esxiHost.id,
 *     networkAdapters: [
 *         "vmnic0",
 *         "vmnic1",
 *     ],
 *     standbyNics: ["vmnic1"],
 * });
 * const pg = new vsphere.HostPortGroup("pg", {
 *     hostSystemId: esxiHost.id,
 *     virtualSwitchName: switchHostVirtualSwitch.name,
 * });
 * ```
 * 
 * **Create a port group with VLAN set and some overrides:**
 * 
 * This example sets the trunk mode VLAN (`4095`, which passes through all tags)
 * and sets
 * [`allow_promiscuous`](https://www.terraform.io/docs/providers/vsphere/r/host_virtual_switch.html#allow_promiscuous)
 * to ensure that all traffic is seen on the port. The latter setting overrides
 * the implicit default of `false` set on the virtual switch.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 * 
 * const datacenter = pulumi.output(vsphere.getDatacenter({
 *     name: "dc1",
 * }));
 * const esxiHost = datacenter.apply(datacenter => vsphere.getHost({
 *     datacenterId: datacenter.id,
 *     name: "esxi1",
 * }));
 * const switchHostVirtualSwitch = new vsphere.HostVirtualSwitch("switch", {
 *     activeNics: ["vmnic0"],
 *     hostSystemId: esxiHost.id,
 *     networkAdapters: [
 *         "vmnic0",
 *         "vmnic1",
 *     ],
 *     standbyNics: ["vmnic1"],
 * });
 * const pg = new vsphere.HostPortGroup("pg", {
 *     allowPromiscuous: true,
 *     hostSystemId: esxiHost.id,
 *     virtualSwitchName: switchHostVirtualSwitch.name,
 *     vlanId: 4095,
 * });
 * ```
 */
export class HostPortGroup extends pulumi.CustomResource {
    /**
     * Get an existing HostPortGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostPortGroupState, opts?: pulumi.CustomResourceOptions): HostPortGroup {
        return new HostPortGroup(name, <any>state, { ...opts, id: id });
    }

    /**
     * List of active network adapters used for load balancing.
     */
    public readonly activeNics!: pulumi.Output<string[] | undefined>;
    /**
     * Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address
     * than that of its own.
     */
    public readonly allowForgedTransmits!: pulumi.Output<boolean | undefined>;
    /**
     * Controls whether or not the Media Access Control (MAC) address can be changed.
     */
    public readonly allowMacChanges!: pulumi.Output<boolean | undefined>;
    /**
     * Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
     */
    public readonly allowPromiscuous!: pulumi.Output<boolean | undefined>;
    /**
     * Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is
     * used only.
     */
    public readonly checkBeacon!: pulumi.Output<boolean | undefined>;
    /**
     * A map with a full set of the [policy
     * options][host-vswitch-policy-options] computed from defaults and overrides,
     * explaining the effective policy for this port group.
     */
    public /*out*/ readonly computedPolicy!: pulumi.Output<{[key: string]: any}>;
    /**
     * If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
     */
    public readonly failback!: pulumi.Output<boolean | undefined>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the host to set the port group up on. Forces a new resource if changed.
     */
    public readonly hostSystemId!: pulumi.Output<string>;
    /**
     * The key for this port group as returned from the vSphere API.
     */
    public /*out*/ readonly key!: pulumi.Output<string>;
    /**
     * The name of the port group.  Forces a new resource if
     * changed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
     */
    public readonly notifySwitches!: pulumi.Output<boolean | undefined>;
    /**
     * A list of ports that currently exist and are used on this port group.
     */
    public /*out*/ readonly ports!: pulumi.Output<{ key: string, macAddresses: string[], type: string }>;
    /**
     * The average bandwidth in bits per second if traffic shaping is enabled.
     */
    public readonly shapingAverageBandwidth!: pulumi.Output<number | undefined>;
    /**
     * The maximum burst size allowed in bytes if traffic shaping is enabled.
     */
    public readonly shapingBurstSize!: pulumi.Output<number | undefined>;
    /**
     * Enable traffic shaping on this virtual switch or port group.
     */
    public readonly shapingEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
     */
    public readonly shapingPeakBandwidth!: pulumi.Output<number | undefined>;
    /**
     * List of standby network adapters used for failover.
     */
    public readonly standbyNics!: pulumi.Output<string[] | undefined>;
    /**
     * The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
     * failover_explicit.
     */
    public readonly teamingPolicy!: pulumi.Output<string | undefined>;
    /**
     * The name of the virtual switch to bind
     * this port group to. Forces a new resource if changed.
     */
    public readonly virtualSwitchName!: pulumi.Output<string>;
    /**
     * The VLAN ID/trunk mode for this port group.  An ID of
     * `0` denotes no tagging, an ID of `1`-`4094` tags with the specific ID, and an
     * ID of `4095` enables trunk mode, allowing the guest to manage its own
     * tagging. Default: `0`.
     */
    public readonly vlanId!: pulumi.Output<number | undefined>;

    /**
     * Create a HostPortGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostPortGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostPortGroupArgs | HostPortGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as HostPortGroupState | undefined;
            inputs["activeNics"] = state ? state.activeNics : undefined;
            inputs["allowForgedTransmits"] = state ? state.allowForgedTransmits : undefined;
            inputs["allowMacChanges"] = state ? state.allowMacChanges : undefined;
            inputs["allowPromiscuous"] = state ? state.allowPromiscuous : undefined;
            inputs["checkBeacon"] = state ? state.checkBeacon : undefined;
            inputs["computedPolicy"] = state ? state.computedPolicy : undefined;
            inputs["failback"] = state ? state.failback : undefined;
            inputs["hostSystemId"] = state ? state.hostSystemId : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notifySwitches"] = state ? state.notifySwitches : undefined;
            inputs["ports"] = state ? state.ports : undefined;
            inputs["shapingAverageBandwidth"] = state ? state.shapingAverageBandwidth : undefined;
            inputs["shapingBurstSize"] = state ? state.shapingBurstSize : undefined;
            inputs["shapingEnabled"] = state ? state.shapingEnabled : undefined;
            inputs["shapingPeakBandwidth"] = state ? state.shapingPeakBandwidth : undefined;
            inputs["standbyNics"] = state ? state.standbyNics : undefined;
            inputs["teamingPolicy"] = state ? state.teamingPolicy : undefined;
            inputs["virtualSwitchName"] = state ? state.virtualSwitchName : undefined;
            inputs["vlanId"] = state ? state.vlanId : undefined;
        } else {
            const args = argsOrState as HostPortGroupArgs | undefined;
            if (!args || args.hostSystemId === undefined) {
                throw new Error("Missing required property 'hostSystemId'");
            }
            if (!args || args.virtualSwitchName === undefined) {
                throw new Error("Missing required property 'virtualSwitchName'");
            }
            inputs["activeNics"] = args ? args.activeNics : undefined;
            inputs["allowForgedTransmits"] = args ? args.allowForgedTransmits : undefined;
            inputs["allowMacChanges"] = args ? args.allowMacChanges : undefined;
            inputs["allowPromiscuous"] = args ? args.allowPromiscuous : undefined;
            inputs["checkBeacon"] = args ? args.checkBeacon : undefined;
            inputs["failback"] = args ? args.failback : undefined;
            inputs["hostSystemId"] = args ? args.hostSystemId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notifySwitches"] = args ? args.notifySwitches : undefined;
            inputs["shapingAverageBandwidth"] = args ? args.shapingAverageBandwidth : undefined;
            inputs["shapingBurstSize"] = args ? args.shapingBurstSize : undefined;
            inputs["shapingEnabled"] = args ? args.shapingEnabled : undefined;
            inputs["shapingPeakBandwidth"] = args ? args.shapingPeakBandwidth : undefined;
            inputs["standbyNics"] = args ? args.standbyNics : undefined;
            inputs["teamingPolicy"] = args ? args.teamingPolicy : undefined;
            inputs["virtualSwitchName"] = args ? args.virtualSwitchName : undefined;
            inputs["vlanId"] = args ? args.vlanId : undefined;
            inputs["computedPolicy"] = undefined /*out*/;
            inputs["key"] = undefined /*out*/;
            inputs["ports"] = undefined /*out*/;
        }
        super("vsphere:index/hostPortGroup:HostPortGroup", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostPortGroup resources.
 */
export interface HostPortGroupState {
    /**
     * List of active network adapters used for load balancing.
     */
    readonly activeNics?: pulumi.Input<pulumi.Input<string>[]>;
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
     * Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is
     * used only.
     */
    readonly checkBeacon?: pulumi.Input<boolean>;
    /**
     * A map with a full set of the [policy
     * options][host-vswitch-policy-options] computed from defaults and overrides,
     * explaining the effective policy for this port group.
     */
    readonly computedPolicy?: pulumi.Input<{[key: string]: any}>;
    /**
     * If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
     */
    readonly failback?: pulumi.Input<boolean>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the host to set the port group up on. Forces a new resource if changed.
     */
    readonly hostSystemId?: pulumi.Input<string>;
    /**
     * The key for this port group as returned from the vSphere API.
     */
    readonly key?: pulumi.Input<string>;
    /**
     * The name of the port group.  Forces a new resource if
     * changed.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
     */
    readonly notifySwitches?: pulumi.Input<boolean>;
    /**
     * A list of ports that currently exist and are used on this port group.
     */
    readonly ports?: pulumi.Input<{ key?: pulumi.Input<string>, macAddresses?: pulumi.Input<pulumi.Input<string>[]>, type?: pulumi.Input<string> }>;
    /**
     * The average bandwidth in bits per second if traffic shaping is enabled.
     */
    readonly shapingAverageBandwidth?: pulumi.Input<number>;
    /**
     * The maximum burst size allowed in bytes if traffic shaping is enabled.
     */
    readonly shapingBurstSize?: pulumi.Input<number>;
    /**
     * Enable traffic shaping on this virtual switch or port group.
     */
    readonly shapingEnabled?: pulumi.Input<boolean>;
    /**
     * The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
     */
    readonly shapingPeakBandwidth?: pulumi.Input<number>;
    /**
     * List of standby network adapters used for failover.
     */
    readonly standbyNics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
     * failover_explicit.
     */
    readonly teamingPolicy?: pulumi.Input<string>;
    /**
     * The name of the virtual switch to bind
     * this port group to. Forces a new resource if changed.
     */
    readonly virtualSwitchName?: pulumi.Input<string>;
    /**
     * The VLAN ID/trunk mode for this port group.  An ID of
     * `0` denotes no tagging, an ID of `1`-`4094` tags with the specific ID, and an
     * ID of `4095` enables trunk mode, allowing the guest to manage its own
     * tagging. Default: `0`.
     */
    readonly vlanId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a HostPortGroup resource.
 */
export interface HostPortGroupArgs {
    /**
     * List of active network adapters used for load balancing.
     */
    readonly activeNics?: pulumi.Input<pulumi.Input<string>[]>;
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
     * Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is
     * used only.
     */
    readonly checkBeacon?: pulumi.Input<boolean>;
    /**
     * If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
     */
    readonly failback?: pulumi.Input<boolean>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the host to set the port group up on. Forces a new resource if changed.
     */
    readonly hostSystemId: pulumi.Input<string>;
    /**
     * The name of the port group.  Forces a new resource if
     * changed.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
     */
    readonly notifySwitches?: pulumi.Input<boolean>;
    /**
     * The average bandwidth in bits per second if traffic shaping is enabled.
     */
    readonly shapingAverageBandwidth?: pulumi.Input<number>;
    /**
     * The maximum burst size allowed in bytes if traffic shaping is enabled.
     */
    readonly shapingBurstSize?: pulumi.Input<number>;
    /**
     * Enable traffic shaping on this virtual switch or port group.
     */
    readonly shapingEnabled?: pulumi.Input<boolean>;
    /**
     * The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
     */
    readonly shapingPeakBandwidth?: pulumi.Input<number>;
    /**
     * List of standby network adapters used for failover.
     */
    readonly standbyNics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
     * failover_explicit.
     */
    readonly teamingPolicy?: pulumi.Input<string>;
    /**
     * The name of the virtual switch to bind
     * this port group to. Forces a new resource if changed.
     */
    readonly virtualSwitchName: pulumi.Input<string>;
    /**
     * The VLAN ID/trunk mode for this port group.  An ID of
     * `0` denotes no tagging, an ID of `1`-`4094` tags with the specific ID, and an
     * ID of `4095` enables trunk mode, allowing the guest to manage its own
     * tagging. Default: `0`.
     */
    readonly vlanId?: pulumi.Input<number>;
}

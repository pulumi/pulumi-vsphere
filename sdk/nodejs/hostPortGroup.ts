// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `vsphere.HostPortGroup` resource can be used to manage port groups on
 * ESXi hosts. These port groups are connected to standard switches, which
 * can be managed by the `vsphere.HostVirtualSwitch`
 * resource.
 *
 * For an overview on vSphere networking concepts, see [the product documentation][ref-vsphere-net-concepts].
 *
 * [ref-vsphere-net-concepts]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.networking.doc/GUID-2B11DBB8-CB3C-4AFF-8885-EFEA0FC562F4.html
 *
 * ## Example Usage
 *
 * **Create a Virtual Switch and Bind a Port Group:**
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const host = datacenter.then(datacenter => vsphere.getHost({
 *     name: "esxi-01.example.com",
 *     datacenterId: datacenter.id,
 * }));
 * const hostVirtualSwitch = new vsphere.HostVirtualSwitch("hostVirtualSwitch", {
 *     hostSystemId: host.then(host => host.id),
 *     networkAdapters: [
 *         "vmnic0",
 *         "vmnic1",
 *     ],
 *     activeNics: ["vmnic0"],
 *     standbyNics: ["vmnic1"],
 * });
 * const pg = new vsphere.HostPortGroup("pg", {
 *     hostSystemId: host.then(host => host.id),
 *     virtualSwitchName: hostVirtualSwitch.name,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * **Create a Port Group with a VLAN and ab Override:**
 *
 * This example sets the trunk mode VLAN (`4095`, which passes through all tags)
 * and sets
 * `allowPromiscuous`
 * to ensure that all traffic is seen on the port. The setting overrides
 * the implicit default of `false` set on the standard switch.
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const host = datacenter.then(datacenter => vsphere.getHost({
 *     name: "esxi-01.example.com",
 *     datacenterId: datacenter.id,
 * }));
 * const hostVirtualSwitch = new vsphere.HostVirtualSwitch("hostVirtualSwitch", {
 *     hostSystemId: host.then(host => host.id),
 *     networkAdapters: [
 *         "vmnic0",
 *         "vmnic1",
 *     ],
 *     activeNics: ["vmnic0"],
 *     standbyNics: ["vmnic1"],
 * });
 * const pg = new vsphere.HostPortGroup("pg", {
 *     hostSystemId: host.then(host => host.id),
 *     virtualSwitchName: hostVirtualSwitch.name,
 *     vlanId: 4095,
 *     allowPromiscuous: true,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Importing
 *
 * An existing host port group can be imported into this resource
 * using the host port group's ID. An example is below:
 *
 * The above would import the `management` host port group from host with ID `host-123`.
 */
export class HostPortGroup extends pulumi.CustomResource {
    /**
     * Get an existing HostPortGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostPortGroupState, opts?: pulumi.CustomResourceOptions): HostPortGroup {
        return new HostPortGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/hostPortGroup:HostPortGroup';

    /**
     * Returns true if the given object is an instance of HostPortGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HostPortGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HostPortGroup.__pulumiType;
    }

    /**
     * List of active network adapters used for load balancing.
     */
    public readonly activeNics!: pulumi.Output<string[] | undefined>;
    /**
     * Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
     * that of its own.
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
     * Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
     * only.
     */
    public readonly checkBeacon!: pulumi.Output<boolean | undefined>;
    /**
     * A map with a full set of the policy
     * options computed from defaults and overrides,
     * explaining the effective policy for this port group.
     */
    public /*out*/ readonly computedPolicy!: pulumi.Output<{[key: string]: string}>;
    /**
     * If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
     */
    public readonly failback!: pulumi.Output<boolean | undefined>;
    /**
     * The managed object ID of
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
    public /*out*/ readonly ports!: pulumi.Output<outputs.HostPortGroupPort[]>;
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HostPortGroupState | undefined;
            resourceInputs["activeNics"] = state ? state.activeNics : undefined;
            resourceInputs["allowForgedTransmits"] = state ? state.allowForgedTransmits : undefined;
            resourceInputs["allowMacChanges"] = state ? state.allowMacChanges : undefined;
            resourceInputs["allowPromiscuous"] = state ? state.allowPromiscuous : undefined;
            resourceInputs["checkBeacon"] = state ? state.checkBeacon : undefined;
            resourceInputs["computedPolicy"] = state ? state.computedPolicy : undefined;
            resourceInputs["failback"] = state ? state.failback : undefined;
            resourceInputs["hostSystemId"] = state ? state.hostSystemId : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notifySwitches"] = state ? state.notifySwitches : undefined;
            resourceInputs["ports"] = state ? state.ports : undefined;
            resourceInputs["shapingAverageBandwidth"] = state ? state.shapingAverageBandwidth : undefined;
            resourceInputs["shapingBurstSize"] = state ? state.shapingBurstSize : undefined;
            resourceInputs["shapingEnabled"] = state ? state.shapingEnabled : undefined;
            resourceInputs["shapingPeakBandwidth"] = state ? state.shapingPeakBandwidth : undefined;
            resourceInputs["standbyNics"] = state ? state.standbyNics : undefined;
            resourceInputs["teamingPolicy"] = state ? state.teamingPolicy : undefined;
            resourceInputs["virtualSwitchName"] = state ? state.virtualSwitchName : undefined;
            resourceInputs["vlanId"] = state ? state.vlanId : undefined;
        } else {
            const args = argsOrState as HostPortGroupArgs | undefined;
            if ((!args || args.hostSystemId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostSystemId'");
            }
            if ((!args || args.virtualSwitchName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualSwitchName'");
            }
            resourceInputs["activeNics"] = args ? args.activeNics : undefined;
            resourceInputs["allowForgedTransmits"] = args ? args.allowForgedTransmits : undefined;
            resourceInputs["allowMacChanges"] = args ? args.allowMacChanges : undefined;
            resourceInputs["allowPromiscuous"] = args ? args.allowPromiscuous : undefined;
            resourceInputs["checkBeacon"] = args ? args.checkBeacon : undefined;
            resourceInputs["failback"] = args ? args.failback : undefined;
            resourceInputs["hostSystemId"] = args ? args.hostSystemId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notifySwitches"] = args ? args.notifySwitches : undefined;
            resourceInputs["shapingAverageBandwidth"] = args ? args.shapingAverageBandwidth : undefined;
            resourceInputs["shapingBurstSize"] = args ? args.shapingBurstSize : undefined;
            resourceInputs["shapingEnabled"] = args ? args.shapingEnabled : undefined;
            resourceInputs["shapingPeakBandwidth"] = args ? args.shapingPeakBandwidth : undefined;
            resourceInputs["standbyNics"] = args ? args.standbyNics : undefined;
            resourceInputs["teamingPolicy"] = args ? args.teamingPolicy : undefined;
            resourceInputs["virtualSwitchName"] = args ? args.virtualSwitchName : undefined;
            resourceInputs["vlanId"] = args ? args.vlanId : undefined;
            resourceInputs["computedPolicy"] = undefined /*out*/;
            resourceInputs["key"] = undefined /*out*/;
            resourceInputs["ports"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HostPortGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostPortGroup resources.
 */
export interface HostPortGroupState {
    /**
     * List of active network adapters used for load balancing.
     */
    activeNics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
     * that of its own.
     */
    allowForgedTransmits?: pulumi.Input<boolean>;
    /**
     * Controls whether or not the Media Access Control (MAC) address can be changed.
     */
    allowMacChanges?: pulumi.Input<boolean>;
    /**
     * Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
     */
    allowPromiscuous?: pulumi.Input<boolean>;
    /**
     * Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
     * only.
     */
    checkBeacon?: pulumi.Input<boolean>;
    /**
     * A map with a full set of the policy
     * options computed from defaults and overrides,
     * explaining the effective policy for this port group.
     */
    computedPolicy?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
     */
    failback?: pulumi.Input<boolean>;
    /**
     * The managed object ID of
     * the host to set the port group up on. Forces a new resource if changed.
     */
    hostSystemId?: pulumi.Input<string>;
    /**
     * The key for this port group as returned from the vSphere API.
     */
    key?: pulumi.Input<string>;
    /**
     * The name of the port group.  Forces a new resource if
     * changed.
     */
    name?: pulumi.Input<string>;
    /**
     * If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
     */
    notifySwitches?: pulumi.Input<boolean>;
    /**
     * A list of ports that currently exist and are used on this port group.
     */
    ports?: pulumi.Input<pulumi.Input<inputs.HostPortGroupPort>[]>;
    /**
     * The average bandwidth in bits per second if traffic shaping is enabled.
     */
    shapingAverageBandwidth?: pulumi.Input<number>;
    /**
     * The maximum burst size allowed in bytes if traffic shaping is enabled.
     */
    shapingBurstSize?: pulumi.Input<number>;
    /**
     * Enable traffic shaping on this virtual switch or port group.
     */
    shapingEnabled?: pulumi.Input<boolean>;
    /**
     * The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
     */
    shapingPeakBandwidth?: pulumi.Input<number>;
    /**
     * List of standby network adapters used for failover.
     */
    standbyNics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
     * failover_explicit.
     */
    teamingPolicy?: pulumi.Input<string>;
    /**
     * The name of the virtual switch to bind
     * this port group to. Forces a new resource if changed.
     */
    virtualSwitchName?: pulumi.Input<string>;
    /**
     * The VLAN ID/trunk mode for this port group.  An ID of
     * `0` denotes no tagging, an ID of `1`-`4094` tags with the specific ID, and an
     * ID of `4095` enables trunk mode, allowing the guest to manage its own
     * tagging. Default: `0`.
     */
    vlanId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a HostPortGroup resource.
 */
export interface HostPortGroupArgs {
    /**
     * List of active network adapters used for load balancing.
     */
    activeNics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
     * that of its own.
     */
    allowForgedTransmits?: pulumi.Input<boolean>;
    /**
     * Controls whether or not the Media Access Control (MAC) address can be changed.
     */
    allowMacChanges?: pulumi.Input<boolean>;
    /**
     * Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
     */
    allowPromiscuous?: pulumi.Input<boolean>;
    /**
     * Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
     * only.
     */
    checkBeacon?: pulumi.Input<boolean>;
    /**
     * If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
     */
    failback?: pulumi.Input<boolean>;
    /**
     * The managed object ID of
     * the host to set the port group up on. Forces a new resource if changed.
     */
    hostSystemId: pulumi.Input<string>;
    /**
     * The name of the port group.  Forces a new resource if
     * changed.
     */
    name?: pulumi.Input<string>;
    /**
     * If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
     */
    notifySwitches?: pulumi.Input<boolean>;
    /**
     * The average bandwidth in bits per second if traffic shaping is enabled.
     */
    shapingAverageBandwidth?: pulumi.Input<number>;
    /**
     * The maximum burst size allowed in bytes if traffic shaping is enabled.
     */
    shapingBurstSize?: pulumi.Input<number>;
    /**
     * Enable traffic shaping on this virtual switch or port group.
     */
    shapingEnabled?: pulumi.Input<boolean>;
    /**
     * The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
     */
    shapingPeakBandwidth?: pulumi.Input<number>;
    /**
     * List of standby network adapters used for failover.
     */
    standbyNics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
     * failover_explicit.
     */
    teamingPolicy?: pulumi.Input<string>;
    /**
     * The name of the virtual switch to bind
     * this port group to. Forces a new resource if changed.
     */
    virtualSwitchName: pulumi.Input<string>;
    /**
     * The VLAN ID/trunk mode for this port group.  An ID of
     * `0` denotes no tagging, an ID of `1`-`4094` tags with the specific ID, and an
     * ID of `4095` enables trunk mode, allowing the guest to manage its own
     * tagging. Default: `0`.
     */
    vlanId?: pulumi.Input<number>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.HostVirtualSwitch` resource can be used to manage vSphere
 * standard switches on an ESXi host. These switches can be used as a backing for
 * standard port groups, which can be managed by the
 * `vsphere.HostPortGroup` resource.
 *
 * For an overview on vSphere networking concepts, see [this
 * page][ref-vsphere-net-concepts].
 *
 * [ref-vsphere-net-concepts]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-networking-8-0/introduction-to-vsphere-networking.html
 *
 * ## Example Usage
 *
 * ### Create a virtual switch with one active and one standby NIC
 *
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
 * const _switch = new vsphere.HostVirtualSwitch("switch", {
 *     name: "vSwitchTest",
 *     hostSystemId: host.then(host => host.id),
 *     networkAdapters: [
 *         "vmnic0",
 *         "vmnic1",
 *     ],
 *     activeNics: ["vmnic0"],
 *     standbyNics: ["vmnic1"],
 * });
 * ```
 *
 * ### Create a virtual switch with extra networking policy options
 *
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
 * const _switch = new vsphere.HostVirtualSwitch("switch", {
 *     name: "vSwitchTest",
 *     hostSystemId: host.then(host => host.id),
 *     networkAdapters: [
 *         "vmnic0",
 *         "vmnic1",
 *     ],
 *     activeNics: ["vmnic0"],
 *     standbyNics: ["vmnic1"],
 *     teamingPolicy: "failover_explicit",
 *     allowPromiscuous: false,
 *     allowForgedTransmits: false,
 *     allowMacChanges: false,
 *     shapingEnabled: true,
 *     shapingAverageBandwidth: 50000000,
 *     shapingPeakBandwidth: 100000000,
 *     shapingBurstSize: 1000000000,
 * });
 * ```
 *
 * ## Import
 *
 * An existing vSwitch can be imported into this resource by its ID.
 *
 * The convention of the id is a prefix, the host system [managed objectID][docs-about-morefs], and the virtual switch
 *
 * name. An example would be `tf-HostVirtualSwitch:host-10:vSwitchTerraformTest`.
 *
 * Import can the be done via the following command:
 *
 * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
 *
 * ```sh
 * $ pulumi import vsphere:index/hostVirtualSwitch:HostVirtualSwitch switch tf-HostVirtualSwitch:host-10:vSwitchTerraformTest
 * ```
 *
 * The above would import the vSwitch named `vSwitchTerraformTest` that is located in the `host-10`
 *
 * vSphere host.
 */
export class HostVirtualSwitch extends pulumi.CustomResource {
    /**
     * Get an existing HostVirtualSwitch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
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
     * List of active network adapters used for load balancing.
     */
    public readonly activeNics!: pulumi.Output<string[]>;
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
     * Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
     */
    public readonly beaconInterval!: pulumi.Output<number | undefined>;
    /**
     * Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
     * only.
     */
    public readonly checkBeacon!: pulumi.Output<boolean | undefined>;
    /**
     * If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
     */
    public readonly failback!: pulumi.Output<boolean | undefined>;
    /**
     * The managed object ID of
     * the host to set the virtual switch up on. Forces a new resource if changed.
     */
    public readonly hostSystemId!: pulumi.Output<string>;
    /**
     * Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
     */
    public readonly linkDiscoveryOperation!: pulumi.Output<string | undefined>;
    /**
     * The discovery protocol type. Valid values are cdp and lldp.
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
     * The list of network adapters to bind to this virtual switch.
     */
    public readonly networkAdapters!: pulumi.Output<string[]>;
    /**
     * If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
     */
    public readonly notifySwitches!: pulumi.Output<boolean | undefined>;
    /**
     * The number of ports to create with this
     * virtual switch. Default: `128`.
     *
     * > **NOTE:** Changing the port count requires a reboot of the host. This provider
     * will not restart the host for you.
     */
    public readonly numberOfPorts!: pulumi.Output<number | undefined>;
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
     * Create a HostVirtualSwitch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostVirtualSwitchArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostVirtualSwitchArgs | HostVirtualSwitchState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HostVirtualSwitchState | undefined;
            resourceInputs["activeNics"] = state ? state.activeNics : undefined;
            resourceInputs["allowForgedTransmits"] = state ? state.allowForgedTransmits : undefined;
            resourceInputs["allowMacChanges"] = state ? state.allowMacChanges : undefined;
            resourceInputs["allowPromiscuous"] = state ? state.allowPromiscuous : undefined;
            resourceInputs["beaconInterval"] = state ? state.beaconInterval : undefined;
            resourceInputs["checkBeacon"] = state ? state.checkBeacon : undefined;
            resourceInputs["failback"] = state ? state.failback : undefined;
            resourceInputs["hostSystemId"] = state ? state.hostSystemId : undefined;
            resourceInputs["linkDiscoveryOperation"] = state ? state.linkDiscoveryOperation : undefined;
            resourceInputs["linkDiscoveryProtocol"] = state ? state.linkDiscoveryProtocol : undefined;
            resourceInputs["mtu"] = state ? state.mtu : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkAdapters"] = state ? state.networkAdapters : undefined;
            resourceInputs["notifySwitches"] = state ? state.notifySwitches : undefined;
            resourceInputs["numberOfPorts"] = state ? state.numberOfPorts : undefined;
            resourceInputs["shapingAverageBandwidth"] = state ? state.shapingAverageBandwidth : undefined;
            resourceInputs["shapingBurstSize"] = state ? state.shapingBurstSize : undefined;
            resourceInputs["shapingEnabled"] = state ? state.shapingEnabled : undefined;
            resourceInputs["shapingPeakBandwidth"] = state ? state.shapingPeakBandwidth : undefined;
            resourceInputs["standbyNics"] = state ? state.standbyNics : undefined;
            resourceInputs["teamingPolicy"] = state ? state.teamingPolicy : undefined;
        } else {
            const args = argsOrState as HostVirtualSwitchArgs | undefined;
            if ((!args || args.activeNics === undefined) && !opts.urn) {
                throw new Error("Missing required property 'activeNics'");
            }
            if ((!args || args.hostSystemId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostSystemId'");
            }
            if ((!args || args.networkAdapters === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkAdapters'");
            }
            resourceInputs["activeNics"] = args ? args.activeNics : undefined;
            resourceInputs["allowForgedTransmits"] = args ? args.allowForgedTransmits : undefined;
            resourceInputs["allowMacChanges"] = args ? args.allowMacChanges : undefined;
            resourceInputs["allowPromiscuous"] = args ? args.allowPromiscuous : undefined;
            resourceInputs["beaconInterval"] = args ? args.beaconInterval : undefined;
            resourceInputs["checkBeacon"] = args ? args.checkBeacon : undefined;
            resourceInputs["failback"] = args ? args.failback : undefined;
            resourceInputs["hostSystemId"] = args ? args.hostSystemId : undefined;
            resourceInputs["linkDiscoveryOperation"] = args ? args.linkDiscoveryOperation : undefined;
            resourceInputs["linkDiscoveryProtocol"] = args ? args.linkDiscoveryProtocol : undefined;
            resourceInputs["mtu"] = args ? args.mtu : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkAdapters"] = args ? args.networkAdapters : undefined;
            resourceInputs["notifySwitches"] = args ? args.notifySwitches : undefined;
            resourceInputs["numberOfPorts"] = args ? args.numberOfPorts : undefined;
            resourceInputs["shapingAverageBandwidth"] = args ? args.shapingAverageBandwidth : undefined;
            resourceInputs["shapingBurstSize"] = args ? args.shapingBurstSize : undefined;
            resourceInputs["shapingEnabled"] = args ? args.shapingEnabled : undefined;
            resourceInputs["shapingPeakBandwidth"] = args ? args.shapingPeakBandwidth : undefined;
            resourceInputs["standbyNics"] = args ? args.standbyNics : undefined;
            resourceInputs["teamingPolicy"] = args ? args.teamingPolicy : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HostVirtualSwitch.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostVirtualSwitch resources.
 */
export interface HostVirtualSwitchState {
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
     * Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
     */
    beaconInterval?: pulumi.Input<number>;
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
     * the host to set the virtual switch up on. Forces a new resource if changed.
     */
    hostSystemId?: pulumi.Input<string>;
    /**
     * Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
     */
    linkDiscoveryOperation?: pulumi.Input<string>;
    /**
     * The discovery protocol type. Valid values are cdp and lldp.
     */
    linkDiscoveryProtocol?: pulumi.Input<string>;
    /**
     * The maximum transmission unit (MTU) for the virtual
     * switch. Default: `1500`.
     */
    mtu?: pulumi.Input<number>;
    /**
     * The name of the virtual switch. Forces a new resource if
     * changed.
     */
    name?: pulumi.Input<string>;
    /**
     * The list of network adapters to bind to this virtual switch.
     */
    networkAdapters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
     */
    notifySwitches?: pulumi.Input<boolean>;
    /**
     * The number of ports to create with this
     * virtual switch. Default: `128`.
     *
     * > **NOTE:** Changing the port count requires a reboot of the host. This provider
     * will not restart the host for you.
     */
    numberOfPorts?: pulumi.Input<number>;
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
}

/**
 * The set of arguments for constructing a HostVirtualSwitch resource.
 */
export interface HostVirtualSwitchArgs {
    /**
     * List of active network adapters used for load balancing.
     */
    activeNics: pulumi.Input<pulumi.Input<string>[]>;
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
     * Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
     */
    beaconInterval?: pulumi.Input<number>;
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
     * the host to set the virtual switch up on. Forces a new resource if changed.
     */
    hostSystemId: pulumi.Input<string>;
    /**
     * Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
     */
    linkDiscoveryOperation?: pulumi.Input<string>;
    /**
     * The discovery protocol type. Valid values are cdp and lldp.
     */
    linkDiscoveryProtocol?: pulumi.Input<string>;
    /**
     * The maximum transmission unit (MTU) for the virtual
     * switch. Default: `1500`.
     */
    mtu?: pulumi.Input<number>;
    /**
     * The name of the virtual switch. Forces a new resource if
     * changed.
     */
    name?: pulumi.Input<string>;
    /**
     * The list of network adapters to bind to this virtual switch.
     */
    networkAdapters: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
     */
    notifySwitches?: pulumi.Input<boolean>;
    /**
     * The number of ports to create with this
     * virtual switch. Default: `128`.
     *
     * > **NOTE:** Changing the port count requires a reboot of the host. This provider
     * will not restart the host for you.
     */
    numberOfPorts?: pulumi.Input<number>;
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
}

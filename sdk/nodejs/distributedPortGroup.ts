// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class DistributedPortGroup extends pulumi.CustomResource {
    /**
     * Get an existing DistributedPortGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DistributedPortGroupState, opts?: pulumi.CustomResourceOptions): DistributedPortGroup {
        return new DistributedPortGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/distributedPortGroup:DistributedPortGroup';

    /**
     * Returns true if the given object is an instance of DistributedPortGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DistributedPortGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DistributedPortGroup.__pulumiType;
    }

    /**
     * List of active uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
     */
    public readonly activeUplinks!: pulumi.Output<string[]>;
    /**
     * Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
     * that of its own.
     */
    public readonly allowForgedTransmits!: pulumi.Output<boolean>;
    /**
     * Controls whether or not the Media Access Control (MAC) address can be changed.
     */
    public readonly allowMacChanges!: pulumi.Output<boolean>;
    /**
     * Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
     */
    public readonly allowPromiscuous!: pulumi.Output<boolean>;
    /**
     * Allows the port group to create additional ports
     * past the limit specified in `numberOfPorts` if necessary. Default: `true`.
     */
    public readonly autoExpand!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether to block all ports by default.
     */
    public readonly blockAllPorts!: pulumi.Output<boolean>;
    /**
     * Allow the port shutdown
     * policy to be overridden on an individual port.
     */
    public readonly blockOverrideAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * Enable beacon probing on the ports this policy applies to.
     */
    public readonly checkBeacon!: pulumi.Output<boolean>;
    /**
     * Version string of the configuration that this spec is trying to change.
     */
    public /*out*/ readonly configVersion!: pulumi.Output<string>;
    /**
     * Map of custom attribute ids to attribute
     * value string to set for port group.
     */
    public readonly customAttributes!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * An optional description for the port group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Allow VMDirectPath Gen2 on the ports this policy applies to.
     */
    public readonly directpathGen2Allowed!: pulumi.Output<boolean>;
    /**
     * The ID of the DVS to add the
     * port group to. Forces a new resource if changed.
     */
    public readonly distributedVirtualSwitchUuid!: pulumi.Output<string>;
    /**
     * The average egress bandwidth in bits per second if egress shaping is enabled on the port.
     */
    public readonly egressShapingAverageBandwidth!: pulumi.Output<number>;
    /**
     * The maximum egress burst size allowed in bytes if egress shaping is enabled on the port.
     */
    public readonly egressShapingBurstSize!: pulumi.Output<number>;
    /**
     * True if the traffic shaper is enabled for egress traffic on the port.
     */
    public readonly egressShapingEnabled!: pulumi.Output<boolean>;
    /**
     * The peak egress bandwidth during bursts in bits per second if egress traffic shaping is enabled on the port.
     */
    public readonly egressShapingPeakBandwidth!: pulumi.Output<number>;
    /**
     * If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
     */
    public readonly failback!: pulumi.Output<boolean>;
    /**
     * The average ingress bandwidth in bits per second if ingress shaping is enabled on the port.
     */
    public readonly ingressShapingAverageBandwidth!: pulumi.Output<number>;
    /**
     * The maximum ingress burst size allowed in bytes if ingress shaping is enabled on the port.
     */
    public readonly ingressShapingBurstSize!: pulumi.Output<number>;
    /**
     * True if the traffic shaper is enabled for ingress traffic on the port.
     */
    public readonly ingressShapingEnabled!: pulumi.Output<boolean>;
    /**
     * The peak ingress bandwidth during bursts in bits per second if ingress traffic shaping is enabled on the port.
     */
    public readonly ingressShapingPeakBandwidth!: pulumi.Output<number>;
    /**
     * The generated UUID of the portgroup.
     */
    public /*out*/ readonly key!: pulumi.Output<string>;
    /**
     * Whether or not to enable LACP on all uplink ports.
     */
    public readonly lacpEnabled!: pulumi.Output<boolean>;
    /**
     * The uplink LACP mode to use. Can be one of active or passive.
     */
    public readonly lacpMode!: pulumi.Output<string>;
    /**
     * Allow a port in this port group to be
     * moved to another port group while it is connected.
     */
    public readonly livePortMovingAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the port group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Indicates whether to enable netflow on all ports.
     */
    public readonly netflowEnabled!: pulumi.Output<boolean>;
    /**
     * Allow the Netflow
     * policy on this port group to be overridden on an individual
     * port.
     */
    public readonly netflowOverrideAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * The key of a network resource pool
     * to associate with this port group. The default is `-1`, which implies no
     * association.
     */
    public readonly networkResourcePoolKey!: pulumi.Output<string | undefined>;
    /**
     * Allow the network
     * resource pool set on this port group to be overridden on an individual port.
     */
    public readonly networkResourcePoolOverrideAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
     */
    public readonly notifySwitches!: pulumi.Output<boolean>;
    /**
     * The number of ports available on this port
     * group. Cannot be decreased below the amount of used ports on the port group.
     */
    public readonly numberOfPorts!: pulumi.Output<number>;
    /**
     * Reset a port's settings to the
     * settings defined on this port group policy when the port disconnects.
     */
    public readonly portConfigResetAtDisconnect!: pulumi.Output<boolean | undefined>;
    /**
     * An optional formatting policy for naming of
     * the ports in this port group. See the `portNameFormat` attribute listed
     * [here][ext-vsphere-portname-format] for details on the format syntax.
     */
    public readonly portNameFormat!: pulumi.Output<string | undefined>;
    /**
     * The secondary VLAN ID for this port.
     */
    public readonly portPrivateSecondaryVlanId!: pulumi.Output<number>;
    /**
     * Allow the security policy
     * settings defined in this port group policy to be
     * overridden on an individual port.
     */
    public readonly securityPolicyOverrideAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * Allow the traffic shaping
     * options on this port group policy to be overridden
     * on an individual port.
     */
    public readonly shapingOverrideAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * List of standby uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
     */
    public readonly standbyUplinks!: pulumi.Output<string[]>;
    /**
     * A list of tag IDs to apply to this object.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid,
     * failover_explicit, or loadbalance_loadbased.
     */
    public readonly teamingPolicy!: pulumi.Output<string>;
    /**
     * Allow any traffic filters on
     * this port group to be overridden on an individual port.
     */
    public readonly trafficFilterOverrideAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * If true, a copy of packets sent to the switch will always be forwarded to an uplink in addition to the regular packet
     * forwarded done by the switch.
     */
    public readonly txUplink!: pulumi.Output<boolean>;
    /**
     * The port group type. Can be one of `earlyBinding` (static
     * binding) or `ephemeral`. Default: `earlyBinding`.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * Allow the uplink teaming
     * options on this port group to be overridden on an
     * individual port.
     */
    public readonly uplinkTeamingOverrideAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * The VLAN ID for single VLAN mode. 0 denotes no VLAN.
     */
    public readonly vlanId!: pulumi.Output<number>;
    /**
     * Allow the VLAN settings
     * on this port group to be overridden on an individual port.
     */
    public readonly vlanOverrideAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * The VLAN ID for single VLAN mode. 0 denotes no VLAN.
     */
    public readonly vlanRanges!: pulumi.Output<outputs.DistributedPortGroupVlanRange[]>;

    /**
     * Create a DistributedPortGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DistributedPortGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DistributedPortGroupArgs | DistributedPortGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DistributedPortGroupState | undefined;
            inputs["activeUplinks"] = state ? state.activeUplinks : undefined;
            inputs["allowForgedTransmits"] = state ? state.allowForgedTransmits : undefined;
            inputs["allowMacChanges"] = state ? state.allowMacChanges : undefined;
            inputs["allowPromiscuous"] = state ? state.allowPromiscuous : undefined;
            inputs["autoExpand"] = state ? state.autoExpand : undefined;
            inputs["blockAllPorts"] = state ? state.blockAllPorts : undefined;
            inputs["blockOverrideAllowed"] = state ? state.blockOverrideAllowed : undefined;
            inputs["checkBeacon"] = state ? state.checkBeacon : undefined;
            inputs["configVersion"] = state ? state.configVersion : undefined;
            inputs["customAttributes"] = state ? state.customAttributes : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["directpathGen2Allowed"] = state ? state.directpathGen2Allowed : undefined;
            inputs["distributedVirtualSwitchUuid"] = state ? state.distributedVirtualSwitchUuid : undefined;
            inputs["egressShapingAverageBandwidth"] = state ? state.egressShapingAverageBandwidth : undefined;
            inputs["egressShapingBurstSize"] = state ? state.egressShapingBurstSize : undefined;
            inputs["egressShapingEnabled"] = state ? state.egressShapingEnabled : undefined;
            inputs["egressShapingPeakBandwidth"] = state ? state.egressShapingPeakBandwidth : undefined;
            inputs["failback"] = state ? state.failback : undefined;
            inputs["ingressShapingAverageBandwidth"] = state ? state.ingressShapingAverageBandwidth : undefined;
            inputs["ingressShapingBurstSize"] = state ? state.ingressShapingBurstSize : undefined;
            inputs["ingressShapingEnabled"] = state ? state.ingressShapingEnabled : undefined;
            inputs["ingressShapingPeakBandwidth"] = state ? state.ingressShapingPeakBandwidth : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["lacpEnabled"] = state ? state.lacpEnabled : undefined;
            inputs["lacpMode"] = state ? state.lacpMode : undefined;
            inputs["livePortMovingAllowed"] = state ? state.livePortMovingAllowed : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["netflowEnabled"] = state ? state.netflowEnabled : undefined;
            inputs["netflowOverrideAllowed"] = state ? state.netflowOverrideAllowed : undefined;
            inputs["networkResourcePoolKey"] = state ? state.networkResourcePoolKey : undefined;
            inputs["networkResourcePoolOverrideAllowed"] = state ? state.networkResourcePoolOverrideAllowed : undefined;
            inputs["notifySwitches"] = state ? state.notifySwitches : undefined;
            inputs["numberOfPorts"] = state ? state.numberOfPorts : undefined;
            inputs["portConfigResetAtDisconnect"] = state ? state.portConfigResetAtDisconnect : undefined;
            inputs["portNameFormat"] = state ? state.portNameFormat : undefined;
            inputs["portPrivateSecondaryVlanId"] = state ? state.portPrivateSecondaryVlanId : undefined;
            inputs["securityPolicyOverrideAllowed"] = state ? state.securityPolicyOverrideAllowed : undefined;
            inputs["shapingOverrideAllowed"] = state ? state.shapingOverrideAllowed : undefined;
            inputs["standbyUplinks"] = state ? state.standbyUplinks : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["teamingPolicy"] = state ? state.teamingPolicy : undefined;
            inputs["trafficFilterOverrideAllowed"] = state ? state.trafficFilterOverrideAllowed : undefined;
            inputs["txUplink"] = state ? state.txUplink : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["uplinkTeamingOverrideAllowed"] = state ? state.uplinkTeamingOverrideAllowed : undefined;
            inputs["vlanId"] = state ? state.vlanId : undefined;
            inputs["vlanOverrideAllowed"] = state ? state.vlanOverrideAllowed : undefined;
            inputs["vlanRanges"] = state ? state.vlanRanges : undefined;
        } else {
            const args = argsOrState as DistributedPortGroupArgs | undefined;
            if (!args || args.distributedVirtualSwitchUuid === undefined) {
                throw new Error("Missing required property 'distributedVirtualSwitchUuid'");
            }
            inputs["activeUplinks"] = args ? args.activeUplinks : undefined;
            inputs["allowForgedTransmits"] = args ? args.allowForgedTransmits : undefined;
            inputs["allowMacChanges"] = args ? args.allowMacChanges : undefined;
            inputs["allowPromiscuous"] = args ? args.allowPromiscuous : undefined;
            inputs["autoExpand"] = args ? args.autoExpand : undefined;
            inputs["blockAllPorts"] = args ? args.blockAllPorts : undefined;
            inputs["blockOverrideAllowed"] = args ? args.blockOverrideAllowed : undefined;
            inputs["checkBeacon"] = args ? args.checkBeacon : undefined;
            inputs["customAttributes"] = args ? args.customAttributes : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["directpathGen2Allowed"] = args ? args.directpathGen2Allowed : undefined;
            inputs["distributedVirtualSwitchUuid"] = args ? args.distributedVirtualSwitchUuid : undefined;
            inputs["egressShapingAverageBandwidth"] = args ? args.egressShapingAverageBandwidth : undefined;
            inputs["egressShapingBurstSize"] = args ? args.egressShapingBurstSize : undefined;
            inputs["egressShapingEnabled"] = args ? args.egressShapingEnabled : undefined;
            inputs["egressShapingPeakBandwidth"] = args ? args.egressShapingPeakBandwidth : undefined;
            inputs["failback"] = args ? args.failback : undefined;
            inputs["ingressShapingAverageBandwidth"] = args ? args.ingressShapingAverageBandwidth : undefined;
            inputs["ingressShapingBurstSize"] = args ? args.ingressShapingBurstSize : undefined;
            inputs["ingressShapingEnabled"] = args ? args.ingressShapingEnabled : undefined;
            inputs["ingressShapingPeakBandwidth"] = args ? args.ingressShapingPeakBandwidth : undefined;
            inputs["lacpEnabled"] = args ? args.lacpEnabled : undefined;
            inputs["lacpMode"] = args ? args.lacpMode : undefined;
            inputs["livePortMovingAllowed"] = args ? args.livePortMovingAllowed : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["netflowEnabled"] = args ? args.netflowEnabled : undefined;
            inputs["netflowOverrideAllowed"] = args ? args.netflowOverrideAllowed : undefined;
            inputs["networkResourcePoolKey"] = args ? args.networkResourcePoolKey : undefined;
            inputs["networkResourcePoolOverrideAllowed"] = args ? args.networkResourcePoolOverrideAllowed : undefined;
            inputs["notifySwitches"] = args ? args.notifySwitches : undefined;
            inputs["numberOfPorts"] = args ? args.numberOfPorts : undefined;
            inputs["portConfigResetAtDisconnect"] = args ? args.portConfigResetAtDisconnect : undefined;
            inputs["portNameFormat"] = args ? args.portNameFormat : undefined;
            inputs["portPrivateSecondaryVlanId"] = args ? args.portPrivateSecondaryVlanId : undefined;
            inputs["securityPolicyOverrideAllowed"] = args ? args.securityPolicyOverrideAllowed : undefined;
            inputs["shapingOverrideAllowed"] = args ? args.shapingOverrideAllowed : undefined;
            inputs["standbyUplinks"] = args ? args.standbyUplinks : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["teamingPolicy"] = args ? args.teamingPolicy : undefined;
            inputs["trafficFilterOverrideAllowed"] = args ? args.trafficFilterOverrideAllowed : undefined;
            inputs["txUplink"] = args ? args.txUplink : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["uplinkTeamingOverrideAllowed"] = args ? args.uplinkTeamingOverrideAllowed : undefined;
            inputs["vlanId"] = args ? args.vlanId : undefined;
            inputs["vlanOverrideAllowed"] = args ? args.vlanOverrideAllowed : undefined;
            inputs["vlanRanges"] = args ? args.vlanRanges : undefined;
            inputs["configVersion"] = undefined /*out*/;
            inputs["key"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DistributedPortGroup.__pulumiType, name, inputs, opts);
    }
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
     * Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
     * that of its own.
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
     * past the limit specified in `numberOfPorts` if necessary. Default: `true`.
     */
    readonly autoExpand?: pulumi.Input<boolean>;
    /**
     * Indicates whether to block all ports by default.
     */
    readonly blockAllPorts?: pulumi.Input<boolean>;
    /**
     * Allow the port shutdown
     * policy to be overridden on an individual port.
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
     * value string to set for port group.
     */
    readonly customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
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
     * Allow the Netflow
     * policy on this port group to be overridden on an individual
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
     * Allow the security policy
     * settings defined in this port group policy to be
     * overridden on an individual port.
     */
    readonly securityPolicyOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * Allow the traffic shaping
     * options on this port group policy to be overridden
     * on an individual port.
     */
    readonly shapingOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * List of standby uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
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
     * If true, a copy of packets sent to the switch will always be forwarded to an uplink in addition to the regular packet
     * forwarded done by the switch.
     */
    readonly txUplink?: pulumi.Input<boolean>;
    /**
     * The port group type. Can be one of `earlyBinding` (static
     * binding) or `ephemeral`. Default: `earlyBinding`.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Allow the uplink teaming
     * options on this port group to be overridden on an
     * individual port.
     */
    readonly uplinkTeamingOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * The VLAN ID for single VLAN mode. 0 denotes no VLAN.
     */
    readonly vlanId?: pulumi.Input<number>;
    /**
     * Allow the VLAN settings
     * on this port group to be overridden on an individual port.
     */
    readonly vlanOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * The VLAN ID for single VLAN mode. 0 denotes no VLAN.
     */
    readonly vlanRanges?: pulumi.Input<pulumi.Input<inputs.DistributedPortGroupVlanRange>[]>;
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
     * Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
     * that of its own.
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
     * past the limit specified in `numberOfPorts` if necessary. Default: `true`.
     */
    readonly autoExpand?: pulumi.Input<boolean>;
    /**
     * Indicates whether to block all ports by default.
     */
    readonly blockAllPorts?: pulumi.Input<boolean>;
    /**
     * Allow the port shutdown
     * policy to be overridden on an individual port.
     */
    readonly blockOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * Enable beacon probing on the ports this policy applies to.
     */
    readonly checkBeacon?: pulumi.Input<boolean>;
    /**
     * Map of custom attribute ids to attribute
     * value string to set for port group.
     */
    readonly customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
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
     * Allow the Netflow
     * policy on this port group to be overridden on an individual
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
     * Allow the security policy
     * settings defined in this port group policy to be
     * overridden on an individual port.
     */
    readonly securityPolicyOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * Allow the traffic shaping
     * options on this port group policy to be overridden
     * on an individual port.
     */
    readonly shapingOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * List of standby uplinks used for load balancing, matching the names of the uplinks assigned in the DVS.
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
     * If true, a copy of packets sent to the switch will always be forwarded to an uplink in addition to the regular packet
     * forwarded done by the switch.
     */
    readonly txUplink?: pulumi.Input<boolean>;
    /**
     * The port group type. Can be one of `earlyBinding` (static
     * binding) or `ephemeral`. Default: `earlyBinding`.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Allow the uplink teaming
     * options on this port group to be overridden on an
     * individual port.
     */
    readonly uplinkTeamingOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * The VLAN ID for single VLAN mode. 0 denotes no VLAN.
     */
    readonly vlanId?: pulumi.Input<number>;
    /**
     * Allow the VLAN settings
     * on this port group to be overridden on an individual port.
     */
    readonly vlanOverrideAllowed?: pulumi.Input<boolean>;
    /**
     * The VLAN ID for single VLAN mode. 0 denotes no VLAN.
     */
    readonly vlanRanges?: pulumi.Input<pulumi.Input<inputs.DistributedPortGroupVlanRange>[]>;
}

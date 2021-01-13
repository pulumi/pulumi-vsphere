// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class HaVmOverride extends pulumi.CustomResource {
    /**
     * Get an existing HaVmOverride resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HaVmOverrideState, opts?: pulumi.CustomResourceOptions): HaVmOverride {
        return new HaVmOverride(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/haVmOverride:HaVmOverride';

    /**
     * Returns true if the given object is an instance of HaVmOverride.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HaVmOverride {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HaVmOverride.__pulumiType;
    }

    /**
     * The managed object reference
     * ID of the cluster to put the override in.  Forces a new
     * resource if changed.
     */
    public readonly computeClusterId!: pulumi.Output<string>;
    /**
     * Controls the action to take
     * on this virtual machine if an APD status on an affected datastore clears in
     * the middle of an APD event. Can be one of `useClusterDefault`, `none` or
     * `reset`.  Default: `useClusterDefault`.
     */
    public readonly haDatastoreApdRecoveryAction!: pulumi.Output<string | undefined>;
    /**
     * Controls the action to take on this
     * virtual machine when the cluster has detected loss to all paths to a relevant
     * datastore. Can be one of `clusterDefault`, `disabled`, `warning`,
     * `restartConservative`, or `restartAggressive`.  Default: `clusterDefault`.
     */
    public readonly haDatastoreApdResponse!: pulumi.Output<string | undefined>;
    /**
     * Controls the delay in minutes
     * to wait after an APD timeout event to execute the response action defined in
     * `haDatastoreApdResponse`. Use `-1` to use
     * the cluster default. Default: `-1`.
     */
    public readonly haDatastoreApdResponseDelay!: pulumi.Output<number | undefined>;
    /**
     * Controls the action to take on this
     * virtual machine when the cluster has detected a permanent device loss to a
     * relevant datastore. Can be one of `clusterDefault`, `disabled`, `warning`, or
     * `restartAggressive`. Default: `clusterDefault`.
     */
    public readonly haDatastorePdlResponse!: pulumi.Output<string | undefined>;
    /**
     * The action to take on this virtual
     * machine when a host has detected that it has been isolated from the rest of
     * the cluster. Can be one of `clusterIsolationResponse`, `none`, `powerOff`, or
     * `shutdown`. Default: `clusterIsolationResponse`.
     */
    public readonly haHostIsolationResponse!: pulumi.Output<string | undefined>;
    /**
     * If a heartbeat from this virtual
     * machine is not received within this configured interval, the virtual machine
     * is marked as failed. The value is in seconds. Default: `30`.
     */
    public readonly haVmFailureInterval!: pulumi.Output<number | undefined>;
    /**
     * The length of the reset window in
     * which `haVmMaximumResets` can operate. When this
     * window expires, no more resets are attempted regardless of the setting
     * configured in `haVmMaximumResets`. `-1` means no window, meaning an
     * unlimited reset time is allotted. The value is specified in seconds. Default:
     * `-1` (no window).
     */
    public readonly haVmMaximumFailureWindow!: pulumi.Output<number | undefined>;
    /**
     * The maximum number of resets that HA will
     * perform to this virtual machine when responding to a failure event. Default:
     * `3`
     */
    public readonly haVmMaximumResets!: pulumi.Output<number | undefined>;
    /**
     * The time, in seconds, that HA waits after
     * powering on this virtual machine before monitoring for heartbeats. Default:
     * `120` (2 minutes).
     */
    public readonly haVmMinimumUptime!: pulumi.Output<number | undefined>;
    /**
     * The type of virtual machine monitoring to use
     * when HA is enabled in the cluster. Can be one of `vmMonitoringDisabled`,
     * `vmMonitoringOnly`, or `vmAndAppMonitoring`. Default: `vmMonitoringDisabled`.
     */
    public readonly haVmMonitoring!: pulumi.Output<string | undefined>;
    /**
     * Determines whether or
     * not the cluster's default settings or the VM override settings specified in
     * this resource are used for virtual machine monitoring. The default is `true`
     * (use cluster defaults) - set to `false` to have overrides take effect.
     */
    public readonly haVmMonitoringUseClusterDefaults!: pulumi.Output<boolean | undefined>;
    /**
     * The restart priority for the virtual
     * machine when vSphere detects a host failure. Can be one of
     * `clusterRestartPriority`, `lowest`, `low`, `medium`, `high`, or `highest`.
     * Default: `clusterRestartPriority`.
     */
    public readonly haVmRestartPriority!: pulumi.Output<string | undefined>;
    /**
     * The maximum time, in seconds, that
     * vSphere HA will wait for this virtual machine to be ready. Use `-1` to
     * specify the cluster default.  Default: `-1`.
     */
    public readonly haVmRestartTimeout!: pulumi.Output<number | undefined>;
    /**
     * The UUID of the virtual machine to create
     * the override for.  Forces a new resource if changed.
     */
    public readonly virtualMachineId!: pulumi.Output<string>;

    /**
     * Create a HaVmOverride resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HaVmOverrideArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HaVmOverrideArgs | HaVmOverrideState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as HaVmOverrideState | undefined;
            inputs["computeClusterId"] = state ? state.computeClusterId : undefined;
            inputs["haDatastoreApdRecoveryAction"] = state ? state.haDatastoreApdRecoveryAction : undefined;
            inputs["haDatastoreApdResponse"] = state ? state.haDatastoreApdResponse : undefined;
            inputs["haDatastoreApdResponseDelay"] = state ? state.haDatastoreApdResponseDelay : undefined;
            inputs["haDatastorePdlResponse"] = state ? state.haDatastorePdlResponse : undefined;
            inputs["haHostIsolationResponse"] = state ? state.haHostIsolationResponse : undefined;
            inputs["haVmFailureInterval"] = state ? state.haVmFailureInterval : undefined;
            inputs["haVmMaximumFailureWindow"] = state ? state.haVmMaximumFailureWindow : undefined;
            inputs["haVmMaximumResets"] = state ? state.haVmMaximumResets : undefined;
            inputs["haVmMinimumUptime"] = state ? state.haVmMinimumUptime : undefined;
            inputs["haVmMonitoring"] = state ? state.haVmMonitoring : undefined;
            inputs["haVmMonitoringUseClusterDefaults"] = state ? state.haVmMonitoringUseClusterDefaults : undefined;
            inputs["haVmRestartPriority"] = state ? state.haVmRestartPriority : undefined;
            inputs["haVmRestartTimeout"] = state ? state.haVmRestartTimeout : undefined;
            inputs["virtualMachineId"] = state ? state.virtualMachineId : undefined;
        } else {
            const args = argsOrState as HaVmOverrideArgs | undefined;
            if ((!args || args.computeClusterId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'computeClusterId'");
            }
            if ((!args || args.virtualMachineId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'virtualMachineId'");
            }
            inputs["computeClusterId"] = args ? args.computeClusterId : undefined;
            inputs["haDatastoreApdRecoveryAction"] = args ? args.haDatastoreApdRecoveryAction : undefined;
            inputs["haDatastoreApdResponse"] = args ? args.haDatastoreApdResponse : undefined;
            inputs["haDatastoreApdResponseDelay"] = args ? args.haDatastoreApdResponseDelay : undefined;
            inputs["haDatastorePdlResponse"] = args ? args.haDatastorePdlResponse : undefined;
            inputs["haHostIsolationResponse"] = args ? args.haHostIsolationResponse : undefined;
            inputs["haVmFailureInterval"] = args ? args.haVmFailureInterval : undefined;
            inputs["haVmMaximumFailureWindow"] = args ? args.haVmMaximumFailureWindow : undefined;
            inputs["haVmMaximumResets"] = args ? args.haVmMaximumResets : undefined;
            inputs["haVmMinimumUptime"] = args ? args.haVmMinimumUptime : undefined;
            inputs["haVmMonitoring"] = args ? args.haVmMonitoring : undefined;
            inputs["haVmMonitoringUseClusterDefaults"] = args ? args.haVmMonitoringUseClusterDefaults : undefined;
            inputs["haVmRestartPriority"] = args ? args.haVmRestartPriority : undefined;
            inputs["haVmRestartTimeout"] = args ? args.haVmRestartTimeout : undefined;
            inputs["virtualMachineId"] = args ? args.virtualMachineId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(HaVmOverride.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HaVmOverride resources.
 */
export interface HaVmOverrideState {
    /**
     * The managed object reference
     * ID of the cluster to put the override in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId?: pulumi.Input<string>;
    /**
     * Controls the action to take
     * on this virtual machine if an APD status on an affected datastore clears in
     * the middle of an APD event. Can be one of `useClusterDefault`, `none` or
     * `reset`.  Default: `useClusterDefault`.
     */
    readonly haDatastoreApdRecoveryAction?: pulumi.Input<string>;
    /**
     * Controls the action to take on this
     * virtual machine when the cluster has detected loss to all paths to a relevant
     * datastore. Can be one of `clusterDefault`, `disabled`, `warning`,
     * `restartConservative`, or `restartAggressive`.  Default: `clusterDefault`.
     */
    readonly haDatastoreApdResponse?: pulumi.Input<string>;
    /**
     * Controls the delay in minutes
     * to wait after an APD timeout event to execute the response action defined in
     * `haDatastoreApdResponse`. Use `-1` to use
     * the cluster default. Default: `-1`.
     */
    readonly haDatastoreApdResponseDelay?: pulumi.Input<number>;
    /**
     * Controls the action to take on this
     * virtual machine when the cluster has detected a permanent device loss to a
     * relevant datastore. Can be one of `clusterDefault`, `disabled`, `warning`, or
     * `restartAggressive`. Default: `clusterDefault`.
     */
    readonly haDatastorePdlResponse?: pulumi.Input<string>;
    /**
     * The action to take on this virtual
     * machine when a host has detected that it has been isolated from the rest of
     * the cluster. Can be one of `clusterIsolationResponse`, `none`, `powerOff`, or
     * `shutdown`. Default: `clusterIsolationResponse`.
     */
    readonly haHostIsolationResponse?: pulumi.Input<string>;
    /**
     * If a heartbeat from this virtual
     * machine is not received within this configured interval, the virtual machine
     * is marked as failed. The value is in seconds. Default: `30`.
     */
    readonly haVmFailureInterval?: pulumi.Input<number>;
    /**
     * The length of the reset window in
     * which `haVmMaximumResets` can operate. When this
     * window expires, no more resets are attempted regardless of the setting
     * configured in `haVmMaximumResets`. `-1` means no window, meaning an
     * unlimited reset time is allotted. The value is specified in seconds. Default:
     * `-1` (no window).
     */
    readonly haVmMaximumFailureWindow?: pulumi.Input<number>;
    /**
     * The maximum number of resets that HA will
     * perform to this virtual machine when responding to a failure event. Default:
     * `3`
     */
    readonly haVmMaximumResets?: pulumi.Input<number>;
    /**
     * The time, in seconds, that HA waits after
     * powering on this virtual machine before monitoring for heartbeats. Default:
     * `120` (2 minutes).
     */
    readonly haVmMinimumUptime?: pulumi.Input<number>;
    /**
     * The type of virtual machine monitoring to use
     * when HA is enabled in the cluster. Can be one of `vmMonitoringDisabled`,
     * `vmMonitoringOnly`, or `vmAndAppMonitoring`. Default: `vmMonitoringDisabled`.
     */
    readonly haVmMonitoring?: pulumi.Input<string>;
    /**
     * Determines whether or
     * not the cluster's default settings or the VM override settings specified in
     * this resource are used for virtual machine monitoring. The default is `true`
     * (use cluster defaults) - set to `false` to have overrides take effect.
     */
    readonly haVmMonitoringUseClusterDefaults?: pulumi.Input<boolean>;
    /**
     * The restart priority for the virtual
     * machine when vSphere detects a host failure. Can be one of
     * `clusterRestartPriority`, `lowest`, `low`, `medium`, `high`, or `highest`.
     * Default: `clusterRestartPriority`.
     */
    readonly haVmRestartPriority?: pulumi.Input<string>;
    /**
     * The maximum time, in seconds, that
     * vSphere HA will wait for this virtual machine to be ready. Use `-1` to
     * specify the cluster default.  Default: `-1`.
     */
    readonly haVmRestartTimeout?: pulumi.Input<number>;
    /**
     * The UUID of the virtual machine to create
     * the override for.  Forces a new resource if changed.
     */
    readonly virtualMachineId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HaVmOverride resource.
 */
export interface HaVmOverrideArgs {
    /**
     * The managed object reference
     * ID of the cluster to put the override in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId: pulumi.Input<string>;
    /**
     * Controls the action to take
     * on this virtual machine if an APD status on an affected datastore clears in
     * the middle of an APD event. Can be one of `useClusterDefault`, `none` or
     * `reset`.  Default: `useClusterDefault`.
     */
    readonly haDatastoreApdRecoveryAction?: pulumi.Input<string>;
    /**
     * Controls the action to take on this
     * virtual machine when the cluster has detected loss to all paths to a relevant
     * datastore. Can be one of `clusterDefault`, `disabled`, `warning`,
     * `restartConservative`, or `restartAggressive`.  Default: `clusterDefault`.
     */
    readonly haDatastoreApdResponse?: pulumi.Input<string>;
    /**
     * Controls the delay in minutes
     * to wait after an APD timeout event to execute the response action defined in
     * `haDatastoreApdResponse`. Use `-1` to use
     * the cluster default. Default: `-1`.
     */
    readonly haDatastoreApdResponseDelay?: pulumi.Input<number>;
    /**
     * Controls the action to take on this
     * virtual machine when the cluster has detected a permanent device loss to a
     * relevant datastore. Can be one of `clusterDefault`, `disabled`, `warning`, or
     * `restartAggressive`. Default: `clusterDefault`.
     */
    readonly haDatastorePdlResponse?: pulumi.Input<string>;
    /**
     * The action to take on this virtual
     * machine when a host has detected that it has been isolated from the rest of
     * the cluster. Can be one of `clusterIsolationResponse`, `none`, `powerOff`, or
     * `shutdown`. Default: `clusterIsolationResponse`.
     */
    readonly haHostIsolationResponse?: pulumi.Input<string>;
    /**
     * If a heartbeat from this virtual
     * machine is not received within this configured interval, the virtual machine
     * is marked as failed. The value is in seconds. Default: `30`.
     */
    readonly haVmFailureInterval?: pulumi.Input<number>;
    /**
     * The length of the reset window in
     * which `haVmMaximumResets` can operate. When this
     * window expires, no more resets are attempted regardless of the setting
     * configured in `haVmMaximumResets`. `-1` means no window, meaning an
     * unlimited reset time is allotted. The value is specified in seconds. Default:
     * `-1` (no window).
     */
    readonly haVmMaximumFailureWindow?: pulumi.Input<number>;
    /**
     * The maximum number of resets that HA will
     * perform to this virtual machine when responding to a failure event. Default:
     * `3`
     */
    readonly haVmMaximumResets?: pulumi.Input<number>;
    /**
     * The time, in seconds, that HA waits after
     * powering on this virtual machine before monitoring for heartbeats. Default:
     * `120` (2 minutes).
     */
    readonly haVmMinimumUptime?: pulumi.Input<number>;
    /**
     * The type of virtual machine monitoring to use
     * when HA is enabled in the cluster. Can be one of `vmMonitoringDisabled`,
     * `vmMonitoringOnly`, or `vmAndAppMonitoring`. Default: `vmMonitoringDisabled`.
     */
    readonly haVmMonitoring?: pulumi.Input<string>;
    /**
     * Determines whether or
     * not the cluster's default settings or the VM override settings specified in
     * this resource are used for virtual machine monitoring. The default is `true`
     * (use cluster defaults) - set to `false` to have overrides take effect.
     */
    readonly haVmMonitoringUseClusterDefaults?: pulumi.Input<boolean>;
    /**
     * The restart priority for the virtual
     * machine when vSphere detects a host failure. Can be one of
     * `clusterRestartPriority`, `lowest`, `low`, `medium`, `high`, or `highest`.
     * Default: `clusterRestartPriority`.
     */
    readonly haVmRestartPriority?: pulumi.Input<string>;
    /**
     * The maximum time, in seconds, that
     * vSphere HA will wait for this virtual machine to be ready. Use `-1` to
     * specify the cluster default.  Default: `-1`.
     */
    readonly haVmRestartTimeout?: pulumi.Input<number>;
    /**
     * The UUID of the virtual machine to create
     * the override for.  Forces a new resource if changed.
     */
    readonly virtualMachineId: pulumi.Input<string>;
}

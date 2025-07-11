// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.HaVmOverride` resource can be used to add an override for
 * vSphere HA settings on a cluster for a specific virtual machine. With this
 * resource, one can control specific HA settings so that they are different than
 * the cluster default, accommodating the needs of that specific virtual machine,
 * while not affecting the rest of the cluster.
 *
 * For more information on vSphere HA, see [this page][ref-vsphere-ha-clusters].
 *
 * [ref-vsphere-ha-clusters]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-availability.html
 *
 * > **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 *
 * ## Example Usage
 *
 * The example below creates a virtual machine in a cluster using the
 * `vsphere.VirtualMachine` resource, creating the
 * virtual machine in the cluster looked up by the
 * `vsphere.ComputeCluster` data source.
 *
 * Considering a scenario where this virtual machine is of high value to the
 * application or organization for which it does its work, it's been determined in
 * the event of a host failure, that this should be one of the first virtual
 * machines to be started by vSphere HA during recovery. Hence, it
 * `haVmRestartPriority` has been set to `highest`,
 * which, assuming that the default restart priority is `medium` and no other
 * virtual machine has been assigned the `highest` priority, will mean that this
 * VM will be started before any other virtual machine in the event of host
 * failure.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const datastore = datacenter.then(datacenter => vsphere.getDatastore({
 *     name: "datastore1",
 *     datacenterId: datacenter.id,
 * }));
 * const cluster = datacenter.then(datacenter => vsphere.getComputeCluster({
 *     name: "cluster-01",
 *     datacenterId: datacenter.id,
 * }));
 * const network = datacenter.then(datacenter => vsphere.getNetwork({
 *     name: "network1",
 *     datacenterId: datacenter.id,
 * }));
 * const vm = new vsphere.VirtualMachine("vm", {
 *     name: "test",
 *     resourcePoolId: cluster.then(cluster => cluster.resourcePoolId),
 *     datastoreId: datastore.then(datastore => datastore.id),
 *     numCpus: 2,
 *     memory: 2048,
 *     guestId: "otherLinux64Guest",
 *     networkInterfaces: [{
 *         networkId: network.then(network => network.id),
 *     }],
 *     disks: [{
 *         label: "disk0",
 *         size: 20,
 *     }],
 * });
 * const haVmOverride = new vsphere.HaVmOverride("ha_vm_override", {
 *     computeClusterId: cluster.then(cluster => cluster.id),
 *     virtualMachineId: vm.id,
 *     haVmRestartPriority: "highest",
 * });
 * ```
 *
 * ## Import
 *
 * An existing override can be imported into this resource by
 *
 * supplying both the path to the cluster, and the path to the virtual machine, to
 *
 * `pulumi import`. If no override exists, an error will be given.  An example
 *
 * is below:
 *
 * [docs-import]: https://developer.hashicorp.com/terraform/cli/import
 *
 * ```sh
 * $ pulumi import vsphere:index/haVmOverride:HaVmOverride ha_vm_override \
 * ```
 *
 *   '{"compute_cluster_path": "/dc1/host/cluster1", \
 *
 *   "virtual_machine_path": "/dc1/vm/srv1"}'
 */
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
     * The managed object ID of the cluster.
     */
    public readonly computeClusterId!: pulumi.Output<string>;
    /**
     * Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an
     * APD event. Can be one of useClusterDefault, none or reset.
     */
    public readonly haDatastoreApdRecoveryAction!: pulumi.Output<string | undefined>;
    /**
     * Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant
     * datastore. Can be one of clusterDefault, disabled, warning, restartConservative, or restartAggressive.
     */
    public readonly haDatastoreApdResponse!: pulumi.Output<string | undefined>;
    /**
     * Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in
     * ha_datastore_apd_response. Specify -1 to use the cluster setting.
     */
    public readonly haDatastoreApdResponseDelay!: pulumi.Output<number | undefined>;
    /**
     * Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant
     * datastore. Can be one of clusterDefault, disabled, warning, or restartAggressive.
     */
    public readonly haDatastorePdlResponse!: pulumi.Output<string | undefined>;
    /**
     * The action to take on this virtual machine when a host is isolated from the rest of the cluster. Can be one of
     * clusterIsolationResponse, none, powerOff, or shutdown.
     */
    public readonly haHostIsolationResponse!: pulumi.Output<string | undefined>;
    /**
     * If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked
     * as failed. The value is in seconds.
     */
    public readonly haVmFailureInterval!: pulumi.Output<number | undefined>;
    /**
     * The length of the reset window in which haVmMaximumResets can operate. When this window expires, no more resets are
     * attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset
     * time is allotted.
     */
    public readonly haVmMaximumFailureWindow!: pulumi.Output<number | undefined>;
    /**
     * The maximum number of resets that HA will perform to this virtual machine when responding to a failure event.
     */
    public readonly haVmMaximumResets!: pulumi.Output<number | undefined>;
    /**
     * The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats.
     */
    public readonly haVmMinimumUptime!: pulumi.Output<number | undefined>;
    /**
     * The type of virtual machine monitoring to use for this virtual machine. Can be one of vmMonitoringDisabled,
     * vmMonitoringOnly, or vmAndAppMonitoring.
     */
    public readonly haVmMonitoring!: pulumi.Output<string | undefined>;
    /**
     * Determines whether or not the cluster's default settings or the VM override settings specified in this resource are used
     * for virtual machine monitoring. The default is true (use cluster defaults) - set to false to have overrides take effect.
     */
    public readonly haVmMonitoringUseClusterDefaults!: pulumi.Output<boolean | undefined>;
    /**
     * The restart priority for this virtual machine when vSphere detects a host failure. Can be one of clusterRestartPriority,
     * lowest, low, medium, high, or highest.
     */
    public readonly haVmRestartPriority!: pulumi.Output<string | undefined>;
    /**
     * The maximum time, in seconds, that vSphere HA will wait for the virtual machine to be ready. Use -1 to use the cluster
     * default.
     */
    public readonly haVmRestartTimeout!: pulumi.Output<number | undefined>;
    /**
     * The managed object ID of the virtual machine.
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HaVmOverrideState | undefined;
            resourceInputs["computeClusterId"] = state ? state.computeClusterId : undefined;
            resourceInputs["haDatastoreApdRecoveryAction"] = state ? state.haDatastoreApdRecoveryAction : undefined;
            resourceInputs["haDatastoreApdResponse"] = state ? state.haDatastoreApdResponse : undefined;
            resourceInputs["haDatastoreApdResponseDelay"] = state ? state.haDatastoreApdResponseDelay : undefined;
            resourceInputs["haDatastorePdlResponse"] = state ? state.haDatastorePdlResponse : undefined;
            resourceInputs["haHostIsolationResponse"] = state ? state.haHostIsolationResponse : undefined;
            resourceInputs["haVmFailureInterval"] = state ? state.haVmFailureInterval : undefined;
            resourceInputs["haVmMaximumFailureWindow"] = state ? state.haVmMaximumFailureWindow : undefined;
            resourceInputs["haVmMaximumResets"] = state ? state.haVmMaximumResets : undefined;
            resourceInputs["haVmMinimumUptime"] = state ? state.haVmMinimumUptime : undefined;
            resourceInputs["haVmMonitoring"] = state ? state.haVmMonitoring : undefined;
            resourceInputs["haVmMonitoringUseClusterDefaults"] = state ? state.haVmMonitoringUseClusterDefaults : undefined;
            resourceInputs["haVmRestartPriority"] = state ? state.haVmRestartPriority : undefined;
            resourceInputs["haVmRestartTimeout"] = state ? state.haVmRestartTimeout : undefined;
            resourceInputs["virtualMachineId"] = state ? state.virtualMachineId : undefined;
        } else {
            const args = argsOrState as HaVmOverrideArgs | undefined;
            if ((!args || args.computeClusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'computeClusterId'");
            }
            if ((!args || args.virtualMachineId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualMachineId'");
            }
            resourceInputs["computeClusterId"] = args ? args.computeClusterId : undefined;
            resourceInputs["haDatastoreApdRecoveryAction"] = args ? args.haDatastoreApdRecoveryAction : undefined;
            resourceInputs["haDatastoreApdResponse"] = args ? args.haDatastoreApdResponse : undefined;
            resourceInputs["haDatastoreApdResponseDelay"] = args ? args.haDatastoreApdResponseDelay : undefined;
            resourceInputs["haDatastorePdlResponse"] = args ? args.haDatastorePdlResponse : undefined;
            resourceInputs["haHostIsolationResponse"] = args ? args.haHostIsolationResponse : undefined;
            resourceInputs["haVmFailureInterval"] = args ? args.haVmFailureInterval : undefined;
            resourceInputs["haVmMaximumFailureWindow"] = args ? args.haVmMaximumFailureWindow : undefined;
            resourceInputs["haVmMaximumResets"] = args ? args.haVmMaximumResets : undefined;
            resourceInputs["haVmMinimumUptime"] = args ? args.haVmMinimumUptime : undefined;
            resourceInputs["haVmMonitoring"] = args ? args.haVmMonitoring : undefined;
            resourceInputs["haVmMonitoringUseClusterDefaults"] = args ? args.haVmMonitoringUseClusterDefaults : undefined;
            resourceInputs["haVmRestartPriority"] = args ? args.haVmRestartPriority : undefined;
            resourceInputs["haVmRestartTimeout"] = args ? args.haVmRestartTimeout : undefined;
            resourceInputs["virtualMachineId"] = args ? args.virtualMachineId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HaVmOverride.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HaVmOverride resources.
 */
export interface HaVmOverrideState {
    /**
     * The managed object ID of the cluster.
     */
    computeClusterId?: pulumi.Input<string>;
    /**
     * Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an
     * APD event. Can be one of useClusterDefault, none or reset.
     */
    haDatastoreApdRecoveryAction?: pulumi.Input<string>;
    /**
     * Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant
     * datastore. Can be one of clusterDefault, disabled, warning, restartConservative, or restartAggressive.
     */
    haDatastoreApdResponse?: pulumi.Input<string>;
    /**
     * Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in
     * ha_datastore_apd_response. Specify -1 to use the cluster setting.
     */
    haDatastoreApdResponseDelay?: pulumi.Input<number>;
    /**
     * Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant
     * datastore. Can be one of clusterDefault, disabled, warning, or restartAggressive.
     */
    haDatastorePdlResponse?: pulumi.Input<string>;
    /**
     * The action to take on this virtual machine when a host is isolated from the rest of the cluster. Can be one of
     * clusterIsolationResponse, none, powerOff, or shutdown.
     */
    haHostIsolationResponse?: pulumi.Input<string>;
    /**
     * If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked
     * as failed. The value is in seconds.
     */
    haVmFailureInterval?: pulumi.Input<number>;
    /**
     * The length of the reset window in which haVmMaximumResets can operate. When this window expires, no more resets are
     * attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset
     * time is allotted.
     */
    haVmMaximumFailureWindow?: pulumi.Input<number>;
    /**
     * The maximum number of resets that HA will perform to this virtual machine when responding to a failure event.
     */
    haVmMaximumResets?: pulumi.Input<number>;
    /**
     * The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats.
     */
    haVmMinimumUptime?: pulumi.Input<number>;
    /**
     * The type of virtual machine monitoring to use for this virtual machine. Can be one of vmMonitoringDisabled,
     * vmMonitoringOnly, or vmAndAppMonitoring.
     */
    haVmMonitoring?: pulumi.Input<string>;
    /**
     * Determines whether or not the cluster's default settings or the VM override settings specified in this resource are used
     * for virtual machine monitoring. The default is true (use cluster defaults) - set to false to have overrides take effect.
     */
    haVmMonitoringUseClusterDefaults?: pulumi.Input<boolean>;
    /**
     * The restart priority for this virtual machine when vSphere detects a host failure. Can be one of clusterRestartPriority,
     * lowest, low, medium, high, or highest.
     */
    haVmRestartPriority?: pulumi.Input<string>;
    /**
     * The maximum time, in seconds, that vSphere HA will wait for the virtual machine to be ready. Use -1 to use the cluster
     * default.
     */
    haVmRestartTimeout?: pulumi.Input<number>;
    /**
     * The managed object ID of the virtual machine.
     */
    virtualMachineId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HaVmOverride resource.
 */
export interface HaVmOverrideArgs {
    /**
     * The managed object ID of the cluster.
     */
    computeClusterId: pulumi.Input<string>;
    /**
     * Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an
     * APD event. Can be one of useClusterDefault, none or reset.
     */
    haDatastoreApdRecoveryAction?: pulumi.Input<string>;
    /**
     * Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant
     * datastore. Can be one of clusterDefault, disabled, warning, restartConservative, or restartAggressive.
     */
    haDatastoreApdResponse?: pulumi.Input<string>;
    /**
     * Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in
     * ha_datastore_apd_response. Specify -1 to use the cluster setting.
     */
    haDatastoreApdResponseDelay?: pulumi.Input<number>;
    /**
     * Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant
     * datastore. Can be one of clusterDefault, disabled, warning, or restartAggressive.
     */
    haDatastorePdlResponse?: pulumi.Input<string>;
    /**
     * The action to take on this virtual machine when a host is isolated from the rest of the cluster. Can be one of
     * clusterIsolationResponse, none, powerOff, or shutdown.
     */
    haHostIsolationResponse?: pulumi.Input<string>;
    /**
     * If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked
     * as failed. The value is in seconds.
     */
    haVmFailureInterval?: pulumi.Input<number>;
    /**
     * The length of the reset window in which haVmMaximumResets can operate. When this window expires, no more resets are
     * attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset
     * time is allotted.
     */
    haVmMaximumFailureWindow?: pulumi.Input<number>;
    /**
     * The maximum number of resets that HA will perform to this virtual machine when responding to a failure event.
     */
    haVmMaximumResets?: pulumi.Input<number>;
    /**
     * The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats.
     */
    haVmMinimumUptime?: pulumi.Input<number>;
    /**
     * The type of virtual machine monitoring to use for this virtual machine. Can be one of vmMonitoringDisabled,
     * vmMonitoringOnly, or vmAndAppMonitoring.
     */
    haVmMonitoring?: pulumi.Input<string>;
    /**
     * Determines whether or not the cluster's default settings or the VM override settings specified in this resource are used
     * for virtual machine monitoring. The default is true (use cluster defaults) - set to false to have overrides take effect.
     */
    haVmMonitoringUseClusterDefaults?: pulumi.Input<boolean>;
    /**
     * The restart priority for this virtual machine when vSphere detects a host failure. Can be one of clusterRestartPriority,
     * lowest, low, medium, high, or highest.
     */
    haVmRestartPriority?: pulumi.Input<string>;
    /**
     * The maximum time, in seconds, that vSphere HA will wait for the virtual machine to be ready. Use -1 to use the cluster
     * default.
     */
    haVmRestartTimeout?: pulumi.Input<number>;
    /**
     * The managed object ID of the virtual machine.
     */
    virtualMachineId: pulumi.Input<string>;
}

import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_ha_vm_override` resource can be used to add an override for
 * vSphere HA settings on a cluster for a specific virtual machine. With this
 * resource, one can control specific HA settings so that they are different than
 * the cluster default, accommodating the needs of that specific virtual machine,
 * while not affecting the rest of the cluster.
 *
 * For more information on vSphere HA, see [this page][ref-vsphere-ha-clusters].
 *
 * [ref-vsphere-ha-clusters]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.avail.doc/GUID-5432CA24-14F1-44E3-87FB-61D937831CF6.html
 *
 * ~> **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 */
export declare class HaVmOverride extends pulumi.CustomResource {
    /**
     * Get an existing HaVmOverride resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HaVmOverrideState): HaVmOverride;
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId: pulumi.Output<string>;
    /**
     * Controls the action to take
     * on this virtual machine if an APD status on an affected datastore clears in
     * the middle of an APD event. Can be one of `useClusterDefault`, `none` or
     * `reset`.  Default: `useClusterDefault`.
     * <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
     */
    readonly haDatastoreApdRecoveryAction: pulumi.Output<string | undefined>;
    /**
     * Controls the action to take on this
     * virtual machine when the cluster has detected loss to all paths to a relevant
     * datastore. Can be one of `clusterDefault`, `disabled`, `warning`,
     * `restartConservative`, or `restartAggressive`.  Default: `clusterDefault`.
     * <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
     */
    readonly haDatastoreApdResponse: pulumi.Output<string | undefined>;
    /**
     * Controls the delay in minutes
     * to wait after an APD timeout event to execute the response action defined in
     * `ha_datastore_apd_response`. Use `-1` to use
     * the cluster default. Default: `-1`.
     * <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
     */
    readonly haDatastoreApdResponseDelay: pulumi.Output<number | undefined>;
    /**
     * Controls the action to take on this
     * virtual machine when the cluster has detected a permanent device loss to a
     * relevant datastore. Can be one of `clusterDefault`, `disabled`, `warning`, or
     * `restartAggressive`. Default: `clusterDefault`.
     * <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
     */
    readonly haDatastorePdlResponse: pulumi.Output<string | undefined>;
    /**
     * The action to take on this virtual
     * machine when a host has detected that it has been isolated from the rest of
     * the cluster. Can be one of `clusterIsolationResponse`, `none`, `powerOff`, or
     * `shutdown`. Default: `clusterIsolationResponse`.
     */
    readonly haHostIsolationResponse: pulumi.Output<string | undefined>;
    /**
     * If a heartbeat from this virtual
     * machine is not received within this configured interval, the virtual machine
     * is marked as failed. The value is in seconds. Default: `30`.
     */
    readonly haVmFailureInterval: pulumi.Output<number | undefined>;
    /**
     * The length of the reset window in
     * which `ha_vm_maximum_resets` can operate. When this
     * window expires, no more resets are attempted regardless of the setting
     * configured in `ha_vm_maximum_resets`. `-1` means no window, meaning an
     * unlimited reset time is allotted. The value is specified in seconds. Default:
     * `-1` (no window).
     */
    readonly haVmMaximumFailureWindow: pulumi.Output<number | undefined>;
    /**
     * The maximum number of resets that HA will
     * perform to this virtual machine when responding to a failure event. Default:
     * `3`
     */
    readonly haVmMaximumResets: pulumi.Output<number | undefined>;
    /**
     * The time, in seconds, that HA waits after
     * powering on this virtual machine before monitoring for heartbeats. Default:
     * `120` (2 minutes).
     */
    readonly haVmMinimumUptime: pulumi.Output<number | undefined>;
    /**
     * The type of virtual machine monitoring to use
     * when HA is enabled in the cluster. Can be one of `vmMonitoringDisabled`,
     * `vmMonitoringOnly`, or `vmAndAppMonitoring`. Default: `vmMonitoringDisabled`.
     */
    readonly haVmMonitoring: pulumi.Output<string | undefined>;
    /**
     * Determines whether or
     * not the cluster's default settings or the VM override settings specified in
     * this resource are used for virtual machine monitoring. The default is `true`
     * (use cluster defaults) - set to `false` to have overrides take effect.
     */
    readonly haVmMonitoringUseClusterDefaults: pulumi.Output<boolean | undefined>;
    /**
     * The restart priority for the virtual
     * machine when vSphere detects a host failure. Can be one of
     * `clusterRestartPriority`, `lowest`, `low`, `medium`, `high`, or `highest`.
     * Default: `clusterRestartPriority`.
     */
    readonly haVmRestartPriority: pulumi.Output<string | undefined>;
    /**
     * The maximum time, in seconds, that
     * vSphere HA will wait for this virtual machine to be ready. Use `-1` to
     * specify the cluster default.  Default: `-1`.
     * <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
     */
    readonly haVmRestartTimeout: pulumi.Output<number | undefined>;
    /**
     * The UUID of the virtual machine to create
     * the override for.  Forces a new resource if changed.
     */
    readonly virtualMachineId: pulumi.Output<string>;
    /**
     * Create a HaVmOverride resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HaVmOverrideArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering HaVmOverride resources.
 */
export interface HaVmOverrideState {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId?: pulumi.Input<string>;
    /**
     * Controls the action to take
     * on this virtual machine if an APD status on an affected datastore clears in
     * the middle of an APD event. Can be one of `useClusterDefault`, `none` or
     * `reset`.  Default: `useClusterDefault`.
     * <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
     */
    readonly haDatastoreApdRecoveryAction?: pulumi.Input<string>;
    /**
     * Controls the action to take on this
     * virtual machine when the cluster has detected loss to all paths to a relevant
     * datastore. Can be one of `clusterDefault`, `disabled`, `warning`,
     * `restartConservative`, or `restartAggressive`.  Default: `clusterDefault`.
     * <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
     */
    readonly haDatastoreApdResponse?: pulumi.Input<string>;
    /**
     * Controls the delay in minutes
     * to wait after an APD timeout event to execute the response action defined in
     * `ha_datastore_apd_response`. Use `-1` to use
     * the cluster default. Default: `-1`.
     * <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
     */
    readonly haDatastoreApdResponseDelay?: pulumi.Input<number>;
    /**
     * Controls the action to take on this
     * virtual machine when the cluster has detected a permanent device loss to a
     * relevant datastore. Can be one of `clusterDefault`, `disabled`, `warning`, or
     * `restartAggressive`. Default: `clusterDefault`.
     * <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
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
     * which `ha_vm_maximum_resets` can operate. When this
     * window expires, no more resets are attempted regardless of the setting
     * configured in `ha_vm_maximum_resets`. `-1` means no window, meaning an
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
     * <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
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
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId: pulumi.Input<string>;
    /**
     * Controls the action to take
     * on this virtual machine if an APD status on an affected datastore clears in
     * the middle of an APD event. Can be one of `useClusterDefault`, `none` or
     * `reset`.  Default: `useClusterDefault`.
     * <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
     */
    readonly haDatastoreApdRecoveryAction?: pulumi.Input<string>;
    /**
     * Controls the action to take on this
     * virtual machine when the cluster has detected loss to all paths to a relevant
     * datastore. Can be one of `clusterDefault`, `disabled`, `warning`,
     * `restartConservative`, or `restartAggressive`.  Default: `clusterDefault`.
     * <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
     */
    readonly haDatastoreApdResponse?: pulumi.Input<string>;
    /**
     * Controls the delay in minutes
     * to wait after an APD timeout event to execute the response action defined in
     * `ha_datastore_apd_response`. Use `-1` to use
     * the cluster default. Default: `-1`.
     * <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
     */
    readonly haDatastoreApdResponseDelay?: pulumi.Input<number>;
    /**
     * Controls the action to take on this
     * virtual machine when the cluster has detected a permanent device loss to a
     * relevant datastore. Can be one of `clusterDefault`, `disabled`, `warning`, or
     * `restartAggressive`. Default: `clusterDefault`.
     * <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
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
     * which `ha_vm_maximum_resets` can operate. When this
     * window expires, no more resets are attempted regardless of the setting
     * configured in `ha_vm_maximum_resets`. `-1` means no window, meaning an
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
     * <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
     */
    readonly haVmRestartTimeout?: pulumi.Input<number>;
    /**
     * The UUID of the virtual machine to create
     * the override for.  Forces a new resource if changed.
     */
    readonly virtualMachineId: pulumi.Input<string>;
}

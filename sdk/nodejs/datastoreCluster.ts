// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class DatastoreCluster extends pulumi.CustomResource {
    /**
     * Get an existing DatastoreCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatastoreClusterState, opts?: pulumi.CustomResourceOptions): DatastoreCluster {
        return new DatastoreCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/datastoreCluster:DatastoreCluster';

    /**
     * Returns true if the given object is an instance of DatastoreCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatastoreCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatastoreCluster.__pulumiType;
    }

    /**
     * A map of custom attribute ids to attribute
     * value strings to set for the datastore cluster. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     */
    public readonly customAttributes!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the datacenter to create the datastore cluster in. Forces a new resource if
     * changed.
     */
    public readonly datacenterId!: pulumi.Output<string>;
    /**
     * The name of the folder to locate the datastore cluster in.
     */
    public readonly folder!: pulumi.Output<string | undefined>;
    /**
     * The name of the datastore cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Advanced configuration options for storage DRS.
     */
    public readonly sdrsAdvancedOptions!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The global automation level for all
     * virtual machines in this datastore cluster. Default: `manual`.
     */
    public readonly sdrsAutomationLevel!: pulumi.Output<string | undefined>;
    /**
     * When `true`, all disks in a
     * single virtual machine will be kept on the same datastore. Default: `true`.
     */
    public readonly sdrsDefaultIntraVmAffinity!: pulumi.Output<boolean | undefined>;
    /**
     * Enable Storage DRS for this datastore cluster.
     * Default: `false`.
     */
    public readonly sdrsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The free space threshold to use.
     * When set to `utilization`, `drsSpaceUtilizationThreshold` is used, and
     * when set to `freeSpace`, `drsFreeSpaceThreshold` is used. Default:
     * `utilization`.
     */
    public readonly sdrsFreeSpaceThreshold!: pulumi.Output<number | undefined>;
    /**
     * The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
     * freeSpace, drs_free_space_threshold is used.
     */
    public readonly sdrsFreeSpaceThresholdMode!: pulumi.Output<string | undefined>;
    /**
     * The threshold, in
     * percent, of difference between space utilization in datastores before storage
     * DRS makes decisions to balance the space. Default: `5` percent.
     */
    public readonly sdrsFreeSpaceUtilizationDifference!: pulumi.Output<number | undefined>;
    /**
     * Overrides the default
     * automation settings when correcting I/O load imbalances.
     */
    public readonly sdrsIoBalanceAutomationLevel!: pulumi.Output<string | undefined>;
    /**
     * The I/O latency threshold, in
     * milliseconds, that storage DRS uses to make recommendations to move disks
     * from this datastore. Default: `15` seconds.
     */
    public readonly sdrsIoLatencyThreshold!: pulumi.Output<number | undefined>;
    /**
     * Enable I/O load balancing for
     * this datastore cluster. Default: `true`.
     */
    public readonly sdrsIoLoadBalanceEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The difference between load
     * in datastores in the cluster before storage DRS makes recommendations to
     * balance the load. Default: `5` percent.
     */
    public readonly sdrsIoLoadImbalanceThreshold!: pulumi.Output<number | undefined>;
    /**
     * The threshold of reservable
     * IOPS of all virtual machines on the datastore before storage DRS makes
     * recommendations to move VMs off of a datastore. Note that this setting should
     * only be set if `sdrsIoReservablePercentThreshold` cannot make an accurate
     * estimate of the capacity of the datastores in your cluster, and should be set
     * to roughly 50-60% of the worst case peak performance of the backing LUNs.
     */
    public readonly sdrsIoReservableIopsThreshold!: pulumi.Output<number | undefined>;
    /**
     * The threshold, in
     * percent, of actual estimated performance of the datastore (in IOPS) that
     * storage DRS uses to make recommendations to move VMs off of a datastore when
     * the total reservable IOPS exceeds the threshold. Default: `60` percent.
     */
    public readonly sdrsIoReservablePercentThreshold!: pulumi.Output<number | undefined>;
    /**
     * The reservable IOPS
     * threshold setting to use, `sdrsIoReservablePercentThreshold` in the event
     * of `automatic`, or `sdrsIoReservableIopsThreshold` in the event of
     * `manual`. Default: `automatic`.
     */
    public readonly sdrsIoReservableThresholdMode!: pulumi.Output<string | undefined>;
    /**
     * The storage DRS poll interval, in
     * minutes. Default: `480` minutes.
     */
    public readonly sdrsLoadBalanceInterval!: pulumi.Output<number | undefined>;
    /**
     * Overrides the default
     * automation settings when correcting storage and VM policy violations.
     */
    public readonly sdrsPolicyEnforcementAutomationLevel!: pulumi.Output<string | undefined>;
    /**
     * Overrides the default
     * automation settings when correcting affinity rule violations.
     */
    public readonly sdrsRuleEnforcementAutomationLevel!: pulumi.Output<string | undefined>;
    /**
     * Overrides the default
     * automation settings when correcting disk space imbalances.
     */
    public readonly sdrsSpaceBalanceAutomationLevel!: pulumi.Output<string | undefined>;
    /**
     * The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     */
    public readonly sdrsSpaceUtilizationThreshold!: pulumi.Output<number | undefined>;
    /**
     * Overrides the default
     * automation settings when generating recommendations for datastore evacuation.
     */
    public readonly sdrsVmEvacuationAutomationLevel!: pulumi.Output<string | undefined>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a DatastoreCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatastoreClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatastoreClusterArgs | DatastoreClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatastoreClusterState | undefined;
            inputs["customAttributes"] = state ? state.customAttributes : undefined;
            inputs["datacenterId"] = state ? state.datacenterId : undefined;
            inputs["folder"] = state ? state.folder : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["sdrsAdvancedOptions"] = state ? state.sdrsAdvancedOptions : undefined;
            inputs["sdrsAutomationLevel"] = state ? state.sdrsAutomationLevel : undefined;
            inputs["sdrsDefaultIntraVmAffinity"] = state ? state.sdrsDefaultIntraVmAffinity : undefined;
            inputs["sdrsEnabled"] = state ? state.sdrsEnabled : undefined;
            inputs["sdrsFreeSpaceThreshold"] = state ? state.sdrsFreeSpaceThreshold : undefined;
            inputs["sdrsFreeSpaceThresholdMode"] = state ? state.sdrsFreeSpaceThresholdMode : undefined;
            inputs["sdrsFreeSpaceUtilizationDifference"] = state ? state.sdrsFreeSpaceUtilizationDifference : undefined;
            inputs["sdrsIoBalanceAutomationLevel"] = state ? state.sdrsIoBalanceAutomationLevel : undefined;
            inputs["sdrsIoLatencyThreshold"] = state ? state.sdrsIoLatencyThreshold : undefined;
            inputs["sdrsIoLoadBalanceEnabled"] = state ? state.sdrsIoLoadBalanceEnabled : undefined;
            inputs["sdrsIoLoadImbalanceThreshold"] = state ? state.sdrsIoLoadImbalanceThreshold : undefined;
            inputs["sdrsIoReservableIopsThreshold"] = state ? state.sdrsIoReservableIopsThreshold : undefined;
            inputs["sdrsIoReservablePercentThreshold"] = state ? state.sdrsIoReservablePercentThreshold : undefined;
            inputs["sdrsIoReservableThresholdMode"] = state ? state.sdrsIoReservableThresholdMode : undefined;
            inputs["sdrsLoadBalanceInterval"] = state ? state.sdrsLoadBalanceInterval : undefined;
            inputs["sdrsPolicyEnforcementAutomationLevel"] = state ? state.sdrsPolicyEnforcementAutomationLevel : undefined;
            inputs["sdrsRuleEnforcementAutomationLevel"] = state ? state.sdrsRuleEnforcementAutomationLevel : undefined;
            inputs["sdrsSpaceBalanceAutomationLevel"] = state ? state.sdrsSpaceBalanceAutomationLevel : undefined;
            inputs["sdrsSpaceUtilizationThreshold"] = state ? state.sdrsSpaceUtilizationThreshold : undefined;
            inputs["sdrsVmEvacuationAutomationLevel"] = state ? state.sdrsVmEvacuationAutomationLevel : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DatastoreClusterArgs | undefined;
            if (!args || args.datacenterId === undefined) {
                throw new Error("Missing required property 'datacenterId'");
            }
            inputs["customAttributes"] = args ? args.customAttributes : undefined;
            inputs["datacenterId"] = args ? args.datacenterId : undefined;
            inputs["folder"] = args ? args.folder : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["sdrsAdvancedOptions"] = args ? args.sdrsAdvancedOptions : undefined;
            inputs["sdrsAutomationLevel"] = args ? args.sdrsAutomationLevel : undefined;
            inputs["sdrsDefaultIntraVmAffinity"] = args ? args.sdrsDefaultIntraVmAffinity : undefined;
            inputs["sdrsEnabled"] = args ? args.sdrsEnabled : undefined;
            inputs["sdrsFreeSpaceThreshold"] = args ? args.sdrsFreeSpaceThreshold : undefined;
            inputs["sdrsFreeSpaceThresholdMode"] = args ? args.sdrsFreeSpaceThresholdMode : undefined;
            inputs["sdrsFreeSpaceUtilizationDifference"] = args ? args.sdrsFreeSpaceUtilizationDifference : undefined;
            inputs["sdrsIoBalanceAutomationLevel"] = args ? args.sdrsIoBalanceAutomationLevel : undefined;
            inputs["sdrsIoLatencyThreshold"] = args ? args.sdrsIoLatencyThreshold : undefined;
            inputs["sdrsIoLoadBalanceEnabled"] = args ? args.sdrsIoLoadBalanceEnabled : undefined;
            inputs["sdrsIoLoadImbalanceThreshold"] = args ? args.sdrsIoLoadImbalanceThreshold : undefined;
            inputs["sdrsIoReservableIopsThreshold"] = args ? args.sdrsIoReservableIopsThreshold : undefined;
            inputs["sdrsIoReservablePercentThreshold"] = args ? args.sdrsIoReservablePercentThreshold : undefined;
            inputs["sdrsIoReservableThresholdMode"] = args ? args.sdrsIoReservableThresholdMode : undefined;
            inputs["sdrsLoadBalanceInterval"] = args ? args.sdrsLoadBalanceInterval : undefined;
            inputs["sdrsPolicyEnforcementAutomationLevel"] = args ? args.sdrsPolicyEnforcementAutomationLevel : undefined;
            inputs["sdrsRuleEnforcementAutomationLevel"] = args ? args.sdrsRuleEnforcementAutomationLevel : undefined;
            inputs["sdrsSpaceBalanceAutomationLevel"] = args ? args.sdrsSpaceBalanceAutomationLevel : undefined;
            inputs["sdrsSpaceUtilizationThreshold"] = args ? args.sdrsSpaceUtilizationThreshold : undefined;
            inputs["sdrsVmEvacuationAutomationLevel"] = args ? args.sdrsVmEvacuationAutomationLevel : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DatastoreCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatastoreCluster resources.
 */
export interface DatastoreClusterState {
    /**
     * A map of custom attribute ids to attribute
     * value strings to set for the datastore cluster. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     */
    readonly customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the datacenter to create the datastore cluster in. Forces a new resource if
     * changed.
     */
    readonly datacenterId?: pulumi.Input<string>;
    /**
     * The name of the folder to locate the datastore cluster in.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * The name of the datastore cluster.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Advanced configuration options for storage DRS.
     */
    readonly sdrsAdvancedOptions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The global automation level for all
     * virtual machines in this datastore cluster. Default: `manual`.
     */
    readonly sdrsAutomationLevel?: pulumi.Input<string>;
    /**
     * When `true`, all disks in a
     * single virtual machine will be kept on the same datastore. Default: `true`.
     */
    readonly sdrsDefaultIntraVmAffinity?: pulumi.Input<boolean>;
    /**
     * Enable Storage DRS for this datastore cluster.
     * Default: `false`.
     */
    readonly sdrsEnabled?: pulumi.Input<boolean>;
    /**
     * The free space threshold to use.
     * When set to `utilization`, `drsSpaceUtilizationThreshold` is used, and
     * when set to `freeSpace`, `drsFreeSpaceThreshold` is used. Default:
     * `utilization`.
     */
    readonly sdrsFreeSpaceThreshold?: pulumi.Input<number>;
    /**
     * The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
     * freeSpace, drs_free_space_threshold is used.
     */
    readonly sdrsFreeSpaceThresholdMode?: pulumi.Input<string>;
    /**
     * The threshold, in
     * percent, of difference between space utilization in datastores before storage
     * DRS makes decisions to balance the space. Default: `5` percent.
     */
    readonly sdrsFreeSpaceUtilizationDifference?: pulumi.Input<number>;
    /**
     * Overrides the default
     * automation settings when correcting I/O load imbalances.
     */
    readonly sdrsIoBalanceAutomationLevel?: pulumi.Input<string>;
    /**
     * The I/O latency threshold, in
     * milliseconds, that storage DRS uses to make recommendations to move disks
     * from this datastore. Default: `15` seconds.
     */
    readonly sdrsIoLatencyThreshold?: pulumi.Input<number>;
    /**
     * Enable I/O load balancing for
     * this datastore cluster. Default: `true`.
     */
    readonly sdrsIoLoadBalanceEnabled?: pulumi.Input<boolean>;
    /**
     * The difference between load
     * in datastores in the cluster before storage DRS makes recommendations to
     * balance the load. Default: `5` percent.
     */
    readonly sdrsIoLoadImbalanceThreshold?: pulumi.Input<number>;
    /**
     * The threshold of reservable
     * IOPS of all virtual machines on the datastore before storage DRS makes
     * recommendations to move VMs off of a datastore. Note that this setting should
     * only be set if `sdrsIoReservablePercentThreshold` cannot make an accurate
     * estimate of the capacity of the datastores in your cluster, and should be set
     * to roughly 50-60% of the worst case peak performance of the backing LUNs.
     */
    readonly sdrsIoReservableIopsThreshold?: pulumi.Input<number>;
    /**
     * The threshold, in
     * percent, of actual estimated performance of the datastore (in IOPS) that
     * storage DRS uses to make recommendations to move VMs off of a datastore when
     * the total reservable IOPS exceeds the threshold. Default: `60` percent.
     */
    readonly sdrsIoReservablePercentThreshold?: pulumi.Input<number>;
    /**
     * The reservable IOPS
     * threshold setting to use, `sdrsIoReservablePercentThreshold` in the event
     * of `automatic`, or `sdrsIoReservableIopsThreshold` in the event of
     * `manual`. Default: `automatic`.
     */
    readonly sdrsIoReservableThresholdMode?: pulumi.Input<string>;
    /**
     * The storage DRS poll interval, in
     * minutes. Default: `480` minutes.
     */
    readonly sdrsLoadBalanceInterval?: pulumi.Input<number>;
    /**
     * Overrides the default
     * automation settings when correcting storage and VM policy violations.
     */
    readonly sdrsPolicyEnforcementAutomationLevel?: pulumi.Input<string>;
    /**
     * Overrides the default
     * automation settings when correcting affinity rule violations.
     */
    readonly sdrsRuleEnforcementAutomationLevel?: pulumi.Input<string>;
    /**
     * Overrides the default
     * automation settings when correcting disk space imbalances.
     */
    readonly sdrsSpaceBalanceAutomationLevel?: pulumi.Input<string>;
    /**
     * The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     */
    readonly sdrsSpaceUtilizationThreshold?: pulumi.Input<number>;
    /**
     * Overrides the default
     * automation settings when generating recommendations for datastore evacuation.
     */
    readonly sdrsVmEvacuationAutomationLevel?: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a DatastoreCluster resource.
 */
export interface DatastoreClusterArgs {
    /**
     * A map of custom attribute ids to attribute
     * value strings to set for the datastore cluster. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     */
    readonly customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the datacenter to create the datastore cluster in. Forces a new resource if
     * changed.
     */
    readonly datacenterId: pulumi.Input<string>;
    /**
     * The name of the folder to locate the datastore cluster in.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * The name of the datastore cluster.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Advanced configuration options for storage DRS.
     */
    readonly sdrsAdvancedOptions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The global automation level for all
     * virtual machines in this datastore cluster. Default: `manual`.
     */
    readonly sdrsAutomationLevel?: pulumi.Input<string>;
    /**
     * When `true`, all disks in a
     * single virtual machine will be kept on the same datastore. Default: `true`.
     */
    readonly sdrsDefaultIntraVmAffinity?: pulumi.Input<boolean>;
    /**
     * Enable Storage DRS for this datastore cluster.
     * Default: `false`.
     */
    readonly sdrsEnabled?: pulumi.Input<boolean>;
    /**
     * The free space threshold to use.
     * When set to `utilization`, `drsSpaceUtilizationThreshold` is used, and
     * when set to `freeSpace`, `drsFreeSpaceThreshold` is used. Default:
     * `utilization`.
     */
    readonly sdrsFreeSpaceThreshold?: pulumi.Input<number>;
    /**
     * The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
     * freeSpace, drs_free_space_threshold is used.
     */
    readonly sdrsFreeSpaceThresholdMode?: pulumi.Input<string>;
    /**
     * The threshold, in
     * percent, of difference between space utilization in datastores before storage
     * DRS makes decisions to balance the space. Default: `5` percent.
     */
    readonly sdrsFreeSpaceUtilizationDifference?: pulumi.Input<number>;
    /**
     * Overrides the default
     * automation settings when correcting I/O load imbalances.
     */
    readonly sdrsIoBalanceAutomationLevel?: pulumi.Input<string>;
    /**
     * The I/O latency threshold, in
     * milliseconds, that storage DRS uses to make recommendations to move disks
     * from this datastore. Default: `15` seconds.
     */
    readonly sdrsIoLatencyThreshold?: pulumi.Input<number>;
    /**
     * Enable I/O load balancing for
     * this datastore cluster. Default: `true`.
     */
    readonly sdrsIoLoadBalanceEnabled?: pulumi.Input<boolean>;
    /**
     * The difference between load
     * in datastores in the cluster before storage DRS makes recommendations to
     * balance the load. Default: `5` percent.
     */
    readonly sdrsIoLoadImbalanceThreshold?: pulumi.Input<number>;
    /**
     * The threshold of reservable
     * IOPS of all virtual machines on the datastore before storage DRS makes
     * recommendations to move VMs off of a datastore. Note that this setting should
     * only be set if `sdrsIoReservablePercentThreshold` cannot make an accurate
     * estimate of the capacity of the datastores in your cluster, and should be set
     * to roughly 50-60% of the worst case peak performance of the backing LUNs.
     */
    readonly sdrsIoReservableIopsThreshold?: pulumi.Input<number>;
    /**
     * The threshold, in
     * percent, of actual estimated performance of the datastore (in IOPS) that
     * storage DRS uses to make recommendations to move VMs off of a datastore when
     * the total reservable IOPS exceeds the threshold. Default: `60` percent.
     */
    readonly sdrsIoReservablePercentThreshold?: pulumi.Input<number>;
    /**
     * The reservable IOPS
     * threshold setting to use, `sdrsIoReservablePercentThreshold` in the event
     * of `automatic`, or `sdrsIoReservableIopsThreshold` in the event of
     * `manual`. Default: `automatic`.
     */
    readonly sdrsIoReservableThresholdMode?: pulumi.Input<string>;
    /**
     * The storage DRS poll interval, in
     * minutes. Default: `480` minutes.
     */
    readonly sdrsLoadBalanceInterval?: pulumi.Input<number>;
    /**
     * Overrides the default
     * automation settings when correcting storage and VM policy violations.
     */
    readonly sdrsPolicyEnforcementAutomationLevel?: pulumi.Input<string>;
    /**
     * Overrides the default
     * automation settings when correcting affinity rule violations.
     */
    readonly sdrsRuleEnforcementAutomationLevel?: pulumi.Input<string>;
    /**
     * Overrides the default
     * automation settings when correcting disk space imbalances.
     */
    readonly sdrsSpaceBalanceAutomationLevel?: pulumi.Input<string>;
    /**
     * The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     */
    readonly sdrsSpaceUtilizationThreshold?: pulumi.Input<number>;
    /**
     * Overrides the default
     * automation settings when generating recommendations for datastore evacuation.
     */
    readonly sdrsVmEvacuationAutomationLevel?: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}

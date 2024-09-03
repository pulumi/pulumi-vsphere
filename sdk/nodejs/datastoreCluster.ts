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
     * @param opts Optional settings to control the behavior of the CustomResource.
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
     * [here](https://www.terraform.io/docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource) for a reference on how to set values
     * for custom attributes.
     *
     *
     * > **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     */
    public readonly customAttributes!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The managed object ID of
     * the datacenter to create the datastore cluster in. Forces a new resource if
     * changed.
     */
    public readonly datacenterId!: pulumi.Output<string>;
    /**
     * The relative path to a folder to put this datastore
     * cluster in.  This is a path relative to the datacenter you are deploying the
     * datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
     * `foo/bar`, The provider will place a datastore cluster named
     * `datastore-cluster-test` in a datastore folder located at
     * `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/datastore-cluster-test`.
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
     * The default automation level for all virtual machines in this storage cluster.
     */
    public readonly sdrsAutomationLevel!: pulumi.Output<string | undefined>;
    /**
     * When true, storage DRS keeps VMDKs for individual VMs on the same datastore by default.
     */
    public readonly sdrsDefaultIntraVmAffinity!: pulumi.Output<boolean | undefined>;
    /**
     * Enable Storage DRS for this datastore cluster.
     * Default: `false`.
     */
    public readonly sdrsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The threshold, in GB, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     */
    public readonly sdrsFreeSpaceThreshold!: pulumi.Output<number | undefined>;
    /**
     * The free space threshold to use. When set to utilization, drsSpaceUtilizationThreshold is used, and when set to
     * freeSpace, drsFreeSpaceThreshold is used.
     */
    public readonly sdrsFreeSpaceThresholdMode!: pulumi.Output<string | undefined>;
    /**
     * The threshold, in percent, of difference between space utilization in datastores before storage DRS makes decisions to
     * balance the space.
     */
    public readonly sdrsFreeSpaceUtilizationDifference!: pulumi.Output<number | undefined>;
    /**
     * Overrides the default automation settings when correcting I/O load imbalances.
     */
    public readonly sdrsIoBalanceAutomationLevel!: pulumi.Output<string | undefined>;
    /**
     * The I/O latency threshold, in milliseconds, that storage DRS uses to make recommendations to move disks from this
     * datastore.
     */
    public readonly sdrsIoLatencyThreshold!: pulumi.Output<number | undefined>;
    /**
     * Enable I/O load balancing for this datastore cluster.
     */
    public readonly sdrsIoLoadBalanceEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The difference between load in datastores in the cluster before storage DRS makes recommendations to balance the load.
     */
    public readonly sdrsIoLoadImbalanceThreshold!: pulumi.Output<number | undefined>;
    /**
     * The threshold of reservable IOPS of all virtual machines on the datastore before storage DRS makes recommendations to
     * move VMs off of a datastore.
     */
    public readonly sdrsIoReservableIopsThreshold!: pulumi.Output<number | undefined>;
    /**
     * The threshold, in percent, of actual estimated performance of the datastore (in IOPS) that storage DRS uses to make
     * recommendations to move VMs off of a datastore when the total reservable IOPS exceeds the threshold.
     */
    public readonly sdrsIoReservablePercentThreshold!: pulumi.Output<number | undefined>;
    /**
     * The reservable IOPS threshold to use, percent in the event of automatic, or manual threshold in the event of manual.
     */
    public readonly sdrsIoReservableThresholdMode!: pulumi.Output<string | undefined>;
    /**
     * The storage DRS poll interval, in minutes.
     */
    public readonly sdrsLoadBalanceInterval!: pulumi.Output<number | undefined>;
    /**
     * Overrides the default automation settings when correcting storage and VM policy violations.
     */
    public readonly sdrsPolicyEnforcementAutomationLevel!: pulumi.Output<string | undefined>;
    /**
     * Overrides the default automation settings when correcting affinity rule violations.
     */
    public readonly sdrsRuleEnforcementAutomationLevel!: pulumi.Output<string | undefined>;
    /**
     * Overrides the default automation settings when correcting disk space imbalances.
     */
    public readonly sdrsSpaceBalanceAutomationLevel!: pulumi.Output<string | undefined>;
    /**
     * The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     */
    public readonly sdrsSpaceUtilizationThreshold!: pulumi.Output<number | undefined>;
    /**
     * Overrides the default automation settings when generating recommendations for datastore evacuation.
     */
    public readonly sdrsVmEvacuationAutomationLevel!: pulumi.Output<string | undefined>;
    /**
     * The IDs of any tags to attach to this resource.
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatastoreClusterState | undefined;
            resourceInputs["customAttributes"] = state ? state.customAttributes : undefined;
            resourceInputs["datacenterId"] = state ? state.datacenterId : undefined;
            resourceInputs["folder"] = state ? state.folder : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["sdrsAdvancedOptions"] = state ? state.sdrsAdvancedOptions : undefined;
            resourceInputs["sdrsAutomationLevel"] = state ? state.sdrsAutomationLevel : undefined;
            resourceInputs["sdrsDefaultIntraVmAffinity"] = state ? state.sdrsDefaultIntraVmAffinity : undefined;
            resourceInputs["sdrsEnabled"] = state ? state.sdrsEnabled : undefined;
            resourceInputs["sdrsFreeSpaceThreshold"] = state ? state.sdrsFreeSpaceThreshold : undefined;
            resourceInputs["sdrsFreeSpaceThresholdMode"] = state ? state.sdrsFreeSpaceThresholdMode : undefined;
            resourceInputs["sdrsFreeSpaceUtilizationDifference"] = state ? state.sdrsFreeSpaceUtilizationDifference : undefined;
            resourceInputs["sdrsIoBalanceAutomationLevel"] = state ? state.sdrsIoBalanceAutomationLevel : undefined;
            resourceInputs["sdrsIoLatencyThreshold"] = state ? state.sdrsIoLatencyThreshold : undefined;
            resourceInputs["sdrsIoLoadBalanceEnabled"] = state ? state.sdrsIoLoadBalanceEnabled : undefined;
            resourceInputs["sdrsIoLoadImbalanceThreshold"] = state ? state.sdrsIoLoadImbalanceThreshold : undefined;
            resourceInputs["sdrsIoReservableIopsThreshold"] = state ? state.sdrsIoReservableIopsThreshold : undefined;
            resourceInputs["sdrsIoReservablePercentThreshold"] = state ? state.sdrsIoReservablePercentThreshold : undefined;
            resourceInputs["sdrsIoReservableThresholdMode"] = state ? state.sdrsIoReservableThresholdMode : undefined;
            resourceInputs["sdrsLoadBalanceInterval"] = state ? state.sdrsLoadBalanceInterval : undefined;
            resourceInputs["sdrsPolicyEnforcementAutomationLevel"] = state ? state.sdrsPolicyEnforcementAutomationLevel : undefined;
            resourceInputs["sdrsRuleEnforcementAutomationLevel"] = state ? state.sdrsRuleEnforcementAutomationLevel : undefined;
            resourceInputs["sdrsSpaceBalanceAutomationLevel"] = state ? state.sdrsSpaceBalanceAutomationLevel : undefined;
            resourceInputs["sdrsSpaceUtilizationThreshold"] = state ? state.sdrsSpaceUtilizationThreshold : undefined;
            resourceInputs["sdrsVmEvacuationAutomationLevel"] = state ? state.sdrsVmEvacuationAutomationLevel : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DatastoreClusterArgs | undefined;
            if ((!args || args.datacenterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datacenterId'");
            }
            resourceInputs["customAttributes"] = args ? args.customAttributes : undefined;
            resourceInputs["datacenterId"] = args ? args.datacenterId : undefined;
            resourceInputs["folder"] = args ? args.folder : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["sdrsAdvancedOptions"] = args ? args.sdrsAdvancedOptions : undefined;
            resourceInputs["sdrsAutomationLevel"] = args ? args.sdrsAutomationLevel : undefined;
            resourceInputs["sdrsDefaultIntraVmAffinity"] = args ? args.sdrsDefaultIntraVmAffinity : undefined;
            resourceInputs["sdrsEnabled"] = args ? args.sdrsEnabled : undefined;
            resourceInputs["sdrsFreeSpaceThreshold"] = args ? args.sdrsFreeSpaceThreshold : undefined;
            resourceInputs["sdrsFreeSpaceThresholdMode"] = args ? args.sdrsFreeSpaceThresholdMode : undefined;
            resourceInputs["sdrsFreeSpaceUtilizationDifference"] = args ? args.sdrsFreeSpaceUtilizationDifference : undefined;
            resourceInputs["sdrsIoBalanceAutomationLevel"] = args ? args.sdrsIoBalanceAutomationLevel : undefined;
            resourceInputs["sdrsIoLatencyThreshold"] = args ? args.sdrsIoLatencyThreshold : undefined;
            resourceInputs["sdrsIoLoadBalanceEnabled"] = args ? args.sdrsIoLoadBalanceEnabled : undefined;
            resourceInputs["sdrsIoLoadImbalanceThreshold"] = args ? args.sdrsIoLoadImbalanceThreshold : undefined;
            resourceInputs["sdrsIoReservableIopsThreshold"] = args ? args.sdrsIoReservableIopsThreshold : undefined;
            resourceInputs["sdrsIoReservablePercentThreshold"] = args ? args.sdrsIoReservablePercentThreshold : undefined;
            resourceInputs["sdrsIoReservableThresholdMode"] = args ? args.sdrsIoReservableThresholdMode : undefined;
            resourceInputs["sdrsLoadBalanceInterval"] = args ? args.sdrsLoadBalanceInterval : undefined;
            resourceInputs["sdrsPolicyEnforcementAutomationLevel"] = args ? args.sdrsPolicyEnforcementAutomationLevel : undefined;
            resourceInputs["sdrsRuleEnforcementAutomationLevel"] = args ? args.sdrsRuleEnforcementAutomationLevel : undefined;
            resourceInputs["sdrsSpaceBalanceAutomationLevel"] = args ? args.sdrsSpaceBalanceAutomationLevel : undefined;
            resourceInputs["sdrsSpaceUtilizationThreshold"] = args ? args.sdrsSpaceUtilizationThreshold : undefined;
            resourceInputs["sdrsVmEvacuationAutomationLevel"] = args ? args.sdrsVmEvacuationAutomationLevel : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatastoreCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatastoreCluster resources.
 */
export interface DatastoreClusterState {
    /**
     * A map of custom attribute ids to attribute
     * value strings to set for the datastore cluster. See
     * [here](https://www.terraform.io/docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource) for a reference on how to set values
     * for custom attributes.
     *
     *
     * > **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     */
    customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The managed object ID of
     * the datacenter to create the datastore cluster in. Forces a new resource if
     * changed.
     */
    datacenterId?: pulumi.Input<string>;
    /**
     * The relative path to a folder to put this datastore
     * cluster in.  This is a path relative to the datacenter you are deploying the
     * datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
     * `foo/bar`, The provider will place a datastore cluster named
     * `datastore-cluster-test` in a datastore folder located at
     * `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/datastore-cluster-test`.
     */
    folder?: pulumi.Input<string>;
    /**
     * The name of the datastore cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * Advanced configuration options for storage DRS.
     */
    sdrsAdvancedOptions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The default automation level for all virtual machines in this storage cluster.
     */
    sdrsAutomationLevel?: pulumi.Input<string>;
    /**
     * When true, storage DRS keeps VMDKs for individual VMs on the same datastore by default.
     */
    sdrsDefaultIntraVmAffinity?: pulumi.Input<boolean>;
    /**
     * Enable Storage DRS for this datastore cluster.
     * Default: `false`.
     */
    sdrsEnabled?: pulumi.Input<boolean>;
    /**
     * The threshold, in GB, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     */
    sdrsFreeSpaceThreshold?: pulumi.Input<number>;
    /**
     * The free space threshold to use. When set to utilization, drsSpaceUtilizationThreshold is used, and when set to
     * freeSpace, drsFreeSpaceThreshold is used.
     */
    sdrsFreeSpaceThresholdMode?: pulumi.Input<string>;
    /**
     * The threshold, in percent, of difference between space utilization in datastores before storage DRS makes decisions to
     * balance the space.
     */
    sdrsFreeSpaceUtilizationDifference?: pulumi.Input<number>;
    /**
     * Overrides the default automation settings when correcting I/O load imbalances.
     */
    sdrsIoBalanceAutomationLevel?: pulumi.Input<string>;
    /**
     * The I/O latency threshold, in milliseconds, that storage DRS uses to make recommendations to move disks from this
     * datastore.
     */
    sdrsIoLatencyThreshold?: pulumi.Input<number>;
    /**
     * Enable I/O load balancing for this datastore cluster.
     */
    sdrsIoLoadBalanceEnabled?: pulumi.Input<boolean>;
    /**
     * The difference between load in datastores in the cluster before storage DRS makes recommendations to balance the load.
     */
    sdrsIoLoadImbalanceThreshold?: pulumi.Input<number>;
    /**
     * The threshold of reservable IOPS of all virtual machines on the datastore before storage DRS makes recommendations to
     * move VMs off of a datastore.
     */
    sdrsIoReservableIopsThreshold?: pulumi.Input<number>;
    /**
     * The threshold, in percent, of actual estimated performance of the datastore (in IOPS) that storage DRS uses to make
     * recommendations to move VMs off of a datastore when the total reservable IOPS exceeds the threshold.
     */
    sdrsIoReservablePercentThreshold?: pulumi.Input<number>;
    /**
     * The reservable IOPS threshold to use, percent in the event of automatic, or manual threshold in the event of manual.
     */
    sdrsIoReservableThresholdMode?: pulumi.Input<string>;
    /**
     * The storage DRS poll interval, in minutes.
     */
    sdrsLoadBalanceInterval?: pulumi.Input<number>;
    /**
     * Overrides the default automation settings when correcting storage and VM policy violations.
     */
    sdrsPolicyEnforcementAutomationLevel?: pulumi.Input<string>;
    /**
     * Overrides the default automation settings when correcting affinity rule violations.
     */
    sdrsRuleEnforcementAutomationLevel?: pulumi.Input<string>;
    /**
     * Overrides the default automation settings when correcting disk space imbalances.
     */
    sdrsSpaceBalanceAutomationLevel?: pulumi.Input<string>;
    /**
     * The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     */
    sdrsSpaceUtilizationThreshold?: pulumi.Input<number>;
    /**
     * Overrides the default automation settings when generating recommendations for datastore evacuation.
     */
    sdrsVmEvacuationAutomationLevel?: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a DatastoreCluster resource.
 */
export interface DatastoreClusterArgs {
    /**
     * A map of custom attribute ids to attribute
     * value strings to set for the datastore cluster. See
     * [here](https://www.terraform.io/docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource) for a reference on how to set values
     * for custom attributes.
     *
     *
     * > **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     */
    customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The managed object ID of
     * the datacenter to create the datastore cluster in. Forces a new resource if
     * changed.
     */
    datacenterId: pulumi.Input<string>;
    /**
     * The relative path to a folder to put this datastore
     * cluster in.  This is a path relative to the datacenter you are deploying the
     * datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
     * `foo/bar`, The provider will place a datastore cluster named
     * `datastore-cluster-test` in a datastore folder located at
     * `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/datastore-cluster-test`.
     */
    folder?: pulumi.Input<string>;
    /**
     * The name of the datastore cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * Advanced configuration options for storage DRS.
     */
    sdrsAdvancedOptions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The default automation level for all virtual machines in this storage cluster.
     */
    sdrsAutomationLevel?: pulumi.Input<string>;
    /**
     * When true, storage DRS keeps VMDKs for individual VMs on the same datastore by default.
     */
    sdrsDefaultIntraVmAffinity?: pulumi.Input<boolean>;
    /**
     * Enable Storage DRS for this datastore cluster.
     * Default: `false`.
     */
    sdrsEnabled?: pulumi.Input<boolean>;
    /**
     * The threshold, in GB, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     */
    sdrsFreeSpaceThreshold?: pulumi.Input<number>;
    /**
     * The free space threshold to use. When set to utilization, drsSpaceUtilizationThreshold is used, and when set to
     * freeSpace, drsFreeSpaceThreshold is used.
     */
    sdrsFreeSpaceThresholdMode?: pulumi.Input<string>;
    /**
     * The threshold, in percent, of difference between space utilization in datastores before storage DRS makes decisions to
     * balance the space.
     */
    sdrsFreeSpaceUtilizationDifference?: pulumi.Input<number>;
    /**
     * Overrides the default automation settings when correcting I/O load imbalances.
     */
    sdrsIoBalanceAutomationLevel?: pulumi.Input<string>;
    /**
     * The I/O latency threshold, in milliseconds, that storage DRS uses to make recommendations to move disks from this
     * datastore.
     */
    sdrsIoLatencyThreshold?: pulumi.Input<number>;
    /**
     * Enable I/O load balancing for this datastore cluster.
     */
    sdrsIoLoadBalanceEnabled?: pulumi.Input<boolean>;
    /**
     * The difference between load in datastores in the cluster before storage DRS makes recommendations to balance the load.
     */
    sdrsIoLoadImbalanceThreshold?: pulumi.Input<number>;
    /**
     * The threshold of reservable IOPS of all virtual machines on the datastore before storage DRS makes recommendations to
     * move VMs off of a datastore.
     */
    sdrsIoReservableIopsThreshold?: pulumi.Input<number>;
    /**
     * The threshold, in percent, of actual estimated performance of the datastore (in IOPS) that storage DRS uses to make
     * recommendations to move VMs off of a datastore when the total reservable IOPS exceeds the threshold.
     */
    sdrsIoReservablePercentThreshold?: pulumi.Input<number>;
    /**
     * The reservable IOPS threshold to use, percent in the event of automatic, or manual threshold in the event of manual.
     */
    sdrsIoReservableThresholdMode?: pulumi.Input<string>;
    /**
     * The storage DRS poll interval, in minutes.
     */
    sdrsLoadBalanceInterval?: pulumi.Input<number>;
    /**
     * Overrides the default automation settings when correcting storage and VM policy violations.
     */
    sdrsPolicyEnforcementAutomationLevel?: pulumi.Input<string>;
    /**
     * Overrides the default automation settings when correcting affinity rule violations.
     */
    sdrsRuleEnforcementAutomationLevel?: pulumi.Input<string>;
    /**
     * Overrides the default automation settings when correcting disk space imbalances.
     */
    sdrsSpaceBalanceAutomationLevel?: pulumi.Input<string>;
    /**
     * The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     */
    sdrsSpaceUtilizationThreshold?: pulumi.Input<number>;
    /**
     * Overrides the default automation settings when generating recommendations for datastore evacuation.
     */
    sdrsVmEvacuationAutomationLevel?: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}

import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_datastore_cluster` resource can be used to create and manage
 * datastore clusters. This can be used to create groups of datastores with a
 * shared management interface, allowing for resource control and load balancing
 * through Storage DRS.
 *
 * For more information on vSphere datastore clusters and Storage DRS, see [this
 * page][ref-vsphere-datastore-clusters].
 *
 * [ref-vsphere-datastore-clusters]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-598DF695-107E-406B-9C95-0AF961FC227A.html
 *
 * ~> **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 *
 * ~> **NOTE:** Storage DRS requires a vSphere Enterprise Plus license.
 */
export declare class DatastoreCluster extends pulumi.CustomResource {
    /**
     * Get an existing DatastoreCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatastoreClusterState): DatastoreCluster;
    /**
     * A map of custom attribute ids to attribute
     * value strings to set for the datastore cluster. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     */
    readonly customAttributes: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the datacenter to create the datastore cluster in. Forces a new resource if
     * changed.
     */
    readonly datacenterId: pulumi.Output<string>;
    /**
     * The relative path to a folder to put this datastore
     * cluster in.  This is a path relative to the datacenter you are deploying the
     * datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
     * `foo/bar`, Terraform will place a datastore cluster named
     * `terraform-datastore-cluster-test` in a datastore folder located at
     * `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/terraform-datastore-cluster-test`.
     */
    readonly folder: pulumi.Output<string | undefined>;
    /**
     * The name of the datastore cluster.
     */
    readonly name: pulumi.Output<string>;
    /**
     * A key/value map of advanced Storage DRS
     * settings that are not exposed via Terraform or the vSphere client.
     */
    readonly sdrsAdvancedOptions: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * The global automation level for all
     * virtual machines in this datastore cluster. Default: `manual`.
     */
    readonly sdrsAutomationLevel: pulumi.Output<string | undefined>;
    /**
     * When `true`, all disks in a
     * single virtual machine will be kept on the same datastore. Default: `true`.
     */
    readonly sdrsDefaultIntraVmAffinity: pulumi.Output<boolean | undefined>;
    /**
     * Enable Storage DRS for this datastore cluster.
     * Default: `false`.
     */
    readonly sdrsEnabled: pulumi.Output<boolean | undefined>;
    /**
     * The free space threshold to use.
     * When set to `utilization`, `drs_space_utilization_threshold` is used, and
     * when set to `freeSpace`, `drs_free_space_threshold` is used. Default:
     * `utilization`.
     */
    readonly sdrsFreeSpaceThreshold: pulumi.Output<number | undefined>;
    /**
     * The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
     * freeSpace, drs_free_space_threshold is used.
     */
    readonly sdrsFreeSpaceThresholdMode: pulumi.Output<string | undefined>;
    /**
     * The threshold, in
     * percent, of difference between space utilization in datastores before storage
     * DRS makes decisions to balance the space. Default: `5` percent.
     */
    readonly sdrsFreeSpaceUtilizationDifference: pulumi.Output<number | undefined>;
    /**
     * Overrides the default
     * automation settings when correcting I/O load imbalances.
     */
    readonly sdrsIoBalanceAutomationLevel: pulumi.Output<string | undefined>;
    /**
     * The I/O latency threshold, in
     * milliseconds, that storage DRS uses to make recommendations to move disks
     * from this datastore. Default: `15` seconds.
     */
    readonly sdrsIoLatencyThreshold: pulumi.Output<number | undefined>;
    /**
     * Enable I/O load balancing for
     * this datastore cluster. Default: `true`.
     */
    readonly sdrsIoLoadBalanceEnabled: pulumi.Output<boolean | undefined>;
    /**
     * The difference between load
     * in datastores in the cluster before storage DRS makes recommendations to
     * balance the load. Default: `5` percent.
     */
    readonly sdrsIoLoadImbalanceThreshold: pulumi.Output<number | undefined>;
    /**
     * The threshold of reservable
     * IOPS of all virtual machines on the datastore before storage DRS makes
     * recommendations to move VMs off of a datastore. Note that this setting should
     * only be set if `sdrs_io_reservable_percent_threshold` cannot make an accurate
     * estimate of the capacity of the datastores in your cluster, and should be set
     * to roughly 50-60% of the worst case peak performance of the backing LUNs.
     */
    readonly sdrsIoReservableIopsThreshold: pulumi.Output<number | undefined>;
    /**
     * The threshold, in
     * percent, of actual estimated performance of the datastore (in IOPS) that
     * storage DRS uses to make recommendations to move VMs off of a datastore when
     * the total reservable IOPS exceeds the threshold. Default: `60` percent.
     */
    readonly sdrsIoReservablePercentThreshold: pulumi.Output<number | undefined>;
    /**
     * The reservable IOPS
     * threshold setting to use, `sdrs_io_reservable_percent_threshold` in the event
     * of `automatic`, or `sdrs_io_reservable_iops_threshold` in the event of
     * `manual`. Default: `automatic`.
     */
    readonly sdrsIoReservableThresholdMode: pulumi.Output<string | undefined>;
    /**
     * The storage DRS poll interval, in
     * minutes. Default: `480` minutes.
     */
    readonly sdrsLoadBalanceInterval: pulumi.Output<number | undefined>;
    /**
     * Overrides the default
     * automation settings when correcting storage and VM policy violations.
     */
    readonly sdrsPolicyEnforcementAutomationLevel: pulumi.Output<string | undefined>;
    /**
     * Overrides the default
     * automation settings when correcting affinity rule violations.
     */
    readonly sdrsRuleEnforcementAutomationLevel: pulumi.Output<string | undefined>;
    /**
     * Overrides the default
     * automation settings when correcting disk space imbalances.
     */
    readonly sdrsSpaceBalanceAutomationLevel: pulumi.Output<string | undefined>;
    /**
     * The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     */
    readonly sdrsSpaceUtilizationThreshold: pulumi.Output<number | undefined>;
    /**
     * Overrides the default
     * automation settings when generating recommendations for datastore evacuation.
     */
    readonly sdrsVmEvacuationAutomationLevel: pulumi.Output<string | undefined>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags: pulumi.Output<string[] | undefined>;
    /**
     * Create a DatastoreCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatastoreClusterArgs, opts?: pulumi.CustomResourceOptions);
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
    readonly customAttributes?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the datacenter to create the datastore cluster in. Forces a new resource if
     * changed.
     */
    readonly datacenterId?: pulumi.Input<string>;
    /**
     * The relative path to a folder to put this datastore
     * cluster in.  This is a path relative to the datacenter you are deploying the
     * datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
     * `foo/bar`, Terraform will place a datastore cluster named
     * `terraform-datastore-cluster-test` in a datastore folder located at
     * `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/terraform-datastore-cluster-test`.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * The name of the datastore cluster.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A key/value map of advanced Storage DRS
     * settings that are not exposed via Terraform or the vSphere client.
     */
    readonly sdrsAdvancedOptions?: pulumi.Input<{
        [key: string]: any;
    }>;
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
     * When set to `utilization`, `drs_space_utilization_threshold` is used, and
     * when set to `freeSpace`, `drs_free_space_threshold` is used. Default:
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
     * only be set if `sdrs_io_reservable_percent_threshold` cannot make an accurate
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
     * threshold setting to use, `sdrs_io_reservable_percent_threshold` in the event
     * of `automatic`, or `sdrs_io_reservable_iops_threshold` in the event of
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
    readonly customAttributes?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the datacenter to create the datastore cluster in. Forces a new resource if
     * changed.
     */
    readonly datacenterId: pulumi.Input<string>;
    /**
     * The relative path to a folder to put this datastore
     * cluster in.  This is a path relative to the datacenter you are deploying the
     * datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
     * `foo/bar`, Terraform will place a datastore cluster named
     * `terraform-datastore-cluster-test` in a datastore folder located at
     * `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/terraform-datastore-cluster-test`.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * The name of the datastore cluster.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A key/value map of advanced Storage DRS
     * settings that are not exposed via Terraform or the vSphere client.
     */
    readonly sdrsAdvancedOptions?: pulumi.Input<{
        [key: string]: any;
    }>;
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
     * When set to `utilization`, `drs_space_utilization_threshold` is used, and
     * when set to `freeSpace`, `drs_free_space_threshold` is used. Default:
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
     * only be set if `sdrs_io_reservable_percent_threshold` cannot make an accurate
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
     * threshold setting to use, `sdrs_io_reservable_percent_threshold` in the event
     * of `automatic`, or `sdrs_io_reservable_iops_threshold` in the event of
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

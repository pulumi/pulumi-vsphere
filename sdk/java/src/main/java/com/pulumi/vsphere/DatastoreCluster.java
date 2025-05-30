// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.DatastoreClusterArgs;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.inputs.DatastoreClusterState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `vsphere.DatastoreCluster` resource can be used to create and manage
 * datastore clusters. This can be used to create groups of datastores with a
 * shared management interface, allowing for resource control and load balancing
 * through Storage DRS.
 * 
 * For more information on vSphere datastore clusters and Storage DRS, see [this
 * page][ref-vsphere-datastore-clusters].
 * 
 * [ref-vsphere-datastore-clusters]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-resource-management-8-0/creating-a-datastore-cluster.html
 * 
 * &gt; **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 * 
 * &gt; **NOTE:** Storage DRS requires a vSphere Enterprise Plus license.
 * 
 * ## Example Usage
 * 
 * The following example sets up a datastore cluster and enables Storage DRS with
 * the default settings. It then creates two NAS datastores using the
 * `vsphere.NasDatastore` resource and assigns them to
 * the datastore cluster.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * An existing datastore cluster can be imported into this resource
 * 
 * via the path to the cluster, via the following command:
 * 
 * [docs-import]: https://developer.hashicorp.com/terraform/cli/import
 * 
 * ```sh
 * $ pulumi import vsphere:index/datastoreCluster:DatastoreCluster datastore_cluster /dc1/datastore/ds-cluster
 * ```
 * 
 * The above would import the datastore cluster named `ds-cluster` that is located
 * 
 * in the `dc1` datacenter.
 * 
 */
@ResourceType(type="vsphere:index/datastoreCluster:DatastoreCluster")
public class DatastoreCluster extends com.pulumi.resources.CustomResource {
    /**
     * A map of custom attribute ids to attribute
     * value strings to set for the datastore cluster. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     * 
     * [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource
     * 
     * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     * 
     */
    @Export(name="customAttributes", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> customAttributes;

    /**
     * @return A map of custom attribute ids to attribute
     * value strings to set for the datastore cluster. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     * 
     * [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource
     * 
     * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     * 
     */
    public Output<Optional<Map<String,String>>> customAttributes() {
        return Codegen.optional(this.customAttributes);
    }
    /**
     * The managed object ID of
     * the datacenter to create the datastore cluster in. Forces a new resource if
     * changed.
     * 
     */
    @Export(name="datacenterId", refs={String.class}, tree="[0]")
    private Output<String> datacenterId;

    /**
     * @return The managed object ID of
     * the datacenter to create the datastore cluster in. Forces a new resource if
     * changed.
     * 
     */
    public Output<String> datacenterId() {
        return this.datacenterId;
    }
    /**
     * The relative path to a folder to put this datastore
     * cluster in.  This is a path relative to the datacenter you are deploying the
     * datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
     * `foo/bar`, The provider will place a datastore cluster named
     * `datastore-cluster-test` in a datastore folder located at
     * `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/datastore-cluster-test`.
     * 
     */
    @Export(name="folder", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> folder;

    /**
     * @return The relative path to a folder to put this datastore
     * cluster in.  This is a path relative to the datacenter you are deploying the
     * datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
     * `foo/bar`, The provider will place a datastore cluster named
     * `datastore-cluster-test` in a datastore folder located at
     * `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/datastore-cluster-test`.
     * 
     */
    public Output<Optional<String>> folder() {
        return Codegen.optional(this.folder);
    }
    /**
     * The name of the datastore cluster.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the datastore cluster.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Advanced configuration options for storage DRS.
     * 
     */
    @Export(name="sdrsAdvancedOptions", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> sdrsAdvancedOptions;

    /**
     * @return Advanced configuration options for storage DRS.
     * 
     */
    public Output<Optional<Map<String,String>>> sdrsAdvancedOptions() {
        return Codegen.optional(this.sdrsAdvancedOptions);
    }
    /**
     * The default automation level for all virtual machines in this storage cluster.
     * 
     */
    @Export(name="sdrsAutomationLevel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sdrsAutomationLevel;

    /**
     * @return The default automation level for all virtual machines in this storage cluster.
     * 
     */
    public Output<Optional<String>> sdrsAutomationLevel() {
        return Codegen.optional(this.sdrsAutomationLevel);
    }
    /**
     * When true, storage DRS keeps VMDKs for individual VMs on the same datastore by default.
     * 
     */
    @Export(name="sdrsDefaultIntraVmAffinity", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> sdrsDefaultIntraVmAffinity;

    /**
     * @return When true, storage DRS keeps VMDKs for individual VMs on the same datastore by default.
     * 
     */
    public Output<Optional<Boolean>> sdrsDefaultIntraVmAffinity() {
        return Codegen.optional(this.sdrsDefaultIntraVmAffinity);
    }
    /**
     * Enable Storage DRS for this datastore cluster.
     * Default: `false`.
     * 
     */
    @Export(name="sdrsEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> sdrsEnabled;

    /**
     * @return Enable Storage DRS for this datastore cluster.
     * Default: `false`.
     * 
     */
    public Output<Optional<Boolean>> sdrsEnabled() {
        return Codegen.optional(this.sdrsEnabled);
    }
    /**
     * The threshold, in GB, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     * 
     */
    @Export(name="sdrsFreeSpaceThreshold", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> sdrsFreeSpaceThreshold;

    /**
     * @return The threshold, in GB, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     * 
     */
    public Output<Optional<Integer>> sdrsFreeSpaceThreshold() {
        return Codegen.optional(this.sdrsFreeSpaceThreshold);
    }
    /**
     * The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
     * freeSpace, drs_free_space_threshold is used.
     * 
     */
    @Export(name="sdrsFreeSpaceThresholdMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sdrsFreeSpaceThresholdMode;

    /**
     * @return The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
     * freeSpace, drs_free_space_threshold is used.
     * 
     */
    public Output<Optional<String>> sdrsFreeSpaceThresholdMode() {
        return Codegen.optional(this.sdrsFreeSpaceThresholdMode);
    }
    /**
     * The threshold, in percent, of difference between space utilization in datastores before storage DRS makes decisions to
     * balance the space.
     * 
     */
    @Export(name="sdrsFreeSpaceUtilizationDifference", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> sdrsFreeSpaceUtilizationDifference;

    /**
     * @return The threshold, in percent, of difference between space utilization in datastores before storage DRS makes decisions to
     * balance the space.
     * 
     */
    public Output<Optional<Integer>> sdrsFreeSpaceUtilizationDifference() {
        return Codegen.optional(this.sdrsFreeSpaceUtilizationDifference);
    }
    /**
     * Overrides the default automation settings when correcting I/O load imbalances.
     * 
     */
    @Export(name="sdrsIoBalanceAutomationLevel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sdrsIoBalanceAutomationLevel;

    /**
     * @return Overrides the default automation settings when correcting I/O load imbalances.
     * 
     */
    public Output<Optional<String>> sdrsIoBalanceAutomationLevel() {
        return Codegen.optional(this.sdrsIoBalanceAutomationLevel);
    }
    /**
     * The I/O latency threshold, in milliseconds, that storage DRS uses to make recommendations to move disks from this
     * datastore.
     * 
     */
    @Export(name="sdrsIoLatencyThreshold", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> sdrsIoLatencyThreshold;

    /**
     * @return The I/O latency threshold, in milliseconds, that storage DRS uses to make recommendations to move disks from this
     * datastore.
     * 
     */
    public Output<Optional<Integer>> sdrsIoLatencyThreshold() {
        return Codegen.optional(this.sdrsIoLatencyThreshold);
    }
    /**
     * Enable I/O load balancing for this datastore cluster.
     * 
     */
    @Export(name="sdrsIoLoadBalanceEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> sdrsIoLoadBalanceEnabled;

    /**
     * @return Enable I/O load balancing for this datastore cluster.
     * 
     */
    public Output<Optional<Boolean>> sdrsIoLoadBalanceEnabled() {
        return Codegen.optional(this.sdrsIoLoadBalanceEnabled);
    }
    /**
     * The difference between load in datastores in the cluster before storage DRS makes recommendations to balance the load.
     * 
     */
    @Export(name="sdrsIoLoadImbalanceThreshold", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> sdrsIoLoadImbalanceThreshold;

    /**
     * @return The difference between load in datastores in the cluster before storage DRS makes recommendations to balance the load.
     * 
     */
    public Output<Optional<Integer>> sdrsIoLoadImbalanceThreshold() {
        return Codegen.optional(this.sdrsIoLoadImbalanceThreshold);
    }
    /**
     * The threshold of reservable IOPS of all virtual machines on the datastore before storage DRS makes recommendations to
     * move VMs off of a datastore.
     * 
     */
    @Export(name="sdrsIoReservableIopsThreshold", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> sdrsIoReservableIopsThreshold;

    /**
     * @return The threshold of reservable IOPS of all virtual machines on the datastore before storage DRS makes recommendations to
     * move VMs off of a datastore.
     * 
     */
    public Output<Optional<Integer>> sdrsIoReservableIopsThreshold() {
        return Codegen.optional(this.sdrsIoReservableIopsThreshold);
    }
    /**
     * The threshold, in percent, of actual estimated performance of the datastore (in IOPS) that storage DRS uses to make
     * recommendations to move VMs off of a datastore when the total reservable IOPS exceeds the threshold.
     * 
     */
    @Export(name="sdrsIoReservablePercentThreshold", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> sdrsIoReservablePercentThreshold;

    /**
     * @return The threshold, in percent, of actual estimated performance of the datastore (in IOPS) that storage DRS uses to make
     * recommendations to move VMs off of a datastore when the total reservable IOPS exceeds the threshold.
     * 
     */
    public Output<Optional<Integer>> sdrsIoReservablePercentThreshold() {
        return Codegen.optional(this.sdrsIoReservablePercentThreshold);
    }
    /**
     * The reservable IOPS threshold to use, percent in the event of automatic, or manual threshold in the event of manual.
     * 
     */
    @Export(name="sdrsIoReservableThresholdMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sdrsIoReservableThresholdMode;

    /**
     * @return The reservable IOPS threshold to use, percent in the event of automatic, or manual threshold in the event of manual.
     * 
     */
    public Output<Optional<String>> sdrsIoReservableThresholdMode() {
        return Codegen.optional(this.sdrsIoReservableThresholdMode);
    }
    /**
     * The storage DRS poll interval, in minutes.
     * 
     */
    @Export(name="sdrsLoadBalanceInterval", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> sdrsLoadBalanceInterval;

    /**
     * @return The storage DRS poll interval, in minutes.
     * 
     */
    public Output<Optional<Integer>> sdrsLoadBalanceInterval() {
        return Codegen.optional(this.sdrsLoadBalanceInterval);
    }
    /**
     * Overrides the default automation settings when correcting storage and VM policy violations.
     * 
     */
    @Export(name="sdrsPolicyEnforcementAutomationLevel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sdrsPolicyEnforcementAutomationLevel;

    /**
     * @return Overrides the default automation settings when correcting storage and VM policy violations.
     * 
     */
    public Output<Optional<String>> sdrsPolicyEnforcementAutomationLevel() {
        return Codegen.optional(this.sdrsPolicyEnforcementAutomationLevel);
    }
    /**
     * Overrides the default automation settings when correcting affinity rule violations.
     * 
     */
    @Export(name="sdrsRuleEnforcementAutomationLevel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sdrsRuleEnforcementAutomationLevel;

    /**
     * @return Overrides the default automation settings when correcting affinity rule violations.
     * 
     */
    public Output<Optional<String>> sdrsRuleEnforcementAutomationLevel() {
        return Codegen.optional(this.sdrsRuleEnforcementAutomationLevel);
    }
    /**
     * Overrides the default automation settings when correcting disk space imbalances.
     * 
     */
    @Export(name="sdrsSpaceBalanceAutomationLevel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sdrsSpaceBalanceAutomationLevel;

    /**
     * @return Overrides the default automation settings when correcting disk space imbalances.
     * 
     */
    public Output<Optional<String>> sdrsSpaceBalanceAutomationLevel() {
        return Codegen.optional(this.sdrsSpaceBalanceAutomationLevel);
    }
    /**
     * The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     * 
     */
    @Export(name="sdrsSpaceUtilizationThreshold", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> sdrsSpaceUtilizationThreshold;

    /**
     * @return The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     * 
     */
    public Output<Optional<Integer>> sdrsSpaceUtilizationThreshold() {
        return Codegen.optional(this.sdrsSpaceUtilizationThreshold);
    }
    /**
     * Overrides the default automation settings when generating recommendations for datastore evacuation.
     * 
     */
    @Export(name="sdrsVmEvacuationAutomationLevel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sdrsVmEvacuationAutomationLevel;

    /**
     * @return Overrides the default automation settings when generating recommendations for datastore evacuation.
     * 
     */
    public Output<Optional<String>> sdrsVmEvacuationAutomationLevel() {
        return Codegen.optional(this.sdrsVmEvacuationAutomationLevel);
    }
    /**
     * The IDs of any tags to attach to this resource.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return The IDs of any tags to attach to this resource.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DatastoreCluster(java.lang.String name) {
        this(name, DatastoreClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DatastoreCluster(java.lang.String name, DatastoreClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DatastoreCluster(java.lang.String name, DatastoreClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/datastoreCluster:DatastoreCluster", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DatastoreCluster(java.lang.String name, Output<java.lang.String> id, @Nullable DatastoreClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/datastoreCluster:DatastoreCluster", name, state, makeResourceOptions(options, id), false);
    }

    private static DatastoreClusterArgs makeArgs(DatastoreClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DatastoreClusterArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static DatastoreCluster get(java.lang.String name, Output<java.lang.String> id, @Nullable DatastoreClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DatastoreCluster(name, id, state, options);
    }
}

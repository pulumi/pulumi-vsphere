// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DatastoreClusterState extends com.pulumi.resources.ResourceArgs {

    public static final DatastoreClusterState Empty = new DatastoreClusterState();

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
    @Import(name="customAttributes")
    private @Nullable Output<Map<String,String>> customAttributes;

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
    public Optional<Output<Map<String,String>>> customAttributes() {
        return Optional.ofNullable(this.customAttributes);
    }

    /**
     * The managed object ID of
     * the datacenter to create the datastore cluster in. Forces a new resource if
     * changed.
     * 
     */
    @Import(name="datacenterId")
    private @Nullable Output<String> datacenterId;

    /**
     * @return The managed object ID of
     * the datacenter to create the datastore cluster in. Forces a new resource if
     * changed.
     * 
     */
    public Optional<Output<String>> datacenterId() {
        return Optional.ofNullable(this.datacenterId);
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
    @Import(name="folder")
    private @Nullable Output<String> folder;

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
    public Optional<Output<String>> folder() {
        return Optional.ofNullable(this.folder);
    }

    /**
     * The name of the datastore cluster.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the datastore cluster.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Advanced configuration options for storage DRS.
     * 
     */
    @Import(name="sdrsAdvancedOptions")
    private @Nullable Output<Map<String,String>> sdrsAdvancedOptions;

    /**
     * @return Advanced configuration options for storage DRS.
     * 
     */
    public Optional<Output<Map<String,String>>> sdrsAdvancedOptions() {
        return Optional.ofNullable(this.sdrsAdvancedOptions);
    }

    /**
     * The default automation level for all virtual machines in this storage cluster.
     * 
     */
    @Import(name="sdrsAutomationLevel")
    private @Nullable Output<String> sdrsAutomationLevel;

    /**
     * @return The default automation level for all virtual machines in this storage cluster.
     * 
     */
    public Optional<Output<String>> sdrsAutomationLevel() {
        return Optional.ofNullable(this.sdrsAutomationLevel);
    }

    /**
     * When true, storage DRS keeps VMDKs for individual VMs on the same datastore by default.
     * 
     */
    @Import(name="sdrsDefaultIntraVmAffinity")
    private @Nullable Output<Boolean> sdrsDefaultIntraVmAffinity;

    /**
     * @return When true, storage DRS keeps VMDKs for individual VMs on the same datastore by default.
     * 
     */
    public Optional<Output<Boolean>> sdrsDefaultIntraVmAffinity() {
        return Optional.ofNullable(this.sdrsDefaultIntraVmAffinity);
    }

    /**
     * Enable Storage DRS for this datastore cluster.
     * Default: `false`.
     * 
     */
    @Import(name="sdrsEnabled")
    private @Nullable Output<Boolean> sdrsEnabled;

    /**
     * @return Enable Storage DRS for this datastore cluster.
     * Default: `false`.
     * 
     */
    public Optional<Output<Boolean>> sdrsEnabled() {
        return Optional.ofNullable(this.sdrsEnabled);
    }

    /**
     * The threshold, in GB, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     * 
     */
    @Import(name="sdrsFreeSpaceThreshold")
    private @Nullable Output<Integer> sdrsFreeSpaceThreshold;

    /**
     * @return The threshold, in GB, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     * 
     */
    public Optional<Output<Integer>> sdrsFreeSpaceThreshold() {
        return Optional.ofNullable(this.sdrsFreeSpaceThreshold);
    }

    /**
     * The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
     * freeSpace, drs_free_space_threshold is used.
     * 
     */
    @Import(name="sdrsFreeSpaceThresholdMode")
    private @Nullable Output<String> sdrsFreeSpaceThresholdMode;

    /**
     * @return The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
     * freeSpace, drs_free_space_threshold is used.
     * 
     */
    public Optional<Output<String>> sdrsFreeSpaceThresholdMode() {
        return Optional.ofNullable(this.sdrsFreeSpaceThresholdMode);
    }

    /**
     * The threshold, in percent, of difference between space utilization in datastores before storage DRS makes decisions to
     * balance the space.
     * 
     */
    @Import(name="sdrsFreeSpaceUtilizationDifference")
    private @Nullable Output<Integer> sdrsFreeSpaceUtilizationDifference;

    /**
     * @return The threshold, in percent, of difference between space utilization in datastores before storage DRS makes decisions to
     * balance the space.
     * 
     */
    public Optional<Output<Integer>> sdrsFreeSpaceUtilizationDifference() {
        return Optional.ofNullable(this.sdrsFreeSpaceUtilizationDifference);
    }

    /**
     * Overrides the default automation settings when correcting I/O load imbalances.
     * 
     */
    @Import(name="sdrsIoBalanceAutomationLevel")
    private @Nullable Output<String> sdrsIoBalanceAutomationLevel;

    /**
     * @return Overrides the default automation settings when correcting I/O load imbalances.
     * 
     */
    public Optional<Output<String>> sdrsIoBalanceAutomationLevel() {
        return Optional.ofNullable(this.sdrsIoBalanceAutomationLevel);
    }

    /**
     * The I/O latency threshold, in milliseconds, that storage DRS uses to make recommendations to move disks from this
     * datastore.
     * 
     */
    @Import(name="sdrsIoLatencyThreshold")
    private @Nullable Output<Integer> sdrsIoLatencyThreshold;

    /**
     * @return The I/O latency threshold, in milliseconds, that storage DRS uses to make recommendations to move disks from this
     * datastore.
     * 
     */
    public Optional<Output<Integer>> sdrsIoLatencyThreshold() {
        return Optional.ofNullable(this.sdrsIoLatencyThreshold);
    }

    /**
     * Enable I/O load balancing for this datastore cluster.
     * 
     */
    @Import(name="sdrsIoLoadBalanceEnabled")
    private @Nullable Output<Boolean> sdrsIoLoadBalanceEnabled;

    /**
     * @return Enable I/O load balancing for this datastore cluster.
     * 
     */
    public Optional<Output<Boolean>> sdrsIoLoadBalanceEnabled() {
        return Optional.ofNullable(this.sdrsIoLoadBalanceEnabled);
    }

    /**
     * The difference between load in datastores in the cluster before storage DRS makes recommendations to balance the load.
     * 
     */
    @Import(name="sdrsIoLoadImbalanceThreshold")
    private @Nullable Output<Integer> sdrsIoLoadImbalanceThreshold;

    /**
     * @return The difference between load in datastores in the cluster before storage DRS makes recommendations to balance the load.
     * 
     */
    public Optional<Output<Integer>> sdrsIoLoadImbalanceThreshold() {
        return Optional.ofNullable(this.sdrsIoLoadImbalanceThreshold);
    }

    /**
     * The threshold of reservable IOPS of all virtual machines on the datastore before storage DRS makes recommendations to
     * move VMs off of a datastore.
     * 
     */
    @Import(name="sdrsIoReservableIopsThreshold")
    private @Nullable Output<Integer> sdrsIoReservableIopsThreshold;

    /**
     * @return The threshold of reservable IOPS of all virtual machines on the datastore before storage DRS makes recommendations to
     * move VMs off of a datastore.
     * 
     */
    public Optional<Output<Integer>> sdrsIoReservableIopsThreshold() {
        return Optional.ofNullable(this.sdrsIoReservableIopsThreshold);
    }

    /**
     * The threshold, in percent, of actual estimated performance of the datastore (in IOPS) that storage DRS uses to make
     * recommendations to move VMs off of a datastore when the total reservable IOPS exceeds the threshold.
     * 
     */
    @Import(name="sdrsIoReservablePercentThreshold")
    private @Nullable Output<Integer> sdrsIoReservablePercentThreshold;

    /**
     * @return The threshold, in percent, of actual estimated performance of the datastore (in IOPS) that storage DRS uses to make
     * recommendations to move VMs off of a datastore when the total reservable IOPS exceeds the threshold.
     * 
     */
    public Optional<Output<Integer>> sdrsIoReservablePercentThreshold() {
        return Optional.ofNullable(this.sdrsIoReservablePercentThreshold);
    }

    /**
     * The reservable IOPS threshold to use, percent in the event of automatic, or manual threshold in the event of manual.
     * 
     */
    @Import(name="sdrsIoReservableThresholdMode")
    private @Nullable Output<String> sdrsIoReservableThresholdMode;

    /**
     * @return The reservable IOPS threshold to use, percent in the event of automatic, or manual threshold in the event of manual.
     * 
     */
    public Optional<Output<String>> sdrsIoReservableThresholdMode() {
        return Optional.ofNullable(this.sdrsIoReservableThresholdMode);
    }

    /**
     * The storage DRS poll interval, in minutes.
     * 
     */
    @Import(name="sdrsLoadBalanceInterval")
    private @Nullable Output<Integer> sdrsLoadBalanceInterval;

    /**
     * @return The storage DRS poll interval, in minutes.
     * 
     */
    public Optional<Output<Integer>> sdrsLoadBalanceInterval() {
        return Optional.ofNullable(this.sdrsLoadBalanceInterval);
    }

    /**
     * Overrides the default automation settings when correcting storage and VM policy violations.
     * 
     */
    @Import(name="sdrsPolicyEnforcementAutomationLevel")
    private @Nullable Output<String> sdrsPolicyEnforcementAutomationLevel;

    /**
     * @return Overrides the default automation settings when correcting storage and VM policy violations.
     * 
     */
    public Optional<Output<String>> sdrsPolicyEnforcementAutomationLevel() {
        return Optional.ofNullable(this.sdrsPolicyEnforcementAutomationLevel);
    }

    /**
     * Overrides the default automation settings when correcting affinity rule violations.
     * 
     */
    @Import(name="sdrsRuleEnforcementAutomationLevel")
    private @Nullable Output<String> sdrsRuleEnforcementAutomationLevel;

    /**
     * @return Overrides the default automation settings when correcting affinity rule violations.
     * 
     */
    public Optional<Output<String>> sdrsRuleEnforcementAutomationLevel() {
        return Optional.ofNullable(this.sdrsRuleEnforcementAutomationLevel);
    }

    /**
     * Overrides the default automation settings when correcting disk space imbalances.
     * 
     */
    @Import(name="sdrsSpaceBalanceAutomationLevel")
    private @Nullable Output<String> sdrsSpaceBalanceAutomationLevel;

    /**
     * @return Overrides the default automation settings when correcting disk space imbalances.
     * 
     */
    public Optional<Output<String>> sdrsSpaceBalanceAutomationLevel() {
        return Optional.ofNullable(this.sdrsSpaceBalanceAutomationLevel);
    }

    /**
     * The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     * 
     */
    @Import(name="sdrsSpaceUtilizationThreshold")
    private @Nullable Output<Integer> sdrsSpaceUtilizationThreshold;

    /**
     * @return The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
     * 
     */
    public Optional<Output<Integer>> sdrsSpaceUtilizationThreshold() {
        return Optional.ofNullable(this.sdrsSpaceUtilizationThreshold);
    }

    /**
     * Overrides the default automation settings when generating recommendations for datastore evacuation.
     * 
     */
    @Import(name="sdrsVmEvacuationAutomationLevel")
    private @Nullable Output<String> sdrsVmEvacuationAutomationLevel;

    /**
     * @return Overrides the default automation settings when generating recommendations for datastore evacuation.
     * 
     */
    public Optional<Output<String>> sdrsVmEvacuationAutomationLevel() {
        return Optional.ofNullable(this.sdrsVmEvacuationAutomationLevel);
    }

    /**
     * The IDs of any tags to attach to this resource.
     * 
     * &gt; **NOTE:** Tagging support requires vCenter 6.0 or higher.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The IDs of any tags to attach to this resource.
     * 
     * &gt; **NOTE:** Tagging support requires vCenter 6.0 or higher.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private DatastoreClusterState() {}

    private DatastoreClusterState(DatastoreClusterState $) {
        this.customAttributes = $.customAttributes;
        this.datacenterId = $.datacenterId;
        this.folder = $.folder;
        this.name = $.name;
        this.sdrsAdvancedOptions = $.sdrsAdvancedOptions;
        this.sdrsAutomationLevel = $.sdrsAutomationLevel;
        this.sdrsDefaultIntraVmAffinity = $.sdrsDefaultIntraVmAffinity;
        this.sdrsEnabled = $.sdrsEnabled;
        this.sdrsFreeSpaceThreshold = $.sdrsFreeSpaceThreshold;
        this.sdrsFreeSpaceThresholdMode = $.sdrsFreeSpaceThresholdMode;
        this.sdrsFreeSpaceUtilizationDifference = $.sdrsFreeSpaceUtilizationDifference;
        this.sdrsIoBalanceAutomationLevel = $.sdrsIoBalanceAutomationLevel;
        this.sdrsIoLatencyThreshold = $.sdrsIoLatencyThreshold;
        this.sdrsIoLoadBalanceEnabled = $.sdrsIoLoadBalanceEnabled;
        this.sdrsIoLoadImbalanceThreshold = $.sdrsIoLoadImbalanceThreshold;
        this.sdrsIoReservableIopsThreshold = $.sdrsIoReservableIopsThreshold;
        this.sdrsIoReservablePercentThreshold = $.sdrsIoReservablePercentThreshold;
        this.sdrsIoReservableThresholdMode = $.sdrsIoReservableThresholdMode;
        this.sdrsLoadBalanceInterval = $.sdrsLoadBalanceInterval;
        this.sdrsPolicyEnforcementAutomationLevel = $.sdrsPolicyEnforcementAutomationLevel;
        this.sdrsRuleEnforcementAutomationLevel = $.sdrsRuleEnforcementAutomationLevel;
        this.sdrsSpaceBalanceAutomationLevel = $.sdrsSpaceBalanceAutomationLevel;
        this.sdrsSpaceUtilizationThreshold = $.sdrsSpaceUtilizationThreshold;
        this.sdrsVmEvacuationAutomationLevel = $.sdrsVmEvacuationAutomationLevel;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatastoreClusterState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatastoreClusterState $;

        public Builder() {
            $ = new DatastoreClusterState();
        }

        public Builder(DatastoreClusterState defaults) {
            $ = new DatastoreClusterState(Objects.requireNonNull(defaults));
        }

        /**
         * @param customAttributes A map of custom attribute ids to attribute
         * value strings to set for the datastore cluster. See
         * [here][docs-setting-custom-attributes] for a reference on how to set values
         * for custom attributes.
         * 
         * [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource
         * 
         * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
         * and require vCenter.
         * 
         * @return builder
         * 
         */
        public Builder customAttributes(@Nullable Output<Map<String,String>> customAttributes) {
            $.customAttributes = customAttributes;
            return this;
        }

        /**
         * @param customAttributes A map of custom attribute ids to attribute
         * value strings to set for the datastore cluster. See
         * [here][docs-setting-custom-attributes] for a reference on how to set values
         * for custom attributes.
         * 
         * [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource
         * 
         * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
         * and require vCenter.
         * 
         * @return builder
         * 
         */
        public Builder customAttributes(Map<String,String> customAttributes) {
            return customAttributes(Output.of(customAttributes));
        }

        /**
         * @param datacenterId The managed object ID of
         * the datacenter to create the datastore cluster in. Forces a new resource if
         * changed.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(@Nullable Output<String> datacenterId) {
            $.datacenterId = datacenterId;
            return this;
        }

        /**
         * @param datacenterId The managed object ID of
         * the datacenter to create the datastore cluster in. Forces a new resource if
         * changed.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(String datacenterId) {
            return datacenterId(Output.of(datacenterId));
        }

        /**
         * @param folder The relative path to a folder to put this datastore
         * cluster in.  This is a path relative to the datacenter you are deploying the
         * datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
         * `foo/bar`, The provider will place a datastore cluster named
         * `datastore-cluster-test` in a datastore folder located at
         * `/dc1/datastore/foo/bar`, with the final inventory path being
         * `/dc1/datastore/foo/bar/datastore-cluster-test`.
         * 
         * @return builder
         * 
         */
        public Builder folder(@Nullable Output<String> folder) {
            $.folder = folder;
            return this;
        }

        /**
         * @param folder The relative path to a folder to put this datastore
         * cluster in.  This is a path relative to the datacenter you are deploying the
         * datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
         * `foo/bar`, The provider will place a datastore cluster named
         * `datastore-cluster-test` in a datastore folder located at
         * `/dc1/datastore/foo/bar`, with the final inventory path being
         * `/dc1/datastore/foo/bar/datastore-cluster-test`.
         * 
         * @return builder
         * 
         */
        public Builder folder(String folder) {
            return folder(Output.of(folder));
        }

        /**
         * @param name The name of the datastore cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the datastore cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param sdrsAdvancedOptions Advanced configuration options for storage DRS.
         * 
         * @return builder
         * 
         */
        public Builder sdrsAdvancedOptions(@Nullable Output<Map<String,String>> sdrsAdvancedOptions) {
            $.sdrsAdvancedOptions = sdrsAdvancedOptions;
            return this;
        }

        /**
         * @param sdrsAdvancedOptions Advanced configuration options for storage DRS.
         * 
         * @return builder
         * 
         */
        public Builder sdrsAdvancedOptions(Map<String,String> sdrsAdvancedOptions) {
            return sdrsAdvancedOptions(Output.of(sdrsAdvancedOptions));
        }

        /**
         * @param sdrsAutomationLevel The default automation level for all virtual machines in this storage cluster.
         * 
         * @return builder
         * 
         */
        public Builder sdrsAutomationLevel(@Nullable Output<String> sdrsAutomationLevel) {
            $.sdrsAutomationLevel = sdrsAutomationLevel;
            return this;
        }

        /**
         * @param sdrsAutomationLevel The default automation level for all virtual machines in this storage cluster.
         * 
         * @return builder
         * 
         */
        public Builder sdrsAutomationLevel(String sdrsAutomationLevel) {
            return sdrsAutomationLevel(Output.of(sdrsAutomationLevel));
        }

        /**
         * @param sdrsDefaultIntraVmAffinity When true, storage DRS keeps VMDKs for individual VMs on the same datastore by default.
         * 
         * @return builder
         * 
         */
        public Builder sdrsDefaultIntraVmAffinity(@Nullable Output<Boolean> sdrsDefaultIntraVmAffinity) {
            $.sdrsDefaultIntraVmAffinity = sdrsDefaultIntraVmAffinity;
            return this;
        }

        /**
         * @param sdrsDefaultIntraVmAffinity When true, storage DRS keeps VMDKs for individual VMs on the same datastore by default.
         * 
         * @return builder
         * 
         */
        public Builder sdrsDefaultIntraVmAffinity(Boolean sdrsDefaultIntraVmAffinity) {
            return sdrsDefaultIntraVmAffinity(Output.of(sdrsDefaultIntraVmAffinity));
        }

        /**
         * @param sdrsEnabled Enable Storage DRS for this datastore cluster.
         * Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder sdrsEnabled(@Nullable Output<Boolean> sdrsEnabled) {
            $.sdrsEnabled = sdrsEnabled;
            return this;
        }

        /**
         * @param sdrsEnabled Enable Storage DRS for this datastore cluster.
         * Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder sdrsEnabled(Boolean sdrsEnabled) {
            return sdrsEnabled(Output.of(sdrsEnabled));
        }

        /**
         * @param sdrsFreeSpaceThreshold The threshold, in GB, that storage DRS uses to make decisions to migrate VMs out of a datastore.
         * 
         * @return builder
         * 
         */
        public Builder sdrsFreeSpaceThreshold(@Nullable Output<Integer> sdrsFreeSpaceThreshold) {
            $.sdrsFreeSpaceThreshold = sdrsFreeSpaceThreshold;
            return this;
        }

        /**
         * @param sdrsFreeSpaceThreshold The threshold, in GB, that storage DRS uses to make decisions to migrate VMs out of a datastore.
         * 
         * @return builder
         * 
         */
        public Builder sdrsFreeSpaceThreshold(Integer sdrsFreeSpaceThreshold) {
            return sdrsFreeSpaceThreshold(Output.of(sdrsFreeSpaceThreshold));
        }

        /**
         * @param sdrsFreeSpaceThresholdMode The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
         * freeSpace, drs_free_space_threshold is used.
         * 
         * @return builder
         * 
         */
        public Builder sdrsFreeSpaceThresholdMode(@Nullable Output<String> sdrsFreeSpaceThresholdMode) {
            $.sdrsFreeSpaceThresholdMode = sdrsFreeSpaceThresholdMode;
            return this;
        }

        /**
         * @param sdrsFreeSpaceThresholdMode The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
         * freeSpace, drs_free_space_threshold is used.
         * 
         * @return builder
         * 
         */
        public Builder sdrsFreeSpaceThresholdMode(String sdrsFreeSpaceThresholdMode) {
            return sdrsFreeSpaceThresholdMode(Output.of(sdrsFreeSpaceThresholdMode));
        }

        /**
         * @param sdrsFreeSpaceUtilizationDifference The threshold, in percent, of difference between space utilization in datastores before storage DRS makes decisions to
         * balance the space.
         * 
         * @return builder
         * 
         */
        public Builder sdrsFreeSpaceUtilizationDifference(@Nullable Output<Integer> sdrsFreeSpaceUtilizationDifference) {
            $.sdrsFreeSpaceUtilizationDifference = sdrsFreeSpaceUtilizationDifference;
            return this;
        }

        /**
         * @param sdrsFreeSpaceUtilizationDifference The threshold, in percent, of difference between space utilization in datastores before storage DRS makes decisions to
         * balance the space.
         * 
         * @return builder
         * 
         */
        public Builder sdrsFreeSpaceUtilizationDifference(Integer sdrsFreeSpaceUtilizationDifference) {
            return sdrsFreeSpaceUtilizationDifference(Output.of(sdrsFreeSpaceUtilizationDifference));
        }

        /**
         * @param sdrsIoBalanceAutomationLevel Overrides the default automation settings when correcting I/O load imbalances.
         * 
         * @return builder
         * 
         */
        public Builder sdrsIoBalanceAutomationLevel(@Nullable Output<String> sdrsIoBalanceAutomationLevel) {
            $.sdrsIoBalanceAutomationLevel = sdrsIoBalanceAutomationLevel;
            return this;
        }

        /**
         * @param sdrsIoBalanceAutomationLevel Overrides the default automation settings when correcting I/O load imbalances.
         * 
         * @return builder
         * 
         */
        public Builder sdrsIoBalanceAutomationLevel(String sdrsIoBalanceAutomationLevel) {
            return sdrsIoBalanceAutomationLevel(Output.of(sdrsIoBalanceAutomationLevel));
        }

        /**
         * @param sdrsIoLatencyThreshold The I/O latency threshold, in milliseconds, that storage DRS uses to make recommendations to move disks from this
         * datastore.
         * 
         * @return builder
         * 
         */
        public Builder sdrsIoLatencyThreshold(@Nullable Output<Integer> sdrsIoLatencyThreshold) {
            $.sdrsIoLatencyThreshold = sdrsIoLatencyThreshold;
            return this;
        }

        /**
         * @param sdrsIoLatencyThreshold The I/O latency threshold, in milliseconds, that storage DRS uses to make recommendations to move disks from this
         * datastore.
         * 
         * @return builder
         * 
         */
        public Builder sdrsIoLatencyThreshold(Integer sdrsIoLatencyThreshold) {
            return sdrsIoLatencyThreshold(Output.of(sdrsIoLatencyThreshold));
        }

        /**
         * @param sdrsIoLoadBalanceEnabled Enable I/O load balancing for this datastore cluster.
         * 
         * @return builder
         * 
         */
        public Builder sdrsIoLoadBalanceEnabled(@Nullable Output<Boolean> sdrsIoLoadBalanceEnabled) {
            $.sdrsIoLoadBalanceEnabled = sdrsIoLoadBalanceEnabled;
            return this;
        }

        /**
         * @param sdrsIoLoadBalanceEnabled Enable I/O load balancing for this datastore cluster.
         * 
         * @return builder
         * 
         */
        public Builder sdrsIoLoadBalanceEnabled(Boolean sdrsIoLoadBalanceEnabled) {
            return sdrsIoLoadBalanceEnabled(Output.of(sdrsIoLoadBalanceEnabled));
        }

        /**
         * @param sdrsIoLoadImbalanceThreshold The difference between load in datastores in the cluster before storage DRS makes recommendations to balance the load.
         * 
         * @return builder
         * 
         */
        public Builder sdrsIoLoadImbalanceThreshold(@Nullable Output<Integer> sdrsIoLoadImbalanceThreshold) {
            $.sdrsIoLoadImbalanceThreshold = sdrsIoLoadImbalanceThreshold;
            return this;
        }

        /**
         * @param sdrsIoLoadImbalanceThreshold The difference between load in datastores in the cluster before storage DRS makes recommendations to balance the load.
         * 
         * @return builder
         * 
         */
        public Builder sdrsIoLoadImbalanceThreshold(Integer sdrsIoLoadImbalanceThreshold) {
            return sdrsIoLoadImbalanceThreshold(Output.of(sdrsIoLoadImbalanceThreshold));
        }

        /**
         * @param sdrsIoReservableIopsThreshold The threshold of reservable IOPS of all virtual machines on the datastore before storage DRS makes recommendations to
         * move VMs off of a datastore.
         * 
         * @return builder
         * 
         */
        public Builder sdrsIoReservableIopsThreshold(@Nullable Output<Integer> sdrsIoReservableIopsThreshold) {
            $.sdrsIoReservableIopsThreshold = sdrsIoReservableIopsThreshold;
            return this;
        }

        /**
         * @param sdrsIoReservableIopsThreshold The threshold of reservable IOPS of all virtual machines on the datastore before storage DRS makes recommendations to
         * move VMs off of a datastore.
         * 
         * @return builder
         * 
         */
        public Builder sdrsIoReservableIopsThreshold(Integer sdrsIoReservableIopsThreshold) {
            return sdrsIoReservableIopsThreshold(Output.of(sdrsIoReservableIopsThreshold));
        }

        /**
         * @param sdrsIoReservablePercentThreshold The threshold, in percent, of actual estimated performance of the datastore (in IOPS) that storage DRS uses to make
         * recommendations to move VMs off of a datastore when the total reservable IOPS exceeds the threshold.
         * 
         * @return builder
         * 
         */
        public Builder sdrsIoReservablePercentThreshold(@Nullable Output<Integer> sdrsIoReservablePercentThreshold) {
            $.sdrsIoReservablePercentThreshold = sdrsIoReservablePercentThreshold;
            return this;
        }

        /**
         * @param sdrsIoReservablePercentThreshold The threshold, in percent, of actual estimated performance of the datastore (in IOPS) that storage DRS uses to make
         * recommendations to move VMs off of a datastore when the total reservable IOPS exceeds the threshold.
         * 
         * @return builder
         * 
         */
        public Builder sdrsIoReservablePercentThreshold(Integer sdrsIoReservablePercentThreshold) {
            return sdrsIoReservablePercentThreshold(Output.of(sdrsIoReservablePercentThreshold));
        }

        /**
         * @param sdrsIoReservableThresholdMode The reservable IOPS threshold to use, percent in the event of automatic, or manual threshold in the event of manual.
         * 
         * @return builder
         * 
         */
        public Builder sdrsIoReservableThresholdMode(@Nullable Output<String> sdrsIoReservableThresholdMode) {
            $.sdrsIoReservableThresholdMode = sdrsIoReservableThresholdMode;
            return this;
        }

        /**
         * @param sdrsIoReservableThresholdMode The reservable IOPS threshold to use, percent in the event of automatic, or manual threshold in the event of manual.
         * 
         * @return builder
         * 
         */
        public Builder sdrsIoReservableThresholdMode(String sdrsIoReservableThresholdMode) {
            return sdrsIoReservableThresholdMode(Output.of(sdrsIoReservableThresholdMode));
        }

        /**
         * @param sdrsLoadBalanceInterval The storage DRS poll interval, in minutes.
         * 
         * @return builder
         * 
         */
        public Builder sdrsLoadBalanceInterval(@Nullable Output<Integer> sdrsLoadBalanceInterval) {
            $.sdrsLoadBalanceInterval = sdrsLoadBalanceInterval;
            return this;
        }

        /**
         * @param sdrsLoadBalanceInterval The storage DRS poll interval, in minutes.
         * 
         * @return builder
         * 
         */
        public Builder sdrsLoadBalanceInterval(Integer sdrsLoadBalanceInterval) {
            return sdrsLoadBalanceInterval(Output.of(sdrsLoadBalanceInterval));
        }

        /**
         * @param sdrsPolicyEnforcementAutomationLevel Overrides the default automation settings when correcting storage and VM policy violations.
         * 
         * @return builder
         * 
         */
        public Builder sdrsPolicyEnforcementAutomationLevel(@Nullable Output<String> sdrsPolicyEnforcementAutomationLevel) {
            $.sdrsPolicyEnforcementAutomationLevel = sdrsPolicyEnforcementAutomationLevel;
            return this;
        }

        /**
         * @param sdrsPolicyEnforcementAutomationLevel Overrides the default automation settings when correcting storage and VM policy violations.
         * 
         * @return builder
         * 
         */
        public Builder sdrsPolicyEnforcementAutomationLevel(String sdrsPolicyEnforcementAutomationLevel) {
            return sdrsPolicyEnforcementAutomationLevel(Output.of(sdrsPolicyEnforcementAutomationLevel));
        }

        /**
         * @param sdrsRuleEnforcementAutomationLevel Overrides the default automation settings when correcting affinity rule violations.
         * 
         * @return builder
         * 
         */
        public Builder sdrsRuleEnforcementAutomationLevel(@Nullable Output<String> sdrsRuleEnforcementAutomationLevel) {
            $.sdrsRuleEnforcementAutomationLevel = sdrsRuleEnforcementAutomationLevel;
            return this;
        }

        /**
         * @param sdrsRuleEnforcementAutomationLevel Overrides the default automation settings when correcting affinity rule violations.
         * 
         * @return builder
         * 
         */
        public Builder sdrsRuleEnforcementAutomationLevel(String sdrsRuleEnforcementAutomationLevel) {
            return sdrsRuleEnforcementAutomationLevel(Output.of(sdrsRuleEnforcementAutomationLevel));
        }

        /**
         * @param sdrsSpaceBalanceAutomationLevel Overrides the default automation settings when correcting disk space imbalances.
         * 
         * @return builder
         * 
         */
        public Builder sdrsSpaceBalanceAutomationLevel(@Nullable Output<String> sdrsSpaceBalanceAutomationLevel) {
            $.sdrsSpaceBalanceAutomationLevel = sdrsSpaceBalanceAutomationLevel;
            return this;
        }

        /**
         * @param sdrsSpaceBalanceAutomationLevel Overrides the default automation settings when correcting disk space imbalances.
         * 
         * @return builder
         * 
         */
        public Builder sdrsSpaceBalanceAutomationLevel(String sdrsSpaceBalanceAutomationLevel) {
            return sdrsSpaceBalanceAutomationLevel(Output.of(sdrsSpaceBalanceAutomationLevel));
        }

        /**
         * @param sdrsSpaceUtilizationThreshold The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
         * 
         * @return builder
         * 
         */
        public Builder sdrsSpaceUtilizationThreshold(@Nullable Output<Integer> sdrsSpaceUtilizationThreshold) {
            $.sdrsSpaceUtilizationThreshold = sdrsSpaceUtilizationThreshold;
            return this;
        }

        /**
         * @param sdrsSpaceUtilizationThreshold The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
         * 
         * @return builder
         * 
         */
        public Builder sdrsSpaceUtilizationThreshold(Integer sdrsSpaceUtilizationThreshold) {
            return sdrsSpaceUtilizationThreshold(Output.of(sdrsSpaceUtilizationThreshold));
        }

        /**
         * @param sdrsVmEvacuationAutomationLevel Overrides the default automation settings when generating recommendations for datastore evacuation.
         * 
         * @return builder
         * 
         */
        public Builder sdrsVmEvacuationAutomationLevel(@Nullable Output<String> sdrsVmEvacuationAutomationLevel) {
            $.sdrsVmEvacuationAutomationLevel = sdrsVmEvacuationAutomationLevel;
            return this;
        }

        /**
         * @param sdrsVmEvacuationAutomationLevel Overrides the default automation settings when generating recommendations for datastore evacuation.
         * 
         * @return builder
         * 
         */
        public Builder sdrsVmEvacuationAutomationLevel(String sdrsVmEvacuationAutomationLevel) {
            return sdrsVmEvacuationAutomationLevel(Output.of(sdrsVmEvacuationAutomationLevel));
        }

        /**
         * @param tags The IDs of any tags to attach to this resource.
         * 
         * &gt; **NOTE:** Tagging support requires vCenter 6.0 or higher.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The IDs of any tags to attach to this resource.
         * 
         * &gt; **NOTE:** Tagging support requires vCenter 6.0 or higher.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags The IDs of any tags to attach to this resource.
         * 
         * &gt; **NOTE:** Tagging support requires vCenter 6.0 or higher.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        public DatastoreClusterState build() {
            return $;
        }
    }

}

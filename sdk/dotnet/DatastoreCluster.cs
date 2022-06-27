// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    [VSphereResourceType("vsphere:index/datastoreCluster:DatastoreCluster")]
    public partial class DatastoreCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// A map of custom attribute ids to attribute
        /// value strings to set for the datastore cluster. See
        /// [here][docs-setting-custom-attributes] for a reference on how to set values
        /// for custom attributes.
        /// </summary>
        [Output("customAttributes")]
        public Output<ImmutableDictionary<string, string>?> CustomAttributes { get; private set; } = null!;

        /// <summary>
        /// The managed object ID of
        /// the datacenter to create the datastore cluster in. Forces a new resource if
        /// changed.
        /// </summary>
        [Output("datacenterId")]
        public Output<string> DatacenterId { get; private set; } = null!;

        /// <summary>
        /// The relative path to a folder to put this datastore
        /// cluster in.  This is a path relative to the datacenter you are deploying the
        /// datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
        /// `foo/bar`, The provider will place a datastore cluster named
        /// `datastore-cluster-test` in a datastore folder located at
        /// `/dc1/datastore/foo/bar`, with the final inventory path being
        /// `/dc1/datastore/foo/bar/datastore-cluster-test`.
        /// </summary>
        [Output("folder")]
        public Output<string?> Folder { get; private set; } = null!;

        /// <summary>
        /// The name of the datastore cluster.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A key/value map of advanced Storage DRS
        /// settings that are not exposed via the provider or the vSphere client.
        /// </summary>
        [Output("sdrsAdvancedOptions")]
        public Output<ImmutableDictionary<string, string>?> SdrsAdvancedOptions { get; private set; } = null!;

        /// <summary>
        /// The global automation level for all
        /// virtual machines in this datastore cluster. Default: `manual`.
        /// </summary>
        [Output("sdrsAutomationLevel")]
        public Output<string?> SdrsAutomationLevel { get; private set; } = null!;

        /// <summary>
        /// When `true`, all disks in a
        /// single virtual machine will be kept on the same datastore. Default: `true`.
        /// </summary>
        [Output("sdrsDefaultIntraVmAffinity")]
        public Output<bool?> SdrsDefaultIntraVmAffinity { get; private set; } = null!;

        /// <summary>
        /// Enable Storage DRS for this datastore cluster.
        /// Default: `false`.
        /// </summary>
        [Output("sdrsEnabled")]
        public Output<bool?> SdrsEnabled { get; private set; } = null!;

        /// <summary>
        /// The free space threshold to use.
        /// When set to `utilization`, `drs_space_utilization_threshold` is used, and
        /// when set to `freeSpace`, `drs_free_space_threshold` is used. Default:
        /// `utilization`.
        /// </summary>
        [Output("sdrsFreeSpaceThreshold")]
        public Output<int?> SdrsFreeSpaceThreshold { get; private set; } = null!;

        /// <summary>
        /// The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
        /// freeSpace, drs_free_space_threshold is used.
        /// </summary>
        [Output("sdrsFreeSpaceThresholdMode")]
        public Output<string?> SdrsFreeSpaceThresholdMode { get; private set; } = null!;

        /// <summary>
        /// The threshold, in
        /// percent, of difference between space utilization in datastores before storage
        /// DRS makes decisions to balance the space. Default: `5` percent.
        /// </summary>
        [Output("sdrsFreeSpaceUtilizationDifference")]
        public Output<int?> SdrsFreeSpaceUtilizationDifference { get; private set; } = null!;

        /// <summary>
        /// Overrides the default
        /// automation settings when correcting I/O load imbalances.
        /// </summary>
        [Output("sdrsIoBalanceAutomationLevel")]
        public Output<string?> SdrsIoBalanceAutomationLevel { get; private set; } = null!;

        /// <summary>
        /// The I/O latency threshold, in
        /// milliseconds, that storage DRS uses to make recommendations to move disks
        /// from this datastore. Default: `15` seconds.
        /// </summary>
        [Output("sdrsIoLatencyThreshold")]
        public Output<int?> SdrsIoLatencyThreshold { get; private set; } = null!;

        /// <summary>
        /// Enable I/O load balancing for
        /// this datastore cluster. Default: `true`.
        /// </summary>
        [Output("sdrsIoLoadBalanceEnabled")]
        public Output<bool?> SdrsIoLoadBalanceEnabled { get; private set; } = null!;

        /// <summary>
        /// The difference between load
        /// in datastores in the cluster before storage DRS makes recommendations to
        /// balance the load. Default: `5` percent.
        /// </summary>
        [Output("sdrsIoLoadImbalanceThreshold")]
        public Output<int?> SdrsIoLoadImbalanceThreshold { get; private set; } = null!;

        /// <summary>
        /// The threshold of reservable
        /// IOPS of all virtual machines on the datastore before storage DRS makes
        /// recommendations to move VMs off of a datastore. Note that this setting should
        /// only be set if `sdrs_io_reservable_percent_threshold` cannot make an accurate
        /// estimate of the capacity of the datastores in your cluster, and should be set
        /// to roughly 50-60% of the worst case peak performance of the backing LUNs.
        /// </summary>
        [Output("sdrsIoReservableIopsThreshold")]
        public Output<int?> SdrsIoReservableIopsThreshold { get; private set; } = null!;

        /// <summary>
        /// The threshold, in
        /// percent, of actual estimated performance of the datastore (in IOPS) that
        /// storage DRS uses to make recommendations to move VMs off of a datastore when
        /// the total reservable IOPS exceeds the threshold. Default: `60` percent.
        /// </summary>
        [Output("sdrsIoReservablePercentThreshold")]
        public Output<int?> SdrsIoReservablePercentThreshold { get; private set; } = null!;

        /// <summary>
        /// The reservable IOPS
        /// threshold setting to use, `sdrs_io_reservable_percent_threshold` in the event
        /// of `automatic`, or `sdrs_io_reservable_iops_threshold` in the event of
        /// `manual`. Default: `automatic`.
        /// </summary>
        [Output("sdrsIoReservableThresholdMode")]
        public Output<string?> SdrsIoReservableThresholdMode { get; private set; } = null!;

        /// <summary>
        /// The storage DRS poll interval, in
        /// minutes. Default: `480` minutes.
        /// </summary>
        [Output("sdrsLoadBalanceInterval")]
        public Output<int?> SdrsLoadBalanceInterval { get; private set; } = null!;

        /// <summary>
        /// Overrides the default
        /// automation settings when correcting storage and VM policy violations.
        /// </summary>
        [Output("sdrsPolicyEnforcementAutomationLevel")]
        public Output<string?> SdrsPolicyEnforcementAutomationLevel { get; private set; } = null!;

        /// <summary>
        /// Overrides the default
        /// automation settings when correcting affinity rule violations.
        /// </summary>
        [Output("sdrsRuleEnforcementAutomationLevel")]
        public Output<string?> SdrsRuleEnforcementAutomationLevel { get; private set; } = null!;

        /// <summary>
        /// Overrides the default
        /// automation settings when correcting disk space imbalances.
        /// </summary>
        [Output("sdrsSpaceBalanceAutomationLevel")]
        public Output<string?> SdrsSpaceBalanceAutomationLevel { get; private set; } = null!;

        /// <summary>
        /// Runtime thresholds govern 
        /// when Storage DRS performs or recommends migrations
        /// (based on the selected automation level). Default: `80` percent.
        /// </summary>
        [Output("sdrsSpaceUtilizationThreshold")]
        public Output<int?> SdrsSpaceUtilizationThreshold { get; private set; } = null!;

        /// <summary>
        /// Overrides the default
        /// automation settings when generating recommendations for datastore evacuation.
        /// </summary>
        [Output("sdrsVmEvacuationAutomationLevel")]
        public Output<string?> SdrsVmEvacuationAutomationLevel { get; private set; } = null!;

        /// <summary>
        /// The IDs of any tags to attach to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DatastoreCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatastoreCluster(string name, DatastoreClusterArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/datastoreCluster:DatastoreCluster", name, args ?? new DatastoreClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatastoreCluster(string name, Input<string> id, DatastoreClusterState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/datastoreCluster:DatastoreCluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DatastoreCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatastoreCluster Get(string name, Input<string> id, DatastoreClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new DatastoreCluster(name, id, state, options);
        }
    }

    public sealed class DatastoreClusterArgs : Pulumi.ResourceArgs
    {
        [Input("customAttributes")]
        private InputMap<string>? _customAttributes;

        /// <summary>
        /// A map of custom attribute ids to attribute
        /// value strings to set for the datastore cluster. See
        /// [here][docs-setting-custom-attributes] for a reference on how to set values
        /// for custom attributes.
        /// </summary>
        public InputMap<string> CustomAttributes
        {
            get => _customAttributes ?? (_customAttributes = new InputMap<string>());
            set => _customAttributes = value;
        }

        /// <summary>
        /// The managed object ID of
        /// the datacenter to create the datastore cluster in. Forces a new resource if
        /// changed.
        /// </summary>
        [Input("datacenterId", required: true)]
        public Input<string> DatacenterId { get; set; } = null!;

        /// <summary>
        /// The relative path to a folder to put this datastore
        /// cluster in.  This is a path relative to the datacenter you are deploying the
        /// datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
        /// `foo/bar`, The provider will place a datastore cluster named
        /// `datastore-cluster-test` in a datastore folder located at
        /// `/dc1/datastore/foo/bar`, with the final inventory path being
        /// `/dc1/datastore/foo/bar/datastore-cluster-test`.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// The name of the datastore cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("sdrsAdvancedOptions")]
        private InputMap<string>? _sdrsAdvancedOptions;

        /// <summary>
        /// A key/value map of advanced Storage DRS
        /// settings that are not exposed via the provider or the vSphere client.
        /// </summary>
        public InputMap<string> SdrsAdvancedOptions
        {
            get => _sdrsAdvancedOptions ?? (_sdrsAdvancedOptions = new InputMap<string>());
            set => _sdrsAdvancedOptions = value;
        }

        /// <summary>
        /// The global automation level for all
        /// virtual machines in this datastore cluster. Default: `manual`.
        /// </summary>
        [Input("sdrsAutomationLevel")]
        public Input<string>? SdrsAutomationLevel { get; set; }

        /// <summary>
        /// When `true`, all disks in a
        /// single virtual machine will be kept on the same datastore. Default: `true`.
        /// </summary>
        [Input("sdrsDefaultIntraVmAffinity")]
        public Input<bool>? SdrsDefaultIntraVmAffinity { get; set; }

        /// <summary>
        /// Enable Storage DRS for this datastore cluster.
        /// Default: `false`.
        /// </summary>
        [Input("sdrsEnabled")]
        public Input<bool>? SdrsEnabled { get; set; }

        /// <summary>
        /// The free space threshold to use.
        /// When set to `utilization`, `drs_space_utilization_threshold` is used, and
        /// when set to `freeSpace`, `drs_free_space_threshold` is used. Default:
        /// `utilization`.
        /// </summary>
        [Input("sdrsFreeSpaceThreshold")]
        public Input<int>? SdrsFreeSpaceThreshold { get; set; }

        /// <summary>
        /// The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
        /// freeSpace, drs_free_space_threshold is used.
        /// </summary>
        [Input("sdrsFreeSpaceThresholdMode")]
        public Input<string>? SdrsFreeSpaceThresholdMode { get; set; }

        /// <summary>
        /// The threshold, in
        /// percent, of difference between space utilization in datastores before storage
        /// DRS makes decisions to balance the space. Default: `5` percent.
        /// </summary>
        [Input("sdrsFreeSpaceUtilizationDifference")]
        public Input<int>? SdrsFreeSpaceUtilizationDifference { get; set; }

        /// <summary>
        /// Overrides the default
        /// automation settings when correcting I/O load imbalances.
        /// </summary>
        [Input("sdrsIoBalanceAutomationLevel")]
        public Input<string>? SdrsIoBalanceAutomationLevel { get; set; }

        /// <summary>
        /// The I/O latency threshold, in
        /// milliseconds, that storage DRS uses to make recommendations to move disks
        /// from this datastore. Default: `15` seconds.
        /// </summary>
        [Input("sdrsIoLatencyThreshold")]
        public Input<int>? SdrsIoLatencyThreshold { get; set; }

        /// <summary>
        /// Enable I/O load balancing for
        /// this datastore cluster. Default: `true`.
        /// </summary>
        [Input("sdrsIoLoadBalanceEnabled")]
        public Input<bool>? SdrsIoLoadBalanceEnabled { get; set; }

        /// <summary>
        /// The difference between load
        /// in datastores in the cluster before storage DRS makes recommendations to
        /// balance the load. Default: `5` percent.
        /// </summary>
        [Input("sdrsIoLoadImbalanceThreshold")]
        public Input<int>? SdrsIoLoadImbalanceThreshold { get; set; }

        /// <summary>
        /// The threshold of reservable
        /// IOPS of all virtual machines on the datastore before storage DRS makes
        /// recommendations to move VMs off of a datastore. Note that this setting should
        /// only be set if `sdrs_io_reservable_percent_threshold` cannot make an accurate
        /// estimate of the capacity of the datastores in your cluster, and should be set
        /// to roughly 50-60% of the worst case peak performance of the backing LUNs.
        /// </summary>
        [Input("sdrsIoReservableIopsThreshold")]
        public Input<int>? SdrsIoReservableIopsThreshold { get; set; }

        /// <summary>
        /// The threshold, in
        /// percent, of actual estimated performance of the datastore (in IOPS) that
        /// storage DRS uses to make recommendations to move VMs off of a datastore when
        /// the total reservable IOPS exceeds the threshold. Default: `60` percent.
        /// </summary>
        [Input("sdrsIoReservablePercentThreshold")]
        public Input<int>? SdrsIoReservablePercentThreshold { get; set; }

        /// <summary>
        /// The reservable IOPS
        /// threshold setting to use, `sdrs_io_reservable_percent_threshold` in the event
        /// of `automatic`, or `sdrs_io_reservable_iops_threshold` in the event of
        /// `manual`. Default: `automatic`.
        /// </summary>
        [Input("sdrsIoReservableThresholdMode")]
        public Input<string>? SdrsIoReservableThresholdMode { get; set; }

        /// <summary>
        /// The storage DRS poll interval, in
        /// minutes. Default: `480` minutes.
        /// </summary>
        [Input("sdrsLoadBalanceInterval")]
        public Input<int>? SdrsLoadBalanceInterval { get; set; }

        /// <summary>
        /// Overrides the default
        /// automation settings when correcting storage and VM policy violations.
        /// </summary>
        [Input("sdrsPolicyEnforcementAutomationLevel")]
        public Input<string>? SdrsPolicyEnforcementAutomationLevel { get; set; }

        /// <summary>
        /// Overrides the default
        /// automation settings when correcting affinity rule violations.
        /// </summary>
        [Input("sdrsRuleEnforcementAutomationLevel")]
        public Input<string>? SdrsRuleEnforcementAutomationLevel { get; set; }

        /// <summary>
        /// Overrides the default
        /// automation settings when correcting disk space imbalances.
        /// </summary>
        [Input("sdrsSpaceBalanceAutomationLevel")]
        public Input<string>? SdrsSpaceBalanceAutomationLevel { get; set; }

        /// <summary>
        /// Runtime thresholds govern 
        /// when Storage DRS performs or recommends migrations
        /// (based on the selected automation level). Default: `80` percent.
        /// </summary>
        [Input("sdrsSpaceUtilizationThreshold")]
        public Input<int>? SdrsSpaceUtilizationThreshold { get; set; }

        /// <summary>
        /// Overrides the default
        /// automation settings when generating recommendations for datastore evacuation.
        /// </summary>
        [Input("sdrsVmEvacuationAutomationLevel")]
        public Input<string>? SdrsVmEvacuationAutomationLevel { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The IDs of any tags to attach to this resource.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public DatastoreClusterArgs()
        {
        }
    }

    public sealed class DatastoreClusterState : Pulumi.ResourceArgs
    {
        [Input("customAttributes")]
        private InputMap<string>? _customAttributes;

        /// <summary>
        /// A map of custom attribute ids to attribute
        /// value strings to set for the datastore cluster. See
        /// [here][docs-setting-custom-attributes] for a reference on how to set values
        /// for custom attributes.
        /// </summary>
        public InputMap<string> CustomAttributes
        {
            get => _customAttributes ?? (_customAttributes = new InputMap<string>());
            set => _customAttributes = value;
        }

        /// <summary>
        /// The managed object ID of
        /// the datacenter to create the datastore cluster in. Forces a new resource if
        /// changed.
        /// </summary>
        [Input("datacenterId")]
        public Input<string>? DatacenterId { get; set; }

        /// <summary>
        /// The relative path to a folder to put this datastore
        /// cluster in.  This is a path relative to the datacenter you are deploying the
        /// datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
        /// `foo/bar`, The provider will place a datastore cluster named
        /// `datastore-cluster-test` in a datastore folder located at
        /// `/dc1/datastore/foo/bar`, with the final inventory path being
        /// `/dc1/datastore/foo/bar/datastore-cluster-test`.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// The name of the datastore cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("sdrsAdvancedOptions")]
        private InputMap<string>? _sdrsAdvancedOptions;

        /// <summary>
        /// A key/value map of advanced Storage DRS
        /// settings that are not exposed via the provider or the vSphere client.
        /// </summary>
        public InputMap<string> SdrsAdvancedOptions
        {
            get => _sdrsAdvancedOptions ?? (_sdrsAdvancedOptions = new InputMap<string>());
            set => _sdrsAdvancedOptions = value;
        }

        /// <summary>
        /// The global automation level for all
        /// virtual machines in this datastore cluster. Default: `manual`.
        /// </summary>
        [Input("sdrsAutomationLevel")]
        public Input<string>? SdrsAutomationLevel { get; set; }

        /// <summary>
        /// When `true`, all disks in a
        /// single virtual machine will be kept on the same datastore. Default: `true`.
        /// </summary>
        [Input("sdrsDefaultIntraVmAffinity")]
        public Input<bool>? SdrsDefaultIntraVmAffinity { get; set; }

        /// <summary>
        /// Enable Storage DRS for this datastore cluster.
        /// Default: `false`.
        /// </summary>
        [Input("sdrsEnabled")]
        public Input<bool>? SdrsEnabled { get; set; }

        /// <summary>
        /// The free space threshold to use.
        /// When set to `utilization`, `drs_space_utilization_threshold` is used, and
        /// when set to `freeSpace`, `drs_free_space_threshold` is used. Default:
        /// `utilization`.
        /// </summary>
        [Input("sdrsFreeSpaceThreshold")]
        public Input<int>? SdrsFreeSpaceThreshold { get; set; }

        /// <summary>
        /// The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
        /// freeSpace, drs_free_space_threshold is used.
        /// </summary>
        [Input("sdrsFreeSpaceThresholdMode")]
        public Input<string>? SdrsFreeSpaceThresholdMode { get; set; }

        /// <summary>
        /// The threshold, in
        /// percent, of difference between space utilization in datastores before storage
        /// DRS makes decisions to balance the space. Default: `5` percent.
        /// </summary>
        [Input("sdrsFreeSpaceUtilizationDifference")]
        public Input<int>? SdrsFreeSpaceUtilizationDifference { get; set; }

        /// <summary>
        /// Overrides the default
        /// automation settings when correcting I/O load imbalances.
        /// </summary>
        [Input("sdrsIoBalanceAutomationLevel")]
        public Input<string>? SdrsIoBalanceAutomationLevel { get; set; }

        /// <summary>
        /// The I/O latency threshold, in
        /// milliseconds, that storage DRS uses to make recommendations to move disks
        /// from this datastore. Default: `15` seconds.
        /// </summary>
        [Input("sdrsIoLatencyThreshold")]
        public Input<int>? SdrsIoLatencyThreshold { get; set; }

        /// <summary>
        /// Enable I/O load balancing for
        /// this datastore cluster. Default: `true`.
        /// </summary>
        [Input("sdrsIoLoadBalanceEnabled")]
        public Input<bool>? SdrsIoLoadBalanceEnabled { get; set; }

        /// <summary>
        /// The difference between load
        /// in datastores in the cluster before storage DRS makes recommendations to
        /// balance the load. Default: `5` percent.
        /// </summary>
        [Input("sdrsIoLoadImbalanceThreshold")]
        public Input<int>? SdrsIoLoadImbalanceThreshold { get; set; }

        /// <summary>
        /// The threshold of reservable
        /// IOPS of all virtual machines on the datastore before storage DRS makes
        /// recommendations to move VMs off of a datastore. Note that this setting should
        /// only be set if `sdrs_io_reservable_percent_threshold` cannot make an accurate
        /// estimate of the capacity of the datastores in your cluster, and should be set
        /// to roughly 50-60% of the worst case peak performance of the backing LUNs.
        /// </summary>
        [Input("sdrsIoReservableIopsThreshold")]
        public Input<int>? SdrsIoReservableIopsThreshold { get; set; }

        /// <summary>
        /// The threshold, in
        /// percent, of actual estimated performance of the datastore (in IOPS) that
        /// storage DRS uses to make recommendations to move VMs off of a datastore when
        /// the total reservable IOPS exceeds the threshold. Default: `60` percent.
        /// </summary>
        [Input("sdrsIoReservablePercentThreshold")]
        public Input<int>? SdrsIoReservablePercentThreshold { get; set; }

        /// <summary>
        /// The reservable IOPS
        /// threshold setting to use, `sdrs_io_reservable_percent_threshold` in the event
        /// of `automatic`, or `sdrs_io_reservable_iops_threshold` in the event of
        /// `manual`. Default: `automatic`.
        /// </summary>
        [Input("sdrsIoReservableThresholdMode")]
        public Input<string>? SdrsIoReservableThresholdMode { get; set; }

        /// <summary>
        /// The storage DRS poll interval, in
        /// minutes. Default: `480` minutes.
        /// </summary>
        [Input("sdrsLoadBalanceInterval")]
        public Input<int>? SdrsLoadBalanceInterval { get; set; }

        /// <summary>
        /// Overrides the default
        /// automation settings when correcting storage and VM policy violations.
        /// </summary>
        [Input("sdrsPolicyEnforcementAutomationLevel")]
        public Input<string>? SdrsPolicyEnforcementAutomationLevel { get; set; }

        /// <summary>
        /// Overrides the default
        /// automation settings when correcting affinity rule violations.
        /// </summary>
        [Input("sdrsRuleEnforcementAutomationLevel")]
        public Input<string>? SdrsRuleEnforcementAutomationLevel { get; set; }

        /// <summary>
        /// Overrides the default
        /// automation settings when correcting disk space imbalances.
        /// </summary>
        [Input("sdrsSpaceBalanceAutomationLevel")]
        public Input<string>? SdrsSpaceBalanceAutomationLevel { get; set; }

        /// <summary>
        /// Runtime thresholds govern 
        /// when Storage DRS performs or recommends migrations
        /// (based on the selected automation level). Default: `80` percent.
        /// </summary>
        [Input("sdrsSpaceUtilizationThreshold")]
        public Input<int>? SdrsSpaceUtilizationThreshold { get; set; }

        /// <summary>
        /// Overrides the default
        /// automation settings when generating recommendations for datastore evacuation.
        /// </summary>
        [Input("sdrsVmEvacuationAutomationLevel")]
        public Input<string>? SdrsVmEvacuationAutomationLevel { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The IDs of any tags to attach to this resource.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public DatastoreClusterState()
        {
        }
    }
}

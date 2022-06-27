// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatastoreCluster struct {
	pulumi.CustomResourceState

	// A map of custom attribute ids to attribute
	// value strings to set for the datastore cluster. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes pulumi.StringMapOutput `pulumi:"customAttributes"`
	// The managed object ID of
	// the datacenter to create the datastore cluster in. Forces a new resource if
	// changed.
	DatacenterId pulumi.StringOutput `pulumi:"datacenterId"`
	// The relative path to a folder to put this datastore
	// cluster in.  This is a path relative to the datacenter you are deploying the
	// datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
	// `foo/bar`, The provider will place a datastore cluster named
	// `datastore-cluster-test` in a datastore folder located at
	// `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/datastore-cluster-test`.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// The name of the datastore cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// A key/value map of advanced Storage DRS
	// settings that are not exposed via the provider or the vSphere client.
	SdrsAdvancedOptions pulumi.StringMapOutput `pulumi:"sdrsAdvancedOptions"`
	// The global automation level for all
	// virtual machines in this datastore cluster. Default: `manual`.
	SdrsAutomationLevel pulumi.StringPtrOutput `pulumi:"sdrsAutomationLevel"`
	// When `true`, all disks in a
	// single virtual machine will be kept on the same datastore. Default: `true`.
	SdrsDefaultIntraVmAffinity pulumi.BoolPtrOutput `pulumi:"sdrsDefaultIntraVmAffinity"`
	// Enable Storage DRS for this datastore cluster.
	// Default: `false`.
	SdrsEnabled pulumi.BoolPtrOutput `pulumi:"sdrsEnabled"`
	// The free space threshold to use.
	// When set to `utilization`, `drsSpaceUtilizationThreshold` is used, and
	// when set to `freeSpace`, `drsFreeSpaceThreshold` is used. Default:
	// `utilization`.
	SdrsFreeSpaceThreshold pulumi.IntPtrOutput `pulumi:"sdrsFreeSpaceThreshold"`
	// The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
	// freeSpace, drs_free_space_threshold is used.
	SdrsFreeSpaceThresholdMode pulumi.StringPtrOutput `pulumi:"sdrsFreeSpaceThresholdMode"`
	// The threshold, in
	// percent, of difference between space utilization in datastores before storage
	// DRS makes decisions to balance the space. Default: `5` percent.
	SdrsFreeSpaceUtilizationDifference pulumi.IntPtrOutput `pulumi:"sdrsFreeSpaceUtilizationDifference"`
	// Overrides the default
	// automation settings when correcting I/O load imbalances.
	SdrsIoBalanceAutomationLevel pulumi.StringPtrOutput `pulumi:"sdrsIoBalanceAutomationLevel"`
	// The I/O latency threshold, in
	// milliseconds, that storage DRS uses to make recommendations to move disks
	// from this datastore. Default: `15` seconds.
	SdrsIoLatencyThreshold pulumi.IntPtrOutput `pulumi:"sdrsIoLatencyThreshold"`
	// Enable I/O load balancing for
	// this datastore cluster. Default: `true`.
	SdrsIoLoadBalanceEnabled pulumi.BoolPtrOutput `pulumi:"sdrsIoLoadBalanceEnabled"`
	// The difference between load
	// in datastores in the cluster before storage DRS makes recommendations to
	// balance the load. Default: `5` percent.
	SdrsIoLoadImbalanceThreshold pulumi.IntPtrOutput `pulumi:"sdrsIoLoadImbalanceThreshold"`
	// The threshold of reservable
	// IOPS of all virtual machines on the datastore before storage DRS makes
	// recommendations to move VMs off of a datastore. Note that this setting should
	// only be set if `sdrsIoReservablePercentThreshold` cannot make an accurate
	// estimate of the capacity of the datastores in your cluster, and should be set
	// to roughly 50-60% of the worst case peak performance of the backing LUNs.
	SdrsIoReservableIopsThreshold pulumi.IntPtrOutput `pulumi:"sdrsIoReservableIopsThreshold"`
	// The threshold, in
	// percent, of actual estimated performance of the datastore (in IOPS) that
	// storage DRS uses to make recommendations to move VMs off of a datastore when
	// the total reservable IOPS exceeds the threshold. Default: `60` percent.
	SdrsIoReservablePercentThreshold pulumi.IntPtrOutput `pulumi:"sdrsIoReservablePercentThreshold"`
	// The reservable IOPS
	// threshold setting to use, `sdrsIoReservablePercentThreshold` in the event
	// of `automatic`, or `sdrsIoReservableIopsThreshold` in the event of
	// `manual`. Default: `automatic`.
	SdrsIoReservableThresholdMode pulumi.StringPtrOutput `pulumi:"sdrsIoReservableThresholdMode"`
	// The storage DRS poll interval, in
	// minutes. Default: `480` minutes.
	SdrsLoadBalanceInterval pulumi.IntPtrOutput `pulumi:"sdrsLoadBalanceInterval"`
	// Overrides the default
	// automation settings when correcting storage and VM policy violations.
	SdrsPolicyEnforcementAutomationLevel pulumi.StringPtrOutput `pulumi:"sdrsPolicyEnforcementAutomationLevel"`
	// Overrides the default
	// automation settings when correcting affinity rule violations.
	SdrsRuleEnforcementAutomationLevel pulumi.StringPtrOutput `pulumi:"sdrsRuleEnforcementAutomationLevel"`
	// Overrides the default
	// automation settings when correcting disk space imbalances.
	SdrsSpaceBalanceAutomationLevel pulumi.StringPtrOutput `pulumi:"sdrsSpaceBalanceAutomationLevel"`
	// Runtime thresholds govern
	// when Storage DRS performs or recommends migrations
	// (based on the selected automation level). Default: `80` percent.
	SdrsSpaceUtilizationThreshold pulumi.IntPtrOutput `pulumi:"sdrsSpaceUtilizationThreshold"`
	// Overrides the default
	// automation settings when generating recommendations for datastore evacuation.
	SdrsVmEvacuationAutomationLevel pulumi.StringPtrOutput `pulumi:"sdrsVmEvacuationAutomationLevel"`
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewDatastoreCluster registers a new resource with the given unique name, arguments, and options.
func NewDatastoreCluster(ctx *pulumi.Context,
	name string, args *DatastoreClusterArgs, opts ...pulumi.ResourceOption) (*DatastoreCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatacenterId == nil {
		return nil, errors.New("invalid value for required argument 'DatacenterId'")
	}
	var resource DatastoreCluster
	err := ctx.RegisterResource("vsphere:index/datastoreCluster:DatastoreCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatastoreCluster gets an existing DatastoreCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatastoreCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatastoreClusterState, opts ...pulumi.ResourceOption) (*DatastoreCluster, error) {
	var resource DatastoreCluster
	err := ctx.ReadResource("vsphere:index/datastoreCluster:DatastoreCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatastoreCluster resources.
type datastoreClusterState struct {
	// A map of custom attribute ids to attribute
	// value strings to set for the datastore cluster. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The managed object ID of
	// the datacenter to create the datastore cluster in. Forces a new resource if
	// changed.
	DatacenterId *string `pulumi:"datacenterId"`
	// The relative path to a folder to put this datastore
	// cluster in.  This is a path relative to the datacenter you are deploying the
	// datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
	// `foo/bar`, The provider will place a datastore cluster named
	// `datastore-cluster-test` in a datastore folder located at
	// `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/datastore-cluster-test`.
	Folder *string `pulumi:"folder"`
	// The name of the datastore cluster.
	Name *string `pulumi:"name"`
	// A key/value map of advanced Storage DRS
	// settings that are not exposed via the provider or the vSphere client.
	SdrsAdvancedOptions map[string]string `pulumi:"sdrsAdvancedOptions"`
	// The global automation level for all
	// virtual machines in this datastore cluster. Default: `manual`.
	SdrsAutomationLevel *string `pulumi:"sdrsAutomationLevel"`
	// When `true`, all disks in a
	// single virtual machine will be kept on the same datastore. Default: `true`.
	SdrsDefaultIntraVmAffinity *bool `pulumi:"sdrsDefaultIntraVmAffinity"`
	// Enable Storage DRS for this datastore cluster.
	// Default: `false`.
	SdrsEnabled *bool `pulumi:"sdrsEnabled"`
	// The free space threshold to use.
	// When set to `utilization`, `drsSpaceUtilizationThreshold` is used, and
	// when set to `freeSpace`, `drsFreeSpaceThreshold` is used. Default:
	// `utilization`.
	SdrsFreeSpaceThreshold *int `pulumi:"sdrsFreeSpaceThreshold"`
	// The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
	// freeSpace, drs_free_space_threshold is used.
	SdrsFreeSpaceThresholdMode *string `pulumi:"sdrsFreeSpaceThresholdMode"`
	// The threshold, in
	// percent, of difference between space utilization in datastores before storage
	// DRS makes decisions to balance the space. Default: `5` percent.
	SdrsFreeSpaceUtilizationDifference *int `pulumi:"sdrsFreeSpaceUtilizationDifference"`
	// Overrides the default
	// automation settings when correcting I/O load imbalances.
	SdrsIoBalanceAutomationLevel *string `pulumi:"sdrsIoBalanceAutomationLevel"`
	// The I/O latency threshold, in
	// milliseconds, that storage DRS uses to make recommendations to move disks
	// from this datastore. Default: `15` seconds.
	SdrsIoLatencyThreshold *int `pulumi:"sdrsIoLatencyThreshold"`
	// Enable I/O load balancing for
	// this datastore cluster. Default: `true`.
	SdrsIoLoadBalanceEnabled *bool `pulumi:"sdrsIoLoadBalanceEnabled"`
	// The difference between load
	// in datastores in the cluster before storage DRS makes recommendations to
	// balance the load. Default: `5` percent.
	SdrsIoLoadImbalanceThreshold *int `pulumi:"sdrsIoLoadImbalanceThreshold"`
	// The threshold of reservable
	// IOPS of all virtual machines on the datastore before storage DRS makes
	// recommendations to move VMs off of a datastore. Note that this setting should
	// only be set if `sdrsIoReservablePercentThreshold` cannot make an accurate
	// estimate of the capacity of the datastores in your cluster, and should be set
	// to roughly 50-60% of the worst case peak performance of the backing LUNs.
	SdrsIoReservableIopsThreshold *int `pulumi:"sdrsIoReservableIopsThreshold"`
	// The threshold, in
	// percent, of actual estimated performance of the datastore (in IOPS) that
	// storage DRS uses to make recommendations to move VMs off of a datastore when
	// the total reservable IOPS exceeds the threshold. Default: `60` percent.
	SdrsIoReservablePercentThreshold *int `pulumi:"sdrsIoReservablePercentThreshold"`
	// The reservable IOPS
	// threshold setting to use, `sdrsIoReservablePercentThreshold` in the event
	// of `automatic`, or `sdrsIoReservableIopsThreshold` in the event of
	// `manual`. Default: `automatic`.
	SdrsIoReservableThresholdMode *string `pulumi:"sdrsIoReservableThresholdMode"`
	// The storage DRS poll interval, in
	// minutes. Default: `480` minutes.
	SdrsLoadBalanceInterval *int `pulumi:"sdrsLoadBalanceInterval"`
	// Overrides the default
	// automation settings when correcting storage and VM policy violations.
	SdrsPolicyEnforcementAutomationLevel *string `pulumi:"sdrsPolicyEnforcementAutomationLevel"`
	// Overrides the default
	// automation settings when correcting affinity rule violations.
	SdrsRuleEnforcementAutomationLevel *string `pulumi:"sdrsRuleEnforcementAutomationLevel"`
	// Overrides the default
	// automation settings when correcting disk space imbalances.
	SdrsSpaceBalanceAutomationLevel *string `pulumi:"sdrsSpaceBalanceAutomationLevel"`
	// Runtime thresholds govern
	// when Storage DRS performs or recommends migrations
	// (based on the selected automation level). Default: `80` percent.
	SdrsSpaceUtilizationThreshold *int `pulumi:"sdrsSpaceUtilizationThreshold"`
	// Overrides the default
	// automation settings when generating recommendations for datastore evacuation.
	SdrsVmEvacuationAutomationLevel *string `pulumi:"sdrsVmEvacuationAutomationLevel"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
}

type DatastoreClusterState struct {
	// A map of custom attribute ids to attribute
	// value strings to set for the datastore cluster. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes pulumi.StringMapInput
	// The managed object ID of
	// the datacenter to create the datastore cluster in. Forces a new resource if
	// changed.
	DatacenterId pulumi.StringPtrInput
	// The relative path to a folder to put this datastore
	// cluster in.  This is a path relative to the datacenter you are deploying the
	// datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
	// `foo/bar`, The provider will place a datastore cluster named
	// `datastore-cluster-test` in a datastore folder located at
	// `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/datastore-cluster-test`.
	Folder pulumi.StringPtrInput
	// The name of the datastore cluster.
	Name pulumi.StringPtrInput
	// A key/value map of advanced Storage DRS
	// settings that are not exposed via the provider or the vSphere client.
	SdrsAdvancedOptions pulumi.StringMapInput
	// The global automation level for all
	// virtual machines in this datastore cluster. Default: `manual`.
	SdrsAutomationLevel pulumi.StringPtrInput
	// When `true`, all disks in a
	// single virtual machine will be kept on the same datastore. Default: `true`.
	SdrsDefaultIntraVmAffinity pulumi.BoolPtrInput
	// Enable Storage DRS for this datastore cluster.
	// Default: `false`.
	SdrsEnabled pulumi.BoolPtrInput
	// The free space threshold to use.
	// When set to `utilization`, `drsSpaceUtilizationThreshold` is used, and
	// when set to `freeSpace`, `drsFreeSpaceThreshold` is used. Default:
	// `utilization`.
	SdrsFreeSpaceThreshold pulumi.IntPtrInput
	// The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
	// freeSpace, drs_free_space_threshold is used.
	SdrsFreeSpaceThresholdMode pulumi.StringPtrInput
	// The threshold, in
	// percent, of difference between space utilization in datastores before storage
	// DRS makes decisions to balance the space. Default: `5` percent.
	SdrsFreeSpaceUtilizationDifference pulumi.IntPtrInput
	// Overrides the default
	// automation settings when correcting I/O load imbalances.
	SdrsIoBalanceAutomationLevel pulumi.StringPtrInput
	// The I/O latency threshold, in
	// milliseconds, that storage DRS uses to make recommendations to move disks
	// from this datastore. Default: `15` seconds.
	SdrsIoLatencyThreshold pulumi.IntPtrInput
	// Enable I/O load balancing for
	// this datastore cluster. Default: `true`.
	SdrsIoLoadBalanceEnabled pulumi.BoolPtrInput
	// The difference between load
	// in datastores in the cluster before storage DRS makes recommendations to
	// balance the load. Default: `5` percent.
	SdrsIoLoadImbalanceThreshold pulumi.IntPtrInput
	// The threshold of reservable
	// IOPS of all virtual machines on the datastore before storage DRS makes
	// recommendations to move VMs off of a datastore. Note that this setting should
	// only be set if `sdrsIoReservablePercentThreshold` cannot make an accurate
	// estimate of the capacity of the datastores in your cluster, and should be set
	// to roughly 50-60% of the worst case peak performance of the backing LUNs.
	SdrsIoReservableIopsThreshold pulumi.IntPtrInput
	// The threshold, in
	// percent, of actual estimated performance of the datastore (in IOPS) that
	// storage DRS uses to make recommendations to move VMs off of a datastore when
	// the total reservable IOPS exceeds the threshold. Default: `60` percent.
	SdrsIoReservablePercentThreshold pulumi.IntPtrInput
	// The reservable IOPS
	// threshold setting to use, `sdrsIoReservablePercentThreshold` in the event
	// of `automatic`, or `sdrsIoReservableIopsThreshold` in the event of
	// `manual`. Default: `automatic`.
	SdrsIoReservableThresholdMode pulumi.StringPtrInput
	// The storage DRS poll interval, in
	// minutes. Default: `480` minutes.
	SdrsLoadBalanceInterval pulumi.IntPtrInput
	// Overrides the default
	// automation settings when correcting storage and VM policy violations.
	SdrsPolicyEnforcementAutomationLevel pulumi.StringPtrInput
	// Overrides the default
	// automation settings when correcting affinity rule violations.
	SdrsRuleEnforcementAutomationLevel pulumi.StringPtrInput
	// Overrides the default
	// automation settings when correcting disk space imbalances.
	SdrsSpaceBalanceAutomationLevel pulumi.StringPtrInput
	// Runtime thresholds govern
	// when Storage DRS performs or recommends migrations
	// (based on the selected automation level). Default: `80` percent.
	SdrsSpaceUtilizationThreshold pulumi.IntPtrInput
	// Overrides the default
	// automation settings when generating recommendations for datastore evacuation.
	SdrsVmEvacuationAutomationLevel pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
}

func (DatastoreClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreClusterState)(nil)).Elem()
}

type datastoreClusterArgs struct {
	// A map of custom attribute ids to attribute
	// value strings to set for the datastore cluster. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The managed object ID of
	// the datacenter to create the datastore cluster in. Forces a new resource if
	// changed.
	DatacenterId string `pulumi:"datacenterId"`
	// The relative path to a folder to put this datastore
	// cluster in.  This is a path relative to the datacenter you are deploying the
	// datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
	// `foo/bar`, The provider will place a datastore cluster named
	// `datastore-cluster-test` in a datastore folder located at
	// `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/datastore-cluster-test`.
	Folder *string `pulumi:"folder"`
	// The name of the datastore cluster.
	Name *string `pulumi:"name"`
	// A key/value map of advanced Storage DRS
	// settings that are not exposed via the provider or the vSphere client.
	SdrsAdvancedOptions map[string]string `pulumi:"sdrsAdvancedOptions"`
	// The global automation level for all
	// virtual machines in this datastore cluster. Default: `manual`.
	SdrsAutomationLevel *string `pulumi:"sdrsAutomationLevel"`
	// When `true`, all disks in a
	// single virtual machine will be kept on the same datastore. Default: `true`.
	SdrsDefaultIntraVmAffinity *bool `pulumi:"sdrsDefaultIntraVmAffinity"`
	// Enable Storage DRS for this datastore cluster.
	// Default: `false`.
	SdrsEnabled *bool `pulumi:"sdrsEnabled"`
	// The free space threshold to use.
	// When set to `utilization`, `drsSpaceUtilizationThreshold` is used, and
	// when set to `freeSpace`, `drsFreeSpaceThreshold` is used. Default:
	// `utilization`.
	SdrsFreeSpaceThreshold *int `pulumi:"sdrsFreeSpaceThreshold"`
	// The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
	// freeSpace, drs_free_space_threshold is used.
	SdrsFreeSpaceThresholdMode *string `pulumi:"sdrsFreeSpaceThresholdMode"`
	// The threshold, in
	// percent, of difference between space utilization in datastores before storage
	// DRS makes decisions to balance the space. Default: `5` percent.
	SdrsFreeSpaceUtilizationDifference *int `pulumi:"sdrsFreeSpaceUtilizationDifference"`
	// Overrides the default
	// automation settings when correcting I/O load imbalances.
	SdrsIoBalanceAutomationLevel *string `pulumi:"sdrsIoBalanceAutomationLevel"`
	// The I/O latency threshold, in
	// milliseconds, that storage DRS uses to make recommendations to move disks
	// from this datastore. Default: `15` seconds.
	SdrsIoLatencyThreshold *int `pulumi:"sdrsIoLatencyThreshold"`
	// Enable I/O load balancing for
	// this datastore cluster. Default: `true`.
	SdrsIoLoadBalanceEnabled *bool `pulumi:"sdrsIoLoadBalanceEnabled"`
	// The difference between load
	// in datastores in the cluster before storage DRS makes recommendations to
	// balance the load. Default: `5` percent.
	SdrsIoLoadImbalanceThreshold *int `pulumi:"sdrsIoLoadImbalanceThreshold"`
	// The threshold of reservable
	// IOPS of all virtual machines on the datastore before storage DRS makes
	// recommendations to move VMs off of a datastore. Note that this setting should
	// only be set if `sdrsIoReservablePercentThreshold` cannot make an accurate
	// estimate of the capacity of the datastores in your cluster, and should be set
	// to roughly 50-60% of the worst case peak performance of the backing LUNs.
	SdrsIoReservableIopsThreshold *int `pulumi:"sdrsIoReservableIopsThreshold"`
	// The threshold, in
	// percent, of actual estimated performance of the datastore (in IOPS) that
	// storage DRS uses to make recommendations to move VMs off of a datastore when
	// the total reservable IOPS exceeds the threshold. Default: `60` percent.
	SdrsIoReservablePercentThreshold *int `pulumi:"sdrsIoReservablePercentThreshold"`
	// The reservable IOPS
	// threshold setting to use, `sdrsIoReservablePercentThreshold` in the event
	// of `automatic`, or `sdrsIoReservableIopsThreshold` in the event of
	// `manual`. Default: `automatic`.
	SdrsIoReservableThresholdMode *string `pulumi:"sdrsIoReservableThresholdMode"`
	// The storage DRS poll interval, in
	// minutes. Default: `480` minutes.
	SdrsLoadBalanceInterval *int `pulumi:"sdrsLoadBalanceInterval"`
	// Overrides the default
	// automation settings when correcting storage and VM policy violations.
	SdrsPolicyEnforcementAutomationLevel *string `pulumi:"sdrsPolicyEnforcementAutomationLevel"`
	// Overrides the default
	// automation settings when correcting affinity rule violations.
	SdrsRuleEnforcementAutomationLevel *string `pulumi:"sdrsRuleEnforcementAutomationLevel"`
	// Overrides the default
	// automation settings when correcting disk space imbalances.
	SdrsSpaceBalanceAutomationLevel *string `pulumi:"sdrsSpaceBalanceAutomationLevel"`
	// Runtime thresholds govern
	// when Storage DRS performs or recommends migrations
	// (based on the selected automation level). Default: `80` percent.
	SdrsSpaceUtilizationThreshold *int `pulumi:"sdrsSpaceUtilizationThreshold"`
	// Overrides the default
	// automation settings when generating recommendations for datastore evacuation.
	SdrsVmEvacuationAutomationLevel *string `pulumi:"sdrsVmEvacuationAutomationLevel"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a DatastoreCluster resource.
type DatastoreClusterArgs struct {
	// A map of custom attribute ids to attribute
	// value strings to set for the datastore cluster. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes pulumi.StringMapInput
	// The managed object ID of
	// the datacenter to create the datastore cluster in. Forces a new resource if
	// changed.
	DatacenterId pulumi.StringInput
	// The relative path to a folder to put this datastore
	// cluster in.  This is a path relative to the datacenter you are deploying the
	// datastore to.  Example: for the `dc1` datacenter, and a provided `folder` of
	// `foo/bar`, The provider will place a datastore cluster named
	// `datastore-cluster-test` in a datastore folder located at
	// `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/datastore-cluster-test`.
	Folder pulumi.StringPtrInput
	// The name of the datastore cluster.
	Name pulumi.StringPtrInput
	// A key/value map of advanced Storage DRS
	// settings that are not exposed via the provider or the vSphere client.
	SdrsAdvancedOptions pulumi.StringMapInput
	// The global automation level for all
	// virtual machines in this datastore cluster. Default: `manual`.
	SdrsAutomationLevel pulumi.StringPtrInput
	// When `true`, all disks in a
	// single virtual machine will be kept on the same datastore. Default: `true`.
	SdrsDefaultIntraVmAffinity pulumi.BoolPtrInput
	// Enable Storage DRS for this datastore cluster.
	// Default: `false`.
	SdrsEnabled pulumi.BoolPtrInput
	// The free space threshold to use.
	// When set to `utilization`, `drsSpaceUtilizationThreshold` is used, and
	// when set to `freeSpace`, `drsFreeSpaceThreshold` is used. Default:
	// `utilization`.
	SdrsFreeSpaceThreshold pulumi.IntPtrInput
	// The free space threshold to use. When set to utilization, drs_space_utilization_threshold is used, and when set to
	// freeSpace, drs_free_space_threshold is used.
	SdrsFreeSpaceThresholdMode pulumi.StringPtrInput
	// The threshold, in
	// percent, of difference between space utilization in datastores before storage
	// DRS makes decisions to balance the space. Default: `5` percent.
	SdrsFreeSpaceUtilizationDifference pulumi.IntPtrInput
	// Overrides the default
	// automation settings when correcting I/O load imbalances.
	SdrsIoBalanceAutomationLevel pulumi.StringPtrInput
	// The I/O latency threshold, in
	// milliseconds, that storage DRS uses to make recommendations to move disks
	// from this datastore. Default: `15` seconds.
	SdrsIoLatencyThreshold pulumi.IntPtrInput
	// Enable I/O load balancing for
	// this datastore cluster. Default: `true`.
	SdrsIoLoadBalanceEnabled pulumi.BoolPtrInput
	// The difference between load
	// in datastores in the cluster before storage DRS makes recommendations to
	// balance the load. Default: `5` percent.
	SdrsIoLoadImbalanceThreshold pulumi.IntPtrInput
	// The threshold of reservable
	// IOPS of all virtual machines on the datastore before storage DRS makes
	// recommendations to move VMs off of a datastore. Note that this setting should
	// only be set if `sdrsIoReservablePercentThreshold` cannot make an accurate
	// estimate of the capacity of the datastores in your cluster, and should be set
	// to roughly 50-60% of the worst case peak performance of the backing LUNs.
	SdrsIoReservableIopsThreshold pulumi.IntPtrInput
	// The threshold, in
	// percent, of actual estimated performance of the datastore (in IOPS) that
	// storage DRS uses to make recommendations to move VMs off of a datastore when
	// the total reservable IOPS exceeds the threshold. Default: `60` percent.
	SdrsIoReservablePercentThreshold pulumi.IntPtrInput
	// The reservable IOPS
	// threshold setting to use, `sdrsIoReservablePercentThreshold` in the event
	// of `automatic`, or `sdrsIoReservableIopsThreshold` in the event of
	// `manual`. Default: `automatic`.
	SdrsIoReservableThresholdMode pulumi.StringPtrInput
	// The storage DRS poll interval, in
	// minutes. Default: `480` minutes.
	SdrsLoadBalanceInterval pulumi.IntPtrInput
	// Overrides the default
	// automation settings when correcting storage and VM policy violations.
	SdrsPolicyEnforcementAutomationLevel pulumi.StringPtrInput
	// Overrides the default
	// automation settings when correcting affinity rule violations.
	SdrsRuleEnforcementAutomationLevel pulumi.StringPtrInput
	// Overrides the default
	// automation settings when correcting disk space imbalances.
	SdrsSpaceBalanceAutomationLevel pulumi.StringPtrInput
	// Runtime thresholds govern
	// when Storage DRS performs or recommends migrations
	// (based on the selected automation level). Default: `80` percent.
	SdrsSpaceUtilizationThreshold pulumi.IntPtrInput
	// Overrides the default
	// automation settings when generating recommendations for datastore evacuation.
	SdrsVmEvacuationAutomationLevel pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
}

func (DatastoreClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreClusterArgs)(nil)).Elem()
}

type DatastoreClusterInput interface {
	pulumi.Input

	ToDatastoreClusterOutput() DatastoreClusterOutput
	ToDatastoreClusterOutputWithContext(ctx context.Context) DatastoreClusterOutput
}

func (*DatastoreCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**DatastoreCluster)(nil)).Elem()
}

func (i *DatastoreCluster) ToDatastoreClusterOutput() DatastoreClusterOutput {
	return i.ToDatastoreClusterOutputWithContext(context.Background())
}

func (i *DatastoreCluster) ToDatastoreClusterOutputWithContext(ctx context.Context) DatastoreClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatastoreClusterOutput)
}

// DatastoreClusterArrayInput is an input type that accepts DatastoreClusterArray and DatastoreClusterArrayOutput values.
// You can construct a concrete instance of `DatastoreClusterArrayInput` via:
//
//          DatastoreClusterArray{ DatastoreClusterArgs{...} }
type DatastoreClusterArrayInput interface {
	pulumi.Input

	ToDatastoreClusterArrayOutput() DatastoreClusterArrayOutput
	ToDatastoreClusterArrayOutputWithContext(context.Context) DatastoreClusterArrayOutput
}

type DatastoreClusterArray []DatastoreClusterInput

func (DatastoreClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatastoreCluster)(nil)).Elem()
}

func (i DatastoreClusterArray) ToDatastoreClusterArrayOutput() DatastoreClusterArrayOutput {
	return i.ToDatastoreClusterArrayOutputWithContext(context.Background())
}

func (i DatastoreClusterArray) ToDatastoreClusterArrayOutputWithContext(ctx context.Context) DatastoreClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatastoreClusterArrayOutput)
}

// DatastoreClusterMapInput is an input type that accepts DatastoreClusterMap and DatastoreClusterMapOutput values.
// You can construct a concrete instance of `DatastoreClusterMapInput` via:
//
//          DatastoreClusterMap{ "key": DatastoreClusterArgs{...} }
type DatastoreClusterMapInput interface {
	pulumi.Input

	ToDatastoreClusterMapOutput() DatastoreClusterMapOutput
	ToDatastoreClusterMapOutputWithContext(context.Context) DatastoreClusterMapOutput
}

type DatastoreClusterMap map[string]DatastoreClusterInput

func (DatastoreClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatastoreCluster)(nil)).Elem()
}

func (i DatastoreClusterMap) ToDatastoreClusterMapOutput() DatastoreClusterMapOutput {
	return i.ToDatastoreClusterMapOutputWithContext(context.Background())
}

func (i DatastoreClusterMap) ToDatastoreClusterMapOutputWithContext(ctx context.Context) DatastoreClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatastoreClusterMapOutput)
}

type DatastoreClusterOutput struct{ *pulumi.OutputState }

func (DatastoreClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatastoreCluster)(nil)).Elem()
}

func (o DatastoreClusterOutput) ToDatastoreClusterOutput() DatastoreClusterOutput {
	return o
}

func (o DatastoreClusterOutput) ToDatastoreClusterOutputWithContext(ctx context.Context) DatastoreClusterOutput {
	return o
}

type DatastoreClusterArrayOutput struct{ *pulumi.OutputState }

func (DatastoreClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatastoreCluster)(nil)).Elem()
}

func (o DatastoreClusterArrayOutput) ToDatastoreClusterArrayOutput() DatastoreClusterArrayOutput {
	return o
}

func (o DatastoreClusterArrayOutput) ToDatastoreClusterArrayOutputWithContext(ctx context.Context) DatastoreClusterArrayOutput {
	return o
}

func (o DatastoreClusterArrayOutput) Index(i pulumi.IntInput) DatastoreClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatastoreCluster {
		return vs[0].([]*DatastoreCluster)[vs[1].(int)]
	}).(DatastoreClusterOutput)
}

type DatastoreClusterMapOutput struct{ *pulumi.OutputState }

func (DatastoreClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatastoreCluster)(nil)).Elem()
}

func (o DatastoreClusterMapOutput) ToDatastoreClusterMapOutput() DatastoreClusterMapOutput {
	return o
}

func (o DatastoreClusterMapOutput) ToDatastoreClusterMapOutputWithContext(ctx context.Context) DatastoreClusterMapOutput {
	return o
}

func (o DatastoreClusterMapOutput) MapIndex(k pulumi.StringInput) DatastoreClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatastoreCluster {
		return vs[0].(map[string]*DatastoreCluster)[vs[1].(string)]
	}).(DatastoreClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatastoreClusterInput)(nil)).Elem(), &DatastoreCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatastoreClusterArrayInput)(nil)).Elem(), DatastoreClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatastoreClusterMapInput)(nil)).Elem(), DatastoreClusterMap{})
	pulumi.RegisterOutputType(DatastoreClusterOutput{})
	pulumi.RegisterOutputType(DatastoreClusterArrayOutput{})
	pulumi.RegisterOutputType(DatastoreClusterMapOutput{})
}

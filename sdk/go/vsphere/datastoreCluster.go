// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `.DatastoreCluster` resource can be used to create and manage
// datastore clusters. This can be used to create groups of datastores with a
// shared management interface, allowing for resource control and load balancing
// through Storage DRS.
//
// For more information on vSphere datastore clusters and Storage DRS, see [this
// page][ref-vsphere-datastore-clusters].
//
// [ref-vsphere-datastore-clusters]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-598DF695-107E-406B-9C95-0AF961FC227A.html
//
// > **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
//
// > **NOTE:** Storage DRS requires a vSphere Enterprise Plus license.
type DatastoreCluster struct {
	pulumi.CustomResourceState

	// A map of custom attribute ids to attribute
	// value strings to set for the datastore cluster. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes pulumi.StringMapOutput `pulumi:"customAttributes"`
	// The [managed object ID][docs-about-morefs] of
	// the datacenter to create the datastore cluster in. Forces a new resource if
	// changed.
	DatacenterId pulumi.StringOutput `pulumi:"datacenterId"`
	// The name of the folder to locate the datastore cluster in.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// The name of the datastore cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// Advanced configuration options for storage DRS.
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
	// The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
	SdrsSpaceUtilizationThreshold pulumi.IntPtrOutput `pulumi:"sdrsSpaceUtilizationThreshold"`
	// Overrides the default
	// automation settings when generating recommendations for datastore evacuation.
	SdrsVmEvacuationAutomationLevel pulumi.StringPtrOutput `pulumi:"sdrsVmEvacuationAutomationLevel"`
	// The IDs of any tags to attach to this resource. See
	// [here][docs-applying-tags] for a reference on how to apply tags.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewDatastoreCluster registers a new resource with the given unique name, arguments, and options.
func NewDatastoreCluster(ctx *pulumi.Context,
	name string, args *DatastoreClusterArgs, opts ...pulumi.ResourceOption) (*DatastoreCluster, error) {
	if args == nil || args.DatacenterId == nil {
		return nil, errors.New("missing required argument 'DatacenterId'")
	}
	if args == nil {
		args = &DatastoreClusterArgs{}
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
	// The [managed object ID][docs-about-morefs] of
	// the datacenter to create the datastore cluster in. Forces a new resource if
	// changed.
	DatacenterId *string `pulumi:"datacenterId"`
	// The name of the folder to locate the datastore cluster in.
	Folder *string `pulumi:"folder"`
	// The name of the datastore cluster.
	Name *string `pulumi:"name"`
	// Advanced configuration options for storage DRS.
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
	// The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
	SdrsSpaceUtilizationThreshold *int `pulumi:"sdrsSpaceUtilizationThreshold"`
	// Overrides the default
	// automation settings when generating recommendations for datastore evacuation.
	SdrsVmEvacuationAutomationLevel *string `pulumi:"sdrsVmEvacuationAutomationLevel"`
	// The IDs of any tags to attach to this resource. See
	// [here][docs-applying-tags] for a reference on how to apply tags.
	Tags []string `pulumi:"tags"`
}

type DatastoreClusterState struct {
	// A map of custom attribute ids to attribute
	// value strings to set for the datastore cluster. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes pulumi.StringMapInput
	// The [managed object ID][docs-about-morefs] of
	// the datacenter to create the datastore cluster in. Forces a new resource if
	// changed.
	DatacenterId pulumi.StringPtrInput
	// The name of the folder to locate the datastore cluster in.
	Folder pulumi.StringPtrInput
	// The name of the datastore cluster.
	Name pulumi.StringPtrInput
	// Advanced configuration options for storage DRS.
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
	// The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
	SdrsSpaceUtilizationThreshold pulumi.IntPtrInput
	// Overrides the default
	// automation settings when generating recommendations for datastore evacuation.
	SdrsVmEvacuationAutomationLevel pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource. See
	// [here][docs-applying-tags] for a reference on how to apply tags.
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
	// The [managed object ID][docs-about-morefs] of
	// the datacenter to create the datastore cluster in. Forces a new resource if
	// changed.
	DatacenterId string `pulumi:"datacenterId"`
	// The name of the folder to locate the datastore cluster in.
	Folder *string `pulumi:"folder"`
	// The name of the datastore cluster.
	Name *string `pulumi:"name"`
	// Advanced configuration options for storage DRS.
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
	// The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
	SdrsSpaceUtilizationThreshold *int `pulumi:"sdrsSpaceUtilizationThreshold"`
	// Overrides the default
	// automation settings when generating recommendations for datastore evacuation.
	SdrsVmEvacuationAutomationLevel *string `pulumi:"sdrsVmEvacuationAutomationLevel"`
	// The IDs of any tags to attach to this resource. See
	// [here][docs-applying-tags] for a reference on how to apply tags.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a DatastoreCluster resource.
type DatastoreClusterArgs struct {
	// A map of custom attribute ids to attribute
	// value strings to set for the datastore cluster. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes pulumi.StringMapInput
	// The [managed object ID][docs-about-morefs] of
	// the datacenter to create the datastore cluster in. Forces a new resource if
	// changed.
	DatacenterId pulumi.StringInput
	// The name of the folder to locate the datastore cluster in.
	Folder pulumi.StringPtrInput
	// The name of the datastore cluster.
	Name pulumi.StringPtrInput
	// Advanced configuration options for storage DRS.
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
	// The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore.
	SdrsSpaceUtilizationThreshold pulumi.IntPtrInput
	// Overrides the default
	// automation settings when generating recommendations for datastore evacuation.
	SdrsVmEvacuationAutomationLevel pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource. See
	// [here][docs-applying-tags] for a reference on how to apply tags.
	Tags pulumi.StringArrayInput
}

func (DatastoreClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreClusterArgs)(nil)).Elem()
}

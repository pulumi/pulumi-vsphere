// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `HaVmOverride` resource can be used to add an override for
// vSphere HA settings on a cluster for a specific virtual machine. With this
// resource, one can control specific HA settings so that they are different than
// the cluster default, accommodating the needs of that specific virtual machine,
// while not affecting the rest of the cluster.
//
// For more information on vSphere HA, see [this page][ref-vsphere-ha-clusters].
//
// > **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
//
// ## Example Usage
//
// The example below creates a virtual machine in a cluster using the
// `VirtualMachine` resource, creating the
// virtual machine in the cluster looked up by the
// `ComputeCluster` data source.
//
// Considering a scenario where this virtual machine is of high value to the
// application or organization for which it does its work, it's been determined in
// the event of a host failure, that this should be one of the first virtual
// machines to be started by vSphere HA during recovery. Hence, it
// `haVmRestartPriority` has been set to `highest`,
// which, assuming that the default restart priority is `medium` and no other
// virtual machine has been assigned the `highest` priority, will mean that this
// VM will be started before any other virtual machine in the event of host
// failure.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			datacenter, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
//				Name: pulumi.StringRef("dc-01"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			datastore, err := vsphere.GetDatastore(ctx, &vsphere.GetDatastoreArgs{
//				Name:         "datastore1",
//				DatacenterId: pulumi.StringRef(datacenter.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			cluster, err := vsphere.LookupComputeCluster(ctx, &vsphere.LookupComputeClusterArgs{
//				Name:         "cluster-01",
//				DatacenterId: pulumi.StringRef(datacenter.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			network, err := vsphere.GetNetwork(ctx, &vsphere.GetNetworkArgs{
//				Name:         "network1",
//				DatacenterId: pulumi.StringRef(datacenter.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vm, err := vsphere.NewVirtualMachine(ctx, "vm", &vsphere.VirtualMachineArgs{
//				Name:           pulumi.String("test"),
//				ResourcePoolId: pulumi.String(cluster.ResourcePoolId),
//				DatastoreId:    pulumi.String(datastore.Id),
//				NumCpus:        pulumi.Int(2),
//				Memory:         pulumi.Int(2048),
//				GuestId:        pulumi.String("otherLinux64Guest"),
//				NetworkInterfaces: vsphere.VirtualMachineNetworkInterfaceArray{
//					&vsphere.VirtualMachineNetworkInterfaceArgs{
//						NetworkId: pulumi.String(network.Id),
//					},
//				},
//				Disks: vsphere.VirtualMachineDiskArray{
//					&vsphere.VirtualMachineDiskArgs{
//						Label: pulumi.String("disk0"),
//						Size:  pulumi.Int(20),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.NewHaVmOverride(ctx, "ha_vm_override", &vsphere.HaVmOverrideArgs{
//				ComputeClusterId:    pulumi.String(cluster.Id),
//				VirtualMachineId:    vm.ID(),
//				HaVmRestartPriority: pulumi.String("highest"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # An existing override can be imported into this resource by
//
// supplying both the path to the cluster, and the path to the virtual machine, to
//
// `pulumi import`. If no override exists, an error will be given.  An example
//
// is below:
//
// ```sh
// $ pulumi import vsphere:index/haVmOverride:HaVmOverride ha_vm_override \
// ```
//
//	'{"compute_cluster_path": "/dc1/host/cluster1", \
//
//	"virtual_machine_path": "/dc1/vm/srv1"}'
//
// [ref-vsphere-ha-clusters]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-availability.html
//
// [docs-import]: https://developer.hashicorp.com/terraform/cli/import
type HaVmOverride struct {
	pulumi.CustomResourceState

	// The managed object ID of the cluster.
	ComputeClusterId pulumi.StringOutput `pulumi:"computeClusterId"`
	// Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an
	// APD event. Can be one of useClusterDefault, none or reset.
	HaDatastoreApdRecoveryAction pulumi.StringPtrOutput `pulumi:"haDatastoreApdRecoveryAction"`
	// Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant
	// datastore. Can be one of clusterDefault, disabled, warning, restartConservative, or restartAggressive.
	HaDatastoreApdResponse pulumi.StringPtrOutput `pulumi:"haDatastoreApdResponse"`
	// Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in
	// ha_datastore_apd_response. Specify -1 to use the cluster setting.
	HaDatastoreApdResponseDelay pulumi.IntPtrOutput `pulumi:"haDatastoreApdResponseDelay"`
	// Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant
	// datastore. Can be one of clusterDefault, disabled, warning, or restartAggressive.
	HaDatastorePdlResponse pulumi.StringPtrOutput `pulumi:"haDatastorePdlResponse"`
	// The action to take on this virtual machine when a host is isolated from the rest of the cluster. Can be one of
	// clusterIsolationResponse, none, powerOff, or shutdown.
	HaHostIsolationResponse pulumi.StringPtrOutput `pulumi:"haHostIsolationResponse"`
	// If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked
	// as failed. The value is in seconds.
	HaVmFailureInterval pulumi.IntPtrOutput `pulumi:"haVmFailureInterval"`
	// The length of the reset window in which haVmMaximumResets can operate. When this window expires, no more resets are
	// attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset
	// time is allotted.
	HaVmMaximumFailureWindow pulumi.IntPtrOutput `pulumi:"haVmMaximumFailureWindow"`
	// The maximum number of resets that HA will perform to this virtual machine when responding to a failure event.
	HaVmMaximumResets pulumi.IntPtrOutput `pulumi:"haVmMaximumResets"`
	// The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats.
	HaVmMinimumUptime pulumi.IntPtrOutput `pulumi:"haVmMinimumUptime"`
	// The type of virtual machine monitoring to use for this virtual machine. Can be one of vmMonitoringDisabled,
	// vmMonitoringOnly, or vmAndAppMonitoring.
	HaVmMonitoring pulumi.StringPtrOutput `pulumi:"haVmMonitoring"`
	// Determines whether or not the cluster's default settings or the VM override settings specified in this resource are used
	// for virtual machine monitoring. The default is true (use cluster defaults) - set to false to have overrides take effect.
	HaVmMonitoringUseClusterDefaults pulumi.BoolPtrOutput `pulumi:"haVmMonitoringUseClusterDefaults"`
	// The restart priority for this virtual machine when vSphere detects a host failure. Can be one of clusterRestartPriority,
	// lowest, low, medium, high, or highest.
	HaVmRestartPriority pulumi.StringPtrOutput `pulumi:"haVmRestartPriority"`
	// The maximum time, in seconds, that vSphere HA will wait for the virtual machine to be ready. Use -1 to use the cluster
	// default.
	HaVmRestartTimeout pulumi.IntPtrOutput `pulumi:"haVmRestartTimeout"`
	// The managed object ID of the virtual machine.
	VirtualMachineId pulumi.StringOutput `pulumi:"virtualMachineId"`
}

// NewHaVmOverride registers a new resource with the given unique name, arguments, and options.
func NewHaVmOverride(ctx *pulumi.Context,
	name string, args *HaVmOverrideArgs, opts ...pulumi.ResourceOption) (*HaVmOverride, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ComputeClusterId'")
	}
	if args.VirtualMachineId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HaVmOverride
	err := ctx.RegisterResource("vsphere:index/haVmOverride:HaVmOverride", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHaVmOverride gets an existing HaVmOverride resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHaVmOverride(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HaVmOverrideState, opts ...pulumi.ResourceOption) (*HaVmOverride, error) {
	var resource HaVmOverride
	err := ctx.ReadResource("vsphere:index/haVmOverride:HaVmOverride", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HaVmOverride resources.
type haVmOverrideState struct {
	// The managed object ID of the cluster.
	ComputeClusterId *string `pulumi:"computeClusterId"`
	// Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an
	// APD event. Can be one of useClusterDefault, none or reset.
	HaDatastoreApdRecoveryAction *string `pulumi:"haDatastoreApdRecoveryAction"`
	// Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant
	// datastore. Can be one of clusterDefault, disabled, warning, restartConservative, or restartAggressive.
	HaDatastoreApdResponse *string `pulumi:"haDatastoreApdResponse"`
	// Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in
	// ha_datastore_apd_response. Specify -1 to use the cluster setting.
	HaDatastoreApdResponseDelay *int `pulumi:"haDatastoreApdResponseDelay"`
	// Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant
	// datastore. Can be one of clusterDefault, disabled, warning, or restartAggressive.
	HaDatastorePdlResponse *string `pulumi:"haDatastorePdlResponse"`
	// The action to take on this virtual machine when a host is isolated from the rest of the cluster. Can be one of
	// clusterIsolationResponse, none, powerOff, or shutdown.
	HaHostIsolationResponse *string `pulumi:"haHostIsolationResponse"`
	// If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked
	// as failed. The value is in seconds.
	HaVmFailureInterval *int `pulumi:"haVmFailureInterval"`
	// The length of the reset window in which haVmMaximumResets can operate. When this window expires, no more resets are
	// attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset
	// time is allotted.
	HaVmMaximumFailureWindow *int `pulumi:"haVmMaximumFailureWindow"`
	// The maximum number of resets that HA will perform to this virtual machine when responding to a failure event.
	HaVmMaximumResets *int `pulumi:"haVmMaximumResets"`
	// The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats.
	HaVmMinimumUptime *int `pulumi:"haVmMinimumUptime"`
	// The type of virtual machine monitoring to use for this virtual machine. Can be one of vmMonitoringDisabled,
	// vmMonitoringOnly, or vmAndAppMonitoring.
	HaVmMonitoring *string `pulumi:"haVmMonitoring"`
	// Determines whether or not the cluster's default settings or the VM override settings specified in this resource are used
	// for virtual machine monitoring. The default is true (use cluster defaults) - set to false to have overrides take effect.
	HaVmMonitoringUseClusterDefaults *bool `pulumi:"haVmMonitoringUseClusterDefaults"`
	// The restart priority for this virtual machine when vSphere detects a host failure. Can be one of clusterRestartPriority,
	// lowest, low, medium, high, or highest.
	HaVmRestartPriority *string `pulumi:"haVmRestartPriority"`
	// The maximum time, in seconds, that vSphere HA will wait for the virtual machine to be ready. Use -1 to use the cluster
	// default.
	HaVmRestartTimeout *int `pulumi:"haVmRestartTimeout"`
	// The managed object ID of the virtual machine.
	VirtualMachineId *string `pulumi:"virtualMachineId"`
}

type HaVmOverrideState struct {
	// The managed object ID of the cluster.
	ComputeClusterId pulumi.StringPtrInput
	// Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an
	// APD event. Can be one of useClusterDefault, none or reset.
	HaDatastoreApdRecoveryAction pulumi.StringPtrInput
	// Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant
	// datastore. Can be one of clusterDefault, disabled, warning, restartConservative, or restartAggressive.
	HaDatastoreApdResponse pulumi.StringPtrInput
	// Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in
	// ha_datastore_apd_response. Specify -1 to use the cluster setting.
	HaDatastoreApdResponseDelay pulumi.IntPtrInput
	// Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant
	// datastore. Can be one of clusterDefault, disabled, warning, or restartAggressive.
	HaDatastorePdlResponse pulumi.StringPtrInput
	// The action to take on this virtual machine when a host is isolated from the rest of the cluster. Can be one of
	// clusterIsolationResponse, none, powerOff, or shutdown.
	HaHostIsolationResponse pulumi.StringPtrInput
	// If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked
	// as failed. The value is in seconds.
	HaVmFailureInterval pulumi.IntPtrInput
	// The length of the reset window in which haVmMaximumResets can operate. When this window expires, no more resets are
	// attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset
	// time is allotted.
	HaVmMaximumFailureWindow pulumi.IntPtrInput
	// The maximum number of resets that HA will perform to this virtual machine when responding to a failure event.
	HaVmMaximumResets pulumi.IntPtrInput
	// The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats.
	HaVmMinimumUptime pulumi.IntPtrInput
	// The type of virtual machine monitoring to use for this virtual machine. Can be one of vmMonitoringDisabled,
	// vmMonitoringOnly, or vmAndAppMonitoring.
	HaVmMonitoring pulumi.StringPtrInput
	// Determines whether or not the cluster's default settings or the VM override settings specified in this resource are used
	// for virtual machine monitoring. The default is true (use cluster defaults) - set to false to have overrides take effect.
	HaVmMonitoringUseClusterDefaults pulumi.BoolPtrInput
	// The restart priority for this virtual machine when vSphere detects a host failure. Can be one of clusterRestartPriority,
	// lowest, low, medium, high, or highest.
	HaVmRestartPriority pulumi.StringPtrInput
	// The maximum time, in seconds, that vSphere HA will wait for the virtual machine to be ready. Use -1 to use the cluster
	// default.
	HaVmRestartTimeout pulumi.IntPtrInput
	// The managed object ID of the virtual machine.
	VirtualMachineId pulumi.StringPtrInput
}

func (HaVmOverrideState) ElementType() reflect.Type {
	return reflect.TypeOf((*haVmOverrideState)(nil)).Elem()
}

type haVmOverrideArgs struct {
	// The managed object ID of the cluster.
	ComputeClusterId string `pulumi:"computeClusterId"`
	// Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an
	// APD event. Can be one of useClusterDefault, none or reset.
	HaDatastoreApdRecoveryAction *string `pulumi:"haDatastoreApdRecoveryAction"`
	// Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant
	// datastore. Can be one of clusterDefault, disabled, warning, restartConservative, or restartAggressive.
	HaDatastoreApdResponse *string `pulumi:"haDatastoreApdResponse"`
	// Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in
	// ha_datastore_apd_response. Specify -1 to use the cluster setting.
	HaDatastoreApdResponseDelay *int `pulumi:"haDatastoreApdResponseDelay"`
	// Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant
	// datastore. Can be one of clusterDefault, disabled, warning, or restartAggressive.
	HaDatastorePdlResponse *string `pulumi:"haDatastorePdlResponse"`
	// The action to take on this virtual machine when a host is isolated from the rest of the cluster. Can be one of
	// clusterIsolationResponse, none, powerOff, or shutdown.
	HaHostIsolationResponse *string `pulumi:"haHostIsolationResponse"`
	// If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked
	// as failed. The value is in seconds.
	HaVmFailureInterval *int `pulumi:"haVmFailureInterval"`
	// The length of the reset window in which haVmMaximumResets can operate. When this window expires, no more resets are
	// attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset
	// time is allotted.
	HaVmMaximumFailureWindow *int `pulumi:"haVmMaximumFailureWindow"`
	// The maximum number of resets that HA will perform to this virtual machine when responding to a failure event.
	HaVmMaximumResets *int `pulumi:"haVmMaximumResets"`
	// The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats.
	HaVmMinimumUptime *int `pulumi:"haVmMinimumUptime"`
	// The type of virtual machine monitoring to use for this virtual machine. Can be one of vmMonitoringDisabled,
	// vmMonitoringOnly, or vmAndAppMonitoring.
	HaVmMonitoring *string `pulumi:"haVmMonitoring"`
	// Determines whether or not the cluster's default settings or the VM override settings specified in this resource are used
	// for virtual machine monitoring. The default is true (use cluster defaults) - set to false to have overrides take effect.
	HaVmMonitoringUseClusterDefaults *bool `pulumi:"haVmMonitoringUseClusterDefaults"`
	// The restart priority for this virtual machine when vSphere detects a host failure. Can be one of clusterRestartPriority,
	// lowest, low, medium, high, or highest.
	HaVmRestartPriority *string `pulumi:"haVmRestartPriority"`
	// The maximum time, in seconds, that vSphere HA will wait for the virtual machine to be ready. Use -1 to use the cluster
	// default.
	HaVmRestartTimeout *int `pulumi:"haVmRestartTimeout"`
	// The managed object ID of the virtual machine.
	VirtualMachineId string `pulumi:"virtualMachineId"`
}

// The set of arguments for constructing a HaVmOverride resource.
type HaVmOverrideArgs struct {
	// The managed object ID of the cluster.
	ComputeClusterId pulumi.StringInput
	// Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an
	// APD event. Can be one of useClusterDefault, none or reset.
	HaDatastoreApdRecoveryAction pulumi.StringPtrInput
	// Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant
	// datastore. Can be one of clusterDefault, disabled, warning, restartConservative, or restartAggressive.
	HaDatastoreApdResponse pulumi.StringPtrInput
	// Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in
	// ha_datastore_apd_response. Specify -1 to use the cluster setting.
	HaDatastoreApdResponseDelay pulumi.IntPtrInput
	// Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant
	// datastore. Can be one of clusterDefault, disabled, warning, or restartAggressive.
	HaDatastorePdlResponse pulumi.StringPtrInput
	// The action to take on this virtual machine when a host is isolated from the rest of the cluster. Can be one of
	// clusterIsolationResponse, none, powerOff, or shutdown.
	HaHostIsolationResponse pulumi.StringPtrInput
	// If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked
	// as failed. The value is in seconds.
	HaVmFailureInterval pulumi.IntPtrInput
	// The length of the reset window in which haVmMaximumResets can operate. When this window expires, no more resets are
	// attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset
	// time is allotted.
	HaVmMaximumFailureWindow pulumi.IntPtrInput
	// The maximum number of resets that HA will perform to this virtual machine when responding to a failure event.
	HaVmMaximumResets pulumi.IntPtrInput
	// The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats.
	HaVmMinimumUptime pulumi.IntPtrInput
	// The type of virtual machine monitoring to use for this virtual machine. Can be one of vmMonitoringDisabled,
	// vmMonitoringOnly, or vmAndAppMonitoring.
	HaVmMonitoring pulumi.StringPtrInput
	// Determines whether or not the cluster's default settings or the VM override settings specified in this resource are used
	// for virtual machine monitoring. The default is true (use cluster defaults) - set to false to have overrides take effect.
	HaVmMonitoringUseClusterDefaults pulumi.BoolPtrInput
	// The restart priority for this virtual machine when vSphere detects a host failure. Can be one of clusterRestartPriority,
	// lowest, low, medium, high, or highest.
	HaVmRestartPriority pulumi.StringPtrInput
	// The maximum time, in seconds, that vSphere HA will wait for the virtual machine to be ready. Use -1 to use the cluster
	// default.
	HaVmRestartTimeout pulumi.IntPtrInput
	// The managed object ID of the virtual machine.
	VirtualMachineId pulumi.StringInput
}

func (HaVmOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*haVmOverrideArgs)(nil)).Elem()
}

type HaVmOverrideInput interface {
	pulumi.Input

	ToHaVmOverrideOutput() HaVmOverrideOutput
	ToHaVmOverrideOutputWithContext(ctx context.Context) HaVmOverrideOutput
}

func (*HaVmOverride) ElementType() reflect.Type {
	return reflect.TypeOf((**HaVmOverride)(nil)).Elem()
}

func (i *HaVmOverride) ToHaVmOverrideOutput() HaVmOverrideOutput {
	return i.ToHaVmOverrideOutputWithContext(context.Background())
}

func (i *HaVmOverride) ToHaVmOverrideOutputWithContext(ctx context.Context) HaVmOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HaVmOverrideOutput)
}

// HaVmOverrideArrayInput is an input type that accepts HaVmOverrideArray and HaVmOverrideArrayOutput values.
// You can construct a concrete instance of `HaVmOverrideArrayInput` via:
//
//	HaVmOverrideArray{ HaVmOverrideArgs{...} }
type HaVmOverrideArrayInput interface {
	pulumi.Input

	ToHaVmOverrideArrayOutput() HaVmOverrideArrayOutput
	ToHaVmOverrideArrayOutputWithContext(context.Context) HaVmOverrideArrayOutput
}

type HaVmOverrideArray []HaVmOverrideInput

func (HaVmOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HaVmOverride)(nil)).Elem()
}

func (i HaVmOverrideArray) ToHaVmOverrideArrayOutput() HaVmOverrideArrayOutput {
	return i.ToHaVmOverrideArrayOutputWithContext(context.Background())
}

func (i HaVmOverrideArray) ToHaVmOverrideArrayOutputWithContext(ctx context.Context) HaVmOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HaVmOverrideArrayOutput)
}

// HaVmOverrideMapInput is an input type that accepts HaVmOverrideMap and HaVmOverrideMapOutput values.
// You can construct a concrete instance of `HaVmOverrideMapInput` via:
//
//	HaVmOverrideMap{ "key": HaVmOverrideArgs{...} }
type HaVmOverrideMapInput interface {
	pulumi.Input

	ToHaVmOverrideMapOutput() HaVmOverrideMapOutput
	ToHaVmOverrideMapOutputWithContext(context.Context) HaVmOverrideMapOutput
}

type HaVmOverrideMap map[string]HaVmOverrideInput

func (HaVmOverrideMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HaVmOverride)(nil)).Elem()
}

func (i HaVmOverrideMap) ToHaVmOverrideMapOutput() HaVmOverrideMapOutput {
	return i.ToHaVmOverrideMapOutputWithContext(context.Background())
}

func (i HaVmOverrideMap) ToHaVmOverrideMapOutputWithContext(ctx context.Context) HaVmOverrideMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HaVmOverrideMapOutput)
}

type HaVmOverrideOutput struct{ *pulumi.OutputState }

func (HaVmOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HaVmOverride)(nil)).Elem()
}

func (o HaVmOverrideOutput) ToHaVmOverrideOutput() HaVmOverrideOutput {
	return o
}

func (o HaVmOverrideOutput) ToHaVmOverrideOutputWithContext(ctx context.Context) HaVmOverrideOutput {
	return o
}

// The managed object ID of the cluster.
func (o HaVmOverrideOutput) ComputeClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *HaVmOverride) pulumi.StringOutput { return v.ComputeClusterId }).(pulumi.StringOutput)
}

// Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an
// APD event. Can be one of useClusterDefault, none or reset.
func (o HaVmOverrideOutput) HaDatastoreApdRecoveryAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HaVmOverride) pulumi.StringPtrOutput { return v.HaDatastoreApdRecoveryAction }).(pulumi.StringPtrOutput)
}

// Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant
// datastore. Can be one of clusterDefault, disabled, warning, restartConservative, or restartAggressive.
func (o HaVmOverrideOutput) HaDatastoreApdResponse() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HaVmOverride) pulumi.StringPtrOutput { return v.HaDatastoreApdResponse }).(pulumi.StringPtrOutput)
}

// Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in
// ha_datastore_apd_response. Specify -1 to use the cluster setting.
func (o HaVmOverrideOutput) HaDatastoreApdResponseDelay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HaVmOverride) pulumi.IntPtrOutput { return v.HaDatastoreApdResponseDelay }).(pulumi.IntPtrOutput)
}

// Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant
// datastore. Can be one of clusterDefault, disabled, warning, or restartAggressive.
func (o HaVmOverrideOutput) HaDatastorePdlResponse() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HaVmOverride) pulumi.StringPtrOutput { return v.HaDatastorePdlResponse }).(pulumi.StringPtrOutput)
}

// The action to take on this virtual machine when a host is isolated from the rest of the cluster. Can be one of
// clusterIsolationResponse, none, powerOff, or shutdown.
func (o HaVmOverrideOutput) HaHostIsolationResponse() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HaVmOverride) pulumi.StringPtrOutput { return v.HaHostIsolationResponse }).(pulumi.StringPtrOutput)
}

// If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked
// as failed. The value is in seconds.
func (o HaVmOverrideOutput) HaVmFailureInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HaVmOverride) pulumi.IntPtrOutput { return v.HaVmFailureInterval }).(pulumi.IntPtrOutput)
}

// The length of the reset window in which haVmMaximumResets can operate. When this window expires, no more resets are
// attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset
// time is allotted.
func (o HaVmOverrideOutput) HaVmMaximumFailureWindow() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HaVmOverride) pulumi.IntPtrOutput { return v.HaVmMaximumFailureWindow }).(pulumi.IntPtrOutput)
}

// The maximum number of resets that HA will perform to this virtual machine when responding to a failure event.
func (o HaVmOverrideOutput) HaVmMaximumResets() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HaVmOverride) pulumi.IntPtrOutput { return v.HaVmMaximumResets }).(pulumi.IntPtrOutput)
}

// The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats.
func (o HaVmOverrideOutput) HaVmMinimumUptime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HaVmOverride) pulumi.IntPtrOutput { return v.HaVmMinimumUptime }).(pulumi.IntPtrOutput)
}

// The type of virtual machine monitoring to use for this virtual machine. Can be one of vmMonitoringDisabled,
// vmMonitoringOnly, or vmAndAppMonitoring.
func (o HaVmOverrideOutput) HaVmMonitoring() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HaVmOverride) pulumi.StringPtrOutput { return v.HaVmMonitoring }).(pulumi.StringPtrOutput)
}

// Determines whether or not the cluster's default settings or the VM override settings specified in this resource are used
// for virtual machine monitoring. The default is true (use cluster defaults) - set to false to have overrides take effect.
func (o HaVmOverrideOutput) HaVmMonitoringUseClusterDefaults() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HaVmOverride) pulumi.BoolPtrOutput { return v.HaVmMonitoringUseClusterDefaults }).(pulumi.BoolPtrOutput)
}

// The restart priority for this virtual machine when vSphere detects a host failure. Can be one of clusterRestartPriority,
// lowest, low, medium, high, or highest.
func (o HaVmOverrideOutput) HaVmRestartPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HaVmOverride) pulumi.StringPtrOutput { return v.HaVmRestartPriority }).(pulumi.StringPtrOutput)
}

// The maximum time, in seconds, that vSphere HA will wait for the virtual machine to be ready. Use -1 to use the cluster
// default.
func (o HaVmOverrideOutput) HaVmRestartTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HaVmOverride) pulumi.IntPtrOutput { return v.HaVmRestartTimeout }).(pulumi.IntPtrOutput)
}

// The managed object ID of the virtual machine.
func (o HaVmOverrideOutput) VirtualMachineId() pulumi.StringOutput {
	return o.ApplyT(func(v *HaVmOverride) pulumi.StringOutput { return v.VirtualMachineId }).(pulumi.StringOutput)
}

type HaVmOverrideArrayOutput struct{ *pulumi.OutputState }

func (HaVmOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HaVmOverride)(nil)).Elem()
}

func (o HaVmOverrideArrayOutput) ToHaVmOverrideArrayOutput() HaVmOverrideArrayOutput {
	return o
}

func (o HaVmOverrideArrayOutput) ToHaVmOverrideArrayOutputWithContext(ctx context.Context) HaVmOverrideArrayOutput {
	return o
}

func (o HaVmOverrideArrayOutput) Index(i pulumi.IntInput) HaVmOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HaVmOverride {
		return vs[0].([]*HaVmOverride)[vs[1].(int)]
	}).(HaVmOverrideOutput)
}

type HaVmOverrideMapOutput struct{ *pulumi.OutputState }

func (HaVmOverrideMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HaVmOverride)(nil)).Elem()
}

func (o HaVmOverrideMapOutput) ToHaVmOverrideMapOutput() HaVmOverrideMapOutput {
	return o
}

func (o HaVmOverrideMapOutput) ToHaVmOverrideMapOutputWithContext(ctx context.Context) HaVmOverrideMapOutput {
	return o
}

func (o HaVmOverrideMapOutput) MapIndex(k pulumi.StringInput) HaVmOverrideOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HaVmOverride {
		return vs[0].(map[string]*HaVmOverride)[vs[1].(string)]
	}).(HaVmOverrideOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HaVmOverrideInput)(nil)).Elem(), &HaVmOverride{})
	pulumi.RegisterInputType(reflect.TypeOf((*HaVmOverrideArrayInput)(nil)).Elem(), HaVmOverrideArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HaVmOverrideMapInput)(nil)).Elem(), HaVmOverrideMap{})
	pulumi.RegisterOutputType(HaVmOverrideOutput{})
	pulumi.RegisterOutputType(HaVmOverrideArrayOutput{})
	pulumi.RegisterOutputType(HaVmOverrideMapOutput{})
}

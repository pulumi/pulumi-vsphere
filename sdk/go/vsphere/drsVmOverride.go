// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `DrsVmOverride` resource can be used to add a DRS override to a
// cluster for a specific virtual machine. With this resource, one can enable or
// disable DRS and control the automation level for a single virtual machine
// without affecting the rest of the cluster.
//
// For more information on vSphere clusters and DRS, see [this
// page][ref-vsphere-drs-clusters].
//
// > **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
//
// ## Example Usage
//
// The example below creates a virtual machine in a cluster using the
// `VirtualMachine` resource, creating the
// virtual machine in the cluster looked up by the
// `ComputeCluster` data source, but also
// pinning the VM to a host defined by the
// `Host` data source, which is assumed to
// be a host within the cluster. To ensure that the VM stays on this host and does
// not need to be migrated back at any point in time, an override is entered using
// the `DrsVmOverride` resource that disables DRS for this virtual
// machine, ensuring that it does not move.
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
//			host, err := vsphere.LookupHost(ctx, &vsphere.LookupHostArgs{
//				Name:         pulumi.StringRef("esxi-01.example.com"),
//				DatacenterId: datacenter.Id,
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
//				HostSystemId:   pulumi.String(host.Id),
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
//			_, err = vsphere.NewDrsVmOverride(ctx, "drs_vm_override", &vsphere.DrsVmOverrideArgs{
//				ComputeClusterId: pulumi.String(cluster.Id),
//				VirtualMachineId: vm.ID(),
//				DrsEnabled:       pulumi.Bool(false),
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
// $ pulumi import vsphere:index/drsVmOverride:DrsVmOverride drs_vm_override \
// ```
//
//	'{"compute_cluster_path": "/dc1/host/cluster1", \
//
//	"virtual_machine_path": "/dc1/vm/srv1"}'
//
// [ref-vsphere-drs-clusters]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-resource-management-8-0/creating-a-drs-cluster.html
type DrsVmOverride struct {
	pulumi.CustomResourceState

	// The managed object reference
	// ID of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringOutput `pulumi:"computeClusterId"`
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	//
	// > **NOTE:** Using this resource _always_ implies an override, even if one of
	// `drsEnabled` or `drsAutomationLevel` is omitted. Take note of the defaults
	// for both options.
	DrsAutomationLevel pulumi.StringPtrOutput `pulumi:"drsAutomationLevel"`
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled pulumi.BoolPtrOutput `pulumi:"drsEnabled"`
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId pulumi.StringOutput `pulumi:"virtualMachineId"`
}

// NewDrsVmOverride registers a new resource with the given unique name, arguments, and options.
func NewDrsVmOverride(ctx *pulumi.Context,
	name string, args *DrsVmOverrideArgs, opts ...pulumi.ResourceOption) (*DrsVmOverride, error) {
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
	var resource DrsVmOverride
	err := ctx.RegisterResource("vsphere:index/drsVmOverride:DrsVmOverride", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDrsVmOverride gets an existing DrsVmOverride resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDrsVmOverride(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DrsVmOverrideState, opts ...pulumi.ResourceOption) (*DrsVmOverride, error) {
	var resource DrsVmOverride
	err := ctx.ReadResource("vsphere:index/drsVmOverride:DrsVmOverride", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DrsVmOverride resources.
type drsVmOverrideState struct {
	// The managed object reference
	// ID of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId *string `pulumi:"computeClusterId"`
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	//
	// > **NOTE:** Using this resource _always_ implies an override, even if one of
	// `drsEnabled` or `drsAutomationLevel` is omitted. Take note of the defaults
	// for both options.
	DrsAutomationLevel *string `pulumi:"drsAutomationLevel"`
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled *bool `pulumi:"drsEnabled"`
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId *string `pulumi:"virtualMachineId"`
}

type DrsVmOverrideState struct {
	// The managed object reference
	// ID of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringPtrInput
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	//
	// > **NOTE:** Using this resource _always_ implies an override, even if one of
	// `drsEnabled` or `drsAutomationLevel` is omitted. Take note of the defaults
	// for both options.
	DrsAutomationLevel pulumi.StringPtrInput
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled pulumi.BoolPtrInput
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId pulumi.StringPtrInput
}

func (DrsVmOverrideState) ElementType() reflect.Type {
	return reflect.TypeOf((*drsVmOverrideState)(nil)).Elem()
}

type drsVmOverrideArgs struct {
	// The managed object reference
	// ID of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId string `pulumi:"computeClusterId"`
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	//
	// > **NOTE:** Using this resource _always_ implies an override, even if one of
	// `drsEnabled` or `drsAutomationLevel` is omitted. Take note of the defaults
	// for both options.
	DrsAutomationLevel *string `pulumi:"drsAutomationLevel"`
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled *bool `pulumi:"drsEnabled"`
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId string `pulumi:"virtualMachineId"`
}

// The set of arguments for constructing a DrsVmOverride resource.
type DrsVmOverrideArgs struct {
	// The managed object reference
	// ID of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringInput
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	//
	// > **NOTE:** Using this resource _always_ implies an override, even if one of
	// `drsEnabled` or `drsAutomationLevel` is omitted. Take note of the defaults
	// for both options.
	DrsAutomationLevel pulumi.StringPtrInput
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled pulumi.BoolPtrInput
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId pulumi.StringInput
}

func (DrsVmOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*drsVmOverrideArgs)(nil)).Elem()
}

type DrsVmOverrideInput interface {
	pulumi.Input

	ToDrsVmOverrideOutput() DrsVmOverrideOutput
	ToDrsVmOverrideOutputWithContext(ctx context.Context) DrsVmOverrideOutput
}

func (*DrsVmOverride) ElementType() reflect.Type {
	return reflect.TypeOf((**DrsVmOverride)(nil)).Elem()
}

func (i *DrsVmOverride) ToDrsVmOverrideOutput() DrsVmOverrideOutput {
	return i.ToDrsVmOverrideOutputWithContext(context.Background())
}

func (i *DrsVmOverride) ToDrsVmOverrideOutputWithContext(ctx context.Context) DrsVmOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrsVmOverrideOutput)
}

// DrsVmOverrideArrayInput is an input type that accepts DrsVmOverrideArray and DrsVmOverrideArrayOutput values.
// You can construct a concrete instance of `DrsVmOverrideArrayInput` via:
//
//	DrsVmOverrideArray{ DrsVmOverrideArgs{...} }
type DrsVmOverrideArrayInput interface {
	pulumi.Input

	ToDrsVmOverrideArrayOutput() DrsVmOverrideArrayOutput
	ToDrsVmOverrideArrayOutputWithContext(context.Context) DrsVmOverrideArrayOutput
}

type DrsVmOverrideArray []DrsVmOverrideInput

func (DrsVmOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DrsVmOverride)(nil)).Elem()
}

func (i DrsVmOverrideArray) ToDrsVmOverrideArrayOutput() DrsVmOverrideArrayOutput {
	return i.ToDrsVmOverrideArrayOutputWithContext(context.Background())
}

func (i DrsVmOverrideArray) ToDrsVmOverrideArrayOutputWithContext(ctx context.Context) DrsVmOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrsVmOverrideArrayOutput)
}

// DrsVmOverrideMapInput is an input type that accepts DrsVmOverrideMap and DrsVmOverrideMapOutput values.
// You can construct a concrete instance of `DrsVmOverrideMapInput` via:
//
//	DrsVmOverrideMap{ "key": DrsVmOverrideArgs{...} }
type DrsVmOverrideMapInput interface {
	pulumi.Input

	ToDrsVmOverrideMapOutput() DrsVmOverrideMapOutput
	ToDrsVmOverrideMapOutputWithContext(context.Context) DrsVmOverrideMapOutput
}

type DrsVmOverrideMap map[string]DrsVmOverrideInput

func (DrsVmOverrideMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DrsVmOverride)(nil)).Elem()
}

func (i DrsVmOverrideMap) ToDrsVmOverrideMapOutput() DrsVmOverrideMapOutput {
	return i.ToDrsVmOverrideMapOutputWithContext(context.Background())
}

func (i DrsVmOverrideMap) ToDrsVmOverrideMapOutputWithContext(ctx context.Context) DrsVmOverrideMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrsVmOverrideMapOutput)
}

type DrsVmOverrideOutput struct{ *pulumi.OutputState }

func (DrsVmOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DrsVmOverride)(nil)).Elem()
}

func (o DrsVmOverrideOutput) ToDrsVmOverrideOutput() DrsVmOverrideOutput {
	return o
}

func (o DrsVmOverrideOutput) ToDrsVmOverrideOutputWithContext(ctx context.Context) DrsVmOverrideOutput {
	return o
}

// The managed object reference
// ID of the cluster to put the override in.  Forces a new
// resource if changed.
func (o DrsVmOverrideOutput) ComputeClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *DrsVmOverride) pulumi.StringOutput { return v.ComputeClusterId }).(pulumi.StringOutput)
}

// Overrides the automation level for this virtual
// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
// `fullyAutomated`. Default: `manual`.
//
// > **NOTE:** Using this resource _always_ implies an override, even if one of
// `drsEnabled` or `drsAutomationLevel` is omitted. Take note of the defaults
// for both options.
func (o DrsVmOverrideOutput) DrsAutomationLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DrsVmOverride) pulumi.StringPtrOutput { return v.DrsAutomationLevel }).(pulumi.StringPtrOutput)
}

// Overrides the default DRS setting for this virtual
// machine. Can be either `true` or `false`. Default: `false`.
func (o DrsVmOverrideOutput) DrsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DrsVmOverride) pulumi.BoolPtrOutput { return v.DrsEnabled }).(pulumi.BoolPtrOutput)
}

// The UUID of the virtual machine to create
// the override for.  Forces a new resource if changed.
func (o DrsVmOverrideOutput) VirtualMachineId() pulumi.StringOutput {
	return o.ApplyT(func(v *DrsVmOverride) pulumi.StringOutput { return v.VirtualMachineId }).(pulumi.StringOutput)
}

type DrsVmOverrideArrayOutput struct{ *pulumi.OutputState }

func (DrsVmOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DrsVmOverride)(nil)).Elem()
}

func (o DrsVmOverrideArrayOutput) ToDrsVmOverrideArrayOutput() DrsVmOverrideArrayOutput {
	return o
}

func (o DrsVmOverrideArrayOutput) ToDrsVmOverrideArrayOutputWithContext(ctx context.Context) DrsVmOverrideArrayOutput {
	return o
}

func (o DrsVmOverrideArrayOutput) Index(i pulumi.IntInput) DrsVmOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DrsVmOverride {
		return vs[0].([]*DrsVmOverride)[vs[1].(int)]
	}).(DrsVmOverrideOutput)
}

type DrsVmOverrideMapOutput struct{ *pulumi.OutputState }

func (DrsVmOverrideMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DrsVmOverride)(nil)).Elem()
}

func (o DrsVmOverrideMapOutput) ToDrsVmOverrideMapOutput() DrsVmOverrideMapOutput {
	return o
}

func (o DrsVmOverrideMapOutput) ToDrsVmOverrideMapOutputWithContext(ctx context.Context) DrsVmOverrideMapOutput {
	return o
}

func (o DrsVmOverrideMapOutput) MapIndex(k pulumi.StringInput) DrsVmOverrideOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DrsVmOverride {
		return vs[0].(map[string]*DrsVmOverride)[vs[1].(string)]
	}).(DrsVmOverrideOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DrsVmOverrideInput)(nil)).Elem(), &DrsVmOverride{})
	pulumi.RegisterInputType(reflect.TypeOf((*DrsVmOverrideArrayInput)(nil)).Elem(), DrsVmOverrideArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DrsVmOverrideMapInput)(nil)).Elem(), DrsVmOverrideMap{})
	pulumi.RegisterOutputType(DrsVmOverrideOutput{})
	pulumi.RegisterOutputType(DrsVmOverrideArrayOutput{})
	pulumi.RegisterOutputType(DrsVmOverrideMapOutput{})
}

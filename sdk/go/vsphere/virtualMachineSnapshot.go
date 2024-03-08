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

// The `VirtualMachineSnapshot` resource can be used to manage snapshots
// for a virtual machine.
//
// For more information on managing snapshots and how they work in VMware, see
// [here][ext-vm-snapshot-management].
//
// > **NOTE:** A snapshot in VMware differs from traditional disk snapshots, and
// can contain the actual running state of the virtual machine, data for all disks
// that have not been set to be independent from the snapshot (including ones that
// have been attached via the `attach`
// parameter to the `VirtualMachine` `disk` block), and even the
// configuration of the virtual machine at the time of the snapshot. Virtual
// machine, disk activity, and configuration changes post-snapshot are not
// included in the original state. Use this resource with care! Neither VMware nor
// HashiCorp recommends retaining snapshots for a extended period of time and does
// NOT recommend using them as as backup feature. For more information on the
// limitation of virtual machine snapshots, see [here][ext-vm-snap-limitations].
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := vsphere.NewVirtualMachineSnapshot(ctx, "demo1", &vsphere.VirtualMachineSnapshotArgs{
//				Consolidate:        pulumi.Bool(true),
//				Description:        pulumi.String("This is Demo Snapshot"),
//				Memory:             pulumi.Bool(true),
//				Quiesce:            pulumi.Bool(true),
//				RemoveChildren:     pulumi.Bool(false),
//				SnapshotName:       pulumi.String("Snapshot Name"),
//				VirtualMachineUuid: pulumi.String("9aac5551-a351-4158-8c5c-15a71e8ec5c9"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// [ext-vm-snapshot-management]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-CA948C69-7F58-4519-AEB1-739545EA94E5.html
// [ext-vm-snap-limitations]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-53F65726-A23B-4CF0-A7D5-48E584B88613.html
type VirtualMachineSnapshot struct {
	pulumi.CustomResourceState

	// If set to `true`, the delta disks involved in this
	// snapshot will be consolidated into the parent when this resource is
	// destroyed.
	Consolidate pulumi.BoolPtrOutput `pulumi:"consolidate"`
	// A description for the snapshot.
	Description pulumi.StringOutput `pulumi:"description"`
	// If set to `true`, a dump of the internal state of the
	// virtual machine is included in the snapshot.
	Memory pulumi.BoolOutput `pulumi:"memory"`
	// If set to `true`, and the virtual machine is powered
	// on when the snapshot is taken, VMware Tools is used to quiesce the file
	// system in the virtual machine.
	Quiesce pulumi.BoolOutput `pulumi:"quiesce"`
	// If set to `true`, the entire snapshot subtree
	// is removed when this resource is destroyed.
	RemoveChildren pulumi.BoolPtrOutput `pulumi:"removeChildren"`
	// The name of the snapshot.
	SnapshotName pulumi.StringOutput `pulumi:"snapshotName"`
	// The virtual machine UUID.
	VirtualMachineUuid pulumi.StringOutput `pulumi:"virtualMachineUuid"`
}

// NewVirtualMachineSnapshot registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineSnapshot(ctx *pulumi.Context,
	name string, args *VirtualMachineSnapshotArgs, opts ...pulumi.ResourceOption) (*VirtualMachineSnapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Memory == nil {
		return nil, errors.New("invalid value for required argument 'Memory'")
	}
	if args.Quiesce == nil {
		return nil, errors.New("invalid value for required argument 'Quiesce'")
	}
	if args.SnapshotName == nil {
		return nil, errors.New("invalid value for required argument 'SnapshotName'")
	}
	if args.VirtualMachineUuid == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineUuid'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualMachineSnapshot
	err := ctx.RegisterResource("vsphere:index/virtualMachineSnapshot:VirtualMachineSnapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineSnapshot gets an existing VirtualMachineSnapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineSnapshotState, opts ...pulumi.ResourceOption) (*VirtualMachineSnapshot, error) {
	var resource VirtualMachineSnapshot
	err := ctx.ReadResource("vsphere:index/virtualMachineSnapshot:VirtualMachineSnapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineSnapshot resources.
type virtualMachineSnapshotState struct {
	// If set to `true`, the delta disks involved in this
	// snapshot will be consolidated into the parent when this resource is
	// destroyed.
	Consolidate *bool `pulumi:"consolidate"`
	// A description for the snapshot.
	Description *string `pulumi:"description"`
	// If set to `true`, a dump of the internal state of the
	// virtual machine is included in the snapshot.
	Memory *bool `pulumi:"memory"`
	// If set to `true`, and the virtual machine is powered
	// on when the snapshot is taken, VMware Tools is used to quiesce the file
	// system in the virtual machine.
	Quiesce *bool `pulumi:"quiesce"`
	// If set to `true`, the entire snapshot subtree
	// is removed when this resource is destroyed.
	RemoveChildren *bool `pulumi:"removeChildren"`
	// The name of the snapshot.
	SnapshotName *string `pulumi:"snapshotName"`
	// The virtual machine UUID.
	VirtualMachineUuid *string `pulumi:"virtualMachineUuid"`
}

type VirtualMachineSnapshotState struct {
	// If set to `true`, the delta disks involved in this
	// snapshot will be consolidated into the parent when this resource is
	// destroyed.
	Consolidate pulumi.BoolPtrInput
	// A description for the snapshot.
	Description pulumi.StringPtrInput
	// If set to `true`, a dump of the internal state of the
	// virtual machine is included in the snapshot.
	Memory pulumi.BoolPtrInput
	// If set to `true`, and the virtual machine is powered
	// on when the snapshot is taken, VMware Tools is used to quiesce the file
	// system in the virtual machine.
	Quiesce pulumi.BoolPtrInput
	// If set to `true`, the entire snapshot subtree
	// is removed when this resource is destroyed.
	RemoveChildren pulumi.BoolPtrInput
	// The name of the snapshot.
	SnapshotName pulumi.StringPtrInput
	// The virtual machine UUID.
	VirtualMachineUuid pulumi.StringPtrInput
}

func (VirtualMachineSnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineSnapshotState)(nil)).Elem()
}

type virtualMachineSnapshotArgs struct {
	// If set to `true`, the delta disks involved in this
	// snapshot will be consolidated into the parent when this resource is
	// destroyed.
	Consolidate *bool `pulumi:"consolidate"`
	// A description for the snapshot.
	Description string `pulumi:"description"`
	// If set to `true`, a dump of the internal state of the
	// virtual machine is included in the snapshot.
	Memory bool `pulumi:"memory"`
	// If set to `true`, and the virtual machine is powered
	// on when the snapshot is taken, VMware Tools is used to quiesce the file
	// system in the virtual machine.
	Quiesce bool `pulumi:"quiesce"`
	// If set to `true`, the entire snapshot subtree
	// is removed when this resource is destroyed.
	RemoveChildren *bool `pulumi:"removeChildren"`
	// The name of the snapshot.
	SnapshotName string `pulumi:"snapshotName"`
	// The virtual machine UUID.
	VirtualMachineUuid string `pulumi:"virtualMachineUuid"`
}

// The set of arguments for constructing a VirtualMachineSnapshot resource.
type VirtualMachineSnapshotArgs struct {
	// If set to `true`, the delta disks involved in this
	// snapshot will be consolidated into the parent when this resource is
	// destroyed.
	Consolidate pulumi.BoolPtrInput
	// A description for the snapshot.
	Description pulumi.StringInput
	// If set to `true`, a dump of the internal state of the
	// virtual machine is included in the snapshot.
	Memory pulumi.BoolInput
	// If set to `true`, and the virtual machine is powered
	// on when the snapshot is taken, VMware Tools is used to quiesce the file
	// system in the virtual machine.
	Quiesce pulumi.BoolInput
	// If set to `true`, the entire snapshot subtree
	// is removed when this resource is destroyed.
	RemoveChildren pulumi.BoolPtrInput
	// The name of the snapshot.
	SnapshotName pulumi.StringInput
	// The virtual machine UUID.
	VirtualMachineUuid pulumi.StringInput
}

func (VirtualMachineSnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineSnapshotArgs)(nil)).Elem()
}

type VirtualMachineSnapshotInput interface {
	pulumi.Input

	ToVirtualMachineSnapshotOutput() VirtualMachineSnapshotOutput
	ToVirtualMachineSnapshotOutputWithContext(ctx context.Context) VirtualMachineSnapshotOutput
}

func (*VirtualMachineSnapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineSnapshot)(nil)).Elem()
}

func (i *VirtualMachineSnapshot) ToVirtualMachineSnapshotOutput() VirtualMachineSnapshotOutput {
	return i.ToVirtualMachineSnapshotOutputWithContext(context.Background())
}

func (i *VirtualMachineSnapshot) ToVirtualMachineSnapshotOutputWithContext(ctx context.Context) VirtualMachineSnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineSnapshotOutput)
}

// VirtualMachineSnapshotArrayInput is an input type that accepts VirtualMachineSnapshotArray and VirtualMachineSnapshotArrayOutput values.
// You can construct a concrete instance of `VirtualMachineSnapshotArrayInput` via:
//
//	VirtualMachineSnapshotArray{ VirtualMachineSnapshotArgs{...} }
type VirtualMachineSnapshotArrayInput interface {
	pulumi.Input

	ToVirtualMachineSnapshotArrayOutput() VirtualMachineSnapshotArrayOutput
	ToVirtualMachineSnapshotArrayOutputWithContext(context.Context) VirtualMachineSnapshotArrayOutput
}

type VirtualMachineSnapshotArray []VirtualMachineSnapshotInput

func (VirtualMachineSnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualMachineSnapshot)(nil)).Elem()
}

func (i VirtualMachineSnapshotArray) ToVirtualMachineSnapshotArrayOutput() VirtualMachineSnapshotArrayOutput {
	return i.ToVirtualMachineSnapshotArrayOutputWithContext(context.Background())
}

func (i VirtualMachineSnapshotArray) ToVirtualMachineSnapshotArrayOutputWithContext(ctx context.Context) VirtualMachineSnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineSnapshotArrayOutput)
}

// VirtualMachineSnapshotMapInput is an input type that accepts VirtualMachineSnapshotMap and VirtualMachineSnapshotMapOutput values.
// You can construct a concrete instance of `VirtualMachineSnapshotMapInput` via:
//
//	VirtualMachineSnapshotMap{ "key": VirtualMachineSnapshotArgs{...} }
type VirtualMachineSnapshotMapInput interface {
	pulumi.Input

	ToVirtualMachineSnapshotMapOutput() VirtualMachineSnapshotMapOutput
	ToVirtualMachineSnapshotMapOutputWithContext(context.Context) VirtualMachineSnapshotMapOutput
}

type VirtualMachineSnapshotMap map[string]VirtualMachineSnapshotInput

func (VirtualMachineSnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualMachineSnapshot)(nil)).Elem()
}

func (i VirtualMachineSnapshotMap) ToVirtualMachineSnapshotMapOutput() VirtualMachineSnapshotMapOutput {
	return i.ToVirtualMachineSnapshotMapOutputWithContext(context.Background())
}

func (i VirtualMachineSnapshotMap) ToVirtualMachineSnapshotMapOutputWithContext(ctx context.Context) VirtualMachineSnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineSnapshotMapOutput)
}

type VirtualMachineSnapshotOutput struct{ *pulumi.OutputState }

func (VirtualMachineSnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineSnapshot)(nil)).Elem()
}

func (o VirtualMachineSnapshotOutput) ToVirtualMachineSnapshotOutput() VirtualMachineSnapshotOutput {
	return o
}

func (o VirtualMachineSnapshotOutput) ToVirtualMachineSnapshotOutputWithContext(ctx context.Context) VirtualMachineSnapshotOutput {
	return o
}

// If set to `true`, the delta disks involved in this
// snapshot will be consolidated into the parent when this resource is
// destroyed.
func (o VirtualMachineSnapshotOutput) Consolidate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSnapshot) pulumi.BoolPtrOutput { return v.Consolidate }).(pulumi.BoolPtrOutput)
}

// A description for the snapshot.
func (o VirtualMachineSnapshotOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineSnapshot) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// If set to `true`, a dump of the internal state of the
// virtual machine is included in the snapshot.
func (o VirtualMachineSnapshotOutput) Memory() pulumi.BoolOutput {
	return o.ApplyT(func(v *VirtualMachineSnapshot) pulumi.BoolOutput { return v.Memory }).(pulumi.BoolOutput)
}

// If set to `true`, and the virtual machine is powered
// on when the snapshot is taken, VMware Tools is used to quiesce the file
// system in the virtual machine.
func (o VirtualMachineSnapshotOutput) Quiesce() pulumi.BoolOutput {
	return o.ApplyT(func(v *VirtualMachineSnapshot) pulumi.BoolOutput { return v.Quiesce }).(pulumi.BoolOutput)
}

// If set to `true`, the entire snapshot subtree
// is removed when this resource is destroyed.
func (o VirtualMachineSnapshotOutput) RemoveChildren() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSnapshot) pulumi.BoolPtrOutput { return v.RemoveChildren }).(pulumi.BoolPtrOutput)
}

// The name of the snapshot.
func (o VirtualMachineSnapshotOutput) SnapshotName() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineSnapshot) pulumi.StringOutput { return v.SnapshotName }).(pulumi.StringOutput)
}

// The virtual machine UUID.
func (o VirtualMachineSnapshotOutput) VirtualMachineUuid() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineSnapshot) pulumi.StringOutput { return v.VirtualMachineUuid }).(pulumi.StringOutput)
}

type VirtualMachineSnapshotArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineSnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualMachineSnapshot)(nil)).Elem()
}

func (o VirtualMachineSnapshotArrayOutput) ToVirtualMachineSnapshotArrayOutput() VirtualMachineSnapshotArrayOutput {
	return o
}

func (o VirtualMachineSnapshotArrayOutput) ToVirtualMachineSnapshotArrayOutputWithContext(ctx context.Context) VirtualMachineSnapshotArrayOutput {
	return o
}

func (o VirtualMachineSnapshotArrayOutput) Index(i pulumi.IntInput) VirtualMachineSnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualMachineSnapshot {
		return vs[0].([]*VirtualMachineSnapshot)[vs[1].(int)]
	}).(VirtualMachineSnapshotOutput)
}

type VirtualMachineSnapshotMapOutput struct{ *pulumi.OutputState }

func (VirtualMachineSnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualMachineSnapshot)(nil)).Elem()
}

func (o VirtualMachineSnapshotMapOutput) ToVirtualMachineSnapshotMapOutput() VirtualMachineSnapshotMapOutput {
	return o
}

func (o VirtualMachineSnapshotMapOutput) ToVirtualMachineSnapshotMapOutputWithContext(ctx context.Context) VirtualMachineSnapshotMapOutput {
	return o
}

func (o VirtualMachineSnapshotMapOutput) MapIndex(k pulumi.StringInput) VirtualMachineSnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualMachineSnapshot {
		return vs[0].(map[string]*VirtualMachineSnapshot)[vs[1].(string)]
	}).(VirtualMachineSnapshotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineSnapshotInput)(nil)).Elem(), &VirtualMachineSnapshot{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineSnapshotArrayInput)(nil)).Elem(), VirtualMachineSnapshotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineSnapshotMapInput)(nil)).Elem(), VirtualMachineSnapshotMap{})
	pulumi.RegisterOutputType(VirtualMachineSnapshotOutput{})
	pulumi.RegisterOutputType(VirtualMachineSnapshotArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineSnapshotMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `VirtualDisk` resource can be used to create virtual disks outside
// of any given `VirtualMachine`
// resource. These disks can be attached to a virtual machine by creating a disk
// block with the `attach` parameter.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vsphere.NewVirtualDisk(ctx, "myDisk", &vsphere.VirtualDiskArgs{
// 			Datacenter: pulumi.String("Datacenter"),
// 			Datastore:  pulumi.String("local"),
// 			Size:       pulumi.Int(2),
// 			Type:       pulumi.String("thin"),
// 			VmdkPath:   pulumi.String("myDisk.vmdk"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type VirtualDisk struct {
	pulumi.CustomResourceState

	// The adapter type for this virtual disk. Can be
	// one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
	//
	// Deprecated: this attribute has no effect on controller types - please use scsi_type in vsphere_virtual_machine instead
	AdapterType pulumi.StringPtrOutput `pulumi:"adapterType"`
	// Tells the resource to create any
	// directories that are a part of the `vmdkPath` parameter if they are missing.
	// Default: `false`.
	CreateDirectories pulumi.BoolPtrOutput `pulumi:"createDirectories"`
	// The name of the datacenter in which to create the
	// disk. Can be omitted when when ESXi or if there is only one datacenter in
	// your infrastructure.
	Datacenter pulumi.StringPtrOutput `pulumi:"datacenter"`
	// The name of the datastore in which to create the
	// disk.
	Datastore pulumi.StringOutput `pulumi:"datastore"`
	// Size of the disk (in GB).
	Size pulumi.IntOutput `pulumi:"size"`
	// The type of disk to create. Can be one of
	// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
	// information on what each kind of disk provisioning policy means, click
	// [here][docs-vmware-vm-disk-provisioning].
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The path, including filename, of the virtual disk to
	// be created.  This needs to end in `.vmdk`.
	VmdkPath pulumi.StringOutput `pulumi:"vmdkPath"`
}

// NewVirtualDisk registers a new resource with the given unique name, arguments, and options.
func NewVirtualDisk(ctx *pulumi.Context,
	name string, args *VirtualDiskArgs, opts ...pulumi.ResourceOption) (*VirtualDisk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Datastore == nil {
		return nil, errors.New("invalid value for required argument 'Datastore'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	if args.VmdkPath == nil {
		return nil, errors.New("invalid value for required argument 'VmdkPath'")
	}
	var resource VirtualDisk
	err := ctx.RegisterResource("vsphere:index/virtualDisk:VirtualDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualDisk gets an existing VirtualDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualDiskState, opts ...pulumi.ResourceOption) (*VirtualDisk, error) {
	var resource VirtualDisk
	err := ctx.ReadResource("vsphere:index/virtualDisk:VirtualDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualDisk resources.
type virtualDiskState struct {
	// The adapter type for this virtual disk. Can be
	// one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
	//
	// Deprecated: this attribute has no effect on controller types - please use scsi_type in vsphere_virtual_machine instead
	AdapterType *string `pulumi:"adapterType"`
	// Tells the resource to create any
	// directories that are a part of the `vmdkPath` parameter if they are missing.
	// Default: `false`.
	CreateDirectories *bool `pulumi:"createDirectories"`
	// The name of the datacenter in which to create the
	// disk. Can be omitted when when ESXi or if there is only one datacenter in
	// your infrastructure.
	Datacenter *string `pulumi:"datacenter"`
	// The name of the datastore in which to create the
	// disk.
	Datastore *string `pulumi:"datastore"`
	// Size of the disk (in GB).
	Size *int `pulumi:"size"`
	// The type of disk to create. Can be one of
	// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
	// information on what each kind of disk provisioning policy means, click
	// [here][docs-vmware-vm-disk-provisioning].
	Type *string `pulumi:"type"`
	// The path, including filename, of the virtual disk to
	// be created.  This needs to end in `.vmdk`.
	VmdkPath *string `pulumi:"vmdkPath"`
}

type VirtualDiskState struct {
	// The adapter type for this virtual disk. Can be
	// one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
	//
	// Deprecated: this attribute has no effect on controller types - please use scsi_type in vsphere_virtual_machine instead
	AdapterType pulumi.StringPtrInput
	// Tells the resource to create any
	// directories that are a part of the `vmdkPath` parameter if they are missing.
	// Default: `false`.
	CreateDirectories pulumi.BoolPtrInput
	// The name of the datacenter in which to create the
	// disk. Can be omitted when when ESXi or if there is only one datacenter in
	// your infrastructure.
	Datacenter pulumi.StringPtrInput
	// The name of the datastore in which to create the
	// disk.
	Datastore pulumi.StringPtrInput
	// Size of the disk (in GB).
	Size pulumi.IntPtrInput
	// The type of disk to create. Can be one of
	// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
	// information on what each kind of disk provisioning policy means, click
	// [here][docs-vmware-vm-disk-provisioning].
	Type pulumi.StringPtrInput
	// The path, including filename, of the virtual disk to
	// be created.  This needs to end in `.vmdk`.
	VmdkPath pulumi.StringPtrInput
}

func (VirtualDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualDiskState)(nil)).Elem()
}

type virtualDiskArgs struct {
	// The adapter type for this virtual disk. Can be
	// one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
	//
	// Deprecated: this attribute has no effect on controller types - please use scsi_type in vsphere_virtual_machine instead
	AdapterType *string `pulumi:"adapterType"`
	// Tells the resource to create any
	// directories that are a part of the `vmdkPath` parameter if they are missing.
	// Default: `false`.
	CreateDirectories *bool `pulumi:"createDirectories"`
	// The name of the datacenter in which to create the
	// disk. Can be omitted when when ESXi or if there is only one datacenter in
	// your infrastructure.
	Datacenter *string `pulumi:"datacenter"`
	// The name of the datastore in which to create the
	// disk.
	Datastore string `pulumi:"datastore"`
	// Size of the disk (in GB).
	Size int `pulumi:"size"`
	// The type of disk to create. Can be one of
	// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
	// information on what each kind of disk provisioning policy means, click
	// [here][docs-vmware-vm-disk-provisioning].
	Type *string `pulumi:"type"`
	// The path, including filename, of the virtual disk to
	// be created.  This needs to end in `.vmdk`.
	VmdkPath string `pulumi:"vmdkPath"`
}

// The set of arguments for constructing a VirtualDisk resource.
type VirtualDiskArgs struct {
	// The adapter type for this virtual disk. Can be
	// one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
	//
	// Deprecated: this attribute has no effect on controller types - please use scsi_type in vsphere_virtual_machine instead
	AdapterType pulumi.StringPtrInput
	// Tells the resource to create any
	// directories that are a part of the `vmdkPath` parameter if they are missing.
	// Default: `false`.
	CreateDirectories pulumi.BoolPtrInput
	// The name of the datacenter in which to create the
	// disk. Can be omitted when when ESXi or if there is only one datacenter in
	// your infrastructure.
	Datacenter pulumi.StringPtrInput
	// The name of the datastore in which to create the
	// disk.
	Datastore pulumi.StringInput
	// Size of the disk (in GB).
	Size pulumi.IntInput
	// The type of disk to create. Can be one of
	// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
	// information on what each kind of disk provisioning policy means, click
	// [here][docs-vmware-vm-disk-provisioning].
	Type pulumi.StringPtrInput
	// The path, including filename, of the virtual disk to
	// be created.  This needs to end in `.vmdk`.
	VmdkPath pulumi.StringInput
}

func (VirtualDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualDiskArgs)(nil)).Elem()
}

type VirtualDiskInput interface {
	pulumi.Input

	ToVirtualDiskOutput() VirtualDiskOutput
	ToVirtualDiskOutputWithContext(ctx context.Context) VirtualDiskOutput
}

func (*VirtualDisk) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDisk)(nil))
}

func (i *VirtualDisk) ToVirtualDiskOutput() VirtualDiskOutput {
	return i.ToVirtualDiskOutputWithContext(context.Background())
}

func (i *VirtualDisk) ToVirtualDiskOutputWithContext(ctx context.Context) VirtualDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskOutput)
}

func (i *VirtualDisk) ToVirtualDiskPtrOutput() VirtualDiskPtrOutput {
	return i.ToVirtualDiskPtrOutputWithContext(context.Background())
}

func (i *VirtualDisk) ToVirtualDiskPtrOutputWithContext(ctx context.Context) VirtualDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskPtrOutput)
}

type VirtualDiskPtrInput interface {
	pulumi.Input

	ToVirtualDiskPtrOutput() VirtualDiskPtrOutput
	ToVirtualDiskPtrOutputWithContext(ctx context.Context) VirtualDiskPtrOutput
}

type virtualDiskPtrType VirtualDiskArgs

func (*virtualDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualDisk)(nil))
}

func (i *virtualDiskPtrType) ToVirtualDiskPtrOutput() VirtualDiskPtrOutput {
	return i.ToVirtualDiskPtrOutputWithContext(context.Background())
}

func (i *virtualDiskPtrType) ToVirtualDiskPtrOutputWithContext(ctx context.Context) VirtualDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskPtrOutput)
}

// VirtualDiskArrayInput is an input type that accepts VirtualDiskArray and VirtualDiskArrayOutput values.
// You can construct a concrete instance of `VirtualDiskArrayInput` via:
//
//          VirtualDiskArray{ VirtualDiskArgs{...} }
type VirtualDiskArrayInput interface {
	pulumi.Input

	ToVirtualDiskArrayOutput() VirtualDiskArrayOutput
	ToVirtualDiskArrayOutputWithContext(context.Context) VirtualDiskArrayOutput
}

type VirtualDiskArray []VirtualDiskInput

func (VirtualDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualDisk)(nil)).Elem()
}

func (i VirtualDiskArray) ToVirtualDiskArrayOutput() VirtualDiskArrayOutput {
	return i.ToVirtualDiskArrayOutputWithContext(context.Background())
}

func (i VirtualDiskArray) ToVirtualDiskArrayOutputWithContext(ctx context.Context) VirtualDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskArrayOutput)
}

// VirtualDiskMapInput is an input type that accepts VirtualDiskMap and VirtualDiskMapOutput values.
// You can construct a concrete instance of `VirtualDiskMapInput` via:
//
//          VirtualDiskMap{ "key": VirtualDiskArgs{...} }
type VirtualDiskMapInput interface {
	pulumi.Input

	ToVirtualDiskMapOutput() VirtualDiskMapOutput
	ToVirtualDiskMapOutputWithContext(context.Context) VirtualDiskMapOutput
}

type VirtualDiskMap map[string]VirtualDiskInput

func (VirtualDiskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualDisk)(nil)).Elem()
}

func (i VirtualDiskMap) ToVirtualDiskMapOutput() VirtualDiskMapOutput {
	return i.ToVirtualDiskMapOutputWithContext(context.Background())
}

func (i VirtualDiskMap) ToVirtualDiskMapOutputWithContext(ctx context.Context) VirtualDiskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskMapOutput)
}

type VirtualDiskOutput struct{ *pulumi.OutputState }

func (VirtualDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDisk)(nil))
}

func (o VirtualDiskOutput) ToVirtualDiskOutput() VirtualDiskOutput {
	return o
}

func (o VirtualDiskOutput) ToVirtualDiskOutputWithContext(ctx context.Context) VirtualDiskOutput {
	return o
}

func (o VirtualDiskOutput) ToVirtualDiskPtrOutput() VirtualDiskPtrOutput {
	return o.ToVirtualDiskPtrOutputWithContext(context.Background())
}

func (o VirtualDiskOutput) ToVirtualDiskPtrOutputWithContext(ctx context.Context) VirtualDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualDisk) *VirtualDisk {
		return &v
	}).(VirtualDiskPtrOutput)
}

type VirtualDiskPtrOutput struct{ *pulumi.OutputState }

func (VirtualDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualDisk)(nil))
}

func (o VirtualDiskPtrOutput) ToVirtualDiskPtrOutput() VirtualDiskPtrOutput {
	return o
}

func (o VirtualDiskPtrOutput) ToVirtualDiskPtrOutputWithContext(ctx context.Context) VirtualDiskPtrOutput {
	return o
}

func (o VirtualDiskPtrOutput) Elem() VirtualDiskOutput {
	return o.ApplyT(func(v *VirtualDisk) VirtualDisk {
		if v != nil {
			return *v
		}
		var ret VirtualDisk
		return ret
	}).(VirtualDiskOutput)
}

type VirtualDiskArrayOutput struct{ *pulumi.OutputState }

func (VirtualDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDisk)(nil))
}

func (o VirtualDiskArrayOutput) ToVirtualDiskArrayOutput() VirtualDiskArrayOutput {
	return o
}

func (o VirtualDiskArrayOutput) ToVirtualDiskArrayOutputWithContext(ctx context.Context) VirtualDiskArrayOutput {
	return o
}

func (o VirtualDiskArrayOutput) Index(i pulumi.IntInput) VirtualDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualDisk {
		return vs[0].([]VirtualDisk)[vs[1].(int)]
	}).(VirtualDiskOutput)
}

type VirtualDiskMapOutput struct{ *pulumi.OutputState }

func (VirtualDiskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VirtualDisk)(nil))
}

func (o VirtualDiskMapOutput) ToVirtualDiskMapOutput() VirtualDiskMapOutput {
	return o
}

func (o VirtualDiskMapOutput) ToVirtualDiskMapOutputWithContext(ctx context.Context) VirtualDiskMapOutput {
	return o
}

func (o VirtualDiskMapOutput) MapIndex(k pulumi.StringInput) VirtualDiskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VirtualDisk {
		return vs[0].(map[string]VirtualDisk)[vs[1].(string)]
	}).(VirtualDiskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDiskInput)(nil)).Elem(), &VirtualDisk{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDiskPtrInput)(nil)).Elem(), &VirtualDisk{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDiskArrayInput)(nil)).Elem(), VirtualDiskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDiskMapInput)(nil)).Elem(), VirtualDiskMap{})
	pulumi.RegisterOutputType(VirtualDiskOutput{})
	pulumi.RegisterOutputType(VirtualDiskPtrOutput{})
	pulumi.RegisterOutputType(VirtualDiskArrayOutput{})
	pulumi.RegisterOutputType(VirtualDiskMapOutput{})
}

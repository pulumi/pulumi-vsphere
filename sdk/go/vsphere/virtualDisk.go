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
//			_, err = vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
//				Name: pulumi.StringRef("datastore-01"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.NewVirtualDisk(ctx, "virtual_disk", &vsphere.VirtualDiskArgs{
//				Size:              pulumi.Int(40),
//				Type:              pulumi.String("thin"),
//				VmdkPath:          pulumi.String("/foo/foo.vmdk"),
//				CreateDirectories: pulumi.Bool(true),
//				Datacenter:        pulumi.String(datacenter.Name),
//				Datastore:         pulumi.Any(datastoreVsphereDatastore.Name),
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
// # An existing virtual disk can be imported into this resource
//
// via supplying the full datastore path to the virtual disk. An example is below:
//
// ```sh
// $ pulumi import vsphere:index/virtualDisk:VirtualDisk virtual_disk \
// ```
//
//	'{"virtual_disk_path": "/dc-01/[datastore-01]foo/bar.vmdk", \ "create_directories": "true"}'
//
// The above would import the virtual disk located at `foo/bar.vmdk` in the `datastore-01`
//
// datastore of the `dc-01` datacenter with `create_directories` set as `true`.
//
// [docs-import]: https://developer.hashicorp.com/terraform/cli/import
type VirtualDisk struct {
	pulumi.CustomResourceState

	// The adapter type for this virtual disk. Can be
	// one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
	//
	// > **NOTE:** `adapterType` is **deprecated**: it does not dictate the type of
	// controller that the virtual disk will be attached to on the virtual machine.
	// Please see the `scsiType` parameter
	// in the `VirtualMachine` resource for information on how to control
	// disk controller types. This parameter will be removed in future versions of the
	// vSphere provider.
	//
	// Deprecated: this attribute has no effect on controller types - please use scsiType in VirtualMachine instead
	AdapterType pulumi.StringPtrOutput `pulumi:"adapterType"`
	// Tells the resource to create any
	// directories that are a part of the `vmdkPath` parameter if they are missing.
	// Default: `false`.
	//
	// > **NOTE:** Any directory created as part of the operation when
	// `createDirectories` is enabled will not be deleted when the resource is
	// destroyed.
	CreateDirectories pulumi.BoolPtrOutput `pulumi:"createDirectories"`
	// The name of the datacenter in which to create the
	// disk. Can be omitted when ESXi or if there is only one datacenter in
	// your infrastructure.
	Datacenter pulumi.StringPtrOutput `pulumi:"datacenter"`
	// The name of the datastore in which to create the
	// disk.
	Datastore pulumi.StringOutput `pulumi:"datastore"`
	// Size of the disk (in GB). Decreasing the size of a disk is not possible.
	// If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
	// created.
	Size pulumi.IntOutput `pulumi:"size"`
	// The type of disk to create. Can be one of
	// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
	// information on what each kind of disk provisioning policy means, click
	// [here][docs-vmware-vm-disk-provisioning].
	//
	// [docs-vmware-vm-disk-provisioning]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-single-host-management-vmware-host-client-8-0/virtual-machine-management-with-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/configuring-virtual-machines-in-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/virtual-disk-configuration-vSphereSingleHostManagementVMwareHostClient/about-virtual-disk-provisioning-policies-vSphereSingleHostManagementVMwareHostClient.html
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
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// > **NOTE:** `adapterType` is **deprecated**: it does not dictate the type of
	// controller that the virtual disk will be attached to on the virtual machine.
	// Please see the `scsiType` parameter
	// in the `VirtualMachine` resource for information on how to control
	// disk controller types. This parameter will be removed in future versions of the
	// vSphere provider.
	//
	// Deprecated: this attribute has no effect on controller types - please use scsiType in VirtualMachine instead
	AdapterType *string `pulumi:"adapterType"`
	// Tells the resource to create any
	// directories that are a part of the `vmdkPath` parameter if they are missing.
	// Default: `false`.
	//
	// > **NOTE:** Any directory created as part of the operation when
	// `createDirectories` is enabled will not be deleted when the resource is
	// destroyed.
	CreateDirectories *bool `pulumi:"createDirectories"`
	// The name of the datacenter in which to create the
	// disk. Can be omitted when ESXi or if there is only one datacenter in
	// your infrastructure.
	Datacenter *string `pulumi:"datacenter"`
	// The name of the datastore in which to create the
	// disk.
	Datastore *string `pulumi:"datastore"`
	// Size of the disk (in GB). Decreasing the size of a disk is not possible.
	// If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
	// created.
	Size *int `pulumi:"size"`
	// The type of disk to create. Can be one of
	// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
	// information on what each kind of disk provisioning policy means, click
	// [here][docs-vmware-vm-disk-provisioning].
	//
	// [docs-vmware-vm-disk-provisioning]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-single-host-management-vmware-host-client-8-0/virtual-machine-management-with-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/configuring-virtual-machines-in-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/virtual-disk-configuration-vSphereSingleHostManagementVMwareHostClient/about-virtual-disk-provisioning-policies-vSphereSingleHostManagementVMwareHostClient.html
	Type *string `pulumi:"type"`
	// The path, including filename, of the virtual disk to
	// be created.  This needs to end in `.vmdk`.
	VmdkPath *string `pulumi:"vmdkPath"`
}

type VirtualDiskState struct {
	// The adapter type for this virtual disk. Can be
	// one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
	//
	// > **NOTE:** `adapterType` is **deprecated**: it does not dictate the type of
	// controller that the virtual disk will be attached to on the virtual machine.
	// Please see the `scsiType` parameter
	// in the `VirtualMachine` resource for information on how to control
	// disk controller types. This parameter will be removed in future versions of the
	// vSphere provider.
	//
	// Deprecated: this attribute has no effect on controller types - please use scsiType in VirtualMachine instead
	AdapterType pulumi.StringPtrInput
	// Tells the resource to create any
	// directories that are a part of the `vmdkPath` parameter if they are missing.
	// Default: `false`.
	//
	// > **NOTE:** Any directory created as part of the operation when
	// `createDirectories` is enabled will not be deleted when the resource is
	// destroyed.
	CreateDirectories pulumi.BoolPtrInput
	// The name of the datacenter in which to create the
	// disk. Can be omitted when ESXi or if there is only one datacenter in
	// your infrastructure.
	Datacenter pulumi.StringPtrInput
	// The name of the datastore in which to create the
	// disk.
	Datastore pulumi.StringPtrInput
	// Size of the disk (in GB). Decreasing the size of a disk is not possible.
	// If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
	// created.
	Size pulumi.IntPtrInput
	// The type of disk to create. Can be one of
	// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
	// information on what each kind of disk provisioning policy means, click
	// [here][docs-vmware-vm-disk-provisioning].
	//
	// [docs-vmware-vm-disk-provisioning]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-single-host-management-vmware-host-client-8-0/virtual-machine-management-with-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/configuring-virtual-machines-in-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/virtual-disk-configuration-vSphereSingleHostManagementVMwareHostClient/about-virtual-disk-provisioning-policies-vSphereSingleHostManagementVMwareHostClient.html
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
	// > **NOTE:** `adapterType` is **deprecated**: it does not dictate the type of
	// controller that the virtual disk will be attached to on the virtual machine.
	// Please see the `scsiType` parameter
	// in the `VirtualMachine` resource for information on how to control
	// disk controller types. This parameter will be removed in future versions of the
	// vSphere provider.
	//
	// Deprecated: this attribute has no effect on controller types - please use scsiType in VirtualMachine instead
	AdapterType *string `pulumi:"adapterType"`
	// Tells the resource to create any
	// directories that are a part of the `vmdkPath` parameter if they are missing.
	// Default: `false`.
	//
	// > **NOTE:** Any directory created as part of the operation when
	// `createDirectories` is enabled will not be deleted when the resource is
	// destroyed.
	CreateDirectories *bool `pulumi:"createDirectories"`
	// The name of the datacenter in which to create the
	// disk. Can be omitted when ESXi or if there is only one datacenter in
	// your infrastructure.
	Datacenter *string `pulumi:"datacenter"`
	// The name of the datastore in which to create the
	// disk.
	Datastore string `pulumi:"datastore"`
	// Size of the disk (in GB). Decreasing the size of a disk is not possible.
	// If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
	// created.
	Size int `pulumi:"size"`
	// The type of disk to create. Can be one of
	// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
	// information on what each kind of disk provisioning policy means, click
	// [here][docs-vmware-vm-disk-provisioning].
	//
	// [docs-vmware-vm-disk-provisioning]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-single-host-management-vmware-host-client-8-0/virtual-machine-management-with-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/configuring-virtual-machines-in-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/virtual-disk-configuration-vSphereSingleHostManagementVMwareHostClient/about-virtual-disk-provisioning-policies-vSphereSingleHostManagementVMwareHostClient.html
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
	// > **NOTE:** `adapterType` is **deprecated**: it does not dictate the type of
	// controller that the virtual disk will be attached to on the virtual machine.
	// Please see the `scsiType` parameter
	// in the `VirtualMachine` resource for information on how to control
	// disk controller types. This parameter will be removed in future versions of the
	// vSphere provider.
	//
	// Deprecated: this attribute has no effect on controller types - please use scsiType in VirtualMachine instead
	AdapterType pulumi.StringPtrInput
	// Tells the resource to create any
	// directories that are a part of the `vmdkPath` parameter if they are missing.
	// Default: `false`.
	//
	// > **NOTE:** Any directory created as part of the operation when
	// `createDirectories` is enabled will not be deleted when the resource is
	// destroyed.
	CreateDirectories pulumi.BoolPtrInput
	// The name of the datacenter in which to create the
	// disk. Can be omitted when ESXi or if there is only one datacenter in
	// your infrastructure.
	Datacenter pulumi.StringPtrInput
	// The name of the datastore in which to create the
	// disk.
	Datastore pulumi.StringInput
	// Size of the disk (in GB). Decreasing the size of a disk is not possible.
	// If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
	// created.
	Size pulumi.IntInput
	// The type of disk to create. Can be one of
	// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
	// information on what each kind of disk provisioning policy means, click
	// [here][docs-vmware-vm-disk-provisioning].
	//
	// [docs-vmware-vm-disk-provisioning]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-single-host-management-vmware-host-client-8-0/virtual-machine-management-with-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/configuring-virtual-machines-in-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/virtual-disk-configuration-vSphereSingleHostManagementVMwareHostClient/about-virtual-disk-provisioning-policies-vSphereSingleHostManagementVMwareHostClient.html
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
	return reflect.TypeOf((**VirtualDisk)(nil)).Elem()
}

func (i *VirtualDisk) ToVirtualDiskOutput() VirtualDiskOutput {
	return i.ToVirtualDiskOutputWithContext(context.Background())
}

func (i *VirtualDisk) ToVirtualDiskOutputWithContext(ctx context.Context) VirtualDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskOutput)
}

// VirtualDiskArrayInput is an input type that accepts VirtualDiskArray and VirtualDiskArrayOutput values.
// You can construct a concrete instance of `VirtualDiskArrayInput` via:
//
//	VirtualDiskArray{ VirtualDiskArgs{...} }
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
//	VirtualDiskMap{ "key": VirtualDiskArgs{...} }
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
	return reflect.TypeOf((**VirtualDisk)(nil)).Elem()
}

func (o VirtualDiskOutput) ToVirtualDiskOutput() VirtualDiskOutput {
	return o
}

func (o VirtualDiskOutput) ToVirtualDiskOutputWithContext(ctx context.Context) VirtualDiskOutput {
	return o
}

// The adapter type for this virtual disk. Can be
// one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
//
// > **NOTE:** `adapterType` is **deprecated**: it does not dictate the type of
// controller that the virtual disk will be attached to on the virtual machine.
// Please see the `scsiType` parameter
// in the `VirtualMachine` resource for information on how to control
// disk controller types. This parameter will be removed in future versions of the
// vSphere provider.
//
// Deprecated: this attribute has no effect on controller types - please use scsiType in VirtualMachine instead
func (o VirtualDiskOutput) AdapterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualDisk) pulumi.StringPtrOutput { return v.AdapterType }).(pulumi.StringPtrOutput)
}

// Tells the resource to create any
// directories that are a part of the `vmdkPath` parameter if they are missing.
// Default: `false`.
//
// > **NOTE:** Any directory created as part of the operation when
// `createDirectories` is enabled will not be deleted when the resource is
// destroyed.
func (o VirtualDiskOutput) CreateDirectories() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualDisk) pulumi.BoolPtrOutput { return v.CreateDirectories }).(pulumi.BoolPtrOutput)
}

// The name of the datacenter in which to create the
// disk. Can be omitted when ESXi or if there is only one datacenter in
// your infrastructure.
func (o VirtualDiskOutput) Datacenter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualDisk) pulumi.StringPtrOutput { return v.Datacenter }).(pulumi.StringPtrOutput)
}

// The name of the datastore in which to create the
// disk.
func (o VirtualDiskOutput) Datastore() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualDisk) pulumi.StringOutput { return v.Datastore }).(pulumi.StringOutput)
}

// Size of the disk (in GB). Decreasing the size of a disk is not possible.
// If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
// created.
func (o VirtualDiskOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualDisk) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The type of disk to create. Can be one of
// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
// information on what each kind of disk provisioning policy means, click
// [here][docs-vmware-vm-disk-provisioning].
//
// [docs-vmware-vm-disk-provisioning]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-single-host-management-vmware-host-client-8-0/virtual-machine-management-with-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/configuring-virtual-machines-in-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/virtual-disk-configuration-vSphereSingleHostManagementVMwareHostClient/about-virtual-disk-provisioning-policies-vSphereSingleHostManagementVMwareHostClient.html
func (o VirtualDiskOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualDisk) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// The path, including filename, of the virtual disk to
// be created.  This needs to end in `.vmdk`.
func (o VirtualDiskOutput) VmdkPath() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualDisk) pulumi.StringOutput { return v.VmdkPath }).(pulumi.StringOutput)
}

type VirtualDiskArrayOutput struct{ *pulumi.OutputState }

func (VirtualDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualDisk)(nil)).Elem()
}

func (o VirtualDiskArrayOutput) ToVirtualDiskArrayOutput() VirtualDiskArrayOutput {
	return o
}

func (o VirtualDiskArrayOutput) ToVirtualDiskArrayOutputWithContext(ctx context.Context) VirtualDiskArrayOutput {
	return o
}

func (o VirtualDiskArrayOutput) Index(i pulumi.IntInput) VirtualDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualDisk {
		return vs[0].([]*VirtualDisk)[vs[1].(int)]
	}).(VirtualDiskOutput)
}

type VirtualDiskMapOutput struct{ *pulumi.OutputState }

func (VirtualDiskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualDisk)(nil)).Elem()
}

func (o VirtualDiskMapOutput) ToVirtualDiskMapOutput() VirtualDiskMapOutput {
	return o
}

func (o VirtualDiskMapOutput) ToVirtualDiskMapOutputWithContext(ctx context.Context) VirtualDiskMapOutput {
	return o
}

func (o VirtualDiskMapOutput) MapIndex(k pulumi.StringInput) VirtualDiskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualDisk {
		return vs[0].(map[string]*VirtualDisk)[vs[1].(string)]
	}).(VirtualDiskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDiskInput)(nil)).Elem(), &VirtualDisk{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDiskArrayInput)(nil)).Elem(), VirtualDiskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualDiskMapInput)(nil)).Elem(), VirtualDiskMap{})
	pulumi.RegisterOutputType(VirtualDiskOutput{})
	pulumi.RegisterOutputType(VirtualDiskArrayOutput{})
	pulumi.RegisterOutputType(VirtualDiskMapOutput{})
}

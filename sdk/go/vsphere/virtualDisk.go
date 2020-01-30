// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `.VirtualDisk` resource can be used to create virtual disks outside
// of any given [`.VirtualMachine`][docs-vsphere-virtual-machine]
// resource. These disks can be attached to a virtual machine by creating a disk
// block with the [`attach`][docs-vsphere-virtual-machine-disk-attach] parameter.
//
// [docs-vsphere-virtual-machine]: /docs/providers/vsphere/r/virtual_machine.html
// [docs-vsphere-virtual-machine-disk-attach]: /docs/providers/vsphere/r/virtual_machine.html#attach
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/virtual_disk.html.markdown.
type VirtualDisk struct {
	s *pulumi.ResourceState
}

// NewVirtualDisk registers a new resource with the given unique name, arguments, and options.
func NewVirtualDisk(ctx *pulumi.Context,
	name string, args *VirtualDiskArgs, opts ...pulumi.ResourceOpt) (*VirtualDisk, error) {
	if args == nil || args.Datastore == nil {
		return nil, errors.New("missing required argument 'Datastore'")
	}
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	if args == nil || args.VmdkPath == nil {
		return nil, errors.New("missing required argument 'VmdkPath'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["adapterType"] = nil
		inputs["createDirectories"] = nil
		inputs["datacenter"] = nil
		inputs["datastore"] = nil
		inputs["size"] = nil
		inputs["type"] = nil
		inputs["vmdkPath"] = nil
	} else {
		inputs["adapterType"] = args.AdapterType
		inputs["createDirectories"] = args.CreateDirectories
		inputs["datacenter"] = args.Datacenter
		inputs["datastore"] = args.Datastore
		inputs["size"] = args.Size
		inputs["type"] = args.Type
		inputs["vmdkPath"] = args.VmdkPath
	}
	s, err := ctx.RegisterResource("vsphere:index/virtualDisk:VirtualDisk", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VirtualDisk{s: s}, nil
}

// GetVirtualDisk gets an existing VirtualDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualDisk(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VirtualDiskState, opts ...pulumi.ResourceOpt) (*VirtualDisk, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["adapterType"] = state.AdapterType
		inputs["createDirectories"] = state.CreateDirectories
		inputs["datacenter"] = state.Datacenter
		inputs["datastore"] = state.Datastore
		inputs["size"] = state.Size
		inputs["type"] = state.Type
		inputs["vmdkPath"] = state.VmdkPath
	}
	s, err := ctx.ReadResource("vsphere:index/virtualDisk:VirtualDisk", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VirtualDisk{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VirtualDisk) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VirtualDisk) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The adapter type for this virtual disk. Can be
// one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
func (r *VirtualDisk) AdapterType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["adapterType"])
}

// Tells the resource to create any
// directories that are a part of the `vmdkPath` parameter if they are missing.
// Default: `false`.
func (r *VirtualDisk) CreateDirectories() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["createDirectories"])
}

// The name of the datacenter in which to create the
// disk. Can be omitted when when ESXi or if there is only one datacenter in
// your infrastructure.
func (r *VirtualDisk) Datacenter() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["datacenter"])
}

// The name of the datastore in which to create the
// disk.
func (r *VirtualDisk) Datastore() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["datastore"])
}

// Size of the disk (in GB).
func (r *VirtualDisk) Size() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["size"])
}

// The type of disk to create. Can be one of
// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
// information on what each kind of disk provisioning policy means, click
// [here][docs-vmware-vm-disk-provisioning].
func (r *VirtualDisk) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

// The path, including filename, of the virtual disk to
// be created.  This needs to end in `.vmdk`.
func (r *VirtualDisk) VmdkPath() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["vmdkPath"])
}

// Input properties used for looking up and filtering VirtualDisk resources.
type VirtualDiskState struct {
	// The adapter type for this virtual disk. Can be
	// one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
	AdapterType interface{}
	// Tells the resource to create any
	// directories that are a part of the `vmdkPath` parameter if they are missing.
	// Default: `false`.
	CreateDirectories interface{}
	// The name of the datacenter in which to create the
	// disk. Can be omitted when when ESXi or if there is only one datacenter in
	// your infrastructure.
	Datacenter interface{}
	// The name of the datastore in which to create the
	// disk.
	Datastore interface{}
	// Size of the disk (in GB).
	Size interface{}
	// The type of disk to create. Can be one of
	// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
	// information on what each kind of disk provisioning policy means, click
	// [here][docs-vmware-vm-disk-provisioning].
	Type interface{}
	// The path, including filename, of the virtual disk to
	// be created.  This needs to end in `.vmdk`.
	VmdkPath interface{}
}

// The set of arguments for constructing a VirtualDisk resource.
type VirtualDiskArgs struct {
	// The adapter type for this virtual disk. Can be
	// one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
	AdapterType interface{}
	// Tells the resource to create any
	// directories that are a part of the `vmdkPath` parameter if they are missing.
	// Default: `false`.
	CreateDirectories interface{}
	// The name of the datacenter in which to create the
	// disk. Can be omitted when when ESXi or if there is only one datacenter in
	// your infrastructure.
	Datacenter interface{}
	// The name of the datastore in which to create the
	// disk.
	Datastore interface{}
	// Size of the disk (in GB).
	Size interface{}
	// The type of disk to create. Can be one of
	// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
	// information on what each kind of disk provisioning policy means, click
	// [here][docs-vmware-vm-disk-provisioning].
	Type interface{}
	// The path, including filename, of the virtual disk to
	// be created.  This needs to end in `.vmdk`.
	VmdkPath interface{}
}

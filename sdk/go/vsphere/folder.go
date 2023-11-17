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

type Folder struct {
	pulumi.CustomResourceState

	// Map of custom attribute ids to attribute
	// value strings to set for folder. See [here][docs-setting-custom-attributes]
	// for a reference on how to set values for custom attributes.
	//
	// > **NOTE:** Custom attributes are unsupported on direct ESXi connections
	// and require vCenter.
	CustomAttributes pulumi.StringMapOutput `pulumi:"customAttributes"`
	// The ID of the datacenter the folder will be created in.
	// Required for all folder types except for datacenter folders. Forces a new
	// resource if changed.
	DatacenterId pulumi.StringPtrOutput `pulumi:"datacenterId"`
	// The path of the folder to be created. This is relative to
	// the root of the type of folder you are creating, and the supplied datacenter.
	// For example, given a default datacenter of `default-dc`, a folder of type
	// `vm` (denoting a virtual machine folder), and a supplied folder of
	// `test-folder`, the resulting path would be
	// `/default-dc/vm/test-folder`.
	//
	// > **NOTE:** `path` can be modified - the resulting behavior is dependent on
	// what section of `path` you are modifying. If you are modifying the parent (so
	// any part before the last `/`), your folder will be moved to that new parent. If
	// modifying the name (the part after the last `/`), your folder will be renamed.
	Path pulumi.StringOutput `pulumi:"path"`
	// The IDs of any tags to attach to this resource.
	//
	// > **NOTE:** Tagging support is unsupported on direct ESXi connections and
	// requires vCenter 6.0 or higher.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The type of folder to create. Allowed options are
	// `datacenter` for datacenter folders, `host` for host and cluster folders,
	// `vm` for virtual machine folders, `datastore` for datastore folders, and
	// `network` for network folders. Forces a new resource if changed.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFolder registers a new resource with the given unique name, arguments, and options.
func NewFolder(ctx *pulumi.Context,
	name string, args *FolderArgs, opts ...pulumi.ResourceOption) (*Folder, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Folder
	err := ctx.RegisterResource("vsphere:index/folder:Folder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolder gets an existing Folder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderState, opts ...pulumi.ResourceOption) (*Folder, error) {
	var resource Folder
	err := ctx.ReadResource("vsphere:index/folder:Folder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Folder resources.
type folderState struct {
	// Map of custom attribute ids to attribute
	// value strings to set for folder. See [here][docs-setting-custom-attributes]
	// for a reference on how to set values for custom attributes.
	//
	// > **NOTE:** Custom attributes are unsupported on direct ESXi connections
	// and require vCenter.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The ID of the datacenter the folder will be created in.
	// Required for all folder types except for datacenter folders. Forces a new
	// resource if changed.
	DatacenterId *string `pulumi:"datacenterId"`
	// The path of the folder to be created. This is relative to
	// the root of the type of folder you are creating, and the supplied datacenter.
	// For example, given a default datacenter of `default-dc`, a folder of type
	// `vm` (denoting a virtual machine folder), and a supplied folder of
	// `test-folder`, the resulting path would be
	// `/default-dc/vm/test-folder`.
	//
	// > **NOTE:** `path` can be modified - the resulting behavior is dependent on
	// what section of `path` you are modifying. If you are modifying the parent (so
	// any part before the last `/`), your folder will be moved to that new parent. If
	// modifying the name (the part after the last `/`), your folder will be renamed.
	Path *string `pulumi:"path"`
	// The IDs of any tags to attach to this resource.
	//
	// > **NOTE:** Tagging support is unsupported on direct ESXi connections and
	// requires vCenter 6.0 or higher.
	Tags []string `pulumi:"tags"`
	// The type of folder to create. Allowed options are
	// `datacenter` for datacenter folders, `host` for host and cluster folders,
	// `vm` for virtual machine folders, `datastore` for datastore folders, and
	// `network` for network folders. Forces a new resource if changed.
	Type *string `pulumi:"type"`
}

type FolderState struct {
	// Map of custom attribute ids to attribute
	// value strings to set for folder. See [here][docs-setting-custom-attributes]
	// for a reference on how to set values for custom attributes.
	//
	// > **NOTE:** Custom attributes are unsupported on direct ESXi connections
	// and require vCenter.
	CustomAttributes pulumi.StringMapInput
	// The ID of the datacenter the folder will be created in.
	// Required for all folder types except for datacenter folders. Forces a new
	// resource if changed.
	DatacenterId pulumi.StringPtrInput
	// The path of the folder to be created. This is relative to
	// the root of the type of folder you are creating, and the supplied datacenter.
	// For example, given a default datacenter of `default-dc`, a folder of type
	// `vm` (denoting a virtual machine folder), and a supplied folder of
	// `test-folder`, the resulting path would be
	// `/default-dc/vm/test-folder`.
	//
	// > **NOTE:** `path` can be modified - the resulting behavior is dependent on
	// what section of `path` you are modifying. If you are modifying the parent (so
	// any part before the last `/`), your folder will be moved to that new parent. If
	// modifying the name (the part after the last `/`), your folder will be renamed.
	Path pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	//
	// > **NOTE:** Tagging support is unsupported on direct ESXi connections and
	// requires vCenter 6.0 or higher.
	Tags pulumi.StringArrayInput
	// The type of folder to create. Allowed options are
	// `datacenter` for datacenter folders, `host` for host and cluster folders,
	// `vm` for virtual machine folders, `datastore` for datastore folders, and
	// `network` for network folders. Forces a new resource if changed.
	Type pulumi.StringPtrInput
}

func (FolderState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderState)(nil)).Elem()
}

type folderArgs struct {
	// Map of custom attribute ids to attribute
	// value strings to set for folder. See [here][docs-setting-custom-attributes]
	// for a reference on how to set values for custom attributes.
	//
	// > **NOTE:** Custom attributes are unsupported on direct ESXi connections
	// and require vCenter.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The ID of the datacenter the folder will be created in.
	// Required for all folder types except for datacenter folders. Forces a new
	// resource if changed.
	DatacenterId *string `pulumi:"datacenterId"`
	// The path of the folder to be created. This is relative to
	// the root of the type of folder you are creating, and the supplied datacenter.
	// For example, given a default datacenter of `default-dc`, a folder of type
	// `vm` (denoting a virtual machine folder), and a supplied folder of
	// `test-folder`, the resulting path would be
	// `/default-dc/vm/test-folder`.
	//
	// > **NOTE:** `path` can be modified - the resulting behavior is dependent on
	// what section of `path` you are modifying. If you are modifying the parent (so
	// any part before the last `/`), your folder will be moved to that new parent. If
	// modifying the name (the part after the last `/`), your folder will be renamed.
	Path string `pulumi:"path"`
	// The IDs of any tags to attach to this resource.
	//
	// > **NOTE:** Tagging support is unsupported on direct ESXi connections and
	// requires vCenter 6.0 or higher.
	Tags []string `pulumi:"tags"`
	// The type of folder to create. Allowed options are
	// `datacenter` for datacenter folders, `host` for host and cluster folders,
	// `vm` for virtual machine folders, `datastore` for datastore folders, and
	// `network` for network folders. Forces a new resource if changed.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Folder resource.
type FolderArgs struct {
	// Map of custom attribute ids to attribute
	// value strings to set for folder. See [here][docs-setting-custom-attributes]
	// for a reference on how to set values for custom attributes.
	//
	// > **NOTE:** Custom attributes are unsupported on direct ESXi connections
	// and require vCenter.
	CustomAttributes pulumi.StringMapInput
	// The ID of the datacenter the folder will be created in.
	// Required for all folder types except for datacenter folders. Forces a new
	// resource if changed.
	DatacenterId pulumi.StringPtrInput
	// The path of the folder to be created. This is relative to
	// the root of the type of folder you are creating, and the supplied datacenter.
	// For example, given a default datacenter of `default-dc`, a folder of type
	// `vm` (denoting a virtual machine folder), and a supplied folder of
	// `test-folder`, the resulting path would be
	// `/default-dc/vm/test-folder`.
	//
	// > **NOTE:** `path` can be modified - the resulting behavior is dependent on
	// what section of `path` you are modifying. If you are modifying the parent (so
	// any part before the last `/`), your folder will be moved to that new parent. If
	// modifying the name (the part after the last `/`), your folder will be renamed.
	Path pulumi.StringInput
	// The IDs of any tags to attach to this resource.
	//
	// > **NOTE:** Tagging support is unsupported on direct ESXi connections and
	// requires vCenter 6.0 or higher.
	Tags pulumi.StringArrayInput
	// The type of folder to create. Allowed options are
	// `datacenter` for datacenter folders, `host` for host and cluster folders,
	// `vm` for virtual machine folders, `datastore` for datastore folders, and
	// `network` for network folders. Forces a new resource if changed.
	Type pulumi.StringInput
}

func (FolderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderArgs)(nil)).Elem()
}

type FolderInput interface {
	pulumi.Input

	ToFolderOutput() FolderOutput
	ToFolderOutputWithContext(ctx context.Context) FolderOutput
}

func (*Folder) ElementType() reflect.Type {
	return reflect.TypeOf((**Folder)(nil)).Elem()
}

func (i *Folder) ToFolderOutput() FolderOutput {
	return i.ToFolderOutputWithContext(context.Background())
}

func (i *Folder) ToFolderOutputWithContext(ctx context.Context) FolderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderOutput)
}

// FolderArrayInput is an input type that accepts FolderArray and FolderArrayOutput values.
// You can construct a concrete instance of `FolderArrayInput` via:
//
//	FolderArray{ FolderArgs{...} }
type FolderArrayInput interface {
	pulumi.Input

	ToFolderArrayOutput() FolderArrayOutput
	ToFolderArrayOutputWithContext(context.Context) FolderArrayOutput
}

type FolderArray []FolderInput

func (FolderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Folder)(nil)).Elem()
}

func (i FolderArray) ToFolderArrayOutput() FolderArrayOutput {
	return i.ToFolderArrayOutputWithContext(context.Background())
}

func (i FolderArray) ToFolderArrayOutputWithContext(ctx context.Context) FolderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderArrayOutput)
}

// FolderMapInput is an input type that accepts FolderMap and FolderMapOutput values.
// You can construct a concrete instance of `FolderMapInput` via:
//
//	FolderMap{ "key": FolderArgs{...} }
type FolderMapInput interface {
	pulumi.Input

	ToFolderMapOutput() FolderMapOutput
	ToFolderMapOutputWithContext(context.Context) FolderMapOutput
}

type FolderMap map[string]FolderInput

func (FolderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Folder)(nil)).Elem()
}

func (i FolderMap) ToFolderMapOutput() FolderMapOutput {
	return i.ToFolderMapOutputWithContext(context.Background())
}

func (i FolderMap) ToFolderMapOutputWithContext(ctx context.Context) FolderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderMapOutput)
}

type FolderOutput struct{ *pulumi.OutputState }

func (FolderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Folder)(nil)).Elem()
}

func (o FolderOutput) ToFolderOutput() FolderOutput {
	return o
}

func (o FolderOutput) ToFolderOutputWithContext(ctx context.Context) FolderOutput {
	return o
}

// Map of custom attribute ids to attribute
// value strings to set for folder. See [here][docs-setting-custom-attributes]
// for a reference on how to set values for custom attributes.
//
// > **NOTE:** Custom attributes are unsupported on direct ESXi connections
// and require vCenter.
func (o FolderOutput) CustomAttributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringMapOutput { return v.CustomAttributes }).(pulumi.StringMapOutput)
}

// The ID of the datacenter the folder will be created in.
// Required for all folder types except for datacenter folders. Forces a new
// resource if changed.
func (o FolderOutput) DatacenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringPtrOutput { return v.DatacenterId }).(pulumi.StringPtrOutput)
}

// The path of the folder to be created. This is relative to
// the root of the type of folder you are creating, and the supplied datacenter.
// For example, given a default datacenter of `default-dc`, a folder of type
// `vm` (denoting a virtual machine folder), and a supplied folder of
// `test-folder`, the resulting path would be
// `/default-dc/vm/test-folder`.
//
// > **NOTE:** `path` can be modified - the resulting behavior is dependent on
// what section of `path` you are modifying. If you are modifying the parent (so
// any part before the last `/`), your folder will be moved to that new parent. If
// modifying the name (the part after the last `/`), your folder will be renamed.
func (o FolderOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// The IDs of any tags to attach to this resource.
//
// > **NOTE:** Tagging support is unsupported on direct ESXi connections and
// requires vCenter 6.0 or higher.
func (o FolderOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of folder to create. Allowed options are
// `datacenter` for datacenter folders, `host` for host and cluster folders,
// `vm` for virtual machine folders, `datastore` for datastore folders, and
// `network` for network folders. Forces a new resource if changed.
func (o FolderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Folder) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type FolderArrayOutput struct{ *pulumi.OutputState }

func (FolderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Folder)(nil)).Elem()
}

func (o FolderArrayOutput) ToFolderArrayOutput() FolderArrayOutput {
	return o
}

func (o FolderArrayOutput) ToFolderArrayOutputWithContext(ctx context.Context) FolderArrayOutput {
	return o
}

func (o FolderArrayOutput) Index(i pulumi.IntInput) FolderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Folder {
		return vs[0].([]*Folder)[vs[1].(int)]
	}).(FolderOutput)
}

type FolderMapOutput struct{ *pulumi.OutputState }

func (FolderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Folder)(nil)).Elem()
}

func (o FolderMapOutput) ToFolderMapOutput() FolderMapOutput {
	return o
}

func (o FolderMapOutput) ToFolderMapOutputWithContext(ctx context.Context) FolderMapOutput {
	return o
}

func (o FolderMapOutput) MapIndex(k pulumi.StringInput) FolderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Folder {
		return vs[0].(map[string]*Folder)[vs[1].(string)]
	}).(FolderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FolderInput)(nil)).Elem(), &Folder{})
	pulumi.RegisterInputType(reflect.TypeOf((*FolderArrayInput)(nil)).Elem(), FolderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FolderMapInput)(nil)).Elem(), FolderMap{})
	pulumi.RegisterOutputType(FolderOutput{})
	pulumi.RegisterOutputType(FolderArrayOutput{})
	pulumi.RegisterOutputType(FolderMapOutput{})
}

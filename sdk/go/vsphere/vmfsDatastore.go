// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VmfsDatastore struct {
	pulumi.CustomResourceState

	// The connectivity status of the datastore. If this is `false`,
	// some other computed attributes may be out of date.
	Accessible pulumi.BoolOutput `pulumi:"accessible"`
	// Maximum capacity of the datastore, in megabytes.
	Capacity pulumi.IntOutput `pulumi:"capacity"`
	// Map of custom attribute ids to attribute
	// value string to set on datastore resource.
	CustomAttributes pulumi.StringMapOutput `pulumi:"customAttributes"`
	// The managed object
	// ID of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId pulumi.StringPtrOutput `pulumi:"datastoreClusterId"`
	// The disks to use with the datastore.
	Disks pulumi.StringArrayOutput `pulumi:"disks"`
	// The relative path to a folder to put this datastore in.
	// This is a path relative to the datacenter you are deploying the datastore to.
	// Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
	// The provider will place a datastore named `test` in a datastore folder
	// located at `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/test`. Conflicts with
	// `datastoreClusterId`.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// Available space of this datastore, in megabytes.
	FreeSpace pulumi.IntOutput `pulumi:"freeSpace"`
	// The managed object ID of
	// the host to set the datastore up on. Note that this is not necessarily the
	// only host that the datastore will be set up on - see
	// here for more info. Forces a
	// new resource if changed.
	HostSystemId pulumi.StringOutput `pulumi:"hostSystemId"`
	// The current maintenance mode state of the datastore.
	MaintenanceMode pulumi.StringOutput `pulumi:"maintenanceMode"`
	// If `true`, more than one host in the datacenter has
	// been configured with access to the datastore.
	MultipleHostAccess pulumi.BoolOutput `pulumi:"multipleHostAccess"`
	// The name of the datastore. Forces a new resource if
	// changed.
	Name pulumi.StringOutput `pulumi:"name"`
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Total additional storage space, in megabytes,
	// potentially used by all virtual machines on this datastore.
	UncommittedSpace pulumi.IntOutput `pulumi:"uncommittedSpace"`
	// The unique locator for the datastore.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewVmfsDatastore registers a new resource with the given unique name, arguments, and options.
func NewVmfsDatastore(ctx *pulumi.Context,
	name string, args *VmfsDatastoreArgs, opts ...pulumi.ResourceOption) (*VmfsDatastore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Disks == nil {
		return nil, errors.New("invalid value for required argument 'Disks'")
	}
	if args.HostSystemId == nil {
		return nil, errors.New("invalid value for required argument 'HostSystemId'")
	}
	var resource VmfsDatastore
	err := ctx.RegisterResource("vsphere:index/vmfsDatastore:VmfsDatastore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVmfsDatastore gets an existing VmfsDatastore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVmfsDatastore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VmfsDatastoreState, opts ...pulumi.ResourceOption) (*VmfsDatastore, error) {
	var resource VmfsDatastore
	err := ctx.ReadResource("vsphere:index/vmfsDatastore:VmfsDatastore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VmfsDatastore resources.
type vmfsDatastoreState struct {
	// The connectivity status of the datastore. If this is `false`,
	// some other computed attributes may be out of date.
	Accessible *bool `pulumi:"accessible"`
	// Maximum capacity of the datastore, in megabytes.
	Capacity *int `pulumi:"capacity"`
	// Map of custom attribute ids to attribute
	// value string to set on datastore resource.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The managed object
	// ID of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId *string `pulumi:"datastoreClusterId"`
	// The disks to use with the datastore.
	Disks []string `pulumi:"disks"`
	// The relative path to a folder to put this datastore in.
	// This is a path relative to the datacenter you are deploying the datastore to.
	// Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
	// The provider will place a datastore named `test` in a datastore folder
	// located at `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/test`. Conflicts with
	// `datastoreClusterId`.
	Folder *string `pulumi:"folder"`
	// Available space of this datastore, in megabytes.
	FreeSpace *int `pulumi:"freeSpace"`
	// The managed object ID of
	// the host to set the datastore up on. Note that this is not necessarily the
	// only host that the datastore will be set up on - see
	// here for more info. Forces a
	// new resource if changed.
	HostSystemId *string `pulumi:"hostSystemId"`
	// The current maintenance mode state of the datastore.
	MaintenanceMode *string `pulumi:"maintenanceMode"`
	// If `true`, more than one host in the datacenter has
	// been configured with access to the datastore.
	MultipleHostAccess *bool `pulumi:"multipleHostAccess"`
	// The name of the datastore. Forces a new resource if
	// changed.
	Name *string `pulumi:"name"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
	// Total additional storage space, in megabytes,
	// potentially used by all virtual machines on this datastore.
	UncommittedSpace *int `pulumi:"uncommittedSpace"`
	// The unique locator for the datastore.
	Url *string `pulumi:"url"`
}

type VmfsDatastoreState struct {
	// The connectivity status of the datastore. If this is `false`,
	// some other computed attributes may be out of date.
	Accessible pulumi.BoolPtrInput
	// Maximum capacity of the datastore, in megabytes.
	Capacity pulumi.IntPtrInput
	// Map of custom attribute ids to attribute
	// value string to set on datastore resource.
	CustomAttributes pulumi.StringMapInput
	// The managed object
	// ID of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId pulumi.StringPtrInput
	// The disks to use with the datastore.
	Disks pulumi.StringArrayInput
	// The relative path to a folder to put this datastore in.
	// This is a path relative to the datacenter you are deploying the datastore to.
	// Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
	// The provider will place a datastore named `test` in a datastore folder
	// located at `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/test`. Conflicts with
	// `datastoreClusterId`.
	Folder pulumi.StringPtrInput
	// Available space of this datastore, in megabytes.
	FreeSpace pulumi.IntPtrInput
	// The managed object ID of
	// the host to set the datastore up on. Note that this is not necessarily the
	// only host that the datastore will be set up on - see
	// here for more info. Forces a
	// new resource if changed.
	HostSystemId pulumi.StringPtrInput
	// The current maintenance mode state of the datastore.
	MaintenanceMode pulumi.StringPtrInput
	// If `true`, more than one host in the datacenter has
	// been configured with access to the datastore.
	MultipleHostAccess pulumi.BoolPtrInput
	// The name of the datastore. Forces a new resource if
	// changed.
	Name pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
	// Total additional storage space, in megabytes,
	// potentially used by all virtual machines on this datastore.
	UncommittedSpace pulumi.IntPtrInput
	// The unique locator for the datastore.
	Url pulumi.StringPtrInput
}

func (VmfsDatastoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmfsDatastoreState)(nil)).Elem()
}

type vmfsDatastoreArgs struct {
	// Map of custom attribute ids to attribute
	// value string to set on datastore resource.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The managed object
	// ID of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId *string `pulumi:"datastoreClusterId"`
	// The disks to use with the datastore.
	Disks []string `pulumi:"disks"`
	// The relative path to a folder to put this datastore in.
	// This is a path relative to the datacenter you are deploying the datastore to.
	// Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
	// The provider will place a datastore named `test` in a datastore folder
	// located at `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/test`. Conflicts with
	// `datastoreClusterId`.
	Folder *string `pulumi:"folder"`
	// The managed object ID of
	// the host to set the datastore up on. Note that this is not necessarily the
	// only host that the datastore will be set up on - see
	// here for more info. Forces a
	// new resource if changed.
	HostSystemId string `pulumi:"hostSystemId"`
	// The name of the datastore. Forces a new resource if
	// changed.
	Name *string `pulumi:"name"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a VmfsDatastore resource.
type VmfsDatastoreArgs struct {
	// Map of custom attribute ids to attribute
	// value string to set on datastore resource.
	CustomAttributes pulumi.StringMapInput
	// The managed object
	// ID of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId pulumi.StringPtrInput
	// The disks to use with the datastore.
	Disks pulumi.StringArrayInput
	// The relative path to a folder to put this datastore in.
	// This is a path relative to the datacenter you are deploying the datastore to.
	// Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
	// The provider will place a datastore named `test` in a datastore folder
	// located at `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/test`. Conflicts with
	// `datastoreClusterId`.
	Folder pulumi.StringPtrInput
	// The managed object ID of
	// the host to set the datastore up on. Note that this is not necessarily the
	// only host that the datastore will be set up on - see
	// here for more info. Forces a
	// new resource if changed.
	HostSystemId pulumi.StringInput
	// The name of the datastore. Forces a new resource if
	// changed.
	Name pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
}

func (VmfsDatastoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmfsDatastoreArgs)(nil)).Elem()
}

type VmfsDatastoreInput interface {
	pulumi.Input

	ToVmfsDatastoreOutput() VmfsDatastoreOutput
	ToVmfsDatastoreOutputWithContext(ctx context.Context) VmfsDatastoreOutput
}

func (*VmfsDatastore) ElementType() reflect.Type {
	return reflect.TypeOf((*VmfsDatastore)(nil))
}

func (i *VmfsDatastore) ToVmfsDatastoreOutput() VmfsDatastoreOutput {
	return i.ToVmfsDatastoreOutputWithContext(context.Background())
}

func (i *VmfsDatastore) ToVmfsDatastoreOutputWithContext(ctx context.Context) VmfsDatastoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmfsDatastoreOutput)
}

func (i *VmfsDatastore) ToVmfsDatastorePtrOutput() VmfsDatastorePtrOutput {
	return i.ToVmfsDatastorePtrOutputWithContext(context.Background())
}

func (i *VmfsDatastore) ToVmfsDatastorePtrOutputWithContext(ctx context.Context) VmfsDatastorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmfsDatastorePtrOutput)
}

type VmfsDatastorePtrInput interface {
	pulumi.Input

	ToVmfsDatastorePtrOutput() VmfsDatastorePtrOutput
	ToVmfsDatastorePtrOutputWithContext(ctx context.Context) VmfsDatastorePtrOutput
}

type vmfsDatastorePtrType VmfsDatastoreArgs

func (*vmfsDatastorePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VmfsDatastore)(nil))
}

func (i *vmfsDatastorePtrType) ToVmfsDatastorePtrOutput() VmfsDatastorePtrOutput {
	return i.ToVmfsDatastorePtrOutputWithContext(context.Background())
}

func (i *vmfsDatastorePtrType) ToVmfsDatastorePtrOutputWithContext(ctx context.Context) VmfsDatastorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmfsDatastorePtrOutput)
}

// VmfsDatastoreArrayInput is an input type that accepts VmfsDatastoreArray and VmfsDatastoreArrayOutput values.
// You can construct a concrete instance of `VmfsDatastoreArrayInput` via:
//
//          VmfsDatastoreArray{ VmfsDatastoreArgs{...} }
type VmfsDatastoreArrayInput interface {
	pulumi.Input

	ToVmfsDatastoreArrayOutput() VmfsDatastoreArrayOutput
	ToVmfsDatastoreArrayOutputWithContext(context.Context) VmfsDatastoreArrayOutput
}

type VmfsDatastoreArray []VmfsDatastoreInput

func (VmfsDatastoreArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VmfsDatastore)(nil))
}

func (i VmfsDatastoreArray) ToVmfsDatastoreArrayOutput() VmfsDatastoreArrayOutput {
	return i.ToVmfsDatastoreArrayOutputWithContext(context.Background())
}

func (i VmfsDatastoreArray) ToVmfsDatastoreArrayOutputWithContext(ctx context.Context) VmfsDatastoreArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmfsDatastoreArrayOutput)
}

// VmfsDatastoreMapInput is an input type that accepts VmfsDatastoreMap and VmfsDatastoreMapOutput values.
// You can construct a concrete instance of `VmfsDatastoreMapInput` via:
//
//          VmfsDatastoreMap{ "key": VmfsDatastoreArgs{...} }
type VmfsDatastoreMapInput interface {
	pulumi.Input

	ToVmfsDatastoreMapOutput() VmfsDatastoreMapOutput
	ToVmfsDatastoreMapOutputWithContext(context.Context) VmfsDatastoreMapOutput
}

type VmfsDatastoreMap map[string]VmfsDatastoreInput

func (VmfsDatastoreMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VmfsDatastore)(nil))
}

func (i VmfsDatastoreMap) ToVmfsDatastoreMapOutput() VmfsDatastoreMapOutput {
	return i.ToVmfsDatastoreMapOutputWithContext(context.Background())
}

func (i VmfsDatastoreMap) ToVmfsDatastoreMapOutputWithContext(ctx context.Context) VmfsDatastoreMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmfsDatastoreMapOutput)
}

type VmfsDatastoreOutput struct {
	*pulumi.OutputState
}

func (VmfsDatastoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmfsDatastore)(nil))
}

func (o VmfsDatastoreOutput) ToVmfsDatastoreOutput() VmfsDatastoreOutput {
	return o
}

func (o VmfsDatastoreOutput) ToVmfsDatastoreOutputWithContext(ctx context.Context) VmfsDatastoreOutput {
	return o
}

func (o VmfsDatastoreOutput) ToVmfsDatastorePtrOutput() VmfsDatastorePtrOutput {
	return o.ToVmfsDatastorePtrOutputWithContext(context.Background())
}

func (o VmfsDatastoreOutput) ToVmfsDatastorePtrOutputWithContext(ctx context.Context) VmfsDatastorePtrOutput {
	return o.ApplyT(func(v VmfsDatastore) *VmfsDatastore {
		return &v
	}).(VmfsDatastorePtrOutput)
}

type VmfsDatastorePtrOutput struct {
	*pulumi.OutputState
}

func (VmfsDatastorePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmfsDatastore)(nil))
}

func (o VmfsDatastorePtrOutput) ToVmfsDatastorePtrOutput() VmfsDatastorePtrOutput {
	return o
}

func (o VmfsDatastorePtrOutput) ToVmfsDatastorePtrOutputWithContext(ctx context.Context) VmfsDatastorePtrOutput {
	return o
}

type VmfsDatastoreArrayOutput struct{ *pulumi.OutputState }

func (VmfsDatastoreArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmfsDatastore)(nil))
}

func (o VmfsDatastoreArrayOutput) ToVmfsDatastoreArrayOutput() VmfsDatastoreArrayOutput {
	return o
}

func (o VmfsDatastoreArrayOutput) ToVmfsDatastoreArrayOutputWithContext(ctx context.Context) VmfsDatastoreArrayOutput {
	return o
}

func (o VmfsDatastoreArrayOutput) Index(i pulumi.IntInput) VmfsDatastoreOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VmfsDatastore {
		return vs[0].([]VmfsDatastore)[vs[1].(int)]
	}).(VmfsDatastoreOutput)
}

type VmfsDatastoreMapOutput struct{ *pulumi.OutputState }

func (VmfsDatastoreMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VmfsDatastore)(nil))
}

func (o VmfsDatastoreMapOutput) ToVmfsDatastoreMapOutput() VmfsDatastoreMapOutput {
	return o
}

func (o VmfsDatastoreMapOutput) ToVmfsDatastoreMapOutputWithContext(ctx context.Context) VmfsDatastoreMapOutput {
	return o
}

func (o VmfsDatastoreMapOutput) MapIndex(k pulumi.StringInput) VmfsDatastoreOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VmfsDatastore {
		return vs[0].(map[string]VmfsDatastore)[vs[1].(string)]
	}).(VmfsDatastoreOutput)
}

func init() {
	pulumi.RegisterOutputType(VmfsDatastoreOutput{})
	pulumi.RegisterOutputType(VmfsDatastorePtrOutput{})
	pulumi.RegisterOutputType(VmfsDatastoreArrayOutput{})
	pulumi.RegisterOutputType(VmfsDatastoreMapOutput{})
}

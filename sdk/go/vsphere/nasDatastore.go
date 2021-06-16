// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NasDatastore struct {
	pulumi.CustomResourceState

	// Access mode for the mount point. Can be one of
	// `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
	// that the datastore will be read-write depending on the permissions of the
	// actual share. Default: `readWrite`. Forces a new resource if changed.
	AccessMode pulumi.StringPtrOutput `pulumi:"accessMode"`
	// The connectivity status of the datastore. If this is `false`,
	// some other computed attributes may be out of date.
	Accessible pulumi.BoolOutput `pulumi:"accessible"`
	// Maximum capacity of the datastore, in megabytes.
	Capacity pulumi.IntOutput `pulumi:"capacity"`
	// Map of custom attribute ids to attribute
	// value strings to set on datasource resource.
	CustomAttributes pulumi.StringMapOutput `pulumi:"customAttributes"`
	// The managed object
	// ID of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId pulumi.StringPtrOutput `pulumi:"datastoreClusterId"`
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
	// The managed object IDs of
	// the hosts to mount the datastore on.
	HostSystemIds pulumi.StringArrayOutput `pulumi:"hostSystemIds"`
	// The current maintenance mode state of the datastore.
	MaintenanceMode pulumi.StringOutput `pulumi:"maintenanceMode"`
	// If `true`, more than one host in the datacenter has
	// been configured with access to the datastore.
	MultipleHostAccess pulumi.BoolOutput `pulumi:"multipleHostAccess"`
	// The name of the datastore. Forces a new resource if
	// changed.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates that this NAS volume is a protocol endpoint.
	// This field is only populated if the host supports virtual datastores.
	ProtocolEndpoint pulumi.BoolOutput `pulumi:"protocolEndpoint"`
	// The hostnames or IP addresses of the remote
	// server or servers. Only one element should be present for NFS v3 but multiple
	// can be present for NFS v4.1. Forces a new resource if changed.
	RemoteHosts pulumi.StringArrayOutput `pulumi:"remoteHosts"`
	// The remote path of the mount point. Forces a new
	// resource if changed.
	RemotePath pulumi.StringOutput `pulumi:"remotePath"`
	// The security type to use when using NFS v4.1.
	// Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
	// if changed.
	SecurityType pulumi.StringPtrOutput `pulumi:"securityType"`
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The type of NAS volume. Can be one of `NFS` (to denote
	// v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
	// changed.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// Total additional storage space, in megabytes,
	// potentially used by all virtual machines on this datastore.
	UncommittedSpace pulumi.IntOutput `pulumi:"uncommittedSpace"`
	// The unique locator for the datastore.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewNasDatastore registers a new resource with the given unique name, arguments, and options.
func NewNasDatastore(ctx *pulumi.Context,
	name string, args *NasDatastoreArgs, opts ...pulumi.ResourceOption) (*NasDatastore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostSystemIds == nil {
		return nil, errors.New("invalid value for required argument 'HostSystemIds'")
	}
	if args.RemoteHosts == nil {
		return nil, errors.New("invalid value for required argument 'RemoteHosts'")
	}
	if args.RemotePath == nil {
		return nil, errors.New("invalid value for required argument 'RemotePath'")
	}
	var resource NasDatastore
	err := ctx.RegisterResource("vsphere:index/nasDatastore:NasDatastore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNasDatastore gets an existing NasDatastore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNasDatastore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NasDatastoreState, opts ...pulumi.ResourceOption) (*NasDatastore, error) {
	var resource NasDatastore
	err := ctx.ReadResource("vsphere:index/nasDatastore:NasDatastore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NasDatastore resources.
type nasDatastoreState struct {
	// Access mode for the mount point. Can be one of
	// `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
	// that the datastore will be read-write depending on the permissions of the
	// actual share. Default: `readWrite`. Forces a new resource if changed.
	AccessMode *string `pulumi:"accessMode"`
	// The connectivity status of the datastore. If this is `false`,
	// some other computed attributes may be out of date.
	Accessible *bool `pulumi:"accessible"`
	// Maximum capacity of the datastore, in megabytes.
	Capacity *int `pulumi:"capacity"`
	// Map of custom attribute ids to attribute
	// value strings to set on datasource resource.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The managed object
	// ID of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId *string `pulumi:"datastoreClusterId"`
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
	// The managed object IDs of
	// the hosts to mount the datastore on.
	HostSystemIds []string `pulumi:"hostSystemIds"`
	// The current maintenance mode state of the datastore.
	MaintenanceMode *string `pulumi:"maintenanceMode"`
	// If `true`, more than one host in the datacenter has
	// been configured with access to the datastore.
	MultipleHostAccess *bool `pulumi:"multipleHostAccess"`
	// The name of the datastore. Forces a new resource if
	// changed.
	Name *string `pulumi:"name"`
	// Indicates that this NAS volume is a protocol endpoint.
	// This field is only populated if the host supports virtual datastores.
	ProtocolEndpoint *bool `pulumi:"protocolEndpoint"`
	// The hostnames or IP addresses of the remote
	// server or servers. Only one element should be present for NFS v3 but multiple
	// can be present for NFS v4.1. Forces a new resource if changed.
	RemoteHosts []string `pulumi:"remoteHosts"`
	// The remote path of the mount point. Forces a new
	// resource if changed.
	RemotePath *string `pulumi:"remotePath"`
	// The security type to use when using NFS v4.1.
	// Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
	// if changed.
	SecurityType *string `pulumi:"securityType"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
	// The type of NAS volume. Can be one of `NFS` (to denote
	// v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
	// changed.
	Type *string `pulumi:"type"`
	// Total additional storage space, in megabytes,
	// potentially used by all virtual machines on this datastore.
	UncommittedSpace *int `pulumi:"uncommittedSpace"`
	// The unique locator for the datastore.
	Url *string `pulumi:"url"`
}

type NasDatastoreState struct {
	// Access mode for the mount point. Can be one of
	// `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
	// that the datastore will be read-write depending on the permissions of the
	// actual share. Default: `readWrite`. Forces a new resource if changed.
	AccessMode pulumi.StringPtrInput
	// The connectivity status of the datastore. If this is `false`,
	// some other computed attributes may be out of date.
	Accessible pulumi.BoolPtrInput
	// Maximum capacity of the datastore, in megabytes.
	Capacity pulumi.IntPtrInput
	// Map of custom attribute ids to attribute
	// value strings to set on datasource resource.
	CustomAttributes pulumi.StringMapInput
	// The managed object
	// ID of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId pulumi.StringPtrInput
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
	// The managed object IDs of
	// the hosts to mount the datastore on.
	HostSystemIds pulumi.StringArrayInput
	// The current maintenance mode state of the datastore.
	MaintenanceMode pulumi.StringPtrInput
	// If `true`, more than one host in the datacenter has
	// been configured with access to the datastore.
	MultipleHostAccess pulumi.BoolPtrInput
	// The name of the datastore. Forces a new resource if
	// changed.
	Name pulumi.StringPtrInput
	// Indicates that this NAS volume is a protocol endpoint.
	// This field is only populated if the host supports virtual datastores.
	ProtocolEndpoint pulumi.BoolPtrInput
	// The hostnames or IP addresses of the remote
	// server or servers. Only one element should be present for NFS v3 but multiple
	// can be present for NFS v4.1. Forces a new resource if changed.
	RemoteHosts pulumi.StringArrayInput
	// The remote path of the mount point. Forces a new
	// resource if changed.
	RemotePath pulumi.StringPtrInput
	// The security type to use when using NFS v4.1.
	// Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
	// if changed.
	SecurityType pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
	// The type of NAS volume. Can be one of `NFS` (to denote
	// v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
	// changed.
	Type pulumi.StringPtrInput
	// Total additional storage space, in megabytes,
	// potentially used by all virtual machines on this datastore.
	UncommittedSpace pulumi.IntPtrInput
	// The unique locator for the datastore.
	Url pulumi.StringPtrInput
}

func (NasDatastoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*nasDatastoreState)(nil)).Elem()
}

type nasDatastoreArgs struct {
	// Access mode for the mount point. Can be one of
	// `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
	// that the datastore will be read-write depending on the permissions of the
	// actual share. Default: `readWrite`. Forces a new resource if changed.
	AccessMode *string `pulumi:"accessMode"`
	// Map of custom attribute ids to attribute
	// value strings to set on datasource resource.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The managed object
	// ID of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId *string `pulumi:"datastoreClusterId"`
	// The relative path to a folder to put this datastore in.
	// This is a path relative to the datacenter you are deploying the datastore to.
	// Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
	// The provider will place a datastore named `test` in a datastore folder
	// located at `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/test`. Conflicts with
	// `datastoreClusterId`.
	Folder *string `pulumi:"folder"`
	// The managed object IDs of
	// the hosts to mount the datastore on.
	HostSystemIds []string `pulumi:"hostSystemIds"`
	// The name of the datastore. Forces a new resource if
	// changed.
	Name *string `pulumi:"name"`
	// The hostnames or IP addresses of the remote
	// server or servers. Only one element should be present for NFS v3 but multiple
	// can be present for NFS v4.1. Forces a new resource if changed.
	RemoteHosts []string `pulumi:"remoteHosts"`
	// The remote path of the mount point. Forces a new
	// resource if changed.
	RemotePath string `pulumi:"remotePath"`
	// The security type to use when using NFS v4.1.
	// Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
	// if changed.
	SecurityType *string `pulumi:"securityType"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
	// The type of NAS volume. Can be one of `NFS` (to denote
	// v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
	// changed.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a NasDatastore resource.
type NasDatastoreArgs struct {
	// Access mode for the mount point. Can be one of
	// `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
	// that the datastore will be read-write depending on the permissions of the
	// actual share. Default: `readWrite`. Forces a new resource if changed.
	AccessMode pulumi.StringPtrInput
	// Map of custom attribute ids to attribute
	// value strings to set on datasource resource.
	CustomAttributes pulumi.StringMapInput
	// The managed object
	// ID of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId pulumi.StringPtrInput
	// The relative path to a folder to put this datastore in.
	// This is a path relative to the datacenter you are deploying the datastore to.
	// Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
	// The provider will place a datastore named `test` in a datastore folder
	// located at `/dc1/datastore/foo/bar`, with the final inventory path being
	// `/dc1/datastore/foo/bar/test`. Conflicts with
	// `datastoreClusterId`.
	Folder pulumi.StringPtrInput
	// The managed object IDs of
	// the hosts to mount the datastore on.
	HostSystemIds pulumi.StringArrayInput
	// The name of the datastore. Forces a new resource if
	// changed.
	Name pulumi.StringPtrInput
	// The hostnames or IP addresses of the remote
	// server or servers. Only one element should be present for NFS v3 but multiple
	// can be present for NFS v4.1. Forces a new resource if changed.
	RemoteHosts pulumi.StringArrayInput
	// The remote path of the mount point. Forces a new
	// resource if changed.
	RemotePath pulumi.StringInput
	// The security type to use when using NFS v4.1.
	// Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
	// if changed.
	SecurityType pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
	// The type of NAS volume. Can be one of `NFS` (to denote
	// v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
	// changed.
	Type pulumi.StringPtrInput
}

func (NasDatastoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nasDatastoreArgs)(nil)).Elem()
}

type NasDatastoreInput interface {
	pulumi.Input

	ToNasDatastoreOutput() NasDatastoreOutput
	ToNasDatastoreOutputWithContext(ctx context.Context) NasDatastoreOutput
}

func (*NasDatastore) ElementType() reflect.Type {
	return reflect.TypeOf((*NasDatastore)(nil))
}

func (i *NasDatastore) ToNasDatastoreOutput() NasDatastoreOutput {
	return i.ToNasDatastoreOutputWithContext(context.Background())
}

func (i *NasDatastore) ToNasDatastoreOutputWithContext(ctx context.Context) NasDatastoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NasDatastoreOutput)
}

func (i *NasDatastore) ToNasDatastorePtrOutput() NasDatastorePtrOutput {
	return i.ToNasDatastorePtrOutputWithContext(context.Background())
}

func (i *NasDatastore) ToNasDatastorePtrOutputWithContext(ctx context.Context) NasDatastorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NasDatastorePtrOutput)
}

type NasDatastorePtrInput interface {
	pulumi.Input

	ToNasDatastorePtrOutput() NasDatastorePtrOutput
	ToNasDatastorePtrOutputWithContext(ctx context.Context) NasDatastorePtrOutput
}

type nasDatastorePtrType NasDatastoreArgs

func (*nasDatastorePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NasDatastore)(nil))
}

func (i *nasDatastorePtrType) ToNasDatastorePtrOutput() NasDatastorePtrOutput {
	return i.ToNasDatastorePtrOutputWithContext(context.Background())
}

func (i *nasDatastorePtrType) ToNasDatastorePtrOutputWithContext(ctx context.Context) NasDatastorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NasDatastorePtrOutput)
}

// NasDatastoreArrayInput is an input type that accepts NasDatastoreArray and NasDatastoreArrayOutput values.
// You can construct a concrete instance of `NasDatastoreArrayInput` via:
//
//          NasDatastoreArray{ NasDatastoreArgs{...} }
type NasDatastoreArrayInput interface {
	pulumi.Input

	ToNasDatastoreArrayOutput() NasDatastoreArrayOutput
	ToNasDatastoreArrayOutputWithContext(context.Context) NasDatastoreArrayOutput
}

type NasDatastoreArray []NasDatastoreInput

func (NasDatastoreArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*NasDatastore)(nil))
}

func (i NasDatastoreArray) ToNasDatastoreArrayOutput() NasDatastoreArrayOutput {
	return i.ToNasDatastoreArrayOutputWithContext(context.Background())
}

func (i NasDatastoreArray) ToNasDatastoreArrayOutputWithContext(ctx context.Context) NasDatastoreArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NasDatastoreArrayOutput)
}

// NasDatastoreMapInput is an input type that accepts NasDatastoreMap and NasDatastoreMapOutput values.
// You can construct a concrete instance of `NasDatastoreMapInput` via:
//
//          NasDatastoreMap{ "key": NasDatastoreArgs{...} }
type NasDatastoreMapInput interface {
	pulumi.Input

	ToNasDatastoreMapOutput() NasDatastoreMapOutput
	ToNasDatastoreMapOutputWithContext(context.Context) NasDatastoreMapOutput
}

type NasDatastoreMap map[string]NasDatastoreInput

func (NasDatastoreMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*NasDatastore)(nil))
}

func (i NasDatastoreMap) ToNasDatastoreMapOutput() NasDatastoreMapOutput {
	return i.ToNasDatastoreMapOutputWithContext(context.Background())
}

func (i NasDatastoreMap) ToNasDatastoreMapOutputWithContext(ctx context.Context) NasDatastoreMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NasDatastoreMapOutput)
}

type NasDatastoreOutput struct {
	*pulumi.OutputState
}

func (NasDatastoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NasDatastore)(nil))
}

func (o NasDatastoreOutput) ToNasDatastoreOutput() NasDatastoreOutput {
	return o
}

func (o NasDatastoreOutput) ToNasDatastoreOutputWithContext(ctx context.Context) NasDatastoreOutput {
	return o
}

func (o NasDatastoreOutput) ToNasDatastorePtrOutput() NasDatastorePtrOutput {
	return o.ToNasDatastorePtrOutputWithContext(context.Background())
}

func (o NasDatastoreOutput) ToNasDatastorePtrOutputWithContext(ctx context.Context) NasDatastorePtrOutput {
	return o.ApplyT(func(v NasDatastore) *NasDatastore {
		return &v
	}).(NasDatastorePtrOutput)
}

type NasDatastorePtrOutput struct {
	*pulumi.OutputState
}

func (NasDatastorePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NasDatastore)(nil))
}

func (o NasDatastorePtrOutput) ToNasDatastorePtrOutput() NasDatastorePtrOutput {
	return o
}

func (o NasDatastorePtrOutput) ToNasDatastorePtrOutputWithContext(ctx context.Context) NasDatastorePtrOutput {
	return o
}

type NasDatastoreArrayOutput struct{ *pulumi.OutputState }

func (NasDatastoreArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NasDatastore)(nil))
}

func (o NasDatastoreArrayOutput) ToNasDatastoreArrayOutput() NasDatastoreArrayOutput {
	return o
}

func (o NasDatastoreArrayOutput) ToNasDatastoreArrayOutputWithContext(ctx context.Context) NasDatastoreArrayOutput {
	return o
}

func (o NasDatastoreArrayOutput) Index(i pulumi.IntInput) NasDatastoreOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NasDatastore {
		return vs[0].([]NasDatastore)[vs[1].(int)]
	}).(NasDatastoreOutput)
}

type NasDatastoreMapOutput struct{ *pulumi.OutputState }

func (NasDatastoreMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NasDatastore)(nil))
}

func (o NasDatastoreMapOutput) ToNasDatastoreMapOutput() NasDatastoreMapOutput {
	return o
}

func (o NasDatastoreMapOutput) ToNasDatastoreMapOutputWithContext(ctx context.Context) NasDatastoreMapOutput {
	return o
}

func (o NasDatastoreMapOutput) MapIndex(k pulumi.StringInput) NasDatastoreOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NasDatastore {
		return vs[0].(map[string]NasDatastore)[vs[1].(string)]
	}).(NasDatastoreOutput)
}

func init() {
	pulumi.RegisterOutputType(NasDatastoreOutput{})
	pulumi.RegisterOutputType(NasDatastorePtrOutput{})
	pulumi.RegisterOutputType(NasDatastoreArrayOutput{})
	pulumi.RegisterOutputType(NasDatastoreMapOutput{})
}

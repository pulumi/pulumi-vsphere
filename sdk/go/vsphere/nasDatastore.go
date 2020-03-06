// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vsphere

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `.NasDatastore` resource can be used to create and manage NAS
// datastores on an ESXi host or a set of hosts. The resource supports mounting
// NFS v3 and v4.1 shares to be used as datastores.
// 
// > **NOTE:** Unlike [`.VmfsDatastore`][resource-vmfs-datastore], a NAS
// datastore is only mounted on the hosts you choose to mount it on. To mount on
// multiple hosts, you must specify each host that you want to add in the
// `hostSystemIds` argument.
// 
// [resource-vmfs-datastore]: /docs/providers/vsphere/r/vmfs_datastore.html
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/nas_datastore.html.markdown.
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
	// value strings to set on datasource resource. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes pulumi.StringMapOutput `pulumi:"customAttributes"`
	// The [managed object
	// ID][docs-about-morefs] of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId pulumi.StringPtrOutput `pulumi:"datastoreClusterId"`
	// The path to the datastore folder to put the datastore in.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// Available space of this datastore, in megabytes.
	FreeSpace pulumi.IntOutput `pulumi:"freeSpace"`
	// The [managed object IDs][docs-about-morefs] of
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
	ProtocolEndpoint pulumi.StringOutput `pulumi:"protocolEndpoint"`
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
	// The IDs of any tags to attach to this resource. See
	// [here][docs-applying-tags] for a reference on how to apply tags.
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
	if args == nil || args.HostSystemIds == nil {
		return nil, errors.New("missing required argument 'HostSystemIds'")
	}
	if args == nil || args.RemoteHosts == nil {
		return nil, errors.New("missing required argument 'RemoteHosts'")
	}
	if args == nil || args.RemotePath == nil {
		return nil, errors.New("missing required argument 'RemotePath'")
	}
	if args == nil {
		args = &NasDatastoreArgs{}
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
	// value strings to set on datasource resource. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The [managed object
	// ID][docs-about-morefs] of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId *string `pulumi:"datastoreClusterId"`
	// The path to the datastore folder to put the datastore in.
	Folder *string `pulumi:"folder"`
	// Available space of this datastore, in megabytes.
	FreeSpace *int `pulumi:"freeSpace"`
	// The [managed object IDs][docs-about-morefs] of
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
	ProtocolEndpoint *string `pulumi:"protocolEndpoint"`
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
	// The IDs of any tags to attach to this resource. See
	// [here][docs-applying-tags] for a reference on how to apply tags.
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
	// value strings to set on datasource resource. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes pulumi.StringMapInput
	// The [managed object
	// ID][docs-about-morefs] of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId pulumi.StringPtrInput
	// The path to the datastore folder to put the datastore in.
	Folder pulumi.StringPtrInput
	// Available space of this datastore, in megabytes.
	FreeSpace pulumi.IntPtrInput
	// The [managed object IDs][docs-about-morefs] of
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
	ProtocolEndpoint pulumi.StringPtrInput
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
	// The IDs of any tags to attach to this resource. See
	// [here][docs-applying-tags] for a reference on how to apply tags.
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
	// value strings to set on datasource resource. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The [managed object
	// ID][docs-about-morefs] of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId *string `pulumi:"datastoreClusterId"`
	// The path to the datastore folder to put the datastore in.
	Folder *string `pulumi:"folder"`
	// The [managed object IDs][docs-about-morefs] of
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
	// The IDs of any tags to attach to this resource. See
	// [here][docs-applying-tags] for a reference on how to apply tags.
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
	// value strings to set on datasource resource. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes pulumi.StringMapInput
	// The [managed object
	// ID][docs-about-morefs] of a datastore cluster to put this datastore in.
	// Conflicts with `folder`.
	DatastoreClusterId pulumi.StringPtrInput
	// The path to the datastore folder to put the datastore in.
	Folder pulumi.StringPtrInput
	// The [managed object IDs][docs-about-morefs] of
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
	// The IDs of any tags to attach to this resource. See
	// [here][docs-applying-tags] for a reference on how to apply tags.
	Tags pulumi.StringArrayInput
	// The type of NAS volume. Can be one of `NFS` (to denote
	// v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
	// changed.
	Type pulumi.StringPtrInput
}

func (NasDatastoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nasDatastoreArgs)(nil)).Elem()
}


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

type Host struct {
	pulumi.CustomResourceState

	// The ID of the Compute Cluster this host should
	// be added to. This should not be set if `datacenter` is set. Conflicts with:
	// `clusterManaged`.
	Cluster pulumi.StringPtrOutput `pulumi:"cluster"`
	// Can be set to `true` if compute cluster
	// membership will be managed through the `computeCluster` resource rather
	// than the`host` resource. Conflicts with: `cluster`.
	ClusterManaged pulumi.BoolPtrOutput `pulumi:"clusterManaged"`
	// If set to false then the host will be disconnected.
	// Default is `false`.
	Connected pulumi.BoolPtrOutput `pulumi:"connected"`
	// A map of custom attribute IDs and string
	// values to apply to the resource. Please refer to the
	// `vsphereCustomAttributes` resource for more information on applying
	// tags to resources.
	//
	// > **NOTE:** Custom attributes are not supported on direct ESXi host
	// connections and require vCenter Server.
	//
	// [docs-host-thumbprint-data-source]: /docs/providers/vsphere/d/host_thumbprint.html
	CustomAttributes pulumi.StringMapOutput `pulumi:"customAttributes"`
	// The ID of the datacenter this host should
	// be added to. This should not be set if `cluster` is set.
	Datacenter pulumi.StringPtrOutput `pulumi:"datacenter"`
	// If set to `true` then it will force the host to be added,
	// even if the host is already connected to a different vCenter Server instance.
	// Default is `false`.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// FQDN or IP address of the host to be added.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The license key that will be applied to the host.
	// The license key is expected to be present in vSphere.
	License pulumi.StringPtrOutput `pulumi:"license"`
	// Set the lockdown state of the host. Valid options are
	// `disabled`, `normal`, and `strict`. Default is `disabled`.
	Lockdown pulumi.StringPtrOutput `pulumi:"lockdown"`
	// Set the management state of the host.
	// Default is `false`.
	Maintenance pulumi.BoolPtrOutput `pulumi:"maintenance"`
	// Password that will be used by vSphere to authenticate
	// to the host.
	Password pulumi.StringOutput `pulumi:"password"`
	// Set Services on host, the settings to be set are based on service being set as part of import.
	Services HostServiceArrayOutput `pulumi:"services"`
	// The IDs of any tags to attach to this resource. Please
	// refer to the `Tag` resource for more information on applying
	// tags to resources.
	//
	// > **NOTE:** Tagging support is not supported on direct ESXi host
	// connections and require vCenter Server.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Host's certificate SHA-1 thumbprint. If not set the
	// CA that signed the host's certificate should be trusted. If the CA is not
	// trusted and no thumbprint is set then the operation will fail. See data source
	// [`getHostThumbprint`][docs-host-thumbprint-data-source].
	Thumbprint pulumi.StringPtrOutput `pulumi:"thumbprint"`
	// Username that will be used by vSphere to authenticate
	// to the host.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewHost registers a new resource with the given unique name, arguments, and options.
func NewHost(ctx *pulumi.Context,
	name string, args *HostArgs, opts ...pulumi.ResourceOption) (*Host, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Hostname == nil {
		return nil, errors.New("invalid value for required argument 'Hostname'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Host
	err := ctx.RegisterResource("vsphere:index/host:Host", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHost gets an existing Host resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostState, opts ...pulumi.ResourceOption) (*Host, error) {
	var resource Host
	err := ctx.ReadResource("vsphere:index/host:Host", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Host resources.
type hostState struct {
	// The ID of the Compute Cluster this host should
	// be added to. This should not be set if `datacenter` is set. Conflicts with:
	// `clusterManaged`.
	Cluster *string `pulumi:"cluster"`
	// Can be set to `true` if compute cluster
	// membership will be managed through the `computeCluster` resource rather
	// than the`host` resource. Conflicts with: `cluster`.
	ClusterManaged *bool `pulumi:"clusterManaged"`
	// If set to false then the host will be disconnected.
	// Default is `false`.
	Connected *bool `pulumi:"connected"`
	// A map of custom attribute IDs and string
	// values to apply to the resource. Please refer to the
	// `vsphereCustomAttributes` resource for more information on applying
	// tags to resources.
	//
	// > **NOTE:** Custom attributes are not supported on direct ESXi host
	// connections and require vCenter Server.
	//
	// [docs-host-thumbprint-data-source]: /docs/providers/vsphere/d/host_thumbprint.html
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The ID of the datacenter this host should
	// be added to. This should not be set if `cluster` is set.
	Datacenter *string `pulumi:"datacenter"`
	// If set to `true` then it will force the host to be added,
	// even if the host is already connected to a different vCenter Server instance.
	// Default is `false`.
	Force *bool `pulumi:"force"`
	// FQDN or IP address of the host to be added.
	Hostname *string `pulumi:"hostname"`
	// The license key that will be applied to the host.
	// The license key is expected to be present in vSphere.
	License *string `pulumi:"license"`
	// Set the lockdown state of the host. Valid options are
	// `disabled`, `normal`, and `strict`. Default is `disabled`.
	Lockdown *string `pulumi:"lockdown"`
	// Set the management state of the host.
	// Default is `false`.
	Maintenance *bool `pulumi:"maintenance"`
	// Password that will be used by vSphere to authenticate
	// to the host.
	Password *string `pulumi:"password"`
	// Set Services on host, the settings to be set are based on service being set as part of import.
	Services []HostService `pulumi:"services"`
	// The IDs of any tags to attach to this resource. Please
	// refer to the `Tag` resource for more information on applying
	// tags to resources.
	//
	// > **NOTE:** Tagging support is not supported on direct ESXi host
	// connections and require vCenter Server.
	Tags []string `pulumi:"tags"`
	// Host's certificate SHA-1 thumbprint. If not set the
	// CA that signed the host's certificate should be trusted. If the CA is not
	// trusted and no thumbprint is set then the operation will fail. See data source
	// [`getHostThumbprint`][docs-host-thumbprint-data-source].
	Thumbprint *string `pulumi:"thumbprint"`
	// Username that will be used by vSphere to authenticate
	// to the host.
	Username *string `pulumi:"username"`
}

type HostState struct {
	// The ID of the Compute Cluster this host should
	// be added to. This should not be set if `datacenter` is set. Conflicts with:
	// `clusterManaged`.
	Cluster pulumi.StringPtrInput
	// Can be set to `true` if compute cluster
	// membership will be managed through the `computeCluster` resource rather
	// than the`host` resource. Conflicts with: `cluster`.
	ClusterManaged pulumi.BoolPtrInput
	// If set to false then the host will be disconnected.
	// Default is `false`.
	Connected pulumi.BoolPtrInput
	// A map of custom attribute IDs and string
	// values to apply to the resource. Please refer to the
	// `vsphereCustomAttributes` resource for more information on applying
	// tags to resources.
	//
	// > **NOTE:** Custom attributes are not supported on direct ESXi host
	// connections and require vCenter Server.
	//
	// [docs-host-thumbprint-data-source]: /docs/providers/vsphere/d/host_thumbprint.html
	CustomAttributes pulumi.StringMapInput
	// The ID of the datacenter this host should
	// be added to. This should not be set if `cluster` is set.
	Datacenter pulumi.StringPtrInput
	// If set to `true` then it will force the host to be added,
	// even if the host is already connected to a different vCenter Server instance.
	// Default is `false`.
	Force pulumi.BoolPtrInput
	// FQDN or IP address of the host to be added.
	Hostname pulumi.StringPtrInput
	// The license key that will be applied to the host.
	// The license key is expected to be present in vSphere.
	License pulumi.StringPtrInput
	// Set the lockdown state of the host. Valid options are
	// `disabled`, `normal`, and `strict`. Default is `disabled`.
	Lockdown pulumi.StringPtrInput
	// Set the management state of the host.
	// Default is `false`.
	Maintenance pulumi.BoolPtrInput
	// Password that will be used by vSphere to authenticate
	// to the host.
	Password pulumi.StringPtrInput
	// Set Services on host, the settings to be set are based on service being set as part of import.
	Services HostServiceArrayInput
	// The IDs of any tags to attach to this resource. Please
	// refer to the `Tag` resource for more information on applying
	// tags to resources.
	//
	// > **NOTE:** Tagging support is not supported on direct ESXi host
	// connections and require vCenter Server.
	Tags pulumi.StringArrayInput
	// Host's certificate SHA-1 thumbprint. If not set the
	// CA that signed the host's certificate should be trusted. If the CA is not
	// trusted and no thumbprint is set then the operation will fail. See data source
	// [`getHostThumbprint`][docs-host-thumbprint-data-source].
	Thumbprint pulumi.StringPtrInput
	// Username that will be used by vSphere to authenticate
	// to the host.
	Username pulumi.StringPtrInput
}

func (HostState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostState)(nil)).Elem()
}

type hostArgs struct {
	// The ID of the Compute Cluster this host should
	// be added to. This should not be set if `datacenter` is set. Conflicts with:
	// `clusterManaged`.
	Cluster *string `pulumi:"cluster"`
	// Can be set to `true` if compute cluster
	// membership will be managed through the `computeCluster` resource rather
	// than the`host` resource. Conflicts with: `cluster`.
	ClusterManaged *bool `pulumi:"clusterManaged"`
	// If set to false then the host will be disconnected.
	// Default is `false`.
	Connected *bool `pulumi:"connected"`
	// A map of custom attribute IDs and string
	// values to apply to the resource. Please refer to the
	// `vsphereCustomAttributes` resource for more information on applying
	// tags to resources.
	//
	// > **NOTE:** Custom attributes are not supported on direct ESXi host
	// connections and require vCenter Server.
	//
	// [docs-host-thumbprint-data-source]: /docs/providers/vsphere/d/host_thumbprint.html
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The ID of the datacenter this host should
	// be added to. This should not be set if `cluster` is set.
	Datacenter *string `pulumi:"datacenter"`
	// If set to `true` then it will force the host to be added,
	// even if the host is already connected to a different vCenter Server instance.
	// Default is `false`.
	Force *bool `pulumi:"force"`
	// FQDN or IP address of the host to be added.
	Hostname string `pulumi:"hostname"`
	// The license key that will be applied to the host.
	// The license key is expected to be present in vSphere.
	License *string `pulumi:"license"`
	// Set the lockdown state of the host. Valid options are
	// `disabled`, `normal`, and `strict`. Default is `disabled`.
	Lockdown *string `pulumi:"lockdown"`
	// Set the management state of the host.
	// Default is `false`.
	Maintenance *bool `pulumi:"maintenance"`
	// Password that will be used by vSphere to authenticate
	// to the host.
	Password string `pulumi:"password"`
	// Set Services on host, the settings to be set are based on service being set as part of import.
	Services []HostService `pulumi:"services"`
	// The IDs of any tags to attach to this resource. Please
	// refer to the `Tag` resource for more information on applying
	// tags to resources.
	//
	// > **NOTE:** Tagging support is not supported on direct ESXi host
	// connections and require vCenter Server.
	Tags []string `pulumi:"tags"`
	// Host's certificate SHA-1 thumbprint. If not set the
	// CA that signed the host's certificate should be trusted. If the CA is not
	// trusted and no thumbprint is set then the operation will fail. See data source
	// [`getHostThumbprint`][docs-host-thumbprint-data-source].
	Thumbprint *string `pulumi:"thumbprint"`
	// Username that will be used by vSphere to authenticate
	// to the host.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a Host resource.
type HostArgs struct {
	// The ID of the Compute Cluster this host should
	// be added to. This should not be set if `datacenter` is set. Conflicts with:
	// `clusterManaged`.
	Cluster pulumi.StringPtrInput
	// Can be set to `true` if compute cluster
	// membership will be managed through the `computeCluster` resource rather
	// than the`host` resource. Conflicts with: `cluster`.
	ClusterManaged pulumi.BoolPtrInput
	// If set to false then the host will be disconnected.
	// Default is `false`.
	Connected pulumi.BoolPtrInput
	// A map of custom attribute IDs and string
	// values to apply to the resource. Please refer to the
	// `vsphereCustomAttributes` resource for more information on applying
	// tags to resources.
	//
	// > **NOTE:** Custom attributes are not supported on direct ESXi host
	// connections and require vCenter Server.
	//
	// [docs-host-thumbprint-data-source]: /docs/providers/vsphere/d/host_thumbprint.html
	CustomAttributes pulumi.StringMapInput
	// The ID of the datacenter this host should
	// be added to. This should not be set if `cluster` is set.
	Datacenter pulumi.StringPtrInput
	// If set to `true` then it will force the host to be added,
	// even if the host is already connected to a different vCenter Server instance.
	// Default is `false`.
	Force pulumi.BoolPtrInput
	// FQDN or IP address of the host to be added.
	Hostname pulumi.StringInput
	// The license key that will be applied to the host.
	// The license key is expected to be present in vSphere.
	License pulumi.StringPtrInput
	// Set the lockdown state of the host. Valid options are
	// `disabled`, `normal`, and `strict`. Default is `disabled`.
	Lockdown pulumi.StringPtrInput
	// Set the management state of the host.
	// Default is `false`.
	Maintenance pulumi.BoolPtrInput
	// Password that will be used by vSphere to authenticate
	// to the host.
	Password pulumi.StringInput
	// Set Services on host, the settings to be set are based on service being set as part of import.
	Services HostServiceArrayInput
	// The IDs of any tags to attach to this resource. Please
	// refer to the `Tag` resource for more information on applying
	// tags to resources.
	//
	// > **NOTE:** Tagging support is not supported on direct ESXi host
	// connections and require vCenter Server.
	Tags pulumi.StringArrayInput
	// Host's certificate SHA-1 thumbprint. If not set the
	// CA that signed the host's certificate should be trusted. If the CA is not
	// trusted and no thumbprint is set then the operation will fail. See data source
	// [`getHostThumbprint`][docs-host-thumbprint-data-source].
	Thumbprint pulumi.StringPtrInput
	// Username that will be used by vSphere to authenticate
	// to the host.
	Username pulumi.StringInput
}

func (HostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostArgs)(nil)).Elem()
}

type HostInput interface {
	pulumi.Input

	ToHostOutput() HostOutput
	ToHostOutputWithContext(ctx context.Context) HostOutput
}

func (*Host) ElementType() reflect.Type {
	return reflect.TypeOf((**Host)(nil)).Elem()
}

func (i *Host) ToHostOutput() HostOutput {
	return i.ToHostOutputWithContext(context.Background())
}

func (i *Host) ToHostOutputWithContext(ctx context.Context) HostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostOutput)
}

// HostArrayInput is an input type that accepts HostArray and HostArrayOutput values.
// You can construct a concrete instance of `HostArrayInput` via:
//
//	HostArray{ HostArgs{...} }
type HostArrayInput interface {
	pulumi.Input

	ToHostArrayOutput() HostArrayOutput
	ToHostArrayOutputWithContext(context.Context) HostArrayOutput
}

type HostArray []HostInput

func (HostArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Host)(nil)).Elem()
}

func (i HostArray) ToHostArrayOutput() HostArrayOutput {
	return i.ToHostArrayOutputWithContext(context.Background())
}

func (i HostArray) ToHostArrayOutputWithContext(ctx context.Context) HostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostArrayOutput)
}

// HostMapInput is an input type that accepts HostMap and HostMapOutput values.
// You can construct a concrete instance of `HostMapInput` via:
//
//	HostMap{ "key": HostArgs{...} }
type HostMapInput interface {
	pulumi.Input

	ToHostMapOutput() HostMapOutput
	ToHostMapOutputWithContext(context.Context) HostMapOutput
}

type HostMap map[string]HostInput

func (HostMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Host)(nil)).Elem()
}

func (i HostMap) ToHostMapOutput() HostMapOutput {
	return i.ToHostMapOutputWithContext(context.Background())
}

func (i HostMap) ToHostMapOutputWithContext(ctx context.Context) HostMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostMapOutput)
}

type HostOutput struct{ *pulumi.OutputState }

func (HostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Host)(nil)).Elem()
}

func (o HostOutput) ToHostOutput() HostOutput {
	return o
}

func (o HostOutput) ToHostOutputWithContext(ctx context.Context) HostOutput {
	return o
}

// The ID of the Compute Cluster this host should
// be added to. This should not be set if `datacenter` is set. Conflicts with:
// `clusterManaged`.
func (o HostOutput) Cluster() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.Cluster }).(pulumi.StringPtrOutput)
}

// Can be set to `true` if compute cluster
// membership will be managed through the `computeCluster` resource rather
// than the`host` resource. Conflicts with: `cluster`.
func (o HostOutput) ClusterManaged() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.BoolPtrOutput { return v.ClusterManaged }).(pulumi.BoolPtrOutput)
}

// If set to false then the host will be disconnected.
// Default is `false`.
func (o HostOutput) Connected() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.BoolPtrOutput { return v.Connected }).(pulumi.BoolPtrOutput)
}

// A map of custom attribute IDs and string
// values to apply to the resource. Please refer to the
// `vsphereCustomAttributes` resource for more information on applying
// tags to resources.
//
// > **NOTE:** Custom attributes are not supported on direct ESXi host
// connections and require vCenter Server.
//
// [docs-host-thumbprint-data-source]: /docs/providers/vsphere/d/host_thumbprint.html
func (o HostOutput) CustomAttributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Host) pulumi.StringMapOutput { return v.CustomAttributes }).(pulumi.StringMapOutput)
}

// The ID of the datacenter this host should
// be added to. This should not be set if `cluster` is set.
func (o HostOutput) Datacenter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.Datacenter }).(pulumi.StringPtrOutput)
}

// If set to `true` then it will force the host to be added,
// even if the host is already connected to a different vCenter Server instance.
// Default is `false`.
func (o HostOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.BoolPtrOutput { return v.Force }).(pulumi.BoolPtrOutput)
}

// FQDN or IP address of the host to be added.
func (o HostOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// The license key that will be applied to the host.
// The license key is expected to be present in vSphere.
func (o HostOutput) License() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.License }).(pulumi.StringPtrOutput)
}

// Set the lockdown state of the host. Valid options are
// `disabled`, `normal`, and `strict`. Default is `disabled`.
func (o HostOutput) Lockdown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.Lockdown }).(pulumi.StringPtrOutput)
}

// Set the management state of the host.
// Default is `false`.
func (o HostOutput) Maintenance() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.BoolPtrOutput { return v.Maintenance }).(pulumi.BoolPtrOutput)
}

// Password that will be used by vSphere to authenticate
// to the host.
func (o HostOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Set Services on host, the settings to be set are based on service being set as part of import.
func (o HostOutput) Services() HostServiceArrayOutput {
	return o.ApplyT(func(v *Host) HostServiceArrayOutput { return v.Services }).(HostServiceArrayOutput)
}

// The IDs of any tags to attach to this resource. Please
// refer to the `Tag` resource for more information on applying
// tags to resources.
//
// > **NOTE:** Tagging support is not supported on direct ESXi host
// connections and require vCenter Server.
func (o HostOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Host) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Host's certificate SHA-1 thumbprint. If not set the
// CA that signed the host's certificate should be trusted. If the CA is not
// trusted and no thumbprint is set then the operation will fail. See data source
// [`getHostThumbprint`][docs-host-thumbprint-data-source].
func (o HostOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Host) pulumi.StringPtrOutput { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

// Username that will be used by vSphere to authenticate
// to the host.
func (o HostOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type HostArrayOutput struct{ *pulumi.OutputState }

func (HostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Host)(nil)).Elem()
}

func (o HostArrayOutput) ToHostArrayOutput() HostArrayOutput {
	return o
}

func (o HostArrayOutput) ToHostArrayOutputWithContext(ctx context.Context) HostArrayOutput {
	return o
}

func (o HostArrayOutput) Index(i pulumi.IntInput) HostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Host {
		return vs[0].([]*Host)[vs[1].(int)]
	}).(HostOutput)
}

type HostMapOutput struct{ *pulumi.OutputState }

func (HostMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Host)(nil)).Elem()
}

func (o HostMapOutput) ToHostMapOutput() HostMapOutput {
	return o
}

func (o HostMapOutput) ToHostMapOutputWithContext(ctx context.Context) HostMapOutput {
	return o
}

func (o HostMapOutput) MapIndex(k pulumi.StringInput) HostOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Host {
		return vs[0].(map[string]*Host)[vs[1].(string)]
	}).(HostOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostInput)(nil)).Elem(), &Host{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostArrayInput)(nil)).Elem(), HostArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostMapInput)(nil)).Elem(), HostMap{})
	pulumi.RegisterOutputType(HostOutput{})
	pulumi.RegisterOutputType(HostArrayOutput{})
	pulumi.RegisterOutputType(HostMapOutput{})
}

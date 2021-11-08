// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VMware vSphere host resource. This represents an ESXi host that
// can be used either as part of a Compute Cluster or Standalone.
//
// ## Example Usage
// ### Create a standalone host
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
// 		opt0 := "my-datacenter"
// 		dc, err := vsphere.LookupDatacenter(ctx, &GetDatacenterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vsphere.NewHost(ctx, "h1", &vsphere.HostArgs{
// 			Hostname:   pulumi.String("10.10.10.1"),
// 			Username:   pulumi.String("root"),
// 			Password:   pulumi.String("password"),
// 			License:    pulumi.String("00000-00000-00000-00000i-00000"),
// 			Datacenter: pulumi.String(dc.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Create host in a compute cluster
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
// 		opt0 := "TfDatacenter"
// 		dc, err := vsphere.LookupDatacenter(ctx, &GetDatacenterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := dc.Id
// 		c1, err := vsphere.LookupComputeCluster(ctx, &GetComputeClusterArgs{
// 			Name:         "DC0_C0",
// 			DatacenterId: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vsphere.NewHost(ctx, "h1", &vsphere.HostArgs{
// 			Hostname: pulumi.String("10.10.10.1"),
// 			Username: pulumi.String("root"),
// 			Password: pulumi.String("password"),
// 			License:  pulumi.String("00000-00000-00000-00000i-00000"),
// 			Cluster:  pulumi.String(c1.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Importing
//
// An existing host can be [imported][docs-import] into this resource
// via supplying the host's ID. An example is below:
//
// [docs-import]: /docs/import/index.html
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
//
// The above would import the host with ID `host-123`.
type Host struct {
	pulumi.CustomResourceState

	// The ID of the Compute Cluster this host should
	// be added to. This should not be set if `datacenter` is set. Conflicts with:
	// `cluster`.
	Cluster pulumi.StringPtrOutput `pulumi:"cluster"`
	// Can be set to `true` if compute cluster
	// membership will be managed through the `computeCluster` resource rather
	// than the`host` resource. Conflicts with: `cluster`.
	ClusterManaged pulumi.BoolPtrOutput `pulumi:"clusterManaged"`
	// If set to false then the host will be disconected.
	// Default is `false`.
	Connected pulumi.BoolPtrOutput `pulumi:"connected"`
	// The ID of the datacenter this host should
	// be added to. This should not be set if `cluster` is set.
	Datacenter pulumi.StringPtrOutput `pulumi:"datacenter"`
	// If set to true then it will force the host to be added, even
	// if the host is already connected to a different vSphere instance. Default is `false`
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// FQDN or IP address of the host to be added.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The license key that will be applied to the host.
	// The license key is expected to be present in vSphere.
	License pulumi.StringPtrOutput `pulumi:"license"`
	// Set the lockdown state of the host. Valid options are
	// `disabled`, `normal`, and `strict`. Default is `disabled`.
	Lockdown pulumi.StringPtrOutput `pulumi:"lockdown"`
	// Set the management state of the host. Default is `false`.
	Maintenance pulumi.BoolPtrOutput `pulumi:"maintenance"`
	// Password that will be used by vSphere to authenticate
	// to the host.
	Password pulumi.StringOutput `pulumi:"password"`
	// Host's certificate SHA-1 thumbprint. If not set the the
	// CA that signed the host's certificate should be trusted. If the CA is not trusted
	// and no thumbprint is set then the operation will fail.
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
	// `cluster`.
	Cluster *string `pulumi:"cluster"`
	// Can be set to `true` if compute cluster
	// membership will be managed through the `computeCluster` resource rather
	// than the`host` resource. Conflicts with: `cluster`.
	ClusterManaged *bool `pulumi:"clusterManaged"`
	// If set to false then the host will be disconected.
	// Default is `false`.
	Connected *bool `pulumi:"connected"`
	// The ID of the datacenter this host should
	// be added to. This should not be set if `cluster` is set.
	Datacenter *string `pulumi:"datacenter"`
	// If set to true then it will force the host to be added, even
	// if the host is already connected to a different vSphere instance. Default is `false`
	Force *bool `pulumi:"force"`
	// FQDN or IP address of the host to be added.
	Hostname *string `pulumi:"hostname"`
	// The license key that will be applied to the host.
	// The license key is expected to be present in vSphere.
	License *string `pulumi:"license"`
	// Set the lockdown state of the host. Valid options are
	// `disabled`, `normal`, and `strict`. Default is `disabled`.
	Lockdown *string `pulumi:"lockdown"`
	// Set the management state of the host. Default is `false`.
	Maintenance *bool `pulumi:"maintenance"`
	// Password that will be used by vSphere to authenticate
	// to the host.
	Password *string `pulumi:"password"`
	// Host's certificate SHA-1 thumbprint. If not set the the
	// CA that signed the host's certificate should be trusted. If the CA is not trusted
	// and no thumbprint is set then the operation will fail.
	Thumbprint *string `pulumi:"thumbprint"`
	// Username that will be used by vSphere to authenticate
	// to the host.
	Username *string `pulumi:"username"`
}

type HostState struct {
	// The ID of the Compute Cluster this host should
	// be added to. This should not be set if `datacenter` is set. Conflicts with:
	// `cluster`.
	Cluster pulumi.StringPtrInput
	// Can be set to `true` if compute cluster
	// membership will be managed through the `computeCluster` resource rather
	// than the`host` resource. Conflicts with: `cluster`.
	ClusterManaged pulumi.BoolPtrInput
	// If set to false then the host will be disconected.
	// Default is `false`.
	Connected pulumi.BoolPtrInput
	// The ID of the datacenter this host should
	// be added to. This should not be set if `cluster` is set.
	Datacenter pulumi.StringPtrInput
	// If set to true then it will force the host to be added, even
	// if the host is already connected to a different vSphere instance. Default is `false`
	Force pulumi.BoolPtrInput
	// FQDN or IP address of the host to be added.
	Hostname pulumi.StringPtrInput
	// The license key that will be applied to the host.
	// The license key is expected to be present in vSphere.
	License pulumi.StringPtrInput
	// Set the lockdown state of the host. Valid options are
	// `disabled`, `normal`, and `strict`. Default is `disabled`.
	Lockdown pulumi.StringPtrInput
	// Set the management state of the host. Default is `false`.
	Maintenance pulumi.BoolPtrInput
	// Password that will be used by vSphere to authenticate
	// to the host.
	Password pulumi.StringPtrInput
	// Host's certificate SHA-1 thumbprint. If not set the the
	// CA that signed the host's certificate should be trusted. If the CA is not trusted
	// and no thumbprint is set then the operation will fail.
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
	// `cluster`.
	Cluster *string `pulumi:"cluster"`
	// Can be set to `true` if compute cluster
	// membership will be managed through the `computeCluster` resource rather
	// than the`host` resource. Conflicts with: `cluster`.
	ClusterManaged *bool `pulumi:"clusterManaged"`
	// If set to false then the host will be disconected.
	// Default is `false`.
	Connected *bool `pulumi:"connected"`
	// The ID of the datacenter this host should
	// be added to. This should not be set if `cluster` is set.
	Datacenter *string `pulumi:"datacenter"`
	// If set to true then it will force the host to be added, even
	// if the host is already connected to a different vSphere instance. Default is `false`
	Force *bool `pulumi:"force"`
	// FQDN or IP address of the host to be added.
	Hostname string `pulumi:"hostname"`
	// The license key that will be applied to the host.
	// The license key is expected to be present in vSphere.
	License *string `pulumi:"license"`
	// Set the lockdown state of the host. Valid options are
	// `disabled`, `normal`, and `strict`. Default is `disabled`.
	Lockdown *string `pulumi:"lockdown"`
	// Set the management state of the host. Default is `false`.
	Maintenance *bool `pulumi:"maintenance"`
	// Password that will be used by vSphere to authenticate
	// to the host.
	Password string `pulumi:"password"`
	// Host's certificate SHA-1 thumbprint. If not set the the
	// CA that signed the host's certificate should be trusted. If the CA is not trusted
	// and no thumbprint is set then the operation will fail.
	Thumbprint *string `pulumi:"thumbprint"`
	// Username that will be used by vSphere to authenticate
	// to the host.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a Host resource.
type HostArgs struct {
	// The ID of the Compute Cluster this host should
	// be added to. This should not be set if `datacenter` is set. Conflicts with:
	// `cluster`.
	Cluster pulumi.StringPtrInput
	// Can be set to `true` if compute cluster
	// membership will be managed through the `computeCluster` resource rather
	// than the`host` resource. Conflicts with: `cluster`.
	ClusterManaged pulumi.BoolPtrInput
	// If set to false then the host will be disconected.
	// Default is `false`.
	Connected pulumi.BoolPtrInput
	// The ID of the datacenter this host should
	// be added to. This should not be set if `cluster` is set.
	Datacenter pulumi.StringPtrInput
	// If set to true then it will force the host to be added, even
	// if the host is already connected to a different vSphere instance. Default is `false`
	Force pulumi.BoolPtrInput
	// FQDN or IP address of the host to be added.
	Hostname pulumi.StringInput
	// The license key that will be applied to the host.
	// The license key is expected to be present in vSphere.
	License pulumi.StringPtrInput
	// Set the lockdown state of the host. Valid options are
	// `disabled`, `normal`, and `strict`. Default is `disabled`.
	Lockdown pulumi.StringPtrInput
	// Set the management state of the host. Default is `false`.
	Maintenance pulumi.BoolPtrInput
	// Password that will be used by vSphere to authenticate
	// to the host.
	Password pulumi.StringInput
	// Host's certificate SHA-1 thumbprint. If not set the the
	// CA that signed the host's certificate should be trusted. If the CA is not trusted
	// and no thumbprint is set then the operation will fail.
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
	return reflect.TypeOf((*Host)(nil))
}

func (i *Host) ToHostOutput() HostOutput {
	return i.ToHostOutputWithContext(context.Background())
}

func (i *Host) ToHostOutputWithContext(ctx context.Context) HostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostOutput)
}

func (i *Host) ToHostPtrOutput() HostPtrOutput {
	return i.ToHostPtrOutputWithContext(context.Background())
}

func (i *Host) ToHostPtrOutputWithContext(ctx context.Context) HostPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostPtrOutput)
}

type HostPtrInput interface {
	pulumi.Input

	ToHostPtrOutput() HostPtrOutput
	ToHostPtrOutputWithContext(ctx context.Context) HostPtrOutput
}

type hostPtrType HostArgs

func (*hostPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Host)(nil))
}

func (i *hostPtrType) ToHostPtrOutput() HostPtrOutput {
	return i.ToHostPtrOutputWithContext(context.Background())
}

func (i *hostPtrType) ToHostPtrOutputWithContext(ctx context.Context) HostPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostPtrOutput)
}

// HostArrayInput is an input type that accepts HostArray and HostArrayOutput values.
// You can construct a concrete instance of `HostArrayInput` via:
//
//          HostArray{ HostArgs{...} }
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
//          HostMap{ "key": HostArgs{...} }
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
	return reflect.TypeOf((*Host)(nil))
}

func (o HostOutput) ToHostOutput() HostOutput {
	return o
}

func (o HostOutput) ToHostOutputWithContext(ctx context.Context) HostOutput {
	return o
}

func (o HostOutput) ToHostPtrOutput() HostPtrOutput {
	return o.ToHostPtrOutputWithContext(context.Background())
}

func (o HostOutput) ToHostPtrOutputWithContext(ctx context.Context) HostPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Host) *Host {
		return &v
	}).(HostPtrOutput)
}

type HostPtrOutput struct{ *pulumi.OutputState }

func (HostPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Host)(nil))
}

func (o HostPtrOutput) ToHostPtrOutput() HostPtrOutput {
	return o
}

func (o HostPtrOutput) ToHostPtrOutputWithContext(ctx context.Context) HostPtrOutput {
	return o
}

func (o HostPtrOutput) Elem() HostOutput {
	return o.ApplyT(func(v *Host) Host {
		if v != nil {
			return *v
		}
		var ret Host
		return ret
	}).(HostOutput)
}

type HostArrayOutput struct{ *pulumi.OutputState }

func (HostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Host)(nil))
}

func (o HostArrayOutput) ToHostArrayOutput() HostArrayOutput {
	return o
}

func (o HostArrayOutput) ToHostArrayOutputWithContext(ctx context.Context) HostArrayOutput {
	return o
}

func (o HostArrayOutput) Index(i pulumi.IntInput) HostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Host {
		return vs[0].([]Host)[vs[1].(int)]
	}).(HostOutput)
}

type HostMapOutput struct{ *pulumi.OutputState }

func (HostMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Host)(nil))
}

func (o HostMapOutput) ToHostMapOutput() HostMapOutput {
	return o
}

func (o HostMapOutput) ToHostMapOutputWithContext(ctx context.Context) HostMapOutput {
	return o
}

func (o HostMapOutput) MapIndex(k pulumi.StringInput) HostOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Host {
		return vs[0].(map[string]Host)[vs[1].(string)]
	}).(HostOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostInput)(nil)).Elem(), &Host{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostPtrInput)(nil)).Elem(), &Host{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostArrayInput)(nil)).Elem(), HostArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostMapInput)(nil)).Elem(), HostMap{})
	pulumi.RegisterOutputType(HostOutput{})
	pulumi.RegisterOutputType(HostPtrOutput{})
	pulumi.RegisterOutputType(HostArrayOutput{})
	pulumi.RegisterOutputType(HostMapOutput{})
}

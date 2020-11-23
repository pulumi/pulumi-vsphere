// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a VMware vSphere vnic resource.
//
// ## Example Usage
//
// ### S
// ### Create a vnic attached to a distributed virtual switch using the vmotion TCP/IP stack
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vsphere/sdk/v2/go/vsphere"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "mydc"
// 		dc, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := "esxi1.host.test"
// 		h1, err := vsphere.LookupHost(ctx, &vsphere.LookupHostArgs{
// 			Name:         &opt1,
// 			DatacenterId: dc.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		d1, err := vsphere.NewDistributedVirtualSwitch(ctx, "d1", &vsphere.DistributedVirtualSwitchArgs{
// 			DatacenterId: pulumi.String(dc.Id),
// 			Hosts: vsphere.DistributedVirtualSwitchHostArray{
// 				&vsphere.DistributedVirtualSwitchHostArgs{
// 					HostSystemId: pulumi.String(h1.Id),
// 					Devices: pulumi.StringArray{
// 						pulumi.String("vnic3"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		p1, err := vsphere.NewDistributedPortGroup(ctx, "p1", &vsphere.DistributedPortGroupArgs{
// 			VlanId:                       pulumi.Int(1234),
// 			DistributedVirtualSwitchUuid: d1.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vsphere.NewVnic(ctx, "v1", &vsphere.VnicArgs{
// 			Host:                  pulumi.String(h1.Id),
// 			DistributedSwitchPort: d1.ID(),
// 			DistributedPortGroup:  p1.ID(),
// 			Ipv4: &vsphere.VnicIpv4Args{
// 				Dhcp: pulumi.Bool(true),
// 			},
// 			Netstack: pulumi.String("vmotion"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Create a vnic attached to a portgroup using the default TCP/IP stack
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vsphere/sdk/v2/go/vsphere"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "mydc"
// 		dc, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := "esxi1.host.test"
// 		h1, err := vsphere.LookupHost(ctx, &vsphere.LookupHostArgs{
// 			Name:         &opt1,
// 			DatacenterId: dc.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		hvs1, err := vsphere.NewHostVirtualSwitch(ctx, "hvs1", &vsphere.HostVirtualSwitchArgs{
// 			HostSystemId: pulumi.String(h1.Id),
// 			NetworkAdapters: pulumi.StringArray{
// 				pulumi.String("vmnic3"),
// 				pulumi.String("vmnic4"),
// 			},
// 			ActiveNics: pulumi.StringArray{
// 				pulumi.String("vmnic3"),
// 			},
// 			StandbyNics: pulumi.StringArray{
// 				pulumi.String("vmnic4"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		p1, err := vsphere.NewHostPortGroup(ctx, "p1", &vsphere.HostPortGroupArgs{
// 			VirtualSwitchName: hvs1.Name,
// 			HostSystemId:      pulumi.String(h1.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vsphere.NewVnic(ctx, "v1", &vsphere.VnicArgs{
// 			Host:      pulumi.String(h1.Id),
// 			Portgroup: p1.Name,
// 			Ipv4: &vsphere.VnicIpv4Args{
// 				Dhcp: pulumi.Bool(true),
// 			},
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
// An existing vNic can be [imported][docs-import] into this resource
// via supplying the vNic's ID. An example is below:
//
// [docs-import]: /docs/import/index.html
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
//
// The above would import the vnic `vmk2` from host with ID `host-123`.
type Vnic struct {
	pulumi.CustomResourceState

	// Key of the distributed portgroup the nic will connect to.
	DistributedPortGroup pulumi.StringPtrOutput `pulumi:"distributedPortGroup"`
	// UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
	DistributedSwitchPort pulumi.StringPtrOutput `pulumi:"distributedSwitchPort"`
	// ESX host the interface belongs to
	Host pulumi.StringOutput `pulumi:"host"`
	// IPv4 settings. Either this or `ipv6` needs to be set. See  ipv4 options below.
	Ipv4 VnicIpv4PtrOutput `pulumi:"ipv4"`
	// IPv6 settings. Either this or `ipv6` needs to be set. See  ipv6 options below.
	Ipv6 VnicIpv6PtrOutput `pulumi:"ipv6"`
	// MAC address of the interface.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// MTU of the interface.
	Mtu pulumi.IntOutput `pulumi:"mtu"`
	// TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default: `defaultTcpipStack`)
	Netstack pulumi.StringPtrOutput `pulumi:"netstack"`
	// Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
	Portgroup pulumi.StringPtrOutput `pulumi:"portgroup"`
}

// NewVnic registers a new resource with the given unique name, arguments, and options.
func NewVnic(ctx *pulumi.Context,
	name string, args *VnicArgs, opts ...pulumi.ResourceOption) (*Vnic, error) {
	if args == nil || args.Host == nil {
		return nil, errors.New("missing required argument 'Host'")
	}
	if args == nil {
		args = &VnicArgs{}
	}
	var resource Vnic
	err := ctx.RegisterResource("vsphere:index/vnic:Vnic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVnic gets an existing Vnic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVnic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VnicState, opts ...pulumi.ResourceOption) (*Vnic, error) {
	var resource Vnic
	err := ctx.ReadResource("vsphere:index/vnic:Vnic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vnic resources.
type vnicState struct {
	// Key of the distributed portgroup the nic will connect to.
	DistributedPortGroup *string `pulumi:"distributedPortGroup"`
	// UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
	DistributedSwitchPort *string `pulumi:"distributedSwitchPort"`
	// ESX host the interface belongs to
	Host *string `pulumi:"host"`
	// IPv4 settings. Either this or `ipv6` needs to be set. See  ipv4 options below.
	Ipv4 *VnicIpv4 `pulumi:"ipv4"`
	// IPv6 settings. Either this or `ipv6` needs to be set. See  ipv6 options below.
	Ipv6 *VnicIpv6 `pulumi:"ipv6"`
	// MAC address of the interface.
	Mac *string `pulumi:"mac"`
	// MTU of the interface.
	Mtu *int `pulumi:"mtu"`
	// TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default: `defaultTcpipStack`)
	Netstack *string `pulumi:"netstack"`
	// Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
	Portgroup *string `pulumi:"portgroup"`
}

type VnicState struct {
	// Key of the distributed portgroup the nic will connect to.
	DistributedPortGroup pulumi.StringPtrInput
	// UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
	DistributedSwitchPort pulumi.StringPtrInput
	// ESX host the interface belongs to
	Host pulumi.StringPtrInput
	// IPv4 settings. Either this or `ipv6` needs to be set. See  ipv4 options below.
	Ipv4 VnicIpv4PtrInput
	// IPv6 settings. Either this or `ipv6` needs to be set. See  ipv6 options below.
	Ipv6 VnicIpv6PtrInput
	// MAC address of the interface.
	Mac pulumi.StringPtrInput
	// MTU of the interface.
	Mtu pulumi.IntPtrInput
	// TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default: `defaultTcpipStack`)
	Netstack pulumi.StringPtrInput
	// Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
	Portgroup pulumi.StringPtrInput
}

func (VnicState) ElementType() reflect.Type {
	return reflect.TypeOf((*vnicState)(nil)).Elem()
}

type vnicArgs struct {
	// Key of the distributed portgroup the nic will connect to.
	DistributedPortGroup *string `pulumi:"distributedPortGroup"`
	// UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
	DistributedSwitchPort *string `pulumi:"distributedSwitchPort"`
	// ESX host the interface belongs to
	Host string `pulumi:"host"`
	// IPv4 settings. Either this or `ipv6` needs to be set. See  ipv4 options below.
	Ipv4 *VnicIpv4 `pulumi:"ipv4"`
	// IPv6 settings. Either this or `ipv6` needs to be set. See  ipv6 options below.
	Ipv6 *VnicIpv6 `pulumi:"ipv6"`
	// MAC address of the interface.
	Mac *string `pulumi:"mac"`
	// MTU of the interface.
	Mtu *int `pulumi:"mtu"`
	// TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default: `defaultTcpipStack`)
	Netstack *string `pulumi:"netstack"`
	// Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
	Portgroup *string `pulumi:"portgroup"`
}

// The set of arguments for constructing a Vnic resource.
type VnicArgs struct {
	// Key of the distributed portgroup the nic will connect to.
	DistributedPortGroup pulumi.StringPtrInput
	// UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
	DistributedSwitchPort pulumi.StringPtrInput
	// ESX host the interface belongs to
	Host pulumi.StringInput
	// IPv4 settings. Either this or `ipv6` needs to be set. See  ipv4 options below.
	Ipv4 VnicIpv4PtrInput
	// IPv6 settings. Either this or `ipv6` needs to be set. See  ipv6 options below.
	Ipv6 VnicIpv6PtrInput
	// MAC address of the interface.
	Mac pulumi.StringPtrInput
	// MTU of the interface.
	Mtu pulumi.IntPtrInput
	// TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default: `defaultTcpipStack`)
	Netstack pulumi.StringPtrInput
	// Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
	Portgroup pulumi.StringPtrInput
}

func (VnicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vnicArgs)(nil)).Elem()
}

type VnicInput interface {
	pulumi.Input

	ToVnicOutput() VnicOutput
	ToVnicOutputWithContext(ctx context.Context) VnicOutput
}

func (Vnic) ElementType() reflect.Type {
	return reflect.TypeOf((*Vnic)(nil)).Elem()
}

func (i Vnic) ToVnicOutput() VnicOutput {
	return i.ToVnicOutputWithContext(context.Background())
}

func (i Vnic) ToVnicOutputWithContext(ctx context.Context) VnicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnicOutput)
}

type VnicOutput struct {
	*pulumi.OutputState
}

func (VnicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VnicOutput)(nil)).Elem()
}

func (o VnicOutput) ToVnicOutput() VnicOutput {
	return o
}

func (o VnicOutput) ToVnicOutputWithContext(ctx context.Context) VnicOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VnicOutput{})
}

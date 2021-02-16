// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `HostPortGroup` resource can be used to manage vSphere standard
// port groups on an ESXi host. These port groups are connected to standard
// virtual switches, which can be managed by the
// `HostVirtualSwitch` resource.
//
// For an overview on vSphere networking concepts, see [this page][ref-vsphere-net-concepts].
//
// [ref-vsphere-net-concepts]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.networking.doc/GUID-2B11DBB8-CB3C-4AFF-8885-EFEA0FC562F4.html
//
// ## Example Usage
// ### Create a virtual switch and bind a port group to it
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
// 		opt0 := "dc1"
// 		datacenter, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := "esxi1"
// 		esxiHost, err := vsphere.LookupHost(ctx, &vsphere.LookupHostArgs{
// 			DatacenterId: datacenter.Id,
// 			Name:         &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vsphere.NewHostVirtualSwitch(ctx, "_switch", &vsphere.HostVirtualSwitchArgs{
// 			ActiveNics: pulumi.StringArray{
// 				pulumi.String("vmnic0"),
// 			},
// 			HostSystemId: pulumi.String(esxiHost.Id),
// 			NetworkAdapters: pulumi.StringArray{
// 				pulumi.String("vmnic0"),
// 				pulumi.String("vmnic1"),
// 			},
// 			StandbyNics: pulumi.StringArray{
// 				pulumi.String("vmnic1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vsphere.NewHostPortGroup(ctx, "pg", &vsphere.HostPortGroupArgs{
// 			HostSystemId:      pulumi.String(esxiHost.Id),
// 			VirtualSwitchName: _switch.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Create a port group with VLAN set and some overrides
//
// This example sets the trunk mode VLAN (`4095`, which passes through all tags)
// and sets
// `allowPromiscuous`
// to ensure that all traffic is seen on the port. The latter setting overrides
// the implicit default of `false` set on the virtual switch.
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
// 		opt0 := "dc1"
// 		datacenter, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := "esxi1"
// 		esxiHost, err := vsphere.LookupHost(ctx, &vsphere.LookupHostArgs{
// 			DatacenterId: datacenter.Id,
// 			Name:         &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vsphere.NewHostVirtualSwitch(ctx, "_switch", &vsphere.HostVirtualSwitchArgs{
// 			ActiveNics: pulumi.StringArray{
// 				pulumi.String("vmnic0"),
// 			},
// 			HostSystemId: pulumi.String(esxiHost.Id),
// 			NetworkAdapters: pulumi.StringArray{
// 				pulumi.String("vmnic0"),
// 				pulumi.String("vmnic1"),
// 			},
// 			StandbyNics: pulumi.StringArray{
// 				pulumi.String("vmnic1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vsphere.NewHostPortGroup(ctx, "pg", &vsphere.HostPortGroupArgs{
// 			AllowPromiscuous:  pulumi.Bool(true),
// 			HostSystemId:      pulumi.String(esxiHost.Id),
// 			VirtualSwitchName: _switch.Name,
// 			VlanId:            pulumi.Int(4095),
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
// An existing host port group can be [imported][docs-import] into this resource
// using the host port group's ID. An example is below:
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
// The above would import the `Management` host port group from host with ID `host-123`.
type HostPortGroup struct {
	pulumi.CustomResourceState

	// List of active network adapters used for load balancing.
	ActiveNics pulumi.StringArrayOutput `pulumi:"activeNics"`
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
	// that of its own.
	AllowForgedTransmits pulumi.BoolPtrOutput `pulumi:"allowForgedTransmits"`
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges pulumi.BoolPtrOutput `pulumi:"allowMacChanges"`
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous pulumi.BoolPtrOutput `pulumi:"allowPromiscuous"`
	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
	// only.
	CheckBeacon pulumi.BoolPtrOutput `pulumi:"checkBeacon"`
	// A map with a full set of the policy
	// options computed from defaults and overrides,
	// explaining the effective policy for this port group.
	ComputedPolicy pulumi.StringMapOutput `pulumi:"computedPolicy"`
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback pulumi.BoolPtrOutput `pulumi:"failback"`
	// The managed object ID of
	// the host to set the port group up on. Forces a new resource if changed.
	HostSystemId pulumi.StringOutput `pulumi:"hostSystemId"`
	// The key for this port group as returned from the vSphere API.
	Key pulumi.StringOutput `pulumi:"key"`
	// The name of the port group.  Forces a new resource if
	// changed.
	Name pulumi.StringOutput `pulumi:"name"`
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches pulumi.BoolPtrOutput `pulumi:"notifySwitches"`
	// A list of ports that currently exist and are used on this port group.
	Ports HostPortGroupPortsOutput `pulumi:"ports"`
	// The average bandwidth in bits per second if traffic shaping is enabled.
	ShapingAverageBandwidth pulumi.IntPtrOutput `pulumi:"shapingAverageBandwidth"`
	// The maximum burst size allowed in bytes if traffic shaping is enabled.
	ShapingBurstSize pulumi.IntPtrOutput `pulumi:"shapingBurstSize"`
	// Enable traffic shaping on this virtual switch or port group.
	ShapingEnabled pulumi.BoolPtrOutput `pulumi:"shapingEnabled"`
	// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
	ShapingPeakBandwidth pulumi.IntPtrOutput `pulumi:"shapingPeakBandwidth"`
	// List of standby network adapters used for failover.
	StandbyNics pulumi.StringArrayOutput `pulumi:"standbyNics"`
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
	// failover_explicit.
	TeamingPolicy pulumi.StringPtrOutput `pulumi:"teamingPolicy"`
	// The name of the virtual switch to bind
	// this port group to. Forces a new resource if changed.
	VirtualSwitchName pulumi.StringOutput `pulumi:"virtualSwitchName"`
	// The VLAN ID/trunk mode for this port group.  An ID of
	// `0` denotes no tagging, an ID of `1`-`4094` tags with the specific ID, and an
	// ID of `4095` enables trunk mode, allowing the guest to manage its own
	// tagging. Default: `0`.
	VlanId pulumi.IntPtrOutput `pulumi:"vlanId"`
}

// NewHostPortGroup registers a new resource with the given unique name, arguments, and options.
func NewHostPortGroup(ctx *pulumi.Context,
	name string, args *HostPortGroupArgs, opts ...pulumi.ResourceOption) (*HostPortGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostSystemId == nil {
		return nil, errors.New("invalid value for required argument 'HostSystemId'")
	}
	if args.VirtualSwitchName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualSwitchName'")
	}
	var resource HostPortGroup
	err := ctx.RegisterResource("vsphere:index/hostPortGroup:HostPortGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostPortGroup gets an existing HostPortGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostPortGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostPortGroupState, opts ...pulumi.ResourceOption) (*HostPortGroup, error) {
	var resource HostPortGroup
	err := ctx.ReadResource("vsphere:index/hostPortGroup:HostPortGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostPortGroup resources.
type hostPortGroupState struct {
	// List of active network adapters used for load balancing.
	ActiveNics []string `pulumi:"activeNics"`
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
	// that of its own.
	AllowForgedTransmits *bool `pulumi:"allowForgedTransmits"`
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges *bool `pulumi:"allowMacChanges"`
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous *bool `pulumi:"allowPromiscuous"`
	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
	// only.
	CheckBeacon *bool `pulumi:"checkBeacon"`
	// A map with a full set of the policy
	// options computed from defaults and overrides,
	// explaining the effective policy for this port group.
	ComputedPolicy map[string]string `pulumi:"computedPolicy"`
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback *bool `pulumi:"failback"`
	// The managed object ID of
	// the host to set the port group up on. Forces a new resource if changed.
	HostSystemId *string `pulumi:"hostSystemId"`
	// The key for this port group as returned from the vSphere API.
	Key *string `pulumi:"key"`
	// The name of the port group.  Forces a new resource if
	// changed.
	Name *string `pulumi:"name"`
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches *bool `pulumi:"notifySwitches"`
	// A list of ports that currently exist and are used on this port group.
	Ports *HostPortGroupPorts `pulumi:"ports"`
	// The average bandwidth in bits per second if traffic shaping is enabled.
	ShapingAverageBandwidth *int `pulumi:"shapingAverageBandwidth"`
	// The maximum burst size allowed in bytes if traffic shaping is enabled.
	ShapingBurstSize *int `pulumi:"shapingBurstSize"`
	// Enable traffic shaping on this virtual switch or port group.
	ShapingEnabled *bool `pulumi:"shapingEnabled"`
	// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
	ShapingPeakBandwidth *int `pulumi:"shapingPeakBandwidth"`
	// List of standby network adapters used for failover.
	StandbyNics []string `pulumi:"standbyNics"`
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
	// failover_explicit.
	TeamingPolicy *string `pulumi:"teamingPolicy"`
	// The name of the virtual switch to bind
	// this port group to. Forces a new resource if changed.
	VirtualSwitchName *string `pulumi:"virtualSwitchName"`
	// The VLAN ID/trunk mode for this port group.  An ID of
	// `0` denotes no tagging, an ID of `1`-`4094` tags with the specific ID, and an
	// ID of `4095` enables trunk mode, allowing the guest to manage its own
	// tagging. Default: `0`.
	VlanId *int `pulumi:"vlanId"`
}

type HostPortGroupState struct {
	// List of active network adapters used for load balancing.
	ActiveNics pulumi.StringArrayInput
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
	// that of its own.
	AllowForgedTransmits pulumi.BoolPtrInput
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges pulumi.BoolPtrInput
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous pulumi.BoolPtrInput
	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
	// only.
	CheckBeacon pulumi.BoolPtrInput
	// A map with a full set of the policy
	// options computed from defaults and overrides,
	// explaining the effective policy for this port group.
	ComputedPolicy pulumi.StringMapInput
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback pulumi.BoolPtrInput
	// The managed object ID of
	// the host to set the port group up on. Forces a new resource if changed.
	HostSystemId pulumi.StringPtrInput
	// The key for this port group as returned from the vSphere API.
	Key pulumi.StringPtrInput
	// The name of the port group.  Forces a new resource if
	// changed.
	Name pulumi.StringPtrInput
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches pulumi.BoolPtrInput
	// A list of ports that currently exist and are used on this port group.
	Ports HostPortGroupPortsPtrInput
	// The average bandwidth in bits per second if traffic shaping is enabled.
	ShapingAverageBandwidth pulumi.IntPtrInput
	// The maximum burst size allowed in bytes if traffic shaping is enabled.
	ShapingBurstSize pulumi.IntPtrInput
	// Enable traffic shaping on this virtual switch or port group.
	ShapingEnabled pulumi.BoolPtrInput
	// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
	ShapingPeakBandwidth pulumi.IntPtrInput
	// List of standby network adapters used for failover.
	StandbyNics pulumi.StringArrayInput
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
	// failover_explicit.
	TeamingPolicy pulumi.StringPtrInput
	// The name of the virtual switch to bind
	// this port group to. Forces a new resource if changed.
	VirtualSwitchName pulumi.StringPtrInput
	// The VLAN ID/trunk mode for this port group.  An ID of
	// `0` denotes no tagging, an ID of `1`-`4094` tags with the specific ID, and an
	// ID of `4095` enables trunk mode, allowing the guest to manage its own
	// tagging. Default: `0`.
	VlanId pulumi.IntPtrInput
}

func (HostPortGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostPortGroupState)(nil)).Elem()
}

type hostPortGroupArgs struct {
	// List of active network adapters used for load balancing.
	ActiveNics []string `pulumi:"activeNics"`
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
	// that of its own.
	AllowForgedTransmits *bool `pulumi:"allowForgedTransmits"`
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges *bool `pulumi:"allowMacChanges"`
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous *bool `pulumi:"allowPromiscuous"`
	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
	// only.
	CheckBeacon *bool `pulumi:"checkBeacon"`
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback *bool `pulumi:"failback"`
	// The managed object ID of
	// the host to set the port group up on. Forces a new resource if changed.
	HostSystemId string `pulumi:"hostSystemId"`
	// The name of the port group.  Forces a new resource if
	// changed.
	Name *string `pulumi:"name"`
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches *bool `pulumi:"notifySwitches"`
	// The average bandwidth in bits per second if traffic shaping is enabled.
	ShapingAverageBandwidth *int `pulumi:"shapingAverageBandwidth"`
	// The maximum burst size allowed in bytes if traffic shaping is enabled.
	ShapingBurstSize *int `pulumi:"shapingBurstSize"`
	// Enable traffic shaping on this virtual switch or port group.
	ShapingEnabled *bool `pulumi:"shapingEnabled"`
	// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
	ShapingPeakBandwidth *int `pulumi:"shapingPeakBandwidth"`
	// List of standby network adapters used for failover.
	StandbyNics []string `pulumi:"standbyNics"`
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
	// failover_explicit.
	TeamingPolicy *string `pulumi:"teamingPolicy"`
	// The name of the virtual switch to bind
	// this port group to. Forces a new resource if changed.
	VirtualSwitchName string `pulumi:"virtualSwitchName"`
	// The VLAN ID/trunk mode for this port group.  An ID of
	// `0` denotes no tagging, an ID of `1`-`4094` tags with the specific ID, and an
	// ID of `4095` enables trunk mode, allowing the guest to manage its own
	// tagging. Default: `0`.
	VlanId *int `pulumi:"vlanId"`
}

// The set of arguments for constructing a HostPortGroup resource.
type HostPortGroupArgs struct {
	// List of active network adapters used for load balancing.
	ActiveNics pulumi.StringArrayInput
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
	// that of its own.
	AllowForgedTransmits pulumi.BoolPtrInput
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges pulumi.BoolPtrInput
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous pulumi.BoolPtrInput
	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
	// only.
	CheckBeacon pulumi.BoolPtrInput
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback pulumi.BoolPtrInput
	// The managed object ID of
	// the host to set the port group up on. Forces a new resource if changed.
	HostSystemId pulumi.StringInput
	// The name of the port group.  Forces a new resource if
	// changed.
	Name pulumi.StringPtrInput
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches pulumi.BoolPtrInput
	// The average bandwidth in bits per second if traffic shaping is enabled.
	ShapingAverageBandwidth pulumi.IntPtrInput
	// The maximum burst size allowed in bytes if traffic shaping is enabled.
	ShapingBurstSize pulumi.IntPtrInput
	// Enable traffic shaping on this virtual switch or port group.
	ShapingEnabled pulumi.BoolPtrInput
	// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
	ShapingPeakBandwidth pulumi.IntPtrInput
	// List of standby network adapters used for failover.
	StandbyNics pulumi.StringArrayInput
	// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
	// failover_explicit.
	TeamingPolicy pulumi.StringPtrInput
	// The name of the virtual switch to bind
	// this port group to. Forces a new resource if changed.
	VirtualSwitchName pulumi.StringInput
	// The VLAN ID/trunk mode for this port group.  An ID of
	// `0` denotes no tagging, an ID of `1`-`4094` tags with the specific ID, and an
	// ID of `4095` enables trunk mode, allowing the guest to manage its own
	// tagging. Default: `0`.
	VlanId pulumi.IntPtrInput
}

func (HostPortGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostPortGroupArgs)(nil)).Elem()
}

type HostPortGroupInput interface {
	pulumi.Input

	ToHostPortGroupOutput() HostPortGroupOutput
	ToHostPortGroupOutputWithContext(ctx context.Context) HostPortGroupOutput
}

func (*HostPortGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*HostPortGroup)(nil))
}

func (i *HostPortGroup) ToHostPortGroupOutput() HostPortGroupOutput {
	return i.ToHostPortGroupOutputWithContext(context.Background())
}

func (i *HostPortGroup) ToHostPortGroupOutputWithContext(ctx context.Context) HostPortGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostPortGroupOutput)
}

func (i *HostPortGroup) ToHostPortGroupPtrOutput() HostPortGroupPtrOutput {
	return i.ToHostPortGroupPtrOutputWithContext(context.Background())
}

func (i *HostPortGroup) ToHostPortGroupPtrOutputWithContext(ctx context.Context) HostPortGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostPortGroupPtrOutput)
}

type HostPortGroupPtrInput interface {
	pulumi.Input

	ToHostPortGroupPtrOutput() HostPortGroupPtrOutput
	ToHostPortGroupPtrOutputWithContext(ctx context.Context) HostPortGroupPtrOutput
}

type hostPortGroupPtrType HostPortGroupArgs

func (*hostPortGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HostPortGroup)(nil))
}

func (i *hostPortGroupPtrType) ToHostPortGroupPtrOutput() HostPortGroupPtrOutput {
	return i.ToHostPortGroupPtrOutputWithContext(context.Background())
}

func (i *hostPortGroupPtrType) ToHostPortGroupPtrOutputWithContext(ctx context.Context) HostPortGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostPortGroupPtrOutput)
}

// HostPortGroupArrayInput is an input type that accepts HostPortGroupArray and HostPortGroupArrayOutput values.
// You can construct a concrete instance of `HostPortGroupArrayInput` via:
//
//          HostPortGroupArray{ HostPortGroupArgs{...} }
type HostPortGroupArrayInput interface {
	pulumi.Input

	ToHostPortGroupArrayOutput() HostPortGroupArrayOutput
	ToHostPortGroupArrayOutputWithContext(context.Context) HostPortGroupArrayOutput
}

type HostPortGroupArray []HostPortGroupInput

func (HostPortGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*HostPortGroup)(nil))
}

func (i HostPortGroupArray) ToHostPortGroupArrayOutput() HostPortGroupArrayOutput {
	return i.ToHostPortGroupArrayOutputWithContext(context.Background())
}

func (i HostPortGroupArray) ToHostPortGroupArrayOutputWithContext(ctx context.Context) HostPortGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostPortGroupArrayOutput)
}

// HostPortGroupMapInput is an input type that accepts HostPortGroupMap and HostPortGroupMapOutput values.
// You can construct a concrete instance of `HostPortGroupMapInput` via:
//
//          HostPortGroupMap{ "key": HostPortGroupArgs{...} }
type HostPortGroupMapInput interface {
	pulumi.Input

	ToHostPortGroupMapOutput() HostPortGroupMapOutput
	ToHostPortGroupMapOutputWithContext(context.Context) HostPortGroupMapOutput
}

type HostPortGroupMap map[string]HostPortGroupInput

func (HostPortGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*HostPortGroup)(nil))
}

func (i HostPortGroupMap) ToHostPortGroupMapOutput() HostPortGroupMapOutput {
	return i.ToHostPortGroupMapOutputWithContext(context.Background())
}

func (i HostPortGroupMap) ToHostPortGroupMapOutputWithContext(ctx context.Context) HostPortGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostPortGroupMapOutput)
}

type HostPortGroupOutput struct {
	*pulumi.OutputState
}

func (HostPortGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostPortGroup)(nil))
}

func (o HostPortGroupOutput) ToHostPortGroupOutput() HostPortGroupOutput {
	return o
}

func (o HostPortGroupOutput) ToHostPortGroupOutputWithContext(ctx context.Context) HostPortGroupOutput {
	return o
}

func (o HostPortGroupOutput) ToHostPortGroupPtrOutput() HostPortGroupPtrOutput {
	return o.ToHostPortGroupPtrOutputWithContext(context.Background())
}

func (o HostPortGroupOutput) ToHostPortGroupPtrOutputWithContext(ctx context.Context) HostPortGroupPtrOutput {
	return o.ApplyT(func(v HostPortGroup) *HostPortGroup {
		return &v
	}).(HostPortGroupPtrOutput)
}

type HostPortGroupPtrOutput struct {
	*pulumi.OutputState
}

func (HostPortGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostPortGroup)(nil))
}

func (o HostPortGroupPtrOutput) ToHostPortGroupPtrOutput() HostPortGroupPtrOutput {
	return o
}

func (o HostPortGroupPtrOutput) ToHostPortGroupPtrOutputWithContext(ctx context.Context) HostPortGroupPtrOutput {
	return o
}

type HostPortGroupArrayOutput struct{ *pulumi.OutputState }

func (HostPortGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostPortGroup)(nil))
}

func (o HostPortGroupArrayOutput) ToHostPortGroupArrayOutput() HostPortGroupArrayOutput {
	return o
}

func (o HostPortGroupArrayOutput) ToHostPortGroupArrayOutputWithContext(ctx context.Context) HostPortGroupArrayOutput {
	return o
}

func (o HostPortGroupArrayOutput) Index(i pulumi.IntInput) HostPortGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostPortGroup {
		return vs[0].([]HostPortGroup)[vs[1].(int)]
	}).(HostPortGroupOutput)
}

type HostPortGroupMapOutput struct{ *pulumi.OutputState }

func (HostPortGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]HostPortGroup)(nil))
}

func (o HostPortGroupMapOutput) ToHostPortGroupMapOutput() HostPortGroupMapOutput {
	return o
}

func (o HostPortGroupMapOutput) ToHostPortGroupMapOutputWithContext(ctx context.Context) HostPortGroupMapOutput {
	return o
}

func (o HostPortGroupMapOutput) MapIndex(k pulumi.StringInput) HostPortGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) HostPortGroup {
		return vs[0].(map[string]HostPortGroup)[vs[1].(string)]
	}).(HostPortGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(HostPortGroupOutput{})
	pulumi.RegisterOutputType(HostPortGroupPtrOutput{})
	pulumi.RegisterOutputType(HostPortGroupArrayOutput{})
	pulumi.RegisterOutputType(HostPortGroupMapOutput{})
}

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

// The `HostVirtualSwitch` resource can be used to manage vSphere
// standard switches on an ESXi host. These switches can be used as a backing for
// standard port groups, which can be managed by the
// `HostPortGroup` resource.
//
// For an overview on vSphere networking concepts, see [this
// page][ref-vsphere-net-concepts].
//
// ## Example Usage
//
// ## Import
//
// An existing vSwitch can be imported into this resource by its ID.
//
// The convention of the id is a prefix, the host system [managed objectID][docs-about-morefs], and the virtual switch
//
// name. An example would be `tf-HostVirtualSwitch:host-10:vSwitchTerraformTest`.
//
// Import can the be done via the following command:
//
// [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
//
// ```sh
// $ pulumi import vsphere:index/hostVirtualSwitch:HostVirtualSwitch switch tf-HostVirtualSwitch:host-10:vSwitchTerraformTest
// ```
//
// The above would import the vSwitch named `vSwitchTerraformTest` that is located in the `host-10`
//
// vSphere host.
//
// [ref-vsphere-net-concepts]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.networking.doc/GUID-2B11DBB8-CB3C-4AFF-8885-EFEA0FC562F4.html
type HostVirtualSwitch struct {
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
	// Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
	BeaconInterval pulumi.IntPtrOutput `pulumi:"beaconInterval"`
	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
	// only.
	CheckBeacon pulumi.BoolPtrOutput `pulumi:"checkBeacon"`
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback pulumi.BoolPtrOutput `pulumi:"failback"`
	// The managed object ID of
	// the host to set the virtual switch up on. Forces a new resource if changed.
	HostSystemId pulumi.StringOutput `pulumi:"hostSystemId"`
	// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
	LinkDiscoveryOperation pulumi.StringPtrOutput `pulumi:"linkDiscoveryOperation"`
	// The discovery protocol type. Valid values are cdp and lldp.
	LinkDiscoveryProtocol pulumi.StringPtrOutput `pulumi:"linkDiscoveryProtocol"`
	// The maximum transmission unit (MTU) for the virtual
	// switch. Default: `1500`.
	Mtu pulumi.IntPtrOutput `pulumi:"mtu"`
	// The name of the virtual switch. Forces a new resource if
	// changed.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of network adapters to bind to this virtual switch.
	NetworkAdapters pulumi.StringArrayOutput `pulumi:"networkAdapters"`
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches pulumi.BoolPtrOutput `pulumi:"notifySwitches"`
	// The number of ports to create with this
	// virtual switch. Default: `128`.
	//
	// > **NOTE:** Changing the port count requires a reboot of the host. This provider
	// will not restart the host for you.
	NumberOfPorts pulumi.IntPtrOutput `pulumi:"numberOfPorts"`
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
}

// NewHostVirtualSwitch registers a new resource with the given unique name, arguments, and options.
func NewHostVirtualSwitch(ctx *pulumi.Context,
	name string, args *HostVirtualSwitchArgs, opts ...pulumi.ResourceOption) (*HostVirtualSwitch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ActiveNics == nil {
		return nil, errors.New("invalid value for required argument 'ActiveNics'")
	}
	if args.HostSystemId == nil {
		return nil, errors.New("invalid value for required argument 'HostSystemId'")
	}
	if args.NetworkAdapters == nil {
		return nil, errors.New("invalid value for required argument 'NetworkAdapters'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HostVirtualSwitch
	err := ctx.RegisterResource("vsphere:index/hostVirtualSwitch:HostVirtualSwitch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostVirtualSwitch gets an existing HostVirtualSwitch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostVirtualSwitch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostVirtualSwitchState, opts ...pulumi.ResourceOption) (*HostVirtualSwitch, error) {
	var resource HostVirtualSwitch
	err := ctx.ReadResource("vsphere:index/hostVirtualSwitch:HostVirtualSwitch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostVirtualSwitch resources.
type hostVirtualSwitchState struct {
	// List of active network adapters used for load balancing.
	ActiveNics []string `pulumi:"activeNics"`
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
	// that of its own.
	AllowForgedTransmits *bool `pulumi:"allowForgedTransmits"`
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges *bool `pulumi:"allowMacChanges"`
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous *bool `pulumi:"allowPromiscuous"`
	// Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
	BeaconInterval *int `pulumi:"beaconInterval"`
	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
	// only.
	CheckBeacon *bool `pulumi:"checkBeacon"`
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback *bool `pulumi:"failback"`
	// The managed object ID of
	// the host to set the virtual switch up on. Forces a new resource if changed.
	HostSystemId *string `pulumi:"hostSystemId"`
	// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
	LinkDiscoveryOperation *string `pulumi:"linkDiscoveryOperation"`
	// The discovery protocol type. Valid values are cdp and lldp.
	LinkDiscoveryProtocol *string `pulumi:"linkDiscoveryProtocol"`
	// The maximum transmission unit (MTU) for the virtual
	// switch. Default: `1500`.
	Mtu *int `pulumi:"mtu"`
	// The name of the virtual switch. Forces a new resource if
	// changed.
	Name *string `pulumi:"name"`
	// The list of network adapters to bind to this virtual switch.
	NetworkAdapters []string `pulumi:"networkAdapters"`
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches *bool `pulumi:"notifySwitches"`
	// The number of ports to create with this
	// virtual switch. Default: `128`.
	//
	// > **NOTE:** Changing the port count requires a reboot of the host. This provider
	// will not restart the host for you.
	NumberOfPorts *int `pulumi:"numberOfPorts"`
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
}

type HostVirtualSwitchState struct {
	// List of active network adapters used for load balancing.
	ActiveNics pulumi.StringArrayInput
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
	// that of its own.
	AllowForgedTransmits pulumi.BoolPtrInput
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges pulumi.BoolPtrInput
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous pulumi.BoolPtrInput
	// Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
	BeaconInterval pulumi.IntPtrInput
	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
	// only.
	CheckBeacon pulumi.BoolPtrInput
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback pulumi.BoolPtrInput
	// The managed object ID of
	// the host to set the virtual switch up on. Forces a new resource if changed.
	HostSystemId pulumi.StringPtrInput
	// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
	LinkDiscoveryOperation pulumi.StringPtrInput
	// The discovery protocol type. Valid values are cdp and lldp.
	LinkDiscoveryProtocol pulumi.StringPtrInput
	// The maximum transmission unit (MTU) for the virtual
	// switch. Default: `1500`.
	Mtu pulumi.IntPtrInput
	// The name of the virtual switch. Forces a new resource if
	// changed.
	Name pulumi.StringPtrInput
	// The list of network adapters to bind to this virtual switch.
	NetworkAdapters pulumi.StringArrayInput
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches pulumi.BoolPtrInput
	// The number of ports to create with this
	// virtual switch. Default: `128`.
	//
	// > **NOTE:** Changing the port count requires a reboot of the host. This provider
	// will not restart the host for you.
	NumberOfPorts pulumi.IntPtrInput
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
}

func (HostVirtualSwitchState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostVirtualSwitchState)(nil)).Elem()
}

type hostVirtualSwitchArgs struct {
	// List of active network adapters used for load balancing.
	ActiveNics []string `pulumi:"activeNics"`
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
	// that of its own.
	AllowForgedTransmits *bool `pulumi:"allowForgedTransmits"`
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges *bool `pulumi:"allowMacChanges"`
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous *bool `pulumi:"allowPromiscuous"`
	// Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
	BeaconInterval *int `pulumi:"beaconInterval"`
	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
	// only.
	CheckBeacon *bool `pulumi:"checkBeacon"`
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback *bool `pulumi:"failback"`
	// The managed object ID of
	// the host to set the virtual switch up on. Forces a new resource if changed.
	HostSystemId string `pulumi:"hostSystemId"`
	// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
	LinkDiscoveryOperation *string `pulumi:"linkDiscoveryOperation"`
	// The discovery protocol type. Valid values are cdp and lldp.
	LinkDiscoveryProtocol *string `pulumi:"linkDiscoveryProtocol"`
	// The maximum transmission unit (MTU) for the virtual
	// switch. Default: `1500`.
	Mtu *int `pulumi:"mtu"`
	// The name of the virtual switch. Forces a new resource if
	// changed.
	Name *string `pulumi:"name"`
	// The list of network adapters to bind to this virtual switch.
	NetworkAdapters []string `pulumi:"networkAdapters"`
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches *bool `pulumi:"notifySwitches"`
	// The number of ports to create with this
	// virtual switch. Default: `128`.
	//
	// > **NOTE:** Changing the port count requires a reboot of the host. This provider
	// will not restart the host for you.
	NumberOfPorts *int `pulumi:"numberOfPorts"`
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
}

// The set of arguments for constructing a HostVirtualSwitch resource.
type HostVirtualSwitchArgs struct {
	// List of active network adapters used for load balancing.
	ActiveNics pulumi.StringArrayInput
	// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
	// that of its own.
	AllowForgedTransmits pulumi.BoolPtrInput
	// Controls whether or not the Media Access Control (MAC) address can be changed.
	AllowMacChanges pulumi.BoolPtrInput
	// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
	AllowPromiscuous pulumi.BoolPtrInput
	// Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
	BeaconInterval pulumi.IntPtrInput
	// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
	// only.
	CheckBeacon pulumi.BoolPtrInput
	// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
	Failback pulumi.BoolPtrInput
	// The managed object ID of
	// the host to set the virtual switch up on. Forces a new resource if changed.
	HostSystemId pulumi.StringInput
	// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
	LinkDiscoveryOperation pulumi.StringPtrInput
	// The discovery protocol type. Valid values are cdp and lldp.
	LinkDiscoveryProtocol pulumi.StringPtrInput
	// The maximum transmission unit (MTU) for the virtual
	// switch. Default: `1500`.
	Mtu pulumi.IntPtrInput
	// The name of the virtual switch. Forces a new resource if
	// changed.
	Name pulumi.StringPtrInput
	// The list of network adapters to bind to this virtual switch.
	NetworkAdapters pulumi.StringArrayInput
	// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
	NotifySwitches pulumi.BoolPtrInput
	// The number of ports to create with this
	// virtual switch. Default: `128`.
	//
	// > **NOTE:** Changing the port count requires a reboot of the host. This provider
	// will not restart the host for you.
	NumberOfPorts pulumi.IntPtrInput
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
}

func (HostVirtualSwitchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostVirtualSwitchArgs)(nil)).Elem()
}

type HostVirtualSwitchInput interface {
	pulumi.Input

	ToHostVirtualSwitchOutput() HostVirtualSwitchOutput
	ToHostVirtualSwitchOutputWithContext(ctx context.Context) HostVirtualSwitchOutput
}

func (*HostVirtualSwitch) ElementType() reflect.Type {
	return reflect.TypeOf((**HostVirtualSwitch)(nil)).Elem()
}

func (i *HostVirtualSwitch) ToHostVirtualSwitchOutput() HostVirtualSwitchOutput {
	return i.ToHostVirtualSwitchOutputWithContext(context.Background())
}

func (i *HostVirtualSwitch) ToHostVirtualSwitchOutputWithContext(ctx context.Context) HostVirtualSwitchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostVirtualSwitchOutput)
}

// HostVirtualSwitchArrayInput is an input type that accepts HostVirtualSwitchArray and HostVirtualSwitchArrayOutput values.
// You can construct a concrete instance of `HostVirtualSwitchArrayInput` via:
//
//	HostVirtualSwitchArray{ HostVirtualSwitchArgs{...} }
type HostVirtualSwitchArrayInput interface {
	pulumi.Input

	ToHostVirtualSwitchArrayOutput() HostVirtualSwitchArrayOutput
	ToHostVirtualSwitchArrayOutputWithContext(context.Context) HostVirtualSwitchArrayOutput
}

type HostVirtualSwitchArray []HostVirtualSwitchInput

func (HostVirtualSwitchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostVirtualSwitch)(nil)).Elem()
}

func (i HostVirtualSwitchArray) ToHostVirtualSwitchArrayOutput() HostVirtualSwitchArrayOutput {
	return i.ToHostVirtualSwitchArrayOutputWithContext(context.Background())
}

func (i HostVirtualSwitchArray) ToHostVirtualSwitchArrayOutputWithContext(ctx context.Context) HostVirtualSwitchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostVirtualSwitchArrayOutput)
}

// HostVirtualSwitchMapInput is an input type that accepts HostVirtualSwitchMap and HostVirtualSwitchMapOutput values.
// You can construct a concrete instance of `HostVirtualSwitchMapInput` via:
//
//	HostVirtualSwitchMap{ "key": HostVirtualSwitchArgs{...} }
type HostVirtualSwitchMapInput interface {
	pulumi.Input

	ToHostVirtualSwitchMapOutput() HostVirtualSwitchMapOutput
	ToHostVirtualSwitchMapOutputWithContext(context.Context) HostVirtualSwitchMapOutput
}

type HostVirtualSwitchMap map[string]HostVirtualSwitchInput

func (HostVirtualSwitchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostVirtualSwitch)(nil)).Elem()
}

func (i HostVirtualSwitchMap) ToHostVirtualSwitchMapOutput() HostVirtualSwitchMapOutput {
	return i.ToHostVirtualSwitchMapOutputWithContext(context.Background())
}

func (i HostVirtualSwitchMap) ToHostVirtualSwitchMapOutputWithContext(ctx context.Context) HostVirtualSwitchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostVirtualSwitchMapOutput)
}

type HostVirtualSwitchOutput struct{ *pulumi.OutputState }

func (HostVirtualSwitchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostVirtualSwitch)(nil)).Elem()
}

func (o HostVirtualSwitchOutput) ToHostVirtualSwitchOutput() HostVirtualSwitchOutput {
	return o
}

func (o HostVirtualSwitchOutput) ToHostVirtualSwitchOutputWithContext(ctx context.Context) HostVirtualSwitchOutput {
	return o
}

// List of active network adapters used for load balancing.
func (o HostVirtualSwitchOutput) ActiveNics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.StringArrayOutput { return v.ActiveNics }).(pulumi.StringArrayOutput)
}

// Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
// that of its own.
func (o HostVirtualSwitchOutput) AllowForgedTransmits() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.BoolPtrOutput { return v.AllowForgedTransmits }).(pulumi.BoolPtrOutput)
}

// Controls whether or not the Media Access Control (MAC) address can be changed.
func (o HostVirtualSwitchOutput) AllowMacChanges() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.BoolPtrOutput { return v.AllowMacChanges }).(pulumi.BoolPtrOutput)
}

// Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
func (o HostVirtualSwitchOutput) AllowPromiscuous() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.BoolPtrOutput { return v.AllowPromiscuous }).(pulumi.BoolPtrOutput)
}

// Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
func (o HostVirtualSwitchOutput) BeaconInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.IntPtrOutput { return v.BeaconInterval }).(pulumi.IntPtrOutput)
}

// Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
// only.
func (o HostVirtualSwitchOutput) CheckBeacon() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.BoolPtrOutput { return v.CheckBeacon }).(pulumi.BoolPtrOutput)
}

// If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
func (o HostVirtualSwitchOutput) Failback() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.BoolPtrOutput { return v.Failback }).(pulumi.BoolPtrOutput)
}

// The managed object ID of
// the host to set the virtual switch up on. Forces a new resource if changed.
func (o HostVirtualSwitchOutput) HostSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.StringOutput { return v.HostSystemId }).(pulumi.StringOutput)
}

// Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
func (o HostVirtualSwitchOutput) LinkDiscoveryOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.StringPtrOutput { return v.LinkDiscoveryOperation }).(pulumi.StringPtrOutput)
}

// The discovery protocol type. Valid values are cdp and lldp.
func (o HostVirtualSwitchOutput) LinkDiscoveryProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.StringPtrOutput { return v.LinkDiscoveryProtocol }).(pulumi.StringPtrOutput)
}

// The maximum transmission unit (MTU) for the virtual
// switch. Default: `1500`.
func (o HostVirtualSwitchOutput) Mtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.IntPtrOutput { return v.Mtu }).(pulumi.IntPtrOutput)
}

// The name of the virtual switch. Forces a new resource if
// changed.
func (o HostVirtualSwitchOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of network adapters to bind to this virtual switch.
func (o HostVirtualSwitchOutput) NetworkAdapters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.StringArrayOutput { return v.NetworkAdapters }).(pulumi.StringArrayOutput)
}

// If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
func (o HostVirtualSwitchOutput) NotifySwitches() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.BoolPtrOutput { return v.NotifySwitches }).(pulumi.BoolPtrOutput)
}

// The number of ports to create with this
// virtual switch. Default: `128`.
//
// > **NOTE:** Changing the port count requires a reboot of the host. This provider
// will not restart the host for you.
func (o HostVirtualSwitchOutput) NumberOfPorts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.IntPtrOutput { return v.NumberOfPorts }).(pulumi.IntPtrOutput)
}

// The average bandwidth in bits per second if traffic shaping is enabled.
func (o HostVirtualSwitchOutput) ShapingAverageBandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.IntPtrOutput { return v.ShapingAverageBandwidth }).(pulumi.IntPtrOutput)
}

// The maximum burst size allowed in bytes if traffic shaping is enabled.
func (o HostVirtualSwitchOutput) ShapingBurstSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.IntPtrOutput { return v.ShapingBurstSize }).(pulumi.IntPtrOutput)
}

// Enable traffic shaping on this virtual switch or port group.
func (o HostVirtualSwitchOutput) ShapingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.BoolPtrOutput { return v.ShapingEnabled }).(pulumi.BoolPtrOutput)
}

// The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
func (o HostVirtualSwitchOutput) ShapingPeakBandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.IntPtrOutput { return v.ShapingPeakBandwidth }).(pulumi.IntPtrOutput)
}

// List of standby network adapters used for failover.
func (o HostVirtualSwitchOutput) StandbyNics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.StringArrayOutput { return v.StandbyNics }).(pulumi.StringArrayOutput)
}

// The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
// failover_explicit.
func (o HostVirtualSwitchOutput) TeamingPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostVirtualSwitch) pulumi.StringPtrOutput { return v.TeamingPolicy }).(pulumi.StringPtrOutput)
}

type HostVirtualSwitchArrayOutput struct{ *pulumi.OutputState }

func (HostVirtualSwitchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostVirtualSwitch)(nil)).Elem()
}

func (o HostVirtualSwitchArrayOutput) ToHostVirtualSwitchArrayOutput() HostVirtualSwitchArrayOutput {
	return o
}

func (o HostVirtualSwitchArrayOutput) ToHostVirtualSwitchArrayOutputWithContext(ctx context.Context) HostVirtualSwitchArrayOutput {
	return o
}

func (o HostVirtualSwitchArrayOutput) Index(i pulumi.IntInput) HostVirtualSwitchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HostVirtualSwitch {
		return vs[0].([]*HostVirtualSwitch)[vs[1].(int)]
	}).(HostVirtualSwitchOutput)
}

type HostVirtualSwitchMapOutput struct{ *pulumi.OutputState }

func (HostVirtualSwitchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostVirtualSwitch)(nil)).Elem()
}

func (o HostVirtualSwitchMapOutput) ToHostVirtualSwitchMapOutput() HostVirtualSwitchMapOutput {
	return o
}

func (o HostVirtualSwitchMapOutput) ToHostVirtualSwitchMapOutputWithContext(ctx context.Context) HostVirtualSwitchMapOutput {
	return o
}

func (o HostVirtualSwitchMapOutput) MapIndex(k pulumi.StringInput) HostVirtualSwitchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HostVirtualSwitch {
		return vs[0].(map[string]*HostVirtualSwitch)[vs[1].(string)]
	}).(HostVirtualSwitchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostVirtualSwitchInput)(nil)).Elem(), &HostVirtualSwitch{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostVirtualSwitchArrayInput)(nil)).Elem(), HostVirtualSwitchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostVirtualSwitchMapInput)(nil)).Elem(), HostVirtualSwitchMap{})
	pulumi.RegisterOutputType(HostVirtualSwitchOutput{})
	pulumi.RegisterOutputType(HostVirtualSwitchArrayOutput{})
	pulumi.RegisterOutputType(HostVirtualSwitchMapOutput{})
}

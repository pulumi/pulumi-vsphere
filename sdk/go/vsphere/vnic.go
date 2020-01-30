// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a VMware vSphere vnic resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/vnic.html.markdown.
type Vnic struct {
	s *pulumi.ResourceState
}

// NewVnic registers a new resource with the given unique name, arguments, and options.
func NewVnic(ctx *pulumi.Context,
	name string, args *VnicArgs, opts ...pulumi.ResourceOpt) (*Vnic, error) {
	if args == nil || args.Host == nil {
		return nil, errors.New("missing required argument 'Host'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["distributedPortGroup"] = nil
		inputs["distributedSwitchPort"] = nil
		inputs["host"] = nil
		inputs["ipv4"] = nil
		inputs["ipv6"] = nil
		inputs["mac"] = nil
		inputs["mtu"] = nil
		inputs["netstack"] = nil
		inputs["portgroup"] = nil
	} else {
		inputs["distributedPortGroup"] = args.DistributedPortGroup
		inputs["distributedSwitchPort"] = args.DistributedSwitchPort
		inputs["host"] = args.Host
		inputs["ipv4"] = args.Ipv4
		inputs["ipv6"] = args.Ipv6
		inputs["mac"] = args.Mac
		inputs["mtu"] = args.Mtu
		inputs["netstack"] = args.Netstack
		inputs["portgroup"] = args.Portgroup
	}
	s, err := ctx.RegisterResource("vsphere:index/vnic:Vnic", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Vnic{s: s}, nil
}

// GetVnic gets an existing Vnic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVnic(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VnicState, opts ...pulumi.ResourceOpt) (*Vnic, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["distributedPortGroup"] = state.DistributedPortGroup
		inputs["distributedSwitchPort"] = state.DistributedSwitchPort
		inputs["host"] = state.Host
		inputs["ipv4"] = state.Ipv4
		inputs["ipv6"] = state.Ipv6
		inputs["mac"] = state.Mac
		inputs["mtu"] = state.Mtu
		inputs["netstack"] = state.Netstack
		inputs["portgroup"] = state.Portgroup
	}
	s, err := ctx.ReadResource("vsphere:index/vnic:Vnic", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Vnic{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Vnic) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Vnic) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Key of the distributed portgroup the nic will connect to.
func (r *Vnic) DistributedPortGroup() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["distributedPortGroup"])
}

// UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
func (r *Vnic) DistributedSwitchPort() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["distributedSwitchPort"])
}

// ESX host the interface belongs to
func (r *Vnic) Host() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["host"])
}

// IPv4 settings. Either this or `ipv6` needs to be set. See  ipv4 options below.
func (r *Vnic) Ipv4() pulumi.Output {
	return r.s.State["ipv4"]
}

// IPv6 settings. Either this or `ipv6` needs to be set. See  ipv6 options below.
func (r *Vnic) Ipv6() pulumi.Output {
	return r.s.State["ipv6"]
}

// MAC address of the interface.
func (r *Vnic) Mac() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["mac"])
}

// MTU of the interface.
func (r *Vnic) Mtu() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["mtu"])
}

// TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default: `defaultTcpipStack`)
func (r *Vnic) Netstack() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["netstack"])
}

// Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
func (r *Vnic) Portgroup() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["portgroup"])
}

// Input properties used for looking up and filtering Vnic resources.
type VnicState struct {
	// Key of the distributed portgroup the nic will connect to.
	DistributedPortGroup interface{}
	// UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
	DistributedSwitchPort interface{}
	// ESX host the interface belongs to
	Host interface{}
	// IPv4 settings. Either this or `ipv6` needs to be set. See  ipv4 options below.
	Ipv4 interface{}
	// IPv6 settings. Either this or `ipv6` needs to be set. See  ipv6 options below.
	Ipv6 interface{}
	// MAC address of the interface.
	Mac interface{}
	// MTU of the interface.
	Mtu interface{}
	// TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default: `defaultTcpipStack`)
	Netstack interface{}
	// Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
	Portgroup interface{}
}

// The set of arguments for constructing a Vnic resource.
type VnicArgs struct {
	// Key of the distributed portgroup the nic will connect to.
	DistributedPortGroup interface{}
	// UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
	DistributedSwitchPort interface{}
	// ESX host the interface belongs to
	Host interface{}
	// IPv4 settings. Either this or `ipv6` needs to be set. See  ipv4 options below.
	Ipv4 interface{}
	// IPv6 settings. Either this or `ipv6` needs to be set. See  ipv6 options below.
	Ipv6 interface{}
	// MAC address of the interface.
	Mac interface{}
	// MTU of the interface.
	Mtu interface{}
	// TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default: `defaultTcpipStack`)
	Netstack interface{}
	// Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
	Portgroup interface{}
}

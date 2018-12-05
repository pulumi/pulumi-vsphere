// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `vsphere_vapp_entity` resource can be used to describe the behavior of an
// entity (virtual machine or sub-vApp container) in a vApp container.
// 
// For more information on vSphere vApps, see [this
// page][ref-vsphere-vapp].
// 
// [ref-vsphere-vapp]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vm_admin.doc/GUID-2A95EBB8-1779-40FA-B4FB-4D0845750879.html
type VappEntity struct {
	s *pulumi.ResourceState
}

// NewVappEntity registers a new resource with the given unique name, arguments, and options.
func NewVappEntity(ctx *pulumi.Context,
	name string, args *VappEntityArgs, opts ...pulumi.ResourceOpt) (*VappEntity, error) {
	if args == nil || args.ContainerId == nil {
		return nil, errors.New("missing required argument 'ContainerId'")
	}
	if args == nil || args.TargetId == nil {
		return nil, errors.New("missing required argument 'TargetId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["containerId"] = nil
		inputs["customAttributes"] = nil
		inputs["startAction"] = nil
		inputs["startDelay"] = nil
		inputs["startOrder"] = nil
		inputs["stopAction"] = nil
		inputs["stopDelay"] = nil
		inputs["tags"] = nil
		inputs["targetId"] = nil
		inputs["waitForGuest"] = nil
	} else {
		inputs["containerId"] = args.ContainerId
		inputs["customAttributes"] = args.CustomAttributes
		inputs["startAction"] = args.StartAction
		inputs["startDelay"] = args.StartDelay
		inputs["startOrder"] = args.StartOrder
		inputs["stopAction"] = args.StopAction
		inputs["stopDelay"] = args.StopDelay
		inputs["tags"] = args.Tags
		inputs["targetId"] = args.TargetId
		inputs["waitForGuest"] = args.WaitForGuest
	}
	s, err := ctx.RegisterResource("vsphere:index/vappEntity:VappEntity", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VappEntity{s: s}, nil
}

// GetVappEntity gets an existing VappEntity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVappEntity(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VappEntityState, opts ...pulumi.ResourceOpt) (*VappEntity, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["containerId"] = state.ContainerId
		inputs["customAttributes"] = state.CustomAttributes
		inputs["startAction"] = state.StartAction
		inputs["startDelay"] = state.StartDelay
		inputs["startOrder"] = state.StartOrder
		inputs["stopAction"] = state.StopAction
		inputs["stopDelay"] = state.StopDelay
		inputs["tags"] = state.Tags
		inputs["targetId"] = state.TargetId
		inputs["waitForGuest"] = state.WaitForGuest
	}
	s, err := ctx.ReadResource("vsphere:index/vappEntity:VappEntity", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VappEntity{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VappEntity) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VappEntity) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// [Managed object ID|docs-about-morefs] of the vApp
// container the entity is a member of.
func (r *VappEntity) ContainerId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["containerId"])
}

// A list of custom attributes to set on this resource.
func (r *VappEntity) CustomAttributes() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["customAttributes"])
}

// How to start the entity. Valid settings are none
// or powerOn. If set to none, then the entity does not participate in auto-start.
// Default: powerOn
func (r *VappEntity) StartAction() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["startAction"])
}

// Delay in seconds before continuing with the next
// entity in the order of entities to be started. Default: 120
func (r *VappEntity) StartDelay() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["startDelay"])
}

// Order to start and stop target in vApp. Default: 1
func (r *VappEntity) StartOrder() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["startOrder"])
}

// Defines the stop action for the entity. Can be set
// to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
// does not participate in auto-stop. Default: powerOff
func (r *VappEntity) StopAction() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["stopAction"])
}

// Delay in seconds before continuing with the next
// entity in the order sequence. This is only used if the stopAction is
// guestShutdown. Default: 120
func (r *VappEntity) StopDelay() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["stopDelay"])
}

// A list of tag IDs to apply to this object.
func (r *VappEntity) Tags() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["tags"])
}

// [Managed object ID|docs-about-morefs] of the entity
// to power on or power off. This can be a virtual machine or a vApp.
func (r *VappEntity) TargetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["targetId"])
}

// Determines if the VM should be marked as being
// started when VMware Tools are ready instead of waiting for `start_delay`. This
// property has no effect for vApps. Default: false
func (r *VappEntity) WaitForGuest() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["waitForGuest"])
}

// Input properties used for looking up and filtering VappEntity resources.
type VappEntityState struct {
	// [Managed object ID|docs-about-morefs] of the vApp
	// container the entity is a member of.
	ContainerId interface{}
	// A list of custom attributes to set on this resource.
	CustomAttributes interface{}
	// How to start the entity. Valid settings are none
	// or powerOn. If set to none, then the entity does not participate in auto-start.
	// Default: powerOn
	StartAction interface{}
	// Delay in seconds before continuing with the next
	// entity in the order of entities to be started. Default: 120
	StartDelay interface{}
	// Order to start and stop target in vApp. Default: 1
	StartOrder interface{}
	// Defines the stop action for the entity. Can be set
	// to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
	// does not participate in auto-stop. Default: powerOff
	StopAction interface{}
	// Delay in seconds before continuing with the next
	// entity in the order sequence. This is only used if the stopAction is
	// guestShutdown. Default: 120
	StopDelay interface{}
	// A list of tag IDs to apply to this object.
	Tags interface{}
	// [Managed object ID|docs-about-morefs] of the entity
	// to power on or power off. This can be a virtual machine or a vApp.
	TargetId interface{}
	// Determines if the VM should be marked as being
	// started when VMware Tools are ready instead of waiting for `start_delay`. This
	// property has no effect for vApps. Default: false
	WaitForGuest interface{}
}

// The set of arguments for constructing a VappEntity resource.
type VappEntityArgs struct {
	// [Managed object ID|docs-about-morefs] of the vApp
	// container the entity is a member of.
	ContainerId interface{}
	// A list of custom attributes to set on this resource.
	CustomAttributes interface{}
	// How to start the entity. Valid settings are none
	// or powerOn. If set to none, then the entity does not participate in auto-start.
	// Default: powerOn
	StartAction interface{}
	// Delay in seconds before continuing with the next
	// entity in the order of entities to be started. Default: 120
	StartDelay interface{}
	// Order to start and stop target in vApp. Default: 1
	StartOrder interface{}
	// Defines the stop action for the entity. Can be set
	// to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
	// does not participate in auto-stop. Default: powerOff
	StopAction interface{}
	// Delay in seconds before continuing with the next
	// entity in the order sequence. This is only used if the stopAction is
	// guestShutdown. Default: 120
	StopDelay interface{}
	// A list of tag IDs to apply to this object.
	Tags interface{}
	// [Managed object ID|docs-about-morefs] of the entity
	// to power on or power off. This can be a virtual machine or a vApp.
	TargetId interface{}
	// Determines if the VM should be marked as being
	// started when VMware Tools are ready instead of waiting for `start_delay`. This
	// property has no effect for vApps. Default: false
	WaitForGuest interface{}
}
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `.DrsVmOverride` resource can be used to add a DRS override to a
// cluster for a specific virtual machine. With this resource, one can enable or
// disable DRS and control the automation level for a single virtual machine
// without affecting the rest of the cluster.
// 
// For more information on vSphere clusters and DRS, see [this
// page][ref-vsphere-drs-clusters].
// 
// [ref-vsphere-drs-clusters]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-8ACF3502-5314-469F-8CC9-4A9BD5925BC2.html
// 
// > **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
// 
// > **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/drs_vm_override.html.markdown.
type DrsVmOverride struct {
	s *pulumi.ResourceState
}

// NewDrsVmOverride registers a new resource with the given unique name, arguments, and options.
func NewDrsVmOverride(ctx *pulumi.Context,
	name string, args *DrsVmOverrideArgs, opts ...pulumi.ResourceOpt) (*DrsVmOverride, error) {
	if args == nil || args.ComputeClusterId == nil {
		return nil, errors.New("missing required argument 'ComputeClusterId'")
	}
	if args == nil || args.VirtualMachineId == nil {
		return nil, errors.New("missing required argument 'VirtualMachineId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["computeClusterId"] = nil
		inputs["drsAutomationLevel"] = nil
		inputs["drsEnabled"] = nil
		inputs["virtualMachineId"] = nil
	} else {
		inputs["computeClusterId"] = args.ComputeClusterId
		inputs["drsAutomationLevel"] = args.DrsAutomationLevel
		inputs["drsEnabled"] = args.DrsEnabled
		inputs["virtualMachineId"] = args.VirtualMachineId
	}
	s, err := ctx.RegisterResource("vsphere:index/drsVmOverride:DrsVmOverride", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DrsVmOverride{s: s}, nil
}

// GetDrsVmOverride gets an existing DrsVmOverride resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDrsVmOverride(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DrsVmOverrideState, opts ...pulumi.ResourceOpt) (*DrsVmOverride, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["computeClusterId"] = state.ComputeClusterId
		inputs["drsAutomationLevel"] = state.DrsAutomationLevel
		inputs["drsEnabled"] = state.DrsEnabled
		inputs["virtualMachineId"] = state.VirtualMachineId
	}
	s, err := ctx.ReadResource("vsphere:index/drsVmOverride:DrsVmOverride", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DrsVmOverride{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DrsVmOverride) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DrsVmOverride) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The [managed object reference
// ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
// resource if changed.
func (r *DrsVmOverride) ComputeClusterId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["computeClusterId"])
}

// Overrides the automation level for this virtual
// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
// `fullyAutomated`. Default: `manual`.
func (r *DrsVmOverride) DrsAutomationLevel() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["drsAutomationLevel"])
}

// Overrides the default DRS setting for this virtual
// machine. Can be either `true` or `false`. Default: `false`.
func (r *DrsVmOverride) DrsEnabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["drsEnabled"])
}

// The UUID of the virtual machine to create
// the override for.  Forces a new resource if changed.
func (r *DrsVmOverride) VirtualMachineId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["virtualMachineId"])
}

// Input properties used for looking up and filtering DrsVmOverride resources.
type DrsVmOverrideState struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId interface{}
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	DrsAutomationLevel interface{}
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled interface{}
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId interface{}
}

// The set of arguments for constructing a DrsVmOverride resource.
type DrsVmOverrideArgs struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId interface{}
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	DrsAutomationLevel interface{}
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled interface{}
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId interface{}
}

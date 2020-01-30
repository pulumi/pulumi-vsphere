// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `.DpmHostOverride` resource can be used to add a DPM override to a
// cluster for a particular host. This allows you to control the power management
// settings for individual hosts in the cluster while leaving any unspecified ones
// at the default power management settings.
// 
// For more information on DPM within vSphere clusters, see [this
// page][ref-vsphere-cluster-dpm].
// 
// [ref-vsphere-cluster-dpm]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-5E5E349A-4644-4C9C-B434-1C0243EBDC80.html
// 
// > **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
// 
// > **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/dpm_host_override.html.markdown.
type DpmHostOverride struct {
	s *pulumi.ResourceState
}

// NewDpmHostOverride registers a new resource with the given unique name, arguments, and options.
func NewDpmHostOverride(ctx *pulumi.Context,
	name string, args *DpmHostOverrideArgs, opts ...pulumi.ResourceOpt) (*DpmHostOverride, error) {
	if args == nil || args.ComputeClusterId == nil {
		return nil, errors.New("missing required argument 'ComputeClusterId'")
	}
	if args == nil || args.HostSystemId == nil {
		return nil, errors.New("missing required argument 'HostSystemId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["computeClusterId"] = nil
		inputs["dpmAutomationLevel"] = nil
		inputs["dpmEnabled"] = nil
		inputs["hostSystemId"] = nil
	} else {
		inputs["computeClusterId"] = args.ComputeClusterId
		inputs["dpmAutomationLevel"] = args.DpmAutomationLevel
		inputs["dpmEnabled"] = args.DpmEnabled
		inputs["hostSystemId"] = args.HostSystemId
	}
	s, err := ctx.RegisterResource("vsphere:index/dpmHostOverride:DpmHostOverride", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DpmHostOverride{s: s}, nil
}

// GetDpmHostOverride gets an existing DpmHostOverride resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDpmHostOverride(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DpmHostOverrideState, opts ...pulumi.ResourceOpt) (*DpmHostOverride, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["computeClusterId"] = state.ComputeClusterId
		inputs["dpmAutomationLevel"] = state.DpmAutomationLevel
		inputs["dpmEnabled"] = state.DpmEnabled
		inputs["hostSystemId"] = state.HostSystemId
	}
	s, err := ctx.ReadResource("vsphere:index/dpmHostOverride:DpmHostOverride", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DpmHostOverride{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DpmHostOverride) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DpmHostOverride) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The [managed object reference
// ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
// resource if changed.
func (r *DpmHostOverride) ComputeClusterId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["computeClusterId"])
}

// The automation level for host power
// operations on this host. Can be one of `manual` or `automated`. Default:
// `manual`.
func (r *DpmHostOverride) DpmAutomationLevel() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dpmAutomationLevel"])
}

// Enable DPM support for this host. Default:
// `false`.
func (r *DpmHostOverride) DpmEnabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["dpmEnabled"])
}

// The managed object ID of the host.
func (r *DpmHostOverride) HostSystemId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["hostSystemId"])
}

// Input properties used for looking up and filtering DpmHostOverride resources.
type DpmHostOverrideState struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId interface{}
	// The automation level for host power
	// operations on this host. Can be one of `manual` or `automated`. Default:
	// `manual`.
	DpmAutomationLevel interface{}
	// Enable DPM support for this host. Default:
	// `false`.
	DpmEnabled interface{}
	// The managed object ID of the host.
	HostSystemId interface{}
}

// The set of arguments for constructing a DpmHostOverride resource.
type DpmHostOverrideArgs struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId interface{}
	// The automation level for host power
	// operations on this host. Can be one of `manual` or `automated`. Default:
	// `manual`.
	DpmAutomationLevel interface{}
	// Enable DPM support for this host. Default:
	// `false`.
	DpmEnabled interface{}
	// The managed object ID of the host.
	HostSystemId interface{}
}

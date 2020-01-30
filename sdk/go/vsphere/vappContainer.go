// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `.VappContainer` resource can be used to create and manage
// vApps.
//
// For more information on vSphere vApps, see [this
// page][ref-vsphere-vapp].
//
// [ref-vsphere-vapp]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vm_admin.doc/GUID-2A95EBB8-1779-40FA-B4FB-4D0845750879.html
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/vapp_container.html.markdown.
type VappContainer struct {
	s *pulumi.ResourceState
}

// NewVappContainer registers a new resource with the given unique name, arguments, and options.
func NewVappContainer(ctx *pulumi.Context,
	name string, args *VappContainerArgs, opts ...pulumi.ResourceOpt) (*VappContainer, error) {
	if args == nil || args.ParentResourcePoolId == nil {
		return nil, errors.New("missing required argument 'ParentResourcePoolId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cpuExpandable"] = nil
		inputs["cpuLimit"] = nil
		inputs["cpuReservation"] = nil
		inputs["cpuShareLevel"] = nil
		inputs["cpuShares"] = nil
		inputs["customAttributes"] = nil
		inputs["memoryExpandable"] = nil
		inputs["memoryLimit"] = nil
		inputs["memoryReservation"] = nil
		inputs["memoryShareLevel"] = nil
		inputs["memoryShares"] = nil
		inputs["name"] = nil
		inputs["parentFolderId"] = nil
		inputs["parentResourcePoolId"] = nil
		inputs["tags"] = nil
	} else {
		inputs["cpuExpandable"] = args.CpuExpandable
		inputs["cpuLimit"] = args.CpuLimit
		inputs["cpuReservation"] = args.CpuReservation
		inputs["cpuShareLevel"] = args.CpuShareLevel
		inputs["cpuShares"] = args.CpuShares
		inputs["customAttributes"] = args.CustomAttributes
		inputs["memoryExpandable"] = args.MemoryExpandable
		inputs["memoryLimit"] = args.MemoryLimit
		inputs["memoryReservation"] = args.MemoryReservation
		inputs["memoryShareLevel"] = args.MemoryShareLevel
		inputs["memoryShares"] = args.MemoryShares
		inputs["name"] = args.Name
		inputs["parentFolderId"] = args.ParentFolderId
		inputs["parentResourcePoolId"] = args.ParentResourcePoolId
		inputs["tags"] = args.Tags
	}
	s, err := ctx.RegisterResource("vsphere:index/vappContainer:VappContainer", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VappContainer{s: s}, nil
}

// GetVappContainer gets an existing VappContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVappContainer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VappContainerState, opts ...pulumi.ResourceOpt) (*VappContainer, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["cpuExpandable"] = state.CpuExpandable
		inputs["cpuLimit"] = state.CpuLimit
		inputs["cpuReservation"] = state.CpuReservation
		inputs["cpuShareLevel"] = state.CpuShareLevel
		inputs["cpuShares"] = state.CpuShares
		inputs["customAttributes"] = state.CustomAttributes
		inputs["memoryExpandable"] = state.MemoryExpandable
		inputs["memoryLimit"] = state.MemoryLimit
		inputs["memoryReservation"] = state.MemoryReservation
		inputs["memoryShareLevel"] = state.MemoryShareLevel
		inputs["memoryShares"] = state.MemoryShares
		inputs["name"] = state.Name
		inputs["parentFolderId"] = state.ParentFolderId
		inputs["parentResourcePoolId"] = state.ParentResourcePoolId
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("vsphere:index/vappContainer:VappContainer", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VappContainer{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VappContainer) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VappContainer) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Determines if the reservation on a vApp
// container can grow beyond the specified value if the parent resource pool has
// unreserved resources. Default: `true`
func (r *VappContainer) CpuExpandable() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["cpuExpandable"])
}

// The CPU utilization of a vApp container will not
// exceed this limit, even if there are available resources. Set to `-1` for
// unlimited.
// Default: `-1`
func (r *VappContainer) CpuLimit() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["cpuLimit"])
}

// Amount of CPU (MHz) that is guaranteed
// available to the vApp container. Default: `0`
func (r *VappContainer) CpuReservation() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["cpuReservation"])
}

// The CPU allocation level. The level is a
// simplified view of shares. Levels map to a pre-determined set of numeric
// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
// `low`, `normal`, or `high` are specified values in `cpuShares` will be
// ignored.  Default: `normal`
func (r *VappContainer) CpuShareLevel() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["cpuShareLevel"])
}

// The number of shares allocated for CPU. Used to
// determine resource allocation in case of resource contention. If this is set,
// `cpuShareLevel` must be `custom`.
func (r *VappContainer) CpuShares() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["cpuShares"])
}

// A list of custom attributes to set on this resource.
func (r *VappContainer) CustomAttributes() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["customAttributes"])
}

// Determines if the reservation on a vApp
// container can grow beyond the specified value if the parent resource pool has
// unreserved resources. Default: `true`
func (r *VappContainer) MemoryExpandable() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["memoryExpandable"])
}

// The CPU utilization of a vApp container will not
// exceed this limit, even if there are available resources. Set to `-1` for
// unlimited.
// Default: `-1`
func (r *VappContainer) MemoryLimit() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["memoryLimit"])
}

// Amount of CPU (MHz) that is guaranteed
// available to the vApp container. Default: `0`
func (r *VappContainer) MemoryReservation() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["memoryReservation"])
}

// The CPU allocation level. The level is a
// simplified view of shares. Levels map to a pre-determined set of numeric
// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
// `low`, `normal`, or `high` are specified values in `memoryShares` will be
// ignored.  Default: `normal`
func (r *VappContainer) MemoryShareLevel() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["memoryShareLevel"])
}

// The number of shares allocated for CPU. Used to
// determine resource allocation in case of resource contention. If this is set,
// `memoryShareLevel` must be `custom`.
func (r *VappContainer) MemoryShares() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["memoryShares"])
}

// The name of the vApp container.
func (r *VappContainer) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The [managed object ID][docs-about-morefs] of
// the vApp container's parent folder.
func (r *VappContainer) ParentFolderId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["parentFolderId"])
}

// The [managed object ID][docs-about-morefs]
// of the parent resource pool. This can be the root resource pool for a cluster
// or standalone host, or a resource pool itself. When moving a vApp container
// from one parent resource pool to another, both must share a common root
// resource pool or the move will fail.
func (r *VappContainer) ParentResourcePoolId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["parentResourcePoolId"])
}

// The IDs of any tags to attach to this resource. See
// [here][docs-applying-tags] for a reference on how to apply tags.
func (r *VappContainer) Tags() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering VappContainer resources.
type VappContainerState struct {
	// Determines if the reservation on a vApp
	// container can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable interface{}
	// The CPU utilization of a vApp container will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited.
	// Default: `-1`
	CpuLimit interface{}
	// Amount of CPU (MHz) that is guaranteed
	// available to the vApp container. Default: `0`
	CpuReservation interface{}
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `cpuShares` will be
	// ignored.  Default: `normal`
	CpuShareLevel interface{}
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `cpuShareLevel` must be `custom`.
	CpuShares interface{}
	// A list of custom attributes to set on this resource.
	CustomAttributes interface{}
	// Determines if the reservation on a vApp
	// container can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable interface{}
	// The CPU utilization of a vApp container will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited.
	// Default: `-1`
	MemoryLimit interface{}
	// Amount of CPU (MHz) that is guaranteed
	// available to the vApp container. Default: `0`
	MemoryReservation interface{}
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `memoryShares` will be
	// ignored.  Default: `normal`
	MemoryShareLevel interface{}
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `memoryShareLevel` must be `custom`.
	MemoryShares interface{}
	// The name of the vApp container.
	Name interface{}
	// The [managed object ID][docs-about-morefs] of
	// the vApp container's parent folder.
	ParentFolderId interface{}
	// The [managed object ID][docs-about-morefs]
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a vApp container
	// from one parent resource pool to another, both must share a common root
	// resource pool or the move will fail.
	ParentResourcePoolId interface{}
	// The IDs of any tags to attach to this resource. See
	// [here][docs-applying-tags] for a reference on how to apply tags.
	Tags interface{}
}

// The set of arguments for constructing a VappContainer resource.
type VappContainerArgs struct {
	// Determines if the reservation on a vApp
	// container can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable interface{}
	// The CPU utilization of a vApp container will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited.
	// Default: `-1`
	CpuLimit interface{}
	// Amount of CPU (MHz) that is guaranteed
	// available to the vApp container. Default: `0`
	CpuReservation interface{}
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `cpuShares` will be
	// ignored.  Default: `normal`
	CpuShareLevel interface{}
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `cpuShareLevel` must be `custom`.
	CpuShares interface{}
	// A list of custom attributes to set on this resource.
	CustomAttributes interface{}
	// Determines if the reservation on a vApp
	// container can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable interface{}
	// The CPU utilization of a vApp container will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited.
	// Default: `-1`
	MemoryLimit interface{}
	// Amount of CPU (MHz) that is guaranteed
	// available to the vApp container. Default: `0`
	MemoryReservation interface{}
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `memoryShares` will be
	// ignored.  Default: `normal`
	MemoryShareLevel interface{}
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `memoryShareLevel` must be `custom`.
	MemoryShares interface{}
	// The name of the vApp container.
	Name interface{}
	// The [managed object ID][docs-about-morefs] of
	// the vApp container's parent folder.
	ParentFolderId interface{}
	// The [managed object ID][docs-about-morefs]
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a vApp container
	// from one parent resource pool to another, both must share a common root
	// resource pool or the move will fail.
	ParentResourcePoolId interface{}
	// The IDs of any tags to attach to this resource. See
	// [here][docs-applying-tags] for a reference on how to apply tags.
	Tags interface{}
}

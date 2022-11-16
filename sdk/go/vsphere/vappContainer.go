// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VappContainer struct {
	pulumi.CustomResourceState

	// Determines if the reservation on a vApp
	// container can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable pulumi.BoolPtrOutput `pulumi:"cpuExpandable"`
	// The CPU utilization of a vApp container will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited.
	// Default: `-1`
	CpuLimit pulumi.IntPtrOutput `pulumi:"cpuLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the vApp container. Default: `0`
	CpuReservation pulumi.IntPtrOutput `pulumi:"cpuReservation"`
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `cpuShares` will be
	// ignored.  Default: `normal`
	CpuShareLevel pulumi.StringPtrOutput `pulumi:"cpuShareLevel"`
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `cpuShareLevel` must be `custom`.
	CpuShares pulumi.IntOutput `pulumi:"cpuShares"`
	// A list of custom attributes to set on this resource.
	CustomAttributes pulumi.StringMapOutput `pulumi:"customAttributes"`
	// Determines if the reservation on a vApp
	// container can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable pulumi.BoolPtrOutput `pulumi:"memoryExpandable"`
	// The CPU utilization of a vApp container will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited. Default: `-1`
	MemoryLimit pulumi.IntPtrOutput `pulumi:"memoryLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the vApp container. Default: `0`
	MemoryReservation pulumi.IntPtrOutput `pulumi:"memoryReservation"`
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `memoryShares` will be
	// ignored.  Default: `normal`
	MemoryShareLevel pulumi.StringPtrOutput `pulumi:"memoryShareLevel"`
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `memoryShareLevel` must be `custom`.
	MemoryShares pulumi.IntOutput `pulumi:"memoryShares"`
	// The name of the vApp container.
	Name pulumi.StringOutput `pulumi:"name"`
	// The managed object ID of
	// the vApp container's parent folder.
	ParentFolderId pulumi.StringPtrOutput `pulumi:"parentFolderId"`
	// The managed object ID
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a vApp container
	// from one parent resource pool to another, both must share a common root
	// resource pool or the move will fail.
	ParentResourcePoolId pulumi.StringOutput `pulumi:"parentResourcePoolId"`
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewVappContainer registers a new resource with the given unique name, arguments, and options.
func NewVappContainer(ctx *pulumi.Context,
	name string, args *VappContainerArgs, opts ...pulumi.ResourceOption) (*VappContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentResourcePoolId == nil {
		return nil, errors.New("invalid value for required argument 'ParentResourcePoolId'")
	}
	var resource VappContainer
	err := ctx.RegisterResource("vsphere:index/vappContainer:VappContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVappContainer gets an existing VappContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVappContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VappContainerState, opts ...pulumi.ResourceOption) (*VappContainer, error) {
	var resource VappContainer
	err := ctx.ReadResource("vsphere:index/vappContainer:VappContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VappContainer resources.
type vappContainerState struct {
	// Determines if the reservation on a vApp
	// container can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable *bool `pulumi:"cpuExpandable"`
	// The CPU utilization of a vApp container will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited.
	// Default: `-1`
	CpuLimit *int `pulumi:"cpuLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the vApp container. Default: `0`
	CpuReservation *int `pulumi:"cpuReservation"`
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `cpuShares` will be
	// ignored.  Default: `normal`
	CpuShareLevel *string `pulumi:"cpuShareLevel"`
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `cpuShareLevel` must be `custom`.
	CpuShares *int `pulumi:"cpuShares"`
	// A list of custom attributes to set on this resource.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// Determines if the reservation on a vApp
	// container can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable *bool `pulumi:"memoryExpandable"`
	// The CPU utilization of a vApp container will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited. Default: `-1`
	MemoryLimit *int `pulumi:"memoryLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the vApp container. Default: `0`
	MemoryReservation *int `pulumi:"memoryReservation"`
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `memoryShares` will be
	// ignored.  Default: `normal`
	MemoryShareLevel *string `pulumi:"memoryShareLevel"`
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `memoryShareLevel` must be `custom`.
	MemoryShares *int `pulumi:"memoryShares"`
	// The name of the vApp container.
	Name *string `pulumi:"name"`
	// The managed object ID of
	// the vApp container's parent folder.
	ParentFolderId *string `pulumi:"parentFolderId"`
	// The managed object ID
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a vApp container
	// from one parent resource pool to another, both must share a common root
	// resource pool or the move will fail.
	ParentResourcePoolId *string `pulumi:"parentResourcePoolId"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
}

type VappContainerState struct {
	// Determines if the reservation on a vApp
	// container can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable pulumi.BoolPtrInput
	// The CPU utilization of a vApp container will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited.
	// Default: `-1`
	CpuLimit pulumi.IntPtrInput
	// Amount of CPU (MHz) that is guaranteed
	// available to the vApp container. Default: `0`
	CpuReservation pulumi.IntPtrInput
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `cpuShares` will be
	// ignored.  Default: `normal`
	CpuShareLevel pulumi.StringPtrInput
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `cpuShareLevel` must be `custom`.
	CpuShares pulumi.IntPtrInput
	// A list of custom attributes to set on this resource.
	CustomAttributes pulumi.StringMapInput
	// Determines if the reservation on a vApp
	// container can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable pulumi.BoolPtrInput
	// The CPU utilization of a vApp container will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited. Default: `-1`
	MemoryLimit pulumi.IntPtrInput
	// Amount of CPU (MHz) that is guaranteed
	// available to the vApp container. Default: `0`
	MemoryReservation pulumi.IntPtrInput
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `memoryShares` will be
	// ignored.  Default: `normal`
	MemoryShareLevel pulumi.StringPtrInput
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `memoryShareLevel` must be `custom`.
	MemoryShares pulumi.IntPtrInput
	// The name of the vApp container.
	Name pulumi.StringPtrInput
	// The managed object ID of
	// the vApp container's parent folder.
	ParentFolderId pulumi.StringPtrInput
	// The managed object ID
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a vApp container
	// from one parent resource pool to another, both must share a common root
	// resource pool or the move will fail.
	ParentResourcePoolId pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
}

func (VappContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*vappContainerState)(nil)).Elem()
}

type vappContainerArgs struct {
	// Determines if the reservation on a vApp
	// container can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable *bool `pulumi:"cpuExpandable"`
	// The CPU utilization of a vApp container will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited.
	// Default: `-1`
	CpuLimit *int `pulumi:"cpuLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the vApp container. Default: `0`
	CpuReservation *int `pulumi:"cpuReservation"`
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `cpuShares` will be
	// ignored.  Default: `normal`
	CpuShareLevel *string `pulumi:"cpuShareLevel"`
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `cpuShareLevel` must be `custom`.
	CpuShares *int `pulumi:"cpuShares"`
	// A list of custom attributes to set on this resource.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// Determines if the reservation on a vApp
	// container can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable *bool `pulumi:"memoryExpandable"`
	// The CPU utilization of a vApp container will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited. Default: `-1`
	MemoryLimit *int `pulumi:"memoryLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the vApp container. Default: `0`
	MemoryReservation *int `pulumi:"memoryReservation"`
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `memoryShares` will be
	// ignored.  Default: `normal`
	MemoryShareLevel *string `pulumi:"memoryShareLevel"`
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `memoryShareLevel` must be `custom`.
	MemoryShares *int `pulumi:"memoryShares"`
	// The name of the vApp container.
	Name *string `pulumi:"name"`
	// The managed object ID of
	// the vApp container's parent folder.
	ParentFolderId *string `pulumi:"parentFolderId"`
	// The managed object ID
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a vApp container
	// from one parent resource pool to another, both must share a common root
	// resource pool or the move will fail.
	ParentResourcePoolId string `pulumi:"parentResourcePoolId"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a VappContainer resource.
type VappContainerArgs struct {
	// Determines if the reservation on a vApp
	// container can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable pulumi.BoolPtrInput
	// The CPU utilization of a vApp container will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited.
	// Default: `-1`
	CpuLimit pulumi.IntPtrInput
	// Amount of CPU (MHz) that is guaranteed
	// available to the vApp container. Default: `0`
	CpuReservation pulumi.IntPtrInput
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `cpuShares` will be
	// ignored.  Default: `normal`
	CpuShareLevel pulumi.StringPtrInput
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `cpuShareLevel` must be `custom`.
	CpuShares pulumi.IntPtrInput
	// A list of custom attributes to set on this resource.
	CustomAttributes pulumi.StringMapInput
	// Determines if the reservation on a vApp
	// container can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable pulumi.BoolPtrInput
	// The CPU utilization of a vApp container will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited. Default: `-1`
	MemoryLimit pulumi.IntPtrInput
	// Amount of CPU (MHz) that is guaranteed
	// available to the vApp container. Default: `0`
	MemoryReservation pulumi.IntPtrInput
	// The CPU allocation level. The level is a
	// simplified view of shares. Levels map to a pre-determined set of numeric
	// values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
	// `low`, `normal`, or `high` are specified values in `memoryShares` will be
	// ignored.  Default: `normal`
	MemoryShareLevel pulumi.StringPtrInput
	// The number of shares allocated for CPU. Used to
	// determine resource allocation in case of resource contention. If this is set,
	// `memoryShareLevel` must be `custom`.
	MemoryShares pulumi.IntPtrInput
	// The name of the vApp container.
	Name pulumi.StringPtrInput
	// The managed object ID of
	// the vApp container's parent folder.
	ParentFolderId pulumi.StringPtrInput
	// The managed object ID
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a vApp container
	// from one parent resource pool to another, both must share a common root
	// resource pool or the move will fail.
	ParentResourcePoolId pulumi.StringInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
}

func (VappContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vappContainerArgs)(nil)).Elem()
}

type VappContainerInput interface {
	pulumi.Input

	ToVappContainerOutput() VappContainerOutput
	ToVappContainerOutputWithContext(ctx context.Context) VappContainerOutput
}

func (*VappContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**VappContainer)(nil)).Elem()
}

func (i *VappContainer) ToVappContainerOutput() VappContainerOutput {
	return i.ToVappContainerOutputWithContext(context.Background())
}

func (i *VappContainer) ToVappContainerOutputWithContext(ctx context.Context) VappContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappContainerOutput)
}

// VappContainerArrayInput is an input type that accepts VappContainerArray and VappContainerArrayOutput values.
// You can construct a concrete instance of `VappContainerArrayInput` via:
//
//	VappContainerArray{ VappContainerArgs{...} }
type VappContainerArrayInput interface {
	pulumi.Input

	ToVappContainerArrayOutput() VappContainerArrayOutput
	ToVappContainerArrayOutputWithContext(context.Context) VappContainerArrayOutput
}

type VappContainerArray []VappContainerInput

func (VappContainerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VappContainer)(nil)).Elem()
}

func (i VappContainerArray) ToVappContainerArrayOutput() VappContainerArrayOutput {
	return i.ToVappContainerArrayOutputWithContext(context.Background())
}

func (i VappContainerArray) ToVappContainerArrayOutputWithContext(ctx context.Context) VappContainerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappContainerArrayOutput)
}

// VappContainerMapInput is an input type that accepts VappContainerMap and VappContainerMapOutput values.
// You can construct a concrete instance of `VappContainerMapInput` via:
//
//	VappContainerMap{ "key": VappContainerArgs{...} }
type VappContainerMapInput interface {
	pulumi.Input

	ToVappContainerMapOutput() VappContainerMapOutput
	ToVappContainerMapOutputWithContext(context.Context) VappContainerMapOutput
}

type VappContainerMap map[string]VappContainerInput

func (VappContainerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VappContainer)(nil)).Elem()
}

func (i VappContainerMap) ToVappContainerMapOutput() VappContainerMapOutput {
	return i.ToVappContainerMapOutputWithContext(context.Background())
}

func (i VappContainerMap) ToVappContainerMapOutputWithContext(ctx context.Context) VappContainerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappContainerMapOutput)
}

type VappContainerOutput struct{ *pulumi.OutputState }

func (VappContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VappContainer)(nil)).Elem()
}

func (o VappContainerOutput) ToVappContainerOutput() VappContainerOutput {
	return o
}

func (o VappContainerOutput) ToVappContainerOutputWithContext(ctx context.Context) VappContainerOutput {
	return o
}

type VappContainerArrayOutput struct{ *pulumi.OutputState }

func (VappContainerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VappContainer)(nil)).Elem()
}

func (o VappContainerArrayOutput) ToVappContainerArrayOutput() VappContainerArrayOutput {
	return o
}

func (o VappContainerArrayOutput) ToVappContainerArrayOutputWithContext(ctx context.Context) VappContainerArrayOutput {
	return o
}

func (o VappContainerArrayOutput) Index(i pulumi.IntInput) VappContainerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VappContainer {
		return vs[0].([]*VappContainer)[vs[1].(int)]
	}).(VappContainerOutput)
}

type VappContainerMapOutput struct{ *pulumi.OutputState }

func (VappContainerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VappContainer)(nil)).Elem()
}

func (o VappContainerMapOutput) ToVappContainerMapOutput() VappContainerMapOutput {
	return o
}

func (o VappContainerMapOutput) ToVappContainerMapOutputWithContext(ctx context.Context) VappContainerMapOutput {
	return o
}

func (o VappContainerMapOutput) MapIndex(k pulumi.StringInput) VappContainerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VappContainer {
		return vs[0].(map[string]*VappContainer)[vs[1].(string)]
	}).(VappContainerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VappContainerInput)(nil)).Elem(), &VappContainer{})
	pulumi.RegisterInputType(reflect.TypeOf((*VappContainerArrayInput)(nil)).Elem(), VappContainerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VappContainerMapInput)(nil)).Elem(), VappContainerMap{})
	pulumi.RegisterOutputType(VappContainerOutput{})
	pulumi.RegisterOutputType(VappContainerArrayOutput{})
	pulumi.RegisterOutputType(VappContainerMapOutput{})
}

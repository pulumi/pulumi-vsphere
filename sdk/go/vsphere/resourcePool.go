// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ResourcePool` resource can be used to create and manage
// resource pools on DRS-enabled vSphere clusters or standalone ESXi hosts.
//
// For more information on vSphere resource pools, please refer to the
// [product documentation][ref-vsphere-resource_pools].
//
// [ref-vsphere-resource_pools]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.resmgmt.doc/GUID-60077B40-66FF-4625-934A-641703ED7601.html
//
// ## Example Usage
//
// The following example sets up a resource pool in an existing compute cluster
// with the default settings for CPU and memory reservations, shares, and limits.
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
// 		datacenter, err := vsphere.LookupDatacenter(ctx, &GetDatacenterArgs{
// 			Name: pulumi.StringRef("dc-01"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		computeCluster, err := vsphere.LookupComputeCluster(ctx, &GetComputeClusterArgs{
// 			Name:         "cluster-01",
// 			DatacenterId: pulumi.StringRef(datacenter.Id),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vsphere.NewResourcePool(ctx, "resourcePool", &vsphere.ResourcePoolArgs{
// 			ParentResourcePoolId: pulumi.String(computeCluster.ResourcePoolId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// A virtual machine resource could be targeted to use the default resource pool
// of the cluster using the following:
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
// 		_, err := vsphere.NewVirtualMachine(ctx, "vm", &vsphere.VirtualMachineArgs{
// 			ResourcePoolId: pulumi.Any(data.Vsphere_compute_cluster.Cluster.Resource_pool_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// The following example sets up a parent resource pool in an existing compute cluster
// with a child resource pool nested below. Each resource pool is configured with
// the default settings for CPU and memory reservations, shares, and limits.
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
// 		datacenter, err := vsphere.LookupDatacenter(ctx, &GetDatacenterArgs{
// 			Name: pulumi.StringRef("dc-01"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		computeCluster, err := vsphere.LookupComputeCluster(ctx, &GetComputeClusterArgs{
// 			Name:         "cluster-01",
// 			DatacenterId: pulumi.StringRef(datacenter.Id),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		resourcePoolParent, err := vsphere.NewResourcePool(ctx, "resourcePoolParent", &vsphere.ResourcePoolArgs{
// 			ParentResourcePoolId: pulumi.String(computeCluster.ResourcePoolId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vsphere.NewResourcePool(ctx, "resourcePoolChild", &vsphere.ResourcePoolArgs{
// 			ParentResourcePoolId: resourcePoolParent.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Importing
// ### Settings that Require vSphere 7.0 or higher
//
// These settings require vSphere 7.0 or higher:
//
// * `scaleDescendantsShares`
type ResourcePool struct {
	pulumi.CustomResourceState

	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable pulumi.BoolPtrOutput `pulumi:"cpuExpandable"`
	// The CPU utilization of a resource pool will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited. Default: `-1`
	CpuLimit pulumi.IntPtrOutput `pulumi:"cpuLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
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
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable pulumi.BoolPtrOutput `pulumi:"memoryExpandable"`
	// The CPU utilization of a resource pool will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited. Default: `-1`
	MemoryLimit pulumi.IntPtrOutput `pulumi:"memoryLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
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
	// The name of the resource pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// The managed object ID
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a resource pool
	// from one parent resource pool to another, both must share a common root
	// resource pool.
	ParentResourcePoolId pulumi.StringOutput `pulumi:"parentResourcePoolId"`
	// Determines if the shares of all
	// descendants of the resource pool are scaled up or down when the shares
	// of the resource pool are scaled up or down. Can be one of `disabled` or
	// `scaleCpuAndMemoryShares`. Default: `disabled`.
	ScaleDescendantsShares pulumi.StringPtrOutput `pulumi:"scaleDescendantsShares"`
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewResourcePool registers a new resource with the given unique name, arguments, and options.
func NewResourcePool(ctx *pulumi.Context,
	name string, args *ResourcePoolArgs, opts ...pulumi.ResourceOption) (*ResourcePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentResourcePoolId == nil {
		return nil, errors.New("invalid value for required argument 'ParentResourcePoolId'")
	}
	var resource ResourcePool
	err := ctx.RegisterResource("vsphere:index/resourcePool:ResourcePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcePool gets an existing ResourcePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcePoolState, opts ...pulumi.ResourceOption) (*ResourcePool, error) {
	var resource ResourcePool
	err := ctx.ReadResource("vsphere:index/resourcePool:ResourcePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcePool resources.
type resourcePoolState struct {
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable *bool `pulumi:"cpuExpandable"`
	// The CPU utilization of a resource pool will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited. Default: `-1`
	CpuLimit *int `pulumi:"cpuLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
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
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable *bool `pulumi:"memoryExpandable"`
	// The CPU utilization of a resource pool will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited. Default: `-1`
	MemoryLimit *int `pulumi:"memoryLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
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
	// The name of the resource pool.
	Name *string `pulumi:"name"`
	// The managed object ID
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a resource pool
	// from one parent resource pool to another, both must share a common root
	// resource pool.
	ParentResourcePoolId *string `pulumi:"parentResourcePoolId"`
	// Determines if the shares of all
	// descendants of the resource pool are scaled up or down when the shares
	// of the resource pool are scaled up or down. Can be one of `disabled` or
	// `scaleCpuAndMemoryShares`. Default: `disabled`.
	ScaleDescendantsShares *string `pulumi:"scaleDescendantsShares"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
}

type ResourcePoolState struct {
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable pulumi.BoolPtrInput
	// The CPU utilization of a resource pool will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited. Default: `-1`
	CpuLimit pulumi.IntPtrInput
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
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
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable pulumi.BoolPtrInput
	// The CPU utilization of a resource pool will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited. Default: `-1`
	MemoryLimit pulumi.IntPtrInput
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
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
	// The name of the resource pool.
	Name pulumi.StringPtrInput
	// The managed object ID
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a resource pool
	// from one parent resource pool to another, both must share a common root
	// resource pool.
	ParentResourcePoolId pulumi.StringPtrInput
	// Determines if the shares of all
	// descendants of the resource pool are scaled up or down when the shares
	// of the resource pool are scaled up or down. Can be one of `disabled` or
	// `scaleCpuAndMemoryShares`. Default: `disabled`.
	ScaleDescendantsShares pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
}

func (ResourcePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePoolState)(nil)).Elem()
}

type resourcePoolArgs struct {
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable *bool `pulumi:"cpuExpandable"`
	// The CPU utilization of a resource pool will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited. Default: `-1`
	CpuLimit *int `pulumi:"cpuLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
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
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable *bool `pulumi:"memoryExpandable"`
	// The CPU utilization of a resource pool will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited. Default: `-1`
	MemoryLimit *int `pulumi:"memoryLimit"`
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
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
	// The name of the resource pool.
	Name *string `pulumi:"name"`
	// The managed object ID
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a resource pool
	// from one parent resource pool to another, both must share a common root
	// resource pool.
	ParentResourcePoolId string `pulumi:"parentResourcePoolId"`
	// Determines if the shares of all
	// descendants of the resource pool are scaled up or down when the shares
	// of the resource pool are scaled up or down. Can be one of `disabled` or
	// `scaleCpuAndMemoryShares`. Default: `disabled`.
	ScaleDescendantsShares *string `pulumi:"scaleDescendantsShares"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a ResourcePool resource.
type ResourcePoolArgs struct {
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	CpuExpandable pulumi.BoolPtrInput
	// The CPU utilization of a resource pool will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited. Default: `-1`
	CpuLimit pulumi.IntPtrInput
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
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
	// Determines if the reservation on a resource
	// pool can grow beyond the specified value if the parent resource pool has
	// unreserved resources. Default: `true`
	MemoryExpandable pulumi.BoolPtrInput
	// The CPU utilization of a resource pool will not
	// exceed this limit, even if there are available resources. Set to `-1` for
	// unlimited. Default: `-1`
	MemoryLimit pulumi.IntPtrInput
	// Amount of CPU (MHz) that is guaranteed
	// available to the resource pool. Default: `0`
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
	// The name of the resource pool.
	Name pulumi.StringPtrInput
	// The managed object ID
	// of the parent resource pool. This can be the root resource pool for a cluster
	// or standalone host, or a resource pool itself. When moving a resource pool
	// from one parent resource pool to another, both must share a common root
	// resource pool.
	ParentResourcePoolId pulumi.StringInput
	// Determines if the shares of all
	// descendants of the resource pool are scaled up or down when the shares
	// of the resource pool are scaled up or down. Can be one of `disabled` or
	// `scaleCpuAndMemoryShares`. Default: `disabled`.
	ScaleDescendantsShares pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
}

func (ResourcePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePoolArgs)(nil)).Elem()
}

type ResourcePoolInput interface {
	pulumi.Input

	ToResourcePoolOutput() ResourcePoolOutput
	ToResourcePoolOutputWithContext(ctx context.Context) ResourcePoolOutput
}

func (*ResourcePool) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePool)(nil)).Elem()
}

func (i *ResourcePool) ToResourcePoolOutput() ResourcePoolOutput {
	return i.ToResourcePoolOutputWithContext(context.Background())
}

func (i *ResourcePool) ToResourcePoolOutputWithContext(ctx context.Context) ResourcePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePoolOutput)
}

// ResourcePoolArrayInput is an input type that accepts ResourcePoolArray and ResourcePoolArrayOutput values.
// You can construct a concrete instance of `ResourcePoolArrayInput` via:
//
//          ResourcePoolArray{ ResourcePoolArgs{...} }
type ResourcePoolArrayInput interface {
	pulumi.Input

	ToResourcePoolArrayOutput() ResourcePoolArrayOutput
	ToResourcePoolArrayOutputWithContext(context.Context) ResourcePoolArrayOutput
}

type ResourcePoolArray []ResourcePoolInput

func (ResourcePoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourcePool)(nil)).Elem()
}

func (i ResourcePoolArray) ToResourcePoolArrayOutput() ResourcePoolArrayOutput {
	return i.ToResourcePoolArrayOutputWithContext(context.Background())
}

func (i ResourcePoolArray) ToResourcePoolArrayOutputWithContext(ctx context.Context) ResourcePoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePoolArrayOutput)
}

// ResourcePoolMapInput is an input type that accepts ResourcePoolMap and ResourcePoolMapOutput values.
// You can construct a concrete instance of `ResourcePoolMapInput` via:
//
//          ResourcePoolMap{ "key": ResourcePoolArgs{...} }
type ResourcePoolMapInput interface {
	pulumi.Input

	ToResourcePoolMapOutput() ResourcePoolMapOutput
	ToResourcePoolMapOutputWithContext(context.Context) ResourcePoolMapOutput
}

type ResourcePoolMap map[string]ResourcePoolInput

func (ResourcePoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourcePool)(nil)).Elem()
}

func (i ResourcePoolMap) ToResourcePoolMapOutput() ResourcePoolMapOutput {
	return i.ToResourcePoolMapOutputWithContext(context.Background())
}

func (i ResourcePoolMap) ToResourcePoolMapOutputWithContext(ctx context.Context) ResourcePoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePoolMapOutput)
}

type ResourcePoolOutput struct{ *pulumi.OutputState }

func (ResourcePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePool)(nil)).Elem()
}

func (o ResourcePoolOutput) ToResourcePoolOutput() ResourcePoolOutput {
	return o
}

func (o ResourcePoolOutput) ToResourcePoolOutputWithContext(ctx context.Context) ResourcePoolOutput {
	return o
}

type ResourcePoolArrayOutput struct{ *pulumi.OutputState }

func (ResourcePoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourcePool)(nil)).Elem()
}

func (o ResourcePoolArrayOutput) ToResourcePoolArrayOutput() ResourcePoolArrayOutput {
	return o
}

func (o ResourcePoolArrayOutput) ToResourcePoolArrayOutputWithContext(ctx context.Context) ResourcePoolArrayOutput {
	return o
}

func (o ResourcePoolArrayOutput) Index(i pulumi.IntInput) ResourcePoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourcePool {
		return vs[0].([]*ResourcePool)[vs[1].(int)]
	}).(ResourcePoolOutput)
}

type ResourcePoolMapOutput struct{ *pulumi.OutputState }

func (ResourcePoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourcePool)(nil)).Elem()
}

func (o ResourcePoolMapOutput) ToResourcePoolMapOutput() ResourcePoolMapOutput {
	return o
}

func (o ResourcePoolMapOutput) ToResourcePoolMapOutputWithContext(ctx context.Context) ResourcePoolMapOutput {
	return o
}

func (o ResourcePoolMapOutput) MapIndex(k pulumi.StringInput) ResourcePoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourcePool {
		return vs[0].(map[string]*ResourcePool)[vs[1].(string)]
	}).(ResourcePoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcePoolInput)(nil)).Elem(), &ResourcePool{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcePoolArrayInput)(nil)).Elem(), ResourcePoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcePoolMapInput)(nil)).Elem(), ResourcePoolMap{})
	pulumi.RegisterOutputType(ResourcePoolOutput{})
	pulumi.RegisterOutputType(ResourcePoolArrayOutput{})
	pulumi.RegisterOutputType(ResourcePoolMapOutput{})
}

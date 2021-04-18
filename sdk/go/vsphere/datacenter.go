// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VMware vSphere datacenter resource. This can be used as the primary
// container of inventory objects such as hosts and virtual machines.
//
// ## Example Usage
// ### Create datacenter on the root folder
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vsphere/sdk/v3/go/vsphere"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vsphere.NewDatacenter(ctx, "prodDatacenter", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Create datacenter on a subfolder
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vsphere/sdk/v3/go/vsphere"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vsphere.NewDatacenter(ctx, "researchDatacenter", &vsphere.DatacenterArgs{
// 			Folder: pulumi.String("/research/"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Datacenter struct {
	pulumi.CustomResourceState

	// Map of custom attribute ids to value
	// strings to set for datacenter resource. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes pulumi.StringMapOutput `pulumi:"customAttributes"`
	// The folder where the datacenter should be created.
	// Forces a new resource if changed.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// Managed object ID of this datacenter.
	Moid pulumi.StringOutput `pulumi:"moid"`
	// The name of the datacenter. This name needs to be unique
	// within the folder. Forces a new resource if changed.
	Name pulumi.StringOutput `pulumi:"name"`
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewDatacenter registers a new resource with the given unique name, arguments, and options.
func NewDatacenter(ctx *pulumi.Context,
	name string, args *DatacenterArgs, opts ...pulumi.ResourceOption) (*Datacenter, error) {
	if args == nil {
		args = &DatacenterArgs{}
	}

	var resource Datacenter
	err := ctx.RegisterResource("vsphere:index/datacenter:Datacenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatacenter gets an existing Datacenter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatacenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatacenterState, opts ...pulumi.ResourceOption) (*Datacenter, error) {
	var resource Datacenter
	err := ctx.ReadResource("vsphere:index/datacenter:Datacenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Datacenter resources.
type datacenterState struct {
	// Map of custom attribute ids to value
	// strings to set for datacenter resource. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The folder where the datacenter should be created.
	// Forces a new resource if changed.
	Folder *string `pulumi:"folder"`
	// Managed object ID of this datacenter.
	Moid *string `pulumi:"moid"`
	// The name of the datacenter. This name needs to be unique
	// within the folder. Forces a new resource if changed.
	Name *string `pulumi:"name"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
}

type DatacenterState struct {
	// Map of custom attribute ids to value
	// strings to set for datacenter resource. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes pulumi.StringMapInput
	// The folder where the datacenter should be created.
	// Forces a new resource if changed.
	Folder pulumi.StringPtrInput
	// Managed object ID of this datacenter.
	Moid pulumi.StringPtrInput
	// The name of the datacenter. This name needs to be unique
	// within the folder. Forces a new resource if changed.
	Name pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
}

func (DatacenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*datacenterState)(nil)).Elem()
}

type datacenterArgs struct {
	// Map of custom attribute ids to value
	// strings to set for datacenter resource. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// The folder where the datacenter should be created.
	// Forces a new resource if changed.
	Folder *string `pulumi:"folder"`
	// The name of the datacenter. This name needs to be unique
	// within the folder. Forces a new resource if changed.
	Name *string `pulumi:"name"`
	// The IDs of any tags to attach to this resource.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a Datacenter resource.
type DatacenterArgs struct {
	// Map of custom attribute ids to value
	// strings to set for datacenter resource. See
	// [here][docs-setting-custom-attributes] for a reference on how to set values
	// for custom attributes.
	CustomAttributes pulumi.StringMapInput
	// The folder where the datacenter should be created.
	// Forces a new resource if changed.
	Folder pulumi.StringPtrInput
	// The name of the datacenter. This name needs to be unique
	// within the folder. Forces a new resource if changed.
	Name pulumi.StringPtrInput
	// The IDs of any tags to attach to this resource.
	Tags pulumi.StringArrayInput
}

func (DatacenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datacenterArgs)(nil)).Elem()
}

type DatacenterInput interface {
	pulumi.Input

	ToDatacenterOutput() DatacenterOutput
	ToDatacenterOutputWithContext(ctx context.Context) DatacenterOutput
}

func (*Datacenter) ElementType() reflect.Type {
	return reflect.TypeOf((*Datacenter)(nil))
}

func (i *Datacenter) ToDatacenterOutput() DatacenterOutput {
	return i.ToDatacenterOutputWithContext(context.Background())
}

func (i *Datacenter) ToDatacenterOutputWithContext(ctx context.Context) DatacenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatacenterOutput)
}

func (i *Datacenter) ToDatacenterPtrOutput() DatacenterPtrOutput {
	return i.ToDatacenterPtrOutputWithContext(context.Background())
}

func (i *Datacenter) ToDatacenterPtrOutputWithContext(ctx context.Context) DatacenterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatacenterPtrOutput)
}

type DatacenterPtrInput interface {
	pulumi.Input

	ToDatacenterPtrOutput() DatacenterPtrOutput
	ToDatacenterPtrOutputWithContext(ctx context.Context) DatacenterPtrOutput
}

type datacenterPtrType DatacenterArgs

func (*datacenterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Datacenter)(nil))
}

func (i *datacenterPtrType) ToDatacenterPtrOutput() DatacenterPtrOutput {
	return i.ToDatacenterPtrOutputWithContext(context.Background())
}

func (i *datacenterPtrType) ToDatacenterPtrOutputWithContext(ctx context.Context) DatacenterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatacenterPtrOutput)
}

// DatacenterArrayInput is an input type that accepts DatacenterArray and DatacenterArrayOutput values.
// You can construct a concrete instance of `DatacenterArrayInput` via:
//
//          DatacenterArray{ DatacenterArgs{...} }
type DatacenterArrayInput interface {
	pulumi.Input

	ToDatacenterArrayOutput() DatacenterArrayOutput
	ToDatacenterArrayOutputWithContext(context.Context) DatacenterArrayOutput
}

type DatacenterArray []DatacenterInput

func (DatacenterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Datacenter)(nil))
}

func (i DatacenterArray) ToDatacenterArrayOutput() DatacenterArrayOutput {
	return i.ToDatacenterArrayOutputWithContext(context.Background())
}

func (i DatacenterArray) ToDatacenterArrayOutputWithContext(ctx context.Context) DatacenterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatacenterArrayOutput)
}

// DatacenterMapInput is an input type that accepts DatacenterMap and DatacenterMapOutput values.
// You can construct a concrete instance of `DatacenterMapInput` via:
//
//          DatacenterMap{ "key": DatacenterArgs{...} }
type DatacenterMapInput interface {
	pulumi.Input

	ToDatacenterMapOutput() DatacenterMapOutput
	ToDatacenterMapOutputWithContext(context.Context) DatacenterMapOutput
}

type DatacenterMap map[string]DatacenterInput

func (DatacenterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Datacenter)(nil))
}

func (i DatacenterMap) ToDatacenterMapOutput() DatacenterMapOutput {
	return i.ToDatacenterMapOutputWithContext(context.Background())
}

func (i DatacenterMap) ToDatacenterMapOutputWithContext(ctx context.Context) DatacenterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatacenterMapOutput)
}

type DatacenterOutput struct {
	*pulumi.OutputState
}

func (DatacenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Datacenter)(nil))
}

func (o DatacenterOutput) ToDatacenterOutput() DatacenterOutput {
	return o
}

func (o DatacenterOutput) ToDatacenterOutputWithContext(ctx context.Context) DatacenterOutput {
	return o
}

func (o DatacenterOutput) ToDatacenterPtrOutput() DatacenterPtrOutput {
	return o.ToDatacenterPtrOutputWithContext(context.Background())
}

func (o DatacenterOutput) ToDatacenterPtrOutputWithContext(ctx context.Context) DatacenterPtrOutput {
	return o.ApplyT(func(v Datacenter) *Datacenter {
		return &v
	}).(DatacenterPtrOutput)
}

type DatacenterPtrOutput struct {
	*pulumi.OutputState
}

func (DatacenterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Datacenter)(nil))
}

func (o DatacenterPtrOutput) ToDatacenterPtrOutput() DatacenterPtrOutput {
	return o
}

func (o DatacenterPtrOutput) ToDatacenterPtrOutputWithContext(ctx context.Context) DatacenterPtrOutput {
	return o
}

type DatacenterArrayOutput struct{ *pulumi.OutputState }

func (DatacenterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Datacenter)(nil))
}

func (o DatacenterArrayOutput) ToDatacenterArrayOutput() DatacenterArrayOutput {
	return o
}

func (o DatacenterArrayOutput) ToDatacenterArrayOutputWithContext(ctx context.Context) DatacenterArrayOutput {
	return o
}

func (o DatacenterArrayOutput) Index(i pulumi.IntInput) DatacenterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Datacenter {
		return vs[0].([]Datacenter)[vs[1].(int)]
	}).(DatacenterOutput)
}

type DatacenterMapOutput struct{ *pulumi.OutputState }

func (DatacenterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Datacenter)(nil))
}

func (o DatacenterMapOutput) ToDatacenterMapOutput() DatacenterMapOutput {
	return o
}

func (o DatacenterMapOutput) ToDatacenterMapOutputWithContext(ctx context.Context) DatacenterMapOutput {
	return o
}

func (o DatacenterMapOutput) MapIndex(k pulumi.StringInput) DatacenterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Datacenter {
		return vs[0].(map[string]Datacenter)[vs[1].(string)]
	}).(DatacenterOutput)
}

func init() {
	pulumi.RegisterOutputType(DatacenterOutput{})
	pulumi.RegisterOutputType(DatacenterPtrOutput{})
	pulumi.RegisterOutputType(DatacenterArrayOutput{})
	pulumi.RegisterOutputType(DatacenterMapOutput{})
}

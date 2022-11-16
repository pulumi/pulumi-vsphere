// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomAttribute struct {
	pulumi.CustomResourceState

	// The object type that this attribute may be
	// applied to. If not set, the custom attribute may be applied to any object
	// type. For a full list, review the Managed Object Types. Forces a new resource if changed.
	ManagedObjectType pulumi.StringPtrOutput `pulumi:"managedObjectType"`
	// The name of the custom attribute.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewCustomAttribute registers a new resource with the given unique name, arguments, and options.
func NewCustomAttribute(ctx *pulumi.Context,
	name string, args *CustomAttributeArgs, opts ...pulumi.ResourceOption) (*CustomAttribute, error) {
	if args == nil {
		args = &CustomAttributeArgs{}
	}

	var resource CustomAttribute
	err := ctx.RegisterResource("vsphere:index/customAttribute:CustomAttribute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomAttribute gets an existing CustomAttribute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomAttribute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomAttributeState, opts ...pulumi.ResourceOption) (*CustomAttribute, error) {
	var resource CustomAttribute
	err := ctx.ReadResource("vsphere:index/customAttribute:CustomAttribute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomAttribute resources.
type customAttributeState struct {
	// The object type that this attribute may be
	// applied to. If not set, the custom attribute may be applied to any object
	// type. For a full list, review the Managed Object Types. Forces a new resource if changed.
	ManagedObjectType *string `pulumi:"managedObjectType"`
	// The name of the custom attribute.
	Name *string `pulumi:"name"`
}

type CustomAttributeState struct {
	// The object type that this attribute may be
	// applied to. If not set, the custom attribute may be applied to any object
	// type. For a full list, review the Managed Object Types. Forces a new resource if changed.
	ManagedObjectType pulumi.StringPtrInput
	// The name of the custom attribute.
	Name pulumi.StringPtrInput
}

func (CustomAttributeState) ElementType() reflect.Type {
	return reflect.TypeOf((*customAttributeState)(nil)).Elem()
}

type customAttributeArgs struct {
	// The object type that this attribute may be
	// applied to. If not set, the custom attribute may be applied to any object
	// type. For a full list, review the Managed Object Types. Forces a new resource if changed.
	ManagedObjectType *string `pulumi:"managedObjectType"`
	// The name of the custom attribute.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a CustomAttribute resource.
type CustomAttributeArgs struct {
	// The object type that this attribute may be
	// applied to. If not set, the custom attribute may be applied to any object
	// type. For a full list, review the Managed Object Types. Forces a new resource if changed.
	ManagedObjectType pulumi.StringPtrInput
	// The name of the custom attribute.
	Name pulumi.StringPtrInput
}

func (CustomAttributeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customAttributeArgs)(nil)).Elem()
}

type CustomAttributeInput interface {
	pulumi.Input

	ToCustomAttributeOutput() CustomAttributeOutput
	ToCustomAttributeOutputWithContext(ctx context.Context) CustomAttributeOutput
}

func (*CustomAttribute) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomAttribute)(nil)).Elem()
}

func (i *CustomAttribute) ToCustomAttributeOutput() CustomAttributeOutput {
	return i.ToCustomAttributeOutputWithContext(context.Background())
}

func (i *CustomAttribute) ToCustomAttributeOutputWithContext(ctx context.Context) CustomAttributeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomAttributeOutput)
}

// CustomAttributeArrayInput is an input type that accepts CustomAttributeArray and CustomAttributeArrayOutput values.
// You can construct a concrete instance of `CustomAttributeArrayInput` via:
//
//	CustomAttributeArray{ CustomAttributeArgs{...} }
type CustomAttributeArrayInput interface {
	pulumi.Input

	ToCustomAttributeArrayOutput() CustomAttributeArrayOutput
	ToCustomAttributeArrayOutputWithContext(context.Context) CustomAttributeArrayOutput
}

type CustomAttributeArray []CustomAttributeInput

func (CustomAttributeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomAttribute)(nil)).Elem()
}

func (i CustomAttributeArray) ToCustomAttributeArrayOutput() CustomAttributeArrayOutput {
	return i.ToCustomAttributeArrayOutputWithContext(context.Background())
}

func (i CustomAttributeArray) ToCustomAttributeArrayOutputWithContext(ctx context.Context) CustomAttributeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomAttributeArrayOutput)
}

// CustomAttributeMapInput is an input type that accepts CustomAttributeMap and CustomAttributeMapOutput values.
// You can construct a concrete instance of `CustomAttributeMapInput` via:
//
//	CustomAttributeMap{ "key": CustomAttributeArgs{...} }
type CustomAttributeMapInput interface {
	pulumi.Input

	ToCustomAttributeMapOutput() CustomAttributeMapOutput
	ToCustomAttributeMapOutputWithContext(context.Context) CustomAttributeMapOutput
}

type CustomAttributeMap map[string]CustomAttributeInput

func (CustomAttributeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomAttribute)(nil)).Elem()
}

func (i CustomAttributeMap) ToCustomAttributeMapOutput() CustomAttributeMapOutput {
	return i.ToCustomAttributeMapOutputWithContext(context.Background())
}

func (i CustomAttributeMap) ToCustomAttributeMapOutputWithContext(ctx context.Context) CustomAttributeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomAttributeMapOutput)
}

type CustomAttributeOutput struct{ *pulumi.OutputState }

func (CustomAttributeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomAttribute)(nil)).Elem()
}

func (o CustomAttributeOutput) ToCustomAttributeOutput() CustomAttributeOutput {
	return o
}

func (o CustomAttributeOutput) ToCustomAttributeOutputWithContext(ctx context.Context) CustomAttributeOutput {
	return o
}

type CustomAttributeArrayOutput struct{ *pulumi.OutputState }

func (CustomAttributeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomAttribute)(nil)).Elem()
}

func (o CustomAttributeArrayOutput) ToCustomAttributeArrayOutput() CustomAttributeArrayOutput {
	return o
}

func (o CustomAttributeArrayOutput) ToCustomAttributeArrayOutputWithContext(ctx context.Context) CustomAttributeArrayOutput {
	return o
}

func (o CustomAttributeArrayOutput) Index(i pulumi.IntInput) CustomAttributeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomAttribute {
		return vs[0].([]*CustomAttribute)[vs[1].(int)]
	}).(CustomAttributeOutput)
}

type CustomAttributeMapOutput struct{ *pulumi.OutputState }

func (CustomAttributeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomAttribute)(nil)).Elem()
}

func (o CustomAttributeMapOutput) ToCustomAttributeMapOutput() CustomAttributeMapOutput {
	return o
}

func (o CustomAttributeMapOutput) ToCustomAttributeMapOutputWithContext(ctx context.Context) CustomAttributeMapOutput {
	return o
}

func (o CustomAttributeMapOutput) MapIndex(k pulumi.StringInput) CustomAttributeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomAttribute {
		return vs[0].(map[string]*CustomAttribute)[vs[1].(string)]
	}).(CustomAttributeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomAttributeInput)(nil)).Elem(), &CustomAttribute{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomAttributeArrayInput)(nil)).Elem(), CustomAttributeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomAttributeMapInput)(nil)).Elem(), CustomAttributeMap{})
	pulumi.RegisterOutputType(CustomAttributeOutput{})
	pulumi.RegisterOutputType(CustomAttributeArrayOutput{})
	pulumi.RegisterOutputType(CustomAttributeMapOutput{})
}

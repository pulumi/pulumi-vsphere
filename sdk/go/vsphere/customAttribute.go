// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type CustomAttribute struct {
	pulumi.CustomResourceState

	// The object type that this attribute may be
	// applied to. If not set, the custom attribute may be applied to any object
	// type. For a full list, click here. Forces a new
	// resource if changed.
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
	// type. For a full list, click here. Forces a new
	// resource if changed.
	ManagedObjectType *string `pulumi:"managedObjectType"`
	// The name of the custom attribute.
	Name *string `pulumi:"name"`
}

type CustomAttributeState struct {
	// The object type that this attribute may be
	// applied to. If not set, the custom attribute may be applied to any object
	// type. For a full list, click here. Forces a new
	// resource if changed.
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
	// type. For a full list, click here. Forces a new
	// resource if changed.
	ManagedObjectType *string `pulumi:"managedObjectType"`
	// The name of the custom attribute.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a CustomAttribute resource.
type CustomAttributeArgs struct {
	// The object type that this attribute may be
	// applied to. If not set, the custom attribute may be applied to any object
	// type. For a full list, click here. Forces a new
	// resource if changed.
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

func (CustomAttribute) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomAttribute)(nil)).Elem()
}

func (i CustomAttribute) ToCustomAttributeOutput() CustomAttributeOutput {
	return i.ToCustomAttributeOutputWithContext(context.Background())
}

func (i CustomAttribute) ToCustomAttributeOutputWithContext(ctx context.Context) CustomAttributeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomAttributeOutput)
}

type CustomAttributeOutput struct {
	*pulumi.OutputState
}

func (CustomAttributeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomAttributeOutput)(nil)).Elem()
}

func (o CustomAttributeOutput) ToCustomAttributeOutput() CustomAttributeOutput {
	return o
}

func (o CustomAttributeOutput) ToCustomAttributeOutputWithContext(ctx context.Context) CustomAttributeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CustomAttributeOutput{})
}

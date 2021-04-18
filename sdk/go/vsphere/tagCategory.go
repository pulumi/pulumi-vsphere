// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TagCategory struct {
	pulumi.CustomResourceState

	// A list object types that this category is
	// valid to be assigned to. For a full list, click
	// here.
	AssociableTypes pulumi.StringArrayOutput `pulumi:"associableTypes"`
	// The number of tags that can be assigned from this
	// category to a single object at once. Can be one of `SINGLE` (object can only
	// be assigned one tag in this category), to `MULTIPLE` (object can be assigned
	// multiple tags in this category). Forces a new resource if changed.
	Cardinality pulumi.StringOutput `pulumi:"cardinality"`
	// A description for the category.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the category.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewTagCategory registers a new resource with the given unique name, arguments, and options.
func NewTagCategory(ctx *pulumi.Context,
	name string, args *TagCategoryArgs, opts ...pulumi.ResourceOption) (*TagCategory, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssociableTypes == nil {
		return nil, errors.New("invalid value for required argument 'AssociableTypes'")
	}
	if args.Cardinality == nil {
		return nil, errors.New("invalid value for required argument 'Cardinality'")
	}
	var resource TagCategory
	err := ctx.RegisterResource("vsphere:index/tagCategory:TagCategory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagCategory gets an existing TagCategory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagCategory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagCategoryState, opts ...pulumi.ResourceOption) (*TagCategory, error) {
	var resource TagCategory
	err := ctx.ReadResource("vsphere:index/tagCategory:TagCategory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagCategory resources.
type tagCategoryState struct {
	// A list object types that this category is
	// valid to be assigned to. For a full list, click
	// here.
	AssociableTypes []string `pulumi:"associableTypes"`
	// The number of tags that can be assigned from this
	// category to a single object at once. Can be one of `SINGLE` (object can only
	// be assigned one tag in this category), to `MULTIPLE` (object can be assigned
	// multiple tags in this category). Forces a new resource if changed.
	Cardinality *string `pulumi:"cardinality"`
	// A description for the category.
	Description *string `pulumi:"description"`
	// The name of the category.
	Name *string `pulumi:"name"`
}

type TagCategoryState struct {
	// A list object types that this category is
	// valid to be assigned to. For a full list, click
	// here.
	AssociableTypes pulumi.StringArrayInput
	// The number of tags that can be assigned from this
	// category to a single object at once. Can be one of `SINGLE` (object can only
	// be assigned one tag in this category), to `MULTIPLE` (object can be assigned
	// multiple tags in this category). Forces a new resource if changed.
	Cardinality pulumi.StringPtrInput
	// A description for the category.
	Description pulumi.StringPtrInput
	// The name of the category.
	Name pulumi.StringPtrInput
}

func (TagCategoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagCategoryState)(nil)).Elem()
}

type tagCategoryArgs struct {
	// A list object types that this category is
	// valid to be assigned to. For a full list, click
	// here.
	AssociableTypes []string `pulumi:"associableTypes"`
	// The number of tags that can be assigned from this
	// category to a single object at once. Can be one of `SINGLE` (object can only
	// be assigned one tag in this category), to `MULTIPLE` (object can be assigned
	// multiple tags in this category). Forces a new resource if changed.
	Cardinality string `pulumi:"cardinality"`
	// A description for the category.
	Description *string `pulumi:"description"`
	// The name of the category.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a TagCategory resource.
type TagCategoryArgs struct {
	// A list object types that this category is
	// valid to be assigned to. For a full list, click
	// here.
	AssociableTypes pulumi.StringArrayInput
	// The number of tags that can be assigned from this
	// category to a single object at once. Can be one of `SINGLE` (object can only
	// be assigned one tag in this category), to `MULTIPLE` (object can be assigned
	// multiple tags in this category). Forces a new resource if changed.
	Cardinality pulumi.StringInput
	// A description for the category.
	Description pulumi.StringPtrInput
	// The name of the category.
	Name pulumi.StringPtrInput
}

func (TagCategoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagCategoryArgs)(nil)).Elem()
}

type TagCategoryInput interface {
	pulumi.Input

	ToTagCategoryOutput() TagCategoryOutput
	ToTagCategoryOutputWithContext(ctx context.Context) TagCategoryOutput
}

func (*TagCategory) ElementType() reflect.Type {
	return reflect.TypeOf((*TagCategory)(nil))
}

func (i *TagCategory) ToTagCategoryOutput() TagCategoryOutput {
	return i.ToTagCategoryOutputWithContext(context.Background())
}

func (i *TagCategory) ToTagCategoryOutputWithContext(ctx context.Context) TagCategoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagCategoryOutput)
}

func (i *TagCategory) ToTagCategoryPtrOutput() TagCategoryPtrOutput {
	return i.ToTagCategoryPtrOutputWithContext(context.Background())
}

func (i *TagCategory) ToTagCategoryPtrOutputWithContext(ctx context.Context) TagCategoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagCategoryPtrOutput)
}

type TagCategoryPtrInput interface {
	pulumi.Input

	ToTagCategoryPtrOutput() TagCategoryPtrOutput
	ToTagCategoryPtrOutputWithContext(ctx context.Context) TagCategoryPtrOutput
}

type tagCategoryPtrType TagCategoryArgs

func (*tagCategoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TagCategory)(nil))
}

func (i *tagCategoryPtrType) ToTagCategoryPtrOutput() TagCategoryPtrOutput {
	return i.ToTagCategoryPtrOutputWithContext(context.Background())
}

func (i *tagCategoryPtrType) ToTagCategoryPtrOutputWithContext(ctx context.Context) TagCategoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagCategoryPtrOutput)
}

// TagCategoryArrayInput is an input type that accepts TagCategoryArray and TagCategoryArrayOutput values.
// You can construct a concrete instance of `TagCategoryArrayInput` via:
//
//          TagCategoryArray{ TagCategoryArgs{...} }
type TagCategoryArrayInput interface {
	pulumi.Input

	ToTagCategoryArrayOutput() TagCategoryArrayOutput
	ToTagCategoryArrayOutputWithContext(context.Context) TagCategoryArrayOutput
}

type TagCategoryArray []TagCategoryInput

func (TagCategoryArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*TagCategory)(nil))
}

func (i TagCategoryArray) ToTagCategoryArrayOutput() TagCategoryArrayOutput {
	return i.ToTagCategoryArrayOutputWithContext(context.Background())
}

func (i TagCategoryArray) ToTagCategoryArrayOutputWithContext(ctx context.Context) TagCategoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagCategoryArrayOutput)
}

// TagCategoryMapInput is an input type that accepts TagCategoryMap and TagCategoryMapOutput values.
// You can construct a concrete instance of `TagCategoryMapInput` via:
//
//          TagCategoryMap{ "key": TagCategoryArgs{...} }
type TagCategoryMapInput interface {
	pulumi.Input

	ToTagCategoryMapOutput() TagCategoryMapOutput
	ToTagCategoryMapOutputWithContext(context.Context) TagCategoryMapOutput
}

type TagCategoryMap map[string]TagCategoryInput

func (TagCategoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*TagCategory)(nil))
}

func (i TagCategoryMap) ToTagCategoryMapOutput() TagCategoryMapOutput {
	return i.ToTagCategoryMapOutputWithContext(context.Background())
}

func (i TagCategoryMap) ToTagCategoryMapOutputWithContext(ctx context.Context) TagCategoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagCategoryMapOutput)
}

type TagCategoryOutput struct {
	*pulumi.OutputState
}

func (TagCategoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagCategory)(nil))
}

func (o TagCategoryOutput) ToTagCategoryOutput() TagCategoryOutput {
	return o
}

func (o TagCategoryOutput) ToTagCategoryOutputWithContext(ctx context.Context) TagCategoryOutput {
	return o
}

func (o TagCategoryOutput) ToTagCategoryPtrOutput() TagCategoryPtrOutput {
	return o.ToTagCategoryPtrOutputWithContext(context.Background())
}

func (o TagCategoryOutput) ToTagCategoryPtrOutputWithContext(ctx context.Context) TagCategoryPtrOutput {
	return o.ApplyT(func(v TagCategory) *TagCategory {
		return &v
	}).(TagCategoryPtrOutput)
}

type TagCategoryPtrOutput struct {
	*pulumi.OutputState
}

func (TagCategoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagCategory)(nil))
}

func (o TagCategoryPtrOutput) ToTagCategoryPtrOutput() TagCategoryPtrOutput {
	return o
}

func (o TagCategoryPtrOutput) ToTagCategoryPtrOutputWithContext(ctx context.Context) TagCategoryPtrOutput {
	return o
}

type TagCategoryArrayOutput struct{ *pulumi.OutputState }

func (TagCategoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagCategory)(nil))
}

func (o TagCategoryArrayOutput) ToTagCategoryArrayOutput() TagCategoryArrayOutput {
	return o
}

func (o TagCategoryArrayOutput) ToTagCategoryArrayOutputWithContext(ctx context.Context) TagCategoryArrayOutput {
	return o
}

func (o TagCategoryArrayOutput) Index(i pulumi.IntInput) TagCategoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagCategory {
		return vs[0].([]TagCategory)[vs[1].(int)]
	}).(TagCategoryOutput)
}

type TagCategoryMapOutput struct{ *pulumi.OutputState }

func (TagCategoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TagCategory)(nil))
}

func (o TagCategoryMapOutput) ToTagCategoryMapOutput() TagCategoryMapOutput {
	return o
}

func (o TagCategoryMapOutput) ToTagCategoryMapOutputWithContext(ctx context.Context) TagCategoryMapOutput {
	return o
}

func (o TagCategoryMapOutput) MapIndex(k pulumi.StringInput) TagCategoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TagCategory {
		return vs[0].(map[string]TagCategory)[vs[1].(string)]
	}).(TagCategoryOutput)
}

func init() {
	pulumi.RegisterOutputType(TagCategoryOutput{})
	pulumi.RegisterOutputType(TagCategoryPtrOutput{})
	pulumi.RegisterOutputType(TagCategoryArrayOutput{})
	pulumi.RegisterOutputType(TagCategoryMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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

func init() {
	pulumi.RegisterOutputType(TagCategoryOutput{})
}

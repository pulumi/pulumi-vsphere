// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `.TagCategory` resource can be used to create and manage tag
// categories, which determine how tags are grouped together and applied to
// specific objects.
//
// For more information about tags, click [here][ext-tags-general]. For more
// information about tag categories specifically, click
// [here][ext-tag-categories].
//
// [ext-tags-general]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vcenterhost.doc/GUID-E8E854DD-AA97-4E0C-8419-CE84F93C4058.html
// [ext-tag-categories]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vcenterhost.doc/GUID-BA3D1794-28F2-43F3-BCE9-3964CB207FB6.html
//
// > **NOTE:** Tagging support is unsupported on direct ESXi connections and
// requires vCenter 6.0 or higher.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/tag_category.html.markdown.
type TagCategory struct {
	s *pulumi.ResourceState
}

// NewTagCategory registers a new resource with the given unique name, arguments, and options.
func NewTagCategory(ctx *pulumi.Context,
	name string, args *TagCategoryArgs, opts ...pulumi.ResourceOpt) (*TagCategory, error) {
	if args == nil || args.AssociableTypes == nil {
		return nil, errors.New("missing required argument 'AssociableTypes'")
	}
	if args == nil || args.Cardinality == nil {
		return nil, errors.New("missing required argument 'Cardinality'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["associableTypes"] = nil
		inputs["cardinality"] = nil
		inputs["description"] = nil
		inputs["name"] = nil
	} else {
		inputs["associableTypes"] = args.AssociableTypes
		inputs["cardinality"] = args.Cardinality
		inputs["description"] = args.Description
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("vsphere:index/tagCategory:TagCategory", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TagCategory{s: s}, nil
}

// GetTagCategory gets an existing TagCategory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagCategory(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TagCategoryState, opts ...pulumi.ResourceOpt) (*TagCategory, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["associableTypes"] = state.AssociableTypes
		inputs["cardinality"] = state.Cardinality
		inputs["description"] = state.Description
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("vsphere:index/tagCategory:TagCategory", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TagCategory{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *TagCategory) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *TagCategory) ID() pulumi.IDOutput {
	return r.s.ID()
}

// A list object types that this category is
// valid to be assigned to. For a full list, click
// here.
func (r *TagCategory) AssociableTypes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["associableTypes"])
}

// The number of tags that can be assigned from this
// category to a single object at once. Can be one of `SINGLE` (object can only
// be assigned one tag in this category), to `MULTIPLE` (object can be assigned
// multiple tags in this category). Forces a new resource if changed.
func (r *TagCategory) Cardinality() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["cardinality"])
}

// A description for the category.
func (r *TagCategory) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The name of the category.
func (r *TagCategory) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering TagCategory resources.
type TagCategoryState struct {
	// A list object types that this category is
	// valid to be assigned to. For a full list, click
	// here.
	AssociableTypes interface{}
	// The number of tags that can be assigned from this
	// category to a single object at once. Can be one of `SINGLE` (object can only
	// be assigned one tag in this category), to `MULTIPLE` (object can be assigned
	// multiple tags in this category). Forces a new resource if changed.
	Cardinality interface{}
	// A description for the category.
	Description interface{}
	// The name of the category.
	Name interface{}
}

// The set of arguments for constructing a TagCategory resource.
type TagCategoryArgs struct {
	// A list object types that this category is
	// valid to be assigned to. For a full list, click
	// here.
	AssociableTypes interface{}
	// The number of tags that can be assigned from this
	// category to a single object at once. Can be one of `SINGLE` (object can only
	// be assigned one tag in this category), to `MULTIPLE` (object can be assigned
	// multiple tags in this category). Forces a new resource if changed.
	Cardinality interface{}
	// A description for the category.
	Description interface{}
	// The name of the category.
	Name interface{}
}

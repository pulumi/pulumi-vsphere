// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vsphere

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `.Tag` resource can be used to create and manage tags, which allow
// you to attach metadata to objects in the vSphere inventory to make these
// objects more sortable and searchable.
//
// For more information about tags, click [here][ext-tags-general].
//
// [ext-tags-general]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vcenterhost.doc/GUID-E8E854DD-AA97-4E0C-8419-CE84F93C4058.html
//
// > **NOTE:** Tagging support is unsupported on direct ESXi connections and
// requires vCenter 6.0 or higher.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/tag.html.markdown.
type Tag struct {
	pulumi.CustomResourceState

	// The unique identifier of the parent category in
	// which this tag will be created. Forces a new resource if changed.
	CategoryId pulumi.StringOutput `pulumi:"categoryId"`
	// A description for the tag.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the tag. The name must be unique
	// within its category.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewTag registers a new resource with the given unique name, arguments, and options.
func NewTag(ctx *pulumi.Context,
	name string, args *TagArgs, opts ...pulumi.ResourceOption) (*Tag, error) {
	if args == nil || args.CategoryId == nil {
		return nil, errors.New("missing required argument 'CategoryId'")
	}
	if args == nil {
		args = &TagArgs{}
	}
	var resource Tag
	err := ctx.RegisterResource("vsphere:index/tag:Tag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTag gets an existing Tag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagState, opts ...pulumi.ResourceOption) (*Tag, error) {
	var resource Tag
	err := ctx.ReadResource("vsphere:index/tag:Tag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tag resources.
type tagState struct {
	// The unique identifier of the parent category in
	// which this tag will be created. Forces a new resource if changed.
	CategoryId *string `pulumi:"categoryId"`
	// A description for the tag.
	Description *string `pulumi:"description"`
	// The display name of the tag. The name must be unique
	// within its category.
	Name *string `pulumi:"name"`
}

type TagState struct {
	// The unique identifier of the parent category in
	// which this tag will be created. Forces a new resource if changed.
	CategoryId pulumi.StringPtrInput
	// A description for the tag.
	Description pulumi.StringPtrInput
	// The display name of the tag. The name must be unique
	// within its category.
	Name pulumi.StringPtrInput
}

func (TagState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagState)(nil)).Elem()
}

type tagArgs struct {
	// The unique identifier of the parent category in
	// which this tag will be created. Forces a new resource if changed.
	CategoryId string `pulumi:"categoryId"`
	// A description for the tag.
	Description *string `pulumi:"description"`
	// The display name of the tag. The name must be unique
	// within its category.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Tag resource.
type TagArgs struct {
	// The unique identifier of the parent category in
	// which this tag will be created. Forces a new resource if changed.
	CategoryId pulumi.StringInput
	// A description for the tag.
	Description pulumi.StringPtrInput
	// The display name of the tag. The name must be unique
	// within its category.
	Name pulumi.StringPtrInput
}

func (TagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagArgs)(nil)).Elem()
}

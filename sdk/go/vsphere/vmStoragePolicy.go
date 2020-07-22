// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `VmStoragePolicy` resource can be used to create and manage storage
// policies. Using this storage policy, tag based placement rules can be created to
// place a VM on a particular tagged datastore.
//
// ## Example Usage
//
// This example creates a storage policy with tagRule having cat1 as tagCategory and
// tag1, tag2 as the tags. While creating a VM, this policy can be referenced to place
// the VM in any of the compatible datastore tagged with these tags.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vsphere/sdk/v2/go/vsphere"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "DC"
// 		_, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		tagCategory, err := vsphere.LookupTagCategory(ctx, &vsphere.LookupTagCategoryArgs{
// 			Name: "cat1",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		tag1, err := vsphere.LookupTag(ctx, &vsphere.LookupTagArgs{
// 			Name:       "tag1",
// 			CategoryId: tagCategory.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		tag2, err := vsphere.LookupTag(ctx, &vsphere.LookupTagArgs{
// 			Name:       "tag2",
// 			CategoryId: tagCategory.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vsphere.NewVmStoragePolicy(ctx, "policyTagBasedPlacement", &vsphere.VmStoragePolicyArgs{
// 			Description: pulumi.String("description"),
// 			TagRules: vsphere.VmStoragePolicyTagRuleArray{
// 				&vsphere.VmStoragePolicyTagRuleArgs{
// 					TagCategory: pulumi.String(tagCategory.Name),
// 					Tags: pulumi.StringArray{
// 						pulumi.String(tag1.Name),
// 						pulumi.String(tag2.Name),
// 					},
// 					IncludeDatastoresWithTags: pulumi.Bool(true),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type VmStoragePolicy struct {
	pulumi.CustomResourceState

	// Description of the storage policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the storage policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of tag rules. The tag category and tags to be associated to this storage policy.
	TagRules VmStoragePolicyTagRuleArrayOutput `pulumi:"tagRules"`
}

// NewVmStoragePolicy registers a new resource with the given unique name, arguments, and options.
func NewVmStoragePolicy(ctx *pulumi.Context,
	name string, args *VmStoragePolicyArgs, opts ...pulumi.ResourceOption) (*VmStoragePolicy, error) {
	if args == nil || args.TagRules == nil {
		return nil, errors.New("missing required argument 'TagRules'")
	}
	if args == nil {
		args = &VmStoragePolicyArgs{}
	}
	var resource VmStoragePolicy
	err := ctx.RegisterResource("vsphere:index/vmStoragePolicy:VmStoragePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVmStoragePolicy gets an existing VmStoragePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVmStoragePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VmStoragePolicyState, opts ...pulumi.ResourceOption) (*VmStoragePolicy, error) {
	var resource VmStoragePolicy
	err := ctx.ReadResource("vsphere:index/vmStoragePolicy:VmStoragePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VmStoragePolicy resources.
type vmStoragePolicyState struct {
	// Description of the storage policy.
	Description *string `pulumi:"description"`
	// The name of the storage policy.
	Name *string `pulumi:"name"`
	// List of tag rules. The tag category and tags to be associated to this storage policy.
	TagRules []VmStoragePolicyTagRule `pulumi:"tagRules"`
}

type VmStoragePolicyState struct {
	// Description of the storage policy.
	Description pulumi.StringPtrInput
	// The name of the storage policy.
	Name pulumi.StringPtrInput
	// List of tag rules. The tag category and tags to be associated to this storage policy.
	TagRules VmStoragePolicyTagRuleArrayInput
}

func (VmStoragePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmStoragePolicyState)(nil)).Elem()
}

type vmStoragePolicyArgs struct {
	// Description of the storage policy.
	Description *string `pulumi:"description"`
	// The name of the storage policy.
	Name *string `pulumi:"name"`
	// List of tag rules. The tag category and tags to be associated to this storage policy.
	TagRules []VmStoragePolicyTagRule `pulumi:"tagRules"`
}

// The set of arguments for constructing a VmStoragePolicy resource.
type VmStoragePolicyArgs struct {
	// Description of the storage policy.
	Description pulumi.StringPtrInput
	// The name of the storage policy.
	Name pulumi.StringPtrInput
	// List of tag rules. The tag category and tags to be associated to this storage policy.
	TagRules VmStoragePolicyTagRuleArrayInput
}

func (VmStoragePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmStoragePolicyArgs)(nil)).Elem()
}
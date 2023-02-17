// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `VmStoragePolicy` resource can be used to create and manage storage
// policies. Using this resource, tag based placement rules can be created to
// place virtual machines on a datastore with matching tags. If storage requirements for the applications on the virtual machine change, you can modify the storage policy that was originally applied to the virtual machine.
//
// ## Example Usage
//
// The following example creates storage policies with `tagRules` base on sets of environment, service level, and replication attributes.
//
// In this example, tags are first applied to datastores.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vsphere.LookupTagCategory(ctx, &vsphere.LookupTagCategoryArgs{
//				Name: "environment",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.LookupTagCategory(ctx, &vsphere.LookupTagCategoryArgs{
//				Name: "service_level",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.LookupTagCategory(ctx, &vsphere.LookupTagCategoryArgs{
//				Name: "replication",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.LookupTag(ctx, &vsphere.LookupTagArgs{
//				CategoryId: "data.vsphere_tag_category.environment.id",
//				Name:       "production",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.LookupTag(ctx, &vsphere.LookupTagArgs{
//				CategoryId: "data.vsphere_tag_category.environment.id",
//				Name:       "development",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.LookupTag(ctx, &vsphere.LookupTagArgs{
//				CategoryId: "data.vsphere_tag_category.service_level.id",
//				Name:       "platinum",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.LookupTag(ctx, &vsphere.LookupTagArgs{
//				CategoryId: "data.vsphere_tag_category.service_level.id",
//				Name:       "platinum",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.LookupTag(ctx, &vsphere.LookupTagArgs{
//				CategoryId: "data.vsphere_tag_category.service_level.id",
//				Name:       "silver",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.LookupTag(ctx, &vsphere.LookupTagArgs{
//				CategoryId: "data.vsphere_tag_category.service_level.id",
//				Name:       "bronze",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.LookupTag(ctx, &vsphere.LookupTagArgs{
//				CategoryId: "data.vsphere_tag_category.replication.id",
//				Name:       "replicated",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.LookupTag(ctx, &vsphere.LookupTagArgs{
//				CategoryId: "data.vsphere_tag_category.replication.id",
//				Name:       "non_replicated",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.NewVmfsDatastore(ctx, "prodDatastore", &vsphere.VmfsDatastoreArgs{
//				Tags: pulumi.StringArray{
//					pulumi.String("data.vsphere_tag.production.id"),
//					pulumi.String("data.vsphere_tag.platinum.id"),
//					pulumi.String("data.vsphere_tag.replicated.id"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.NewNasDatastore(ctx, "devDatastore", &vsphere.NasDatastoreArgs{
//				Tags: pulumi.StringArray{
//					pulumi.String("data.vsphere_tag.development.id"),
//					pulumi.String("data.vsphere_tag.silver.id"),
//					pulumi.String("data.vsphere_tag.non_replicated.id"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Next, storage policies are created and `tagRules` are applied.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vsphere.NewVmStoragePolicy(ctx, "prodPlatinumReplicated", &vsphere.VmStoragePolicyArgs{
//				Description: pulumi.String("prod_platinum_replicated"),
//				TagRules: vsphere.VmStoragePolicyTagRuleArray{
//					&vsphere.VmStoragePolicyTagRuleArgs{
//						TagCategory: pulumi.Any(data.Vsphere_tag_category.Environment.Name),
//						Tags: pulumi.StringArray{
//							data.Vsphere_tag.Production.Name,
//						},
//						IncludeDatastoresWithTags: pulumi.Bool(true),
//					},
//					&vsphere.VmStoragePolicyTagRuleArgs{
//						TagCategory: pulumi.Any(data.Vsphere_tag_category.Service_level.Name),
//						Tags: pulumi.StringArray{
//							data.Vsphere_tag.Platinum.Name,
//						},
//						IncludeDatastoresWithTags: pulumi.Bool(true),
//					},
//					&vsphere.VmStoragePolicyTagRuleArgs{
//						TagCategory: pulumi.Any(data.Vsphere_tag_category.Replication.Name),
//						Tags: pulumi.StringArray{
//							data.Vsphere_tag.Replicated.Name,
//						},
//						IncludeDatastoresWithTags: pulumi.Bool(true),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.NewVmStoragePolicy(ctx, "devSilverNonreplicated", &vsphere.VmStoragePolicyArgs{
//				Description: pulumi.String("dev_silver_nonreplicated"),
//				TagRules: vsphere.VmStoragePolicyTagRuleArray{
//					&vsphere.VmStoragePolicyTagRuleArgs{
//						TagCategory: pulumi.Any(data.Vsphere_tag_category.Environment.Name),
//						Tags: pulumi.StringArray{
//							data.Vsphere_tag.Development.Name,
//						},
//						IncludeDatastoresWithTags: pulumi.Bool(true),
//					},
//					&vsphere.VmStoragePolicyTagRuleArgs{
//						TagCategory: pulumi.Any(data.Vsphere_tag_category.Service_level.Name),
//						Tags: pulumi.StringArray{
//							data.Vsphere_tag.Silver.Name,
//						},
//						IncludeDatastoresWithTags: pulumi.Bool(true),
//					},
//					&vsphere.VmStoragePolicyTagRuleArgs{
//						TagCategory: pulumi.Any(data.Vsphere_tag_category.Replication.Name),
//						Tags: pulumi.StringArray{
//							data.Vsphere_tag.Non_replicated.Name,
//						},
//						IncludeDatastoresWithTags: pulumi.Bool(true),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Lasttly, when creating a virtual machine resource, a storage policy can be specificed to direct virtual machine placement to a datastore which matches the policy's `tagsRules`.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vsphere.GetPolicy(ctx, &vsphere.GetPolicyArgs{
//				Name: "prod_platinum_replicated",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.GetPolicy(ctx, &vsphere.GetPolicyArgs{
//				Name: "dev_silver_nonreplicated",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.NewVirtualMachine(ctx, "prodVm", &vsphere.VirtualMachineArgs{
//				StoragePolicyId: pulumi.Any(data.Vsphere_storage_policy.Storage_policy.Prod_platinum_replicated.Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.NewVirtualMachine(ctx, "devVm", &vsphere.VirtualMachineArgs{
//				StoragePolicyId: pulumi.Any(data.Vsphere_storage_policy.Storage_policy.Dev_silver_nonreplicated.Id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TagRules == nil {
		return nil, errors.New("invalid value for required argument 'TagRules'")
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

type VmStoragePolicyInput interface {
	pulumi.Input

	ToVmStoragePolicyOutput() VmStoragePolicyOutput
	ToVmStoragePolicyOutputWithContext(ctx context.Context) VmStoragePolicyOutput
}

func (*VmStoragePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**VmStoragePolicy)(nil)).Elem()
}

func (i *VmStoragePolicy) ToVmStoragePolicyOutput() VmStoragePolicyOutput {
	return i.ToVmStoragePolicyOutputWithContext(context.Background())
}

func (i *VmStoragePolicy) ToVmStoragePolicyOutputWithContext(ctx context.Context) VmStoragePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmStoragePolicyOutput)
}

// VmStoragePolicyArrayInput is an input type that accepts VmStoragePolicyArray and VmStoragePolicyArrayOutput values.
// You can construct a concrete instance of `VmStoragePolicyArrayInput` via:
//
//	VmStoragePolicyArray{ VmStoragePolicyArgs{...} }
type VmStoragePolicyArrayInput interface {
	pulumi.Input

	ToVmStoragePolicyArrayOutput() VmStoragePolicyArrayOutput
	ToVmStoragePolicyArrayOutputWithContext(context.Context) VmStoragePolicyArrayOutput
}

type VmStoragePolicyArray []VmStoragePolicyInput

func (VmStoragePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VmStoragePolicy)(nil)).Elem()
}

func (i VmStoragePolicyArray) ToVmStoragePolicyArrayOutput() VmStoragePolicyArrayOutput {
	return i.ToVmStoragePolicyArrayOutputWithContext(context.Background())
}

func (i VmStoragePolicyArray) ToVmStoragePolicyArrayOutputWithContext(ctx context.Context) VmStoragePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmStoragePolicyArrayOutput)
}

// VmStoragePolicyMapInput is an input type that accepts VmStoragePolicyMap and VmStoragePolicyMapOutput values.
// You can construct a concrete instance of `VmStoragePolicyMapInput` via:
//
//	VmStoragePolicyMap{ "key": VmStoragePolicyArgs{...} }
type VmStoragePolicyMapInput interface {
	pulumi.Input

	ToVmStoragePolicyMapOutput() VmStoragePolicyMapOutput
	ToVmStoragePolicyMapOutputWithContext(context.Context) VmStoragePolicyMapOutput
}

type VmStoragePolicyMap map[string]VmStoragePolicyInput

func (VmStoragePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VmStoragePolicy)(nil)).Elem()
}

func (i VmStoragePolicyMap) ToVmStoragePolicyMapOutput() VmStoragePolicyMapOutput {
	return i.ToVmStoragePolicyMapOutputWithContext(context.Background())
}

func (i VmStoragePolicyMap) ToVmStoragePolicyMapOutputWithContext(ctx context.Context) VmStoragePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmStoragePolicyMapOutput)
}

type VmStoragePolicyOutput struct{ *pulumi.OutputState }

func (VmStoragePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmStoragePolicy)(nil)).Elem()
}

func (o VmStoragePolicyOutput) ToVmStoragePolicyOutput() VmStoragePolicyOutput {
	return o
}

func (o VmStoragePolicyOutput) ToVmStoragePolicyOutputWithContext(ctx context.Context) VmStoragePolicyOutput {
	return o
}

// Description of the storage policy.
func (o VmStoragePolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VmStoragePolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the storage policy.
func (o VmStoragePolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VmStoragePolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of tag rules. The tag category and tags to be associated to this storage policy.
func (o VmStoragePolicyOutput) TagRules() VmStoragePolicyTagRuleArrayOutput {
	return o.ApplyT(func(v *VmStoragePolicy) VmStoragePolicyTagRuleArrayOutput { return v.TagRules }).(VmStoragePolicyTagRuleArrayOutput)
}

type VmStoragePolicyArrayOutput struct{ *pulumi.OutputState }

func (VmStoragePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VmStoragePolicy)(nil)).Elem()
}

func (o VmStoragePolicyArrayOutput) ToVmStoragePolicyArrayOutput() VmStoragePolicyArrayOutput {
	return o
}

func (o VmStoragePolicyArrayOutput) ToVmStoragePolicyArrayOutputWithContext(ctx context.Context) VmStoragePolicyArrayOutput {
	return o
}

func (o VmStoragePolicyArrayOutput) Index(i pulumi.IntInput) VmStoragePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VmStoragePolicy {
		return vs[0].([]*VmStoragePolicy)[vs[1].(int)]
	}).(VmStoragePolicyOutput)
}

type VmStoragePolicyMapOutput struct{ *pulumi.OutputState }

func (VmStoragePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VmStoragePolicy)(nil)).Elem()
}

func (o VmStoragePolicyMapOutput) ToVmStoragePolicyMapOutput() VmStoragePolicyMapOutput {
	return o
}

func (o VmStoragePolicyMapOutput) ToVmStoragePolicyMapOutputWithContext(ctx context.Context) VmStoragePolicyMapOutput {
	return o
}

func (o VmStoragePolicyMapOutput) MapIndex(k pulumi.StringInput) VmStoragePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VmStoragePolicy {
		return vs[0].(map[string]*VmStoragePolicy)[vs[1].(string)]
	}).(VmStoragePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VmStoragePolicyInput)(nil)).Elem(), &VmStoragePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmStoragePolicyArrayInput)(nil)).Elem(), VmStoragePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmStoragePolicyMapInput)(nil)).Elem(), VmStoragePolicyMap{})
	pulumi.RegisterOutputType(VmStoragePolicyOutput{})
	pulumi.RegisterOutputType(VmStoragePolicyArrayOutput{})
	pulumi.RegisterOutputType(VmStoragePolicyMapOutput{})
}

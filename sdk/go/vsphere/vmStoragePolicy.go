// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The `VmStoragePolicy` resource can be used to create and manage storage
// policies. Using this resource, tag based placement rules can be created to
// place virtual machines on a datastore with matching tags. If storage requirements for the applications on the virtual machine change, you can modify the storage policy that was originally applied to the virtual machine.
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
	opts = internal.PkgResourceDefaultOpts(opts)
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

func (i *VmStoragePolicy) ToOutput(ctx context.Context) pulumix.Output[*VmStoragePolicy] {
	return pulumix.Output[*VmStoragePolicy]{
		OutputState: i.ToVmStoragePolicyOutputWithContext(ctx).OutputState,
	}
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

func (i VmStoragePolicyArray) ToOutput(ctx context.Context) pulumix.Output[[]*VmStoragePolicy] {
	return pulumix.Output[[]*VmStoragePolicy]{
		OutputState: i.ToVmStoragePolicyArrayOutputWithContext(ctx).OutputState,
	}
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

func (i VmStoragePolicyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VmStoragePolicy] {
	return pulumix.Output[map[string]*VmStoragePolicy]{
		OutputState: i.ToVmStoragePolicyMapOutputWithContext(ctx).OutputState,
	}
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

func (o VmStoragePolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*VmStoragePolicy] {
	return pulumix.Output[*VmStoragePolicy]{
		OutputState: o.OutputState,
	}
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

func (o VmStoragePolicyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VmStoragePolicy] {
	return pulumix.Output[[]*VmStoragePolicy]{
		OutputState: o.OutputState,
	}
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

func (o VmStoragePolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VmStoragePolicy] {
	return pulumix.Output[map[string]*VmStoragePolicy]{
		OutputState: o.OutputState,
	}
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

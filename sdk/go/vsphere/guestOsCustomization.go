// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `GuestOsCustomization` resource can be used to a customization specification for a guest operating system.
//
// > **NOTE:** The name attribute is unique identifier for the guest OS spec per VC.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := vsphere.NewGuestOsCustomization(ctx, "windows_customization", &vsphere.GuestOsCustomizationArgs{
//				Name: pulumi.String("windows-spec"),
//				Type: pulumi.String("Windows"),
//				Spec: &vsphere.GuestOsCustomizationSpecArgs{
//					WindowsOptions: &vsphere.GuestOsCustomizationSpecWindowsOptionsArgs{
//						RunOnceCommandLists: pulumi.StringArray{
//							pulumi.String("command-1"),
//							pulumi.String("command-2"),
//						},
//						ComputerName:   pulumi.String("windows"),
//						AutoLogon:      pulumi.Bool(false),
//						AutoLogonCount: pulumi.Int(0),
//						AdminPassword:  pulumi.String("VMware1!"),
//						TimeZone:       pulumi.Int(4),
//						Workgroup:      pulumi.String("workgroup"),
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
// <!--End PulumiCodeChooser -->
type GuestOsCustomization struct {
	pulumi.CustomResourceState

	// The number of last changed version to the customization specification.
	ChangeVersion pulumi.StringOutput `pulumi:"changeVersion"`
	// The description for the customization specification.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The time of last modification to the customization specification.
	LastUpdateTime pulumi.StringOutput `pulumi:"lastUpdateTime"`
	// The name of the customization specification is the unique identifier per vCenter Server instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// Container object for the Guest OS properties about to be customized . See virtual machine customizations
	Spec GuestOsCustomizationSpecOutput `pulumi:"spec"`
	// The type of customization specification: One among: Windows, Linux.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGuestOsCustomization registers a new resource with the given unique name, arguments, and options.
func NewGuestOsCustomization(ctx *pulumi.Context,
	name string, args *GuestOsCustomizationArgs, opts ...pulumi.ResourceOption) (*GuestOsCustomization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GuestOsCustomization
	err := ctx.RegisterResource("vsphere:index/guestOsCustomization:GuestOsCustomization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGuestOsCustomization gets an existing GuestOsCustomization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGuestOsCustomization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestOsCustomizationState, opts ...pulumi.ResourceOption) (*GuestOsCustomization, error) {
	var resource GuestOsCustomization
	err := ctx.ReadResource("vsphere:index/guestOsCustomization:GuestOsCustomization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GuestOsCustomization resources.
type guestOsCustomizationState struct {
	// The number of last changed version to the customization specification.
	ChangeVersion *string `pulumi:"changeVersion"`
	// The description for the customization specification.
	Description *string `pulumi:"description"`
	// The time of last modification to the customization specification.
	LastUpdateTime *string `pulumi:"lastUpdateTime"`
	// The name of the customization specification is the unique identifier per vCenter Server instance.
	Name *string `pulumi:"name"`
	// Container object for the Guest OS properties about to be customized . See virtual machine customizations
	Spec *GuestOsCustomizationSpec `pulumi:"spec"`
	// The type of customization specification: One among: Windows, Linux.
	Type *string `pulumi:"type"`
}

type GuestOsCustomizationState struct {
	// The number of last changed version to the customization specification.
	ChangeVersion pulumi.StringPtrInput
	// The description for the customization specification.
	Description pulumi.StringPtrInput
	// The time of last modification to the customization specification.
	LastUpdateTime pulumi.StringPtrInput
	// The name of the customization specification is the unique identifier per vCenter Server instance.
	Name pulumi.StringPtrInput
	// Container object for the Guest OS properties about to be customized . See virtual machine customizations
	Spec GuestOsCustomizationSpecPtrInput
	// The type of customization specification: One among: Windows, Linux.
	Type pulumi.StringPtrInput
}

func (GuestOsCustomizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*guestOsCustomizationState)(nil)).Elem()
}

type guestOsCustomizationArgs struct {
	// The description for the customization specification.
	Description *string `pulumi:"description"`
	// The name of the customization specification is the unique identifier per vCenter Server instance.
	Name *string `pulumi:"name"`
	// Container object for the Guest OS properties about to be customized . See virtual machine customizations
	Spec GuestOsCustomizationSpec `pulumi:"spec"`
	// The type of customization specification: One among: Windows, Linux.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a GuestOsCustomization resource.
type GuestOsCustomizationArgs struct {
	// The description for the customization specification.
	Description pulumi.StringPtrInput
	// The name of the customization specification is the unique identifier per vCenter Server instance.
	Name pulumi.StringPtrInput
	// Container object for the Guest OS properties about to be customized . See virtual machine customizations
	Spec GuestOsCustomizationSpecInput
	// The type of customization specification: One among: Windows, Linux.
	Type pulumi.StringInput
}

func (GuestOsCustomizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*guestOsCustomizationArgs)(nil)).Elem()
}

type GuestOsCustomizationInput interface {
	pulumi.Input

	ToGuestOsCustomizationOutput() GuestOsCustomizationOutput
	ToGuestOsCustomizationOutputWithContext(ctx context.Context) GuestOsCustomizationOutput
}

func (*GuestOsCustomization) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestOsCustomization)(nil)).Elem()
}

func (i *GuestOsCustomization) ToGuestOsCustomizationOutput() GuestOsCustomizationOutput {
	return i.ToGuestOsCustomizationOutputWithContext(context.Background())
}

func (i *GuestOsCustomization) ToGuestOsCustomizationOutputWithContext(ctx context.Context) GuestOsCustomizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestOsCustomizationOutput)
}

// GuestOsCustomizationArrayInput is an input type that accepts GuestOsCustomizationArray and GuestOsCustomizationArrayOutput values.
// You can construct a concrete instance of `GuestOsCustomizationArrayInput` via:
//
//	GuestOsCustomizationArray{ GuestOsCustomizationArgs{...} }
type GuestOsCustomizationArrayInput interface {
	pulumi.Input

	ToGuestOsCustomizationArrayOutput() GuestOsCustomizationArrayOutput
	ToGuestOsCustomizationArrayOutputWithContext(context.Context) GuestOsCustomizationArrayOutput
}

type GuestOsCustomizationArray []GuestOsCustomizationInput

func (GuestOsCustomizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GuestOsCustomization)(nil)).Elem()
}

func (i GuestOsCustomizationArray) ToGuestOsCustomizationArrayOutput() GuestOsCustomizationArrayOutput {
	return i.ToGuestOsCustomizationArrayOutputWithContext(context.Background())
}

func (i GuestOsCustomizationArray) ToGuestOsCustomizationArrayOutputWithContext(ctx context.Context) GuestOsCustomizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestOsCustomizationArrayOutput)
}

// GuestOsCustomizationMapInput is an input type that accepts GuestOsCustomizationMap and GuestOsCustomizationMapOutput values.
// You can construct a concrete instance of `GuestOsCustomizationMapInput` via:
//
//	GuestOsCustomizationMap{ "key": GuestOsCustomizationArgs{...} }
type GuestOsCustomizationMapInput interface {
	pulumi.Input

	ToGuestOsCustomizationMapOutput() GuestOsCustomizationMapOutput
	ToGuestOsCustomizationMapOutputWithContext(context.Context) GuestOsCustomizationMapOutput
}

type GuestOsCustomizationMap map[string]GuestOsCustomizationInput

func (GuestOsCustomizationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GuestOsCustomization)(nil)).Elem()
}

func (i GuestOsCustomizationMap) ToGuestOsCustomizationMapOutput() GuestOsCustomizationMapOutput {
	return i.ToGuestOsCustomizationMapOutputWithContext(context.Background())
}

func (i GuestOsCustomizationMap) ToGuestOsCustomizationMapOutputWithContext(ctx context.Context) GuestOsCustomizationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestOsCustomizationMapOutput)
}

type GuestOsCustomizationOutput struct{ *pulumi.OutputState }

func (GuestOsCustomizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestOsCustomization)(nil)).Elem()
}

func (o GuestOsCustomizationOutput) ToGuestOsCustomizationOutput() GuestOsCustomizationOutput {
	return o
}

func (o GuestOsCustomizationOutput) ToGuestOsCustomizationOutputWithContext(ctx context.Context) GuestOsCustomizationOutput {
	return o
}

// The number of last changed version to the customization specification.
func (o GuestOsCustomizationOutput) ChangeVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestOsCustomization) pulumi.StringOutput { return v.ChangeVersion }).(pulumi.StringOutput)
}

// The description for the customization specification.
func (o GuestOsCustomizationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestOsCustomization) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The time of last modification to the customization specification.
func (o GuestOsCustomizationOutput) LastUpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestOsCustomization) pulumi.StringOutput { return v.LastUpdateTime }).(pulumi.StringOutput)
}

// The name of the customization specification is the unique identifier per vCenter Server instance.
func (o GuestOsCustomizationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestOsCustomization) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Container object for the Guest OS properties about to be customized . See virtual machine customizations
func (o GuestOsCustomizationOutput) Spec() GuestOsCustomizationSpecOutput {
	return o.ApplyT(func(v *GuestOsCustomization) GuestOsCustomizationSpecOutput { return v.Spec }).(GuestOsCustomizationSpecOutput)
}

// The type of customization specification: One among: Windows, Linux.
func (o GuestOsCustomizationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestOsCustomization) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type GuestOsCustomizationArrayOutput struct{ *pulumi.OutputState }

func (GuestOsCustomizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GuestOsCustomization)(nil)).Elem()
}

func (o GuestOsCustomizationArrayOutput) ToGuestOsCustomizationArrayOutput() GuestOsCustomizationArrayOutput {
	return o
}

func (o GuestOsCustomizationArrayOutput) ToGuestOsCustomizationArrayOutputWithContext(ctx context.Context) GuestOsCustomizationArrayOutput {
	return o
}

func (o GuestOsCustomizationArrayOutput) Index(i pulumi.IntInput) GuestOsCustomizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GuestOsCustomization {
		return vs[0].([]*GuestOsCustomization)[vs[1].(int)]
	}).(GuestOsCustomizationOutput)
}

type GuestOsCustomizationMapOutput struct{ *pulumi.OutputState }

func (GuestOsCustomizationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GuestOsCustomization)(nil)).Elem()
}

func (o GuestOsCustomizationMapOutput) ToGuestOsCustomizationMapOutput() GuestOsCustomizationMapOutput {
	return o
}

func (o GuestOsCustomizationMapOutput) ToGuestOsCustomizationMapOutputWithContext(ctx context.Context) GuestOsCustomizationMapOutput {
	return o
}

func (o GuestOsCustomizationMapOutput) MapIndex(k pulumi.StringInput) GuestOsCustomizationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GuestOsCustomization {
		return vs[0].(map[string]*GuestOsCustomization)[vs[1].(string)]
	}).(GuestOsCustomizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GuestOsCustomizationInput)(nil)).Elem(), &GuestOsCustomization{})
	pulumi.RegisterInputType(reflect.TypeOf((*GuestOsCustomizationArrayInput)(nil)).Elem(), GuestOsCustomizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GuestOsCustomizationMapInput)(nil)).Elem(), GuestOsCustomizationMap{})
	pulumi.RegisterOutputType(GuestOsCustomizationOutput{})
	pulumi.RegisterOutputType(GuestOsCustomizationArrayOutput{})
	pulumi.RegisterOutputType(GuestOsCustomizationMapOutput{})
}

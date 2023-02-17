// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContentLibrary struct {
	pulumi.CustomResourceState

	// A description for the content library.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the content library.
	Name pulumi.StringOutput `pulumi:"name"`
	// Options to publish a local content library.
	Publication ContentLibraryPublicationOutput `pulumi:"publication"`
	// The managed object reference ID of the datastore on which to store the content library items.
	StorageBackings pulumi.StringArrayOutput `pulumi:"storageBackings"`
	// Options subscribe to a published content library.
	Subscription ContentLibrarySubscriptionPtrOutput `pulumi:"subscription"`
}

// NewContentLibrary registers a new resource with the given unique name, arguments, and options.
func NewContentLibrary(ctx *pulumi.Context,
	name string, args *ContentLibraryArgs, opts ...pulumi.ResourceOption) (*ContentLibrary, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StorageBackings == nil {
		return nil, errors.New("invalid value for required argument 'StorageBackings'")
	}
	var resource ContentLibrary
	err := ctx.RegisterResource("vsphere:index/contentLibrary:ContentLibrary", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContentLibrary gets an existing ContentLibrary resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContentLibrary(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContentLibraryState, opts ...pulumi.ResourceOption) (*ContentLibrary, error) {
	var resource ContentLibrary
	err := ctx.ReadResource("vsphere:index/contentLibrary:ContentLibrary", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContentLibrary resources.
type contentLibraryState struct {
	// A description for the content library.
	Description *string `pulumi:"description"`
	// The name of the content library.
	Name *string `pulumi:"name"`
	// Options to publish a local content library.
	Publication *ContentLibraryPublication `pulumi:"publication"`
	// The managed object reference ID of the datastore on which to store the content library items.
	StorageBackings []string `pulumi:"storageBackings"`
	// Options subscribe to a published content library.
	Subscription *ContentLibrarySubscription `pulumi:"subscription"`
}

type ContentLibraryState struct {
	// A description for the content library.
	Description pulumi.StringPtrInput
	// The name of the content library.
	Name pulumi.StringPtrInput
	// Options to publish a local content library.
	Publication ContentLibraryPublicationPtrInput
	// The managed object reference ID of the datastore on which to store the content library items.
	StorageBackings pulumi.StringArrayInput
	// Options subscribe to a published content library.
	Subscription ContentLibrarySubscriptionPtrInput
}

func (ContentLibraryState) ElementType() reflect.Type {
	return reflect.TypeOf((*contentLibraryState)(nil)).Elem()
}

type contentLibraryArgs struct {
	// A description for the content library.
	Description *string `pulumi:"description"`
	// The name of the content library.
	Name *string `pulumi:"name"`
	// Options to publish a local content library.
	Publication *ContentLibraryPublication `pulumi:"publication"`
	// The managed object reference ID of the datastore on which to store the content library items.
	StorageBackings []string `pulumi:"storageBackings"`
	// Options subscribe to a published content library.
	Subscription *ContentLibrarySubscription `pulumi:"subscription"`
}

// The set of arguments for constructing a ContentLibrary resource.
type ContentLibraryArgs struct {
	// A description for the content library.
	Description pulumi.StringPtrInput
	// The name of the content library.
	Name pulumi.StringPtrInput
	// Options to publish a local content library.
	Publication ContentLibraryPublicationPtrInput
	// The managed object reference ID of the datastore on which to store the content library items.
	StorageBackings pulumi.StringArrayInput
	// Options subscribe to a published content library.
	Subscription ContentLibrarySubscriptionPtrInput
}

func (ContentLibraryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contentLibraryArgs)(nil)).Elem()
}

type ContentLibraryInput interface {
	pulumi.Input

	ToContentLibraryOutput() ContentLibraryOutput
	ToContentLibraryOutputWithContext(ctx context.Context) ContentLibraryOutput
}

func (*ContentLibrary) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLibrary)(nil)).Elem()
}

func (i *ContentLibrary) ToContentLibraryOutput() ContentLibraryOutput {
	return i.ToContentLibraryOutputWithContext(context.Background())
}

func (i *ContentLibrary) ToContentLibraryOutputWithContext(ctx context.Context) ContentLibraryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLibraryOutput)
}

// ContentLibraryArrayInput is an input type that accepts ContentLibraryArray and ContentLibraryArrayOutput values.
// You can construct a concrete instance of `ContentLibraryArrayInput` via:
//
//	ContentLibraryArray{ ContentLibraryArgs{...} }
type ContentLibraryArrayInput interface {
	pulumi.Input

	ToContentLibraryArrayOutput() ContentLibraryArrayOutput
	ToContentLibraryArrayOutputWithContext(context.Context) ContentLibraryArrayOutput
}

type ContentLibraryArray []ContentLibraryInput

func (ContentLibraryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContentLibrary)(nil)).Elem()
}

func (i ContentLibraryArray) ToContentLibraryArrayOutput() ContentLibraryArrayOutput {
	return i.ToContentLibraryArrayOutputWithContext(context.Background())
}

func (i ContentLibraryArray) ToContentLibraryArrayOutputWithContext(ctx context.Context) ContentLibraryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLibraryArrayOutput)
}

// ContentLibraryMapInput is an input type that accepts ContentLibraryMap and ContentLibraryMapOutput values.
// You can construct a concrete instance of `ContentLibraryMapInput` via:
//
//	ContentLibraryMap{ "key": ContentLibraryArgs{...} }
type ContentLibraryMapInput interface {
	pulumi.Input

	ToContentLibraryMapOutput() ContentLibraryMapOutput
	ToContentLibraryMapOutputWithContext(context.Context) ContentLibraryMapOutput
}

type ContentLibraryMap map[string]ContentLibraryInput

func (ContentLibraryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContentLibrary)(nil)).Elem()
}

func (i ContentLibraryMap) ToContentLibraryMapOutput() ContentLibraryMapOutput {
	return i.ToContentLibraryMapOutputWithContext(context.Background())
}

func (i ContentLibraryMap) ToContentLibraryMapOutputWithContext(ctx context.Context) ContentLibraryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLibraryMapOutput)
}

type ContentLibraryOutput struct{ *pulumi.OutputState }

func (ContentLibraryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLibrary)(nil)).Elem()
}

func (o ContentLibraryOutput) ToContentLibraryOutput() ContentLibraryOutput {
	return o
}

func (o ContentLibraryOutput) ToContentLibraryOutputWithContext(ctx context.Context) ContentLibraryOutput {
	return o
}

// A description for the content library.
func (o ContentLibraryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLibrary) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the content library.
func (o ContentLibraryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentLibrary) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Options to publish a local content library.
func (o ContentLibraryOutput) Publication() ContentLibraryPublicationOutput {
	return o.ApplyT(func(v *ContentLibrary) ContentLibraryPublicationOutput { return v.Publication }).(ContentLibraryPublicationOutput)
}

// The managed object reference ID of the datastore on which to store the content library items.
func (o ContentLibraryOutput) StorageBackings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContentLibrary) pulumi.StringArrayOutput { return v.StorageBackings }).(pulumi.StringArrayOutput)
}

// Options subscribe to a published content library.
func (o ContentLibraryOutput) Subscription() ContentLibrarySubscriptionPtrOutput {
	return o.ApplyT(func(v *ContentLibrary) ContentLibrarySubscriptionPtrOutput { return v.Subscription }).(ContentLibrarySubscriptionPtrOutput)
}

type ContentLibraryArrayOutput struct{ *pulumi.OutputState }

func (ContentLibraryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContentLibrary)(nil)).Elem()
}

func (o ContentLibraryArrayOutput) ToContentLibraryArrayOutput() ContentLibraryArrayOutput {
	return o
}

func (o ContentLibraryArrayOutput) ToContentLibraryArrayOutputWithContext(ctx context.Context) ContentLibraryArrayOutput {
	return o
}

func (o ContentLibraryArrayOutput) Index(i pulumi.IntInput) ContentLibraryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContentLibrary {
		return vs[0].([]*ContentLibrary)[vs[1].(int)]
	}).(ContentLibraryOutput)
}

type ContentLibraryMapOutput struct{ *pulumi.OutputState }

func (ContentLibraryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContentLibrary)(nil)).Elem()
}

func (o ContentLibraryMapOutput) ToContentLibraryMapOutput() ContentLibraryMapOutput {
	return o
}

func (o ContentLibraryMapOutput) ToContentLibraryMapOutputWithContext(ctx context.Context) ContentLibraryMapOutput {
	return o
}

func (o ContentLibraryMapOutput) MapIndex(k pulumi.StringInput) ContentLibraryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContentLibrary {
		return vs[0].(map[string]*ContentLibrary)[vs[1].(string)]
	}).(ContentLibraryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContentLibraryInput)(nil)).Elem(), &ContentLibrary{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentLibraryArrayInput)(nil)).Elem(), ContentLibraryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentLibraryMapInput)(nil)).Elem(), ContentLibraryMap{})
	pulumi.RegisterOutputType(ContentLibraryOutput{})
	pulumi.RegisterOutputType(ContentLibraryArrayOutput{})
	pulumi.RegisterOutputType(ContentLibraryMapOutput{})
}

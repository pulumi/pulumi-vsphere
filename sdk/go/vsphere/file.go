// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type File struct {
	pulumi.CustomResourceState

	// Create directories in `destinationFile`
	// path parameter on first apply if any are missing for copy operation.
	CreateDirectories pulumi.BoolPtrOutput `pulumi:"createDirectories"`
	// The name of a datacenter to which the file will be
	// uploaded.
	Datacenter pulumi.StringPtrOutput `pulumi:"datacenter"`
	// The name of the datastore to which to upload the
	// file.
	Datastore pulumi.StringOutput `pulumi:"datastore"`
	// The path to where the file should be uploaded
	// or copied to on the destination `datastore` in vSphere.
	DestinationFile pulumi.StringOutput `pulumi:"destinationFile"`
	// The name of a datacenter from which the file
	// will be copied. Forces a new resource if changed.
	SourceDatacenter pulumi.StringPtrOutput `pulumi:"sourceDatacenter"`
	// The name of the datastore from which file will
	// be copied. Forces a new resource if changed.
	SourceDatastore pulumi.StringPtrOutput `pulumi:"sourceDatastore"`
	SourceFile      pulumi.StringOutput    `pulumi:"sourceFile"`
}

// NewFile registers a new resource with the given unique name, arguments, and options.
func NewFile(ctx *pulumi.Context,
	name string, args *FileArgs, opts ...pulumi.ResourceOption) (*File, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Datastore == nil {
		return nil, errors.New("invalid value for required argument 'Datastore'")
	}
	if args.DestinationFile == nil {
		return nil, errors.New("invalid value for required argument 'DestinationFile'")
	}
	if args.SourceFile == nil {
		return nil, errors.New("invalid value for required argument 'SourceFile'")
	}
	var resource File
	err := ctx.RegisterResource("vsphere:index/file:File", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFile gets an existing File resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileState, opts ...pulumi.ResourceOption) (*File, error) {
	var resource File
	err := ctx.ReadResource("vsphere:index/file:File", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering File resources.
type fileState struct {
	// Create directories in `destinationFile`
	// path parameter on first apply if any are missing for copy operation.
	CreateDirectories *bool `pulumi:"createDirectories"`
	// The name of a datacenter to which the file will be
	// uploaded.
	Datacenter *string `pulumi:"datacenter"`
	// The name of the datastore to which to upload the
	// file.
	Datastore *string `pulumi:"datastore"`
	// The path to where the file should be uploaded
	// or copied to on the destination `datastore` in vSphere.
	DestinationFile *string `pulumi:"destinationFile"`
	// The name of a datacenter from which the file
	// will be copied. Forces a new resource if changed.
	SourceDatacenter *string `pulumi:"sourceDatacenter"`
	// The name of the datastore from which file will
	// be copied. Forces a new resource if changed.
	SourceDatastore *string `pulumi:"sourceDatastore"`
	SourceFile      *string `pulumi:"sourceFile"`
}

type FileState struct {
	// Create directories in `destinationFile`
	// path parameter on first apply if any are missing for copy operation.
	CreateDirectories pulumi.BoolPtrInput
	// The name of a datacenter to which the file will be
	// uploaded.
	Datacenter pulumi.StringPtrInput
	// The name of the datastore to which to upload the
	// file.
	Datastore pulumi.StringPtrInput
	// The path to where the file should be uploaded
	// or copied to on the destination `datastore` in vSphere.
	DestinationFile pulumi.StringPtrInput
	// The name of a datacenter from which the file
	// will be copied. Forces a new resource if changed.
	SourceDatacenter pulumi.StringPtrInput
	// The name of the datastore from which file will
	// be copied. Forces a new resource if changed.
	SourceDatastore pulumi.StringPtrInput
	SourceFile      pulumi.StringPtrInput
}

func (FileState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileState)(nil)).Elem()
}

type fileArgs struct {
	// Create directories in `destinationFile`
	// path parameter on first apply if any are missing for copy operation.
	CreateDirectories *bool `pulumi:"createDirectories"`
	// The name of a datacenter to which the file will be
	// uploaded.
	Datacenter *string `pulumi:"datacenter"`
	// The name of the datastore to which to upload the
	// file.
	Datastore string `pulumi:"datastore"`
	// The path to where the file should be uploaded
	// or copied to on the destination `datastore` in vSphere.
	DestinationFile string `pulumi:"destinationFile"`
	// The name of a datacenter from which the file
	// will be copied. Forces a new resource if changed.
	SourceDatacenter *string `pulumi:"sourceDatacenter"`
	// The name of the datastore from which file will
	// be copied. Forces a new resource if changed.
	SourceDatastore *string `pulumi:"sourceDatastore"`
	SourceFile      string  `pulumi:"sourceFile"`
}

// The set of arguments for constructing a File resource.
type FileArgs struct {
	// Create directories in `destinationFile`
	// path parameter on first apply if any are missing for copy operation.
	CreateDirectories pulumi.BoolPtrInput
	// The name of a datacenter to which the file will be
	// uploaded.
	Datacenter pulumi.StringPtrInput
	// The name of the datastore to which to upload the
	// file.
	Datastore pulumi.StringInput
	// The path to where the file should be uploaded
	// or copied to on the destination `datastore` in vSphere.
	DestinationFile pulumi.StringInput
	// The name of a datacenter from which the file
	// will be copied. Forces a new resource if changed.
	SourceDatacenter pulumi.StringPtrInput
	// The name of the datastore from which file will
	// be copied. Forces a new resource if changed.
	SourceDatastore pulumi.StringPtrInput
	SourceFile      pulumi.StringInput
}

func (FileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileArgs)(nil)).Elem()
}

type FileInput interface {
	pulumi.Input

	ToFileOutput() FileOutput
	ToFileOutputWithContext(ctx context.Context) FileOutput
}

func (*File) ElementType() reflect.Type {
	return reflect.TypeOf((**File)(nil)).Elem()
}

func (i *File) ToFileOutput() FileOutput {
	return i.ToFileOutputWithContext(context.Background())
}

func (i *File) ToFileOutputWithContext(ctx context.Context) FileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileOutput)
}

// FileArrayInput is an input type that accepts FileArray and FileArrayOutput values.
// You can construct a concrete instance of `FileArrayInput` via:
//
//          FileArray{ FileArgs{...} }
type FileArrayInput interface {
	pulumi.Input

	ToFileArrayOutput() FileArrayOutput
	ToFileArrayOutputWithContext(context.Context) FileArrayOutput
}

type FileArray []FileInput

func (FileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*File)(nil)).Elem()
}

func (i FileArray) ToFileArrayOutput() FileArrayOutput {
	return i.ToFileArrayOutputWithContext(context.Background())
}

func (i FileArray) ToFileArrayOutputWithContext(ctx context.Context) FileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileArrayOutput)
}

// FileMapInput is an input type that accepts FileMap and FileMapOutput values.
// You can construct a concrete instance of `FileMapInput` via:
//
//          FileMap{ "key": FileArgs{...} }
type FileMapInput interface {
	pulumi.Input

	ToFileMapOutput() FileMapOutput
	ToFileMapOutputWithContext(context.Context) FileMapOutput
}

type FileMap map[string]FileInput

func (FileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*File)(nil)).Elem()
}

func (i FileMap) ToFileMapOutput() FileMapOutput {
	return i.ToFileMapOutputWithContext(context.Background())
}

func (i FileMap) ToFileMapOutputWithContext(ctx context.Context) FileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileMapOutput)
}

type FileOutput struct{ *pulumi.OutputState }

func (FileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**File)(nil)).Elem()
}

func (o FileOutput) ToFileOutput() FileOutput {
	return o
}

func (o FileOutput) ToFileOutputWithContext(ctx context.Context) FileOutput {
	return o
}

type FileArrayOutput struct{ *pulumi.OutputState }

func (FileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*File)(nil)).Elem()
}

func (o FileArrayOutput) ToFileArrayOutput() FileArrayOutput {
	return o
}

func (o FileArrayOutput) ToFileArrayOutputWithContext(ctx context.Context) FileArrayOutput {
	return o
}

func (o FileArrayOutput) Index(i pulumi.IntInput) FileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *File {
		return vs[0].([]*File)[vs[1].(int)]
	}).(FileOutput)
}

type FileMapOutput struct{ *pulumi.OutputState }

func (FileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*File)(nil)).Elem()
}

func (o FileMapOutput) ToFileMapOutput() FileMapOutput {
	return o
}

func (o FileMapOutput) ToFileMapOutputWithContext(ctx context.Context) FileMapOutput {
	return o
}

func (o FileMapOutput) MapIndex(k pulumi.StringInput) FileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *File {
		return vs[0].(map[string]*File)[vs[1].(string)]
	}).(FileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FileInput)(nil)).Elem(), &File{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileArrayInput)(nil)).Elem(), FileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileMapInput)(nil)).Elem(), FileMap{})
	pulumi.RegisterOutputType(FileOutput{})
	pulumi.RegisterOutputType(FileArrayOutput{})
	pulumi.RegisterOutputType(FileMapOutput{})
}

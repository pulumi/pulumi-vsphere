// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EntityPermissions struct {
	pulumi.CustomResourceState

	// The managed object id (uuid for some entities) on which permissions are to be created.
	EntityId pulumi.StringOutput `pulumi:"entityId"`
	// The managed object type, types can be found in the managed object type section
	// [here](https://code.vmware.com/apis/968/vsphere).
	EntityType pulumi.StringOutput `pulumi:"entityType"`
	// The permissions to be given on this entity. Keep the permissions sorted
	// alphabetically on `userOrGroup` for a better user experience.
	Permissions EntityPermissionsPermissionArrayOutput `pulumi:"permissions"`
}

// NewEntityPermissions registers a new resource with the given unique name, arguments, and options.
func NewEntityPermissions(ctx *pulumi.Context,
	name string, args *EntityPermissionsArgs, opts ...pulumi.ResourceOption) (*EntityPermissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EntityId == nil {
		return nil, errors.New("invalid value for required argument 'EntityId'")
	}
	if args.EntityType == nil {
		return nil, errors.New("invalid value for required argument 'EntityType'")
	}
	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	var resource EntityPermissions
	err := ctx.RegisterResource("vsphere:index/entityPermissions:EntityPermissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntityPermissions gets an existing EntityPermissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntityPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntityPermissionsState, opts ...pulumi.ResourceOption) (*EntityPermissions, error) {
	var resource EntityPermissions
	err := ctx.ReadResource("vsphere:index/entityPermissions:EntityPermissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EntityPermissions resources.
type entityPermissionsState struct {
	// The managed object id (uuid for some entities) on which permissions are to be created.
	EntityId *string `pulumi:"entityId"`
	// The managed object type, types can be found in the managed object type section
	// [here](https://code.vmware.com/apis/968/vsphere).
	EntityType *string `pulumi:"entityType"`
	// The permissions to be given on this entity. Keep the permissions sorted
	// alphabetically on `userOrGroup` for a better user experience.
	Permissions []EntityPermissionsPermission `pulumi:"permissions"`
}

type EntityPermissionsState struct {
	// The managed object id (uuid for some entities) on which permissions are to be created.
	EntityId pulumi.StringPtrInput
	// The managed object type, types can be found in the managed object type section
	// [here](https://code.vmware.com/apis/968/vsphere).
	EntityType pulumi.StringPtrInput
	// The permissions to be given on this entity. Keep the permissions sorted
	// alphabetically on `userOrGroup` for a better user experience.
	Permissions EntityPermissionsPermissionArrayInput
}

func (EntityPermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*entityPermissionsState)(nil)).Elem()
}

type entityPermissionsArgs struct {
	// The managed object id (uuid for some entities) on which permissions are to be created.
	EntityId string `pulumi:"entityId"`
	// The managed object type, types can be found in the managed object type section
	// [here](https://code.vmware.com/apis/968/vsphere).
	EntityType string `pulumi:"entityType"`
	// The permissions to be given on this entity. Keep the permissions sorted
	// alphabetically on `userOrGroup` for a better user experience.
	Permissions []EntityPermissionsPermission `pulumi:"permissions"`
}

// The set of arguments for constructing a EntityPermissions resource.
type EntityPermissionsArgs struct {
	// The managed object id (uuid for some entities) on which permissions are to be created.
	EntityId pulumi.StringInput
	// The managed object type, types can be found in the managed object type section
	// [here](https://code.vmware.com/apis/968/vsphere).
	EntityType pulumi.StringInput
	// The permissions to be given on this entity. Keep the permissions sorted
	// alphabetically on `userOrGroup` for a better user experience.
	Permissions EntityPermissionsPermissionArrayInput
}

func (EntityPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entityPermissionsArgs)(nil)).Elem()
}

type EntityPermissionsInput interface {
	pulumi.Input

	ToEntityPermissionsOutput() EntityPermissionsOutput
	ToEntityPermissionsOutputWithContext(ctx context.Context) EntityPermissionsOutput
}

func (*EntityPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityPermissions)(nil))
}

func (i *EntityPermissions) ToEntityPermissionsOutput() EntityPermissionsOutput {
	return i.ToEntityPermissionsOutputWithContext(context.Background())
}

func (i *EntityPermissions) ToEntityPermissionsOutputWithContext(ctx context.Context) EntityPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityPermissionsOutput)
}

func (i *EntityPermissions) ToEntityPermissionsPtrOutput() EntityPermissionsPtrOutput {
	return i.ToEntityPermissionsPtrOutputWithContext(context.Background())
}

func (i *EntityPermissions) ToEntityPermissionsPtrOutputWithContext(ctx context.Context) EntityPermissionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityPermissionsPtrOutput)
}

type EntityPermissionsPtrInput interface {
	pulumi.Input

	ToEntityPermissionsPtrOutput() EntityPermissionsPtrOutput
	ToEntityPermissionsPtrOutputWithContext(ctx context.Context) EntityPermissionsPtrOutput
}

type entityPermissionsPtrType EntityPermissionsArgs

func (*entityPermissionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityPermissions)(nil))
}

func (i *entityPermissionsPtrType) ToEntityPermissionsPtrOutput() EntityPermissionsPtrOutput {
	return i.ToEntityPermissionsPtrOutputWithContext(context.Background())
}

func (i *entityPermissionsPtrType) ToEntityPermissionsPtrOutputWithContext(ctx context.Context) EntityPermissionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityPermissionsPtrOutput)
}

// EntityPermissionsArrayInput is an input type that accepts EntityPermissionsArray and EntityPermissionsArrayOutput values.
// You can construct a concrete instance of `EntityPermissionsArrayInput` via:
//
//          EntityPermissionsArray{ EntityPermissionsArgs{...} }
type EntityPermissionsArrayInput interface {
	pulumi.Input

	ToEntityPermissionsArrayOutput() EntityPermissionsArrayOutput
	ToEntityPermissionsArrayOutputWithContext(context.Context) EntityPermissionsArrayOutput
}

type EntityPermissionsArray []EntityPermissionsInput

func (EntityPermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EntityPermissions)(nil)).Elem()
}

func (i EntityPermissionsArray) ToEntityPermissionsArrayOutput() EntityPermissionsArrayOutput {
	return i.ToEntityPermissionsArrayOutputWithContext(context.Background())
}

func (i EntityPermissionsArray) ToEntityPermissionsArrayOutputWithContext(ctx context.Context) EntityPermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityPermissionsArrayOutput)
}

// EntityPermissionsMapInput is an input type that accepts EntityPermissionsMap and EntityPermissionsMapOutput values.
// You can construct a concrete instance of `EntityPermissionsMapInput` via:
//
//          EntityPermissionsMap{ "key": EntityPermissionsArgs{...} }
type EntityPermissionsMapInput interface {
	pulumi.Input

	ToEntityPermissionsMapOutput() EntityPermissionsMapOutput
	ToEntityPermissionsMapOutputWithContext(context.Context) EntityPermissionsMapOutput
}

type EntityPermissionsMap map[string]EntityPermissionsInput

func (EntityPermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EntityPermissions)(nil)).Elem()
}

func (i EntityPermissionsMap) ToEntityPermissionsMapOutput() EntityPermissionsMapOutput {
	return i.ToEntityPermissionsMapOutputWithContext(context.Background())
}

func (i EntityPermissionsMap) ToEntityPermissionsMapOutputWithContext(ctx context.Context) EntityPermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityPermissionsMapOutput)
}

type EntityPermissionsOutput struct{ *pulumi.OutputState }

func (EntityPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityPermissions)(nil))
}

func (o EntityPermissionsOutput) ToEntityPermissionsOutput() EntityPermissionsOutput {
	return o
}

func (o EntityPermissionsOutput) ToEntityPermissionsOutputWithContext(ctx context.Context) EntityPermissionsOutput {
	return o
}

func (o EntityPermissionsOutput) ToEntityPermissionsPtrOutput() EntityPermissionsPtrOutput {
	return o.ToEntityPermissionsPtrOutputWithContext(context.Background())
}

func (o EntityPermissionsOutput) ToEntityPermissionsPtrOutputWithContext(ctx context.Context) EntityPermissionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EntityPermissions) *EntityPermissions {
		return &v
	}).(EntityPermissionsPtrOutput)
}

type EntityPermissionsPtrOutput struct{ *pulumi.OutputState }

func (EntityPermissionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityPermissions)(nil))
}

func (o EntityPermissionsPtrOutput) ToEntityPermissionsPtrOutput() EntityPermissionsPtrOutput {
	return o
}

func (o EntityPermissionsPtrOutput) ToEntityPermissionsPtrOutputWithContext(ctx context.Context) EntityPermissionsPtrOutput {
	return o
}

func (o EntityPermissionsPtrOutput) Elem() EntityPermissionsOutput {
	return o.ApplyT(func(v *EntityPermissions) EntityPermissions {
		if v != nil {
			return *v
		}
		var ret EntityPermissions
		return ret
	}).(EntityPermissionsOutput)
}

type EntityPermissionsArrayOutput struct{ *pulumi.OutputState }

func (EntityPermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EntityPermissions)(nil))
}

func (o EntityPermissionsArrayOutput) ToEntityPermissionsArrayOutput() EntityPermissionsArrayOutput {
	return o
}

func (o EntityPermissionsArrayOutput) ToEntityPermissionsArrayOutputWithContext(ctx context.Context) EntityPermissionsArrayOutput {
	return o
}

func (o EntityPermissionsArrayOutput) Index(i pulumi.IntInput) EntityPermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EntityPermissions {
		return vs[0].([]EntityPermissions)[vs[1].(int)]
	}).(EntityPermissionsOutput)
}

type EntityPermissionsMapOutput struct{ *pulumi.OutputState }

func (EntityPermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EntityPermissions)(nil))
}

func (o EntityPermissionsMapOutput) ToEntityPermissionsMapOutput() EntityPermissionsMapOutput {
	return o
}

func (o EntityPermissionsMapOutput) ToEntityPermissionsMapOutputWithContext(ctx context.Context) EntityPermissionsMapOutput {
	return o
}

func (o EntityPermissionsMapOutput) MapIndex(k pulumi.StringInput) EntityPermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EntityPermissions {
		return vs[0].(map[string]EntityPermissions)[vs[1].(string)]
	}).(EntityPermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EntityPermissionsInput)(nil)).Elem(), &EntityPermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntityPermissionsPtrInput)(nil)).Elem(), &EntityPermissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntityPermissionsArrayInput)(nil)).Elem(), EntityPermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntityPermissionsMapInput)(nil)).Elem(), EntityPermissionsMap{})
	pulumi.RegisterOutputType(EntityPermissionsOutput{})
	pulumi.RegisterOutputType(EntityPermissionsPtrOutput{})
	pulumi.RegisterOutputType(EntityPermissionsArrayOutput{})
	pulumi.RegisterOutputType(EntityPermissionsMapOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vsphere

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a VMware vSphere license resource. This can be used to add and remove license keys.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/license.html.markdown.
type License struct {
	pulumi.CustomResourceState

	// The product edition of the license key.
	EditionKey pulumi.StringOutput `pulumi:"editionKey"`
	// A map of key/value pairs to be attached as labels (tags) to the license key.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The license key to add.
	LicenseKey pulumi.StringOutput `pulumi:"licenseKey"`
	// The display name for the license.
	Name pulumi.StringOutput `pulumi:"name"`
	// Total number of units (example: CPUs) contained in the license.
	Total pulumi.IntOutput `pulumi:"total"`
	// The number of units (example: CPUs) assigned to this license.
	Used pulumi.IntOutput `pulumi:"used"`
}

// NewLicense registers a new resource with the given unique name, arguments, and options.
func NewLicense(ctx *pulumi.Context,
	name string, args *LicenseArgs, opts ...pulumi.ResourceOption) (*License, error) {
	if args == nil || args.LicenseKey == nil {
		return nil, errors.New("missing required argument 'LicenseKey'")
	}
	if args == nil {
		args = &LicenseArgs{}
	}
	var resource License
	err := ctx.RegisterResource("vsphere:index/license:License", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLicense gets an existing License resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLicense(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LicenseState, opts ...pulumi.ResourceOption) (*License, error) {
	var resource License
	err := ctx.ReadResource("vsphere:index/license:License", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering License resources.
type licenseState struct {
	// The product edition of the license key.
	EditionKey *string `pulumi:"editionKey"`
	// A map of key/value pairs to be attached as labels (tags) to the license key.
	Labels map[string]string `pulumi:"labels"`
	// The license key to add.
	LicenseKey *string `pulumi:"licenseKey"`
	// The display name for the license.
	Name *string `pulumi:"name"`
	// Total number of units (example: CPUs) contained in the license.
	Total *int `pulumi:"total"`
	// The number of units (example: CPUs) assigned to this license.
	Used *int `pulumi:"used"`
}

type LicenseState struct {
	// The product edition of the license key.
	EditionKey pulumi.StringPtrInput
	// A map of key/value pairs to be attached as labels (tags) to the license key.
	Labels pulumi.StringMapInput
	// The license key to add.
	LicenseKey pulumi.StringPtrInput
	// The display name for the license.
	Name pulumi.StringPtrInput
	// Total number of units (example: CPUs) contained in the license.
	Total pulumi.IntPtrInput
	// The number of units (example: CPUs) assigned to this license.
	Used pulumi.IntPtrInput
}

func (LicenseState) ElementType() reflect.Type {
	return reflect.TypeOf((*licenseState)(nil)).Elem()
}

type licenseArgs struct {
	// A map of key/value pairs to be attached as labels (tags) to the license key.
	Labels map[string]string `pulumi:"labels"`
	// The license key to add.
	LicenseKey string `pulumi:"licenseKey"`
}

// The set of arguments for constructing a License resource.
type LicenseArgs struct {
	// A map of key/value pairs to be attached as labels (tags) to the license key.
	Labels pulumi.StringMapInput
	// The license key to add.
	LicenseKey pulumi.StringInput
}

func (LicenseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*licenseArgs)(nil)).Elem()
}

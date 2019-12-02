// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a VMware vSphere license resource. This can be used to add and remove license keys.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/license.html.markdown.
type License struct {
	s *pulumi.ResourceState
}

// NewLicense registers a new resource with the given unique name, arguments, and options.
func NewLicense(ctx *pulumi.Context,
	name string, args *LicenseArgs, opts ...pulumi.ResourceOpt) (*License, error) {
	if args == nil || args.LicenseKey == nil {
		return nil, errors.New("missing required argument 'LicenseKey'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["labels"] = nil
		inputs["licenseKey"] = nil
	} else {
		inputs["labels"] = args.Labels
		inputs["licenseKey"] = args.LicenseKey
	}
	inputs["editionKey"] = nil
	inputs["name"] = nil
	inputs["total"] = nil
	inputs["used"] = nil
	s, err := ctx.RegisterResource("vsphere:index/license:License", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &License{s: s}, nil
}

// GetLicense gets an existing License resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLicense(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LicenseState, opts ...pulumi.ResourceOpt) (*License, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["editionKey"] = state.EditionKey
		inputs["labels"] = state.Labels
		inputs["licenseKey"] = state.LicenseKey
		inputs["name"] = state.Name
		inputs["total"] = state.Total
		inputs["used"] = state.Used
	}
	s, err := ctx.ReadResource("vsphere:index/license:License", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &License{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *License) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *License) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The product edition of the license key.
func (r *License) EditionKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["editionKey"])
}

// A map of key/value pairs to be attached as labels (tags) to the license key.
func (r *License) Labels() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["labels"])
}

// The license key to add.
func (r *License) LicenseKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["licenseKey"])
}

// The display name for the license.
func (r *License) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Total number of units (example: CPUs) contained in the license.
func (r *License) Total() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["total"])
}

// The number of units (example: CPUs) assigned to this license.
func (r *License) Used() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["used"])
}

// Input properties used for looking up and filtering License resources.
type LicenseState struct {
	// The product edition of the license key.
	EditionKey interface{}
	// A map of key/value pairs to be attached as labels (tags) to the license key.
	Labels interface{}
	// The license key to add.
	LicenseKey interface{}
	// The display name for the license.
	Name interface{}
	// Total number of units (example: CPUs) contained in the license.
	Total interface{}
	// The number of units (example: CPUs) assigned to this license.
	Used interface{}
}

// The set of arguments for constructing a License resource.
type LicenseArgs struct {
	// A map of key/value pairs to be attached as labels (tags) to the license key.
	Labels interface{}
	// The license key to add.
	LicenseKey interface{}
}

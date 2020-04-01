// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vsphere

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `.DrsVmOverride` resource can be used to add a DRS override to a
// cluster for a specific virtual machine. With this resource, one can enable or
// disable DRS and control the automation level for a single virtual machine
// without affecting the rest of the cluster.
//
// For more information on vSphere clusters and DRS, see [this
// page][ref-vsphere-drs-clusters].
//
// [ref-vsphere-drs-clusters]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-8ACF3502-5314-469F-8CC9-4A9BD5925BC2.html
//
// > **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
//
// > **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/drs_vm_override.html.markdown.
type DrsVmOverride struct {
	pulumi.CustomResourceState

	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringOutput `pulumi:"computeClusterId"`
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	DrsAutomationLevel pulumi.StringPtrOutput `pulumi:"drsAutomationLevel"`
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled pulumi.BoolPtrOutput `pulumi:"drsEnabled"`
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId pulumi.StringOutput `pulumi:"virtualMachineId"`
}

// NewDrsVmOverride registers a new resource with the given unique name, arguments, and options.
func NewDrsVmOverride(ctx *pulumi.Context,
	name string, args *DrsVmOverrideArgs, opts ...pulumi.ResourceOption) (*DrsVmOverride, error) {
	if args == nil || args.ComputeClusterId == nil {
		return nil, errors.New("missing required argument 'ComputeClusterId'")
	}
	if args == nil || args.VirtualMachineId == nil {
		return nil, errors.New("missing required argument 'VirtualMachineId'")
	}
	if args == nil {
		args = &DrsVmOverrideArgs{}
	}
	var resource DrsVmOverride
	err := ctx.RegisterResource("vsphere:index/drsVmOverride:DrsVmOverride", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDrsVmOverride gets an existing DrsVmOverride resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDrsVmOverride(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DrsVmOverrideState, opts ...pulumi.ResourceOption) (*DrsVmOverride, error) {
	var resource DrsVmOverride
	err := ctx.ReadResource("vsphere:index/drsVmOverride:DrsVmOverride", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DrsVmOverride resources.
type drsVmOverrideState struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId *string `pulumi:"computeClusterId"`
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	DrsAutomationLevel *string `pulumi:"drsAutomationLevel"`
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled *bool `pulumi:"drsEnabled"`
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId *string `pulumi:"virtualMachineId"`
}

type DrsVmOverrideState struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringPtrInput
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	DrsAutomationLevel pulumi.StringPtrInput
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled pulumi.BoolPtrInput
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId pulumi.StringPtrInput
}

func (DrsVmOverrideState) ElementType() reflect.Type {
	return reflect.TypeOf((*drsVmOverrideState)(nil)).Elem()
}

type drsVmOverrideArgs struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId string `pulumi:"computeClusterId"`
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	DrsAutomationLevel *string `pulumi:"drsAutomationLevel"`
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled *bool `pulumi:"drsEnabled"`
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId string `pulumi:"virtualMachineId"`
}

// The set of arguments for constructing a DrsVmOverride resource.
type DrsVmOverrideArgs struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringInput
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	DrsAutomationLevel pulumi.StringPtrInput
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled pulumi.BoolPtrInput
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId pulumi.StringInput
}

func (DrsVmOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*drsVmOverrideArgs)(nil)).Elem()
}

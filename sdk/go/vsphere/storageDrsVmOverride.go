// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `.StorageDrsVmOverride` resource can be used to add a Storage DRS
// override to a datastore cluster for a specific virtual machine. With this
// resource, one can enable or disable Storage DRS, and control the automation
// level and disk affinity for a single virtual machine without affecting the rest
// of the datastore cluster.
// 
// For more information on vSphere datastore clusters and Storage DRS, see [this
// page][ref-vsphere-datastore-clusters].
// 
// [ref-vsphere-datastore-clusters]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-598DF695-107E-406B-9C95-0AF961FC227A.html
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/storage_drs_vm_override.html.markdown.
type StorageDrsVmOverride struct {
	s *pulumi.ResourceState
}

// NewStorageDrsVmOverride registers a new resource with the given unique name, arguments, and options.
func NewStorageDrsVmOverride(ctx *pulumi.Context,
	name string, args *StorageDrsVmOverrideArgs, opts ...pulumi.ResourceOpt) (*StorageDrsVmOverride, error) {
	if args == nil || args.DatastoreClusterId == nil {
		return nil, errors.New("missing required argument 'DatastoreClusterId'")
	}
	if args == nil || args.VirtualMachineId == nil {
		return nil, errors.New("missing required argument 'VirtualMachineId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["datastoreClusterId"] = nil
		inputs["sdrsAutomationLevel"] = nil
		inputs["sdrsEnabled"] = nil
		inputs["sdrsIntraVmAffinity"] = nil
		inputs["virtualMachineId"] = nil
	} else {
		inputs["datastoreClusterId"] = args.DatastoreClusterId
		inputs["sdrsAutomationLevel"] = args.SdrsAutomationLevel
		inputs["sdrsEnabled"] = args.SdrsEnabled
		inputs["sdrsIntraVmAffinity"] = args.SdrsIntraVmAffinity
		inputs["virtualMachineId"] = args.VirtualMachineId
	}
	s, err := ctx.RegisterResource("vsphere:index/storageDrsVmOverride:StorageDrsVmOverride", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &StorageDrsVmOverride{s: s}, nil
}

// GetStorageDrsVmOverride gets an existing StorageDrsVmOverride resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageDrsVmOverride(ctx *pulumi.Context,
	name string, id pulumi.ID, state *StorageDrsVmOverrideState, opts ...pulumi.ResourceOpt) (*StorageDrsVmOverride, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["datastoreClusterId"] = state.DatastoreClusterId
		inputs["sdrsAutomationLevel"] = state.SdrsAutomationLevel
		inputs["sdrsEnabled"] = state.SdrsEnabled
		inputs["sdrsIntraVmAffinity"] = state.SdrsIntraVmAffinity
		inputs["virtualMachineId"] = state.VirtualMachineId
	}
	s, err := ctx.ReadResource("vsphere:index/storageDrsVmOverride:StorageDrsVmOverride", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &StorageDrsVmOverride{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *StorageDrsVmOverride) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *StorageDrsVmOverride) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The [managed object reference
// ID][docs-about-morefs] of the datastore cluster to put the override in.
// Forces a new resource if changed.
func (r *StorageDrsVmOverride) DatastoreClusterId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["datastoreClusterId"])
}

// Overrides any Storage DRS automation
// levels for this virtual machine. Can be one of `automated` or `manual`. When
// not specified, the datastore cluster's settings are used according to the
// [specific SDRS subsystem][tf-vsphere-datastore-cluster-sdrs-levels].
func (r *StorageDrsVmOverride) SdrsAutomationLevel() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sdrsAutomationLevel"])
}

// Overrides the default Storage DRS setting for
// this virtual machine. When not specified, the datastore cluster setting is
// used.
func (r *StorageDrsVmOverride) SdrsEnabled() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sdrsEnabled"])
}

// Overrides the intra-VM affinity setting
// for this virtual machine. When `true`, all disks for this virtual machine
// will be kept on the same datastore. When `false`, Storage DRS may locate
// individual disks on different datastores if it helps satisfy cluster
// requirements. When not specified, the datastore cluster's settings are used.
func (r *StorageDrsVmOverride) SdrsIntraVmAffinity() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sdrsIntraVmAffinity"])
}

// The UUID of the virtual machine to create
// the override for.  Forces a new resource if changed.
func (r *StorageDrsVmOverride) VirtualMachineId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["virtualMachineId"])
}

// Input properties used for looking up and filtering StorageDrsVmOverride resources.
type StorageDrsVmOverrideState struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the datastore cluster to put the override in.
	// Forces a new resource if changed.
	DatastoreClusterId interface{}
	// Overrides any Storage DRS automation
	// levels for this virtual machine. Can be one of `automated` or `manual`. When
	// not specified, the datastore cluster's settings are used according to the
	// [specific SDRS subsystem][tf-vsphere-datastore-cluster-sdrs-levels].
	SdrsAutomationLevel interface{}
	// Overrides the default Storage DRS setting for
	// this virtual machine. When not specified, the datastore cluster setting is
	// used.
	SdrsEnabled interface{}
	// Overrides the intra-VM affinity setting
	// for this virtual machine. When `true`, all disks for this virtual machine
	// will be kept on the same datastore. When `false`, Storage DRS may locate
	// individual disks on different datastores if it helps satisfy cluster
	// requirements. When not specified, the datastore cluster's settings are used.
	SdrsIntraVmAffinity interface{}
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId interface{}
}

// The set of arguments for constructing a StorageDrsVmOverride resource.
type StorageDrsVmOverrideArgs struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the datastore cluster to put the override in.
	// Forces a new resource if changed.
	DatastoreClusterId interface{}
	// Overrides any Storage DRS automation
	// levels for this virtual machine. Can be one of `automated` or `manual`. When
	// not specified, the datastore cluster's settings are used according to the
	// [specific SDRS subsystem][tf-vsphere-datastore-cluster-sdrs-levels].
	SdrsAutomationLevel interface{}
	// Overrides the default Storage DRS setting for
	// this virtual machine. When not specified, the datastore cluster setting is
	// used.
	SdrsEnabled interface{}
	// Overrides the intra-VM affinity setting
	// for this virtual machine. When `true`, all disks for this virtual machine
	// will be kept on the same datastore. When `false`, Storage DRS may locate
	// individual disks on different datastores if it helps satisfy cluster
	// requirements. When not specified, the datastore cluster's settings are used.
	SdrsIntraVmAffinity interface{}
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId interface{}
}

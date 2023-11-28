// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "vsphere:index/computeCluster:ComputeCluster":
		r = &ComputeCluster{}
	case "vsphere:index/computeClusterHostGroup:ComputeClusterHostGroup":
		r = &ComputeClusterHostGroup{}
	case "vsphere:index/computeClusterVmAffinityRule:ComputeClusterVmAffinityRule":
		r = &ComputeClusterVmAffinityRule{}
	case "vsphere:index/computeClusterVmAntiAffinityRule:ComputeClusterVmAntiAffinityRule":
		r = &ComputeClusterVmAntiAffinityRule{}
	case "vsphere:index/computeClusterVmDependencyRule:ComputeClusterVmDependencyRule":
		r = &ComputeClusterVmDependencyRule{}
	case "vsphere:index/computeClusterVmGroup:ComputeClusterVmGroup":
		r = &ComputeClusterVmGroup{}
	case "vsphere:index/computeClusterVmHostRule:ComputeClusterVmHostRule":
		r = &ComputeClusterVmHostRule{}
	case "vsphere:index/contentLibrary:ContentLibrary":
		r = &ContentLibrary{}
	case "vsphere:index/contentLibraryItem:ContentLibraryItem":
		r = &ContentLibraryItem{}
	case "vsphere:index/customAttribute:CustomAttribute":
		r = &CustomAttribute{}
	case "vsphere:index/datacenter:Datacenter":
		r = &Datacenter{}
	case "vsphere:index/datastoreCluster:DatastoreCluster":
		r = &DatastoreCluster{}
	case "vsphere:index/datastoreClusterVmAntiAffinityRule:DatastoreClusterVmAntiAffinityRule":
		r = &DatastoreClusterVmAntiAffinityRule{}
	case "vsphere:index/distributedPortGroup:DistributedPortGroup":
		r = &DistributedPortGroup{}
	case "vsphere:index/distributedVirtualSwitch:DistributedVirtualSwitch":
		r = &DistributedVirtualSwitch{}
	case "vsphere:index/dpmHostOverride:DpmHostOverride":
		r = &DpmHostOverride{}
	case "vsphere:index/drsVmOverride:DrsVmOverride":
		r = &DrsVmOverride{}
	case "vsphere:index/entityPermissions:EntityPermissions":
		r = &EntityPermissions{}
	case "vsphere:index/file:File":
		r = &File{}
	case "vsphere:index/folder:Folder":
		r = &Folder{}
	case "vsphere:index/guestOsCustomization:GuestOsCustomization":
		r = &GuestOsCustomization{}
	case "vsphere:index/haVmOverride:HaVmOverride":
		r = &HaVmOverride{}
	case "vsphere:index/host:Host":
		r = &Host{}
	case "vsphere:index/hostPortGroup:HostPortGroup":
		r = &HostPortGroup{}
	case "vsphere:index/hostVirtualSwitch:HostVirtualSwitch":
		r = &HostVirtualSwitch{}
	case "vsphere:index/license:License":
		r = &License{}
	case "vsphere:index/nasDatastore:NasDatastore":
		r = &NasDatastore{}
	case "vsphere:index/resourcePool:ResourcePool":
		r = &ResourcePool{}
	case "vsphere:index/role:Role":
		r = &Role{}
	case "vsphere:index/storageDrsVmOverride:StorageDrsVmOverride":
		r = &StorageDrsVmOverride{}
	case "vsphere:index/tag:Tag":
		r = &Tag{}
	case "vsphere:index/tagCategory:TagCategory":
		r = &TagCategory{}
	case "vsphere:index/vappContainer:VappContainer":
		r = &VappContainer{}
	case "vsphere:index/vappEntity:VappEntity":
		r = &VappEntity{}
	case "vsphere:index/virtualDisk:VirtualDisk":
		r = &VirtualDisk{}
	case "vsphere:index/virtualMachine:VirtualMachine":
		r = &VirtualMachine{}
	case "vsphere:index/virtualMachineSnapshot:VirtualMachineSnapshot":
		r = &VirtualMachineSnapshot{}
	case "vsphere:index/vmStoragePolicy:VmStoragePolicy":
		r = &VmStoragePolicy{}
	case "vsphere:index/vmfsDatastore:VmfsDatastore":
		r = &VmfsDatastore{}
	case "vsphere:index/vnic:Vnic":
		r = &Vnic{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:vsphere" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/computeCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/computeClusterHostGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/computeClusterVmAffinityRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/computeClusterVmAntiAffinityRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/computeClusterVmDependencyRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/computeClusterVmGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/computeClusterVmHostRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/contentLibrary",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/contentLibraryItem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/customAttribute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/datacenter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/datastoreCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/datastoreClusterVmAntiAffinityRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/distributedPortGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/distributedVirtualSwitch",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/dpmHostOverride",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/drsVmOverride",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/entityPermissions",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/file",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/folder",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/guestOsCustomization",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/haVmOverride",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/host",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/hostPortGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/hostVirtualSwitch",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/license",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/nasDatastore",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/resourcePool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/role",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/storageDrsVmOverride",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/tag",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/tagCategory",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/vappContainer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/vappEntity",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/virtualDisk",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/virtualMachine",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/virtualMachineSnapshot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/vmStoragePolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/vmfsDatastore",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vsphere",
		"index/vnic",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"vsphere",
		&pkg{version},
	)
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./computeCluster";
export * from "./computeClusterHostGroup";
export * from "./computeClusterVmAffinityRule";
export * from "./computeClusterVmAntiAffinityRule";
export * from "./computeClusterVmDependencyRule";
export * from "./computeClusterVmGroup";
export * from "./computeClusterVmHostRule";
export * from "./contentLibrary";
export * from "./contentLibraryItem";
export * from "./customAttribute";
export * from "./datacenter";
export * from "./datastoreCluster";
export * from "./datastoreClusterVmAntiAffinityRule";
export * from "./distributedPortGroup";
export * from "./distributedVirtualSwitch";
export * from "./dpmHostOverride";
export * from "./drsVmOverride";
export * from "./entityPermissions";
export * from "./file";
export * from "./folder";
export * from "./getComputeCluster";
export * from "./getContentLibrary";
export * from "./getContentLibraryItem";
export * from "./getCustomAttribute";
export * from "./getDatacenter";
export * from "./getDatastore";
export * from "./getDatastoreCluster";
export * from "./getDistributedVirtualSwitch";
export * from "./getDynamic";
export * from "./getFolder";
export * from "./getHost";
export * from "./getHostPciDevice";
export * from "./getHostThumbprint";
export * from "./getNetwork";
export * from "./getOvfVmTemplate";
export * from "./getPolicy";
export * from "./getResourcePool";
export * from "./getRole";
export * from "./getTag";
export * from "./getTagCategory";
export * from "./getVappContainer";
export * from "./getVirtualMachine";
export * from "./getVmfsDisks";
export * from "./haVmOverride";
export * from "./host";
export * from "./hostPortGroup";
export * from "./hostVirtualSwitch";
export * from "./license";
export * from "./nasDatastore";
export * from "./provider";
export * from "./resourcePool";
export * from "./role";
export * from "./storageDrsVmOverride";
export * from "./tag";
export * from "./tagCategory";
export * from "./vappContainer";
export * from "./vappEntity";
export * from "./virtualDisk";
export * from "./virtualMachine";
export * from "./virtualMachineSnapshot";
export * from "./vmStoragePolicy";
export * from "./vmfsDatastore";
export * from "./vnic";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { ComputeCluster } from "./computeCluster";
import { ComputeClusterHostGroup } from "./computeClusterHostGroup";
import { ComputeClusterVmAffinityRule } from "./computeClusterVmAffinityRule";
import { ComputeClusterVmAntiAffinityRule } from "./computeClusterVmAntiAffinityRule";
import { ComputeClusterVmDependencyRule } from "./computeClusterVmDependencyRule";
import { ComputeClusterVmGroup } from "./computeClusterVmGroup";
import { ComputeClusterVmHostRule } from "./computeClusterVmHostRule";
import { ContentLibrary } from "./contentLibrary";
import { ContentLibraryItem } from "./contentLibraryItem";
import { CustomAttribute } from "./customAttribute";
import { Datacenter } from "./datacenter";
import { DatastoreCluster } from "./datastoreCluster";
import { DatastoreClusterVmAntiAffinityRule } from "./datastoreClusterVmAntiAffinityRule";
import { DistributedPortGroup } from "./distributedPortGroup";
import { DistributedVirtualSwitch } from "./distributedVirtualSwitch";
import { DpmHostOverride } from "./dpmHostOverride";
import { DrsVmOverride } from "./drsVmOverride";
import { EntityPermissions } from "./entityPermissions";
import { File } from "./file";
import { Folder } from "./folder";
import { HaVmOverride } from "./haVmOverride";
import { Host } from "./host";
import { HostPortGroup } from "./hostPortGroup";
import { HostVirtualSwitch } from "./hostVirtualSwitch";
import { License } from "./license";
import { NasDatastore } from "./nasDatastore";
import { ResourcePool } from "./resourcePool";
import { Role } from "./role";
import { StorageDrsVmOverride } from "./storageDrsVmOverride";
import { Tag } from "./tag";
import { TagCategory } from "./tagCategory";
import { VappContainer } from "./vappContainer";
import { VappEntity } from "./vappEntity";
import { VirtualDisk } from "./virtualDisk";
import { VirtualMachine } from "./virtualMachine";
import { VirtualMachineSnapshot } from "./virtualMachineSnapshot";
import { VmStoragePolicy } from "./vmStoragePolicy";
import { VmfsDatastore } from "./vmfsDatastore";
import { Vnic } from "./vnic";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vsphere:index/computeCluster:ComputeCluster":
                return new ComputeCluster(name, <any>undefined, { urn })
            case "vsphere:index/computeClusterHostGroup:ComputeClusterHostGroup":
                return new ComputeClusterHostGroup(name, <any>undefined, { urn })
            case "vsphere:index/computeClusterVmAffinityRule:ComputeClusterVmAffinityRule":
                return new ComputeClusterVmAffinityRule(name, <any>undefined, { urn })
            case "vsphere:index/computeClusterVmAntiAffinityRule:ComputeClusterVmAntiAffinityRule":
                return new ComputeClusterVmAntiAffinityRule(name, <any>undefined, { urn })
            case "vsphere:index/computeClusterVmDependencyRule:ComputeClusterVmDependencyRule":
                return new ComputeClusterVmDependencyRule(name, <any>undefined, { urn })
            case "vsphere:index/computeClusterVmGroup:ComputeClusterVmGroup":
                return new ComputeClusterVmGroup(name, <any>undefined, { urn })
            case "vsphere:index/computeClusterVmHostRule:ComputeClusterVmHostRule":
                return new ComputeClusterVmHostRule(name, <any>undefined, { urn })
            case "vsphere:index/contentLibrary:ContentLibrary":
                return new ContentLibrary(name, <any>undefined, { urn })
            case "vsphere:index/contentLibraryItem:ContentLibraryItem":
                return new ContentLibraryItem(name, <any>undefined, { urn })
            case "vsphere:index/customAttribute:CustomAttribute":
                return new CustomAttribute(name, <any>undefined, { urn })
            case "vsphere:index/datacenter:Datacenter":
                return new Datacenter(name, <any>undefined, { urn })
            case "vsphere:index/datastoreCluster:DatastoreCluster":
                return new DatastoreCluster(name, <any>undefined, { urn })
            case "vsphere:index/datastoreClusterVmAntiAffinityRule:DatastoreClusterVmAntiAffinityRule":
                return new DatastoreClusterVmAntiAffinityRule(name, <any>undefined, { urn })
            case "vsphere:index/distributedPortGroup:DistributedPortGroup":
                return new DistributedPortGroup(name, <any>undefined, { urn })
            case "vsphere:index/distributedVirtualSwitch:DistributedVirtualSwitch":
                return new DistributedVirtualSwitch(name, <any>undefined, { urn })
            case "vsphere:index/dpmHostOverride:DpmHostOverride":
                return new DpmHostOverride(name, <any>undefined, { urn })
            case "vsphere:index/drsVmOverride:DrsVmOverride":
                return new DrsVmOverride(name, <any>undefined, { urn })
            case "vsphere:index/entityPermissions:EntityPermissions":
                return new EntityPermissions(name, <any>undefined, { urn })
            case "vsphere:index/file:File":
                return new File(name, <any>undefined, { urn })
            case "vsphere:index/folder:Folder":
                return new Folder(name, <any>undefined, { urn })
            case "vsphere:index/haVmOverride:HaVmOverride":
                return new HaVmOverride(name, <any>undefined, { urn })
            case "vsphere:index/host:Host":
                return new Host(name, <any>undefined, { urn })
            case "vsphere:index/hostPortGroup:HostPortGroup":
                return new HostPortGroup(name, <any>undefined, { urn })
            case "vsphere:index/hostVirtualSwitch:HostVirtualSwitch":
                return new HostVirtualSwitch(name, <any>undefined, { urn })
            case "vsphere:index/license:License":
                return new License(name, <any>undefined, { urn })
            case "vsphere:index/nasDatastore:NasDatastore":
                return new NasDatastore(name, <any>undefined, { urn })
            case "vsphere:index/resourcePool:ResourcePool":
                return new ResourcePool(name, <any>undefined, { urn })
            case "vsphere:index/role:Role":
                return new Role(name, <any>undefined, { urn })
            case "vsphere:index/storageDrsVmOverride:StorageDrsVmOverride":
                return new StorageDrsVmOverride(name, <any>undefined, { urn })
            case "vsphere:index/tag:Tag":
                return new Tag(name, <any>undefined, { urn })
            case "vsphere:index/tagCategory:TagCategory":
                return new TagCategory(name, <any>undefined, { urn })
            case "vsphere:index/vappContainer:VappContainer":
                return new VappContainer(name, <any>undefined, { urn })
            case "vsphere:index/vappEntity:VappEntity":
                return new VappEntity(name, <any>undefined, { urn })
            case "vsphere:index/virtualDisk:VirtualDisk":
                return new VirtualDisk(name, <any>undefined, { urn })
            case "vsphere:index/virtualMachine:VirtualMachine":
                return new VirtualMachine(name, <any>undefined, { urn })
            case "vsphere:index/virtualMachineSnapshot:VirtualMachineSnapshot":
                return new VirtualMachineSnapshot(name, <any>undefined, { urn })
            case "vsphere:index/vmStoragePolicy:VmStoragePolicy":
                return new VmStoragePolicy(name, <any>undefined, { urn })
            case "vsphere:index/vmfsDatastore:VmfsDatastore":
                return new VmfsDatastore(name, <any>undefined, { urn })
            case "vsphere:index/vnic:Vnic":
                return new Vnic(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vsphere", "index/computeCluster", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/computeClusterHostGroup", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/computeClusterVmAffinityRule", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/computeClusterVmAntiAffinityRule", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/computeClusterVmDependencyRule", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/computeClusterVmGroup", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/computeClusterVmHostRule", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/contentLibrary", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/contentLibraryItem", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/customAttribute", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/datacenter", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/datastoreCluster", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/datastoreClusterVmAntiAffinityRule", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/distributedPortGroup", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/distributedVirtualSwitch", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/dpmHostOverride", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/drsVmOverride", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/entityPermissions", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/file", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/folder", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/haVmOverride", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/host", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/hostPortGroup", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/hostVirtualSwitch", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/license", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/nasDatastore", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/resourcePool", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/role", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/storageDrsVmOverride", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/tag", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/tagCategory", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/vappContainer", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/vappEntity", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/virtualDisk", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/virtualMachine", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/virtualMachineSnapshot", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/vmStoragePolicy", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/vmfsDatastore", _module)
pulumi.runtime.registerResourceModule("vsphere", "index/vnic", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("vsphere", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:vsphere") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});

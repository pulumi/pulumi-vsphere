// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

import * as vsphere from "@pulumi/vsphere";

// The configuration below is designed to match what is needed to provision resources in a VMWare Cloud on AWS
// installation of vSphere.  Similar (and possibly simpler) configuration will work with other environments.
const dc = vsphere.getDatacenter({ name: "SDDC-Datacenter" });
const datastoreId = dc.then(dc => vsphere.getDatastore({ name: "WorkloadDatastore", datacenterId: dc.id })).then(d => d.id);
const poolId = dc.then(dc => vsphere.getResourcePool({ name: "Compute-ResourcePool", datacenterId: dc.id })).then(p => p.id);
const networkId = dc.then(dc => vsphere.getNetwork({ name: "sddc-cgw-network-1", datacenterId: dc.id })).then(n => n.id);
const template = dc.then(dc => vsphere.getVirtualMachine({ name: "Workloads/OtherVM", datacenterId: dc.id }));

const vm = new vsphere.VirtualMachine("vm", {
    resourcePoolId: poolId,
    datastoreId: datastoreId,
    folder: "Workloads",
    numCpus: 2,
    memory: 2048,
    guestId: template.then(t => t.guestId),
    networkInterfaces: [{
        networkId: networkId,
        adapterType: template.then(t => t.networkInterfaceTypes[0]),
    }],
    disks: [{
        label: "disk0",
        size: template.then(t => t.disks[0].size),
        eagerlyScrub: template.then(t => t.disks[0].eagerlyScrub),
        thinProvisioned: template.then(t => t.disks[0].thinProvisioned),
    }],
    clone: {
        templateUuid: template.then(t => t.id), 
    },
});

export let defaultIp = vm.defaultIpAddress;

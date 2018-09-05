// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

import * as vsphere from "@pulumi/vsphere";

const dc = vsphere.getDatacenter({ name: "dc1" });
const datastoreId = dc.then(dc => vsphere.getDatastore({ name: "datastore1", datacenterId: dc.id })).then(d => d.id);
const poolId = dc.then(dc => vsphere.getResourcePool({ name: "cluster1/Resources", datacenterId: dc.id })).then(p => p.id);
const networkId = dc.then(dc => vsphere.getNetwork({ name: "public", datacenterId: dc.id })).then(n => n.id);

const vm = new vsphere.VirtualMachine("vm", {
    resourcePoolId: poolId,
    datastoreId: datastoreId,
    numCpus: 2,
    memory: 1024,
    guestId: "other3xLinux64Guest",
    networkInterfaces: [{
        networkId: networkId,
    }],
    disks: [{
        label: "disk0",
        size: 20,
    }],
});

export let defaultIp = vm.defaultIpAddress;

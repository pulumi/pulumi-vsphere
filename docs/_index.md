---
title: Vsphere Provider
meta_desc: Provides an overview on how to configure the Pulumi Vsphere provider.
layout: package
---
## Installation

The vsphere provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/vsphere`](https://www.npmjs.com/package/@pulumi/vsphere)
* Python: [`pulumi-vsphere`](https://pypi.org/project/pulumi-vsphere/)
* Go: [`github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere`](https://github.com/pulumi/pulumi-vsphere)
* .NET: [`Pulumi.Vsphere`](https://www.nuget.org/packages/Pulumi.Vsphere)
* Java: [`com.pulumi/vsphere`](https://central.sonatype.com/artifact/com.pulumi/vsphere)
## Overview

This provider gives Pulumi the ability to work with VMware vSphere,
notably [vCenter Server](https://www.vmware.com/products/vcenter.html) and [ESXi](https://www.vmware.com/content/vmware/vmware-published-sites/us/products/esxi-and-esx.html).
This provider can be used to manage many aspects of a vSphere environment,
including virtual machines, standard and distributed switches, datastores,
content libraries, and more.

Use the navigation to read about the resources and functions supported by
this provider.

> **NOTE:** This provider requires API write access and hence is not supported
on a free ESXi license.
## Example Usage

The following abridged example demonstrates basic usage of the provider to
provision a virtual machine using the
[`vsphere.VirtualMachine`](https://www.terraform.io/docs/providers/vsphere/r/virtual_machine.html) resource.
The datacenter, datastore, resource pool, and network are discovered using the
[`vsphere.Datacenter`](https://www.terraform.io/docs/providers/vsphere/d/datacenter.html),
[`vsphere.getDatastore`](https://www.terraform.io/docs/providers/vsphere/d/datastore.html),
[`vsphere.ResourcePool`](https://www.terraform.io/docs/providers/vsphere/d/resource_pool.html), and
[`vsphere.getNetwork`](https://www.terraform.io/docs/providers/vsphere/d/network.html) functions respectively.

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    vsphere:allowUnverifiedSsl:
        value: true
    vsphere:apiTimeout:
        value: 10
    vsphere:password:
        value: 'TODO: var.vsphere_password'
    vsphere:user:
        value: 'TODO: var.vsphere_user'
    vsphere:vsphereServer:
        value: 'TODO: var.vsphere_server'

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as vsphere from "@pulumi/vsphere";

const datacenter = vsphere.getDatacenter({
    name: "dc-01",
});
const datastore = datacenter.then(datacenter => vsphere.getDatastore({
    name: "datastore-01",
    datacenterId: datacenter.id,
}));
const cluster = datacenter.then(datacenter => vsphere.getComputeCluster({
    name: "cluster-01",
    datacenterId: datacenter.id,
}));
const network = datacenter.then(datacenter => vsphere.getNetwork({
    name: "VM Network",
    datacenterId: datacenter.id,
}));
const vm = new vsphere.VirtualMachine("vm", {
    name: "foo",
    resourcePoolId: cluster.then(cluster => cluster.resourcePoolId),
    datastoreId: datastore.then(datastore => datastore.id),
    numCpus: 1,
    memory: 1024,
    guestId: "otherLinux64Guest",
    networkInterfaces: [{
        networkId: network.then(network => network.id),
    }],
    disks: [{
        label: "disk0",
        size: 20,
    }],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    vsphere:allowUnverifiedSsl:
        value: true
    vsphere:apiTimeout:
        value: 10
    vsphere:password:
        value: 'TODO: var.vsphere_password'
    vsphere:user:
        value: 'TODO: var.vsphere_user'
    vsphere:vsphereServer:
        value: 'TODO: var.vsphere_server'

```
```python
import pulumi
import pulumi_vsphere as vsphere

datacenter = vsphere.get_datacenter(name="dc-01")
datastore = vsphere.get_datastore(name="datastore-01",
    datacenter_id=datacenter.id)
cluster = vsphere.get_compute_cluster(name="cluster-01",
    datacenter_id=datacenter.id)
network = vsphere.get_network(name="VM Network",
    datacenter_id=datacenter.id)
vm = vsphere.VirtualMachine("vm",
    name="foo",
    resource_pool_id=cluster.resource_pool_id,
    datastore_id=datastore.id,
    num_cpus=1,
    memory=1024,
    guest_id="otherLinux64Guest",
    network_interfaces=[{
        "network_id": network.id,
    }],
    disks=[{
        "label": "disk0",
        "size": 20,
    }])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    vsphere:allowUnverifiedSsl:
        value: true
    vsphere:apiTimeout:
        value: 10
    vsphere:password:
        value: 'TODO: var.vsphere_password'
    vsphere:user:
        value: 'TODO: var.vsphere_user'
    vsphere:vsphereServer:
        value: 'TODO: var.vsphere_server'

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using VSphere = Pulumi.VSphere;

return await Deployment.RunAsync(() =>
{
    var datacenter = VSphere.GetDatacenter.Invoke(new()
    {
        Name = "dc-01",
    });

    var datastore = VSphere.GetDatastore.Invoke(new()
    {
        Name = "datastore-01",
        DatacenterId = datacenter.Apply(getDatacenterResult => getDatacenterResult.Id),
    });

    var cluster = VSphere.GetComputeCluster.Invoke(new()
    {
        Name = "cluster-01",
        DatacenterId = datacenter.Apply(getDatacenterResult => getDatacenterResult.Id),
    });

    var network = VSphere.GetNetwork.Invoke(new()
    {
        Name = "VM Network",
        DatacenterId = datacenter.Apply(getDatacenterResult => getDatacenterResult.Id),
    });

    var vm = new VSphere.VirtualMachine("vm", new()
    {
        Name = "foo",
        ResourcePoolId = cluster.Apply(getComputeClusterResult => getComputeClusterResult.ResourcePoolId),
        DatastoreId = datastore.Apply(getDatastoreResult => getDatastoreResult.Id),
        NumCpus = 1,
        Memory = 1024,
        GuestId = "otherLinux64Guest",
        NetworkInterfaces = new[]
        {
            new VSphere.Inputs.VirtualMachineNetworkInterfaceArgs
            {
                NetworkId = network.Apply(getNetworkResult => getNetworkResult.Id),
            },
        },
        Disks = new[]
        {
            new VSphere.Inputs.VirtualMachineDiskArgs
            {
                Label = "disk0",
                Size = 20,
            },
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    vsphere:allowUnverifiedSsl:
        value: true
    vsphere:apiTimeout:
        value: 10
    vsphere:password:
        value: 'TODO: var.vsphere_password'
    vsphere:user:
        value: 'TODO: var.vsphere_user'
    vsphere:vsphereServer:
        value: 'TODO: var.vsphere_server'

```
```go
package main

import (
	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		datacenter, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
			Name: pulumi.StringRef("dc-01"),
		}, nil)
		if err != nil {
			return err
		}
		datastore, err := vsphere.GetDatastore(ctx, &vsphere.GetDatastoreArgs{
			Name:         "datastore-01",
			DatacenterId: pulumi.StringRef(datacenter.Id),
		}, nil)
		if err != nil {
			return err
		}
		cluster, err := vsphere.LookupComputeCluster(ctx, &vsphere.LookupComputeClusterArgs{
			Name:         "cluster-01",
			DatacenterId: pulumi.StringRef(datacenter.Id),
		}, nil)
		if err != nil {
			return err
		}
		network, err := vsphere.GetNetwork(ctx, &vsphere.GetNetworkArgs{
			Name:         "VM Network",
			DatacenterId: pulumi.StringRef(datacenter.Id),
		}, nil)
		if err != nil {
			return err
		}
		_, err = vsphere.NewVirtualMachine(ctx, "vm", &vsphere.VirtualMachineArgs{
			Name:           pulumi.String("foo"),
			ResourcePoolId: pulumi.String(cluster.ResourcePoolId),
			DatastoreId:    pulumi.String(datastore.Id),
			NumCpus:        pulumi.Int(1),
			Memory:         pulumi.Int(1024),
			GuestId:        pulumi.String("otherLinux64Guest"),
			NetworkInterfaces: vsphere.VirtualMachineNetworkInterfaceArray{
				&vsphere.VirtualMachineNetworkInterfaceArgs{
					NetworkId: pulumi.String(network.Id),
				},
			},
			Disks: vsphere.VirtualMachineDiskArray{
				&vsphere.VirtualMachineDiskArgs{
					Label: pulumi.String("disk0"),
					Size:  pulumi.Int(20),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    vsphere:allowUnverifiedSsl:
        value: true
    vsphere:apiTimeout:
        value: 10
    vsphere:password:
        value: 'TODO: var.vsphere_password'
    vsphere:user:
        value: 'TODO: var.vsphere_user'
    vsphere:vsphereServer:
        value: 'TODO: var.vsphere_server'

```
```yaml
resources:
  vm:
    type: vsphere:VirtualMachine
    properties:
      name: foo
      resourcePoolId: ${cluster.resourcePoolId}
      datastoreId: ${datastore.id}
      numCpus: 1
      memory: 1024
      guestId: otherLinux64Guest
      networkInterfaces:
        - networkId: ${network.id}
      disks:
        - label: disk0
          size: 20
variables:
  datacenter:
    fn::invoke:
      Function: vsphere:getDatacenter
      Arguments:
        name: dc-01
  datastore:
    fn::invoke:
      Function: vsphere:getDatastore
      Arguments:
        name: datastore-01
        datacenterId: ${datacenter.id}
  cluster:
    fn::invoke:
      Function: vsphere:getComputeCluster
      Arguments:
        name: cluster-01
        datacenterId: ${datacenter.id}
  network:
    fn::invoke:
      Function: vsphere:getNetwork
      Arguments:
        name: VM Network
        datacenterId: ${datacenter.id}
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    vsphere:allowUnverifiedSsl:
        value: true
    vsphere:apiTimeout:
        value: 10
    vsphere:password:
        value: 'TODO: var.vsphere_password'
    vsphere:user:
        value: 'TODO: var.vsphere_user'
    vsphere:vsphereServer:
        value: 'TODO: var.vsphere_server'

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.vsphere.VsphereFunctions;
import com.pulumi.vsphere.inputs.GetDatacenterArgs;
import com.pulumi.vsphere.inputs.GetDatastoreArgs;
import com.pulumi.vsphere.inputs.GetComputeClusterArgs;
import com.pulumi.vsphere.inputs.GetNetworkArgs;
import com.pulumi.vsphere.VirtualMachine;
import com.pulumi.vsphere.VirtualMachineArgs;
import com.pulumi.vsphere.inputs.VirtualMachineNetworkInterfaceArgs;
import com.pulumi.vsphere.inputs.VirtualMachineDiskArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        final var datacenter = VsphereFunctions.getDatacenter(GetDatacenterArgs.builder()
            .name("dc-01")
            .build());

        final var datastore = VsphereFunctions.getDatastore(GetDatastoreArgs.builder()
            .name("datastore-01")
            .datacenterId(datacenter.applyValue(getDatacenterResult -> getDatacenterResult.id()))
            .build());

        final var cluster = VsphereFunctions.getComputeCluster(GetComputeClusterArgs.builder()
            .name("cluster-01")
            .datacenterId(datacenter.applyValue(getDatacenterResult -> getDatacenterResult.id()))
            .build());

        final var network = VsphereFunctions.getNetwork(GetNetworkArgs.builder()
            .name("VM Network")
            .datacenterId(datacenter.applyValue(getDatacenterResult -> getDatacenterResult.id()))
            .build());

        var vm = new VirtualMachine("vm", VirtualMachineArgs.builder()
            .name("foo")
            .resourcePoolId(cluster.applyValue(getComputeClusterResult -> getComputeClusterResult.resourcePoolId()))
            .datastoreId(datastore.applyValue(getDatastoreResult -> getDatastoreResult.id()))
            .numCpus(1)
            .memory(1024)
            .guestId("otherLinux64Guest")
            .networkInterfaces(VirtualMachineNetworkInterfaceArgs.builder()
                .networkId(network.applyValue(getNetworkResult -> getNetworkResult.id()))
                .build())
            .disks(VirtualMachineDiskArgs.builder()
                .label("disk0")
                .size(20)
                .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

Refer to the provider documentation for information on all of the resources
and functions supported by this provider. Each includes a detailed
description of the purpose and how to use it.
## Configuration Reference

The following arguments are used to configure the provider:

* `user` - (Required) This is the username for vSphere API operations. Can also
  be specified with the `VSPHERE_USER` environment variable.
* `password` - (Required) This is the password for vSphere API operations. Can
  also be specified with the `VSPHERE_PASSWORD` environment variable.
* `vsphereServer` - (Required) This is the vCenter Server FQDN or IP Address
  for vSphere API operations. Can also be specified with the `VSPHERE_SERVER`
  environment variable.
* `allowUnverifiedSsl` - (Optional) Boolean that can be set to true to
  disable SSL certificate verification. This should be used with care as it
  could allow an attacker to intercept your authentication token. If omitted,
  default value is `false`. Can also be specified with the
  `VSPHERE_ALLOW_UNVERIFIED_SSL` environment variable.
* `vimKeepAlive` - (Optional) Keep alive interval in minutes for the VIM
  session. Standard session timeout in vSphere is 30 minutes. This defaults to
  10 minutes to ensure that operations that take a longer than 30 minutes
  without API interaction do not result in a session timeout. Can also be
  specified with the `VSPHERE_VIM_KEEP_ALIVE` environment variable.
* `apiTimeout` - (Optional) Sets the number of minutes to wait for operations
  to complete. The default timeout is 5 minutes. Can also be
  specified with the `VSPHERE_API_TIMEOUT` environment variable.

> **NOTE:** Use of the `apiTimeout` option to extend the timeout from the
default is recommended when creating virtual machines with large disks.
### Session Persistence Options

The provider also provides session persistence options that can be configured
below. These can help when using Pulumi in a way where session limits could
be normally reached by creating a new session for every run, such as a large
amount of concurrent or consecutive Pulumi runs in a short period of time.

> **NOTE:** Session keys are as good as user credentials for as long as the
session is valid for - handle them with care and delete them when you know you
will no longer need them.

* `persistSession` - (Optional) Persist the SOAP and REST client sessions to
  disk. Default: `false`. Can also be specified by the
  `VSPHERE_PERSIST_SESSION` environment variable.
* `vimSessionPath` - (Optional) The directory to save the VIM SOAP API
  session to. Default: `${HOME}/.govmomi/sessions`. Can also be specified by
  the `VSPHERE_VIM_SESSION_PATH` environment variable.
* `restSessionPath` - The directory to save the REST API session to.
  Default: `${HOME}/.govmomi/rest_sessions`. Can also be specified by the
  `VSPHERE_REST_SESSION_PATH` environment variable.
#### Session Interoperability for vmware/govc and the Provider

The session format used to save VIM SOAP sessions is the same used
with [vmware/govc](https://github.com/vmware/govmomi/tree/main/govc). If you use govc as part of your provisioning
process, Pulumi will use the saved session if present and if
`persistSession` is enabled.
## Notes on Required Privileges

When using a non-administrator account to perform provider operations, consider
that most Pulumi resources perform operations in a CRUD-like fashion and
require both read and write privileges to the resources they are managing. Make
sure that the user has appropriate read-write access to the resources you need
to work with. Read-only access should be sufficient when only using data
sources on some features. You can read more about vSphere permissions and user
management [here](https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-security/GUID-5372F580-5C23-4E9C-8A4E-EF1B4DD9033E.html).

There are a some notable exceptions to keep in mind when setting up a restricted
provisioning user:
### Tags

The provider will always attempt to read [tags](https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vcenter-esxi-management/GUID-E8E854DD-AA97-4E0C-8419-CE84F93C4058.html) from a
resource, even if you do not have any tags defined. Ensure that your user has
access to read tags.
### Events

The provider will attempt to read event data from vSphere to check for certain
events, such as, virtual machine customization or power events. Ensure that the
user account used for Pulumi has the privilege to be able to read event data.
### Storage

The provider implementation requires the ability to read storage profiles
from vSphere for some resource and function operations. Ensure that the
user account used for Pulumi is provided the Profile-driven Storage > View
(`StorageProfile.View`) privilege to be able to read the available storage
policies.
### Virtual Machine

The provider implementation requires the ability to set a default swap
placement policy on a virtual machine resource. Ensure that the user account
used for Pulumi is provided the Virtual Machine > Change Configuration >
Change Swapfile Placement (`VirtualMachine.Config.SwapPlacement`) privilege.
## Use of Managed Object References by the Provider

Unlike the vSphere client, many resources managed by the provider
use managed object IDs (or UUIDs when provided and practical) when referring
to placement parameters and upstream resources. This provides a stable
interface for providing necessary data to downstream resources, in addition to
minimizing the issues that can arise from the flexibility in how an individual
object's name or inventory path can be supplied.

There are several functions (such as
[`vsphere.Datacenter`](https://www.terraform.io/docs/providers/vsphere/d/datacenter.html),
[`vsphere.Host`](https://www.terraform.io/docs/providers/vsphere/d/host.html),
[`vsphere.ResourcePool`](https://www.terraform.io/docs/providers/vsphere/d/resource_pool.html),
[`vsphere.getDatastore`](https://www.terraform.io/docs/providers/vsphere/d/datastore.html), and
[`vsphere.getNetwork`](https://www.terraform.io/docs/providers/vsphere/d/network.html)) that assist with searching for a
specific resource. For usage details on a specific function, look for a
link in the provider documentation. In addition, most vSphere
resources in Pulumi supply the managed object ID (or UUID, when it makes
more sense) as the `id` attribute, which can be supplied to downstream
resources that should depend on the parent.
### Locating Managed Object IDs

There are certain points in time that you may need to locate the managed object
ID of a specific vSphere resource yourself. A couple of methods are documented
below.
#### Using `govc`

[`govc`](https://github.com/vmware/govmomi/tree/main/govc) is an vSphere CLI built on [govmomi](https://github.com/vmware/govmomi), the
vSphere Go SDK. It has a robust inventory browser command that can also be used
to list managed object IDs.

To get all the necessary data in a single output, use `govc ls -l -i PATH`.

Sample output is below:

```shell
$ govc ls -l -i /dc-01/vm
VirtualMachine:vm-123 /dc-01/vm/foobar
Folder:group-v234 /dc-01/vm/subfolder
```

To do a reverse search, supply the `-L` switch:

```shell
$ govc ls -i -l -L VirtualMachine:vm-123
VirtualMachine:vm-123 /dc-01/vm/foo
```

For details on setting up `govc`, see the [GitHub project](https://github.com/vmware/govmomi/tree/main/govc).
#### Using the vSphere Managed Object Browser (MOB)

The Managed Object Browser (MOB) allows one to browse the entire vSphere
inventory as it's presented to the API. It's normally accessed using
`https://<vcenter_fqdn>/mob`.

> **NOTE:** The MOB also offers API method invocation capabilities, and for
security reasons should be used sparingly. Modern vSphere installations may
have the MOB disabled by default, at the very least on ESXi systems. For more
information on current security best practices related to the MOB on ESXi,
click [here](https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-security/GUID-0EF83EA7-277C-400B-B697-04BDC9173EA3.html).
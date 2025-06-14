# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['DrsVmOverrideArgs', 'DrsVmOverride']

@pulumi.input_type
class DrsVmOverrideArgs:
    def __init__(__self__, *,
                 compute_cluster_id: pulumi.Input[builtins.str],
                 virtual_machine_id: pulumi.Input[builtins.str],
                 drs_automation_level: Optional[pulumi.Input[builtins.str]] = None,
                 drs_enabled: Optional[pulumi.Input[builtins.bool]] = None):
        """
        The set of arguments for constructing a DrsVmOverride resource.
        :param pulumi.Input[builtins.str] compute_cluster_id: The managed object reference
               ID of the cluster to put the override in.  Forces a new
               resource if changed.
        :param pulumi.Input[builtins.str] virtual_machine_id: The UUID of the virtual machine to create
               the override for.  Forces a new resource if changed.
        :param pulumi.Input[builtins.str] drs_automation_level: Overrides the automation level for this virtual
               machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
               `fullyAutomated`. Default: `manual`.
               
               > **NOTE:** Using this resource _always_ implies an override, even if one of
               `drs_enabled` or `drs_automation_level` is omitted. Take note of the defaults
               for both options.
        :param pulumi.Input[builtins.bool] drs_enabled: Overrides the default DRS setting for this virtual
               machine. Can be either `true` or `false`. Default: `false`.
        """
        pulumi.set(__self__, "compute_cluster_id", compute_cluster_id)
        pulumi.set(__self__, "virtual_machine_id", virtual_machine_id)
        if drs_automation_level is not None:
            pulumi.set(__self__, "drs_automation_level", drs_automation_level)
        if drs_enabled is not None:
            pulumi.set(__self__, "drs_enabled", drs_enabled)

    @property
    @pulumi.getter(name="computeClusterId")
    def compute_cluster_id(self) -> pulumi.Input[builtins.str]:
        """
        The managed object reference
        ID of the cluster to put the override in.  Forces a new
        resource if changed.
        """
        return pulumi.get(self, "compute_cluster_id")

    @compute_cluster_id.setter
    def compute_cluster_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "compute_cluster_id", value)

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> pulumi.Input[builtins.str]:
        """
        The UUID of the virtual machine to create
        the override for.  Forces a new resource if changed.
        """
        return pulumi.get(self, "virtual_machine_id")

    @virtual_machine_id.setter
    def virtual_machine_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "virtual_machine_id", value)

    @property
    @pulumi.getter(name="drsAutomationLevel")
    def drs_automation_level(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Overrides the automation level for this virtual
        machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
        `fullyAutomated`. Default: `manual`.

        > **NOTE:** Using this resource _always_ implies an override, even if one of
        `drs_enabled` or `drs_automation_level` is omitted. Take note of the defaults
        for both options.
        """
        return pulumi.get(self, "drs_automation_level")

    @drs_automation_level.setter
    def drs_automation_level(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "drs_automation_level", value)

    @property
    @pulumi.getter(name="drsEnabled")
    def drs_enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Overrides the default DRS setting for this virtual
        machine. Can be either `true` or `false`. Default: `false`.
        """
        return pulumi.get(self, "drs_enabled")

    @drs_enabled.setter
    def drs_enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "drs_enabled", value)


@pulumi.input_type
class _DrsVmOverrideState:
    def __init__(__self__, *,
                 compute_cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 drs_automation_level: Optional[pulumi.Input[builtins.str]] = None,
                 drs_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 virtual_machine_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering DrsVmOverride resources.
        :param pulumi.Input[builtins.str] compute_cluster_id: The managed object reference
               ID of the cluster to put the override in.  Forces a new
               resource if changed.
        :param pulumi.Input[builtins.str] drs_automation_level: Overrides the automation level for this virtual
               machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
               `fullyAutomated`. Default: `manual`.
               
               > **NOTE:** Using this resource _always_ implies an override, even if one of
               `drs_enabled` or `drs_automation_level` is omitted. Take note of the defaults
               for both options.
        :param pulumi.Input[builtins.bool] drs_enabled: Overrides the default DRS setting for this virtual
               machine. Can be either `true` or `false`. Default: `false`.
        :param pulumi.Input[builtins.str] virtual_machine_id: The UUID of the virtual machine to create
               the override for.  Forces a new resource if changed.
        """
        if compute_cluster_id is not None:
            pulumi.set(__self__, "compute_cluster_id", compute_cluster_id)
        if drs_automation_level is not None:
            pulumi.set(__self__, "drs_automation_level", drs_automation_level)
        if drs_enabled is not None:
            pulumi.set(__self__, "drs_enabled", drs_enabled)
        if virtual_machine_id is not None:
            pulumi.set(__self__, "virtual_machine_id", virtual_machine_id)

    @property
    @pulumi.getter(name="computeClusterId")
    def compute_cluster_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The managed object reference
        ID of the cluster to put the override in.  Forces a new
        resource if changed.
        """
        return pulumi.get(self, "compute_cluster_id")

    @compute_cluster_id.setter
    def compute_cluster_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "compute_cluster_id", value)

    @property
    @pulumi.getter(name="drsAutomationLevel")
    def drs_automation_level(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Overrides the automation level for this virtual
        machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
        `fullyAutomated`. Default: `manual`.

        > **NOTE:** Using this resource _always_ implies an override, even if one of
        `drs_enabled` or `drs_automation_level` is omitted. Take note of the defaults
        for both options.
        """
        return pulumi.get(self, "drs_automation_level")

    @drs_automation_level.setter
    def drs_automation_level(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "drs_automation_level", value)

    @property
    @pulumi.getter(name="drsEnabled")
    def drs_enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Overrides the default DRS setting for this virtual
        machine. Can be either `true` or `false`. Default: `false`.
        """
        return pulumi.get(self, "drs_enabled")

    @drs_enabled.setter
    def drs_enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "drs_enabled", value)

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The UUID of the virtual machine to create
        the override for.  Forces a new resource if changed.
        """
        return pulumi.get(self, "virtual_machine_id")

    @virtual_machine_id.setter
    def virtual_machine_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "virtual_machine_id", value)


@pulumi.type_token("vsphere:index/drsVmOverride:DrsVmOverride")
class DrsVmOverride(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 drs_automation_level: Optional[pulumi.Input[builtins.str]] = None,
                 drs_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 virtual_machine_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The `DrsVmOverride` resource can be used to add a DRS override to a
        cluster for a specific virtual machine. With this resource, one can enable or
        disable DRS and control the automation level for a single virtual machine
        without affecting the rest of the cluster.

        For more information on vSphere clusters and DRS, see [this
        page][ref-vsphere-drs-clusters].

        [ref-vsphere-drs-clusters]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-resource-management-8-0/creating-a-drs-cluster.html

        > **NOTE:** This resource requires vCenter and is not available on direct ESXi
        connections.

        ## Example Usage

        The example below creates a virtual machine in a cluster using the
        `VirtualMachine` resource, creating the
        virtual machine in the cluster looked up by the
        `ComputeCluster` data source, but also
        pinning the VM to a host defined by the
        `Host` data source, which is assumed to
        be a host within the cluster. To ensure that the VM stays on this host and does
        not need to be migrated back at any point in time, an override is entered using
        the `DrsVmOverride` resource that disables DRS for this virtual
        machine, ensuring that it does not move.

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        datacenter = vsphere.get_datacenter(name="dc-01")
        datastore = vsphere.get_datastore(name="datastore1",
            datacenter_id=datacenter.id)
        cluster = vsphere.get_compute_cluster(name="cluster-01",
            datacenter_id=datacenter.id)
        host = vsphere.get_host(name="esxi-01.example.com",
            datacenter_id=datacenter.id)
        network = vsphere.get_network(name="network1",
            datacenter_id=datacenter.id)
        vm = vsphere.VirtualMachine("vm",
            name="test",
            resource_pool_id=cluster.resource_pool_id,
            host_system_id=host.id,
            datastore_id=datastore.id,
            num_cpus=2,
            memory=2048,
            guest_id="otherLinux64Guest",
            network_interfaces=[{
                "network_id": network.id,
            }],
            disks=[{
                "label": "disk0",
                "size": 20,
            }])
        drs_vm_override = vsphere.DrsVmOverride("drs_vm_override",
            compute_cluster_id=cluster.id,
            virtual_machine_id=vm.id,
            drs_enabled=False)
        ```

        ## Import

        An existing override can be imported into this resource by

        supplying both the path to the cluster, and the path to the virtual machine, to

        `pulumi import`. If no override exists, an error will be given.  An example

        is below:

        [docs-import]: https://developer.hashicorp.com/terraform/cli/import

        ```sh
        $ pulumi import vsphere:index/drsVmOverride:DrsVmOverride drs_vm_override \\
        ```

          '{"compute_cluster_path": "/dc1/host/cluster1", \\

          "virtual_machine_path": "/dc1/vm/srv1"}'

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] compute_cluster_id: The managed object reference
               ID of the cluster to put the override in.  Forces a new
               resource if changed.
        :param pulumi.Input[builtins.str] drs_automation_level: Overrides the automation level for this virtual
               machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
               `fullyAutomated`. Default: `manual`.
               
               > **NOTE:** Using this resource _always_ implies an override, even if one of
               `drs_enabled` or `drs_automation_level` is omitted. Take note of the defaults
               for both options.
        :param pulumi.Input[builtins.bool] drs_enabled: Overrides the default DRS setting for this virtual
               machine. Can be either `true` or `false`. Default: `false`.
        :param pulumi.Input[builtins.str] virtual_machine_id: The UUID of the virtual machine to create
               the override for.  Forces a new resource if changed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DrsVmOverrideArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `DrsVmOverride` resource can be used to add a DRS override to a
        cluster for a specific virtual machine. With this resource, one can enable or
        disable DRS and control the automation level for a single virtual machine
        without affecting the rest of the cluster.

        For more information on vSphere clusters and DRS, see [this
        page][ref-vsphere-drs-clusters].

        [ref-vsphere-drs-clusters]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-resource-management-8-0/creating-a-drs-cluster.html

        > **NOTE:** This resource requires vCenter and is not available on direct ESXi
        connections.

        ## Example Usage

        The example below creates a virtual machine in a cluster using the
        `VirtualMachine` resource, creating the
        virtual machine in the cluster looked up by the
        `ComputeCluster` data source, but also
        pinning the VM to a host defined by the
        `Host` data source, which is assumed to
        be a host within the cluster. To ensure that the VM stays on this host and does
        not need to be migrated back at any point in time, an override is entered using
        the `DrsVmOverride` resource that disables DRS for this virtual
        machine, ensuring that it does not move.

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        datacenter = vsphere.get_datacenter(name="dc-01")
        datastore = vsphere.get_datastore(name="datastore1",
            datacenter_id=datacenter.id)
        cluster = vsphere.get_compute_cluster(name="cluster-01",
            datacenter_id=datacenter.id)
        host = vsphere.get_host(name="esxi-01.example.com",
            datacenter_id=datacenter.id)
        network = vsphere.get_network(name="network1",
            datacenter_id=datacenter.id)
        vm = vsphere.VirtualMachine("vm",
            name="test",
            resource_pool_id=cluster.resource_pool_id,
            host_system_id=host.id,
            datastore_id=datastore.id,
            num_cpus=2,
            memory=2048,
            guest_id="otherLinux64Guest",
            network_interfaces=[{
                "network_id": network.id,
            }],
            disks=[{
                "label": "disk0",
                "size": 20,
            }])
        drs_vm_override = vsphere.DrsVmOverride("drs_vm_override",
            compute_cluster_id=cluster.id,
            virtual_machine_id=vm.id,
            drs_enabled=False)
        ```

        ## Import

        An existing override can be imported into this resource by

        supplying both the path to the cluster, and the path to the virtual machine, to

        `pulumi import`. If no override exists, an error will be given.  An example

        is below:

        [docs-import]: https://developer.hashicorp.com/terraform/cli/import

        ```sh
        $ pulumi import vsphere:index/drsVmOverride:DrsVmOverride drs_vm_override \\
        ```

          '{"compute_cluster_path": "/dc1/host/cluster1", \\

          "virtual_machine_path": "/dc1/vm/srv1"}'

        :param str resource_name: The name of the resource.
        :param DrsVmOverrideArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DrsVmOverrideArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 drs_automation_level: Optional[pulumi.Input[builtins.str]] = None,
                 drs_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 virtual_machine_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DrsVmOverrideArgs.__new__(DrsVmOverrideArgs)

            if compute_cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'compute_cluster_id'")
            __props__.__dict__["compute_cluster_id"] = compute_cluster_id
            __props__.__dict__["drs_automation_level"] = drs_automation_level
            __props__.__dict__["drs_enabled"] = drs_enabled
            if virtual_machine_id is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_machine_id'")
            __props__.__dict__["virtual_machine_id"] = virtual_machine_id
        super(DrsVmOverride, __self__).__init__(
            'vsphere:index/drsVmOverride:DrsVmOverride',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            compute_cluster_id: Optional[pulumi.Input[builtins.str]] = None,
            drs_automation_level: Optional[pulumi.Input[builtins.str]] = None,
            drs_enabled: Optional[pulumi.Input[builtins.bool]] = None,
            virtual_machine_id: Optional[pulumi.Input[builtins.str]] = None) -> 'DrsVmOverride':
        """
        Get an existing DrsVmOverride resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] compute_cluster_id: The managed object reference
               ID of the cluster to put the override in.  Forces a new
               resource if changed.
        :param pulumi.Input[builtins.str] drs_automation_level: Overrides the automation level for this virtual
               machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
               `fullyAutomated`. Default: `manual`.
               
               > **NOTE:** Using this resource _always_ implies an override, even if one of
               `drs_enabled` or `drs_automation_level` is omitted. Take note of the defaults
               for both options.
        :param pulumi.Input[builtins.bool] drs_enabled: Overrides the default DRS setting for this virtual
               machine. Can be either `true` or `false`. Default: `false`.
        :param pulumi.Input[builtins.str] virtual_machine_id: The UUID of the virtual machine to create
               the override for.  Forces a new resource if changed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DrsVmOverrideState.__new__(_DrsVmOverrideState)

        __props__.__dict__["compute_cluster_id"] = compute_cluster_id
        __props__.__dict__["drs_automation_level"] = drs_automation_level
        __props__.__dict__["drs_enabled"] = drs_enabled
        __props__.__dict__["virtual_machine_id"] = virtual_machine_id
        return DrsVmOverride(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="computeClusterId")
    def compute_cluster_id(self) -> pulumi.Output[builtins.str]:
        """
        The managed object reference
        ID of the cluster to put the override in.  Forces a new
        resource if changed.
        """
        return pulumi.get(self, "compute_cluster_id")

    @property
    @pulumi.getter(name="drsAutomationLevel")
    def drs_automation_level(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Overrides the automation level for this virtual
        machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
        `fullyAutomated`. Default: `manual`.

        > **NOTE:** Using this resource _always_ implies an override, even if one of
        `drs_enabled` or `drs_automation_level` is omitted. Take note of the defaults
        for both options.
        """
        return pulumi.get(self, "drs_automation_level")

    @property
    @pulumi.getter(name="drsEnabled")
    def drs_enabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Overrides the default DRS setting for this virtual
        machine. Can be either `true` or `false`. Default: `false`.
        """
        return pulumi.get(self, "drs_enabled")

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> pulumi.Output[builtins.str]:
        """
        The UUID of the virtual machine to create
        the override for.  Forces a new resource if changed.
        """
        return pulumi.get(self, "virtual_machine_id")


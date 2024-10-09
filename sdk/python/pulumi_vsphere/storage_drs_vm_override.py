# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = ['StorageDrsVmOverrideArgs', 'StorageDrsVmOverride']

@pulumi.input_type
class StorageDrsVmOverrideArgs:
    def __init__(__self__, *,
                 datastore_cluster_id: pulumi.Input[str],
                 virtual_machine_id: pulumi.Input[str],
                 sdrs_automation_level: Optional[pulumi.Input[str]] = None,
                 sdrs_enabled: Optional[pulumi.Input[str]] = None,
                 sdrs_intra_vm_affinity: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a StorageDrsVmOverride resource.
        :param pulumi.Input[str] datastore_cluster_id: The managed object reference
               ID of the datastore cluster to put the override in.
               Forces a new resource if changed.
        :param pulumi.Input[str] virtual_machine_id: The UUID of the virtual machine to create
               the override for.  Forces a new resource if changed.
        :param pulumi.Input[str] sdrs_automation_level: Overrides any Storage DRS automation
               levels for this virtual machine. Can be one of `automated` or `manual`. When
               not specified, the datastore cluster's settings are used according to the
               specific SDRS subsystem.
        :param pulumi.Input[str] sdrs_enabled: Overrides the default Storage DRS setting for
               this virtual machine. When not specified, the datastore cluster setting is
               used.
        :param pulumi.Input[str] sdrs_intra_vm_affinity: Overrides the intra-VM affinity setting
               for this virtual machine. When `true`, all disks for this virtual machine
               will be kept on the same datastore. When `false`, Storage DRS may locate
               individual disks on different datastores if it helps satisfy cluster
               requirements. When not specified, the datastore cluster's settings are used.
        """
        pulumi.set(__self__, "datastore_cluster_id", datastore_cluster_id)
        pulumi.set(__self__, "virtual_machine_id", virtual_machine_id)
        if sdrs_automation_level is not None:
            pulumi.set(__self__, "sdrs_automation_level", sdrs_automation_level)
        if sdrs_enabled is not None:
            pulumi.set(__self__, "sdrs_enabled", sdrs_enabled)
        if sdrs_intra_vm_affinity is not None:
            pulumi.set(__self__, "sdrs_intra_vm_affinity", sdrs_intra_vm_affinity)

    @property
    @pulumi.getter(name="datastoreClusterId")
    def datastore_cluster_id(self) -> pulumi.Input[str]:
        """
        The managed object reference
        ID of the datastore cluster to put the override in.
        Forces a new resource if changed.
        """
        return pulumi.get(self, "datastore_cluster_id")

    @datastore_cluster_id.setter
    def datastore_cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "datastore_cluster_id", value)

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> pulumi.Input[str]:
        """
        The UUID of the virtual machine to create
        the override for.  Forces a new resource if changed.
        """
        return pulumi.get(self, "virtual_machine_id")

    @virtual_machine_id.setter
    def virtual_machine_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "virtual_machine_id", value)

    @property
    @pulumi.getter(name="sdrsAutomationLevel")
    def sdrs_automation_level(self) -> Optional[pulumi.Input[str]]:
        """
        Overrides any Storage DRS automation
        levels for this virtual machine. Can be one of `automated` or `manual`. When
        not specified, the datastore cluster's settings are used according to the
        specific SDRS subsystem.
        """
        return pulumi.get(self, "sdrs_automation_level")

    @sdrs_automation_level.setter
    def sdrs_automation_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sdrs_automation_level", value)

    @property
    @pulumi.getter(name="sdrsEnabled")
    def sdrs_enabled(self) -> Optional[pulumi.Input[str]]:
        """
        Overrides the default Storage DRS setting for
        this virtual machine. When not specified, the datastore cluster setting is
        used.
        """
        return pulumi.get(self, "sdrs_enabled")

    @sdrs_enabled.setter
    def sdrs_enabled(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sdrs_enabled", value)

    @property
    @pulumi.getter(name="sdrsIntraVmAffinity")
    def sdrs_intra_vm_affinity(self) -> Optional[pulumi.Input[str]]:
        """
        Overrides the intra-VM affinity setting
        for this virtual machine. When `true`, all disks for this virtual machine
        will be kept on the same datastore. When `false`, Storage DRS may locate
        individual disks on different datastores if it helps satisfy cluster
        requirements. When not specified, the datastore cluster's settings are used.
        """
        return pulumi.get(self, "sdrs_intra_vm_affinity")

    @sdrs_intra_vm_affinity.setter
    def sdrs_intra_vm_affinity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sdrs_intra_vm_affinity", value)


@pulumi.input_type
class _StorageDrsVmOverrideState:
    def __init__(__self__, *,
                 datastore_cluster_id: Optional[pulumi.Input[str]] = None,
                 sdrs_automation_level: Optional[pulumi.Input[str]] = None,
                 sdrs_enabled: Optional[pulumi.Input[str]] = None,
                 sdrs_intra_vm_affinity: Optional[pulumi.Input[str]] = None,
                 virtual_machine_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StorageDrsVmOverride resources.
        :param pulumi.Input[str] datastore_cluster_id: The managed object reference
               ID of the datastore cluster to put the override in.
               Forces a new resource if changed.
        :param pulumi.Input[str] sdrs_automation_level: Overrides any Storage DRS automation
               levels for this virtual machine. Can be one of `automated` or `manual`. When
               not specified, the datastore cluster's settings are used according to the
               specific SDRS subsystem.
        :param pulumi.Input[str] sdrs_enabled: Overrides the default Storage DRS setting for
               this virtual machine. When not specified, the datastore cluster setting is
               used.
        :param pulumi.Input[str] sdrs_intra_vm_affinity: Overrides the intra-VM affinity setting
               for this virtual machine. When `true`, all disks for this virtual machine
               will be kept on the same datastore. When `false`, Storage DRS may locate
               individual disks on different datastores if it helps satisfy cluster
               requirements. When not specified, the datastore cluster's settings are used.
        :param pulumi.Input[str] virtual_machine_id: The UUID of the virtual machine to create
               the override for.  Forces a new resource if changed.
        """
        if datastore_cluster_id is not None:
            pulumi.set(__self__, "datastore_cluster_id", datastore_cluster_id)
        if sdrs_automation_level is not None:
            pulumi.set(__self__, "sdrs_automation_level", sdrs_automation_level)
        if sdrs_enabled is not None:
            pulumi.set(__self__, "sdrs_enabled", sdrs_enabled)
        if sdrs_intra_vm_affinity is not None:
            pulumi.set(__self__, "sdrs_intra_vm_affinity", sdrs_intra_vm_affinity)
        if virtual_machine_id is not None:
            pulumi.set(__self__, "virtual_machine_id", virtual_machine_id)

    @property
    @pulumi.getter(name="datastoreClusterId")
    def datastore_cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed object reference
        ID of the datastore cluster to put the override in.
        Forces a new resource if changed.
        """
        return pulumi.get(self, "datastore_cluster_id")

    @datastore_cluster_id.setter
    def datastore_cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "datastore_cluster_id", value)

    @property
    @pulumi.getter(name="sdrsAutomationLevel")
    def sdrs_automation_level(self) -> Optional[pulumi.Input[str]]:
        """
        Overrides any Storage DRS automation
        levels for this virtual machine. Can be one of `automated` or `manual`. When
        not specified, the datastore cluster's settings are used according to the
        specific SDRS subsystem.
        """
        return pulumi.get(self, "sdrs_automation_level")

    @sdrs_automation_level.setter
    def sdrs_automation_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sdrs_automation_level", value)

    @property
    @pulumi.getter(name="sdrsEnabled")
    def sdrs_enabled(self) -> Optional[pulumi.Input[str]]:
        """
        Overrides the default Storage DRS setting for
        this virtual machine. When not specified, the datastore cluster setting is
        used.
        """
        return pulumi.get(self, "sdrs_enabled")

    @sdrs_enabled.setter
    def sdrs_enabled(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sdrs_enabled", value)

    @property
    @pulumi.getter(name="sdrsIntraVmAffinity")
    def sdrs_intra_vm_affinity(self) -> Optional[pulumi.Input[str]]:
        """
        Overrides the intra-VM affinity setting
        for this virtual machine. When `true`, all disks for this virtual machine
        will be kept on the same datastore. When `false`, Storage DRS may locate
        individual disks on different datastores if it helps satisfy cluster
        requirements. When not specified, the datastore cluster's settings are used.
        """
        return pulumi.get(self, "sdrs_intra_vm_affinity")

    @sdrs_intra_vm_affinity.setter
    def sdrs_intra_vm_affinity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sdrs_intra_vm_affinity", value)

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID of the virtual machine to create
        the override for.  Forces a new resource if changed.
        """
        return pulumi.get(self, "virtual_machine_id")

    @virtual_machine_id.setter
    def virtual_machine_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_machine_id", value)


class StorageDrsVmOverride(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 datastore_cluster_id: Optional[pulumi.Input[str]] = None,
                 sdrs_automation_level: Optional[pulumi.Input[str]] = None,
                 sdrs_enabled: Optional[pulumi.Input[str]] = None,
                 sdrs_intra_vm_affinity: Optional[pulumi.Input[str]] = None,
                 virtual_machine_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `StorageDrsVmOverride` resource can be used to add a Storage DRS
        override to a datastore cluster for a specific virtual machine. With this
        resource, one can enable or disable Storage DRS, and control the automation
        level and disk affinity for a single virtual machine without affecting the rest
        of the datastore cluster.

        For more information on vSphere datastore clusters and Storage DRS, see [this
        page][ref-vsphere-datastore-clusters].

        [ref-vsphere-datastore-clusters]: https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-resource-management/GUID-598DF695-107E-406B-9C95-0AF961FC227A.html

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        datacenter = vsphere.get_datacenter(name="dc-01")
        datastore_cluster = vsphere.get_datastore_cluster(name="datastore-cluster1",
            datacenter_id=datacenter.id)
        member_datastore = vsphere.get_datastore(name="datastore-cluster1-member1",
            datacenter_id=datacenter.id)
        pool = vsphere.get_resource_pool(name="cluster1/Resources",
            datacenter_id=datacenter.id)
        network = vsphere.get_network(name="public",
            datacenter_id=datacenter.id)
        vm = vsphere.VirtualMachine("vm",
            name="test",
            resource_pool_id=pool.id,
            datastore_id=member_datastore.id,
            num_cpus=2,
            memory=1024,
            guest_id="otherLinux64Guest",
            network_interfaces=[{
                "network_id": network.id,
            }],
            disks=[{
                "label": "disk0",
                "size": 20,
            }])
        drs_vm_override = vsphere.StorageDrsVmOverride("drs_vm_override",
            datastore_cluster_id=datastore_cluster.id,
            virtual_machine_id=vm.id,
            sdrs_enabled="false")
        ```

        ## Import

        An existing override can be imported into this resource by

        supplying both the path to the datastore cluster and the path to the virtual

        machine to `pulumi import`. If no override exists, an error will be given.

        An example is below:

        ```sh
        $ pulumi import vsphere:index/storageDrsVmOverride:StorageDrsVmOverride drs_vm_override \\
        ```

          '{"datastore_cluster_path": "/dc1/datastore/ds-cluster", \\

          "virtual_machine_path": "/dc1/vm/srv1"}'

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] datastore_cluster_id: The managed object reference
               ID of the datastore cluster to put the override in.
               Forces a new resource if changed.
        :param pulumi.Input[str] sdrs_automation_level: Overrides any Storage DRS automation
               levels for this virtual machine. Can be one of `automated` or `manual`. When
               not specified, the datastore cluster's settings are used according to the
               specific SDRS subsystem.
        :param pulumi.Input[str] sdrs_enabled: Overrides the default Storage DRS setting for
               this virtual machine. When not specified, the datastore cluster setting is
               used.
        :param pulumi.Input[str] sdrs_intra_vm_affinity: Overrides the intra-VM affinity setting
               for this virtual machine. When `true`, all disks for this virtual machine
               will be kept on the same datastore. When `false`, Storage DRS may locate
               individual disks on different datastores if it helps satisfy cluster
               requirements. When not specified, the datastore cluster's settings are used.
        :param pulumi.Input[str] virtual_machine_id: The UUID of the virtual machine to create
               the override for.  Forces a new resource if changed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StorageDrsVmOverrideArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `StorageDrsVmOverride` resource can be used to add a Storage DRS
        override to a datastore cluster for a specific virtual machine. With this
        resource, one can enable or disable Storage DRS, and control the automation
        level and disk affinity for a single virtual machine without affecting the rest
        of the datastore cluster.

        For more information on vSphere datastore clusters and Storage DRS, see [this
        page][ref-vsphere-datastore-clusters].

        [ref-vsphere-datastore-clusters]: https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-resource-management/GUID-598DF695-107E-406B-9C95-0AF961FC227A.html

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        datacenter = vsphere.get_datacenter(name="dc-01")
        datastore_cluster = vsphere.get_datastore_cluster(name="datastore-cluster1",
            datacenter_id=datacenter.id)
        member_datastore = vsphere.get_datastore(name="datastore-cluster1-member1",
            datacenter_id=datacenter.id)
        pool = vsphere.get_resource_pool(name="cluster1/Resources",
            datacenter_id=datacenter.id)
        network = vsphere.get_network(name="public",
            datacenter_id=datacenter.id)
        vm = vsphere.VirtualMachine("vm",
            name="test",
            resource_pool_id=pool.id,
            datastore_id=member_datastore.id,
            num_cpus=2,
            memory=1024,
            guest_id="otherLinux64Guest",
            network_interfaces=[{
                "network_id": network.id,
            }],
            disks=[{
                "label": "disk0",
                "size": 20,
            }])
        drs_vm_override = vsphere.StorageDrsVmOverride("drs_vm_override",
            datastore_cluster_id=datastore_cluster.id,
            virtual_machine_id=vm.id,
            sdrs_enabled="false")
        ```

        ## Import

        An existing override can be imported into this resource by

        supplying both the path to the datastore cluster and the path to the virtual

        machine to `pulumi import`. If no override exists, an error will be given.

        An example is below:

        ```sh
        $ pulumi import vsphere:index/storageDrsVmOverride:StorageDrsVmOverride drs_vm_override \\
        ```

          '{"datastore_cluster_path": "/dc1/datastore/ds-cluster", \\

          "virtual_machine_path": "/dc1/vm/srv1"}'

        :param str resource_name: The name of the resource.
        :param StorageDrsVmOverrideArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StorageDrsVmOverrideArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 datastore_cluster_id: Optional[pulumi.Input[str]] = None,
                 sdrs_automation_level: Optional[pulumi.Input[str]] = None,
                 sdrs_enabled: Optional[pulumi.Input[str]] = None,
                 sdrs_intra_vm_affinity: Optional[pulumi.Input[str]] = None,
                 virtual_machine_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StorageDrsVmOverrideArgs.__new__(StorageDrsVmOverrideArgs)

            if datastore_cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'datastore_cluster_id'")
            __props__.__dict__["datastore_cluster_id"] = datastore_cluster_id
            __props__.__dict__["sdrs_automation_level"] = sdrs_automation_level
            __props__.__dict__["sdrs_enabled"] = sdrs_enabled
            __props__.__dict__["sdrs_intra_vm_affinity"] = sdrs_intra_vm_affinity
            if virtual_machine_id is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_machine_id'")
            __props__.__dict__["virtual_machine_id"] = virtual_machine_id
        super(StorageDrsVmOverride, __self__).__init__(
            'vsphere:index/storageDrsVmOverride:StorageDrsVmOverride',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            datastore_cluster_id: Optional[pulumi.Input[str]] = None,
            sdrs_automation_level: Optional[pulumi.Input[str]] = None,
            sdrs_enabled: Optional[pulumi.Input[str]] = None,
            sdrs_intra_vm_affinity: Optional[pulumi.Input[str]] = None,
            virtual_machine_id: Optional[pulumi.Input[str]] = None) -> 'StorageDrsVmOverride':
        """
        Get an existing StorageDrsVmOverride resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] datastore_cluster_id: The managed object reference
               ID of the datastore cluster to put the override in.
               Forces a new resource if changed.
        :param pulumi.Input[str] sdrs_automation_level: Overrides any Storage DRS automation
               levels for this virtual machine. Can be one of `automated` or `manual`. When
               not specified, the datastore cluster's settings are used according to the
               specific SDRS subsystem.
        :param pulumi.Input[str] sdrs_enabled: Overrides the default Storage DRS setting for
               this virtual machine. When not specified, the datastore cluster setting is
               used.
        :param pulumi.Input[str] sdrs_intra_vm_affinity: Overrides the intra-VM affinity setting
               for this virtual machine. When `true`, all disks for this virtual machine
               will be kept on the same datastore. When `false`, Storage DRS may locate
               individual disks on different datastores if it helps satisfy cluster
               requirements. When not specified, the datastore cluster's settings are used.
        :param pulumi.Input[str] virtual_machine_id: The UUID of the virtual machine to create
               the override for.  Forces a new resource if changed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StorageDrsVmOverrideState.__new__(_StorageDrsVmOverrideState)

        __props__.__dict__["datastore_cluster_id"] = datastore_cluster_id
        __props__.__dict__["sdrs_automation_level"] = sdrs_automation_level
        __props__.__dict__["sdrs_enabled"] = sdrs_enabled
        __props__.__dict__["sdrs_intra_vm_affinity"] = sdrs_intra_vm_affinity
        __props__.__dict__["virtual_machine_id"] = virtual_machine_id
        return StorageDrsVmOverride(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="datastoreClusterId")
    def datastore_cluster_id(self) -> pulumi.Output[str]:
        """
        The managed object reference
        ID of the datastore cluster to put the override in.
        Forces a new resource if changed.
        """
        return pulumi.get(self, "datastore_cluster_id")

    @property
    @pulumi.getter(name="sdrsAutomationLevel")
    def sdrs_automation_level(self) -> pulumi.Output[Optional[str]]:
        """
        Overrides any Storage DRS automation
        levels for this virtual machine. Can be one of `automated` or `manual`. When
        not specified, the datastore cluster's settings are used according to the
        specific SDRS subsystem.
        """
        return pulumi.get(self, "sdrs_automation_level")

    @property
    @pulumi.getter(name="sdrsEnabled")
    def sdrs_enabled(self) -> pulumi.Output[Optional[str]]:
        """
        Overrides the default Storage DRS setting for
        this virtual machine. When not specified, the datastore cluster setting is
        used.
        """
        return pulumi.get(self, "sdrs_enabled")

    @property
    @pulumi.getter(name="sdrsIntraVmAffinity")
    def sdrs_intra_vm_affinity(self) -> pulumi.Output[Optional[str]]:
        """
        Overrides the intra-VM affinity setting
        for this virtual machine. When `true`, all disks for this virtual machine
        will be kept on the same datastore. When `false`, Storage DRS may locate
        individual disks on different datastores if it helps satisfy cluster
        requirements. When not specified, the datastore cluster's settings are used.
        """
        return pulumi.get(self, "sdrs_intra_vm_affinity")

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> pulumi.Output[str]:
        """
        The UUID of the virtual machine to create
        the override for.  Forces a new resource if changed.
        """
        return pulumi.get(self, "virtual_machine_id")


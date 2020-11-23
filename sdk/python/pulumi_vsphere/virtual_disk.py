# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['VirtualDisk']


class VirtualDisk(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adapter_type: Optional[pulumi.Input[str]] = None,
                 create_directories: Optional[pulumi.Input[bool]] = None,
                 datacenter: Optional[pulumi.Input[str]] = None,
                 datastore: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vmdk_path: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The `VirtualDisk` resource can be used to create virtual disks outside
        of any given `VirtualMachine`
        resource. These disks can be attached to a virtual machine by creating a disk
        block with the `attach` parameter.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        my_disk = vsphere.VirtualDisk("myDisk",
            datacenter="Datacenter",
            datastore="local",
            size=2,
            type="thin",
            vmdk_path="myDisk.vmdk")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adapter_type: The adapter type for this virtual disk. Can be
               one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
        :param pulumi.Input[bool] create_directories: Tells the resource to create any
               directories that are a part of the `vmdk_path` parameter if they are missing.
               Default: `false`.
        :param pulumi.Input[str] datacenter: The name of the datacenter in which to create the
               disk. Can be omitted when when ESXi or if there is only one datacenter in
               your infrastructure.
        :param pulumi.Input[str] datastore: The name of the datastore in which to create the
               disk.
        :param pulumi.Input[int] size: Size of the disk (in GB).
        :param pulumi.Input[str] type: The type of disk to create. Can be one of
               `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
               information on what each kind of disk provisioning policy means, click
               [here][docs-vmware-vm-disk-provisioning].
        :param pulumi.Input[str] vmdk_path: The path, including filename, of the virtual disk to
               be created.  This needs to end in `.vmdk`.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if adapter_type is not None:
                warnings.warn("""this attribute has no effect on controller types - please use scsi_type in vsphere_virtual_machine instead""", DeprecationWarning)
                pulumi.log.warn("adapter_type is deprecated: this attribute has no effect on controller types - please use scsi_type in vsphere_virtual_machine instead")
            __props__['adapter_type'] = adapter_type
            __props__['create_directories'] = create_directories
            __props__['datacenter'] = datacenter
            if datastore is None:
                raise TypeError("Missing required property 'datastore'")
            __props__['datastore'] = datastore
            if size is None:
                raise TypeError("Missing required property 'size'")
            __props__['size'] = size
            __props__['type'] = type
            if vmdk_path is None:
                raise TypeError("Missing required property 'vmdk_path'")
            __props__['vmdk_path'] = vmdk_path
        super(VirtualDisk, __self__).__init__(
            'vsphere:index/virtualDisk:VirtualDisk',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            adapter_type: Optional[pulumi.Input[str]] = None,
            create_directories: Optional[pulumi.Input[bool]] = None,
            datacenter: Optional[pulumi.Input[str]] = None,
            datastore: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            vmdk_path: Optional[pulumi.Input[str]] = None) -> 'VirtualDisk':
        """
        Get an existing VirtualDisk resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adapter_type: The adapter type for this virtual disk. Can be
               one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
        :param pulumi.Input[bool] create_directories: Tells the resource to create any
               directories that are a part of the `vmdk_path` parameter if they are missing.
               Default: `false`.
        :param pulumi.Input[str] datacenter: The name of the datacenter in which to create the
               disk. Can be omitted when when ESXi or if there is only one datacenter in
               your infrastructure.
        :param pulumi.Input[str] datastore: The name of the datastore in which to create the
               disk.
        :param pulumi.Input[int] size: Size of the disk (in GB).
        :param pulumi.Input[str] type: The type of disk to create. Can be one of
               `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
               information on what each kind of disk provisioning policy means, click
               [here][docs-vmware-vm-disk-provisioning].
        :param pulumi.Input[str] vmdk_path: The path, including filename, of the virtual disk to
               be created.  This needs to end in `.vmdk`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["adapter_type"] = adapter_type
        __props__["create_directories"] = create_directories
        __props__["datacenter"] = datacenter
        __props__["datastore"] = datastore
        __props__["size"] = size
        __props__["type"] = type
        __props__["vmdk_path"] = vmdk_path
        return VirtualDisk(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adapterType")
    def adapter_type(self) -> pulumi.Output[Optional[str]]:
        """
        The adapter type for this virtual disk. Can be
        one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
        """
        return pulumi.get(self, "adapter_type")

    @property
    @pulumi.getter(name="createDirectories")
    def create_directories(self) -> pulumi.Output[Optional[bool]]:
        """
        Tells the resource to create any
        directories that are a part of the `vmdk_path` parameter if they are missing.
        Default: `false`.
        """
        return pulumi.get(self, "create_directories")

    @property
    @pulumi.getter
    def datacenter(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the datacenter in which to create the
        disk. Can be omitted when when ESXi or if there is only one datacenter in
        your infrastructure.
        """
        return pulumi.get(self, "datacenter")

    @property
    @pulumi.getter
    def datastore(self) -> pulumi.Output[str]:
        """
        The name of the datastore in which to create the
        disk.
        """
        return pulumi.get(self, "datastore")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        Size of the disk (in GB).
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of disk to create. Can be one of
        `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
        information on what each kind of disk provisioning policy means, click
        [here][docs-vmware-vm-disk-provisioning].
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vmdkPath")
    def vmdk_path(self) -> pulumi.Output[str]:
        """
        The path, including filename, of the virtual disk to
        be created.  This needs to end in `.vmdk`.
        """
        return pulumi.get(self, "vmdk_path")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


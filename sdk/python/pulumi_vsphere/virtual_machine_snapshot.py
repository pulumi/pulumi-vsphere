# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import _utilities, _tables


class VirtualMachineSnapshot(pulumi.CustomResource):
    consolidate: pulumi.Output[bool]
    """
    If set to `true`, the delta disks involved in this
    snapshot will be consolidated into the parent when this resource is
    destroyed.
    """
    description: pulumi.Output[str]
    """
    A description for the snapshot.
    """
    memory: pulumi.Output[bool]
    """
    If set to `true`, a dump of the internal state of the
    virtual machine is included in the snapshot.
    """
    quiesce: pulumi.Output[bool]
    """
    If set to `true`, and the virtual machine is powered
    on when the snapshot is taken, VMware Tools is used to quiesce the file
    system in the virtual machine.
    """
    remove_children: pulumi.Output[bool]
    """
    If set to `true`, the entire snapshot subtree
    is removed when this resource is destroyed.
    """
    snapshot_name: pulumi.Output[str]
    """
    The name of the snapshot.
    """
    virtual_machine_uuid: pulumi.Output[str]
    """
    The virtual machine UUID.
    """
    def __init__(__self__, resource_name, opts=None, consolidate=None, description=None, memory=None, quiesce=None, remove_children=None, snapshot_name=None, virtual_machine_uuid=None, __props__=None, __name__=None, __opts__=None):
        """
        The `VirtualMachineSnapshot` resource can be used to manage snapshots
        for a virtual machine.

        For more information on managing snapshots and how they work in VMware, see
        [here][ext-vm-snapshot-management].

        [ext-vm-snapshot-management]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vm_admin.doc/GUID-CA948C69-7F58-4519-AEB1-739545EA94E5.html

        > **NOTE:** A snapshot in VMware differs from traditional disk snapshots, and
        can contain the actual running state of the virtual machine, data for all disks
        that have not been set to be independent from the snapshot (including ones that
        have been attached via the `attach`
        parameter to the `VirtualMachine` `disk` block), and even the
        configuration of the virtual machine at the time of the snapshot. Virtual
        machine, disk activity, and configuration changes post-snapshot are not
        included in the original state. Use this resource with care! Neither VMware nor
        HashiCorp recommends retaining snapshots for a extended period of time and does
        NOT recommend using them as as backup feature. For more information on the
        limitation of virtual machine snapshots, see [here][ext-vm-snap-limitations].

        [ext-vm-snap-limitations]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vm_admin.doc/GUID-53F65726-A23B-4CF0-A7D5-48E584B88613.html

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        demo1 = vsphere.VirtualMachineSnapshot("demo1",
            consolidate="true",
            description="This is Demo Snapshot",
            memory="true",
            quiesce="true",
            remove_children="false",
            snapshot_name="Snapshot Name",
            virtual_machine_uuid="9aac5551-a351-4158-8c5c-15a71e8ec5c9")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] consolidate: If set to `true`, the delta disks involved in this
               snapshot will be consolidated into the parent when this resource is
               destroyed.
        :param pulumi.Input[str] description: A description for the snapshot.
        :param pulumi.Input[bool] memory: If set to `true`, a dump of the internal state of the
               virtual machine is included in the snapshot.
        :param pulumi.Input[bool] quiesce: If set to `true`, and the virtual machine is powered
               on when the snapshot is taken, VMware Tools is used to quiesce the file
               system in the virtual machine.
        :param pulumi.Input[bool] remove_children: If set to `true`, the entire snapshot subtree
               is removed when this resource is destroyed.
        :param pulumi.Input[str] snapshot_name: The name of the snapshot.
        :param pulumi.Input[str] virtual_machine_uuid: The virtual machine UUID.
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

            __props__['consolidate'] = consolidate
            if description is None:
                raise TypeError("Missing required property 'description'")
            __props__['description'] = description
            if memory is None:
                raise TypeError("Missing required property 'memory'")
            __props__['memory'] = memory
            if quiesce is None:
                raise TypeError("Missing required property 'quiesce'")
            __props__['quiesce'] = quiesce
            __props__['remove_children'] = remove_children
            if snapshot_name is None:
                raise TypeError("Missing required property 'snapshot_name'")
            __props__['snapshot_name'] = snapshot_name
            if virtual_machine_uuid is None:
                raise TypeError("Missing required property 'virtual_machine_uuid'")
            __props__['virtual_machine_uuid'] = virtual_machine_uuid
        super(VirtualMachineSnapshot, __self__).__init__(
            'vsphere:index/virtualMachineSnapshot:VirtualMachineSnapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, consolidate=None, description=None, memory=None, quiesce=None, remove_children=None, snapshot_name=None, virtual_machine_uuid=None):
        """
        Get an existing VirtualMachineSnapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] consolidate: If set to `true`, the delta disks involved in this
               snapshot will be consolidated into the parent when this resource is
               destroyed.
        :param pulumi.Input[str] description: A description for the snapshot.
        :param pulumi.Input[bool] memory: If set to `true`, a dump of the internal state of the
               virtual machine is included in the snapshot.
        :param pulumi.Input[bool] quiesce: If set to `true`, and the virtual machine is powered
               on when the snapshot is taken, VMware Tools is used to quiesce the file
               system in the virtual machine.
        :param pulumi.Input[bool] remove_children: If set to `true`, the entire snapshot subtree
               is removed when this resource is destroyed.
        :param pulumi.Input[str] snapshot_name: The name of the snapshot.
        :param pulumi.Input[str] virtual_machine_uuid: The virtual machine UUID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["consolidate"] = consolidate
        __props__["description"] = description
        __props__["memory"] = memory
        __props__["quiesce"] = quiesce
        __props__["remove_children"] = remove_children
        __props__["snapshot_name"] = snapshot_name
        __props__["virtual_machine_uuid"] = virtual_machine_uuid
        return VirtualMachineSnapshot(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

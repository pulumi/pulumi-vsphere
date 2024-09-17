# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ComputeClusterVmGroupArgs', 'ComputeClusterVmGroup']

@pulumi.input_type
class ComputeClusterVmGroupArgs:
    def __init__(__self__, *,
                 compute_cluster_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 virtual_machine_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ComputeClusterVmGroup resource.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the group in.  Forces a new
               resource if changed.
        :param pulumi.Input[str] name: The name of the VM group. This must be unique in the
               cluster. Forces a new resource if changed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] virtual_machine_ids: The UUIDs of the virtual machines in this
               group.
               
               > **NOTE:** The namespace for cluster names on this resource (defined by the
               `name` argument) is shared with the
               `ComputeClusterHostGroup`
               resource. Make sure your names are unique across both resources.
               
               > **NOTE:** To update a existing VM group, you must first import the group with `import` command in
               import section. When importing a VM group, validate that all virtual machines that
               need to be in the group are included in the `virtual_machine_ids`; otherwise, any virtual machines
               that are not in `virtual_machine_ids` the included will be removed from the group.
        """
        pulumi.set(__self__, "compute_cluster_id", compute_cluster_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if virtual_machine_ids is not None:
            pulumi.set(__self__, "virtual_machine_ids", virtual_machine_ids)

    @property
    @pulumi.getter(name="computeClusterId")
    def compute_cluster_id(self) -> pulumi.Input[str]:
        """
        The managed object reference
        ID of the cluster to put the group in.  Forces a new
        resource if changed.
        """
        return pulumi.get(self, "compute_cluster_id")

    @compute_cluster_id.setter
    def compute_cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "compute_cluster_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the VM group. This must be unique in the
        cluster. Forces a new resource if changed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="virtualMachineIds")
    def virtual_machine_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The UUIDs of the virtual machines in this
        group.

        > **NOTE:** The namespace for cluster names on this resource (defined by the
        `name` argument) is shared with the
        `ComputeClusterHostGroup`
        resource. Make sure your names are unique across both resources.

        > **NOTE:** To update a existing VM group, you must first import the group with `import` command in
        import section. When importing a VM group, validate that all virtual machines that
        need to be in the group are included in the `virtual_machine_ids`; otherwise, any virtual machines
        that are not in `virtual_machine_ids` the included will be removed from the group.
        """
        return pulumi.get(self, "virtual_machine_ids")

    @virtual_machine_ids.setter
    def virtual_machine_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "virtual_machine_ids", value)


@pulumi.input_type
class _ComputeClusterVmGroupState:
    def __init__(__self__, *,
                 compute_cluster_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 virtual_machine_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ComputeClusterVmGroup resources.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the group in.  Forces a new
               resource if changed.
        :param pulumi.Input[str] name: The name of the VM group. This must be unique in the
               cluster. Forces a new resource if changed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] virtual_machine_ids: The UUIDs of the virtual machines in this
               group.
               
               > **NOTE:** The namespace for cluster names on this resource (defined by the
               `name` argument) is shared with the
               `ComputeClusterHostGroup`
               resource. Make sure your names are unique across both resources.
               
               > **NOTE:** To update a existing VM group, you must first import the group with `import` command in
               import section. When importing a VM group, validate that all virtual machines that
               need to be in the group are included in the `virtual_machine_ids`; otherwise, any virtual machines
               that are not in `virtual_machine_ids` the included will be removed from the group.
        """
        if compute_cluster_id is not None:
            pulumi.set(__self__, "compute_cluster_id", compute_cluster_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if virtual_machine_ids is not None:
            pulumi.set(__self__, "virtual_machine_ids", virtual_machine_ids)

    @property
    @pulumi.getter(name="computeClusterId")
    def compute_cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed object reference
        ID of the cluster to put the group in.  Forces a new
        resource if changed.
        """
        return pulumi.get(self, "compute_cluster_id")

    @compute_cluster_id.setter
    def compute_cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compute_cluster_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the VM group. This must be unique in the
        cluster. Forces a new resource if changed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="virtualMachineIds")
    def virtual_machine_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The UUIDs of the virtual machines in this
        group.

        > **NOTE:** The namespace for cluster names on this resource (defined by the
        `name` argument) is shared with the
        `ComputeClusterHostGroup`
        resource. Make sure your names are unique across both resources.

        > **NOTE:** To update a existing VM group, you must first import the group with `import` command in
        import section. When importing a VM group, validate that all virtual machines that
        need to be in the group are included in the `virtual_machine_ids`; otherwise, any virtual machines
        that are not in `virtual_machine_ids` the included will be removed from the group.
        """
        return pulumi.get(self, "virtual_machine_ids")

    @virtual_machine_ids.setter
    def virtual_machine_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "virtual_machine_ids", value)


class ComputeClusterVmGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_cluster_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 virtual_machine_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a ComputeClusterVmGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the group in.  Forces a new
               resource if changed.
        :param pulumi.Input[str] name: The name of the VM group. This must be unique in the
               cluster. Forces a new resource if changed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] virtual_machine_ids: The UUIDs of the virtual machines in this
               group.
               
               > **NOTE:** The namespace for cluster names on this resource (defined by the
               `name` argument) is shared with the
               `ComputeClusterHostGroup`
               resource. Make sure your names are unique across both resources.
               
               > **NOTE:** To update a existing VM group, you must first import the group with `import` command in
               import section. When importing a VM group, validate that all virtual machines that
               need to be in the group are included in the `virtual_machine_ids`; otherwise, any virtual machines
               that are not in `virtual_machine_ids` the included will be removed from the group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ComputeClusterVmGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ComputeClusterVmGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ComputeClusterVmGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ComputeClusterVmGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_cluster_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 virtual_machine_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ComputeClusterVmGroupArgs.__new__(ComputeClusterVmGroupArgs)

            if compute_cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'compute_cluster_id'")
            __props__.__dict__["compute_cluster_id"] = compute_cluster_id
            __props__.__dict__["name"] = name
            __props__.__dict__["virtual_machine_ids"] = virtual_machine_ids
        super(ComputeClusterVmGroup, __self__).__init__(
            'vsphere:index/computeClusterVmGroup:ComputeClusterVmGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            compute_cluster_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            virtual_machine_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'ComputeClusterVmGroup':
        """
        Get an existing ComputeClusterVmGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the group in.  Forces a new
               resource if changed.
        :param pulumi.Input[str] name: The name of the VM group. This must be unique in the
               cluster. Forces a new resource if changed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] virtual_machine_ids: The UUIDs of the virtual machines in this
               group.
               
               > **NOTE:** The namespace for cluster names on this resource (defined by the
               `name` argument) is shared with the
               `ComputeClusterHostGroup`
               resource. Make sure your names are unique across both resources.
               
               > **NOTE:** To update a existing VM group, you must first import the group with `import` command in
               import section. When importing a VM group, validate that all virtual machines that
               need to be in the group are included in the `virtual_machine_ids`; otherwise, any virtual machines
               that are not in `virtual_machine_ids` the included will be removed from the group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ComputeClusterVmGroupState.__new__(_ComputeClusterVmGroupState)

        __props__.__dict__["compute_cluster_id"] = compute_cluster_id
        __props__.__dict__["name"] = name
        __props__.__dict__["virtual_machine_ids"] = virtual_machine_ids
        return ComputeClusterVmGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="computeClusterId")
    def compute_cluster_id(self) -> pulumi.Output[str]:
        """
        The managed object reference
        ID of the cluster to put the group in.  Forces a new
        resource if changed.
        """
        return pulumi.get(self, "compute_cluster_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the VM group. This must be unique in the
        cluster. Forces a new resource if changed.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="virtualMachineIds")
    def virtual_machine_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The UUIDs of the virtual machines in this
        group.

        > **NOTE:** The namespace for cluster names on this resource (defined by the
        `name` argument) is shared with the
        `ComputeClusterHostGroup`
        resource. Make sure your names are unique across both resources.

        > **NOTE:** To update a existing VM group, you must first import the group with `import` command in
        import section. When importing a VM group, validate that all virtual machines that
        need to be in the group are included in the `virtual_machine_ids`; otherwise, any virtual machines
        that are not in `virtual_machine_ids` the included will be removed from the group.
        """
        return pulumi.get(self, "virtual_machine_ids")


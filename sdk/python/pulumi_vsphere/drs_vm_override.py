# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DrsVmOverrideArgs', 'DrsVmOverride']

@pulumi.input_type
class DrsVmOverrideArgs:
    def __init__(__self__, *,
                 compute_cluster_id: pulumi.Input[str],
                 virtual_machine_id: pulumi.Input[str],
                 drs_automation_level: Optional[pulumi.Input[str]] = None,
                 drs_enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a DrsVmOverride resource.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the override in.  Forces a new
               resource if changed.
        :param pulumi.Input[str] virtual_machine_id: The UUID of the virtual machine to create
               the override for.  Forces a new resource if changed.
        :param pulumi.Input[str] drs_automation_level: Overrides the automation level for this virtual
               machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
               `fullyAutomated`. Default: `manual`.
               
               > **NOTE:** Using this resource _always_ implies an override, even if one of
               `drs_enabled` or `drs_automation_level` is omitted. Take note of the defaults
               for both options.
        :param pulumi.Input[bool] drs_enabled: Overrides the default DRS setting for this virtual
               machine. Can be either `true` or `false`. Default: `false`.
        """
        DrsVmOverrideArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            compute_cluster_id=compute_cluster_id,
            virtual_machine_id=virtual_machine_id,
            drs_automation_level=drs_automation_level,
            drs_enabled=drs_enabled,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             compute_cluster_id: pulumi.Input[str],
             virtual_machine_id: pulumi.Input[str],
             drs_automation_level: Optional[pulumi.Input[str]] = None,
             drs_enabled: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("compute_cluster_id", compute_cluster_id)
        _setter("virtual_machine_id", virtual_machine_id)
        if drs_automation_level is not None:
            _setter("drs_automation_level", drs_automation_level)
        if drs_enabled is not None:
            _setter("drs_enabled", drs_enabled)

    @property
    @pulumi.getter(name="computeClusterId")
    def compute_cluster_id(self) -> pulumi.Input[str]:
        """
        The managed object reference
        ID of the cluster to put the override in.  Forces a new
        resource if changed.
        """
        return pulumi.get(self, "compute_cluster_id")

    @compute_cluster_id.setter
    def compute_cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "compute_cluster_id", value)

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
    @pulumi.getter(name="drsAutomationLevel")
    def drs_automation_level(self) -> Optional[pulumi.Input[str]]:
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
    def drs_automation_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "drs_automation_level", value)

    @property
    @pulumi.getter(name="drsEnabled")
    def drs_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Overrides the default DRS setting for this virtual
        machine. Can be either `true` or `false`. Default: `false`.
        """
        return pulumi.get(self, "drs_enabled")

    @drs_enabled.setter
    def drs_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "drs_enabled", value)


@pulumi.input_type
class _DrsVmOverrideState:
    def __init__(__self__, *,
                 compute_cluster_id: Optional[pulumi.Input[str]] = None,
                 drs_automation_level: Optional[pulumi.Input[str]] = None,
                 drs_enabled: Optional[pulumi.Input[bool]] = None,
                 virtual_machine_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DrsVmOverride resources.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the override in.  Forces a new
               resource if changed.
        :param pulumi.Input[str] drs_automation_level: Overrides the automation level for this virtual
               machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
               `fullyAutomated`. Default: `manual`.
               
               > **NOTE:** Using this resource _always_ implies an override, even if one of
               `drs_enabled` or `drs_automation_level` is omitted. Take note of the defaults
               for both options.
        :param pulumi.Input[bool] drs_enabled: Overrides the default DRS setting for this virtual
               machine. Can be either `true` or `false`. Default: `false`.
        :param pulumi.Input[str] virtual_machine_id: The UUID of the virtual machine to create
               the override for.  Forces a new resource if changed.
        """
        _DrsVmOverrideState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            compute_cluster_id=compute_cluster_id,
            drs_automation_level=drs_automation_level,
            drs_enabled=drs_enabled,
            virtual_machine_id=virtual_machine_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             compute_cluster_id: Optional[pulumi.Input[str]] = None,
             drs_automation_level: Optional[pulumi.Input[str]] = None,
             drs_enabled: Optional[pulumi.Input[bool]] = None,
             virtual_machine_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if compute_cluster_id is not None:
            _setter("compute_cluster_id", compute_cluster_id)
        if drs_automation_level is not None:
            _setter("drs_automation_level", drs_automation_level)
        if drs_enabled is not None:
            _setter("drs_enabled", drs_enabled)
        if virtual_machine_id is not None:
            _setter("virtual_machine_id", virtual_machine_id)

    @property
    @pulumi.getter(name="computeClusterId")
    def compute_cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed object reference
        ID of the cluster to put the override in.  Forces a new
        resource if changed.
        """
        return pulumi.get(self, "compute_cluster_id")

    @compute_cluster_id.setter
    def compute_cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compute_cluster_id", value)

    @property
    @pulumi.getter(name="drsAutomationLevel")
    def drs_automation_level(self) -> Optional[pulumi.Input[str]]:
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
    def drs_automation_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "drs_automation_level", value)

    @property
    @pulumi.getter(name="drsEnabled")
    def drs_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Overrides the default DRS setting for this virtual
        machine. Can be either `true` or `false`. Default: `false`.
        """
        return pulumi.get(self, "drs_enabled")

    @drs_enabled.setter
    def drs_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "drs_enabled", value)

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


class DrsVmOverride(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_cluster_id: Optional[pulumi.Input[str]] = None,
                 drs_automation_level: Optional[pulumi.Input[str]] = None,
                 drs_enabled: Optional[pulumi.Input[bool]] = None,
                 virtual_machine_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a DrsVmOverride resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the override in.  Forces a new
               resource if changed.
        :param pulumi.Input[str] drs_automation_level: Overrides the automation level for this virtual
               machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
               `fullyAutomated`. Default: `manual`.
               
               > **NOTE:** Using this resource _always_ implies an override, even if one of
               `drs_enabled` or `drs_automation_level` is omitted. Take note of the defaults
               for both options.
        :param pulumi.Input[bool] drs_enabled: Overrides the default DRS setting for this virtual
               machine. Can be either `true` or `false`. Default: `false`.
        :param pulumi.Input[str] virtual_machine_id: The UUID of the virtual machine to create
               the override for.  Forces a new resource if changed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DrsVmOverrideArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DrsVmOverride resource with the given unique name, props, and options.
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
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            DrsVmOverrideArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_cluster_id: Optional[pulumi.Input[str]] = None,
                 drs_automation_level: Optional[pulumi.Input[str]] = None,
                 drs_enabled: Optional[pulumi.Input[bool]] = None,
                 virtual_machine_id: Optional[pulumi.Input[str]] = None,
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
            compute_cluster_id: Optional[pulumi.Input[str]] = None,
            drs_automation_level: Optional[pulumi.Input[str]] = None,
            drs_enabled: Optional[pulumi.Input[bool]] = None,
            virtual_machine_id: Optional[pulumi.Input[str]] = None) -> 'DrsVmOverride':
        """
        Get an existing DrsVmOverride resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the override in.  Forces a new
               resource if changed.
        :param pulumi.Input[str] drs_automation_level: Overrides the automation level for this virtual
               machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
               `fullyAutomated`. Default: `manual`.
               
               > **NOTE:** Using this resource _always_ implies an override, even if one of
               `drs_enabled` or `drs_automation_level` is omitted. Take note of the defaults
               for both options.
        :param pulumi.Input[bool] drs_enabled: Overrides the default DRS setting for this virtual
               machine. Can be either `true` or `false`. Default: `false`.
        :param pulumi.Input[str] virtual_machine_id: The UUID of the virtual machine to create
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
    def compute_cluster_id(self) -> pulumi.Output[str]:
        """
        The managed object reference
        ID of the cluster to put the override in.  Forces a new
        resource if changed.
        """
        return pulumi.get(self, "compute_cluster_id")

    @property
    @pulumi.getter(name="drsAutomationLevel")
    def drs_automation_level(self) -> pulumi.Output[Optional[str]]:
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
    def drs_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Overrides the default DRS setting for this virtual
        machine. Can be either `true` or `false`. Default: `false`.
        """
        return pulumi.get(self, "drs_enabled")

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> pulumi.Output[str]:
        """
        The UUID of the virtual machine to create
        the override for.  Forces a new resource if changed.
        """
        return pulumi.get(self, "virtual_machine_id")


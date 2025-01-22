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

__all__ = ['DistributedVirtualSwitchPvlanMappingInitArgs', 'DistributedVirtualSwitchPvlanMapping']

@pulumi.input_type
class DistributedVirtualSwitchPvlanMappingInitArgs:
    def __init__(__self__, *,
                 distributed_virtual_switch_id: pulumi.Input[str],
                 primary_vlan_id: pulumi.Input[int],
                 pvlan_type: pulumi.Input[str],
                 secondary_vlan_id: pulumi.Input[int]):
        """
        The set of arguments for constructing a DistributedVirtualSwitchPvlanMapping resource.
        :param pulumi.Input[str] distributed_virtual_switch_id: The ID of the distributed virtual switch to attach this mapping to.
        :param pulumi.Input[int] primary_vlan_id: The primary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        :param pulumi.Input[str] pvlan_type: The private VLAN type. Valid values are promiscuous, community and isolated.
        :param pulumi.Input[int] secondary_vlan_id: The secondary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        """
        pulumi.set(__self__, "distributed_virtual_switch_id", distributed_virtual_switch_id)
        pulumi.set(__self__, "primary_vlan_id", primary_vlan_id)
        pulumi.set(__self__, "pvlan_type", pvlan_type)
        pulumi.set(__self__, "secondary_vlan_id", secondary_vlan_id)

    @property
    @pulumi.getter(name="distributedVirtualSwitchId")
    def distributed_virtual_switch_id(self) -> pulumi.Input[str]:
        """
        The ID of the distributed virtual switch to attach this mapping to.
        """
        return pulumi.get(self, "distributed_virtual_switch_id")

    @distributed_virtual_switch_id.setter
    def distributed_virtual_switch_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "distributed_virtual_switch_id", value)

    @property
    @pulumi.getter(name="primaryVlanId")
    def primary_vlan_id(self) -> pulumi.Input[int]:
        """
        The primary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        """
        return pulumi.get(self, "primary_vlan_id")

    @primary_vlan_id.setter
    def primary_vlan_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "primary_vlan_id", value)

    @property
    @pulumi.getter(name="pvlanType")
    def pvlan_type(self) -> pulumi.Input[str]:
        """
        The private VLAN type. Valid values are promiscuous, community and isolated.
        """
        return pulumi.get(self, "pvlan_type")

    @pvlan_type.setter
    def pvlan_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "pvlan_type", value)

    @property
    @pulumi.getter(name="secondaryVlanId")
    def secondary_vlan_id(self) -> pulumi.Input[int]:
        """
        The secondary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        """
        return pulumi.get(self, "secondary_vlan_id")

    @secondary_vlan_id.setter
    def secondary_vlan_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "secondary_vlan_id", value)


@pulumi.input_type
class _DistributedVirtualSwitchPvlanMappingState:
    def __init__(__self__, *,
                 distributed_virtual_switch_id: Optional[pulumi.Input[str]] = None,
                 primary_vlan_id: Optional[pulumi.Input[int]] = None,
                 pvlan_type: Optional[pulumi.Input[str]] = None,
                 secondary_vlan_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering DistributedVirtualSwitchPvlanMapping resources.
        :param pulumi.Input[str] distributed_virtual_switch_id: The ID of the distributed virtual switch to attach this mapping to.
        :param pulumi.Input[int] primary_vlan_id: The primary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        :param pulumi.Input[str] pvlan_type: The private VLAN type. Valid values are promiscuous, community and isolated.
        :param pulumi.Input[int] secondary_vlan_id: The secondary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        """
        if distributed_virtual_switch_id is not None:
            pulumi.set(__self__, "distributed_virtual_switch_id", distributed_virtual_switch_id)
        if primary_vlan_id is not None:
            pulumi.set(__self__, "primary_vlan_id", primary_vlan_id)
        if pvlan_type is not None:
            pulumi.set(__self__, "pvlan_type", pvlan_type)
        if secondary_vlan_id is not None:
            pulumi.set(__self__, "secondary_vlan_id", secondary_vlan_id)

    @property
    @pulumi.getter(name="distributedVirtualSwitchId")
    def distributed_virtual_switch_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the distributed virtual switch to attach this mapping to.
        """
        return pulumi.get(self, "distributed_virtual_switch_id")

    @distributed_virtual_switch_id.setter
    def distributed_virtual_switch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "distributed_virtual_switch_id", value)

    @property
    @pulumi.getter(name="primaryVlanId")
    def primary_vlan_id(self) -> Optional[pulumi.Input[int]]:
        """
        The primary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        """
        return pulumi.get(self, "primary_vlan_id")

    @primary_vlan_id.setter
    def primary_vlan_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "primary_vlan_id", value)

    @property
    @pulumi.getter(name="pvlanType")
    def pvlan_type(self) -> Optional[pulumi.Input[str]]:
        """
        The private VLAN type. Valid values are promiscuous, community and isolated.
        """
        return pulumi.get(self, "pvlan_type")

    @pvlan_type.setter
    def pvlan_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pvlan_type", value)

    @property
    @pulumi.getter(name="secondaryVlanId")
    def secondary_vlan_id(self) -> Optional[pulumi.Input[int]]:
        """
        The secondary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        """
        return pulumi.get(self, "secondary_vlan_id")

    @secondary_vlan_id.setter
    def secondary_vlan_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "secondary_vlan_id", value)


class DistributedVirtualSwitchPvlanMapping(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 distributed_virtual_switch_id: Optional[pulumi.Input[str]] = None,
                 primary_vlan_id: Optional[pulumi.Input[int]] = None,
                 pvlan_type: Optional[pulumi.Input[str]] = None,
                 secondary_vlan_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a DistributedVirtualSwitchPvlanMapping resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] distributed_virtual_switch_id: The ID of the distributed virtual switch to attach this mapping to.
        :param pulumi.Input[int] primary_vlan_id: The primary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        :param pulumi.Input[str] pvlan_type: The private VLAN type. Valid values are promiscuous, community and isolated.
        :param pulumi.Input[int] secondary_vlan_id: The secondary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DistributedVirtualSwitchPvlanMappingInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DistributedVirtualSwitchPvlanMapping resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DistributedVirtualSwitchPvlanMappingInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DistributedVirtualSwitchPvlanMappingInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 distributed_virtual_switch_id: Optional[pulumi.Input[str]] = None,
                 primary_vlan_id: Optional[pulumi.Input[int]] = None,
                 pvlan_type: Optional[pulumi.Input[str]] = None,
                 secondary_vlan_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DistributedVirtualSwitchPvlanMappingInitArgs.__new__(DistributedVirtualSwitchPvlanMappingInitArgs)

            if distributed_virtual_switch_id is None and not opts.urn:
                raise TypeError("Missing required property 'distributed_virtual_switch_id'")
            __props__.__dict__["distributed_virtual_switch_id"] = distributed_virtual_switch_id
            if primary_vlan_id is None and not opts.urn:
                raise TypeError("Missing required property 'primary_vlan_id'")
            __props__.__dict__["primary_vlan_id"] = primary_vlan_id
            if pvlan_type is None and not opts.urn:
                raise TypeError("Missing required property 'pvlan_type'")
            __props__.__dict__["pvlan_type"] = pvlan_type
            if secondary_vlan_id is None and not opts.urn:
                raise TypeError("Missing required property 'secondary_vlan_id'")
            __props__.__dict__["secondary_vlan_id"] = secondary_vlan_id
        super(DistributedVirtualSwitchPvlanMapping, __self__).__init__(
            'vsphere:index/distributedVirtualSwitchPvlanMapping:DistributedVirtualSwitchPvlanMapping',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            distributed_virtual_switch_id: Optional[pulumi.Input[str]] = None,
            primary_vlan_id: Optional[pulumi.Input[int]] = None,
            pvlan_type: Optional[pulumi.Input[str]] = None,
            secondary_vlan_id: Optional[pulumi.Input[int]] = None) -> 'DistributedVirtualSwitchPvlanMapping':
        """
        Get an existing DistributedVirtualSwitchPvlanMapping resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] distributed_virtual_switch_id: The ID of the distributed virtual switch to attach this mapping to.
        :param pulumi.Input[int] primary_vlan_id: The primary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        :param pulumi.Input[str] pvlan_type: The private VLAN type. Valid values are promiscuous, community and isolated.
        :param pulumi.Input[int] secondary_vlan_id: The secondary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DistributedVirtualSwitchPvlanMappingState.__new__(_DistributedVirtualSwitchPvlanMappingState)

        __props__.__dict__["distributed_virtual_switch_id"] = distributed_virtual_switch_id
        __props__.__dict__["primary_vlan_id"] = primary_vlan_id
        __props__.__dict__["pvlan_type"] = pvlan_type
        __props__.__dict__["secondary_vlan_id"] = secondary_vlan_id
        return DistributedVirtualSwitchPvlanMapping(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="distributedVirtualSwitchId")
    def distributed_virtual_switch_id(self) -> pulumi.Output[str]:
        """
        The ID of the distributed virtual switch to attach this mapping to.
        """
        return pulumi.get(self, "distributed_virtual_switch_id")

    @property
    @pulumi.getter(name="primaryVlanId")
    def primary_vlan_id(self) -> pulumi.Output[int]:
        """
        The primary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        """
        return pulumi.get(self, "primary_vlan_id")

    @property
    @pulumi.getter(name="pvlanType")
    def pvlan_type(self) -> pulumi.Output[str]:
        """
        The private VLAN type. Valid values are promiscuous, community and isolated.
        """
        return pulumi.get(self, "pvlan_type")

    @property
    @pulumi.getter(name="secondaryVlanId")
    def secondary_vlan_id(self) -> pulumi.Output[int]:
        """
        The secondary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
        """
        return pulumi.get(self, "secondary_vlan_id")


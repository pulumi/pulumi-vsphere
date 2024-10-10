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
from . import outputs
from ._inputs import *

__all__ = ['OfflineSoftwareDepotArgs', 'OfflineSoftwareDepot']

@pulumi.input_type
class OfflineSoftwareDepotArgs:
    def __init__(__self__, *,
                 location: pulumi.Input[str]):
        """
        The set of arguments for constructing a OfflineSoftwareDepot resource.
        :param pulumi.Input[str] location: The URL where the depot source is hosted.
        """
        pulumi.set(__self__, "location", location)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        The URL where the depot source is hosted.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)


@pulumi.input_type
class _OfflineSoftwareDepotState:
    def __init__(__self__, *,
                 components: Optional[pulumi.Input[Sequence[pulumi.Input['OfflineSoftwareDepotComponentArgs']]]] = None,
                 location: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OfflineSoftwareDepot resources.
        :param pulumi.Input[Sequence[pulumi.Input['OfflineSoftwareDepotComponentArgs']]] components: The list of custom components in the depot.
        :param pulumi.Input[str] location: The URL where the depot source is hosted.
        """
        if components is not None:
            pulumi.set(__self__, "components", components)
        if location is not None:
            pulumi.set(__self__, "location", location)

    @property
    @pulumi.getter
    def components(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OfflineSoftwareDepotComponentArgs']]]]:
        """
        The list of custom components in the depot.
        """
        return pulumi.get(self, "components")

    @components.setter
    def components(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OfflineSoftwareDepotComponentArgs']]]]):
        pulumi.set(self, "components", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The URL where the depot source is hosted.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)


class OfflineSoftwareDepot(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a VMware vSphere offline software depot resource.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The URL where the depot source is hosted.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OfflineSoftwareDepotArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a VMware vSphere offline software depot resource.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param OfflineSoftwareDepotArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OfflineSoftwareDepotArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OfflineSoftwareDepotArgs.__new__(OfflineSoftwareDepotArgs)

            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["components"] = None
        super(OfflineSoftwareDepot, __self__).__init__(
            'vsphere:index/offlineSoftwareDepot:OfflineSoftwareDepot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            components: Optional[pulumi.Input[Sequence[pulumi.Input[Union['OfflineSoftwareDepotComponentArgs', 'OfflineSoftwareDepotComponentArgsDict']]]]] = None,
            location: Optional[pulumi.Input[str]] = None) -> 'OfflineSoftwareDepot':
        """
        Get an existing OfflineSoftwareDepot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['OfflineSoftwareDepotComponentArgs', 'OfflineSoftwareDepotComponentArgsDict']]]] components: The list of custom components in the depot.
        :param pulumi.Input[str] location: The URL where the depot source is hosted.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OfflineSoftwareDepotState.__new__(_OfflineSoftwareDepotState)

        __props__.__dict__["components"] = components
        __props__.__dict__["location"] = location
        return OfflineSoftwareDepot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def components(self) -> pulumi.Output[Sequence['outputs.OfflineSoftwareDepotComponent']]:
        """
        The list of custom components in the depot.
        """
        return pulumi.get(self, "components")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The URL where the depot source is hosted.
        """
        return pulumi.get(self, "location")


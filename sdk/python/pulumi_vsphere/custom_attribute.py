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

__all__ = ['CustomAttributeArgs', 'CustomAttribute']

@pulumi.input_type
class CustomAttributeArgs:
    def __init__(__self__, *,
                 managed_object_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CustomAttribute resource.
        :param pulumi.Input[str] managed_object_type: The object type that this attribute may be
               applied to. If not set, the custom attribute may be applied to any object
               type. For a full list, review the Managed Object Types. Forces a new resource if changed.
        :param pulumi.Input[str] name: The name of the custom attribute.
        """
        if managed_object_type is not None:
            pulumi.set(__self__, "managed_object_type", managed_object_type)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="managedObjectType")
    def managed_object_type(self) -> Optional[pulumi.Input[str]]:
        """
        The object type that this attribute may be
        applied to. If not set, the custom attribute may be applied to any object
        type. For a full list, review the Managed Object Types. Forces a new resource if changed.
        """
        return pulumi.get(self, "managed_object_type")

    @managed_object_type.setter
    def managed_object_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "managed_object_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the custom attribute.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _CustomAttributeState:
    def __init__(__self__, *,
                 managed_object_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CustomAttribute resources.
        :param pulumi.Input[str] managed_object_type: The object type that this attribute may be
               applied to. If not set, the custom attribute may be applied to any object
               type. For a full list, review the Managed Object Types. Forces a new resource if changed.
        :param pulumi.Input[str] name: The name of the custom attribute.
        """
        if managed_object_type is not None:
            pulumi.set(__self__, "managed_object_type", managed_object_type)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="managedObjectType")
    def managed_object_type(self) -> Optional[pulumi.Input[str]]:
        """
        The object type that this attribute may be
        applied to. If not set, the custom attribute may be applied to any object
        type. For a full list, review the Managed Object Types. Forces a new resource if changed.
        """
        return pulumi.get(self, "managed_object_type")

    @managed_object_type.setter
    def managed_object_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "managed_object_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the custom attribute.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class CustomAttribute(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 managed_object_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `CustomAttribute` resource can be used to create and manage custom
        attributes, which allow users to associate user-specific meta-information with
        vSphere managed objects. Custom attribute values must be strings and are stored
        on the vCenter Server and not the managed object.

        For more information about custom attributes, click [here][ext-custom-attributes].

        [ext-custom-attributes]: https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vcenter-esxi-management/GUID-73606C4C-763C-4E27-A1DA-032E4C46219D.html

        > **NOTE:** Custom attributes are unsupported on direct ESXi host connections
        and require vCenter Server.

        ## Example Usage

        This example creates a custom attribute named `test-attribute`. The
        resulting custom attribute can be assigned to VMs only.

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        attribute = vsphere.CustomAttribute("attribute",
            name="test-attribute",
            managed_object_type="VirtualMachine")
        ```

        ## Import

        An existing custom attribute can be imported into this resource

        via its name, using the following command:

        ```sh
        $ pulumi import vsphere:index/customAttribute:CustomAttribute attribute terraform-test-attribute
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] managed_object_type: The object type that this attribute may be
               applied to. If not set, the custom attribute may be applied to any object
               type. For a full list, review the Managed Object Types. Forces a new resource if changed.
        :param pulumi.Input[str] name: The name of the custom attribute.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[CustomAttributeArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `CustomAttribute` resource can be used to create and manage custom
        attributes, which allow users to associate user-specific meta-information with
        vSphere managed objects. Custom attribute values must be strings and are stored
        on the vCenter Server and not the managed object.

        For more information about custom attributes, click [here][ext-custom-attributes].

        [ext-custom-attributes]: https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vcenter-esxi-management/GUID-73606C4C-763C-4E27-A1DA-032E4C46219D.html

        > **NOTE:** Custom attributes are unsupported on direct ESXi host connections
        and require vCenter Server.

        ## Example Usage

        This example creates a custom attribute named `test-attribute`. The
        resulting custom attribute can be assigned to VMs only.

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        attribute = vsphere.CustomAttribute("attribute",
            name="test-attribute",
            managed_object_type="VirtualMachine")
        ```

        ## Import

        An existing custom attribute can be imported into this resource

        via its name, using the following command:

        ```sh
        $ pulumi import vsphere:index/customAttribute:CustomAttribute attribute terraform-test-attribute
        ```

        :param str resource_name: The name of the resource.
        :param CustomAttributeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomAttributeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 managed_object_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CustomAttributeArgs.__new__(CustomAttributeArgs)

            __props__.__dict__["managed_object_type"] = managed_object_type
            __props__.__dict__["name"] = name
        super(CustomAttribute, __self__).__init__(
            'vsphere:index/customAttribute:CustomAttribute',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            managed_object_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'CustomAttribute':
        """
        Get an existing CustomAttribute resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] managed_object_type: The object type that this attribute may be
               applied to. If not set, the custom attribute may be applied to any object
               type. For a full list, review the Managed Object Types. Forces a new resource if changed.
        :param pulumi.Input[str] name: The name of the custom attribute.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CustomAttributeState.__new__(_CustomAttributeState)

        __props__.__dict__["managed_object_type"] = managed_object_type
        __props__.__dict__["name"] = name
        return CustomAttribute(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="managedObjectType")
    def managed_object_type(self) -> pulumi.Output[Optional[str]]:
        """
        The object type that this attribute may be
        applied to. If not set, the custom attribute may be applied to any object
        type. For a full list, review the Managed Object Types. Forces a new resource if changed.
        """
        return pulumi.get(self, "managed_object_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the custom attribute.
        """
        return pulumi.get(self, "name")


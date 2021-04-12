# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables

__all__ = ['FolderArgs', 'Folder']

@pulumi.input_type
class FolderArgs:
    def __init__(__self__, *,
                 path: pulumi.Input[str],
                 type: pulumi.Input[str],
                 custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 datacenter_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Folder resource.
        :param pulumi.Input[str] path: The path of the folder to be created. This is relative to
               the root of the type of folder you are creating, and the supplied datacenter.
               For example, given a default datacenter of `default-dc`, a folder of type
               `vm` (denoting a virtual machine folder), and a supplied folder of
               `test-folder`, the resulting path would be
               `/default-dc/vm/test-folder`.
        :param pulumi.Input[str] type: The type of folder to create. Allowed options are
               `datacenter` for datacenter folders, `host` for host and cluster folders,
               `vm` for virtual machine folders, `datastore` for datastore folders, and
               `network` for network folders. Forces a new resource if changed.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: Map of custom attribute ids to attribute 
               value strings to set for folder. See [here][docs-setting-custom-attributes]
               for a reference on how to set values for custom attributes.
        :param pulumi.Input[str] datacenter_id: The ID of the datacenter the folder will be created in.
               Required for all folder types except for datacenter folders. Forces a new
               resource if changed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The IDs of any tags to attach to this resource.
        """
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "type", type)
        if custom_attributes is not None:
            pulumi.set(__self__, "custom_attributes", custom_attributes)
        if datacenter_id is not None:
            pulumi.set(__self__, "datacenter_id", datacenter_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        The path of the folder to be created. This is relative to
        the root of the type of folder you are creating, and the supplied datacenter.
        For example, given a default datacenter of `default-dc`, a folder of type
        `vm` (denoting a virtual machine folder), and a supplied folder of
        `test-folder`, the resulting path would be
        `/default-dc/vm/test-folder`.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of folder to create. Allowed options are
        `datacenter` for datacenter folders, `host` for host and cluster folders,
        `vm` for virtual machine folders, `datastore` for datastore folders, and
        `network` for network folders. Forces a new resource if changed.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="customAttributes")
    def custom_attributes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of custom attribute ids to attribute 
        value strings to set for folder. See [here][docs-setting-custom-attributes]
        for a reference on how to set values for custom attributes.
        """
        return pulumi.get(self, "custom_attributes")

    @custom_attributes.setter
    def custom_attributes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "custom_attributes", value)

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the datacenter the folder will be created in.
        Required for all folder types except for datacenter folders. Forces a new
        resource if changed.
        """
        return pulumi.get(self, "datacenter_id")

    @datacenter_id.setter
    def datacenter_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "datacenter_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The IDs of any tags to attach to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class Folder(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 datacenter_id: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Folder resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: Map of custom attribute ids to attribute 
               value strings to set for folder. See [here][docs-setting-custom-attributes]
               for a reference on how to set values for custom attributes.
        :param pulumi.Input[str] datacenter_id: The ID of the datacenter the folder will be created in.
               Required for all folder types except for datacenter folders. Forces a new
               resource if changed.
        :param pulumi.Input[str] path: The path of the folder to be created. This is relative to
               the root of the type of folder you are creating, and the supplied datacenter.
               For example, given a default datacenter of `default-dc`, a folder of type
               `vm` (denoting a virtual machine folder), and a supplied folder of
               `test-folder`, the resulting path would be
               `/default-dc/vm/test-folder`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The IDs of any tags to attach to this resource.
        :param pulumi.Input[str] type: The type of folder to create. Allowed options are
               `datacenter` for datacenter folders, `host` for host and cluster folders,
               `vm` for virtual machine folders, `datastore` for datastore folders, and
               `network` for network folders. Forces a new resource if changed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FolderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Folder resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FolderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FolderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 datacenter_id: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            __props__['custom_attributes'] = custom_attributes
            __props__['datacenter_id'] = datacenter_id
            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__['path'] = path
            __props__['tags'] = tags
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
        super(Folder, __self__).__init__(
            'vsphere:index/folder:Folder',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            datacenter_id: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Folder':
        """
        Get an existing Folder resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: Map of custom attribute ids to attribute 
               value strings to set for folder. See [here][docs-setting-custom-attributes]
               for a reference on how to set values for custom attributes.
        :param pulumi.Input[str] datacenter_id: The ID of the datacenter the folder will be created in.
               Required for all folder types except for datacenter folders. Forces a new
               resource if changed.
        :param pulumi.Input[str] path: The path of the folder to be created. This is relative to
               the root of the type of folder you are creating, and the supplied datacenter.
               For example, given a default datacenter of `default-dc`, a folder of type
               `vm` (denoting a virtual machine folder), and a supplied folder of
               `test-folder`, the resulting path would be
               `/default-dc/vm/test-folder`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The IDs of any tags to attach to this resource.
        :param pulumi.Input[str] type: The type of folder to create. Allowed options are
               `datacenter` for datacenter folders, `host` for host and cluster folders,
               `vm` for virtual machine folders, `datastore` for datastore folders, and
               `network` for network folders. Forces a new resource if changed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["custom_attributes"] = custom_attributes
        __props__["datacenter_id"] = datacenter_id
        __props__["path"] = path
        __props__["tags"] = tags
        __props__["type"] = type
        return Folder(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="customAttributes")
    def custom_attributes(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map of custom attribute ids to attribute 
        value strings to set for folder. See [here][docs-setting-custom-attributes]
        for a reference on how to set values for custom attributes.
        """
        return pulumi.get(self, "custom_attributes")

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the datacenter the folder will be created in.
        Required for all folder types except for datacenter folders. Forces a new
        resource if changed.
        """
        return pulumi.get(self, "datacenter_id")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The path of the folder to be created. This is relative to
        the root of the type of folder you are creating, and the supplied datacenter.
        For example, given a default datacenter of `default-dc`, a folder of type
        `vm` (denoting a virtual machine folder), and a supplied folder of
        `test-folder`, the resulting path would be
        `/default-dc/vm/test-folder`.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The IDs of any tags to attach to this resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of folder to create. Allowed options are
        `datacenter` for datacenter folders, `host` for host and cluster folders,
        `vm` for virtual machine folders, `datastore` for datastore folders, and
        `network` for network folders. Forces a new resource if changed.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


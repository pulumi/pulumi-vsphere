# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['EntityPermissions']


class EntityPermissions(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 entity_type: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EntityPermissionsPermissionArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a EntityPermissions resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] entity_id: The managed object id (uuid for some entities) on which permissions are to be created.
        :param pulumi.Input[str] entity_type: The managed object type, types can be found in the managed object type section 
               [here](https://code.vmware.com/apis/968/vsphere).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EntityPermissionsPermissionArgs']]]] permissions: The permissions to be given on this entity. Keep the permissions sorted
               alphabetically on `user_or_group` for a better user experience.
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

            if entity_id is None:
                raise TypeError("Missing required property 'entity_id'")
            __props__['entity_id'] = entity_id
            if entity_type is None:
                raise TypeError("Missing required property 'entity_type'")
            __props__['entity_type'] = entity_type
            if permissions is None:
                raise TypeError("Missing required property 'permissions'")
            __props__['permissions'] = permissions
        super(EntityPermissions, __self__).__init__(
            'vsphere:index/entityPermissions:EntityPermissions',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            entity_id: Optional[pulumi.Input[str]] = None,
            entity_type: Optional[pulumi.Input[str]] = None,
            permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EntityPermissionsPermissionArgs']]]]] = None) -> 'EntityPermissions':
        """
        Get an existing EntityPermissions resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] entity_id: The managed object id (uuid for some entities) on which permissions are to be created.
        :param pulumi.Input[str] entity_type: The managed object type, types can be found in the managed object type section 
               [here](https://code.vmware.com/apis/968/vsphere).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EntityPermissionsPermissionArgs']]]] permissions: The permissions to be given on this entity. Keep the permissions sorted
               alphabetically on `user_or_group` for a better user experience.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["entity_id"] = entity_id
        __props__["entity_type"] = entity_type
        __props__["permissions"] = permissions
        return EntityPermissions(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> pulumi.Output[str]:
        """
        The managed object id (uuid for some entities) on which permissions are to be created.
        """
        return pulumi.get(self, "entity_id")

    @property
    @pulumi.getter(name="entityType")
    def entity_type(self) -> pulumi.Output[str]:
        """
        The managed object type, types can be found in the managed object type section 
        [here](https://code.vmware.com/apis/968/vsphere).
        """
        return pulumi.get(self, "entity_type")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Sequence['outputs.EntityPermissionsPermission']]:
        """
        The permissions to be given on this entity. Keep the permissions sorted
        alphabetically on `user_or_group` for a better user experience.
        """
        return pulumi.get(self, "permissions")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


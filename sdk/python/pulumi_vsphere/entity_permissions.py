# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['EntityPermissionsArgs', 'EntityPermissions']

@pulumi.input_type
class EntityPermissionsArgs:
    def __init__(__self__, *,
                 entity_id: pulumi.Input[str],
                 entity_type: pulumi.Input[str],
                 permissions: pulumi.Input[Sequence[pulumi.Input['EntityPermissionsPermissionArgs']]]):
        """
        The set of arguments for constructing a EntityPermissions resource.
        :param pulumi.Input[str] entity_id: The managed object id (uuid for some entities) on which permissions are to be created.
        :param pulumi.Input[str] entity_type: The managed object type, types can be found in the managed object type section 
               [here](https://developer.vmware.com/apis/968/vsphere).
        :param pulumi.Input[Sequence[pulumi.Input['EntityPermissionsPermissionArgs']]] permissions: The permissions to be given on this entity. Keep the permissions sorted
               alphabetically on `user_or_group` for a better user experience.
        """
        EntityPermissionsArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            entity_id=entity_id,
            entity_type=entity_type,
            permissions=permissions,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             entity_id: pulumi.Input[str],
             entity_type: pulumi.Input[str],
             permissions: pulumi.Input[Sequence[pulumi.Input['EntityPermissionsPermissionArgs']]],
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("entity_id", entity_id)
        _setter("entity_type", entity_type)
        _setter("permissions", permissions)

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> pulumi.Input[str]:
        """
        The managed object id (uuid for some entities) on which permissions are to be created.
        """
        return pulumi.get(self, "entity_id")

    @entity_id.setter
    def entity_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "entity_id", value)

    @property
    @pulumi.getter(name="entityType")
    def entity_type(self) -> pulumi.Input[str]:
        """
        The managed object type, types can be found in the managed object type section 
        [here](https://developer.vmware.com/apis/968/vsphere).
        """
        return pulumi.get(self, "entity_type")

    @entity_type.setter
    def entity_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "entity_type", value)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Input[Sequence[pulumi.Input['EntityPermissionsPermissionArgs']]]:
        """
        The permissions to be given on this entity. Keep the permissions sorted
        alphabetically on `user_or_group` for a better user experience.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: pulumi.Input[Sequence[pulumi.Input['EntityPermissionsPermissionArgs']]]):
        pulumi.set(self, "permissions", value)


@pulumi.input_type
class _EntityPermissionsState:
    def __init__(__self__, *,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 entity_type: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['EntityPermissionsPermissionArgs']]]] = None):
        """
        Input properties used for looking up and filtering EntityPermissions resources.
        :param pulumi.Input[str] entity_id: The managed object id (uuid for some entities) on which permissions are to be created.
        :param pulumi.Input[str] entity_type: The managed object type, types can be found in the managed object type section 
               [here](https://developer.vmware.com/apis/968/vsphere).
        :param pulumi.Input[Sequence[pulumi.Input['EntityPermissionsPermissionArgs']]] permissions: The permissions to be given on this entity. Keep the permissions sorted
               alphabetically on `user_or_group` for a better user experience.
        """
        _EntityPermissionsState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            entity_id=entity_id,
            entity_type=entity_type,
            permissions=permissions,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             entity_id: Optional[pulumi.Input[str]] = None,
             entity_type: Optional[pulumi.Input[str]] = None,
             permissions: Optional[pulumi.Input[Sequence[pulumi.Input['EntityPermissionsPermissionArgs']]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if entity_id is not None:
            _setter("entity_id", entity_id)
        if entity_type is not None:
            _setter("entity_type", entity_type)
        if permissions is not None:
            _setter("permissions", permissions)

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed object id (uuid for some entities) on which permissions are to be created.
        """
        return pulumi.get(self, "entity_id")

    @entity_id.setter
    def entity_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_id", value)

    @property
    @pulumi.getter(name="entityType")
    def entity_type(self) -> Optional[pulumi.Input[str]]:
        """
        The managed object type, types can be found in the managed object type section 
        [here](https://developer.vmware.com/apis/968/vsphere).
        """
        return pulumi.get(self, "entity_type")

    @entity_type.setter
    def entity_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_type", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EntityPermissionsPermissionArgs']]]]:
        """
        The permissions to be given on this entity. Keep the permissions sorted
        alphabetically on `user_or_group` for a better user experience.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EntityPermissionsPermissionArgs']]]]):
        pulumi.set(self, "permissions", value)


class EntityPermissions(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 entity_type: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EntityPermissionsPermissionArgs']]]]] = None,
                 __props__=None):
        """
        Create a EntityPermissions resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] entity_id: The managed object id (uuid for some entities) on which permissions are to be created.
        :param pulumi.Input[str] entity_type: The managed object type, types can be found in the managed object type section 
               [here](https://developer.vmware.com/apis/968/vsphere).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EntityPermissionsPermissionArgs']]]] permissions: The permissions to be given on this entity. Keep the permissions sorted
               alphabetically on `user_or_group` for a better user experience.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EntityPermissionsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a EntityPermissions resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param EntityPermissionsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EntityPermissionsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            EntityPermissionsArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 entity_type: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EntityPermissionsPermissionArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EntityPermissionsArgs.__new__(EntityPermissionsArgs)

            if entity_id is None and not opts.urn:
                raise TypeError("Missing required property 'entity_id'")
            __props__.__dict__["entity_id"] = entity_id
            if entity_type is None and not opts.urn:
                raise TypeError("Missing required property 'entity_type'")
            __props__.__dict__["entity_type"] = entity_type
            if permissions is None and not opts.urn:
                raise TypeError("Missing required property 'permissions'")
            __props__.__dict__["permissions"] = permissions
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
               [here](https://developer.vmware.com/apis/968/vsphere).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EntityPermissionsPermissionArgs']]]] permissions: The permissions to be given on this entity. Keep the permissions sorted
               alphabetically on `user_or_group` for a better user experience.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EntityPermissionsState.__new__(_EntityPermissionsState)

        __props__.__dict__["entity_id"] = entity_id
        __props__.__dict__["entity_type"] = entity_type
        __props__.__dict__["permissions"] = permissions
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
        [here](https://developer.vmware.com/apis/968/vsphere).
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


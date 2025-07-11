# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = ['RoleArgs', 'Role']

@pulumi.input_type
class RoleArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 role_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a Role resource.
        :param pulumi.Input[builtins.str] name: The name of the role.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] role_privileges: The privileges to be associated with this role.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if role_privileges is not None:
            pulumi.set(__self__, "role_privileges", role_privileges)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the role.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rolePrivileges")
    def role_privileges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The privileges to be associated with this role.
        """
        return pulumi.get(self, "role_privileges")

    @role_privileges.setter
    def role_privileges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "role_privileges", value)


@pulumi.input_type
class _RoleState:
    def __init__(__self__, *,
                 label: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 role_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering Role resources.
        :param pulumi.Input[builtins.str] label: The display label of the role.
        :param pulumi.Input[builtins.str] name: The name of the role.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] role_privileges: The privileges to be associated with this role.
        """
        if label is not None:
            pulumi.set(__self__, "label", label)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if role_privileges is not None:
            pulumi.set(__self__, "role_privileges", role_privileges)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The display label of the role.
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the role.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rolePrivileges")
    def role_privileges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The privileges to be associated with this role.
        """
        return pulumi.get(self, "role_privileges")

    @role_privileges.setter
    def role_privileges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "role_privileges", value)


@pulumi.type_token("vsphere:index/role:Role")
class Role(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 role_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        ## Import

        An existing role can be imported into this resource by supplying the role id. An example is below:

        ```sh
        $ pulumi import vsphere:index/role:Role role1 -709298051
        ```

        Use [`vsphere_role` data source][ref-vsphere-role-data-source]

        to read information about system roles.

        [ref-vsphere-role-data-source]: /docs/providers/vsphere/d/vsphere_role.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] name: The name of the role.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] role_privileges: The privileges to be associated with this role.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RoleArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        An existing role can be imported into this resource by supplying the role id. An example is below:

        ```sh
        $ pulumi import vsphere:index/role:Role role1 -709298051
        ```

        Use [`vsphere_role` data source][ref-vsphere-role-data-source]

        to read information about system roles.

        [ref-vsphere-role-data-source]: /docs/providers/vsphere/d/vsphere_role.html

        :param str resource_name: The name of the resource.
        :param RoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 role_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RoleArgs.__new__(RoleArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["role_privileges"] = role_privileges
            __props__.__dict__["label"] = None
        super(Role, __self__).__init__(
            'vsphere:index/role:Role',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            label: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            role_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None) -> 'Role':
        """
        Get an existing Role resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] label: The display label of the role.
        :param pulumi.Input[builtins.str] name: The name of the role.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] role_privileges: The privileges to be associated with this role.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RoleState.__new__(_RoleState)

        __props__.__dict__["label"] = label
        __props__.__dict__["name"] = name
        __props__.__dict__["role_privileges"] = role_privileges
        return Role(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def label(self) -> pulumi.Output[builtins.str]:
        """
        The display label of the role.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the role.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="rolePrivileges")
    def role_privileges(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        The privileges to be associated with this role.
        """
        return pulumi.get(self, "role_privileges")


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 password: pulumi.Input[str],
                 user: pulumi.Input[str],
                 allow_unverified_ssl: Optional[pulumi.Input[bool]] = None,
                 api_timeout: Optional[pulumi.Input[int]] = None,
                 client_debug: Optional[pulumi.Input[bool]] = None,
                 client_debug_path: Optional[pulumi.Input[str]] = None,
                 client_debug_path_run: Optional[pulumi.Input[str]] = None,
                 persist_session: Optional[pulumi.Input[bool]] = None,
                 rest_session_path: Optional[pulumi.Input[str]] = None,
                 vcenter_server: Optional[pulumi.Input[str]] = None,
                 vim_keep_alive: Optional[pulumi.Input[int]] = None,
                 vim_session_path: Optional[pulumi.Input[str]] = None,
                 vsphere_server: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] password: The user password for vSphere API operations.
        :param pulumi.Input[str] user: The user name for vSphere API operations.
        :param pulumi.Input[bool] allow_unverified_ssl: If set, VMware vSphere client will permit unverifiable SSL certificates.
        :param pulumi.Input[int] api_timeout: API timeout in minutes (Default: 5)
        :param pulumi.Input[bool] client_debug: govmomi debug
        :param pulumi.Input[str] client_debug_path: govmomi debug path for debug
        :param pulumi.Input[str] client_debug_path_run: govmomi debug path for a single run
        :param pulumi.Input[bool] persist_session: Persist vSphere client sessions to disk
        :param pulumi.Input[str] rest_session_path: The directory to save vSphere REST API sessions to
        :param pulumi.Input[int] vim_keep_alive: Keep alive interval for the VIM session in minutes
        :param pulumi.Input[str] vim_session_path: The directory to save vSphere SOAP API sessions to
        :param pulumi.Input[str] vsphere_server: The vSphere Server name for vSphere API operations.
        """
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "user", user)
        if allow_unverified_ssl is None:
            allow_unverified_ssl = _utilities.get_env_bool('VSPHERE_ALLOW_UNVERIFIED_SSL')
        if allow_unverified_ssl is not None:
            pulumi.set(__self__, "allow_unverified_ssl", allow_unverified_ssl)
        if api_timeout is not None:
            pulumi.set(__self__, "api_timeout", api_timeout)
        if client_debug is None:
            client_debug = _utilities.get_env_bool('VSPHERE_CLIENT_DEBUG')
        if client_debug is not None:
            pulumi.set(__self__, "client_debug", client_debug)
        if client_debug_path is None:
            client_debug_path = _utilities.get_env('VSPHERE_CLIENT_DEBUG_PATH')
        if client_debug_path is not None:
            pulumi.set(__self__, "client_debug_path", client_debug_path)
        if client_debug_path_run is None:
            client_debug_path_run = _utilities.get_env('VSPHERE_CLIENT_DEBUG_PATH_RUN')
        if client_debug_path_run is not None:
            pulumi.set(__self__, "client_debug_path_run", client_debug_path_run)
        if persist_session is None:
            persist_session = _utilities.get_env_bool('VSPHERE_PERSIST_SESSION')
        if persist_session is not None:
            pulumi.set(__self__, "persist_session", persist_session)
        if rest_session_path is None:
            rest_session_path = _utilities.get_env('VSPHERE_REST_SESSION_PATH')
        if rest_session_path is not None:
            pulumi.set(__self__, "rest_session_path", rest_session_path)
        if vcenter_server is not None:
            warnings.warn("""This field has been renamed to vsphere_server.""", DeprecationWarning)
            pulumi.log.warn("""vcenter_server is deprecated: This field has been renamed to vsphere_server.""")
        if vcenter_server is not None:
            pulumi.set(__self__, "vcenter_server", vcenter_server)
        if vim_keep_alive is None:
            vim_keep_alive = _utilities.get_env_int('VSPHERE_VIM_KEEP_ALIVE')
        if vim_keep_alive is not None:
            pulumi.set(__self__, "vim_keep_alive", vim_keep_alive)
        if vim_session_path is None:
            vim_session_path = _utilities.get_env('VSPHERE_VIM_SESSION_PATH')
        if vim_session_path is not None:
            pulumi.set(__self__, "vim_session_path", vim_session_path)
        if vsphere_server is not None:
            pulumi.set(__self__, "vsphere_server", vsphere_server)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        The user password for vSphere API operations.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def user(self) -> pulumi.Input[str]:
        """
        The user name for vSphere API operations.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: pulumi.Input[str]):
        pulumi.set(self, "user", value)

    @property
    @pulumi.getter(name="allowUnverifiedSsl")
    def allow_unverified_ssl(self) -> Optional[pulumi.Input[bool]]:
        """
        If set, VMware vSphere client will permit unverifiable SSL certificates.
        """
        return pulumi.get(self, "allow_unverified_ssl")

    @allow_unverified_ssl.setter
    def allow_unverified_ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_unverified_ssl", value)

    @property
    @pulumi.getter(name="apiTimeout")
    def api_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        API timeout in minutes (Default: 5)
        """
        return pulumi.get(self, "api_timeout")

    @api_timeout.setter
    def api_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "api_timeout", value)

    @property
    @pulumi.getter(name="clientDebug")
    def client_debug(self) -> Optional[pulumi.Input[bool]]:
        """
        govmomi debug
        """
        return pulumi.get(self, "client_debug")

    @client_debug.setter
    def client_debug(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "client_debug", value)

    @property
    @pulumi.getter(name="clientDebugPath")
    def client_debug_path(self) -> Optional[pulumi.Input[str]]:
        """
        govmomi debug path for debug
        """
        return pulumi.get(self, "client_debug_path")

    @client_debug_path.setter
    def client_debug_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_debug_path", value)

    @property
    @pulumi.getter(name="clientDebugPathRun")
    def client_debug_path_run(self) -> Optional[pulumi.Input[str]]:
        """
        govmomi debug path for a single run
        """
        return pulumi.get(self, "client_debug_path_run")

    @client_debug_path_run.setter
    def client_debug_path_run(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_debug_path_run", value)

    @property
    @pulumi.getter(name="persistSession")
    def persist_session(self) -> Optional[pulumi.Input[bool]]:
        """
        Persist vSphere client sessions to disk
        """
        return pulumi.get(self, "persist_session")

    @persist_session.setter
    def persist_session(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "persist_session", value)

    @property
    @pulumi.getter(name="restSessionPath")
    def rest_session_path(self) -> Optional[pulumi.Input[str]]:
        """
        The directory to save vSphere REST API sessions to
        """
        return pulumi.get(self, "rest_session_path")

    @rest_session_path.setter
    def rest_session_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rest_session_path", value)

    @property
    @pulumi.getter(name="vcenterServer")
    def vcenter_server(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vcenter_server")

    @vcenter_server.setter
    def vcenter_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vcenter_server", value)

    @property
    @pulumi.getter(name="vimKeepAlive")
    def vim_keep_alive(self) -> Optional[pulumi.Input[int]]:
        """
        Keep alive interval for the VIM session in minutes
        """
        return pulumi.get(self, "vim_keep_alive")

    @vim_keep_alive.setter
    def vim_keep_alive(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vim_keep_alive", value)

    @property
    @pulumi.getter(name="vimSessionPath")
    def vim_session_path(self) -> Optional[pulumi.Input[str]]:
        """
        The directory to save vSphere SOAP API sessions to
        """
        return pulumi.get(self, "vim_session_path")

    @vim_session_path.setter
    def vim_session_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vim_session_path", value)

    @property
    @pulumi.getter(name="vsphereServer")
    def vsphere_server(self) -> Optional[pulumi.Input[str]]:
        """
        The vSphere Server name for vSphere API operations.
        """
        return pulumi.get(self, "vsphere_server")

    @vsphere_server.setter
    def vsphere_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vsphere_server", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_unverified_ssl: Optional[pulumi.Input[bool]] = None,
                 api_timeout: Optional[pulumi.Input[int]] = None,
                 client_debug: Optional[pulumi.Input[bool]] = None,
                 client_debug_path: Optional[pulumi.Input[str]] = None,
                 client_debug_path_run: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 persist_session: Optional[pulumi.Input[bool]] = None,
                 rest_session_path: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 vcenter_server: Optional[pulumi.Input[str]] = None,
                 vim_keep_alive: Optional[pulumi.Input[int]] = None,
                 vim_session_path: Optional[pulumi.Input[str]] = None,
                 vsphere_server: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The provider type for the vsphere package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_unverified_ssl: If set, VMware vSphere client will permit unverifiable SSL certificates.
        :param pulumi.Input[int] api_timeout: API timeout in minutes (Default: 5)
        :param pulumi.Input[bool] client_debug: govmomi debug
        :param pulumi.Input[str] client_debug_path: govmomi debug path for debug
        :param pulumi.Input[str] client_debug_path_run: govmomi debug path for a single run
        :param pulumi.Input[str] password: The user password for vSphere API operations.
        :param pulumi.Input[bool] persist_session: Persist vSphere client sessions to disk
        :param pulumi.Input[str] rest_session_path: The directory to save vSphere REST API sessions to
        :param pulumi.Input[str] user: The user name for vSphere API operations.
        :param pulumi.Input[int] vim_keep_alive: Keep alive interval for the VIM session in minutes
        :param pulumi.Input[str] vim_session_path: The directory to save vSphere SOAP API sessions to
        :param pulumi.Input[str] vsphere_server: The vSphere Server name for vSphere API operations.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the vsphere package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_unverified_ssl: Optional[pulumi.Input[bool]] = None,
                 api_timeout: Optional[pulumi.Input[int]] = None,
                 client_debug: Optional[pulumi.Input[bool]] = None,
                 client_debug_path: Optional[pulumi.Input[str]] = None,
                 client_debug_path_run: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 persist_session: Optional[pulumi.Input[bool]] = None,
                 rest_session_path: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 vcenter_server: Optional[pulumi.Input[str]] = None,
                 vim_keep_alive: Optional[pulumi.Input[int]] = None,
                 vim_session_path: Optional[pulumi.Input[str]] = None,
                 vsphere_server: Optional[pulumi.Input[str]] = None,
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

            if allow_unverified_ssl is None:
                allow_unverified_ssl = _utilities.get_env_bool('VSPHERE_ALLOW_UNVERIFIED_SSL')
            __props__['allow_unverified_ssl'] = pulumi.Output.from_input(allow_unverified_ssl).apply(pulumi.runtime.to_json) if allow_unverified_ssl is not None else None
            __props__['api_timeout'] = pulumi.Output.from_input(api_timeout).apply(pulumi.runtime.to_json) if api_timeout is not None else None
            if client_debug is None:
                client_debug = _utilities.get_env_bool('VSPHERE_CLIENT_DEBUG')
            __props__['client_debug'] = pulumi.Output.from_input(client_debug).apply(pulumi.runtime.to_json) if client_debug is not None else None
            if client_debug_path is None:
                client_debug_path = _utilities.get_env('VSPHERE_CLIENT_DEBUG_PATH')
            __props__['client_debug_path'] = client_debug_path
            if client_debug_path_run is None:
                client_debug_path_run = _utilities.get_env('VSPHERE_CLIENT_DEBUG_PATH_RUN')
            __props__['client_debug_path_run'] = client_debug_path_run
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__['password'] = password
            if persist_session is None:
                persist_session = _utilities.get_env_bool('VSPHERE_PERSIST_SESSION')
            __props__['persist_session'] = pulumi.Output.from_input(persist_session).apply(pulumi.runtime.to_json) if persist_session is not None else None
            if rest_session_path is None:
                rest_session_path = _utilities.get_env('VSPHERE_REST_SESSION_PATH')
            __props__['rest_session_path'] = rest_session_path
            if user is None and not opts.urn:
                raise TypeError("Missing required property 'user'")
            __props__['user'] = user
            if vcenter_server is not None and not opts.urn:
                warnings.warn("""This field has been renamed to vsphere_server.""", DeprecationWarning)
                pulumi.log.warn("""vcenter_server is deprecated: This field has been renamed to vsphere_server.""")
            __props__['vcenter_server'] = vcenter_server
            if vim_keep_alive is None:
                vim_keep_alive = _utilities.get_env_int('VSPHERE_VIM_KEEP_ALIVE')
            __props__['vim_keep_alive'] = pulumi.Output.from_input(vim_keep_alive).apply(pulumi.runtime.to_json) if vim_keep_alive is not None else None
            if vim_session_path is None:
                vim_session_path = _utilities.get_env('VSPHERE_VIM_SESSION_PATH')
            __props__['vim_session_path'] = vim_session_path
            __props__['vsphere_server'] = vsphere_server
        super(Provider, __self__).__init__(
            'vsphere',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


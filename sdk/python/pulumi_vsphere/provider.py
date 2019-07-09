# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class Provider(pulumi.ProviderResource):
    def __init__(__self__, resource_name, opts=None, allow_unverified_ssl=None, client_debug=None, client_debug_path=None, client_debug_path_run=None, password=None, persist_session=None, rest_session_path=None, user=None, vcenter_server=None, vim_keep_alive=None, vim_session_path=None, vsphere_server=None, __name__=None, __opts__=None):
        """
        The provider type for the vsphere package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://pulumi.io/reference/programming-model.html#providers) for more information.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/index.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['allow_unverified_ssl'] = pulumi.Output.from_input(allow_unverified_ssl).apply(json.dumps) if allow_unverified_ssl is not None else None

        __props__['client_debug'] = pulumi.Output.from_input(client_debug).apply(json.dumps) if client_debug is not None else None

        __props__['client_debug_path'] = client_debug_path

        __props__['client_debug_path_run'] = client_debug_path_run

        if password is None:
            raise TypeError("Missing required property 'password'")
        __props__['password'] = password

        __props__['persist_session'] = pulumi.Output.from_input(persist_session).apply(json.dumps) if persist_session is not None else None

        __props__['rest_session_path'] = rest_session_path

        if user is None:
            raise TypeError("Missing required property 'user'")
        __props__['user'] = user

        __props__['vcenter_server'] = vcenter_server

        __props__['vim_keep_alive'] = pulumi.Output.from_input(vim_keep_alive).apply(json.dumps) if vim_keep_alive is not None else None

        __props__['vim_session_path'] = vim_session_path

        __props__['vsphere_server'] = vsphere_server

        super(Provider, __self__).__init__(
            'vsphere',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


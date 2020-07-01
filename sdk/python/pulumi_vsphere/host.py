# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class Host(pulumi.CustomResource):
    cluster: pulumi.Output[str]
    """
    The ID of the Compute Cluster this host should
    be added to. This should not be set if `datacenter` is set. Conflicts with:
    `cluster`.
    """
    cluster_managed: pulumi.Output[bool]
    """
    Can be set to `true` if compute cluster
    membership will be managed through the `compute_cluster` resource rather
    than the`host` resource. Conflicts with: `cluster`.
    """
    connected: pulumi.Output[bool]
    """
    If set to false then the host will be disconected.
    Default is `false`.
    """
    datacenter: pulumi.Output[str]
    """
    The ID of the datacenter this host should
    be added to. This should not be set if `cluster` is set.
    """
    force: pulumi.Output[bool]
    """
    If set to true then it will force the host to be added, even
    if the host is already connected to a different vSphere instance. Default is `false`
    """
    hostname: pulumi.Output[str]
    """
    FQDN or IP address of the host to be added.
    """
    license: pulumi.Output[str]
    """
    The license key that will be applied to the host.
    The license key is expected to be present in vSphere.
    """
    lockdown: pulumi.Output[str]
    """
    Set the lockdown state of the host. Valid options are
    `disabled`, `normal`, and `strict`. Default is `disabled`.
    """
    maintenance: pulumi.Output[bool]
    """
    Set the management state of the host. Default is `false`.
    """
    password: pulumi.Output[str]
    """
    Password that will be used by vSphere to authenticate
    to the host.
    """
    thumbprint: pulumi.Output[str]
    """
    Host's certificate SHA-1 thumbprint. If not set the the
    CA that signed the host's certificate should be trusted. If the CA is not trusted
    and no thumbprint is set then the operation will fail.
    """
    username: pulumi.Output[str]
    """
    Username that will be used by vSphere to authenticate
    to the host.
    """
    def __init__(__self__, resource_name, opts=None, cluster=None, cluster_managed=None, connected=None, datacenter=None, force=None, hostname=None, license=None, lockdown=None, maintenance=None, password=None, thumbprint=None, username=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a VMware vSphere host resource. This represents an ESXi host that
        can be used either as part of a Compute Cluster or Standalone.

        ## Example Usage
        ### Create a standalone host

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        dc = vsphere.get_datacenter(name="my-datacenter")
        h1 = vsphere.Host("h1",
            hostname="10.10.10.1",
            username="root",
            password="password",
            license="00000-00000-00000-00000i-00000",
            datacenter=dc.id)
        ```
        ### Create host in a compute cluster

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        dc = vsphere.get_datacenter(name="TfDatacenter")
        c1 = vsphere.get_compute_cluster(name="DC0_C0",
            datacenter_id=dc.id)
        h1 = vsphere.Host("h1",
            hostname="10.10.10.1",
            username="root",
            password="password",
            license="00000-00000-00000-00000i-00000",
            cluster=c1.id)
        ```
        ## Importing

        An existing host can be [imported][docs-import] into this resource
        via supplying the host's ID. An example is below:

        [docs-import]: /docs/import/index.html

        ```python
        import pulumi
        ```

        The above would import the host with ID `host-123`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster: The ID of the Compute Cluster this host should
               be added to. This should not be set if `datacenter` is set. Conflicts with:
               `cluster`.
        :param pulumi.Input[bool] cluster_managed: Can be set to `true` if compute cluster
               membership will be managed through the `compute_cluster` resource rather
               than the`host` resource. Conflicts with: `cluster`.
        :param pulumi.Input[bool] connected: If set to false then the host will be disconected.
               Default is `false`.
        :param pulumi.Input[str] datacenter: The ID of the datacenter this host should
               be added to. This should not be set if `cluster` is set.
        :param pulumi.Input[bool] force: If set to true then it will force the host to be added, even
               if the host is already connected to a different vSphere instance. Default is `false`
        :param pulumi.Input[str] hostname: FQDN or IP address of the host to be added.
        :param pulumi.Input[str] license: The license key that will be applied to the host.
               The license key is expected to be present in vSphere.
        :param pulumi.Input[str] lockdown: Set the lockdown state of the host. Valid options are
               `disabled`, `normal`, and `strict`. Default is `disabled`.
        :param pulumi.Input[bool] maintenance: Set the management state of the host. Default is `false`.
        :param pulumi.Input[str] password: Password that will be used by vSphere to authenticate
               to the host.
        :param pulumi.Input[str] thumbprint: Host's certificate SHA-1 thumbprint. If not set the the
               CA that signed the host's certificate should be trusted. If the CA is not trusted
               and no thumbprint is set then the operation will fail.
        :param pulumi.Input[str] username: Username that will be used by vSphere to authenticate
               to the host.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['cluster'] = cluster
            __props__['cluster_managed'] = cluster_managed
            __props__['connected'] = connected
            __props__['datacenter'] = datacenter
            __props__['force'] = force
            if hostname is None:
                raise TypeError("Missing required property 'hostname'")
            __props__['hostname'] = hostname
            __props__['license'] = license
            __props__['lockdown'] = lockdown
            __props__['maintenance'] = maintenance
            if password is None:
                raise TypeError("Missing required property 'password'")
            __props__['password'] = password
            __props__['thumbprint'] = thumbprint
            if username is None:
                raise TypeError("Missing required property 'username'")
            __props__['username'] = username
        super(Host, __self__).__init__(
            'vsphere:index/host:Host',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, cluster=None, cluster_managed=None, connected=None, datacenter=None, force=None, hostname=None, license=None, lockdown=None, maintenance=None, password=None, thumbprint=None, username=None):
        """
        Get an existing Host resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster: The ID of the Compute Cluster this host should
               be added to. This should not be set if `datacenter` is set. Conflicts with:
               `cluster`.
        :param pulumi.Input[bool] cluster_managed: Can be set to `true` if compute cluster
               membership will be managed through the `compute_cluster` resource rather
               than the`host` resource. Conflicts with: `cluster`.
        :param pulumi.Input[bool] connected: If set to false then the host will be disconected.
               Default is `false`.
        :param pulumi.Input[str] datacenter: The ID of the datacenter this host should
               be added to. This should not be set if `cluster` is set.
        :param pulumi.Input[bool] force: If set to true then it will force the host to be added, even
               if the host is already connected to a different vSphere instance. Default is `false`
        :param pulumi.Input[str] hostname: FQDN or IP address of the host to be added.
        :param pulumi.Input[str] license: The license key that will be applied to the host.
               The license key is expected to be present in vSphere.
        :param pulumi.Input[str] lockdown: Set the lockdown state of the host. Valid options are
               `disabled`, `normal`, and `strict`. Default is `disabled`.
        :param pulumi.Input[bool] maintenance: Set the management state of the host. Default is `false`.
        :param pulumi.Input[str] password: Password that will be used by vSphere to authenticate
               to the host.
        :param pulumi.Input[str] thumbprint: Host's certificate SHA-1 thumbprint. If not set the the
               CA that signed the host's certificate should be trusted. If the CA is not trusted
               and no thumbprint is set then the operation will fail.
        :param pulumi.Input[str] username: Username that will be used by vSphere to authenticate
               to the host.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cluster"] = cluster
        __props__["cluster_managed"] = cluster_managed
        __props__["connected"] = connected
        __props__["datacenter"] = datacenter
        __props__["force"] = force
        __props__["hostname"] = hostname
        __props__["license"] = license
        __props__["lockdown"] = lockdown
        __props__["maintenance"] = maintenance
        __props__["password"] = password
        __props__["thumbprint"] = thumbprint
        __props__["username"] = username
        return Host(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

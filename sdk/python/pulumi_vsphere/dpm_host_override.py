# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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

__all__ = ['DpmHostOverrideArgs', 'DpmHostOverride']

@pulumi.input_type
class DpmHostOverrideArgs:
    def __init__(__self__, *,
                 compute_cluster_id: pulumi.Input[builtins.str],
                 host_system_id: pulumi.Input[builtins.str],
                 dpm_automation_level: Optional[pulumi.Input[builtins.str]] = None,
                 dpm_enabled: Optional[pulumi.Input[builtins.bool]] = None):
        """
        The set of arguments for constructing a DpmHostOverride resource.
        :param pulumi.Input[builtins.str] compute_cluster_id: The managed object reference
               ID of the cluster to put the override in.  Forces a new
               resource if changed.
        :param pulumi.Input[builtins.str] host_system_id: The managed object ID of the host.
        :param pulumi.Input[builtins.str] dpm_automation_level: The automation level for host power
               operations on this host. Can be one of `manual` or `automated`. Default:
               `manual`.
               
               > **NOTE:** Using this resource _always_ implies an override, even if one of
               `dpm_enabled` or `dpm_automation_level` is omitted. Take note of the defaults
               for both options.
        :param pulumi.Input[builtins.bool] dpm_enabled: Enable DPM support for this host. Default:
               `false`.
        """
        pulumi.set(__self__, "compute_cluster_id", compute_cluster_id)
        pulumi.set(__self__, "host_system_id", host_system_id)
        if dpm_automation_level is not None:
            pulumi.set(__self__, "dpm_automation_level", dpm_automation_level)
        if dpm_enabled is not None:
            pulumi.set(__self__, "dpm_enabled", dpm_enabled)

    @property
    @pulumi.getter(name="computeClusterId")
    def compute_cluster_id(self) -> pulumi.Input[builtins.str]:
        """
        The managed object reference
        ID of the cluster to put the override in.  Forces a new
        resource if changed.
        """
        return pulumi.get(self, "compute_cluster_id")

    @compute_cluster_id.setter
    def compute_cluster_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "compute_cluster_id", value)

    @property
    @pulumi.getter(name="hostSystemId")
    def host_system_id(self) -> pulumi.Input[builtins.str]:
        """
        The managed object ID of the host.
        """
        return pulumi.get(self, "host_system_id")

    @host_system_id.setter
    def host_system_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "host_system_id", value)

    @property
    @pulumi.getter(name="dpmAutomationLevel")
    def dpm_automation_level(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The automation level for host power
        operations on this host. Can be one of `manual` or `automated`. Default:
        `manual`.

        > **NOTE:** Using this resource _always_ implies an override, even if one of
        `dpm_enabled` or `dpm_automation_level` is omitted. Take note of the defaults
        for both options.
        """
        return pulumi.get(self, "dpm_automation_level")

    @dpm_automation_level.setter
    def dpm_automation_level(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "dpm_automation_level", value)

    @property
    @pulumi.getter(name="dpmEnabled")
    def dpm_enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Enable DPM support for this host. Default:
        `false`.
        """
        return pulumi.get(self, "dpm_enabled")

    @dpm_enabled.setter
    def dpm_enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "dpm_enabled", value)


@pulumi.input_type
class _DpmHostOverrideState:
    def __init__(__self__, *,
                 compute_cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 dpm_automation_level: Optional[pulumi.Input[builtins.str]] = None,
                 dpm_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 host_system_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering DpmHostOverride resources.
        :param pulumi.Input[builtins.str] compute_cluster_id: The managed object reference
               ID of the cluster to put the override in.  Forces a new
               resource if changed.
        :param pulumi.Input[builtins.str] dpm_automation_level: The automation level for host power
               operations on this host. Can be one of `manual` or `automated`. Default:
               `manual`.
               
               > **NOTE:** Using this resource _always_ implies an override, even if one of
               `dpm_enabled` or `dpm_automation_level` is omitted. Take note of the defaults
               for both options.
        :param pulumi.Input[builtins.bool] dpm_enabled: Enable DPM support for this host. Default:
               `false`.
        :param pulumi.Input[builtins.str] host_system_id: The managed object ID of the host.
        """
        if compute_cluster_id is not None:
            pulumi.set(__self__, "compute_cluster_id", compute_cluster_id)
        if dpm_automation_level is not None:
            pulumi.set(__self__, "dpm_automation_level", dpm_automation_level)
        if dpm_enabled is not None:
            pulumi.set(__self__, "dpm_enabled", dpm_enabled)
        if host_system_id is not None:
            pulumi.set(__self__, "host_system_id", host_system_id)

    @property
    @pulumi.getter(name="computeClusterId")
    def compute_cluster_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The managed object reference
        ID of the cluster to put the override in.  Forces a new
        resource if changed.
        """
        return pulumi.get(self, "compute_cluster_id")

    @compute_cluster_id.setter
    def compute_cluster_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "compute_cluster_id", value)

    @property
    @pulumi.getter(name="dpmAutomationLevel")
    def dpm_automation_level(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The automation level for host power
        operations on this host. Can be one of `manual` or `automated`. Default:
        `manual`.

        > **NOTE:** Using this resource _always_ implies an override, even if one of
        `dpm_enabled` or `dpm_automation_level` is omitted. Take note of the defaults
        for both options.
        """
        return pulumi.get(self, "dpm_automation_level")

    @dpm_automation_level.setter
    def dpm_automation_level(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "dpm_automation_level", value)

    @property
    @pulumi.getter(name="dpmEnabled")
    def dpm_enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Enable DPM support for this host. Default:
        `false`.
        """
        return pulumi.get(self, "dpm_enabled")

    @dpm_enabled.setter
    def dpm_enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "dpm_enabled", value)

    @property
    @pulumi.getter(name="hostSystemId")
    def host_system_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The managed object ID of the host.
        """
        return pulumi.get(self, "host_system_id")

    @host_system_id.setter
    def host_system_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "host_system_id", value)


class DpmHostOverride(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 dpm_automation_level: Optional[pulumi.Input[builtins.str]] = None,
                 dpm_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 host_system_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The `DpmHostOverride` resource can be used to add a DPM override to a
        cluster for a particular host. This allows you to control the power management
        settings for individual hosts in the cluster while leaving any unspecified ones
        at the default power management settings.

        For more information on DPM within vSphere clusters, see [this
        page][ref-vsphere-cluster-dpm].

        [ref-vsphere-cluster-dpm]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-resource-management-8-0/using-drs-clusters-to-manage-resources/managing-power-resources.html

        > **NOTE:** This resource requires vCenter and is not available on direct ESXi
        connections.

        ## Import

        An existing override can be imported into this resource by

        supplying both the path to the cluster, and the path to the host, to `terraform

        import`. If no override exists, an error will be given.  An example is below:

        ```sh
        $ pulumi import vsphere:index/dpmHostOverride:DpmHostOverride dpm_host_override \\
        ```

          '{"compute_cluster_path": "/dc1/host/cluster1", \\

          "host_path": "/dc1/host/esxi1"}'

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] compute_cluster_id: The managed object reference
               ID of the cluster to put the override in.  Forces a new
               resource if changed.
        :param pulumi.Input[builtins.str] dpm_automation_level: The automation level for host power
               operations on this host. Can be one of `manual` or `automated`. Default:
               `manual`.
               
               > **NOTE:** Using this resource _always_ implies an override, even if one of
               `dpm_enabled` or `dpm_automation_level` is omitted. Take note of the defaults
               for both options.
        :param pulumi.Input[builtins.bool] dpm_enabled: Enable DPM support for this host. Default:
               `false`.
        :param pulumi.Input[builtins.str] host_system_id: The managed object ID of the host.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DpmHostOverrideArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `DpmHostOverride` resource can be used to add a DPM override to a
        cluster for a particular host. This allows you to control the power management
        settings for individual hosts in the cluster while leaving any unspecified ones
        at the default power management settings.

        For more information on DPM within vSphere clusters, see [this
        page][ref-vsphere-cluster-dpm].

        [ref-vsphere-cluster-dpm]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-resource-management-8-0/using-drs-clusters-to-manage-resources/managing-power-resources.html

        > **NOTE:** This resource requires vCenter and is not available on direct ESXi
        connections.

        ## Import

        An existing override can be imported into this resource by

        supplying both the path to the cluster, and the path to the host, to `terraform

        import`. If no override exists, an error will be given.  An example is below:

        ```sh
        $ pulumi import vsphere:index/dpmHostOverride:DpmHostOverride dpm_host_override \\
        ```

          '{"compute_cluster_path": "/dc1/host/cluster1", \\

          "host_path": "/dc1/host/esxi1"}'

        :param str resource_name: The name of the resource.
        :param DpmHostOverrideArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DpmHostOverrideArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 dpm_automation_level: Optional[pulumi.Input[builtins.str]] = None,
                 dpm_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 host_system_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DpmHostOverrideArgs.__new__(DpmHostOverrideArgs)

            if compute_cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'compute_cluster_id'")
            __props__.__dict__["compute_cluster_id"] = compute_cluster_id
            __props__.__dict__["dpm_automation_level"] = dpm_automation_level
            __props__.__dict__["dpm_enabled"] = dpm_enabled
            if host_system_id is None and not opts.urn:
                raise TypeError("Missing required property 'host_system_id'")
            __props__.__dict__["host_system_id"] = host_system_id
        super(DpmHostOverride, __self__).__init__(
            'vsphere:index/dpmHostOverride:DpmHostOverride',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            compute_cluster_id: Optional[pulumi.Input[builtins.str]] = None,
            dpm_automation_level: Optional[pulumi.Input[builtins.str]] = None,
            dpm_enabled: Optional[pulumi.Input[builtins.bool]] = None,
            host_system_id: Optional[pulumi.Input[builtins.str]] = None) -> 'DpmHostOverride':
        """
        Get an existing DpmHostOverride resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] compute_cluster_id: The managed object reference
               ID of the cluster to put the override in.  Forces a new
               resource if changed.
        :param pulumi.Input[builtins.str] dpm_automation_level: The automation level for host power
               operations on this host. Can be one of `manual` or `automated`. Default:
               `manual`.
               
               > **NOTE:** Using this resource _always_ implies an override, even if one of
               `dpm_enabled` or `dpm_automation_level` is omitted. Take note of the defaults
               for both options.
        :param pulumi.Input[builtins.bool] dpm_enabled: Enable DPM support for this host. Default:
               `false`.
        :param pulumi.Input[builtins.str] host_system_id: The managed object ID of the host.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DpmHostOverrideState.__new__(_DpmHostOverrideState)

        __props__.__dict__["compute_cluster_id"] = compute_cluster_id
        __props__.__dict__["dpm_automation_level"] = dpm_automation_level
        __props__.__dict__["dpm_enabled"] = dpm_enabled
        __props__.__dict__["host_system_id"] = host_system_id
        return DpmHostOverride(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="computeClusterId")
    def compute_cluster_id(self) -> pulumi.Output[builtins.str]:
        """
        The managed object reference
        ID of the cluster to put the override in.  Forces a new
        resource if changed.
        """
        return pulumi.get(self, "compute_cluster_id")

    @property
    @pulumi.getter(name="dpmAutomationLevel")
    def dpm_automation_level(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The automation level for host power
        operations on this host. Can be one of `manual` or `automated`. Default:
        `manual`.

        > **NOTE:** Using this resource _always_ implies an override, even if one of
        `dpm_enabled` or `dpm_automation_level` is omitted. Take note of the defaults
        for both options.
        """
        return pulumi.get(self, "dpm_automation_level")

    @property
    @pulumi.getter(name="dpmEnabled")
    def dpm_enabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Enable DPM support for this host. Default:
        `false`.
        """
        return pulumi.get(self, "dpm_enabled")

    @property
    @pulumi.getter(name="hostSystemId")
    def host_system_id(self) -> pulumi.Output[builtins.str]:
        """
        The managed object ID of the host.
        """
        return pulumi.get(self, "host_system_id")


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ComputeClusterVmHostRuleArgs', 'ComputeClusterVmHostRule']

@pulumi.input_type
class ComputeClusterVmHostRuleArgs:
    def __init__(__self__, *,
                 compute_cluster_id: pulumi.Input[str],
                 vm_group_name: pulumi.Input[str],
                 affinity_host_group_name: Optional[pulumi.Input[str]] = None,
                 anti_affinity_host_group_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 mandatory: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ComputeClusterVmHostRule resource.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the group in.  Forces a new
               resource if changed.
        :param pulumi.Input[str] vm_group_name: The name of the virtual machine group to use
               with this rule.
        :param pulumi.Input[str] affinity_host_group_name: When this field is used, the virtual
               machines defined in `vm_group_name` will be run on the
               hosts defined in this host group.
        :param pulumi.Input[str] anti_affinity_host_group_name: When this field is used, the
               virtual machines defined in `vm_group_name` will _not_ be
               run on the hosts defined in this host group.
        :param pulumi.Input[bool] enabled: Enable this rule in the cluster. Default: `true`.
        :param pulumi.Input[bool] mandatory: When this value is `true`, prevents any virtual
               machine operations that may violate this rule. Default: `false`.
               
               > **NOTE:** One of `affinity_host_group_name` or
               `anti_affinity_host_group_name` must be
               defined, but not both.
               
               > **NOTE:** The namespace for rule names on this resource (defined by the
               `name` argument) is shared with all rules in the cluster - consider
               this when naming your rules.
        :param pulumi.Input[str] name: The name of the rule. This must be unique in the
               cluster.
        """
        ComputeClusterVmHostRuleArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            compute_cluster_id=compute_cluster_id,
            vm_group_name=vm_group_name,
            affinity_host_group_name=affinity_host_group_name,
            anti_affinity_host_group_name=anti_affinity_host_group_name,
            enabled=enabled,
            mandatory=mandatory,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             compute_cluster_id: Optional[pulumi.Input[str]] = None,
             vm_group_name: Optional[pulumi.Input[str]] = None,
             affinity_host_group_name: Optional[pulumi.Input[str]] = None,
             anti_affinity_host_group_name: Optional[pulumi.Input[str]] = None,
             enabled: Optional[pulumi.Input[bool]] = None,
             mandatory: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if compute_cluster_id is None and 'computeClusterId' in kwargs:
            compute_cluster_id = kwargs['computeClusterId']
        if compute_cluster_id is None:
            raise TypeError("Missing 'compute_cluster_id' argument")
        if vm_group_name is None and 'vmGroupName' in kwargs:
            vm_group_name = kwargs['vmGroupName']
        if vm_group_name is None:
            raise TypeError("Missing 'vm_group_name' argument")
        if affinity_host_group_name is None and 'affinityHostGroupName' in kwargs:
            affinity_host_group_name = kwargs['affinityHostGroupName']
        if anti_affinity_host_group_name is None and 'antiAffinityHostGroupName' in kwargs:
            anti_affinity_host_group_name = kwargs['antiAffinityHostGroupName']

        _setter("compute_cluster_id", compute_cluster_id)
        _setter("vm_group_name", vm_group_name)
        if affinity_host_group_name is not None:
            _setter("affinity_host_group_name", affinity_host_group_name)
        if anti_affinity_host_group_name is not None:
            _setter("anti_affinity_host_group_name", anti_affinity_host_group_name)
        if enabled is not None:
            _setter("enabled", enabled)
        if mandatory is not None:
            _setter("mandatory", mandatory)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter(name="computeClusterId")
    def compute_cluster_id(self) -> pulumi.Input[str]:
        """
        The managed object reference
        ID of the cluster to put the group in.  Forces a new
        resource if changed.
        """
        return pulumi.get(self, "compute_cluster_id")

    @compute_cluster_id.setter
    def compute_cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "compute_cluster_id", value)

    @property
    @pulumi.getter(name="vmGroupName")
    def vm_group_name(self) -> pulumi.Input[str]:
        """
        The name of the virtual machine group to use
        with this rule.
        """
        return pulumi.get(self, "vm_group_name")

    @vm_group_name.setter
    def vm_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "vm_group_name", value)

    @property
    @pulumi.getter(name="affinityHostGroupName")
    def affinity_host_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        When this field is used, the virtual
        machines defined in `vm_group_name` will be run on the
        hosts defined in this host group.
        """
        return pulumi.get(self, "affinity_host_group_name")

    @affinity_host_group_name.setter
    def affinity_host_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "affinity_host_group_name", value)

    @property
    @pulumi.getter(name="antiAffinityHostGroupName")
    def anti_affinity_host_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        When this field is used, the
        virtual machines defined in `vm_group_name` will _not_ be
        run on the hosts defined in this host group.
        """
        return pulumi.get(self, "anti_affinity_host_group_name")

    @anti_affinity_host_group_name.setter
    def anti_affinity_host_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "anti_affinity_host_group_name", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable this rule in the cluster. Default: `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def mandatory(self) -> Optional[pulumi.Input[bool]]:
        """
        When this value is `true`, prevents any virtual
        machine operations that may violate this rule. Default: `false`.

        > **NOTE:** One of `affinity_host_group_name` or
        `anti_affinity_host_group_name` must be
        defined, but not both.

        > **NOTE:** The namespace for rule names on this resource (defined by the
        `name` argument) is shared with all rules in the cluster - consider
        this when naming your rules.
        """
        return pulumi.get(self, "mandatory")

    @mandatory.setter
    def mandatory(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "mandatory", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the rule. This must be unique in the
        cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ComputeClusterVmHostRuleState:
    def __init__(__self__, *,
                 affinity_host_group_name: Optional[pulumi.Input[str]] = None,
                 anti_affinity_host_group_name: Optional[pulumi.Input[str]] = None,
                 compute_cluster_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 mandatory: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vm_group_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ComputeClusterVmHostRule resources.
        :param pulumi.Input[str] affinity_host_group_name: When this field is used, the virtual
               machines defined in `vm_group_name` will be run on the
               hosts defined in this host group.
        :param pulumi.Input[str] anti_affinity_host_group_name: When this field is used, the
               virtual machines defined in `vm_group_name` will _not_ be
               run on the hosts defined in this host group.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the group in.  Forces a new
               resource if changed.
        :param pulumi.Input[bool] enabled: Enable this rule in the cluster. Default: `true`.
        :param pulumi.Input[bool] mandatory: When this value is `true`, prevents any virtual
               machine operations that may violate this rule. Default: `false`.
               
               > **NOTE:** One of `affinity_host_group_name` or
               `anti_affinity_host_group_name` must be
               defined, but not both.
               
               > **NOTE:** The namespace for rule names on this resource (defined by the
               `name` argument) is shared with all rules in the cluster - consider
               this when naming your rules.
        :param pulumi.Input[str] name: The name of the rule. This must be unique in the
               cluster.
        :param pulumi.Input[str] vm_group_name: The name of the virtual machine group to use
               with this rule.
        """
        _ComputeClusterVmHostRuleState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            affinity_host_group_name=affinity_host_group_name,
            anti_affinity_host_group_name=anti_affinity_host_group_name,
            compute_cluster_id=compute_cluster_id,
            enabled=enabled,
            mandatory=mandatory,
            name=name,
            vm_group_name=vm_group_name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             affinity_host_group_name: Optional[pulumi.Input[str]] = None,
             anti_affinity_host_group_name: Optional[pulumi.Input[str]] = None,
             compute_cluster_id: Optional[pulumi.Input[str]] = None,
             enabled: Optional[pulumi.Input[bool]] = None,
             mandatory: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             vm_group_name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if affinity_host_group_name is None and 'affinityHostGroupName' in kwargs:
            affinity_host_group_name = kwargs['affinityHostGroupName']
        if anti_affinity_host_group_name is None and 'antiAffinityHostGroupName' in kwargs:
            anti_affinity_host_group_name = kwargs['antiAffinityHostGroupName']
        if compute_cluster_id is None and 'computeClusterId' in kwargs:
            compute_cluster_id = kwargs['computeClusterId']
        if vm_group_name is None and 'vmGroupName' in kwargs:
            vm_group_name = kwargs['vmGroupName']

        if affinity_host_group_name is not None:
            _setter("affinity_host_group_name", affinity_host_group_name)
        if anti_affinity_host_group_name is not None:
            _setter("anti_affinity_host_group_name", anti_affinity_host_group_name)
        if compute_cluster_id is not None:
            _setter("compute_cluster_id", compute_cluster_id)
        if enabled is not None:
            _setter("enabled", enabled)
        if mandatory is not None:
            _setter("mandatory", mandatory)
        if name is not None:
            _setter("name", name)
        if vm_group_name is not None:
            _setter("vm_group_name", vm_group_name)

    @property
    @pulumi.getter(name="affinityHostGroupName")
    def affinity_host_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        When this field is used, the virtual
        machines defined in `vm_group_name` will be run on the
        hosts defined in this host group.
        """
        return pulumi.get(self, "affinity_host_group_name")

    @affinity_host_group_name.setter
    def affinity_host_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "affinity_host_group_name", value)

    @property
    @pulumi.getter(name="antiAffinityHostGroupName")
    def anti_affinity_host_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        When this field is used, the
        virtual machines defined in `vm_group_name` will _not_ be
        run on the hosts defined in this host group.
        """
        return pulumi.get(self, "anti_affinity_host_group_name")

    @anti_affinity_host_group_name.setter
    def anti_affinity_host_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "anti_affinity_host_group_name", value)

    @property
    @pulumi.getter(name="computeClusterId")
    def compute_cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed object reference
        ID of the cluster to put the group in.  Forces a new
        resource if changed.
        """
        return pulumi.get(self, "compute_cluster_id")

    @compute_cluster_id.setter
    def compute_cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compute_cluster_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable this rule in the cluster. Default: `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def mandatory(self) -> Optional[pulumi.Input[bool]]:
        """
        When this value is `true`, prevents any virtual
        machine operations that may violate this rule. Default: `false`.

        > **NOTE:** One of `affinity_host_group_name` or
        `anti_affinity_host_group_name` must be
        defined, but not both.

        > **NOTE:** The namespace for rule names on this resource (defined by the
        `name` argument) is shared with all rules in the cluster - consider
        this when naming your rules.
        """
        return pulumi.get(self, "mandatory")

    @mandatory.setter
    def mandatory(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "mandatory", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the rule. This must be unique in the
        cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="vmGroupName")
    def vm_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the virtual machine group to use
        with this rule.
        """
        return pulumi.get(self, "vm_group_name")

    @vm_group_name.setter
    def vm_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vm_group_name", value)


class ComputeClusterVmHostRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 affinity_host_group_name: Optional[pulumi.Input[str]] = None,
                 anti_affinity_host_group_name: Optional[pulumi.Input[str]] = None,
                 compute_cluster_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 mandatory: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vm_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ComputeClusterVmHostRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] affinity_host_group_name: When this field is used, the virtual
               machines defined in `vm_group_name` will be run on the
               hosts defined in this host group.
        :param pulumi.Input[str] anti_affinity_host_group_name: When this field is used, the
               virtual machines defined in `vm_group_name` will _not_ be
               run on the hosts defined in this host group.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the group in.  Forces a new
               resource if changed.
        :param pulumi.Input[bool] enabled: Enable this rule in the cluster. Default: `true`.
        :param pulumi.Input[bool] mandatory: When this value is `true`, prevents any virtual
               machine operations that may violate this rule. Default: `false`.
               
               > **NOTE:** One of `affinity_host_group_name` or
               `anti_affinity_host_group_name` must be
               defined, but not both.
               
               > **NOTE:** The namespace for rule names on this resource (defined by the
               `name` argument) is shared with all rules in the cluster - consider
               this when naming your rules.
        :param pulumi.Input[str] name: The name of the rule. This must be unique in the
               cluster.
        :param pulumi.Input[str] vm_group_name: The name of the virtual machine group to use
               with this rule.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ComputeClusterVmHostRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ComputeClusterVmHostRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ComputeClusterVmHostRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ComputeClusterVmHostRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ComputeClusterVmHostRuleArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 affinity_host_group_name: Optional[pulumi.Input[str]] = None,
                 anti_affinity_host_group_name: Optional[pulumi.Input[str]] = None,
                 compute_cluster_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 mandatory: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vm_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ComputeClusterVmHostRuleArgs.__new__(ComputeClusterVmHostRuleArgs)

            __props__.__dict__["affinity_host_group_name"] = affinity_host_group_name
            __props__.__dict__["anti_affinity_host_group_name"] = anti_affinity_host_group_name
            if compute_cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'compute_cluster_id'")
            __props__.__dict__["compute_cluster_id"] = compute_cluster_id
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["mandatory"] = mandatory
            __props__.__dict__["name"] = name
            if vm_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'vm_group_name'")
            __props__.__dict__["vm_group_name"] = vm_group_name
        super(ComputeClusterVmHostRule, __self__).__init__(
            'vsphere:index/computeClusterVmHostRule:ComputeClusterVmHostRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            affinity_host_group_name: Optional[pulumi.Input[str]] = None,
            anti_affinity_host_group_name: Optional[pulumi.Input[str]] = None,
            compute_cluster_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            mandatory: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vm_group_name: Optional[pulumi.Input[str]] = None) -> 'ComputeClusterVmHostRule':
        """
        Get an existing ComputeClusterVmHostRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] affinity_host_group_name: When this field is used, the virtual
               machines defined in `vm_group_name` will be run on the
               hosts defined in this host group.
        :param pulumi.Input[str] anti_affinity_host_group_name: When this field is used, the
               virtual machines defined in `vm_group_name` will _not_ be
               run on the hosts defined in this host group.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the group in.  Forces a new
               resource if changed.
        :param pulumi.Input[bool] enabled: Enable this rule in the cluster. Default: `true`.
        :param pulumi.Input[bool] mandatory: When this value is `true`, prevents any virtual
               machine operations that may violate this rule. Default: `false`.
               
               > **NOTE:** One of `affinity_host_group_name` or
               `anti_affinity_host_group_name` must be
               defined, but not both.
               
               > **NOTE:** The namespace for rule names on this resource (defined by the
               `name` argument) is shared with all rules in the cluster - consider
               this when naming your rules.
        :param pulumi.Input[str] name: The name of the rule. This must be unique in the
               cluster.
        :param pulumi.Input[str] vm_group_name: The name of the virtual machine group to use
               with this rule.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ComputeClusterVmHostRuleState.__new__(_ComputeClusterVmHostRuleState)

        __props__.__dict__["affinity_host_group_name"] = affinity_host_group_name
        __props__.__dict__["anti_affinity_host_group_name"] = anti_affinity_host_group_name
        __props__.__dict__["compute_cluster_id"] = compute_cluster_id
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["mandatory"] = mandatory
        __props__.__dict__["name"] = name
        __props__.__dict__["vm_group_name"] = vm_group_name
        return ComputeClusterVmHostRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="affinityHostGroupName")
    def affinity_host_group_name(self) -> pulumi.Output[Optional[str]]:
        """
        When this field is used, the virtual
        machines defined in `vm_group_name` will be run on the
        hosts defined in this host group.
        """
        return pulumi.get(self, "affinity_host_group_name")

    @property
    @pulumi.getter(name="antiAffinityHostGroupName")
    def anti_affinity_host_group_name(self) -> pulumi.Output[Optional[str]]:
        """
        When this field is used, the
        virtual machines defined in `vm_group_name` will _not_ be
        run on the hosts defined in this host group.
        """
        return pulumi.get(self, "anti_affinity_host_group_name")

    @property
    @pulumi.getter(name="computeClusterId")
    def compute_cluster_id(self) -> pulumi.Output[str]:
        """
        The managed object reference
        ID of the cluster to put the group in.  Forces a new
        resource if changed.
        """
        return pulumi.get(self, "compute_cluster_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable this rule in the cluster. Default: `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def mandatory(self) -> pulumi.Output[Optional[bool]]:
        """
        When this value is `true`, prevents any virtual
        machine operations that may violate this rule. Default: `false`.

        > **NOTE:** One of `affinity_host_group_name` or
        `anti_affinity_host_group_name` must be
        defined, but not both.

        > **NOTE:** The namespace for rule names on this resource (defined by the
        `name` argument) is shared with all rules in the cluster - consider
        this when naming your rules.
        """
        return pulumi.get(self, "mandatory")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the rule. This must be unique in the
        cluster.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="vmGroupName")
    def vm_group_name(self) -> pulumi.Output[str]:
        """
        The name of the virtual machine group to use
        with this rule.
        """
        return pulumi.get(self, "vm_group_name")


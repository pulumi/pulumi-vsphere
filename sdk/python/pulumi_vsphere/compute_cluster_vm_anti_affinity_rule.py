# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ComputeClusterVmAntiAffinityRuleArgs', 'ComputeClusterVmAntiAffinityRule']

@pulumi.input_type
class ComputeClusterVmAntiAffinityRuleArgs:
    def __init__(__self__, *,
                 compute_cluster_id: pulumi.Input[str],
                 virtual_machine_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 enabled: Optional[pulumi.Input[bool]] = None,
                 mandatory: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ComputeClusterVmAntiAffinityRule resource.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the group in.  Forces a new
               resource if changed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] virtual_machine_ids: The UUIDs of the virtual machines to run
               on hosts different from each other.
        :param pulumi.Input[bool] enabled: Enable this rule in the cluster. Default: `true`.
        :param pulumi.Input[bool] mandatory: When this value is `true`, prevents any virtual
               machine operations that may violate this rule. Default: `false`.
               
               > **NOTE:** The namespace for rule names on this resource (defined by the
               `name` argument) is shared with all rules in the cluster - consider
               this when naming your rules.
        :param pulumi.Input[str] name: The name of the rule. This must be unique in the cluster.
        """
        pulumi.set(__self__, "compute_cluster_id", compute_cluster_id)
        pulumi.set(__self__, "virtual_machine_ids", virtual_machine_ids)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if mandatory is not None:
            pulumi.set(__self__, "mandatory", mandatory)
        if name is not None:
            pulumi.set(__self__, "name", name)

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
    @pulumi.getter(name="virtualMachineIds")
    def virtual_machine_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The UUIDs of the virtual machines to run
        on hosts different from each other.
        """
        return pulumi.get(self, "virtual_machine_ids")

    @virtual_machine_ids.setter
    def virtual_machine_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "virtual_machine_ids", value)

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
        The name of the rule. This must be unique in the cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ComputeClusterVmAntiAffinityRuleState:
    def __init__(__self__, *,
                 compute_cluster_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 mandatory: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 virtual_machine_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ComputeClusterVmAntiAffinityRule resources.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the group in.  Forces a new
               resource if changed.
        :param pulumi.Input[bool] enabled: Enable this rule in the cluster. Default: `true`.
        :param pulumi.Input[bool] mandatory: When this value is `true`, prevents any virtual
               machine operations that may violate this rule. Default: `false`.
               
               > **NOTE:** The namespace for rule names on this resource (defined by the
               `name` argument) is shared with all rules in the cluster - consider
               this when naming your rules.
        :param pulumi.Input[str] name: The name of the rule. This must be unique in the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] virtual_machine_ids: The UUIDs of the virtual machines to run
               on hosts different from each other.
        """
        if compute_cluster_id is not None:
            pulumi.set(__self__, "compute_cluster_id", compute_cluster_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if mandatory is not None:
            pulumi.set(__self__, "mandatory", mandatory)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if virtual_machine_ids is not None:
            pulumi.set(__self__, "virtual_machine_ids", virtual_machine_ids)

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
        The name of the rule. This must be unique in the cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="virtualMachineIds")
    def virtual_machine_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The UUIDs of the virtual machines to run
        on hosts different from each other.
        """
        return pulumi.get(self, "virtual_machine_ids")

    @virtual_machine_ids.setter
    def virtual_machine_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "virtual_machine_ids", value)


class ComputeClusterVmAntiAffinityRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_cluster_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 mandatory: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 virtual_machine_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        ## Import

        An existing rule can be imported into this resource by supplying

        both the path to the cluster, and the name the rule. If the name or cluster is

        not found, or if the rule is of a different type, an error will be returned. An

        example is below:

        ```sh
        $ pulumi import vsphere:index/computeClusterVmAntiAffinityRule:ComputeClusterVmAntiAffinityRule vm_anti_affinity_rule \\
        ```

          '{"compute_cluster_path": "/dc-01/host/cluster-01", \\

          "name": "vm-anti-affinity-rule"}'

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the group in.  Forces a new
               resource if changed.
        :param pulumi.Input[bool] enabled: Enable this rule in the cluster. Default: `true`.
        :param pulumi.Input[bool] mandatory: When this value is `true`, prevents any virtual
               machine operations that may violate this rule. Default: `false`.
               
               > **NOTE:** The namespace for rule names on this resource (defined by the
               `name` argument) is shared with all rules in the cluster - consider
               this when naming your rules.
        :param pulumi.Input[str] name: The name of the rule. This must be unique in the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] virtual_machine_ids: The UUIDs of the virtual machines to run
               on hosts different from each other.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ComputeClusterVmAntiAffinityRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        An existing rule can be imported into this resource by supplying

        both the path to the cluster, and the name the rule. If the name or cluster is

        not found, or if the rule is of a different type, an error will be returned. An

        example is below:

        ```sh
        $ pulumi import vsphere:index/computeClusterVmAntiAffinityRule:ComputeClusterVmAntiAffinityRule vm_anti_affinity_rule \\
        ```

          '{"compute_cluster_path": "/dc-01/host/cluster-01", \\

          "name": "vm-anti-affinity-rule"}'

        :param str resource_name: The name of the resource.
        :param ComputeClusterVmAntiAffinityRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ComputeClusterVmAntiAffinityRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_cluster_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 mandatory: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 virtual_machine_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ComputeClusterVmAntiAffinityRuleArgs.__new__(ComputeClusterVmAntiAffinityRuleArgs)

            if compute_cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'compute_cluster_id'")
            __props__.__dict__["compute_cluster_id"] = compute_cluster_id
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["mandatory"] = mandatory
            __props__.__dict__["name"] = name
            if virtual_machine_ids is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_machine_ids'")
            __props__.__dict__["virtual_machine_ids"] = virtual_machine_ids
        super(ComputeClusterVmAntiAffinityRule, __self__).__init__(
            'vsphere:index/computeClusterVmAntiAffinityRule:ComputeClusterVmAntiAffinityRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            compute_cluster_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            mandatory: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            virtual_machine_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'ComputeClusterVmAntiAffinityRule':
        """
        Get an existing ComputeClusterVmAntiAffinityRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the group in.  Forces a new
               resource if changed.
        :param pulumi.Input[bool] enabled: Enable this rule in the cluster. Default: `true`.
        :param pulumi.Input[bool] mandatory: When this value is `true`, prevents any virtual
               machine operations that may violate this rule. Default: `false`.
               
               > **NOTE:** The namespace for rule names on this resource (defined by the
               `name` argument) is shared with all rules in the cluster - consider
               this when naming your rules.
        :param pulumi.Input[str] name: The name of the rule. This must be unique in the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] virtual_machine_ids: The UUIDs of the virtual machines to run
               on hosts different from each other.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ComputeClusterVmAntiAffinityRuleState.__new__(_ComputeClusterVmAntiAffinityRuleState)

        __props__.__dict__["compute_cluster_id"] = compute_cluster_id
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["mandatory"] = mandatory
        __props__.__dict__["name"] = name
        __props__.__dict__["virtual_machine_ids"] = virtual_machine_ids
        return ComputeClusterVmAntiAffinityRule(resource_name, opts=opts, __props__=__props__)

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

        > **NOTE:** The namespace for rule names on this resource (defined by the
        `name` argument) is shared with all rules in the cluster - consider
        this when naming your rules.
        """
        return pulumi.get(self, "mandatory")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the rule. This must be unique in the cluster.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="virtualMachineIds")
    def virtual_machine_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        The UUIDs of the virtual machines to run
        on hosts different from each other.
        """
        return pulumi.get(self, "virtual_machine_ids")


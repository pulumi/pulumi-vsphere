# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['VmStoragePolicyArgs', 'VmStoragePolicy']

@pulumi.input_type
class VmStoragePolicyArgs:
    def __init__(__self__, *,
                 tag_rules: pulumi.Input[Sequence[pulumi.Input['VmStoragePolicyTagRuleArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VmStoragePolicy resource.
        :param pulumi.Input[Sequence[pulumi.Input['VmStoragePolicyTagRuleArgs']]] tag_rules: List of tag rules. The tag category and tags to be associated to this storage policy.
        :param pulumi.Input[str] description: Description of the storage policy.
        :param pulumi.Input[str] name: The name of the storage policy.
        """
        pulumi.set(__self__, "tag_rules", tag_rules)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="tagRules")
    def tag_rules(self) -> pulumi.Input[Sequence[pulumi.Input['VmStoragePolicyTagRuleArgs']]]:
        """
        List of tag rules. The tag category and tags to be associated to this storage policy.
        """
        return pulumi.get(self, "tag_rules")

    @tag_rules.setter
    def tag_rules(self, value: pulumi.Input[Sequence[pulumi.Input['VmStoragePolicyTagRuleArgs']]]):
        pulumi.set(self, "tag_rules", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the storage policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the storage policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _VmStoragePolicyState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tag_rules: Optional[pulumi.Input[Sequence[pulumi.Input['VmStoragePolicyTagRuleArgs']]]] = None):
        """
        Input properties used for looking up and filtering VmStoragePolicy resources.
        :param pulumi.Input[str] description: Description of the storage policy.
        :param pulumi.Input[str] name: The name of the storage policy.
        :param pulumi.Input[Sequence[pulumi.Input['VmStoragePolicyTagRuleArgs']]] tag_rules: List of tag rules. The tag category and tags to be associated to this storage policy.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tag_rules is not None:
            pulumi.set(__self__, "tag_rules", tag_rules)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the storage policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the storage policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="tagRules")
    def tag_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VmStoragePolicyTagRuleArgs']]]]:
        """
        List of tag rules. The tag category and tags to be associated to this storage policy.
        """
        return pulumi.get(self, "tag_rules")

    @tag_rules.setter
    def tag_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VmStoragePolicyTagRuleArgs']]]]):
        pulumi.set(self, "tag_rules", value)


class VmStoragePolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tag_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VmStoragePolicyTagRuleArgs']]]]] = None,
                 __props__=None):
        """
        The `VmStoragePolicy` resource can be used to create and manage storage
        policies. Using this storage policy, tag based placement rules can be created to
        place a VM on a particular tagged datastore.

        ## Example Usage

        This example creates a storage policy with tag_rule having cat1 as tag_category and
        tag1, tag2 as the tags. While creating a VM, this policy can be referenced to place
        the VM in any of the compatible datastore tagged with these tags.

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        dc = vsphere.get_datacenter(name="DC")
        tag_category = vsphere.get_tag_category(name="cat1")
        tag1 = vsphere.get_tag(name="tag1",
            category_id=tag_category.id)
        tag2 = vsphere.get_tag(name="tag2",
            category_id=tag_category.id)
        policy_tag_based_placement = vsphere.VmStoragePolicy("policyTagBasedPlacement",
            description="description",
            tag_rules=[vsphere.VmStoragePolicyTagRuleArgs(
                tag_category=tag_category.name,
                tags=[
                    tag1.name,
                    tag2.name,
                ],
                include_datastores_with_tags=True,
            )])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the storage policy.
        :param pulumi.Input[str] name: The name of the storage policy.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VmStoragePolicyTagRuleArgs']]]] tag_rules: List of tag rules. The tag category and tags to be associated to this storage policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VmStoragePolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `VmStoragePolicy` resource can be used to create and manage storage
        policies. Using this storage policy, tag based placement rules can be created to
        place a VM on a particular tagged datastore.

        ## Example Usage

        This example creates a storage policy with tag_rule having cat1 as tag_category and
        tag1, tag2 as the tags. While creating a VM, this policy can be referenced to place
        the VM in any of the compatible datastore tagged with these tags.

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        dc = vsphere.get_datacenter(name="DC")
        tag_category = vsphere.get_tag_category(name="cat1")
        tag1 = vsphere.get_tag(name="tag1",
            category_id=tag_category.id)
        tag2 = vsphere.get_tag(name="tag2",
            category_id=tag_category.id)
        policy_tag_based_placement = vsphere.VmStoragePolicy("policyTagBasedPlacement",
            description="description",
            tag_rules=[vsphere.VmStoragePolicyTagRuleArgs(
                tag_category=tag_category.name,
                tags=[
                    tag1.name,
                    tag2.name,
                ],
                include_datastores_with_tags=True,
            )])
        ```

        :param str resource_name: The name of the resource.
        :param VmStoragePolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VmStoragePolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tag_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VmStoragePolicyTagRuleArgs']]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VmStoragePolicyArgs.__new__(VmStoragePolicyArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if tag_rules is None and not opts.urn:
                raise TypeError("Missing required property 'tag_rules'")
            __props__.__dict__["tag_rules"] = tag_rules
        super(VmStoragePolicy, __self__).__init__(
            'vsphere:index/vmStoragePolicy:VmStoragePolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tag_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VmStoragePolicyTagRuleArgs']]]]] = None) -> 'VmStoragePolicy':
        """
        Get an existing VmStoragePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the storage policy.
        :param pulumi.Input[str] name: The name of the storage policy.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VmStoragePolicyTagRuleArgs']]]] tag_rules: List of tag rules. The tag category and tags to be associated to this storage policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VmStoragePolicyState.__new__(_VmStoragePolicyState)

        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["tag_rules"] = tag_rules
        return VmStoragePolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the storage policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the storage policy.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="tagRules")
    def tag_rules(self) -> pulumi.Output[Sequence['outputs.VmStoragePolicyTagRule']]:
        """
        List of tag rules. The tag category and tags to be associated to this storage policy.
        """
        return pulumi.get(self, "tag_rules")


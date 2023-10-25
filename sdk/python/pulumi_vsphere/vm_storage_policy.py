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
        VmStoragePolicyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            tag_rules=tag_rules,
            description=description,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             tag_rules: Optional[pulumi.Input[Sequence[pulumi.Input['VmStoragePolicyTagRuleArgs']]]] = None,
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if tag_rules is None and 'tagRules' in kwargs:
            tag_rules = kwargs['tagRules']
        if tag_rules is None:
            raise TypeError("Missing 'tag_rules' argument")

        _setter("tag_rules", tag_rules)
        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)

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
        _VmStoragePolicyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            name=name,
            tag_rules=tag_rules,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             tag_rules: Optional[pulumi.Input[Sequence[pulumi.Input['VmStoragePolicyTagRuleArgs']]]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if tag_rules is None and 'tagRules' in kwargs:
            tag_rules = kwargs['tagRules']

        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)
        if tag_rules is not None:
            _setter("tag_rules", tag_rules)

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
        policies. Using this resource, tag based placement rules can be created to
        place virtual machines on a datastore with matching tags. If storage requirements for the applications on the virtual machine change, you can modify the storage policy that was originally applied to the virtual machine.

        ## Example Usage

        The following example creates storage policies with `tag_rules` base on sets of environment, service level, and replication attributes.

        In this example, tags are first applied to datastores.

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        environment = vsphere.get_tag_category(name="environment")
        service_level = vsphere.get_tag_category(name="service_level")
        replication = vsphere.get_tag_category(name="replication")
        production = vsphere.get_tag(category_id="data.vsphere_tag_category.environment.id",
            name="production")
        development = vsphere.get_tag(category_id="data.vsphere_tag_category.environment.id",
            name="development")
        platinum = vsphere.get_tag(category_id="data.vsphere_tag_category.service_level.id",
            name="platinum")
        gold = vsphere.get_tag(category_id="data.vsphere_tag_category.service_level.id",
            name="platinum")
        silver = vsphere.get_tag(category_id="data.vsphere_tag_category.service_level.id",
            name="silver")
        bronze = vsphere.get_tag(category_id="data.vsphere_tag_category.service_level.id",
            name="bronze")
        replicated = vsphere.get_tag(category_id="data.vsphere_tag_category.replication.id",
            name="replicated")
        non_replicated = vsphere.get_tag(category_id="data.vsphere_tag_category.replication.id",
            name="non_replicated")
        prod_datastore = vsphere.VmfsDatastore("prodDatastore", tags=[
            "data.vsphere_tag.production.id",
            "data.vsphere_tag.platinum.id",
            "data.vsphere_tag.replicated.id",
        ])
        dev_datastore = vsphere.NasDatastore("devDatastore", tags=[
            "data.vsphere_tag.development.id",
            "data.vsphere_tag.silver.id",
            "data.vsphere_tag.non_replicated.id",
        ])
        ```

        Next, storage policies are created and `tag_rules` are applied.

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        prod_platinum_replicated = vsphere.VmStoragePolicy("prodPlatinumReplicated",
            description="prod_platinum_replicated",
            tag_rules=[
                vsphere.VmStoragePolicyTagRuleArgs(
                    tag_category=data["vsphere_tag_category"]["environment"]["name"],
                    tags=[data["vsphere_tag"]["production"]["name"]],
                    include_datastores_with_tags=True,
                ),
                vsphere.VmStoragePolicyTagRuleArgs(
                    tag_category=data["vsphere_tag_category"]["service_level"]["name"],
                    tags=[data["vsphere_tag"]["platinum"]["name"]],
                    include_datastores_with_tags=True,
                ),
                vsphere.VmStoragePolicyTagRuleArgs(
                    tag_category=data["vsphere_tag_category"]["replication"]["name"],
                    tags=[data["vsphere_tag"]["replicated"]["name"]],
                    include_datastores_with_tags=True,
                ),
            ])
        dev_silver_nonreplicated = vsphere.VmStoragePolicy("devSilverNonreplicated",
            description="dev_silver_nonreplicated",
            tag_rules=[
                vsphere.VmStoragePolicyTagRuleArgs(
                    tag_category=data["vsphere_tag_category"]["environment"]["name"],
                    tags=[data["vsphere_tag"]["development"]["name"]],
                    include_datastores_with_tags=True,
                ),
                vsphere.VmStoragePolicyTagRuleArgs(
                    tag_category=data["vsphere_tag_category"]["service_level"]["name"],
                    tags=[data["vsphere_tag"]["silver"]["name"]],
                    include_datastores_with_tags=True,
                ),
                vsphere.VmStoragePolicyTagRuleArgs(
                    tag_category=data["vsphere_tag_category"]["replication"]["name"],
                    tags=[data["vsphere_tag"]["non_replicated"]["name"]],
                    include_datastores_with_tags=True,
                ),
            ])
        ```

        Lasttly, when creating a virtual machine resource, a storage policy can be specificed to direct virtual machine placement to a datastore which matches the policy's `tags_rules`.

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        prod_platinum_replicated = vsphere.get_policy(name="prod_platinum_replicated")
        dev_silver_nonreplicated = vsphere.get_policy(name="dev_silver_nonreplicated")
        prod_vm = vsphere.VirtualMachine("prodVm", storage_policy_id=data["vsphere_storage_policy"]["storage_policy"]["prod_platinum_replicated"]["id"])
        # ... other configuration ...
        dev_vm = vsphere.VirtualMachine("devVm", storage_policy_id=data["vsphere_storage_policy"]["storage_policy"]["dev_silver_nonreplicated"]["id"])
        # ... other configuration ...
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
        policies. Using this resource, tag based placement rules can be created to
        place virtual machines on a datastore with matching tags. If storage requirements for the applications on the virtual machine change, you can modify the storage policy that was originally applied to the virtual machine.

        ## Example Usage

        The following example creates storage policies with `tag_rules` base on sets of environment, service level, and replication attributes.

        In this example, tags are first applied to datastores.

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        environment = vsphere.get_tag_category(name="environment")
        service_level = vsphere.get_tag_category(name="service_level")
        replication = vsphere.get_tag_category(name="replication")
        production = vsphere.get_tag(category_id="data.vsphere_tag_category.environment.id",
            name="production")
        development = vsphere.get_tag(category_id="data.vsphere_tag_category.environment.id",
            name="development")
        platinum = vsphere.get_tag(category_id="data.vsphere_tag_category.service_level.id",
            name="platinum")
        gold = vsphere.get_tag(category_id="data.vsphere_tag_category.service_level.id",
            name="platinum")
        silver = vsphere.get_tag(category_id="data.vsphere_tag_category.service_level.id",
            name="silver")
        bronze = vsphere.get_tag(category_id="data.vsphere_tag_category.service_level.id",
            name="bronze")
        replicated = vsphere.get_tag(category_id="data.vsphere_tag_category.replication.id",
            name="replicated")
        non_replicated = vsphere.get_tag(category_id="data.vsphere_tag_category.replication.id",
            name="non_replicated")
        prod_datastore = vsphere.VmfsDatastore("prodDatastore", tags=[
            "data.vsphere_tag.production.id",
            "data.vsphere_tag.platinum.id",
            "data.vsphere_tag.replicated.id",
        ])
        dev_datastore = vsphere.NasDatastore("devDatastore", tags=[
            "data.vsphere_tag.development.id",
            "data.vsphere_tag.silver.id",
            "data.vsphere_tag.non_replicated.id",
        ])
        ```

        Next, storage policies are created and `tag_rules` are applied.

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        prod_platinum_replicated = vsphere.VmStoragePolicy("prodPlatinumReplicated",
            description="prod_platinum_replicated",
            tag_rules=[
                vsphere.VmStoragePolicyTagRuleArgs(
                    tag_category=data["vsphere_tag_category"]["environment"]["name"],
                    tags=[data["vsphere_tag"]["production"]["name"]],
                    include_datastores_with_tags=True,
                ),
                vsphere.VmStoragePolicyTagRuleArgs(
                    tag_category=data["vsphere_tag_category"]["service_level"]["name"],
                    tags=[data["vsphere_tag"]["platinum"]["name"]],
                    include_datastores_with_tags=True,
                ),
                vsphere.VmStoragePolicyTagRuleArgs(
                    tag_category=data["vsphere_tag_category"]["replication"]["name"],
                    tags=[data["vsphere_tag"]["replicated"]["name"]],
                    include_datastores_with_tags=True,
                ),
            ])
        dev_silver_nonreplicated = vsphere.VmStoragePolicy("devSilverNonreplicated",
            description="dev_silver_nonreplicated",
            tag_rules=[
                vsphere.VmStoragePolicyTagRuleArgs(
                    tag_category=data["vsphere_tag_category"]["environment"]["name"],
                    tags=[data["vsphere_tag"]["development"]["name"]],
                    include_datastores_with_tags=True,
                ),
                vsphere.VmStoragePolicyTagRuleArgs(
                    tag_category=data["vsphere_tag_category"]["service_level"]["name"],
                    tags=[data["vsphere_tag"]["silver"]["name"]],
                    include_datastores_with_tags=True,
                ),
                vsphere.VmStoragePolicyTagRuleArgs(
                    tag_category=data["vsphere_tag_category"]["replication"]["name"],
                    tags=[data["vsphere_tag"]["non_replicated"]["name"]],
                    include_datastores_with_tags=True,
                ),
            ])
        ```

        Lasttly, when creating a virtual machine resource, a storage policy can be specificed to direct virtual machine placement to a datastore which matches the policy's `tags_rules`.

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        prod_platinum_replicated = vsphere.get_policy(name="prod_platinum_replicated")
        dev_silver_nonreplicated = vsphere.get_policy(name="dev_silver_nonreplicated")
        prod_vm = vsphere.VirtualMachine("prodVm", storage_policy_id=data["vsphere_storage_policy"]["storage_policy"]["prod_platinum_replicated"]["id"])
        # ... other configuration ...
        dev_vm = vsphere.VirtualMachine("devVm", storage_policy_id=data["vsphere_storage_policy"]["storage_policy"]["dev_silver_nonreplicated"]["id"])
        # ... other configuration ...
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
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            VmStoragePolicyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tag_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VmStoragePolicyTagRuleArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
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


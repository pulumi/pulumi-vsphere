# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TagCategoryArgs', 'TagCategory']

@pulumi.input_type
class TagCategoryArgs:
    def __init__(__self__, *,
                 associable_types: pulumi.Input[Sequence[pulumi.Input[str]]],
                 cardinality: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TagCategory resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] associable_types: A list object types that this category is
               valid to be assigned to. For a full list, click
               here.
        :param pulumi.Input[str] cardinality: The number of tags that can be assigned from this
               category to a single object at once. Can be one of `SINGLE` (object can only
               be assigned one tag in this category), to `MULTIPLE` (object can be assigned
               multiple tags in this category). Forces a new resource if changed.
        :param pulumi.Input[str] description: A description for the category.
               
               > **NOTE:** You can add associable types to a category, but you cannot remove
               them. Attempting to do so will result in an error.
        :param pulumi.Input[str] name: The name of the category.
        """
        TagCategoryArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            associable_types=associable_types,
            cardinality=cardinality,
            description=description,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             associable_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             cardinality: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if associable_types is None and 'associableTypes' in kwargs:
            associable_types = kwargs['associableTypes']
        if associable_types is None:
            raise TypeError("Missing 'associable_types' argument")
        if cardinality is None:
            raise TypeError("Missing 'cardinality' argument")

        _setter("associable_types", associable_types)
        _setter("cardinality", cardinality)
        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter(name="associableTypes")
    def associable_types(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list object types that this category is
        valid to be assigned to. For a full list, click
        here.
        """
        return pulumi.get(self, "associable_types")

    @associable_types.setter
    def associable_types(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "associable_types", value)

    @property
    @pulumi.getter
    def cardinality(self) -> pulumi.Input[str]:
        """
        The number of tags that can be assigned from this
        category to a single object at once. Can be one of `SINGLE` (object can only
        be assigned one tag in this category), to `MULTIPLE` (object can be assigned
        multiple tags in this category). Forces a new resource if changed.
        """
        return pulumi.get(self, "cardinality")

    @cardinality.setter
    def cardinality(self, value: pulumi.Input[str]):
        pulumi.set(self, "cardinality", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the category.

        > **NOTE:** You can add associable types to a category, but you cannot remove
        them. Attempting to do so will result in an error.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the category.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _TagCategoryState:
    def __init__(__self__, *,
                 associable_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cardinality: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TagCategory resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] associable_types: A list object types that this category is
               valid to be assigned to. For a full list, click
               here.
        :param pulumi.Input[str] cardinality: The number of tags that can be assigned from this
               category to a single object at once. Can be one of `SINGLE` (object can only
               be assigned one tag in this category), to `MULTIPLE` (object can be assigned
               multiple tags in this category). Forces a new resource if changed.
        :param pulumi.Input[str] description: A description for the category.
               
               > **NOTE:** You can add associable types to a category, but you cannot remove
               them. Attempting to do so will result in an error.
        :param pulumi.Input[str] name: The name of the category.
        """
        _TagCategoryState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            associable_types=associable_types,
            cardinality=cardinality,
            description=description,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             associable_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             cardinality: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if associable_types is None and 'associableTypes' in kwargs:
            associable_types = kwargs['associableTypes']

        if associable_types is not None:
            _setter("associable_types", associable_types)
        if cardinality is not None:
            _setter("cardinality", cardinality)
        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter(name="associableTypes")
    def associable_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list object types that this category is
        valid to be assigned to. For a full list, click
        here.
        """
        return pulumi.get(self, "associable_types")

    @associable_types.setter
    def associable_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "associable_types", value)

    @property
    @pulumi.getter
    def cardinality(self) -> Optional[pulumi.Input[str]]:
        """
        The number of tags that can be assigned from this
        category to a single object at once. Can be one of `SINGLE` (object can only
        be assigned one tag in this category), to `MULTIPLE` (object can be assigned
        multiple tags in this category). Forces a new resource if changed.
        """
        return pulumi.get(self, "cardinality")

    @cardinality.setter
    def cardinality(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cardinality", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the category.

        > **NOTE:** You can add associable types to a category, but you cannot remove
        them. Attempting to do so will result in an error.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the category.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class TagCategory(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 associable_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cardinality: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a TagCategory resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] associable_types: A list object types that this category is
               valid to be assigned to. For a full list, click
               here.
        :param pulumi.Input[str] cardinality: The number of tags that can be assigned from this
               category to a single object at once. Can be one of `SINGLE` (object can only
               be assigned one tag in this category), to `MULTIPLE` (object can be assigned
               multiple tags in this category). Forces a new resource if changed.
        :param pulumi.Input[str] description: A description for the category.
               
               > **NOTE:** You can add associable types to a category, but you cannot remove
               them. Attempting to do so will result in an error.
        :param pulumi.Input[str] name: The name of the category.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TagCategoryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a TagCategory resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param TagCategoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TagCategoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            TagCategoryArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 associable_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cardinality: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TagCategoryArgs.__new__(TagCategoryArgs)

            if associable_types is None and not opts.urn:
                raise TypeError("Missing required property 'associable_types'")
            __props__.__dict__["associable_types"] = associable_types
            if cardinality is None and not opts.urn:
                raise TypeError("Missing required property 'cardinality'")
            __props__.__dict__["cardinality"] = cardinality
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
        super(TagCategory, __self__).__init__(
            'vsphere:index/tagCategory:TagCategory',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            associable_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            cardinality: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'TagCategory':
        """
        Get an existing TagCategory resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] associable_types: A list object types that this category is
               valid to be assigned to. For a full list, click
               here.
        :param pulumi.Input[str] cardinality: The number of tags that can be assigned from this
               category to a single object at once. Can be one of `SINGLE` (object can only
               be assigned one tag in this category), to `MULTIPLE` (object can be assigned
               multiple tags in this category). Forces a new resource if changed.
        :param pulumi.Input[str] description: A description for the category.
               
               > **NOTE:** You can add associable types to a category, but you cannot remove
               them. Attempting to do so will result in an error.
        :param pulumi.Input[str] name: The name of the category.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TagCategoryState.__new__(_TagCategoryState)

        __props__.__dict__["associable_types"] = associable_types
        __props__.__dict__["cardinality"] = cardinality
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        return TagCategory(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="associableTypes")
    def associable_types(self) -> pulumi.Output[Sequence[str]]:
        """
        A list object types that this category is
        valid to be assigned to. For a full list, click
        here.
        """
        return pulumi.get(self, "associable_types")

    @property
    @pulumi.getter
    def cardinality(self) -> pulumi.Output[str]:
        """
        The number of tags that can be assigned from this
        category to a single object at once. Can be one of `SINGLE` (object can only
        be assigned one tag in this category), to `MULTIPLE` (object can be assigned
        multiple tags in this category). Forces a new resource if changed.
        """
        return pulumi.get(self, "cardinality")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description for the category.

        > **NOTE:** You can add associable types to a category, but you cannot remove
        them. Attempting to do so will result in an error.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the category.
        """
        return pulumi.get(self, "name")


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class GetTagCategoryResult:
    """
    A collection of values returned by getTagCategory.
    """
    def __init__(__self__, associable_types=None, cardinality=None, description=None, id=None):
        if associable_types and not isinstance(associable_types, list):
            raise TypeError('Expected argument associable_types to be a list')
        __self__.associable_types = associable_types
        if cardinality and not isinstance(cardinality, str):
            raise TypeError('Expected argument cardinality to be a str')
        __self__.cardinality = cardinality
        if description and not isinstance(description, str):
            raise TypeError('Expected argument description to be a str')
        __self__.description = description
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_tag_category(name=None,opts=None):
    """
    The `vsphere_tag_category` data source can be used to reference tag categories
    that are not managed by Terraform. Its attributes are exactly the same as the
    [`vsphere_tag_category` resource][resource-tag-category], and, like importing,
    the data source takes a name to search on. The `id` and other attributes are
    then populated with the data found by the search.
    
    [resource-tag-category]: /docs/providers/vsphere/r/tag_category.html
    
    > **NOTE:** Tagging support is unsupported on direct ESXi connections and
    requires vCenter 6.0 or higher.
    """
    __args__ = dict()

    __args__['name'] = name
    __ret__ = await pulumi.runtime.invoke('vsphere:index/getTagCategory:getTagCategory', __args__, opts=opts)

    return GetTagCategoryResult(
        associable_types=__ret__.get('associableTypes'),
        cardinality=__ret__.get('cardinality'),
        description=__ret__.get('description'),
        id=__ret__.get('id'))

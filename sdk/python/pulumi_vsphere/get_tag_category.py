# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetTagCategoryResult:
    """
    A collection of values returned by getTagCategory.
    """
    def __init__(__self__, associable_types=None, cardinality=None, description=None, id=None, name=None):
        if associable_types and not isinstance(associable_types, list):
            raise TypeError("Expected argument 'associable_types' to be a list")
        __self__.associable_types = associable_types
        if cardinality and not isinstance(cardinality, str):
            raise TypeError("Expected argument 'cardinality' to be a str")
        __self__.cardinality = cardinality
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
class AwaitableGetTagCategoryResult(GetTagCategoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTagCategoryResult(
            associable_types=self.associable_types,
            cardinality=self.cardinality,
            description=self.description,
            id=self.id,
            name=self.name)

def get_tag_category(name=None,opts=None):
    """
    The `.TagCategory` data source can be used to reference tag categories
    that are not managed by this provider. Its attributes are exactly the same as the
    `.TagCategory` resource, and, like importing,
    the data source takes a name to search on. The `id` and other attributes are
    then populated with the data found by the search.

    > **NOTE:** Tagging support is unsupported on direct ESXi connections and
    requires vCenter 6.0 or higher.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    category = vsphere.get_tag_category(name="test-category")
    ```



    :param str name: The name of the tag category.
    """
    __args__ = dict()


    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vsphere:index/getTagCategory:getTagCategory', __args__, opts=opts).value

    return AwaitableGetTagCategoryResult(
        associable_types=__ret__.get('associableTypes'),
        cardinality=__ret__.get('cardinality'),
        description=__ret__.get('description'),
        id=__ret__.get('id'),
        name=__ret__.get('name'))

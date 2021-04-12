# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables

__all__ = [
    'GetTagResult',
    'AwaitableGetTagResult',
    'get_tag',
]

@pulumi.output_type
class GetTagResult:
    """
    A collection of values returned by getTag.
    """
    def __init__(__self__, category_id=None, description=None, id=None, name=None):
        if category_id and not isinstance(category_id, str):
            raise TypeError("Expected argument 'category_id' to be a str")
        pulumi.set(__self__, "category_id", category_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="categoryId")
    def category_id(self) -> str:
        return pulumi.get(self, "category_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


class AwaitableGetTagResult(GetTagResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTagResult(
            category_id=self.category_id,
            description=self.description,
            id=self.id,
            name=self.name)


def get_tag(category_id: Optional[str] = None,
            name: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTagResult:
    """
    The `Tag` data source can be used to reference tags that are not
    managed by this provider. Its attributes are exactly the same as the `Tag`
    resource, and, like importing, the data source takes a name and
    category to search on. The `id` and other attributes are then populated with
    the data found by the search.

    > **NOTE:** Tagging support is unsupported on direct ESXi connections and
    requires vCenter 6.0 or higher.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    category = vsphere.get_tag_category(name="test-category")
    tag = vsphere.get_tag(category_id=category.id,
        name="test-tag")
    ```


    :param str category_id: The ID of the tag category the tag is located in.
    :param str name: The name of the tag.
    """
    __args__ = dict()
    __args__['categoryId'] = category_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vsphere:index/getTag:getTag', __args__, opts=opts, typ=GetTagResult).value

    return AwaitableGetTagResult(
        category_id=__ret__.category_id,
        description=__ret__.description,
        id=__ret__.id,
        name=__ret__.name)

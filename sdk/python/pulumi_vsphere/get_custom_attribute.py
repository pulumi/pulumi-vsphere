# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables

__all__ = [
    'GetCustomAttributeResult',
    'AwaitableGetCustomAttributeResult',
    'get_custom_attribute',
]

@pulumi.output_type
class GetCustomAttributeResult:
    """
    A collection of values returned by getCustomAttribute.
    """
    def __init__(__self__, id=None, managed_object_type=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if managed_object_type and not isinstance(managed_object_type, str):
            raise TypeError("Expected argument 'managed_object_type' to be a str")
        pulumi.set(__self__, "managed_object_type", managed_object_type)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="managedObjectType")
    def managed_object_type(self) -> str:
        return pulumi.get(self, "managed_object_type")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


class AwaitableGetCustomAttributeResult(GetCustomAttributeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCustomAttributeResult(
            id=self.id,
            managed_object_type=self.managed_object_type,
            name=self.name)


def get_custom_attribute(name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCustomAttributeResult:
    """
    The `CustomAttribute` data source can be used to reference custom
    attributes that are not managed by this provider. Its attributes are exactly the
    same as the `CustomAttribute` resource,
    and, like importing, the data source takes a name to search on. The `id` and
    other attributes are then populated with the data found by the search.

    > **NOTE:** Custom attributes are unsupported on direct ESXi connections
    and require vCenter.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    attribute = vsphere.get_custom_attribute(name="test-attribute")
    ```


    :param str name: The name of the custom attribute.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vsphere:index/getCustomAttribute:getCustomAttribute', __args__, opts=opts, typ=GetCustomAttributeResult).value

    return AwaitableGetCustomAttributeResult(
        id=__ret__.id,
        managed_object_type=__ret__.managed_object_type,
        name=__ret__.name)

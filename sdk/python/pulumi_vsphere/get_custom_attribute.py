# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class GetCustomAttributeResult:
    """
    A collection of values returned by getCustomAttribute.
    """
    def __init__(__self__, managed_object_type=None, name=None, id=None):
        if managed_object_type and not isinstance(managed_object_type, str):
            raise TypeError("Expected argument 'managed_object_type' to be a str")
        __self__.managed_object_type = managed_object_type
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetCustomAttributeResult(GetCustomAttributeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCustomAttributeResult(
            managed_object_type=self.managed_object_type,
            name=self.name,
            id=self.id)

def get_custom_attribute(name=None,opts=None):
    """
    > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/custom_attribute.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vsphere:index/getCustomAttribute:getCustomAttribute', __args__, opts=opts).value

    return AwaitableGetCustomAttributeResult(
        managed_object_type=__ret__.get('managedObjectType'),
        name=__ret__.get('name'),
        id=__ret__.get('id'))

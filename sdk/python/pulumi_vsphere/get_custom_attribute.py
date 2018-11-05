# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from . import utilities

class GetCustomAttributeResult(object):
    """
    A collection of values returned by getCustomAttribute.
    """
    def __init__(__self__, managed_object_type=None, id=None):
        if managed_object_type and not isinstance(managed_object_type, str):
            raise TypeError('Expected argument managed_object_type to be a str')
        __self__.managed_object_type = managed_object_type
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_custom_attribute(name=None):
    """
    The `vsphere_custom_attribute` data source can be used to reference custom 
    attributes that are not managed by Terraform. Its attributes are exactly the 
    same as the [`vsphere_custom_attribute` resource][resource-custom-attribute], 
    and, like importing, the data source takes a name to search on. The `id` and 
    other attributes are then populated with the data found by the search.
    
    [resource-custom-attribute]: /docs/providers/vsphere/r/custom_attribute.html
    
    ~> **NOTE:** Custom attributes are unsupported on direct ESXi connections 
    and require vCenter.
    """
    __args__ = dict()

    __args__['name'] = name
    __ret__ = pulumi.runtime.invoke('vsphere:index/getCustomAttribute:getCustomAttribute', __args__)

    return GetCustomAttributeResult(
        managed_object_type=__ret__.get('managedObjectType'),
        id=__ret__.get('id'))

# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetDynamicResult:
    """
    A collection of values returned by getDynamic.
    """
    def __init__(__self__, filters=None, id=None, name_regex=None, type=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
class AwaitableGetDynamicResult(GetDynamicResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDynamicResult(
            filters=self.filters,
            id=self.id,
            name_regex=self.name_regex,
            type=self.type)

def get_dynamic(filters=None,name_regex=None,type=None,opts=None):
    """
    [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider

    The `.getDynamic` data source can be used to get the [managed object 
      reference ID][docs-about-morefs] of any tagged managed object in vCenter
      by providing a list of tag IDs and an optional regular expression to filter
      objects by name.
       



    :param list filters: A list of tag IDs that must be present on an object to
           be a match.
    :param str name_regex: A regular expression that will be used to match
           the object's name.
    :param str type: The managed object type the returned object must match.
           For a full list, click [here](https://code.vmware.com/apis/196/vsphere).
    """
    __args__ = dict()


    __args__['filters'] = filters
    __args__['nameRegex'] = name_regex
    __args__['type'] = type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vsphere:index/getDynamic:getDynamic', __args__, opts=opts).value

    return AwaitableGetDynamicResult(
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        name_regex=__ret__.get('nameRegex'),
        type=__ret__.get('type'))

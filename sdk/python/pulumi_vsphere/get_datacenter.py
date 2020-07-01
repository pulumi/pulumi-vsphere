# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetDatacenterResult:
    """
    A collection of values returned by getDatacenter.
    """
    def __init__(__self__, id=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
class AwaitableGetDatacenterResult(GetDatacenterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatacenterResult(
            id=self.id,
            name=self.name)

def get_datacenter(name=None,opts=None):
    """
    The `Datacenter` data source can be used to discover the ID of a
    vSphere datacenter. This can then be used with resources or data sources that
    require a datacenter, such as the `Host`
    data source.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc1")
    ```


    :param str name: The name of the datacenter. This can be a name or path.
           Can be omitted if there is only one datacenter in your inventory.
    """
    __args__ = dict()


    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vsphere:index/getDatacenter:getDatacenter', __args__, opts=opts).value

    return AwaitableGetDatacenterResult(
        id=__ret__.get('id'),
        name=__ret__.get('name'))

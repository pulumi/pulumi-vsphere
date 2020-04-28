# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetContentLibraryItemResult:
    """
    A collection of values returned by getContentLibraryItem.
    """
    def __init__(__self__, id=None, library_id=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if library_id and not isinstance(library_id, str):
            raise TypeError("Expected argument 'library_id' to be a str")
        __self__.library_id = library_id
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
class AwaitableGetContentLibraryItemResult(GetContentLibraryItemResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContentLibraryItemResult(
            id=self.id,
            library_id=self.library_id,
            name=self.name)

def get_content_library_item(library_id=None,name=None,opts=None):
    """
    The `.ContentLibraryItem` data source can be used to discover the ID of a Content Library item.

    > **NOTE:** This resource requires vCenter and is not available on direct ESXi
    connections.




    :param str library_id: The ID of the Content Library the item exists in.
    :param str name: The name of the Content Library.
    """
    __args__ = dict()


    __args__['libraryId'] = library_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vsphere:index/getContentLibraryItem:getContentLibraryItem', __args__, opts=opts).value

    return AwaitableGetContentLibraryItemResult(
        id=__ret__.get('id'),
        library_id=__ret__.get('libraryId'),
        name=__ret__.get('name'))

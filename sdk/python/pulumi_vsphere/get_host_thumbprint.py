# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetHostThumbprintResult',
    'AwaitableGetHostThumbprintResult',
    'get_host_thumbprint',
    'get_host_thumbprint_output',
]

@pulumi.output_type
class GetHostThumbprintResult:
    """
    A collection of values returned by getHostThumbprint.
    """
    def __init__(__self__, address=None, id=None, insecure=None, port=None):
        if address and not isinstance(address, str):
            raise TypeError("Expected argument 'address' to be a str")
        pulumi.set(__self__, "address", address)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if insecure and not isinstance(insecure, bool):
            raise TypeError("Expected argument 'insecure' to be a bool")
        pulumi.set(__self__, "insecure", insecure)
        if port and not isinstance(port, str):
            raise TypeError("Expected argument 'port' to be a str")
        pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def address(self) -> str:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def insecure(self) -> Optional[bool]:
        return pulumi.get(self, "insecure")

    @property
    @pulumi.getter
    def port(self) -> Optional[str]:
        return pulumi.get(self, "port")


class AwaitableGetHostThumbprintResult(GetHostThumbprintResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHostThumbprintResult(
            address=self.address,
            id=self.id,
            insecure=self.insecure,
            port=self.port)


def get_host_thumbprint(address: Optional[str] = None,
                        insecure: Optional[bool] = None,
                        port: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHostThumbprintResult:
    """
    The `vsphere_thumbprint` data source can be used to discover the host
    thumbprint of an ESXi host. This can be used when adding the `Host`
    resource. If the ESXi host is using a certificate chain, the first one returned
    will be used to generate the thumbprint.


    :param str address: The address of the ESXi host to retrieve the
           thumbprint from.
    :param bool insecure: Disables SSL certificate verification.
           Default: `false`
    :param str port: The port to use connecting to the ESXi host. Default: 443
    """
    __args__ = dict()
    __args__['address'] = address
    __args__['insecure'] = insecure
    __args__['port'] = port
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vsphere:index/getHostThumbprint:getHostThumbprint', __args__, opts=opts, typ=GetHostThumbprintResult).value

    return AwaitableGetHostThumbprintResult(
        address=pulumi.get(__ret__, 'address'),
        id=pulumi.get(__ret__, 'id'),
        insecure=pulumi.get(__ret__, 'insecure'),
        port=pulumi.get(__ret__, 'port'))


@_utilities.lift_output_func(get_host_thumbprint)
def get_host_thumbprint_output(address: Optional[pulumi.Input[str]] = None,
                               insecure: Optional[pulumi.Input[Optional[bool]]] = None,
                               port: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHostThumbprintResult]:
    """
    The `vsphere_thumbprint` data source can be used to discover the host
    thumbprint of an ESXi host. This can be used when adding the `Host`
    resource. If the ESXi host is using a certificate chain, the first one returned
    will be used to generate the thumbprint.


    :param str address: The address of the ESXi host to retrieve the
           thumbprint from.
    :param bool insecure: Disables SSL certificate verification.
           Default: `false`
    :param str port: The port to use connecting to the ESXi host. Default: 443
    """
    ...

# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs

__all__ = [
    'GetHostVgpuProfileResult',
    'AwaitableGetHostVgpuProfileResult',
    'get_host_vgpu_profile',
    'get_host_vgpu_profile_output',
]

@pulumi.output_type
class GetHostVgpuProfileResult:
    """
    A collection of values returned by getHostVgpuProfile.
    """
    def __init__(__self__, host_id=None, id=None, name_regex=None, vgpu_profiles=None):
        if host_id and not isinstance(host_id, str):
            raise TypeError("Expected argument 'host_id' to be a str")
        pulumi.set(__self__, "host_id", host_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if vgpu_profiles and not isinstance(vgpu_profiles, list):
            raise TypeError("Expected argument 'vgpu_profiles' to be a list")
        pulumi.set(__self__, "vgpu_profiles", vgpu_profiles)

    @property
    @pulumi.getter(name="hostId")
    def host_id(self) -> str:
        """
        The [managed objectID][docs-about-morefs] of the ESXi host.
        """
        return pulumi.get(self, "host_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        """
        (Optional) A regular expression that will be used to match the
        host vGPU profile name.
        """
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="vgpuProfiles")
    def vgpu_profiles(self) -> Sequence['outputs.GetHostVgpuProfileVgpuProfileResult']:
        """
        The list of available vGPU profiles on the ESXi host.
        This may be and empty array if no vGPU profile are identified.
        """
        return pulumi.get(self, "vgpu_profiles")


class AwaitableGetHostVgpuProfileResult(GetHostVgpuProfileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHostVgpuProfileResult(
            host_id=self.host_id,
            id=self.id,
            name_regex=self.name_regex,
            vgpu_profiles=self.vgpu_profiles)


def get_host_vgpu_profile(host_id: Optional[str] = None,
                          name_regex: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHostVgpuProfileResult:
    """
    The `get_host_vgpu_profile` data source can be used to discover the
    available vGPU profiles of a vSphere host.

    ## Example Usage

    ### To Return All VGPU Profiles

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    host = vsphere.get_host(name="esxi-01.example.com",
        datacenter_id=datacenter.id)
    vgpu_profile = vsphere.get_host_vgpu_profile(host_id=host.id)
    ```

    ### With VGPU Profile Name_regex

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    host = vsphere.get_host(name="esxi-01.example.com",
        datacenter_id=datacenter.id)
    vgpu_profile = vsphere.get_host_vgpu_profile(host_id=host.id,
        name_regex="a100")
    ```


    :param str host_id: The [managed object reference ID][docs-about-morefs] of
           a host.
    :param str name_regex: A regular expression that will be used to match the
           host vGPU profile name.
           
           [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
    """
    __args__ = dict()
    __args__['hostId'] = host_id
    __args__['nameRegex'] = name_regex
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vsphere:index/getHostVgpuProfile:getHostVgpuProfile', __args__, opts=opts, typ=GetHostVgpuProfileResult).value

    return AwaitableGetHostVgpuProfileResult(
        host_id=pulumi.get(__ret__, 'host_id'),
        id=pulumi.get(__ret__, 'id'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        vgpu_profiles=pulumi.get(__ret__, 'vgpu_profiles'))
def get_host_vgpu_profile_output(host_id: Optional[pulumi.Input[str]] = None,
                                 name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetHostVgpuProfileResult]:
    """
    The `get_host_vgpu_profile` data source can be used to discover the
    available vGPU profiles of a vSphere host.

    ## Example Usage

    ### To Return All VGPU Profiles

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    host = vsphere.get_host(name="esxi-01.example.com",
        datacenter_id=datacenter.id)
    vgpu_profile = vsphere.get_host_vgpu_profile(host_id=host.id)
    ```

    ### With VGPU Profile Name_regex

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    host = vsphere.get_host(name="esxi-01.example.com",
        datacenter_id=datacenter.id)
    vgpu_profile = vsphere.get_host_vgpu_profile(host_id=host.id,
        name_regex="a100")
    ```


    :param str host_id: The [managed object reference ID][docs-about-morefs] of
           a host.
    :param str name_regex: A regular expression that will be used to match the
           host vGPU profile name.
           
           [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
    """
    __args__ = dict()
    __args__['hostId'] = host_id
    __args__['nameRegex'] = name_regex
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vsphere:index/getHostVgpuProfile:getHostVgpuProfile', __args__, opts=opts, typ=GetHostVgpuProfileResult)
    return __ret__.apply(lambda __response__: GetHostVgpuProfileResult(
        host_id=pulumi.get(__response__, 'host_id'),
        id=pulumi.get(__response__, 'id'),
        name_regex=pulumi.get(__response__, 'name_regex'),
        vgpu_profiles=pulumi.get(__response__, 'vgpu_profiles')))

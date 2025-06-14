# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = [
    'GetLicenseResult',
    'AwaitableGetLicenseResult',
    'get_license',
    'get_license_output',
]

@pulumi.output_type
class GetLicenseResult:
    """
    A collection of values returned by getLicense.
    """
    def __init__(__self__, edition_key=None, id=None, labels=None, license_key=None, name=None, total=None, used=None):
        if edition_key and not isinstance(edition_key, str):
            raise TypeError("Expected argument 'edition_key' to be a str")
        pulumi.set(__self__, "edition_key", edition_key)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if license_key and not isinstance(license_key, str):
            raise TypeError("Expected argument 'license_key' to be a str")
        pulumi.set(__self__, "license_key", license_key)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if total and not isinstance(total, int):
            raise TypeError("Expected argument 'total' to be a int")
        pulumi.set(__self__, "total", total)
        if used and not isinstance(used, int):
            raise TypeError("Expected argument 'used' to be a int")
        pulumi.set(__self__, "used", used)

    @property
    @pulumi.getter(name="editionKey")
    def edition_key(self) -> builtins.str:
        """
        The product edition of the license key.
        """
        return pulumi.get(self, "edition_key")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The license key ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, builtins.str]:
        """
        A map of labels applied to the license key.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="licenseKey")
    def license_key(self) -> builtins.str:
        return pulumi.get(self, "license_key")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The display name for the license.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def total(self) -> builtins.int:
        """
        The total number of units contained in the license key.
        """
        return pulumi.get(self, "total")

    @property
    @pulumi.getter
    def used(self) -> builtins.int:
        """
        The number of units assigned to this license key.
        """
        return pulumi.get(self, "used")


class AwaitableGetLicenseResult(GetLicenseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLicenseResult(
            edition_key=self.edition_key,
            id=self.id,
            labels=self.labels,
            license_key=self.license_key,
            name=self.name,
            total=self.total,
            used=self.used)


def get_license(license_key: Optional[builtins.str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLicenseResult:
    """
    The `License` data source can be used to get the general attributes of
    a license keys from a vCenter Server instance.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    license = vsphere.get_license(license_key="00000-00000-00000-00000-00000")
    ```


    :param builtins.str license_key: The license key value.
    """
    __args__ = dict()
    __args__['licenseKey'] = license_key
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vsphere:index/getLicense:getLicense', __args__, opts=opts, typ=GetLicenseResult).value

    return AwaitableGetLicenseResult(
        edition_key=pulumi.get(__ret__, 'edition_key'),
        id=pulumi.get(__ret__, 'id'),
        labels=pulumi.get(__ret__, 'labels'),
        license_key=pulumi.get(__ret__, 'license_key'),
        name=pulumi.get(__ret__, 'name'),
        total=pulumi.get(__ret__, 'total'),
        used=pulumi.get(__ret__, 'used'))
def get_license_output(license_key: Optional[pulumi.Input[builtins.str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetLicenseResult]:
    """
    The `License` data source can be used to get the general attributes of
    a license keys from a vCenter Server instance.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    license = vsphere.get_license(license_key="00000-00000-00000-00000-00000")
    ```


    :param builtins.str license_key: The license key value.
    """
    __args__ = dict()
    __args__['licenseKey'] = license_key
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vsphere:index/getLicense:getLicense', __args__, opts=opts, typ=GetLicenseResult)
    return __ret__.apply(lambda __response__: GetLicenseResult(
        edition_key=pulumi.get(__response__, 'edition_key'),
        id=pulumi.get(__response__, 'id'),
        labels=pulumi.get(__response__, 'labels'),
        license_key=pulumi.get(__response__, 'license_key'),
        name=pulumi.get(__response__, 'name'),
        total=pulumi.get(__response__, 'total'),
        used=pulumi.get(__response__, 'used')))

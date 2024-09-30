# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['VnicArgs', 'Vnic']

@pulumi.input_type
class VnicArgs:
    def __init__(__self__, *,
                 host: pulumi.Input[str],
                 distributed_port_group: Optional[pulumi.Input[str]] = None,
                 distributed_switch_port: Optional[pulumi.Input[str]] = None,
                 ipv4: Optional[pulumi.Input['VnicIpv4Args']] = None,
                 ipv6: Optional[pulumi.Input['VnicIpv6Args']] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 netstack: Optional[pulumi.Input[str]] = None,
                 portgroup: Optional[pulumi.Input[str]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Vnic resource.
        :param pulumi.Input[str] host: ESX host the interface belongs to
        :param pulumi.Input[str] distributed_port_group: Key of the distributed portgroup the nic will connect to.
        :param pulumi.Input[str] distributed_switch_port: UUID of the vdswitch the nic will be attached to. Do not set if you set portgroup.
        :param pulumi.Input['VnicIpv4Args'] ipv4: IPv4 settings. Either this or `ipv6` needs to be set. See IPv4 options below.
        :param pulumi.Input['VnicIpv6Args'] ipv6: IPv6 settings. Either this or `ipv6` needs to be set. See IPv6 options below.
        :param pulumi.Input[str] mac: MAC address of the interface.
        :param pulumi.Input[int] mtu: MTU of the interface.
        :param pulumi.Input[str] netstack: TCP/IP stack setting for this interface. Possible values are `defaultTcpipStack``, 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default:`defaultTcpipStack`)
        :param pulumi.Input[str] portgroup: Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: Enabled services setting for this interface. Currently support values are `vmotion`, `management`, and `vsan`.
        """
        pulumi.set(__self__, "host", host)
        if distributed_port_group is not None:
            pulumi.set(__self__, "distributed_port_group", distributed_port_group)
        if distributed_switch_port is not None:
            pulumi.set(__self__, "distributed_switch_port", distributed_switch_port)
        if ipv4 is not None:
            pulumi.set(__self__, "ipv4", ipv4)
        if ipv6 is not None:
            pulumi.set(__self__, "ipv6", ipv6)
        if mac is not None:
            pulumi.set(__self__, "mac", mac)
        if mtu is not None:
            pulumi.set(__self__, "mtu", mtu)
        if netstack is not None:
            pulumi.set(__self__, "netstack", netstack)
        if portgroup is not None:
            pulumi.set(__self__, "portgroup", portgroup)
        if services is not None:
            pulumi.set(__self__, "services", services)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Input[str]:
        """
        ESX host the interface belongs to
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: pulumi.Input[str]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="distributedPortGroup")
    def distributed_port_group(self) -> Optional[pulumi.Input[str]]:
        """
        Key of the distributed portgroup the nic will connect to.
        """
        return pulumi.get(self, "distributed_port_group")

    @distributed_port_group.setter
    def distributed_port_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "distributed_port_group", value)

    @property
    @pulumi.getter(name="distributedSwitchPort")
    def distributed_switch_port(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the vdswitch the nic will be attached to. Do not set if you set portgroup.
        """
        return pulumi.get(self, "distributed_switch_port")

    @distributed_switch_port.setter
    def distributed_switch_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "distributed_switch_port", value)

    @property
    @pulumi.getter
    def ipv4(self) -> Optional[pulumi.Input['VnicIpv4Args']]:
        """
        IPv4 settings. Either this or `ipv6` needs to be set. See IPv4 options below.
        """
        return pulumi.get(self, "ipv4")

    @ipv4.setter
    def ipv4(self, value: Optional[pulumi.Input['VnicIpv4Args']]):
        pulumi.set(self, "ipv4", value)

    @property
    @pulumi.getter
    def ipv6(self) -> Optional[pulumi.Input['VnicIpv6Args']]:
        """
        IPv6 settings. Either this or `ipv6` needs to be set. See IPv6 options below.
        """
        return pulumi.get(self, "ipv6")

    @ipv6.setter
    def ipv6(self, value: Optional[pulumi.Input['VnicIpv6Args']]):
        pulumi.set(self, "ipv6", value)

    @property
    @pulumi.getter
    def mac(self) -> Optional[pulumi.Input[str]]:
        """
        MAC address of the interface.
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac", value)

    @property
    @pulumi.getter
    def mtu(self) -> Optional[pulumi.Input[int]]:
        """
        MTU of the interface.
        """
        return pulumi.get(self, "mtu")

    @mtu.setter
    def mtu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mtu", value)

    @property
    @pulumi.getter
    def netstack(self) -> Optional[pulumi.Input[str]]:
        """
        TCP/IP stack setting for this interface. Possible values are `defaultTcpipStack``, 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default:`defaultTcpipStack`)
        """
        return pulumi.get(self, "netstack")

    @netstack.setter
    def netstack(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netstack", value)

    @property
    @pulumi.getter
    def portgroup(self) -> Optional[pulumi.Input[str]]:
        """
        Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
        """
        return pulumi.get(self, "portgroup")

    @portgroup.setter
    def portgroup(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "portgroup", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Enabled services setting for this interface. Currently support values are `vmotion`, `management`, and `vsan`.
        """
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "services", value)


@pulumi.input_type
class _VnicState:
    def __init__(__self__, *,
                 distributed_port_group: Optional[pulumi.Input[str]] = None,
                 distributed_switch_port: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 ipv4: Optional[pulumi.Input['VnicIpv4Args']] = None,
                 ipv6: Optional[pulumi.Input['VnicIpv6Args']] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 netstack: Optional[pulumi.Input[str]] = None,
                 portgroup: Optional[pulumi.Input[str]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Vnic resources.
        :param pulumi.Input[str] distributed_port_group: Key of the distributed portgroup the nic will connect to.
        :param pulumi.Input[str] distributed_switch_port: UUID of the vdswitch the nic will be attached to. Do not set if you set portgroup.
        :param pulumi.Input[str] host: ESX host the interface belongs to
        :param pulumi.Input['VnicIpv4Args'] ipv4: IPv4 settings. Either this or `ipv6` needs to be set. See IPv4 options below.
        :param pulumi.Input['VnicIpv6Args'] ipv6: IPv6 settings. Either this or `ipv6` needs to be set. See IPv6 options below.
        :param pulumi.Input[str] mac: MAC address of the interface.
        :param pulumi.Input[int] mtu: MTU of the interface.
        :param pulumi.Input[str] netstack: TCP/IP stack setting for this interface. Possible values are `defaultTcpipStack``, 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default:`defaultTcpipStack`)
        :param pulumi.Input[str] portgroup: Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: Enabled services setting for this interface. Currently support values are `vmotion`, `management`, and `vsan`.
        """
        if distributed_port_group is not None:
            pulumi.set(__self__, "distributed_port_group", distributed_port_group)
        if distributed_switch_port is not None:
            pulumi.set(__self__, "distributed_switch_port", distributed_switch_port)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if ipv4 is not None:
            pulumi.set(__self__, "ipv4", ipv4)
        if ipv6 is not None:
            pulumi.set(__self__, "ipv6", ipv6)
        if mac is not None:
            pulumi.set(__self__, "mac", mac)
        if mtu is not None:
            pulumi.set(__self__, "mtu", mtu)
        if netstack is not None:
            pulumi.set(__self__, "netstack", netstack)
        if portgroup is not None:
            pulumi.set(__self__, "portgroup", portgroup)
        if services is not None:
            pulumi.set(__self__, "services", services)

    @property
    @pulumi.getter(name="distributedPortGroup")
    def distributed_port_group(self) -> Optional[pulumi.Input[str]]:
        """
        Key of the distributed portgroup the nic will connect to.
        """
        return pulumi.get(self, "distributed_port_group")

    @distributed_port_group.setter
    def distributed_port_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "distributed_port_group", value)

    @property
    @pulumi.getter(name="distributedSwitchPort")
    def distributed_switch_port(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the vdswitch the nic will be attached to. Do not set if you set portgroup.
        """
        return pulumi.get(self, "distributed_switch_port")

    @distributed_switch_port.setter
    def distributed_switch_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "distributed_switch_port", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        ESX host the interface belongs to
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def ipv4(self) -> Optional[pulumi.Input['VnicIpv4Args']]:
        """
        IPv4 settings. Either this or `ipv6` needs to be set. See IPv4 options below.
        """
        return pulumi.get(self, "ipv4")

    @ipv4.setter
    def ipv4(self, value: Optional[pulumi.Input['VnicIpv4Args']]):
        pulumi.set(self, "ipv4", value)

    @property
    @pulumi.getter
    def ipv6(self) -> Optional[pulumi.Input['VnicIpv6Args']]:
        """
        IPv6 settings. Either this or `ipv6` needs to be set. See IPv6 options below.
        """
        return pulumi.get(self, "ipv6")

    @ipv6.setter
    def ipv6(self, value: Optional[pulumi.Input['VnicIpv6Args']]):
        pulumi.set(self, "ipv6", value)

    @property
    @pulumi.getter
    def mac(self) -> Optional[pulumi.Input[str]]:
        """
        MAC address of the interface.
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac", value)

    @property
    @pulumi.getter
    def mtu(self) -> Optional[pulumi.Input[int]]:
        """
        MTU of the interface.
        """
        return pulumi.get(self, "mtu")

    @mtu.setter
    def mtu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mtu", value)

    @property
    @pulumi.getter
    def netstack(self) -> Optional[pulumi.Input[str]]:
        """
        TCP/IP stack setting for this interface. Possible values are `defaultTcpipStack``, 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default:`defaultTcpipStack`)
        """
        return pulumi.get(self, "netstack")

    @netstack.setter
    def netstack(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netstack", value)

    @property
    @pulumi.getter
    def portgroup(self) -> Optional[pulumi.Input[str]]:
        """
        Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
        """
        return pulumi.get(self, "portgroup")

    @portgroup.setter
    def portgroup(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "portgroup", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Enabled services setting for this interface. Currently support values are `vmotion`, `management`, and `vsan`.
        """
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "services", value)


class Vnic(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 distributed_port_group: Optional[pulumi.Input[str]] = None,
                 distributed_switch_port: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 ipv4: Optional[pulumi.Input[Union['VnicIpv4Args', 'VnicIpv4ArgsDict']]] = None,
                 ipv6: Optional[pulumi.Input[Union['VnicIpv6Args', 'VnicIpv6ArgsDict']]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 netstack: Optional[pulumi.Input[str]] = None,
                 portgroup: Optional[pulumi.Input[str]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a VMware vSphere vnic resource.

        ## Example Usage

        ### Create a vnic attached to a distributed virtual switch using the vmotion TCP/IP stack

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        datacenter = vsphere.get_datacenter(name="dc-01")
        host = vsphere.get_host(name="esxi-01.example.com",
            datacenter_id=datacenter.id)
        vds = vsphere.DistributedVirtualSwitch("vds",
            name="vds-01",
            datacenter_id=datacenter.id,
            hosts=[{
                "host_system_id": host.id,
                "devices": ["vnic3"],
            }])
        pg = vsphere.DistributedPortGroup("pg",
            name="pg-01",
            vlan_id=1234,
            distributed_virtual_switch_uuid=vds.id)
        vnic = vsphere.Vnic("vnic",
            host=host.id,
            distributed_switch_port=vds.id,
            distributed_port_group=pg.id,
            ipv4={
                "dhcp": True,
            },
            netstack="vmotion")
        ```

        ### Create a vnic attached to a portgroup using the default TCP/IP stack

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        datacenter = vsphere.get_datacenter(name="dc-01")
        host = vsphere.get_host(name="esxi-01.example.com",
            datacenter_id=datacenter.id)
        hvs = vsphere.HostVirtualSwitch("hvs",
            name="hvs-01",
            host_system_id=host.id,
            network_adapters=[
                "vmnic3",
                "vmnic4",
            ],
            active_nics=["vmnic3"],
            standby_nics=["vmnic4"])
        pg = vsphere.HostPortGroup("pg",
            name="pg-01",
            virtual_switch_name=hvs.name,
            host_system_id=host.id)
        vnic = vsphere.Vnic("vnic",
            host=host.id,
            portgroup=pg.name,
            ipv4={
                "dhcp": True,
            },
            services=[
                "vsan",
                "management",
            ])
        ```

        ## Import

        An existing vNic can be imported into this resource

        via supplying the vNic's ID. An example is below:

        [docs-import]: /docs/import/index.html

        ```sh
        $ pulumi import vsphere:index/vnic:Vnic vnic host-123_vmk2
        ```

        The above would import the vnic `vmk2` from host with ID `host-123`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] distributed_port_group: Key of the distributed portgroup the nic will connect to.
        :param pulumi.Input[str] distributed_switch_port: UUID of the vdswitch the nic will be attached to. Do not set if you set portgroup.
        :param pulumi.Input[str] host: ESX host the interface belongs to
        :param pulumi.Input[Union['VnicIpv4Args', 'VnicIpv4ArgsDict']] ipv4: IPv4 settings. Either this or `ipv6` needs to be set. See IPv4 options below.
        :param pulumi.Input[Union['VnicIpv6Args', 'VnicIpv6ArgsDict']] ipv6: IPv6 settings. Either this or `ipv6` needs to be set. See IPv6 options below.
        :param pulumi.Input[str] mac: MAC address of the interface.
        :param pulumi.Input[int] mtu: MTU of the interface.
        :param pulumi.Input[str] netstack: TCP/IP stack setting for this interface. Possible values are `defaultTcpipStack``, 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default:`defaultTcpipStack`)
        :param pulumi.Input[str] portgroup: Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: Enabled services setting for this interface. Currently support values are `vmotion`, `management`, and `vsan`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VnicArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a VMware vSphere vnic resource.

        ## Example Usage

        ### Create a vnic attached to a distributed virtual switch using the vmotion TCP/IP stack

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        datacenter = vsphere.get_datacenter(name="dc-01")
        host = vsphere.get_host(name="esxi-01.example.com",
            datacenter_id=datacenter.id)
        vds = vsphere.DistributedVirtualSwitch("vds",
            name="vds-01",
            datacenter_id=datacenter.id,
            hosts=[{
                "host_system_id": host.id,
                "devices": ["vnic3"],
            }])
        pg = vsphere.DistributedPortGroup("pg",
            name="pg-01",
            vlan_id=1234,
            distributed_virtual_switch_uuid=vds.id)
        vnic = vsphere.Vnic("vnic",
            host=host.id,
            distributed_switch_port=vds.id,
            distributed_port_group=pg.id,
            ipv4={
                "dhcp": True,
            },
            netstack="vmotion")
        ```

        ### Create a vnic attached to a portgroup using the default TCP/IP stack

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        datacenter = vsphere.get_datacenter(name="dc-01")
        host = vsphere.get_host(name="esxi-01.example.com",
            datacenter_id=datacenter.id)
        hvs = vsphere.HostVirtualSwitch("hvs",
            name="hvs-01",
            host_system_id=host.id,
            network_adapters=[
                "vmnic3",
                "vmnic4",
            ],
            active_nics=["vmnic3"],
            standby_nics=["vmnic4"])
        pg = vsphere.HostPortGroup("pg",
            name="pg-01",
            virtual_switch_name=hvs.name,
            host_system_id=host.id)
        vnic = vsphere.Vnic("vnic",
            host=host.id,
            portgroup=pg.name,
            ipv4={
                "dhcp": True,
            },
            services=[
                "vsan",
                "management",
            ])
        ```

        ## Import

        An existing vNic can be imported into this resource

        via supplying the vNic's ID. An example is below:

        [docs-import]: /docs/import/index.html

        ```sh
        $ pulumi import vsphere:index/vnic:Vnic vnic host-123_vmk2
        ```

        The above would import the vnic `vmk2` from host with ID `host-123`.

        :param str resource_name: The name of the resource.
        :param VnicArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VnicArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 distributed_port_group: Optional[pulumi.Input[str]] = None,
                 distributed_switch_port: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 ipv4: Optional[pulumi.Input[Union['VnicIpv4Args', 'VnicIpv4ArgsDict']]] = None,
                 ipv6: Optional[pulumi.Input[Union['VnicIpv6Args', 'VnicIpv6ArgsDict']]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 netstack: Optional[pulumi.Input[str]] = None,
                 portgroup: Optional[pulumi.Input[str]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VnicArgs.__new__(VnicArgs)

            __props__.__dict__["distributed_port_group"] = distributed_port_group
            __props__.__dict__["distributed_switch_port"] = distributed_switch_port
            if host is None and not opts.urn:
                raise TypeError("Missing required property 'host'")
            __props__.__dict__["host"] = host
            __props__.__dict__["ipv4"] = ipv4
            __props__.__dict__["ipv6"] = ipv6
            __props__.__dict__["mac"] = mac
            __props__.__dict__["mtu"] = mtu
            __props__.__dict__["netstack"] = netstack
            __props__.__dict__["portgroup"] = portgroup
            __props__.__dict__["services"] = services
        super(Vnic, __self__).__init__(
            'vsphere:index/vnic:Vnic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            distributed_port_group: Optional[pulumi.Input[str]] = None,
            distributed_switch_port: Optional[pulumi.Input[str]] = None,
            host: Optional[pulumi.Input[str]] = None,
            ipv4: Optional[pulumi.Input[Union['VnicIpv4Args', 'VnicIpv4ArgsDict']]] = None,
            ipv6: Optional[pulumi.Input[Union['VnicIpv6Args', 'VnicIpv6ArgsDict']]] = None,
            mac: Optional[pulumi.Input[str]] = None,
            mtu: Optional[pulumi.Input[int]] = None,
            netstack: Optional[pulumi.Input[str]] = None,
            portgroup: Optional[pulumi.Input[str]] = None,
            services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Vnic':
        """
        Get an existing Vnic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] distributed_port_group: Key of the distributed portgroup the nic will connect to.
        :param pulumi.Input[str] distributed_switch_port: UUID of the vdswitch the nic will be attached to. Do not set if you set portgroup.
        :param pulumi.Input[str] host: ESX host the interface belongs to
        :param pulumi.Input[Union['VnicIpv4Args', 'VnicIpv4ArgsDict']] ipv4: IPv4 settings. Either this or `ipv6` needs to be set. See IPv4 options below.
        :param pulumi.Input[Union['VnicIpv6Args', 'VnicIpv6ArgsDict']] ipv6: IPv6 settings. Either this or `ipv6` needs to be set. See IPv6 options below.
        :param pulumi.Input[str] mac: MAC address of the interface.
        :param pulumi.Input[int] mtu: MTU of the interface.
        :param pulumi.Input[str] netstack: TCP/IP stack setting for this interface. Possible values are `defaultTcpipStack``, 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default:`defaultTcpipStack`)
        :param pulumi.Input[str] portgroup: Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: Enabled services setting for this interface. Currently support values are `vmotion`, `management`, and `vsan`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VnicState.__new__(_VnicState)

        __props__.__dict__["distributed_port_group"] = distributed_port_group
        __props__.__dict__["distributed_switch_port"] = distributed_switch_port
        __props__.__dict__["host"] = host
        __props__.__dict__["ipv4"] = ipv4
        __props__.__dict__["ipv6"] = ipv6
        __props__.__dict__["mac"] = mac
        __props__.__dict__["mtu"] = mtu
        __props__.__dict__["netstack"] = netstack
        __props__.__dict__["portgroup"] = portgroup
        __props__.__dict__["services"] = services
        return Vnic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="distributedPortGroup")
    def distributed_port_group(self) -> pulumi.Output[Optional[str]]:
        """
        Key of the distributed portgroup the nic will connect to.
        """
        return pulumi.get(self, "distributed_port_group")

    @property
    @pulumi.getter(name="distributedSwitchPort")
    def distributed_switch_port(self) -> pulumi.Output[Optional[str]]:
        """
        UUID of the vdswitch the nic will be attached to. Do not set if you set portgroup.
        """
        return pulumi.get(self, "distributed_switch_port")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        ESX host the interface belongs to
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def ipv4(self) -> pulumi.Output[Optional['outputs.VnicIpv4']]:
        """
        IPv4 settings. Either this or `ipv6` needs to be set. See IPv4 options below.
        """
        return pulumi.get(self, "ipv4")

    @property
    @pulumi.getter
    def ipv6(self) -> pulumi.Output[Optional['outputs.VnicIpv6']]:
        """
        IPv6 settings. Either this or `ipv6` needs to be set. See IPv6 options below.
        """
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter
    def mac(self) -> pulumi.Output[str]:
        """
        MAC address of the interface.
        """
        return pulumi.get(self, "mac")

    @property
    @pulumi.getter
    def mtu(self) -> pulumi.Output[int]:
        """
        MTU of the interface.
        """
        return pulumi.get(self, "mtu")

    @property
    @pulumi.getter
    def netstack(self) -> pulumi.Output[Optional[str]]:
        """
        TCP/IP stack setting for this interface. Possible values are `defaultTcpipStack``, 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default:`defaultTcpipStack`)
        """
        return pulumi.get(self, "netstack")

    @property
    @pulumi.getter
    def portgroup(self) -> pulumi.Output[Optional[str]]:
        """
        Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
        """
        return pulumi.get(self, "portgroup")

    @property
    @pulumi.getter
    def services(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Enabled services setting for this interface. Currently support values are `vmotion`, `management`, and `vsan`.
        """
        return pulumi.get(self, "services")


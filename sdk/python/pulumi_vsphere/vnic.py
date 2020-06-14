# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Vnic(pulumi.CustomResource):
    distributed_port_group: pulumi.Output[str]
    """
    Key of the distributed portgroup the nic will connect to.
    """
    distributed_switch_port: pulumi.Output[str]
    """
    UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
    """
    host: pulumi.Output[str]
    """
    ESX host the interface belongs to
    """
    ipv4: pulumi.Output[dict]
    """
    IPv4 settings. Either this or `ipv6` needs to be set. See  ipv4 options below.

      * `dhcp` (`bool`) - Use DHCP to configure the interface's IPv4 stack.
      * `gw` (`str`) - IP address of the default gateway, if DHCP or autoconfig is not set.
      * `ip` (`str`) - Address of the interface, if DHCP is not set.
      * `netmask` (`str`) - Netmask of the interface, if DHCP is not set.
    """
    ipv6: pulumi.Output[dict]
    """
    IPv6 settings. Either this or `ipv6` needs to be set. See  ipv6 options below.

      * `addresses` (`list`) - List of IPv6 addresses
      * `autoconfig` (`bool`) - Use IPv6 Autoconfiguration (RFC2462).
      * `dhcp` (`bool`) - Use DHCP to configure the interface's IPv4 stack.
      * `gw` (`str`) - IP address of the default gateway, if DHCP or autoconfig is not set.
    """
    mac: pulumi.Output[str]
    """
    MAC address of the interface.
    """
    mtu: pulumi.Output[float]
    """
    MTU of the interface.
    """
    netstack: pulumi.Output[str]
    """
    TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default: `defaultTcpipStack`)
    """
    portgroup: pulumi.Output[str]
    """
    Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
    """
    def __init__(__self__, resource_name, opts=None, distributed_port_group=None, distributed_switch_port=None, host=None, ipv4=None, ipv6=None, mac=None, mtu=None, netstack=None, portgroup=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a VMware vSphere vnic resource.

        ## Example Usage

        ### Create a vnic attached to a distributed virtual switch using the vmotion TCP/IP stack

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        dc = vsphere.get_datacenter(name="mydc")
        h1 = vsphere.get_host(name="esxi1.host.test",
            datacenter_id=dc.id)
        d1 = vsphere.DistributedVirtualSwitch("d1",
            datacenter_id=dc.id,
            host=[{
                "host_system_id": h1.id,
                "devices": ["vnic3"],
            }])
        p1 = vsphere.DistributedPortGroup("p1",
            vlan_id=1234,
            distributed_virtual_switch_uuid=d1.id)
        v1 = vsphere.Vnic("v1",
            host=h1.id,
            distributed_switch_port=d1.id,
            distributed_port_group=p1.id,
            ipv4={
                "dhcp": True,
            },
            netstack="vmotion")
        ```

        ### Create a vnic attached to a portgroup using the default TCP/IP stack

        ```python
        import pulumi
        import pulumi_vsphere as vsphere

        dc = vsphere.get_datacenter(name="mydc")
        h1 = vsphere.get_host(name="esxi1.host.test",
            datacenter_id=dc.id)
        hvs1 = vsphere.HostVirtualSwitch("hvs1",
            host_system_id=h1.id,
            network_adapters=[
                "vmnic3",
                "vmnic4",
            ],
            active_nics=["vmnic3"],
            standby_nics=["vmnic4"])
        p1 = vsphere.HostPortGroup("p1",
            virtual_switch_name=hvs1.name,
            host_system_id=h1.id)
        v1 = vsphere.Vnic("v1",
            host=h1.id,
            portgroup=p1.name,
            ipv4={
                "dhcp": True,
            })
        ```

        ## Importing 

        An existing vNic can be [imported][docs-import] into this resource
        via supplying the vNic's ID. An example is below:

        [docs-import]: /docs/import/index.html

        ```python
        import pulumi
        ```

        The above would import the the vnic `vmk2` from host with ID `host-123`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] distributed_port_group: Key of the distributed portgroup the nic will connect to.
        :param pulumi.Input[str] distributed_switch_port: UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
        :param pulumi.Input[str] host: ESX host the interface belongs to
        :param pulumi.Input[dict] ipv4: IPv4 settings. Either this or `ipv6` needs to be set. See  ipv4 options below.
        :param pulumi.Input[dict] ipv6: IPv6 settings. Either this or `ipv6` needs to be set. See  ipv6 options below.
        :param pulumi.Input[str] mac: MAC address of the interface.
        :param pulumi.Input[float] mtu: MTU of the interface.
        :param pulumi.Input[str] netstack: TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default: `defaultTcpipStack`)
        :param pulumi.Input[str] portgroup: Portgroup to attach the nic to. Do not set if you set distributed_switch_port.

        The **ipv4** object supports the following:

          * `dhcp` (`pulumi.Input[bool]`) - Use DHCP to configure the interface's IPv4 stack.
          * `gw` (`pulumi.Input[str]`) - IP address of the default gateway, if DHCP or autoconfig is not set.
          * `ip` (`pulumi.Input[str]`) - Address of the interface, if DHCP is not set.
          * `netmask` (`pulumi.Input[str]`) - Netmask of the interface, if DHCP is not set.

        The **ipv6** object supports the following:

          * `addresses` (`pulumi.Input[list]`) - List of IPv6 addresses
          * `autoconfig` (`pulumi.Input[bool]`) - Use IPv6 Autoconfiguration (RFC2462).
          * `dhcp` (`pulumi.Input[bool]`) - Use DHCP to configure the interface's IPv4 stack.
          * `gw` (`pulumi.Input[str]`) - IP address of the default gateway, if DHCP or autoconfig is not set.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['distributed_port_group'] = distributed_port_group
            __props__['distributed_switch_port'] = distributed_switch_port
            if host is None:
                raise TypeError("Missing required property 'host'")
            __props__['host'] = host
            __props__['ipv4'] = ipv4
            __props__['ipv6'] = ipv6
            __props__['mac'] = mac
            __props__['mtu'] = mtu
            __props__['netstack'] = netstack
            __props__['portgroup'] = portgroup
        super(Vnic, __self__).__init__(
            'vsphere:index/vnic:Vnic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, distributed_port_group=None, distributed_switch_port=None, host=None, ipv4=None, ipv6=None, mac=None, mtu=None, netstack=None, portgroup=None):
        """
        Get an existing Vnic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] distributed_port_group: Key of the distributed portgroup the nic will connect to.
        :param pulumi.Input[str] distributed_switch_port: UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
        :param pulumi.Input[str] host: ESX host the interface belongs to
        :param pulumi.Input[dict] ipv4: IPv4 settings. Either this or `ipv6` needs to be set. See  ipv4 options below.
        :param pulumi.Input[dict] ipv6: IPv6 settings. Either this or `ipv6` needs to be set. See  ipv6 options below.
        :param pulumi.Input[str] mac: MAC address of the interface.
        :param pulumi.Input[float] mtu: MTU of the interface.
        :param pulumi.Input[str] netstack: TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default: `defaultTcpipStack`)
        :param pulumi.Input[str] portgroup: Portgroup to attach the nic to. Do not set if you set distributed_switch_port.

        The **ipv4** object supports the following:

          * `dhcp` (`pulumi.Input[bool]`) - Use DHCP to configure the interface's IPv4 stack.
          * `gw` (`pulumi.Input[str]`) - IP address of the default gateway, if DHCP or autoconfig is not set.
          * `ip` (`pulumi.Input[str]`) - Address of the interface, if DHCP is not set.
          * `netmask` (`pulumi.Input[str]`) - Netmask of the interface, if DHCP is not set.

        The **ipv6** object supports the following:

          * `addresses` (`pulumi.Input[list]`) - List of IPv6 addresses
          * `autoconfig` (`pulumi.Input[bool]`) - Use IPv6 Autoconfiguration (RFC2462).
          * `dhcp` (`pulumi.Input[bool]`) - Use DHCP to configure the interface's IPv4 stack.
          * `gw` (`pulumi.Input[str]`) - IP address of the default gateway, if DHCP or autoconfig is not set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["distributed_port_group"] = distributed_port_group
        __props__["distributed_switch_port"] = distributed_switch_port
        __props__["host"] = host
        __props__["ipv4"] = ipv4
        __props__["ipv6"] = ipv6
        __props__["mac"] = mac
        __props__["mtu"] = mtu
        __props__["netstack"] = netstack
        __props__["portgroup"] = portgroup
        return Vnic(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from . import utilities, tables

class HostVirtualSwitch(pulumi.CustomResource):
    active_nics: pulumi.Output[list]
    """
    The list of active network adapters used for load
    balancing.
    """
    allow_forged_transmits: pulumi.Output[bool]
    """
    Controls whether or not the virtual
    network adapter is allowed to send network traffic with a different MAC
    address than that of its own. Default: `true`.
    """
    allow_mac_changes: pulumi.Output[bool]
    """
    Controls whether or not the Media Access
    Control (MAC) address can be changed. Default: `true`.
    """
    allow_promiscuous: pulumi.Output[bool]
    """
    Enable promiscuous mode on the network. This
    flag indicates whether or not all traffic is seen on a given port. Default:
    `false`.
    """
    beacon_interval: pulumi.Output[int]
    """
    The interval, in seconds, that a NIC beacon
    packet is sent out. This can be used with `check_beacon` to
    offer link failure capability beyond link status only. Default: `1`.
    """
    check_beacon: pulumi.Output[bool]
    """
    Enable beacon probing - this requires that the
    `beacon_interval` option has been set in the bridge
    options. If this is set to `false`, only link status is used to check for
    failed NICs.  Default: `false`.
    """
    failback: pulumi.Output[bool]
    """
    If set to `true`, the teaming policy will re-activate
    failed interfaces higher in precedence when they come back up.  Default:
    `true`.
    """
    host_system_id: pulumi.Output[str]
    """
    The [managed object ID][docs-about-morefs] of
    the host to set the virtual switch up on. Forces a new resource if changed.
    """
    link_discovery_operation: pulumi.Output[str]
    """
    Whether to `advertise` or `listen`
    for link discovery traffic. Default: `listen`.
    """
    link_discovery_protocol: pulumi.Output[str]
    """
    The discovery protocol type.  Valid
    types are `cpd` and `lldp`. Default: `cdp`.
    """
    mtu: pulumi.Output[int]
    """
    The maximum transmission unit (MTU) for the virtual
    switch. Default: `1500`.
    """
    name: pulumi.Output[str]
    """
    The name of the virtual switch. Forces a new resource if
    changed.
    """
    network_adapters: pulumi.Output[list]
    """
    The network interfaces to bind to the bridge.
    """
    notify_switches: pulumi.Output[bool]
    """
    If set to `true`, the teaming policy will
    notify the broadcast network of a NIC failover, triggering cache updates.
    Default: `true`.
    """
    number_of_ports: pulumi.Output[int]
    """
    The number of ports to create with this
    virtual switch. Default: `128`.
    """
    shaping_average_bandwidth: pulumi.Output[int]
    """
    The average bandwidth in bits per
    second if traffic shaping is enabled. Default: `0`
    """
    shaping_burst_size: pulumi.Output[int]
    """
    The maximum burst size allowed in bytes if
    shaping is enabled. Default: `0`
    """
    shaping_enabled: pulumi.Output[bool]
    """
    Set to `true` to enable the traffic shaper for
    ports managed by this virtual switch. Default: `false`.
    """
    shaping_peak_bandwidth: pulumi.Output[int]
    """
    The peak bandwidth during bursts in
    bits per second if traffic shaping is enabled. Default: `0`
    """
    standby_nics: pulumi.Output[list]
    """
    The list of standby network adapters used for
    failover.
    """
    teaming_policy: pulumi.Output[str]
    """
    The network adapter teaming policy. Can be one
    of `loadbalance_ip`, `loadbalance_srcmac`, `loadbalance_srcid`, or
    `failover_explicit`. Default: `loadbalance_srcid`.
    """
    def __init__(__self__, __name__, __opts__=None, active_nics=None, allow_forged_transmits=None, allow_mac_changes=None, allow_promiscuous=None, beacon_interval=None, check_beacon=None, failback=None, host_system_id=None, link_discovery_operation=None, link_discovery_protocol=None, mtu=None, name=None, network_adapters=None, notify_switches=None, number_of_ports=None, shaping_average_bandwidth=None, shaping_burst_size=None, shaping_enabled=None, shaping_peak_bandwidth=None, standby_nics=None, teaming_policy=None):
        """
        The `vsphere_host_virtual_switch` resource can be used to manage vSphere
        standard switches on an ESXi host. These switches can be used as a backing for
        standard port groups, which can be managed by the
        [`vsphere_host_port_group`][host-port-group] resource.
        
        For an overview on vSphere networking concepts, see [this
        page][ref-vsphere-net-concepts].
        
        [host-port-group]: /docs/providers/vsphere/r/host_port_group.html
        [ref-vsphere-net-concepts]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.networking.doc/GUID-2B11DBB8-CB3C-4AFF-8885-EFEA0FC562F4.html
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[list] active_nics: The list of active network adapters used for load
               balancing.
        :param pulumi.Input[bool] allow_forged_transmits: Controls whether or not the virtual
               network adapter is allowed to send network traffic with a different MAC
               address than that of its own. Default: `true`.
        :param pulumi.Input[bool] allow_mac_changes: Controls whether or not the Media Access
               Control (MAC) address can be changed. Default: `true`.
        :param pulumi.Input[bool] allow_promiscuous: Enable promiscuous mode on the network. This
               flag indicates whether or not all traffic is seen on a given port. Default:
               `false`.
        :param pulumi.Input[int] beacon_interval: The interval, in seconds, that a NIC beacon
               packet is sent out. This can be used with `check_beacon` to
               offer link failure capability beyond link status only. Default: `1`.
        :param pulumi.Input[bool] check_beacon: Enable beacon probing - this requires that the
               `beacon_interval` option has been set in the bridge
               options. If this is set to `false`, only link status is used to check for
               failed NICs.  Default: `false`.
        :param pulumi.Input[bool] failback: If set to `true`, the teaming policy will re-activate
               failed interfaces higher in precedence when they come back up.  Default:
               `true`.
        :param pulumi.Input[str] host_system_id: The [managed object ID][docs-about-morefs] of
               the host to set the virtual switch up on. Forces a new resource if changed.
        :param pulumi.Input[str] link_discovery_operation: Whether to `advertise` or `listen`
               for link discovery traffic. Default: `listen`.
        :param pulumi.Input[str] link_discovery_protocol: The discovery protocol type.  Valid
               types are `cpd` and `lldp`. Default: `cdp`.
        :param pulumi.Input[int] mtu: The maximum transmission unit (MTU) for the virtual
               switch. Default: `1500`.
        :param pulumi.Input[str] name: The name of the virtual switch. Forces a new resource if
               changed.
        :param pulumi.Input[list] network_adapters: The network interfaces to bind to the bridge.
        :param pulumi.Input[bool] notify_switches: If set to `true`, the teaming policy will
               notify the broadcast network of a NIC failover, triggering cache updates.
               Default: `true`.
        :param pulumi.Input[int] number_of_ports: The number of ports to create with this
               virtual switch. Default: `128`.
        :param pulumi.Input[int] shaping_average_bandwidth: The average bandwidth in bits per
               second if traffic shaping is enabled. Default: `0`
        :param pulumi.Input[int] shaping_burst_size: The maximum burst size allowed in bytes if
               shaping is enabled. Default: `0`
        :param pulumi.Input[bool] shaping_enabled: Set to `true` to enable the traffic shaper for
               ports managed by this virtual switch. Default: `false`.
        :param pulumi.Input[int] shaping_peak_bandwidth: The peak bandwidth during bursts in
               bits per second if traffic shaping is enabled. Default: `0`
        :param pulumi.Input[list] standby_nics: The list of standby network adapters used for
               failover.
        :param pulumi.Input[str] teaming_policy: The network adapter teaming policy. Can be one
               of `loadbalance_ip`, `loadbalance_srcmac`, `loadbalance_srcid`, or
               `failover_explicit`. Default: `loadbalance_srcid`.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not active_nics:
            raise TypeError('Missing required property active_nics')
        __props__['active_nics'] = active_nics

        __props__['allow_forged_transmits'] = allow_forged_transmits

        __props__['allow_mac_changes'] = allow_mac_changes

        __props__['allow_promiscuous'] = allow_promiscuous

        __props__['beacon_interval'] = beacon_interval

        __props__['check_beacon'] = check_beacon

        __props__['failback'] = failback

        if not host_system_id:
            raise TypeError('Missing required property host_system_id')
        __props__['host_system_id'] = host_system_id

        __props__['link_discovery_operation'] = link_discovery_operation

        __props__['link_discovery_protocol'] = link_discovery_protocol

        __props__['mtu'] = mtu

        __props__['name'] = name

        if not network_adapters:
            raise TypeError('Missing required property network_adapters')
        __props__['network_adapters'] = network_adapters

        __props__['notify_switches'] = notify_switches

        __props__['number_of_ports'] = number_of_ports

        __props__['shaping_average_bandwidth'] = shaping_average_bandwidth

        __props__['shaping_burst_size'] = shaping_burst_size

        __props__['shaping_enabled'] = shaping_enabled

        __props__['shaping_peak_bandwidth'] = shaping_peak_bandwidth

        if not standby_nics:
            raise TypeError('Missing required property standby_nics')
        __props__['standby_nics'] = standby_nics

        __props__['teaming_policy'] = teaming_policy

        super(HostVirtualSwitch, __self__).__init__(
            'vsphere:index/hostVirtualSwitch:HostVirtualSwitch',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


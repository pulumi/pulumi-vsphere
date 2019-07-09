# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class HaVmOverride(pulumi.CustomResource):
    compute_cluster_id: pulumi.Output[str]
    """
    The [managed object reference
    ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
    resource if changed.
    """
    ha_datastore_apd_recovery_action: pulumi.Output[str]
    """
    Controls the action to take
    on this virtual machine if an APD status on an affected datastore clears in
    the middle of an APD event. Can be one of `useClusterDefault`, `none` or
    `reset`.  Default: `useClusterDefault`.
    <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
    """
    ha_datastore_apd_response: pulumi.Output[str]
    """
    Controls the action to take on this
    virtual machine when the cluster has detected loss to all paths to a relevant
    datastore. Can be one of `clusterDefault`, `disabled`, `warning`,
    `restartConservative`, or `restartAggressive`.  Default: `clusterDefault`.
    <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
    """
    ha_datastore_apd_response_delay: pulumi.Output[float]
    """
    Controls the delay in minutes
    to wait after an APD timeout event to execute the response action defined in
    `ha_datastore_apd_response`. Use `-1` to use
    the cluster default. Default: `-1`.
    <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
    """
    ha_datastore_pdl_response: pulumi.Output[str]
    """
    Controls the action to take on this
    virtual machine when the cluster has detected a permanent device loss to a
    relevant datastore. Can be one of `clusterDefault`, `disabled`, `warning`, or
    `restartAggressive`. Default: `clusterDefault`.
    <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
    """
    ha_host_isolation_response: pulumi.Output[str]
    """
    The action to take on this virtual
    machine when a host has detected that it has been isolated from the rest of
    the cluster. Can be one of `clusterIsolationResponse`, `none`, `powerOff`, or
    `shutdown`. Default: `clusterIsolationResponse`.
    """
    ha_vm_failure_interval: pulumi.Output[float]
    """
    If a heartbeat from this virtual
    machine is not received within this configured interval, the virtual machine
    is marked as failed. The value is in seconds. Default: `30`.
    """
    ha_vm_maximum_failure_window: pulumi.Output[float]
    """
    The length of the reset window in
    which `ha_vm_maximum_resets` can operate. When this
    window expires, no more resets are attempted regardless of the setting
    configured in `ha_vm_maximum_resets`. `-1` means no window, meaning an
    unlimited reset time is allotted. The value is specified in seconds. Default:
    `-1` (no window).
    """
    ha_vm_maximum_resets: pulumi.Output[float]
    """
    The maximum number of resets that HA will
    perform to this virtual machine when responding to a failure event. Default:
    `3`
    """
    ha_vm_minimum_uptime: pulumi.Output[float]
    """
    The time, in seconds, that HA waits after
    powering on this virtual machine before monitoring for heartbeats. Default:
    `120` (2 minutes).
    """
    ha_vm_monitoring: pulumi.Output[str]
    """
    The type of virtual machine monitoring to use
    when HA is enabled in the cluster. Can be one of `vmMonitoringDisabled`,
    `vmMonitoringOnly`, or `vmAndAppMonitoring`. Default: `vmMonitoringDisabled`.
    """
    ha_vm_monitoring_use_cluster_defaults: pulumi.Output[bool]
    """
    Determines whether or
    not the cluster's default settings or the VM override settings specified in
    this resource are used for virtual machine monitoring. The default is `true`
    (use cluster defaults) - set to `false` to have overrides take effect.
    """
    ha_vm_restart_priority: pulumi.Output[str]
    """
    The restart priority for the virtual
    machine when vSphere detects a host failure. Can be one of
    `clusterRestartPriority`, `lowest`, `low`, `medium`, `high`, or `highest`.
    Default: `clusterRestartPriority`.
    """
    ha_vm_restart_timeout: pulumi.Output[float]
    """
    The maximum time, in seconds, that
    vSphere HA will wait for this virtual machine to be ready. Use `-1` to
    specify the cluster default.  Default: `-1`.
    <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
    """
    virtual_machine_id: pulumi.Output[str]
    """
    The UUID of the virtual machine to create
    the override for.  Forces a new resource if changed.
    """
    def __init__(__self__, resource_name, opts=None, compute_cluster_id=None, ha_datastore_apd_recovery_action=None, ha_datastore_apd_response=None, ha_datastore_apd_response_delay=None, ha_datastore_pdl_response=None, ha_host_isolation_response=None, ha_vm_failure_interval=None, ha_vm_maximum_failure_window=None, ha_vm_maximum_resets=None, ha_vm_minimum_uptime=None, ha_vm_monitoring=None, ha_vm_monitoring_use_cluster_defaults=None, ha_vm_restart_priority=None, ha_vm_restart_timeout=None, virtual_machine_id=None, __name__=None, __opts__=None):
        """
        The `vsphere_ha_vm_override` resource can be used to add an override for
        vSphere HA settings on a cluster for a specific virtual machine. With this
        resource, one can control specific HA settings so that they are different than
        the cluster default, accommodating the needs of that specific virtual machine,
        while not affecting the rest of the cluster.
        
        For more information on vSphere HA, see [this page][ref-vsphere-ha-clusters].
        
        [ref-vsphere-ha-clusters]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.avail.doc/GUID-5432CA24-14F1-44E3-87FB-61D937831CF6.html
        
        > **NOTE:** This resource requires vCenter and is not available on direct ESXi
        connections.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_cluster_id: The [managed object reference
               ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
               resource if changed.
        :param pulumi.Input[str] ha_datastore_apd_recovery_action: Controls the action to take
               on this virtual machine if an APD status on an affected datastore clears in
               the middle of an APD event. Can be one of `useClusterDefault`, `none` or
               `reset`.  Default: `useClusterDefault`.
               <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
        :param pulumi.Input[str] ha_datastore_apd_response: Controls the action to take on this
               virtual machine when the cluster has detected loss to all paths to a relevant
               datastore. Can be one of `clusterDefault`, `disabled`, `warning`,
               `restartConservative`, or `restartAggressive`.  Default: `clusterDefault`.
               <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
        :param pulumi.Input[float] ha_datastore_apd_response_delay: Controls the delay in minutes
               to wait after an APD timeout event to execute the response action defined in
               `ha_datastore_apd_response`. Use `-1` to use
               the cluster default. Default: `-1`.
               <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
        :param pulumi.Input[str] ha_datastore_pdl_response: Controls the action to take on this
               virtual machine when the cluster has detected a permanent device loss to a
               relevant datastore. Can be one of `clusterDefault`, `disabled`, `warning`, or
               `restartAggressive`. Default: `clusterDefault`.
               <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
        :param pulumi.Input[str] ha_host_isolation_response: The action to take on this virtual
               machine when a host has detected that it has been isolated from the rest of
               the cluster. Can be one of `clusterIsolationResponse`, `none`, `powerOff`, or
               `shutdown`. Default: `clusterIsolationResponse`.
        :param pulumi.Input[float] ha_vm_failure_interval: If a heartbeat from this virtual
               machine is not received within this configured interval, the virtual machine
               is marked as failed. The value is in seconds. Default: `30`.
        :param pulumi.Input[float] ha_vm_maximum_failure_window: The length of the reset window in
               which `ha_vm_maximum_resets` can operate. When this
               window expires, no more resets are attempted regardless of the setting
               configured in `ha_vm_maximum_resets`. `-1` means no window, meaning an
               unlimited reset time is allotted. The value is specified in seconds. Default:
               `-1` (no window).
        :param pulumi.Input[float] ha_vm_maximum_resets: The maximum number of resets that HA will
               perform to this virtual machine when responding to a failure event. Default:
               `3`
        :param pulumi.Input[float] ha_vm_minimum_uptime: The time, in seconds, that HA waits after
               powering on this virtual machine before monitoring for heartbeats. Default:
               `120` (2 minutes).
        :param pulumi.Input[str] ha_vm_monitoring: The type of virtual machine monitoring to use
               when HA is enabled in the cluster. Can be one of `vmMonitoringDisabled`,
               `vmMonitoringOnly`, or `vmAndAppMonitoring`. Default: `vmMonitoringDisabled`.
        :param pulumi.Input[bool] ha_vm_monitoring_use_cluster_defaults: Determines whether or
               not the cluster's default settings or the VM override settings specified in
               this resource are used for virtual machine monitoring. The default is `true`
               (use cluster defaults) - set to `false` to have overrides take effect.
        :param pulumi.Input[str] ha_vm_restart_priority: The restart priority for the virtual
               machine when vSphere detects a host failure. Can be one of
               `clusterRestartPriority`, `lowest`, `low`, `medium`, `high`, or `highest`.
               Default: `clusterRestartPriority`.
        :param pulumi.Input[float] ha_vm_restart_timeout: The maximum time, in seconds, that
               vSphere HA will wait for this virtual machine to be ready. Use `-1` to
               specify the cluster default.  Default: `-1`.
               <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
        :param pulumi.Input[str] virtual_machine_id: The UUID of the virtual machine to create
               the override for.  Forces a new resource if changed.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/ha_vm_override.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if compute_cluster_id is None:
            raise TypeError("Missing required property 'compute_cluster_id'")
        __props__['compute_cluster_id'] = compute_cluster_id

        __props__['ha_datastore_apd_recovery_action'] = ha_datastore_apd_recovery_action

        __props__['ha_datastore_apd_response'] = ha_datastore_apd_response

        __props__['ha_datastore_apd_response_delay'] = ha_datastore_apd_response_delay

        __props__['ha_datastore_pdl_response'] = ha_datastore_pdl_response

        __props__['ha_host_isolation_response'] = ha_host_isolation_response

        __props__['ha_vm_failure_interval'] = ha_vm_failure_interval

        __props__['ha_vm_maximum_failure_window'] = ha_vm_maximum_failure_window

        __props__['ha_vm_maximum_resets'] = ha_vm_maximum_resets

        __props__['ha_vm_minimum_uptime'] = ha_vm_minimum_uptime

        __props__['ha_vm_monitoring'] = ha_vm_monitoring

        __props__['ha_vm_monitoring_use_cluster_defaults'] = ha_vm_monitoring_use_cluster_defaults

        __props__['ha_vm_restart_priority'] = ha_vm_restart_priority

        __props__['ha_vm_restart_timeout'] = ha_vm_restart_timeout

        if virtual_machine_id is None:
            raise TypeError("Missing required property 'virtual_machine_id'")
        __props__['virtual_machine_id'] = virtual_machine_id

        super(HaVmOverride, __self__).__init__(
            'vsphere:index/haVmOverride:HaVmOverride',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


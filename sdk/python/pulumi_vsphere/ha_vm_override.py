# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['HaVmOverride']


class HaVmOverride(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_cluster_id: Optional[pulumi.Input[str]] = None,
                 ha_datastore_apd_recovery_action: Optional[pulumi.Input[str]] = None,
                 ha_datastore_apd_response: Optional[pulumi.Input[str]] = None,
                 ha_datastore_apd_response_delay: Optional[pulumi.Input[int]] = None,
                 ha_datastore_pdl_response: Optional[pulumi.Input[str]] = None,
                 ha_host_isolation_response: Optional[pulumi.Input[str]] = None,
                 ha_vm_failure_interval: Optional[pulumi.Input[int]] = None,
                 ha_vm_maximum_failure_window: Optional[pulumi.Input[int]] = None,
                 ha_vm_maximum_resets: Optional[pulumi.Input[int]] = None,
                 ha_vm_minimum_uptime: Optional[pulumi.Input[int]] = None,
                 ha_vm_monitoring: Optional[pulumi.Input[str]] = None,
                 ha_vm_monitoring_use_cluster_defaults: Optional[pulumi.Input[bool]] = None,
                 ha_vm_restart_priority: Optional[pulumi.Input[str]] = None,
                 ha_vm_restart_timeout: Optional[pulumi.Input[int]] = None,
                 virtual_machine_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a HaVmOverride resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the override in.  Forces a new
               resource if changed.
        :param pulumi.Input[str] ha_datastore_apd_recovery_action: Controls the action to take
               on this virtual machine if an APD status on an affected datastore clears in
               the middle of an APD event. Can be one of `useClusterDefault`, `none` or
               `reset`.  Default: `useClusterDefault`.
        :param pulumi.Input[str] ha_datastore_apd_response: Controls the action to take on this
               virtual machine when the cluster has detected loss to all paths to a relevant
               datastore. Can be one of `clusterDefault`, `disabled`, `warning`,
               `restartConservative`, or `restartAggressive`.  Default: `clusterDefault`.
        :param pulumi.Input[int] ha_datastore_apd_response_delay: Controls the delay in minutes
               to wait after an APD timeout event to execute the response action defined in
               `ha_datastore_apd_response`. Use `-1` to use
               the cluster default. Default: `-1`.
        :param pulumi.Input[str] ha_datastore_pdl_response: Controls the action to take on this
               virtual machine when the cluster has detected a permanent device loss to a
               relevant datastore. Can be one of `clusterDefault`, `disabled`, `warning`, or
               `restartAggressive`. Default: `clusterDefault`.
        :param pulumi.Input[str] ha_host_isolation_response: The action to take on this virtual
               machine when a host has detected that it has been isolated from the rest of
               the cluster. Can be one of `clusterIsolationResponse`, `none`, `powerOff`, or
               `shutdown`. Default: `clusterIsolationResponse`.
        :param pulumi.Input[int] ha_vm_failure_interval: If a heartbeat from this virtual
               machine is not received within this configured interval, the virtual machine
               is marked as failed. The value is in seconds. Default: `30`.
        :param pulumi.Input[int] ha_vm_maximum_failure_window: The length of the reset window in
               which `ha_vm_maximum_resets` can operate. When this
               window expires, no more resets are attempted regardless of the setting
               configured in `ha_vm_maximum_resets`. `-1` means no window, meaning an
               unlimited reset time is allotted. The value is specified in seconds. Default:
               `-1` (no window).
        :param pulumi.Input[int] ha_vm_maximum_resets: The maximum number of resets that HA will
               perform to this virtual machine when responding to a failure event. Default:
               `3`
        :param pulumi.Input[int] ha_vm_minimum_uptime: The time, in seconds, that HA waits after
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
        :param pulumi.Input[int] ha_vm_restart_timeout: The maximum time, in seconds, that
               vSphere HA will wait for this virtual machine to be ready. Use `-1` to
               specify the cluster default.  Default: `-1`.
        :param pulumi.Input[str] virtual_machine_id: The UUID of the virtual machine to create
               the override for.  Forces a new resource if changed.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
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

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            compute_cluster_id: Optional[pulumi.Input[str]] = None,
            ha_datastore_apd_recovery_action: Optional[pulumi.Input[str]] = None,
            ha_datastore_apd_response: Optional[pulumi.Input[str]] = None,
            ha_datastore_apd_response_delay: Optional[pulumi.Input[int]] = None,
            ha_datastore_pdl_response: Optional[pulumi.Input[str]] = None,
            ha_host_isolation_response: Optional[pulumi.Input[str]] = None,
            ha_vm_failure_interval: Optional[pulumi.Input[int]] = None,
            ha_vm_maximum_failure_window: Optional[pulumi.Input[int]] = None,
            ha_vm_maximum_resets: Optional[pulumi.Input[int]] = None,
            ha_vm_minimum_uptime: Optional[pulumi.Input[int]] = None,
            ha_vm_monitoring: Optional[pulumi.Input[str]] = None,
            ha_vm_monitoring_use_cluster_defaults: Optional[pulumi.Input[bool]] = None,
            ha_vm_restart_priority: Optional[pulumi.Input[str]] = None,
            ha_vm_restart_timeout: Optional[pulumi.Input[int]] = None,
            virtual_machine_id: Optional[pulumi.Input[str]] = None) -> 'HaVmOverride':
        """
        Get an existing HaVmOverride resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_cluster_id: The managed object reference
               ID of the cluster to put the override in.  Forces a new
               resource if changed.
        :param pulumi.Input[str] ha_datastore_apd_recovery_action: Controls the action to take
               on this virtual machine if an APD status on an affected datastore clears in
               the middle of an APD event. Can be one of `useClusterDefault`, `none` or
               `reset`.  Default: `useClusterDefault`.
        :param pulumi.Input[str] ha_datastore_apd_response: Controls the action to take on this
               virtual machine when the cluster has detected loss to all paths to a relevant
               datastore. Can be one of `clusterDefault`, `disabled`, `warning`,
               `restartConservative`, or `restartAggressive`.  Default: `clusterDefault`.
        :param pulumi.Input[int] ha_datastore_apd_response_delay: Controls the delay in minutes
               to wait after an APD timeout event to execute the response action defined in
               `ha_datastore_apd_response`. Use `-1` to use
               the cluster default. Default: `-1`.
        :param pulumi.Input[str] ha_datastore_pdl_response: Controls the action to take on this
               virtual machine when the cluster has detected a permanent device loss to a
               relevant datastore. Can be one of `clusterDefault`, `disabled`, `warning`, or
               `restartAggressive`. Default: `clusterDefault`.
        :param pulumi.Input[str] ha_host_isolation_response: The action to take on this virtual
               machine when a host has detected that it has been isolated from the rest of
               the cluster. Can be one of `clusterIsolationResponse`, `none`, `powerOff`, or
               `shutdown`. Default: `clusterIsolationResponse`.
        :param pulumi.Input[int] ha_vm_failure_interval: If a heartbeat from this virtual
               machine is not received within this configured interval, the virtual machine
               is marked as failed. The value is in seconds. Default: `30`.
        :param pulumi.Input[int] ha_vm_maximum_failure_window: The length of the reset window in
               which `ha_vm_maximum_resets` can operate. When this
               window expires, no more resets are attempted regardless of the setting
               configured in `ha_vm_maximum_resets`. `-1` means no window, meaning an
               unlimited reset time is allotted. The value is specified in seconds. Default:
               `-1` (no window).
        :param pulumi.Input[int] ha_vm_maximum_resets: The maximum number of resets that HA will
               perform to this virtual machine when responding to a failure event. Default:
               `3`
        :param pulumi.Input[int] ha_vm_minimum_uptime: The time, in seconds, that HA waits after
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
        :param pulumi.Input[int] ha_vm_restart_timeout: The maximum time, in seconds, that
               vSphere HA will wait for this virtual machine to be ready. Use `-1` to
               specify the cluster default.  Default: `-1`.
        :param pulumi.Input[str] virtual_machine_id: The UUID of the virtual machine to create
               the override for.  Forces a new resource if changed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["compute_cluster_id"] = compute_cluster_id
        __props__["ha_datastore_apd_recovery_action"] = ha_datastore_apd_recovery_action
        __props__["ha_datastore_apd_response"] = ha_datastore_apd_response
        __props__["ha_datastore_apd_response_delay"] = ha_datastore_apd_response_delay
        __props__["ha_datastore_pdl_response"] = ha_datastore_pdl_response
        __props__["ha_host_isolation_response"] = ha_host_isolation_response
        __props__["ha_vm_failure_interval"] = ha_vm_failure_interval
        __props__["ha_vm_maximum_failure_window"] = ha_vm_maximum_failure_window
        __props__["ha_vm_maximum_resets"] = ha_vm_maximum_resets
        __props__["ha_vm_minimum_uptime"] = ha_vm_minimum_uptime
        __props__["ha_vm_monitoring"] = ha_vm_monitoring
        __props__["ha_vm_monitoring_use_cluster_defaults"] = ha_vm_monitoring_use_cluster_defaults
        __props__["ha_vm_restart_priority"] = ha_vm_restart_priority
        __props__["ha_vm_restart_timeout"] = ha_vm_restart_timeout
        __props__["virtual_machine_id"] = virtual_machine_id
        return HaVmOverride(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="computeClusterId")
    def compute_cluster_id(self) -> pulumi.Output[str]:
        """
        The managed object reference
        ID of the cluster to put the override in.  Forces a new
        resource if changed.
        """
        return pulumi.get(self, "compute_cluster_id")

    @property
    @pulumi.getter(name="haDatastoreApdRecoveryAction")
    def ha_datastore_apd_recovery_action(self) -> pulumi.Output[Optional[str]]:
        """
        Controls the action to take
        on this virtual machine if an APD status on an affected datastore clears in
        the middle of an APD event. Can be one of `useClusterDefault`, `none` or
        `reset`.  Default: `useClusterDefault`.
        """
        return pulumi.get(self, "ha_datastore_apd_recovery_action")

    @property
    @pulumi.getter(name="haDatastoreApdResponse")
    def ha_datastore_apd_response(self) -> pulumi.Output[Optional[str]]:
        """
        Controls the action to take on this
        virtual machine when the cluster has detected loss to all paths to a relevant
        datastore. Can be one of `clusterDefault`, `disabled`, `warning`,
        `restartConservative`, or `restartAggressive`.  Default: `clusterDefault`.
        """
        return pulumi.get(self, "ha_datastore_apd_response")

    @property
    @pulumi.getter(name="haDatastoreApdResponseDelay")
    def ha_datastore_apd_response_delay(self) -> pulumi.Output[Optional[int]]:
        """
        Controls the delay in minutes
        to wait after an APD timeout event to execute the response action defined in
        `ha_datastore_apd_response`. Use `-1` to use
        the cluster default. Default: `-1`.
        """
        return pulumi.get(self, "ha_datastore_apd_response_delay")

    @property
    @pulumi.getter(name="haDatastorePdlResponse")
    def ha_datastore_pdl_response(self) -> pulumi.Output[Optional[str]]:
        """
        Controls the action to take on this
        virtual machine when the cluster has detected a permanent device loss to a
        relevant datastore. Can be one of `clusterDefault`, `disabled`, `warning`, or
        `restartAggressive`. Default: `clusterDefault`.
        """
        return pulumi.get(self, "ha_datastore_pdl_response")

    @property
    @pulumi.getter(name="haHostIsolationResponse")
    def ha_host_isolation_response(self) -> pulumi.Output[Optional[str]]:
        """
        The action to take on this virtual
        machine when a host has detected that it has been isolated from the rest of
        the cluster. Can be one of `clusterIsolationResponse`, `none`, `powerOff`, or
        `shutdown`. Default: `clusterIsolationResponse`.
        """
        return pulumi.get(self, "ha_host_isolation_response")

    @property
    @pulumi.getter(name="haVmFailureInterval")
    def ha_vm_failure_interval(self) -> pulumi.Output[Optional[int]]:
        """
        If a heartbeat from this virtual
        machine is not received within this configured interval, the virtual machine
        is marked as failed. The value is in seconds. Default: `30`.
        """
        return pulumi.get(self, "ha_vm_failure_interval")

    @property
    @pulumi.getter(name="haVmMaximumFailureWindow")
    def ha_vm_maximum_failure_window(self) -> pulumi.Output[Optional[int]]:
        """
        The length of the reset window in
        which `ha_vm_maximum_resets` can operate. When this
        window expires, no more resets are attempted regardless of the setting
        configured in `ha_vm_maximum_resets`. `-1` means no window, meaning an
        unlimited reset time is allotted. The value is specified in seconds. Default:
        `-1` (no window).
        """
        return pulumi.get(self, "ha_vm_maximum_failure_window")

    @property
    @pulumi.getter(name="haVmMaximumResets")
    def ha_vm_maximum_resets(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum number of resets that HA will
        perform to this virtual machine when responding to a failure event. Default:
        `3`
        """
        return pulumi.get(self, "ha_vm_maximum_resets")

    @property
    @pulumi.getter(name="haVmMinimumUptime")
    def ha_vm_minimum_uptime(self) -> pulumi.Output[Optional[int]]:
        """
        The time, in seconds, that HA waits after
        powering on this virtual machine before monitoring for heartbeats. Default:
        `120` (2 minutes).
        """
        return pulumi.get(self, "ha_vm_minimum_uptime")

    @property
    @pulumi.getter(name="haVmMonitoring")
    def ha_vm_monitoring(self) -> pulumi.Output[Optional[str]]:
        """
        The type of virtual machine monitoring to use
        when HA is enabled in the cluster. Can be one of `vmMonitoringDisabled`,
        `vmMonitoringOnly`, or `vmAndAppMonitoring`. Default: `vmMonitoringDisabled`.
        """
        return pulumi.get(self, "ha_vm_monitoring")

    @property
    @pulumi.getter(name="haVmMonitoringUseClusterDefaults")
    def ha_vm_monitoring_use_cluster_defaults(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines whether or
        not the cluster's default settings or the VM override settings specified in
        this resource are used for virtual machine monitoring. The default is `true`
        (use cluster defaults) - set to `false` to have overrides take effect.
        """
        return pulumi.get(self, "ha_vm_monitoring_use_cluster_defaults")

    @property
    @pulumi.getter(name="haVmRestartPriority")
    def ha_vm_restart_priority(self) -> pulumi.Output[Optional[str]]:
        """
        The restart priority for the virtual
        machine when vSphere detects a host failure. Can be one of
        `clusterRestartPriority`, `lowest`, `low`, `medium`, `high`, or `highest`.
        Default: `clusterRestartPriority`.
        """
        return pulumi.get(self, "ha_vm_restart_priority")

    @property
    @pulumi.getter(name="haVmRestartTimeout")
    def ha_vm_restart_timeout(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum time, in seconds, that
        vSphere HA will wait for this virtual machine to be ready. Use `-1` to
        specify the cluster default.  Default: `-1`.
        """
        return pulumi.get(self, "ha_vm_restart_timeout")

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> pulumi.Output[str]:
        """
        The UUID of the virtual machine to create
        the override for.  Forces a new resource if changed.
        """
        return pulumi.get(self, "virtual_machine_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


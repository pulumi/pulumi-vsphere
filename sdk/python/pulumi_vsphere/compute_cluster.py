# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class ComputeCluster(pulumi.CustomResource):
    custom_attributes: pulumi.Output[dict]
    """
    A map of custom attribute ids to attribute
    value strings to set for the datastore cluster. See
    [here][docs-setting-custom-attributes] for a reference on how to set values
    for custom attributes.
    """
    datacenter_id: pulumi.Output[str]
    """
    The [managed object ID][docs-about-morefs] of
    the datacenter to create the cluster in. Forces a new resource if changed.
    """
    dpm_automation_level: pulumi.Output[str]
    """
    The automation level for host power
    operations in this cluster. Can be one of `manual` or `automated`. Default:
    `manual`.
    """
    dpm_enabled: pulumi.Output[bool]
    """
    Enable DPM support for DRS in this cluster.
    Requires `drs_enabled` to be `true` in order to be effective.
    Default: `false`.
    """
    dpm_threshold: pulumi.Output[float]
    """
    A value between `1` and `5` indicating the
    threshold of load within the cluster that influences host power operations.
    This affects both power on and power off operations - a lower setting will
    tolerate more of a surplus/deficit than a higher setting. Default: `3`.
    """
    drs_advanced_options: pulumi.Output[dict]
    """
    A key/value map that specifies advanced
    options for DRS and DPM.
    """
    drs_automation_level: pulumi.Output[str]
    """
    The default automation level for all
    virtual machines in this cluster. Can be one of `manual`,
    `partiallyAutomated`, or `fullyAutomated`. Default: `manual`.
    """
    drs_enable_predictive_drs: pulumi.Output[bool]
    """
    When `true`, enables DRS to use data
    from [vRealize Operations Manager][ref-vsphere-vro] to make proactive DRS
    recommendations. <sup>\*</sup>
    """
    drs_enable_vm_overrides: pulumi.Output[bool]
    """
    Allow individual DRS overrides to be
    set for virtual machines in the cluster. Default: `true`.
    """
    drs_enabled: pulumi.Output[bool]
    """
    Enable DRS for this cluster. Default: `false`.
    """
    drs_migration_threshold: pulumi.Output[float]
    """
    A value between `1` and `5` indicating
    the threshold of imbalance tolerated between hosts. A lower setting will
    tolerate more imbalance while a higher setting will tolerate less. Default:
    `3`.
    """
    folder: pulumi.Output[str]
    force_evacuate_on_destroy: pulumi.Output[bool]
    ha_admission_control_failover_host_system_ids: pulumi.Output[list]
    """
    Defines the
    [managed object IDs][docs-about-morefs] of hosts to use as dedicated failover
    hosts. These hosts are kept as available as possible - admission control will
    block access to the host, and DRS will ignore the host when making
    recommendations.
    """
    ha_admission_control_host_failure_tolerance: pulumi.Output[float]
    """
    The maximum number
    of failed hosts that admission control tolerates when making decisions on
    whether to permit virtual machine operations. The maximum is one less than
    the number of hosts in the cluster. Default: `1`.
    <sup>\*</sup>
    """
    ha_admission_control_performance_tolerance: pulumi.Output[float]
    """
    The percentage of
    resource reduction that a cluster of virtual machines can tolerate in case of
    a failover. A value of 0 produces warnings only, whereas a value of 100
    disables the setting. Default: `100` (disabled).
    """
    ha_admission_control_policy: pulumi.Output[str]
    """
    The type of admission control
    policy to use with vSphere HA. Can be one of `resourcePercentage`,
    `slotPolicy`, `failoverHosts`, or `disabled`. Default: `resourcePercentage`.
    """
    ha_admission_control_resource_percentage_auto_compute: pulumi.Output[bool]
    """
    
    Automatically determine available resource percentages by subtracting the
    average number of host resources represented by the
    `ha_admission_control_host_failure_tolerance`
    setting from the total amount of resources in the cluster. Disable to supply
    user-defined values. Default: `true`.
    <sup>\*</sup>
    """
    ha_admission_control_resource_percentage_cpu: pulumi.Output[float]
    """
    Controls the
    user-defined percentage of CPU resources in the cluster to reserve for
    failover. Default: `100`.
    """
    ha_admission_control_resource_percentage_memory: pulumi.Output[float]
    """
    Controls the
    user-defined percentage of memory resources in the cluster to reserve for
    failover. Default: `100`.
    """
    ha_admission_control_slot_policy_explicit_cpu: pulumi.Output[float]
    """
    Controls the
    user-defined CPU slot size, in MHz. Default: `32`.
    """
    ha_admission_control_slot_policy_explicit_memory: pulumi.Output[float]
    """
    Controls the
    user-defined memory slot size, in MB. Default: `100`.
    """
    ha_admission_control_slot_policy_use_explicit_size: pulumi.Output[bool]
    """
    Controls
    whether or not you wish to supply explicit values to CPU and memory slot
    sizes. The default is `false`, which tells vSphere to gather a automatic
    average based on all powered-on virtual machines currently in the cluster.
    """
    ha_advanced_options: pulumi.Output[dict]
    """
    A key/value map that specifies advanced
    options for vSphere HA.
    """
    ha_datastore_apd_recovery_action: pulumi.Output[str]
    """
    Controls the action to take
    on virtual machines if an APD status on an affected datastore clears in the
    middle of an APD event. Can be one of `none` or `reset`. Default: `none`.
    <sup>\*</sup>
    """
    ha_datastore_apd_response: pulumi.Output[str]
    """
    Controls the action to take on
    virtual machines when the cluster has detected loss to all paths to a
    relevant datastore. Can be one of `disabled`, `warning`,
    `restartConservative`, or `restartAggressive`.  Default: `disabled`.
    <sup>\*</sup>
    """
    ha_datastore_apd_response_delay: pulumi.Output[float]
    """
    Controls the delay in minutes
    to wait after an APD timeout event to execute the response action defined in
    `ha_datastore_apd_response`. Default: `3`
    minutes. <sup>\*</sup>
    """
    ha_datastore_pdl_response: pulumi.Output[str]
    """
    Controls the action to take on
    virtual machines when the cluster has detected a permanent device loss to a
    relevant datastore. Can be one of `disabled`, `warning`, or
    `restartAggressive`. Default: `disabled`.
    <sup>\*</sup>
    """
    ha_enabled: pulumi.Output[bool]
    """
    Enable vSphere HA for this cluster. Default:
    `false`.
    """
    ha_heartbeat_datastore_ids: pulumi.Output[list]
    """
    The list of managed object IDs for
    preferred datastores to use for HA heartbeating. This setting is only useful
    when `ha_heartbeat_datastore_policy` is set
    to either `userSelectedDs` or `allFeasibleDsWithUserPreference`.
    """
    ha_heartbeat_datastore_policy: pulumi.Output[str]
    """
    The selection policy for HA
    heartbeat datastores. Can be one of `allFeasibleDs`, `userSelectedDs`, or
    `allFeasibleDsWithUserPreference`. Default:
    `allFeasibleDsWithUserPreference`.
    """
    ha_host_isolation_response: pulumi.Output[str]
    """
    The action to take on virtual
    machines when a host has detected that it has been isolated from the rest of
    the cluster. Can be one of `none`, `powerOff`, or `shutdown`. Default:
    `none`.
    """
    ha_host_monitoring: pulumi.Output[str]
    """
    Global setting that controls whether
    vSphere HA remediates virtual machines on host failure. Can be one of `enabled`
    or `disabled`. Default: `enabled`.
    """
    ha_vm_component_protection: pulumi.Output[str]
    """
    Controls vSphere VM component
    protection for virtual machines in this cluster. Can be one of `enabled` or
    `disabled`. Default: `enabled`.
    <sup>\*</sup>
    """
    ha_vm_dependency_restart_condition: pulumi.Output[str]
    """
    The condition used to
    determine whether or not virtual machines in a certain restart priority class
    are online, allowing HA to move on to restarting virtual machines on the next
    priority. Can be one of `none`, `poweredOn`, `guestHbStatusGreen`, or
    `appHbStatusGreen`. The default is `none`, which means that a virtual machine
    is considered ready immediately after a host is found to start it on.
    <sup>\*</sup>
    """
    ha_vm_failure_interval: pulumi.Output[float]
    """
    If a heartbeat from a virtual machine
    is not received within this configured interval, the virtual machine is
    marked as failed. The value is in seconds. Default: `30`.
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
    perform to a virtual machine when responding to a failure event. Default: `3`
    """
    ha_vm_minimum_uptime: pulumi.Output[float]
    """
    The time, in seconds, that HA waits after
    powering on a virtual machine before monitoring for heartbeats. Default:
    `120` (2 minutes).
    """
    ha_vm_monitoring: pulumi.Output[str]
    """
    The type of virtual machine monitoring to use
    when HA is enabled in the cluster. Can be one of `vmMonitoringDisabled`,
    `vmMonitoringOnly`, or `vmAndAppMonitoring`. Default: `vmMonitoringDisabled`.
    """
    ha_vm_restart_additional_delay: pulumi.Output[float]
    """
    Additional delay in seconds
    after ready condition is met. A VM is considered ready at this point.
    Default: `0` (no delay). <sup>\*</sup>
    """
    ha_vm_restart_priority: pulumi.Output[str]
    """
    The default restart priority
    for affected virtual machines when vSphere detects a host failure. Can be one
    of `lowest`, `low`, `medium`, `high`, or `highest`. Default: `medium`.
    """
    ha_vm_restart_timeout: pulumi.Output[float]
    """
    The maximum time, in seconds,
    that vSphere HA will wait for virtual machines in one priority to be ready
    before proceeding with the next priority. Default: `600` (10 minutes).
    <sup>\*</sup>
    """
    host_cluster_exit_timeout: pulumi.Output[float]
    """
    The timeout for each host maintenance mode
    operation when removing hosts from a cluster. The value is specified in
    seconds. Default: `3600` (1 hour).
    """
    host_system_ids: pulumi.Output[list]
    """
    The [managed object IDs][docs-about-morefs] of
    the hosts to put in the cluster.
    """
    name: pulumi.Output[str]
    """
    The name of the cluster.
    """
    proactive_ha_automation_level: pulumi.Output[str]
    """
    Determines how the host
    quarantine, maintenance mode, or virtual machine migration recommendations
    made by proactive HA are to be handled. Can be one of `Automated` or
    `Manual`. Default: `Manual`. <sup>\*</sup>
    """
    proactive_ha_enabled: pulumi.Output[bool]
    """
    Enables Proactive HA. Default: `false`.
    <sup>\*</sup>
    """
    proactive_ha_moderate_remediation: pulumi.Output[str]
    """
    The configured remediation
    for moderately degraded hosts. Can be one of `MaintenanceMode` or
    `QuarantineMode`. Note that this cannot be set to `MaintenanceMode` when
    `proactive_ha_severe_remediation` is set
    to `QuarantineMode`. Default: `QuarantineMode`.
    <sup>\*</sup>
    """
    proactive_ha_provider_ids: pulumi.Output[list]
    """
    The list of IDs for health update
    providers configured for this cluster.
    <sup>\*</sup>
    """
    proactive_ha_severe_remediation: pulumi.Output[str]
    """
    The configured remediation for
    severely degraded hosts. Can be one of `MaintenanceMode` or `QuarantineMode`.
    Note that this cannot be set to `QuarantineMode` when
    `proactive_ha_moderate_remediation` is
    set to `MaintenanceMode`. Default: `QuarantineMode`.
    <sup>\*</sup>
    """
    resource_pool_id: pulumi.Output[str]
    tags: pulumi.Output[list]
    """
    The IDs of any tags to attach to this resource. See
    [here][docs-applying-tags] for a reference on how to apply tags.
    """
    def __init__(__self__, resource_name, opts=None, custom_attributes=None, datacenter_id=None, dpm_automation_level=None, dpm_enabled=None, dpm_threshold=None, drs_advanced_options=None, drs_automation_level=None, drs_enable_predictive_drs=None, drs_enable_vm_overrides=None, drs_enabled=None, drs_migration_threshold=None, folder=None, force_evacuate_on_destroy=None, ha_admission_control_failover_host_system_ids=None, ha_admission_control_host_failure_tolerance=None, ha_admission_control_performance_tolerance=None, ha_admission_control_policy=None, ha_admission_control_resource_percentage_auto_compute=None, ha_admission_control_resource_percentage_cpu=None, ha_admission_control_resource_percentage_memory=None, ha_admission_control_slot_policy_explicit_cpu=None, ha_admission_control_slot_policy_explicit_memory=None, ha_admission_control_slot_policy_use_explicit_size=None, ha_advanced_options=None, ha_datastore_apd_recovery_action=None, ha_datastore_apd_response=None, ha_datastore_apd_response_delay=None, ha_datastore_pdl_response=None, ha_enabled=None, ha_heartbeat_datastore_ids=None, ha_heartbeat_datastore_policy=None, ha_host_isolation_response=None, ha_host_monitoring=None, ha_vm_component_protection=None, ha_vm_dependency_restart_condition=None, ha_vm_failure_interval=None, ha_vm_maximum_failure_window=None, ha_vm_maximum_resets=None, ha_vm_minimum_uptime=None, ha_vm_monitoring=None, ha_vm_restart_additional_delay=None, ha_vm_restart_priority=None, ha_vm_restart_timeout=None, host_cluster_exit_timeout=None, host_system_ids=None, name=None, proactive_ha_automation_level=None, proactive_ha_enabled=None, proactive_ha_moderate_remediation=None, proactive_ha_provider_ids=None, proactive_ha_severe_remediation=None, tags=None, __name__=None, __opts__=None):
        """
        Create a ComputeCluster resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] custom_attributes: A map of custom attribute ids to attribute
               value strings to set for the datastore cluster. See
               [here][docs-setting-custom-attributes] for a reference on how to set values
               for custom attributes.
        :param pulumi.Input[str] datacenter_id: The [managed object ID][docs-about-morefs] of
               the datacenter to create the cluster in. Forces a new resource if changed.
        :param pulumi.Input[str] dpm_automation_level: The automation level for host power
               operations in this cluster. Can be one of `manual` or `automated`. Default:
               `manual`.
        :param pulumi.Input[bool] dpm_enabled: Enable DPM support for DRS in this cluster.
               Requires `drs_enabled` to be `true` in order to be effective.
               Default: `false`.
        :param pulumi.Input[float] dpm_threshold: A value between `1` and `5` indicating the
               threshold of load within the cluster that influences host power operations.
               This affects both power on and power off operations - a lower setting will
               tolerate more of a surplus/deficit than a higher setting. Default: `3`.
        :param pulumi.Input[dict] drs_advanced_options: A key/value map that specifies advanced
               options for DRS and DPM.
        :param pulumi.Input[str] drs_automation_level: The default automation level for all
               virtual machines in this cluster. Can be one of `manual`,
               `partiallyAutomated`, or `fullyAutomated`. Default: `manual`.
        :param pulumi.Input[bool] drs_enable_predictive_drs: When `true`, enables DRS to use data
               from [vRealize Operations Manager][ref-vsphere-vro] to make proactive DRS
               recommendations. <sup>\*</sup>
        :param pulumi.Input[bool] drs_enable_vm_overrides: Allow individual DRS overrides to be
               set for virtual machines in the cluster. Default: `true`.
        :param pulumi.Input[bool] drs_enabled: Enable DRS for this cluster. Default: `false`.
        :param pulumi.Input[float] drs_migration_threshold: A value between `1` and `5` indicating
               the threshold of imbalance tolerated between hosts. A lower setting will
               tolerate more imbalance while a higher setting will tolerate less. Default:
               `3`.
        :param pulumi.Input[list] ha_admission_control_failover_host_system_ids: Defines the
               [managed object IDs][docs-about-morefs] of hosts to use as dedicated failover
               hosts. These hosts are kept as available as possible - admission control will
               block access to the host, and DRS will ignore the host when making
               recommendations.
        :param pulumi.Input[float] ha_admission_control_host_failure_tolerance: The maximum number
               of failed hosts that admission control tolerates when making decisions on
               whether to permit virtual machine operations. The maximum is one less than
               the number of hosts in the cluster. Default: `1`.
               <sup>\*</sup>
        :param pulumi.Input[float] ha_admission_control_performance_tolerance: The percentage of
               resource reduction that a cluster of virtual machines can tolerate in case of
               a failover. A value of 0 produces warnings only, whereas a value of 100
               disables the setting. Default: `100` (disabled).
        :param pulumi.Input[str] ha_admission_control_policy: The type of admission control
               policy to use with vSphere HA. Can be one of `resourcePercentage`,
               `slotPolicy`, `failoverHosts`, or `disabled`. Default: `resourcePercentage`.
        :param pulumi.Input[bool] ha_admission_control_resource_percentage_auto_compute: 
               Automatically determine available resource percentages by subtracting the
               average number of host resources represented by the
               `ha_admission_control_host_failure_tolerance`
               setting from the total amount of resources in the cluster. Disable to supply
               user-defined values. Default: `true`.
               <sup>\*</sup>
        :param pulumi.Input[float] ha_admission_control_resource_percentage_cpu: Controls the
               user-defined percentage of CPU resources in the cluster to reserve for
               failover. Default: `100`.
        :param pulumi.Input[float] ha_admission_control_resource_percentage_memory: Controls the
               user-defined percentage of memory resources in the cluster to reserve for
               failover. Default: `100`.
        :param pulumi.Input[float] ha_admission_control_slot_policy_explicit_cpu: Controls the
               user-defined CPU slot size, in MHz. Default: `32`.
        :param pulumi.Input[float] ha_admission_control_slot_policy_explicit_memory: Controls the
               user-defined memory slot size, in MB. Default: `100`.
        :param pulumi.Input[bool] ha_admission_control_slot_policy_use_explicit_size: Controls
               whether or not you wish to supply explicit values to CPU and memory slot
               sizes. The default is `false`, which tells vSphere to gather a automatic
               average based on all powered-on virtual machines currently in the cluster.
        :param pulumi.Input[dict] ha_advanced_options: A key/value map that specifies advanced
               options for vSphere HA.
        :param pulumi.Input[str] ha_datastore_apd_recovery_action: Controls the action to take
               on virtual machines if an APD status on an affected datastore clears in the
               middle of an APD event. Can be one of `none` or `reset`. Default: `none`.
               <sup>\*</sup>
        :param pulumi.Input[str] ha_datastore_apd_response: Controls the action to take on
               virtual machines when the cluster has detected loss to all paths to a
               relevant datastore. Can be one of `disabled`, `warning`,
               `restartConservative`, or `restartAggressive`.  Default: `disabled`.
               <sup>\*</sup>
        :param pulumi.Input[float] ha_datastore_apd_response_delay: Controls the delay in minutes
               to wait after an APD timeout event to execute the response action defined in
               `ha_datastore_apd_response`. Default: `3`
               minutes. <sup>\*</sup>
        :param pulumi.Input[str] ha_datastore_pdl_response: Controls the action to take on
               virtual machines when the cluster has detected a permanent device loss to a
               relevant datastore. Can be one of `disabled`, `warning`, or
               `restartAggressive`. Default: `disabled`.
               <sup>\*</sup>
        :param pulumi.Input[bool] ha_enabled: Enable vSphere HA for this cluster. Default:
               `false`.
        :param pulumi.Input[list] ha_heartbeat_datastore_ids: The list of managed object IDs for
               preferred datastores to use for HA heartbeating. This setting is only useful
               when `ha_heartbeat_datastore_policy` is set
               to either `userSelectedDs` or `allFeasibleDsWithUserPreference`.
        :param pulumi.Input[str] ha_heartbeat_datastore_policy: The selection policy for HA
               heartbeat datastores. Can be one of `allFeasibleDs`, `userSelectedDs`, or
               `allFeasibleDsWithUserPreference`. Default:
               `allFeasibleDsWithUserPreference`.
        :param pulumi.Input[str] ha_host_isolation_response: The action to take on virtual
               machines when a host has detected that it has been isolated from the rest of
               the cluster. Can be one of `none`, `powerOff`, or `shutdown`. Default:
               `none`.
        :param pulumi.Input[str] ha_host_monitoring: Global setting that controls whether
               vSphere HA remediates virtual machines on host failure. Can be one of `enabled`
               or `disabled`. Default: `enabled`.
        :param pulumi.Input[str] ha_vm_component_protection: Controls vSphere VM component
               protection for virtual machines in this cluster. Can be one of `enabled` or
               `disabled`. Default: `enabled`.
               <sup>\*</sup>
        :param pulumi.Input[str] ha_vm_dependency_restart_condition: The condition used to
               determine whether or not virtual machines in a certain restart priority class
               are online, allowing HA to move on to restarting virtual machines on the next
               priority. Can be one of `none`, `poweredOn`, `guestHbStatusGreen`, or
               `appHbStatusGreen`. The default is `none`, which means that a virtual machine
               is considered ready immediately after a host is found to start it on.
               <sup>\*</sup>
        :param pulumi.Input[float] ha_vm_failure_interval: If a heartbeat from a virtual machine
               is not received within this configured interval, the virtual machine is
               marked as failed. The value is in seconds. Default: `30`.
        :param pulumi.Input[float] ha_vm_maximum_failure_window: The length of the reset window in
               which `ha_vm_maximum_resets` can operate. When this
               window expires, no more resets are attempted regardless of the setting
               configured in `ha_vm_maximum_resets`. `-1` means no window, meaning an
               unlimited reset time is allotted. The value is specified in seconds. Default:
               `-1` (no window).
        :param pulumi.Input[float] ha_vm_maximum_resets: The maximum number of resets that HA will
               perform to a virtual machine when responding to a failure event. Default: `3`
        :param pulumi.Input[float] ha_vm_minimum_uptime: The time, in seconds, that HA waits after
               powering on a virtual machine before monitoring for heartbeats. Default:
               `120` (2 minutes).
        :param pulumi.Input[str] ha_vm_monitoring: The type of virtual machine monitoring to use
               when HA is enabled in the cluster. Can be one of `vmMonitoringDisabled`,
               `vmMonitoringOnly`, or `vmAndAppMonitoring`. Default: `vmMonitoringDisabled`.
        :param pulumi.Input[float] ha_vm_restart_additional_delay: Additional delay in seconds
               after ready condition is met. A VM is considered ready at this point.
               Default: `0` (no delay). <sup>\*</sup>
        :param pulumi.Input[str] ha_vm_restart_priority: The default restart priority
               for affected virtual machines when vSphere detects a host failure. Can be one
               of `lowest`, `low`, `medium`, `high`, or `highest`. Default: `medium`.
        :param pulumi.Input[float] ha_vm_restart_timeout: The maximum time, in seconds,
               that vSphere HA will wait for virtual machines in one priority to be ready
               before proceeding with the next priority. Default: `600` (10 minutes).
               <sup>\*</sup>
        :param pulumi.Input[float] host_cluster_exit_timeout: The timeout for each host maintenance mode
               operation when removing hosts from a cluster. The value is specified in
               seconds. Default: `3600` (1 hour).
        :param pulumi.Input[list] host_system_ids: The [managed object IDs][docs-about-morefs] of
               the hosts to put in the cluster.
        :param pulumi.Input[str] name: The name of the cluster.
        :param pulumi.Input[str] proactive_ha_automation_level: Determines how the host
               quarantine, maintenance mode, or virtual machine migration recommendations
               made by proactive HA are to be handled. Can be one of `Automated` or
               `Manual`. Default: `Manual`. <sup>\*</sup>
        :param pulumi.Input[bool] proactive_ha_enabled: Enables Proactive HA. Default: `false`.
               <sup>\*</sup>
        :param pulumi.Input[str] proactive_ha_moderate_remediation: The configured remediation
               for moderately degraded hosts. Can be one of `MaintenanceMode` or
               `QuarantineMode`. Note that this cannot be set to `MaintenanceMode` when
               `proactive_ha_severe_remediation` is set
               to `QuarantineMode`. Default: `QuarantineMode`.
               <sup>\*</sup>
        :param pulumi.Input[list] proactive_ha_provider_ids: The list of IDs for health update
               providers configured for this cluster.
               <sup>\*</sup>
        :param pulumi.Input[str] proactive_ha_severe_remediation: The configured remediation for
               severely degraded hosts. Can be one of `MaintenanceMode` or `QuarantineMode`.
               Note that this cannot be set to `QuarantineMode` when
               `proactive_ha_moderate_remediation` is
               set to `MaintenanceMode`. Default: `QuarantineMode`.
               <sup>\*</sup>
        :param pulumi.Input[list] tags: The IDs of any tags to attach to this resource. See
               [here][docs-applying-tags] for a reference on how to apply tags.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/compute_cluster.html.markdown.
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

        __props__['custom_attributes'] = custom_attributes

        if datacenter_id is None:
            raise TypeError("Missing required property 'datacenter_id'")
        __props__['datacenter_id'] = datacenter_id

        __props__['dpm_automation_level'] = dpm_automation_level

        __props__['dpm_enabled'] = dpm_enabled

        __props__['dpm_threshold'] = dpm_threshold

        __props__['drs_advanced_options'] = drs_advanced_options

        __props__['drs_automation_level'] = drs_automation_level

        __props__['drs_enable_predictive_drs'] = drs_enable_predictive_drs

        __props__['drs_enable_vm_overrides'] = drs_enable_vm_overrides

        __props__['drs_enabled'] = drs_enabled

        __props__['drs_migration_threshold'] = drs_migration_threshold

        __props__['folder'] = folder

        __props__['force_evacuate_on_destroy'] = force_evacuate_on_destroy

        __props__['ha_admission_control_failover_host_system_ids'] = ha_admission_control_failover_host_system_ids

        __props__['ha_admission_control_host_failure_tolerance'] = ha_admission_control_host_failure_tolerance

        __props__['ha_admission_control_performance_tolerance'] = ha_admission_control_performance_tolerance

        __props__['ha_admission_control_policy'] = ha_admission_control_policy

        __props__['ha_admission_control_resource_percentage_auto_compute'] = ha_admission_control_resource_percentage_auto_compute

        __props__['ha_admission_control_resource_percentage_cpu'] = ha_admission_control_resource_percentage_cpu

        __props__['ha_admission_control_resource_percentage_memory'] = ha_admission_control_resource_percentage_memory

        __props__['ha_admission_control_slot_policy_explicit_cpu'] = ha_admission_control_slot_policy_explicit_cpu

        __props__['ha_admission_control_slot_policy_explicit_memory'] = ha_admission_control_slot_policy_explicit_memory

        __props__['ha_admission_control_slot_policy_use_explicit_size'] = ha_admission_control_slot_policy_use_explicit_size

        __props__['ha_advanced_options'] = ha_advanced_options

        __props__['ha_datastore_apd_recovery_action'] = ha_datastore_apd_recovery_action

        __props__['ha_datastore_apd_response'] = ha_datastore_apd_response

        __props__['ha_datastore_apd_response_delay'] = ha_datastore_apd_response_delay

        __props__['ha_datastore_pdl_response'] = ha_datastore_pdl_response

        __props__['ha_enabled'] = ha_enabled

        __props__['ha_heartbeat_datastore_ids'] = ha_heartbeat_datastore_ids

        __props__['ha_heartbeat_datastore_policy'] = ha_heartbeat_datastore_policy

        __props__['ha_host_isolation_response'] = ha_host_isolation_response

        __props__['ha_host_monitoring'] = ha_host_monitoring

        __props__['ha_vm_component_protection'] = ha_vm_component_protection

        __props__['ha_vm_dependency_restart_condition'] = ha_vm_dependency_restart_condition

        __props__['ha_vm_failure_interval'] = ha_vm_failure_interval

        __props__['ha_vm_maximum_failure_window'] = ha_vm_maximum_failure_window

        __props__['ha_vm_maximum_resets'] = ha_vm_maximum_resets

        __props__['ha_vm_minimum_uptime'] = ha_vm_minimum_uptime

        __props__['ha_vm_monitoring'] = ha_vm_monitoring

        __props__['ha_vm_restart_additional_delay'] = ha_vm_restart_additional_delay

        __props__['ha_vm_restart_priority'] = ha_vm_restart_priority

        __props__['ha_vm_restart_timeout'] = ha_vm_restart_timeout

        __props__['host_cluster_exit_timeout'] = host_cluster_exit_timeout

        __props__['host_system_ids'] = host_system_ids

        __props__['name'] = name

        __props__['proactive_ha_automation_level'] = proactive_ha_automation_level

        __props__['proactive_ha_enabled'] = proactive_ha_enabled

        __props__['proactive_ha_moderate_remediation'] = proactive_ha_moderate_remediation

        __props__['proactive_ha_provider_ids'] = proactive_ha_provider_ids

        __props__['proactive_ha_severe_remediation'] = proactive_ha_severe_remediation

        __props__['tags'] = tags

        __props__['resource_pool_id'] = None

        super(ComputeCluster, __self__).__init__(
            'vsphere:index/computeCluster:ComputeCluster',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `.HaVmOverride` resource can be used to add an override for
// vSphere HA settings on a cluster for a specific virtual machine. With this
// resource, one can control specific HA settings so that they are different than
// the cluster default, accommodating the needs of that specific virtual machine,
// while not affecting the rest of the cluster.
// 
// For more information on vSphere HA, see [this page][ref-vsphere-ha-clusters].
// 
// [ref-vsphere-ha-clusters]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.avail.doc/GUID-5432CA24-14F1-44E3-87FB-61D937831CF6.html
// 
// > **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/ha_vm_override.html.markdown.
type HaVmOverride struct {
	s *pulumi.ResourceState
}

// NewHaVmOverride registers a new resource with the given unique name, arguments, and options.
func NewHaVmOverride(ctx *pulumi.Context,
	name string, args *HaVmOverrideArgs, opts ...pulumi.ResourceOpt) (*HaVmOverride, error) {
	if args == nil || args.ComputeClusterId == nil {
		return nil, errors.New("missing required argument 'ComputeClusterId'")
	}
	if args == nil || args.VirtualMachineId == nil {
		return nil, errors.New("missing required argument 'VirtualMachineId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["computeClusterId"] = nil
		inputs["haDatastoreApdRecoveryAction"] = nil
		inputs["haDatastoreApdResponse"] = nil
		inputs["haDatastoreApdResponseDelay"] = nil
		inputs["haDatastorePdlResponse"] = nil
		inputs["haHostIsolationResponse"] = nil
		inputs["haVmFailureInterval"] = nil
		inputs["haVmMaximumFailureWindow"] = nil
		inputs["haVmMaximumResets"] = nil
		inputs["haVmMinimumUptime"] = nil
		inputs["haVmMonitoring"] = nil
		inputs["haVmMonitoringUseClusterDefaults"] = nil
		inputs["haVmRestartPriority"] = nil
		inputs["haVmRestartTimeout"] = nil
		inputs["virtualMachineId"] = nil
	} else {
		inputs["computeClusterId"] = args.ComputeClusterId
		inputs["haDatastoreApdRecoveryAction"] = args.HaDatastoreApdRecoveryAction
		inputs["haDatastoreApdResponse"] = args.HaDatastoreApdResponse
		inputs["haDatastoreApdResponseDelay"] = args.HaDatastoreApdResponseDelay
		inputs["haDatastorePdlResponse"] = args.HaDatastorePdlResponse
		inputs["haHostIsolationResponse"] = args.HaHostIsolationResponse
		inputs["haVmFailureInterval"] = args.HaVmFailureInterval
		inputs["haVmMaximumFailureWindow"] = args.HaVmMaximumFailureWindow
		inputs["haVmMaximumResets"] = args.HaVmMaximumResets
		inputs["haVmMinimumUptime"] = args.HaVmMinimumUptime
		inputs["haVmMonitoring"] = args.HaVmMonitoring
		inputs["haVmMonitoringUseClusterDefaults"] = args.HaVmMonitoringUseClusterDefaults
		inputs["haVmRestartPriority"] = args.HaVmRestartPriority
		inputs["haVmRestartTimeout"] = args.HaVmRestartTimeout
		inputs["virtualMachineId"] = args.VirtualMachineId
	}
	s, err := ctx.RegisterResource("vsphere:index/haVmOverride:HaVmOverride", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &HaVmOverride{s: s}, nil
}

// GetHaVmOverride gets an existing HaVmOverride resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHaVmOverride(ctx *pulumi.Context,
	name string, id pulumi.ID, state *HaVmOverrideState, opts ...pulumi.ResourceOpt) (*HaVmOverride, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["computeClusterId"] = state.ComputeClusterId
		inputs["haDatastoreApdRecoveryAction"] = state.HaDatastoreApdRecoveryAction
		inputs["haDatastoreApdResponse"] = state.HaDatastoreApdResponse
		inputs["haDatastoreApdResponseDelay"] = state.HaDatastoreApdResponseDelay
		inputs["haDatastorePdlResponse"] = state.HaDatastorePdlResponse
		inputs["haHostIsolationResponse"] = state.HaHostIsolationResponse
		inputs["haVmFailureInterval"] = state.HaVmFailureInterval
		inputs["haVmMaximumFailureWindow"] = state.HaVmMaximumFailureWindow
		inputs["haVmMaximumResets"] = state.HaVmMaximumResets
		inputs["haVmMinimumUptime"] = state.HaVmMinimumUptime
		inputs["haVmMonitoring"] = state.HaVmMonitoring
		inputs["haVmMonitoringUseClusterDefaults"] = state.HaVmMonitoringUseClusterDefaults
		inputs["haVmRestartPriority"] = state.HaVmRestartPriority
		inputs["haVmRestartTimeout"] = state.HaVmRestartTimeout
		inputs["virtualMachineId"] = state.VirtualMachineId
	}
	s, err := ctx.ReadResource("vsphere:index/haVmOverride:HaVmOverride", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &HaVmOverride{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *HaVmOverride) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *HaVmOverride) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The [managed object reference
// ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
// resource if changed.
func (r *HaVmOverride) ComputeClusterId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["computeClusterId"])
}

// Controls the action to take
// on this virtual machine if an APD status on an affected datastore clears in
// the middle of an APD event. Can be one of `useClusterDefault`, `none` or
// `reset`.  Default: `useClusterDefault`.
// <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
func (r *HaVmOverride) HaDatastoreApdRecoveryAction() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["haDatastoreApdRecoveryAction"])
}

// Controls the action to take on this
// virtual machine when the cluster has detected loss to all paths to a relevant
// datastore. Can be one of `clusterDefault`, `disabled`, `warning`,
// `restartConservative`, or `restartAggressive`.  Default: `clusterDefault`.
// <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
func (r *HaVmOverride) HaDatastoreApdResponse() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["haDatastoreApdResponse"])
}

// Controls the delay in minutes
// to wait after an APD timeout event to execute the response action defined in
// `haDatastoreApdResponse`. Use `-1` to use
// the cluster default. Default: `-1`.
// <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
func (r *HaVmOverride) HaDatastoreApdResponseDelay() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["haDatastoreApdResponseDelay"])
}

// Controls the action to take on this
// virtual machine when the cluster has detected a permanent device loss to a
// relevant datastore. Can be one of `clusterDefault`, `disabled`, `warning`, or
// `restartAggressive`. Default: `clusterDefault`.
// <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
func (r *HaVmOverride) HaDatastorePdlResponse() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["haDatastorePdlResponse"])
}

// The action to take on this virtual
// machine when a host has detected that it has been isolated from the rest of
// the cluster. Can be one of `clusterIsolationResponse`, `none`, `powerOff`, or
// `shutdown`. Default: `clusterIsolationResponse`.
func (r *HaVmOverride) HaHostIsolationResponse() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["haHostIsolationResponse"])
}

// If a heartbeat from this virtual
// machine is not received within this configured interval, the virtual machine
// is marked as failed. The value is in seconds. Default: `30`.
func (r *HaVmOverride) HaVmFailureInterval() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["haVmFailureInterval"])
}

// The length of the reset window in
// which `haVmMaximumResets` can operate. When this
// window expires, no more resets are attempted regardless of the setting
// configured in `haVmMaximumResets`. `-1` means no window, meaning an
// unlimited reset time is allotted. The value is specified in seconds. Default:
// `-1` (no window).
func (r *HaVmOverride) HaVmMaximumFailureWindow() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["haVmMaximumFailureWindow"])
}

// The maximum number of resets that HA will
// perform to this virtual machine when responding to a failure event. Default:
// `3`
func (r *HaVmOverride) HaVmMaximumResets() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["haVmMaximumResets"])
}

// The time, in seconds, that HA waits after
// powering on this virtual machine before monitoring for heartbeats. Default:
// `120` (2 minutes).
func (r *HaVmOverride) HaVmMinimumUptime() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["haVmMinimumUptime"])
}

// The type of virtual machine monitoring to use
// when HA is enabled in the cluster. Can be one of `vmMonitoringDisabled`,
// `vmMonitoringOnly`, or `vmAndAppMonitoring`. Default: `vmMonitoringDisabled`.
func (r *HaVmOverride) HaVmMonitoring() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["haVmMonitoring"])
}

// Determines whether or
// not the cluster's default settings or the VM override settings specified in
// this resource are used for virtual machine monitoring. The default is `true`
// (use cluster defaults) - set to `false` to have overrides take effect.
func (r *HaVmOverride) HaVmMonitoringUseClusterDefaults() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["haVmMonitoringUseClusterDefaults"])
}

// The restart priority for the virtual
// machine when vSphere detects a host failure. Can be one of
// `clusterRestartPriority`, `lowest`, `low`, `medium`, `high`, or `highest`.
// Default: `clusterRestartPriority`.
func (r *HaVmOverride) HaVmRestartPriority() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["haVmRestartPriority"])
}

// The maximum time, in seconds, that
// vSphere HA will wait for this virtual machine to be ready. Use `-1` to
// specify the cluster default.  Default: `-1`.
// <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
func (r *HaVmOverride) HaVmRestartTimeout() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["haVmRestartTimeout"])
}

// The UUID of the virtual machine to create
// the override for.  Forces a new resource if changed.
func (r *HaVmOverride) VirtualMachineId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["virtualMachineId"])
}

// Input properties used for looking up and filtering HaVmOverride resources.
type HaVmOverrideState struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId interface{}
	// Controls the action to take
	// on this virtual machine if an APD status on an affected datastore clears in
	// the middle of an APD event. Can be one of `useClusterDefault`, `none` or
	// `reset`.  Default: `useClusterDefault`.
	// <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
	HaDatastoreApdRecoveryAction interface{}
	// Controls the action to take on this
	// virtual machine when the cluster has detected loss to all paths to a relevant
	// datastore. Can be one of `clusterDefault`, `disabled`, `warning`,
	// `restartConservative`, or `restartAggressive`.  Default: `clusterDefault`.
	// <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
	HaDatastoreApdResponse interface{}
	// Controls the delay in minutes
	// to wait after an APD timeout event to execute the response action defined in
	// `haDatastoreApdResponse`. Use `-1` to use
	// the cluster default. Default: `-1`.
	// <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
	HaDatastoreApdResponseDelay interface{}
	// Controls the action to take on this
	// virtual machine when the cluster has detected a permanent device loss to a
	// relevant datastore. Can be one of `clusterDefault`, `disabled`, `warning`, or
	// `restartAggressive`. Default: `clusterDefault`.
	// <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
	HaDatastorePdlResponse interface{}
	// The action to take on this virtual
	// machine when a host has detected that it has been isolated from the rest of
	// the cluster. Can be one of `clusterIsolationResponse`, `none`, `powerOff`, or
	// `shutdown`. Default: `clusterIsolationResponse`.
	HaHostIsolationResponse interface{}
	// If a heartbeat from this virtual
	// machine is not received within this configured interval, the virtual machine
	// is marked as failed. The value is in seconds. Default: `30`.
	HaVmFailureInterval interface{}
	// The length of the reset window in
	// which `haVmMaximumResets` can operate. When this
	// window expires, no more resets are attempted regardless of the setting
	// configured in `haVmMaximumResets`. `-1` means no window, meaning an
	// unlimited reset time is allotted. The value is specified in seconds. Default:
	// `-1` (no window).
	HaVmMaximumFailureWindow interface{}
	// The maximum number of resets that HA will
	// perform to this virtual machine when responding to a failure event. Default:
	// `3`
	HaVmMaximumResets interface{}
	// The time, in seconds, that HA waits after
	// powering on this virtual machine before monitoring for heartbeats. Default:
	// `120` (2 minutes).
	HaVmMinimumUptime interface{}
	// The type of virtual machine monitoring to use
	// when HA is enabled in the cluster. Can be one of `vmMonitoringDisabled`,
	// `vmMonitoringOnly`, or `vmAndAppMonitoring`. Default: `vmMonitoringDisabled`.
	HaVmMonitoring interface{}
	// Determines whether or
	// not the cluster's default settings or the VM override settings specified in
	// this resource are used for virtual machine monitoring. The default is `true`
	// (use cluster defaults) - set to `false` to have overrides take effect.
	HaVmMonitoringUseClusterDefaults interface{}
	// The restart priority for the virtual
	// machine when vSphere detects a host failure. Can be one of
	// `clusterRestartPriority`, `lowest`, `low`, `medium`, `high`, or `highest`.
	// Default: `clusterRestartPriority`.
	HaVmRestartPriority interface{}
	// The maximum time, in seconds, that
	// vSphere HA will wait for this virtual machine to be ready. Use `-1` to
	// specify the cluster default.  Default: `-1`.
	// <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
	HaVmRestartTimeout interface{}
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId interface{}
}

// The set of arguments for constructing a HaVmOverride resource.
type HaVmOverrideArgs struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId interface{}
	// Controls the action to take
	// on this virtual machine if an APD status on an affected datastore clears in
	// the middle of an APD event. Can be one of `useClusterDefault`, `none` or
	// `reset`.  Default: `useClusterDefault`.
	// <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
	HaDatastoreApdRecoveryAction interface{}
	// Controls the action to take on this
	// virtual machine when the cluster has detected loss to all paths to a relevant
	// datastore. Can be one of `clusterDefault`, `disabled`, `warning`,
	// `restartConservative`, or `restartAggressive`.  Default: `clusterDefault`.
	// <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
	HaDatastoreApdResponse interface{}
	// Controls the delay in minutes
	// to wait after an APD timeout event to execute the response action defined in
	// `haDatastoreApdResponse`. Use `-1` to use
	// the cluster default. Default: `-1`.
	// <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
	HaDatastoreApdResponseDelay interface{}
	// Controls the action to take on this
	// virtual machine when the cluster has detected a permanent device loss to a
	// relevant datastore. Can be one of `clusterDefault`, `disabled`, `warning`, or
	// `restartAggressive`. Default: `clusterDefault`.
	// <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
	HaDatastorePdlResponse interface{}
	// The action to take on this virtual
	// machine when a host has detected that it has been isolated from the rest of
	// the cluster. Can be one of `clusterIsolationResponse`, `none`, `powerOff`, or
	// `shutdown`. Default: `clusterIsolationResponse`.
	HaHostIsolationResponse interface{}
	// If a heartbeat from this virtual
	// machine is not received within this configured interval, the virtual machine
	// is marked as failed. The value is in seconds. Default: `30`.
	HaVmFailureInterval interface{}
	// The length of the reset window in
	// which `haVmMaximumResets` can operate. When this
	// window expires, no more resets are attempted regardless of the setting
	// configured in `haVmMaximumResets`. `-1` means no window, meaning an
	// unlimited reset time is allotted. The value is specified in seconds. Default:
	// `-1` (no window).
	HaVmMaximumFailureWindow interface{}
	// The maximum number of resets that HA will
	// perform to this virtual machine when responding to a failure event. Default:
	// `3`
	HaVmMaximumResets interface{}
	// The time, in seconds, that HA waits after
	// powering on this virtual machine before monitoring for heartbeats. Default:
	// `120` (2 minutes).
	HaVmMinimumUptime interface{}
	// The type of virtual machine monitoring to use
	// when HA is enabled in the cluster. Can be one of `vmMonitoringDisabled`,
	// `vmMonitoringOnly`, or `vmAndAppMonitoring`. Default: `vmMonitoringDisabled`.
	HaVmMonitoring interface{}
	// Determines whether or
	// not the cluster's default settings or the VM override settings specified in
	// this resource are used for virtual machine monitoring. The default is `true`
	// (use cluster defaults) - set to `false` to have overrides take effect.
	HaVmMonitoringUseClusterDefaults interface{}
	// The restart priority for the virtual
	// machine when vSphere detects a host failure. Can be one of
	// `clusterRestartPriority`, `lowest`, `low`, `medium`, `high`, or `highest`.
	// Default: `clusterRestartPriority`.
	HaVmRestartPriority interface{}
	// The maximum time, in seconds, that
	// vSphere HA will wait for this virtual machine to be ready. Use `-1` to
	// specify the cluster default.  Default: `-1`.
	// <sup>[\*][tf-vsphere-cluster-resource-version-restrictions]</sup>
	HaVmRestartTimeout interface{}
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId interface{}
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HaVmOverrideArgs extends com.pulumi.resources.ResourceArgs {

    public static final HaVmOverrideArgs Empty = new HaVmOverrideArgs();

    /**
     * The managed object ID of the cluster.
     * 
     */
    @Import(name="computeClusterId", required=true)
    private Output<String> computeClusterId;

    /**
     * @return The managed object ID of the cluster.
     * 
     */
    public Output<String> computeClusterId() {
        return this.computeClusterId;
    }

    /**
     * Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an
     * APD event. Can be one of useClusterDefault, none or reset.
     * 
     */
    @Import(name="haDatastoreApdRecoveryAction")
    private @Nullable Output<String> haDatastoreApdRecoveryAction;

    /**
     * @return Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an
     * APD event. Can be one of useClusterDefault, none or reset.
     * 
     */
    public Optional<Output<String>> haDatastoreApdRecoveryAction() {
        return Optional.ofNullable(this.haDatastoreApdRecoveryAction);
    }

    /**
     * Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant
     * datastore. Can be one of clusterDefault, disabled, warning, restartConservative, or restartAggressive.
     * 
     */
    @Import(name="haDatastoreApdResponse")
    private @Nullable Output<String> haDatastoreApdResponse;

    /**
     * @return Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant
     * datastore. Can be one of clusterDefault, disabled, warning, restartConservative, or restartAggressive.
     * 
     */
    public Optional<Output<String>> haDatastoreApdResponse() {
        return Optional.ofNullable(this.haDatastoreApdResponse);
    }

    /**
     * Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in
     * ha_datastore_apd_response. Specify -1 to use the cluster setting.
     * 
     */
    @Import(name="haDatastoreApdResponseDelay")
    private @Nullable Output<Integer> haDatastoreApdResponseDelay;

    /**
     * @return Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in
     * ha_datastore_apd_response. Specify -1 to use the cluster setting.
     * 
     */
    public Optional<Output<Integer>> haDatastoreApdResponseDelay() {
        return Optional.ofNullable(this.haDatastoreApdResponseDelay);
    }

    /**
     * Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant
     * datastore. Can be one of clusterDefault, disabled, warning, or restartAggressive.
     * 
     */
    @Import(name="haDatastorePdlResponse")
    private @Nullable Output<String> haDatastorePdlResponse;

    /**
     * @return Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant
     * datastore. Can be one of clusterDefault, disabled, warning, or restartAggressive.
     * 
     */
    public Optional<Output<String>> haDatastorePdlResponse() {
        return Optional.ofNullable(this.haDatastorePdlResponse);
    }

    /**
     * The action to take on this virtual machine when a host is isolated from the rest of the cluster. Can be one of
     * clusterIsolationResponse, none, powerOff, or shutdown.
     * 
     */
    @Import(name="haHostIsolationResponse")
    private @Nullable Output<String> haHostIsolationResponse;

    /**
     * @return The action to take on this virtual machine when a host is isolated from the rest of the cluster. Can be one of
     * clusterIsolationResponse, none, powerOff, or shutdown.
     * 
     */
    public Optional<Output<String>> haHostIsolationResponse() {
        return Optional.ofNullable(this.haHostIsolationResponse);
    }

    /**
     * If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked
     * as failed. The value is in seconds.
     * 
     */
    @Import(name="haVmFailureInterval")
    private @Nullable Output<Integer> haVmFailureInterval;

    /**
     * @return If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked
     * as failed. The value is in seconds.
     * 
     */
    public Optional<Output<Integer>> haVmFailureInterval() {
        return Optional.ofNullable(this.haVmFailureInterval);
    }

    /**
     * The length of the reset window in which ha_vm_maximum_resets can operate. When this window expires, no more resets are
     * attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset
     * time is allotted.
     * 
     */
    @Import(name="haVmMaximumFailureWindow")
    private @Nullable Output<Integer> haVmMaximumFailureWindow;

    /**
     * @return The length of the reset window in which ha_vm_maximum_resets can operate. When this window expires, no more resets are
     * attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset
     * time is allotted.
     * 
     */
    public Optional<Output<Integer>> haVmMaximumFailureWindow() {
        return Optional.ofNullable(this.haVmMaximumFailureWindow);
    }

    /**
     * The maximum number of resets that HA will perform to this virtual machine when responding to a failure event.
     * 
     */
    @Import(name="haVmMaximumResets")
    private @Nullable Output<Integer> haVmMaximumResets;

    /**
     * @return The maximum number of resets that HA will perform to this virtual machine when responding to a failure event.
     * 
     */
    public Optional<Output<Integer>> haVmMaximumResets() {
        return Optional.ofNullable(this.haVmMaximumResets);
    }

    /**
     * The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats.
     * 
     */
    @Import(name="haVmMinimumUptime")
    private @Nullable Output<Integer> haVmMinimumUptime;

    /**
     * @return The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats.
     * 
     */
    public Optional<Output<Integer>> haVmMinimumUptime() {
        return Optional.ofNullable(this.haVmMinimumUptime);
    }

    /**
     * The type of virtual machine monitoring to use for this virtual machine. Can be one of vmMonitoringDisabled,
     * vmMonitoringOnly, or vmAndAppMonitoring.
     * 
     */
    @Import(name="haVmMonitoring")
    private @Nullable Output<String> haVmMonitoring;

    /**
     * @return The type of virtual machine monitoring to use for this virtual machine. Can be one of vmMonitoringDisabled,
     * vmMonitoringOnly, or vmAndAppMonitoring.
     * 
     */
    public Optional<Output<String>> haVmMonitoring() {
        return Optional.ofNullable(this.haVmMonitoring);
    }

    /**
     * Determines whether or not the cluster&#39;s default settings or the VM override settings specified in this resource are used
     * for virtual machine monitoring. The default is true (use cluster defaults) - set to false to have overrides take effect.
     * 
     */
    @Import(name="haVmMonitoringUseClusterDefaults")
    private @Nullable Output<Boolean> haVmMonitoringUseClusterDefaults;

    /**
     * @return Determines whether or not the cluster&#39;s default settings or the VM override settings specified in this resource are used
     * for virtual machine monitoring. The default is true (use cluster defaults) - set to false to have overrides take effect.
     * 
     */
    public Optional<Output<Boolean>> haVmMonitoringUseClusterDefaults() {
        return Optional.ofNullable(this.haVmMonitoringUseClusterDefaults);
    }

    /**
     * The restart priority for this virtual machine when vSphere detects a host failure. Can be one of clusterRestartPriority,
     * lowest, low, medium, high, or highest.
     * 
     */
    @Import(name="haVmRestartPriority")
    private @Nullable Output<String> haVmRestartPriority;

    /**
     * @return The restart priority for this virtual machine when vSphere detects a host failure. Can be one of clusterRestartPriority,
     * lowest, low, medium, high, or highest.
     * 
     */
    public Optional<Output<String>> haVmRestartPriority() {
        return Optional.ofNullable(this.haVmRestartPriority);
    }

    /**
     * The maximum time, in seconds, that vSphere HA will wait for the virtual machine to be ready. Use -1 to use the cluster
     * default.
     * 
     */
    @Import(name="haVmRestartTimeout")
    private @Nullable Output<Integer> haVmRestartTimeout;

    /**
     * @return The maximum time, in seconds, that vSphere HA will wait for the virtual machine to be ready. Use -1 to use the cluster
     * default.
     * 
     */
    public Optional<Output<Integer>> haVmRestartTimeout() {
        return Optional.ofNullable(this.haVmRestartTimeout);
    }

    /**
     * The managed object ID of the virtual machine.
     * 
     */
    @Import(name="virtualMachineId", required=true)
    private Output<String> virtualMachineId;

    /**
     * @return The managed object ID of the virtual machine.
     * 
     */
    public Output<String> virtualMachineId() {
        return this.virtualMachineId;
    }

    private HaVmOverrideArgs() {}

    private HaVmOverrideArgs(HaVmOverrideArgs $) {
        this.computeClusterId = $.computeClusterId;
        this.haDatastoreApdRecoveryAction = $.haDatastoreApdRecoveryAction;
        this.haDatastoreApdResponse = $.haDatastoreApdResponse;
        this.haDatastoreApdResponseDelay = $.haDatastoreApdResponseDelay;
        this.haDatastorePdlResponse = $.haDatastorePdlResponse;
        this.haHostIsolationResponse = $.haHostIsolationResponse;
        this.haVmFailureInterval = $.haVmFailureInterval;
        this.haVmMaximumFailureWindow = $.haVmMaximumFailureWindow;
        this.haVmMaximumResets = $.haVmMaximumResets;
        this.haVmMinimumUptime = $.haVmMinimumUptime;
        this.haVmMonitoring = $.haVmMonitoring;
        this.haVmMonitoringUseClusterDefaults = $.haVmMonitoringUseClusterDefaults;
        this.haVmRestartPriority = $.haVmRestartPriority;
        this.haVmRestartTimeout = $.haVmRestartTimeout;
        this.virtualMachineId = $.virtualMachineId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HaVmOverrideArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HaVmOverrideArgs $;

        public Builder() {
            $ = new HaVmOverrideArgs();
        }

        public Builder(HaVmOverrideArgs defaults) {
            $ = new HaVmOverrideArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param computeClusterId The managed object ID of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder computeClusterId(Output<String> computeClusterId) {
            $.computeClusterId = computeClusterId;
            return this;
        }

        /**
         * @param computeClusterId The managed object ID of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder computeClusterId(String computeClusterId) {
            return computeClusterId(Output.of(computeClusterId));
        }

        /**
         * @param haDatastoreApdRecoveryAction Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an
         * APD event. Can be one of useClusterDefault, none or reset.
         * 
         * @return builder
         * 
         */
        public Builder haDatastoreApdRecoveryAction(@Nullable Output<String> haDatastoreApdRecoveryAction) {
            $.haDatastoreApdRecoveryAction = haDatastoreApdRecoveryAction;
            return this;
        }

        /**
         * @param haDatastoreApdRecoveryAction Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an
         * APD event. Can be one of useClusterDefault, none or reset.
         * 
         * @return builder
         * 
         */
        public Builder haDatastoreApdRecoveryAction(String haDatastoreApdRecoveryAction) {
            return haDatastoreApdRecoveryAction(Output.of(haDatastoreApdRecoveryAction));
        }

        /**
         * @param haDatastoreApdResponse Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant
         * datastore. Can be one of clusterDefault, disabled, warning, restartConservative, or restartAggressive.
         * 
         * @return builder
         * 
         */
        public Builder haDatastoreApdResponse(@Nullable Output<String> haDatastoreApdResponse) {
            $.haDatastoreApdResponse = haDatastoreApdResponse;
            return this;
        }

        /**
         * @param haDatastoreApdResponse Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant
         * datastore. Can be one of clusterDefault, disabled, warning, restartConservative, or restartAggressive.
         * 
         * @return builder
         * 
         */
        public Builder haDatastoreApdResponse(String haDatastoreApdResponse) {
            return haDatastoreApdResponse(Output.of(haDatastoreApdResponse));
        }

        /**
         * @param haDatastoreApdResponseDelay Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in
         * ha_datastore_apd_response. Specify -1 to use the cluster setting.
         * 
         * @return builder
         * 
         */
        public Builder haDatastoreApdResponseDelay(@Nullable Output<Integer> haDatastoreApdResponseDelay) {
            $.haDatastoreApdResponseDelay = haDatastoreApdResponseDelay;
            return this;
        }

        /**
         * @param haDatastoreApdResponseDelay Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in
         * ha_datastore_apd_response. Specify -1 to use the cluster setting.
         * 
         * @return builder
         * 
         */
        public Builder haDatastoreApdResponseDelay(Integer haDatastoreApdResponseDelay) {
            return haDatastoreApdResponseDelay(Output.of(haDatastoreApdResponseDelay));
        }

        /**
         * @param haDatastorePdlResponse Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant
         * datastore. Can be one of clusterDefault, disabled, warning, or restartAggressive.
         * 
         * @return builder
         * 
         */
        public Builder haDatastorePdlResponse(@Nullable Output<String> haDatastorePdlResponse) {
            $.haDatastorePdlResponse = haDatastorePdlResponse;
            return this;
        }

        /**
         * @param haDatastorePdlResponse Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant
         * datastore. Can be one of clusterDefault, disabled, warning, or restartAggressive.
         * 
         * @return builder
         * 
         */
        public Builder haDatastorePdlResponse(String haDatastorePdlResponse) {
            return haDatastorePdlResponse(Output.of(haDatastorePdlResponse));
        }

        /**
         * @param haHostIsolationResponse The action to take on this virtual machine when a host is isolated from the rest of the cluster. Can be one of
         * clusterIsolationResponse, none, powerOff, or shutdown.
         * 
         * @return builder
         * 
         */
        public Builder haHostIsolationResponse(@Nullable Output<String> haHostIsolationResponse) {
            $.haHostIsolationResponse = haHostIsolationResponse;
            return this;
        }

        /**
         * @param haHostIsolationResponse The action to take on this virtual machine when a host is isolated from the rest of the cluster. Can be one of
         * clusterIsolationResponse, none, powerOff, or shutdown.
         * 
         * @return builder
         * 
         */
        public Builder haHostIsolationResponse(String haHostIsolationResponse) {
            return haHostIsolationResponse(Output.of(haHostIsolationResponse));
        }

        /**
         * @param haVmFailureInterval If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked
         * as failed. The value is in seconds.
         * 
         * @return builder
         * 
         */
        public Builder haVmFailureInterval(@Nullable Output<Integer> haVmFailureInterval) {
            $.haVmFailureInterval = haVmFailureInterval;
            return this;
        }

        /**
         * @param haVmFailureInterval If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked
         * as failed. The value is in seconds.
         * 
         * @return builder
         * 
         */
        public Builder haVmFailureInterval(Integer haVmFailureInterval) {
            return haVmFailureInterval(Output.of(haVmFailureInterval));
        }

        /**
         * @param haVmMaximumFailureWindow The length of the reset window in which ha_vm_maximum_resets can operate. When this window expires, no more resets are
         * attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset
         * time is allotted.
         * 
         * @return builder
         * 
         */
        public Builder haVmMaximumFailureWindow(@Nullable Output<Integer> haVmMaximumFailureWindow) {
            $.haVmMaximumFailureWindow = haVmMaximumFailureWindow;
            return this;
        }

        /**
         * @param haVmMaximumFailureWindow The length of the reset window in which ha_vm_maximum_resets can operate. When this window expires, no more resets are
         * attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset
         * time is allotted.
         * 
         * @return builder
         * 
         */
        public Builder haVmMaximumFailureWindow(Integer haVmMaximumFailureWindow) {
            return haVmMaximumFailureWindow(Output.of(haVmMaximumFailureWindow));
        }

        /**
         * @param haVmMaximumResets The maximum number of resets that HA will perform to this virtual machine when responding to a failure event.
         * 
         * @return builder
         * 
         */
        public Builder haVmMaximumResets(@Nullable Output<Integer> haVmMaximumResets) {
            $.haVmMaximumResets = haVmMaximumResets;
            return this;
        }

        /**
         * @param haVmMaximumResets The maximum number of resets that HA will perform to this virtual machine when responding to a failure event.
         * 
         * @return builder
         * 
         */
        public Builder haVmMaximumResets(Integer haVmMaximumResets) {
            return haVmMaximumResets(Output.of(haVmMaximumResets));
        }

        /**
         * @param haVmMinimumUptime The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats.
         * 
         * @return builder
         * 
         */
        public Builder haVmMinimumUptime(@Nullable Output<Integer> haVmMinimumUptime) {
            $.haVmMinimumUptime = haVmMinimumUptime;
            return this;
        }

        /**
         * @param haVmMinimumUptime The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats.
         * 
         * @return builder
         * 
         */
        public Builder haVmMinimumUptime(Integer haVmMinimumUptime) {
            return haVmMinimumUptime(Output.of(haVmMinimumUptime));
        }

        /**
         * @param haVmMonitoring The type of virtual machine monitoring to use for this virtual machine. Can be one of vmMonitoringDisabled,
         * vmMonitoringOnly, or vmAndAppMonitoring.
         * 
         * @return builder
         * 
         */
        public Builder haVmMonitoring(@Nullable Output<String> haVmMonitoring) {
            $.haVmMonitoring = haVmMonitoring;
            return this;
        }

        /**
         * @param haVmMonitoring The type of virtual machine monitoring to use for this virtual machine. Can be one of vmMonitoringDisabled,
         * vmMonitoringOnly, or vmAndAppMonitoring.
         * 
         * @return builder
         * 
         */
        public Builder haVmMonitoring(String haVmMonitoring) {
            return haVmMonitoring(Output.of(haVmMonitoring));
        }

        /**
         * @param haVmMonitoringUseClusterDefaults Determines whether or not the cluster&#39;s default settings or the VM override settings specified in this resource are used
         * for virtual machine monitoring. The default is true (use cluster defaults) - set to false to have overrides take effect.
         * 
         * @return builder
         * 
         */
        public Builder haVmMonitoringUseClusterDefaults(@Nullable Output<Boolean> haVmMonitoringUseClusterDefaults) {
            $.haVmMonitoringUseClusterDefaults = haVmMonitoringUseClusterDefaults;
            return this;
        }

        /**
         * @param haVmMonitoringUseClusterDefaults Determines whether or not the cluster&#39;s default settings or the VM override settings specified in this resource are used
         * for virtual machine monitoring. The default is true (use cluster defaults) - set to false to have overrides take effect.
         * 
         * @return builder
         * 
         */
        public Builder haVmMonitoringUseClusterDefaults(Boolean haVmMonitoringUseClusterDefaults) {
            return haVmMonitoringUseClusterDefaults(Output.of(haVmMonitoringUseClusterDefaults));
        }

        /**
         * @param haVmRestartPriority The restart priority for this virtual machine when vSphere detects a host failure. Can be one of clusterRestartPriority,
         * lowest, low, medium, high, or highest.
         * 
         * @return builder
         * 
         */
        public Builder haVmRestartPriority(@Nullable Output<String> haVmRestartPriority) {
            $.haVmRestartPriority = haVmRestartPriority;
            return this;
        }

        /**
         * @param haVmRestartPriority The restart priority for this virtual machine when vSphere detects a host failure. Can be one of clusterRestartPriority,
         * lowest, low, medium, high, or highest.
         * 
         * @return builder
         * 
         */
        public Builder haVmRestartPriority(String haVmRestartPriority) {
            return haVmRestartPriority(Output.of(haVmRestartPriority));
        }

        /**
         * @param haVmRestartTimeout The maximum time, in seconds, that vSphere HA will wait for the virtual machine to be ready. Use -1 to use the cluster
         * default.
         * 
         * @return builder
         * 
         */
        public Builder haVmRestartTimeout(@Nullable Output<Integer> haVmRestartTimeout) {
            $.haVmRestartTimeout = haVmRestartTimeout;
            return this;
        }

        /**
         * @param haVmRestartTimeout The maximum time, in seconds, that vSphere HA will wait for the virtual machine to be ready. Use -1 to use the cluster
         * default.
         * 
         * @return builder
         * 
         */
        public Builder haVmRestartTimeout(Integer haVmRestartTimeout) {
            return haVmRestartTimeout(Output.of(haVmRestartTimeout));
        }

        /**
         * @param virtualMachineId The managed object ID of the virtual machine.
         * 
         * @return builder
         * 
         */
        public Builder virtualMachineId(Output<String> virtualMachineId) {
            $.virtualMachineId = virtualMachineId;
            return this;
        }

        /**
         * @param virtualMachineId The managed object ID of the virtual machine.
         * 
         * @return builder
         * 
         */
        public Builder virtualMachineId(String virtualMachineId) {
            return virtualMachineId(Output.of(virtualMachineId));
        }

        public HaVmOverrideArgs build() {
            if ($.computeClusterId == null) {
                throw new MissingRequiredPropertyException("HaVmOverrideArgs", "computeClusterId");
            }
            if ($.virtualMachineId == null) {
                throw new MissingRequiredPropertyException("HaVmOverrideArgs", "virtualMachineId");
            }
            return $;
        }
    }

}

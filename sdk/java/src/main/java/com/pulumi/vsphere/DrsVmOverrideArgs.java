// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DrsVmOverrideArgs extends com.pulumi.resources.ResourceArgs {

    public static final DrsVmOverrideArgs Empty = new DrsVmOverrideArgs();

    /**
     * The managed object reference
     * ID of the cluster to put the override in.  Forces a new
     * resource if changed.
     * 
     */
    @Import(name="computeClusterId", required=true)
    private Output<String> computeClusterId;

    /**
     * @return The managed object reference
     * ID of the cluster to put the override in.  Forces a new
     * resource if changed.
     * 
     */
    public Output<String> computeClusterId() {
        return this.computeClusterId;
    }

    /**
     * Overrides the automation level for this virtual
     * machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
     * `fullyAutomated`. Default: `manual`.
     * 
     * &gt; **NOTE:** Using this resource _always_ implies an override, even if one of
     * `drs_enabled` or `drs_automation_level` is omitted. Take note of the defaults
     * for both options.
     * 
     */
    @Import(name="drsAutomationLevel")
    private @Nullable Output<String> drsAutomationLevel;

    /**
     * @return Overrides the automation level for this virtual
     * machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
     * `fullyAutomated`. Default: `manual`.
     * 
     * &gt; **NOTE:** Using this resource _always_ implies an override, even if one of
     * `drs_enabled` or `drs_automation_level` is omitted. Take note of the defaults
     * for both options.
     * 
     */
    public Optional<Output<String>> drsAutomationLevel() {
        return Optional.ofNullable(this.drsAutomationLevel);
    }

    /**
     * Overrides the default DRS setting for this virtual
     * machine. Can be either `true` or `false`. Default: `false`.
     * 
     */
    @Import(name="drsEnabled")
    private @Nullable Output<Boolean> drsEnabled;

    /**
     * @return Overrides the default DRS setting for this virtual
     * machine. Can be either `true` or `false`. Default: `false`.
     * 
     */
    public Optional<Output<Boolean>> drsEnabled() {
        return Optional.ofNullable(this.drsEnabled);
    }

    /**
     * The UUID of the virtual machine to create
     * the override for.  Forces a new resource if changed.
     * 
     */
    @Import(name="virtualMachineId", required=true)
    private Output<String> virtualMachineId;

    /**
     * @return The UUID of the virtual machine to create
     * the override for.  Forces a new resource if changed.
     * 
     */
    public Output<String> virtualMachineId() {
        return this.virtualMachineId;
    }

    private DrsVmOverrideArgs() {}

    private DrsVmOverrideArgs(DrsVmOverrideArgs $) {
        this.computeClusterId = $.computeClusterId;
        this.drsAutomationLevel = $.drsAutomationLevel;
        this.drsEnabled = $.drsEnabled;
        this.virtualMachineId = $.virtualMachineId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DrsVmOverrideArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DrsVmOverrideArgs $;

        public Builder() {
            $ = new DrsVmOverrideArgs();
        }

        public Builder(DrsVmOverrideArgs defaults) {
            $ = new DrsVmOverrideArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param computeClusterId The managed object reference
         * ID of the cluster to put the override in.  Forces a new
         * resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder computeClusterId(Output<String> computeClusterId) {
            $.computeClusterId = computeClusterId;
            return this;
        }

        /**
         * @param computeClusterId The managed object reference
         * ID of the cluster to put the override in.  Forces a new
         * resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder computeClusterId(String computeClusterId) {
            return computeClusterId(Output.of(computeClusterId));
        }

        /**
         * @param drsAutomationLevel Overrides the automation level for this virtual
         * machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
         * `fullyAutomated`. Default: `manual`.
         * 
         * &gt; **NOTE:** Using this resource _always_ implies an override, even if one of
         * `drs_enabled` or `drs_automation_level` is omitted. Take note of the defaults
         * for both options.
         * 
         * @return builder
         * 
         */
        public Builder drsAutomationLevel(@Nullable Output<String> drsAutomationLevel) {
            $.drsAutomationLevel = drsAutomationLevel;
            return this;
        }

        /**
         * @param drsAutomationLevel Overrides the automation level for this virtual
         * machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
         * `fullyAutomated`. Default: `manual`.
         * 
         * &gt; **NOTE:** Using this resource _always_ implies an override, even if one of
         * `drs_enabled` or `drs_automation_level` is omitted. Take note of the defaults
         * for both options.
         * 
         * @return builder
         * 
         */
        public Builder drsAutomationLevel(String drsAutomationLevel) {
            return drsAutomationLevel(Output.of(drsAutomationLevel));
        }

        /**
         * @param drsEnabled Overrides the default DRS setting for this virtual
         * machine. Can be either `true` or `false`. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder drsEnabled(@Nullable Output<Boolean> drsEnabled) {
            $.drsEnabled = drsEnabled;
            return this;
        }

        /**
         * @param drsEnabled Overrides the default DRS setting for this virtual
         * machine. Can be either `true` or `false`. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder drsEnabled(Boolean drsEnabled) {
            return drsEnabled(Output.of(drsEnabled));
        }

        /**
         * @param virtualMachineId The UUID of the virtual machine to create
         * the override for.  Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder virtualMachineId(Output<String> virtualMachineId) {
            $.virtualMachineId = virtualMachineId;
            return this;
        }

        /**
         * @param virtualMachineId The UUID of the virtual machine to create
         * the override for.  Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder virtualMachineId(String virtualMachineId) {
            return virtualMachineId(Output.of(virtualMachineId));
        }

        public DrsVmOverrideArgs build() {
            if ($.computeClusterId == null) {
                throw new MissingRequiredPropertyException("DrsVmOverrideArgs", "computeClusterId");
            }
            if ($.virtualMachineId == null) {
                throw new MissingRequiredPropertyException("DrsVmOverrideArgs", "virtualMachineId");
            }
            return $;
        }
    }

}

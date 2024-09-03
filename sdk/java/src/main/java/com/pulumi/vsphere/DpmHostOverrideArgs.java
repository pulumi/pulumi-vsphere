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


public final class DpmHostOverrideArgs extends com.pulumi.resources.ResourceArgs {

    public static final DpmHostOverrideArgs Empty = new DpmHostOverrideArgs();

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
     * The automation level for host power
     * operations on this host. Can be one of `manual` or `automated`. Default:
     * `manual`.
     * 
     * &gt; **NOTE:** Using this resource *always* implies an override, even if one of
     * `dpm_enabled` or `dpm_automation_level` is omitted. Take note of the defaults
     * for both options.
     * 
     */
    @Import(name="dpmAutomationLevel")
    private @Nullable Output<String> dpmAutomationLevel;

    /**
     * @return The automation level for host power
     * operations on this host. Can be one of `manual` or `automated`. Default:
     * `manual`.
     * 
     * &gt; **NOTE:** Using this resource *always* implies an override, even if one of
     * `dpm_enabled` or `dpm_automation_level` is omitted. Take note of the defaults
     * for both options.
     * 
     */
    public Optional<Output<String>> dpmAutomationLevel() {
        return Optional.ofNullable(this.dpmAutomationLevel);
    }

    /**
     * Enable DPM support for this host. Default:
     * `false`.
     * 
     */
    @Import(name="dpmEnabled")
    private @Nullable Output<Boolean> dpmEnabled;

    /**
     * @return Enable DPM support for this host. Default:
     * `false`.
     * 
     */
    public Optional<Output<Boolean>> dpmEnabled() {
        return Optional.ofNullable(this.dpmEnabled);
    }

    /**
     * The managed object ID of the host.
     * 
     */
    @Import(name="hostSystemId", required=true)
    private Output<String> hostSystemId;

    /**
     * @return The managed object ID of the host.
     * 
     */
    public Output<String> hostSystemId() {
        return this.hostSystemId;
    }

    private DpmHostOverrideArgs() {}

    private DpmHostOverrideArgs(DpmHostOverrideArgs $) {
        this.computeClusterId = $.computeClusterId;
        this.dpmAutomationLevel = $.dpmAutomationLevel;
        this.dpmEnabled = $.dpmEnabled;
        this.hostSystemId = $.hostSystemId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DpmHostOverrideArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DpmHostOverrideArgs $;

        public Builder() {
            $ = new DpmHostOverrideArgs();
        }

        public Builder(DpmHostOverrideArgs defaults) {
            $ = new DpmHostOverrideArgs(Objects.requireNonNull(defaults));
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
         * @param dpmAutomationLevel The automation level for host power
         * operations on this host. Can be one of `manual` or `automated`. Default:
         * `manual`.
         * 
         * &gt; **NOTE:** Using this resource *always* implies an override, even if one of
         * `dpm_enabled` or `dpm_automation_level` is omitted. Take note of the defaults
         * for both options.
         * 
         * @return builder
         * 
         */
        public Builder dpmAutomationLevel(@Nullable Output<String> dpmAutomationLevel) {
            $.dpmAutomationLevel = dpmAutomationLevel;
            return this;
        }

        /**
         * @param dpmAutomationLevel The automation level for host power
         * operations on this host. Can be one of `manual` or `automated`. Default:
         * `manual`.
         * 
         * &gt; **NOTE:** Using this resource *always* implies an override, even if one of
         * `dpm_enabled` or `dpm_automation_level` is omitted. Take note of the defaults
         * for both options.
         * 
         * @return builder
         * 
         */
        public Builder dpmAutomationLevel(String dpmAutomationLevel) {
            return dpmAutomationLevel(Output.of(dpmAutomationLevel));
        }

        /**
         * @param dpmEnabled Enable DPM support for this host. Default:
         * `false`.
         * 
         * @return builder
         * 
         */
        public Builder dpmEnabled(@Nullable Output<Boolean> dpmEnabled) {
            $.dpmEnabled = dpmEnabled;
            return this;
        }

        /**
         * @param dpmEnabled Enable DPM support for this host. Default:
         * `false`.
         * 
         * @return builder
         * 
         */
        public Builder dpmEnabled(Boolean dpmEnabled) {
            return dpmEnabled(Output.of(dpmEnabled));
        }

        /**
         * @param hostSystemId The managed object ID of the host.
         * 
         * @return builder
         * 
         */
        public Builder hostSystemId(Output<String> hostSystemId) {
            $.hostSystemId = hostSystemId;
            return this;
        }

        /**
         * @param hostSystemId The managed object ID of the host.
         * 
         * @return builder
         * 
         */
        public Builder hostSystemId(String hostSystemId) {
            return hostSystemId(Output.of(hostSystemId));
        }

        public DpmHostOverrideArgs build() {
            if ($.computeClusterId == null) {
                throw new MissingRequiredPropertyException("DpmHostOverrideArgs", "computeClusterId");
            }
            if ($.hostSystemId == null) {
                throw new MissingRequiredPropertyException("DpmHostOverrideArgs", "hostSystemId");
            }
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DpmHostOverrideState extends com.pulumi.resources.ResourceArgs {

    public static final DpmHostOverrideState Empty = new DpmHostOverrideState();

    /**
     * The managed object reference
     * ID of the cluster to put the override in.  Forces a new
     * resource if changed.
     * 
     */
    @Import(name="computeClusterId")
    private @Nullable Output<String> computeClusterId;

    /**
     * @return The managed object reference
     * ID of the cluster to put the override in.  Forces a new
     * resource if changed.
     * 
     */
    public Optional<Output<String>> computeClusterId() {
        return Optional.ofNullable(this.computeClusterId);
    }

    /**
     * The automation level for host power
     * operations on this host. Can be one of `manual` or `automated`. Default:
     * `manual`.
     * 
     */
    @Import(name="dpmAutomationLevel")
    private @Nullable Output<String> dpmAutomationLevel;

    /**
     * @return The automation level for host power
     * operations on this host. Can be one of `manual` or `automated`. Default:
     * `manual`.
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
    @Import(name="hostSystemId")
    private @Nullable Output<String> hostSystemId;

    /**
     * @return The managed object ID of the host.
     * 
     */
    public Optional<Output<String>> hostSystemId() {
        return Optional.ofNullable(this.hostSystemId);
    }

    private DpmHostOverrideState() {}

    private DpmHostOverrideState(DpmHostOverrideState $) {
        this.computeClusterId = $.computeClusterId;
        this.dpmAutomationLevel = $.dpmAutomationLevel;
        this.dpmEnabled = $.dpmEnabled;
        this.hostSystemId = $.hostSystemId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DpmHostOverrideState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DpmHostOverrideState $;

        public Builder() {
            $ = new DpmHostOverrideState();
        }

        public Builder(DpmHostOverrideState defaults) {
            $ = new DpmHostOverrideState(Objects.requireNonNull(defaults));
        }

        /**
         * @param computeClusterId The managed object reference
         * ID of the cluster to put the override in.  Forces a new
         * resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder computeClusterId(@Nullable Output<String> computeClusterId) {
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
        public Builder hostSystemId(@Nullable Output<String> hostSystemId) {
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

        public DpmHostOverrideState build() {
            return $;
        }
    }

}

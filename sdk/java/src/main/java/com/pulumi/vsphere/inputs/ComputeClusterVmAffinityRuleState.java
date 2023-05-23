// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ComputeClusterVmAffinityRuleState extends com.pulumi.resources.ResourceArgs {

    public static final ComputeClusterVmAffinityRuleState Empty = new ComputeClusterVmAffinityRuleState();

    /**
     * The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     * 
     */
    @Import(name="computeClusterId")
    private @Nullable Output<String> computeClusterId;

    /**
     * @return The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     * 
     */
    public Optional<Output<String>> computeClusterId() {
        return Optional.ofNullable(this.computeClusterId);
    }

    /**
     * Enable this rule in the cluster. Default: `true`.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Enable this rule in the cluster. Default: `true`.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * When this value is `true`, prevents any virtual
     * machine operations that may violate this rule. Default: `false`.
     * 
     * &gt; **NOTE:** The namespace for rule names on this resource (defined by the
     * `name` argument) is shared with all rules in the cluster - consider
     * this when naming your rules.
     * 
     */
    @Import(name="mandatory")
    private @Nullable Output<Boolean> mandatory;

    /**
     * @return When this value is `true`, prevents any virtual
     * machine operations that may violate this rule. Default: `false`.
     * 
     * &gt; **NOTE:** The namespace for rule names on this resource (defined by the
     * `name` argument) is shared with all rules in the cluster - consider
     * this when naming your rules.
     * 
     */
    public Optional<Output<Boolean>> mandatory() {
        return Optional.ofNullable(this.mandatory);
    }

    /**
     * The name of the rule. This must be unique in the cluster.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the rule. This must be unique in the cluster.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The UUIDs of the virtual machines to run
     * on the same host together.
     * 
     */
    @Import(name="virtualMachineIds")
    private @Nullable Output<List<String>> virtualMachineIds;

    /**
     * @return The UUIDs of the virtual machines to run
     * on the same host together.
     * 
     */
    public Optional<Output<List<String>>> virtualMachineIds() {
        return Optional.ofNullable(this.virtualMachineIds);
    }

    private ComputeClusterVmAffinityRuleState() {}

    private ComputeClusterVmAffinityRuleState(ComputeClusterVmAffinityRuleState $) {
        this.computeClusterId = $.computeClusterId;
        this.enabled = $.enabled;
        this.mandatory = $.mandatory;
        this.name = $.name;
        this.virtualMachineIds = $.virtualMachineIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ComputeClusterVmAffinityRuleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ComputeClusterVmAffinityRuleState $;

        public Builder() {
            $ = new ComputeClusterVmAffinityRuleState();
        }

        public Builder(ComputeClusterVmAffinityRuleState defaults) {
            $ = new ComputeClusterVmAffinityRuleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param computeClusterId The managed object reference
         * ID of the cluster to put the group in.  Forces a new
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
         * ID of the cluster to put the group in.  Forces a new
         * resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder computeClusterId(String computeClusterId) {
            return computeClusterId(Output.of(computeClusterId));
        }

        /**
         * @param enabled Enable this rule in the cluster. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Enable this rule in the cluster. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param mandatory When this value is `true`, prevents any virtual
         * machine operations that may violate this rule. Default: `false`.
         * 
         * &gt; **NOTE:** The namespace for rule names on this resource (defined by the
         * `name` argument) is shared with all rules in the cluster - consider
         * this when naming your rules.
         * 
         * @return builder
         * 
         */
        public Builder mandatory(@Nullable Output<Boolean> mandatory) {
            $.mandatory = mandatory;
            return this;
        }

        /**
         * @param mandatory When this value is `true`, prevents any virtual
         * machine operations that may violate this rule. Default: `false`.
         * 
         * &gt; **NOTE:** The namespace for rule names on this resource (defined by the
         * `name` argument) is shared with all rules in the cluster - consider
         * this when naming your rules.
         * 
         * @return builder
         * 
         */
        public Builder mandatory(Boolean mandatory) {
            return mandatory(Output.of(mandatory));
        }

        /**
         * @param name The name of the rule. This must be unique in the cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the rule. This must be unique in the cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param virtualMachineIds The UUIDs of the virtual machines to run
         * on the same host together.
         * 
         * @return builder
         * 
         */
        public Builder virtualMachineIds(@Nullable Output<List<String>> virtualMachineIds) {
            $.virtualMachineIds = virtualMachineIds;
            return this;
        }

        /**
         * @param virtualMachineIds The UUIDs of the virtual machines to run
         * on the same host together.
         * 
         * @return builder
         * 
         */
        public Builder virtualMachineIds(List<String> virtualMachineIds) {
            return virtualMachineIds(Output.of(virtualMachineIds));
        }

        /**
         * @param virtualMachineIds The UUIDs of the virtual machines to run
         * on the same host together.
         * 
         * @return builder
         * 
         */
        public Builder virtualMachineIds(String... virtualMachineIds) {
            return virtualMachineIds(List.of(virtualMachineIds));
        }

        public ComputeClusterVmAffinityRuleState build() {
            return $;
        }
    }

}

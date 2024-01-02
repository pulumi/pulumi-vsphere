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


public final class ComputeClusterVmHostRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final ComputeClusterVmHostRuleArgs Empty = new ComputeClusterVmHostRuleArgs();

    /**
     * When this field is used, the virtual
     * machines defined in `vm_group_name` will be run on the
     * hosts defined in this host group.
     * 
     */
    @Import(name="affinityHostGroupName")
    private @Nullable Output<String> affinityHostGroupName;

    /**
     * @return When this field is used, the virtual
     * machines defined in `vm_group_name` will be run on the
     * hosts defined in this host group.
     * 
     */
    public Optional<Output<String>> affinityHostGroupName() {
        return Optional.ofNullable(this.affinityHostGroupName);
    }

    /**
     * When this field is used, the
     * virtual machines defined in `vm_group_name` will _not_ be
     * run on the hosts defined in this host group.
     * 
     */
    @Import(name="antiAffinityHostGroupName")
    private @Nullable Output<String> antiAffinityHostGroupName;

    /**
     * @return When this field is used, the
     * virtual machines defined in `vm_group_name` will _not_ be
     * run on the hosts defined in this host group.
     * 
     */
    public Optional<Output<String>> antiAffinityHostGroupName() {
        return Optional.ofNullable(this.antiAffinityHostGroupName);
    }

    /**
     * The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     * 
     */
    @Import(name="computeClusterId", required=true)
    private Output<String> computeClusterId;

    /**
     * @return The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     * 
     */
    public Output<String> computeClusterId() {
        return this.computeClusterId;
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
     * &gt; **NOTE:** One of `affinity_host_group_name` or
     * `anti_affinity_host_group_name` must be
     * defined, but not both.
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
     * &gt; **NOTE:** One of `affinity_host_group_name` or
     * `anti_affinity_host_group_name` must be
     * defined, but not both.
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
     * The name of the rule. This must be unique in the
     * cluster.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the rule. This must be unique in the
     * cluster.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The name of the virtual machine group to use
     * with this rule.
     * 
     */
    @Import(name="vmGroupName", required=true)
    private Output<String> vmGroupName;

    /**
     * @return The name of the virtual machine group to use
     * with this rule.
     * 
     */
    public Output<String> vmGroupName() {
        return this.vmGroupName;
    }

    private ComputeClusterVmHostRuleArgs() {}

    private ComputeClusterVmHostRuleArgs(ComputeClusterVmHostRuleArgs $) {
        this.affinityHostGroupName = $.affinityHostGroupName;
        this.antiAffinityHostGroupName = $.antiAffinityHostGroupName;
        this.computeClusterId = $.computeClusterId;
        this.enabled = $.enabled;
        this.mandatory = $.mandatory;
        this.name = $.name;
        this.vmGroupName = $.vmGroupName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ComputeClusterVmHostRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ComputeClusterVmHostRuleArgs $;

        public Builder() {
            $ = new ComputeClusterVmHostRuleArgs();
        }

        public Builder(ComputeClusterVmHostRuleArgs defaults) {
            $ = new ComputeClusterVmHostRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param affinityHostGroupName When this field is used, the virtual
         * machines defined in `vm_group_name` will be run on the
         * hosts defined in this host group.
         * 
         * @return builder
         * 
         */
        public Builder affinityHostGroupName(@Nullable Output<String> affinityHostGroupName) {
            $.affinityHostGroupName = affinityHostGroupName;
            return this;
        }

        /**
         * @param affinityHostGroupName When this field is used, the virtual
         * machines defined in `vm_group_name` will be run on the
         * hosts defined in this host group.
         * 
         * @return builder
         * 
         */
        public Builder affinityHostGroupName(String affinityHostGroupName) {
            return affinityHostGroupName(Output.of(affinityHostGroupName));
        }

        /**
         * @param antiAffinityHostGroupName When this field is used, the
         * virtual machines defined in `vm_group_name` will _not_ be
         * run on the hosts defined in this host group.
         * 
         * @return builder
         * 
         */
        public Builder antiAffinityHostGroupName(@Nullable Output<String> antiAffinityHostGroupName) {
            $.antiAffinityHostGroupName = antiAffinityHostGroupName;
            return this;
        }

        /**
         * @param antiAffinityHostGroupName When this field is used, the
         * virtual machines defined in `vm_group_name` will _not_ be
         * run on the hosts defined in this host group.
         * 
         * @return builder
         * 
         */
        public Builder antiAffinityHostGroupName(String antiAffinityHostGroupName) {
            return antiAffinityHostGroupName(Output.of(antiAffinityHostGroupName));
        }

        /**
         * @param computeClusterId The managed object reference
         * ID of the cluster to put the group in.  Forces a new
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
         * &gt; **NOTE:** One of `affinity_host_group_name` or
         * `anti_affinity_host_group_name` must be
         * defined, but not both.
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
         * &gt; **NOTE:** One of `affinity_host_group_name` or
         * `anti_affinity_host_group_name` must be
         * defined, but not both.
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
         * @param name The name of the rule. This must be unique in the
         * cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the rule. This must be unique in the
         * cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param vmGroupName The name of the virtual machine group to use
         * with this rule.
         * 
         * @return builder
         * 
         */
        public Builder vmGroupName(Output<String> vmGroupName) {
            $.vmGroupName = vmGroupName;
            return this;
        }

        /**
         * @param vmGroupName The name of the virtual machine group to use
         * with this rule.
         * 
         * @return builder
         * 
         */
        public Builder vmGroupName(String vmGroupName) {
            return vmGroupName(Output.of(vmGroupName));
        }

        public ComputeClusterVmHostRuleArgs build() {
            if ($.computeClusterId == null) {
                throw new MissingRequiredPropertyException("ComputeClusterVmHostRuleArgs", "computeClusterId");
            }
            if ($.vmGroupName == null) {
                throw new MissingRequiredPropertyException("ComputeClusterVmHostRuleArgs", "vmGroupName");
            }
            return $;
        }
    }

}

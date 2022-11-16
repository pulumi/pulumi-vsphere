// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ComputeClusterVmDependencyRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final ComputeClusterVmDependencyRuleArgs Empty = new ComputeClusterVmDependencyRuleArgs();

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
     * The name of the VM group that this
     * rule depends on. The VMs defined in the group specified by
     * `vm_group_name` will not be started until the VMs in this
     * group are started.
     * 
     */
    @Import(name="dependencyVmGroupName", required=true)
    private Output<String> dependencyVmGroupName;

    /**
     * @return The name of the VM group that this
     * rule depends on. The VMs defined in the group specified by
     * `vm_group_name` will not be started until the VMs in this
     * group are started.
     * 
     */
    public Output<String> dependencyVmGroupName() {
        return this.dependencyVmGroupName;
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
     */
    @Import(name="mandatory")
    private @Nullable Output<Boolean> mandatory;

    /**
     * @return When this value is `true`, prevents any virtual
     * machine operations that may violate this rule. Default: `false`.
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
     * The name of the VM group that is the subject of
     * this rule. The VMs defined in this group will not be started until the VMs in
     * the group specified by
     * `dependency_vm_group_name` are started.
     * 
     */
    @Import(name="vmGroupName", required=true)
    private Output<String> vmGroupName;

    /**
     * @return The name of the VM group that is the subject of
     * this rule. The VMs defined in this group will not be started until the VMs in
     * the group specified by
     * `dependency_vm_group_name` are started.
     * 
     */
    public Output<String> vmGroupName() {
        return this.vmGroupName;
    }

    private ComputeClusterVmDependencyRuleArgs() {}

    private ComputeClusterVmDependencyRuleArgs(ComputeClusterVmDependencyRuleArgs $) {
        this.computeClusterId = $.computeClusterId;
        this.dependencyVmGroupName = $.dependencyVmGroupName;
        this.enabled = $.enabled;
        this.mandatory = $.mandatory;
        this.name = $.name;
        this.vmGroupName = $.vmGroupName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ComputeClusterVmDependencyRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ComputeClusterVmDependencyRuleArgs $;

        public Builder() {
            $ = new ComputeClusterVmDependencyRuleArgs();
        }

        public Builder(ComputeClusterVmDependencyRuleArgs defaults) {
            $ = new ComputeClusterVmDependencyRuleArgs(Objects.requireNonNull(defaults));
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
         * @param dependencyVmGroupName The name of the VM group that this
         * rule depends on. The VMs defined in the group specified by
         * `vm_group_name` will not be started until the VMs in this
         * group are started.
         * 
         * @return builder
         * 
         */
        public Builder dependencyVmGroupName(Output<String> dependencyVmGroupName) {
            $.dependencyVmGroupName = dependencyVmGroupName;
            return this;
        }

        /**
         * @param dependencyVmGroupName The name of the VM group that this
         * rule depends on. The VMs defined in the group specified by
         * `vm_group_name` will not be started until the VMs in this
         * group are started.
         * 
         * @return builder
         * 
         */
        public Builder dependencyVmGroupName(String dependencyVmGroupName) {
            return dependencyVmGroupName(Output.of(dependencyVmGroupName));
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
         * @param vmGroupName The name of the VM group that is the subject of
         * this rule. The VMs defined in this group will not be started until the VMs in
         * the group specified by
         * `dependency_vm_group_name` are started.
         * 
         * @return builder
         * 
         */
        public Builder vmGroupName(Output<String> vmGroupName) {
            $.vmGroupName = vmGroupName;
            return this;
        }

        /**
         * @param vmGroupName The name of the VM group that is the subject of
         * this rule. The VMs defined in this group will not be started until the VMs in
         * the group specified by
         * `dependency_vm_group_name` are started.
         * 
         * @return builder
         * 
         */
        public Builder vmGroupName(String vmGroupName) {
            return vmGroupName(Output.of(vmGroupName));
        }

        public ComputeClusterVmDependencyRuleArgs build() {
            $.computeClusterId = Objects.requireNonNull($.computeClusterId, "expected parameter 'computeClusterId' to be non-null");
            $.dependencyVmGroupName = Objects.requireNonNull($.dependencyVmGroupName, "expected parameter 'dependencyVmGroupName' to be non-null");
            $.vmGroupName = Objects.requireNonNull($.vmGroupName, "expected parameter 'vmGroupName' to be non-null");
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ComputeClusterVmGroupState extends com.pulumi.resources.ResourceArgs {

    public static final ComputeClusterVmGroupState Empty = new ComputeClusterVmGroupState();

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
     * The name of the VM group. This must be unique in the
     * cluster. Forces a new resource if changed.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the VM group. This must be unique in the
     * cluster. Forces a new resource if changed.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The UUIDs of the virtual machines in this
     * group.
     * 
     * &gt; **NOTE:** The namespace for cluster names on this resource (defined by the
     * `name` argument) is shared with the
     * `vsphere.ComputeClusterHostGroup`
     * resource. Make sure your names are unique across both resources.
     * 
     * &gt; **NOTE:** To update a existing VM group, you must first import the group with `import` command in
     * import section. When importing a VM group, validate that all virtual machines that
     * need to be in the group are included in the `virtual_machine_ids`; otherwise, any virtual machines
     * that are not in `virtual_machine_ids` the included will be removed from the group.
     * 
     */
    @Import(name="virtualMachineIds")
    private @Nullable Output<List<String>> virtualMachineIds;

    /**
     * @return The UUIDs of the virtual machines in this
     * group.
     * 
     * &gt; **NOTE:** The namespace for cluster names on this resource (defined by the
     * `name` argument) is shared with the
     * `vsphere.ComputeClusterHostGroup`
     * resource. Make sure your names are unique across both resources.
     * 
     * &gt; **NOTE:** To update a existing VM group, you must first import the group with `import` command in
     * import section. When importing a VM group, validate that all virtual machines that
     * need to be in the group are included in the `virtual_machine_ids`; otherwise, any virtual machines
     * that are not in `virtual_machine_ids` the included will be removed from the group.
     * 
     */
    public Optional<Output<List<String>>> virtualMachineIds() {
        return Optional.ofNullable(this.virtualMachineIds);
    }

    private ComputeClusterVmGroupState() {}

    private ComputeClusterVmGroupState(ComputeClusterVmGroupState $) {
        this.computeClusterId = $.computeClusterId;
        this.name = $.name;
        this.virtualMachineIds = $.virtualMachineIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ComputeClusterVmGroupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ComputeClusterVmGroupState $;

        public Builder() {
            $ = new ComputeClusterVmGroupState();
        }

        public Builder(ComputeClusterVmGroupState defaults) {
            $ = new ComputeClusterVmGroupState(Objects.requireNonNull(defaults));
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
         * @param name The name of the VM group. This must be unique in the
         * cluster. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the VM group. This must be unique in the
         * cluster. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param virtualMachineIds The UUIDs of the virtual machines in this
         * group.
         * 
         * &gt; **NOTE:** The namespace for cluster names on this resource (defined by the
         * `name` argument) is shared with the
         * `vsphere.ComputeClusterHostGroup`
         * resource. Make sure your names are unique across both resources.
         * 
         * &gt; **NOTE:** To update a existing VM group, you must first import the group with `import` command in
         * import section. When importing a VM group, validate that all virtual machines that
         * need to be in the group are included in the `virtual_machine_ids`; otherwise, any virtual machines
         * that are not in `virtual_machine_ids` the included will be removed from the group.
         * 
         * @return builder
         * 
         */
        public Builder virtualMachineIds(@Nullable Output<List<String>> virtualMachineIds) {
            $.virtualMachineIds = virtualMachineIds;
            return this;
        }

        /**
         * @param virtualMachineIds The UUIDs of the virtual machines in this
         * group.
         * 
         * &gt; **NOTE:** The namespace for cluster names on this resource (defined by the
         * `name` argument) is shared with the
         * `vsphere.ComputeClusterHostGroup`
         * resource. Make sure your names are unique across both resources.
         * 
         * &gt; **NOTE:** To update a existing VM group, you must first import the group with `import` command in
         * import section. When importing a VM group, validate that all virtual machines that
         * need to be in the group are included in the `virtual_machine_ids`; otherwise, any virtual machines
         * that are not in `virtual_machine_ids` the included will be removed from the group.
         * 
         * @return builder
         * 
         */
        public Builder virtualMachineIds(List<String> virtualMachineIds) {
            return virtualMachineIds(Output.of(virtualMachineIds));
        }

        /**
         * @param virtualMachineIds The UUIDs of the virtual machines in this
         * group.
         * 
         * &gt; **NOTE:** The namespace for cluster names on this resource (defined by the
         * `name` argument) is shared with the
         * `vsphere.ComputeClusterHostGroup`
         * resource. Make sure your names are unique across both resources.
         * 
         * &gt; **NOTE:** To update a existing VM group, you must first import the group with `import` command in
         * import section. When importing a VM group, validate that all virtual machines that
         * need to be in the group are included in the `virtual_machine_ids`; otherwise, any virtual machines
         * that are not in `virtual_machine_ids` the included will be removed from the group.
         * 
         * @return builder
         * 
         */
        public Builder virtualMachineIds(String... virtualMachineIds) {
            return virtualMachineIds(List.of(virtualMachineIds));
        }

        public ComputeClusterVmGroupState build() {
            return $;
        }
    }

}

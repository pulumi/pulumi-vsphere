// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.vsphere.inputs.ComputeClusterHostImageComponentArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ComputeClusterHostImageArgs extends com.pulumi.resources.ResourceArgs {

    public static final ComputeClusterHostImageArgs Empty = new ComputeClusterHostImageArgs();

    /**
     * List of custom components.
     * 
     */
    @Import(name="components")
    private @Nullable Output<List<ComputeClusterHostImageComponentArgs>> components;

    /**
     * @return List of custom components.
     * 
     */
    public Optional<Output<List<ComputeClusterHostImageComponentArgs>>> components() {
        return Optional.ofNullable(this.components);
    }

    /**
     * The ESXi version which the image is based on.
     * 
     */
    @Import(name="esxVersion")
    private @Nullable Output<String> esxVersion;

    /**
     * @return The ESXi version which the image is based on.
     * 
     */
    public Optional<Output<String>> esxVersion() {
        return Optional.ofNullable(this.esxVersion);
    }

    private ComputeClusterHostImageArgs() {}

    private ComputeClusterHostImageArgs(ComputeClusterHostImageArgs $) {
        this.components = $.components;
        this.esxVersion = $.esxVersion;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ComputeClusterHostImageArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ComputeClusterHostImageArgs $;

        public Builder() {
            $ = new ComputeClusterHostImageArgs();
        }

        public Builder(ComputeClusterHostImageArgs defaults) {
            $ = new ComputeClusterHostImageArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param components List of custom components.
         * 
         * @return builder
         * 
         */
        public Builder components(@Nullable Output<List<ComputeClusterHostImageComponentArgs>> components) {
            $.components = components;
            return this;
        }

        /**
         * @param components List of custom components.
         * 
         * @return builder
         * 
         */
        public Builder components(List<ComputeClusterHostImageComponentArgs> components) {
            return components(Output.of(components));
        }

        /**
         * @param components List of custom components.
         * 
         * @return builder
         * 
         */
        public Builder components(ComputeClusterHostImageComponentArgs... components) {
            return components(List.of(components));
        }

        /**
         * @param esxVersion The ESXi version which the image is based on.
         * 
         * @return builder
         * 
         */
        public Builder esxVersion(@Nullable Output<String> esxVersion) {
            $.esxVersion = esxVersion;
            return this;
        }

        /**
         * @param esxVersion The ESXi version which the image is based on.
         * 
         * @return builder
         * 
         */
        public Builder esxVersion(String esxVersion) {
            return esxVersion(Output.of(esxVersion));
        }

        public ComputeClusterHostImageArgs build() {
            return $;
        }
    }

}

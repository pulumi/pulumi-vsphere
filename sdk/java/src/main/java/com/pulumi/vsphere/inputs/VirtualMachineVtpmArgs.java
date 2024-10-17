// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VirtualMachineVtpmArgs extends com.pulumi.resources.ResourceArgs {

    public static final VirtualMachineVtpmArgs Empty = new VirtualMachineVtpmArgs();

    /**
     * The version of the TPM device. Default is 2.0.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return The version of the TPM device. Default is 2.0.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private VirtualMachineVtpmArgs() {}

    private VirtualMachineVtpmArgs(VirtualMachineVtpmArgs $) {
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VirtualMachineVtpmArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VirtualMachineVtpmArgs $;

        public Builder() {
            $ = new VirtualMachineVtpmArgs();
        }

        public Builder(VirtualMachineVtpmArgs defaults) {
            $ = new VirtualMachineVtpmArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param version The version of the TPM device. Default is 2.0.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version The version of the TPM device. Default is 2.0.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public VirtualMachineVtpmArgs build() {
            return $;
        }
    }

}

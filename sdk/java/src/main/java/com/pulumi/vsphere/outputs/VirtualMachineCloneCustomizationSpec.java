// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VirtualMachineCloneCustomizationSpec {
    /**
     * @return The UUID of the virtual machine.
     * 
     */
    private String id;
    private @Nullable Integer timeout;

    private VirtualMachineCloneCustomizationSpec() {}
    /**
     * @return The UUID of the virtual machine.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<Integer> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VirtualMachineCloneCustomizationSpec defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private @Nullable Integer timeout;
        public Builder() {}
        public Builder(VirtualMachineCloneCustomizationSpec defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.timeout = defaults.timeout;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder timeout(@Nullable Integer timeout) {
            this.timeout = timeout;
            return this;
        }
        public VirtualMachineCloneCustomizationSpec build() {
            final var o = new VirtualMachineCloneCustomizationSpec();
            o.id = id;
            o.timeout = timeout;
            return o;
        }
    }
}
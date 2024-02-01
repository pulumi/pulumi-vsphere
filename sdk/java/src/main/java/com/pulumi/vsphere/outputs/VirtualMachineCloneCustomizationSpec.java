// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
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
    /**
     * @return The amount of time, in minutes, to wait for guest OS customization to complete before returning with an error. Setting this value to 0 or a negative value skips the waiter. Default: 10.
     * 
     */
    private @Nullable Integer timeout;

    private VirtualMachineCloneCustomizationSpec() {}
    /**
     * @return The UUID of the virtual machine.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The amount of time, in minutes, to wait for guest OS customization to complete before returning with an error. Setting this value to 0 or a negative value skips the waiter. Default: 10.
     * 
     */
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
            if (id == null) {
              throw new MissingRequiredPropertyException("VirtualMachineCloneCustomizationSpec", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder timeout(@Nullable Integer timeout) {

            this.timeout = timeout;
            return this;
        }
        public VirtualMachineCloneCustomizationSpec build() {
            final var _resultValue = new VirtualMachineCloneCustomizationSpec();
            _resultValue.id = id;
            _resultValue.timeout = timeout;
            return _resultValue;
        }
    }
}

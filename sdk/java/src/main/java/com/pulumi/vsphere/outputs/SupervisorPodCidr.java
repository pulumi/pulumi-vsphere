// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class SupervisorPodCidr {
    /**
     * @return Network address.
     * 
     */
    private String address;
    /**
     * @return Subnet prefix.
     * 
     */
    private Integer prefix;

    private SupervisorPodCidr() {}
    /**
     * @return Network address.
     * 
     */
    public String address() {
        return this.address;
    }
    /**
     * @return Subnet prefix.
     * 
     */
    public Integer prefix() {
        return this.prefix;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SupervisorPodCidr defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String address;
        private Integer prefix;
        public Builder() {}
        public Builder(SupervisorPodCidr defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.address = defaults.address;
    	      this.prefix = defaults.prefix;
        }

        @CustomType.Setter
        public Builder address(String address) {
            if (address == null) {
              throw new MissingRequiredPropertyException("SupervisorPodCidr", "address");
            }
            this.address = address;
            return this;
        }
        @CustomType.Setter
        public Builder prefix(Integer prefix) {
            if (prefix == null) {
              throw new MissingRequiredPropertyException("SupervisorPodCidr", "prefix");
            }
            this.prefix = prefix;
            return this;
        }
        public SupervisorPodCidr build() {
            final var _resultValue = new SupervisorPodCidr();
            _resultValue.address = address;
            _resultValue.prefix = prefix;
            return _resultValue;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class DistributedPortGroupVlanRange {
    private Integer maxVlan;
    private Integer minVlan;

    private DistributedPortGroupVlanRange() {}
    public Integer maxVlan() {
        return this.maxVlan;
    }
    public Integer minVlan() {
        return this.minVlan;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DistributedPortGroupVlanRange defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer maxVlan;
        private Integer minVlan;
        public Builder() {}
        public Builder(DistributedPortGroupVlanRange defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.maxVlan = defaults.maxVlan;
    	      this.minVlan = defaults.minVlan;
        }

        @CustomType.Setter
        public Builder maxVlan(Integer maxVlan) {
            this.maxVlan = Objects.requireNonNull(maxVlan);
            return this;
        }
        @CustomType.Setter
        public Builder minVlan(Integer minVlan) {
            this.minVlan = Objects.requireNonNull(minVlan);
            return this;
        }
        public DistributedPortGroupVlanRange build() {
            final var _resultValue = new DistributedPortGroupVlanRange();
            _resultValue.maxVlan = maxVlan;
            _resultValue.minVlan = minVlan;
            return _resultValue;
        }
    }
}

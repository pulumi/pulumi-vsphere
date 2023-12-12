// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetDistributedVirtualSwitchResult {
    private @Nullable String datacenterId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String name;
    /**
     * @return The list of the uplinks on this vSphere distributed switch, as per the
     * `uplinks` argument to the
     * `vsphere.DistributedVirtualSwitch`
     * resource.
     * 
     */
    private List<String> uplinks;

    private GetDistributedVirtualSwitchResult() {}
    public Optional<String> datacenterId() {
        return Optional.ofNullable(this.datacenterId);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }
    /**
     * @return The list of the uplinks on this vSphere distributed switch, as per the
     * `uplinks` argument to the
     * `vsphere.DistributedVirtualSwitch`
     * resource.
     * 
     */
    public List<String> uplinks() {
        return this.uplinks;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDistributedVirtualSwitchResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String datacenterId;
        private String id;
        private String name;
        private List<String> uplinks;
        public Builder() {}
        public Builder(GetDistributedVirtualSwitchResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.datacenterId = defaults.datacenterId;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.uplinks = defaults.uplinks;
        }

        @CustomType.Setter
        public Builder datacenterId(@Nullable String datacenterId) {
            this.datacenterId = datacenterId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder uplinks(List<String> uplinks) {
            this.uplinks = Objects.requireNonNull(uplinks);
            return this;
        }
        public Builder uplinks(String... uplinks) {
            return uplinks(List.of(uplinks));
        }
        public GetDistributedVirtualSwitchResult build() {
            final var _resultValue = new GetDistributedVirtualSwitchResult();
            _resultValue.datacenterId = datacenterId;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.uplinks = uplinks;
            return _resultValue;
        }
    }
}

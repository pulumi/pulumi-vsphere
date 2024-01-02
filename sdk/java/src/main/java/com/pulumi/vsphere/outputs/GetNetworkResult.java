// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetNetworkResult {
    private @Nullable String datacenterId;
    private @Nullable String distributedVirtualSwitchUuid;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String name;
    /**
     * @return The managed object type for the discovered network. This will be one
     * of `DistributedVirtualPortgroup` for distributed port groups, `Network` for
     * standard (host-based) port groups, or `OpaqueNetwork` for networks managed
     * externally, such as those managed by NSX.
     * 
     */
    private String type;

    private GetNetworkResult() {}
    public Optional<String> datacenterId() {
        return Optional.ofNullable(this.datacenterId);
    }
    public Optional<String> distributedVirtualSwitchUuid() {
        return Optional.ofNullable(this.distributedVirtualSwitchUuid);
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
     * @return The managed object type for the discovered network. This will be one
     * of `DistributedVirtualPortgroup` for distributed port groups, `Network` for
     * standard (host-based) port groups, or `OpaqueNetwork` for networks managed
     * externally, such as those managed by NSX.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetNetworkResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String datacenterId;
        private @Nullable String distributedVirtualSwitchUuid;
        private String id;
        private String name;
        private String type;
        public Builder() {}
        public Builder(GetNetworkResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.datacenterId = defaults.datacenterId;
    	      this.distributedVirtualSwitchUuid = defaults.distributedVirtualSwitchUuid;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder datacenterId(@Nullable String datacenterId) {

            this.datacenterId = datacenterId;
            return this;
        }
        @CustomType.Setter
        public Builder distributedVirtualSwitchUuid(@Nullable String distributedVirtualSwitchUuid) {

            this.distributedVirtualSwitchUuid = distributedVirtualSwitchUuid;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetNetworkResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetNetworkResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetNetworkResult", "type");
            }
            this.type = type;
            return this;
        }
        public GetNetworkResult build() {
            final var _resultValue = new GetNetworkResult();
            _resultValue.datacenterId = datacenterId;
            _resultValue.distributedVirtualSwitchUuid = distributedVirtualSwitchUuid;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}

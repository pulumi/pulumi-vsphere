// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class DistributedVirtualSwitchHost {
    /**
     * @return The list of NIC devices to map to uplinks on the VDS,
     * added in order they are specified.
     * 
     */
    private @Nullable List<String> devices;
    /**
     * @return The host system ID of the host to add to the
     * VDS.
     * 
     */
    private String hostSystemId;

    private DistributedVirtualSwitchHost() {}
    /**
     * @return The list of NIC devices to map to uplinks on the VDS,
     * added in order they are specified.
     * 
     */
    public List<String> devices() {
        return this.devices == null ? List.of() : this.devices;
    }
    /**
     * @return The host system ID of the host to add to the
     * VDS.
     * 
     */
    public String hostSystemId() {
        return this.hostSystemId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DistributedVirtualSwitchHost defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> devices;
        private String hostSystemId;
        public Builder() {}
        public Builder(DistributedVirtualSwitchHost defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.devices = defaults.devices;
    	      this.hostSystemId = defaults.hostSystemId;
        }

        @CustomType.Setter
        public Builder devices(@Nullable List<String> devices) {
            this.devices = devices;
            return this;
        }
        public Builder devices(String... devices) {
            return devices(List.of(devices));
        }
        @CustomType.Setter
        public Builder hostSystemId(String hostSystemId) {
            this.hostSystemId = Objects.requireNonNull(hostSystemId);
            return this;
        }
        public DistributedVirtualSwitchHost build() {
            final var _resultValue = new DistributedVirtualSwitchHost();
            _resultValue.devices = devices;
            _resultValue.hostSystemId = hostSystemId;
            return _resultValue;
        }
    }
}

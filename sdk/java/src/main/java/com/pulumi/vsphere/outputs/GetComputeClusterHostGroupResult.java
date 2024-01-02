// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetComputeClusterHostGroupResult {
    private String computeClusterId;
    /**
     * @return The [managed object reference ID][docs-about-morefs] of
     * the ESXi hosts in the host group.
     * 
     */
    private List<String> hostSystemIds;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String name;

    private GetComputeClusterHostGroupResult() {}
    public String computeClusterId() {
        return this.computeClusterId;
    }
    /**
     * @return The [managed object reference ID][docs-about-morefs] of
     * the ESXi hosts in the host group.
     * 
     */
    public List<String> hostSystemIds() {
        return this.hostSystemIds;
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

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetComputeClusterHostGroupResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String computeClusterId;
        private List<String> hostSystemIds;
        private String id;
        private String name;
        public Builder() {}
        public Builder(GetComputeClusterHostGroupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.computeClusterId = defaults.computeClusterId;
    	      this.hostSystemIds = defaults.hostSystemIds;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder computeClusterId(String computeClusterId) {
            if (computeClusterId == null) {
              throw new MissingRequiredPropertyException("GetComputeClusterHostGroupResult", "computeClusterId");
            }
            this.computeClusterId = computeClusterId;
            return this;
        }
        @CustomType.Setter
        public Builder hostSystemIds(List<String> hostSystemIds) {
            if (hostSystemIds == null) {
              throw new MissingRequiredPropertyException("GetComputeClusterHostGroupResult", "hostSystemIds");
            }
            this.hostSystemIds = hostSystemIds;
            return this;
        }
        public Builder hostSystemIds(String... hostSystemIds) {
            return hostSystemIds(List.of(hostSystemIds));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetComputeClusterHostGroupResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetComputeClusterHostGroupResult", "name");
            }
            this.name = name;
            return this;
        }
        public GetComputeClusterHostGroupResult build() {
            final var _resultValue = new GetComputeClusterHostGroupResult();
            _resultValue.computeClusterId = computeClusterId;
            _resultValue.hostSystemIds = hostSystemIds;
            _resultValue.id = id;
            _resultValue.name = name;
            return _resultValue;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetDatastoreResult {
    private @Nullable String datacenterId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String name;

    private GetDatastoreResult() {}
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

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDatastoreResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String datacenterId;
        private String id;
        private String name;
        public Builder() {}
        public Builder(GetDatastoreResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.datacenterId = defaults.datacenterId;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
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
        public GetDatastoreResult build() {
            final var o = new GetDatastoreResult();
            o.datacenterId = datacenterId;
            o.id = id;
            o.name = name;
            return o;
        }
    }
}

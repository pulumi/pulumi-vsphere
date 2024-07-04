// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
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
    /**
     * @return The disk space usage statistics for the specific datastore. The
     * total datastore capacity is represented as `capacity` and the free remaining
     * disk is represented as `free`.
     * 
     */
    private @Nullable Map<String,Object> stats;

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
    /**
     * @return The disk space usage statistics for the specific datastore. The
     * total datastore capacity is represented as `capacity` and the free remaining
     * disk is represented as `free`.
     * 
     */
    public Map<String,Object> stats() {
        return this.stats == null ? Map.of() : this.stats;
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
        private @Nullable Map<String,Object> stats;
        public Builder() {}
        public Builder(GetDatastoreResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.datacenterId = defaults.datacenterId;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.stats = defaults.stats;
        }

        @CustomType.Setter
        public Builder datacenterId(@Nullable String datacenterId) {

            this.datacenterId = datacenterId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetDatastoreResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetDatastoreResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder stats(@Nullable Map<String,Object> stats) {

            this.stats = stats;
            return this;
        }
        public GetDatastoreResult build() {
            final var _resultValue = new GetDatastoreResult();
            _resultValue.datacenterId = datacenterId;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.stats = stats;
            return _resultValue;
        }
    }
}

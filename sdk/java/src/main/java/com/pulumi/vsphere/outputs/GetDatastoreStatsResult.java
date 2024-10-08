// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetDatastoreStatsResult {
    /**
     * @return A mapping of the capacity for all datastore in the datacenter,
     * where the name of the datastore is used as key and the capacity as value.
     * 
     */
    private @Nullable Map<String,String> capacity;
    /**
     * @return The [managed object reference ID][docs-about-morefs] of the
     * datacenter the datastores are located in.
     * 
     */
    private String datacenterId;
    /**
     * @return A mapping of the free space for each datastore in the
     * datacenter, where the name of the datastore is used as key and the free space
     * as value.
     * 
     */
    private @Nullable Map<String,String> freeSpace;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;

    private GetDatastoreStatsResult() {}
    /**
     * @return A mapping of the capacity for all datastore in the datacenter,
     * where the name of the datastore is used as key and the capacity as value.
     * 
     */
    public Map<String,String> capacity() {
        return this.capacity == null ? Map.of() : this.capacity;
    }
    /**
     * @return The [managed object reference ID][docs-about-morefs] of the
     * datacenter the datastores are located in.
     * 
     */
    public String datacenterId() {
        return this.datacenterId;
    }
    /**
     * @return A mapping of the free space for each datastore in the
     * datacenter, where the name of the datastore is used as key and the free space
     * as value.
     * 
     */
    public Map<String,String> freeSpace() {
        return this.freeSpace == null ? Map.of() : this.freeSpace;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDatastoreStatsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,String> capacity;
        private String datacenterId;
        private @Nullable Map<String,String> freeSpace;
        private String id;
        public Builder() {}
        public Builder(GetDatastoreStatsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.capacity = defaults.capacity;
    	      this.datacenterId = defaults.datacenterId;
    	      this.freeSpace = defaults.freeSpace;
    	      this.id = defaults.id;
        }

        @CustomType.Setter
        public Builder capacity(@Nullable Map<String,String> capacity) {

            this.capacity = capacity;
            return this;
        }
        @CustomType.Setter
        public Builder datacenterId(String datacenterId) {
            if (datacenterId == null) {
              throw new MissingRequiredPropertyException("GetDatastoreStatsResult", "datacenterId");
            }
            this.datacenterId = datacenterId;
            return this;
        }
        @CustomType.Setter
        public Builder freeSpace(@Nullable Map<String,String> freeSpace) {

            this.freeSpace = freeSpace;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetDatastoreStatsResult", "id");
            }
            this.id = id;
            return this;
        }
        public GetDatastoreStatsResult build() {
            final var _resultValue = new GetDatastoreStatsResult();
            _resultValue.capacity = capacity;
            _resultValue.datacenterId = datacenterId;
            _resultValue.freeSpace = freeSpace;
            _resultValue.id = id;
            return _resultValue;
        }
    }
}

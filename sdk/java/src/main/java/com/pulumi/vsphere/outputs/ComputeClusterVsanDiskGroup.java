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
public final class ComputeClusterVsanDiskGroup {
    /**
     * @return The canonical name of the disk to use for vSAN cache.
     * 
     */
    private @Nullable String cache;
    /**
     * @return An array of disk canonical names for vSAN storage.
     * 
     */
    private @Nullable List<String> storages;

    private ComputeClusterVsanDiskGroup() {}
    /**
     * @return The canonical name of the disk to use for vSAN cache.
     * 
     */
    public Optional<String> cache() {
        return Optional.ofNullable(this.cache);
    }
    /**
     * @return An array of disk canonical names for vSAN storage.
     * 
     */
    public List<String> storages() {
        return this.storages == null ? List.of() : this.storages;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ComputeClusterVsanDiskGroup defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String cache;
        private @Nullable List<String> storages;
        public Builder() {}
        public Builder(ComputeClusterVsanDiskGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cache = defaults.cache;
    	      this.storages = defaults.storages;
        }

        @CustomType.Setter
        public Builder cache(@Nullable String cache) {
            this.cache = cache;
            return this;
        }
        @CustomType.Setter
        public Builder storages(@Nullable List<String> storages) {
            this.storages = storages;
            return this;
        }
        public Builder storages(String... storages) {
            return storages(List.of(storages));
        }
        public ComputeClusterVsanDiskGroup build() {
            final var _resultValue = new ComputeClusterVsanDiskGroup();
            _resultValue.cache = cache;
            _resultValue.storages = storages;
            return _resultValue;
        }
    }
}

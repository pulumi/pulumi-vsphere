// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetVmfsDisksResult {
    /**
     * @return A lexicographically sorted list of devices discovered by the
     * operation, matching the supplied `filter`, if provided.
     * 
     */
    private List<String> disks;
    private @Nullable String filter;
    private String hostSystemId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable Boolean rescan;

    private GetVmfsDisksResult() {}
    /**
     * @return A lexicographically sorted list of devices discovered by the
     * operation, matching the supplied `filter`, if provided.
     * 
     */
    public List<String> disks() {
        return this.disks;
    }
    public Optional<String> filter() {
        return Optional.ofNullable(this.filter);
    }
    public String hostSystemId() {
        return this.hostSystemId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<Boolean> rescan() {
        return Optional.ofNullable(this.rescan);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVmfsDisksResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> disks;
        private @Nullable String filter;
        private String hostSystemId;
        private String id;
        private @Nullable Boolean rescan;
        public Builder() {}
        public Builder(GetVmfsDisksResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.disks = defaults.disks;
    	      this.filter = defaults.filter;
    	      this.hostSystemId = defaults.hostSystemId;
    	      this.id = defaults.id;
    	      this.rescan = defaults.rescan;
        }

        @CustomType.Setter
        public Builder disks(List<String> disks) {
            if (disks == null) {
              throw new MissingRequiredPropertyException("GetVmfsDisksResult", "disks");
            }
            this.disks = disks;
            return this;
        }
        public Builder disks(String... disks) {
            return disks(List.of(disks));
        }
        @CustomType.Setter
        public Builder filter(@Nullable String filter) {

            this.filter = filter;
            return this;
        }
        @CustomType.Setter
        public Builder hostSystemId(String hostSystemId) {
            if (hostSystemId == null) {
              throw new MissingRequiredPropertyException("GetVmfsDisksResult", "hostSystemId");
            }
            this.hostSystemId = hostSystemId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetVmfsDisksResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder rescan(@Nullable Boolean rescan) {

            this.rescan = rescan;
            return this;
        }
        public GetVmfsDisksResult build() {
            final var _resultValue = new GetVmfsDisksResult();
            _resultValue.disks = disks;
            _resultValue.filter = filter;
            _resultValue.hostSystemId = hostSystemId;
            _resultValue.id = id;
            _resultValue.rescan = rescan;
            return _resultValue;
        }
    }
}

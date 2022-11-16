// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.vsphere.outputs.VirtualMachineCloneCustomize;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VirtualMachineClone {
    private @Nullable VirtualMachineCloneCustomize customize;
    private @Nullable Boolean linkedClone;
    private @Nullable Map<String,String> ovfNetworkMap;
    private @Nullable Map<String,String> ovfStorageMap;
    private String templateUuid;
    private @Nullable Integer timeout;

    private VirtualMachineClone() {}
    public Optional<VirtualMachineCloneCustomize> customize() {
        return Optional.ofNullable(this.customize);
    }
    public Optional<Boolean> linkedClone() {
        return Optional.ofNullable(this.linkedClone);
    }
    public Map<String,String> ovfNetworkMap() {
        return this.ovfNetworkMap == null ? Map.of() : this.ovfNetworkMap;
    }
    public Map<String,String> ovfStorageMap() {
        return this.ovfStorageMap == null ? Map.of() : this.ovfStorageMap;
    }
    public String templateUuid() {
        return this.templateUuid;
    }
    public Optional<Integer> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VirtualMachineClone defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable VirtualMachineCloneCustomize customize;
        private @Nullable Boolean linkedClone;
        private @Nullable Map<String,String> ovfNetworkMap;
        private @Nullable Map<String,String> ovfStorageMap;
        private String templateUuid;
        private @Nullable Integer timeout;
        public Builder() {}
        public Builder(VirtualMachineClone defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.customize = defaults.customize;
    	      this.linkedClone = defaults.linkedClone;
    	      this.ovfNetworkMap = defaults.ovfNetworkMap;
    	      this.ovfStorageMap = defaults.ovfStorageMap;
    	      this.templateUuid = defaults.templateUuid;
    	      this.timeout = defaults.timeout;
        }

        @CustomType.Setter
        public Builder customize(@Nullable VirtualMachineCloneCustomize customize) {
            this.customize = customize;
            return this;
        }
        @CustomType.Setter
        public Builder linkedClone(@Nullable Boolean linkedClone) {
            this.linkedClone = linkedClone;
            return this;
        }
        @CustomType.Setter
        public Builder ovfNetworkMap(@Nullable Map<String,String> ovfNetworkMap) {
            this.ovfNetworkMap = ovfNetworkMap;
            return this;
        }
        @CustomType.Setter
        public Builder ovfStorageMap(@Nullable Map<String,String> ovfStorageMap) {
            this.ovfStorageMap = ovfStorageMap;
            return this;
        }
        @CustomType.Setter
        public Builder templateUuid(String templateUuid) {
            this.templateUuid = Objects.requireNonNull(templateUuid);
            return this;
        }
        @CustomType.Setter
        public Builder timeout(@Nullable Integer timeout) {
            this.timeout = timeout;
            return this;
        }
        public VirtualMachineClone build() {
            final var o = new VirtualMachineClone();
            o.customize = customize;
            o.linkedClone = linkedClone;
            o.ovfNetworkMap = ovfNetworkMap;
            o.ovfStorageMap = ovfStorageMap;
            o.templateUuid = templateUuid;
            o.timeout = timeout;
            return o;
        }
    }
}
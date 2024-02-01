// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.vsphere.outputs.VirtualMachineCloneCustomizationSpec;
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
    /**
     * @return The customization specification for the virtual machine post-clone.
     * 
     */
    private @Nullable VirtualMachineCloneCustomizationSpec customizationSpec;
    /**
     * @return The customization specification for the virtual machine post-clone.
     * 
     */
    private @Nullable VirtualMachineCloneCustomize customize;
    /**
     * @return Whether or not to create a linked clone when cloning. When this option is used, the source VM must have a single snapshot associated with it.
     * 
     */
    private @Nullable Boolean linkedClone;
    /**
     * @return Mapping of ovf networks to the networks to use in vSphere.
     * 
     */
    private @Nullable Map<String,String> ovfNetworkMap;
    /**
     * @return Mapping of ovf storage to the datastores to use in vSphere.
     * 
     */
    private @Nullable Map<String,String> ovfStorageMap;
    /**
     * @return The UUID of the source virtual machine or template.
     * 
     */
    private String templateUuid;
    /**
     * @return The timeout, in minutes, to wait for the virtual machine clone to complete.
     * 
     */
    private @Nullable Integer timeout;

    private VirtualMachineClone() {}
    /**
     * @return The customization specification for the virtual machine post-clone.
     * 
     */
    public Optional<VirtualMachineCloneCustomizationSpec> customizationSpec() {
        return Optional.ofNullable(this.customizationSpec);
    }
    /**
     * @return The customization specification for the virtual machine post-clone.
     * 
     */
    public Optional<VirtualMachineCloneCustomize> customize() {
        return Optional.ofNullable(this.customize);
    }
    /**
     * @return Whether or not to create a linked clone when cloning. When this option is used, the source VM must have a single snapshot associated with it.
     * 
     */
    public Optional<Boolean> linkedClone() {
        return Optional.ofNullable(this.linkedClone);
    }
    /**
     * @return Mapping of ovf networks to the networks to use in vSphere.
     * 
     */
    public Map<String,String> ovfNetworkMap() {
        return this.ovfNetworkMap == null ? Map.of() : this.ovfNetworkMap;
    }
    /**
     * @return Mapping of ovf storage to the datastores to use in vSphere.
     * 
     */
    public Map<String,String> ovfStorageMap() {
        return this.ovfStorageMap == null ? Map.of() : this.ovfStorageMap;
    }
    /**
     * @return The UUID of the source virtual machine or template.
     * 
     */
    public String templateUuid() {
        return this.templateUuid;
    }
    /**
     * @return The timeout, in minutes, to wait for the virtual machine clone to complete.
     * 
     */
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
        private @Nullable VirtualMachineCloneCustomizationSpec customizationSpec;
        private @Nullable VirtualMachineCloneCustomize customize;
        private @Nullable Boolean linkedClone;
        private @Nullable Map<String,String> ovfNetworkMap;
        private @Nullable Map<String,String> ovfStorageMap;
        private String templateUuid;
        private @Nullable Integer timeout;
        public Builder() {}
        public Builder(VirtualMachineClone defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.customizationSpec = defaults.customizationSpec;
    	      this.customize = defaults.customize;
    	      this.linkedClone = defaults.linkedClone;
    	      this.ovfNetworkMap = defaults.ovfNetworkMap;
    	      this.ovfStorageMap = defaults.ovfStorageMap;
    	      this.templateUuid = defaults.templateUuid;
    	      this.timeout = defaults.timeout;
        }

        @CustomType.Setter
        public Builder customizationSpec(@Nullable VirtualMachineCloneCustomizationSpec customizationSpec) {

            this.customizationSpec = customizationSpec;
            return this;
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
            if (templateUuid == null) {
              throw new MissingRequiredPropertyException("VirtualMachineClone", "templateUuid");
            }
            this.templateUuid = templateUuid;
            return this;
        }
        @CustomType.Setter
        public Builder timeout(@Nullable Integer timeout) {

            this.timeout = timeout;
            return this;
        }
        public VirtualMachineClone build() {
            final var _resultValue = new VirtualMachineClone();
            _resultValue.customizationSpec = customizationSpec;
            _resultValue.customize = customize;
            _resultValue.linkedClone = linkedClone;
            _resultValue.ovfNetworkMap = ovfNetworkMap;
            _resultValue.ovfStorageMap = ovfStorageMap;
            _resultValue.templateUuid = templateUuid;
            _resultValue.timeout = timeout;
            return _resultValue;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.vsphere.inputs.VirtualMachineCloneCustomizationSpecArgs;
import com.pulumi.vsphere.inputs.VirtualMachineCloneCustomizeArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VirtualMachineCloneArgs extends com.pulumi.resources.ResourceArgs {

    public static final VirtualMachineCloneArgs Empty = new VirtualMachineCloneArgs();

    @Import(name="customizationSpec")
    private @Nullable Output<VirtualMachineCloneCustomizationSpecArgs> customizationSpec;

    public Optional<Output<VirtualMachineCloneCustomizationSpecArgs>> customizationSpec() {
        return Optional.ofNullable(this.customizationSpec);
    }

    @Import(name="customize")
    private @Nullable Output<VirtualMachineCloneCustomizeArgs> customize;

    public Optional<Output<VirtualMachineCloneCustomizeArgs>> customize() {
        return Optional.ofNullable(this.customize);
    }

    @Import(name="linkedClone")
    private @Nullable Output<Boolean> linkedClone;

    public Optional<Output<Boolean>> linkedClone() {
        return Optional.ofNullable(this.linkedClone);
    }

    @Import(name="ovfNetworkMap")
    private @Nullable Output<Map<String,String>> ovfNetworkMap;

    public Optional<Output<Map<String,String>>> ovfNetworkMap() {
        return Optional.ofNullable(this.ovfNetworkMap);
    }

    @Import(name="ovfStorageMap")
    private @Nullable Output<Map<String,String>> ovfStorageMap;

    public Optional<Output<Map<String,String>>> ovfStorageMap() {
        return Optional.ofNullable(this.ovfStorageMap);
    }

    @Import(name="templateUuid", required=true)
    private Output<String> templateUuid;

    public Output<String> templateUuid() {
        return this.templateUuid;
    }

    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private VirtualMachineCloneArgs() {}

    private VirtualMachineCloneArgs(VirtualMachineCloneArgs $) {
        this.customizationSpec = $.customizationSpec;
        this.customize = $.customize;
        this.linkedClone = $.linkedClone;
        this.ovfNetworkMap = $.ovfNetworkMap;
        this.ovfStorageMap = $.ovfStorageMap;
        this.templateUuid = $.templateUuid;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VirtualMachineCloneArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VirtualMachineCloneArgs $;

        public Builder() {
            $ = new VirtualMachineCloneArgs();
        }

        public Builder(VirtualMachineCloneArgs defaults) {
            $ = new VirtualMachineCloneArgs(Objects.requireNonNull(defaults));
        }

        public Builder customizationSpec(@Nullable Output<VirtualMachineCloneCustomizationSpecArgs> customizationSpec) {
            $.customizationSpec = customizationSpec;
            return this;
        }

        public Builder customizationSpec(VirtualMachineCloneCustomizationSpecArgs customizationSpec) {
            return customizationSpec(Output.of(customizationSpec));
        }

        public Builder customize(@Nullable Output<VirtualMachineCloneCustomizeArgs> customize) {
            $.customize = customize;
            return this;
        }

        public Builder customize(VirtualMachineCloneCustomizeArgs customize) {
            return customize(Output.of(customize));
        }

        public Builder linkedClone(@Nullable Output<Boolean> linkedClone) {
            $.linkedClone = linkedClone;
            return this;
        }

        public Builder linkedClone(Boolean linkedClone) {
            return linkedClone(Output.of(linkedClone));
        }

        public Builder ovfNetworkMap(@Nullable Output<Map<String,String>> ovfNetworkMap) {
            $.ovfNetworkMap = ovfNetworkMap;
            return this;
        }

        public Builder ovfNetworkMap(Map<String,String> ovfNetworkMap) {
            return ovfNetworkMap(Output.of(ovfNetworkMap));
        }

        public Builder ovfStorageMap(@Nullable Output<Map<String,String>> ovfStorageMap) {
            $.ovfStorageMap = ovfStorageMap;
            return this;
        }

        public Builder ovfStorageMap(Map<String,String> ovfStorageMap) {
            return ovfStorageMap(Output.of(ovfStorageMap));
        }

        public Builder templateUuid(Output<String> templateUuid) {
            $.templateUuid = templateUuid;
            return this;
        }

        public Builder templateUuid(String templateUuid) {
            return templateUuid(Output.of(templateUuid));
        }

        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        public VirtualMachineCloneArgs build() {
            if ($.templateUuid == null) {
                throw new MissingRequiredPropertyException("VirtualMachineCloneArgs", "templateUuid");
            }
            return $;
        }
    }

}

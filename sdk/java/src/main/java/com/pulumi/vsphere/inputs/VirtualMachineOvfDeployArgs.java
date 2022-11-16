// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VirtualMachineOvfDeployArgs extends com.pulumi.resources.ResourceArgs {

    public static final VirtualMachineOvfDeployArgs Empty = new VirtualMachineOvfDeployArgs();

    @Import(name="allowUnverifiedSslCert")
    private @Nullable Output<Boolean> allowUnverifiedSslCert;

    public Optional<Output<Boolean>> allowUnverifiedSslCert() {
        return Optional.ofNullable(this.allowUnverifiedSslCert);
    }

    @Import(name="deploymentOption")
    private @Nullable Output<String> deploymentOption;

    public Optional<Output<String>> deploymentOption() {
        return Optional.ofNullable(this.deploymentOption);
    }

    @Import(name="diskProvisioning")
    private @Nullable Output<String> diskProvisioning;

    public Optional<Output<String>> diskProvisioning() {
        return Optional.ofNullable(this.diskProvisioning);
    }

    @Import(name="enableHiddenProperties")
    private @Nullable Output<Boolean> enableHiddenProperties;

    public Optional<Output<Boolean>> enableHiddenProperties() {
        return Optional.ofNullable(this.enableHiddenProperties);
    }

    @Import(name="ipAllocationPolicy")
    private @Nullable Output<String> ipAllocationPolicy;

    public Optional<Output<String>> ipAllocationPolicy() {
        return Optional.ofNullable(this.ipAllocationPolicy);
    }

    @Import(name="ipProtocol")
    private @Nullable Output<String> ipProtocol;

    public Optional<Output<String>> ipProtocol() {
        return Optional.ofNullable(this.ipProtocol);
    }

    @Import(name="localOvfPath")
    private @Nullable Output<String> localOvfPath;

    public Optional<Output<String>> localOvfPath() {
        return Optional.ofNullable(this.localOvfPath);
    }

    @Import(name="ovfNetworkMap")
    private @Nullable Output<Map<String,String>> ovfNetworkMap;

    public Optional<Output<Map<String,String>>> ovfNetworkMap() {
        return Optional.ofNullable(this.ovfNetworkMap);
    }

    @Import(name="remoteOvfUrl")
    private @Nullable Output<String> remoteOvfUrl;

    public Optional<Output<String>> remoteOvfUrl() {
        return Optional.ofNullable(this.remoteOvfUrl);
    }

    private VirtualMachineOvfDeployArgs() {}

    private VirtualMachineOvfDeployArgs(VirtualMachineOvfDeployArgs $) {
        this.allowUnverifiedSslCert = $.allowUnverifiedSslCert;
        this.deploymentOption = $.deploymentOption;
        this.diskProvisioning = $.diskProvisioning;
        this.enableHiddenProperties = $.enableHiddenProperties;
        this.ipAllocationPolicy = $.ipAllocationPolicy;
        this.ipProtocol = $.ipProtocol;
        this.localOvfPath = $.localOvfPath;
        this.ovfNetworkMap = $.ovfNetworkMap;
        this.remoteOvfUrl = $.remoteOvfUrl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VirtualMachineOvfDeployArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VirtualMachineOvfDeployArgs $;

        public Builder() {
            $ = new VirtualMachineOvfDeployArgs();
        }

        public Builder(VirtualMachineOvfDeployArgs defaults) {
            $ = new VirtualMachineOvfDeployArgs(Objects.requireNonNull(defaults));
        }

        public Builder allowUnverifiedSslCert(@Nullable Output<Boolean> allowUnverifiedSslCert) {
            $.allowUnverifiedSslCert = allowUnverifiedSslCert;
            return this;
        }

        public Builder allowUnverifiedSslCert(Boolean allowUnverifiedSslCert) {
            return allowUnverifiedSslCert(Output.of(allowUnverifiedSslCert));
        }

        public Builder deploymentOption(@Nullable Output<String> deploymentOption) {
            $.deploymentOption = deploymentOption;
            return this;
        }

        public Builder deploymentOption(String deploymentOption) {
            return deploymentOption(Output.of(deploymentOption));
        }

        public Builder diskProvisioning(@Nullable Output<String> diskProvisioning) {
            $.diskProvisioning = diskProvisioning;
            return this;
        }

        public Builder diskProvisioning(String diskProvisioning) {
            return diskProvisioning(Output.of(diskProvisioning));
        }

        public Builder enableHiddenProperties(@Nullable Output<Boolean> enableHiddenProperties) {
            $.enableHiddenProperties = enableHiddenProperties;
            return this;
        }

        public Builder enableHiddenProperties(Boolean enableHiddenProperties) {
            return enableHiddenProperties(Output.of(enableHiddenProperties));
        }

        public Builder ipAllocationPolicy(@Nullable Output<String> ipAllocationPolicy) {
            $.ipAllocationPolicy = ipAllocationPolicy;
            return this;
        }

        public Builder ipAllocationPolicy(String ipAllocationPolicy) {
            return ipAllocationPolicy(Output.of(ipAllocationPolicy));
        }

        public Builder ipProtocol(@Nullable Output<String> ipProtocol) {
            $.ipProtocol = ipProtocol;
            return this;
        }

        public Builder ipProtocol(String ipProtocol) {
            return ipProtocol(Output.of(ipProtocol));
        }

        public Builder localOvfPath(@Nullable Output<String> localOvfPath) {
            $.localOvfPath = localOvfPath;
            return this;
        }

        public Builder localOvfPath(String localOvfPath) {
            return localOvfPath(Output.of(localOvfPath));
        }

        public Builder ovfNetworkMap(@Nullable Output<Map<String,String>> ovfNetworkMap) {
            $.ovfNetworkMap = ovfNetworkMap;
            return this;
        }

        public Builder ovfNetworkMap(Map<String,String> ovfNetworkMap) {
            return ovfNetworkMap(Output.of(ovfNetworkMap));
        }

        public Builder remoteOvfUrl(@Nullable Output<String> remoteOvfUrl) {
            $.remoteOvfUrl = remoteOvfUrl;
            return this;
        }

        public Builder remoteOvfUrl(String remoteOvfUrl) {
            return remoteOvfUrl(Output.of(remoteOvfUrl));
        }

        public VirtualMachineOvfDeployArgs build() {
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.vsphere.outputs.VirtualMachineCloneCustomizeLinuxOptions;
import com.pulumi.vsphere.outputs.VirtualMachineCloneCustomizeNetworkInterface;
import com.pulumi.vsphere.outputs.VirtualMachineCloneCustomizeWindowsOptions;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VirtualMachineCloneCustomize {
    private @Nullable List<String> dnsServerLists;
    private @Nullable List<String> dnsSuffixLists;
    private @Nullable String ipv4Gateway;
    private @Nullable String ipv6Gateway;
    private @Nullable VirtualMachineCloneCustomizeLinuxOptions linuxOptions;
    /**
     * @return A specification for a virtual NIC on the virtual machine. See network interface options for more information.
     * 
     */
    private @Nullable List<VirtualMachineCloneCustomizeNetworkInterface> networkInterfaces;
    private @Nullable Integer timeout;
    private @Nullable VirtualMachineCloneCustomizeWindowsOptions windowsOptions;
    private @Nullable String windowsSysprepText;

    private VirtualMachineCloneCustomize() {}
    public List<String> dnsServerLists() {
        return this.dnsServerLists == null ? List.of() : this.dnsServerLists;
    }
    public List<String> dnsSuffixLists() {
        return this.dnsSuffixLists == null ? List.of() : this.dnsSuffixLists;
    }
    public Optional<String> ipv4Gateway() {
        return Optional.ofNullable(this.ipv4Gateway);
    }
    public Optional<String> ipv6Gateway() {
        return Optional.ofNullable(this.ipv6Gateway);
    }
    public Optional<VirtualMachineCloneCustomizeLinuxOptions> linuxOptions() {
        return Optional.ofNullable(this.linuxOptions);
    }
    /**
     * @return A specification for a virtual NIC on the virtual machine. See network interface options for more information.
     * 
     */
    public List<VirtualMachineCloneCustomizeNetworkInterface> networkInterfaces() {
        return this.networkInterfaces == null ? List.of() : this.networkInterfaces;
    }
    public Optional<Integer> timeout() {
        return Optional.ofNullable(this.timeout);
    }
    public Optional<VirtualMachineCloneCustomizeWindowsOptions> windowsOptions() {
        return Optional.ofNullable(this.windowsOptions);
    }
    public Optional<String> windowsSysprepText() {
        return Optional.ofNullable(this.windowsSysprepText);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VirtualMachineCloneCustomize defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> dnsServerLists;
        private @Nullable List<String> dnsSuffixLists;
        private @Nullable String ipv4Gateway;
        private @Nullable String ipv6Gateway;
        private @Nullable VirtualMachineCloneCustomizeLinuxOptions linuxOptions;
        private @Nullable List<VirtualMachineCloneCustomizeNetworkInterface> networkInterfaces;
        private @Nullable Integer timeout;
        private @Nullable VirtualMachineCloneCustomizeWindowsOptions windowsOptions;
        private @Nullable String windowsSysprepText;
        public Builder() {}
        public Builder(VirtualMachineCloneCustomize defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dnsServerLists = defaults.dnsServerLists;
    	      this.dnsSuffixLists = defaults.dnsSuffixLists;
    	      this.ipv4Gateway = defaults.ipv4Gateway;
    	      this.ipv6Gateway = defaults.ipv6Gateway;
    	      this.linuxOptions = defaults.linuxOptions;
    	      this.networkInterfaces = defaults.networkInterfaces;
    	      this.timeout = defaults.timeout;
    	      this.windowsOptions = defaults.windowsOptions;
    	      this.windowsSysprepText = defaults.windowsSysprepText;
        }

        @CustomType.Setter
        public Builder dnsServerLists(@Nullable List<String> dnsServerLists) {
            this.dnsServerLists = dnsServerLists;
            return this;
        }
        public Builder dnsServerLists(String... dnsServerLists) {
            return dnsServerLists(List.of(dnsServerLists));
        }
        @CustomType.Setter
        public Builder dnsSuffixLists(@Nullable List<String> dnsSuffixLists) {
            this.dnsSuffixLists = dnsSuffixLists;
            return this;
        }
        public Builder dnsSuffixLists(String... dnsSuffixLists) {
            return dnsSuffixLists(List.of(dnsSuffixLists));
        }
        @CustomType.Setter
        public Builder ipv4Gateway(@Nullable String ipv4Gateway) {
            this.ipv4Gateway = ipv4Gateway;
            return this;
        }
        @CustomType.Setter
        public Builder ipv6Gateway(@Nullable String ipv6Gateway) {
            this.ipv6Gateway = ipv6Gateway;
            return this;
        }
        @CustomType.Setter
        public Builder linuxOptions(@Nullable VirtualMachineCloneCustomizeLinuxOptions linuxOptions) {
            this.linuxOptions = linuxOptions;
            return this;
        }
        @CustomType.Setter
        public Builder networkInterfaces(@Nullable List<VirtualMachineCloneCustomizeNetworkInterface> networkInterfaces) {
            this.networkInterfaces = networkInterfaces;
            return this;
        }
        public Builder networkInterfaces(VirtualMachineCloneCustomizeNetworkInterface... networkInterfaces) {
            return networkInterfaces(List.of(networkInterfaces));
        }
        @CustomType.Setter
        public Builder timeout(@Nullable Integer timeout) {
            this.timeout = timeout;
            return this;
        }
        @CustomType.Setter
        public Builder windowsOptions(@Nullable VirtualMachineCloneCustomizeWindowsOptions windowsOptions) {
            this.windowsOptions = windowsOptions;
            return this;
        }
        @CustomType.Setter
        public Builder windowsSysprepText(@Nullable String windowsSysprepText) {
            this.windowsSysprepText = windowsSysprepText;
            return this;
        }
        public VirtualMachineCloneCustomize build() {
            final var o = new VirtualMachineCloneCustomize();
            o.dnsServerLists = dnsServerLists;
            o.dnsSuffixLists = dnsSuffixLists;
            o.ipv4Gateway = ipv4Gateway;
            o.ipv6Gateway = ipv6Gateway;
            o.linuxOptions = linuxOptions;
            o.networkInterfaces = networkInterfaces;
            o.timeout = timeout;
            o.windowsOptions = windowsOptions;
            o.windowsSysprepText = windowsSysprepText;
            return o;
        }
    }
}
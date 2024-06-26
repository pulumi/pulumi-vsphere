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
    /**
     * @return The list of DNS servers for a virtual network adapter with a static IP address.
     * 
     */
    private @Nullable List<String> dnsServerLists;
    /**
     * @return A list of DNS search domains to add to the DNS configuration on the virtual machine.
     * 
     */
    private @Nullable List<String> dnsSuffixLists;
    /**
     * @return The IPv4 default gateway when using network_interface customization on the virtual machine. This address must be local to a static IPv4 address configured in an interface sub-resource.
     * 
     */
    private @Nullable String ipv4Gateway;
    /**
     * @return The IPv6 default gateway when using network_interface customization on the virtual machine. This address must be local to a static IPv4 address configured in an interface sub-resource.
     * 
     */
    private @Nullable String ipv6Gateway;
    /**
     * @return A list of configuration options specific to Linux virtual machines.
     * 
     */
    private @Nullable VirtualMachineCloneCustomizeLinuxOptions linuxOptions;
    /**
     * @return A specification of network interface configuration options.
     * 
     */
    private @Nullable List<VirtualMachineCloneCustomizeNetworkInterface> networkInterfaces;
    /**
     * @return The amount of time, in minutes, to wait for guest OS customization to complete before returning with an error. Setting this value to 0 or a negative value skips the waiter. Default: 10.
     * 
     */
    private @Nullable Integer timeout;
    /**
     * @return A list of configuration options specific to Windows virtual machines.
     * 
     */
    private @Nullable VirtualMachineCloneCustomizeWindowsOptions windowsOptions;
    /**
     * @return Use this option to specify a windows sysprep file directly.
     * 
     */
    private @Nullable String windowsSysprepText;

    private VirtualMachineCloneCustomize() {}
    /**
     * @return The list of DNS servers for a virtual network adapter with a static IP address.
     * 
     */
    public List<String> dnsServerLists() {
        return this.dnsServerLists == null ? List.of() : this.dnsServerLists;
    }
    /**
     * @return A list of DNS search domains to add to the DNS configuration on the virtual machine.
     * 
     */
    public List<String> dnsSuffixLists() {
        return this.dnsSuffixLists == null ? List.of() : this.dnsSuffixLists;
    }
    /**
     * @return The IPv4 default gateway when using network_interface customization on the virtual machine. This address must be local to a static IPv4 address configured in an interface sub-resource.
     * 
     */
    public Optional<String> ipv4Gateway() {
        return Optional.ofNullable(this.ipv4Gateway);
    }
    /**
     * @return The IPv6 default gateway when using network_interface customization on the virtual machine. This address must be local to a static IPv4 address configured in an interface sub-resource.
     * 
     */
    public Optional<String> ipv6Gateway() {
        return Optional.ofNullable(this.ipv6Gateway);
    }
    /**
     * @return A list of configuration options specific to Linux virtual machines.
     * 
     */
    public Optional<VirtualMachineCloneCustomizeLinuxOptions> linuxOptions() {
        return Optional.ofNullable(this.linuxOptions);
    }
    /**
     * @return A specification of network interface configuration options.
     * 
     */
    public List<VirtualMachineCloneCustomizeNetworkInterface> networkInterfaces() {
        return this.networkInterfaces == null ? List.of() : this.networkInterfaces;
    }
    /**
     * @return The amount of time, in minutes, to wait for guest OS customization to complete before returning with an error. Setting this value to 0 or a negative value skips the waiter. Default: 10.
     * 
     */
    public Optional<Integer> timeout() {
        return Optional.ofNullable(this.timeout);
    }
    /**
     * @return A list of configuration options specific to Windows virtual machines.
     * 
     */
    public Optional<VirtualMachineCloneCustomizeWindowsOptions> windowsOptions() {
        return Optional.ofNullable(this.windowsOptions);
    }
    /**
     * @return Use this option to specify a windows sysprep file directly.
     * 
     */
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
            final var _resultValue = new VirtualMachineCloneCustomize();
            _resultValue.dnsServerLists = dnsServerLists;
            _resultValue.dnsSuffixLists = dnsSuffixLists;
            _resultValue.ipv4Gateway = ipv4Gateway;
            _resultValue.ipv6Gateway = ipv6Gateway;
            _resultValue.linuxOptions = linuxOptions;
            _resultValue.networkInterfaces = networkInterfaces;
            _resultValue.timeout = timeout;
            _resultValue.windowsOptions = windowsOptions;
            _resultValue.windowsSysprepText = windowsSysprepText;
            return _resultValue;
        }
    }
}

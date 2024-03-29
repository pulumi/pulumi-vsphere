// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VirtualMachineCloneCustomizeNetworkInterface {
    /**
     * @return A DNS search domain to add to the DNS configuration on the virtual machine.
     * 
     */
    private @Nullable String dnsDomain;
    /**
     * @return Network-interface specific DNS settings for Windows operating systems. Ignored on Linux.
     * 
     */
    private @Nullable List<String> dnsServerLists;
    /**
     * @return The IPv4 address assigned to this network adapter. If left blank, DHCP is used.
     * 
     */
    private @Nullable String ipv4Address;
    /**
     * @return The IPv4 CIDR netmask for the supplied IP address. Ignored if DHCP is selected.
     * 
     */
    private @Nullable Integer ipv4Netmask;
    /**
     * @return The IPv6 address assigned to this network adapter. If left blank, default auto-configuration is used.
     * 
     */
    private @Nullable String ipv6Address;
    /**
     * @return The IPv6 CIDR netmask for the supplied IP address. Ignored if auto-configuration is selected.
     * 
     */
    private @Nullable Integer ipv6Netmask;

    private VirtualMachineCloneCustomizeNetworkInterface() {}
    /**
     * @return A DNS search domain to add to the DNS configuration on the virtual machine.
     * 
     */
    public Optional<String> dnsDomain() {
        return Optional.ofNullable(this.dnsDomain);
    }
    /**
     * @return Network-interface specific DNS settings for Windows operating systems. Ignored on Linux.
     * 
     */
    public List<String> dnsServerLists() {
        return this.dnsServerLists == null ? List.of() : this.dnsServerLists;
    }
    /**
     * @return The IPv4 address assigned to this network adapter. If left blank, DHCP is used.
     * 
     */
    public Optional<String> ipv4Address() {
        return Optional.ofNullable(this.ipv4Address);
    }
    /**
     * @return The IPv4 CIDR netmask for the supplied IP address. Ignored if DHCP is selected.
     * 
     */
    public Optional<Integer> ipv4Netmask() {
        return Optional.ofNullable(this.ipv4Netmask);
    }
    /**
     * @return The IPv6 address assigned to this network adapter. If left blank, default auto-configuration is used.
     * 
     */
    public Optional<String> ipv6Address() {
        return Optional.ofNullable(this.ipv6Address);
    }
    /**
     * @return The IPv6 CIDR netmask for the supplied IP address. Ignored if auto-configuration is selected.
     * 
     */
    public Optional<Integer> ipv6Netmask() {
        return Optional.ofNullable(this.ipv6Netmask);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VirtualMachineCloneCustomizeNetworkInterface defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String dnsDomain;
        private @Nullable List<String> dnsServerLists;
        private @Nullable String ipv4Address;
        private @Nullable Integer ipv4Netmask;
        private @Nullable String ipv6Address;
        private @Nullable Integer ipv6Netmask;
        public Builder() {}
        public Builder(VirtualMachineCloneCustomizeNetworkInterface defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dnsDomain = defaults.dnsDomain;
    	      this.dnsServerLists = defaults.dnsServerLists;
    	      this.ipv4Address = defaults.ipv4Address;
    	      this.ipv4Netmask = defaults.ipv4Netmask;
    	      this.ipv6Address = defaults.ipv6Address;
    	      this.ipv6Netmask = defaults.ipv6Netmask;
        }

        @CustomType.Setter
        public Builder dnsDomain(@Nullable String dnsDomain) {

            this.dnsDomain = dnsDomain;
            return this;
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
        public Builder ipv4Address(@Nullable String ipv4Address) {

            this.ipv4Address = ipv4Address;
            return this;
        }
        @CustomType.Setter
        public Builder ipv4Netmask(@Nullable Integer ipv4Netmask) {

            this.ipv4Netmask = ipv4Netmask;
            return this;
        }
        @CustomType.Setter
        public Builder ipv6Address(@Nullable String ipv6Address) {

            this.ipv6Address = ipv6Address;
            return this;
        }
        @CustomType.Setter
        public Builder ipv6Netmask(@Nullable Integer ipv6Netmask) {

            this.ipv6Netmask = ipv6Netmask;
            return this;
        }
        public VirtualMachineCloneCustomizeNetworkInterface build() {
            final var _resultValue = new VirtualMachineCloneCustomizeNetworkInterface();
            _resultValue.dnsDomain = dnsDomain;
            _resultValue.dnsServerLists = dnsServerLists;
            _resultValue.ipv4Address = ipv4Address;
            _resultValue.ipv4Netmask = ipv4Netmask;
            _resultValue.ipv6Address = ipv6Address;
            _resultValue.ipv6Netmask = ipv6Netmask;
            return _resultValue;
        }
    }
}

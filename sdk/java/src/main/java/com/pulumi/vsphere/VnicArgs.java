// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.vsphere.inputs.VnicIpv4Args;
import com.pulumi.vsphere.inputs.VnicIpv6Args;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VnicArgs extends com.pulumi.resources.ResourceArgs {

    public static final VnicArgs Empty = new VnicArgs();

    /**
     * Key of the distributed portgroup the nic will connect to.
     * 
     */
    @Import(name="distributedPortGroup")
    private @Nullable Output<String> distributedPortGroup;

    /**
     * @return Key of the distributed portgroup the nic will connect to.
     * 
     */
    public Optional<Output<String>> distributedPortGroup() {
        return Optional.ofNullable(this.distributedPortGroup);
    }

    /**
     * UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
     * 
     */
    @Import(name="distributedSwitchPort")
    private @Nullable Output<String> distributedSwitchPort;

    /**
     * @return UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
     * 
     */
    public Optional<Output<String>> distributedSwitchPort() {
        return Optional.ofNullable(this.distributedSwitchPort);
    }

    /**
     * ESX host the interface belongs to
     * 
     */
    @Import(name="host", required=true)
    private Output<String> host;

    /**
     * @return ESX host the interface belongs to
     * 
     */
    public Output<String> host() {
        return this.host;
    }

    /**
     * IPv4 settings. Either this or `ipv6` needs to be set. See IPv4 options below.
     * 
     */
    @Import(name="ipv4")
    private @Nullable Output<VnicIpv4Args> ipv4;

    /**
     * @return IPv4 settings. Either this or `ipv6` needs to be set. See IPv4 options below.
     * 
     */
    public Optional<Output<VnicIpv4Args>> ipv4() {
        return Optional.ofNullable(this.ipv4);
    }

    /**
     * IPv6 settings. Either this or `ipv6` needs to be set. See IPv6 options below.
     * 
     */
    @Import(name="ipv6")
    private @Nullable Output<VnicIpv6Args> ipv6;

    /**
     * @return IPv6 settings. Either this or `ipv6` needs to be set. See IPv6 options below.
     * 
     */
    public Optional<Output<VnicIpv6Args>> ipv6() {
        return Optional.ofNullable(this.ipv6);
    }

    /**
     * MAC address of the interface.
     * 
     */
    @Import(name="mac")
    private @Nullable Output<String> mac;

    /**
     * @return MAC address of the interface.
     * 
     */
    public Optional<Output<String>> mac() {
        return Optional.ofNullable(this.mac);
    }

    /**
     * MTU of the interface.
     * 
     */
    @Import(name="mtu")
    private @Nullable Output<Integer> mtu;

    /**
     * @return MTU of the interface.
     * 
     */
    public Optional<Output<Integer>> mtu() {
        return Optional.ofNullable(this.mtu);
    }

    /**
     * TCP/IP stack setting for this interface. Possible values are `defaultTcpipStack``, &#39;vmotion&#39;, &#39;vSphereProvisioning&#39;. Changing this will force the creation of a new interface since it&#39;s not possible to change the stack once it gets created. (Default:`defaultTcpipStack`)
     * 
     */
    @Import(name="netstack")
    private @Nullable Output<String> netstack;

    /**
     * @return TCP/IP stack setting for this interface. Possible values are `defaultTcpipStack``, &#39;vmotion&#39;, &#39;vSphereProvisioning&#39;. Changing this will force the creation of a new interface since it&#39;s not possible to change the stack once it gets created. (Default:`defaultTcpipStack`)
     * 
     */
    public Optional<Output<String>> netstack() {
        return Optional.ofNullable(this.netstack);
    }

    /**
     * Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
     * 
     */
    @Import(name="portgroup")
    private @Nullable Output<String> portgroup;

    /**
     * @return Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
     * 
     */
    public Optional<Output<String>> portgroup() {
        return Optional.ofNullable(this.portgroup);
    }

    /**
     * Enabled services setting for this interface. Currently support values are `vmotion`, `management`, and `vsan`.
     * 
     */
    @Import(name="services")
    private @Nullable Output<List<String>> services;

    /**
     * @return Enabled services setting for this interface. Currently support values are `vmotion`, `management`, and `vsan`.
     * 
     */
    public Optional<Output<List<String>>> services() {
        return Optional.ofNullable(this.services);
    }

    private VnicArgs() {}

    private VnicArgs(VnicArgs $) {
        this.distributedPortGroup = $.distributedPortGroup;
        this.distributedSwitchPort = $.distributedSwitchPort;
        this.host = $.host;
        this.ipv4 = $.ipv4;
        this.ipv6 = $.ipv6;
        this.mac = $.mac;
        this.mtu = $.mtu;
        this.netstack = $.netstack;
        this.portgroup = $.portgroup;
        this.services = $.services;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VnicArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VnicArgs $;

        public Builder() {
            $ = new VnicArgs();
        }

        public Builder(VnicArgs defaults) {
            $ = new VnicArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param distributedPortGroup Key of the distributed portgroup the nic will connect to.
         * 
         * @return builder
         * 
         */
        public Builder distributedPortGroup(@Nullable Output<String> distributedPortGroup) {
            $.distributedPortGroup = distributedPortGroup;
            return this;
        }

        /**
         * @param distributedPortGroup Key of the distributed portgroup the nic will connect to.
         * 
         * @return builder
         * 
         */
        public Builder distributedPortGroup(String distributedPortGroup) {
            return distributedPortGroup(Output.of(distributedPortGroup));
        }

        /**
         * @param distributedSwitchPort UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
         * 
         * @return builder
         * 
         */
        public Builder distributedSwitchPort(@Nullable Output<String> distributedSwitchPort) {
            $.distributedSwitchPort = distributedSwitchPort;
            return this;
        }

        /**
         * @param distributedSwitchPort UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
         * 
         * @return builder
         * 
         */
        public Builder distributedSwitchPort(String distributedSwitchPort) {
            return distributedSwitchPort(Output.of(distributedSwitchPort));
        }

        /**
         * @param host ESX host the interface belongs to
         * 
         * @return builder
         * 
         */
        public Builder host(Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host ESX host the interface belongs to
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param ipv4 IPv4 settings. Either this or `ipv6` needs to be set. See IPv4 options below.
         * 
         * @return builder
         * 
         */
        public Builder ipv4(@Nullable Output<VnicIpv4Args> ipv4) {
            $.ipv4 = ipv4;
            return this;
        }

        /**
         * @param ipv4 IPv4 settings. Either this or `ipv6` needs to be set. See IPv4 options below.
         * 
         * @return builder
         * 
         */
        public Builder ipv4(VnicIpv4Args ipv4) {
            return ipv4(Output.of(ipv4));
        }

        /**
         * @param ipv6 IPv6 settings. Either this or `ipv6` needs to be set. See IPv6 options below.
         * 
         * @return builder
         * 
         */
        public Builder ipv6(@Nullable Output<VnicIpv6Args> ipv6) {
            $.ipv6 = ipv6;
            return this;
        }

        /**
         * @param ipv6 IPv6 settings. Either this or `ipv6` needs to be set. See IPv6 options below.
         * 
         * @return builder
         * 
         */
        public Builder ipv6(VnicIpv6Args ipv6) {
            return ipv6(Output.of(ipv6));
        }

        /**
         * @param mac MAC address of the interface.
         * 
         * @return builder
         * 
         */
        public Builder mac(@Nullable Output<String> mac) {
            $.mac = mac;
            return this;
        }

        /**
         * @param mac MAC address of the interface.
         * 
         * @return builder
         * 
         */
        public Builder mac(String mac) {
            return mac(Output.of(mac));
        }

        /**
         * @param mtu MTU of the interface.
         * 
         * @return builder
         * 
         */
        public Builder mtu(@Nullable Output<Integer> mtu) {
            $.mtu = mtu;
            return this;
        }

        /**
         * @param mtu MTU of the interface.
         * 
         * @return builder
         * 
         */
        public Builder mtu(Integer mtu) {
            return mtu(Output.of(mtu));
        }

        /**
         * @param netstack TCP/IP stack setting for this interface. Possible values are `defaultTcpipStack``, &#39;vmotion&#39;, &#39;vSphereProvisioning&#39;. Changing this will force the creation of a new interface since it&#39;s not possible to change the stack once it gets created. (Default:`defaultTcpipStack`)
         * 
         * @return builder
         * 
         */
        public Builder netstack(@Nullable Output<String> netstack) {
            $.netstack = netstack;
            return this;
        }

        /**
         * @param netstack TCP/IP stack setting for this interface. Possible values are `defaultTcpipStack``, &#39;vmotion&#39;, &#39;vSphereProvisioning&#39;. Changing this will force the creation of a new interface since it&#39;s not possible to change the stack once it gets created. (Default:`defaultTcpipStack`)
         * 
         * @return builder
         * 
         */
        public Builder netstack(String netstack) {
            return netstack(Output.of(netstack));
        }

        /**
         * @param portgroup Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
         * 
         * @return builder
         * 
         */
        public Builder portgroup(@Nullable Output<String> portgroup) {
            $.portgroup = portgroup;
            return this;
        }

        /**
         * @param portgroup Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
         * 
         * @return builder
         * 
         */
        public Builder portgroup(String portgroup) {
            return portgroup(Output.of(portgroup));
        }

        /**
         * @param services Enabled services setting for this interface. Currently support values are `vmotion`, `management`, and `vsan`.
         * 
         * @return builder
         * 
         */
        public Builder services(@Nullable Output<List<String>> services) {
            $.services = services;
            return this;
        }

        /**
         * @param services Enabled services setting for this interface. Currently support values are `vmotion`, `management`, and `vsan`.
         * 
         * @return builder
         * 
         */
        public Builder services(List<String> services) {
            return services(Output.of(services));
        }

        /**
         * @param services Enabled services setting for this interface. Currently support values are `vmotion`, `management`, and `vsan`.
         * 
         * @return builder
         * 
         */
        public Builder services(String... services) {
            return services(List.of(services));
        }

        public VnicArgs build() {
            $.host = Objects.requireNonNull($.host, "expected parameter 'host' to be non-null");
            return $;
        }
    }

}

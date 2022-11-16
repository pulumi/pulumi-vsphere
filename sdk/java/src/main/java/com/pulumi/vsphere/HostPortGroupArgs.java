// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HostPortGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final HostPortGroupArgs Empty = new HostPortGroupArgs();

    /**
     * List of active network adapters used for load balancing.
     * 
     */
    @Import(name="activeNics")
    private @Nullable Output<List<String>> activeNics;

    /**
     * @return List of active network adapters used for load balancing.
     * 
     */
    public Optional<Output<List<String>>> activeNics() {
        return Optional.ofNullable(this.activeNics);
    }

    /**
     * Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
     * that of its own.
     * 
     */
    @Import(name="allowForgedTransmits")
    private @Nullable Output<Boolean> allowForgedTransmits;

    /**
     * @return Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
     * that of its own.
     * 
     */
    public Optional<Output<Boolean>> allowForgedTransmits() {
        return Optional.ofNullable(this.allowForgedTransmits);
    }

    /**
     * Controls whether or not the Media Access Control (MAC) address can be changed.
     * 
     */
    @Import(name="allowMacChanges")
    private @Nullable Output<Boolean> allowMacChanges;

    /**
     * @return Controls whether or not the Media Access Control (MAC) address can be changed.
     * 
     */
    public Optional<Output<Boolean>> allowMacChanges() {
        return Optional.ofNullable(this.allowMacChanges);
    }

    /**
     * Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
     * 
     */
    @Import(name="allowPromiscuous")
    private @Nullable Output<Boolean> allowPromiscuous;

    /**
     * @return Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
     * 
     */
    public Optional<Output<Boolean>> allowPromiscuous() {
        return Optional.ofNullable(this.allowPromiscuous);
    }

    /**
     * Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
     * only.
     * 
     */
    @Import(name="checkBeacon")
    private @Nullable Output<Boolean> checkBeacon;

    /**
     * @return Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
     * only.
     * 
     */
    public Optional<Output<Boolean>> checkBeacon() {
        return Optional.ofNullable(this.checkBeacon);
    }

    /**
     * If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
     * 
     */
    @Import(name="failback")
    private @Nullable Output<Boolean> failback;

    /**
     * @return If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
     * 
     */
    public Optional<Output<Boolean>> failback() {
        return Optional.ofNullable(this.failback);
    }

    /**
     * The managed object ID of
     * the host to set the port group up on. Forces a new resource if changed.
     * 
     */
    @Import(name="hostSystemId", required=true)
    private Output<String> hostSystemId;

    /**
     * @return The managed object ID of
     * the host to set the port group up on. Forces a new resource if changed.
     * 
     */
    public Output<String> hostSystemId() {
        return this.hostSystemId;
    }

    /**
     * The name of the port group.  Forces a new resource if
     * changed.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the port group.  Forces a new resource if
     * changed.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
     * 
     */
    @Import(name="notifySwitches")
    private @Nullable Output<Boolean> notifySwitches;

    /**
     * @return If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
     * 
     */
    public Optional<Output<Boolean>> notifySwitches() {
        return Optional.ofNullable(this.notifySwitches);
    }

    /**
     * The average bandwidth in bits per second if traffic shaping is enabled.
     * 
     */
    @Import(name="shapingAverageBandwidth")
    private @Nullable Output<Integer> shapingAverageBandwidth;

    /**
     * @return The average bandwidth in bits per second if traffic shaping is enabled.
     * 
     */
    public Optional<Output<Integer>> shapingAverageBandwidth() {
        return Optional.ofNullable(this.shapingAverageBandwidth);
    }

    /**
     * The maximum burst size allowed in bytes if traffic shaping is enabled.
     * 
     */
    @Import(name="shapingBurstSize")
    private @Nullable Output<Integer> shapingBurstSize;

    /**
     * @return The maximum burst size allowed in bytes if traffic shaping is enabled.
     * 
     */
    public Optional<Output<Integer>> shapingBurstSize() {
        return Optional.ofNullable(this.shapingBurstSize);
    }

    /**
     * Enable traffic shaping on this virtual switch or port group.
     * 
     */
    @Import(name="shapingEnabled")
    private @Nullable Output<Boolean> shapingEnabled;

    /**
     * @return Enable traffic shaping on this virtual switch or port group.
     * 
     */
    public Optional<Output<Boolean>> shapingEnabled() {
        return Optional.ofNullable(this.shapingEnabled);
    }

    /**
     * The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
     * 
     */
    @Import(name="shapingPeakBandwidth")
    private @Nullable Output<Integer> shapingPeakBandwidth;

    /**
     * @return The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
     * 
     */
    public Optional<Output<Integer>> shapingPeakBandwidth() {
        return Optional.ofNullable(this.shapingPeakBandwidth);
    }

    /**
     * List of standby network adapters used for failover.
     * 
     */
    @Import(name="standbyNics")
    private @Nullable Output<List<String>> standbyNics;

    /**
     * @return List of standby network adapters used for failover.
     * 
     */
    public Optional<Output<List<String>>> standbyNics() {
        return Optional.ofNullable(this.standbyNics);
    }

    /**
     * The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
     * failover_explicit.
     * 
     */
    @Import(name="teamingPolicy")
    private @Nullable Output<String> teamingPolicy;

    /**
     * @return The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
     * failover_explicit.
     * 
     */
    public Optional<Output<String>> teamingPolicy() {
        return Optional.ofNullable(this.teamingPolicy);
    }

    /**
     * The name of the virtual switch to bind
     * this port group to. Forces a new resource if changed.
     * 
     */
    @Import(name="virtualSwitchName", required=true)
    private Output<String> virtualSwitchName;

    /**
     * @return The name of the virtual switch to bind
     * this port group to. Forces a new resource if changed.
     * 
     */
    public Output<String> virtualSwitchName() {
        return this.virtualSwitchName;
    }

    /**
     * The VLAN ID/trunk mode for this port group.  An ID of
     * `0` denotes no tagging, an ID of `1`-`4094` tags with the specific ID, and an
     * ID of `4095` enables trunk mode, allowing the guest to manage its own
     * tagging. Default: `0`.
     * 
     */
    @Import(name="vlanId")
    private @Nullable Output<Integer> vlanId;

    /**
     * @return The VLAN ID/trunk mode for this port group.  An ID of
     * `0` denotes no tagging, an ID of `1`-`4094` tags with the specific ID, and an
     * ID of `4095` enables trunk mode, allowing the guest to manage its own
     * tagging. Default: `0`.
     * 
     */
    public Optional<Output<Integer>> vlanId() {
        return Optional.ofNullable(this.vlanId);
    }

    private HostPortGroupArgs() {}

    private HostPortGroupArgs(HostPortGroupArgs $) {
        this.activeNics = $.activeNics;
        this.allowForgedTransmits = $.allowForgedTransmits;
        this.allowMacChanges = $.allowMacChanges;
        this.allowPromiscuous = $.allowPromiscuous;
        this.checkBeacon = $.checkBeacon;
        this.failback = $.failback;
        this.hostSystemId = $.hostSystemId;
        this.name = $.name;
        this.notifySwitches = $.notifySwitches;
        this.shapingAverageBandwidth = $.shapingAverageBandwidth;
        this.shapingBurstSize = $.shapingBurstSize;
        this.shapingEnabled = $.shapingEnabled;
        this.shapingPeakBandwidth = $.shapingPeakBandwidth;
        this.standbyNics = $.standbyNics;
        this.teamingPolicy = $.teamingPolicy;
        this.virtualSwitchName = $.virtualSwitchName;
        this.vlanId = $.vlanId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HostPortGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HostPortGroupArgs $;

        public Builder() {
            $ = new HostPortGroupArgs();
        }

        public Builder(HostPortGroupArgs defaults) {
            $ = new HostPortGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param activeNics List of active network adapters used for load balancing.
         * 
         * @return builder
         * 
         */
        public Builder activeNics(@Nullable Output<List<String>> activeNics) {
            $.activeNics = activeNics;
            return this;
        }

        /**
         * @param activeNics List of active network adapters used for load balancing.
         * 
         * @return builder
         * 
         */
        public Builder activeNics(List<String> activeNics) {
            return activeNics(Output.of(activeNics));
        }

        /**
         * @param activeNics List of active network adapters used for load balancing.
         * 
         * @return builder
         * 
         */
        public Builder activeNics(String... activeNics) {
            return activeNics(List.of(activeNics));
        }

        /**
         * @param allowForgedTransmits Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
         * that of its own.
         * 
         * @return builder
         * 
         */
        public Builder allowForgedTransmits(@Nullable Output<Boolean> allowForgedTransmits) {
            $.allowForgedTransmits = allowForgedTransmits;
            return this;
        }

        /**
         * @param allowForgedTransmits Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
         * that of its own.
         * 
         * @return builder
         * 
         */
        public Builder allowForgedTransmits(Boolean allowForgedTransmits) {
            return allowForgedTransmits(Output.of(allowForgedTransmits));
        }

        /**
         * @param allowMacChanges Controls whether or not the Media Access Control (MAC) address can be changed.
         * 
         * @return builder
         * 
         */
        public Builder allowMacChanges(@Nullable Output<Boolean> allowMacChanges) {
            $.allowMacChanges = allowMacChanges;
            return this;
        }

        /**
         * @param allowMacChanges Controls whether or not the Media Access Control (MAC) address can be changed.
         * 
         * @return builder
         * 
         */
        public Builder allowMacChanges(Boolean allowMacChanges) {
            return allowMacChanges(Output.of(allowMacChanges));
        }

        /**
         * @param allowPromiscuous Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
         * 
         * @return builder
         * 
         */
        public Builder allowPromiscuous(@Nullable Output<Boolean> allowPromiscuous) {
            $.allowPromiscuous = allowPromiscuous;
            return this;
        }

        /**
         * @param allowPromiscuous Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
         * 
         * @return builder
         * 
         */
        public Builder allowPromiscuous(Boolean allowPromiscuous) {
            return allowPromiscuous(Output.of(allowPromiscuous));
        }

        /**
         * @param checkBeacon Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
         * only.
         * 
         * @return builder
         * 
         */
        public Builder checkBeacon(@Nullable Output<Boolean> checkBeacon) {
            $.checkBeacon = checkBeacon;
            return this;
        }

        /**
         * @param checkBeacon Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
         * only.
         * 
         * @return builder
         * 
         */
        public Builder checkBeacon(Boolean checkBeacon) {
            return checkBeacon(Output.of(checkBeacon));
        }

        /**
         * @param failback If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
         * 
         * @return builder
         * 
         */
        public Builder failback(@Nullable Output<Boolean> failback) {
            $.failback = failback;
            return this;
        }

        /**
         * @param failback If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
         * 
         * @return builder
         * 
         */
        public Builder failback(Boolean failback) {
            return failback(Output.of(failback));
        }

        /**
         * @param hostSystemId The managed object ID of
         * the host to set the port group up on. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder hostSystemId(Output<String> hostSystemId) {
            $.hostSystemId = hostSystemId;
            return this;
        }

        /**
         * @param hostSystemId The managed object ID of
         * the host to set the port group up on. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder hostSystemId(String hostSystemId) {
            return hostSystemId(Output.of(hostSystemId));
        }

        /**
         * @param name The name of the port group.  Forces a new resource if
         * changed.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the port group.  Forces a new resource if
         * changed.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param notifySwitches If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
         * 
         * @return builder
         * 
         */
        public Builder notifySwitches(@Nullable Output<Boolean> notifySwitches) {
            $.notifySwitches = notifySwitches;
            return this;
        }

        /**
         * @param notifySwitches If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
         * 
         * @return builder
         * 
         */
        public Builder notifySwitches(Boolean notifySwitches) {
            return notifySwitches(Output.of(notifySwitches));
        }

        /**
         * @param shapingAverageBandwidth The average bandwidth in bits per second if traffic shaping is enabled.
         * 
         * @return builder
         * 
         */
        public Builder shapingAverageBandwidth(@Nullable Output<Integer> shapingAverageBandwidth) {
            $.shapingAverageBandwidth = shapingAverageBandwidth;
            return this;
        }

        /**
         * @param shapingAverageBandwidth The average bandwidth in bits per second if traffic shaping is enabled.
         * 
         * @return builder
         * 
         */
        public Builder shapingAverageBandwidth(Integer shapingAverageBandwidth) {
            return shapingAverageBandwidth(Output.of(shapingAverageBandwidth));
        }

        /**
         * @param shapingBurstSize The maximum burst size allowed in bytes if traffic shaping is enabled.
         * 
         * @return builder
         * 
         */
        public Builder shapingBurstSize(@Nullable Output<Integer> shapingBurstSize) {
            $.shapingBurstSize = shapingBurstSize;
            return this;
        }

        /**
         * @param shapingBurstSize The maximum burst size allowed in bytes if traffic shaping is enabled.
         * 
         * @return builder
         * 
         */
        public Builder shapingBurstSize(Integer shapingBurstSize) {
            return shapingBurstSize(Output.of(shapingBurstSize));
        }

        /**
         * @param shapingEnabled Enable traffic shaping on this virtual switch or port group.
         * 
         * @return builder
         * 
         */
        public Builder shapingEnabled(@Nullable Output<Boolean> shapingEnabled) {
            $.shapingEnabled = shapingEnabled;
            return this;
        }

        /**
         * @param shapingEnabled Enable traffic shaping on this virtual switch or port group.
         * 
         * @return builder
         * 
         */
        public Builder shapingEnabled(Boolean shapingEnabled) {
            return shapingEnabled(Output.of(shapingEnabled));
        }

        /**
         * @param shapingPeakBandwidth The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
         * 
         * @return builder
         * 
         */
        public Builder shapingPeakBandwidth(@Nullable Output<Integer> shapingPeakBandwidth) {
            $.shapingPeakBandwidth = shapingPeakBandwidth;
            return this;
        }

        /**
         * @param shapingPeakBandwidth The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
         * 
         * @return builder
         * 
         */
        public Builder shapingPeakBandwidth(Integer shapingPeakBandwidth) {
            return shapingPeakBandwidth(Output.of(shapingPeakBandwidth));
        }

        /**
         * @param standbyNics List of standby network adapters used for failover.
         * 
         * @return builder
         * 
         */
        public Builder standbyNics(@Nullable Output<List<String>> standbyNics) {
            $.standbyNics = standbyNics;
            return this;
        }

        /**
         * @param standbyNics List of standby network adapters used for failover.
         * 
         * @return builder
         * 
         */
        public Builder standbyNics(List<String> standbyNics) {
            return standbyNics(Output.of(standbyNics));
        }

        /**
         * @param standbyNics List of standby network adapters used for failover.
         * 
         * @return builder
         * 
         */
        public Builder standbyNics(String... standbyNics) {
            return standbyNics(List.of(standbyNics));
        }

        /**
         * @param teamingPolicy The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
         * failover_explicit.
         * 
         * @return builder
         * 
         */
        public Builder teamingPolicy(@Nullable Output<String> teamingPolicy) {
            $.teamingPolicy = teamingPolicy;
            return this;
        }

        /**
         * @param teamingPolicy The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
         * failover_explicit.
         * 
         * @return builder
         * 
         */
        public Builder teamingPolicy(String teamingPolicy) {
            return teamingPolicy(Output.of(teamingPolicy));
        }

        /**
         * @param virtualSwitchName The name of the virtual switch to bind
         * this port group to. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder virtualSwitchName(Output<String> virtualSwitchName) {
            $.virtualSwitchName = virtualSwitchName;
            return this;
        }

        /**
         * @param virtualSwitchName The name of the virtual switch to bind
         * this port group to. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder virtualSwitchName(String virtualSwitchName) {
            return virtualSwitchName(Output.of(virtualSwitchName));
        }

        /**
         * @param vlanId The VLAN ID/trunk mode for this port group.  An ID of
         * `0` denotes no tagging, an ID of `1`-`4094` tags with the specific ID, and an
         * ID of `4095` enables trunk mode, allowing the guest to manage its own
         * tagging. Default: `0`.
         * 
         * @return builder
         * 
         */
        public Builder vlanId(@Nullable Output<Integer> vlanId) {
            $.vlanId = vlanId;
            return this;
        }

        /**
         * @param vlanId The VLAN ID/trunk mode for this port group.  An ID of
         * `0` denotes no tagging, an ID of `1`-`4094` tags with the specific ID, and an
         * ID of `4095` enables trunk mode, allowing the guest to manage its own
         * tagging. Default: `0`.
         * 
         * @return builder
         * 
         */
        public Builder vlanId(Integer vlanId) {
            return vlanId(Output.of(vlanId));
        }

        public HostPortGroupArgs build() {
            $.hostSystemId = Objects.requireNonNull($.hostSystemId, "expected parameter 'hostSystemId' to be non-null");
            $.virtualSwitchName = Objects.requireNonNull($.virtualSwitchName, "expected parameter 'virtualSwitchName' to be non-null");
            return $;
        }
    }

}

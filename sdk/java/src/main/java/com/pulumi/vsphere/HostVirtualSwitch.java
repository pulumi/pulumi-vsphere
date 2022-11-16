// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.HostVirtualSwitchArgs;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.inputs.HostVirtualSwitchState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="vsphere:index/hostVirtualSwitch:HostVirtualSwitch")
public class HostVirtualSwitch extends com.pulumi.resources.CustomResource {
    /**
     * The list of active network adapters used for load
     * balancing.
     * 
     */
    @Export(name="activeNics", type=List.class, parameters={String.class})
    private Output<List<String>> activeNics;

    /**
     * @return The list of active network adapters used for load
     * balancing.
     * 
     */
    public Output<List<String>> activeNics() {
        return this.activeNics;
    }
    /**
     * Controls whether or not the virtual
     * network adapter is allowed to send network traffic with a different MAC
     * address than that of its own. Default: `true`.
     * 
     */
    @Export(name="allowForgedTransmits", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> allowForgedTransmits;

    /**
     * @return Controls whether or not the virtual
     * network adapter is allowed to send network traffic with a different MAC
     * address than that of its own. Default: `true`.
     * 
     */
    public Output<Optional<Boolean>> allowForgedTransmits() {
        return Codegen.optional(this.allowForgedTransmits);
    }
    /**
     * Controls whether or not the Media Access
     * Control (MAC) address can be changed. Default: `true`.
     * 
     */
    @Export(name="allowMacChanges", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> allowMacChanges;

    /**
     * @return Controls whether or not the Media Access
     * Control (MAC) address can be changed. Default: `true`.
     * 
     */
    public Output<Optional<Boolean>> allowMacChanges() {
        return Codegen.optional(this.allowMacChanges);
    }
    /**
     * Enable promiscuous mode on the network. This
     * flag indicates whether or not all traffic is seen on a given port. Default:
     * `false`.
     * 
     */
    @Export(name="allowPromiscuous", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> allowPromiscuous;

    /**
     * @return Enable promiscuous mode on the network. This
     * flag indicates whether or not all traffic is seen on a given port. Default:
     * `false`.
     * 
     */
    public Output<Optional<Boolean>> allowPromiscuous() {
        return Codegen.optional(this.allowPromiscuous);
    }
    /**
     * The interval, in seconds, that a NIC beacon
     * packet is sent out. This can be used with `check_beacon` to
     * offer link failure capability beyond link status only. Default: `1`.
     * 
     */
    @Export(name="beaconInterval", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> beaconInterval;

    /**
     * @return The interval, in seconds, that a NIC beacon
     * packet is sent out. This can be used with `check_beacon` to
     * offer link failure capability beyond link status only. Default: `1`.
     * 
     */
    public Output<Optional<Integer>> beaconInterval() {
        return Codegen.optional(this.beaconInterval);
    }
    /**
     * Enable beacon probing - this requires that the
     * `beacon_interval` option has been set in the bridge
     * options. If this is set to `false`, only link status is used to check for
     * failed NICs.  Default: `false`.
     * 
     */
    @Export(name="checkBeacon", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> checkBeacon;

    /**
     * @return Enable beacon probing - this requires that the
     * `beacon_interval` option has been set in the bridge
     * options. If this is set to `false`, only link status is used to check for
     * failed NICs.  Default: `false`.
     * 
     */
    public Output<Optional<Boolean>> checkBeacon() {
        return Codegen.optional(this.checkBeacon);
    }
    /**
     * If set to `true`, the teaming policy will re-activate
     * failed interfaces higher in precedence when they come back up.  Default:
     * `true`.
     * 
     */
    @Export(name="failback", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> failback;

    /**
     * @return If set to `true`, the teaming policy will re-activate
     * failed interfaces higher in precedence when they come back up.  Default:
     * `true`.
     * 
     */
    public Output<Optional<Boolean>> failback() {
        return Codegen.optional(this.failback);
    }
    /**
     * The managed object ID of
     * the host to set the virtual switch up on. Forces a new resource if changed.
     * 
     */
    @Export(name="hostSystemId", type=String.class, parameters={})
    private Output<String> hostSystemId;

    /**
     * @return The managed object ID of
     * the host to set the virtual switch up on. Forces a new resource if changed.
     * 
     */
    public Output<String> hostSystemId() {
        return this.hostSystemId;
    }
    /**
     * Whether to `advertise` or `listen`
     * for link discovery traffic. Default: `listen`.
     * 
     */
    @Export(name="linkDiscoveryOperation", type=String.class, parameters={})
    private Output</* @Nullable */ String> linkDiscoveryOperation;

    /**
     * @return Whether to `advertise` or `listen`
     * for link discovery traffic. Default: `listen`.
     * 
     */
    public Output<Optional<String>> linkDiscoveryOperation() {
        return Codegen.optional(this.linkDiscoveryOperation);
    }
    /**
     * The discovery protocol type.  Valid
     * types are `cpd` and `lldp`. Default: `cdp`.
     * 
     */
    @Export(name="linkDiscoveryProtocol", type=String.class, parameters={})
    private Output</* @Nullable */ String> linkDiscoveryProtocol;

    /**
     * @return The discovery protocol type.  Valid
     * types are `cpd` and `lldp`. Default: `cdp`.
     * 
     */
    public Output<Optional<String>> linkDiscoveryProtocol() {
        return Codegen.optional(this.linkDiscoveryProtocol);
    }
    /**
     * The maximum transmission unit (MTU) for the virtual
     * switch. Default: `1500`.
     * 
     */
    @Export(name="mtu", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> mtu;

    /**
     * @return The maximum transmission unit (MTU) for the virtual
     * switch. Default: `1500`.
     * 
     */
    public Output<Optional<Integer>> mtu() {
        return Codegen.optional(this.mtu);
    }
    /**
     * The name of the virtual switch. Forces a new resource if
     * changed.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the virtual switch. Forces a new resource if
     * changed.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The network interfaces to bind to the bridge.
     * 
     */
    @Export(name="networkAdapters", type=List.class, parameters={String.class})
    private Output<List<String>> networkAdapters;

    /**
     * @return The network interfaces to bind to the bridge.
     * 
     */
    public Output<List<String>> networkAdapters() {
        return this.networkAdapters;
    }
    /**
     * If set to `true`, the teaming policy will
     * notify the broadcast network of a NIC failover, triggering cache updates.
     * Default: `true`.
     * 
     */
    @Export(name="notifySwitches", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> notifySwitches;

    /**
     * @return If set to `true`, the teaming policy will
     * notify the broadcast network of a NIC failover, triggering cache updates.
     * Default: `true`.
     * 
     */
    public Output<Optional<Boolean>> notifySwitches() {
        return Codegen.optional(this.notifySwitches);
    }
    /**
     * The number of ports to create with this
     * virtual switch. Default: `128`.
     * 
     */
    @Export(name="numberOfPorts", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> numberOfPorts;

    /**
     * @return The number of ports to create with this
     * virtual switch. Default: `128`.
     * 
     */
    public Output<Optional<Integer>> numberOfPorts() {
        return Codegen.optional(this.numberOfPorts);
    }
    /**
     * The average bandwidth in bits per
     * second if traffic shaping is enabled. Default: `0`
     * 
     */
    @Export(name="shapingAverageBandwidth", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> shapingAverageBandwidth;

    /**
     * @return The average bandwidth in bits per
     * second if traffic shaping is enabled. Default: `0`
     * 
     */
    public Output<Optional<Integer>> shapingAverageBandwidth() {
        return Codegen.optional(this.shapingAverageBandwidth);
    }
    /**
     * The maximum burst size allowed in bytes if
     * shaping is enabled. Default: `0`
     * 
     */
    @Export(name="shapingBurstSize", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> shapingBurstSize;

    /**
     * @return The maximum burst size allowed in bytes if
     * shaping is enabled. Default: `0`
     * 
     */
    public Output<Optional<Integer>> shapingBurstSize() {
        return Codegen.optional(this.shapingBurstSize);
    }
    /**
     * Set to `true` to enable the traffic shaper for
     * ports managed by this virtual switch. Default: `false`.
     * 
     */
    @Export(name="shapingEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> shapingEnabled;

    /**
     * @return Set to `true` to enable the traffic shaper for
     * ports managed by this virtual switch. Default: `false`.
     * 
     */
    public Output<Optional<Boolean>> shapingEnabled() {
        return Codegen.optional(this.shapingEnabled);
    }
    /**
     * The peak bandwidth during bursts in
     * bits per second if traffic shaping is enabled. Default: `0`
     * 
     */
    @Export(name="shapingPeakBandwidth", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> shapingPeakBandwidth;

    /**
     * @return The peak bandwidth during bursts in
     * bits per second if traffic shaping is enabled. Default: `0`
     * 
     */
    public Output<Optional<Integer>> shapingPeakBandwidth() {
        return Codegen.optional(this.shapingPeakBandwidth);
    }
    /**
     * The list of standby network adapters used for
     * failover.
     * 
     */
    @Export(name="standbyNics", type=List.class, parameters={String.class})
    private Output<List<String>> standbyNics;

    /**
     * @return The list of standby network adapters used for
     * failover.
     * 
     */
    public Output<List<String>> standbyNics() {
        return this.standbyNics;
    }
    /**
     * The network adapter teaming policy. Can be one
     * of `loadbalance_ip`, `loadbalance_srcmac`, `loadbalance_srcid`, or
     * `failover_explicit`. Default: `loadbalance_srcid`.
     * 
     */
    @Export(name="teamingPolicy", type=String.class, parameters={})
    private Output</* @Nullable */ String> teamingPolicy;

    /**
     * @return The network adapter teaming policy. Can be one
     * of `loadbalance_ip`, `loadbalance_srcmac`, `loadbalance_srcid`, or
     * `failover_explicit`. Default: `loadbalance_srcid`.
     * 
     */
    public Output<Optional<String>> teamingPolicy() {
        return Codegen.optional(this.teamingPolicy);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HostVirtualSwitch(String name) {
        this(name, HostVirtualSwitchArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HostVirtualSwitch(String name, HostVirtualSwitchArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HostVirtualSwitch(String name, HostVirtualSwitchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/hostVirtualSwitch:HostVirtualSwitch", name, args == null ? HostVirtualSwitchArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HostVirtualSwitch(String name, Output<String> id, @Nullable HostVirtualSwitchState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/hostVirtualSwitch:HostVirtualSwitch", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static HostVirtualSwitch get(String name, Output<String> id, @Nullable HostVirtualSwitchState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HostVirtualSwitch(name, id, state, options);
    }
}
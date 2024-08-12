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
     * List of active network adapters used for load balancing.
     * 
     */
    @Export(name="activeNics", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> activeNics;

    /**
     * @return List of active network adapters used for load balancing.
     * 
     */
    public Output<List<String>> activeNics() {
        return this.activeNics;
    }
    /**
     * Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
     * that of its own.
     * 
     */
    @Export(name="allowForgedTransmits", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowForgedTransmits;

    /**
     * @return Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than
     * that of its own.
     * 
     */
    public Output<Optional<Boolean>> allowForgedTransmits() {
        return Codegen.optional(this.allowForgedTransmits);
    }
    /**
     * Controls whether or not the Media Access Control (MAC) address can be changed.
     * 
     */
    @Export(name="allowMacChanges", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowMacChanges;

    /**
     * @return Controls whether or not the Media Access Control (MAC) address can be changed.
     * 
     */
    public Output<Optional<Boolean>> allowMacChanges() {
        return Codegen.optional(this.allowMacChanges);
    }
    /**
     * Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
     * 
     */
    @Export(name="allowPromiscuous", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowPromiscuous;

    /**
     * @return Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port.
     * 
     */
    public Output<Optional<Boolean>> allowPromiscuous() {
        return Codegen.optional(this.allowPromiscuous);
    }
    /**
     * Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
     * 
     */
    @Export(name="beaconInterval", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> beaconInterval;

    /**
     * @return Determines how often, in seconds, a beacon should be sent to probe for the validity of a link.
     * 
     */
    public Output<Optional<Integer>> beaconInterval() {
        return Codegen.optional(this.beaconInterval);
    }
    /**
     * Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
     * only.
     * 
     */
    @Export(name="checkBeacon", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> checkBeacon;

    /**
     * @return Enable beacon probing. Requires that the vSwitch has been configured to use a beacon. If disabled, link status is used
     * only.
     * 
     */
    public Output<Optional<Boolean>> checkBeacon() {
        return Codegen.optional(this.checkBeacon);
    }
    /**
     * If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
     * 
     */
    @Export(name="failback", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> failback;

    /**
     * @return If true, the teaming policy will re-activate failed interfaces higher in precedence when they come back up.
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
    @Export(name="hostSystemId", refs={String.class}, tree="[0]")
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
     * Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
     * 
     */
    @Export(name="linkDiscoveryOperation", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> linkDiscoveryOperation;

    /**
     * @return Whether to advertise or listen for link discovery. Valid values are advertise, both, listen, and none.
     * 
     */
    public Output<Optional<String>> linkDiscoveryOperation() {
        return Codegen.optional(this.linkDiscoveryOperation);
    }
    /**
     * The discovery protocol type. Valid values are cdp and lldp.
     * 
     */
    @Export(name="linkDiscoveryProtocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> linkDiscoveryProtocol;

    /**
     * @return The discovery protocol type. Valid values are cdp and lldp.
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
    @Export(name="mtu", refs={Integer.class}, tree="[0]")
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
    @Export(name="name", refs={String.class}, tree="[0]")
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
     * The list of network adapters to bind to this virtual switch.
     * 
     */
    @Export(name="networkAdapters", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> networkAdapters;

    /**
     * @return The list of network adapters to bind to this virtual switch.
     * 
     */
    public Output<List<String>> networkAdapters() {
        return this.networkAdapters;
    }
    /**
     * If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
     * 
     */
    @Export(name="notifySwitches", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> notifySwitches;

    /**
     * @return If true, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates.
     * 
     */
    public Output<Optional<Boolean>> notifySwitches() {
        return Codegen.optional(this.notifySwitches);
    }
    /**
     * The number of ports to create with this
     * virtual switch. Default: `128`.
     * 
     * &gt; **NOTE:** Changing the port count requires a reboot of the host. This provider
     * will not restart the host for you.
     * 
     */
    @Export(name="numberOfPorts", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> numberOfPorts;

    /**
     * @return The number of ports to create with this
     * virtual switch. Default: `128`.
     * 
     * &gt; **NOTE:** Changing the port count requires a reboot of the host. This provider
     * will not restart the host for you.
     * 
     */
    public Output<Optional<Integer>> numberOfPorts() {
        return Codegen.optional(this.numberOfPorts);
    }
    /**
     * The average bandwidth in bits per second if traffic shaping is enabled.
     * 
     */
    @Export(name="shapingAverageBandwidth", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> shapingAverageBandwidth;

    /**
     * @return The average bandwidth in bits per second if traffic shaping is enabled.
     * 
     */
    public Output<Optional<Integer>> shapingAverageBandwidth() {
        return Codegen.optional(this.shapingAverageBandwidth);
    }
    /**
     * The maximum burst size allowed in bytes if traffic shaping is enabled.
     * 
     */
    @Export(name="shapingBurstSize", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> shapingBurstSize;

    /**
     * @return The maximum burst size allowed in bytes if traffic shaping is enabled.
     * 
     */
    public Output<Optional<Integer>> shapingBurstSize() {
        return Codegen.optional(this.shapingBurstSize);
    }
    /**
     * Enable traffic shaping on this virtual switch or port group.
     * 
     */
    @Export(name="shapingEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> shapingEnabled;

    /**
     * @return Enable traffic shaping on this virtual switch or port group.
     * 
     */
    public Output<Optional<Boolean>> shapingEnabled() {
        return Codegen.optional(this.shapingEnabled);
    }
    /**
     * The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
     * 
     */
    @Export(name="shapingPeakBandwidth", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> shapingPeakBandwidth;

    /**
     * @return The peak bandwidth during bursts in bits per second if traffic shaping is enabled.
     * 
     */
    public Output<Optional<Integer>> shapingPeakBandwidth() {
        return Codegen.optional(this.shapingPeakBandwidth);
    }
    /**
     * List of standby network adapters used for failover.
     * 
     */
    @Export(name="standbyNics", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> standbyNics;

    /**
     * @return List of standby network adapters used for failover.
     * 
     */
    public Output<Optional<List<String>>> standbyNics() {
        return Codegen.optional(this.standbyNics);
    }
    /**
     * The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
     * failover_explicit.
     * 
     */
    @Export(name="teamingPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> teamingPolicy;

    /**
     * @return The network adapter teaming policy. Can be one of loadbalance_ip, loadbalance_srcmac, loadbalance_srcid, or
     * failover_explicit.
     * 
     */
    public Output<Optional<String>> teamingPolicy() {
        return Codegen.optional(this.teamingPolicy);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HostVirtualSwitch(java.lang.String name) {
        this(name, HostVirtualSwitchArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HostVirtualSwitch(java.lang.String name, HostVirtualSwitchArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HostVirtualSwitch(java.lang.String name, HostVirtualSwitchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/hostVirtualSwitch:HostVirtualSwitch", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private HostVirtualSwitch(java.lang.String name, Output<java.lang.String> id, @Nullable HostVirtualSwitchState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/hostVirtualSwitch:HostVirtualSwitch", name, state, makeResourceOptions(options, id), false);
    }

    private static HostVirtualSwitchArgs makeArgs(HostVirtualSwitchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? HostVirtualSwitchArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static HostVirtualSwitch get(java.lang.String name, Output<java.lang.String> id, @Nullable HostVirtualSwitchState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HostVirtualSwitch(name, id, state, options);
    }
}

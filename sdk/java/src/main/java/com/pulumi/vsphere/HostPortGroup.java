// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.HostPortGroupArgs;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.inputs.HostPortGroupState;
import com.pulumi.vsphere.outputs.HostPortGroupPort;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `vsphere.HostPortGroup` resource can be used to manage port groups on
 * ESXi hosts. These port groups are connected to standard switches, which
 * can be managed by the `vsphere.HostVirtualSwitch`
 * resource.
 * 
 * For an overview on vSphere networking concepts, see [the product documentation][ref-vsphere-net-concepts].
 * 
 * [ref-vsphere-net-concepts]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.networking.doc/GUID-2B11DBB8-CB3C-4AFF-8885-EFEA0FC562F4.html
 * 
 * ## Example Usage
 * 
 * **Create a Virtual Switch and Bind a Port Group:**
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vsphere.VsphereFunctions;
 * import com.pulumi.vsphere.inputs.GetDatacenterArgs;
 * import com.pulumi.vsphere.inputs.GetHostArgs;
 * import com.pulumi.vsphere.HostVirtualSwitch;
 * import com.pulumi.vsphere.HostVirtualSwitchArgs;
 * import com.pulumi.vsphere.HostPortGroup;
 * import com.pulumi.vsphere.HostPortGroupArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var datacenter = VsphereFunctions.getDatacenter(GetDatacenterArgs.builder()
 *             .name("dc-01")
 *             .build());
 * 
 *         final var host = VsphereFunctions.getHost(GetHostArgs.builder()
 *             .name("esxi-01.example.com")
 *             .datacenterId(datacenter.applyValue(getDatacenterResult -> getDatacenterResult.id()))
 *             .build());
 * 
 *         var hostVirtualSwitch = new HostVirtualSwitch("hostVirtualSwitch", HostVirtualSwitchArgs.builder()
 *             .name("switch-01")
 *             .hostSystemId(host.applyValue(getHostResult -> getHostResult.id()))
 *             .networkAdapters(            
 *                 "vmnic0",
 *                 "vmnic1")
 *             .activeNics("vmnic0")
 *             .standbyNics("vmnic1")
 *             .build());
 * 
 *         var pg = new HostPortGroup("pg", HostPortGroupArgs.builder()
 *             .name("portgroup-01")
 *             .hostSystemId(host.applyValue(getHostResult -> getHostResult.id()))
 *             .virtualSwitchName(hostVirtualSwitch.name())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * **Create a Port Group with a VLAN and ab Override:**
 * 
 * This example sets the trunk mode VLAN (`4095`, which passes through all tags)
 * and sets
 * `allow_promiscuous`
 * to ensure that all traffic is seen on the port. The setting overrides
 * the implicit default of `false` set on the standard switch.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vsphere.VsphereFunctions;
 * import com.pulumi.vsphere.inputs.GetDatacenterArgs;
 * import com.pulumi.vsphere.inputs.GetHostArgs;
 * import com.pulumi.vsphere.HostVirtualSwitch;
 * import com.pulumi.vsphere.HostVirtualSwitchArgs;
 * import com.pulumi.vsphere.HostPortGroup;
 * import com.pulumi.vsphere.HostPortGroupArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var datacenter = VsphereFunctions.getDatacenter(GetDatacenterArgs.builder()
 *             .name("dc-01")
 *             .build());
 * 
 *         final var host = VsphereFunctions.getHost(GetHostArgs.builder()
 *             .name("esxi-01.example.com")
 *             .datacenterId(datacenter.applyValue(getDatacenterResult -> getDatacenterResult.id()))
 *             .build());
 * 
 *         var hostVirtualSwitch = new HostVirtualSwitch("hostVirtualSwitch", HostVirtualSwitchArgs.builder()
 *             .name("switch-01")
 *             .hostSystemId(host.applyValue(getHostResult -> getHostResult.id()))
 *             .networkAdapters(            
 *                 "vmnic0",
 *                 "vmnic1")
 *             .activeNics("vmnic0")
 *             .standbyNics("vmnic1")
 *             .build());
 * 
 *         var pg = new HostPortGroup("pg", HostPortGroupArgs.builder()
 *             .name("portgroup-01")
 *             .hostSystemId(host.applyValue(getHostResult -> getHostResult.id()))
 *             .virtualSwitchName(hostVirtualSwitch.name())
 *             .vlanId(4095)
 *             .allowPromiscuous(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Importing
 * 
 * An existing host port group can be imported into this resource
 * using the host port group&#39;s ID. An example is below:
 * 
 * The above would import the `management` host port group from host with ID `host-123`.
 * 
 */
@ResourceType(type="vsphere:index/hostPortGroup:HostPortGroup")
public class HostPortGroup extends com.pulumi.resources.CustomResource {
    /**
     * List of active network adapters used for load balancing.
     * 
     */
    @Export(name="activeNics", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> activeNics;

    /**
     * @return List of active network adapters used for load balancing.
     * 
     */
    public Output<Optional<List<String>>> activeNics() {
        return Codegen.optional(this.activeNics);
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
     * A map with a full set of the policy
     * options computed from defaults and overrides,
     * explaining the effective policy for this port group.
     * 
     */
    @Export(name="computedPolicy", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> computedPolicy;

    /**
     * @return A map with a full set of the policy
     * options computed from defaults and overrides,
     * explaining the effective policy for this port group.
     * 
     */
    public Output<Map<String,String>> computedPolicy() {
        return this.computedPolicy;
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
     * the host to set the port group up on. Forces a new resource if changed.
     * 
     */
    @Export(name="hostSystemId", refs={String.class}, tree="[0]")
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
     * The key for this port group as returned from the vSphere API.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return The key for this port group as returned from the vSphere API.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * The name of the port group.  Forces a new resource if
     * changed.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the port group.  Forces a new resource if
     * changed.
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * A list of ports that currently exist and are used on this port group.
     * 
     */
    @Export(name="ports", refs={List.class,HostPortGroupPort.class}, tree="[0,1]")
    private Output<List<HostPortGroupPort>> ports;

    /**
     * @return A list of ports that currently exist and are used on this port group.
     * 
     */
    public Output<List<HostPortGroupPort>> ports() {
        return this.ports;
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
     * The name of the virtual switch to bind
     * this port group to. Forces a new resource if changed.
     * 
     */
    @Export(name="virtualSwitchName", refs={String.class}, tree="[0]")
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
    @Export(name="vlanId", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> vlanId;

    /**
     * @return The VLAN ID/trunk mode for this port group.  An ID of
     * `0` denotes no tagging, an ID of `1`-`4094` tags with the specific ID, and an
     * ID of `4095` enables trunk mode, allowing the guest to manage its own
     * tagging. Default: `0`.
     * 
     */
    public Output<Optional<Integer>> vlanId() {
        return Codegen.optional(this.vlanId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HostPortGroup(java.lang.String name) {
        this(name, HostPortGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HostPortGroup(java.lang.String name, HostPortGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HostPortGroup(java.lang.String name, HostPortGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/hostPortGroup:HostPortGroup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private HostPortGroup(java.lang.String name, Output<java.lang.String> id, @Nullable HostPortGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/hostPortGroup:HostPortGroup", name, state, makeResourceOptions(options, id), false);
    }

    private static HostPortGroupArgs makeArgs(HostPortGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? HostPortGroupArgs.Empty : args;
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
    public static HostPortGroup get(java.lang.String name, Output<java.lang.String> id, @Nullable HostPortGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HostPortGroup(name, id, state, options);
    }
}

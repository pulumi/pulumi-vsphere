// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.VnicArgs;
import com.pulumi.vsphere.inputs.VnicState;
import com.pulumi.vsphere.outputs.VnicIpv4;
import com.pulumi.vsphere.outputs.VnicIpv6;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a VMware vSphere vnic resource.
 * 
 * ## Example Usage
 * 
 * ### Create a vnic attached to a distributed virtual switch using the vmotion TCP/IP stack
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
 * import com.pulumi.vsphere.DistributedVirtualSwitch;
 * import com.pulumi.vsphere.DistributedVirtualSwitchArgs;
 * import com.pulumi.vsphere.inputs.DistributedVirtualSwitchHostArgs;
 * import com.pulumi.vsphere.DistributedPortGroup;
 * import com.pulumi.vsphere.DistributedPortGroupArgs;
 * import com.pulumi.vsphere.Vnic;
 * import com.pulumi.vsphere.VnicArgs;
 * import com.pulumi.vsphere.inputs.VnicIpv4Args;
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
 *             .datacenterId(datacenter.id())
 *             .build());
 * 
 *         var vds = new DistributedVirtualSwitch("vds", DistributedVirtualSwitchArgs.builder()
 *             .name("vds-01")
 *             .datacenterId(datacenter.id())
 *             .hosts(DistributedVirtualSwitchHostArgs.builder()
 *                 .hostSystemId(host.id())
 *                 .devices("vnic3")
 *                 .build())
 *             .build());
 * 
 *         var pg = new DistributedPortGroup("pg", DistributedPortGroupArgs.builder()
 *             .name("pg-01")
 *             .vlanId(1234)
 *             .distributedVirtualSwitchUuid(vds.id())
 *             .build());
 * 
 *         var vnic = new Vnic("vnic", VnicArgs.builder()
 *             .host(host.id())
 *             .distributedSwitchPort(vds.id())
 *             .distributedPortGroup(pg.id())
 *             .ipv4(VnicIpv4Args.builder()
 *                 .dhcp(true)
 *                 .build())
 *             .netstack("vmotion")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Create a vnic attached to a portgroup using the default TCP/IP stack
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
 * import com.pulumi.vsphere.Vnic;
 * import com.pulumi.vsphere.VnicArgs;
 * import com.pulumi.vsphere.inputs.VnicIpv4Args;
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
 *             .datacenterId(datacenter.id())
 *             .build());
 * 
 *         var hvs = new HostVirtualSwitch("hvs", HostVirtualSwitchArgs.builder()
 *             .name("hvs-01")
 *             .hostSystemId(host.id())
 *             .networkAdapters(            
 *                 "vmnic3",
 *                 "vmnic4")
 *             .activeNics("vmnic3")
 *             .standbyNics("vmnic4")
 *             .build());
 * 
 *         var pg = new HostPortGroup("pg", HostPortGroupArgs.builder()
 *             .name("pg-01")
 *             .virtualSwitchName(hvs.name())
 *             .hostSystemId(host.id())
 *             .build());
 * 
 *         var vnic = new Vnic("vnic", VnicArgs.builder()
 *             .host(host.id())
 *             .portgroup(pg.name())
 *             .ipv4(VnicIpv4Args.builder()
 *                 .dhcp(true)
 *                 .build())
 *             .services(            
 *                 "vsan",
 *                 "management")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * An existing vNic can be imported into this resource
 * 
 * via supplying the vNic&#39;s ID. An example is below:
 * 
 * [docs-import]: /docs/import/index.html
 * 
 * ```sh
 * $ pulumi import vsphere:index/vnic:Vnic vnic host-123_vmk2
 * ```
 * 
 * The above would import the vnic `vmk2` from host with ID `host-123`.
 * 
 */
@ResourceType(type="vsphere:index/vnic:Vnic")
public class Vnic extends com.pulumi.resources.CustomResource {
    /**
     * Key of the distributed portgroup the nic will connect to.
     * 
     */
    @Export(name="distributedPortGroup", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> distributedPortGroup;

    /**
     * @return Key of the distributed portgroup the nic will connect to.
     * 
     */
    public Output<Optional<String>> distributedPortGroup() {
        return Codegen.optional(this.distributedPortGroup);
    }
    /**
     * UUID of the vdswitch the nic will be attached to. Do not set if you set portgroup.
     * 
     */
    @Export(name="distributedSwitchPort", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> distributedSwitchPort;

    /**
     * @return UUID of the vdswitch the nic will be attached to. Do not set if you set portgroup.
     * 
     */
    public Output<Optional<String>> distributedSwitchPort() {
        return Codegen.optional(this.distributedSwitchPort);
    }
    /**
     * ESX host the interface belongs to
     * 
     */
    @Export(name="host", refs={String.class}, tree="[0]")
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
    @Export(name="ipv4", refs={VnicIpv4.class}, tree="[0]")
    private Output</* @Nullable */ VnicIpv4> ipv4;

    /**
     * @return IPv4 settings. Either this or `ipv6` needs to be set. See IPv4 options below.
     * 
     */
    public Output<Optional<VnicIpv4>> ipv4() {
        return Codegen.optional(this.ipv4);
    }
    /**
     * IPv6 settings. Either this or `ipv6` needs to be set. See IPv6 options below.
     * 
     */
    @Export(name="ipv6", refs={VnicIpv6.class}, tree="[0]")
    private Output</* @Nullable */ VnicIpv6> ipv6;

    /**
     * @return IPv6 settings. Either this or `ipv6` needs to be set. See IPv6 options below.
     * 
     */
    public Output<Optional<VnicIpv6>> ipv6() {
        return Codegen.optional(this.ipv6);
    }
    /**
     * MAC address of the interface.
     * 
     */
    @Export(name="mac", refs={String.class}, tree="[0]")
    private Output<String> mac;

    /**
     * @return MAC address of the interface.
     * 
     */
    public Output<String> mac() {
        return this.mac;
    }
    /**
     * MTU of the interface.
     * 
     */
    @Export(name="mtu", refs={Integer.class}, tree="[0]")
    private Output<Integer> mtu;

    /**
     * @return MTU of the interface.
     * 
     */
    public Output<Integer> mtu() {
        return this.mtu;
    }
    /**
     * TCP/IP stack setting for this interface. Possible values are `defaultTcpipStack``, &#39;vmotion&#39;, &#39;vSphereProvisioning&#39;. Changing this will force the creation of a new interface since it&#39;s not possible to change the stack once it gets created. (Default:`defaultTcpipStack`)
     * 
     */
    @Export(name="netstack", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> netstack;

    /**
     * @return TCP/IP stack setting for this interface. Possible values are `defaultTcpipStack``, &#39;vmotion&#39;, &#39;vSphereProvisioning&#39;. Changing this will force the creation of a new interface since it&#39;s not possible to change the stack once it gets created. (Default:`defaultTcpipStack`)
     * 
     */
    public Output<Optional<String>> netstack() {
        return Codegen.optional(this.netstack);
    }
    /**
     * Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
     * 
     */
    @Export(name="portgroup", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> portgroup;

    /**
     * @return Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
     * 
     */
    public Output<Optional<String>> portgroup() {
        return Codegen.optional(this.portgroup);
    }
    /**
     * Enabled services setting for this interface. Currently support values are `vmotion`, `management`, and `vsan`.
     * 
     */
    @Export(name="services", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> services;

    /**
     * @return Enabled services setting for this interface. Currently support values are `vmotion`, `management`, and `vsan`.
     * 
     */
    public Output<Optional<List<String>>> services() {
        return Codegen.optional(this.services);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Vnic(java.lang.String name) {
        this(name, VnicArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Vnic(java.lang.String name, VnicArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Vnic(java.lang.String name, VnicArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/vnic:Vnic", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Vnic(java.lang.String name, Output<java.lang.String> id, @Nullable VnicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/vnic:Vnic", name, state, makeResourceOptions(options, id), false);
    }

    private static VnicArgs makeArgs(VnicArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VnicArgs.Empty : args;
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
    public static Vnic get(java.lang.String name, Output<java.lang.String> id, @Nullable VnicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Vnic(name, id, state, options);
    }
}

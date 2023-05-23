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
 * ### S
 * ### Create a vnic attached to a distributed virtual switch using the vmotion TCP/IP stack
 * ```java
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
 *         final var dc = VsphereFunctions.getDatacenter(GetDatacenterArgs.builder()
 *             .name(&#34;mydc&#34;)
 *             .build());
 * 
 *         final var h1 = VsphereFunctions.getHost(GetHostArgs.builder()
 *             .name(&#34;esxi1.host.test&#34;)
 *             .datacenterId(dc.applyValue(getDatacenterResult -&gt; getDatacenterResult.id()))
 *             .build());
 * 
 *         var d1 = new DistributedVirtualSwitch(&#34;d1&#34;, DistributedVirtualSwitchArgs.builder()        
 *             .datacenterId(dc.applyValue(getDatacenterResult -&gt; getDatacenterResult.id()))
 *             .hosts(DistributedVirtualSwitchHostArgs.builder()
 *                 .hostSystemId(h1.applyValue(getHostResult -&gt; getHostResult.id()))
 *                 .devices(&#34;vnic3&#34;)
 *                 .build())
 *             .build());
 * 
 *         var p1 = new DistributedPortGroup(&#34;p1&#34;, DistributedPortGroupArgs.builder()        
 *             .vlanId(1234)
 *             .distributedVirtualSwitchUuid(d1.id())
 *             .build());
 * 
 *         var v1 = new Vnic(&#34;v1&#34;, VnicArgs.builder()        
 *             .host(h1.applyValue(getHostResult -&gt; getHostResult.id()))
 *             .distributedSwitchPort(d1.id())
 *             .distributedPortGroup(p1.id())
 *             .ipv4(VnicIpv4Args.builder()
 *                 .dhcp(true)
 *                 .build())
 *             .netstack(&#34;vmotion&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Create a vnic attached to a portgroup using the default TCP/IP stack
 * 
 * ```java
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
 *         final var dc = VsphereFunctions.getDatacenter(GetDatacenterArgs.builder()
 *             .name(&#34;mydc&#34;)
 *             .build());
 * 
 *         final var h1 = VsphereFunctions.getHost(GetHostArgs.builder()
 *             .name(&#34;esxi1.host.test&#34;)
 *             .datacenterId(dc.applyValue(getDatacenterResult -&gt; getDatacenterResult.id()))
 *             .build());
 * 
 *         var hvs1 = new HostVirtualSwitch(&#34;hvs1&#34;, HostVirtualSwitchArgs.builder()        
 *             .hostSystemId(h1.applyValue(getHostResult -&gt; getHostResult.id()))
 *             .networkAdapters(            
 *                 &#34;vmnic3&#34;,
 *                 &#34;vmnic4&#34;)
 *             .activeNics(&#34;vmnic3&#34;)
 *             .standbyNics(&#34;vmnic4&#34;)
 *             .build());
 * 
 *         var p1 = new HostPortGroup(&#34;p1&#34;, HostPortGroupArgs.builder()        
 *             .virtualSwitchName(hvs1.name())
 *             .hostSystemId(h1.applyValue(getHostResult -&gt; getHostResult.id()))
 *             .build());
 * 
 *         var v1 = new Vnic(&#34;v1&#34;, VnicArgs.builder()        
 *             .host(h1.applyValue(getHostResult -&gt; getHostResult.id()))
 *             .portgroup(p1.name())
 *             .ipv4(VnicIpv4Args.builder()
 *                 .dhcp(true)
 *                 .build())
 *             .enabledServices(            
 *                 &#34;vsan&#34;,
 *                 &#34;management&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Importing
 * 
 * An existing vNic can be [imported][docs-import] into this resource
 * via supplying the vNic&#39;s ID. An example is below:
 * 
 * [docs-import]: /docs/import/index.html
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
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
 *     }
 * }
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
    @Export(name="distributedPortGroup", type=String.class, parameters={})
    private Output</* @Nullable */ String> distributedPortGroup;

    /**
     * @return Key of the distributed portgroup the nic will connect to.
     * 
     */
    public Output<Optional<String>> distributedPortGroup() {
        return Codegen.optional(this.distributedPortGroup);
    }
    /**
     * UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
     * 
     */
    @Export(name="distributedSwitchPort", type=String.class, parameters={})
    private Output</* @Nullable */ String> distributedSwitchPort;

    /**
     * @return UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.
     * 
     */
    public Output<Optional<String>> distributedSwitchPort() {
        return Codegen.optional(this.distributedSwitchPort);
    }
    /**
     * ESX host the interface belongs to
     * 
     */
    @Export(name="host", type=String.class, parameters={})
    private Output<String> host;

    /**
     * @return ESX host the interface belongs to
     * 
     */
    public Output<String> host() {
        return this.host;
    }
    /**
     * IPv4 settings. Either this or `ipv6` needs to be set. See  ipv4 options below.
     * 
     */
    @Export(name="ipv4", type=VnicIpv4.class, parameters={})
    private Output</* @Nullable */ VnicIpv4> ipv4;

    /**
     * @return IPv4 settings. Either this or `ipv6` needs to be set. See  ipv4 options below.
     * 
     */
    public Output<Optional<VnicIpv4>> ipv4() {
        return Codegen.optional(this.ipv4);
    }
    /**
     * IPv6 settings. Either this or `ipv6` needs to be set. See  ipv6 options below.
     * 
     */
    @Export(name="ipv6", type=VnicIpv6.class, parameters={})
    private Output</* @Nullable */ VnicIpv6> ipv6;

    /**
     * @return IPv6 settings. Either this or `ipv6` needs to be set. See  ipv6 options below.
     * 
     */
    public Output<Optional<VnicIpv6>> ipv6() {
        return Codegen.optional(this.ipv6);
    }
    /**
     * MAC address of the interface.
     * 
     */
    @Export(name="mac", type=String.class, parameters={})
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
    @Export(name="mtu", type=Integer.class, parameters={})
    private Output<Integer> mtu;

    /**
     * @return MTU of the interface.
     * 
     */
    public Output<Integer> mtu() {
        return this.mtu;
    }
    /**
     * TCP/IP stack setting for this interface. Possible values are &#39;defaultTcpipStack&#39;, &#39;vmotion&#39;, &#39;vSphereProvisioning&#39;. Changing this will force the creation of a new interface since it&#39;s not possible to change the stack once it gets created. (Default: `defaultTcpipStack`)
     * 
     */
    @Export(name="netstack", type=String.class, parameters={})
    private Output</* @Nullable */ String> netstack;

    /**
     * @return TCP/IP stack setting for this interface. Possible values are &#39;defaultTcpipStack&#39;, &#39;vmotion&#39;, &#39;vSphereProvisioning&#39;. Changing this will force the creation of a new interface since it&#39;s not possible to change the stack once it gets created. (Default: `defaultTcpipStack`)
     * 
     */
    public Output<Optional<String>> netstack() {
        return Codegen.optional(this.netstack);
    }
    /**
     * Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
     * 
     */
    @Export(name="portgroup", type=String.class, parameters={})
    private Output</* @Nullable */ String> portgroup;

    /**
     * @return Portgroup to attach the nic to. Do not set if you set distributed_switch_port.
     * 
     */
    public Output<Optional<String>> portgroup() {
        return Codegen.optional(this.portgroup);
    }
    /**
     * Enabled services setting for this interface. Current possible values are &#39;vmotion&#39;, &#39;management&#39;, and &#39;vsan&#39;.
     * 
     */
    @Export(name="services", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> services;

    /**
     * @return Enabled services setting for this interface. Current possible values are &#39;vmotion&#39;, &#39;management&#39;, and &#39;vsan&#39;.
     * 
     */
    public Output<Optional<List<String>>> services() {
        return Codegen.optional(this.services);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Vnic(String name) {
        this(name, VnicArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Vnic(String name, VnicArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Vnic(String name, VnicArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/vnic:Vnic", name, args == null ? VnicArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Vnic(String name, Output<String> id, @Nullable VnicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/vnic:Vnic", name, state, makeResourceOptions(options, id));
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
    public static Vnic get(String name, Output<String> id, @Nullable VnicState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Vnic(name, id, state, options);
    }
}

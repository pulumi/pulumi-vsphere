// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.VappEntityArgs;
import com.pulumi.vsphere.inputs.VappEntityState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `vsphere.VappEntity` resource can be used to describe the behavior of an
 * entity (virtual machine or sub-vApp container) in a vApp container.
 * 
 * For more information on vSphere vApps, see [this
 * page][ref-vsphere-vapp].
 * 
 * [ref-vsphere-vapp]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/create-a-vapp-h5-and-flex.html
 * 
 * ## Example Usage
 * 
 * The basic example below sets up a vApp container and a virtual machine in a
 * compute cluster and then creates a vApp entity to change the virtual machine&#39;s
 * power on behavior in the vApp container.
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
 * import com.pulumi.vsphere.inputs.GetComputeClusterArgs;
 * import com.pulumi.vsphere.inputs.GetNetworkArgs;
 * import com.pulumi.vsphere.inputs.GetDatastoreArgs;
 * import com.pulumi.vsphere.VappContainer;
 * import com.pulumi.vsphere.VappContainerArgs;
 * import com.pulumi.vsphere.VirtualMachine;
 * import com.pulumi.vsphere.VirtualMachineArgs;
 * import com.pulumi.vsphere.inputs.VirtualMachineDiskArgs;
 * import com.pulumi.vsphere.inputs.VirtualMachineNetworkInterfaceArgs;
 * import com.pulumi.vsphere.VappEntity;
 * import com.pulumi.vsphere.VappEntityArgs;
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
 *         final var config = ctx.config();
 *         final var datacenter = config.get("datacenter").orElse("dc-01");
 *         final var cluster = config.get("cluster").orElse("cluster-01");
 *         final var datacenterGetDatacenter = VsphereFunctions.getDatacenter(GetDatacenterArgs.builder()
 *             .name(datacenter)
 *             .build());
 * 
 *         final var computeCluster = VsphereFunctions.getComputeCluster(GetComputeClusterArgs.builder()
 *             .name(cluster)
 *             .datacenterId(datacenterGetDatacenter.applyValue(getDatacenterResult -> getDatacenterResult.id()))
 *             .build());
 * 
 *         final var network = VsphereFunctions.getNetwork(GetNetworkArgs.builder()
 *             .name("network1")
 *             .datacenterId(datacenterGetDatacenter.applyValue(getDatacenterResult -> getDatacenterResult.id()))
 *             .build());
 * 
 *         final var datastore = VsphereFunctions.getDatastore(GetDatastoreArgs.builder()
 *             .name("datastore1")
 *             .datacenterId(datacenterGetDatacenter.applyValue(getDatacenterResult -> getDatacenterResult.id()))
 *             .build());
 * 
 *         var vappContainer = new VappContainer("vappContainer", VappContainerArgs.builder()
 *             .name("vapp-container-test")
 *             .parentResourcePoolId(computeCluster.applyValue(getComputeClusterResult -> getComputeClusterResult.id()))
 *             .build());
 * 
 *         var vm = new VirtualMachine("vm", VirtualMachineArgs.builder()
 *             .name("virtual-machine-test")
 *             .resourcePoolId(vappContainer.id())
 *             .datastoreId(datastore.applyValue(getDatastoreResult -> getDatastoreResult.id()))
 *             .numCpus(2)
 *             .memory(1024)
 *             .guestId("ubuntu64Guest")
 *             .disks(VirtualMachineDiskArgs.builder()
 *                 .label("disk0")
 *                 .size(1)
 *                 .build())
 *             .networkInterfaces(VirtualMachineNetworkInterfaceArgs.builder()
 *                 .networkId(network.applyValue(getNetworkResult -> getNetworkResult.id()))
 *                 .build())
 *             .build());
 * 
 *         var vappEntity = new VappEntity("vappEntity", VappEntityArgs.builder()
 *             .targetId(vm.moid())
 *             .containerId(vappContainer.id())
 *             .startAction("none")
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
 * An existing vApp entity can be imported into this resource via
 * 
 * the ID of the vApp Entity.
 * 
 * ```sh
 * $ pulumi import vsphere:index/vappEntity:VappEntity vapp_entity vm-123:res-456
 * ```
 * 
 * The above would import the vApp entity that governs the behavior of the virtual
 * 
 * machine with a [managed object ID][docs-about-morefs] of vm-123 in the vApp
 * 
 * container with the [managed object ID][docs-about-morefs] res-456.
 * 
 */
@ResourceType(type="vsphere:index/vappEntity:VappEntity")
public class VappEntity extends com.pulumi.resources.CustomResource {
    /**
     * Managed object ID of the vApp
     * container the entity is a member of.
     * 
     */
    @Export(name="containerId", refs={String.class}, tree="[0]")
    private Output<String> containerId;

    /**
     * @return Managed object ID of the vApp
     * container the entity is a member of.
     * 
     */
    public Output<String> containerId() {
        return this.containerId;
    }
    /**
     * A list of custom attributes to set on this resource.
     * 
     */
    @Export(name="customAttributes", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> customAttributes;

    /**
     * @return A list of custom attributes to set on this resource.
     * 
     */
    public Output<Optional<Map<String,String>>> customAttributes() {
        return Codegen.optional(this.customAttributes);
    }
    /**
     * How to start the entity. Valid settings are none
     * or powerOn. If set to none, then the entity does not participate in auto-start.
     * Default: powerOn
     * 
     */
    @Export(name="startAction", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> startAction;

    /**
     * @return How to start the entity. Valid settings are none
     * or powerOn. If set to none, then the entity does not participate in auto-start.
     * Default: powerOn
     * 
     */
    public Output<Optional<String>> startAction() {
        return Codegen.optional(this.startAction);
    }
    /**
     * Delay in seconds before continuing with the next
     * entity in the order of entities to be started. Default: 120
     * 
     */
    @Export(name="startDelay", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> startDelay;

    /**
     * @return Delay in seconds before continuing with the next
     * entity in the order of entities to be started. Default: 120
     * 
     */
    public Output<Optional<Integer>> startDelay() {
        return Codegen.optional(this.startDelay);
    }
    /**
     * Order to start and stop target in vApp. Default: 1
     * 
     */
    @Export(name="startOrder", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> startOrder;

    /**
     * @return Order to start and stop target in vApp. Default: 1
     * 
     */
    public Output<Optional<Integer>> startOrder() {
        return Codegen.optional(this.startOrder);
    }
    /**
     * Defines the stop action for the entity. Can be set
     * to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
     * does not participate in auto-stop. Default: powerOff
     * 
     */
    @Export(name="stopAction", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stopAction;

    /**
     * @return Defines the stop action for the entity. Can be set
     * to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
     * does not participate in auto-stop. Default: powerOff
     * 
     */
    public Output<Optional<String>> stopAction() {
        return Codegen.optional(this.stopAction);
    }
    /**
     * Delay in seconds before continuing with the next
     * entity in the order sequence. This is only used if the stopAction is
     * guestShutdown. Default: 120
     * 
     */
    @Export(name="stopDelay", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> stopDelay;

    /**
     * @return Delay in seconds before continuing with the next
     * entity in the order sequence. This is only used if the stopAction is
     * guestShutdown. Default: 120
     * 
     */
    public Output<Optional<Integer>> stopDelay() {
        return Codegen.optional(this.stopDelay);
    }
    /**
     * A list of tag IDs to apply to this object.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A list of tag IDs to apply to this object.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Managed object ID of the entity
     * to power on or power off. This can be a virtual machine or a vApp.
     * 
     */
    @Export(name="targetId", refs={String.class}, tree="[0]")
    private Output<String> targetId;

    /**
     * @return Managed object ID of the entity
     * to power on or power off. This can be a virtual machine or a vApp.
     * 
     */
    public Output<String> targetId() {
        return this.targetId;
    }
    /**
     * Determines if the VM should be marked as being
     * started when VMware Tools are ready instead of waiting for `start_delay`. This
     * property has no effect for vApps. Default: false
     * 
     */
    @Export(name="waitForGuest", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> waitForGuest;

    /**
     * @return Determines if the VM should be marked as being
     * started when VMware Tools are ready instead of waiting for `start_delay`. This
     * property has no effect for vApps. Default: false
     * 
     */
    public Output<Optional<Boolean>> waitForGuest() {
        return Codegen.optional(this.waitForGuest);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VappEntity(java.lang.String name) {
        this(name, VappEntityArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VappEntity(java.lang.String name, VappEntityArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VappEntity(java.lang.String name, VappEntityArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/vappEntity:VappEntity", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private VappEntity(java.lang.String name, Output<java.lang.String> id, @Nullable VappEntityState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/vappEntity:VappEntity", name, state, makeResourceOptions(options, id), false);
    }

    private static VappEntityArgs makeArgs(VappEntityArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VappEntityArgs.Empty : args;
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
    public static VappEntity get(java.lang.String name, Output<java.lang.String> id, @Nullable VappEntityState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VappEntity(name, id, state, options);
    }
}

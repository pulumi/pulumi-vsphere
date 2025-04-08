// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.HaVmOverrideArgs;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.inputs.HaVmOverrideState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `vsphere.HaVmOverride` resource can be used to add an override for
 * vSphere HA settings on a cluster for a specific virtual machine. With this
 * resource, one can control specific HA settings so that they are different than
 * the cluster default, accommodating the needs of that specific virtual machine,
 * while not affecting the rest of the cluster.
 * 
 * For more information on vSphere HA, see [this page][ref-vsphere-ha-clusters].
 * 
 * [ref-vsphere-ha-clusters]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-availability.html
 * 
 * &gt; **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 * 
 * ## Example Usage
 * 
 * The example below creates a virtual machine in a cluster using the
 * `vsphere.VirtualMachine` resource, creating the
 * virtual machine in the cluster looked up by the
 * `vsphere.ComputeCluster` data source.
 * 
 * Considering a scenario where this virtual machine is of high value to the
 * application or organization for which it does its work, it&#39;s been determined in
 * the event of a host failure, that this should be one of the first virtual
 * machines to be started by vSphere HA during recovery. Hence, it
 * `ha_vm_restart_priority` has been set to `highest`,
 * which, assuming that the default restart priority is `medium` and no other
 * virtual machine has been assigned the `highest` priority, will mean that this
 * VM will be started before any other virtual machine in the event of host
 * failure.
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
 * import com.pulumi.vsphere.inputs.GetDatastoreArgs;
 * import com.pulumi.vsphere.inputs.GetComputeClusterArgs;
 * import com.pulumi.vsphere.inputs.GetNetworkArgs;
 * import com.pulumi.vsphere.VirtualMachine;
 * import com.pulumi.vsphere.VirtualMachineArgs;
 * import com.pulumi.vsphere.inputs.VirtualMachineNetworkInterfaceArgs;
 * import com.pulumi.vsphere.inputs.VirtualMachineDiskArgs;
 * import com.pulumi.vsphere.HaVmOverride;
 * import com.pulumi.vsphere.HaVmOverrideArgs;
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
 *         final var datastore = VsphereFunctions.getDatastore(GetDatastoreArgs.builder()
 *             .name("datastore1")
 *             .datacenterId(datacenter.id())
 *             .build());
 * 
 *         final var cluster = VsphereFunctions.getComputeCluster(GetComputeClusterArgs.builder()
 *             .name("cluster-01")
 *             .datacenterId(datacenter.id())
 *             .build());
 * 
 *         final var network = VsphereFunctions.getNetwork(GetNetworkArgs.builder()
 *             .name("network1")
 *             .datacenterId(datacenter.id())
 *             .build());
 * 
 *         var vm = new VirtualMachine("vm", VirtualMachineArgs.builder()
 *             .name("test")
 *             .resourcePoolId(cluster.resourcePoolId())
 *             .datastoreId(datastore.id())
 *             .numCpus(2)
 *             .memory(2048)
 *             .guestId("otherLinux64Guest")
 *             .networkInterfaces(VirtualMachineNetworkInterfaceArgs.builder()
 *                 .networkId(network.id())
 *                 .build())
 *             .disks(VirtualMachineDiskArgs.builder()
 *                 .label("disk0")
 *                 .size(20)
 *                 .build())
 *             .build());
 * 
 *         var haVmOverride = new HaVmOverride("haVmOverride", HaVmOverrideArgs.builder()
 *             .computeClusterId(cluster.id())
 *             .virtualMachineId(vm.id())
 *             .haVmRestartPriority("highest")
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
 * An existing override can be imported into this resource by
 * 
 * supplying both the path to the cluster, and the path to the virtual machine, to
 * 
 * `pulumi import`. If no override exists, an error will be given.  An example
 * 
 * is below:
 * 
 * ```sh
 * $ pulumi import vsphere:index/haVmOverride:HaVmOverride ha_vm_override \
 * ```
 * 
 *   &#39;{&#34;compute_cluster_path&#34;: &#34;/dc1/host/cluster1&#34;, \
 * 
 *   &#34;virtual_machine_path&#34;: &#34;/dc1/vm/srv1&#34;}&#39;
 * 
 */
@ResourceType(type="vsphere:index/haVmOverride:HaVmOverride")
public class HaVmOverride extends com.pulumi.resources.CustomResource {
    /**
     * The managed object ID of the cluster.
     * 
     */
    @Export(name="computeClusterId", refs={String.class}, tree="[0]")
    private Output<String> computeClusterId;

    /**
     * @return The managed object ID of the cluster.
     * 
     */
    public Output<String> computeClusterId() {
        return this.computeClusterId;
    }
    /**
     * Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an
     * APD event. Can be one of useClusterDefault, none or reset.
     * 
     */
    @Export(name="haDatastoreApdRecoveryAction", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> haDatastoreApdRecoveryAction;

    /**
     * @return Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an
     * APD event. Can be one of useClusterDefault, none or reset.
     * 
     */
    public Output<Optional<String>> haDatastoreApdRecoveryAction() {
        return Codegen.optional(this.haDatastoreApdRecoveryAction);
    }
    /**
     * Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant
     * datastore. Can be one of clusterDefault, disabled, warning, restartConservative, or restartAggressive.
     * 
     */
    @Export(name="haDatastoreApdResponse", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> haDatastoreApdResponse;

    /**
     * @return Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant
     * datastore. Can be one of clusterDefault, disabled, warning, restartConservative, or restartAggressive.
     * 
     */
    public Output<Optional<String>> haDatastoreApdResponse() {
        return Codegen.optional(this.haDatastoreApdResponse);
    }
    /**
     * Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in
     * ha_datastore_apd_response. Specify -1 to use the cluster setting.
     * 
     */
    @Export(name="haDatastoreApdResponseDelay", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> haDatastoreApdResponseDelay;

    /**
     * @return Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in
     * ha_datastore_apd_response. Specify -1 to use the cluster setting.
     * 
     */
    public Output<Optional<Integer>> haDatastoreApdResponseDelay() {
        return Codegen.optional(this.haDatastoreApdResponseDelay);
    }
    /**
     * Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant
     * datastore. Can be one of clusterDefault, disabled, warning, or restartAggressive.
     * 
     */
    @Export(name="haDatastorePdlResponse", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> haDatastorePdlResponse;

    /**
     * @return Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant
     * datastore. Can be one of clusterDefault, disabled, warning, or restartAggressive.
     * 
     */
    public Output<Optional<String>> haDatastorePdlResponse() {
        return Codegen.optional(this.haDatastorePdlResponse);
    }
    /**
     * The action to take on this virtual machine when a host is isolated from the rest of the cluster. Can be one of
     * clusterIsolationResponse, none, powerOff, or shutdown.
     * 
     */
    @Export(name="haHostIsolationResponse", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> haHostIsolationResponse;

    /**
     * @return The action to take on this virtual machine when a host is isolated from the rest of the cluster. Can be one of
     * clusterIsolationResponse, none, powerOff, or shutdown.
     * 
     */
    public Output<Optional<String>> haHostIsolationResponse() {
        return Codegen.optional(this.haHostIsolationResponse);
    }
    /**
     * If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked
     * as failed. The value is in seconds.
     * 
     */
    @Export(name="haVmFailureInterval", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> haVmFailureInterval;

    /**
     * @return If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked
     * as failed. The value is in seconds.
     * 
     */
    public Output<Optional<Integer>> haVmFailureInterval() {
        return Codegen.optional(this.haVmFailureInterval);
    }
    /**
     * The length of the reset window in which ha_vm_maximum_resets can operate. When this window expires, no more resets are
     * attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset
     * time is allotted.
     * 
     */
    @Export(name="haVmMaximumFailureWindow", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> haVmMaximumFailureWindow;

    /**
     * @return The length of the reset window in which ha_vm_maximum_resets can operate. When this window expires, no more resets are
     * attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset
     * time is allotted.
     * 
     */
    public Output<Optional<Integer>> haVmMaximumFailureWindow() {
        return Codegen.optional(this.haVmMaximumFailureWindow);
    }
    /**
     * The maximum number of resets that HA will perform to this virtual machine when responding to a failure event.
     * 
     */
    @Export(name="haVmMaximumResets", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> haVmMaximumResets;

    /**
     * @return The maximum number of resets that HA will perform to this virtual machine when responding to a failure event.
     * 
     */
    public Output<Optional<Integer>> haVmMaximumResets() {
        return Codegen.optional(this.haVmMaximumResets);
    }
    /**
     * The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats.
     * 
     */
    @Export(name="haVmMinimumUptime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> haVmMinimumUptime;

    /**
     * @return The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats.
     * 
     */
    public Output<Optional<Integer>> haVmMinimumUptime() {
        return Codegen.optional(this.haVmMinimumUptime);
    }
    /**
     * The type of virtual machine monitoring to use for this virtual machine. Can be one of vmMonitoringDisabled,
     * vmMonitoringOnly, or vmAndAppMonitoring.
     * 
     */
    @Export(name="haVmMonitoring", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> haVmMonitoring;

    /**
     * @return The type of virtual machine monitoring to use for this virtual machine. Can be one of vmMonitoringDisabled,
     * vmMonitoringOnly, or vmAndAppMonitoring.
     * 
     */
    public Output<Optional<String>> haVmMonitoring() {
        return Codegen.optional(this.haVmMonitoring);
    }
    /**
     * Determines whether or not the cluster&#39;s default settings or the VM override settings specified in this resource are used
     * for virtual machine monitoring. The default is true (use cluster defaults) - set to false to have overrides take effect.
     * 
     */
    @Export(name="haVmMonitoringUseClusterDefaults", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> haVmMonitoringUseClusterDefaults;

    /**
     * @return Determines whether or not the cluster&#39;s default settings or the VM override settings specified in this resource are used
     * for virtual machine monitoring. The default is true (use cluster defaults) - set to false to have overrides take effect.
     * 
     */
    public Output<Optional<Boolean>> haVmMonitoringUseClusterDefaults() {
        return Codegen.optional(this.haVmMonitoringUseClusterDefaults);
    }
    /**
     * The restart priority for this virtual machine when vSphere detects a host failure. Can be one of clusterRestartPriority,
     * lowest, low, medium, high, or highest.
     * 
     */
    @Export(name="haVmRestartPriority", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> haVmRestartPriority;

    /**
     * @return The restart priority for this virtual machine when vSphere detects a host failure. Can be one of clusterRestartPriority,
     * lowest, low, medium, high, or highest.
     * 
     */
    public Output<Optional<String>> haVmRestartPriority() {
        return Codegen.optional(this.haVmRestartPriority);
    }
    /**
     * The maximum time, in seconds, that vSphere HA will wait for the virtual machine to be ready. Use -1 to use the cluster
     * default.
     * 
     */
    @Export(name="haVmRestartTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> haVmRestartTimeout;

    /**
     * @return The maximum time, in seconds, that vSphere HA will wait for the virtual machine to be ready. Use -1 to use the cluster
     * default.
     * 
     */
    public Output<Optional<Integer>> haVmRestartTimeout() {
        return Codegen.optional(this.haVmRestartTimeout);
    }
    /**
     * The managed object ID of the virtual machine.
     * 
     */
    @Export(name="virtualMachineId", refs={String.class}, tree="[0]")
    private Output<String> virtualMachineId;

    /**
     * @return The managed object ID of the virtual machine.
     * 
     */
    public Output<String> virtualMachineId() {
        return this.virtualMachineId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HaVmOverride(java.lang.String name) {
        this(name, HaVmOverrideArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HaVmOverride(java.lang.String name, HaVmOverrideArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HaVmOverride(java.lang.String name, HaVmOverrideArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/haVmOverride:HaVmOverride", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private HaVmOverride(java.lang.String name, Output<java.lang.String> id, @Nullable HaVmOverrideState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/haVmOverride:HaVmOverride", name, state, makeResourceOptions(options, id), false);
    }

    private static HaVmOverrideArgs makeArgs(HaVmOverrideArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? HaVmOverrideArgs.Empty : args;
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
    public static HaVmOverride get(java.lang.String name, Output<java.lang.String> id, @Nullable HaVmOverrideState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HaVmOverride(name, id, state, options);
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.ResourcePoolArgs;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.inputs.ResourcePoolState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `vsphere.ResourcePool` resource can be used to create and manage
 * resource pools on DRS-enabled vSphere clusters or standalone ESXi hosts.
 * 
 * For more information on vSphere resource pools, please refer to the
 * [product documentation][ref-vsphere-resource_pools].
 * 
 * [ref-vsphere-resource_pools]: https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-resource-management/GUID-60077B40-66FF-4625-934A-641703ED7601.html
 * 
 * ## Example Usage
 * 
 * The following example sets up a resource pool in an existing compute cluster
 * with the default settings for CPU and memory reservations, shares, and limits.
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
 * import com.pulumi.vsphere.ResourcePool;
 * import com.pulumi.vsphere.ResourcePoolArgs;
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
 *         final var computeCluster = VsphereFunctions.getComputeCluster(GetComputeClusterArgs.builder()
 *             .name("cluster-01")
 *             .datacenterId(datacenter.applyValue(getDatacenterResult -> getDatacenterResult.id()))
 *             .build());
 * 
 *         var resourcePool = new ResourcePool("resourcePool", ResourcePoolArgs.builder()
 *             .name("resource-pool-01")
 *             .parentResourcePoolId(computeCluster.applyValue(getComputeClusterResult -> getComputeClusterResult.resourcePoolId()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * A virtual machine resource could be targeted to use the default resource pool
 * of the cluster using the following:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vsphere.VirtualMachine;
 * import com.pulumi.vsphere.VirtualMachineArgs;
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
 *         var vm = new VirtualMachine("vm", VirtualMachineArgs.builder()
 *             .resourcePoolId(cluster.resourcePoolId())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * The following example sets up a parent resource pool in an existing compute cluster
 * with a child resource pool nested below. Each resource pool is configured with
 * the default settings for CPU and memory reservations, shares, and limits.
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
 * import com.pulumi.vsphere.ResourcePool;
 * import com.pulumi.vsphere.ResourcePoolArgs;
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
 *         final var computeCluster = VsphereFunctions.getComputeCluster(GetComputeClusterArgs.builder()
 *             .name("cluster-01")
 *             .datacenterId(datacenter.applyValue(getDatacenterResult -> getDatacenterResult.id()))
 *             .build());
 * 
 *         var resourcePoolParent = new ResourcePool("resourcePoolParent", ResourcePoolArgs.builder()
 *             .name("parent")
 *             .parentResourcePoolId(computeCluster.applyValue(getComputeClusterResult -> getComputeClusterResult.resourcePoolId()))
 *             .build());
 * 
 *         var resourcePoolChild = new ResourcePool("resourcePoolChild", ResourcePoolArgs.builder()
 *             .name("child")
 *             .parentResourcePoolId(resourcePoolParent.id())
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
 * An existing resource pool can be imported into this resource via
 * the path to the resource pool, using the following command:
 * 
 * The above would import the resource pool named `resource-pool-01` that is located
 * in the compute cluster `cluster-01` in the `dc-01` datacenter.
 * 
 * ### Settings that Require vSphere 7.0 or higher
 * 
 * These settings require vSphere 7.0 or higher:
 * 
 * * `scale_descendants_shares`
 * 
 */
@ResourceType(type="vsphere:index/resourcePool:ResourcePool")
public class ResourcePool extends com.pulumi.resources.CustomResource {
    /**
     * Determines if the reservation on a resource
     * pool can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     * 
     */
    @Export(name="cpuExpandable", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> cpuExpandable;

    /**
     * @return Determines if the reservation on a resource
     * pool can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     * 
     */
    public Output<Optional<Boolean>> cpuExpandable() {
        return Codegen.optional(this.cpuExpandable);
    }
    /**
     * The CPU utilization of a resource pool will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited. Default: `-1`
     * 
     */
    @Export(name="cpuLimit", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> cpuLimit;

    /**
     * @return The CPU utilization of a resource pool will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited. Default: `-1`
     * 
     */
    public Output<Optional<Integer>> cpuLimit() {
        return Codegen.optional(this.cpuLimit);
    }
    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the resource pool. Default: `0`
     * 
     */
    @Export(name="cpuReservation", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> cpuReservation;

    /**
     * @return Amount of CPU (MHz) that is guaranteed
     * available to the resource pool. Default: `0`
     * 
     */
    public Output<Optional<Integer>> cpuReservation() {
        return Codegen.optional(this.cpuReservation);
    }
    /**
     * The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `cpu_shares` will be
     * ignored.  Default: `normal`
     * 
     */
    @Export(name="cpuShareLevel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cpuShareLevel;

    /**
     * @return The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `cpu_shares` will be
     * ignored.  Default: `normal`
     * 
     */
    public Output<Optional<String>> cpuShareLevel() {
        return Codegen.optional(this.cpuShareLevel);
    }
    /**
     * The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `cpu_share_level` must be `custom`.
     * 
     */
    @Export(name="cpuShares", refs={Integer.class}, tree="[0]")
    private Output<Integer> cpuShares;

    /**
     * @return The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `cpu_share_level` must be `custom`.
     * 
     */
    public Output<Integer> cpuShares() {
        return this.cpuShares;
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
     * Determines if the reservation on a resource
     * pool can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     * 
     */
    @Export(name="memoryExpandable", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> memoryExpandable;

    /**
     * @return Determines if the reservation on a resource
     * pool can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     * 
     */
    public Output<Optional<Boolean>> memoryExpandable() {
        return Codegen.optional(this.memoryExpandable);
    }
    /**
     * The CPU utilization of a resource pool will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited. Default: `-1`
     * 
     */
    @Export(name="memoryLimit", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> memoryLimit;

    /**
     * @return The CPU utilization of a resource pool will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited. Default: `-1`
     * 
     */
    public Output<Optional<Integer>> memoryLimit() {
        return Codegen.optional(this.memoryLimit);
    }
    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the resource pool. Default: `0`
     * 
     */
    @Export(name="memoryReservation", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> memoryReservation;

    /**
     * @return Amount of CPU (MHz) that is guaranteed
     * available to the resource pool. Default: `0`
     * 
     */
    public Output<Optional<Integer>> memoryReservation() {
        return Codegen.optional(this.memoryReservation);
    }
    /**
     * The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `memory_shares` will be
     * ignored.  Default: `normal`
     * 
     */
    @Export(name="memoryShareLevel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> memoryShareLevel;

    /**
     * @return The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `memory_shares` will be
     * ignored.  Default: `normal`
     * 
     */
    public Output<Optional<String>> memoryShareLevel() {
        return Codegen.optional(this.memoryShareLevel);
    }
    /**
     * The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `memory_share_level` must be `custom`.
     * 
     */
    @Export(name="memoryShares", refs={Integer.class}, tree="[0]")
    private Output<Integer> memoryShares;

    /**
     * @return The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `memory_share_level` must be `custom`.
     * 
     */
    public Output<Integer> memoryShares() {
        return this.memoryShares;
    }
    /**
     * The name of the resource pool.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the resource pool.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The managed object ID
     * of the parent resource pool. This can be the root resource pool for a cluster
     * or standalone host, or a resource pool itself. When moving a resource pool
     * from one parent resource pool to another, both must share a common root
     * resource pool.
     * 
     */
    @Export(name="parentResourcePoolId", refs={String.class}, tree="[0]")
    private Output<String> parentResourcePoolId;

    /**
     * @return The managed object ID
     * of the parent resource pool. This can be the root resource pool for a cluster
     * or standalone host, or a resource pool itself. When moving a resource pool
     * from one parent resource pool to another, both must share a common root
     * resource pool.
     * 
     */
    public Output<String> parentResourcePoolId() {
        return this.parentResourcePoolId;
    }
    /**
     * Determines if the shares of all
     * descendants of the resource pool are scaled up or down when the shares
     * of the resource pool are scaled up or down. Can be one of `disabled` or
     * `scaleCpuAndMemoryShares`. Default: `disabled`.
     * 
     */
    @Export(name="scaleDescendantsShares", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> scaleDescendantsShares;

    /**
     * @return Determines if the shares of all
     * descendants of the resource pool are scaled up or down when the shares
     * of the resource pool are scaled up or down. Can be one of `disabled` or
     * `scaleCpuAndMemoryShares`. Default: `disabled`.
     * 
     */
    public Output<Optional<String>> scaleDescendantsShares() {
        return Codegen.optional(this.scaleDescendantsShares);
    }
    /**
     * The IDs of any tags to attach to this resource.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return The IDs of any tags to attach to this resource.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ResourcePool(String name) {
        this(name, ResourcePoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ResourcePool(String name, ResourcePoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ResourcePool(String name, ResourcePoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/resourcePool:ResourcePool", name, args == null ? ResourcePoolArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ResourcePool(String name, Output<String> id, @Nullable ResourcePoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/resourcePool:ResourcePool", name, state, makeResourceOptions(options, id));
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
    public static ResourcePool get(String name, Output<String> id, @Nullable ResourcePoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ResourcePool(name, id, state, options);
    }
}

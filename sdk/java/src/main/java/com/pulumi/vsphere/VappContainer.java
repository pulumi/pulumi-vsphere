// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.VappContainerArgs;
import com.pulumi.vsphere.inputs.VappContainerState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="vsphere:index/vappContainer:VappContainer")
public class VappContainer extends com.pulumi.resources.CustomResource {
    /**
     * Determines if the reservation on a vApp
     * container can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     * 
     */
    @Export(name="cpuExpandable", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> cpuExpandable;

    /**
     * @return Determines if the reservation on a vApp
     * container can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     * 
     */
    public Output<Optional<Boolean>> cpuExpandable() {
        return Codegen.optional(this.cpuExpandable);
    }
    /**
     * The CPU utilization of a vApp container will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited.
     * Default: `-1`
     * 
     */
    @Export(name="cpuLimit", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> cpuLimit;

    /**
     * @return The CPU utilization of a vApp container will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited.
     * Default: `-1`
     * 
     */
    public Output<Optional<Integer>> cpuLimit() {
        return Codegen.optional(this.cpuLimit);
    }
    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the vApp container. Default: `0`
     * 
     */
    @Export(name="cpuReservation", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> cpuReservation;

    /**
     * @return Amount of CPU (MHz) that is guaranteed
     * available to the vApp container. Default: `0`
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
     * Determines if the reservation on a vApp
     * container can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     * 
     */
    @Export(name="memoryExpandable", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> memoryExpandable;

    /**
     * @return Determines if the reservation on a vApp
     * container can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     * 
     */
    public Output<Optional<Boolean>> memoryExpandable() {
        return Codegen.optional(this.memoryExpandable);
    }
    /**
     * The CPU utilization of a vApp container will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited. Default: `-1`
     * 
     */
    @Export(name="memoryLimit", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> memoryLimit;

    /**
     * @return The CPU utilization of a vApp container will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited. Default: `-1`
     * 
     */
    public Output<Optional<Integer>> memoryLimit() {
        return Codegen.optional(this.memoryLimit);
    }
    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the vApp container. Default: `0`
     * 
     */
    @Export(name="memoryReservation", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> memoryReservation;

    /**
     * @return Amount of CPU (MHz) that is guaranteed
     * available to the vApp container. Default: `0`
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
     * The name of the vApp container.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the vApp container.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The managed object ID of
     * the vApp container&#39;s parent folder.
     * 
     */
    @Export(name="parentFolderId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> parentFolderId;

    /**
     * @return The managed object ID of
     * the vApp container&#39;s parent folder.
     * 
     */
    public Output<Optional<String>> parentFolderId() {
        return Codegen.optional(this.parentFolderId);
    }
    /**
     * The managed object ID
     * of the parent resource pool. This can be the root resource pool for a cluster
     * or standalone host, or a resource pool itself. When moving a vApp container
     * from one parent resource pool to another, both must share a common root
     * resource pool or the move will fail.
     * 
     */
    @Export(name="parentResourcePoolId", refs={String.class}, tree="[0]")
    private Output<String> parentResourcePoolId;

    /**
     * @return The managed object ID
     * of the parent resource pool. This can be the root resource pool for a cluster
     * or standalone host, or a resource pool itself. When moving a vApp container
     * from one parent resource pool to another, both must share a common root
     * resource pool or the move will fail.
     * 
     */
    public Output<String> parentResourcePoolId() {
        return this.parentResourcePoolId;
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
    public VappContainer(java.lang.String name) {
        this(name, VappContainerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VappContainer(java.lang.String name, VappContainerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VappContainer(java.lang.String name, VappContainerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/vappContainer:VappContainer", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private VappContainer(java.lang.String name, Output<java.lang.String> id, @Nullable VappContainerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/vappContainer:VappContainer", name, state, makeResourceOptions(options, id), false);
    }

    private static VappContainerArgs makeArgs(VappContainerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VappContainerArgs.Empty : args;
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
    public static VappContainer get(java.lang.String name, Output<java.lang.String> id, @Nullable VappContainerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VappContainer(name, id, state, options);
    }
}

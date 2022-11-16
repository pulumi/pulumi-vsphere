// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VappContainerState extends com.pulumi.resources.ResourceArgs {

    public static final VappContainerState Empty = new VappContainerState();

    /**
     * Determines if the reservation on a vApp
     * container can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     * 
     */
    @Import(name="cpuExpandable")
    private @Nullable Output<Boolean> cpuExpandable;

    /**
     * @return Determines if the reservation on a vApp
     * container can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     * 
     */
    public Optional<Output<Boolean>> cpuExpandable() {
        return Optional.ofNullable(this.cpuExpandable);
    }

    /**
     * The CPU utilization of a vApp container will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited.
     * Default: `-1`
     * 
     */
    @Import(name="cpuLimit")
    private @Nullable Output<Integer> cpuLimit;

    /**
     * @return The CPU utilization of a vApp container will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited.
     * Default: `-1`
     * 
     */
    public Optional<Output<Integer>> cpuLimit() {
        return Optional.ofNullable(this.cpuLimit);
    }

    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the vApp container. Default: `0`
     * 
     */
    @Import(name="cpuReservation")
    private @Nullable Output<Integer> cpuReservation;

    /**
     * @return Amount of CPU (MHz) that is guaranteed
     * available to the vApp container. Default: `0`
     * 
     */
    public Optional<Output<Integer>> cpuReservation() {
        return Optional.ofNullable(this.cpuReservation);
    }

    /**
     * The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `cpu_shares` will be
     * ignored.  Default: `normal`
     * 
     */
    @Import(name="cpuShareLevel")
    private @Nullable Output<String> cpuShareLevel;

    /**
     * @return The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `cpu_shares` will be
     * ignored.  Default: `normal`
     * 
     */
    public Optional<Output<String>> cpuShareLevel() {
        return Optional.ofNullable(this.cpuShareLevel);
    }

    /**
     * The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `cpu_share_level` must be `custom`.
     * 
     */
    @Import(name="cpuShares")
    private @Nullable Output<Integer> cpuShares;

    /**
     * @return The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `cpu_share_level` must be `custom`.
     * 
     */
    public Optional<Output<Integer>> cpuShares() {
        return Optional.ofNullable(this.cpuShares);
    }

    /**
     * A list of custom attributes to set on this resource.
     * 
     */
    @Import(name="customAttributes")
    private @Nullable Output<Map<String,String>> customAttributes;

    /**
     * @return A list of custom attributes to set on this resource.
     * 
     */
    public Optional<Output<Map<String,String>>> customAttributes() {
        return Optional.ofNullable(this.customAttributes);
    }

    /**
     * Determines if the reservation on a vApp
     * container can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     * 
     */
    @Import(name="memoryExpandable")
    private @Nullable Output<Boolean> memoryExpandable;

    /**
     * @return Determines if the reservation on a vApp
     * container can grow beyond the specified value if the parent resource pool has
     * unreserved resources. Default: `true`
     * 
     */
    public Optional<Output<Boolean>> memoryExpandable() {
        return Optional.ofNullable(this.memoryExpandable);
    }

    /**
     * The CPU utilization of a vApp container will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited. Default: `-1`
     * 
     */
    @Import(name="memoryLimit")
    private @Nullable Output<Integer> memoryLimit;

    /**
     * @return The CPU utilization of a vApp container will not
     * exceed this limit, even if there are available resources. Set to `-1` for
     * unlimited. Default: `-1`
     * 
     */
    public Optional<Output<Integer>> memoryLimit() {
        return Optional.ofNullable(this.memoryLimit);
    }

    /**
     * Amount of CPU (MHz) that is guaranteed
     * available to the vApp container. Default: `0`
     * 
     */
    @Import(name="memoryReservation")
    private @Nullable Output<Integer> memoryReservation;

    /**
     * @return Amount of CPU (MHz) that is guaranteed
     * available to the vApp container. Default: `0`
     * 
     */
    public Optional<Output<Integer>> memoryReservation() {
        return Optional.ofNullable(this.memoryReservation);
    }

    /**
     * The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `memory_shares` will be
     * ignored.  Default: `normal`
     * 
     */
    @Import(name="memoryShareLevel")
    private @Nullable Output<String> memoryShareLevel;

    /**
     * @return The CPU allocation level. The level is a
     * simplified view of shares. Levels map to a pre-determined set of numeric
     * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
     * `low`, `normal`, or `high` are specified values in `memory_shares` will be
     * ignored.  Default: `normal`
     * 
     */
    public Optional<Output<String>> memoryShareLevel() {
        return Optional.ofNullable(this.memoryShareLevel);
    }

    /**
     * The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `memory_share_level` must be `custom`.
     * 
     */
    @Import(name="memoryShares")
    private @Nullable Output<Integer> memoryShares;

    /**
     * @return The number of shares allocated for CPU. Used to
     * determine resource allocation in case of resource contention. If this is set,
     * `memory_share_level` must be `custom`.
     * 
     */
    public Optional<Output<Integer>> memoryShares() {
        return Optional.ofNullable(this.memoryShares);
    }

    /**
     * The name of the vApp container.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the vApp container.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The managed object ID of
     * the vApp container&#39;s parent folder.
     * 
     */
    @Import(name="parentFolderId")
    private @Nullable Output<String> parentFolderId;

    /**
     * @return The managed object ID of
     * the vApp container&#39;s parent folder.
     * 
     */
    public Optional<Output<String>> parentFolderId() {
        return Optional.ofNullable(this.parentFolderId);
    }

    /**
     * The managed object ID
     * of the parent resource pool. This can be the root resource pool for a cluster
     * or standalone host, or a resource pool itself. When moving a vApp container
     * from one parent resource pool to another, both must share a common root
     * resource pool or the move will fail.
     * 
     */
    @Import(name="parentResourcePoolId")
    private @Nullable Output<String> parentResourcePoolId;

    /**
     * @return The managed object ID
     * of the parent resource pool. This can be the root resource pool for a cluster
     * or standalone host, or a resource pool itself. When moving a vApp container
     * from one parent resource pool to another, both must share a common root
     * resource pool or the move will fail.
     * 
     */
    public Optional<Output<String>> parentResourcePoolId() {
        return Optional.ofNullable(this.parentResourcePoolId);
    }

    /**
     * The IDs of any tags to attach to this resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The IDs of any tags to attach to this resource.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private VappContainerState() {}

    private VappContainerState(VappContainerState $) {
        this.cpuExpandable = $.cpuExpandable;
        this.cpuLimit = $.cpuLimit;
        this.cpuReservation = $.cpuReservation;
        this.cpuShareLevel = $.cpuShareLevel;
        this.cpuShares = $.cpuShares;
        this.customAttributes = $.customAttributes;
        this.memoryExpandable = $.memoryExpandable;
        this.memoryLimit = $.memoryLimit;
        this.memoryReservation = $.memoryReservation;
        this.memoryShareLevel = $.memoryShareLevel;
        this.memoryShares = $.memoryShares;
        this.name = $.name;
        this.parentFolderId = $.parentFolderId;
        this.parentResourcePoolId = $.parentResourcePoolId;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VappContainerState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VappContainerState $;

        public Builder() {
            $ = new VappContainerState();
        }

        public Builder(VappContainerState defaults) {
            $ = new VappContainerState(Objects.requireNonNull(defaults));
        }

        /**
         * @param cpuExpandable Determines if the reservation on a vApp
         * container can grow beyond the specified value if the parent resource pool has
         * unreserved resources. Default: `true`
         * 
         * @return builder
         * 
         */
        public Builder cpuExpandable(@Nullable Output<Boolean> cpuExpandable) {
            $.cpuExpandable = cpuExpandable;
            return this;
        }

        /**
         * @param cpuExpandable Determines if the reservation on a vApp
         * container can grow beyond the specified value if the parent resource pool has
         * unreserved resources. Default: `true`
         * 
         * @return builder
         * 
         */
        public Builder cpuExpandable(Boolean cpuExpandable) {
            return cpuExpandable(Output.of(cpuExpandable));
        }

        /**
         * @param cpuLimit The CPU utilization of a vApp container will not
         * exceed this limit, even if there are available resources. Set to `-1` for
         * unlimited.
         * Default: `-1`
         * 
         * @return builder
         * 
         */
        public Builder cpuLimit(@Nullable Output<Integer> cpuLimit) {
            $.cpuLimit = cpuLimit;
            return this;
        }

        /**
         * @param cpuLimit The CPU utilization of a vApp container will not
         * exceed this limit, even if there are available resources. Set to `-1` for
         * unlimited.
         * Default: `-1`
         * 
         * @return builder
         * 
         */
        public Builder cpuLimit(Integer cpuLimit) {
            return cpuLimit(Output.of(cpuLimit));
        }

        /**
         * @param cpuReservation Amount of CPU (MHz) that is guaranteed
         * available to the vApp container. Default: `0`
         * 
         * @return builder
         * 
         */
        public Builder cpuReservation(@Nullable Output<Integer> cpuReservation) {
            $.cpuReservation = cpuReservation;
            return this;
        }

        /**
         * @param cpuReservation Amount of CPU (MHz) that is guaranteed
         * available to the vApp container. Default: `0`
         * 
         * @return builder
         * 
         */
        public Builder cpuReservation(Integer cpuReservation) {
            return cpuReservation(Output.of(cpuReservation));
        }

        /**
         * @param cpuShareLevel The CPU allocation level. The level is a
         * simplified view of shares. Levels map to a pre-determined set of numeric
         * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
         * `low`, `normal`, or `high` are specified values in `cpu_shares` will be
         * ignored.  Default: `normal`
         * 
         * @return builder
         * 
         */
        public Builder cpuShareLevel(@Nullable Output<String> cpuShareLevel) {
            $.cpuShareLevel = cpuShareLevel;
            return this;
        }

        /**
         * @param cpuShareLevel The CPU allocation level. The level is a
         * simplified view of shares. Levels map to a pre-determined set of numeric
         * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
         * `low`, `normal`, or `high` are specified values in `cpu_shares` will be
         * ignored.  Default: `normal`
         * 
         * @return builder
         * 
         */
        public Builder cpuShareLevel(String cpuShareLevel) {
            return cpuShareLevel(Output.of(cpuShareLevel));
        }

        /**
         * @param cpuShares The number of shares allocated for CPU. Used to
         * determine resource allocation in case of resource contention. If this is set,
         * `cpu_share_level` must be `custom`.
         * 
         * @return builder
         * 
         */
        public Builder cpuShares(@Nullable Output<Integer> cpuShares) {
            $.cpuShares = cpuShares;
            return this;
        }

        /**
         * @param cpuShares The number of shares allocated for CPU. Used to
         * determine resource allocation in case of resource contention. If this is set,
         * `cpu_share_level` must be `custom`.
         * 
         * @return builder
         * 
         */
        public Builder cpuShares(Integer cpuShares) {
            return cpuShares(Output.of(cpuShares));
        }

        /**
         * @param customAttributes A list of custom attributes to set on this resource.
         * 
         * @return builder
         * 
         */
        public Builder customAttributes(@Nullable Output<Map<String,String>> customAttributes) {
            $.customAttributes = customAttributes;
            return this;
        }

        /**
         * @param customAttributes A list of custom attributes to set on this resource.
         * 
         * @return builder
         * 
         */
        public Builder customAttributes(Map<String,String> customAttributes) {
            return customAttributes(Output.of(customAttributes));
        }

        /**
         * @param memoryExpandable Determines if the reservation on a vApp
         * container can grow beyond the specified value if the parent resource pool has
         * unreserved resources. Default: `true`
         * 
         * @return builder
         * 
         */
        public Builder memoryExpandable(@Nullable Output<Boolean> memoryExpandable) {
            $.memoryExpandable = memoryExpandable;
            return this;
        }

        /**
         * @param memoryExpandable Determines if the reservation on a vApp
         * container can grow beyond the specified value if the parent resource pool has
         * unreserved resources. Default: `true`
         * 
         * @return builder
         * 
         */
        public Builder memoryExpandable(Boolean memoryExpandable) {
            return memoryExpandable(Output.of(memoryExpandable));
        }

        /**
         * @param memoryLimit The CPU utilization of a vApp container will not
         * exceed this limit, even if there are available resources. Set to `-1` for
         * unlimited. Default: `-1`
         * 
         * @return builder
         * 
         */
        public Builder memoryLimit(@Nullable Output<Integer> memoryLimit) {
            $.memoryLimit = memoryLimit;
            return this;
        }

        /**
         * @param memoryLimit The CPU utilization of a vApp container will not
         * exceed this limit, even if there are available resources. Set to `-1` for
         * unlimited. Default: `-1`
         * 
         * @return builder
         * 
         */
        public Builder memoryLimit(Integer memoryLimit) {
            return memoryLimit(Output.of(memoryLimit));
        }

        /**
         * @param memoryReservation Amount of CPU (MHz) that is guaranteed
         * available to the vApp container. Default: `0`
         * 
         * @return builder
         * 
         */
        public Builder memoryReservation(@Nullable Output<Integer> memoryReservation) {
            $.memoryReservation = memoryReservation;
            return this;
        }

        /**
         * @param memoryReservation Amount of CPU (MHz) that is guaranteed
         * available to the vApp container. Default: `0`
         * 
         * @return builder
         * 
         */
        public Builder memoryReservation(Integer memoryReservation) {
            return memoryReservation(Output.of(memoryReservation));
        }

        /**
         * @param memoryShareLevel The CPU allocation level. The level is a
         * simplified view of shares. Levels map to a pre-determined set of numeric
         * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
         * `low`, `normal`, or `high` are specified values in `memory_shares` will be
         * ignored.  Default: `normal`
         * 
         * @return builder
         * 
         */
        public Builder memoryShareLevel(@Nullable Output<String> memoryShareLevel) {
            $.memoryShareLevel = memoryShareLevel;
            return this;
        }

        /**
         * @param memoryShareLevel The CPU allocation level. The level is a
         * simplified view of shares. Levels map to a pre-determined set of numeric
         * values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
         * `low`, `normal`, or `high` are specified values in `memory_shares` will be
         * ignored.  Default: `normal`
         * 
         * @return builder
         * 
         */
        public Builder memoryShareLevel(String memoryShareLevel) {
            return memoryShareLevel(Output.of(memoryShareLevel));
        }

        /**
         * @param memoryShares The number of shares allocated for CPU. Used to
         * determine resource allocation in case of resource contention. If this is set,
         * `memory_share_level` must be `custom`.
         * 
         * @return builder
         * 
         */
        public Builder memoryShares(@Nullable Output<Integer> memoryShares) {
            $.memoryShares = memoryShares;
            return this;
        }

        /**
         * @param memoryShares The number of shares allocated for CPU. Used to
         * determine resource allocation in case of resource contention. If this is set,
         * `memory_share_level` must be `custom`.
         * 
         * @return builder
         * 
         */
        public Builder memoryShares(Integer memoryShares) {
            return memoryShares(Output.of(memoryShares));
        }

        /**
         * @param name The name of the vApp container.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the vApp container.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param parentFolderId The managed object ID of
         * the vApp container&#39;s parent folder.
         * 
         * @return builder
         * 
         */
        public Builder parentFolderId(@Nullable Output<String> parentFolderId) {
            $.parentFolderId = parentFolderId;
            return this;
        }

        /**
         * @param parentFolderId The managed object ID of
         * the vApp container&#39;s parent folder.
         * 
         * @return builder
         * 
         */
        public Builder parentFolderId(String parentFolderId) {
            return parentFolderId(Output.of(parentFolderId));
        }

        /**
         * @param parentResourcePoolId The managed object ID
         * of the parent resource pool. This can be the root resource pool for a cluster
         * or standalone host, or a resource pool itself. When moving a vApp container
         * from one parent resource pool to another, both must share a common root
         * resource pool or the move will fail.
         * 
         * @return builder
         * 
         */
        public Builder parentResourcePoolId(@Nullable Output<String> parentResourcePoolId) {
            $.parentResourcePoolId = parentResourcePoolId;
            return this;
        }

        /**
         * @param parentResourcePoolId The managed object ID
         * of the parent resource pool. This can be the root resource pool for a cluster
         * or standalone host, or a resource pool itself. When moving a vApp container
         * from one parent resource pool to another, both must share a common root
         * resource pool or the move will fail.
         * 
         * @return builder
         * 
         */
        public Builder parentResourcePoolId(String parentResourcePoolId) {
            return parentResourcePoolId(Output.of(parentResourcePoolId));
        }

        /**
         * @param tags The IDs of any tags to attach to this resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The IDs of any tags to attach to this resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags The IDs of any tags to attach to this resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        public VappContainerState build() {
            return $;
        }
    }

}
// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VmfsDatastoreArgs extends com.pulumi.resources.ResourceArgs {

    public static final VmfsDatastoreArgs Empty = new VmfsDatastoreArgs();

    /**
     * Map of custom attribute ids to attribute
     * value string to set on datastore resource.
     * 
     * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     * 
     */
    @Import(name="customAttributes")
    private @Nullable Output<Map<String,String>> customAttributes;

    /**
     * @return Map of custom attribute ids to attribute
     * value string to set on datastore resource.
     * 
     * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     * 
     */
    public Optional<Output<Map<String,String>>> customAttributes() {
        return Optional.ofNullable(this.customAttributes);
    }

    /**
     * The managed object
     * ID of a datastore cluster to put this datastore in.
     * Conflicts with `folder`.
     * 
     */
    @Import(name="datastoreClusterId")
    private @Nullable Output<String> datastoreClusterId;

    /**
     * @return The managed object
     * ID of a datastore cluster to put this datastore in.
     * Conflicts with `folder`.
     * 
     */
    public Optional<Output<String>> datastoreClusterId() {
        return Optional.ofNullable(this.datastoreClusterId);
    }

    /**
     * The disks to use with the datastore.
     * 
     */
    @Import(name="disks", required=true)
    private Output<List<String>> disks;

    /**
     * @return The disks to use with the datastore.
     * 
     */
    public Output<List<String>> disks() {
        return this.disks;
    }

    /**
     * The relative path to a folder to put this datastore in.
     * This is a path relative to the datacenter you are deploying the datastore to.
     * Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
     * The provider will place a datastore named `test` in a datastore folder
     * located at `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/test`. Conflicts with
     * `datastore_cluster_id`.
     * 
     */
    @Import(name="folder")
    private @Nullable Output<String> folder;

    /**
     * @return The relative path to a folder to put this datastore in.
     * This is a path relative to the datacenter you are deploying the datastore to.
     * Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
     * The provider will place a datastore named `test` in a datastore folder
     * located at `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/test`. Conflicts with
     * `datastore_cluster_id`.
     * 
     */
    public Optional<Output<String>> folder() {
        return Optional.ofNullable(this.folder);
    }

    /**
     * The managed object ID of
     * the host to set the datastore up on. Note that this is not necessarily the
     * only host that the datastore will be set up on - see
     * here for more info. Forces a
     * new resource if changed.
     * 
     */
    @Import(name="hostSystemId", required=true)
    private Output<String> hostSystemId;

    /**
     * @return The managed object ID of
     * the host to set the datastore up on. Note that this is not necessarily the
     * only host that the datastore will be set up on - see
     * here for more info. Forces a
     * new resource if changed.
     * 
     */
    public Output<String> hostSystemId() {
        return this.hostSystemId;
    }

    /**
     * The name of the datastore. Forces a new resource if
     * changed.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the datastore. Forces a new resource if
     * changed.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The IDs of any tags to attach to this resource.
     * 
     * &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The IDs of any tags to attach to this resource.
     * 
     * &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private VmfsDatastoreArgs() {}

    private VmfsDatastoreArgs(VmfsDatastoreArgs $) {
        this.customAttributes = $.customAttributes;
        this.datastoreClusterId = $.datastoreClusterId;
        this.disks = $.disks;
        this.folder = $.folder;
        this.hostSystemId = $.hostSystemId;
        this.name = $.name;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VmfsDatastoreArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VmfsDatastoreArgs $;

        public Builder() {
            $ = new VmfsDatastoreArgs();
        }

        public Builder(VmfsDatastoreArgs defaults) {
            $ = new VmfsDatastoreArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param customAttributes Map of custom attribute ids to attribute
         * value string to set on datastore resource.
         * 
         * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
         * and require vCenter.
         * 
         * @return builder
         * 
         */
        public Builder customAttributes(@Nullable Output<Map<String,String>> customAttributes) {
            $.customAttributes = customAttributes;
            return this;
        }

        /**
         * @param customAttributes Map of custom attribute ids to attribute
         * value string to set on datastore resource.
         * 
         * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
         * and require vCenter.
         * 
         * @return builder
         * 
         */
        public Builder customAttributes(Map<String,String> customAttributes) {
            return customAttributes(Output.of(customAttributes));
        }

        /**
         * @param datastoreClusterId The managed object
         * ID of a datastore cluster to put this datastore in.
         * Conflicts with `folder`.
         * 
         * @return builder
         * 
         */
        public Builder datastoreClusterId(@Nullable Output<String> datastoreClusterId) {
            $.datastoreClusterId = datastoreClusterId;
            return this;
        }

        /**
         * @param datastoreClusterId The managed object
         * ID of a datastore cluster to put this datastore in.
         * Conflicts with `folder`.
         * 
         * @return builder
         * 
         */
        public Builder datastoreClusterId(String datastoreClusterId) {
            return datastoreClusterId(Output.of(datastoreClusterId));
        }

        /**
         * @param disks The disks to use with the datastore.
         * 
         * @return builder
         * 
         */
        public Builder disks(Output<List<String>> disks) {
            $.disks = disks;
            return this;
        }

        /**
         * @param disks The disks to use with the datastore.
         * 
         * @return builder
         * 
         */
        public Builder disks(List<String> disks) {
            return disks(Output.of(disks));
        }

        /**
         * @param disks The disks to use with the datastore.
         * 
         * @return builder
         * 
         */
        public Builder disks(String... disks) {
            return disks(List.of(disks));
        }

        /**
         * @param folder The relative path to a folder to put this datastore in.
         * This is a path relative to the datacenter you are deploying the datastore to.
         * Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
         * The provider will place a datastore named `test` in a datastore folder
         * located at `/dc1/datastore/foo/bar`, with the final inventory path being
         * `/dc1/datastore/foo/bar/test`. Conflicts with
         * `datastore_cluster_id`.
         * 
         * @return builder
         * 
         */
        public Builder folder(@Nullable Output<String> folder) {
            $.folder = folder;
            return this;
        }

        /**
         * @param folder The relative path to a folder to put this datastore in.
         * This is a path relative to the datacenter you are deploying the datastore to.
         * Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
         * The provider will place a datastore named `test` in a datastore folder
         * located at `/dc1/datastore/foo/bar`, with the final inventory path being
         * `/dc1/datastore/foo/bar/test`. Conflicts with
         * `datastore_cluster_id`.
         * 
         * @return builder
         * 
         */
        public Builder folder(String folder) {
            return folder(Output.of(folder));
        }

        /**
         * @param hostSystemId The managed object ID of
         * the host to set the datastore up on. Note that this is not necessarily the
         * only host that the datastore will be set up on - see
         * here for more info. Forces a
         * new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder hostSystemId(Output<String> hostSystemId) {
            $.hostSystemId = hostSystemId;
            return this;
        }

        /**
         * @param hostSystemId The managed object ID of
         * the host to set the datastore up on. Note that this is not necessarily the
         * only host that the datastore will be set up on - see
         * here for more info. Forces a
         * new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder hostSystemId(String hostSystemId) {
            return hostSystemId(Output.of(hostSystemId));
        }

        /**
         * @param name The name of the datastore. Forces a new resource if
         * changed.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the datastore. Forces a new resource if
         * changed.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param tags The IDs of any tags to attach to this resource.
         * 
         * &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
         * requires vCenter 6.0 or higher.
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
         * &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
         * requires vCenter 6.0 or higher.
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
         * &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
         * requires vCenter 6.0 or higher.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        public VmfsDatastoreArgs build() {
            $.disks = Objects.requireNonNull($.disks, "expected parameter 'disks' to be non-null");
            $.hostSystemId = Objects.requireNonNull($.hostSystemId, "expected parameter 'hostSystemId' to be non-null");
            return $;
        }
    }

}

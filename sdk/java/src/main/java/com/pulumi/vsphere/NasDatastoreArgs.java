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


public final class NasDatastoreArgs extends com.pulumi.resources.ResourceArgs {

    public static final NasDatastoreArgs Empty = new NasDatastoreArgs();

    /**
     * Access mode for the mount point. Can be one of
     * `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
     * that the datastore will be read-write depending on the permissions of the
     * actual share. Default: `readWrite`. Forces a new resource if changed.
     * 
     */
    @Import(name="accessMode")
    private @Nullable Output<String> accessMode;

    /**
     * @return Access mode for the mount point. Can be one of
     * `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
     * that the datastore will be read-write depending on the permissions of the
     * actual share. Default: `readWrite`. Forces a new resource if changed.
     * 
     */
    public Optional<Output<String>> accessMode() {
        return Optional.ofNullable(this.accessMode);
    }

    /**
     * Map of custom attribute ids to attribute
     * value strings to set on datasource resource.
     * 
     * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     * 
     */
    @Import(name="customAttributes")
    private @Nullable Output<Map<String,String>> customAttributes;

    /**
     * @return Map of custom attribute ids to attribute
     * value strings to set on datasource resource.
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
     * The managed object IDs of
     * the hosts to mount the datastore on.
     * 
     */
    @Import(name="hostSystemIds", required=true)
    private Output<List<String>> hostSystemIds;

    /**
     * @return The managed object IDs of
     * the hosts to mount the datastore on.
     * 
     */
    public Output<List<String>> hostSystemIds() {
        return this.hostSystemIds;
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
     * The hostnames or IP addresses of the remote
     * server or servers. Only one element should be present for NFS v3 but multiple
     * can be present for NFS v4.1. Forces a new resource if changed.
     * 
     */
    @Import(name="remoteHosts", required=true)
    private Output<List<String>> remoteHosts;

    /**
     * @return The hostnames or IP addresses of the remote
     * server or servers. Only one element should be present for NFS v3 but multiple
     * can be present for NFS v4.1. Forces a new resource if changed.
     * 
     */
    public Output<List<String>> remoteHosts() {
        return this.remoteHosts;
    }

    /**
     * The remote path of the mount point. Forces a new
     * resource if changed.
     * 
     */
    @Import(name="remotePath", required=true)
    private Output<String> remotePath;

    /**
     * @return The remote path of the mount point. Forces a new
     * resource if changed.
     * 
     */
    public Output<String> remotePath() {
        return this.remotePath;
    }

    /**
     * The security type to use when using NFS v4.1.
     * Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
     * if changed.
     * 
     */
    @Import(name="securityType")
    private @Nullable Output<String> securityType;

    /**
     * @return The security type to use when using NFS v4.1.
     * Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
     * if changed.
     * 
     */
    public Optional<Output<String>> securityType() {
        return Optional.ofNullable(this.securityType);
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

    /**
     * The type of NAS volume. Can be one of `NFS` (to denote
     * v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
     * changed.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The type of NAS volume. Can be one of `NFS` (to denote
     * v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
     * changed.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private NasDatastoreArgs() {}

    private NasDatastoreArgs(NasDatastoreArgs $) {
        this.accessMode = $.accessMode;
        this.customAttributes = $.customAttributes;
        this.datastoreClusterId = $.datastoreClusterId;
        this.folder = $.folder;
        this.hostSystemIds = $.hostSystemIds;
        this.name = $.name;
        this.remoteHosts = $.remoteHosts;
        this.remotePath = $.remotePath;
        this.securityType = $.securityType;
        this.tags = $.tags;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NasDatastoreArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NasDatastoreArgs $;

        public Builder() {
            $ = new NasDatastoreArgs();
        }

        public Builder(NasDatastoreArgs defaults) {
            $ = new NasDatastoreArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessMode Access mode for the mount point. Can be one of
         * `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
         * that the datastore will be read-write depending on the permissions of the
         * actual share. Default: `readWrite`. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder accessMode(@Nullable Output<String> accessMode) {
            $.accessMode = accessMode;
            return this;
        }

        /**
         * @param accessMode Access mode for the mount point. Can be one of
         * `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
         * that the datastore will be read-write depending on the permissions of the
         * actual share. Default: `readWrite`. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder accessMode(String accessMode) {
            return accessMode(Output.of(accessMode));
        }

        /**
         * @param customAttributes Map of custom attribute ids to attribute
         * value strings to set on datasource resource.
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
         * value strings to set on datasource resource.
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
         * @param hostSystemIds The managed object IDs of
         * the hosts to mount the datastore on.
         * 
         * @return builder
         * 
         */
        public Builder hostSystemIds(Output<List<String>> hostSystemIds) {
            $.hostSystemIds = hostSystemIds;
            return this;
        }

        /**
         * @param hostSystemIds The managed object IDs of
         * the hosts to mount the datastore on.
         * 
         * @return builder
         * 
         */
        public Builder hostSystemIds(List<String> hostSystemIds) {
            return hostSystemIds(Output.of(hostSystemIds));
        }

        /**
         * @param hostSystemIds The managed object IDs of
         * the hosts to mount the datastore on.
         * 
         * @return builder
         * 
         */
        public Builder hostSystemIds(String... hostSystemIds) {
            return hostSystemIds(List.of(hostSystemIds));
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
         * @param remoteHosts The hostnames or IP addresses of the remote
         * server or servers. Only one element should be present for NFS v3 but multiple
         * can be present for NFS v4.1. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder remoteHosts(Output<List<String>> remoteHosts) {
            $.remoteHosts = remoteHosts;
            return this;
        }

        /**
         * @param remoteHosts The hostnames or IP addresses of the remote
         * server or servers. Only one element should be present for NFS v3 but multiple
         * can be present for NFS v4.1. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder remoteHosts(List<String> remoteHosts) {
            return remoteHosts(Output.of(remoteHosts));
        }

        /**
         * @param remoteHosts The hostnames or IP addresses of the remote
         * server or servers. Only one element should be present for NFS v3 but multiple
         * can be present for NFS v4.1. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder remoteHosts(String... remoteHosts) {
            return remoteHosts(List.of(remoteHosts));
        }

        /**
         * @param remotePath The remote path of the mount point. Forces a new
         * resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder remotePath(Output<String> remotePath) {
            $.remotePath = remotePath;
            return this;
        }

        /**
         * @param remotePath The remote path of the mount point. Forces a new
         * resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder remotePath(String remotePath) {
            return remotePath(Output.of(remotePath));
        }

        /**
         * @param securityType The security type to use when using NFS v4.1.
         * Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
         * if changed.
         * 
         * @return builder
         * 
         */
        public Builder securityType(@Nullable Output<String> securityType) {
            $.securityType = securityType;
            return this;
        }

        /**
         * @param securityType The security type to use when using NFS v4.1.
         * Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
         * if changed.
         * 
         * @return builder
         * 
         */
        public Builder securityType(String securityType) {
            return securityType(Output.of(securityType));
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

        /**
         * @param type The type of NAS volume. Can be one of `NFS` (to denote
         * v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
         * changed.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of NAS volume. Can be one of `NFS` (to denote
         * v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
         * changed.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public NasDatastoreArgs build() {
            $.hostSystemIds = Objects.requireNonNull($.hostSystemIds, "expected parameter 'hostSystemIds' to be non-null");
            $.remoteHosts = Objects.requireNonNull($.remoteHosts, "expected parameter 'remoteHosts' to be non-null");
            $.remotePath = Objects.requireNonNull($.remotePath, "expected parameter 'remotePath' to be non-null");
            return $;
        }
    }

}

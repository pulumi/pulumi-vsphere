// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.NasDatastoreArgs;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.inputs.NasDatastoreState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="vsphere:index/nasDatastore:NasDatastore")
public class NasDatastore extends com.pulumi.resources.CustomResource {
    /**
     * Access mode for the mount point. Can be one of
     * `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
     * that the datastore will be read-write depending on the permissions of the
     * actual share. Default: `readWrite`. Forces a new resource if changed.
     * 
     */
    @Export(name="accessMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessMode;

    /**
     * @return Access mode for the mount point. Can be one of
     * `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
     * that the datastore will be read-write depending on the permissions of the
     * actual share. Default: `readWrite`. Forces a new resource if changed.
     * 
     */
    public Output<Optional<String>> accessMode() {
        return Codegen.optional(this.accessMode);
    }
    /**
     * The connectivity status of the datastore. If this is `false`,
     * some other computed attributes may be out of date.
     * 
     */
    @Export(name="accessible", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> accessible;

    /**
     * @return The connectivity status of the datastore. If this is `false`,
     * some other computed attributes may be out of date.
     * 
     */
    public Output<Boolean> accessible() {
        return this.accessible;
    }
    /**
     * Maximum capacity of the datastore, in megabytes.
     * 
     */
    @Export(name="capacity", refs={Integer.class}, tree="[0]")
    private Output<Integer> capacity;

    /**
     * @return Maximum capacity of the datastore, in megabytes.
     * 
     */
    public Output<Integer> capacity() {
        return this.capacity;
    }
    /**
     * Map of custom attribute ids to attribute
     * value strings to set on datasource resource.
     * 
     * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     * 
     */
    @Export(name="customAttributes", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> customAttributes;

    /**
     * @return Map of custom attribute ids to attribute
     * value strings to set on datasource resource.
     * 
     * &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     * 
     */
    public Output<Optional<Map<String,String>>> customAttributes() {
        return Codegen.optional(this.customAttributes);
    }
    /**
     * The managed object
     * ID of a datastore cluster to put this datastore in.
     * Conflicts with `folder`.
     * 
     */
    @Export(name="datastoreClusterId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> datastoreClusterId;

    /**
     * @return The managed object
     * ID of a datastore cluster to put this datastore in.
     * Conflicts with `folder`.
     * 
     */
    public Output<Optional<String>> datastoreClusterId() {
        return Codegen.optional(this.datastoreClusterId);
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
    @Export(name="folder", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> folder;

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
    public Output<Optional<String>> folder() {
        return Codegen.optional(this.folder);
    }
    /**
     * Available space of this datastore, in megabytes.
     * 
     */
    @Export(name="freeSpace", refs={Integer.class}, tree="[0]")
    private Output<Integer> freeSpace;

    /**
     * @return Available space of this datastore, in megabytes.
     * 
     */
    public Output<Integer> freeSpace() {
        return this.freeSpace;
    }
    /**
     * The managed object IDs of
     * the hosts to mount the datastore on.
     * 
     */
    @Export(name="hostSystemIds", refs={List.class,String.class}, tree="[0,1]")
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
     * The current maintenance mode state of the datastore.
     * 
     */
    @Export(name="maintenanceMode", refs={String.class}, tree="[0]")
    private Output<String> maintenanceMode;

    /**
     * @return The current maintenance mode state of the datastore.
     * 
     */
    public Output<String> maintenanceMode() {
        return this.maintenanceMode;
    }
    /**
     * If `true`, more than one host in the datacenter has
     * been configured with access to the datastore.
     * 
     */
    @Export(name="multipleHostAccess", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> multipleHostAccess;

    /**
     * @return If `true`, more than one host in the datacenter has
     * been configured with access to the datastore.
     * 
     */
    public Output<Boolean> multipleHostAccess() {
        return this.multipleHostAccess;
    }
    /**
     * The name of the datastore. Forces a new resource if
     * changed.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the datastore. Forces a new resource if
     * changed.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Indicates that this NAS volume is a protocol endpoint.
     * This field is only populated if the host supports virtual datastores.
     * 
     */
    @Export(name="protocolEndpoint", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> protocolEndpoint;

    /**
     * @return Indicates that this NAS volume is a protocol endpoint.
     * This field is only populated if the host supports virtual datastores.
     * 
     */
    public Output<Boolean> protocolEndpoint() {
        return this.protocolEndpoint;
    }
    /**
     * The hostnames or IP addresses of the remote
     * servers. Only one element should be present for NFS v3 but multiple
     * can be present for NFS v4.1. Forces a new resource if changed.
     * 
     */
    @Export(name="remoteHosts", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> remoteHosts;

    /**
     * @return The hostnames or IP addresses of the remote
     * servers. Only one element should be present for NFS v3 but multiple
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
    @Export(name="remotePath", refs={String.class}, tree="[0]")
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
    @Export(name="securityType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> securityType;

    /**
     * @return The security type to use when using NFS v4.1.
     * Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
     * if changed.
     * 
     */
    public Output<Optional<String>> securityType() {
        return Codegen.optional(this.securityType);
    }
    /**
     * The IDs of any tags to attach to this resource.
     * 
     * &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return The IDs of any tags to attach to this resource.
     * 
     * &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The type of NAS volume. Can be one of `NFS` (to denote
     * v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
     * changed.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return The type of NAS volume. Can be one of `NFS` (to denote
     * v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
     * changed.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }
    /**
     * Total additional storage space, in megabytes,
     * potentially used by all virtual machines on this datastore.
     * 
     */
    @Export(name="uncommittedSpace", refs={Integer.class}, tree="[0]")
    private Output<Integer> uncommittedSpace;

    /**
     * @return Total additional storage space, in megabytes,
     * potentially used by all virtual machines on this datastore.
     * 
     */
    public Output<Integer> uncommittedSpace() {
        return this.uncommittedSpace;
    }
    /**
     * The unique locator for the datastore.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The unique locator for the datastore.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NasDatastore(java.lang.String name) {
        this(name, NasDatastoreArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NasDatastore(java.lang.String name, NasDatastoreArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NasDatastore(java.lang.String name, NasDatastoreArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/nasDatastore:NasDatastore", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private NasDatastore(java.lang.String name, Output<java.lang.String> id, @Nullable NasDatastoreState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/nasDatastore:NasDatastore", name, state, makeResourceOptions(options, id), false);
    }

    private static NasDatastoreArgs makeArgs(NasDatastoreArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? NasDatastoreArgs.Empty : args;
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
    public static NasDatastore get(java.lang.String name, Output<java.lang.String> id, @Nullable NasDatastoreState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NasDatastore(name, id, state, options);
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.VmfsDatastoreArgs;
import com.pulumi.vsphere.inputs.VmfsDatastoreState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="vsphere:index/vmfsDatastore:VmfsDatastore")
public class VmfsDatastore extends com.pulumi.resources.CustomResource {
    /**
     * The connectivity status of the datastore. If this is `false`,
     * some other computed attributes may be out of date.
     * 
     */
    @Export(name="accessible", type=Boolean.class, parameters={})
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
    @Export(name="capacity", type=Integer.class, parameters={})
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
     * value string to set on datastore resource.
     * 
     */
    @Export(name="customAttributes", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> customAttributes;

    /**
     * @return Map of custom attribute ids to attribute
     * value string to set on datastore resource.
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
    @Export(name="datastoreClusterId", type=String.class, parameters={})
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
     * The disks to use with the datastore.
     * 
     */
    @Export(name="disks", type=List.class, parameters={String.class})
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
    @Export(name="folder", type=String.class, parameters={})
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
    @Export(name="freeSpace", type=Integer.class, parameters={})
    private Output<Integer> freeSpace;

    /**
     * @return Available space of this datastore, in megabytes.
     * 
     */
    public Output<Integer> freeSpace() {
        return this.freeSpace;
    }
    /**
     * The managed object ID of
     * the host to set the datastore up on. Note that this is not necessarily the
     * only host that the datastore will be set up on - see
     * here for more info. Forces a
     * new resource if changed.
     * 
     */
    @Export(name="hostSystemId", type=String.class, parameters={})
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
     * The current maintenance mode state of the datastore.
     * 
     */
    @Export(name="maintenanceMode", type=String.class, parameters={})
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
    @Export(name="multipleHostAccess", type=Boolean.class, parameters={})
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
    @Export(name="name", type=String.class, parameters={})
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
     * The IDs of any tags to attach to this resource.
     * 
     */
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return The IDs of any tags to attach to this resource.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Total additional storage space, in megabytes,
     * potentially used by all virtual machines on this datastore.
     * 
     */
    @Export(name="uncommittedSpace", type=Integer.class, parameters={})
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
    @Export(name="url", type=String.class, parameters={})
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
    public VmfsDatastore(String name) {
        this(name, VmfsDatastoreArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VmfsDatastore(String name, VmfsDatastoreArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VmfsDatastore(String name, VmfsDatastoreArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/vmfsDatastore:VmfsDatastore", name, args == null ? VmfsDatastoreArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VmfsDatastore(String name, Output<String> id, @Nullable VmfsDatastoreState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/vmfsDatastore:VmfsDatastore", name, state, makeResourceOptions(options, id));
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
    public static VmfsDatastore get(String name, Output<String> id, @Nullable VmfsDatastoreState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VmfsDatastore(name, id, state, options);
    }
}

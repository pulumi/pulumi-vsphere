// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.ComputeClusterHostGroupArgs;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.inputs.ComputeClusterHostGroupState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `vsphere.ComputeClusterHostGroup` resource can be used to manage groups
 * of hosts in a cluster, either created by the
 * `vsphere.ComputeCluster` resource or looked up
 * by the `vsphere.ComputeCluster` data source.
 * 
 * This resource mainly serves as an input to the
 * `vsphere.ComputeClusterVmHostRule`
 * resource - see the documentation for that resource for further details on how
 * to use host groups.
 * 
 * &gt; **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 * 
 * ## Import
 * 
 * An existing group can be imported into this resource by
 * 
 * supplying both the path to the cluster, and the name of the host group. If the
 * 
 * name or cluster is not found, or if the group is of a different type, an error
 * 
 * will be returned. An example is below:
 * 
 * ```sh
 * $ pulumi import vsphere:index/computeClusterHostGroup:ComputeClusterHostGroup cluster_host_group \
 * ```
 * 
 *   &#39;{&#34;compute_cluster_path&#34;: &#34;/dc1/host/cluster1&#34;, \
 * 
 *   &#34;name&#34;: &#34;pulumi-test-cluster-host-group&#34;}&#39;
 * 
 */
@ResourceType(type="vsphere:index/computeClusterHostGroup:ComputeClusterHostGroup")
public class ComputeClusterHostGroup extends com.pulumi.resources.CustomResource {
    /**
     * The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     * 
     */
    @Export(name="computeClusterId", refs={String.class}, tree="[0]")
    private Output<String> computeClusterId;

    /**
     * @return The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     * 
     */
    public Output<String> computeClusterId() {
        return this.computeClusterId;
    }
    /**
     * The managed object IDs of
     * the hosts to put in the cluster.
     * 
     * &gt; **NOTE:** The namespace for cluster names on this resource (defined by the
     * `name` argument) is shared with the
     * `vsphere.ComputeClusterVmGroup`
     * resource. Make sure your names are unique across both resources.
     * 
     */
    @Export(name="hostSystemIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> hostSystemIds;

    /**
     * @return The managed object IDs of
     * the hosts to put in the cluster.
     * 
     * &gt; **NOTE:** The namespace for cluster names on this resource (defined by the
     * `name` argument) is shared with the
     * `vsphere.ComputeClusterVmGroup`
     * resource. Make sure your names are unique across both resources.
     * 
     */
    public Output<Optional<List<String>>> hostSystemIds() {
        return Codegen.optional(this.hostSystemIds);
    }
    /**
     * The name of the host group. This must be unique in the
     * cluster. Forces a new resource if changed.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the host group. This must be unique in the
     * cluster. Forces a new resource if changed.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ComputeClusterHostGroup(java.lang.String name) {
        this(name, ComputeClusterHostGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ComputeClusterHostGroup(java.lang.String name, ComputeClusterHostGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ComputeClusterHostGroup(java.lang.String name, ComputeClusterHostGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/computeClusterHostGroup:ComputeClusterHostGroup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ComputeClusterHostGroup(java.lang.String name, Output<java.lang.String> id, @Nullable ComputeClusterHostGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/computeClusterHostGroup:ComputeClusterHostGroup", name, state, makeResourceOptions(options, id), false);
    }

    private static ComputeClusterHostGroupArgs makeArgs(ComputeClusterHostGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ComputeClusterHostGroupArgs.Empty : args;
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
    public static ComputeClusterHostGroup get(java.lang.String name, Output<java.lang.String> id, @Nullable ComputeClusterHostGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ComputeClusterHostGroup(name, id, state, options);
    }
}

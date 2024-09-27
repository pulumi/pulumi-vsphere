// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.DpmHostOverrideArgs;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.inputs.DpmHostOverrideState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The `vsphere.DpmHostOverride` resource can be used to add a DPM override to a
 * cluster for a particular host. This allows you to control the power management
 * settings for individual hosts in the cluster while leaving any unspecified ones
 * at the default power management settings.
 * 
 * For more information on DPM within vSphere clusters, see [this
 * page][ref-vsphere-cluster-dpm].
 * 
 * [ref-vsphere-cluster-dpm]: https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-resource-management/GUID-5E5E349A-4644-4C9C-B434-1C0243EBDC80.html
 * 
 * &gt; **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 * 
 * ## Import
 * 
 * An existing override can be imported into this resource by
 * 
 * supplying both the path to the cluster, and the path to the host, to `terraform
 * 
 * import`. If no override exists, an error will be given.  An example is below:
 * 
 * ```sh
 * $ pulumi import vsphere:index/dpmHostOverride:DpmHostOverride dpm_host_override \
 * ```
 * 
 *   &#39;{&#34;compute_cluster_path&#34;: &#34;/dc1/host/cluster1&#34;, \
 * 
 *   &#34;host_path&#34;: &#34;/dc1/host/esxi1&#34;}&#39;
 * 
 */
@ResourceType(type="vsphere:index/dpmHostOverride:DpmHostOverride")
public class DpmHostOverride extends com.pulumi.resources.CustomResource {
    /**
     * The managed object reference
     * ID of the cluster to put the override in.  Forces a new
     * resource if changed.
     * 
     */
    @Export(name="computeClusterId", refs={String.class}, tree="[0]")
    private Output<String> computeClusterId;

    /**
     * @return The managed object reference
     * ID of the cluster to put the override in.  Forces a new
     * resource if changed.
     * 
     */
    public Output<String> computeClusterId() {
        return this.computeClusterId;
    }
    /**
     * The automation level for host power
     * operations on this host. Can be one of `manual` or `automated`. Default:
     * `manual`.
     * 
     * &gt; **NOTE:** Using this resource _always_ implies an override, even if one of
     * `dpm_enabled` or `dpm_automation_level` is omitted. Take note of the defaults
     * for both options.
     * 
     */
    @Export(name="dpmAutomationLevel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dpmAutomationLevel;

    /**
     * @return The automation level for host power
     * operations on this host. Can be one of `manual` or `automated`. Default:
     * `manual`.
     * 
     * &gt; **NOTE:** Using this resource _always_ implies an override, even if one of
     * `dpm_enabled` or `dpm_automation_level` is omitted. Take note of the defaults
     * for both options.
     * 
     */
    public Output<Optional<String>> dpmAutomationLevel() {
        return Codegen.optional(this.dpmAutomationLevel);
    }
    /**
     * Enable DPM support for this host. Default:
     * `false`.
     * 
     */
    @Export(name="dpmEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dpmEnabled;

    /**
     * @return Enable DPM support for this host. Default:
     * `false`.
     * 
     */
    public Output<Optional<Boolean>> dpmEnabled() {
        return Codegen.optional(this.dpmEnabled);
    }
    /**
     * The managed object ID of the host.
     * 
     */
    @Export(name="hostSystemId", refs={String.class}, tree="[0]")
    private Output<String> hostSystemId;

    /**
     * @return The managed object ID of the host.
     * 
     */
    public Output<String> hostSystemId() {
        return this.hostSystemId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DpmHostOverride(java.lang.String name) {
        this(name, DpmHostOverrideArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DpmHostOverride(java.lang.String name, DpmHostOverrideArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DpmHostOverride(java.lang.String name, DpmHostOverrideArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/dpmHostOverride:DpmHostOverride", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DpmHostOverride(java.lang.String name, Output<java.lang.String> id, @Nullable DpmHostOverrideState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/dpmHostOverride:DpmHostOverride", name, state, makeResourceOptions(options, id), false);
    }

    private static DpmHostOverrideArgs makeArgs(DpmHostOverrideArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DpmHostOverrideArgs.Empty : args;
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
    public static DpmHostOverride get(java.lang.String name, Output<java.lang.String> id, @Nullable DpmHostOverrideState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DpmHostOverride(name, id, state, options);
    }
}

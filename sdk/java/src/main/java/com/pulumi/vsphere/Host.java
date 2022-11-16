// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.HostArgs;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.inputs.HostState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a VMware vSphere host resource. This represents an ESXi host that
 * can be used either as a member of a cluster or as a standalone host.
 * 
 * ## Example Usage
 * ## Importing
 * 
 * An existing host can be [imported][docs-import] into this resource by supplying
 * the host&#39;s ID. An example is below:
 * 
 * [docs-import]: /docs/import/index.html
 * 
 * The above would import the host with ID `host-123`.
 * 
 */
@ResourceType(type="vsphere:index/host:Host")
public class Host extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the Compute Cluster this host should
     * be added to. This should not be set if `datacenter` is set. Conflicts with:
     * `cluster`.
     * 
     */
    @Export(name="cluster", type=String.class, parameters={})
    private Output</* @Nullable */ String> cluster;

    /**
     * @return The ID of the Compute Cluster this host should
     * be added to. This should not be set if `datacenter` is set. Conflicts with:
     * `cluster`.
     * 
     */
    public Output<Optional<String>> cluster() {
        return Codegen.optional(this.cluster);
    }
    /**
     * Can be set to `true` if compute cluster
     * membership will be managed through the `compute_cluster` resource rather
     * than the`host` resource. Conflicts with: `cluster`.
     * 
     */
    @Export(name="clusterManaged", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> clusterManaged;

    /**
     * @return Can be set to `true` if compute cluster
     * membership will be managed through the `compute_cluster` resource rather
     * than the`host` resource. Conflicts with: `cluster`.
     * 
     */
    public Output<Optional<Boolean>> clusterManaged() {
        return Codegen.optional(this.clusterManaged);
    }
    /**
     * If set to false then the host will be disconected.
     * Default is `false`.
     * 
     */
    @Export(name="connected", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> connected;

    /**
     * @return If set to false then the host will be disconected.
     * Default is `false`.
     * 
     */
    public Output<Optional<Boolean>> connected() {
        return Codegen.optional(this.connected);
    }
    /**
     * A map of custom attribute IDs and string
     * values to apply to the resource. Please refer to the
     * `vsphere_custom_attributes` resource for more information on applying
     * tags to resources.
     * 
     */
    @Export(name="customAttributes", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> customAttributes;

    /**
     * @return A map of custom attribute IDs and string
     * values to apply to the resource. Please refer to the
     * `vsphere_custom_attributes` resource for more information on applying
     * tags to resources.
     * 
     */
    public Output<Optional<Map<String,String>>> customAttributes() {
        return Codegen.optional(this.customAttributes);
    }
    /**
     * The ID of the datacenter this host should
     * be added to. This should not be set if `cluster` is set.
     * 
     */
    @Export(name="datacenter", type=String.class, parameters={})
    private Output</* @Nullable */ String> datacenter;

    /**
     * @return The ID of the datacenter this host should
     * be added to. This should not be set if `cluster` is set.
     * 
     */
    public Output<Optional<String>> datacenter() {
        return Codegen.optional(this.datacenter);
    }
    /**
     * If set to `true` then it will force the host to be added,
     * even if the host is already connected to a different vCenter Server instance.
     * Default is `false`.
     * 
     */
    @Export(name="force", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> force;

    /**
     * @return If set to `true` then it will force the host to be added,
     * even if the host is already connected to a different vCenter Server instance.
     * Default is `false`.
     * 
     */
    public Output<Optional<Boolean>> force() {
        return Codegen.optional(this.force);
    }
    /**
     * FQDN or IP address of the host to be added.
     * 
     */
    @Export(name="hostname", type=String.class, parameters={})
    private Output<String> hostname;

    /**
     * @return FQDN or IP address of the host to be added.
     * 
     */
    public Output<String> hostname() {
        return this.hostname;
    }
    /**
     * The license key that will be applied to the host.
     * The license key is expected to be present in vSphere.
     * 
     */
    @Export(name="license", type=String.class, parameters={})
    private Output</* @Nullable */ String> license;

    /**
     * @return The license key that will be applied to the host.
     * The license key is expected to be present in vSphere.
     * 
     */
    public Output<Optional<String>> license() {
        return Codegen.optional(this.license);
    }
    /**
     * Set the lockdown state of the host. Valid options are
     * `disabled`, `normal`, and `strict`. Default is `disabled`.
     * 
     */
    @Export(name="lockdown", type=String.class, parameters={})
    private Output</* @Nullable */ String> lockdown;

    /**
     * @return Set the lockdown state of the host. Valid options are
     * `disabled`, `normal`, and `strict`. Default is `disabled`.
     * 
     */
    public Output<Optional<String>> lockdown() {
        return Codegen.optional(this.lockdown);
    }
    /**
     * Set the management state of the host.
     * Default is `false`.
     * 
     */
    @Export(name="maintenance", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> maintenance;

    /**
     * @return Set the management state of the host.
     * Default is `false`.
     * 
     */
    public Output<Optional<Boolean>> maintenance() {
        return Codegen.optional(this.maintenance);
    }
    /**
     * Password that will be used by vSphere to authenticate
     * to the host.
     * 
     */
    @Export(name="password", type=String.class, parameters={})
    private Output<String> password;

    /**
     * @return Password that will be used by vSphere to authenticate
     * to the host.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * The IDs of any tags to attach to this resource. Please
     * refer to the `vsphere.Tag` resource for more information on applying
     * tags to resources.
     * 
     */
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return The IDs of any tags to attach to this resource. Please
     * refer to the `vsphere.Tag` resource for more information on applying
     * tags to resources.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Host&#39;s certificate SHA-1 thumbprint. If not set the
     * CA that signed the host&#39;s certificate should be trusted. If the CA is not
     * trusted and no thumbprint is set then the operation will fail.
     * 
     */
    @Export(name="thumbprint", type=String.class, parameters={})
    private Output</* @Nullable */ String> thumbprint;

    /**
     * @return Host&#39;s certificate SHA-1 thumbprint. If not set the
     * CA that signed the host&#39;s certificate should be trusted. If the CA is not
     * trusted and no thumbprint is set then the operation will fail.
     * 
     */
    public Output<Optional<String>> thumbprint() {
        return Codegen.optional(this.thumbprint);
    }
    /**
     * Username that will be used by vSphere to authenticate
     * to the host.
     * 
     */
    @Export(name="username", type=String.class, parameters={})
    private Output<String> username;

    /**
     * @return Username that will be used by vSphere to authenticate
     * to the host.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Host(String name) {
        this(name, HostArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Host(String name, HostArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Host(String name, HostArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/host:Host", name, args == null ? HostArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Host(String name, Output<String> id, @Nullable HostState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/host:Host", name, state, makeResourceOptions(options, id));
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
    public static Host get(String name, Output<String> id, @Nullable HostState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Host(name, id, state, options);
    }
}

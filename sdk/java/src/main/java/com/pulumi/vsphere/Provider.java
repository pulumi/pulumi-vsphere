// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.ProviderArgs;
import com.pulumi.vsphere.Utilities;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The provider type for the vsphere package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 * 
 */
@ResourceType(type="pulumi:providers:vsphere")
public class Provider extends com.pulumi.resources.ProviderResource {
    /**
     * govmomi debug path for debug
     * 
     */
    @Export(name="clientDebugPath", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientDebugPath;

    /**
     * @return govmomi debug path for debug
     * 
     */
    public Output<Optional<String>> clientDebugPath() {
        return Codegen.optional(this.clientDebugPath);
    }
    /**
     * govmomi debug path for a single run
     * 
     */
    @Export(name="clientDebugPathRun", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientDebugPathRun;

    /**
     * @return govmomi debug path for a single run
     * 
     */
    public Output<Optional<String>> clientDebugPathRun() {
        return Codegen.optional(this.clientDebugPathRun);
    }
    /**
     * The user password for vSphere API operations.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return The user password for vSphere API operations.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * The directory to save vSphere REST API sessions to
     * 
     */
    @Export(name="restSessionPath", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> restSessionPath;

    /**
     * @return The directory to save vSphere REST API sessions to
     * 
     */
    public Output<Optional<String>> restSessionPath() {
        return Codegen.optional(this.restSessionPath);
    }
    /**
     * The user name for vSphere API operations.
     * 
     */
    @Export(name="user", refs={String.class}, tree="[0]")
    private Output<String> user;

    /**
     * @return The user name for vSphere API operations.
     * 
     */
    public Output<String> user() {
        return this.user;
    }
    /**
     * @deprecated
     * This field has been renamed to vsphere_server.
     * 
     */
    @Deprecated /* This field has been renamed to vsphere_server. */
    @Export(name="vcenterServer", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vcenterServer;

    public Output<Optional<String>> vcenterServer() {
        return Codegen.optional(this.vcenterServer);
    }
    /**
     * The directory to save vSphere SOAP API sessions to
     * 
     */
    @Export(name="vimSessionPath", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vimSessionPath;

    /**
     * @return The directory to save vSphere SOAP API sessions to
     * 
     */
    public Output<Optional<String>> vimSessionPath() {
        return Codegen.optional(this.vimSessionPath);
    }
    /**
     * The vSphere Server name for vSphere API operations.
     * 
     */
    @Export(name="vsphereServer", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vsphereServer;

    /**
     * @return The vSphere Server name for vSphere API operations.
     * 
     */
    public Output<Optional<String>> vsphereServer() {
        return Codegen.optional(this.vsphereServer);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Provider(String name) {
        this(name, ProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Provider(String name, ProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Provider(String name, ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere", name, args == null ? ProviderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

}

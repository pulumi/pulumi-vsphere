// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.CustomAttributeArgs;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.inputs.CustomAttributeState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="vsphere:index/customAttribute:CustomAttribute")
public class CustomAttribute extends com.pulumi.resources.CustomResource {
    /**
     * The object type that this attribute may be
     * applied to. If not set, the custom attribute may be applied to any object
     * type. For a full list, review the Managed Object Types. Forces a new resource if changed.
     * 
     */
    @Export(name="managedObjectType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> managedObjectType;

    /**
     * @return The object type that this attribute may be
     * applied to. If not set, the custom attribute may be applied to any object
     * type. For a full list, review the Managed Object Types. Forces a new resource if changed.
     * 
     */
    public Output<Optional<String>> managedObjectType() {
        return Codegen.optional(this.managedObjectType);
    }
    /**
     * The name of the custom attribute.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the custom attribute.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CustomAttribute(java.lang.String name) {
        this(name, CustomAttributeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CustomAttribute(java.lang.String name, @Nullable CustomAttributeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CustomAttribute(java.lang.String name, @Nullable CustomAttributeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/customAttribute:CustomAttribute", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private CustomAttribute(java.lang.String name, Output<java.lang.String> id, @Nullable CustomAttributeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/customAttribute:CustomAttribute", name, state, makeResourceOptions(options, id), false);
    }

    private static CustomAttributeArgs makeArgs(@Nullable CustomAttributeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? CustomAttributeArgs.Empty : args;
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
    public static CustomAttribute get(java.lang.String name, Output<java.lang.String> id, @Nullable CustomAttributeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CustomAttribute(name, id, state, options);
    }
}

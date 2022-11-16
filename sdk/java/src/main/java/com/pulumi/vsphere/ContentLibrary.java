// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.ContentLibraryArgs;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.inputs.ContentLibraryState;
import com.pulumi.vsphere.outputs.ContentLibraryPublication;
import com.pulumi.vsphere.outputs.ContentLibrarySubscription;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="vsphere:index/contentLibrary:ContentLibrary")
public class ContentLibrary extends com.pulumi.resources.CustomResource {
    /**
     * A description for the content library.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return A description for the content library.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The name of the content library.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the content library.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Options to publish a local content library.
     * 
     */
    @Export(name="publication", type=ContentLibraryPublication.class, parameters={})
    private Output<ContentLibraryPublication> publication;

    /**
     * @return Options to publish a local content library.
     * 
     */
    public Output<ContentLibraryPublication> publication() {
        return this.publication;
    }
    /**
     * The managed object reference ID of the datastore on which to store the content library items.
     * 
     */
    @Export(name="storageBackings", type=List.class, parameters={String.class})
    private Output<List<String>> storageBackings;

    /**
     * @return The managed object reference ID of the datastore on which to store the content library items.
     * 
     */
    public Output<List<String>> storageBackings() {
        return this.storageBackings;
    }
    /**
     * Options subscribe to a published content library.
     * 
     */
    @Export(name="subscription", type=ContentLibrarySubscription.class, parameters={})
    private Output</* @Nullable */ ContentLibrarySubscription> subscription;

    /**
     * @return Options subscribe to a published content library.
     * 
     */
    public Output<Optional<ContentLibrarySubscription>> subscription() {
        return Codegen.optional(this.subscription);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ContentLibrary(String name) {
        this(name, ContentLibraryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ContentLibrary(String name, ContentLibraryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ContentLibrary(String name, ContentLibraryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/contentLibrary:ContentLibrary", name, args == null ? ContentLibraryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ContentLibrary(String name, Output<String> id, @Nullable ContentLibraryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/contentLibrary:ContentLibrary", name, state, makeResourceOptions(options, id));
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
    public static ContentLibrary get(String name, Output<String> id, @Nullable ContentLibraryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ContentLibrary(name, id, state, options);
    }
}
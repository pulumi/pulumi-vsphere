// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vsphere.ContentLibraryItemArgs;
import com.pulumi.vsphere.Utilities;
import com.pulumi.vsphere.inputs.ContentLibraryItemState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="vsphere:index/contentLibraryItem:ContentLibraryItem")
public class ContentLibraryItem extends com.pulumi.resources.CustomResource {
    /**
     * A description for the content library item.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description for the content library item.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * File to import as the content library item.
     * 
     */
    @Export(name="fileUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> fileUrl;

    /**
     * @return File to import as the content library item.
     * 
     */
    public Output<Optional<String>> fileUrl() {
        return Codegen.optional(this.fileUrl);
    }
    /**
     * The ID of the content library in which to create the item.
     * 
     */
    @Export(name="libraryId", refs={String.class}, tree="[0]")
    private Output<String> libraryId;

    /**
     * @return The ID of the content library in which to create the item.
     * 
     */
    public Output<String> libraryId() {
        return this.libraryId;
    }
    /**
     * The name of the item to be created in the content library.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the item to be created in the content library.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Virtual machine UUID to clone to content library.
     * 
     */
    @Export(name="sourceUuid", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceUuid;

    /**
     * @return Virtual machine UUID to clone to content library.
     * 
     */
    public Output<Optional<String>> sourceUuid() {
        return Codegen.optional(this.sourceUuid);
    }
    /**
     * Type of content library item.
     * One of &#34;ovf&#34;, &#34;iso&#34;, or &#34;vm-template&#34;. Default: `ovf`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return Type of content library item.
     * One of &#34;ovf&#34;, &#34;iso&#34;, or &#34;vm-template&#34;. Default: `ovf`.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ContentLibraryItem(String name) {
        this(name, ContentLibraryItemArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ContentLibraryItem(String name, ContentLibraryItemArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ContentLibraryItem(String name, ContentLibraryItemArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/contentLibraryItem:ContentLibraryItem", name, args == null ? ContentLibraryItemArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ContentLibraryItem(String name, Output<String> id, @Nullable ContentLibraryItemState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vsphere:index/contentLibraryItem:ContentLibraryItem", name, state, makeResourceOptions(options, id));
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
    public static ContentLibraryItem get(String name, Output<String> id, @Nullable ContentLibraryItemState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ContentLibraryItem(name, id, state, options);
    }
}

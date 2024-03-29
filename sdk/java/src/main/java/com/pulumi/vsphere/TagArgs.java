// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TagArgs extends com.pulumi.resources.ResourceArgs {

    public static final TagArgs Empty = new TagArgs();

    /**
     * The unique identifier of the parent category in
     * which this tag will be created. Forces a new resource if changed.
     * 
     */
    @Import(name="categoryId", required=true)
    private Output<String> categoryId;

    /**
     * @return The unique identifier of the parent category in
     * which this tag will be created. Forces a new resource if changed.
     * 
     */
    public Output<String> categoryId() {
        return this.categoryId;
    }

    /**
     * A description for the tag.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description for the tag.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The display name of the tag. The name must be unique
     * within its category.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The display name of the tag. The name must be unique
     * within its category.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private TagArgs() {}

    private TagArgs(TagArgs $) {
        this.categoryId = $.categoryId;
        this.description = $.description;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TagArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TagArgs $;

        public Builder() {
            $ = new TagArgs();
        }

        public Builder(TagArgs defaults) {
            $ = new TagArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param categoryId The unique identifier of the parent category in
         * which this tag will be created. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder categoryId(Output<String> categoryId) {
            $.categoryId = categoryId;
            return this;
        }

        /**
         * @param categoryId The unique identifier of the parent category in
         * which this tag will be created. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder categoryId(String categoryId) {
            return categoryId(Output.of(categoryId));
        }

        /**
         * @param description A description for the tag.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description for the tag.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name The display name of the tag. The name must be unique
         * within its category.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The display name of the tag. The name must be unique
         * within its category.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public TagArgs build() {
            if ($.categoryId == null) {
                throw new MissingRequiredPropertyException("TagArgs", "categoryId");
            }
            return $;
        }
    }

}

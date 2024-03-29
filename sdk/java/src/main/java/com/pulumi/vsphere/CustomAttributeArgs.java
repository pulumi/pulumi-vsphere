// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CustomAttributeArgs extends com.pulumi.resources.ResourceArgs {

    public static final CustomAttributeArgs Empty = new CustomAttributeArgs();

    /**
     * The object type that this attribute may be
     * applied to. If not set, the custom attribute may be applied to any object
     * type. For a full list, review the Managed Object Types. Forces a new resource if changed.
     * 
     */
    @Import(name="managedObjectType")
    private @Nullable Output<String> managedObjectType;

    /**
     * @return The object type that this attribute may be
     * applied to. If not set, the custom attribute may be applied to any object
     * type. For a full list, review the Managed Object Types. Forces a new resource if changed.
     * 
     */
    public Optional<Output<String>> managedObjectType() {
        return Optional.ofNullable(this.managedObjectType);
    }

    /**
     * The name of the custom attribute.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the custom attribute.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private CustomAttributeArgs() {}

    private CustomAttributeArgs(CustomAttributeArgs $) {
        this.managedObjectType = $.managedObjectType;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CustomAttributeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CustomAttributeArgs $;

        public Builder() {
            $ = new CustomAttributeArgs();
        }

        public Builder(CustomAttributeArgs defaults) {
            $ = new CustomAttributeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param managedObjectType The object type that this attribute may be
         * applied to. If not set, the custom attribute may be applied to any object
         * type. For a full list, review the Managed Object Types. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder managedObjectType(@Nullable Output<String> managedObjectType) {
            $.managedObjectType = managedObjectType;
            return this;
        }

        /**
         * @param managedObjectType The object type that this attribute may be
         * applied to. If not set, the custom attribute may be applied to any object
         * type. For a full list, review the Managed Object Types. Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder managedObjectType(String managedObjectType) {
            return managedObjectType(Output.of(managedObjectType));
        }

        /**
         * @param name The name of the custom attribute.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the custom attribute.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public CustomAttributeArgs build() {
            return $;
        }
    }

}

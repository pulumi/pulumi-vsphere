// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.vsphere.inputs.GuestOsCustomizationSpecArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GuestOsCustomizationArgs extends com.pulumi.resources.ResourceArgs {

    public static final GuestOsCustomizationArgs Empty = new GuestOsCustomizationArgs();

    /**
     * The description for the customization specification.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description for the customization specification.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name of the customization specification is the unique identifier per vCenter Server instance.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the customization specification is the unique identifier per vCenter Server instance.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Container object for the Guest OS properties about to be customized . See virtual machine customizations
     * 
     */
    @Import(name="spec", required=true)
    private Output<GuestOsCustomizationSpecArgs> spec;

    /**
     * @return Container object for the Guest OS properties about to be customized . See virtual machine customizations
     * 
     */
    public Output<GuestOsCustomizationSpecArgs> spec() {
        return this.spec;
    }

    /**
     * The type of customization specification: One among: Windows, Linux.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The type of customization specification: One among: Windows, Linux.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private GuestOsCustomizationArgs() {}

    private GuestOsCustomizationArgs(GuestOsCustomizationArgs $) {
        this.description = $.description;
        this.name = $.name;
        this.spec = $.spec;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GuestOsCustomizationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GuestOsCustomizationArgs $;

        public Builder() {
            $ = new GuestOsCustomizationArgs();
        }

        public Builder(GuestOsCustomizationArgs defaults) {
            $ = new GuestOsCustomizationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description for the customization specification.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description for the customization specification.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name The name of the customization specification is the unique identifier per vCenter Server instance.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the customization specification is the unique identifier per vCenter Server instance.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param spec Container object for the Guest OS properties about to be customized . See virtual machine customizations
         * 
         * @return builder
         * 
         */
        public Builder spec(Output<GuestOsCustomizationSpecArgs> spec) {
            $.spec = spec;
            return this;
        }

        /**
         * @param spec Container object for the Guest OS properties about to be customized . See virtual machine customizations
         * 
         * @return builder
         * 
         */
        public Builder spec(GuestOsCustomizationSpecArgs spec) {
            return spec(Output.of(spec));
        }

        /**
         * @param type The type of customization specification: One among: Windows, Linux.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of customization specification: One among: Windows, Linux.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public GuestOsCustomizationArgs build() {
            $.spec = Objects.requireNonNull($.spec, "expected parameter 'spec' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
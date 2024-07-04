// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetGuestOsCustomizationArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGuestOsCustomizationArgs Empty = new GetGuestOsCustomizationArgs();

    /**
     * The name of the customization specification is the unique
     * identifier per vCenter Server instance. ## Attribute Reference
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the customization specification is the unique
     * identifier per vCenter Server instance. ## Attribute Reference
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private GetGuestOsCustomizationArgs() {}

    private GetGuestOsCustomizationArgs(GetGuestOsCustomizationArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGuestOsCustomizationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGuestOsCustomizationArgs $;

        public Builder() {
            $ = new GetGuestOsCustomizationArgs();
        }

        public Builder(GetGuestOsCustomizationArgs defaults) {
            $ = new GetGuestOsCustomizationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the customization specification is the unique
         * identifier per vCenter Server instance. ## Attribute Reference
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the customization specification is the unique
         * identifier per vCenter Server instance. ## Attribute Reference
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetGuestOsCustomizationArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetGuestOsCustomizationArgs", "name");
            }
            return $;
        }
    }

}

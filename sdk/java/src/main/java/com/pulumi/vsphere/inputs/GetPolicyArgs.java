// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetPolicyArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPolicyArgs Empty = new GetPolicyArgs();

    /**
     * The name of the storage policy.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the storage policy.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private GetPolicyArgs() {}

    private GetPolicyArgs(GetPolicyArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPolicyArgs $;

        public Builder() {
            $ = new GetPolicyArgs();
        }

        public Builder(GetPolicyArgs defaults) {
            $ = new GetPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the storage policy.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the storage policy.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetPolicyArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetPolicyArgs", "name");
            }
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetVappContainerArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVappContainerArgs Empty = new GetVappContainerArgs();

    /**
     * The managed object reference ID
     * of the datacenter in which the vApp container is located.
     * 
     */
    @Import(name="datacenterId", required=true)
    private Output<String> datacenterId;

    /**
     * @return The managed object reference ID
     * of the datacenter in which the vApp container is located.
     * 
     */
    public Output<String> datacenterId() {
        return this.datacenterId;
    }

    /**
     * The name of the vApp container. This can be a name or
     * path.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the vApp container. This can be a name or
     * path.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private GetVappContainerArgs() {}

    private GetVappContainerArgs(GetVappContainerArgs $) {
        this.datacenterId = $.datacenterId;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVappContainerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVappContainerArgs $;

        public Builder() {
            $ = new GetVappContainerArgs();
        }

        public Builder(GetVappContainerArgs defaults) {
            $ = new GetVappContainerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param datacenterId The managed object reference ID
         * of the datacenter in which the vApp container is located.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(Output<String> datacenterId) {
            $.datacenterId = datacenterId;
            return this;
        }

        /**
         * @param datacenterId The managed object reference ID
         * of the datacenter in which the vApp container is located.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(String datacenterId) {
            return datacenterId(Output.of(datacenterId));
        }

        /**
         * @param name The name of the vApp container. This can be a name or
         * path.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the vApp container. This can be a name or
         * path.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetVappContainerArgs build() {
            if ($.datacenterId == null) {
                throw new MissingRequiredPropertyException("GetVappContainerArgs", "datacenterId");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetVappContainerArgs", "name");
            }
            return $;
        }
    }

}

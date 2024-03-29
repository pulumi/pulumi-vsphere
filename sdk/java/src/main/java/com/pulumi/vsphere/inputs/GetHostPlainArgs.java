// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetHostPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetHostPlainArgs Empty = new GetHostPlainArgs();

    /**
     * The managed object reference ID
     * of a vSphere datacenter object.
     * 
     */
    @Import(name="datacenterId", required=true)
    private String datacenterId;

    /**
     * @return The managed object reference ID
     * of a vSphere datacenter object.
     * 
     */
    public String datacenterId() {
        return this.datacenterId;
    }

    /**
     * The name of the ESXI host. This can be a name or path.
     * Can be omitted if there is only one host in your inventory.
     * 
     * &gt; **NOTE:** When used against an ESXi host directly, this data source _always_
     * returns the ESXi host&#39;s object ID, regardless of what is entered into `name`.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the ESXI host. This can be a name or path.
     * Can be omitted if there is only one host in your inventory.
     * 
     * &gt; **NOTE:** When used against an ESXi host directly, this data source _always_
     * returns the ESXi host&#39;s object ID, regardless of what is entered into `name`.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    private GetHostPlainArgs() {}

    private GetHostPlainArgs(GetHostPlainArgs $) {
        this.datacenterId = $.datacenterId;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetHostPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetHostPlainArgs $;

        public Builder() {
            $ = new GetHostPlainArgs();
        }

        public Builder(GetHostPlainArgs defaults) {
            $ = new GetHostPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param datacenterId The managed object reference ID
         * of a vSphere datacenter object.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(String datacenterId) {
            $.datacenterId = datacenterId;
            return this;
        }

        /**
         * @param name The name of the ESXI host. This can be a name or path.
         * Can be omitted if there is only one host in your inventory.
         * 
         * &gt; **NOTE:** When used against an ESXi host directly, this data source _always_
         * returns the ESXi host&#39;s object ID, regardless of what is entered into `name`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public GetHostPlainArgs build() {
            if ($.datacenterId == null) {
                throw new MissingRequiredPropertyException("GetHostPlainArgs", "datacenterId");
            }
            return $;
        }
    }

}

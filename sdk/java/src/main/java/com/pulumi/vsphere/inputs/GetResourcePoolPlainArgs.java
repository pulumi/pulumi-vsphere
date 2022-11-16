// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetResourcePoolPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetResourcePoolPlainArgs Empty = new GetResourcePoolPlainArgs();

    /**
     * The managed object reference ID
     * of the datacenter in which the resource pool is located. This can be omitted
     * if the search path used in `name` is an absolute path. For default
     * datacenters, use the id attribute from an empty `vsphere.Datacenter` data
     * source.
     * 
     */
    @Import(name="datacenterId")
    private @Nullable String datacenterId;

    /**
     * @return The managed object reference ID
     * of the datacenter in which the resource pool is located. This can be omitted
     * if the search path used in `name` is an absolute path. For default
     * datacenters, use the id attribute from an empty `vsphere.Datacenter` data
     * source.
     * 
     */
    public Optional<String> datacenterId() {
        return Optional.ofNullable(this.datacenterId);
    }

    /**
     * The name of the resource pool. This can be a name or
     * path. This is required when using vCenter.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the resource pool. This can be a name or
     * path. This is required when using vCenter.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    private GetResourcePoolPlainArgs() {}

    private GetResourcePoolPlainArgs(GetResourcePoolPlainArgs $) {
        this.datacenterId = $.datacenterId;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetResourcePoolPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetResourcePoolPlainArgs $;

        public Builder() {
            $ = new GetResourcePoolPlainArgs();
        }

        public Builder(GetResourcePoolPlainArgs defaults) {
            $ = new GetResourcePoolPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param datacenterId The managed object reference ID
         * of the datacenter in which the resource pool is located. This can be omitted
         * if the search path used in `name` is an absolute path. For default
         * datacenters, use the id attribute from an empty `vsphere.Datacenter` data
         * source.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(@Nullable String datacenterId) {
            $.datacenterId = datacenterId;
            return this;
        }

        /**
         * @param name The name of the resource pool. This can be a name or
         * path. This is required when using vCenter.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public GetResourcePoolPlainArgs build() {
            return $;
        }
    }

}
// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetResourcePoolArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetResourcePoolArgs Empty = new GetResourcePoolArgs();

    /**
     * The managed object reference ID
     * of the datacenter in which the resource pool is located. This can be omitted
     * if the search path used in `name` is an absolute path. For default
     * datacenters, use the id attribute from an empty `vsphere.Datacenter` data
     * source.
     * 
     * &gt; **Note:** When using ESXi without a vCenter Server instance, you do not
     * need to specify either attribute to use this data source. An empty declaration
     * will load the ESXi host&#39;s root resource pool.
     * 
     */
    @Import(name="datacenterId")
    private @Nullable Output<String> datacenterId;

    /**
     * @return The managed object reference ID
     * of the datacenter in which the resource pool is located. This can be omitted
     * if the search path used in `name` is an absolute path. For default
     * datacenters, use the id attribute from an empty `vsphere.Datacenter` data
     * source.
     * 
     * &gt; **Note:** When using ESXi without a vCenter Server instance, you do not
     * need to specify either attribute to use this data source. An empty declaration
     * will load the ESXi host&#39;s root resource pool.
     * 
     */
    public Optional<Output<String>> datacenterId() {
        return Optional.ofNullable(this.datacenterId);
    }

    /**
     * The name of the resource pool. This can be a name or
     * path. This is required when using vCenter.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the resource pool. This can be a name or
     * path. This is required when using vCenter.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private GetResourcePoolArgs() {}

    private GetResourcePoolArgs(GetResourcePoolArgs $) {
        this.datacenterId = $.datacenterId;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetResourcePoolArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetResourcePoolArgs $;

        public Builder() {
            $ = new GetResourcePoolArgs();
        }

        public Builder(GetResourcePoolArgs defaults) {
            $ = new GetResourcePoolArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param datacenterId The managed object reference ID
         * of the datacenter in which the resource pool is located. This can be omitted
         * if the search path used in `name` is an absolute path. For default
         * datacenters, use the id attribute from an empty `vsphere.Datacenter` data
         * source.
         * 
         * &gt; **Note:** When using ESXi without a vCenter Server instance, you do not
         * need to specify either attribute to use this data source. An empty declaration
         * will load the ESXi host&#39;s root resource pool.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(@Nullable Output<String> datacenterId) {
            $.datacenterId = datacenterId;
            return this;
        }

        /**
         * @param datacenterId The managed object reference ID
         * of the datacenter in which the resource pool is located. This can be omitted
         * if the search path used in `name` is an absolute path. For default
         * datacenters, use the id attribute from an empty `vsphere.Datacenter` data
         * source.
         * 
         * &gt; **Note:** When using ESXi without a vCenter Server instance, you do not
         * need to specify either attribute to use this data source. An empty declaration
         * will load the ESXi host&#39;s root resource pool.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(String datacenterId) {
            return datacenterId(Output.of(datacenterId));
        }

        /**
         * @param name The name of the resource pool. This can be a name or
         * path. This is required when using vCenter.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the resource pool. This can be a name or
         * path. This is required when using vCenter.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetResourcePoolArgs build() {
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDatastoreClusterPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDatastoreClusterPlainArgs Empty = new GetDatastoreClusterPlainArgs();

    /**
     * The managed object reference
     * ID of the datacenter the datastore cluster is located in.
     * This can be omitted if the search path used in `name` is an absolute path.
     * For default datacenters, use the id attribute from an empty
     * `vsphere.Datacenter` data source.
     * 
     */
    @Import(name="datacenterId")
    private @Nullable String datacenterId;

    /**
     * @return The managed object reference
     * ID of the datacenter the datastore cluster is located in.
     * This can be omitted if the search path used in `name` is an absolute path.
     * For default datacenters, use the id attribute from an empty
     * `vsphere.Datacenter` data source.
     * 
     */
    public Optional<String> datacenterId() {
        return Optional.ofNullable(this.datacenterId);
    }

    /**
     * The name or absolute path to the datastore cluster.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return The name or absolute path to the datastore cluster.
     * 
     */
    public String name() {
        return this.name;
    }

    private GetDatastoreClusterPlainArgs() {}

    private GetDatastoreClusterPlainArgs(GetDatastoreClusterPlainArgs $) {
        this.datacenterId = $.datacenterId;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDatastoreClusterPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDatastoreClusterPlainArgs $;

        public Builder() {
            $ = new GetDatastoreClusterPlainArgs();
        }

        public Builder(GetDatastoreClusterPlainArgs defaults) {
            $ = new GetDatastoreClusterPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param datacenterId The managed object reference
         * ID of the datacenter the datastore cluster is located in.
         * This can be omitted if the search path used in `name` is an absolute path.
         * For default datacenters, use the id attribute from an empty
         * `vsphere.Datacenter` data source.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(@Nullable String datacenterId) {
            $.datacenterId = datacenterId;
            return this;
        }

        /**
         * @param name The name or absolute path to the datastore cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public GetDatastoreClusterPlainArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetDatastoreClusterPlainArgs", "name");
            }
            return $;
        }
    }

}

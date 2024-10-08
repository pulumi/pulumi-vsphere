// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDatastoreArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDatastoreArgs Empty = new GetDatastoreArgs();

    /**
     * The managed object reference ID
     * of the datacenter the datastore is located in. This can be omitted if the
     * search path used in `name` is an absolute path. For default datacenters, use
     * the `id` attribute from an empty `vsphere.Datacenter` data source.
     * 
     */
    @Import(name="datacenterId")
    private @Nullable Output<String> datacenterId;

    /**
     * @return The managed object reference ID
     * of the datacenter the datastore is located in. This can be omitted if the
     * search path used in `name` is an absolute path. For default datacenters, use
     * the `id` attribute from an empty `vsphere.Datacenter` data source.
     * 
     */
    public Optional<Output<String>> datacenterId() {
        return Optional.ofNullable(this.datacenterId);
    }

    /**
     * The name of the datastore. This can be a name or path.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the datastore. This can be a name or path.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The disk space usage statistics for the specific datastore. The
     * total datastore capacity is represented as `capacity` and the free remaining
     * disk is represented as `free`.
     * 
     */
    @Import(name="stats")
    private @Nullable Output<Map<String,String>> stats;

    /**
     * @return The disk space usage statistics for the specific datastore. The
     * total datastore capacity is represented as `capacity` and the free remaining
     * disk is represented as `free`.
     * 
     */
    public Optional<Output<Map<String,String>>> stats() {
        return Optional.ofNullable(this.stats);
    }

    private GetDatastoreArgs() {}

    private GetDatastoreArgs(GetDatastoreArgs $) {
        this.datacenterId = $.datacenterId;
        this.name = $.name;
        this.stats = $.stats;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDatastoreArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDatastoreArgs $;

        public Builder() {
            $ = new GetDatastoreArgs();
        }

        public Builder(GetDatastoreArgs defaults) {
            $ = new GetDatastoreArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param datacenterId The managed object reference ID
         * of the datacenter the datastore is located in. This can be omitted if the
         * search path used in `name` is an absolute path. For default datacenters, use
         * the `id` attribute from an empty `vsphere.Datacenter` data source.
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
         * of the datacenter the datastore is located in. This can be omitted if the
         * search path used in `name` is an absolute path. For default datacenters, use
         * the `id` attribute from an empty `vsphere.Datacenter` data source.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(String datacenterId) {
            return datacenterId(Output.of(datacenterId));
        }

        /**
         * @param name The name of the datastore. This can be a name or path.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the datastore. This can be a name or path.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param stats The disk space usage statistics for the specific datastore. The
         * total datastore capacity is represented as `capacity` and the free remaining
         * disk is represented as `free`.
         * 
         * @return builder
         * 
         */
        public Builder stats(@Nullable Output<Map<String,String>> stats) {
            $.stats = stats;
            return this;
        }

        /**
         * @param stats The disk space usage statistics for the specific datastore. The
         * total datastore capacity is represented as `capacity` and the free remaining
         * disk is represented as `free`.
         * 
         * @return builder
         * 
         */
        public Builder stats(Map<String,String> stats) {
            return stats(Output.of(stats));
        }

        public GetDatastoreArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetDatastoreArgs", "name");
            }
            return $;
        }
    }

}

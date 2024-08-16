// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDatastorePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDatastorePlainArgs Empty = new GetDatastorePlainArgs();

    /**
     * The managed object reference ID
     * of the datacenter the datastore is located in. This can be omitted if the
     * search path used in `name` is an absolute path. For default datacenters, use
     * the `id` attribute from an empty `vsphere.Datacenter` data source.
     * 
     */
    @Import(name="datacenterId")
    private @Nullable String datacenterId;

    /**
     * @return The managed object reference ID
     * of the datacenter the datastore is located in. This can be omitted if the
     * search path used in `name` is an absolute path. For default datacenters, use
     * the `id` attribute from an empty `vsphere.Datacenter` data source.
     * 
     */
    public Optional<String> datacenterId() {
        return Optional.ofNullable(this.datacenterId);
    }

    /**
     * The name of the datastore. This can be a name or path.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return The name of the datastore. This can be a name or path.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * The disk space usage statistics for the specific datastore. The
     * total datastore capacity is represented as `capacity` and the free remaining
     * disk is represented as `free`.
     * 
     */
    @Import(name="stats")
    private @Nullable Map<String,String> stats;

    /**
     * @return The disk space usage statistics for the specific datastore. The
     * total datastore capacity is represented as `capacity` and the free remaining
     * disk is represented as `free`.
     * 
     */
    public Optional<Map<String,String>> stats() {
        return Optional.ofNullable(this.stats);
    }

    private GetDatastorePlainArgs() {}

    private GetDatastorePlainArgs(GetDatastorePlainArgs $) {
        this.datacenterId = $.datacenterId;
        this.name = $.name;
        this.stats = $.stats;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDatastorePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDatastorePlainArgs $;

        public Builder() {
            $ = new GetDatastorePlainArgs();
        }

        public Builder(GetDatastorePlainArgs defaults) {
            $ = new GetDatastorePlainArgs(Objects.requireNonNull(defaults));
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
        public Builder datacenterId(@Nullable String datacenterId) {
            $.datacenterId = datacenterId;
            return this;
        }

        /**
         * @param name The name of the datastore. This can be a name or path.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
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
        public Builder stats(@Nullable Map<String,String> stats) {
            $.stats = stats;
            return this;
        }

        public GetDatastorePlainArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetDatastorePlainArgs", "name");
            }
            return $;
        }
    }

}

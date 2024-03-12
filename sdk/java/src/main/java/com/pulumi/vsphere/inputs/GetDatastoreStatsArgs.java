// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDatastoreStatsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDatastoreStatsArgs Empty = new GetDatastoreStatsArgs();

    /**
     * A mapping of the capacity for all datastore in the datacenter
     * , where the name of the datastore is used as key and the capacity as value.
     * 
     */
    @Import(name="capacity")
    private @Nullable Output<Map<String,Object>> capacity;

    /**
     * @return A mapping of the capacity for all datastore in the datacenter
     * , where the name of the datastore is used as key and the capacity as value.
     * 
     */
    public Optional<Output<Map<String,Object>>> capacity() {
        return Optional.ofNullable(this.capacity);
    }

    /**
     * The [managed object reference ID][docs-about-morefs]
     * of the datacenter the datastores are located in. For default datacenters, use
     * the `id` attribute from an empty `vsphere.Datacenter` data source.
     * 
     * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
     * 
     */
    @Import(name="datacenterId", required=true)
    private Output<String> datacenterId;

    /**
     * @return The [managed object reference ID][docs-about-morefs]
     * of the datacenter the datastores are located in. For default datacenters, use
     * the `id` attribute from an empty `vsphere.Datacenter` data source.
     * 
     * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
     * 
     */
    public Output<String> datacenterId() {
        return this.datacenterId;
    }

    /**
     * A mapping of the free space for each datastore in the
     * datacenter, where the name of the datastore is used as key and the free
     * space as value.
     * 
     */
    @Import(name="freeSpace")
    private @Nullable Output<Map<String,Object>> freeSpace;

    /**
     * @return A mapping of the free space for each datastore in the
     * datacenter, where the name of the datastore is used as key and the free
     * space as value.
     * 
     */
    public Optional<Output<Map<String,Object>>> freeSpace() {
        return Optional.ofNullable(this.freeSpace);
    }

    private GetDatastoreStatsArgs() {}

    private GetDatastoreStatsArgs(GetDatastoreStatsArgs $) {
        this.capacity = $.capacity;
        this.datacenterId = $.datacenterId;
        this.freeSpace = $.freeSpace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDatastoreStatsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDatastoreStatsArgs $;

        public Builder() {
            $ = new GetDatastoreStatsArgs();
        }

        public Builder(GetDatastoreStatsArgs defaults) {
            $ = new GetDatastoreStatsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param capacity A mapping of the capacity for all datastore in the datacenter
         * , where the name of the datastore is used as key and the capacity as value.
         * 
         * @return builder
         * 
         */
        public Builder capacity(@Nullable Output<Map<String,Object>> capacity) {
            $.capacity = capacity;
            return this;
        }

        /**
         * @param capacity A mapping of the capacity for all datastore in the datacenter
         * , where the name of the datastore is used as key and the capacity as value.
         * 
         * @return builder
         * 
         */
        public Builder capacity(Map<String,Object> capacity) {
            return capacity(Output.of(capacity));
        }

        /**
         * @param datacenterId The [managed object reference ID][docs-about-morefs]
         * of the datacenter the datastores are located in. For default datacenters, use
         * the `id` attribute from an empty `vsphere.Datacenter` data source.
         * 
         * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(Output<String> datacenterId) {
            $.datacenterId = datacenterId;
            return this;
        }

        /**
         * @param datacenterId The [managed object reference ID][docs-about-morefs]
         * of the datacenter the datastores are located in. For default datacenters, use
         * the `id` attribute from an empty `vsphere.Datacenter` data source.
         * 
         * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(String datacenterId) {
            return datacenterId(Output.of(datacenterId));
        }

        /**
         * @param freeSpace A mapping of the free space for each datastore in the
         * datacenter, where the name of the datastore is used as key and the free
         * space as value.
         * 
         * @return builder
         * 
         */
        public Builder freeSpace(@Nullable Output<Map<String,Object>> freeSpace) {
            $.freeSpace = freeSpace;
            return this;
        }

        /**
         * @param freeSpace A mapping of the free space for each datastore in the
         * datacenter, where the name of the datastore is used as key and the free
         * space as value.
         * 
         * @return builder
         * 
         */
        public Builder freeSpace(Map<String,Object> freeSpace) {
            return freeSpace(Output.of(freeSpace));
        }

        public GetDatastoreStatsArgs build() {
            if ($.datacenterId == null) {
                throw new MissingRequiredPropertyException("GetDatastoreStatsArgs", "datacenterId");
            }
            return $;
        }
    }

}

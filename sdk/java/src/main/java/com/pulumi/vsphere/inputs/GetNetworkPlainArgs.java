// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.vsphere.inputs.GetNetworkFilter;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetNetworkPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNetworkPlainArgs Empty = new GetNetworkPlainArgs();

    /**
     * The managed object reference ID
     * of the datacenter the network is located in. This can be omitted if the
     * search path used in `name` is an absolute path. For default datacenters,
     * use the `id` attribute from an empty `vsphere.Datacenter` data source.
     * 
     */
    @Import(name="datacenterId")
    private @Nullable String datacenterId;

    /**
     * @return The managed object reference ID
     * of the datacenter the network is located in. This can be omitted if the
     * search path used in `name` is an absolute path. For default datacenters,
     * use the `id` attribute from an empty `vsphere.Datacenter` data source.
     * 
     */
    public Optional<String> datacenterId() {
        return Optional.ofNullable(this.datacenterId);
    }

    /**
     * For distributed port group type
     * network objects, the ID of the distributed virtual switch for which the port
     * group belongs. It is useful to differentiate port groups with same name using
     * the distributed virtual switch ID.
     * 
     */
    @Import(name="distributedVirtualSwitchUuid")
    private @Nullable String distributedVirtualSwitchUuid;

    /**
     * @return For distributed port group type
     * network objects, the ID of the distributed virtual switch for which the port
     * group belongs. It is useful to differentiate port groups with same name using
     * the distributed virtual switch ID.
     * 
     */
    public Optional<String> distributedVirtualSwitchUuid() {
        return Optional.ofNullable(this.distributedVirtualSwitchUuid);
    }

    /**
     * Apply a filter for the discovered network.
     * 
     */
    @Import(name="filters")
    private @Nullable List<GetNetworkFilter> filters;

    /**
     * @return Apply a filter for the discovered network.
     * 
     */
    public Optional<List<GetNetworkFilter>> filters() {
        return Optional.ofNullable(this.filters);
    }

    /**
     * The name of the network. This can be a name or path.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return The name of the network. This can be a name or path.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * The interval in milliseconds to retry the read operation if `retry_timeout` is set. Default: 500.
     * 
     */
    @Import(name="retryInterval")
    private @Nullable Integer retryInterval;

    /**
     * @return The interval in milliseconds to retry the read operation if `retry_timeout` is set. Default: 500.
     * 
     */
    public Optional<Integer> retryInterval() {
        return Optional.ofNullable(this.retryInterval);
    }

    /**
     * The timeout duration in seconds for the data source to retry read operations.
     * 
     */
    @Import(name="retryTimeout")
    private @Nullable Integer retryTimeout;

    /**
     * @return The timeout duration in seconds for the data source to retry read operations.
     * 
     */
    public Optional<Integer> retryTimeout() {
        return Optional.ofNullable(this.retryTimeout);
    }

    private GetNetworkPlainArgs() {}

    private GetNetworkPlainArgs(GetNetworkPlainArgs $) {
        this.datacenterId = $.datacenterId;
        this.distributedVirtualSwitchUuid = $.distributedVirtualSwitchUuid;
        this.filters = $.filters;
        this.name = $.name;
        this.retryInterval = $.retryInterval;
        this.retryTimeout = $.retryTimeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNetworkPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNetworkPlainArgs $;

        public Builder() {
            $ = new GetNetworkPlainArgs();
        }

        public Builder(GetNetworkPlainArgs defaults) {
            $ = new GetNetworkPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param datacenterId The managed object reference ID
         * of the datacenter the network is located in. This can be omitted if the
         * search path used in `name` is an absolute path. For default datacenters,
         * use the `id` attribute from an empty `vsphere.Datacenter` data source.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(@Nullable String datacenterId) {
            $.datacenterId = datacenterId;
            return this;
        }

        /**
         * @param distributedVirtualSwitchUuid For distributed port group type
         * network objects, the ID of the distributed virtual switch for which the port
         * group belongs. It is useful to differentiate port groups with same name using
         * the distributed virtual switch ID.
         * 
         * @return builder
         * 
         */
        public Builder distributedVirtualSwitchUuid(@Nullable String distributedVirtualSwitchUuid) {
            $.distributedVirtualSwitchUuid = distributedVirtualSwitchUuid;
            return this;
        }

        /**
         * @param filters Apply a filter for the discovered network.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable List<GetNetworkFilter> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters Apply a filter for the discovered network.
         * 
         * @return builder
         * 
         */
        public Builder filters(GetNetworkFilter... filters) {
            return filters(List.of(filters));
        }

        /**
         * @param name The name of the network. This can be a name or path.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param retryInterval The interval in milliseconds to retry the read operation if `retry_timeout` is set. Default: 500.
         * 
         * @return builder
         * 
         */
        public Builder retryInterval(@Nullable Integer retryInterval) {
            $.retryInterval = retryInterval;
            return this;
        }

        /**
         * @param retryTimeout The timeout duration in seconds for the data source to retry read operations.
         * 
         * @return builder
         * 
         */
        public Builder retryTimeout(@Nullable Integer retryTimeout) {
            $.retryTimeout = retryTimeout;
            return this;
        }

        public GetNetworkPlainArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetNetworkPlainArgs", "name");
            }
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetNetworkArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNetworkArgs Empty = new GetNetworkArgs();

    /**
     * The managed object reference ID
     * of the datacenter the network is located in. This can be omitted if the
     * search path used in `name` is an absolute path. For default datacenters,
     * use the `id` attribute from an empty `vsphere.Datacenter` data source.
     * 
     */
    @Import(name="datacenterId")
    private @Nullable Output<String> datacenterId;

    /**
     * @return The managed object reference ID
     * of the datacenter the network is located in. This can be omitted if the
     * search path used in `name` is an absolute path. For default datacenters,
     * use the `id` attribute from an empty `vsphere.Datacenter` data source.
     * 
     */
    public Optional<Output<String>> datacenterId() {
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
    private @Nullable Output<String> distributedVirtualSwitchUuid;

    /**
     * @return For distributed port group type
     * network objects, the ID of the distributed virtual switch for which the port
     * group belongs. It is useful to differentiate port groups with same name using
     * the distributed virtual switch ID.
     * 
     */
    public Optional<Output<String>> distributedVirtualSwitchUuid() {
        return Optional.ofNullable(this.distributedVirtualSwitchUuid);
    }

    /**
     * The name of the network. This can be a name or path.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the network. This can be a name or path.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private GetNetworkArgs() {}

    private GetNetworkArgs(GetNetworkArgs $) {
        this.datacenterId = $.datacenterId;
        this.distributedVirtualSwitchUuid = $.distributedVirtualSwitchUuid;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNetworkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNetworkArgs $;

        public Builder() {
            $ = new GetNetworkArgs();
        }

        public Builder(GetNetworkArgs defaults) {
            $ = new GetNetworkArgs(Objects.requireNonNull(defaults));
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
        public Builder datacenterId(@Nullable Output<String> datacenterId) {
            $.datacenterId = datacenterId;
            return this;
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
        public Builder datacenterId(String datacenterId) {
            return datacenterId(Output.of(datacenterId));
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
        public Builder distributedVirtualSwitchUuid(@Nullable Output<String> distributedVirtualSwitchUuid) {
            $.distributedVirtualSwitchUuid = distributedVirtualSwitchUuid;
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
        public Builder distributedVirtualSwitchUuid(String distributedVirtualSwitchUuid) {
            return distributedVirtualSwitchUuid(Output.of(distributedVirtualSwitchUuid));
        }

        /**
         * @param name The name of the network. This can be a name or path.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the network. This can be a name or path.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetNetworkArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetNetworkArgs", "name");
            }
            return $;
        }
    }

}

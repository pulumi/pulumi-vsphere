// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetNetworkFilter extends com.pulumi.resources.InvokeArgs {

    public static final GetNetworkFilter Empty = new GetNetworkFilter();

    /**
     * This is required if you have multiple port groups with the same name. This will be one of `DistributedVirtualPortgroup` for distributed port groups, `Network` for standard (host-based) port groups, or `OpaqueNetwork` for networks managed externally, such as those managed by NSX.
     * 
     */
    @Import(name="networkType")
    private @Nullable String networkType;

    /**
     * @return This is required if you have multiple port groups with the same name. This will be one of `DistributedVirtualPortgroup` for distributed port groups, `Network` for standard (host-based) port groups, or `OpaqueNetwork` for networks managed externally, such as those managed by NSX.
     * 
     */
    public Optional<String> networkType() {
        return Optional.ofNullable(this.networkType);
    }

    private GetNetworkFilter() {}

    private GetNetworkFilter(GetNetworkFilter $) {
        this.networkType = $.networkType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNetworkFilter defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNetworkFilter $;

        public Builder() {
            $ = new GetNetworkFilter();
        }

        public Builder(GetNetworkFilter defaults) {
            $ = new GetNetworkFilter(Objects.requireNonNull(defaults));
        }

        /**
         * @param networkType This is required if you have multiple port groups with the same name. This will be one of `DistributedVirtualPortgroup` for distributed port groups, `Network` for standard (host-based) port groups, or `OpaqueNetwork` for networks managed externally, such as those managed by NSX.
         * 
         * @return builder
         * 
         */
        public Builder networkType(@Nullable String networkType) {
            $.networkType = networkType;
            return this;
        }

        public GetNetworkFilter build() {
            return $;
        }
    }

}

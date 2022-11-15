// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class DistributedVirtualSwitchPvlanMapping {
    /**
     * @return The primary VLAN ID. The VLAN IDs of 0 and
     * 4095 are reserved and cannot be used in this property.
     * 
     */
    private Integer primaryVlanId;
    /**
     * @return The private VLAN type. Valid values are
     * promiscuous, community and isolated.
     * 
     */
    private String pvlanType;
    /**
     * @return The secondary VLAN ID. The VLAN IDs of 0
     * and 4095 are reserved and cannot be used in this property.
     * 
     */
    private Integer secondaryVlanId;

    private DistributedVirtualSwitchPvlanMapping() {}
    /**
     * @return The primary VLAN ID. The VLAN IDs of 0 and
     * 4095 are reserved and cannot be used in this property.
     * 
     */
    public Integer primaryVlanId() {
        return this.primaryVlanId;
    }
    /**
     * @return The private VLAN type. Valid values are
     * promiscuous, community and isolated.
     * 
     */
    public String pvlanType() {
        return this.pvlanType;
    }
    /**
     * @return The secondary VLAN ID. The VLAN IDs of 0
     * and 4095 are reserved and cannot be used in this property.
     * 
     */
    public Integer secondaryVlanId() {
        return this.secondaryVlanId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DistributedVirtualSwitchPvlanMapping defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer primaryVlanId;
        private String pvlanType;
        private Integer secondaryVlanId;
        public Builder() {}
        public Builder(DistributedVirtualSwitchPvlanMapping defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.primaryVlanId = defaults.primaryVlanId;
    	      this.pvlanType = defaults.pvlanType;
    	      this.secondaryVlanId = defaults.secondaryVlanId;
        }

        @CustomType.Setter
        public Builder primaryVlanId(Integer primaryVlanId) {
            this.primaryVlanId = Objects.requireNonNull(primaryVlanId);
            return this;
        }
        @CustomType.Setter
        public Builder pvlanType(String pvlanType) {
            this.pvlanType = Objects.requireNonNull(pvlanType);
            return this;
        }
        @CustomType.Setter
        public Builder secondaryVlanId(Integer secondaryVlanId) {
            this.secondaryVlanId = Objects.requireNonNull(secondaryVlanId);
            return this;
        }
        public DistributedVirtualSwitchPvlanMapping build() {
            final var o = new DistributedVirtualSwitchPvlanMapping();
            o.primaryVlanId = primaryVlanId;
            o.pvlanType = pvlanType;
            o.secondaryVlanId = secondaryVlanId;
            return o;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetVirtualMachineNetworkInterface {
    /**
     * @return The network interface types for each network interface found
     * on the virtual machine, in device bus order. Will be one of `e1000`, `e1000e` or
     * `vmxnet3`.
     * 
     */
    private String adapterType;
    /**
     * @return The upper bandwidth limit of this network interface,
     * in Mbits/sec.
     * 
     */
    private @Nullable Integer bandwidthLimit;
    /**
     * @return The bandwidth reservation of this network interface,
     * in Mbits/sec.
     * 
     */
    private @Nullable Integer bandwidthReservation;
    /**
     * @return The share count for this network interface when the
     * share level is custom.
     * 
     */
    private Integer bandwidthShareCount;
    /**
     * @return The bandwidth share allocation level for this interface.
     * Can be one of `low`, `normal`, `high`, or `custom`.
     * 
     */
    private @Nullable String bandwidthShareLevel;
    /**
     * @return The MAC address of this network interface.
     * 
     */
    private String macAddress;
    /**
     * @return The managed object reference ID of the network this interface is
     * connected to.
     * 
     */
    private String networkId;

    private GetVirtualMachineNetworkInterface() {}
    /**
     * @return The network interface types for each network interface found
     * on the virtual machine, in device bus order. Will be one of `e1000`, `e1000e` or
     * `vmxnet3`.
     * 
     */
    public String adapterType() {
        return this.adapterType;
    }
    /**
     * @return The upper bandwidth limit of this network interface,
     * in Mbits/sec.
     * 
     */
    public Optional<Integer> bandwidthLimit() {
        return Optional.ofNullable(this.bandwidthLimit);
    }
    /**
     * @return The bandwidth reservation of this network interface,
     * in Mbits/sec.
     * 
     */
    public Optional<Integer> bandwidthReservation() {
        return Optional.ofNullable(this.bandwidthReservation);
    }
    /**
     * @return The share count for this network interface when the
     * share level is custom.
     * 
     */
    public Integer bandwidthShareCount() {
        return this.bandwidthShareCount;
    }
    /**
     * @return The bandwidth share allocation level for this interface.
     * Can be one of `low`, `normal`, `high`, or `custom`.
     * 
     */
    public Optional<String> bandwidthShareLevel() {
        return Optional.ofNullable(this.bandwidthShareLevel);
    }
    /**
     * @return The MAC address of this network interface.
     * 
     */
    public String macAddress() {
        return this.macAddress;
    }
    /**
     * @return The managed object reference ID of the network this interface is
     * connected to.
     * 
     */
    public String networkId() {
        return this.networkId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVirtualMachineNetworkInterface defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String adapterType;
        private @Nullable Integer bandwidthLimit;
        private @Nullable Integer bandwidthReservation;
        private Integer bandwidthShareCount;
        private @Nullable String bandwidthShareLevel;
        private String macAddress;
        private String networkId;
        public Builder() {}
        public Builder(GetVirtualMachineNetworkInterface defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adapterType = defaults.adapterType;
    	      this.bandwidthLimit = defaults.bandwidthLimit;
    	      this.bandwidthReservation = defaults.bandwidthReservation;
    	      this.bandwidthShareCount = defaults.bandwidthShareCount;
    	      this.bandwidthShareLevel = defaults.bandwidthShareLevel;
    	      this.macAddress = defaults.macAddress;
    	      this.networkId = defaults.networkId;
        }

        @CustomType.Setter
        public Builder adapterType(String adapterType) {
            this.adapterType = Objects.requireNonNull(adapterType);
            return this;
        }
        @CustomType.Setter
        public Builder bandwidthLimit(@Nullable Integer bandwidthLimit) {
            this.bandwidthLimit = bandwidthLimit;
            return this;
        }
        @CustomType.Setter
        public Builder bandwidthReservation(@Nullable Integer bandwidthReservation) {
            this.bandwidthReservation = bandwidthReservation;
            return this;
        }
        @CustomType.Setter
        public Builder bandwidthShareCount(Integer bandwidthShareCount) {
            this.bandwidthShareCount = Objects.requireNonNull(bandwidthShareCount);
            return this;
        }
        @CustomType.Setter
        public Builder bandwidthShareLevel(@Nullable String bandwidthShareLevel) {
            this.bandwidthShareLevel = bandwidthShareLevel;
            return this;
        }
        @CustomType.Setter
        public Builder macAddress(String macAddress) {
            this.macAddress = Objects.requireNonNull(macAddress);
            return this;
        }
        @CustomType.Setter
        public Builder networkId(String networkId) {
            this.networkId = Objects.requireNonNull(networkId);
            return this;
        }
        public GetVirtualMachineNetworkInterface build() {
            final var o = new GetVirtualMachineNetworkInterface();
            o.adapterType = adapterType;
            o.bandwidthLimit = bandwidthLimit;
            o.bandwidthReservation = bandwidthReservation;
            o.bandwidthShareCount = bandwidthShareCount;
            o.bandwidthShareLevel = bandwidthShareLevel;
            o.macAddress = macAddress;
            o.networkId = networkId;
            return o;
        }
    }
}
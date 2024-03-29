// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VirtualMachineNetworkInterface {
    /**
     * @return The network interface type. One of `e1000`, `e1000e`, `sriov`, or `vmxnet3`. Default: `vmxnet3`.
     * 
     */
    private @Nullable String adapterType;
    /**
     * @return The upper bandwidth limit of the network interface, in Mbits/sec. The default is no limit. Ignored if `adapter_type` is set to `sriov`.
     * 
     */
    private @Nullable Integer bandwidthLimit;
    /**
     * @return The bandwidth reservation of the network interface, in Mbits/sec. The default is no reservation.
     * 
     */
    private @Nullable Integer bandwidthReservation;
    /**
     * @return The share count for the network interface when the share level is `custom`. Ignored if `adapter_type` is set to `sriov`.
     * 
     */
    private @Nullable Integer bandwidthShareCount;
    /**
     * @return The bandwidth share allocation level for the network interface. One of `low`, `normal`, `high`, or `custom`. Default: `normal`. Ignored if `adapter_type` is set to `sriov`.
     * 
     */
    private @Nullable String bandwidthShareLevel;
    /**
     * @return The internally-computed address of this device, such as scsi:0:1, denoting scsi bus #0 and device unit 1.
     * 
     */
    private @Nullable String deviceAddress;
    /**
     * @return The ID of the device within the virtual machine.
     * 
     */
    private @Nullable Integer key;
    /**
     * @return The MAC address of the network interface. Can only be manually set if `use_static_mac` is `true`. Otherwise, the value is computed and presents the assigned MAC address for the interface.
     * 
     */
    private @Nullable String macAddress;
    /**
     * @return The [managed object reference ID][docs-about-morefs] of the network on which to connect the virtual machine network interface.
     * 
     */
    private String networkId;
    /**
     * @return Specifies which NIC in an OVF/OVA the `network_interface` should be associated. Only applies at creation when deploying from an OVF/OVA.
     * 
     */
    private @Nullable String ovfMapping;
    /**
     * @return The ID of the Physical SR-IOV NIC to attach to, e.g. &#39;0000:d8:00.0&#39;
     * 
     */
    private @Nullable String physicalFunction;
    /**
     * @return If true, the `mac_address` field is treated as a static MAC address and set accordingly. Setting this to `true` requires `mac_address` to be set. Default: `false`.
     * 
     */
    private @Nullable Boolean useStaticMac;

    private VirtualMachineNetworkInterface() {}
    /**
     * @return The network interface type. One of `e1000`, `e1000e`, `sriov`, or `vmxnet3`. Default: `vmxnet3`.
     * 
     */
    public Optional<String> adapterType() {
        return Optional.ofNullable(this.adapterType);
    }
    /**
     * @return The upper bandwidth limit of the network interface, in Mbits/sec. The default is no limit. Ignored if `adapter_type` is set to `sriov`.
     * 
     */
    public Optional<Integer> bandwidthLimit() {
        return Optional.ofNullable(this.bandwidthLimit);
    }
    /**
     * @return The bandwidth reservation of the network interface, in Mbits/sec. The default is no reservation.
     * 
     */
    public Optional<Integer> bandwidthReservation() {
        return Optional.ofNullable(this.bandwidthReservation);
    }
    /**
     * @return The share count for the network interface when the share level is `custom`. Ignored if `adapter_type` is set to `sriov`.
     * 
     */
    public Optional<Integer> bandwidthShareCount() {
        return Optional.ofNullable(this.bandwidthShareCount);
    }
    /**
     * @return The bandwidth share allocation level for the network interface. One of `low`, `normal`, `high`, or `custom`. Default: `normal`. Ignored if `adapter_type` is set to `sriov`.
     * 
     */
    public Optional<String> bandwidthShareLevel() {
        return Optional.ofNullable(this.bandwidthShareLevel);
    }
    /**
     * @return The internally-computed address of this device, such as scsi:0:1, denoting scsi bus #0 and device unit 1.
     * 
     */
    public Optional<String> deviceAddress() {
        return Optional.ofNullable(this.deviceAddress);
    }
    /**
     * @return The ID of the device within the virtual machine.
     * 
     */
    public Optional<Integer> key() {
        return Optional.ofNullable(this.key);
    }
    /**
     * @return The MAC address of the network interface. Can only be manually set if `use_static_mac` is `true`. Otherwise, the value is computed and presents the assigned MAC address for the interface.
     * 
     */
    public Optional<String> macAddress() {
        return Optional.ofNullable(this.macAddress);
    }
    /**
     * @return The [managed object reference ID][docs-about-morefs] of the network on which to connect the virtual machine network interface.
     * 
     */
    public String networkId() {
        return this.networkId;
    }
    /**
     * @return Specifies which NIC in an OVF/OVA the `network_interface` should be associated. Only applies at creation when deploying from an OVF/OVA.
     * 
     */
    public Optional<String> ovfMapping() {
        return Optional.ofNullable(this.ovfMapping);
    }
    /**
     * @return The ID of the Physical SR-IOV NIC to attach to, e.g. &#39;0000:d8:00.0&#39;
     * 
     */
    public Optional<String> physicalFunction() {
        return Optional.ofNullable(this.physicalFunction);
    }
    /**
     * @return If true, the `mac_address` field is treated as a static MAC address and set accordingly. Setting this to `true` requires `mac_address` to be set. Default: `false`.
     * 
     */
    public Optional<Boolean> useStaticMac() {
        return Optional.ofNullable(this.useStaticMac);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VirtualMachineNetworkInterface defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String adapterType;
        private @Nullable Integer bandwidthLimit;
        private @Nullable Integer bandwidthReservation;
        private @Nullable Integer bandwidthShareCount;
        private @Nullable String bandwidthShareLevel;
        private @Nullable String deviceAddress;
        private @Nullable Integer key;
        private @Nullable String macAddress;
        private String networkId;
        private @Nullable String ovfMapping;
        private @Nullable String physicalFunction;
        private @Nullable Boolean useStaticMac;
        public Builder() {}
        public Builder(VirtualMachineNetworkInterface defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adapterType = defaults.adapterType;
    	      this.bandwidthLimit = defaults.bandwidthLimit;
    	      this.bandwidthReservation = defaults.bandwidthReservation;
    	      this.bandwidthShareCount = defaults.bandwidthShareCount;
    	      this.bandwidthShareLevel = defaults.bandwidthShareLevel;
    	      this.deviceAddress = defaults.deviceAddress;
    	      this.key = defaults.key;
    	      this.macAddress = defaults.macAddress;
    	      this.networkId = defaults.networkId;
    	      this.ovfMapping = defaults.ovfMapping;
    	      this.physicalFunction = defaults.physicalFunction;
    	      this.useStaticMac = defaults.useStaticMac;
        }

        @CustomType.Setter
        public Builder adapterType(@Nullable String adapterType) {

            this.adapterType = adapterType;
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
        public Builder bandwidthShareCount(@Nullable Integer bandwidthShareCount) {

            this.bandwidthShareCount = bandwidthShareCount;
            return this;
        }
        @CustomType.Setter
        public Builder bandwidthShareLevel(@Nullable String bandwidthShareLevel) {

            this.bandwidthShareLevel = bandwidthShareLevel;
            return this;
        }
        @CustomType.Setter
        public Builder deviceAddress(@Nullable String deviceAddress) {

            this.deviceAddress = deviceAddress;
            return this;
        }
        @CustomType.Setter
        public Builder key(@Nullable Integer key) {

            this.key = key;
            return this;
        }
        @CustomType.Setter
        public Builder macAddress(@Nullable String macAddress) {

            this.macAddress = macAddress;
            return this;
        }
        @CustomType.Setter
        public Builder networkId(String networkId) {
            if (networkId == null) {
              throw new MissingRequiredPropertyException("VirtualMachineNetworkInterface", "networkId");
            }
            this.networkId = networkId;
            return this;
        }
        @CustomType.Setter
        public Builder ovfMapping(@Nullable String ovfMapping) {

            this.ovfMapping = ovfMapping;
            return this;
        }
        @CustomType.Setter
        public Builder physicalFunction(@Nullable String physicalFunction) {

            this.physicalFunction = physicalFunction;
            return this;
        }
        @CustomType.Setter
        public Builder useStaticMac(@Nullable Boolean useStaticMac) {

            this.useStaticMac = useStaticMac;
            return this;
        }
        public VirtualMachineNetworkInterface build() {
            final var _resultValue = new VirtualMachineNetworkInterface();
            _resultValue.adapterType = adapterType;
            _resultValue.bandwidthLimit = bandwidthLimit;
            _resultValue.bandwidthReservation = bandwidthReservation;
            _resultValue.bandwidthShareCount = bandwidthShareCount;
            _resultValue.bandwidthShareLevel = bandwidthShareLevel;
            _resultValue.deviceAddress = deviceAddress;
            _resultValue.key = key;
            _resultValue.macAddress = macAddress;
            _resultValue.networkId = networkId;
            _resultValue.ovfMapping = ovfMapping;
            _resultValue.physicalFunction = physicalFunction;
            _resultValue.useStaticMac = useStaticMac;
            return _resultValue;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VirtualMachineNetworkInterfaceArgs extends com.pulumi.resources.ResourceArgs {

    public static final VirtualMachineNetworkInterfaceArgs Empty = new VirtualMachineNetworkInterfaceArgs();

    /**
     * The network interface type. One of `e1000`, `e1000e`, `sriov`, or `vmxnet3`. Default: `vmxnet3`.
     * 
     */
    @Import(name="adapterType")
    private @Nullable Output<String> adapterType;

    /**
     * @return The network interface type. One of `e1000`, `e1000e`, `sriov`, or `vmxnet3`. Default: `vmxnet3`.
     * 
     */
    public Optional<Output<String>> adapterType() {
        return Optional.ofNullable(this.adapterType);
    }

    /**
     * The upper bandwidth limit of the network interface, in Mbits/sec. The default is no limit. Ignored if `adapter_type` is set to `sriov`.
     * 
     */
    @Import(name="bandwidthLimit")
    private @Nullable Output<Integer> bandwidthLimit;

    /**
     * @return The upper bandwidth limit of the network interface, in Mbits/sec. The default is no limit. Ignored if `adapter_type` is set to `sriov`.
     * 
     */
    public Optional<Output<Integer>> bandwidthLimit() {
        return Optional.ofNullable(this.bandwidthLimit);
    }

    /**
     * The bandwidth reservation of the network interface, in Mbits/sec. The default is no reservation.
     * 
     */
    @Import(name="bandwidthReservation")
    private @Nullable Output<Integer> bandwidthReservation;

    /**
     * @return The bandwidth reservation of the network interface, in Mbits/sec. The default is no reservation.
     * 
     */
    public Optional<Output<Integer>> bandwidthReservation() {
        return Optional.ofNullable(this.bandwidthReservation);
    }

    /**
     * The share count for the network interface when the share level is `custom`. Ignored if `adapter_type` is set to `sriov`.
     * 
     */
    @Import(name="bandwidthShareCount")
    private @Nullable Output<Integer> bandwidthShareCount;

    /**
     * @return The share count for the network interface when the share level is `custom`. Ignored if `adapter_type` is set to `sriov`.
     * 
     */
    public Optional<Output<Integer>> bandwidthShareCount() {
        return Optional.ofNullable(this.bandwidthShareCount);
    }

    /**
     * The bandwidth share allocation level for the network interface. One of `low`, `normal`, `high`, or `custom`. Default: `normal`. Ignored if `adapter_type` is set to `sriov`.
     * 
     */
    @Import(name="bandwidthShareLevel")
    private @Nullable Output<String> bandwidthShareLevel;

    /**
     * @return The bandwidth share allocation level for the network interface. One of `low`, `normal`, `high`, or `custom`. Default: `normal`. Ignored if `adapter_type` is set to `sriov`.
     * 
     */
    public Optional<Output<String>> bandwidthShareLevel() {
        return Optional.ofNullable(this.bandwidthShareLevel);
    }

    /**
     * The internally-computed address of this device, such as scsi:0:1, denoting scsi bus #0 and device unit 1.
     * 
     */
    @Import(name="deviceAddress")
    private @Nullable Output<String> deviceAddress;

    /**
     * @return The internally-computed address of this device, such as scsi:0:1, denoting scsi bus #0 and device unit 1.
     * 
     */
    public Optional<Output<String>> deviceAddress() {
        return Optional.ofNullable(this.deviceAddress);
    }

    /**
     * The ID of the device within the virtual machine.
     * 
     */
    @Import(name="key")
    private @Nullable Output<Integer> key;

    /**
     * @return The ID of the device within the virtual machine.
     * 
     */
    public Optional<Output<Integer>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * The MAC address of the network interface. Can only be manually set if `use_static_mac` is `true`. Otherwise, the value is computed and presents the assigned MAC address for the interface.
     * 
     */
    @Import(name="macAddress")
    private @Nullable Output<String> macAddress;

    /**
     * @return The MAC address of the network interface. Can only be manually set if `use_static_mac` is `true`. Otherwise, the value is computed and presents the assigned MAC address for the interface.
     * 
     */
    public Optional<Output<String>> macAddress() {
        return Optional.ofNullable(this.macAddress);
    }

    /**
     * The [managed object reference ID][docs-about-morefs] of the network on which to connect the virtual machine network interface.
     * 
     */
    @Import(name="networkId", required=true)
    private Output<String> networkId;

    /**
     * @return The [managed object reference ID][docs-about-morefs] of the network on which to connect the virtual machine network interface.
     * 
     */
    public Output<String> networkId() {
        return this.networkId;
    }

    /**
     * Specifies which NIC in an OVF/OVA the `network_interface` should be associated. Only applies at creation when deploying from an OVF/OVA.
     * 
     */
    @Import(name="ovfMapping")
    private @Nullable Output<String> ovfMapping;

    /**
     * @return Specifies which NIC in an OVF/OVA the `network_interface` should be associated. Only applies at creation when deploying from an OVF/OVA.
     * 
     */
    public Optional<Output<String>> ovfMapping() {
        return Optional.ofNullable(this.ovfMapping);
    }

    /**
     * The ID of the Physical SR-IOV NIC to attach to, e.g. &#39;0000:d8:00.0&#39;
     * 
     */
    @Import(name="physicalFunction")
    private @Nullable Output<String> physicalFunction;

    /**
     * @return The ID of the Physical SR-IOV NIC to attach to, e.g. &#39;0000:d8:00.0&#39;
     * 
     */
    public Optional<Output<String>> physicalFunction() {
        return Optional.ofNullable(this.physicalFunction);
    }

    /**
     * If true, the `mac_address` field is treated as a static MAC address and set accordingly. Setting this to `true` requires `mac_address` to be set. Default: `false`.
     * 
     */
    @Import(name="useStaticMac")
    private @Nullable Output<Boolean> useStaticMac;

    /**
     * @return If true, the `mac_address` field is treated as a static MAC address and set accordingly. Setting this to `true` requires `mac_address` to be set. Default: `false`.
     * 
     */
    public Optional<Output<Boolean>> useStaticMac() {
        return Optional.ofNullable(this.useStaticMac);
    }

    private VirtualMachineNetworkInterfaceArgs() {}

    private VirtualMachineNetworkInterfaceArgs(VirtualMachineNetworkInterfaceArgs $) {
        this.adapterType = $.adapterType;
        this.bandwidthLimit = $.bandwidthLimit;
        this.bandwidthReservation = $.bandwidthReservation;
        this.bandwidthShareCount = $.bandwidthShareCount;
        this.bandwidthShareLevel = $.bandwidthShareLevel;
        this.deviceAddress = $.deviceAddress;
        this.key = $.key;
        this.macAddress = $.macAddress;
        this.networkId = $.networkId;
        this.ovfMapping = $.ovfMapping;
        this.physicalFunction = $.physicalFunction;
        this.useStaticMac = $.useStaticMac;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VirtualMachineNetworkInterfaceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VirtualMachineNetworkInterfaceArgs $;

        public Builder() {
            $ = new VirtualMachineNetworkInterfaceArgs();
        }

        public Builder(VirtualMachineNetworkInterfaceArgs defaults) {
            $ = new VirtualMachineNetworkInterfaceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param adapterType The network interface type. One of `e1000`, `e1000e`, `sriov`, or `vmxnet3`. Default: `vmxnet3`.
         * 
         * @return builder
         * 
         */
        public Builder adapterType(@Nullable Output<String> adapterType) {
            $.adapterType = adapterType;
            return this;
        }

        /**
         * @param adapterType The network interface type. One of `e1000`, `e1000e`, `sriov`, or `vmxnet3`. Default: `vmxnet3`.
         * 
         * @return builder
         * 
         */
        public Builder adapterType(String adapterType) {
            return adapterType(Output.of(adapterType));
        }

        /**
         * @param bandwidthLimit The upper bandwidth limit of the network interface, in Mbits/sec. The default is no limit. Ignored if `adapter_type` is set to `sriov`.
         * 
         * @return builder
         * 
         */
        public Builder bandwidthLimit(@Nullable Output<Integer> bandwidthLimit) {
            $.bandwidthLimit = bandwidthLimit;
            return this;
        }

        /**
         * @param bandwidthLimit The upper bandwidth limit of the network interface, in Mbits/sec. The default is no limit. Ignored if `adapter_type` is set to `sriov`.
         * 
         * @return builder
         * 
         */
        public Builder bandwidthLimit(Integer bandwidthLimit) {
            return bandwidthLimit(Output.of(bandwidthLimit));
        }

        /**
         * @param bandwidthReservation The bandwidth reservation of the network interface, in Mbits/sec. The default is no reservation.
         * 
         * @return builder
         * 
         */
        public Builder bandwidthReservation(@Nullable Output<Integer> bandwidthReservation) {
            $.bandwidthReservation = bandwidthReservation;
            return this;
        }

        /**
         * @param bandwidthReservation The bandwidth reservation of the network interface, in Mbits/sec. The default is no reservation.
         * 
         * @return builder
         * 
         */
        public Builder bandwidthReservation(Integer bandwidthReservation) {
            return bandwidthReservation(Output.of(bandwidthReservation));
        }

        /**
         * @param bandwidthShareCount The share count for the network interface when the share level is `custom`. Ignored if `adapter_type` is set to `sriov`.
         * 
         * @return builder
         * 
         */
        public Builder bandwidthShareCount(@Nullable Output<Integer> bandwidthShareCount) {
            $.bandwidthShareCount = bandwidthShareCount;
            return this;
        }

        /**
         * @param bandwidthShareCount The share count for the network interface when the share level is `custom`. Ignored if `adapter_type` is set to `sriov`.
         * 
         * @return builder
         * 
         */
        public Builder bandwidthShareCount(Integer bandwidthShareCount) {
            return bandwidthShareCount(Output.of(bandwidthShareCount));
        }

        /**
         * @param bandwidthShareLevel The bandwidth share allocation level for the network interface. One of `low`, `normal`, `high`, or `custom`. Default: `normal`. Ignored if `adapter_type` is set to `sriov`.
         * 
         * @return builder
         * 
         */
        public Builder bandwidthShareLevel(@Nullable Output<String> bandwidthShareLevel) {
            $.bandwidthShareLevel = bandwidthShareLevel;
            return this;
        }

        /**
         * @param bandwidthShareLevel The bandwidth share allocation level for the network interface. One of `low`, `normal`, `high`, or `custom`. Default: `normal`. Ignored if `adapter_type` is set to `sriov`.
         * 
         * @return builder
         * 
         */
        public Builder bandwidthShareLevel(String bandwidthShareLevel) {
            return bandwidthShareLevel(Output.of(bandwidthShareLevel));
        }

        /**
         * @param deviceAddress The internally-computed address of this device, such as scsi:0:1, denoting scsi bus #0 and device unit 1.
         * 
         * @return builder
         * 
         */
        public Builder deviceAddress(@Nullable Output<String> deviceAddress) {
            $.deviceAddress = deviceAddress;
            return this;
        }

        /**
         * @param deviceAddress The internally-computed address of this device, such as scsi:0:1, denoting scsi bus #0 and device unit 1.
         * 
         * @return builder
         * 
         */
        public Builder deviceAddress(String deviceAddress) {
            return deviceAddress(Output.of(deviceAddress));
        }

        /**
         * @param key The ID of the device within the virtual machine.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<Integer> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The ID of the device within the virtual machine.
         * 
         * @return builder
         * 
         */
        public Builder key(Integer key) {
            return key(Output.of(key));
        }

        /**
         * @param macAddress The MAC address of the network interface. Can only be manually set if `use_static_mac` is `true`. Otherwise, the value is computed and presents the assigned MAC address for the interface.
         * 
         * @return builder
         * 
         */
        public Builder macAddress(@Nullable Output<String> macAddress) {
            $.macAddress = macAddress;
            return this;
        }

        /**
         * @param macAddress The MAC address of the network interface. Can only be manually set if `use_static_mac` is `true`. Otherwise, the value is computed and presents the assigned MAC address for the interface.
         * 
         * @return builder
         * 
         */
        public Builder macAddress(String macAddress) {
            return macAddress(Output.of(macAddress));
        }

        /**
         * @param networkId The [managed object reference ID][docs-about-morefs] of the network on which to connect the virtual machine network interface.
         * 
         * @return builder
         * 
         */
        public Builder networkId(Output<String> networkId) {
            $.networkId = networkId;
            return this;
        }

        /**
         * @param networkId The [managed object reference ID][docs-about-morefs] of the network on which to connect the virtual machine network interface.
         * 
         * @return builder
         * 
         */
        public Builder networkId(String networkId) {
            return networkId(Output.of(networkId));
        }

        /**
         * @param ovfMapping Specifies which NIC in an OVF/OVA the `network_interface` should be associated. Only applies at creation when deploying from an OVF/OVA.
         * 
         * @return builder
         * 
         */
        public Builder ovfMapping(@Nullable Output<String> ovfMapping) {
            $.ovfMapping = ovfMapping;
            return this;
        }

        /**
         * @param ovfMapping Specifies which NIC in an OVF/OVA the `network_interface` should be associated. Only applies at creation when deploying from an OVF/OVA.
         * 
         * @return builder
         * 
         */
        public Builder ovfMapping(String ovfMapping) {
            return ovfMapping(Output.of(ovfMapping));
        }

        /**
         * @param physicalFunction The ID of the Physical SR-IOV NIC to attach to, e.g. &#39;0000:d8:00.0&#39;
         * 
         * @return builder
         * 
         */
        public Builder physicalFunction(@Nullable Output<String> physicalFunction) {
            $.physicalFunction = physicalFunction;
            return this;
        }

        /**
         * @param physicalFunction The ID of the Physical SR-IOV NIC to attach to, e.g. &#39;0000:d8:00.0&#39;
         * 
         * @return builder
         * 
         */
        public Builder physicalFunction(String physicalFunction) {
            return physicalFunction(Output.of(physicalFunction));
        }

        /**
         * @param useStaticMac If true, the `mac_address` field is treated as a static MAC address and set accordingly. Setting this to `true` requires `mac_address` to be set. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder useStaticMac(@Nullable Output<Boolean> useStaticMac) {
            $.useStaticMac = useStaticMac;
            return this;
        }

        /**
         * @param useStaticMac If true, the `mac_address` field is treated as a static MAC address and set accordingly. Setting this to `true` requires `mac_address` to be set. Default: `false`.
         * 
         * @return builder
         * 
         */
        public Builder useStaticMac(Boolean useStaticMac) {
            return useStaticMac(Output.of(useStaticMac));
        }

        public VirtualMachineNetworkInterfaceArgs build() {
            if ($.networkId == null) {
                throw new MissingRequiredPropertyException("VirtualMachineNetworkInterfaceArgs", "networkId");
            }
            return $;
        }
    }

}

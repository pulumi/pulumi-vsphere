// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VirtualMachineCdromArgs extends com.pulumi.resources.ResourceArgs {

    public static final VirtualMachineCdromArgs Empty = new VirtualMachineCdromArgs();

    /**
     * Indicates whether the device should be mapped to a remote client device
     * 
     */
    @Import(name="clientDevice")
    private @Nullable Output<Boolean> clientDevice;

    /**
     * @return Indicates whether the device should be mapped to a remote client device
     * 
     */
    public Optional<Output<Boolean>> clientDevice() {
        return Optional.ofNullable(this.clientDevice);
    }

    /**
     * The datastore ID the ISO is located on.
     * 
     */
    @Import(name="datastoreId")
    private @Nullable Output<String> datastoreId;

    /**
     * @return The datastore ID the ISO is located on.
     * 
     */
    public Optional<Output<String>> datastoreId() {
        return Optional.ofNullable(this.datastoreId);
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
     * The path to the ISO file on the datastore.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The path to the ISO file on the datastore.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    private VirtualMachineCdromArgs() {}

    private VirtualMachineCdromArgs(VirtualMachineCdromArgs $) {
        this.clientDevice = $.clientDevice;
        this.datastoreId = $.datastoreId;
        this.deviceAddress = $.deviceAddress;
        this.key = $.key;
        this.path = $.path;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VirtualMachineCdromArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VirtualMachineCdromArgs $;

        public Builder() {
            $ = new VirtualMachineCdromArgs();
        }

        public Builder(VirtualMachineCdromArgs defaults) {
            $ = new VirtualMachineCdromArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientDevice Indicates whether the device should be mapped to a remote client device
         * 
         * @return builder
         * 
         */
        public Builder clientDevice(@Nullable Output<Boolean> clientDevice) {
            $.clientDevice = clientDevice;
            return this;
        }

        /**
         * @param clientDevice Indicates whether the device should be mapped to a remote client device
         * 
         * @return builder
         * 
         */
        public Builder clientDevice(Boolean clientDevice) {
            return clientDevice(Output.of(clientDevice));
        }

        /**
         * @param datastoreId The datastore ID the ISO is located on.
         * 
         * @return builder
         * 
         */
        public Builder datastoreId(@Nullable Output<String> datastoreId) {
            $.datastoreId = datastoreId;
            return this;
        }

        /**
         * @param datastoreId The datastore ID the ISO is located on.
         * 
         * @return builder
         * 
         */
        public Builder datastoreId(String datastoreId) {
            return datastoreId(Output.of(datastoreId));
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
         * @param path The path to the ISO file on the datastore.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path to the ISO file on the datastore.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        public VirtualMachineCdromArgs build() {
            return $;
        }
    }

}

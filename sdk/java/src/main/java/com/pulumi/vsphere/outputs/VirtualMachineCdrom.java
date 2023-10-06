// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VirtualMachineCdrom {
    /**
     * @return Indicates whether the device should be backed by remote client device. Conflicts with `datastore_id` and `path`.
     * 
     */
    private @Nullable Boolean clientDevice;
    /**
     * @return The datastore ID that on which the ISO is located. Required for using a datastore ISO. Conflicts with `client_device`.
     * 
     */
    private @Nullable String datastoreId;
    private @Nullable String deviceAddress;
    /**
     * @return The ID of the device within the virtual machine.
     * 
     */
    private @Nullable Integer key;
    /**
     * @return The path to the ISO file. Required for using a datastore ISO. Conflicts with `client_device`.
     * 
     * &gt; **NOTE:** Either `client_device` (for a remote backed CD-ROM) or `datastore_id` and `path` (for a datastore ISO backed CD-ROM) are required to .
     * 
     * &gt; **NOTE:** Some CD-ROM drive types are not supported by this resource, such as pass-through devices. If these drives are present in a cloned template, or added outside of the provider, the desired state will be corrected to the defined device, or removed if no `cdrom` block is present.
     * 
     */
    private @Nullable String path;

    private VirtualMachineCdrom() {}
    /**
     * @return Indicates whether the device should be backed by remote client device. Conflicts with `datastore_id` and `path`.
     * 
     */
    public Optional<Boolean> clientDevice() {
        return Optional.ofNullable(this.clientDevice);
    }
    /**
     * @return The datastore ID that on which the ISO is located. Required for using a datastore ISO. Conflicts with `client_device`.
     * 
     */
    public Optional<String> datastoreId() {
        return Optional.ofNullable(this.datastoreId);
    }
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
     * @return The path to the ISO file. Required for using a datastore ISO. Conflicts with `client_device`.
     * 
     * &gt; **NOTE:** Either `client_device` (for a remote backed CD-ROM) or `datastore_id` and `path` (for a datastore ISO backed CD-ROM) are required to .
     * 
     * &gt; **NOTE:** Some CD-ROM drive types are not supported by this resource, such as pass-through devices. If these drives are present in a cloned template, or added outside of the provider, the desired state will be corrected to the defined device, or removed if no `cdrom` block is present.
     * 
     */
    public Optional<String> path() {
        return Optional.ofNullable(this.path);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VirtualMachineCdrom defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean clientDevice;
        private @Nullable String datastoreId;
        private @Nullable String deviceAddress;
        private @Nullable Integer key;
        private @Nullable String path;
        public Builder() {}
        public Builder(VirtualMachineCdrom defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clientDevice = defaults.clientDevice;
    	      this.datastoreId = defaults.datastoreId;
    	      this.deviceAddress = defaults.deviceAddress;
    	      this.key = defaults.key;
    	      this.path = defaults.path;
        }

        @CustomType.Setter
        public Builder clientDevice(@Nullable Boolean clientDevice) {
            this.clientDevice = clientDevice;
            return this;
        }
        @CustomType.Setter
        public Builder datastoreId(@Nullable String datastoreId) {
            this.datastoreId = datastoreId;
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
        public Builder path(@Nullable String path) {
            this.path = path;
            return this;
        }
        public VirtualMachineCdrom build() {
            final var o = new VirtualMachineCdrom();
            o.clientDevice = clientDevice;
            o.datastoreId = datastoreId;
            o.deviceAddress = deviceAddress;
            o.key = key;
            o.path = path;
            return o;
        }
    }
}

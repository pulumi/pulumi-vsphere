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
public final class VirtualMachineDisk {
    /**
     * @return Attach an external disk instead of creating a new one. Implies and conflicts with `keep_on_remove`. If set, you cannot set `size`, `eagerly_scrub`, or `thin_provisioned`. Must set `path` if used.
     * 
     */
    private @Nullable Boolean attach;
    /**
     * @return The type of storage controller to attach the  disk to. Can be `scsi`, `sata`, or `ide`. You must have the appropriate number of controllers enabled for the selected type. Default `scsi`.
     * 
     */
    private @Nullable String controllerType;
    /**
     * @return The managed object reference ID of the datastore in which to place the virtual machine. The virtual machine configuration files is placed here, along with any virtual disks that are created where a datastore is not explicitly specified. See the section on virtual machine migration for more information on modifying this value.
     * 
     */
    private @Nullable String datastoreId;
    private @Nullable String deviceAddress;
    /**
     * @return The mode of this this virtual disk for purposes of writes and snapshots. One of `append`, `independent_nonpersistent`, `independent_persistent`, `nonpersistent`, `persistent`, or `undoable`. Default: `persistent`. For more information on these option, please refer to the [product documentation][vmware-docs-disk-mode].
     * 
     */
    private @Nullable String diskMode;
    /**
     * @return The sharing mode of this virtual disk. One of `sharingMultiWriter` or `sharingNone`. Default: `sharingNone`.
     * 
     */
    private @Nullable String diskSharing;
    /**
     * @return If set to `true`, the disk space is zeroed out when the virtual machine is created. This will delay the creation of the virtual disk. Cannot be set to `true` when `thin_provisioned` is `true`.  See the section on picking a disk type for more information.  Default: `false`.
     * 
     */
    private @Nullable Boolean eagerlyScrub;
    /**
     * @return The upper limit of IOPS that this disk can use. The default is no limit.
     * 
     */
    private @Nullable Integer ioLimit;
    /**
     * @return The I/O reservation (guarantee) for the virtual disk has, in IOPS.  The default is no reservation.
     * 
     */
    private @Nullable Integer ioReservation;
    /**
     * @return The share count for the virtual disk when the share level is `custom`.
     * 
     */
    private @Nullable Integer ioShareCount;
    /**
     * @return The share allocation level for the virtual disk. One of `low`, `normal`, `high`, or `custom`. Default: `normal`.
     * 
     */
    private @Nullable String ioShareLevel;
    /**
     * @return Keep this disk when removing the device or destroying the virtual machine. Default: `false`.
     * 
     */
    private @Nullable Boolean keepOnRemove;
    /**
     * @return The ID of the device within the virtual machine.
     * 
     */
    private @Nullable Integer key;
    /**
     * @return A label for the virtual disk. Forces a new disk, if changed.
     * 
     */
    private String label;
    /**
     * @return When using `attach`, this parameter controls the path of a virtual disk to attach externally. Otherwise, it is a computed attribute that contains the virtual disk filename.
     * 
     */
    private @Nullable String path;
    /**
     * @return The size of the disk, in GB. Must be a whole number.
     * 
     */
    private @Nullable Integer size;
    /**
     * @return The ID of the storage policy to assign to the home directory of a virtual machine.
     * 
     */
    private @Nullable String storagePolicyId;
    /**
     * @return If `true`, the disk is thin provisioned, with space for the file being allocated on an as-needed basis. Cannot be set to `true` when `eagerly_scrub` is `true`. See the section on selecting a disk type for more information. Default: `true`.
     * 
     */
    private @Nullable Boolean thinProvisioned;
    /**
     * @return The disk number on the storage bus. The maximum value for this setting is the value of the controller count times the controller capacity (15 for SCSI, 30 for SATA, and 2 for IDE). Duplicate unit numbers are not allowed. Default `0`, for which one disk must be set to.
     * 
     */
    private @Nullable Integer unitNumber;
    /**
     * @return The UUID of the virtual disk VMDK file. This is used to track the virtual disk on the virtual machine.
     * 
     */
    private @Nullable String uuid;
    /**
     * @return If `true`, writes for this disk are sent directly to the filesystem immediately instead of being buffered. Default: `false`.
     * 
     */
    private @Nullable Boolean writeThrough;

    private VirtualMachineDisk() {}
    /**
     * @return Attach an external disk instead of creating a new one. Implies and conflicts with `keep_on_remove`. If set, you cannot set `size`, `eagerly_scrub`, or `thin_provisioned`. Must set `path` if used.
     * 
     */
    public Optional<Boolean> attach() {
        return Optional.ofNullable(this.attach);
    }
    /**
     * @return The type of storage controller to attach the  disk to. Can be `scsi`, `sata`, or `ide`. You must have the appropriate number of controllers enabled for the selected type. Default `scsi`.
     * 
     */
    public Optional<String> controllerType() {
        return Optional.ofNullable(this.controllerType);
    }
    /**
     * @return The managed object reference ID of the datastore in which to place the virtual machine. The virtual machine configuration files is placed here, along with any virtual disks that are created where a datastore is not explicitly specified. See the section on virtual machine migration for more information on modifying this value.
     * 
     */
    public Optional<String> datastoreId() {
        return Optional.ofNullable(this.datastoreId);
    }
    public Optional<String> deviceAddress() {
        return Optional.ofNullable(this.deviceAddress);
    }
    /**
     * @return The mode of this this virtual disk for purposes of writes and snapshots. One of `append`, `independent_nonpersistent`, `independent_persistent`, `nonpersistent`, `persistent`, or `undoable`. Default: `persistent`. For more information on these option, please refer to the [product documentation][vmware-docs-disk-mode].
     * 
     */
    public Optional<String> diskMode() {
        return Optional.ofNullable(this.diskMode);
    }
    /**
     * @return The sharing mode of this virtual disk. One of `sharingMultiWriter` or `sharingNone`. Default: `sharingNone`.
     * 
     */
    public Optional<String> diskSharing() {
        return Optional.ofNullable(this.diskSharing);
    }
    /**
     * @return If set to `true`, the disk space is zeroed out when the virtual machine is created. This will delay the creation of the virtual disk. Cannot be set to `true` when `thin_provisioned` is `true`.  See the section on picking a disk type for more information.  Default: `false`.
     * 
     */
    public Optional<Boolean> eagerlyScrub() {
        return Optional.ofNullable(this.eagerlyScrub);
    }
    /**
     * @return The upper limit of IOPS that this disk can use. The default is no limit.
     * 
     */
    public Optional<Integer> ioLimit() {
        return Optional.ofNullable(this.ioLimit);
    }
    /**
     * @return The I/O reservation (guarantee) for the virtual disk has, in IOPS.  The default is no reservation.
     * 
     */
    public Optional<Integer> ioReservation() {
        return Optional.ofNullable(this.ioReservation);
    }
    /**
     * @return The share count for the virtual disk when the share level is `custom`.
     * 
     */
    public Optional<Integer> ioShareCount() {
        return Optional.ofNullable(this.ioShareCount);
    }
    /**
     * @return The share allocation level for the virtual disk. One of `low`, `normal`, `high`, or `custom`. Default: `normal`.
     * 
     */
    public Optional<String> ioShareLevel() {
        return Optional.ofNullable(this.ioShareLevel);
    }
    /**
     * @return Keep this disk when removing the device or destroying the virtual machine. Default: `false`.
     * 
     */
    public Optional<Boolean> keepOnRemove() {
        return Optional.ofNullable(this.keepOnRemove);
    }
    /**
     * @return The ID of the device within the virtual machine.
     * 
     */
    public Optional<Integer> key() {
        return Optional.ofNullable(this.key);
    }
    /**
     * @return A label for the virtual disk. Forces a new disk, if changed.
     * 
     */
    public String label() {
        return this.label;
    }
    /**
     * @return When using `attach`, this parameter controls the path of a virtual disk to attach externally. Otherwise, it is a computed attribute that contains the virtual disk filename.
     * 
     */
    public Optional<String> path() {
        return Optional.ofNullable(this.path);
    }
    /**
     * @return The size of the disk, in GB. Must be a whole number.
     * 
     */
    public Optional<Integer> size() {
        return Optional.ofNullable(this.size);
    }
    /**
     * @return The ID of the storage policy to assign to the home directory of a virtual machine.
     * 
     */
    public Optional<String> storagePolicyId() {
        return Optional.ofNullable(this.storagePolicyId);
    }
    /**
     * @return If `true`, the disk is thin provisioned, with space for the file being allocated on an as-needed basis. Cannot be set to `true` when `eagerly_scrub` is `true`. See the section on selecting a disk type for more information. Default: `true`.
     * 
     */
    public Optional<Boolean> thinProvisioned() {
        return Optional.ofNullable(this.thinProvisioned);
    }
    /**
     * @return The disk number on the storage bus. The maximum value for this setting is the value of the controller count times the controller capacity (15 for SCSI, 30 for SATA, and 2 for IDE). Duplicate unit numbers are not allowed. Default `0`, for which one disk must be set to.
     * 
     */
    public Optional<Integer> unitNumber() {
        return Optional.ofNullable(this.unitNumber);
    }
    /**
     * @return The UUID of the virtual disk VMDK file. This is used to track the virtual disk on the virtual machine.
     * 
     */
    public Optional<String> uuid() {
        return Optional.ofNullable(this.uuid);
    }
    /**
     * @return If `true`, writes for this disk are sent directly to the filesystem immediately instead of being buffered. Default: `false`.
     * 
     */
    public Optional<Boolean> writeThrough() {
        return Optional.ofNullable(this.writeThrough);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VirtualMachineDisk defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean attach;
        private @Nullable String controllerType;
        private @Nullable String datastoreId;
        private @Nullable String deviceAddress;
        private @Nullable String diskMode;
        private @Nullable String diskSharing;
        private @Nullable Boolean eagerlyScrub;
        private @Nullable Integer ioLimit;
        private @Nullable Integer ioReservation;
        private @Nullable Integer ioShareCount;
        private @Nullable String ioShareLevel;
        private @Nullable Boolean keepOnRemove;
        private @Nullable Integer key;
        private String label;
        private @Nullable String path;
        private @Nullable Integer size;
        private @Nullable String storagePolicyId;
        private @Nullable Boolean thinProvisioned;
        private @Nullable Integer unitNumber;
        private @Nullable String uuid;
        private @Nullable Boolean writeThrough;
        public Builder() {}
        public Builder(VirtualMachineDisk defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.attach = defaults.attach;
    	      this.controllerType = defaults.controllerType;
    	      this.datastoreId = defaults.datastoreId;
    	      this.deviceAddress = defaults.deviceAddress;
    	      this.diskMode = defaults.diskMode;
    	      this.diskSharing = defaults.diskSharing;
    	      this.eagerlyScrub = defaults.eagerlyScrub;
    	      this.ioLimit = defaults.ioLimit;
    	      this.ioReservation = defaults.ioReservation;
    	      this.ioShareCount = defaults.ioShareCount;
    	      this.ioShareLevel = defaults.ioShareLevel;
    	      this.keepOnRemove = defaults.keepOnRemove;
    	      this.key = defaults.key;
    	      this.label = defaults.label;
    	      this.path = defaults.path;
    	      this.size = defaults.size;
    	      this.storagePolicyId = defaults.storagePolicyId;
    	      this.thinProvisioned = defaults.thinProvisioned;
    	      this.unitNumber = defaults.unitNumber;
    	      this.uuid = defaults.uuid;
    	      this.writeThrough = defaults.writeThrough;
        }

        @CustomType.Setter
        public Builder attach(@Nullable Boolean attach) {
            this.attach = attach;
            return this;
        }
        @CustomType.Setter
        public Builder controllerType(@Nullable String controllerType) {
            this.controllerType = controllerType;
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
        public Builder diskMode(@Nullable String diskMode) {
            this.diskMode = diskMode;
            return this;
        }
        @CustomType.Setter
        public Builder diskSharing(@Nullable String diskSharing) {
            this.diskSharing = diskSharing;
            return this;
        }
        @CustomType.Setter
        public Builder eagerlyScrub(@Nullable Boolean eagerlyScrub) {
            this.eagerlyScrub = eagerlyScrub;
            return this;
        }
        @CustomType.Setter
        public Builder ioLimit(@Nullable Integer ioLimit) {
            this.ioLimit = ioLimit;
            return this;
        }
        @CustomType.Setter
        public Builder ioReservation(@Nullable Integer ioReservation) {
            this.ioReservation = ioReservation;
            return this;
        }
        @CustomType.Setter
        public Builder ioShareCount(@Nullable Integer ioShareCount) {
            this.ioShareCount = ioShareCount;
            return this;
        }
        @CustomType.Setter
        public Builder ioShareLevel(@Nullable String ioShareLevel) {
            this.ioShareLevel = ioShareLevel;
            return this;
        }
        @CustomType.Setter
        public Builder keepOnRemove(@Nullable Boolean keepOnRemove) {
            this.keepOnRemove = keepOnRemove;
            return this;
        }
        @CustomType.Setter
        public Builder key(@Nullable Integer key) {
            this.key = key;
            return this;
        }
        @CustomType.Setter
        public Builder label(String label) {
            this.label = Objects.requireNonNull(label);
            return this;
        }
        @CustomType.Setter
        public Builder path(@Nullable String path) {
            this.path = path;
            return this;
        }
        @CustomType.Setter
        public Builder size(@Nullable Integer size) {
            this.size = size;
            return this;
        }
        @CustomType.Setter
        public Builder storagePolicyId(@Nullable String storagePolicyId) {
            this.storagePolicyId = storagePolicyId;
            return this;
        }
        @CustomType.Setter
        public Builder thinProvisioned(@Nullable Boolean thinProvisioned) {
            this.thinProvisioned = thinProvisioned;
            return this;
        }
        @CustomType.Setter
        public Builder unitNumber(@Nullable Integer unitNumber) {
            this.unitNumber = unitNumber;
            return this;
        }
        @CustomType.Setter
        public Builder uuid(@Nullable String uuid) {
            this.uuid = uuid;
            return this;
        }
        @CustomType.Setter
        public Builder writeThrough(@Nullable Boolean writeThrough) {
            this.writeThrough = writeThrough;
            return this;
        }
        public VirtualMachineDisk build() {
            final var o = new VirtualMachineDisk();
            o.attach = attach;
            o.controllerType = controllerType;
            o.datastoreId = datastoreId;
            o.deviceAddress = deviceAddress;
            o.diskMode = diskMode;
            o.diskSharing = diskSharing;
            o.eagerlyScrub = eagerlyScrub;
            o.ioLimit = ioLimit;
            o.ioReservation = ioReservation;
            o.ioShareCount = ioShareCount;
            o.ioShareLevel = ioShareLevel;
            o.keepOnRemove = keepOnRemove;
            o.key = key;
            o.label = label;
            o.path = path;
            o.size = size;
            o.storagePolicyId = storagePolicyId;
            o.thinProvisioned = thinProvisioned;
            o.unitNumber = unitNumber;
            o.uuid = uuid;
            o.writeThrough = writeThrough;
            return o;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VirtualDiskArgs extends com.pulumi.resources.ResourceArgs {

    public static final VirtualDiskArgs Empty = new VirtualDiskArgs();

    /**
     * The adapter type for this virtual disk. Can be
     * one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
     * 
     * &gt; **NOTE:** `adapter_type` is **deprecated**: it does not dictate the type of
     * controller that the virtual disk will be attached to on the virtual machine.
     * Please see the `scsi_type` parameter
     * in the `vsphere.VirtualMachine` resource for information on how to control
     * disk controller types. This parameter will be removed in future versions of the
     * vSphere provider.
     * 
     * @deprecated
     * this attribute has no effect on controller types - please use scsi_type in vsphere_virtual_machine instead
     * 
     */
    @Deprecated /* this attribute has no effect on controller types - please use scsi_type in vsphere_virtual_machine instead */
    @Import(name="adapterType")
    private @Nullable Output<String> adapterType;

    /**
     * @return The adapter type for this virtual disk. Can be
     * one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
     * 
     * &gt; **NOTE:** `adapter_type` is **deprecated**: it does not dictate the type of
     * controller that the virtual disk will be attached to on the virtual machine.
     * Please see the `scsi_type` parameter
     * in the `vsphere.VirtualMachine` resource for information on how to control
     * disk controller types. This parameter will be removed in future versions of the
     * vSphere provider.
     * 
     * @deprecated
     * this attribute has no effect on controller types - please use scsi_type in vsphere_virtual_machine instead
     * 
     */
    @Deprecated /* this attribute has no effect on controller types - please use scsi_type in vsphere_virtual_machine instead */
    public Optional<Output<String>> adapterType() {
        return Optional.ofNullable(this.adapterType);
    }

    /**
     * Tells the resource to create any
     * directories that are a part of the `vmdk_path` parameter if they are missing.
     * Default: `false`.
     * 
     * &gt; **NOTE:** Any directory created as part of the operation when
     * `create_directories` is enabled will not be deleted when the resource is
     * destroyed.
     * 
     */
    @Import(name="createDirectories")
    private @Nullable Output<Boolean> createDirectories;

    /**
     * @return Tells the resource to create any
     * directories that are a part of the `vmdk_path` parameter if they are missing.
     * Default: `false`.
     * 
     * &gt; **NOTE:** Any directory created as part of the operation when
     * `create_directories` is enabled will not be deleted when the resource is
     * destroyed.
     * 
     */
    public Optional<Output<Boolean>> createDirectories() {
        return Optional.ofNullable(this.createDirectories);
    }

    /**
     * The name of the datacenter in which to create the
     * disk. Can be omitted when when ESXi or if there is only one datacenter in
     * your infrastructure.
     * 
     */
    @Import(name="datacenter")
    private @Nullable Output<String> datacenter;

    /**
     * @return The name of the datacenter in which to create the
     * disk. Can be omitted when when ESXi or if there is only one datacenter in
     * your infrastructure.
     * 
     */
    public Optional<Output<String>> datacenter() {
        return Optional.ofNullable(this.datacenter);
    }

    /**
     * The name of the datastore in which to create the
     * disk.
     * 
     */
    @Import(name="datastore", required=true)
    private Output<String> datastore;

    /**
     * @return The name of the datastore in which to create the
     * disk.
     * 
     */
    public Output<String> datastore() {
        return this.datastore;
    }

    /**
     * Size of the disk (in GB).
     * 
     */
    @Import(name="size", required=true)
    private Output<Integer> size;

    /**
     * @return Size of the disk (in GB).
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }

    /**
     * The type of disk to create. Can be one of
     * `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
     * information on what each kind of disk provisioning policy means, click
     * [here][docs-vmware-vm-disk-provisioning].
     * 
     * [docs-vmware-vm-disk-provisioning]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-4C0F4D73-82F2-4B81-8AA7-1DD752A8A5AC.html
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The type of disk to create. Can be one of
     * `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
     * information on what each kind of disk provisioning policy means, click
     * [here][docs-vmware-vm-disk-provisioning].
     * 
     * [docs-vmware-vm-disk-provisioning]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-4C0F4D73-82F2-4B81-8AA7-1DD752A8A5AC.html
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * The path, including filename, of the virtual disk to
     * be created.  This needs to end in `.vmdk`.
     * 
     */
    @Import(name="vmdkPath", required=true)
    private Output<String> vmdkPath;

    /**
     * @return The path, including filename, of the virtual disk to
     * be created.  This needs to end in `.vmdk`.
     * 
     */
    public Output<String> vmdkPath() {
        return this.vmdkPath;
    }

    private VirtualDiskArgs() {}

    private VirtualDiskArgs(VirtualDiskArgs $) {
        this.adapterType = $.adapterType;
        this.createDirectories = $.createDirectories;
        this.datacenter = $.datacenter;
        this.datastore = $.datastore;
        this.size = $.size;
        this.type = $.type;
        this.vmdkPath = $.vmdkPath;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VirtualDiskArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VirtualDiskArgs $;

        public Builder() {
            $ = new VirtualDiskArgs();
        }

        public Builder(VirtualDiskArgs defaults) {
            $ = new VirtualDiskArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param adapterType The adapter type for this virtual disk. Can be
         * one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
         * 
         * &gt; **NOTE:** `adapter_type` is **deprecated**: it does not dictate the type of
         * controller that the virtual disk will be attached to on the virtual machine.
         * Please see the `scsi_type` parameter
         * in the `vsphere.VirtualMachine` resource for information on how to control
         * disk controller types. This parameter will be removed in future versions of the
         * vSphere provider.
         * 
         * @return builder
         * 
         * @deprecated
         * this attribute has no effect on controller types - please use scsi_type in vsphere_virtual_machine instead
         * 
         */
        @Deprecated /* this attribute has no effect on controller types - please use scsi_type in vsphere_virtual_machine instead */
        public Builder adapterType(@Nullable Output<String> adapterType) {
            $.adapterType = adapterType;
            return this;
        }

        /**
         * @param adapterType The adapter type for this virtual disk. Can be
         * one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
         * 
         * &gt; **NOTE:** `adapter_type` is **deprecated**: it does not dictate the type of
         * controller that the virtual disk will be attached to on the virtual machine.
         * Please see the `scsi_type` parameter
         * in the `vsphere.VirtualMachine` resource for information on how to control
         * disk controller types. This parameter will be removed in future versions of the
         * vSphere provider.
         * 
         * @return builder
         * 
         * @deprecated
         * this attribute has no effect on controller types - please use scsi_type in vsphere_virtual_machine instead
         * 
         */
        @Deprecated /* this attribute has no effect on controller types - please use scsi_type in vsphere_virtual_machine instead */
        public Builder adapterType(String adapterType) {
            return adapterType(Output.of(adapterType));
        }

        /**
         * @param createDirectories Tells the resource to create any
         * directories that are a part of the `vmdk_path` parameter if they are missing.
         * Default: `false`.
         * 
         * &gt; **NOTE:** Any directory created as part of the operation when
         * `create_directories` is enabled will not be deleted when the resource is
         * destroyed.
         * 
         * @return builder
         * 
         */
        public Builder createDirectories(@Nullable Output<Boolean> createDirectories) {
            $.createDirectories = createDirectories;
            return this;
        }

        /**
         * @param createDirectories Tells the resource to create any
         * directories that are a part of the `vmdk_path` parameter if they are missing.
         * Default: `false`.
         * 
         * &gt; **NOTE:** Any directory created as part of the operation when
         * `create_directories` is enabled will not be deleted when the resource is
         * destroyed.
         * 
         * @return builder
         * 
         */
        public Builder createDirectories(Boolean createDirectories) {
            return createDirectories(Output.of(createDirectories));
        }

        /**
         * @param datacenter The name of the datacenter in which to create the
         * disk. Can be omitted when when ESXi or if there is only one datacenter in
         * your infrastructure.
         * 
         * @return builder
         * 
         */
        public Builder datacenter(@Nullable Output<String> datacenter) {
            $.datacenter = datacenter;
            return this;
        }

        /**
         * @param datacenter The name of the datacenter in which to create the
         * disk. Can be omitted when when ESXi or if there is only one datacenter in
         * your infrastructure.
         * 
         * @return builder
         * 
         */
        public Builder datacenter(String datacenter) {
            return datacenter(Output.of(datacenter));
        }

        /**
         * @param datastore The name of the datastore in which to create the
         * disk.
         * 
         * @return builder
         * 
         */
        public Builder datastore(Output<String> datastore) {
            $.datastore = datastore;
            return this;
        }

        /**
         * @param datastore The name of the datastore in which to create the
         * disk.
         * 
         * @return builder
         * 
         */
        public Builder datastore(String datastore) {
            return datastore(Output.of(datastore));
        }

        /**
         * @param size Size of the disk (in GB).
         * 
         * @return builder
         * 
         */
        public Builder size(Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size Size of the disk (in GB).
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param type The type of disk to create. Can be one of
         * `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
         * information on what each kind of disk provisioning policy means, click
         * [here][docs-vmware-vm-disk-provisioning].
         * 
         * [docs-vmware-vm-disk-provisioning]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-4C0F4D73-82F2-4B81-8AA7-1DD752A8A5AC.html
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of disk to create. Can be one of
         * `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
         * information on what each kind of disk provisioning policy means, click
         * [here][docs-vmware-vm-disk-provisioning].
         * 
         * [docs-vmware-vm-disk-provisioning]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-4C0F4D73-82F2-4B81-8AA7-1DD752A8A5AC.html
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param vmdkPath The path, including filename, of the virtual disk to
         * be created.  This needs to end in `.vmdk`.
         * 
         * @return builder
         * 
         */
        public Builder vmdkPath(Output<String> vmdkPath) {
            $.vmdkPath = vmdkPath;
            return this;
        }

        /**
         * @param vmdkPath The path, including filename, of the virtual disk to
         * be created.  This needs to end in `.vmdk`.
         * 
         * @return builder
         * 
         */
        public Builder vmdkPath(String vmdkPath) {
            return vmdkPath(Output.of(vmdkPath));
        }

        public VirtualDiskArgs build() {
            if ($.datastore == null) {
                throw new MissingRequiredPropertyException("VirtualDiskArgs", "datastore");
            }
            if ($.size == null) {
                throw new MissingRequiredPropertyException("VirtualDiskArgs", "size");
            }
            if ($.vmdkPath == null) {
                throw new MissingRequiredPropertyException("VirtualDiskArgs", "vmdkPath");
            }
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class StorageDrsVmOverrideArgs extends com.pulumi.resources.ResourceArgs {

    public static final StorageDrsVmOverrideArgs Empty = new StorageDrsVmOverrideArgs();

    /**
     * The managed object reference
     * ID of the datastore cluster to put the override in.
     * Forces a new resource if changed.
     * 
     */
    @Import(name="datastoreClusterId", required=true)
    private Output<String> datastoreClusterId;

    /**
     * @return The managed object reference
     * ID of the datastore cluster to put the override in.
     * Forces a new resource if changed.
     * 
     */
    public Output<String> datastoreClusterId() {
        return this.datastoreClusterId;
    }

    /**
     * Overrides any Storage DRS automation
     * levels for this virtual machine. Can be one of `automated` or `manual`. When
     * not specified, the datastore cluster&#39;s settings are used according to the
     * specific SDRS subsystem.
     * 
     */
    @Import(name="sdrsAutomationLevel")
    private @Nullable Output<String> sdrsAutomationLevel;

    /**
     * @return Overrides any Storage DRS automation
     * levels for this virtual machine. Can be one of `automated` or `manual`. When
     * not specified, the datastore cluster&#39;s settings are used according to the
     * specific SDRS subsystem.
     * 
     */
    public Optional<Output<String>> sdrsAutomationLevel() {
        return Optional.ofNullable(this.sdrsAutomationLevel);
    }

    /**
     * Overrides the default Storage DRS setting for
     * this virtual machine. When not specified, the datastore cluster setting is
     * used.
     * 
     */
    @Import(name="sdrsEnabled")
    private @Nullable Output<String> sdrsEnabled;

    /**
     * @return Overrides the default Storage DRS setting for
     * this virtual machine. When not specified, the datastore cluster setting is
     * used.
     * 
     */
    public Optional<Output<String>> sdrsEnabled() {
        return Optional.ofNullable(this.sdrsEnabled);
    }

    /**
     * Overrides the intra-VM affinity setting
     * for this virtual machine. When `true`, all disks for this virtual machine
     * will be kept on the same datastore. When `false`, Storage DRS may locate
     * individual disks on different datastores if it helps satisfy cluster
     * requirements. When not specified, the datastore cluster&#39;s settings are used.
     * 
     */
    @Import(name="sdrsIntraVmAffinity")
    private @Nullable Output<String> sdrsIntraVmAffinity;

    /**
     * @return Overrides the intra-VM affinity setting
     * for this virtual machine. When `true`, all disks for this virtual machine
     * will be kept on the same datastore. When `false`, Storage DRS may locate
     * individual disks on different datastores if it helps satisfy cluster
     * requirements. When not specified, the datastore cluster&#39;s settings are used.
     * 
     */
    public Optional<Output<String>> sdrsIntraVmAffinity() {
        return Optional.ofNullable(this.sdrsIntraVmAffinity);
    }

    /**
     * The UUID of the virtual machine to create
     * the override for.  Forces a new resource if changed.
     * 
     */
    @Import(name="virtualMachineId", required=true)
    private Output<String> virtualMachineId;

    /**
     * @return The UUID of the virtual machine to create
     * the override for.  Forces a new resource if changed.
     * 
     */
    public Output<String> virtualMachineId() {
        return this.virtualMachineId;
    }

    private StorageDrsVmOverrideArgs() {}

    private StorageDrsVmOverrideArgs(StorageDrsVmOverrideArgs $) {
        this.datastoreClusterId = $.datastoreClusterId;
        this.sdrsAutomationLevel = $.sdrsAutomationLevel;
        this.sdrsEnabled = $.sdrsEnabled;
        this.sdrsIntraVmAffinity = $.sdrsIntraVmAffinity;
        this.virtualMachineId = $.virtualMachineId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(StorageDrsVmOverrideArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private StorageDrsVmOverrideArgs $;

        public Builder() {
            $ = new StorageDrsVmOverrideArgs();
        }

        public Builder(StorageDrsVmOverrideArgs defaults) {
            $ = new StorageDrsVmOverrideArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param datastoreClusterId The managed object reference
         * ID of the datastore cluster to put the override in.
         * Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder datastoreClusterId(Output<String> datastoreClusterId) {
            $.datastoreClusterId = datastoreClusterId;
            return this;
        }

        /**
         * @param datastoreClusterId The managed object reference
         * ID of the datastore cluster to put the override in.
         * Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder datastoreClusterId(String datastoreClusterId) {
            return datastoreClusterId(Output.of(datastoreClusterId));
        }

        /**
         * @param sdrsAutomationLevel Overrides any Storage DRS automation
         * levels for this virtual machine. Can be one of `automated` or `manual`. When
         * not specified, the datastore cluster&#39;s settings are used according to the
         * specific SDRS subsystem.
         * 
         * @return builder
         * 
         */
        public Builder sdrsAutomationLevel(@Nullable Output<String> sdrsAutomationLevel) {
            $.sdrsAutomationLevel = sdrsAutomationLevel;
            return this;
        }

        /**
         * @param sdrsAutomationLevel Overrides any Storage DRS automation
         * levels for this virtual machine. Can be one of `automated` or `manual`. When
         * not specified, the datastore cluster&#39;s settings are used according to the
         * specific SDRS subsystem.
         * 
         * @return builder
         * 
         */
        public Builder sdrsAutomationLevel(String sdrsAutomationLevel) {
            return sdrsAutomationLevel(Output.of(sdrsAutomationLevel));
        }

        /**
         * @param sdrsEnabled Overrides the default Storage DRS setting for
         * this virtual machine. When not specified, the datastore cluster setting is
         * used.
         * 
         * @return builder
         * 
         */
        public Builder sdrsEnabled(@Nullable Output<String> sdrsEnabled) {
            $.sdrsEnabled = sdrsEnabled;
            return this;
        }

        /**
         * @param sdrsEnabled Overrides the default Storage DRS setting for
         * this virtual machine. When not specified, the datastore cluster setting is
         * used.
         * 
         * @return builder
         * 
         */
        public Builder sdrsEnabled(String sdrsEnabled) {
            return sdrsEnabled(Output.of(sdrsEnabled));
        }

        /**
         * @param sdrsIntraVmAffinity Overrides the intra-VM affinity setting
         * for this virtual machine. When `true`, all disks for this virtual machine
         * will be kept on the same datastore. When `false`, Storage DRS may locate
         * individual disks on different datastores if it helps satisfy cluster
         * requirements. When not specified, the datastore cluster&#39;s settings are used.
         * 
         * @return builder
         * 
         */
        public Builder sdrsIntraVmAffinity(@Nullable Output<String> sdrsIntraVmAffinity) {
            $.sdrsIntraVmAffinity = sdrsIntraVmAffinity;
            return this;
        }

        /**
         * @param sdrsIntraVmAffinity Overrides the intra-VM affinity setting
         * for this virtual machine. When `true`, all disks for this virtual machine
         * will be kept on the same datastore. When `false`, Storage DRS may locate
         * individual disks on different datastores if it helps satisfy cluster
         * requirements. When not specified, the datastore cluster&#39;s settings are used.
         * 
         * @return builder
         * 
         */
        public Builder sdrsIntraVmAffinity(String sdrsIntraVmAffinity) {
            return sdrsIntraVmAffinity(Output.of(sdrsIntraVmAffinity));
        }

        /**
         * @param virtualMachineId The UUID of the virtual machine to create
         * the override for.  Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder virtualMachineId(Output<String> virtualMachineId) {
            $.virtualMachineId = virtualMachineId;
            return this;
        }

        /**
         * @param virtualMachineId The UUID of the virtual machine to create
         * the override for.  Forces a new resource if changed.
         * 
         * @return builder
         * 
         */
        public Builder virtualMachineId(String virtualMachineId) {
            return virtualMachineId(Output.of(virtualMachineId));
        }

        public StorageDrsVmOverrideArgs build() {
            $.datastoreClusterId = Objects.requireNonNull($.datastoreClusterId, "expected parameter 'datastoreClusterId' to be non-null");
            $.virtualMachineId = Objects.requireNonNull($.virtualMachineId, "expected parameter 'virtualMachineId' to be non-null");
            return $;
        }
    }

}

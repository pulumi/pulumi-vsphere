// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vsphere.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.vsphere.inputs.GetVirtualMachineVapp;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVirtualMachinePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVirtualMachinePlainArgs Empty = new GetVirtualMachinePlainArgs();

    /**
     * The alternate guest name of the virtual machine when
     * `guest_id` is a non-specific operating system, like `otherGuest` or
     * `otherGuest64`.
     * 
     */
    @Import(name="alternateGuestName")
    private @Nullable String alternateGuestName;

    /**
     * @return The alternate guest name of the virtual machine when
     * `guest_id` is a non-specific operating system, like `otherGuest` or
     * `otherGuest64`.
     * 
     */
    public Optional<String> alternateGuestName() {
        return Optional.ofNullable(this.alternateGuestName);
    }

    /**
     * The user-provided description of this virtual machine.
     * 
     */
    @Import(name="annotation")
    private @Nullable String annotation;

    /**
     * @return The user-provided description of this virtual machine.
     * 
     */
    public Optional<String> annotation() {
        return Optional.ofNullable(this.annotation);
    }

    @Import(name="bootDelay")
    private @Nullable Integer bootDelay;

    public Optional<Integer> bootDelay() {
        return Optional.ofNullable(this.bootDelay);
    }

    @Import(name="bootRetryDelay")
    private @Nullable Integer bootRetryDelay;

    public Optional<Integer> bootRetryDelay() {
        return Optional.ofNullable(this.bootRetryDelay);
    }

    @Import(name="bootRetryEnabled")
    private @Nullable Boolean bootRetryEnabled;

    public Optional<Boolean> bootRetryEnabled() {
        return Optional.ofNullable(this.bootRetryEnabled);
    }

    @Import(name="cpuHotAddEnabled")
    private @Nullable Boolean cpuHotAddEnabled;

    public Optional<Boolean> cpuHotAddEnabled() {
        return Optional.ofNullable(this.cpuHotAddEnabled);
    }

    @Import(name="cpuHotRemoveEnabled")
    private @Nullable Boolean cpuHotRemoveEnabled;

    public Optional<Boolean> cpuHotRemoveEnabled() {
        return Optional.ofNullable(this.cpuHotRemoveEnabled);
    }

    @Import(name="cpuLimit")
    private @Nullable Integer cpuLimit;

    public Optional<Integer> cpuLimit() {
        return Optional.ofNullable(this.cpuLimit);
    }

    @Import(name="cpuPerformanceCountersEnabled")
    private @Nullable Boolean cpuPerformanceCountersEnabled;

    public Optional<Boolean> cpuPerformanceCountersEnabled() {
        return Optional.ofNullable(this.cpuPerformanceCountersEnabled);
    }

    @Import(name="cpuReservation")
    private @Nullable Integer cpuReservation;

    public Optional<Integer> cpuReservation() {
        return Optional.ofNullable(this.cpuReservation);
    }

    @Import(name="cpuShareCount")
    private @Nullable Integer cpuShareCount;

    public Optional<Integer> cpuShareCount() {
        return Optional.ofNullable(this.cpuShareCount);
    }

    @Import(name="cpuShareLevel")
    private @Nullable String cpuShareLevel;

    public Optional<String> cpuShareLevel() {
        return Optional.ofNullable(this.cpuShareLevel);
    }

    /**
     * The managed object reference
     * ID of the datacenter the virtual machine is located in.
     * This can be omitted if the search path used in `name` is an absolute path.
     * For default datacenters, use the `id` attribute from an empty
     * `vsphere.Datacenter` data source.
     * 
     */
    @Import(name="datacenterId")
    private @Nullable String datacenterId;

    /**
     * @return The managed object reference
     * ID of the datacenter the virtual machine is located in.
     * This can be omitted if the search path used in `name` is an absolute path.
     * For default datacenters, use the `id` attribute from an empty
     * `vsphere.Datacenter` data source.
     * 
     */
    public Optional<String> datacenterId() {
        return Optional.ofNullable(this.datacenterId);
    }

    @Import(name="efiSecureBootEnabled")
    private @Nullable Boolean efiSecureBootEnabled;

    public Optional<Boolean> efiSecureBootEnabled() {
        return Optional.ofNullable(this.efiSecureBootEnabled);
    }

    @Import(name="enableDiskUuid")
    private @Nullable Boolean enableDiskUuid;

    public Optional<Boolean> enableDiskUuid() {
        return Optional.ofNullable(this.enableDiskUuid);
    }

    @Import(name="enableLogging")
    private @Nullable Boolean enableLogging;

    public Optional<Boolean> enableLogging() {
        return Optional.ofNullable(this.enableLogging);
    }

    @Import(name="eptRviMode")
    private @Nullable String eptRviMode;

    public Optional<String> eptRviMode() {
        return Optional.ofNullable(this.eptRviMode);
    }

    @Import(name="extraConfig")
    private @Nullable Map<String,String> extraConfig;

    public Optional<Map<String,String>> extraConfig() {
        return Optional.ofNullable(this.extraConfig);
    }

    @Import(name="extraConfigRebootRequired")
    private @Nullable Boolean extraConfigRebootRequired;

    public Optional<Boolean> extraConfigRebootRequired() {
        return Optional.ofNullable(this.extraConfigRebootRequired);
    }

    /**
     * The firmware type for this virtual machine. Can be `bios` or
     * `efi`.
     * 
     */
    @Import(name="firmware")
    private @Nullable String firmware;

    /**
     * @return The firmware type for this virtual machine. Can be `bios` or
     * `efi`.
     * 
     */
    public Optional<String> firmware() {
        return Optional.ofNullable(this.firmware);
    }

    /**
     * The name of the virtual machine folder where the virtual machine is located. The `name` argument is limited to 80 characters. If the `name` argument includes the full path to the virtual machine and exceeds the 80 characters limit, the `folder` folder argument can be used.
     * 
     */
    @Import(name="folder")
    private @Nullable String folder;

    /**
     * @return The name of the virtual machine folder where the virtual machine is located. The `name` argument is limited to 80 characters. If the `name` argument includes the full path to the virtual machine and exceeds the 80 characters limit, the `folder` folder argument can be used.
     * 
     */
    public Optional<String> folder() {
        return Optional.ofNullable(this.folder);
    }

    /**
     * The guest ID of the virtual machine or template.
     * 
     */
    @Import(name="guestId")
    private @Nullable String guestId;

    /**
     * @return The guest ID of the virtual machine or template.
     * 
     */
    public Optional<String> guestId() {
        return Optional.ofNullable(this.guestId);
    }

    /**
     * The hardware version number on this virtual machine.
     * 
     */
    @Import(name="hardwareVersion")
    private @Nullable Integer hardwareVersion;

    /**
     * @return The hardware version number on this virtual machine.
     * 
     */
    public Optional<Integer> hardwareVersion() {
        return Optional.ofNullable(this.hardwareVersion);
    }

    @Import(name="hvMode")
    private @Nullable String hvMode;

    public Optional<String> hvMode() {
        return Optional.ofNullable(this.hvMode);
    }

    @Import(name="ideControllerScanCount")
    private @Nullable Integer ideControllerScanCount;

    public Optional<Integer> ideControllerScanCount() {
        return Optional.ofNullable(this.ideControllerScanCount);
    }

    @Import(name="latencySensitivity")
    private @Nullable String latencySensitivity;

    public Optional<String> latencySensitivity() {
        return Optional.ofNullable(this.latencySensitivity);
    }

    /**
     * The size of the virtual machine&#39;s memory, in MB.
     * 
     */
    @Import(name="memory")
    private @Nullable Integer memory;

    /**
     * @return The size of the virtual machine&#39;s memory, in MB.
     * 
     */
    public Optional<Integer> memory() {
        return Optional.ofNullable(this.memory);
    }

    @Import(name="memoryHotAddEnabled")
    private @Nullable Boolean memoryHotAddEnabled;

    public Optional<Boolean> memoryHotAddEnabled() {
        return Optional.ofNullable(this.memoryHotAddEnabled);
    }

    @Import(name="memoryLimit")
    private @Nullable Integer memoryLimit;

    public Optional<Integer> memoryLimit() {
        return Optional.ofNullable(this.memoryLimit);
    }

    @Import(name="memoryReservation")
    private @Nullable Integer memoryReservation;

    public Optional<Integer> memoryReservation() {
        return Optional.ofNullable(this.memoryReservation);
    }

    @Import(name="memoryReservationLockedToMax")
    private @Nullable Boolean memoryReservationLockedToMax;

    public Optional<Boolean> memoryReservationLockedToMax() {
        return Optional.ofNullable(this.memoryReservationLockedToMax);
    }

    @Import(name="memoryShareCount")
    private @Nullable Integer memoryShareCount;

    public Optional<Integer> memoryShareCount() {
        return Optional.ofNullable(this.memoryShareCount);
    }

    @Import(name="memoryShareLevel")
    private @Nullable String memoryShareLevel;

    public Optional<String> memoryShareLevel() {
        return Optional.ofNullable(this.memoryShareLevel);
    }

    @Import(name="moid")
    private @Nullable String moid;

    public Optional<String> moid() {
        return Optional.ofNullable(this.moid);
    }

    /**
     * The name of the virtual machine. This can be a name or
     * the full path relative to the datacenter. This is required if a UUID lookup
     * is not performed.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the virtual machine. This can be a name or
     * the full path relative to the datacenter. This is required if a UUID lookup
     * is not performed.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="nestedHvEnabled")
    private @Nullable Boolean nestedHvEnabled;

    public Optional<Boolean> nestedHvEnabled() {
        return Optional.ofNullable(this.nestedHvEnabled);
    }

    /**
     * The number of cores per socket for this virtual
     * machine.
     * 
     */
    @Import(name="numCoresPerSocket")
    private @Nullable Integer numCoresPerSocket;

    /**
     * @return The number of cores per socket for this virtual
     * machine.
     * 
     */
    public Optional<Integer> numCoresPerSocket() {
        return Optional.ofNullable(this.numCoresPerSocket);
    }

    /**
     * The total number of virtual processor cores assigned to this
     * virtual machine.
     * 
     */
    @Import(name="numCpus")
    private @Nullable Integer numCpus;

    /**
     * @return The total number of virtual processor cores assigned to this
     * virtual machine.
     * 
     */
    public Optional<Integer> numCpus() {
        return Optional.ofNullable(this.numCpus);
    }

    /**
     * The number of NVMe controllers to
     * scan for disk attributes and controller types on. Default: `1`.
     * 
     * &gt; **NOTE:** For best results, ensure that all the disks on any templates you
     * use with this data source reside on the primary controller, and leave this
     * value at the default. See the `vsphere.VirtualMachine`
     * resource documentation for the significance of this setting, specifically the
     * additional requirements and notes for cloning section.
     * 
     */
    @Import(name="nvmeControllerScanCount")
    private @Nullable Integer nvmeControllerScanCount;

    /**
     * @return The number of NVMe controllers to
     * scan for disk attributes and controller types on. Default: `1`.
     * 
     * &gt; **NOTE:** For best results, ensure that all the disks on any templates you
     * use with this data source reside on the primary controller, and leave this
     * value at the default. See the `vsphere.VirtualMachine`
     * resource documentation for the significance of this setting, specifically the
     * additional requirements and notes for cloning section.
     * 
     */
    public Optional<Integer> nvmeControllerScanCount() {
        return Optional.ofNullable(this.nvmeControllerScanCount);
    }

    @Import(name="replaceTrigger")
    private @Nullable String replaceTrigger;

    public Optional<String> replaceTrigger() {
        return Optional.ofNullable(this.replaceTrigger);
    }

    @Import(name="runToolsScriptsAfterPowerOn")
    private @Nullable Boolean runToolsScriptsAfterPowerOn;

    public Optional<Boolean> runToolsScriptsAfterPowerOn() {
        return Optional.ofNullable(this.runToolsScriptsAfterPowerOn);
    }

    @Import(name="runToolsScriptsAfterResume")
    private @Nullable Boolean runToolsScriptsAfterResume;

    public Optional<Boolean> runToolsScriptsAfterResume() {
        return Optional.ofNullable(this.runToolsScriptsAfterResume);
    }

    @Import(name="runToolsScriptsBeforeGuestReboot")
    private @Nullable Boolean runToolsScriptsBeforeGuestReboot;

    public Optional<Boolean> runToolsScriptsBeforeGuestReboot() {
        return Optional.ofNullable(this.runToolsScriptsBeforeGuestReboot);
    }

    @Import(name="runToolsScriptsBeforeGuestShutdown")
    private @Nullable Boolean runToolsScriptsBeforeGuestShutdown;

    public Optional<Boolean> runToolsScriptsBeforeGuestShutdown() {
        return Optional.ofNullable(this.runToolsScriptsBeforeGuestShutdown);
    }

    @Import(name="runToolsScriptsBeforeGuestStandby")
    private @Nullable Boolean runToolsScriptsBeforeGuestStandby;

    public Optional<Boolean> runToolsScriptsBeforeGuestStandby() {
        return Optional.ofNullable(this.runToolsScriptsBeforeGuestStandby);
    }

    @Import(name="sataControllerScanCount")
    private @Nullable Integer sataControllerScanCount;

    public Optional<Integer> sataControllerScanCount() {
        return Optional.ofNullable(this.sataControllerScanCount);
    }

    /**
     * The number of SCSI controllers to
     * scan for disk attributes and controller types on. Default: `1`.
     * 
     */
    @Import(name="scsiControllerScanCount")
    private @Nullable Integer scsiControllerScanCount;

    /**
     * @return The number of SCSI controllers to
     * scan for disk attributes and controller types on. Default: `1`.
     * 
     */
    public Optional<Integer> scsiControllerScanCount() {
        return Optional.ofNullable(this.scsiControllerScanCount);
    }

    @Import(name="storagePolicyId")
    private @Nullable String storagePolicyId;

    public Optional<String> storagePolicyId() {
        return Optional.ofNullable(this.storagePolicyId);
    }

    @Import(name="swapPlacementPolicy")
    private @Nullable String swapPlacementPolicy;

    public Optional<String> swapPlacementPolicy() {
        return Optional.ofNullable(this.swapPlacementPolicy);
    }

    @Import(name="syncTimeWithHost")
    private @Nullable Boolean syncTimeWithHost;

    public Optional<Boolean> syncTimeWithHost() {
        return Optional.ofNullable(this.syncTimeWithHost);
    }

    @Import(name="syncTimeWithHostPeriodically")
    private @Nullable Boolean syncTimeWithHostPeriodically;

    public Optional<Boolean> syncTimeWithHostPeriodically() {
        return Optional.ofNullable(this.syncTimeWithHostPeriodically);
    }

    @Import(name="toolsUpgradePolicy")
    private @Nullable String toolsUpgradePolicy;

    public Optional<String> toolsUpgradePolicy() {
        return Optional.ofNullable(this.toolsUpgradePolicy);
    }

    /**
     * Specify this field for a UUID lookup, `name` and `datacenter_id`
     * are not required if this is specified.
     * 
     */
    @Import(name="uuid")
    private @Nullable String uuid;

    /**
     * @return Specify this field for a UUID lookup, `name` and `datacenter_id`
     * are not required if this is specified.
     * 
     */
    public Optional<String> uuid() {
        return Optional.ofNullable(this.uuid);
    }

    @Import(name="vapp")
    private @Nullable GetVirtualMachineVapp vapp;

    public Optional<GetVirtualMachineVapp> vapp() {
        return Optional.ofNullable(this.vapp);
    }

    @Import(name="vbsEnabled")
    private @Nullable Boolean vbsEnabled;

    public Optional<Boolean> vbsEnabled() {
        return Optional.ofNullable(this.vbsEnabled);
    }

    @Import(name="vvtdEnabled")
    private @Nullable Boolean vvtdEnabled;

    public Optional<Boolean> vvtdEnabled() {
        return Optional.ofNullable(this.vvtdEnabled);
    }

    private GetVirtualMachinePlainArgs() {}

    private GetVirtualMachinePlainArgs(GetVirtualMachinePlainArgs $) {
        this.alternateGuestName = $.alternateGuestName;
        this.annotation = $.annotation;
        this.bootDelay = $.bootDelay;
        this.bootRetryDelay = $.bootRetryDelay;
        this.bootRetryEnabled = $.bootRetryEnabled;
        this.cpuHotAddEnabled = $.cpuHotAddEnabled;
        this.cpuHotRemoveEnabled = $.cpuHotRemoveEnabled;
        this.cpuLimit = $.cpuLimit;
        this.cpuPerformanceCountersEnabled = $.cpuPerformanceCountersEnabled;
        this.cpuReservation = $.cpuReservation;
        this.cpuShareCount = $.cpuShareCount;
        this.cpuShareLevel = $.cpuShareLevel;
        this.datacenterId = $.datacenterId;
        this.efiSecureBootEnabled = $.efiSecureBootEnabled;
        this.enableDiskUuid = $.enableDiskUuid;
        this.enableLogging = $.enableLogging;
        this.eptRviMode = $.eptRviMode;
        this.extraConfig = $.extraConfig;
        this.extraConfigRebootRequired = $.extraConfigRebootRequired;
        this.firmware = $.firmware;
        this.folder = $.folder;
        this.guestId = $.guestId;
        this.hardwareVersion = $.hardwareVersion;
        this.hvMode = $.hvMode;
        this.ideControllerScanCount = $.ideControllerScanCount;
        this.latencySensitivity = $.latencySensitivity;
        this.memory = $.memory;
        this.memoryHotAddEnabled = $.memoryHotAddEnabled;
        this.memoryLimit = $.memoryLimit;
        this.memoryReservation = $.memoryReservation;
        this.memoryReservationLockedToMax = $.memoryReservationLockedToMax;
        this.memoryShareCount = $.memoryShareCount;
        this.memoryShareLevel = $.memoryShareLevel;
        this.moid = $.moid;
        this.name = $.name;
        this.nestedHvEnabled = $.nestedHvEnabled;
        this.numCoresPerSocket = $.numCoresPerSocket;
        this.numCpus = $.numCpus;
        this.nvmeControllerScanCount = $.nvmeControllerScanCount;
        this.replaceTrigger = $.replaceTrigger;
        this.runToolsScriptsAfterPowerOn = $.runToolsScriptsAfterPowerOn;
        this.runToolsScriptsAfterResume = $.runToolsScriptsAfterResume;
        this.runToolsScriptsBeforeGuestReboot = $.runToolsScriptsBeforeGuestReboot;
        this.runToolsScriptsBeforeGuestShutdown = $.runToolsScriptsBeforeGuestShutdown;
        this.runToolsScriptsBeforeGuestStandby = $.runToolsScriptsBeforeGuestStandby;
        this.sataControllerScanCount = $.sataControllerScanCount;
        this.scsiControllerScanCount = $.scsiControllerScanCount;
        this.storagePolicyId = $.storagePolicyId;
        this.swapPlacementPolicy = $.swapPlacementPolicy;
        this.syncTimeWithHost = $.syncTimeWithHost;
        this.syncTimeWithHostPeriodically = $.syncTimeWithHostPeriodically;
        this.toolsUpgradePolicy = $.toolsUpgradePolicy;
        this.uuid = $.uuid;
        this.vapp = $.vapp;
        this.vbsEnabled = $.vbsEnabled;
        this.vvtdEnabled = $.vvtdEnabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVirtualMachinePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVirtualMachinePlainArgs $;

        public Builder() {
            $ = new GetVirtualMachinePlainArgs();
        }

        public Builder(GetVirtualMachinePlainArgs defaults) {
            $ = new GetVirtualMachinePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param alternateGuestName The alternate guest name of the virtual machine when
         * `guest_id` is a non-specific operating system, like `otherGuest` or
         * `otherGuest64`.
         * 
         * @return builder
         * 
         */
        public Builder alternateGuestName(@Nullable String alternateGuestName) {
            $.alternateGuestName = alternateGuestName;
            return this;
        }

        /**
         * @param annotation The user-provided description of this virtual machine.
         * 
         * @return builder
         * 
         */
        public Builder annotation(@Nullable String annotation) {
            $.annotation = annotation;
            return this;
        }

        public Builder bootDelay(@Nullable Integer bootDelay) {
            $.bootDelay = bootDelay;
            return this;
        }

        public Builder bootRetryDelay(@Nullable Integer bootRetryDelay) {
            $.bootRetryDelay = bootRetryDelay;
            return this;
        }

        public Builder bootRetryEnabled(@Nullable Boolean bootRetryEnabled) {
            $.bootRetryEnabled = bootRetryEnabled;
            return this;
        }

        public Builder cpuHotAddEnabled(@Nullable Boolean cpuHotAddEnabled) {
            $.cpuHotAddEnabled = cpuHotAddEnabled;
            return this;
        }

        public Builder cpuHotRemoveEnabled(@Nullable Boolean cpuHotRemoveEnabled) {
            $.cpuHotRemoveEnabled = cpuHotRemoveEnabled;
            return this;
        }

        public Builder cpuLimit(@Nullable Integer cpuLimit) {
            $.cpuLimit = cpuLimit;
            return this;
        }

        public Builder cpuPerformanceCountersEnabled(@Nullable Boolean cpuPerformanceCountersEnabled) {
            $.cpuPerformanceCountersEnabled = cpuPerformanceCountersEnabled;
            return this;
        }

        public Builder cpuReservation(@Nullable Integer cpuReservation) {
            $.cpuReservation = cpuReservation;
            return this;
        }

        public Builder cpuShareCount(@Nullable Integer cpuShareCount) {
            $.cpuShareCount = cpuShareCount;
            return this;
        }

        public Builder cpuShareLevel(@Nullable String cpuShareLevel) {
            $.cpuShareLevel = cpuShareLevel;
            return this;
        }

        /**
         * @param datacenterId The managed object reference
         * ID of the datacenter the virtual machine is located in.
         * This can be omitted if the search path used in `name` is an absolute path.
         * For default datacenters, use the `id` attribute from an empty
         * `vsphere.Datacenter` data source.
         * 
         * @return builder
         * 
         */
        public Builder datacenterId(@Nullable String datacenterId) {
            $.datacenterId = datacenterId;
            return this;
        }

        public Builder efiSecureBootEnabled(@Nullable Boolean efiSecureBootEnabled) {
            $.efiSecureBootEnabled = efiSecureBootEnabled;
            return this;
        }

        public Builder enableDiskUuid(@Nullable Boolean enableDiskUuid) {
            $.enableDiskUuid = enableDiskUuid;
            return this;
        }

        public Builder enableLogging(@Nullable Boolean enableLogging) {
            $.enableLogging = enableLogging;
            return this;
        }

        public Builder eptRviMode(@Nullable String eptRviMode) {
            $.eptRviMode = eptRviMode;
            return this;
        }

        public Builder extraConfig(@Nullable Map<String,String> extraConfig) {
            $.extraConfig = extraConfig;
            return this;
        }

        public Builder extraConfigRebootRequired(@Nullable Boolean extraConfigRebootRequired) {
            $.extraConfigRebootRequired = extraConfigRebootRequired;
            return this;
        }

        /**
         * @param firmware The firmware type for this virtual machine. Can be `bios` or
         * `efi`.
         * 
         * @return builder
         * 
         */
        public Builder firmware(@Nullable String firmware) {
            $.firmware = firmware;
            return this;
        }

        /**
         * @param folder The name of the virtual machine folder where the virtual machine is located. The `name` argument is limited to 80 characters. If the `name` argument includes the full path to the virtual machine and exceeds the 80 characters limit, the `folder` folder argument can be used.
         * 
         * @return builder
         * 
         */
        public Builder folder(@Nullable String folder) {
            $.folder = folder;
            return this;
        }

        /**
         * @param guestId The guest ID of the virtual machine or template.
         * 
         * @return builder
         * 
         */
        public Builder guestId(@Nullable String guestId) {
            $.guestId = guestId;
            return this;
        }

        /**
         * @param hardwareVersion The hardware version number on this virtual machine.
         * 
         * @return builder
         * 
         */
        public Builder hardwareVersion(@Nullable Integer hardwareVersion) {
            $.hardwareVersion = hardwareVersion;
            return this;
        }

        public Builder hvMode(@Nullable String hvMode) {
            $.hvMode = hvMode;
            return this;
        }

        public Builder ideControllerScanCount(@Nullable Integer ideControllerScanCount) {
            $.ideControllerScanCount = ideControllerScanCount;
            return this;
        }

        public Builder latencySensitivity(@Nullable String latencySensitivity) {
            $.latencySensitivity = latencySensitivity;
            return this;
        }

        /**
         * @param memory The size of the virtual machine&#39;s memory, in MB.
         * 
         * @return builder
         * 
         */
        public Builder memory(@Nullable Integer memory) {
            $.memory = memory;
            return this;
        }

        public Builder memoryHotAddEnabled(@Nullable Boolean memoryHotAddEnabled) {
            $.memoryHotAddEnabled = memoryHotAddEnabled;
            return this;
        }

        public Builder memoryLimit(@Nullable Integer memoryLimit) {
            $.memoryLimit = memoryLimit;
            return this;
        }

        public Builder memoryReservation(@Nullable Integer memoryReservation) {
            $.memoryReservation = memoryReservation;
            return this;
        }

        public Builder memoryReservationLockedToMax(@Nullable Boolean memoryReservationLockedToMax) {
            $.memoryReservationLockedToMax = memoryReservationLockedToMax;
            return this;
        }

        public Builder memoryShareCount(@Nullable Integer memoryShareCount) {
            $.memoryShareCount = memoryShareCount;
            return this;
        }

        public Builder memoryShareLevel(@Nullable String memoryShareLevel) {
            $.memoryShareLevel = memoryShareLevel;
            return this;
        }

        public Builder moid(@Nullable String moid) {
            $.moid = moid;
            return this;
        }

        /**
         * @param name The name of the virtual machine. This can be a name or
         * the full path relative to the datacenter. This is required if a UUID lookup
         * is not performed.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public Builder nestedHvEnabled(@Nullable Boolean nestedHvEnabled) {
            $.nestedHvEnabled = nestedHvEnabled;
            return this;
        }

        /**
         * @param numCoresPerSocket The number of cores per socket for this virtual
         * machine.
         * 
         * @return builder
         * 
         */
        public Builder numCoresPerSocket(@Nullable Integer numCoresPerSocket) {
            $.numCoresPerSocket = numCoresPerSocket;
            return this;
        }

        /**
         * @param numCpus The total number of virtual processor cores assigned to this
         * virtual machine.
         * 
         * @return builder
         * 
         */
        public Builder numCpus(@Nullable Integer numCpus) {
            $.numCpus = numCpus;
            return this;
        }

        /**
         * @param nvmeControllerScanCount The number of NVMe controllers to
         * scan for disk attributes and controller types on. Default: `1`.
         * 
         * &gt; **NOTE:** For best results, ensure that all the disks on any templates you
         * use with this data source reside on the primary controller, and leave this
         * value at the default. See the `vsphere.VirtualMachine`
         * resource documentation for the significance of this setting, specifically the
         * additional requirements and notes for cloning section.
         * 
         * @return builder
         * 
         */
        public Builder nvmeControllerScanCount(@Nullable Integer nvmeControllerScanCount) {
            $.nvmeControllerScanCount = nvmeControllerScanCount;
            return this;
        }

        public Builder replaceTrigger(@Nullable String replaceTrigger) {
            $.replaceTrigger = replaceTrigger;
            return this;
        }

        public Builder runToolsScriptsAfterPowerOn(@Nullable Boolean runToolsScriptsAfterPowerOn) {
            $.runToolsScriptsAfterPowerOn = runToolsScriptsAfterPowerOn;
            return this;
        }

        public Builder runToolsScriptsAfterResume(@Nullable Boolean runToolsScriptsAfterResume) {
            $.runToolsScriptsAfterResume = runToolsScriptsAfterResume;
            return this;
        }

        public Builder runToolsScriptsBeforeGuestReboot(@Nullable Boolean runToolsScriptsBeforeGuestReboot) {
            $.runToolsScriptsBeforeGuestReboot = runToolsScriptsBeforeGuestReboot;
            return this;
        }

        public Builder runToolsScriptsBeforeGuestShutdown(@Nullable Boolean runToolsScriptsBeforeGuestShutdown) {
            $.runToolsScriptsBeforeGuestShutdown = runToolsScriptsBeforeGuestShutdown;
            return this;
        }

        public Builder runToolsScriptsBeforeGuestStandby(@Nullable Boolean runToolsScriptsBeforeGuestStandby) {
            $.runToolsScriptsBeforeGuestStandby = runToolsScriptsBeforeGuestStandby;
            return this;
        }

        public Builder sataControllerScanCount(@Nullable Integer sataControllerScanCount) {
            $.sataControllerScanCount = sataControllerScanCount;
            return this;
        }

        /**
         * @param scsiControllerScanCount The number of SCSI controllers to
         * scan for disk attributes and controller types on. Default: `1`.
         * 
         * @return builder
         * 
         */
        public Builder scsiControllerScanCount(@Nullable Integer scsiControllerScanCount) {
            $.scsiControllerScanCount = scsiControllerScanCount;
            return this;
        }

        public Builder storagePolicyId(@Nullable String storagePolicyId) {
            $.storagePolicyId = storagePolicyId;
            return this;
        }

        public Builder swapPlacementPolicy(@Nullable String swapPlacementPolicy) {
            $.swapPlacementPolicy = swapPlacementPolicy;
            return this;
        }

        public Builder syncTimeWithHost(@Nullable Boolean syncTimeWithHost) {
            $.syncTimeWithHost = syncTimeWithHost;
            return this;
        }

        public Builder syncTimeWithHostPeriodically(@Nullable Boolean syncTimeWithHostPeriodically) {
            $.syncTimeWithHostPeriodically = syncTimeWithHostPeriodically;
            return this;
        }

        public Builder toolsUpgradePolicy(@Nullable String toolsUpgradePolicy) {
            $.toolsUpgradePolicy = toolsUpgradePolicy;
            return this;
        }

        /**
         * @param uuid Specify this field for a UUID lookup, `name` and `datacenter_id`
         * are not required if this is specified.
         * 
         * @return builder
         * 
         */
        public Builder uuid(@Nullable String uuid) {
            $.uuid = uuid;
            return this;
        }

        public Builder vapp(@Nullable GetVirtualMachineVapp vapp) {
            $.vapp = vapp;
            return this;
        }

        public Builder vbsEnabled(@Nullable Boolean vbsEnabled) {
            $.vbsEnabled = vbsEnabled;
            return this;
        }

        public Builder vvtdEnabled(@Nullable Boolean vvtdEnabled) {
            $.vvtdEnabled = vvtdEnabled;
            return this;
        }

        public GetVirtualMachinePlainArgs build() {
            return $;
        }
    }

}

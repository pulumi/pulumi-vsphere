// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetVirtualMachine
    {
        /// <summary>
        /// The `vsphere.VirtualMachine` data source can be used to find the UUID of an
        /// existing virtual machine or template. The most common purpose is for finding
        /// the UUID of a template to be used as the source for cloning to a new
        /// `vsphere.VirtualMachine` resource. It also
        /// reads the guest ID so that can be supplied as well.
        /// 
        /// ## Example Usage
        /// 
        /// In the following example, a virtual machine template is returned by its
        /// unique name within the `vsphere.Datacenter`.
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using VSphere = Pulumi.VSphere;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var datacenter = VSphere.GetDatacenter.Invoke(new()
        ///     {
        ///         Name = "dc-01",
        ///     });
        /// 
        ///     var template = VSphere.GetVirtualMachine.Invoke(new()
        ///     {
        ///         Name = "ubuntu-server-template",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// In the following example, each virtual machine template is returned by its
        /// unique full path within the `vsphere.Datacenter`.
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using VSphere = Pulumi.VSphere;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var datacenter = VSphere.GetDatacenter.Invoke(new()
        ///     {
        ///         Name = "dc-01",
        ///     });
        /// 
        ///     var productionTemplate = VSphere.GetVirtualMachine.Invoke(new()
        ///     {
        ///         Name = "production/templates/ubuntu-server-template",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        ///     var developmentTemplate = VSphere.GetVirtualMachine.Invoke(new()
        ///     {
        ///         Name = "development/templates/ubuntu-server-template",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetVirtualMachineResult> InvokeAsync(GetVirtualMachineArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualMachineResult>("vsphere:index/getVirtualMachine:getVirtualMachine", args ?? new GetVirtualMachineArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.VirtualMachine` data source can be used to find the UUID of an
        /// existing virtual machine or template. The most common purpose is for finding
        /// the UUID of a template to be used as the source for cloning to a new
        /// `vsphere.VirtualMachine` resource. It also
        /// reads the guest ID so that can be supplied as well.
        /// 
        /// ## Example Usage
        /// 
        /// In the following example, a virtual machine template is returned by its
        /// unique name within the `vsphere.Datacenter`.
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using VSphere = Pulumi.VSphere;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var datacenter = VSphere.GetDatacenter.Invoke(new()
        ///     {
        ///         Name = "dc-01",
        ///     });
        /// 
        ///     var template = VSphere.GetVirtualMachine.Invoke(new()
        ///     {
        ///         Name = "ubuntu-server-template",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// In the following example, each virtual machine template is returned by its
        /// unique full path within the `vsphere.Datacenter`.
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using VSphere = Pulumi.VSphere;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var datacenter = VSphere.GetDatacenter.Invoke(new()
        ///     {
        ///         Name = "dc-01",
        ///     });
        /// 
        ///     var productionTemplate = VSphere.GetVirtualMachine.Invoke(new()
        ///     {
        ///         Name = "production/templates/ubuntu-server-template",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        ///     var developmentTemplate = VSphere.GetVirtualMachine.Invoke(new()
        ///     {
        ///         Name = "development/templates/ubuntu-server-template",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetVirtualMachineResult> Invoke(GetVirtualMachineInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualMachineResult>("vsphere:index/getVirtualMachine:getVirtualMachine", args ?? new GetVirtualMachineInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualMachineArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The alternate guest name of the virtual machine when
        /// `guest_id` is a non-specific operating system, like `otherGuest` or `otherGuest64`.
        /// </summary>
        [Input("alternateGuestName")]
        public string? AlternateGuestName { get; set; }

        /// <summary>
        /// The user-provided description of this virtual machine.
        /// </summary>
        [Input("annotation")]
        public string? Annotation { get; set; }

        [Input("bootDelay")]
        public int? BootDelay { get; set; }

        [Input("bootRetryDelay")]
        public int? BootRetryDelay { get; set; }

        [Input("bootRetryEnabled")]
        public bool? BootRetryEnabled { get; set; }

        [Input("cpuHotAddEnabled")]
        public bool? CpuHotAddEnabled { get; set; }

        [Input("cpuHotRemoveEnabled")]
        public bool? CpuHotRemoveEnabled { get; set; }

        [Input("cpuLimit")]
        public int? CpuLimit { get; set; }

        [Input("cpuPerformanceCountersEnabled")]
        public bool? CpuPerformanceCountersEnabled { get; set; }

        [Input("cpuReservation")]
        public int? CpuReservation { get; set; }

        [Input("cpuShareCount")]
        public int? CpuShareCount { get; set; }

        [Input("cpuShareLevel")]
        public string? CpuShareLevel { get; set; }

        /// <summary>
        /// The managed object reference
        /// ID of the datacenter the virtual machine is located in.
        /// This can be omitted if the search path used in `name` is an absolute path.
        /// For default datacenters, use the `id` attribute from an empty
        /// `vsphere.Datacenter` data source.
        /// </summary>
        [Input("datacenterId")]
        public string? DatacenterId { get; set; }

        [Input("efiSecureBootEnabled")]
        public bool? EfiSecureBootEnabled { get; set; }

        [Input("enableDiskUuid")]
        public bool? EnableDiskUuid { get; set; }

        [Input("enableLogging")]
        public bool? EnableLogging { get; set; }

        [Input("eptRviMode")]
        public string? EptRviMode { get; set; }

        [Input("extraConfig")]
        private Dictionary<string, string>? _extraConfig;
        public Dictionary<string, string> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new Dictionary<string, string>());
            set => _extraConfig = value;
        }

        [Input("extraConfigRebootRequired")]
        public bool? ExtraConfigRebootRequired { get; set; }

        /// <summary>
        /// The firmware type for this virtual machine. Can be `bios` or `efi`.
        /// </summary>
        [Input("firmware")]
        public string? Firmware { get; set; }

        /// <summary>
        /// The name of the virtual machine folder where the virtual machine is located. The `name` argument is limited to 80 characters. If the `name` argument includes the full path to the virtual machine and exceeds the 80 characters limit, the `folder` folder argument can be used.
        /// </summary>
        [Input("folder")]
        public string? Folder { get; set; }

        /// <summary>
        /// The guest ID of the virtual machine or template.
        /// </summary>
        [Input("guestId")]
        public string? GuestId { get; set; }

        /// <summary>
        /// The hardware version number on this virtual machine.
        /// </summary>
        [Input("hardwareVersion")]
        public int? HardwareVersion { get; set; }

        [Input("hvMode")]
        public string? HvMode { get; set; }

        [Input("ideControllerScanCount")]
        public int? IdeControllerScanCount { get; set; }

        [Input("latencySensitivity")]
        public string? LatencySensitivity { get; set; }

        /// <summary>
        /// The size of the virtual machine's memory, in MB.
        /// </summary>
        [Input("memory")]
        public int? Memory { get; set; }

        [Input("memoryHotAddEnabled")]
        public bool? MemoryHotAddEnabled { get; set; }

        [Input("memoryLimit")]
        public int? MemoryLimit { get; set; }

        [Input("memoryReservation")]
        public int? MemoryReservation { get; set; }

        [Input("memoryReservationLockedToMax")]
        public bool? MemoryReservationLockedToMax { get; set; }

        [Input("memoryShareCount")]
        public int? MemoryShareCount { get; set; }

        [Input("memoryShareLevel")]
        public string? MemoryShareLevel { get; set; }

        [Input("moid")]
        public string? Moid { get; set; }

        /// <summary>
        /// The name of the virtual machine. This can be a name or
        /// the full path relative to the datacenter. This is required if a UUID lookup
        /// is not performed.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("nestedHvEnabled")]
        public bool? NestedHvEnabled { get; set; }

        /// <summary>
        /// The number of cores per socket for this virtual machine.
        /// </summary>
        [Input("numCoresPerSocket")]
        public int? NumCoresPerSocket { get; set; }

        /// <summary>
        /// The total number of virtual processor cores assigned to this
        /// virtual machine.
        /// </summary>
        [Input("numCpus")]
        public int? NumCpus { get; set; }

        [Input("replaceTrigger")]
        public string? ReplaceTrigger { get; set; }

        [Input("runToolsScriptsAfterPowerOn")]
        public bool? RunToolsScriptsAfterPowerOn { get; set; }

        [Input("runToolsScriptsAfterResume")]
        public bool? RunToolsScriptsAfterResume { get; set; }

        [Input("runToolsScriptsBeforeGuestReboot")]
        public bool? RunToolsScriptsBeforeGuestReboot { get; set; }

        [Input("runToolsScriptsBeforeGuestShutdown")]
        public bool? RunToolsScriptsBeforeGuestShutdown { get; set; }

        [Input("runToolsScriptsBeforeGuestStandby")]
        public bool? RunToolsScriptsBeforeGuestStandby { get; set; }

        [Input("sataControllerScanCount")]
        public int? SataControllerScanCount { get; set; }

        /// <summary>
        /// The number of SCSI controllers to
        /// scan for disk attributes and controller types on. Default: `1`.
        /// 
        /// &gt; **NOTE:** For best results, ensure that all the disks on any templates you
        /// use with this data source reside on the primary controller, and leave this
        /// value at the default. See the `vsphere.VirtualMachine`
        /// resource documentation for the significance of this setting, specifically the
        /// additional requirements and notes for cloning section.
        /// </summary>
        [Input("scsiControllerScanCount")]
        public int? ScsiControllerScanCount { get; set; }

        [Input("storagePolicyId")]
        public string? StoragePolicyId { get; set; }

        [Input("swapPlacementPolicy")]
        public string? SwapPlacementPolicy { get; set; }

        [Input("syncTimeWithHost")]
        public bool? SyncTimeWithHost { get; set; }

        [Input("syncTimeWithHostPeriodically")]
        public bool? SyncTimeWithHostPeriodically { get; set; }

        [Input("toolsUpgradePolicy")]
        public string? ToolsUpgradePolicy { get; set; }

        /// <summary>
        /// Specify this field for a UUID lookup, `name` and `datacenter_id`
        /// are not required if this is specified.
        /// </summary>
        [Input("uuid")]
        public string? Uuid { get; set; }

        [Input("vapp")]
        public Inputs.GetVirtualMachineVappArgs? Vapp { get; set; }

        [Input("vbsEnabled")]
        public bool? VbsEnabled { get; set; }

        [Input("vvtdEnabled")]
        public bool? VvtdEnabled { get; set; }

        public GetVirtualMachineArgs()
        {
        }
        public static new GetVirtualMachineArgs Empty => new GetVirtualMachineArgs();
    }

    public sealed class GetVirtualMachineInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The alternate guest name of the virtual machine when
        /// `guest_id` is a non-specific operating system, like `otherGuest` or `otherGuest64`.
        /// </summary>
        [Input("alternateGuestName")]
        public Input<string>? AlternateGuestName { get; set; }

        /// <summary>
        /// The user-provided description of this virtual machine.
        /// </summary>
        [Input("annotation")]
        public Input<string>? Annotation { get; set; }

        [Input("bootDelay")]
        public Input<int>? BootDelay { get; set; }

        [Input("bootRetryDelay")]
        public Input<int>? BootRetryDelay { get; set; }

        [Input("bootRetryEnabled")]
        public Input<bool>? BootRetryEnabled { get; set; }

        [Input("cpuHotAddEnabled")]
        public Input<bool>? CpuHotAddEnabled { get; set; }

        [Input("cpuHotRemoveEnabled")]
        public Input<bool>? CpuHotRemoveEnabled { get; set; }

        [Input("cpuLimit")]
        public Input<int>? CpuLimit { get; set; }

        [Input("cpuPerformanceCountersEnabled")]
        public Input<bool>? CpuPerformanceCountersEnabled { get; set; }

        [Input("cpuReservation")]
        public Input<int>? CpuReservation { get; set; }

        [Input("cpuShareCount")]
        public Input<int>? CpuShareCount { get; set; }

        [Input("cpuShareLevel")]
        public Input<string>? CpuShareLevel { get; set; }

        /// <summary>
        /// The managed object reference
        /// ID of the datacenter the virtual machine is located in.
        /// This can be omitted if the search path used in `name` is an absolute path.
        /// For default datacenters, use the `id` attribute from an empty
        /// `vsphere.Datacenter` data source.
        /// </summary>
        [Input("datacenterId")]
        public Input<string>? DatacenterId { get; set; }

        [Input("efiSecureBootEnabled")]
        public Input<bool>? EfiSecureBootEnabled { get; set; }

        [Input("enableDiskUuid")]
        public Input<bool>? EnableDiskUuid { get; set; }

        [Input("enableLogging")]
        public Input<bool>? EnableLogging { get; set; }

        [Input("eptRviMode")]
        public Input<string>? EptRviMode { get; set; }

        [Input("extraConfig")]
        private InputMap<string>? _extraConfig;
        public InputMap<string> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputMap<string>());
            set => _extraConfig = value;
        }

        [Input("extraConfigRebootRequired")]
        public Input<bool>? ExtraConfigRebootRequired { get; set; }

        /// <summary>
        /// The firmware type for this virtual machine. Can be `bios` or `efi`.
        /// </summary>
        [Input("firmware")]
        public Input<string>? Firmware { get; set; }

        /// <summary>
        /// The name of the virtual machine folder where the virtual machine is located. The `name` argument is limited to 80 characters. If the `name` argument includes the full path to the virtual machine and exceeds the 80 characters limit, the `folder` folder argument can be used.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// The guest ID of the virtual machine or template.
        /// </summary>
        [Input("guestId")]
        public Input<string>? GuestId { get; set; }

        /// <summary>
        /// The hardware version number on this virtual machine.
        /// </summary>
        [Input("hardwareVersion")]
        public Input<int>? HardwareVersion { get; set; }

        [Input("hvMode")]
        public Input<string>? HvMode { get; set; }

        [Input("ideControllerScanCount")]
        public Input<int>? IdeControllerScanCount { get; set; }

        [Input("latencySensitivity")]
        public Input<string>? LatencySensitivity { get; set; }

        /// <summary>
        /// The size of the virtual machine's memory, in MB.
        /// </summary>
        [Input("memory")]
        public Input<int>? Memory { get; set; }

        [Input("memoryHotAddEnabled")]
        public Input<bool>? MemoryHotAddEnabled { get; set; }

        [Input("memoryLimit")]
        public Input<int>? MemoryLimit { get; set; }

        [Input("memoryReservation")]
        public Input<int>? MemoryReservation { get; set; }

        [Input("memoryReservationLockedToMax")]
        public Input<bool>? MemoryReservationLockedToMax { get; set; }

        [Input("memoryShareCount")]
        public Input<int>? MemoryShareCount { get; set; }

        [Input("memoryShareLevel")]
        public Input<string>? MemoryShareLevel { get; set; }

        [Input("moid")]
        public Input<string>? Moid { get; set; }

        /// <summary>
        /// The name of the virtual machine. This can be a name or
        /// the full path relative to the datacenter. This is required if a UUID lookup
        /// is not performed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nestedHvEnabled")]
        public Input<bool>? NestedHvEnabled { get; set; }

        /// <summary>
        /// The number of cores per socket for this virtual machine.
        /// </summary>
        [Input("numCoresPerSocket")]
        public Input<int>? NumCoresPerSocket { get; set; }

        /// <summary>
        /// The total number of virtual processor cores assigned to this
        /// virtual machine.
        /// </summary>
        [Input("numCpus")]
        public Input<int>? NumCpus { get; set; }

        [Input("replaceTrigger")]
        public Input<string>? ReplaceTrigger { get; set; }

        [Input("runToolsScriptsAfterPowerOn")]
        public Input<bool>? RunToolsScriptsAfterPowerOn { get; set; }

        [Input("runToolsScriptsAfterResume")]
        public Input<bool>? RunToolsScriptsAfterResume { get; set; }

        [Input("runToolsScriptsBeforeGuestReboot")]
        public Input<bool>? RunToolsScriptsBeforeGuestReboot { get; set; }

        [Input("runToolsScriptsBeforeGuestShutdown")]
        public Input<bool>? RunToolsScriptsBeforeGuestShutdown { get; set; }

        [Input("runToolsScriptsBeforeGuestStandby")]
        public Input<bool>? RunToolsScriptsBeforeGuestStandby { get; set; }

        [Input("sataControllerScanCount")]
        public Input<int>? SataControllerScanCount { get; set; }

        /// <summary>
        /// The number of SCSI controllers to
        /// scan for disk attributes and controller types on. Default: `1`.
        /// 
        /// &gt; **NOTE:** For best results, ensure that all the disks on any templates you
        /// use with this data source reside on the primary controller, and leave this
        /// value at the default. See the `vsphere.VirtualMachine`
        /// resource documentation for the significance of this setting, specifically the
        /// additional requirements and notes for cloning section.
        /// </summary>
        [Input("scsiControllerScanCount")]
        public Input<int>? ScsiControllerScanCount { get; set; }

        [Input("storagePolicyId")]
        public Input<string>? StoragePolicyId { get; set; }

        [Input("swapPlacementPolicy")]
        public Input<string>? SwapPlacementPolicy { get; set; }

        [Input("syncTimeWithHost")]
        public Input<bool>? SyncTimeWithHost { get; set; }

        [Input("syncTimeWithHostPeriodically")]
        public Input<bool>? SyncTimeWithHostPeriodically { get; set; }

        [Input("toolsUpgradePolicy")]
        public Input<string>? ToolsUpgradePolicy { get; set; }

        /// <summary>
        /// Specify this field for a UUID lookup, `name` and `datacenter_id`
        /// are not required if this is specified.
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        [Input("vapp")]
        public Input<Inputs.GetVirtualMachineVappInputArgs>? Vapp { get; set; }

        [Input("vbsEnabled")]
        public Input<bool>? VbsEnabled { get; set; }

        [Input("vvtdEnabled")]
        public Input<bool>? VvtdEnabled { get; set; }

        public GetVirtualMachineInvokeArgs()
        {
        }
        public static new GetVirtualMachineInvokeArgs Empty => new GetVirtualMachineInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualMachineResult
    {
        /// <summary>
        /// The alternate guest name of the virtual machine when
        /// `guest_id` is a non-specific operating system, like `otherGuest` or `otherGuest64`.
        /// </summary>
        public readonly string? AlternateGuestName;
        /// <summary>
        /// The user-provided description of this virtual machine.
        /// </summary>
        public readonly string Annotation;
        public readonly int? BootDelay;
        public readonly int? BootRetryDelay;
        public readonly bool? BootRetryEnabled;
        public readonly string ChangeVersion;
        public readonly bool? CpuHotAddEnabled;
        public readonly bool? CpuHotRemoveEnabled;
        public readonly int? CpuLimit;
        public readonly bool? CpuPerformanceCountersEnabled;
        public readonly int? CpuReservation;
        public readonly int CpuShareCount;
        public readonly string? CpuShareLevel;
        public readonly string? DatacenterId;
        /// <summary>
        /// Whenever possible, this is the first IPv4 address that is reachable through
        /// the default gateway configured on the machine, then the first reachable IPv6
        /// address, and then the first general discovered address if neither exist. If
        /// VMware Tools is not running on the virtual machine, or if the VM is powered
        /// off, this value will be blank.
        /// </summary>
        public readonly string DefaultIpAddress;
        /// <summary>
        /// Information about each of the disks on this virtual machine or
        /// template. These are sorted by bus and unit number so that they can be applied
        /// to a `vsphere.VirtualMachine` resource in the order the resource expects
        /// while cloning. This is useful for discovering certain disk settings while
        /// performing a linked clone, as all settings that are output by this data
        /// source must be the same on the destination virtual machine as the source.
        /// Only the first number of controllers defined by `scsi_controller_scan_count`
        /// are scanned for disks. The sub-attributes are:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualMachineDiskResult> Disks;
        public readonly bool? EfiSecureBootEnabled;
        public readonly bool? EnableDiskUuid;
        public readonly bool? EnableLogging;
        public readonly string? EptRviMode;
        public readonly ImmutableDictionary<string, string>? ExtraConfig;
        public readonly bool? ExtraConfigRebootRequired;
        /// <summary>
        /// The firmware type for this virtual machine. Can be `bios` or `efi`.
        /// </summary>
        public readonly string? Firmware;
        public readonly string? Folder;
        /// <summary>
        /// The guest ID of the virtual machine or template.
        /// </summary>
        public readonly string GuestId;
        /// <summary>
        /// A list of IP addresses as reported by VMware Tools.
        /// </summary>
        public readonly ImmutableArray<string> GuestIpAddresses;
        /// <summary>
        /// The hardware version number on this virtual machine.
        /// </summary>
        public readonly int HardwareVersion;
        public readonly string? HvMode;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int? IdeControllerScanCount;
        public readonly string? LatencySensitivity;
        /// <summary>
        /// The size of the virtual machine's memory, in MB.
        /// </summary>
        public readonly int? Memory;
        public readonly bool? MemoryHotAddEnabled;
        public readonly int? MemoryLimit;
        public readonly int? MemoryReservation;
        public readonly bool? MemoryReservationLockedToMax;
        public readonly int MemoryShareCount;
        public readonly string? MemoryShareLevel;
        public readonly string Moid;
        public readonly string? Name;
        public readonly bool? NestedHvEnabled;
        /// <summary>
        /// The network interface types for each network
        /// interface found on the virtual machine, in device bus order. Will be one of
        /// `e1000`, `e1000e`, `pcnet32`, `sriov`, `vmxnet2`, `vmxnet3vrdma`, or `vmxnet3`.
        /// </summary>
        public readonly ImmutableArray<string> NetworkInterfaceTypes;
        /// <summary>
        /// Information about each of the network interfaces on this 
        /// virtual machine or template. These are sorted by device bus order so that they
        /// can be applied to a `vsphere.VirtualMachine` resource in the order the resource
        /// expects while cloning. This is useful for discovering certain network interface
        /// settings while performing a linked clone, as all settings that are output by this
        /// data source must be the same on the destination virtual machine as the source.
        /// The sub-attributes are:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualMachineNetworkInterfaceResult> NetworkInterfaces;
        /// <summary>
        /// The number of cores per socket for this virtual machine.
        /// </summary>
        public readonly int? NumCoresPerSocket;
        /// <summary>
        /// The total number of virtual processor cores assigned to this
        /// virtual machine.
        /// </summary>
        public readonly int? NumCpus;
        public readonly string? ReplaceTrigger;
        public readonly bool? RunToolsScriptsAfterPowerOn;
        public readonly bool? RunToolsScriptsAfterResume;
        public readonly bool? RunToolsScriptsBeforeGuestReboot;
        public readonly bool? RunToolsScriptsBeforeGuestShutdown;
        public readonly bool? RunToolsScriptsBeforeGuestStandby;
        public readonly int? SataControllerScanCount;
        /// <summary>
        /// Mode for sharing the SCSI bus. The modes are
        /// physicalSharing, virtualSharing, and noSharing. Only the first number of
        /// controllers defined by `scsi_controller_scan_count` are scanned.
        /// </summary>
        public readonly string ScsiBusSharing;
        public readonly int? ScsiControllerScanCount;
        /// <summary>
        /// The common type of all SCSI controllers on this virtual machine.
        /// Will be one of `lsilogic` (LSI Logic Parallel), `lsilogic-sas` (LSI Logic
        /// SAS), `pvscsi` (VMware Paravirtual), `buslogic` (BusLogic), or `mixed` when
        /// there are multiple controller types. Only the first number of controllers
        /// defined by `scsi_controller_scan_count` are scanned.
        /// </summary>
        public readonly string ScsiType;
        public readonly string StoragePolicyId;
        public readonly string? SwapPlacementPolicy;
        public readonly bool? SyncTimeWithHost;
        public readonly bool? SyncTimeWithHostPeriodically;
        public readonly string? ToolsUpgradePolicy;
        public readonly string Uuid;
        public readonly Outputs.GetVirtualMachineVappResult? Vapp;
        public readonly ImmutableArray<string> VappTransports;
        public readonly bool? VbsEnabled;
        public readonly bool? VvtdEnabled;

        [OutputConstructor]
        private GetVirtualMachineResult(
            string? alternateGuestName,

            string annotation,

            int? bootDelay,

            int? bootRetryDelay,

            bool? bootRetryEnabled,

            string changeVersion,

            bool? cpuHotAddEnabled,

            bool? cpuHotRemoveEnabled,

            int? cpuLimit,

            bool? cpuPerformanceCountersEnabled,

            int? cpuReservation,

            int cpuShareCount,

            string? cpuShareLevel,

            string? datacenterId,

            string defaultIpAddress,

            ImmutableArray<Outputs.GetVirtualMachineDiskResult> disks,

            bool? efiSecureBootEnabled,

            bool? enableDiskUuid,

            bool? enableLogging,

            string? eptRviMode,

            ImmutableDictionary<string, string>? extraConfig,

            bool? extraConfigRebootRequired,

            string? firmware,

            string? folder,

            string guestId,

            ImmutableArray<string> guestIpAddresses,

            int hardwareVersion,

            string? hvMode,

            string id,

            int? ideControllerScanCount,

            string? latencySensitivity,

            int? memory,

            bool? memoryHotAddEnabled,

            int? memoryLimit,

            int? memoryReservation,

            bool? memoryReservationLockedToMax,

            int memoryShareCount,

            string? memoryShareLevel,

            string moid,

            string? name,

            bool? nestedHvEnabled,

            ImmutableArray<string> networkInterfaceTypes,

            ImmutableArray<Outputs.GetVirtualMachineNetworkInterfaceResult> networkInterfaces,

            int? numCoresPerSocket,

            int? numCpus,

            string? replaceTrigger,

            bool? runToolsScriptsAfterPowerOn,

            bool? runToolsScriptsAfterResume,

            bool? runToolsScriptsBeforeGuestReboot,

            bool? runToolsScriptsBeforeGuestShutdown,

            bool? runToolsScriptsBeforeGuestStandby,

            int? sataControllerScanCount,

            string scsiBusSharing,

            int? scsiControllerScanCount,

            string scsiType,

            string storagePolicyId,

            string? swapPlacementPolicy,

            bool? syncTimeWithHost,

            bool? syncTimeWithHostPeriodically,

            string? toolsUpgradePolicy,

            string uuid,

            Outputs.GetVirtualMachineVappResult? vapp,

            ImmutableArray<string> vappTransports,

            bool? vbsEnabled,

            bool? vvtdEnabled)
        {
            AlternateGuestName = alternateGuestName;
            Annotation = annotation;
            BootDelay = bootDelay;
            BootRetryDelay = bootRetryDelay;
            BootRetryEnabled = bootRetryEnabled;
            ChangeVersion = changeVersion;
            CpuHotAddEnabled = cpuHotAddEnabled;
            CpuHotRemoveEnabled = cpuHotRemoveEnabled;
            CpuLimit = cpuLimit;
            CpuPerformanceCountersEnabled = cpuPerformanceCountersEnabled;
            CpuReservation = cpuReservation;
            CpuShareCount = cpuShareCount;
            CpuShareLevel = cpuShareLevel;
            DatacenterId = datacenterId;
            DefaultIpAddress = defaultIpAddress;
            Disks = disks;
            EfiSecureBootEnabled = efiSecureBootEnabled;
            EnableDiskUuid = enableDiskUuid;
            EnableLogging = enableLogging;
            EptRviMode = eptRviMode;
            ExtraConfig = extraConfig;
            ExtraConfigRebootRequired = extraConfigRebootRequired;
            Firmware = firmware;
            Folder = folder;
            GuestId = guestId;
            GuestIpAddresses = guestIpAddresses;
            HardwareVersion = hardwareVersion;
            HvMode = hvMode;
            Id = id;
            IdeControllerScanCount = ideControllerScanCount;
            LatencySensitivity = latencySensitivity;
            Memory = memory;
            MemoryHotAddEnabled = memoryHotAddEnabled;
            MemoryLimit = memoryLimit;
            MemoryReservation = memoryReservation;
            MemoryReservationLockedToMax = memoryReservationLockedToMax;
            MemoryShareCount = memoryShareCount;
            MemoryShareLevel = memoryShareLevel;
            Moid = moid;
            Name = name;
            NestedHvEnabled = nestedHvEnabled;
            NetworkInterfaceTypes = networkInterfaceTypes;
            NetworkInterfaces = networkInterfaces;
            NumCoresPerSocket = numCoresPerSocket;
            NumCpus = numCpus;
            ReplaceTrigger = replaceTrigger;
            RunToolsScriptsAfterPowerOn = runToolsScriptsAfterPowerOn;
            RunToolsScriptsAfterResume = runToolsScriptsAfterResume;
            RunToolsScriptsBeforeGuestReboot = runToolsScriptsBeforeGuestReboot;
            RunToolsScriptsBeforeGuestShutdown = runToolsScriptsBeforeGuestShutdown;
            RunToolsScriptsBeforeGuestStandby = runToolsScriptsBeforeGuestStandby;
            SataControllerScanCount = sataControllerScanCount;
            ScsiBusSharing = scsiBusSharing;
            ScsiControllerScanCount = scsiControllerScanCount;
            ScsiType = scsiType;
            StoragePolicyId = storagePolicyId;
            SwapPlacementPolicy = swapPlacementPolicy;
            SyncTimeWithHost = syncTimeWithHost;
            SyncTimeWithHostPeriodically = syncTimeWithHostPeriodically;
            ToolsUpgradePolicy = toolsUpgradePolicy;
            Uuid = uuid;
            Vapp = vapp;
            VappTransports = vappTransports;
            VbsEnabled = vbsEnabled;
            VvtdEnabled = vvtdEnabled;
        }
    }
}

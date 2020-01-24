// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static partial class Invokes
    {
        /// <summary>
        /// The `vsphere..VirtualMachine` data source can be used to find the UUID of an
        /// existing virtual machine or template. Its most relevant purpose is for finding
        /// the UUID of a template to be used as the source for cloning into a new
        /// [`vsphere..VirtualMachine`][docs-virtual-machine-resource] resource. It also
        /// reads the guest ID so that can be supplied as well.
        /// 
        /// [docs-virtual-machine-resource]: /docs/providers/vsphere/r/virtual_machine.html
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/virtual_machine.html.markdown.
        /// </summary>
        public static Task<GetVirtualMachineResult> GetVirtualMachine(GetVirtualMachineArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualMachineResult>("vsphere:index/getVirtualMachine:getVirtualMachine", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetVirtualMachineArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [managed object reference
        /// ID][docs-about-morefs] of the datacenter the virtual machine is located in.
        /// This can be omitted if the search path used in `name` is an absolute path.
        /// For default datacenters, use the `id` attribute from an empty
        /// `vsphere..Datacenter` data source.
        /// </summary>
        [Input("datacenterId")]
        public Input<string>? DatacenterId { get; set; }

        /// <summary>
        /// The name of the virtual machine. This can be a name or
        /// path.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The number of SCSI controllers to
        /// scan for disk attributes and controller types on. Default: `1`.
        /// </summary>
        [Input("scsiControllerScanCount")]
        public Input<int>? ScsiControllerScanCount { get; set; }

        public GetVirtualMachineArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetVirtualMachineResult
    {
        /// <summary>
        /// The alternate guest name of the virtual machine when
        /// guest_id is a non-specific operating system, like `otherGuest`.
        /// </summary>
        public readonly string AlternateGuestName;
        public readonly string? DatacenterId;
        /// <summary>
        /// Information about each of the disks on this virtual machine or
        /// template. These are sorted by bus and unit number so that they can be applied
        /// to a `vsphere..VirtualMachine` resource in the order the resource expects
        /// while cloning. This is useful for discovering certain disk settings while
        /// performing a linked clone, as all settings that are output by this data
        /// source must be the same on the destination virtual machine as the source.
        /// Only the first number of controllers defined by `scsi_controller_scan_count`
        /// are scanned for disks. The sub-attributes are:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVirtualMachineDisksResult> Disks;
        /// <summary>
        /// The firmware type for this virtual machine. Can be `bios` or `efi`.
        /// </summary>
        public readonly string Firmware;
        /// <summary>
        /// The guest ID of the virtual machine or template.
        /// </summary>
        public readonly string GuestId;
        /// <summary>
        /// A list of IP addresses as reported by VMWare tools.
        /// </summary>
        public readonly ImmutableArray<string> GuestIpAddresses;
        public readonly string Name;
        /// <summary>
        /// The network interface types for each network
        /// interface found on the virtual machine, in device bus order. Will be one of
        /// `e1000`, `e1000e`, `pcnet32`, `sriov`, `vmxnet2`, or `vmxnet3`.
        /// </summary>
        public readonly ImmutableArray<string> NetworkInterfaceTypes;
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
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetVirtualMachineResult(
            string alternateGuestName,
            string? datacenterId,
            ImmutableArray<Outputs.GetVirtualMachineDisksResult> disks,
            string firmware,
            string guestId,
            ImmutableArray<string> guestIpAddresses,
            string name,
            ImmutableArray<string> networkInterfaceTypes,
            string scsiBusSharing,
            int? scsiControllerScanCount,
            string scsiType,
            string id)
        {
            AlternateGuestName = alternateGuestName;
            DatacenterId = datacenterId;
            Disks = disks;
            Firmware = firmware;
            GuestId = guestId;
            GuestIpAddresses = guestIpAddresses;
            Name = name;
            NetworkInterfaceTypes = networkInterfaceTypes;
            ScsiBusSharing = scsiBusSharing;
            ScsiControllerScanCount = scsiControllerScanCount;
            ScsiType = scsiType;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetVirtualMachineDisksResult
    {
        /// <summary>
        /// Set to `true` if the disk has been eager zeroed.
        /// </summary>
        public readonly bool EagerlyScrub;
        /// <summary>
        /// The size of the disk, in GIB.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// Set to `true` if the disk has been thin provisioned.
        /// </summary>
        public readonly bool ThinProvisioned;

        [OutputConstructor]
        private GetVirtualMachineDisksResult(
            bool eagerlyScrub,
            int size,
            bool thinProvisioned)
        {
            EagerlyScrub = eagerlyScrub;
            Size = size;
            ThinProvisioned = thinProvisioned;
        }
    }
    }
}

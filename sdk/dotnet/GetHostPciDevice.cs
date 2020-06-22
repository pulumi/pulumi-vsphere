// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetHostPciDevice
    {
        /// <summary>
        /// The `vsphere..getHostPciDevice` data source can be used to discover the DeviceID
        /// of a vSphere host's PCI device. This can then be used with 
        /// `vsphere..VirtualMachine`'s `pci_device_id`.
        /// 
        /// ## Example Usage With Name Regular Expression
        ///  
        ///  ```hcl
        ///  data "vsphere..Datacenter" "datacenter" {
        ///    name = "dc1"
        ///  }
        ///  
        ///  data "vsphere..Host" "host" {
        ///    name          = "esxi1"
        ///    datacenter_id = data.vsphere_datacenter.datacenter.id
        ///  }
        ///  
        ///  data "vsphere..getHostPciDevice" "dev" {
        ///    host_id    = data.vsphere_host.host.id
        ///    name_regex = "MMC"
        ///  }
        ///  ```
        /// </summary>
        public static Task<GetHostPciDeviceResult> InvokeAsync(GetHostPciDeviceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHostPciDeviceResult>("vsphere:index/getHostPciDevice:getHostPciDevice", args ?? new GetHostPciDeviceArgs(), options.WithVersion());
    }


    public sealed class GetHostPciDeviceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The hexadecimal PCI device class ID
        /// </summary>
        [Input("classId")]
        public string? ClassId { get; set; }

        /// <summary>
        /// The [managed object reference
        /// ID][docs-about-morefs] of a host.
        /// </summary>
        [Input("hostId", required: true)]
        public string HostId { get; set; } = null!;

        /// <summary>
        /// A regular expression that will be used to match
        /// the host PCI device name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// The hexadecimal PCI device vendor ID.
        /// </summary>
        [Input("vendorId")]
        public string? VendorId { get; set; }

        public GetHostPciDeviceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetHostPciDeviceResult
    {
        public readonly string? ClassId;
        public readonly string HostId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the PCI device.
        /// </summary>
        public readonly string Name;
        public readonly string? NameRegex;
        public readonly string? VendorId;

        [OutputConstructor]
        private GetHostPciDeviceResult(
            string? classId,

            string hostId,

            string id,

            string name,

            string? nameRegex,

            string? vendorId)
        {
            ClassId = classId;
            HostId = hostId;
            Id = id;
            Name = name;
            NameRegex = nameRegex;
            VendorId = vendorId;
        }
    }
}

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
        /// The `vsphere.getHostPciDevice` data source can be used to discover the device ID
        /// of a vSphere host's PCI device. This can then be used with
        /// `vsphere.VirtualMachine`'s `pci_device_id`.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### With Vendor ID And Class ID
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
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
        ///     var host = VSphere.GetHost.Invoke(new()
        ///     {
        ///         Name = "esxi-01.example.com",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        ///     var dev = VSphere.GetHostPciDevice.Invoke(new()
        ///     {
        ///         HostId = host.Apply(getHostResult =&gt; getHostResult.Id),
        ///         ClassId = "123",
        ///         VendorId = "456",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// ### With Name Regular Expression
        ///  
        ///  ```hcl
        ///  data "vsphere_datacenter" "datacenter" {
        ///    name = "dc-01"
        ///  }
        ///  
        ///  data "vsphere_host" "host" {
        ///    name          = "esxi-01.example.com"
        ///    datacenter_id = data.vsphere_datacenter.datacenter.id
        ///  }
        ///  
        ///  data "vsphere_host_pci_device" "dev" {
        ///    host_id    = data.vsphere_host.host.id
        ///    name_regex = "MMC"
        ///  }
        ///  ```
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetHostPciDeviceResult> InvokeAsync(GetHostPciDeviceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHostPciDeviceResult>("vsphere:index/getHostPciDevice:getHostPciDevice", args ?? new GetHostPciDeviceArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.getHostPciDevice` data source can be used to discover the device ID
        /// of a vSphere host's PCI device. This can then be used with
        /// `vsphere.VirtualMachine`'s `pci_device_id`.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### With Vendor ID And Class ID
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
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
        ///     var host = VSphere.GetHost.Invoke(new()
        ///     {
        ///         Name = "esxi-01.example.com",
        ///         DatacenterId = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Id),
        ///     });
        /// 
        ///     var dev = VSphere.GetHostPciDevice.Invoke(new()
        ///     {
        ///         HostId = host.Apply(getHostResult =&gt; getHostResult.Id),
        ///         ClassId = "123",
        ///         VendorId = "456",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// ### With Name Regular Expression
        ///  
        ///  ```hcl
        ///  data "vsphere_datacenter" "datacenter" {
        ///    name = "dc-01"
        ///  }
        ///  
        ///  data "vsphere_host" "host" {
        ///    name          = "esxi-01.example.com"
        ///    datacenter_id = data.vsphere_datacenter.datacenter.id
        ///  }
        ///  
        ///  data "vsphere_host_pci_device" "dev" {
        ///    host_id    = data.vsphere_host.host.id
        ///    name_regex = "MMC"
        ///  }
        ///  ```
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetHostPciDeviceResult> Invoke(GetHostPciDeviceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHostPciDeviceResult>("vsphere:index/getHostPciDevice:getHostPciDevice", args ?? new GetHostPciDeviceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHostPciDeviceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The hexadecimal PCI device class ID
        /// </summary>
        [Input("classId")]
        public string? ClassId { get; set; }

        /// <summary>
        /// The [managed object reference ID][docs-about-morefs] of a host.
        /// </summary>
        [Input("hostId", required: true)]
        public string HostId { get; set; } = null!;

        /// <summary>
        /// A regular expression that will be used to match the
        /// host PCI device name.
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
        public static new GetHostPciDeviceArgs Empty => new GetHostPciDeviceArgs();
    }

    public sealed class GetHostPciDeviceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The hexadecimal PCI device class ID
        /// </summary>
        [Input("classId")]
        public Input<string>? ClassId { get; set; }

        /// <summary>
        /// The [managed object reference ID][docs-about-morefs] of a host.
        /// </summary>
        [Input("hostId", required: true)]
        public Input<string> HostId { get; set; } = null!;

        /// <summary>
        /// A regular expression that will be used to match the
        /// host PCI device name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// The hexadecimal PCI device vendor ID.
        /// </summary>
        [Input("vendorId")]
        public Input<string>? VendorId { get; set; }

        public GetHostPciDeviceInvokeArgs()
        {
        }
        public static new GetHostPciDeviceInvokeArgs Empty => new GetHostPciDeviceInvokeArgs();
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

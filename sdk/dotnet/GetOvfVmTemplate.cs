// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    public static class GetOvfVmTemplate
    {
        /// <summary>
        /// The `vsphere.getOvfVmTemplate` data source can be used to submit an OVF to vSphere and extract its hardware
        /// settings in a form that can be then used as inputs for a `vsphere.VirtualMachine` resource.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using VSphere = Pulumi.VSphere;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var ovf = Output.Create(VSphere.GetOvfVmTemplate.InvokeAsync(new VSphere.GetOvfVmTemplateArgs
        ///         {
        ///             Name = "testOVF",
        ///             ResourcePoolId = vsphere_resource_pool.Rp.Id,
        ///             DatastoreId = data.Vsphere_datastore.Ds.Id,
        ///             HostSystemId = data.Vsphere_host.Hs.Id,
        ///             RemoteOvfUrl = "https://download3.vmware.com/software/vmw-tools/nested-esxi/Nested_ESXi7.0_Appliance_Template_v1.ova",
        ///             OvfNetworkMap = 
        ///             {
        ///                 { "Network 1", data.Vsphere_network.Net.Id },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOvfVmTemplateResult> InvokeAsync(GetOvfVmTemplateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOvfVmTemplateResult>("vsphere:index/getOvfVmTemplate:getOvfVmTemplate", args ?? new GetOvfVmTemplateArgs(), options.WithVersion());
    }


    public sealed class GetOvfVmTemplateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Allow unverified ssl certificates while deploying ovf/ova from url.
        /// </summary>
        [Input("allowUnverifiedSslCert")]
        public bool? AllowUnverifiedSslCert { get; set; }

        /// <summary>
        /// The ID of the virtual machine's datastore. The virtual machine configuration is placed here, along with any virtual disks that are created without datastores.
        /// </summary>
        [Input("datastoreId")]
        public string? DatastoreId { get; set; }

        /// <summary>
        /// The key of the chosen deployment option. If empty, the default option is chosen.
        /// </summary>
        [Input("deploymentOption")]
        public string? DeploymentOption { get; set; }

        /// <summary>
        /// The disk provisioning. If set, all the disks in the deployed OVF will have
        /// the same specified disk type (accepted values {thin, flat, thick, sameAsSource}).
        /// </summary>
        [Input("diskProvisioning")]
        public string? DiskProvisioning { get; set; }

        /// <summary>
        /// The name of the folder to locate the virtual machine in.
        /// </summary>
        [Input("folder")]
        public string? Folder { get; set; }

        /// <summary>
        /// The ID of an optional host system to pin the virtual machine to.
        /// </summary>
        [Input("hostSystemId", required: true)]
        public string HostSystemId { get; set; } = null!;

        /// <summary>
        /// The IP allocation policy.
        /// </summary>
        [Input("ipAllocationPolicy")]
        public string? IpAllocationPolicy { get; set; }

        /// <summary>
        /// The IP protocol.
        /// </summary>
        [Input("ipProtocol")]
        public string? IpProtocol { get; set; }

        /// <summary>
        /// The absolute path to the ovf/ova file in the local system. While deploying from ovf,
        /// make sure the other necessary files like the .vmdk files are also in the same directory as the given ovf file.
        /// </summary>
        [Input("localOvfPath")]
        public string? LocalOvfPath { get; set; }

        /// <summary>
        /// Name of the virtual machine to create.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("ovfNetworkMap")]
        private Dictionary<string, string>? _ovfNetworkMap;

        /// <summary>
        /// The mapping of name of network identifiers from the ovf descriptor to network UUID in the
        /// VI infrastructure.
        /// </summary>
        public Dictionary<string, string> OvfNetworkMap
        {
            get => _ovfNetworkMap ?? (_ovfNetworkMap = new Dictionary<string, string>());
            set => _ovfNetworkMap = value;
        }

        /// <summary>
        /// URL to the remote ovf/ova file to be deployed.
        /// </summary>
        [Input("remoteOvfUrl")]
        public string? RemoteOvfUrl { get; set; }

        /// <summary>
        /// The ID of a resource pool to put the virtual machine in.
        /// </summary>
        [Input("resourcePoolId", required: true)]
        public string ResourcePoolId { get; set; } = null!;

        public GetOvfVmTemplateArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOvfVmTemplateResult
    {
        public readonly bool? AllowUnverifiedSslCert;
        /// <summary>
        /// The guest name for the operating system .
        /// </summary>
        public readonly string AlternateGuestName;
        /// <summary>
        /// User-provided description of the virtual machine.
        /// </summary>
        public readonly string Annotation;
        /// <summary>
        /// Allow CPUs to be added to this virtual machine while it is running.
        /// </summary>
        public readonly bool CpuHotAddEnabled;
        /// <summary>
        /// Allow CPUs to be added to this virtual machine while it is running.
        /// </summary>
        public readonly bool CpuHotRemoveEnabled;
        public readonly bool CpuPerformanceCountersEnabled;
        public readonly string? DatastoreId;
        public readonly string? DeploymentOption;
        public readonly string? DiskProvisioning;
        /// <summary>
        /// The firmware interface to use on the virtual machine.
        /// </summary>
        public readonly string Firmware;
        public readonly string? Folder;
        /// <summary>
        /// The guest ID for the operating system
        /// </summary>
        public readonly string GuestId;
        public readonly string HostSystemId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int IdeControllerCount;
        public readonly string? IpAllocationPolicy;
        public readonly string? IpProtocol;
        public readonly string? LocalOvfPath;
        /// <summary>
        /// The size of the virtual machine's memory, in MB.
        /// </summary>
        public readonly int Memory;
        /// <summary>
        /// Allow memory to be added to this virtual machine while it is running.
        /// </summary>
        public readonly bool MemoryHotAddEnabled;
        public readonly string Name;
        /// <summary>
        /// Enable nested hardware virtualization on this virtual machine, facilitating nested virtualization in the guest.
        /// </summary>
        public readonly bool NestedHvEnabled;
        /// <summary>
        /// The number of cores to distribute amongst the CPUs in this virtual machine.
        /// </summary>
        public readonly int NumCoresPerSocket;
        /// <summary>
        /// The number of virtual processors to assign to this virtual machine.
        /// </summary>
        public readonly int NumCpus;
        public readonly ImmutableDictionary<string, string>? OvfNetworkMap;
        public readonly string? RemoteOvfUrl;
        public readonly string ResourcePoolId;
        public readonly int SataControllerCount;
        public readonly int ScsiControllerCount;
        public readonly string ScsiType;
        /// <summary>
        /// The swap file placement policy for this virtual machine.
        /// </summary>
        public readonly string SwapPlacementPolicy;

        [OutputConstructor]
        private GetOvfVmTemplateResult(
            bool? allowUnverifiedSslCert,

            string alternateGuestName,

            string annotation,

            bool cpuHotAddEnabled,

            bool cpuHotRemoveEnabled,

            bool cpuPerformanceCountersEnabled,

            string? datastoreId,

            string? deploymentOption,

            string? diskProvisioning,

            string firmware,

            string? folder,

            string guestId,

            string hostSystemId,

            string id,

            int ideControllerCount,

            string? ipAllocationPolicy,

            string? ipProtocol,

            string? localOvfPath,

            int memory,

            bool memoryHotAddEnabled,

            string name,

            bool nestedHvEnabled,

            int numCoresPerSocket,

            int numCpus,

            ImmutableDictionary<string, string>? ovfNetworkMap,

            string? remoteOvfUrl,

            string resourcePoolId,

            int sataControllerCount,

            int scsiControllerCount,

            string scsiType,

            string swapPlacementPolicy)
        {
            AllowUnverifiedSslCert = allowUnverifiedSslCert;
            AlternateGuestName = alternateGuestName;
            Annotation = annotation;
            CpuHotAddEnabled = cpuHotAddEnabled;
            CpuHotRemoveEnabled = cpuHotRemoveEnabled;
            CpuPerformanceCountersEnabled = cpuPerformanceCountersEnabled;
            DatastoreId = datastoreId;
            DeploymentOption = deploymentOption;
            DiskProvisioning = diskProvisioning;
            Firmware = firmware;
            Folder = folder;
            GuestId = guestId;
            HostSystemId = hostSystemId;
            Id = id;
            IdeControllerCount = ideControllerCount;
            IpAllocationPolicy = ipAllocationPolicy;
            IpProtocol = ipProtocol;
            LocalOvfPath = localOvfPath;
            Memory = memory;
            MemoryHotAddEnabled = memoryHotAddEnabled;
            Name = name;
            NestedHvEnabled = nestedHvEnabled;
            NumCoresPerSocket = numCoresPerSocket;
            NumCpus = numCpus;
            OvfNetworkMap = ovfNetworkMap;
            RemoteOvfUrl = remoteOvfUrl;
            ResourcePoolId = resourcePoolId;
            SataControllerCount = sataControllerCount;
            ScsiControllerCount = scsiControllerCount;
            ScsiType = scsiType;
            SwapPlacementPolicy = swapPlacementPolicy;
        }
    }
}

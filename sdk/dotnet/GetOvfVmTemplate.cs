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
        /// The `vsphere.getOvfVmTemplate` data source can be used to submit an OVF to
        /// vSphere and extract its hardware settings in a form that can be then used as
        /// inputs for a `vsphere.VirtualMachine` resource.
        /// </summary>
        public static Task<GetOvfVmTemplateResult> InvokeAsync(GetOvfVmTemplateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOvfVmTemplateResult>("vsphere:index/getOvfVmTemplate:getOvfVmTemplate", args ?? new GetOvfVmTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// The `vsphere.getOvfVmTemplate` data source can be used to submit an OVF to
        /// vSphere and extract its hardware settings in a form that can be then used as
        /// inputs for a `vsphere.VirtualMachine` resource.
        /// </summary>
        public static Output<GetOvfVmTemplateResult> Invoke(GetOvfVmTemplateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOvfVmTemplateResult>("vsphere:index/getOvfVmTemplate:getOvfVmTemplate", args ?? new GetOvfVmTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOvfVmTemplateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Allow unverified SSL certificates
        /// when deploying OVF/OVA from a URL.
        /// </summary>
        [Input("allowUnverifiedSslCert")]
        public bool? AllowUnverifiedSslCert { get; set; }

        /// <summary>
        /// The ID of the virtual machine's datastore. The
        /// virtual machine configuration is placed here, along with any virtual disks
        /// that are created without datastores.
        /// </summary>
        [Input("datastoreId")]
        public string? DatastoreId { get; set; }

        /// <summary>
        /// The key of the chosen deployment option. If
        /// empty, the default option is chosen.
        /// </summary>
        [Input("deploymentOption")]
        public string? DeploymentOption { get; set; }

        /// <summary>
        /// The disk provisioning type. If set, all the
        /// disks in the deployed OVA/OVF will have the same specified disk type. Can be
        /// one of `thin`, `flat`, `thick` or `sameAsSource`.
        /// </summary>
        [Input("diskProvisioning")]
        public string? DiskProvisioning { get; set; }

        /// <summary>
        /// Allow properties with
        /// `ovf:userConfigurable=false` to be set.
        /// </summary>
        [Input("enableHiddenProperties")]
        public bool? EnableHiddenProperties { get; set; }

        /// <summary>
        /// The name of the folder in which to place the virtual
        /// machine.
        /// </summary>
        [Input("folder")]
        public string? Folder { get; set; }

        /// <summary>
        /// The ID of the ESXi host system to deploy the
        /// virtual machine.
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
        /// The absolute path to the OVF/OVA file on the
        /// local system. When deploying from an OVF, ensure all necessary files such as
        /// the `.vmdk` files are present in the same directory as the OVF.
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
        /// The mapping of name of network identifiers
        /// from the OVF descriptor to network UUID in the environment.
        /// </summary>
        public Dictionary<string, string> OvfNetworkMap
        {
            get => _ovfNetworkMap ?? (_ovfNetworkMap = new Dictionary<string, string>());
            set => _ovfNetworkMap = value;
        }

        /// <summary>
        /// URL of the remote OVF/OVA file to be deployed.
        /// 
        /// &gt; **NOTE:** Either `local_ovf_path` or `remote_ovf_url` is required, both can
        /// not be empty.
        /// </summary>
        [Input("remoteOvfUrl")]
        public string? RemoteOvfUrl { get; set; }

        /// <summary>
        /// The ID of a resource pool in which to place
        /// the virtual machine.
        /// </summary>
        [Input("resourcePoolId", required: true)]
        public string ResourcePoolId { get; set; } = null!;

        public GetOvfVmTemplateArgs()
        {
        }
        public static new GetOvfVmTemplateArgs Empty => new GetOvfVmTemplateArgs();
    }

    public sealed class GetOvfVmTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Allow unverified SSL certificates
        /// when deploying OVF/OVA from a URL.
        /// </summary>
        [Input("allowUnverifiedSslCert")]
        public Input<bool>? AllowUnverifiedSslCert { get; set; }

        /// <summary>
        /// The ID of the virtual machine's datastore. The
        /// virtual machine configuration is placed here, along with any virtual disks
        /// that are created without datastores.
        /// </summary>
        [Input("datastoreId")]
        public Input<string>? DatastoreId { get; set; }

        /// <summary>
        /// The key of the chosen deployment option. If
        /// empty, the default option is chosen.
        /// </summary>
        [Input("deploymentOption")]
        public Input<string>? DeploymentOption { get; set; }

        /// <summary>
        /// The disk provisioning type. If set, all the
        /// disks in the deployed OVA/OVF will have the same specified disk type. Can be
        /// one of `thin`, `flat`, `thick` or `sameAsSource`.
        /// </summary>
        [Input("diskProvisioning")]
        public Input<string>? DiskProvisioning { get; set; }

        /// <summary>
        /// Allow properties with
        /// `ovf:userConfigurable=false` to be set.
        /// </summary>
        [Input("enableHiddenProperties")]
        public Input<bool>? EnableHiddenProperties { get; set; }

        /// <summary>
        /// The name of the folder in which to place the virtual
        /// machine.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// The ID of the ESXi host system to deploy the
        /// virtual machine.
        /// </summary>
        [Input("hostSystemId", required: true)]
        public Input<string> HostSystemId { get; set; } = null!;

        /// <summary>
        /// The IP allocation policy.
        /// </summary>
        [Input("ipAllocationPolicy")]
        public Input<string>? IpAllocationPolicy { get; set; }

        /// <summary>
        /// The IP protocol.
        /// </summary>
        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        /// <summary>
        /// The absolute path to the OVF/OVA file on the
        /// local system. When deploying from an OVF, ensure all necessary files such as
        /// the `.vmdk` files are present in the same directory as the OVF.
        /// </summary>
        [Input("localOvfPath")]
        public Input<string>? LocalOvfPath { get; set; }

        /// <summary>
        /// Name of the virtual machine to create.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("ovfNetworkMap")]
        private InputMap<string>? _ovfNetworkMap;

        /// <summary>
        /// The mapping of name of network identifiers
        /// from the OVF descriptor to network UUID in the environment.
        /// </summary>
        public InputMap<string> OvfNetworkMap
        {
            get => _ovfNetworkMap ?? (_ovfNetworkMap = new InputMap<string>());
            set => _ovfNetworkMap = value;
        }

        /// <summary>
        /// URL of the remote OVF/OVA file to be deployed.
        /// 
        /// &gt; **NOTE:** Either `local_ovf_path` or `remote_ovf_url` is required, both can
        /// not be empty.
        /// </summary>
        [Input("remoteOvfUrl")]
        public Input<string>? RemoteOvfUrl { get; set; }

        /// <summary>
        /// The ID of a resource pool in which to place
        /// the virtual machine.
        /// </summary>
        [Input("resourcePoolId", required: true)]
        public Input<string> ResourcePoolId { get; set; } = null!;

        public GetOvfVmTemplateInvokeArgs()
        {
        }
        public static new GetOvfVmTemplateInvokeArgs Empty => new GetOvfVmTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetOvfVmTemplateResult
    {
        public readonly bool? AllowUnverifiedSslCert;
        /// <summary>
        /// An alternate guest operating system name.
        /// </summary>
        public readonly string AlternateGuestName;
        /// <summary>
        /// A description of the virtual machine.
        /// </summary>
        public readonly string Annotation;
        /// <summary>
        /// Allow CPUs to be added to the virtual machine while
        /// powered on.
        /// </summary>
        public readonly bool CpuHotAddEnabled;
        /// <summary>
        /// Allow CPUs to be removed from the virtual machine
        /// while powered on.
        /// </summary>
        public readonly bool CpuHotRemoveEnabled;
        public readonly bool CpuPerformanceCountersEnabled;
        public readonly string? DatastoreId;
        public readonly string? DeploymentOption;
        public readonly string? DiskProvisioning;
        public readonly bool? EnableHiddenProperties;
        /// <summary>
        /// The firmware to use on the virtual machine.
        /// </summary>
        public readonly string Firmware;
        public readonly string? Folder;
        /// <summary>
        /// The ID for the guest operating system
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
        /// The size of the virtual machine memory, in MB.
        /// </summary>
        public readonly int Memory;
        /// <summary>
        /// Allow memory to be added to the virtual machine
        /// while powered on.
        /// </summary>
        public readonly bool MemoryHotAddEnabled;
        public readonly string Name;
        /// <summary>
        /// Enable nested hardware virtualization on the virtual
        /// machine, facilitating nested virtualization in the guest.
        /// </summary>
        public readonly bool NestedHvEnabled;
        /// <summary>
        /// The number of cores per virtual CPU in the virtual
        /// machine.
        /// </summary>
        public readonly int NumCoresPerSocket;
        /// <summary>
        /// The number of virtual CPUs to assign to the virtual machine.
        /// </summary>
        public readonly int NumCpus;
        public readonly ImmutableDictionary<string, string>? OvfNetworkMap;
        public readonly string? RemoteOvfUrl;
        public readonly string ResourcePoolId;
        public readonly int SataControllerCount;
        public readonly int ScsiControllerCount;
        public readonly string ScsiType;
        /// <summary>
        /// The swap file placement policy for the virtual
        /// machine.
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

            bool? enableHiddenProperties,

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
            EnableHiddenProperties = enableHiddenProperties;
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

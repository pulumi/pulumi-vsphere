// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    /// <summary>
    /// The `vsphere.VirtualDisk` resource can be used to create virtual disks outside
    /// of any given `vsphere.VirtualMachine`
    /// resource. These disks can be attached to a virtual machine by creating a disk
    /// block with the `attach` parameter.
    /// 
    /// ## Example Usage
    /// 
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
    ///     var datastore = VSphere.GetDatacenter.Invoke(new()
    ///     {
    ///         Name = "datastore-01",
    ///     });
    /// 
    ///     var virtualDisk = new VSphere.VirtualDisk("virtual_disk", new()
    ///     {
    ///         Size = 40,
    ///         Type = "thin",
    ///         VmdkPath = "/foo/foo.vmdk",
    ///         CreateDirectories = true,
    ///         Datacenter = datacenter.Apply(getDatacenterResult =&gt; getDatacenterResult.Name),
    ///         Datastore = datastoreVsphereDatastore.Name,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// An existing virtual disk can be imported into this resource
    /// 
    /// via supplying the full datastore path to the virtual disk. An example is below:
    /// 
    /// [docs-import]: https://developer.hashicorp.com/terraform/cli/import
    /// 
    /// ```sh
    /// $ pulumi import vsphere:index/virtualDisk:VirtualDisk virtual_disk \
    /// ```
    /// 
    ///   '{"virtual_disk_path": "/dc-01/[datastore-01]foo/bar.vmdk", \ "create_directories": "true"}'
    /// 
    /// The above would import the virtual disk located at `foo/bar.vmdk` in the `datastore-01`
    /// 
    /// datastore of the `dc-01` datacenter with `create_directories` set as `true`.
    /// </summary>
    [VSphereResourceType("vsphere:index/virtualDisk:VirtualDisk")]
    public partial class VirtualDisk : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The adapter type for this virtual disk. Can be
        /// one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
        /// 
        /// &gt; **NOTE:** `adapter_type` is **deprecated**: it does not dictate the type of
        /// controller that the virtual disk will be attached to on the virtual machine.
        /// Please see the `scsi_type` parameter
        /// in the `vsphere.VirtualMachine` resource for information on how to control
        /// disk controller types. This parameter will be removed in future versions of the
        /// vSphere provider.
        /// </summary>
        [Output("adapterType")]
        public Output<string?> AdapterType { get; private set; } = null!;

        /// <summary>
        /// Tells the resource to create any
        /// directories that are a part of the `vmdk_path` parameter if they are missing.
        /// Default: `false`.
        /// 
        /// &gt; **NOTE:** Any directory created as part of the operation when
        /// `create_directories` is enabled will not be deleted when the resource is
        /// destroyed.
        /// </summary>
        [Output("createDirectories")]
        public Output<bool?> CreateDirectories { get; private set; } = null!;

        /// <summary>
        /// The name of the datacenter in which to create the
        /// disk. Can be omitted when ESXi or if there is only one datacenter in
        /// your infrastructure.
        /// </summary>
        [Output("datacenter")]
        public Output<string?> Datacenter { get; private set; } = null!;

        /// <summary>
        /// The name of the datastore in which to create the
        /// disk.
        /// </summary>
        [Output("datastore")]
        public Output<string> Datastore { get; private set; } = null!;

        /// <summary>
        /// Size of the disk (in GB). Decreasing the size of a disk is not possible.
        /// If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
        /// created.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// The type of disk to create. Can be one of
        /// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
        /// information on what each kind of disk provisioning policy means, click
        /// [here][docs-vmware-vm-disk-provisioning].
        /// 
        /// [docs-vmware-vm-disk-provisioning]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-single-host-management-vmware-host-client-8-0/virtual-machine-management-with-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/configuring-virtual-machines-in-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/virtual-disk-configuration-vSphereSingleHostManagementVMwareHostClient/about-virtual-disk-provisioning-policies-vSphereSingleHostManagementVMwareHostClient.html
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// The path, including filename, of the virtual disk to
        /// be created.  This needs to end in `.vmdk`.
        /// </summary>
        [Output("vmdkPath")]
        public Output<string> VmdkPath { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualDisk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualDisk(string name, VirtualDiskArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/virtualDisk:VirtualDisk", name, args ?? new VirtualDiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualDisk(string name, Input<string> id, VirtualDiskState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/virtualDisk:VirtualDisk", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualDisk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualDisk Get(string name, Input<string> id, VirtualDiskState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualDisk(name, id, state, options);
        }
    }

    public sealed class VirtualDiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The adapter type for this virtual disk. Can be
        /// one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
        /// 
        /// &gt; **NOTE:** `adapter_type` is **deprecated**: it does not dictate the type of
        /// controller that the virtual disk will be attached to on the virtual machine.
        /// Please see the `scsi_type` parameter
        /// in the `vsphere.VirtualMachine` resource for information on how to control
        /// disk controller types. This parameter will be removed in future versions of the
        /// vSphere provider.
        /// </summary>
        [Input("adapterType")]
        public Input<string>? AdapterType { get; set; }

        /// <summary>
        /// Tells the resource to create any
        /// directories that are a part of the `vmdk_path` parameter if they are missing.
        /// Default: `false`.
        /// 
        /// &gt; **NOTE:** Any directory created as part of the operation when
        /// `create_directories` is enabled will not be deleted when the resource is
        /// destroyed.
        /// </summary>
        [Input("createDirectories")]
        public Input<bool>? CreateDirectories { get; set; }

        /// <summary>
        /// The name of the datacenter in which to create the
        /// disk. Can be omitted when ESXi or if there is only one datacenter in
        /// your infrastructure.
        /// </summary>
        [Input("datacenter")]
        public Input<string>? Datacenter { get; set; }

        /// <summary>
        /// The name of the datastore in which to create the
        /// disk.
        /// </summary>
        [Input("datastore", required: true)]
        public Input<string> Datastore { get; set; } = null!;

        /// <summary>
        /// Size of the disk (in GB). Decreasing the size of a disk is not possible.
        /// If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
        /// created.
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        /// <summary>
        /// The type of disk to create. Can be one of
        /// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
        /// information on what each kind of disk provisioning policy means, click
        /// [here][docs-vmware-vm-disk-provisioning].
        /// 
        /// [docs-vmware-vm-disk-provisioning]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-single-host-management-vmware-host-client-8-0/virtual-machine-management-with-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/configuring-virtual-machines-in-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/virtual-disk-configuration-vSphereSingleHostManagementVMwareHostClient/about-virtual-disk-provisioning-policies-vSphereSingleHostManagementVMwareHostClient.html
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The path, including filename, of the virtual disk to
        /// be created.  This needs to end in `.vmdk`.
        /// </summary>
        [Input("vmdkPath", required: true)]
        public Input<string> VmdkPath { get; set; } = null!;

        public VirtualDiskArgs()
        {
        }
        public static new VirtualDiskArgs Empty => new VirtualDiskArgs();
    }

    public sealed class VirtualDiskState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The adapter type for this virtual disk. Can be
        /// one of `ide`, `lsiLogic`, or `busLogic`.  Default: `lsiLogic`.
        /// 
        /// &gt; **NOTE:** `adapter_type` is **deprecated**: it does not dictate the type of
        /// controller that the virtual disk will be attached to on the virtual machine.
        /// Please see the `scsi_type` parameter
        /// in the `vsphere.VirtualMachine` resource for information on how to control
        /// disk controller types. This parameter will be removed in future versions of the
        /// vSphere provider.
        /// </summary>
        [Input("adapterType")]
        public Input<string>? AdapterType { get; set; }

        /// <summary>
        /// Tells the resource to create any
        /// directories that are a part of the `vmdk_path` parameter if they are missing.
        /// Default: `false`.
        /// 
        /// &gt; **NOTE:** Any directory created as part of the operation when
        /// `create_directories` is enabled will not be deleted when the resource is
        /// destroyed.
        /// </summary>
        [Input("createDirectories")]
        public Input<bool>? CreateDirectories { get; set; }

        /// <summary>
        /// The name of the datacenter in which to create the
        /// disk. Can be omitted when ESXi or if there is only one datacenter in
        /// your infrastructure.
        /// </summary>
        [Input("datacenter")]
        public Input<string>? Datacenter { get; set; }

        /// <summary>
        /// The name of the datastore in which to create the
        /// disk.
        /// </summary>
        [Input("datastore")]
        public Input<string>? Datastore { get; set; }

        /// <summary>
        /// Size of the disk (in GB). Decreasing the size of a disk is not possible.
        /// If a disk of a smaller size is required then the original has to be destroyed along with its data and a new one has to be
        /// created.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The type of disk to create. Can be one of
        /// `eagerZeroedThick`, `lazy`, or `thin`. Default: `eagerZeroedThick`. For
        /// information on what each kind of disk provisioning policy means, click
        /// [here][docs-vmware-vm-disk-provisioning].
        /// 
        /// [docs-vmware-vm-disk-provisioning]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-single-host-management-vmware-host-client-8-0/virtual-machine-management-with-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/configuring-virtual-machines-in-the-vsphere-host-client-vSphereSingleHostManagementVMwareHostClient/virtual-disk-configuration-vSphereSingleHostManagementVMwareHostClient/about-virtual-disk-provisioning-policies-vSphereSingleHostManagementVMwareHostClient.html
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The path, including filename, of the virtual disk to
        /// be created.  This needs to end in `.vmdk`.
        /// </summary>
        [Input("vmdkPath")]
        public Input<string>? VmdkPath { get; set; }

        public VirtualDiskState()
        {
        }
        public static new VirtualDiskState Empty => new VirtualDiskState();
    }
}

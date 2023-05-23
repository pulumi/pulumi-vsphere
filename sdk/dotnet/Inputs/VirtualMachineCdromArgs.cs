// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class VirtualMachineCdromArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the device should be backed by remote client device. Conflicts with `datastore_id` and `path`.
        /// </summary>
        [Input("clientDevice")]
        public Input<bool>? ClientDevice { get; set; }

        /// <summary>
        /// The managed object reference ID of the datastore in which to place the virtual machine. The virtual machine configuration files is placed here, along with any virtual disks that are created where a datastore is not explicitly specified. See the section on virtual machine migration for more information on modifying this value.
        /// 
        /// &gt; **NOTE:** Datastores cannot be assigned to individual disks when `datastore_cluster_id` is used.
        /// </summary>
        [Input("datastoreId")]
        public Input<string>? DatastoreId { get; set; }

        [Input("deviceAddress")]
        public Input<string>? DeviceAddress { get; set; }

        /// <summary>
        /// The ID of the device within the virtual machine.
        /// </summary>
        [Input("key")]
        public Input<int>? Key { get; set; }

        /// <summary>
        /// When using `attach`, this parameter controls the path of a virtual disk to attach externally. Otherwise, it is a computed attribute that contains the virtual disk filename.
        /// 
        /// &gt; **NOTE:** Either `client_device` (for a remote backed CD-ROM) or `datastore_id` and `path` (for a datastore ISO backed CD-ROM) are required to .
        /// 
        /// &gt; **NOTE:** Some CD-ROM drive types are not supported by this resource, such as pass-through devices. If these drives are present in a cloned template, or added outside of the provider, the desired state will be corrected to the defined device, or removed if no `cdrom` block is present.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public VirtualMachineCdromArgs()
        {
        }
        public static new VirtualMachineCdromArgs Empty => new VirtualMachineCdromArgs();
    }
}

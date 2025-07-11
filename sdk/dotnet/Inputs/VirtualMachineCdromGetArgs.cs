// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class VirtualMachineCdromGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the device should be mapped to a remote client device
        /// </summary>
        [Input("clientDevice")]
        public Input<bool>? ClientDevice { get; set; }

        /// <summary>
        /// The datastore ID the ISO is located on.
        /// </summary>
        [Input("datastoreId")]
        public Input<string>? DatastoreId { get; set; }

        /// <summary>
        /// The internally-computed address of this device, such as scsi:0:1, denoting scsi bus #0 and device unit 1.
        /// </summary>
        [Input("deviceAddress")]
        public Input<string>? DeviceAddress { get; set; }

        /// <summary>
        /// The ID of the device within the virtual machine.
        /// </summary>
        [Input("key")]
        public Input<int>? Key { get; set; }

        /// <summary>
        /// The path to the ISO file on the datastore.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public VirtualMachineCdromGetArgs()
        {
        }
        public static new VirtualMachineCdromGetArgs Empty => new VirtualMachineCdromGetArgs();
    }
}

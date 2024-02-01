// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class VirtualMachineCloneCustomizationSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UUID of the virtual machine.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The amount of time, in minutes, to wait for guest OS customization to complete before returning with an error. Setting this value to 0 or a negative value skips the waiter. Default: 10.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public VirtualMachineCloneCustomizationSpecArgs()
        {
        }
        public static new VirtualMachineCloneCustomizationSpecArgs Empty => new VirtualMachineCloneCustomizationSpecArgs();
    }
}

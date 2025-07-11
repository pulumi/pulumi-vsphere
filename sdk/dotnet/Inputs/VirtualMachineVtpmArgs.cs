// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class VirtualMachineVtpmArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The version of the TPM device. Default is 2.0.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public VirtualMachineVtpmArgs()
        {
        }
        public static new VirtualMachineVtpmArgs Empty => new VirtualMachineVtpmArgs();
    }
}

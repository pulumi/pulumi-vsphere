// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class ComputeClusterHostImageArgs : global::Pulumi.ResourceArgs
    {
        [Input("components")]
        private InputList<Inputs.ComputeClusterHostImageComponentArgs>? _components;

        /// <summary>
        /// List of custom components.
        /// </summary>
        public InputList<Inputs.ComputeClusterHostImageComponentArgs> Components
        {
            get => _components ?? (_components = new InputList<Inputs.ComputeClusterHostImageComponentArgs>());
            set => _components = value;
        }

        /// <summary>
        /// The ESXi version which the image is based on.
        /// </summary>
        [Input("esxVersion")]
        public Input<string>? EsxVersion { get; set; }

        public ComputeClusterHostImageArgs()
        {
        }
        public static new ComputeClusterHostImageArgs Empty => new ComputeClusterHostImageArgs();
    }
}

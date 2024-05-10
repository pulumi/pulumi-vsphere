// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class ComputeClusterHostImageComponentGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identifier for the component.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The version to use.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ComputeClusterHostImageComponentGetArgs()
        {
        }
        public static new ComputeClusterHostImageComponentGetArgs Empty => new ComputeClusterHostImageComponentGetArgs();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Outputs
{

    [OutputType]
    public sealed class SupervisorServiceCidr
    {
        /// <summary>
        /// Network address.
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// Subnet prefix.
        /// </summary>
        public readonly int Prefix;

        [OutputConstructor]
        private SupervisorServiceCidr(
            string address,

            int prefix)
        {
            Address = address;
            Prefix = prefix;
        }
    }
}

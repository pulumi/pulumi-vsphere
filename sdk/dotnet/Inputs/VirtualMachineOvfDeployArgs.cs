// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class VirtualMachineOvfDeployArgs : Pulumi.ResourceArgs
    {
        [Input("allowUnverifiedSslCert")]
        public Input<bool>? AllowUnverifiedSslCert { get; set; }

        [Input("diskProvisioning")]
        public Input<string>? DiskProvisioning { get; set; }

        [Input("ipAllocationPolicy")]
        public Input<string>? IpAllocationPolicy { get; set; }

        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        [Input("localOvfPath")]
        public Input<string>? LocalOvfPath { get; set; }

        [Input("ovfNetworkMap")]
        private InputMap<string>? _ovfNetworkMap;
        public InputMap<string> OvfNetworkMap
        {
            get => _ovfNetworkMap ?? (_ovfNetworkMap = new InputMap<string>());
            set => _ovfNetworkMap = value;
        }

        [Input("remoteOvfUrl")]
        public Input<string>? RemoteOvfUrl { get; set; }

        public VirtualMachineOvfDeployArgs()
        {
        }
    }
}

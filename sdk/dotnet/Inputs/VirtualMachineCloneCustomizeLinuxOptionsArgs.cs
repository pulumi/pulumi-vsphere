// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class VirtualMachineCloneCustomizeLinuxOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("hostName", required: true)]
        public Input<string> HostName { get; set; } = null!;

        [Input("hwClockUtc")]
        public Input<bool>? HwClockUtc { get; set; }

        [Input("scriptText")]
        private Input<string>? _scriptText;
        public Input<string>? ScriptText
        {
            get => _scriptText;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _scriptText = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public VirtualMachineCloneCustomizeLinuxOptionsArgs()
        {
        }
        public static new VirtualMachineCloneCustomizeLinuxOptionsArgs Empty => new VirtualMachineCloneCustomizeLinuxOptionsArgs();
    }
}

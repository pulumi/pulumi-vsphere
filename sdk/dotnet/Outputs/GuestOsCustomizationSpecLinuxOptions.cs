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
    public sealed class GuestOsCustomizationSpecLinuxOptions
    {
        public readonly string Domain;
        public readonly string HostName;
        public readonly bool? HwClockUtc;
        public readonly string? ScriptText;
        public readonly string? TimeZone;

        [OutputConstructor]
        private GuestOsCustomizationSpecLinuxOptions(
            string domain,

            string hostName,

            bool? hwClockUtc,

            string? scriptText,

            string? timeZone)
        {
            Domain = domain;
            HostName = hostName;
            HwClockUtc = hwClockUtc;
            ScriptText = scriptText;
            TimeZone = timeZone;
        }
    }
}

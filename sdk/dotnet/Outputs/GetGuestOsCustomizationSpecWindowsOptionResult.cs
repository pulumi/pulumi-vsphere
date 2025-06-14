// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Outputs
{

    [OutputType]
    public sealed class GetGuestOsCustomizationSpecWindowsOptionResult
    {
        /// <summary>
        /// The new administrator password for this virtual machine.
        /// </summary>
        public readonly string AdminPassword;
        /// <summary>
        /// Specifies whether or not the guest operating system automatically logs on as Administrator.
        /// </summary>
        public readonly bool AutoLogon;
        /// <summary>
        /// Specifies how many times the guest operating system should auto-logon the Administrator account when `auto_logon` is `true`.
        /// </summary>
        public readonly int AutoLogonCount;
        /// <summary>
        /// The hostname for this virtual machine.
        /// </summary>
        public readonly string ComputerName;
        /// <summary>
        /// The user account used to join this virtual machine to the Active Directory domain.
        /// </summary>
        public readonly string? DomainAdminPassword;
        /// <summary>
        /// The user account of the domain administrator used to join this virtual machine to the domain.
        /// </summary>
        public readonly string DomainAdminUser;
        /// <summary>
        /// The MachineObjectOU which specifies the full LDAP path name of the OU to which the virtual machine belongs.
        /// </summary>
        public readonly string DomainOu;
        /// <summary>
        /// The full name of the user of this virtual machine.
        /// </summary>
        public readonly string FullName;
        /// <summary>
        /// The Active Directory domain for the virtual machine to join.
        /// </summary>
        public readonly string JoinDomain;
        /// <summary>
        /// The organization name this virtual machine is being installed for.
        /// </summary>
        public readonly string OrganizationName;
        /// <summary>
        /// The product key for this virtual machine.
        /// </summary>
        public readonly string ProductKey;
        /// <summary>
        /// A list of commands to run at first user logon, after guest customization.
        /// </summary>
        public readonly ImmutableArray<string> RunOnceCommandLists;
        /// <summary>
        /// The new time zone for the virtual machine. This is a sysprep-dictated timezone code.
        /// </summary>
        public readonly int TimeZone;
        /// <summary>
        /// The workgroup for this virtual machine if not joining an Active Directory domain.
        /// </summary>
        public readonly string Workgroup;

        [OutputConstructor]
        private GetGuestOsCustomizationSpecWindowsOptionResult(
            string adminPassword,

            bool autoLogon,

            int autoLogonCount,

            string computerName,

            string? domainAdminPassword,

            string domainAdminUser,

            string domainOu,

            string fullName,

            string joinDomain,

            string organizationName,

            string productKey,

            ImmutableArray<string> runOnceCommandLists,

            int timeZone,

            string workgroup)
        {
            AdminPassword = adminPassword;
            AutoLogon = autoLogon;
            AutoLogonCount = autoLogonCount;
            ComputerName = computerName;
            DomainAdminPassword = domainAdminPassword;
            DomainAdminUser = domainAdminUser;
            DomainOu = domainOu;
            FullName = fullName;
            JoinDomain = joinDomain;
            OrganizationName = organizationName;
            ProductKey = productKey;
            RunOnceCommandLists = runOnceCommandLists;
            TimeZone = timeZone;
            Workgroup = workgroup;
        }
    }
}

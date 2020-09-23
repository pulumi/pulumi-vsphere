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
    public sealed class ContentLibrarySubscription
    {
        /// <summary>
        /// Method to log into remote Content Library. Must be `NONE` or `BASIC`.
        /// </summary>
        public readonly string? AuthenticationMethod;
        /// <summary>
        /// Enable automatic synchronization with the external content library.
        /// </summary>
        public readonly bool? AutomaticSync;
        /// <summary>
        /// Download all library content immediately.
        /// </summary>
        public readonly bool? OnDemand;
        /// <summary>
        /// Password to log in with.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// URL of remote Content Library.
        /// </summary>
        public readonly string? SubscriptionUrl;
        /// <summary>
        /// User name to log in with.
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private ContentLibrarySubscription(
            string? authenticationMethod,

            bool? automaticSync,

            bool? onDemand,

            string? password,

            string? subscriptionUrl,

            string? username)
        {
            AuthenticationMethod = authenticationMethod;
            AutomaticSync = automaticSync;
            OnDemand = onDemand;
            Password = password;
            SubscriptionUrl = subscriptionUrl;
            Username = username;
        }
    }
}
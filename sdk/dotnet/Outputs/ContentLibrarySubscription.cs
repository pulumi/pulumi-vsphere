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
        /// Authentication method to connect ro a published content library. Must be `NONE` or `BASIC`.
        /// </summary>
        public readonly string? AuthenticationMethod;
        /// <summary>
        /// Enable automatic synchronization with the published library. Default `false`.
        /// </summary>
        public readonly bool? AutomaticSync;
        /// <summary>
        /// Download the library from a content only when needed. Default `true`.
        /// </summary>
        public readonly bool? OnDemand;
        /// <summary>
        /// Password used for authentication.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// URL of the published content library.
        /// </summary>
        public readonly string? SubscriptionUrl;
        /// <summary>
        /// Username used for authentication.
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

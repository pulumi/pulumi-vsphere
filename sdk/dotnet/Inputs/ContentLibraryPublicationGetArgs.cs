// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class ContentLibraryPublicationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication method to connect ro a published content library. Must be `NONE` or `BASIC`.
        /// </summary>
        [Input("authenticationMethod")]
        public Input<string>? AuthenticationMethod { get; set; }

        /// <summary>
        /// Password used for authentication.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The URL of the published content library.
        /// </summary>
        [Input("publishUrl")]
        public Input<string>? PublishUrl { get; set; }

        /// <summary>
        /// Publish the content library. Default `false`.
        /// </summary>
        [Input("published")]
        public Input<bool>? Published { get; set; }

        /// <summary>
        /// Username used for authentication.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ContentLibraryPublicationGetArgs()
        {
        }
    }
}

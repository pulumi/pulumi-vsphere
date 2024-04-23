// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class ContentLibraryPublicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Method to authenticate users. Must be `NONE` or `BASIC`.
        /// </summary>
        [Input("authenticationMethod")]
        public Input<string>? AuthenticationMethod { get; set; }

        /// <summary>
        /// Password used by subscribers to authenticate.
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
        /// Username used by subscribers to authenticate. Currently can only be `vcsp`.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ContentLibraryPublicationArgs()
        {
        }
        public static new ContentLibraryPublicationArgs Empty => new ContentLibraryPublicationArgs();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    [VSphereResourceType("vsphere:index/contentLibrary:ContentLibrary")]
    public partial class ContentLibrary : Pulumi.CustomResource
    {
        /// <summary>
        /// A description for the content library.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the content library.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Options to publish a local content library.
        /// </summary>
        [Output("publication")]
        public Output<Outputs.ContentLibraryPublication> Publication { get; private set; } = null!;

        /// <summary>
        /// The managed object reference ID of the datastore on which to store the content library items.
        /// </summary>
        [Output("storageBackings")]
        public Output<ImmutableArray<string>> StorageBackings { get; private set; } = null!;

        /// <summary>
        /// Options subscribe to a published content library.
        /// </summary>
        [Output("subscription")]
        public Output<Outputs.ContentLibrarySubscription?> Subscription { get; private set; } = null!;


        /// <summary>
        /// Create a ContentLibrary resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContentLibrary(string name, ContentLibraryArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/contentLibrary:ContentLibrary", name, args ?? new ContentLibraryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContentLibrary(string name, Input<string> id, ContentLibraryState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/contentLibrary:ContentLibrary", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ContentLibrary resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContentLibrary Get(string name, Input<string> id, ContentLibraryState? state = null, CustomResourceOptions? options = null)
        {
            return new ContentLibrary(name, id, state, options);
        }
    }

    public sealed class ContentLibraryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the content library.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the content library.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Options to publish a local content library.
        /// </summary>
        [Input("publication")]
        public Input<Inputs.ContentLibraryPublicationArgs>? Publication { get; set; }

        [Input("storageBackings", required: true)]
        private InputList<string>? _storageBackings;

        /// <summary>
        /// The managed object reference ID of the datastore on which to store the content library items.
        /// </summary>
        public InputList<string> StorageBackings
        {
            get => _storageBackings ?? (_storageBackings = new InputList<string>());
            set => _storageBackings = value;
        }

        /// <summary>
        /// Options subscribe to a published content library.
        /// </summary>
        [Input("subscription")]
        public Input<Inputs.ContentLibrarySubscriptionArgs>? Subscription { get; set; }

        public ContentLibraryArgs()
        {
        }
    }

    public sealed class ContentLibraryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the content library.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the content library.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Options to publish a local content library.
        /// </summary>
        [Input("publication")]
        public Input<Inputs.ContentLibraryPublicationGetArgs>? Publication { get; set; }

        [Input("storageBackings")]
        private InputList<string>? _storageBackings;

        /// <summary>
        /// The managed object reference ID of the datastore on which to store the content library items.
        /// </summary>
        public InputList<string> StorageBackings
        {
            get => _storageBackings ?? (_storageBackings = new InputList<string>());
            set => _storageBackings = value;
        }

        /// <summary>
        /// Options subscribe to a published content library.
        /// </summary>
        [Input("subscription")]
        public Input<Inputs.ContentLibrarySubscriptionGetArgs>? Subscription { get; set; }

        public ContentLibraryState()
        {
        }
    }
}

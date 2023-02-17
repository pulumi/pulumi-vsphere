// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    [VSphereResourceType("vsphere:index/contentLibraryItem:ContentLibraryItem")]
    public partial class ContentLibraryItem : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A description for the content library item.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// File to import as the content library item.
        /// </summary>
        [Output("fileUrl")]
        public Output<string?> FileUrl { get; private set; } = null!;

        /// <summary>
        /// The ID of the content library in which to create the item.
        /// </summary>
        [Output("libraryId")]
        public Output<string> LibraryId { get; private set; } = null!;

        /// <summary>
        /// The name of the item to be created in the content library.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Virtual machine UUID to clone to content library.
        /// </summary>
        [Output("sourceUuid")]
        public Output<string?> SourceUuid { get; private set; } = null!;

        /// <summary>
        /// Type of content library item.
        /// One of "ovf", "iso", or "vm-template". Default: `ovf`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ContentLibraryItem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContentLibraryItem(string name, ContentLibraryItemArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/contentLibraryItem:ContentLibraryItem", name, args ?? new ContentLibraryItemArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContentLibraryItem(string name, Input<string> id, ContentLibraryItemState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/contentLibraryItem:ContentLibraryItem", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ContentLibraryItem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContentLibraryItem Get(string name, Input<string> id, ContentLibraryItemState? state = null, CustomResourceOptions? options = null)
        {
            return new ContentLibraryItem(name, id, state, options);
        }
    }

    public sealed class ContentLibraryItemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the content library item.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// File to import as the content library item.
        /// </summary>
        [Input("fileUrl")]
        public Input<string>? FileUrl { get; set; }

        /// <summary>
        /// The ID of the content library in which to create the item.
        /// </summary>
        [Input("libraryId", required: true)]
        public Input<string> LibraryId { get; set; } = null!;

        /// <summary>
        /// The name of the item to be created in the content library.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Virtual machine UUID to clone to content library.
        /// </summary>
        [Input("sourceUuid")]
        public Input<string>? SourceUuid { get; set; }

        /// <summary>
        /// Type of content library item.
        /// One of "ovf", "iso", or "vm-template". Default: `ovf`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ContentLibraryItemArgs()
        {
        }
        public static new ContentLibraryItemArgs Empty => new ContentLibraryItemArgs();
    }

    public sealed class ContentLibraryItemState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the content library item.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// File to import as the content library item.
        /// </summary>
        [Input("fileUrl")]
        public Input<string>? FileUrl { get; set; }

        /// <summary>
        /// The ID of the content library in which to create the item.
        /// </summary>
        [Input("libraryId")]
        public Input<string>? LibraryId { get; set; }

        /// <summary>
        /// The name of the item to be created in the content library.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Virtual machine UUID to clone to content library.
        /// </summary>
        [Input("sourceUuid")]
        public Input<string>? SourceUuid { get; set; }

        /// <summary>
        /// Type of content library item.
        /// One of "ovf", "iso", or "vm-template". Default: `ovf`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ContentLibraryItemState()
        {
        }
        public static new ContentLibraryItemState Empty => new ContentLibraryItemState();
    }
}

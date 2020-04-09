// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    /// <summary>
    /// The `vsphere..ContentLibraryItem` resource can be used to create items in a Content Library. Each item can contain 
    /// multiple files. Each `file_url` must be accessible from the vSphere environment as it will be downloaded from the
    /// specified location and stored on the Content Library's storage backing.
    /// 
    /// To make a `content_library_item` a functioning template, the template must be in OVF format. The .ovf and .vmdk
    /// file(s) can then be set as the `file_url` list.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/content_library_item.html.markdown.
    /// </summary>
    public partial class ContentLibraryItem : Pulumi.CustomResource
    {
        /// <summary>
        /// A description for the item.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A list of files to download for the Content Library item.
        /// </summary>
        [Output("fileUrls")]
        public Output<ImmutableArray<string>> FileUrls { get; private set; } = null!;

        /// <summary>
        /// The ID of the Content Library the item should be created in.
        /// </summary>
        [Output("libraryId")]
        public Output<string> LibraryId { get; private set; } = null!;

        /// <summary>
        /// The name of the item to be created in the Content Library.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Type of content library item.
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
            : base("vsphere:index/contentLibraryItem:ContentLibraryItem", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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

    public sealed class ContentLibraryItemArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the item.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("fileUrls", required: true)]
        private InputList<string>? _fileUrls;

        /// <summary>
        /// A list of files to download for the Content Library item.
        /// </summary>
        public InputList<string> FileUrls
        {
            get => _fileUrls ?? (_fileUrls = new InputList<string>());
            set => _fileUrls = value;
        }

        /// <summary>
        /// The ID of the Content Library the item should be created in.
        /// </summary>
        [Input("libraryId", required: true)]
        public Input<string> LibraryId { get; set; } = null!;

        /// <summary>
        /// The name of the item to be created in the Content Library.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Type of content library item.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ContentLibraryItemArgs()
        {
        }
    }

    public sealed class ContentLibraryItemState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the item.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("fileUrls")]
        private InputList<string>? _fileUrls;

        /// <summary>
        /// A list of files to download for the Content Library item.
        /// </summary>
        public InputList<string> FileUrls
        {
            get => _fileUrls ?? (_fileUrls = new InputList<string>());
            set => _fileUrls = value;
        }

        /// <summary>
        /// The ID of the Content Library the item should be created in.
        /// </summary>
        [Input("libraryId")]
        public Input<string>? LibraryId { get; set; }

        /// <summary>
        /// The name of the item to be created in the Content Library.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Type of content library item.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ContentLibraryItemState()
        {
        }
    }
}
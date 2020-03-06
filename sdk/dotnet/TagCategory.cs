// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere
{
    /// <summary>
    /// The `vsphere..TagCategory` resource can be used to create and manage tag
    /// categories, which determine how tags are grouped together and applied to
    /// specific objects.
    /// 
    /// For more information about tags, click [here][ext-tags-general]. For more
    /// information about tag categories specifically, click
    /// [here][ext-tag-categories].
    /// 
    /// [ext-tags-general]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vcenterhost.doc/GUID-E8E854DD-AA97-4E0C-8419-CE84F93C4058.html
    /// [ext-tag-categories]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vcenterhost.doc/GUID-BA3D1794-28F2-43F3-BCE9-3964CB207FB6.html
    /// 
    /// &gt; **NOTE:** Tagging support is unsupported on direct ESXi connections and
    /// requires vCenter 6.0 or higher.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/tag_category.html.markdown.
    /// </summary>
    public partial class TagCategory : Pulumi.CustomResource
    {
        /// <summary>
        /// A list object types that this category is
        /// valid to be assigned to. For a full list, click
        /// here.
        /// </summary>
        [Output("associableTypes")]
        public Output<ImmutableArray<string>> AssociableTypes { get; private set; } = null!;

        /// <summary>
        /// The number of tags that can be assigned from this
        /// category to a single object at once. Can be one of `SINGLE` (object can only
        /// be assigned one tag in this category), to `MULTIPLE` (object can be assigned
        /// multiple tags in this category). Forces a new resource if changed.
        /// </summary>
        [Output("cardinality")]
        public Output<string> Cardinality { get; private set; } = null!;

        /// <summary>
        /// A description for the category.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the category.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a TagCategory resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TagCategory(string name, TagCategoryArgs args, CustomResourceOptions? options = null)
            : base("vsphere:index/tagCategory:TagCategory", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private TagCategory(string name, Input<string> id, TagCategoryState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/tagCategory:TagCategory", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TagCategory resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TagCategory Get(string name, Input<string> id, TagCategoryState? state = null, CustomResourceOptions? options = null)
        {
            return new TagCategory(name, id, state, options);
        }
    }

    public sealed class TagCategoryArgs : Pulumi.ResourceArgs
    {
        [Input("associableTypes", required: true)]
        private InputList<string>? _associableTypes;

        /// <summary>
        /// A list object types that this category is
        /// valid to be assigned to. For a full list, click
        /// here.
        /// </summary>
        public InputList<string> AssociableTypes
        {
            get => _associableTypes ?? (_associableTypes = new InputList<string>());
            set => _associableTypes = value;
        }

        /// <summary>
        /// The number of tags that can be assigned from this
        /// category to a single object at once. Can be one of `SINGLE` (object can only
        /// be assigned one tag in this category), to `MULTIPLE` (object can be assigned
        /// multiple tags in this category). Forces a new resource if changed.
        /// </summary>
        [Input("cardinality", required: true)]
        public Input<string> Cardinality { get; set; } = null!;

        /// <summary>
        /// A description for the category.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the category.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public TagCategoryArgs()
        {
        }
    }

    public sealed class TagCategoryState : Pulumi.ResourceArgs
    {
        [Input("associableTypes")]
        private InputList<string>? _associableTypes;

        /// <summary>
        /// A list object types that this category is
        /// valid to be assigned to. For a full list, click
        /// here.
        /// </summary>
        public InputList<string> AssociableTypes
        {
            get => _associableTypes ?? (_associableTypes = new InputList<string>());
            set => _associableTypes = value;
        }

        /// <summary>
        /// The number of tags that can be assigned from this
        /// category to a single object at once. Can be one of `SINGLE` (object can only
        /// be assigned one tag in this category), to `MULTIPLE` (object can be assigned
        /// multiple tags in this category). Forces a new resource if changed.
        /// </summary>
        [Input("cardinality")]
        public Input<string>? Cardinality { get; set; }

        /// <summary>
        /// A description for the category.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the category.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public TagCategoryState()
        {
        }
    }
}

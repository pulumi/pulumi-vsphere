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
    /// Provides a VMware vSphere datacenter resource. This can be used as the primary
    /// container of inventory objects such as hosts and virtual machines.
    /// 
    /// ## Example Usage
    /// 
    /// ### Create datacenter on the root folder
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using VSphere = Pulumi.VSphere;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var prodDatacenter = new VSphere.Datacenter("prod_datacenter", new()
    ///     {
    ///         Name = "my_prod_datacenter",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Create datacenter on a subfolder
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using VSphere = Pulumi.VSphere;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var researchDatacenter = new VSphere.Datacenter("research_datacenter", new()
    ///     {
    ///         Name = "my_research_datacenter",
    ///         Folder = "/research/",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Importing
    /// 
    /// An existing datacenter can be [imported][docs-import] into this resource
    /// via supplying the full path to the datacenter. An example is below:
    /// 
    /// [docs-import]: /docs/import/index.html
    /// 
    /// The above would import the datacenter named `dc1`.
    /// </summary>
    [VSphereResourceType("vsphere:index/datacenter:Datacenter")]
    public partial class Datacenter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Map of custom attribute ids to value
        /// strings to set for datacenter resource. See
        /// [here][docs-setting-custom-attributes] for a reference on how to set values
        /// for custom attributes.
        /// 
        /// [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource
        /// 
        /// &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
        /// and require vCenter.
        /// </summary>
        [Output("customAttributes")]
        public Output<ImmutableDictionary<string, string>?> CustomAttributes { get; private set; } = null!;

        /// <summary>
        /// The folder where the datacenter should be created.
        /// Forces a new resource if changed.
        /// </summary>
        [Output("folder")]
        public Output<string?> Folder { get; private set; } = null!;

        /// <summary>
        /// Managed object ID of this datacenter.
        /// </summary>
        [Output("moid")]
        public Output<string> Moid { get; private set; } = null!;

        /// <summary>
        /// The name of the datacenter. This name needs to be unique
        /// within the folder. Forces a new resource if changed.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The IDs of any tags to attach to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Datacenter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Datacenter(string name, DatacenterArgs? args = null, CustomResourceOptions? options = null)
            : base("vsphere:index/datacenter:Datacenter", name, args ?? new DatacenterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Datacenter(string name, Input<string> id, DatacenterState? state = null, CustomResourceOptions? options = null)
            : base("vsphere:index/datacenter:Datacenter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Datacenter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Datacenter Get(string name, Input<string> id, DatacenterState? state = null, CustomResourceOptions? options = null)
        {
            return new Datacenter(name, id, state, options);
        }
    }

    public sealed class DatacenterArgs : global::Pulumi.ResourceArgs
    {
        [Input("customAttributes")]
        private InputMap<string>? _customAttributes;

        /// <summary>
        /// Map of custom attribute ids to value
        /// strings to set for datacenter resource. See
        /// [here][docs-setting-custom-attributes] for a reference on how to set values
        /// for custom attributes.
        /// 
        /// [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource
        /// 
        /// &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
        /// and require vCenter.
        /// </summary>
        public InputMap<string> CustomAttributes
        {
            get => _customAttributes ?? (_customAttributes = new InputMap<string>());
            set => _customAttributes = value;
        }

        /// <summary>
        /// The folder where the datacenter should be created.
        /// Forces a new resource if changed.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// The name of the datacenter. This name needs to be unique
        /// within the folder. Forces a new resource if changed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The IDs of any tags to attach to this resource.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public DatacenterArgs()
        {
        }
        public static new DatacenterArgs Empty => new DatacenterArgs();
    }

    public sealed class DatacenterState : global::Pulumi.ResourceArgs
    {
        [Input("customAttributes")]
        private InputMap<string>? _customAttributes;

        /// <summary>
        /// Map of custom attribute ids to value
        /// strings to set for datacenter resource. See
        /// [here][docs-setting-custom-attributes] for a reference on how to set values
        /// for custom attributes.
        /// 
        /// [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource
        /// 
        /// &gt; **NOTE:** Custom attributes are unsupported on direct ESXi connections
        /// and require vCenter.
        /// </summary>
        public InputMap<string> CustomAttributes
        {
            get => _customAttributes ?? (_customAttributes = new InputMap<string>());
            set => _customAttributes = value;
        }

        /// <summary>
        /// The folder where the datacenter should be created.
        /// Forces a new resource if changed.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// Managed object ID of this datacenter.
        /// </summary>
        [Input("moid")]
        public Input<string>? Moid { get; set; }

        /// <summary>
        /// The name of the datacenter. This name needs to be unique
        /// within the folder. Forces a new resource if changed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The IDs of any tags to attach to this resource.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public DatacenterState()
        {
        }
        public static new DatacenterState Empty => new DatacenterState();
    }
}
